package docspring

// Checks a request against API v2's request schemas (request_schemas.json, the schemas the API validates requests
// with) before it's sent, and says what the API would say about it. api_clients/validator_conformance checks that it
// agrees with the API.

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

//go:embed request_schemas.json
var requestSchemasJSON []byte

// ValidationErrorDetail is one of the API's error details: Pointer points into the body, Parameter names a query
// parameter.
type ValidationErrorDetail struct {
	Code      string  `json:"code"`
	Message   string  `json:"message"`
	Pointer   *string `json:"pointer,omitempty"`
	Parameter *string `json:"parameter,omitempty"`
}

// ValidationError is returned before a request is sent when the API would reject it with a 422. Details are the API's
// error details.
type ValidationError struct {
	Details []ValidationErrorDetail
}

func (e *ValidationError) Error() string {
	shown := []string{}
	for index, detail := range e.Details {
		if index == 3 {
			break
		}
		subject := ""
		if detail.Pointer != nil {
			subject = *detail.Pointer
			if subject == "" {
				subject = "the request body"
			}
		} else if detail.Parameter != nil {
			subject = *detail.Parameter
		}
		shown = append(shown, subject+" "+detail.Message)
	}
	message := "Invalid request: " + strings.Join(shown, "; ")
	if len(e.Details) > 3 {
		message += fmt.Sprintf(" (and %d more)", len(e.Details)-3)
	}
	return message
}

var validatorTypeNames = map[string]string{
	"string":  "a string",
	"integer": "an integer",
	"number":  "a number",
	"boolean": "true or false",
	"object":  "an object",
	"array":   "an array",
	"null":    "null",
}

// What a multipart body's file field holds when the API checks it
const validatorFilePlaceholder = "(file)"

// A date and time as the API reads it (Ruby's DateTime.rfc3339, then json_schemer's checks)
var validatorDateTime = regexp.MustCompile(`\A[ \t\n\v\f\r]*(-?\d{4})-(\d{2})-(\d{2})[Tt \t\n\v\f\r](\d{2}):(\d{2}):(\d{2})(?:\.\d+)?(?:[Zz]|([+-])(\d{2}):(\d{2}))\z`)

// An absolute URI as the API reads it (RFC 3986, as Ruby's URI parser has it), a little looser for IP addresses
var validatorURI = regexp.MustCompile(
	`\A[A-Za-z][+\-.0-9A-Za-z]*:` +
		`(?://(?:(?:%[0-9A-Fa-f]{2}|[!$&-.0-9:;=A-Z_a-z~])*@)?` +
		`(?:\[(?:[0-9A-Fa-f:.]+|[vV][0-9A-Fa-f]+\.[!$&-.0-9:;=A-Z_a-z~]+)\]|(?:%[0-9A-Fa-f]{2}|[!$&-.0-9;=A-Z_a-z~])*)` +
		`(?::[0-9]*)?(?:/(?:%[0-9A-Fa-f]{2}|[!$&-.0-9:;=@A-Z_a-z~/])*)?` +
		`|/(?:(?:%[0-9A-Fa-f]{2}|[!$&-.0-9:;=@A-Z_a-z~])(?:%[0-9A-Fa-f]{2}|[!$&-.0-9:;=@A-Z_a-z~/])*)?` +
		`|(?:%[0-9A-Fa-f]{2}|[!$&-.0-9:;=@A-Z_a-z~])(?:%[0-9A-Fa-f]{2}|[!$&-.0-9:;=@A-Z_a-z~/])*` +
		`|)` +
		`(?:\?[^#]*)?` +
		`(?:#(?:%[0-9A-Fa-f]{2}|[!$&-.0-9:;=@A-Z_a-z~/?])*)?\z`)
var validatorWhitespace = regexp.MustCompile(`[ \t\n\v\f\r]`)

var (
	validatorSchemas     map[string]interface{}
	validatorSchemasOnce sync.Once
	validatorPatterns    sync.Map
)

func requestSchemas() map[string]interface{} {
	validatorSchemasOnce.Do(func() {
		decoder := json.NewDecoder(bytes.NewReader(requestSchemasJSON))
		decoder.UseNumber()
		if err := decoder.Decode(&validatorSchemas); err != nil {
			panic(fmt.Sprintf("docspring: can't read request_schemas.json: %v", err))
		}
	})
	return validatorSchemas
}

// validateRequest returns a *ValidationError when the API would reject a request to an operation (createSubmission),
// unless the configuration's SkipRequestValidation is set
func (c *APIClient) validateRequest(operationID string, query url.Values, body interface{}, form url.Values, files []formFile) error {
	if c.cfg.SkipRequestValidation {
		return nil
	}
	queryValues := map[string]interface{}{}
	for name, values := range query {
		if len(values) > 0 {
			queryValues[name] = values[0]
		}
	}
	var requestBody interface{}
	if len(form) > 0 || len(files) > 0 {
		values := map[string]interface{}{}
		for name, items := range form {
			if len(items) > 0 {
				values[name] = items[0]
			}
		}
		for _, file := range files {
			values[file.formFileName] = validatorFilePlaceholder
		}
		requestBody = values
	} else if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return err
		}
		decoder := json.NewDecoder(bytes.NewReader(encoded))
		decoder.UseNumber()
		if err := decoder.Decode(&requestBody); err != nil {
			return err
		}
	}
	details := RequestErrors(operationID, queryValues, requestBody)
	if len(details) > 0 {
		return &ValidationError{Details: details}
	}
	return nil
}

// RequestErrors returns the details for a request to an operation (createSubmission). Numbers in the query and body
// are json.Number values (a json.Decoder with UseNumber), or Go numbers.
func RequestErrors(operationID string, query map[string]interface{}, body interface{}) []ValidationErrorDetail {
	operations, _ := requestSchemas()["operations"].(map[string]interface{})
	entry, ok := operations[operationID].(map[string]interface{})
	if !ok {
		return nil
	}
	details := []ValidationErrorDetail{}
	if params, ok := entry["query"].(map[string]interface{}); ok && len(query) > 0 {
		properties := map[string]interface{}{}
		required := []interface{}{}
		for name, param := range params {
			param := param.(map[string]interface{})
			properties[name] = param["schema"]
			if param["required"] == true {
				required = append(required, name)
			}
		}
		schema := map[string]interface{}{"type": "object", "properties": properties, "required": required, "additionalProperties": false}
		details = append(details, validateParams(schema, query, "query")...)
	}
	if bodyEntry, ok := entry["body"].(map[string]interface{}); ok && body != nil {
		schema := bodyEntry["schema"].(map[string]interface{})
		if bodyEntry["mediaType"] == "multipart/form-data" {
			if values, ok := body.(map[string]interface{}); ok {
				details = append(details, validateParams(schema, values, "body")...)
			}
		} else {
			details = append(details, ValidateValue(schema, body, "body")...)
		}
	}
	return details
}

// ValidateValue returns the details for data checked against a schema at a location (body or query)
func ValidateValue(schema interface{}, data interface{}, location string) []ValidationErrorDetail {
	details := []ValidationErrorDetail{}
	validatorCheck(schema, data, "", location, &details)
	return details
}

// Query and multipart params arrive as strings, which the API converts where the schema types them
func validateParams(schema map[string]interface{}, data map[string]interface{}, location string) []ValidationErrorDetail {
	properties, _ := schema["properties"].(map[string]interface{})
	converted := map[string]interface{}{}
	for name, value := range data {
		converted[name] = convertParam(value, properties[name])
	}
	return ValidateValue(schema, converted, location)
}

var (
	validatorIntegerString = regexp.MustCompile(`\A-?[0-9]+\z`)
	validatorNumberString  = regexp.MustCompile(`\A-?[0-9]+(?:\.[0-9]+)?\z`)
)

func convertParam(value interface{}, rawSchema interface{}) interface{} {
	schema, ok := rawSchema.(map[string]interface{})
	if !ok {
		return value
	}
	types := schemaTypes(schema)
	if items, ok := value.([]interface{}); ok && containsString(types, "array") {
		converted := make([]interface{}, len(items))
		for index, item := range items {
			converted[index] = convertParam(item, schema["items"])
		}
		return converted
	}
	text, ok := value.(string)
	if !ok {
		return value
	}
	if containsString(types, "integer") && validatorIntegerString.MatchString(text) {
		return json.Number(text)
	}
	if containsString(types, "number") && validatorNumberString.MatchString(text) {
		return json.Number(text)
	}
	if containsString(types, "boolean") && (text == "true" || text == "false") {
		return text == "true"
	}
	return value
}

func schemaTypes(schema map[string]interface{}) []string {
	switch types := schema["type"].(type) {
	case string:
		return []string{types}
	case []interface{}:
		names := []string{}
		for _, name := range types {
			if text, ok := name.(string); ok {
				names = append(names, text)
			}
		}
		return names
	default:
		return nil
	}
}

func containsString(values []string, value string) bool {
	for _, item := range values {
		if item == value {
			return true
		}
	}
	return false
}

func resolveSchema(schema interface{}) interface{} {
	for {
		object, ok := schema.(map[string]interface{})
		if !ok {
			return schema
		}
		ref, ok := object["$ref"].(string)
		if !ok {
			return schema
		}
		components := requestSchemas()["components"].(map[string]interface{})["schemas"].(map[string]interface{})
		schema = components[strings.TrimPrefix(ref, "#/components/schemas/")]
	}
}

func validatorCheck(rawSchema interface{}, data interface{}, pointer string, location string, details *[]ValidationErrorDetail) {
	schema, ok := resolveSchema(rawSchema).(map[string]interface{})
	if !ok {
		return
	}
	if discriminator, ok := schema["discriminator"].(map[string]interface{}); ok {
		checkUnion(discriminator, data, pointer, location, details)
		return
	}

	if _, ok := schema["type"]; ok {
		matches := false
		for _, name := range schemaTypes(schema) {
			matches = matches || isJSONType(data, name)
		}
		if !matches {
			message := typeMessage(schemaTypes(schema))
			if pointer == "" && location == "body" {
				message = "must be a JSON object"
			}
			*details = append(*details, validatorDetail(location, pointer, "type", message))
		}
	}
	if values, ok := schema["enum"].([]interface{}); ok {
		found := false
		shown := []string{}
		for _, value := range values {
			found = found || jsonEqual(value, data)
			shown = append(shown, toJSON(value))
		}
		if !found {
			*details = append(*details, validatorDetail(location, pointer, "enum", "must be one of: "+strings.Join(shown, ", ")))
		}
	}
	if value, ok := schema["const"]; ok && !jsonEqual(value, data) {
		*details = append(*details, validatorDetail(location, pointer, "const", "must be "+toJSON(value)))
	}

	switch value := data.(type) {
	case string:
		checkString(schema, value, pointer, location, details)
	case json.Number, float64, float32, int, int32, int64:
		if number, ok := toFloat(value); ok {
			checkNumber(schema, number, pointer, location, details)
		}
	case []interface{}:
		checkArray(schema, value, pointer, location, details)
	case map[string]interface{}:
		checkObject(schema, value, pointer, location, details)
	}
}

// A union (oneOf with a discriminator): the discriminator property's value picks the variant, and only that one is
// checked, as the API does. The mapping is [value, $ref] pairs, in the order the API lists them.
func checkUnion(discriminator map[string]interface{}, data interface{}, pointer string, location string, details *[]ValidationErrorDetail) {
	object, ok := data.(map[string]interface{})
	if !ok {
		*details = append(*details, validatorDetail(location, pointer, "type", "must be "+validatorTypeNames["object"]))
		return
	}

	name, _ := discriminator["propertyName"].(string)
	child := pointer + "/" + escapePointer(name)
	value, present := object[name]
	if !present {
		*details = append(*details, validatorDetail(location, child, "required", "is required"))
		return
	}

	mapping, _ := discriminator["mapping"].([]interface{})
	shown := []string{}
	for _, rawPair := range mapping {
		pair, _ := rawPair.([]interface{})
		if len(pair) != 2 {
			continue
		}
		if text, ok := value.(string); ok && pair[0] == text {
			validatorCheck(map[string]interface{}{"$ref": pair[1]}, data, pointer, location, details)
			return
		}
		shown = append(shown, toJSON(pair[0]))
	}
	*details = append(*details, validatorDetail(location, child, "enum", "must be one of: "+strings.Join(shown, ", ")))
}

func checkString(schema map[string]interface{}, data string, pointer string, location string, details *[]ValidationErrorDetail) {
	length := float64(utf8.RuneCountInString(data))
	if limit, ok := schemaNumber(schema, "minLength"); ok && length < limit {
		message := atLeast(limit, "must be at least "+countNoun(schema["minLength"], limit, "character"))
		*details = append(*details, validatorDetail(location, pointer, "minLength", message))
	}
	if limit, ok := schemaNumber(schema, "maxLength"); ok && length > limit {
		message := "must be at most " + countNoun(schema["maxLength"], limit, "character")
		*details = append(*details, validatorDetail(location, pointer, "maxLength", message))
	}
	if source, ok := schema["pattern"].(string); ok {
		if compiled := validatorPattern(source); compiled != nil && !compiled.MatchString(data) {
			message, ok := schema["x-pattern-message"].(string)
			if !ok {
				message = "must match the pattern " + source
			}
			*details = append(*details, validatorDetail(location, pointer, "pattern", message))
		}
	}
	if name, ok := schema["format"].(string); ok && !validFormat(name, data) {
		message, ok := requestSchemas()["formatMessages"].(map[string]interface{})[name].(string)
		if !ok {
			message = "must be a valid " + name
		}
		*details = append(*details, validatorDetail(location, pointer, "format", message))
	}
}

func checkNumber(schema map[string]interface{}, data float64, pointer string, location string, details *[]ValidationErrorDetail) {
	if limit, ok := schemaNumber(schema, "minimum"); ok && data < limit {
		*details = append(*details, validatorDetail(location, pointer, "minimum", "must be at least "+numberText(schema["minimum"])))
	}
	if limit, ok := schemaNumber(schema, "maximum"); ok && data > limit {
		*details = append(*details, validatorDetail(location, pointer, "maximum", "must be at most "+numberText(schema["maximum"])))
	}
	if limit, ok := schemaNumber(schema, "exclusiveMinimum"); ok && data <= limit {
		*details = append(*details, validatorDetail(location, pointer, "exclusiveMinimum", "must be greater than "+numberText(schema["exclusiveMinimum"])))
	}
	if limit, ok := schemaNumber(schema, "exclusiveMaximum"); ok && data >= limit {
		*details = append(*details, validatorDetail(location, pointer, "exclusiveMaximum", "must be less than "+numberText(schema["exclusiveMaximum"])))
	}
}

func checkArray(schema map[string]interface{}, data []interface{}, pointer string, location string, details *[]ValidationErrorDetail) {
	size := float64(len(data))
	if limit, ok := schemaNumber(schema, "minItems"); ok && size < limit {
		message := atLeast(limit, "must have at least "+countNoun(schema["minItems"], limit, "item"))
		*details = append(*details, validatorDetail(location, pointer, "minItems", message))
	}
	if limit, ok := schemaNumber(schema, "maxItems"); ok && size > limit {
		message := "must have at most " + countNoun(schema["maxItems"], limit, "item")
		*details = append(*details, validatorDetail(location, pointer, "maxItems", message))
	}
	if schema["uniqueItems"] == true && !uniqueItems(data) {
		message := "can't contain the same item twice"
		if duplicate, ok := firstDuplicate(data); ok {
			message = "has " + toJSON(duplicate) + " more than once"
		}
		*details = append(*details, validatorDetail(location, pointer, "uniqueItems", message))
	}
	if items, ok := schema["items"]; ok {
		for index, item := range data {
			validatorCheck(items, item, pointer+"/"+strconv.Itoa(index), location, details)
		}
	}
}

func checkObject(schema map[string]interface{}, data map[string]interface{}, pointer string, location string, details *[]ValidationErrorDetail) {
	if required, ok := schema["required"].([]interface{}); ok {
		for _, rawName := range required {
			name, _ := rawName.(string)
			if _, present := data[name]; !present {
				*details = append(*details, validatorDetail(location, pointer+"/"+escapePointer(name), "required", "is required"))
			}
		}
	}

	properties, _ := schema["properties"].(map[string]interface{})
	names := make([]string, 0, len(data))
	for name := range data {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		child := pointer + "/" + escapePointer(name)
		if property, ok := properties[name]; ok {
			validatorCheck(property, data[name], child, location, details)
		} else if schema["additionalProperties"] == false {
			*details = append(*details, unknownProperty(schema, name, child, location))
		} else if additional, ok := schema["additionalProperties"].(map[string]interface{}); ok {
			validatorCheck(additional, data[name], child, location, details)
		}
	}
}

func unknownProperty(schema map[string]interface{}, name string, pointer string, location string) ValidationErrorDetail {
	if notAllowed, ok := schema["x-not-allowed-properties"].(map[string]interface{}); ok {
		if message, ok := notAllowed[name].(string); ok {
			return validatorDetail(location, pointer, "not_allowed", message)
		}
	}

	noun := "property"
	if location == "query" {
		noun = "query parameter"
	}
	custom, hasCustom := schema["x-unknown-property-message"].(string)
	message := "isn't a documented " + noun
	if hasCustom {
		message = custom
	}
	snake := underscore(name)
	properties, _ := schema["properties"].(map[string]interface{})
	if _, ok := properties[snake]; ok && snake != name {
		plural := "properties"
		if hasCustom {
			plural = "settings"
		} else if noun == "query parameter" {
			plural = "query parameters"
		}
		message += ". API v2 names " + plural + " in snake_case: '" + snake + "'"
	}
	return validatorDetail(location, pointer, "unknown_property", message)
}

func validatorDetail(location string, pointer string, code string, message string) ValidationErrorDetail {
	if location != "query" {
		value := pointer
		return ValidationErrorDetail{Code: code, Message: message, Pointer: &value}
	}
	parameter := strings.SplitN(strings.TrimPrefix(pointer, "/"), "/", 2)[0]
	parameter = strings.ReplaceAll(strings.ReplaceAll(parameter, "~1", "/"), "~0", "~")
	return ValidationErrorDetail{Code: code, Message: message, Parameter: &parameter}
}

func toFloat(value interface{}) (float64, bool) {
	switch number := value.(type) {
	case json.Number:
		parsed, err := number.Float64()
		return parsed, err == nil
	case float64:
		return number, true
	case float32:
		return float64(number), true
	case int:
		return float64(number), true
	case int32:
		return float64(number), true
	case int64:
		return float64(number), true
	default:
		return 0, false
	}
}

func schemaNumber(schema map[string]interface{}, keyword string) (float64, bool) {
	value, ok := schema[keyword]
	if !ok {
		return 0, false
	}
	return toFloat(value)
}

func isJSONType(data interface{}, name string) bool {
	switch name {
	case "string":
		_, ok := data.(string)
		return ok
	case "integer":
		number, ok := toFloat(data)
		return ok && !math.IsInf(number, 0) && !math.IsNaN(number) && number == math.Floor(number)
	case "number":
		_, ok := toFloat(data)
		return ok
	case "boolean":
		_, ok := data.(bool)
		return ok
	case "object":
		_, ok := data.(map[string]interface{})
		return ok
	case "array":
		_, ok := data.([]interface{})
		return ok
	case "null":
		return data == nil
	default:
		return false
	}
}

// Patterns are ECMA-262 regexps, as OpenAPI says: ^ and $ are the start and end of the value, not of a line. The
// API's patterns only use ^ as the first character and $ as the last (api_clients/scripts/extract_request_schemas).
// One Go can't compile isn't checked.
func validatorPattern(source string) *regexp.Regexp {
	if compiled, ok := validatorPatterns.Load(source); ok {
		return compiled.(*regexp.Regexp)
	}
	anchored := source
	if strings.HasPrefix(anchored, "^") {
		anchored = `\A` + anchored[1:]
	}
	// A $ after an even number of backslashes is the anchor
	if strings.HasSuffix(anchored, "$") {
		trimmed := strings.TrimRight(anchored[:len(anchored)-1], `\`)
		if (len(anchored)-1-len(trimmed))%2 == 0 {
			anchored = anchored[:len(anchored)-1] + `\z`
		}
	}
	compiled, err := regexp.Compile(anchored)
	if err != nil {
		return nil
	}
	validatorPatterns.Store(source, compiled)
	return compiled
}

func validFormat(name string, data string) bool {
	switch name {
	case "uri":
		for _, character := range data {
			if character > 0x7f {
				return false
			}
		}
		return !validatorWhitespace.MatchString(data) && validatorURI.MatchString(data)
	case "date-time":
		return validDateTime(data)
	default:
		return true
	}
}

func validDateTime(data string) bool {
	match := validatorDateTime.FindStringSubmatch(data)
	if match == nil {
		return false
	}
	numbers := make([]int, 6)
	for index := range numbers {
		numbers[index], _ = strconv.Atoi(match[index+1])
	}
	year, month, day, hour, minute, second := numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5]
	if month < 1 || month > 12 || day < 1 || day > daysInMonth(year, month) {
		return false
	}
	if hour > 23 || minute > 59 || second > 60 {
		return false
	}
	offset := 0
	if match[7] != "" {
		offsetHours, _ := strconv.Atoi(match[8])
		offsetMinutes, _ := strconv.Atoi(match[9])
		if offsetHours > 23 || offsetMinutes > 59 {
			return false
		}
		offset = offsetHours*60 + offsetMinutes
		if match[7] == "-" {
			offset = -offset
		}
	}
	// A leap second is only at the end of a UTC day
	return second < 60 || ((hour*60+minute-offset)%1440+1440)%1440 == 1439
}

// Before the Gregorian calendar (1582), the API's dates are Julian: every fourth year is a leap year
func daysInMonth(year int, month int) int {
	leap := year%4 == 0 && (year < 1583 || year%100 != 0 || year%400 == 0)
	if month == 2 && leap {
		return 29
	}
	return []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}[month-1]
}

func uniqueItems(data []interface{}) bool {
	for index, item := range data {
		for _, other := range data[index+1:] {
			if jsonEqual(item, other) {
				return false
			}
		}
	}
	return true
}

// The first value that's in the list twice, as the API names it
func firstDuplicate(data []interface{}) (interface{}, bool) {
	keys := []string{}
	counts := map[string]int{}
	firsts := map[string]interface{}{}
	for _, item := range data {
		key := tallyKey(item)
		if _, seen := counts[key]; !seen {
			keys = append(keys, key)
			firsts[key] = item
		}
		counts[key]++
	}
	for _, key := range keys {
		if counts[key] > 1 {
			return firsts[key], true
		}
	}
	return nil, false
}

func tallyKey(value interface{}) string {
	switch item := value.(type) {
	case []interface{}:
		parts := make([]string, len(item))
		for index, element := range item {
			parts[index] = tallyKey(element)
		}
		return "[" + strings.Join(parts, ",") + "]"
	case map[string]interface{}:
		names := make([]string, 0, len(item))
		for name := range item {
			names = append(names, name)
		}
		sort.Strings(names)
		parts := make([]string, len(names))
		for index, name := range names {
			parts[index] = toJSON(name) + ":" + tallyKey(item[name])
		}
		return "{" + strings.Join(parts, ",") + "}"
	default:
		return toJSON(value)
	}
}

// JSON equality: numbers by value, objects regardless of key order
func jsonEqual(left interface{}, right interface{}) bool {
	leftNumber, leftIsNumber := toFloat(left)
	rightNumber, rightIsNumber := toFloat(right)
	if leftIsNumber || rightIsNumber {
		return leftIsNumber && rightIsNumber && leftNumber == rightNumber
	}
	switch leftValue := left.(type) {
	case []interface{}:
		rightValue, ok := right.([]interface{})
		if !ok || len(leftValue) != len(rightValue) {
			return false
		}
		for index := range leftValue {
			if !jsonEqual(leftValue[index], rightValue[index]) {
				return false
			}
		}
		return true
	case map[string]interface{}:
		rightValue, ok := right.(map[string]interface{})
		if !ok || len(leftValue) != len(rightValue) {
			return false
		}
		for name, value := range leftValue {
			other, present := rightValue[name]
			if !present || !jsonEqual(value, other) {
				return false
			}
		}
		return true
	default:
		return left == right
	}
}

func typeMessage(types []string) string {
	names := make([]string, len(types))
	connector := " or "
	for index, name := range types {
		names[index] = name
		if text, ok := validatorTypeNames[name]; ok {
			names[index] = text
		}
	}
	for _, name := range names {
		if len(names) > 1 && strings.Contains(name, " or ") {
			connector = ", or "
		}
	}
	switch len(names) {
	case 0:
		return "must be valid"
	case 1:
		return "must be " + names[0]
	case 2:
		return "must be " + strings.Join(names, connector)
	default:
		return "must be " + strings.Join(names[:len(names)-1], ", ") + connector + names[len(names)-1]
	}
}

func atLeast(limit float64, message string) string {
	if limit == 1 {
		return "can't be empty"
	}
	return message
}

func countNoun(raw interface{}, number float64, noun string) string {
	if number == 1 {
		return numberText(raw) + " " + noun
	}
	return numberText(raw) + " " + noun + "s"
}

// A number as the schema writes it
func numberText(value interface{}) string {
	if number, ok := value.(json.Number); ok {
		return number.String()
	}
	return fmt.Sprint(value)
}

func toJSON(value interface{}) string {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(value); err != nil {
		return fmt.Sprint(value)
	}
	return strings.TrimSuffix(buffer.String(), "\n")
}

func escapePointer(token string) string {
	return strings.ReplaceAll(strings.ReplaceAll(token, "~", "~0"), "/", "~1")
}

var (
	validatorAcronymBoundary = regexp.MustCompile(`([A-Z\d]+)([A-Z][a-z])`)
	validatorWordBoundary    = regexp.MustCompile(`([a-z\d])([A-Z])`)
)

func underscore(name string) string {
	name = validatorAcronymBoundary.ReplaceAllString(name, "${1}_${2}")
	name = validatorWordBoundary.ReplaceAllString(name, "${1}_${2}")
	return strings.ToLower(strings.ReplaceAll(name, "-", "_"))
}
