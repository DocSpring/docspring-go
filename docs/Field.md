# Field

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The field's ID, an integer that's unique within the template. | 
**Name** | **NullableString** | The field's name, the key of its value in submission data. A `/` nests the key in objects and arrays (`client/address`, `items/0/price`); write a literal `/` as `~1` and `~` as `~0`. Fields with the same name share one value and their shared settings, and array indexes inside a name are ignored for settings, so `items/0/price` and `items/1/price` share settings but not values. A `#` and a line number (`notes#0`, `notes#1`) splits one value over several fields as lines of text, and names that start with `@@` print submission details such as `@@date`. | 
**Page** | **int32** | The number of the page the field is on, starting at 1. A field added without one goes on page 1, and a field moved to another page without a new `y` keeps its place on the page. | 
**AcroformName** | **NullableString** | The name of the PDF form field this field was made from, when the template's PDF had fillable fields. Null for fields added in DocSpring. | 
**Type** | [**FieldType**](FieldType.md) | The data type of the field's value: `string`, `number`, `boolean`, `date`, `address`, `country`, `email`, `url`, `image`, `signature`, `initials`, `barcode`, or `combined` (made from other fields' values, never sent in submission data). It sets how the value is validated and which `display_type`s the field can use. `object` and `array` are the types of field groups (the template's `field_groups`), not of fields. Shared by every field with the same `name`. | 
**Static** | **bool** | The field prints fixed content, its `default`, instead of a value from submission data, and it's left out of the template's data schema and forms. Static check marks and shapes always print unless `hidden`, and a static number field's `default` can be a formula. Shared by every field with the same `name`. | 
**Editable** | **bool** | Text and check display only: prints the field as a fillable PDF form field instead of flat content (a text field with an `option_list` becomes a dropdown). A submission's `editable` option, when set, overrides this for every field. Shared by every field with the same `name`. | 
**Default** | **NullableString** | The value used when a submission leaves the field out, and the value forms are pre-filled with. A required field still needs a value in each submission. An image field's default is an image URL, a number field's can be a formula, and a static field's is the content it prints. Shared by every field with the same `name`. | 
**Title** | **string** | The field's title in the template's data schema and on forms. When it's empty, a title is made from the field's name. Shared by every field with the same `name`. | 
**Description** | **string** | Help text for the field, added to the template's data schema and shown on forms. Shared by every field with the same `name`. | 
**Placeholder** | **string** | Placeholder text for the field's input on forms, added to the template's data schema as `x-placeholder`. Shared by every field with the same `name`. | 
**Metadata** | **string** | Text for your own use, added to the template's data schema as `x-metadata`. DocSpring doesn't use it. Shared by every field with the same `name`. | 
**Required** | **bool** | Whether each submission must include a value for the field. A submission without one fails validation, and no PDF is generated. Shared by every field with the same `name`. | 
**OptionList** | **NullableString** | String fields only: a comma-separated list of the values the field accepts. It's added to the template's data schema as an `enum`, forms show it as a select list, and an `editable` field becomes a dropdown. Shared by every field with the same `name`. | 
**CombinedFieldNames** | **string** | Combined fields only: a comma-separated list of the names of the fields whose values are combined, in order. Shared by every field with the same `name`. | 
**CombinedFieldType** | [**CombinedFieldType**](CombinedFieldType.md) | Combined fields only: `separator` joins the values with `combined_field_separator`, and `format` puts them into `combined_field_format`. Shared by every field with the same `name`. | 
**CombinedFieldSeparator** | **string** | Combined fields only: the text placed between the values when `combined_field_type` is `separator`. Empty values are skipped. Defaults to a space. Shared by every field with the same `name`. | 
**CombinedFieldFormat** | **string** | Combined fields only: the text for `combined_field_type` `format`, where `$0`, `$1` and so on are replaced with the values of `combined_field_names` in order, such as `Full name: $0, Occupation: $1`. Shared by every field with the same `name`. | 
**Integer** | **bool** | Number fields only: the value must be a whole number. Shared by every field with the same `name`. | 
**Minimum** | **NullableFloat32** | Number fields only: the smallest value allowed. Shared by every field with the same `name`. | 
**Maximum** | **NullableFloat32** | Number fields only: the largest value allowed. Shared by every field with the same `name`. | 
**ExclusiveMinimum** | **bool** | Number fields only: the value must be greater than `minimum`, not equal to it. Shared by every field with the same `name`. | 
**ExclusiveMaximum** | **bool** | Number fields only: the value must be less than `maximum`, not equal to it. Shared by every field with the same `name`. | 
**IncludeTime** | **NullableBool** | Date fields only: the value is a date and time (`YYYY-MM-DDThh:mm:ss.fffZ`) instead of a date (`YYYY-MM-DD`), and it prints with the template's date-time format unless `date_time_format` is set. Shared by every field with the same `name`. | 
**SignatureAllowDraw** | **bool** | Signature and initials fields only: lets signers draw their signature on forms and in data requests. Shared by every field with the same `name`. | 
**SignatureAllowType** | **bool** | Signature and initials fields only: lets signers type their name and choose a script font instead of drawing, on forms and in data requests. Shared by every field with the same `name`. | 
**BarcodeSymbology** | [**BarcodeSymbology**](BarcodeSymbology.md) | Barcode fields only: the barcode format, which the value must match. `code_128` and `pdf417` take any ASCII text; `code_39` and `code_93` take capital letters, digits, spaces and `- . $ / + %`; `code_25` takes digits; `ean_13` takes 12 digits, `ean_8` 7 digits and `upc_a` 11 digits (the check digit is added). Shared by every field with the same `name`. | 
**Hidden** | **bool** | Leaves the field off the PDF. It stays in the template's data schema (marked `x-hidden`) and on forms, so other fields, such as combined fields, can use its value. A hidden static check mark or shape isn't printed. | 
**X** | **float32** | The distance from the page's left edge to the field's left edge. Positions and sizes are in units where every page is 2480 wide: one unit is the page's width in points divided by 2480 (about 0.247 points on a US Letter page). | 
**Y** | **float32** | The distance from the top edge of the field's `page` down to the field's top edge, in the units of `x`. A page is 2480 times its height divided by its width units tall: about 3209 for US Letter and 3508 for A4. | 
**Width** | **float32** | The field's width, in the units of `x`. A check mark is sized to the field's width. | 
**Height** | **float32** | The field's height, in the units of `x`. The template editor keeps check marks, QR codes, and square and circle shapes square, with `height` equal to `width`. | 
**Rotation** | **float32** | Rotates the field clockwise by this many degrees, from -360 to 360 (negative turns it counterclockwise). | 
**DisplayType** | [**FieldDisplayType**](FieldDisplayType.md) | How the value appears on the PDF: `text`, `check` (a check mark), `shape`, `qrcode`, `barcode`, or `image` (image, signature and initials fields). Check marks and shapes print only when the value meets the field's `condition`, and the field's `type` limits which display types it can use. | 
**Opacity** | **float32** | The opacity of the field's content and background, from 0 (invisible) to 1. QR codes and barcodes always print at full opacity. | 
**DateTimeFormat** | **NullableString** | Date fields only: how the date prints, as a strftime format such as `%m/%d/%Y`, or a calculation such as `@years_since`, `@years_until` or `@diff_years` (also months, weeks, days, hours, minutes and seconds), counted to when the PDF is generated. Empty uses the template's date format, or its date-time format with `include_time`. | 
**Typeface** | [**NullableTypeface**](Typeface.md) | Text display only: the font, one of `Helvetica`, `Arial`, `Calibri`, `Open Sans`, `Roboto`, `Lato`, `Poppins`, `Public Sans`, `Questrial`, `Times New Roman`, `Dancing Script`, `Lovtony Script`, or the monospace `Courier`, `Inconsolata`, `Menlo`, `Roboto Mono` and `Ubuntu Mono`. Empty uses the template's default typeface. | 
**FontSize** | **NullableInt32** | Text display only: the font size in points. Empty uses the template's default font size. | 
**CharacterSpacing** | **float32** | Text display only: extra space between characters, in points. Empty uses the template's default. | 
**Bold** | **bool** | Text display only: prints the text in bold. | 
**Uppercase** | **bool** | Text display only: prints the text in capital letters. | 
**Strikethrough** | **bool** | Text display only: draws a line through the text. | 
**KeepLeadingSpaces** | **bool** | Text display only: keeps the spaces at the start of the text and of each line, so text can be lined up with spaces. | 
**Overflow** | [**TextOverflow**](TextOverflow.md) | Text display only: what happens to text that doesn't fit the field. `shrink_to_fit` makes the font smaller, down to 5 points; `truncate` cuts the text off and returns the rest in the submission's `truncated_text`. | 
**Multiline** | **bool** | Text display only: wraps the text over as many lines as fit in the field's height instead of printing one line, and forms show a multi-line input. | 
**MultilineLines** | **NullableInt32** | Text display only, with `multiline`: spreads the text over exactly this many evenly spaced lines in the field's height (up to 300), to line up with lines printed on the form. Empty uses the font's normal line spacing. | 
**MinLength** | **float32** | Text display only: the fewest characters the value can have, added to the template's data schema. Not used for date, address, country or number fields, or fields with an `option_list`. | 
**MaxLength** | **NullableInt32** | Text display only: the most characters the value can have, added to the template's data schema; it also limits an `editable` field. Not used for date, address, country or number fields, or fields with an `option_list`. | 
**AutoCalculateMaxLength** | **bool** | Text display only, with a monospace `typeface` (`Courier`, `Inconsolata`, `Menlo`, `Roboto Mono` or `Ubuntu Mono`): adds the number of characters that fit in the field to the template's data schema as its maximum length. | 
**Color** | **NullableString** | Text display only: the text color, as a hex code such as `444444`. Empty uses the template's default text color. | 
**ColorFieldName** | **NullableString** | Text display only: the name of a data field whose value, a hex code such as `FF0000`, sets the text color for each submission. That field is added to the template's data schema, with `color` as its default. | 
**ColorFieldRequired** | **bool** | Makes the `color_field_name` field required. When it isn't required and a submission leaves it out, `color` is used. | 
**BackgroundColor** | **NullableString** | Text and check display only: a color that fills the field's area behind its content, as a hex code, to cover what the page shows there. Check fields print it even when unchecked; text fields print it only when they have a value. | 
**BackgroundColorFieldName** | **NullableString** | Text and check display only: the name of a data field whose value, a hex code, sets the background color for each submission. That field is added to the template's data schema, with `background_color` as its default. | 
**BackgroundColorFieldRequired** | **bool** | Makes the `background_color_field_name` field required. When it isn't required and a submission leaves it out, `background_color` is used. | 
**Alignment** | [**TextAlignment**](TextAlignment.md) | Text display only: how the text is aligned across the field: `left`, `center`, `right` or `justify` (both edges of every line of multiline text except the last of each paragraph). | 
**VAlignment** | [**TextVerticalAlignment**](TextVerticalAlignment.md) | Text display only: how the text is aligned vertically in the field: `bottom`, `center` or `top`. | 
**Comb** | **bool** | Text display only: splits the field into `comb_number_of_cells` boxes of equal width and prints one character, centered, in each. | 
**CombNumberOfCells** | **int32** | With `comb`: how many boxes the field has. The value's maximum length in the template's data schema is this plus `comb_value_offset`. | 
**CombValueOffset** | **int32** | With `comb`: the position, from 0, of the first character of the value that this field prints. For an SSN printed in boxes of 3, 2 and 4 digits, three fields named `ssn` use offsets 0, 3 and 5. | 
**ShapeType** | [**ShapeType**](ShapeType.md) | Shape display only: `square`, `rectangle`, `circle`, `ellipse`, `triangle` or `polygon`. A square or circle uses the smaller of `width` and `height`. | 
**ShapeSides** | **int32** | Shape display only, with `shape_type` `polygon`: the number of sides, 3 to 12. | 
**ShapeBorderWidth** | **int32** | Shape display only: the width of the shape's border, in points. | 
**ShapeBorderColor** | **string** | Shape display only: the shape's border color, as a hex code. Null draws no border. | 
**ShapeBorderColorFieldName** | **NullableString** | Shape display only: the name of a data field whose value, a hex code, sets the border color for each submission. That field is added to the template's data schema, with `shape_border_color` as its default. | 
**ShapeBorderColorFieldRequired** | **bool** | Makes the `shape_border_color_field_name` field required. When it isn't required and a submission leaves it out, `shape_border_color` is used. | 
**ShapeFillColor** | **NullableString** | Shape display only: the color that fills the shape, as a hex code. Null leaves it unfilled. | 
**ShapeFillColorFieldName** | **NullableString** | Shape display only: the name of a data field whose value, a hex code, sets the fill color for each submission. That field is added to the template's data schema, with `shape_fill_color` as its default. | 
**ShapeFillColorFieldRequired** | **bool** | Makes the `shape_fill_color_field_name` field required. When it isn't required and a submission leaves it out, `shape_fill_color` is used. | 
**CheckCharacter** | [**CheckCharacter**](CheckCharacter.md) | Check display only: the character printed, as an HTML entity: `&amp;#10003;` (✓), `&amp;#10004;` (✔, the default), `&amp;#10006;` (✖), `&amp;#10007;` (✗) or `&amp;#10008;` (✘). | 
**CheckColor** | **NullableString** | Check display only: the check mark's color, as a hex code. Empty uses the template's default check color. | 
**CheckColorFieldName** | **NullableString** | Check display only: the name of a data field whose value, a hex code, sets the check mark's color for each submission. That field is added to the template's data schema, with `check_color` as its default. | 
**CheckColorFieldRequired** | **bool** | Makes the `check_color_field_name` field required. When it isn't required and a submission leaves it out, `check_color` is used. | 
**QrcodeColor** | **NullableString** | QR code display only: the QR code's color, as a hex code, printed on a white background. Empty prints it in black. | 
**QrcodeColorFieldName** | **NullableString** | QR code display only: the name of a data field whose value, a hex code, sets the QR code's color for each submission. That field is added to the template's data schema, with `qrcode_color` as its default. | 
**QrcodeColorFieldRequired** | **bool** | Makes the `qrcode_color_field_name` field required. When it isn't required and a submission leaves it out, `qrcode_color` is used. | 
**TrueText** | **NullableString** | Boolean fields with text display: the text printed when the value is `true`. Empty uses the template's `true_text`; when neither is set, nothing prints. | 
**FalseText** | **NullableString** | Boolean fields with text display: the text printed when the value is `false`. Empty uses the template's `false_text`; when neither is set, nothing prints. | 
**Condition** | **NullableString** | Check and shape display only: the value a string or number field's value is compared with, by `string_condition_type` or `number_condition_type`, to decide whether the check mark or shape prints. Boolean fields print when `true` (see `invert_boolean_condition`), and other types print when they have a value. | 
**StringConditionType** | [**StringConditionType**](StringConditionType.md) | String fields with check or shape display: how the value is compared with `condition`: `equals`, `contains`, `starts_with`, `ends_with`, or `regex`, where `condition` is a regular expression the value must match. | 
**NumberConditionType** | [**NumberConditionType**](NumberConditionType.md) | Number fields with check or shape display: how the value is compared with `condition`: `equals`, `gt` (greater than), `gte` (greater than or equal to), `lt` (less than), `lte` (less than or equal to), or `range`, between `number_condition_range_min` and `number_condition_range_max`. | 
**NumberConditionRangeMin** | **NullableFloat32** | With `number_condition_type` `range`: the lower end of the range, included unless `number_condition_range_exclusive_min` is set. | 
**NumberConditionRangeMax** | **NullableFloat32** | With `number_condition_type` `range`: the upper end of the range, included unless `number_condition_range_exclusive_max` is set. | 
**NumberConditionRangeExclusiveMin** | **bool** | With `number_condition_type` `range`: leaves `number_condition_range_min` itself out of the range. | 
**NumberConditionRangeExclusiveMax** | **bool** | With `number_condition_type` `range`: leaves `number_condition_range_max` itself out of the range. | 
**InvertBooleanCondition** | **bool** | Boolean fields with check or shape display: prints the check mark or shape when the value is `false` or missing, instead of `true`. | 
**Currency** | **bool** | Number fields only: forms show a currency input, and the template's data schema gives the field `format: currency`. It doesn't change how the number prints. | 
**DecimalPlaces** | **NullableInt32** | Number fields with text display: prints the number with this many decimal places, from 0 (rounds to a whole number) to 99. Empty prints the number as it was sent. | 
**ImageScaleType** | [**ImageScaleType**](ImageScaleType.md) | Image fields only: how the image is resized to the field. `fit` keeps its proportions and fits it inside the field, which can leave empty space; `fill` keeps its proportions and crops it to fill the field; `stretch` stretches it to the field's size. Signatures and initials always fit. | 
**ImageGravity** | [**ImageGravity**](ImageGravity.md) | Image, signature and initials fields: where the image sits in the field with `fit`, or which part is kept with `fill`: `NorthWest` (top left), `North`, `NorthEast`, `West`, `Center`, `East`, `SouthWest`, `South` or `SouthEast` (bottom right). It also aligns typed signatures, and isn't used with `stretch`. | 
**Dependencies** | [**NullableFieldDependency**](FieldDependency.md) | The field's own conditional logic rule: when other fields' values meet its conditions, the field is shown, made required, or both. Null when it has none. The template's rules (`field_dependencies`) can apply to the field too. Shared by every field with the same `name`. | 

## Methods

### NewField

`func NewField(id int32, name NullableString, page int32, acroformName NullableString, type_ FieldType, static bool, editable bool, default_ NullableString, title string, description string, placeholder string, metadata string, required bool, optionList NullableString, combinedFieldNames string, combinedFieldType CombinedFieldType, combinedFieldSeparator string, combinedFieldFormat string, integer bool, minimum NullableFloat32, maximum NullableFloat32, exclusiveMinimum bool, exclusiveMaximum bool, includeTime NullableBool, signatureAllowDraw bool, signatureAllowType bool, barcodeSymbology BarcodeSymbology, hidden bool, x float32, y float32, width float32, height float32, rotation float32, displayType FieldDisplayType, opacity float32, dateTimeFormat NullableString, typeface NullableTypeface, fontSize NullableInt32, characterSpacing float32, bold bool, uppercase bool, strikethrough bool, keepLeadingSpaces bool, overflow TextOverflow, multiline bool, multilineLines NullableInt32, minLength float32, maxLength NullableInt32, autoCalculateMaxLength bool, color NullableString, colorFieldName NullableString, colorFieldRequired bool, backgroundColor NullableString, backgroundColorFieldName NullableString, backgroundColorFieldRequired bool, alignment TextAlignment, vAlignment TextVerticalAlignment, comb bool, combNumberOfCells int32, combValueOffset int32, shapeType ShapeType, shapeSides int32, shapeBorderWidth int32, shapeBorderColor string, shapeBorderColorFieldName NullableString, shapeBorderColorFieldRequired bool, shapeFillColor NullableString, shapeFillColorFieldName NullableString, shapeFillColorFieldRequired bool, checkCharacter CheckCharacter, checkColor NullableString, checkColorFieldName NullableString, checkColorFieldRequired bool, qrcodeColor NullableString, qrcodeColorFieldName NullableString, qrcodeColorFieldRequired bool, trueText NullableString, falseText NullableString, condition NullableString, stringConditionType StringConditionType, numberConditionType NumberConditionType, numberConditionRangeMin NullableFloat32, numberConditionRangeMax NullableFloat32, numberConditionRangeExclusiveMin bool, numberConditionRangeExclusiveMax bool, invertBooleanCondition bool, currency bool, decimalPlaces NullableInt32, imageScaleType ImageScaleType, imageGravity ImageGravity, dependencies NullableFieldDependency, ) *Field`

NewField instantiates a new Field object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldWithDefaults

`func NewFieldWithDefaults() *Field`

NewFieldWithDefaults instantiates a new Field object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Field) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Field) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Field) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Field) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Field) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Field) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Field) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Field) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPage

`func (o *Field) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *Field) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *Field) SetPage(v int32)`

SetPage sets Page field to given value.


### GetAcroformName

`func (o *Field) GetAcroformName() string`

GetAcroformName returns the AcroformName field if non-nil, zero value otherwise.

### GetAcroformNameOk

`func (o *Field) GetAcroformNameOk() (*string, bool)`

GetAcroformNameOk returns a tuple with the AcroformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcroformName

`func (o *Field) SetAcroformName(v string)`

SetAcroformName sets AcroformName field to given value.


### SetAcroformNameNil

`func (o *Field) SetAcroformNameNil(b bool)`

 SetAcroformNameNil sets the value for AcroformName to be an explicit nil

### UnsetAcroformName
`func (o *Field) UnsetAcroformName()`

UnsetAcroformName ensures that no value is present for AcroformName, not even an explicit nil
### GetType

`func (o *Field) GetType() FieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Field) GetTypeOk() (*FieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Field) SetType(v FieldType)`

SetType sets Type field to given value.


### GetStatic

`func (o *Field) GetStatic() bool`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *Field) GetStaticOk() (*bool, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *Field) SetStatic(v bool)`

SetStatic sets Static field to given value.


### GetEditable

`func (o *Field) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *Field) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *Field) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetDefault

`func (o *Field) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Field) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Field) SetDefault(v string)`

SetDefault sets Default field to given value.


### SetDefaultNil

`func (o *Field) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *Field) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetTitle

`func (o *Field) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Field) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Field) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Field) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Field) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Field) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPlaceholder

`func (o *Field) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *Field) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *Field) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.


### GetMetadata

`func (o *Field) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Field) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Field) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetRequired

`func (o *Field) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Field) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Field) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetOptionList

`func (o *Field) GetOptionList() string`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *Field) GetOptionListOk() (*string, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *Field) SetOptionList(v string)`

SetOptionList sets OptionList field to given value.


### SetOptionListNil

`func (o *Field) SetOptionListNil(b bool)`

 SetOptionListNil sets the value for OptionList to be an explicit nil

### UnsetOptionList
`func (o *Field) UnsetOptionList()`

UnsetOptionList ensures that no value is present for OptionList, not even an explicit nil
### GetCombinedFieldNames

`func (o *Field) GetCombinedFieldNames() string`

GetCombinedFieldNames returns the CombinedFieldNames field if non-nil, zero value otherwise.

### GetCombinedFieldNamesOk

`func (o *Field) GetCombinedFieldNamesOk() (*string, bool)`

GetCombinedFieldNamesOk returns a tuple with the CombinedFieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldNames

`func (o *Field) SetCombinedFieldNames(v string)`

SetCombinedFieldNames sets CombinedFieldNames field to given value.


### GetCombinedFieldType

`func (o *Field) GetCombinedFieldType() CombinedFieldType`

GetCombinedFieldType returns the CombinedFieldType field if non-nil, zero value otherwise.

### GetCombinedFieldTypeOk

`func (o *Field) GetCombinedFieldTypeOk() (*CombinedFieldType, bool)`

GetCombinedFieldTypeOk returns a tuple with the CombinedFieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldType

`func (o *Field) SetCombinedFieldType(v CombinedFieldType)`

SetCombinedFieldType sets CombinedFieldType field to given value.


### GetCombinedFieldSeparator

`func (o *Field) GetCombinedFieldSeparator() string`

GetCombinedFieldSeparator returns the CombinedFieldSeparator field if non-nil, zero value otherwise.

### GetCombinedFieldSeparatorOk

`func (o *Field) GetCombinedFieldSeparatorOk() (*string, bool)`

GetCombinedFieldSeparatorOk returns a tuple with the CombinedFieldSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldSeparator

`func (o *Field) SetCombinedFieldSeparator(v string)`

SetCombinedFieldSeparator sets CombinedFieldSeparator field to given value.


### GetCombinedFieldFormat

`func (o *Field) GetCombinedFieldFormat() string`

GetCombinedFieldFormat returns the CombinedFieldFormat field if non-nil, zero value otherwise.

### GetCombinedFieldFormatOk

`func (o *Field) GetCombinedFieldFormatOk() (*string, bool)`

GetCombinedFieldFormatOk returns a tuple with the CombinedFieldFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldFormat

`func (o *Field) SetCombinedFieldFormat(v string)`

SetCombinedFieldFormat sets CombinedFieldFormat field to given value.


### GetInteger

`func (o *Field) GetInteger() bool`

GetInteger returns the Integer field if non-nil, zero value otherwise.

### GetIntegerOk

`func (o *Field) GetIntegerOk() (*bool, bool)`

GetIntegerOk returns a tuple with the Integer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteger

`func (o *Field) SetInteger(v bool)`

SetInteger sets Integer field to given value.


### GetMinimum

`func (o *Field) GetMinimum() float32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *Field) GetMinimumOk() (*float32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *Field) SetMinimum(v float32)`

SetMinimum sets Minimum field to given value.


### SetMinimumNil

`func (o *Field) SetMinimumNil(b bool)`

 SetMinimumNil sets the value for Minimum to be an explicit nil

### UnsetMinimum
`func (o *Field) UnsetMinimum()`

UnsetMinimum ensures that no value is present for Minimum, not even an explicit nil
### GetMaximum

`func (o *Field) GetMaximum() float32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *Field) GetMaximumOk() (*float32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *Field) SetMaximum(v float32)`

SetMaximum sets Maximum field to given value.


### SetMaximumNil

`func (o *Field) SetMaximumNil(b bool)`

 SetMaximumNil sets the value for Maximum to be an explicit nil

### UnsetMaximum
`func (o *Field) UnsetMaximum()`

UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil
### GetExclusiveMinimum

`func (o *Field) GetExclusiveMinimum() bool`

GetExclusiveMinimum returns the ExclusiveMinimum field if non-nil, zero value otherwise.

### GetExclusiveMinimumOk

`func (o *Field) GetExclusiveMinimumOk() (*bool, bool)`

GetExclusiveMinimumOk returns a tuple with the ExclusiveMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMinimum

`func (o *Field) SetExclusiveMinimum(v bool)`

SetExclusiveMinimum sets ExclusiveMinimum field to given value.


### GetExclusiveMaximum

`func (o *Field) GetExclusiveMaximum() bool`

GetExclusiveMaximum returns the ExclusiveMaximum field if non-nil, zero value otherwise.

### GetExclusiveMaximumOk

`func (o *Field) GetExclusiveMaximumOk() (*bool, bool)`

GetExclusiveMaximumOk returns a tuple with the ExclusiveMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMaximum

`func (o *Field) SetExclusiveMaximum(v bool)`

SetExclusiveMaximum sets ExclusiveMaximum field to given value.


### GetIncludeTime

`func (o *Field) GetIncludeTime() bool`

GetIncludeTime returns the IncludeTime field if non-nil, zero value otherwise.

### GetIncludeTimeOk

`func (o *Field) GetIncludeTimeOk() (*bool, bool)`

GetIncludeTimeOk returns a tuple with the IncludeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTime

`func (o *Field) SetIncludeTime(v bool)`

SetIncludeTime sets IncludeTime field to given value.


### SetIncludeTimeNil

`func (o *Field) SetIncludeTimeNil(b bool)`

 SetIncludeTimeNil sets the value for IncludeTime to be an explicit nil

### UnsetIncludeTime
`func (o *Field) UnsetIncludeTime()`

UnsetIncludeTime ensures that no value is present for IncludeTime, not even an explicit nil
### GetSignatureAllowDraw

`func (o *Field) GetSignatureAllowDraw() bool`

GetSignatureAllowDraw returns the SignatureAllowDraw field if non-nil, zero value otherwise.

### GetSignatureAllowDrawOk

`func (o *Field) GetSignatureAllowDrawOk() (*bool, bool)`

GetSignatureAllowDrawOk returns a tuple with the SignatureAllowDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAllowDraw

`func (o *Field) SetSignatureAllowDraw(v bool)`

SetSignatureAllowDraw sets SignatureAllowDraw field to given value.


### GetSignatureAllowType

`func (o *Field) GetSignatureAllowType() bool`

GetSignatureAllowType returns the SignatureAllowType field if non-nil, zero value otherwise.

### GetSignatureAllowTypeOk

`func (o *Field) GetSignatureAllowTypeOk() (*bool, bool)`

GetSignatureAllowTypeOk returns a tuple with the SignatureAllowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAllowType

`func (o *Field) SetSignatureAllowType(v bool)`

SetSignatureAllowType sets SignatureAllowType field to given value.


### GetBarcodeSymbology

`func (o *Field) GetBarcodeSymbology() BarcodeSymbology`

GetBarcodeSymbology returns the BarcodeSymbology field if non-nil, zero value otherwise.

### GetBarcodeSymbologyOk

`func (o *Field) GetBarcodeSymbologyOk() (*BarcodeSymbology, bool)`

GetBarcodeSymbologyOk returns a tuple with the BarcodeSymbology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeSymbology

`func (o *Field) SetBarcodeSymbology(v BarcodeSymbology)`

SetBarcodeSymbology sets BarcodeSymbology field to given value.


### GetHidden

`func (o *Field) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Field) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Field) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetX

`func (o *Field) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *Field) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *Field) SetX(v float32)`

SetX sets X field to given value.


### GetY

`func (o *Field) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *Field) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *Field) SetY(v float32)`

SetY sets Y field to given value.


### GetWidth

`func (o *Field) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Field) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Field) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *Field) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Field) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Field) SetHeight(v float32)`

SetHeight sets Height field to given value.


### GetRotation

`func (o *Field) GetRotation() float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *Field) GetRotationOk() (*float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *Field) SetRotation(v float32)`

SetRotation sets Rotation field to given value.


### GetDisplayType

`func (o *Field) GetDisplayType() FieldDisplayType`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *Field) GetDisplayTypeOk() (*FieldDisplayType, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *Field) SetDisplayType(v FieldDisplayType)`

SetDisplayType sets DisplayType field to given value.


### GetOpacity

`func (o *Field) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *Field) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *Field) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.


### GetDateTimeFormat

`func (o *Field) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *Field) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *Field) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.


### SetDateTimeFormatNil

`func (o *Field) SetDateTimeFormatNil(b bool)`

 SetDateTimeFormatNil sets the value for DateTimeFormat to be an explicit nil

### UnsetDateTimeFormat
`func (o *Field) UnsetDateTimeFormat()`

UnsetDateTimeFormat ensures that no value is present for DateTimeFormat, not even an explicit nil
### GetTypeface

`func (o *Field) GetTypeface() Typeface`

GetTypeface returns the Typeface field if non-nil, zero value otherwise.

### GetTypefaceOk

`func (o *Field) GetTypefaceOk() (*Typeface, bool)`

GetTypefaceOk returns a tuple with the Typeface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeface

`func (o *Field) SetTypeface(v Typeface)`

SetTypeface sets Typeface field to given value.


### SetTypefaceNil

`func (o *Field) SetTypefaceNil(b bool)`

 SetTypefaceNil sets the value for Typeface to be an explicit nil

### UnsetTypeface
`func (o *Field) UnsetTypeface()`

UnsetTypeface ensures that no value is present for Typeface, not even an explicit nil
### GetFontSize

`func (o *Field) GetFontSize() int32`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *Field) GetFontSizeOk() (*int32, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *Field) SetFontSize(v int32)`

SetFontSize sets FontSize field to given value.


### SetFontSizeNil

`func (o *Field) SetFontSizeNil(b bool)`

 SetFontSizeNil sets the value for FontSize to be an explicit nil

### UnsetFontSize
`func (o *Field) UnsetFontSize()`

UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil
### GetCharacterSpacing

`func (o *Field) GetCharacterSpacing() float32`

GetCharacterSpacing returns the CharacterSpacing field if non-nil, zero value otherwise.

### GetCharacterSpacingOk

`func (o *Field) GetCharacterSpacingOk() (*float32, bool)`

GetCharacterSpacingOk returns a tuple with the CharacterSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterSpacing

`func (o *Field) SetCharacterSpacing(v float32)`

SetCharacterSpacing sets CharacterSpacing field to given value.


### GetBold

`func (o *Field) GetBold() bool`

GetBold returns the Bold field if non-nil, zero value otherwise.

### GetBoldOk

`func (o *Field) GetBoldOk() (*bool, bool)`

GetBoldOk returns a tuple with the Bold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBold

`func (o *Field) SetBold(v bool)`

SetBold sets Bold field to given value.


### GetUppercase

`func (o *Field) GetUppercase() bool`

GetUppercase returns the Uppercase field if non-nil, zero value otherwise.

### GetUppercaseOk

`func (o *Field) GetUppercaseOk() (*bool, bool)`

GetUppercaseOk returns a tuple with the Uppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUppercase

`func (o *Field) SetUppercase(v bool)`

SetUppercase sets Uppercase field to given value.


### GetStrikethrough

`func (o *Field) GetStrikethrough() bool`

GetStrikethrough returns the Strikethrough field if non-nil, zero value otherwise.

### GetStrikethroughOk

`func (o *Field) GetStrikethroughOk() (*bool, bool)`

GetStrikethroughOk returns a tuple with the Strikethrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikethrough

`func (o *Field) SetStrikethrough(v bool)`

SetStrikethrough sets Strikethrough field to given value.


### GetKeepLeadingSpaces

`func (o *Field) GetKeepLeadingSpaces() bool`

GetKeepLeadingSpaces returns the KeepLeadingSpaces field if non-nil, zero value otherwise.

### GetKeepLeadingSpacesOk

`func (o *Field) GetKeepLeadingSpacesOk() (*bool, bool)`

GetKeepLeadingSpacesOk returns a tuple with the KeepLeadingSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepLeadingSpaces

`func (o *Field) SetKeepLeadingSpaces(v bool)`

SetKeepLeadingSpaces sets KeepLeadingSpaces field to given value.


### GetOverflow

`func (o *Field) GetOverflow() TextOverflow`

GetOverflow returns the Overflow field if non-nil, zero value otherwise.

### GetOverflowOk

`func (o *Field) GetOverflowOk() (*TextOverflow, bool)`

GetOverflowOk returns a tuple with the Overflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflow

`func (o *Field) SetOverflow(v TextOverflow)`

SetOverflow sets Overflow field to given value.


### GetMultiline

`func (o *Field) GetMultiline() bool`

GetMultiline returns the Multiline field if non-nil, zero value otherwise.

### GetMultilineOk

`func (o *Field) GetMultilineOk() (*bool, bool)`

GetMultilineOk returns a tuple with the Multiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiline

`func (o *Field) SetMultiline(v bool)`

SetMultiline sets Multiline field to given value.


### GetMultilineLines

`func (o *Field) GetMultilineLines() int32`

GetMultilineLines returns the MultilineLines field if non-nil, zero value otherwise.

### GetMultilineLinesOk

`func (o *Field) GetMultilineLinesOk() (*int32, bool)`

GetMultilineLinesOk returns a tuple with the MultilineLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultilineLines

`func (o *Field) SetMultilineLines(v int32)`

SetMultilineLines sets MultilineLines field to given value.


### SetMultilineLinesNil

`func (o *Field) SetMultilineLinesNil(b bool)`

 SetMultilineLinesNil sets the value for MultilineLines to be an explicit nil

### UnsetMultilineLines
`func (o *Field) UnsetMultilineLines()`

UnsetMultilineLines ensures that no value is present for MultilineLines, not even an explicit nil
### GetMinLength

`func (o *Field) GetMinLength() float32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *Field) GetMinLengthOk() (*float32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *Field) SetMinLength(v float32)`

SetMinLength sets MinLength field to given value.


### GetMaxLength

`func (o *Field) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *Field) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *Field) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.


### SetMaxLengthNil

`func (o *Field) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *Field) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetAutoCalculateMaxLength

`func (o *Field) GetAutoCalculateMaxLength() bool`

GetAutoCalculateMaxLength returns the AutoCalculateMaxLength field if non-nil, zero value otherwise.

### GetAutoCalculateMaxLengthOk

`func (o *Field) GetAutoCalculateMaxLengthOk() (*bool, bool)`

GetAutoCalculateMaxLengthOk returns a tuple with the AutoCalculateMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCalculateMaxLength

`func (o *Field) SetAutoCalculateMaxLength(v bool)`

SetAutoCalculateMaxLength sets AutoCalculateMaxLength field to given value.


### GetColor

`func (o *Field) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Field) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Field) SetColor(v string)`

SetColor sets Color field to given value.


### SetColorNil

`func (o *Field) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *Field) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetColorFieldName

`func (o *Field) GetColorFieldName() string`

GetColorFieldName returns the ColorFieldName field if non-nil, zero value otherwise.

### GetColorFieldNameOk

`func (o *Field) GetColorFieldNameOk() (*string, bool)`

GetColorFieldNameOk returns a tuple with the ColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorFieldName

`func (o *Field) SetColorFieldName(v string)`

SetColorFieldName sets ColorFieldName field to given value.


### SetColorFieldNameNil

`func (o *Field) SetColorFieldNameNil(b bool)`

 SetColorFieldNameNil sets the value for ColorFieldName to be an explicit nil

### UnsetColorFieldName
`func (o *Field) UnsetColorFieldName()`

UnsetColorFieldName ensures that no value is present for ColorFieldName, not even an explicit nil
### GetColorFieldRequired

`func (o *Field) GetColorFieldRequired() bool`

GetColorFieldRequired returns the ColorFieldRequired field if non-nil, zero value otherwise.

### GetColorFieldRequiredOk

`func (o *Field) GetColorFieldRequiredOk() (*bool, bool)`

GetColorFieldRequiredOk returns a tuple with the ColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorFieldRequired

`func (o *Field) SetColorFieldRequired(v bool)`

SetColorFieldRequired sets ColorFieldRequired field to given value.


### GetBackgroundColor

`func (o *Field) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *Field) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *Field) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.


### SetBackgroundColorNil

`func (o *Field) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *Field) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetBackgroundColorFieldName

`func (o *Field) GetBackgroundColorFieldName() string`

GetBackgroundColorFieldName returns the BackgroundColorFieldName field if non-nil, zero value otherwise.

### GetBackgroundColorFieldNameOk

`func (o *Field) GetBackgroundColorFieldNameOk() (*string, bool)`

GetBackgroundColorFieldNameOk returns a tuple with the BackgroundColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColorFieldName

`func (o *Field) SetBackgroundColorFieldName(v string)`

SetBackgroundColorFieldName sets BackgroundColorFieldName field to given value.


### SetBackgroundColorFieldNameNil

`func (o *Field) SetBackgroundColorFieldNameNil(b bool)`

 SetBackgroundColorFieldNameNil sets the value for BackgroundColorFieldName to be an explicit nil

### UnsetBackgroundColorFieldName
`func (o *Field) UnsetBackgroundColorFieldName()`

UnsetBackgroundColorFieldName ensures that no value is present for BackgroundColorFieldName, not even an explicit nil
### GetBackgroundColorFieldRequired

`func (o *Field) GetBackgroundColorFieldRequired() bool`

GetBackgroundColorFieldRequired returns the BackgroundColorFieldRequired field if non-nil, zero value otherwise.

### GetBackgroundColorFieldRequiredOk

`func (o *Field) GetBackgroundColorFieldRequiredOk() (*bool, bool)`

GetBackgroundColorFieldRequiredOk returns a tuple with the BackgroundColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColorFieldRequired

`func (o *Field) SetBackgroundColorFieldRequired(v bool)`

SetBackgroundColorFieldRequired sets BackgroundColorFieldRequired field to given value.


### GetAlignment

`func (o *Field) GetAlignment() TextAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *Field) GetAlignmentOk() (*TextAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *Field) SetAlignment(v TextAlignment)`

SetAlignment sets Alignment field to given value.


### GetVAlignment

`func (o *Field) GetVAlignment() TextVerticalAlignment`

GetVAlignment returns the VAlignment field if non-nil, zero value otherwise.

### GetVAlignmentOk

`func (o *Field) GetVAlignmentOk() (*TextVerticalAlignment, bool)`

GetVAlignmentOk returns a tuple with the VAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAlignment

`func (o *Field) SetVAlignment(v TextVerticalAlignment)`

SetVAlignment sets VAlignment field to given value.


### GetComb

`func (o *Field) GetComb() bool`

GetComb returns the Comb field if non-nil, zero value otherwise.

### GetCombOk

`func (o *Field) GetCombOk() (*bool, bool)`

GetCombOk returns a tuple with the Comb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComb

`func (o *Field) SetComb(v bool)`

SetComb sets Comb field to given value.


### GetCombNumberOfCells

`func (o *Field) GetCombNumberOfCells() int32`

GetCombNumberOfCells returns the CombNumberOfCells field if non-nil, zero value otherwise.

### GetCombNumberOfCellsOk

`func (o *Field) GetCombNumberOfCellsOk() (*int32, bool)`

GetCombNumberOfCellsOk returns a tuple with the CombNumberOfCells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombNumberOfCells

`func (o *Field) SetCombNumberOfCells(v int32)`

SetCombNumberOfCells sets CombNumberOfCells field to given value.


### GetCombValueOffset

`func (o *Field) GetCombValueOffset() int32`

GetCombValueOffset returns the CombValueOffset field if non-nil, zero value otherwise.

### GetCombValueOffsetOk

`func (o *Field) GetCombValueOffsetOk() (*int32, bool)`

GetCombValueOffsetOk returns a tuple with the CombValueOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombValueOffset

`func (o *Field) SetCombValueOffset(v int32)`

SetCombValueOffset sets CombValueOffset field to given value.


### GetShapeType

`func (o *Field) GetShapeType() ShapeType`

GetShapeType returns the ShapeType field if non-nil, zero value otherwise.

### GetShapeTypeOk

`func (o *Field) GetShapeTypeOk() (*ShapeType, bool)`

GetShapeTypeOk returns a tuple with the ShapeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeType

`func (o *Field) SetShapeType(v ShapeType)`

SetShapeType sets ShapeType field to given value.


### GetShapeSides

`func (o *Field) GetShapeSides() int32`

GetShapeSides returns the ShapeSides field if non-nil, zero value otherwise.

### GetShapeSidesOk

`func (o *Field) GetShapeSidesOk() (*int32, bool)`

GetShapeSidesOk returns a tuple with the ShapeSides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeSides

`func (o *Field) SetShapeSides(v int32)`

SetShapeSides sets ShapeSides field to given value.


### GetShapeBorderWidth

`func (o *Field) GetShapeBorderWidth() int32`

GetShapeBorderWidth returns the ShapeBorderWidth field if non-nil, zero value otherwise.

### GetShapeBorderWidthOk

`func (o *Field) GetShapeBorderWidthOk() (*int32, bool)`

GetShapeBorderWidthOk returns a tuple with the ShapeBorderWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderWidth

`func (o *Field) SetShapeBorderWidth(v int32)`

SetShapeBorderWidth sets ShapeBorderWidth field to given value.


### GetShapeBorderColor

`func (o *Field) GetShapeBorderColor() string`

GetShapeBorderColor returns the ShapeBorderColor field if non-nil, zero value otherwise.

### GetShapeBorderColorOk

`func (o *Field) GetShapeBorderColorOk() (*string, bool)`

GetShapeBorderColorOk returns a tuple with the ShapeBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderColor

`func (o *Field) SetShapeBorderColor(v string)`

SetShapeBorderColor sets ShapeBorderColor field to given value.


### GetShapeBorderColorFieldName

`func (o *Field) GetShapeBorderColorFieldName() string`

GetShapeBorderColorFieldName returns the ShapeBorderColorFieldName field if non-nil, zero value otherwise.

### GetShapeBorderColorFieldNameOk

`func (o *Field) GetShapeBorderColorFieldNameOk() (*string, bool)`

GetShapeBorderColorFieldNameOk returns a tuple with the ShapeBorderColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderColorFieldName

`func (o *Field) SetShapeBorderColorFieldName(v string)`

SetShapeBorderColorFieldName sets ShapeBorderColorFieldName field to given value.


### SetShapeBorderColorFieldNameNil

`func (o *Field) SetShapeBorderColorFieldNameNil(b bool)`

 SetShapeBorderColorFieldNameNil sets the value for ShapeBorderColorFieldName to be an explicit nil

### UnsetShapeBorderColorFieldName
`func (o *Field) UnsetShapeBorderColorFieldName()`

UnsetShapeBorderColorFieldName ensures that no value is present for ShapeBorderColorFieldName, not even an explicit nil
### GetShapeBorderColorFieldRequired

`func (o *Field) GetShapeBorderColorFieldRequired() bool`

GetShapeBorderColorFieldRequired returns the ShapeBorderColorFieldRequired field if non-nil, zero value otherwise.

### GetShapeBorderColorFieldRequiredOk

`func (o *Field) GetShapeBorderColorFieldRequiredOk() (*bool, bool)`

GetShapeBorderColorFieldRequiredOk returns a tuple with the ShapeBorderColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderColorFieldRequired

`func (o *Field) SetShapeBorderColorFieldRequired(v bool)`

SetShapeBorderColorFieldRequired sets ShapeBorderColorFieldRequired field to given value.


### GetShapeFillColor

`func (o *Field) GetShapeFillColor() string`

GetShapeFillColor returns the ShapeFillColor field if non-nil, zero value otherwise.

### GetShapeFillColorOk

`func (o *Field) GetShapeFillColorOk() (*string, bool)`

GetShapeFillColorOk returns a tuple with the ShapeFillColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeFillColor

`func (o *Field) SetShapeFillColor(v string)`

SetShapeFillColor sets ShapeFillColor field to given value.


### SetShapeFillColorNil

`func (o *Field) SetShapeFillColorNil(b bool)`

 SetShapeFillColorNil sets the value for ShapeFillColor to be an explicit nil

### UnsetShapeFillColor
`func (o *Field) UnsetShapeFillColor()`

UnsetShapeFillColor ensures that no value is present for ShapeFillColor, not even an explicit nil
### GetShapeFillColorFieldName

`func (o *Field) GetShapeFillColorFieldName() string`

GetShapeFillColorFieldName returns the ShapeFillColorFieldName field if non-nil, zero value otherwise.

### GetShapeFillColorFieldNameOk

`func (o *Field) GetShapeFillColorFieldNameOk() (*string, bool)`

GetShapeFillColorFieldNameOk returns a tuple with the ShapeFillColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeFillColorFieldName

`func (o *Field) SetShapeFillColorFieldName(v string)`

SetShapeFillColorFieldName sets ShapeFillColorFieldName field to given value.


### SetShapeFillColorFieldNameNil

`func (o *Field) SetShapeFillColorFieldNameNil(b bool)`

 SetShapeFillColorFieldNameNil sets the value for ShapeFillColorFieldName to be an explicit nil

### UnsetShapeFillColorFieldName
`func (o *Field) UnsetShapeFillColorFieldName()`

UnsetShapeFillColorFieldName ensures that no value is present for ShapeFillColorFieldName, not even an explicit nil
### GetShapeFillColorFieldRequired

`func (o *Field) GetShapeFillColorFieldRequired() bool`

GetShapeFillColorFieldRequired returns the ShapeFillColorFieldRequired field if non-nil, zero value otherwise.

### GetShapeFillColorFieldRequiredOk

`func (o *Field) GetShapeFillColorFieldRequiredOk() (*bool, bool)`

GetShapeFillColorFieldRequiredOk returns a tuple with the ShapeFillColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeFillColorFieldRequired

`func (o *Field) SetShapeFillColorFieldRequired(v bool)`

SetShapeFillColorFieldRequired sets ShapeFillColorFieldRequired field to given value.


### GetCheckCharacter

`func (o *Field) GetCheckCharacter() CheckCharacter`

GetCheckCharacter returns the CheckCharacter field if non-nil, zero value otherwise.

### GetCheckCharacterOk

`func (o *Field) GetCheckCharacterOk() (*CheckCharacter, bool)`

GetCheckCharacterOk returns a tuple with the CheckCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCharacter

`func (o *Field) SetCheckCharacter(v CheckCharacter)`

SetCheckCharacter sets CheckCharacter field to given value.


### GetCheckColor

`func (o *Field) GetCheckColor() string`

GetCheckColor returns the CheckColor field if non-nil, zero value otherwise.

### GetCheckColorOk

`func (o *Field) GetCheckColorOk() (*string, bool)`

GetCheckColorOk returns a tuple with the CheckColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckColor

`func (o *Field) SetCheckColor(v string)`

SetCheckColor sets CheckColor field to given value.


### SetCheckColorNil

`func (o *Field) SetCheckColorNil(b bool)`

 SetCheckColorNil sets the value for CheckColor to be an explicit nil

### UnsetCheckColor
`func (o *Field) UnsetCheckColor()`

UnsetCheckColor ensures that no value is present for CheckColor, not even an explicit nil
### GetCheckColorFieldName

`func (o *Field) GetCheckColorFieldName() string`

GetCheckColorFieldName returns the CheckColorFieldName field if non-nil, zero value otherwise.

### GetCheckColorFieldNameOk

`func (o *Field) GetCheckColorFieldNameOk() (*string, bool)`

GetCheckColorFieldNameOk returns a tuple with the CheckColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckColorFieldName

`func (o *Field) SetCheckColorFieldName(v string)`

SetCheckColorFieldName sets CheckColorFieldName field to given value.


### SetCheckColorFieldNameNil

`func (o *Field) SetCheckColorFieldNameNil(b bool)`

 SetCheckColorFieldNameNil sets the value for CheckColorFieldName to be an explicit nil

### UnsetCheckColorFieldName
`func (o *Field) UnsetCheckColorFieldName()`

UnsetCheckColorFieldName ensures that no value is present for CheckColorFieldName, not even an explicit nil
### GetCheckColorFieldRequired

`func (o *Field) GetCheckColorFieldRequired() bool`

GetCheckColorFieldRequired returns the CheckColorFieldRequired field if non-nil, zero value otherwise.

### GetCheckColorFieldRequiredOk

`func (o *Field) GetCheckColorFieldRequiredOk() (*bool, bool)`

GetCheckColorFieldRequiredOk returns a tuple with the CheckColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckColorFieldRequired

`func (o *Field) SetCheckColorFieldRequired(v bool)`

SetCheckColorFieldRequired sets CheckColorFieldRequired field to given value.


### GetQrcodeColor

`func (o *Field) GetQrcodeColor() string`

GetQrcodeColor returns the QrcodeColor field if non-nil, zero value otherwise.

### GetQrcodeColorOk

`func (o *Field) GetQrcodeColorOk() (*string, bool)`

GetQrcodeColorOk returns a tuple with the QrcodeColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrcodeColor

`func (o *Field) SetQrcodeColor(v string)`

SetQrcodeColor sets QrcodeColor field to given value.


### SetQrcodeColorNil

`func (o *Field) SetQrcodeColorNil(b bool)`

 SetQrcodeColorNil sets the value for QrcodeColor to be an explicit nil

### UnsetQrcodeColor
`func (o *Field) UnsetQrcodeColor()`

UnsetQrcodeColor ensures that no value is present for QrcodeColor, not even an explicit nil
### GetQrcodeColorFieldName

`func (o *Field) GetQrcodeColorFieldName() string`

GetQrcodeColorFieldName returns the QrcodeColorFieldName field if non-nil, zero value otherwise.

### GetQrcodeColorFieldNameOk

`func (o *Field) GetQrcodeColorFieldNameOk() (*string, bool)`

GetQrcodeColorFieldNameOk returns a tuple with the QrcodeColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrcodeColorFieldName

`func (o *Field) SetQrcodeColorFieldName(v string)`

SetQrcodeColorFieldName sets QrcodeColorFieldName field to given value.


### SetQrcodeColorFieldNameNil

`func (o *Field) SetQrcodeColorFieldNameNil(b bool)`

 SetQrcodeColorFieldNameNil sets the value for QrcodeColorFieldName to be an explicit nil

### UnsetQrcodeColorFieldName
`func (o *Field) UnsetQrcodeColorFieldName()`

UnsetQrcodeColorFieldName ensures that no value is present for QrcodeColorFieldName, not even an explicit nil
### GetQrcodeColorFieldRequired

`func (o *Field) GetQrcodeColorFieldRequired() bool`

GetQrcodeColorFieldRequired returns the QrcodeColorFieldRequired field if non-nil, zero value otherwise.

### GetQrcodeColorFieldRequiredOk

`func (o *Field) GetQrcodeColorFieldRequiredOk() (*bool, bool)`

GetQrcodeColorFieldRequiredOk returns a tuple with the QrcodeColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrcodeColorFieldRequired

`func (o *Field) SetQrcodeColorFieldRequired(v bool)`

SetQrcodeColorFieldRequired sets QrcodeColorFieldRequired field to given value.


### GetTrueText

`func (o *Field) GetTrueText() string`

GetTrueText returns the TrueText field if non-nil, zero value otherwise.

### GetTrueTextOk

`func (o *Field) GetTrueTextOk() (*string, bool)`

GetTrueTextOk returns a tuple with the TrueText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueText

`func (o *Field) SetTrueText(v string)`

SetTrueText sets TrueText field to given value.


### SetTrueTextNil

`func (o *Field) SetTrueTextNil(b bool)`

 SetTrueTextNil sets the value for TrueText to be an explicit nil

### UnsetTrueText
`func (o *Field) UnsetTrueText()`

UnsetTrueText ensures that no value is present for TrueText, not even an explicit nil
### GetFalseText

`func (o *Field) GetFalseText() string`

GetFalseText returns the FalseText field if non-nil, zero value otherwise.

### GetFalseTextOk

`func (o *Field) GetFalseTextOk() (*string, bool)`

GetFalseTextOk returns a tuple with the FalseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseText

`func (o *Field) SetFalseText(v string)`

SetFalseText sets FalseText field to given value.


### SetFalseTextNil

`func (o *Field) SetFalseTextNil(b bool)`

 SetFalseTextNil sets the value for FalseText to be an explicit nil

### UnsetFalseText
`func (o *Field) UnsetFalseText()`

UnsetFalseText ensures that no value is present for FalseText, not even an explicit nil
### GetCondition

`func (o *Field) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *Field) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *Field) SetCondition(v string)`

SetCondition sets Condition field to given value.


### SetConditionNil

`func (o *Field) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *Field) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetStringConditionType

`func (o *Field) GetStringConditionType() StringConditionType`

GetStringConditionType returns the StringConditionType field if non-nil, zero value otherwise.

### GetStringConditionTypeOk

`func (o *Field) GetStringConditionTypeOk() (*StringConditionType, bool)`

GetStringConditionTypeOk returns a tuple with the StringConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringConditionType

`func (o *Field) SetStringConditionType(v StringConditionType)`

SetStringConditionType sets StringConditionType field to given value.


### GetNumberConditionType

`func (o *Field) GetNumberConditionType() NumberConditionType`

GetNumberConditionType returns the NumberConditionType field if non-nil, zero value otherwise.

### GetNumberConditionTypeOk

`func (o *Field) GetNumberConditionTypeOk() (*NumberConditionType, bool)`

GetNumberConditionTypeOk returns a tuple with the NumberConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionType

`func (o *Field) SetNumberConditionType(v NumberConditionType)`

SetNumberConditionType sets NumberConditionType field to given value.


### GetNumberConditionRangeMin

`func (o *Field) GetNumberConditionRangeMin() float32`

GetNumberConditionRangeMin returns the NumberConditionRangeMin field if non-nil, zero value otherwise.

### GetNumberConditionRangeMinOk

`func (o *Field) GetNumberConditionRangeMinOk() (*float32, bool)`

GetNumberConditionRangeMinOk returns a tuple with the NumberConditionRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeMin

`func (o *Field) SetNumberConditionRangeMin(v float32)`

SetNumberConditionRangeMin sets NumberConditionRangeMin field to given value.


### SetNumberConditionRangeMinNil

`func (o *Field) SetNumberConditionRangeMinNil(b bool)`

 SetNumberConditionRangeMinNil sets the value for NumberConditionRangeMin to be an explicit nil

### UnsetNumberConditionRangeMin
`func (o *Field) UnsetNumberConditionRangeMin()`

UnsetNumberConditionRangeMin ensures that no value is present for NumberConditionRangeMin, not even an explicit nil
### GetNumberConditionRangeMax

`func (o *Field) GetNumberConditionRangeMax() float32`

GetNumberConditionRangeMax returns the NumberConditionRangeMax field if non-nil, zero value otherwise.

### GetNumberConditionRangeMaxOk

`func (o *Field) GetNumberConditionRangeMaxOk() (*float32, bool)`

GetNumberConditionRangeMaxOk returns a tuple with the NumberConditionRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeMax

`func (o *Field) SetNumberConditionRangeMax(v float32)`

SetNumberConditionRangeMax sets NumberConditionRangeMax field to given value.


### SetNumberConditionRangeMaxNil

`func (o *Field) SetNumberConditionRangeMaxNil(b bool)`

 SetNumberConditionRangeMaxNil sets the value for NumberConditionRangeMax to be an explicit nil

### UnsetNumberConditionRangeMax
`func (o *Field) UnsetNumberConditionRangeMax()`

UnsetNumberConditionRangeMax ensures that no value is present for NumberConditionRangeMax, not even an explicit nil
### GetNumberConditionRangeExclusiveMin

`func (o *Field) GetNumberConditionRangeExclusiveMin() bool`

GetNumberConditionRangeExclusiveMin returns the NumberConditionRangeExclusiveMin field if non-nil, zero value otherwise.

### GetNumberConditionRangeExclusiveMinOk

`func (o *Field) GetNumberConditionRangeExclusiveMinOk() (*bool, bool)`

GetNumberConditionRangeExclusiveMinOk returns a tuple with the NumberConditionRangeExclusiveMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeExclusiveMin

`func (o *Field) SetNumberConditionRangeExclusiveMin(v bool)`

SetNumberConditionRangeExclusiveMin sets NumberConditionRangeExclusiveMin field to given value.


### GetNumberConditionRangeExclusiveMax

`func (o *Field) GetNumberConditionRangeExclusiveMax() bool`

GetNumberConditionRangeExclusiveMax returns the NumberConditionRangeExclusiveMax field if non-nil, zero value otherwise.

### GetNumberConditionRangeExclusiveMaxOk

`func (o *Field) GetNumberConditionRangeExclusiveMaxOk() (*bool, bool)`

GetNumberConditionRangeExclusiveMaxOk returns a tuple with the NumberConditionRangeExclusiveMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeExclusiveMax

`func (o *Field) SetNumberConditionRangeExclusiveMax(v bool)`

SetNumberConditionRangeExclusiveMax sets NumberConditionRangeExclusiveMax field to given value.


### GetInvertBooleanCondition

`func (o *Field) GetInvertBooleanCondition() bool`

GetInvertBooleanCondition returns the InvertBooleanCondition field if non-nil, zero value otherwise.

### GetInvertBooleanConditionOk

`func (o *Field) GetInvertBooleanConditionOk() (*bool, bool)`

GetInvertBooleanConditionOk returns a tuple with the InvertBooleanCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertBooleanCondition

`func (o *Field) SetInvertBooleanCondition(v bool)`

SetInvertBooleanCondition sets InvertBooleanCondition field to given value.


### GetCurrency

`func (o *Field) GetCurrency() bool`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Field) GetCurrencyOk() (*bool, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Field) SetCurrency(v bool)`

SetCurrency sets Currency field to given value.


### GetDecimalPlaces

`func (o *Field) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *Field) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *Field) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.


### SetDecimalPlacesNil

`func (o *Field) SetDecimalPlacesNil(b bool)`

 SetDecimalPlacesNil sets the value for DecimalPlaces to be an explicit nil

### UnsetDecimalPlaces
`func (o *Field) UnsetDecimalPlaces()`

UnsetDecimalPlaces ensures that no value is present for DecimalPlaces, not even an explicit nil
### GetImageScaleType

`func (o *Field) GetImageScaleType() ImageScaleType`

GetImageScaleType returns the ImageScaleType field if non-nil, zero value otherwise.

### GetImageScaleTypeOk

`func (o *Field) GetImageScaleTypeOk() (*ImageScaleType, bool)`

GetImageScaleTypeOk returns a tuple with the ImageScaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScaleType

`func (o *Field) SetImageScaleType(v ImageScaleType)`

SetImageScaleType sets ImageScaleType field to given value.


### GetImageGravity

`func (o *Field) GetImageGravity() ImageGravity`

GetImageGravity returns the ImageGravity field if non-nil, zero value otherwise.

### GetImageGravityOk

`func (o *Field) GetImageGravityOk() (*ImageGravity, bool)`

GetImageGravityOk returns a tuple with the ImageGravity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageGravity

`func (o *Field) SetImageGravity(v ImageGravity)`

SetImageGravity sets ImageGravity field to given value.


### GetDependencies

`func (o *Field) GetDependencies() FieldDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *Field) GetDependenciesOk() (*FieldDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *Field) SetDependencies(v FieldDependency)`

SetDependencies sets Dependencies field to given value.


### SetDependenciesNil

`func (o *Field) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *Field) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


