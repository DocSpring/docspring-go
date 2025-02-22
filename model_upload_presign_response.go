/*
DocSpring API

DocSpring provides an API that helps you fill out and sign PDF templates.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package docspring

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UploadPresignResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadPresignResponse{}

// UploadPresignResponse struct for UploadPresignResponse
type UploadPresignResponse struct {
	Fields map[string]interface{} `json:"fields"`
	Headers map[string]interface{} `json:"headers"`
	Url string `json:"url"`
	Method *string `json:"method,omitempty"`
}

type _UploadPresignResponse UploadPresignResponse

// NewUploadPresignResponse instantiates a new UploadPresignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadPresignResponse(fields map[string]interface{}, headers map[string]interface{}, url string) *UploadPresignResponse {
	this := UploadPresignResponse{}
	this.Fields = fields
	this.Headers = headers
	this.Url = url
	return &this
}

// NewUploadPresignResponseWithDefaults instantiates a new UploadPresignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadPresignResponseWithDefaults() *UploadPresignResponse {
	this := UploadPresignResponse{}
	return &this
}

// GetFields returns the Fields field value
func (o *UploadPresignResponse) GetFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *UploadPresignResponse) GetFieldsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *UploadPresignResponse) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetHeaders returns the Headers field value
func (o *UploadPresignResponse) GetHeaders() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *UploadPresignResponse) GetHeadersOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Headers, true
}

// SetHeaders sets field value
func (o *UploadPresignResponse) SetHeaders(v map[string]interface{}) {
	o.Headers = v
}

// GetUrl returns the Url field value
func (o *UploadPresignResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UploadPresignResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UploadPresignResponse) SetUrl(v string) {
	o.Url = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *UploadPresignResponse) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadPresignResponse) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *UploadPresignResponse) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *UploadPresignResponse) SetMethod(v string) {
	o.Method = &v
}

func (o UploadPresignResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadPresignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	toSerialize["headers"] = o.Headers
	toSerialize["url"] = o.Url
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	return toSerialize, nil
}

func (o *UploadPresignResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields",
		"headers",
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUploadPresignResponse := _UploadPresignResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUploadPresignResponse)

	if err != nil {
		return err
	}

	*o = UploadPresignResponse(varUploadPresignResponse)

	return err
}

type NullableUploadPresignResponse struct {
	value *UploadPresignResponse
	isSet bool
}

func (v NullableUploadPresignResponse) Get() *UploadPresignResponse {
	return v.value
}

func (v *NullableUploadPresignResponse) Set(val *UploadPresignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadPresignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadPresignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadPresignResponse(val *UploadPresignResponse) *NullableUploadPresignResponse {
	return &NullableUploadPresignResponse{value: val, isSet: true}
}

func (v NullableUploadPresignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadPresignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


