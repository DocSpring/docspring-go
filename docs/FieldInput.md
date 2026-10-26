# FieldInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The field's name, the key of its value in submission data. A `/` nests the key in objects and arrays (`client/address`, `items/0/price`); write a literal `/` as `~1` and `~` as `~0`. Fields with the same name share one value and their shared settings, and array indexes inside a name are ignored for settings, so `items/0/price` and `items/1/price` share settings but not values. A `#` and a line number (`notes#0`, `notes#1`) splits one value over several fields as lines of text, and names that start with `@@` print submission details such as `@@date`. | [optional] 
**Page** | Pointer to **int32** | The number of the page the field is on, starting at 1. A field added without one goes on page 1, and a field moved to another page without a new `y` keeps its place on the page. | [optional] 
**Type** | Pointer to [**FieldType**](FieldType.md) | The data type of the field's value: `string`, `number`, `boolean`, `date`, `address`, `country`, `email`, `url`, `image`, `signature`, `initials`, `barcode`, or `combined` (made from other fields' values, never sent in submission data). It sets how the value is validated and which `display_type`s the field can use. `object` and `array` are the types of field groups (the template's `field_groups`), not of fields. Shared by every field with the same `name`. | [optional] 
**Static** | Pointer to **bool** | The field prints fixed content, its `default`, instead of a value from submission data, and it's left out of the template's data schema and forms. Static check marks and shapes always print unless `hidden`, and a static number field's `default` can be a formula. Shared by every field with the same `name`. | [optional] 
**Editable** | Pointer to **bool** | Text and check display only: prints the field as a fillable PDF form field instead of flat content (a text field with an `option_list` becomes a dropdown). A submission's `editable` option, when set, overrides this for every field. Shared by every field with the same `name`. | [optional] 
**Default** | Pointer to **NullableString** | The value used when a submission leaves the field out, and the value forms are pre-filled with. A required field still needs a value in each submission. An image field's default is an image URL, a number field's can be a formula, and a static field's is the content it prints. Shared by every field with the same `name`. | [optional] 
**Title** | Pointer to **string** | The field's title in the template's data schema and on forms. When it's empty, a title is made from the field's name. Shared by every field with the same `name`. | [optional] 
**Description** | Pointer to **string** | Help text for the field, added to the template's data schema and shown on forms. Shared by every field with the same `name`. | [optional] 
**Placeholder** | Pointer to **string** | Placeholder text for the field's input on forms, added to the template's data schema as `x-placeholder`. Shared by every field with the same `name`. | [optional] 
**Metadata** | Pointer to **string** | Text for your own use, added to the template's data schema as `x-metadata`. DocSpring doesn't use it. Shared by every field with the same `name`. | [optional] 
**Required** | Pointer to **bool** | Whether each submission must include a value for the field. A submission without one fails validation, and no PDF is generated. Shared by every field with the same `name`. | [optional] 
**OptionList** | Pointer to **NullableString** | String fields only: a comma-separated list of the values the field accepts. It's added to the template's data schema as an `enum`, forms show it as a select list, and an `editable` field becomes a dropdown. Shared by every field with the same `name`. | [optional] 
**CombinedFieldNames** | Pointer to **string** | Combined fields only: a comma-separated list of the names of the fields whose values are combined, in order. Shared by every field with the same `name`. | [optional] 
**CombinedFieldType** | Pointer to [**CombinedFieldType**](CombinedFieldType.md) | Combined fields only: `separator` joins the values with `combined_field_separator`, and `format` puts them into `combined_field_format`. Shared by every field with the same `name`. | [optional] 
**CombinedFieldSeparator** | Pointer to **string** | Combined fields only: the text placed between the values when `combined_field_type` is `separator`. Empty values are skipped. Defaults to a space. Shared by every field with the same `name`. | [optional] 
**CombinedFieldFormat** | Pointer to **string** | Combined fields only: the text for `combined_field_type` `format`, where `$0`, `$1` and so on are replaced with the values of `combined_field_names` in order, such as `Full name: $0, Occupation: $1`. Shared by every field with the same `name`. | [optional] 
**Integer** | Pointer to **bool** | Number fields only: the value must be a whole number. Shared by every field with the same `name`. | [optional] 
**Minimum** | Pointer to **NullableFloat32** | Number fields only: the smallest value allowed. Shared by every field with the same `name`. | [optional] 
**Maximum** | Pointer to **NullableFloat32** | Number fields only: the largest value allowed. Shared by every field with the same `name`. | [optional] 
**ExclusiveMinimum** | Pointer to **bool** | Number fields only: the value must be greater than `minimum`, not equal to it. Shared by every field with the same `name`. | [optional] 
**ExclusiveMaximum** | Pointer to **bool** | Number fields only: the value must be less than `maximum`, not equal to it. Shared by every field with the same `name`. | [optional] 
**IncludeTime** | Pointer to **NullableBool** | Date fields only: the value is a date and time (`YYYY-MM-DDThh:mm:ss.fffZ`) instead of a date (`YYYY-MM-DD`), and it prints with the template's date-time format unless `date_time_format` is set. Shared by every field with the same `name`. | [optional] 
**SignatureAllowDraw** | Pointer to **bool** | Signature and initials fields only: lets signers draw their signature on forms and in data requests. Shared by every field with the same `name`. | [optional] 
**SignatureAllowType** | Pointer to **bool** | Signature and initials fields only: lets signers type their name and choose a script font instead of drawing, on forms and in data requests. Shared by every field with the same `name`. | [optional] 
**BarcodeSymbology** | Pointer to [**BarcodeSymbology**](BarcodeSymbology.md) | Barcode fields only: the barcode format, which the value must match. `code_128` and `pdf417` take any ASCII text; `code_39` and `code_93` take capital letters, digits, spaces and `- . $ / + %`; `code_25` takes digits; `ean_13` takes 12 digits, `ean_8` 7 digits and `upc_a` 11 digits (the check digit is added). Shared by every field with the same `name`. | [optional] 
**Hidden** | Pointer to **bool** | Leaves the field off the PDF. It stays in the template's data schema (marked `x-hidden`) and on forms, so other fields, such as combined fields, can use its value. A hidden static check mark or shape isn't printed. | [optional] 
**X** | Pointer to **float32** | The distance from the page's left edge to the field's left edge. Positions and sizes are in units where every page is 2480 wide: one unit is the page's width in points divided by 2480 (about 0.247 points on a US Letter page). | [optional] 
**Y** | Pointer to **float32** | The distance from the top edge of the field's `page` down to the field's top edge, in the units of `x`. A page is 2480 times its height divided by its width units tall: about 3209 for US Letter and 3508 for A4. | [optional] 
**Width** | Pointer to **float32** | The field's width, in the units of `x`. A check mark is sized to the field's width. | [optional] 
**Height** | Pointer to **float32** | The field's height, in the units of `x`. The template editor keeps check marks, QR codes, and square and circle shapes square, with `height` equal to `width`. | [optional] 
**Rotation** | Pointer to **float32** | Rotates the field clockwise by this many degrees, from -360 to 360 (negative turns it counterclockwise). | [optional] 
**DisplayType** | Pointer to [**FieldDisplayType**](FieldDisplayType.md) | How the value appears on the PDF: `text`, `check` (a check mark), `shape`, `qrcode`, `barcode`, or `image` (image, signature and initials fields). Check marks and shapes print only when the value meets the field's `condition`, and the field's `type` limits which display types it can use. | [optional] 
**Opacity** | Pointer to **float32** | The opacity of the field's content and background, from 0 (invisible) to 1. QR codes and barcodes always print at full opacity. | [optional] 
**DateTimeFormat** | Pointer to **NullableString** | Date fields only: how the date prints, as a strftime format such as `%m/%d/%Y`, or a calculation such as `@years_since`, `@years_until` or `@diff_years` (also months, weeks, days, hours, minutes and seconds), counted to when the PDF is generated. Empty uses the template's date format, or its date-time format with `include_time`. | [optional] 
**Typeface** | Pointer to [**NullableTypeface**](Typeface.md) | Text display only: the font, one of `Helvetica`, `Arial`, `Calibri`, `Open Sans`, `Roboto`, `Lato`, `Poppins`, `Public Sans`, `Questrial`, `Times New Roman`, `Dancing Script`, `Lovtony Script`, or the monospace `Courier`, `Inconsolata`, `Menlo`, `Roboto Mono` and `Ubuntu Mono`. Empty uses the template's default typeface. | [optional] 
**FontSize** | Pointer to **NullableInt32** | Text display only: the font size in points. Empty uses the template's default font size. | [optional] 
**CharacterSpacing** | Pointer to **float32** | Text display only: extra space between characters, in points. Empty uses the template's default. | [optional] 
**Bold** | Pointer to **bool** | Text display only: prints the text in bold. | [optional] 
**Uppercase** | Pointer to **bool** | Text display only: prints the text in capital letters. | [optional] 
**Strikethrough** | Pointer to **bool** | Text display only: draws a line through the text. | [optional] 
**KeepLeadingSpaces** | Pointer to **bool** | Text display only: keeps the spaces at the start of the text and of each line, so text can be lined up with spaces. | [optional] 
**Overflow** | Pointer to [**TextOverflow**](TextOverflow.md) | Text display only: what happens to text that doesn't fit the field. `shrink_to_fit` makes the font smaller, down to 5 points; `truncate` cuts the text off and returns the rest in the submission's `truncated_text`. | [optional] 
**Multiline** | Pointer to **bool** | Text display only: wraps the text over as many lines as fit in the field's height instead of printing one line, and forms show a multi-line input. | [optional] 
**MultilineLines** | Pointer to **NullableInt32** | Text display only, with `multiline`: spreads the text over exactly this many evenly spaced lines in the field's height (up to 300), to line up with lines printed on the form. Empty uses the font's normal line spacing. | [optional] 
**MinLength** | Pointer to **float32** | Text display only: the fewest characters the value can have, added to the template's data schema. Not used for date, address, country or number fields, or fields with an `option_list`. | [optional] 
**MaxLength** | Pointer to **NullableInt32** | Text display only: the most characters the value can have, added to the template's data schema; it also limits an `editable` field. Not used for date, address, country or number fields, or fields with an `option_list`. | [optional] 
**AutoCalculateMaxLength** | Pointer to **bool** | Text display only, with a monospace `typeface` (`Courier`, `Inconsolata`, `Menlo`, `Roboto Mono` or `Ubuntu Mono`): adds the number of characters that fit in the field to the template's data schema as its maximum length. | [optional] 
**Color** | Pointer to **NullableString** | Text display only: the text color, as a hex code such as `444444`. Empty uses the template's default text color. | [optional] 
**ColorFieldName** | Pointer to **NullableString** | Text display only: the name of a data field whose value, a hex code such as `FF0000`, sets the text color for each submission. That field is added to the template's data schema, with `color` as its default. | [optional] 
**ColorFieldRequired** | Pointer to **bool** | Makes the `color_field_name` field required. When it isn't required and a submission leaves it out, `color` is used. | [optional] 
**BackgroundColor** | Pointer to **NullableString** | Text and check display only: a color that fills the field's area behind its content, as a hex code, to cover what the page shows there. Check fields print it even when unchecked; text fields print it only when they have a value. | [optional] 
**BackgroundColorFieldName** | Pointer to **NullableString** | Text and check display only: the name of a data field whose value, a hex code, sets the background color for each submission. That field is added to the template's data schema, with `background_color` as its default. | [optional] 
**BackgroundColorFieldRequired** | Pointer to **bool** | Makes the `background_color_field_name` field required. When it isn't required and a submission leaves it out, `background_color` is used. | [optional] 
**Alignment** | Pointer to [**TextAlignment**](TextAlignment.md) | Text display only: how the text is aligned across the field: `left`, `center`, `right` or `justify` (both edges of every line of multiline text except the last of each paragraph). | [optional] 
**VAlignment** | Pointer to [**TextVerticalAlignment**](TextVerticalAlignment.md) | Text display only: how the text is aligned vertically in the field: `bottom`, `center` or `top`. | [optional] 
**Comb** | Pointer to **bool** | Text display only: splits the field into `comb_number_of_cells` boxes of equal width and prints one character, centered, in each. | [optional] 
**CombNumberOfCells** | Pointer to **int32** | With `comb`: how many boxes the field has. The value's maximum length in the template's data schema is this plus `comb_value_offset`. | [optional] 
**CombValueOffset** | Pointer to **int32** | With `comb`: the position, from 0, of the first character of the value that this field prints. For an SSN printed in boxes of 3, 2 and 4 digits, three fields named `ssn` use offsets 0, 3 and 5. | [optional] 
**ShapeType** | Pointer to [**ShapeType**](ShapeType.md) | Shape display only: `square`, `rectangle`, `circle`, `ellipse`, `triangle` or `polygon`. A square or circle uses the smaller of `width` and `height`. | [optional] 
**ShapeSides** | Pointer to **int32** | Shape display only, with `shape_type` `polygon`: the number of sides, 3 to 12. | [optional] 
**ShapeBorderWidth** | Pointer to **int32** | Shape display only: the width of the shape's border, in points. | [optional] 
**ShapeBorderColor** | Pointer to **string** | Shape display only: the shape's border color, as a hex code. Null draws no border. | [optional] 
**ShapeBorderColorFieldName** | Pointer to **NullableString** | Shape display only: the name of a data field whose value, a hex code, sets the border color for each submission. That field is added to the template's data schema, with `shape_border_color` as its default. | [optional] 
**ShapeBorderColorFieldRequired** | Pointer to **bool** | Makes the `shape_border_color_field_name` field required. When it isn't required and a submission leaves it out, `shape_border_color` is used. | [optional] 
**ShapeFillColor** | Pointer to **NullableString** | Shape display only: the color that fills the shape, as a hex code. Null leaves it unfilled. | [optional] 
**ShapeFillColorFieldName** | Pointer to **NullableString** | Shape display only: the name of a data field whose value, a hex code, sets the fill color for each submission. That field is added to the template's data schema, with `shape_fill_color` as its default. | [optional] 
**ShapeFillColorFieldRequired** | Pointer to **bool** | Makes the `shape_fill_color_field_name` field required. When it isn't required and a submission leaves it out, `shape_fill_color` is used. | [optional] 
**CheckCharacter** | Pointer to [**CheckCharacter**](CheckCharacter.md) | Check display only: the character printed, as an HTML entity: `&amp;#10003;` (✓), `&amp;#10004;` (✔, the default), `&amp;#10006;` (✖), `&amp;#10007;` (✗) or `&amp;#10008;` (✘). | [optional] 
**CheckColor** | Pointer to **NullableString** | Check display only: the check mark's color, as a hex code. Empty uses the template's default check color. | [optional] 
**CheckColorFieldName** | Pointer to **NullableString** | Check display only: the name of a data field whose value, a hex code, sets the check mark's color for each submission. That field is added to the template's data schema, with `check_color` as its default. | [optional] 
**CheckColorFieldRequired** | Pointer to **bool** | Makes the `check_color_field_name` field required. When it isn't required and a submission leaves it out, `check_color` is used. | [optional] 
**QrcodeColor** | Pointer to **NullableString** | QR code display only: the QR code's color, as a hex code, printed on a white background. Empty prints it in black. | [optional] 
**QrcodeColorFieldName** | Pointer to **NullableString** | QR code display only: the name of a data field whose value, a hex code, sets the QR code's color for each submission. That field is added to the template's data schema, with `qrcode_color` as its default. | [optional] 
**QrcodeColorFieldRequired** | Pointer to **bool** | Makes the `qrcode_color_field_name` field required. When it isn't required and a submission leaves it out, `qrcode_color` is used. | [optional] 
**TrueText** | Pointer to **NullableString** | Boolean fields with text display: the text printed when the value is `true`. Empty uses the template's `true_text`; when neither is set, nothing prints. | [optional] 
**FalseText** | Pointer to **NullableString** | Boolean fields with text display: the text printed when the value is `false`. Empty uses the template's `false_text`; when neither is set, nothing prints. | [optional] 
**Condition** | Pointer to **NullableString** | Check and shape display only: the value a string or number field's value is compared with, by `string_condition_type` or `number_condition_type`, to decide whether the check mark or shape prints. Boolean fields print when `true` (see `invert_boolean_condition`), and other types print when they have a value. | [optional] 
**StringConditionType** | Pointer to [**StringConditionType**](StringConditionType.md) | String fields with check or shape display: how the value is compared with `condition`: `equals`, `contains`, `starts_with`, `ends_with`, or `regex`, where `condition` is a regular expression the value must match. | [optional] 
**NumberConditionType** | Pointer to [**NumberConditionType**](NumberConditionType.md) | Number fields with check or shape display: how the value is compared with `condition`: `equals`, `gt` (greater than), `gte` (greater than or equal to), `lt` (less than), `lte` (less than or equal to), or `range`, between `number_condition_range_min` and `number_condition_range_max`. | [optional] 
**NumberConditionRangeMin** | Pointer to **NullableFloat32** | With `number_condition_type` `range`: the lower end of the range, included unless `number_condition_range_exclusive_min` is set. | [optional] 
**NumberConditionRangeMax** | Pointer to **NullableFloat32** | With `number_condition_type` `range`: the upper end of the range, included unless `number_condition_range_exclusive_max` is set. | [optional] 
**NumberConditionRangeExclusiveMin** | Pointer to **bool** | With `number_condition_type` `range`: leaves `number_condition_range_min` itself out of the range. | [optional] 
**NumberConditionRangeExclusiveMax** | Pointer to **bool** | With `number_condition_type` `range`: leaves `number_condition_range_max` itself out of the range. | [optional] 
**InvertBooleanCondition** | Pointer to **bool** | Boolean fields with check or shape display: prints the check mark or shape when the value is `false` or missing, instead of `true`. | [optional] 
**Currency** | Pointer to **bool** | Number fields only: forms show a currency input, and the template's data schema gives the field `format: currency`. It doesn't change how the number prints. | [optional] 
**DecimalPlaces** | Pointer to **NullableInt32** | Number fields with text display: prints the number with this many decimal places, from 0 (rounds to a whole number) to 99. Empty prints the number as it was sent. | [optional] 
**ImageScaleType** | Pointer to [**ImageScaleType**](ImageScaleType.md) | Image fields only: how the image is resized to the field. `fit` keeps its proportions and fits it inside the field, which can leave empty space; `fill` keeps its proportions and crops it to fill the field; `stretch` stretches it to the field's size. Signatures and initials always fit. | [optional] 
**ImageGravity** | Pointer to [**ImageGravity**](ImageGravity.md) | Image, signature and initials fields: where the image sits in the field with `fit`, or which part is kept with `fill`: `NorthWest` (top left), `North`, `NorthEast`, `West`, `Center`, `East`, `SouthWest`, `South` or `SouthEast` (bottom right). It also aligns typed signatures, and isn't used with `stretch`. | [optional] 
**Dependencies** | Pointer to [**NullableFieldDependencyInput**](FieldDependencyInput.md) | The field's own conditional logic rule: when other fields' values meet its conditions, the field is shown, made required, or both. Null when it has none. The template's rules (`field_dependencies`) can apply to the field too. Shared by every field with the same `name`. | [optional] 
**Id** | Pointer to **int32** | The field's ID. Required when updating a field. When adding one, leave it out for the next free ID. | [optional] 

## Methods

### NewFieldInput

`func NewFieldInput() *FieldInput`

NewFieldInput instantiates a new FieldInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldInputWithDefaults

`func NewFieldInputWithDefaults() *FieldInput`

NewFieldInputWithDefaults instantiates a new FieldInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FieldInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FieldInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FieldInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FieldInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPage

`func (o *FieldInput) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *FieldInput) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *FieldInput) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *FieldInput) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetType

`func (o *FieldInput) GetType() FieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldInput) GetTypeOk() (*FieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldInput) SetType(v FieldType)`

SetType sets Type field to given value.

### HasType

`func (o *FieldInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatic

`func (o *FieldInput) GetStatic() bool`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *FieldInput) GetStaticOk() (*bool, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *FieldInput) SetStatic(v bool)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *FieldInput) HasStatic() bool`

HasStatic returns a boolean if a field has been set.

### GetEditable

`func (o *FieldInput) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *FieldInput) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *FieldInput) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *FieldInput) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetDefault

`func (o *FieldInput) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FieldInput) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FieldInput) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FieldInput) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *FieldInput) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *FieldInput) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetTitle

`func (o *FieldInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FieldInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FieldInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FieldInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *FieldInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlaceholder

`func (o *FieldInput) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *FieldInput) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *FieldInput) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *FieldInput) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetMetadata

`func (o *FieldInput) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FieldInput) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FieldInput) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FieldInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRequired

`func (o *FieldInput) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldInput) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldInput) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FieldInput) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetOptionList

`func (o *FieldInput) GetOptionList() string`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *FieldInput) GetOptionListOk() (*string, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *FieldInput) SetOptionList(v string)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *FieldInput) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### SetOptionListNil

`func (o *FieldInput) SetOptionListNil(b bool)`

 SetOptionListNil sets the value for OptionList to be an explicit nil

### UnsetOptionList
`func (o *FieldInput) UnsetOptionList()`

UnsetOptionList ensures that no value is present for OptionList, not even an explicit nil
### GetCombinedFieldNames

`func (o *FieldInput) GetCombinedFieldNames() string`

GetCombinedFieldNames returns the CombinedFieldNames field if non-nil, zero value otherwise.

### GetCombinedFieldNamesOk

`func (o *FieldInput) GetCombinedFieldNamesOk() (*string, bool)`

GetCombinedFieldNamesOk returns a tuple with the CombinedFieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldNames

`func (o *FieldInput) SetCombinedFieldNames(v string)`

SetCombinedFieldNames sets CombinedFieldNames field to given value.

### HasCombinedFieldNames

`func (o *FieldInput) HasCombinedFieldNames() bool`

HasCombinedFieldNames returns a boolean if a field has been set.

### GetCombinedFieldType

`func (o *FieldInput) GetCombinedFieldType() CombinedFieldType`

GetCombinedFieldType returns the CombinedFieldType field if non-nil, zero value otherwise.

### GetCombinedFieldTypeOk

`func (o *FieldInput) GetCombinedFieldTypeOk() (*CombinedFieldType, bool)`

GetCombinedFieldTypeOk returns a tuple with the CombinedFieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldType

`func (o *FieldInput) SetCombinedFieldType(v CombinedFieldType)`

SetCombinedFieldType sets CombinedFieldType field to given value.

### HasCombinedFieldType

`func (o *FieldInput) HasCombinedFieldType() bool`

HasCombinedFieldType returns a boolean if a field has been set.

### GetCombinedFieldSeparator

`func (o *FieldInput) GetCombinedFieldSeparator() string`

GetCombinedFieldSeparator returns the CombinedFieldSeparator field if non-nil, zero value otherwise.

### GetCombinedFieldSeparatorOk

`func (o *FieldInput) GetCombinedFieldSeparatorOk() (*string, bool)`

GetCombinedFieldSeparatorOk returns a tuple with the CombinedFieldSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldSeparator

`func (o *FieldInput) SetCombinedFieldSeparator(v string)`

SetCombinedFieldSeparator sets CombinedFieldSeparator field to given value.

### HasCombinedFieldSeparator

`func (o *FieldInput) HasCombinedFieldSeparator() bool`

HasCombinedFieldSeparator returns a boolean if a field has been set.

### GetCombinedFieldFormat

`func (o *FieldInput) GetCombinedFieldFormat() string`

GetCombinedFieldFormat returns the CombinedFieldFormat field if non-nil, zero value otherwise.

### GetCombinedFieldFormatOk

`func (o *FieldInput) GetCombinedFieldFormatOk() (*string, bool)`

GetCombinedFieldFormatOk returns a tuple with the CombinedFieldFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldFormat

`func (o *FieldInput) SetCombinedFieldFormat(v string)`

SetCombinedFieldFormat sets CombinedFieldFormat field to given value.

### HasCombinedFieldFormat

`func (o *FieldInput) HasCombinedFieldFormat() bool`

HasCombinedFieldFormat returns a boolean if a field has been set.

### GetInteger

`func (o *FieldInput) GetInteger() bool`

GetInteger returns the Integer field if non-nil, zero value otherwise.

### GetIntegerOk

`func (o *FieldInput) GetIntegerOk() (*bool, bool)`

GetIntegerOk returns a tuple with the Integer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteger

`func (o *FieldInput) SetInteger(v bool)`

SetInteger sets Integer field to given value.

### HasInteger

`func (o *FieldInput) HasInteger() bool`

HasInteger returns a boolean if a field has been set.

### GetMinimum

`func (o *FieldInput) GetMinimum() float32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *FieldInput) GetMinimumOk() (*float32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *FieldInput) SetMinimum(v float32)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *FieldInput) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### SetMinimumNil

`func (o *FieldInput) SetMinimumNil(b bool)`

 SetMinimumNil sets the value for Minimum to be an explicit nil

### UnsetMinimum
`func (o *FieldInput) UnsetMinimum()`

UnsetMinimum ensures that no value is present for Minimum, not even an explicit nil
### GetMaximum

`func (o *FieldInput) GetMaximum() float32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *FieldInput) GetMaximumOk() (*float32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *FieldInput) SetMaximum(v float32)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *FieldInput) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### SetMaximumNil

`func (o *FieldInput) SetMaximumNil(b bool)`

 SetMaximumNil sets the value for Maximum to be an explicit nil

### UnsetMaximum
`func (o *FieldInput) UnsetMaximum()`

UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil
### GetExclusiveMinimum

`func (o *FieldInput) GetExclusiveMinimum() bool`

GetExclusiveMinimum returns the ExclusiveMinimum field if non-nil, zero value otherwise.

### GetExclusiveMinimumOk

`func (o *FieldInput) GetExclusiveMinimumOk() (*bool, bool)`

GetExclusiveMinimumOk returns a tuple with the ExclusiveMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMinimum

`func (o *FieldInput) SetExclusiveMinimum(v bool)`

SetExclusiveMinimum sets ExclusiveMinimum field to given value.

### HasExclusiveMinimum

`func (o *FieldInput) HasExclusiveMinimum() bool`

HasExclusiveMinimum returns a boolean if a field has been set.

### GetExclusiveMaximum

`func (o *FieldInput) GetExclusiveMaximum() bool`

GetExclusiveMaximum returns the ExclusiveMaximum field if non-nil, zero value otherwise.

### GetExclusiveMaximumOk

`func (o *FieldInput) GetExclusiveMaximumOk() (*bool, bool)`

GetExclusiveMaximumOk returns a tuple with the ExclusiveMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMaximum

`func (o *FieldInput) SetExclusiveMaximum(v bool)`

SetExclusiveMaximum sets ExclusiveMaximum field to given value.

### HasExclusiveMaximum

`func (o *FieldInput) HasExclusiveMaximum() bool`

HasExclusiveMaximum returns a boolean if a field has been set.

### GetIncludeTime

`func (o *FieldInput) GetIncludeTime() bool`

GetIncludeTime returns the IncludeTime field if non-nil, zero value otherwise.

### GetIncludeTimeOk

`func (o *FieldInput) GetIncludeTimeOk() (*bool, bool)`

GetIncludeTimeOk returns a tuple with the IncludeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTime

`func (o *FieldInput) SetIncludeTime(v bool)`

SetIncludeTime sets IncludeTime field to given value.

### HasIncludeTime

`func (o *FieldInput) HasIncludeTime() bool`

HasIncludeTime returns a boolean if a field has been set.

### SetIncludeTimeNil

`func (o *FieldInput) SetIncludeTimeNil(b bool)`

 SetIncludeTimeNil sets the value for IncludeTime to be an explicit nil

### UnsetIncludeTime
`func (o *FieldInput) UnsetIncludeTime()`

UnsetIncludeTime ensures that no value is present for IncludeTime, not even an explicit nil
### GetSignatureAllowDraw

`func (o *FieldInput) GetSignatureAllowDraw() bool`

GetSignatureAllowDraw returns the SignatureAllowDraw field if non-nil, zero value otherwise.

### GetSignatureAllowDrawOk

`func (o *FieldInput) GetSignatureAllowDrawOk() (*bool, bool)`

GetSignatureAllowDrawOk returns a tuple with the SignatureAllowDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAllowDraw

`func (o *FieldInput) SetSignatureAllowDraw(v bool)`

SetSignatureAllowDraw sets SignatureAllowDraw field to given value.

### HasSignatureAllowDraw

`func (o *FieldInput) HasSignatureAllowDraw() bool`

HasSignatureAllowDraw returns a boolean if a field has been set.

### GetSignatureAllowType

`func (o *FieldInput) GetSignatureAllowType() bool`

GetSignatureAllowType returns the SignatureAllowType field if non-nil, zero value otherwise.

### GetSignatureAllowTypeOk

`func (o *FieldInput) GetSignatureAllowTypeOk() (*bool, bool)`

GetSignatureAllowTypeOk returns a tuple with the SignatureAllowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAllowType

`func (o *FieldInput) SetSignatureAllowType(v bool)`

SetSignatureAllowType sets SignatureAllowType field to given value.

### HasSignatureAllowType

`func (o *FieldInput) HasSignatureAllowType() bool`

HasSignatureAllowType returns a boolean if a field has been set.

### GetBarcodeSymbology

`func (o *FieldInput) GetBarcodeSymbology() BarcodeSymbology`

GetBarcodeSymbology returns the BarcodeSymbology field if non-nil, zero value otherwise.

### GetBarcodeSymbologyOk

`func (o *FieldInput) GetBarcodeSymbologyOk() (*BarcodeSymbology, bool)`

GetBarcodeSymbologyOk returns a tuple with the BarcodeSymbology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeSymbology

`func (o *FieldInput) SetBarcodeSymbology(v BarcodeSymbology)`

SetBarcodeSymbology sets BarcodeSymbology field to given value.

### HasBarcodeSymbology

`func (o *FieldInput) HasBarcodeSymbology() bool`

HasBarcodeSymbology returns a boolean if a field has been set.

### GetHidden

`func (o *FieldInput) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *FieldInput) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *FieldInput) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *FieldInput) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetX

`func (o *FieldInput) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *FieldInput) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *FieldInput) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *FieldInput) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *FieldInput) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *FieldInput) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *FieldInput) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *FieldInput) HasY() bool`

HasY returns a boolean if a field has been set.

### GetWidth

`func (o *FieldInput) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *FieldInput) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *FieldInput) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *FieldInput) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *FieldInput) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *FieldInput) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *FieldInput) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *FieldInput) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetRotation

`func (o *FieldInput) GetRotation() float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *FieldInput) GetRotationOk() (*float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *FieldInput) SetRotation(v float32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *FieldInput) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetDisplayType

`func (o *FieldInput) GetDisplayType() FieldDisplayType`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *FieldInput) GetDisplayTypeOk() (*FieldDisplayType, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *FieldInput) SetDisplayType(v FieldDisplayType)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *FieldInput) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetOpacity

`func (o *FieldInput) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *FieldInput) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *FieldInput) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *FieldInput) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *FieldInput) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *FieldInput) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *FieldInput) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *FieldInput) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### SetDateTimeFormatNil

`func (o *FieldInput) SetDateTimeFormatNil(b bool)`

 SetDateTimeFormatNil sets the value for DateTimeFormat to be an explicit nil

### UnsetDateTimeFormat
`func (o *FieldInput) UnsetDateTimeFormat()`

UnsetDateTimeFormat ensures that no value is present for DateTimeFormat, not even an explicit nil
### GetTypeface

`func (o *FieldInput) GetTypeface() Typeface`

GetTypeface returns the Typeface field if non-nil, zero value otherwise.

### GetTypefaceOk

`func (o *FieldInput) GetTypefaceOk() (*Typeface, bool)`

GetTypefaceOk returns a tuple with the Typeface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeface

`func (o *FieldInput) SetTypeface(v Typeface)`

SetTypeface sets Typeface field to given value.

### HasTypeface

`func (o *FieldInput) HasTypeface() bool`

HasTypeface returns a boolean if a field has been set.

### SetTypefaceNil

`func (o *FieldInput) SetTypefaceNil(b bool)`

 SetTypefaceNil sets the value for Typeface to be an explicit nil

### UnsetTypeface
`func (o *FieldInput) UnsetTypeface()`

UnsetTypeface ensures that no value is present for Typeface, not even an explicit nil
### GetFontSize

`func (o *FieldInput) GetFontSize() int32`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *FieldInput) GetFontSizeOk() (*int32, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *FieldInput) SetFontSize(v int32)`

SetFontSize sets FontSize field to given value.

### HasFontSize

`func (o *FieldInput) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.

### SetFontSizeNil

`func (o *FieldInput) SetFontSizeNil(b bool)`

 SetFontSizeNil sets the value for FontSize to be an explicit nil

### UnsetFontSize
`func (o *FieldInput) UnsetFontSize()`

UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil
### GetCharacterSpacing

`func (o *FieldInput) GetCharacterSpacing() float32`

GetCharacterSpacing returns the CharacterSpacing field if non-nil, zero value otherwise.

### GetCharacterSpacingOk

`func (o *FieldInput) GetCharacterSpacingOk() (*float32, bool)`

GetCharacterSpacingOk returns a tuple with the CharacterSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterSpacing

`func (o *FieldInput) SetCharacterSpacing(v float32)`

SetCharacterSpacing sets CharacterSpacing field to given value.

### HasCharacterSpacing

`func (o *FieldInput) HasCharacterSpacing() bool`

HasCharacterSpacing returns a boolean if a field has been set.

### GetBold

`func (o *FieldInput) GetBold() bool`

GetBold returns the Bold field if non-nil, zero value otherwise.

### GetBoldOk

`func (o *FieldInput) GetBoldOk() (*bool, bool)`

GetBoldOk returns a tuple with the Bold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBold

`func (o *FieldInput) SetBold(v bool)`

SetBold sets Bold field to given value.

### HasBold

`func (o *FieldInput) HasBold() bool`

HasBold returns a boolean if a field has been set.

### GetUppercase

`func (o *FieldInput) GetUppercase() bool`

GetUppercase returns the Uppercase field if non-nil, zero value otherwise.

### GetUppercaseOk

`func (o *FieldInput) GetUppercaseOk() (*bool, bool)`

GetUppercaseOk returns a tuple with the Uppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUppercase

`func (o *FieldInput) SetUppercase(v bool)`

SetUppercase sets Uppercase field to given value.

### HasUppercase

`func (o *FieldInput) HasUppercase() bool`

HasUppercase returns a boolean if a field has been set.

### GetStrikethrough

`func (o *FieldInput) GetStrikethrough() bool`

GetStrikethrough returns the Strikethrough field if non-nil, zero value otherwise.

### GetStrikethroughOk

`func (o *FieldInput) GetStrikethroughOk() (*bool, bool)`

GetStrikethroughOk returns a tuple with the Strikethrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikethrough

`func (o *FieldInput) SetStrikethrough(v bool)`

SetStrikethrough sets Strikethrough field to given value.

### HasStrikethrough

`func (o *FieldInput) HasStrikethrough() bool`

HasStrikethrough returns a boolean if a field has been set.

### GetKeepLeadingSpaces

`func (o *FieldInput) GetKeepLeadingSpaces() bool`

GetKeepLeadingSpaces returns the KeepLeadingSpaces field if non-nil, zero value otherwise.

### GetKeepLeadingSpacesOk

`func (o *FieldInput) GetKeepLeadingSpacesOk() (*bool, bool)`

GetKeepLeadingSpacesOk returns a tuple with the KeepLeadingSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepLeadingSpaces

`func (o *FieldInput) SetKeepLeadingSpaces(v bool)`

SetKeepLeadingSpaces sets KeepLeadingSpaces field to given value.

### HasKeepLeadingSpaces

`func (o *FieldInput) HasKeepLeadingSpaces() bool`

HasKeepLeadingSpaces returns a boolean if a field has been set.

### GetOverflow

`func (o *FieldInput) GetOverflow() TextOverflow`

GetOverflow returns the Overflow field if non-nil, zero value otherwise.

### GetOverflowOk

`func (o *FieldInput) GetOverflowOk() (*TextOverflow, bool)`

GetOverflowOk returns a tuple with the Overflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflow

`func (o *FieldInput) SetOverflow(v TextOverflow)`

SetOverflow sets Overflow field to given value.

### HasOverflow

`func (o *FieldInput) HasOverflow() bool`

HasOverflow returns a boolean if a field has been set.

### GetMultiline

`func (o *FieldInput) GetMultiline() bool`

GetMultiline returns the Multiline field if non-nil, zero value otherwise.

### GetMultilineOk

`func (o *FieldInput) GetMultilineOk() (*bool, bool)`

GetMultilineOk returns a tuple with the Multiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiline

`func (o *FieldInput) SetMultiline(v bool)`

SetMultiline sets Multiline field to given value.

### HasMultiline

`func (o *FieldInput) HasMultiline() bool`

HasMultiline returns a boolean if a field has been set.

### GetMultilineLines

`func (o *FieldInput) GetMultilineLines() int32`

GetMultilineLines returns the MultilineLines field if non-nil, zero value otherwise.

### GetMultilineLinesOk

`func (o *FieldInput) GetMultilineLinesOk() (*int32, bool)`

GetMultilineLinesOk returns a tuple with the MultilineLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultilineLines

`func (o *FieldInput) SetMultilineLines(v int32)`

SetMultilineLines sets MultilineLines field to given value.

### HasMultilineLines

`func (o *FieldInput) HasMultilineLines() bool`

HasMultilineLines returns a boolean if a field has been set.

### SetMultilineLinesNil

`func (o *FieldInput) SetMultilineLinesNil(b bool)`

 SetMultilineLinesNil sets the value for MultilineLines to be an explicit nil

### UnsetMultilineLines
`func (o *FieldInput) UnsetMultilineLines()`

UnsetMultilineLines ensures that no value is present for MultilineLines, not even an explicit nil
### GetMinLength

`func (o *FieldInput) GetMinLength() float32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *FieldInput) GetMinLengthOk() (*float32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *FieldInput) SetMinLength(v float32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *FieldInput) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *FieldInput) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *FieldInput) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *FieldInput) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *FieldInput) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *FieldInput) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *FieldInput) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetAutoCalculateMaxLength

`func (o *FieldInput) GetAutoCalculateMaxLength() bool`

GetAutoCalculateMaxLength returns the AutoCalculateMaxLength field if non-nil, zero value otherwise.

### GetAutoCalculateMaxLengthOk

`func (o *FieldInput) GetAutoCalculateMaxLengthOk() (*bool, bool)`

GetAutoCalculateMaxLengthOk returns a tuple with the AutoCalculateMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCalculateMaxLength

`func (o *FieldInput) SetAutoCalculateMaxLength(v bool)`

SetAutoCalculateMaxLength sets AutoCalculateMaxLength field to given value.

### HasAutoCalculateMaxLength

`func (o *FieldInput) HasAutoCalculateMaxLength() bool`

HasAutoCalculateMaxLength returns a boolean if a field has been set.

### GetColor

`func (o *FieldInput) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *FieldInput) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *FieldInput) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *FieldInput) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *FieldInput) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *FieldInput) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetColorFieldName

`func (o *FieldInput) GetColorFieldName() string`

GetColorFieldName returns the ColorFieldName field if non-nil, zero value otherwise.

### GetColorFieldNameOk

`func (o *FieldInput) GetColorFieldNameOk() (*string, bool)`

GetColorFieldNameOk returns a tuple with the ColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorFieldName

`func (o *FieldInput) SetColorFieldName(v string)`

SetColorFieldName sets ColorFieldName field to given value.

### HasColorFieldName

`func (o *FieldInput) HasColorFieldName() bool`

HasColorFieldName returns a boolean if a field has been set.

### SetColorFieldNameNil

`func (o *FieldInput) SetColorFieldNameNil(b bool)`

 SetColorFieldNameNil sets the value for ColorFieldName to be an explicit nil

### UnsetColorFieldName
`func (o *FieldInput) UnsetColorFieldName()`

UnsetColorFieldName ensures that no value is present for ColorFieldName, not even an explicit nil
### GetColorFieldRequired

`func (o *FieldInput) GetColorFieldRequired() bool`

GetColorFieldRequired returns the ColorFieldRequired field if non-nil, zero value otherwise.

### GetColorFieldRequiredOk

`func (o *FieldInput) GetColorFieldRequiredOk() (*bool, bool)`

GetColorFieldRequiredOk returns a tuple with the ColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorFieldRequired

`func (o *FieldInput) SetColorFieldRequired(v bool)`

SetColorFieldRequired sets ColorFieldRequired field to given value.

### HasColorFieldRequired

`func (o *FieldInput) HasColorFieldRequired() bool`

HasColorFieldRequired returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *FieldInput) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *FieldInput) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *FieldInput) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *FieldInput) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *FieldInput) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *FieldInput) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetBackgroundColorFieldName

`func (o *FieldInput) GetBackgroundColorFieldName() string`

GetBackgroundColorFieldName returns the BackgroundColorFieldName field if non-nil, zero value otherwise.

### GetBackgroundColorFieldNameOk

`func (o *FieldInput) GetBackgroundColorFieldNameOk() (*string, bool)`

GetBackgroundColorFieldNameOk returns a tuple with the BackgroundColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColorFieldName

`func (o *FieldInput) SetBackgroundColorFieldName(v string)`

SetBackgroundColorFieldName sets BackgroundColorFieldName field to given value.

### HasBackgroundColorFieldName

`func (o *FieldInput) HasBackgroundColorFieldName() bool`

HasBackgroundColorFieldName returns a boolean if a field has been set.

### SetBackgroundColorFieldNameNil

`func (o *FieldInput) SetBackgroundColorFieldNameNil(b bool)`

 SetBackgroundColorFieldNameNil sets the value for BackgroundColorFieldName to be an explicit nil

### UnsetBackgroundColorFieldName
`func (o *FieldInput) UnsetBackgroundColorFieldName()`

UnsetBackgroundColorFieldName ensures that no value is present for BackgroundColorFieldName, not even an explicit nil
### GetBackgroundColorFieldRequired

`func (o *FieldInput) GetBackgroundColorFieldRequired() bool`

GetBackgroundColorFieldRequired returns the BackgroundColorFieldRequired field if non-nil, zero value otherwise.

### GetBackgroundColorFieldRequiredOk

`func (o *FieldInput) GetBackgroundColorFieldRequiredOk() (*bool, bool)`

GetBackgroundColorFieldRequiredOk returns a tuple with the BackgroundColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColorFieldRequired

`func (o *FieldInput) SetBackgroundColorFieldRequired(v bool)`

SetBackgroundColorFieldRequired sets BackgroundColorFieldRequired field to given value.

### HasBackgroundColorFieldRequired

`func (o *FieldInput) HasBackgroundColorFieldRequired() bool`

HasBackgroundColorFieldRequired returns a boolean if a field has been set.

### GetAlignment

`func (o *FieldInput) GetAlignment() TextAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *FieldInput) GetAlignmentOk() (*TextAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *FieldInput) SetAlignment(v TextAlignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *FieldInput) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetVAlignment

`func (o *FieldInput) GetVAlignment() TextVerticalAlignment`

GetVAlignment returns the VAlignment field if non-nil, zero value otherwise.

### GetVAlignmentOk

`func (o *FieldInput) GetVAlignmentOk() (*TextVerticalAlignment, bool)`

GetVAlignmentOk returns a tuple with the VAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAlignment

`func (o *FieldInput) SetVAlignment(v TextVerticalAlignment)`

SetVAlignment sets VAlignment field to given value.

### HasVAlignment

`func (o *FieldInput) HasVAlignment() bool`

HasVAlignment returns a boolean if a field has been set.

### GetComb

`func (o *FieldInput) GetComb() bool`

GetComb returns the Comb field if non-nil, zero value otherwise.

### GetCombOk

`func (o *FieldInput) GetCombOk() (*bool, bool)`

GetCombOk returns a tuple with the Comb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComb

`func (o *FieldInput) SetComb(v bool)`

SetComb sets Comb field to given value.

### HasComb

`func (o *FieldInput) HasComb() bool`

HasComb returns a boolean if a field has been set.

### GetCombNumberOfCells

`func (o *FieldInput) GetCombNumberOfCells() int32`

GetCombNumberOfCells returns the CombNumberOfCells field if non-nil, zero value otherwise.

### GetCombNumberOfCellsOk

`func (o *FieldInput) GetCombNumberOfCellsOk() (*int32, bool)`

GetCombNumberOfCellsOk returns a tuple with the CombNumberOfCells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombNumberOfCells

`func (o *FieldInput) SetCombNumberOfCells(v int32)`

SetCombNumberOfCells sets CombNumberOfCells field to given value.

### HasCombNumberOfCells

`func (o *FieldInput) HasCombNumberOfCells() bool`

HasCombNumberOfCells returns a boolean if a field has been set.

### GetCombValueOffset

`func (o *FieldInput) GetCombValueOffset() int32`

GetCombValueOffset returns the CombValueOffset field if non-nil, zero value otherwise.

### GetCombValueOffsetOk

`func (o *FieldInput) GetCombValueOffsetOk() (*int32, bool)`

GetCombValueOffsetOk returns a tuple with the CombValueOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombValueOffset

`func (o *FieldInput) SetCombValueOffset(v int32)`

SetCombValueOffset sets CombValueOffset field to given value.

### HasCombValueOffset

`func (o *FieldInput) HasCombValueOffset() bool`

HasCombValueOffset returns a boolean if a field has been set.

### GetShapeType

`func (o *FieldInput) GetShapeType() ShapeType`

GetShapeType returns the ShapeType field if non-nil, zero value otherwise.

### GetShapeTypeOk

`func (o *FieldInput) GetShapeTypeOk() (*ShapeType, bool)`

GetShapeTypeOk returns a tuple with the ShapeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeType

`func (o *FieldInput) SetShapeType(v ShapeType)`

SetShapeType sets ShapeType field to given value.

### HasShapeType

`func (o *FieldInput) HasShapeType() bool`

HasShapeType returns a boolean if a field has been set.

### GetShapeSides

`func (o *FieldInput) GetShapeSides() int32`

GetShapeSides returns the ShapeSides field if non-nil, zero value otherwise.

### GetShapeSidesOk

`func (o *FieldInput) GetShapeSidesOk() (*int32, bool)`

GetShapeSidesOk returns a tuple with the ShapeSides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeSides

`func (o *FieldInput) SetShapeSides(v int32)`

SetShapeSides sets ShapeSides field to given value.

### HasShapeSides

`func (o *FieldInput) HasShapeSides() bool`

HasShapeSides returns a boolean if a field has been set.

### GetShapeBorderWidth

`func (o *FieldInput) GetShapeBorderWidth() int32`

GetShapeBorderWidth returns the ShapeBorderWidth field if non-nil, zero value otherwise.

### GetShapeBorderWidthOk

`func (o *FieldInput) GetShapeBorderWidthOk() (*int32, bool)`

GetShapeBorderWidthOk returns a tuple with the ShapeBorderWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderWidth

`func (o *FieldInput) SetShapeBorderWidth(v int32)`

SetShapeBorderWidth sets ShapeBorderWidth field to given value.

### HasShapeBorderWidth

`func (o *FieldInput) HasShapeBorderWidth() bool`

HasShapeBorderWidth returns a boolean if a field has been set.

### GetShapeBorderColor

`func (o *FieldInput) GetShapeBorderColor() string`

GetShapeBorderColor returns the ShapeBorderColor field if non-nil, zero value otherwise.

### GetShapeBorderColorOk

`func (o *FieldInput) GetShapeBorderColorOk() (*string, bool)`

GetShapeBorderColorOk returns a tuple with the ShapeBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderColor

`func (o *FieldInput) SetShapeBorderColor(v string)`

SetShapeBorderColor sets ShapeBorderColor field to given value.

### HasShapeBorderColor

`func (o *FieldInput) HasShapeBorderColor() bool`

HasShapeBorderColor returns a boolean if a field has been set.

### GetShapeBorderColorFieldName

`func (o *FieldInput) GetShapeBorderColorFieldName() string`

GetShapeBorderColorFieldName returns the ShapeBorderColorFieldName field if non-nil, zero value otherwise.

### GetShapeBorderColorFieldNameOk

`func (o *FieldInput) GetShapeBorderColorFieldNameOk() (*string, bool)`

GetShapeBorderColorFieldNameOk returns a tuple with the ShapeBorderColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderColorFieldName

`func (o *FieldInput) SetShapeBorderColorFieldName(v string)`

SetShapeBorderColorFieldName sets ShapeBorderColorFieldName field to given value.

### HasShapeBorderColorFieldName

`func (o *FieldInput) HasShapeBorderColorFieldName() bool`

HasShapeBorderColorFieldName returns a boolean if a field has been set.

### SetShapeBorderColorFieldNameNil

`func (o *FieldInput) SetShapeBorderColorFieldNameNil(b bool)`

 SetShapeBorderColorFieldNameNil sets the value for ShapeBorderColorFieldName to be an explicit nil

### UnsetShapeBorderColorFieldName
`func (o *FieldInput) UnsetShapeBorderColorFieldName()`

UnsetShapeBorderColorFieldName ensures that no value is present for ShapeBorderColorFieldName, not even an explicit nil
### GetShapeBorderColorFieldRequired

`func (o *FieldInput) GetShapeBorderColorFieldRequired() bool`

GetShapeBorderColorFieldRequired returns the ShapeBorderColorFieldRequired field if non-nil, zero value otherwise.

### GetShapeBorderColorFieldRequiredOk

`func (o *FieldInput) GetShapeBorderColorFieldRequiredOk() (*bool, bool)`

GetShapeBorderColorFieldRequiredOk returns a tuple with the ShapeBorderColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderColorFieldRequired

`func (o *FieldInput) SetShapeBorderColorFieldRequired(v bool)`

SetShapeBorderColorFieldRequired sets ShapeBorderColorFieldRequired field to given value.

### HasShapeBorderColorFieldRequired

`func (o *FieldInput) HasShapeBorderColorFieldRequired() bool`

HasShapeBorderColorFieldRequired returns a boolean if a field has been set.

### GetShapeFillColor

`func (o *FieldInput) GetShapeFillColor() string`

GetShapeFillColor returns the ShapeFillColor field if non-nil, zero value otherwise.

### GetShapeFillColorOk

`func (o *FieldInput) GetShapeFillColorOk() (*string, bool)`

GetShapeFillColorOk returns a tuple with the ShapeFillColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeFillColor

`func (o *FieldInput) SetShapeFillColor(v string)`

SetShapeFillColor sets ShapeFillColor field to given value.

### HasShapeFillColor

`func (o *FieldInput) HasShapeFillColor() bool`

HasShapeFillColor returns a boolean if a field has been set.

### SetShapeFillColorNil

`func (o *FieldInput) SetShapeFillColorNil(b bool)`

 SetShapeFillColorNil sets the value for ShapeFillColor to be an explicit nil

### UnsetShapeFillColor
`func (o *FieldInput) UnsetShapeFillColor()`

UnsetShapeFillColor ensures that no value is present for ShapeFillColor, not even an explicit nil
### GetShapeFillColorFieldName

`func (o *FieldInput) GetShapeFillColorFieldName() string`

GetShapeFillColorFieldName returns the ShapeFillColorFieldName field if non-nil, zero value otherwise.

### GetShapeFillColorFieldNameOk

`func (o *FieldInput) GetShapeFillColorFieldNameOk() (*string, bool)`

GetShapeFillColorFieldNameOk returns a tuple with the ShapeFillColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeFillColorFieldName

`func (o *FieldInput) SetShapeFillColorFieldName(v string)`

SetShapeFillColorFieldName sets ShapeFillColorFieldName field to given value.

### HasShapeFillColorFieldName

`func (o *FieldInput) HasShapeFillColorFieldName() bool`

HasShapeFillColorFieldName returns a boolean if a field has been set.

### SetShapeFillColorFieldNameNil

`func (o *FieldInput) SetShapeFillColorFieldNameNil(b bool)`

 SetShapeFillColorFieldNameNil sets the value for ShapeFillColorFieldName to be an explicit nil

### UnsetShapeFillColorFieldName
`func (o *FieldInput) UnsetShapeFillColorFieldName()`

UnsetShapeFillColorFieldName ensures that no value is present for ShapeFillColorFieldName, not even an explicit nil
### GetShapeFillColorFieldRequired

`func (o *FieldInput) GetShapeFillColorFieldRequired() bool`

GetShapeFillColorFieldRequired returns the ShapeFillColorFieldRequired field if non-nil, zero value otherwise.

### GetShapeFillColorFieldRequiredOk

`func (o *FieldInput) GetShapeFillColorFieldRequiredOk() (*bool, bool)`

GetShapeFillColorFieldRequiredOk returns a tuple with the ShapeFillColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeFillColorFieldRequired

`func (o *FieldInput) SetShapeFillColorFieldRequired(v bool)`

SetShapeFillColorFieldRequired sets ShapeFillColorFieldRequired field to given value.

### HasShapeFillColorFieldRequired

`func (o *FieldInput) HasShapeFillColorFieldRequired() bool`

HasShapeFillColorFieldRequired returns a boolean if a field has been set.

### GetCheckCharacter

`func (o *FieldInput) GetCheckCharacter() CheckCharacter`

GetCheckCharacter returns the CheckCharacter field if non-nil, zero value otherwise.

### GetCheckCharacterOk

`func (o *FieldInput) GetCheckCharacterOk() (*CheckCharacter, bool)`

GetCheckCharacterOk returns a tuple with the CheckCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCharacter

`func (o *FieldInput) SetCheckCharacter(v CheckCharacter)`

SetCheckCharacter sets CheckCharacter field to given value.

### HasCheckCharacter

`func (o *FieldInput) HasCheckCharacter() bool`

HasCheckCharacter returns a boolean if a field has been set.

### GetCheckColor

`func (o *FieldInput) GetCheckColor() string`

GetCheckColor returns the CheckColor field if non-nil, zero value otherwise.

### GetCheckColorOk

`func (o *FieldInput) GetCheckColorOk() (*string, bool)`

GetCheckColorOk returns a tuple with the CheckColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckColor

`func (o *FieldInput) SetCheckColor(v string)`

SetCheckColor sets CheckColor field to given value.

### HasCheckColor

`func (o *FieldInput) HasCheckColor() bool`

HasCheckColor returns a boolean if a field has been set.

### SetCheckColorNil

`func (o *FieldInput) SetCheckColorNil(b bool)`

 SetCheckColorNil sets the value for CheckColor to be an explicit nil

### UnsetCheckColor
`func (o *FieldInput) UnsetCheckColor()`

UnsetCheckColor ensures that no value is present for CheckColor, not even an explicit nil
### GetCheckColorFieldName

`func (o *FieldInput) GetCheckColorFieldName() string`

GetCheckColorFieldName returns the CheckColorFieldName field if non-nil, zero value otherwise.

### GetCheckColorFieldNameOk

`func (o *FieldInput) GetCheckColorFieldNameOk() (*string, bool)`

GetCheckColorFieldNameOk returns a tuple with the CheckColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckColorFieldName

`func (o *FieldInput) SetCheckColorFieldName(v string)`

SetCheckColorFieldName sets CheckColorFieldName field to given value.

### HasCheckColorFieldName

`func (o *FieldInput) HasCheckColorFieldName() bool`

HasCheckColorFieldName returns a boolean if a field has been set.

### SetCheckColorFieldNameNil

`func (o *FieldInput) SetCheckColorFieldNameNil(b bool)`

 SetCheckColorFieldNameNil sets the value for CheckColorFieldName to be an explicit nil

### UnsetCheckColorFieldName
`func (o *FieldInput) UnsetCheckColorFieldName()`

UnsetCheckColorFieldName ensures that no value is present for CheckColorFieldName, not even an explicit nil
### GetCheckColorFieldRequired

`func (o *FieldInput) GetCheckColorFieldRequired() bool`

GetCheckColorFieldRequired returns the CheckColorFieldRequired field if non-nil, zero value otherwise.

### GetCheckColorFieldRequiredOk

`func (o *FieldInput) GetCheckColorFieldRequiredOk() (*bool, bool)`

GetCheckColorFieldRequiredOk returns a tuple with the CheckColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckColorFieldRequired

`func (o *FieldInput) SetCheckColorFieldRequired(v bool)`

SetCheckColorFieldRequired sets CheckColorFieldRequired field to given value.

### HasCheckColorFieldRequired

`func (o *FieldInput) HasCheckColorFieldRequired() bool`

HasCheckColorFieldRequired returns a boolean if a field has been set.

### GetQrcodeColor

`func (o *FieldInput) GetQrcodeColor() string`

GetQrcodeColor returns the QrcodeColor field if non-nil, zero value otherwise.

### GetQrcodeColorOk

`func (o *FieldInput) GetQrcodeColorOk() (*string, bool)`

GetQrcodeColorOk returns a tuple with the QrcodeColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrcodeColor

`func (o *FieldInput) SetQrcodeColor(v string)`

SetQrcodeColor sets QrcodeColor field to given value.

### HasQrcodeColor

`func (o *FieldInput) HasQrcodeColor() bool`

HasQrcodeColor returns a boolean if a field has been set.

### SetQrcodeColorNil

`func (o *FieldInput) SetQrcodeColorNil(b bool)`

 SetQrcodeColorNil sets the value for QrcodeColor to be an explicit nil

### UnsetQrcodeColor
`func (o *FieldInput) UnsetQrcodeColor()`

UnsetQrcodeColor ensures that no value is present for QrcodeColor, not even an explicit nil
### GetQrcodeColorFieldName

`func (o *FieldInput) GetQrcodeColorFieldName() string`

GetQrcodeColorFieldName returns the QrcodeColorFieldName field if non-nil, zero value otherwise.

### GetQrcodeColorFieldNameOk

`func (o *FieldInput) GetQrcodeColorFieldNameOk() (*string, bool)`

GetQrcodeColorFieldNameOk returns a tuple with the QrcodeColorFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrcodeColorFieldName

`func (o *FieldInput) SetQrcodeColorFieldName(v string)`

SetQrcodeColorFieldName sets QrcodeColorFieldName field to given value.

### HasQrcodeColorFieldName

`func (o *FieldInput) HasQrcodeColorFieldName() bool`

HasQrcodeColorFieldName returns a boolean if a field has been set.

### SetQrcodeColorFieldNameNil

`func (o *FieldInput) SetQrcodeColorFieldNameNil(b bool)`

 SetQrcodeColorFieldNameNil sets the value for QrcodeColorFieldName to be an explicit nil

### UnsetQrcodeColorFieldName
`func (o *FieldInput) UnsetQrcodeColorFieldName()`

UnsetQrcodeColorFieldName ensures that no value is present for QrcodeColorFieldName, not even an explicit nil
### GetQrcodeColorFieldRequired

`func (o *FieldInput) GetQrcodeColorFieldRequired() bool`

GetQrcodeColorFieldRequired returns the QrcodeColorFieldRequired field if non-nil, zero value otherwise.

### GetQrcodeColorFieldRequiredOk

`func (o *FieldInput) GetQrcodeColorFieldRequiredOk() (*bool, bool)`

GetQrcodeColorFieldRequiredOk returns a tuple with the QrcodeColorFieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrcodeColorFieldRequired

`func (o *FieldInput) SetQrcodeColorFieldRequired(v bool)`

SetQrcodeColorFieldRequired sets QrcodeColorFieldRequired field to given value.

### HasQrcodeColorFieldRequired

`func (o *FieldInput) HasQrcodeColorFieldRequired() bool`

HasQrcodeColorFieldRequired returns a boolean if a field has been set.

### GetTrueText

`func (o *FieldInput) GetTrueText() string`

GetTrueText returns the TrueText field if non-nil, zero value otherwise.

### GetTrueTextOk

`func (o *FieldInput) GetTrueTextOk() (*string, bool)`

GetTrueTextOk returns a tuple with the TrueText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueText

`func (o *FieldInput) SetTrueText(v string)`

SetTrueText sets TrueText field to given value.

### HasTrueText

`func (o *FieldInput) HasTrueText() bool`

HasTrueText returns a boolean if a field has been set.

### SetTrueTextNil

`func (o *FieldInput) SetTrueTextNil(b bool)`

 SetTrueTextNil sets the value for TrueText to be an explicit nil

### UnsetTrueText
`func (o *FieldInput) UnsetTrueText()`

UnsetTrueText ensures that no value is present for TrueText, not even an explicit nil
### GetFalseText

`func (o *FieldInput) GetFalseText() string`

GetFalseText returns the FalseText field if non-nil, zero value otherwise.

### GetFalseTextOk

`func (o *FieldInput) GetFalseTextOk() (*string, bool)`

GetFalseTextOk returns a tuple with the FalseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseText

`func (o *FieldInput) SetFalseText(v string)`

SetFalseText sets FalseText field to given value.

### HasFalseText

`func (o *FieldInput) HasFalseText() bool`

HasFalseText returns a boolean if a field has been set.

### SetFalseTextNil

`func (o *FieldInput) SetFalseTextNil(b bool)`

 SetFalseTextNil sets the value for FalseText to be an explicit nil

### UnsetFalseText
`func (o *FieldInput) UnsetFalseText()`

UnsetFalseText ensures that no value is present for FalseText, not even an explicit nil
### GetCondition

`func (o *FieldInput) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *FieldInput) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *FieldInput) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *FieldInput) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *FieldInput) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *FieldInput) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetStringConditionType

`func (o *FieldInput) GetStringConditionType() StringConditionType`

GetStringConditionType returns the StringConditionType field if non-nil, zero value otherwise.

### GetStringConditionTypeOk

`func (o *FieldInput) GetStringConditionTypeOk() (*StringConditionType, bool)`

GetStringConditionTypeOk returns a tuple with the StringConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringConditionType

`func (o *FieldInput) SetStringConditionType(v StringConditionType)`

SetStringConditionType sets StringConditionType field to given value.

### HasStringConditionType

`func (o *FieldInput) HasStringConditionType() bool`

HasStringConditionType returns a boolean if a field has been set.

### GetNumberConditionType

`func (o *FieldInput) GetNumberConditionType() NumberConditionType`

GetNumberConditionType returns the NumberConditionType field if non-nil, zero value otherwise.

### GetNumberConditionTypeOk

`func (o *FieldInput) GetNumberConditionTypeOk() (*NumberConditionType, bool)`

GetNumberConditionTypeOk returns a tuple with the NumberConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionType

`func (o *FieldInput) SetNumberConditionType(v NumberConditionType)`

SetNumberConditionType sets NumberConditionType field to given value.

### HasNumberConditionType

`func (o *FieldInput) HasNumberConditionType() bool`

HasNumberConditionType returns a boolean if a field has been set.

### GetNumberConditionRangeMin

`func (o *FieldInput) GetNumberConditionRangeMin() float32`

GetNumberConditionRangeMin returns the NumberConditionRangeMin field if non-nil, zero value otherwise.

### GetNumberConditionRangeMinOk

`func (o *FieldInput) GetNumberConditionRangeMinOk() (*float32, bool)`

GetNumberConditionRangeMinOk returns a tuple with the NumberConditionRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeMin

`func (o *FieldInput) SetNumberConditionRangeMin(v float32)`

SetNumberConditionRangeMin sets NumberConditionRangeMin field to given value.

### HasNumberConditionRangeMin

`func (o *FieldInput) HasNumberConditionRangeMin() bool`

HasNumberConditionRangeMin returns a boolean if a field has been set.

### SetNumberConditionRangeMinNil

`func (o *FieldInput) SetNumberConditionRangeMinNil(b bool)`

 SetNumberConditionRangeMinNil sets the value for NumberConditionRangeMin to be an explicit nil

### UnsetNumberConditionRangeMin
`func (o *FieldInput) UnsetNumberConditionRangeMin()`

UnsetNumberConditionRangeMin ensures that no value is present for NumberConditionRangeMin, not even an explicit nil
### GetNumberConditionRangeMax

`func (o *FieldInput) GetNumberConditionRangeMax() float32`

GetNumberConditionRangeMax returns the NumberConditionRangeMax field if non-nil, zero value otherwise.

### GetNumberConditionRangeMaxOk

`func (o *FieldInput) GetNumberConditionRangeMaxOk() (*float32, bool)`

GetNumberConditionRangeMaxOk returns a tuple with the NumberConditionRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeMax

`func (o *FieldInput) SetNumberConditionRangeMax(v float32)`

SetNumberConditionRangeMax sets NumberConditionRangeMax field to given value.

### HasNumberConditionRangeMax

`func (o *FieldInput) HasNumberConditionRangeMax() bool`

HasNumberConditionRangeMax returns a boolean if a field has been set.

### SetNumberConditionRangeMaxNil

`func (o *FieldInput) SetNumberConditionRangeMaxNil(b bool)`

 SetNumberConditionRangeMaxNil sets the value for NumberConditionRangeMax to be an explicit nil

### UnsetNumberConditionRangeMax
`func (o *FieldInput) UnsetNumberConditionRangeMax()`

UnsetNumberConditionRangeMax ensures that no value is present for NumberConditionRangeMax, not even an explicit nil
### GetNumberConditionRangeExclusiveMin

`func (o *FieldInput) GetNumberConditionRangeExclusiveMin() bool`

GetNumberConditionRangeExclusiveMin returns the NumberConditionRangeExclusiveMin field if non-nil, zero value otherwise.

### GetNumberConditionRangeExclusiveMinOk

`func (o *FieldInput) GetNumberConditionRangeExclusiveMinOk() (*bool, bool)`

GetNumberConditionRangeExclusiveMinOk returns a tuple with the NumberConditionRangeExclusiveMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeExclusiveMin

`func (o *FieldInput) SetNumberConditionRangeExclusiveMin(v bool)`

SetNumberConditionRangeExclusiveMin sets NumberConditionRangeExclusiveMin field to given value.

### HasNumberConditionRangeExclusiveMin

`func (o *FieldInput) HasNumberConditionRangeExclusiveMin() bool`

HasNumberConditionRangeExclusiveMin returns a boolean if a field has been set.

### GetNumberConditionRangeExclusiveMax

`func (o *FieldInput) GetNumberConditionRangeExclusiveMax() bool`

GetNumberConditionRangeExclusiveMax returns the NumberConditionRangeExclusiveMax field if non-nil, zero value otherwise.

### GetNumberConditionRangeExclusiveMaxOk

`func (o *FieldInput) GetNumberConditionRangeExclusiveMaxOk() (*bool, bool)`

GetNumberConditionRangeExclusiveMaxOk returns a tuple with the NumberConditionRangeExclusiveMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeExclusiveMax

`func (o *FieldInput) SetNumberConditionRangeExclusiveMax(v bool)`

SetNumberConditionRangeExclusiveMax sets NumberConditionRangeExclusiveMax field to given value.

### HasNumberConditionRangeExclusiveMax

`func (o *FieldInput) HasNumberConditionRangeExclusiveMax() bool`

HasNumberConditionRangeExclusiveMax returns a boolean if a field has been set.

### GetInvertBooleanCondition

`func (o *FieldInput) GetInvertBooleanCondition() bool`

GetInvertBooleanCondition returns the InvertBooleanCondition field if non-nil, zero value otherwise.

### GetInvertBooleanConditionOk

`func (o *FieldInput) GetInvertBooleanConditionOk() (*bool, bool)`

GetInvertBooleanConditionOk returns a tuple with the InvertBooleanCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertBooleanCondition

`func (o *FieldInput) SetInvertBooleanCondition(v bool)`

SetInvertBooleanCondition sets InvertBooleanCondition field to given value.

### HasInvertBooleanCondition

`func (o *FieldInput) HasInvertBooleanCondition() bool`

HasInvertBooleanCondition returns a boolean if a field has been set.

### GetCurrency

`func (o *FieldInput) GetCurrency() bool`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FieldInput) GetCurrencyOk() (*bool, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FieldInput) SetCurrency(v bool)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FieldInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *FieldInput) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *FieldInput) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *FieldInput) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *FieldInput) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### SetDecimalPlacesNil

`func (o *FieldInput) SetDecimalPlacesNil(b bool)`

 SetDecimalPlacesNil sets the value for DecimalPlaces to be an explicit nil

### UnsetDecimalPlaces
`func (o *FieldInput) UnsetDecimalPlaces()`

UnsetDecimalPlaces ensures that no value is present for DecimalPlaces, not even an explicit nil
### GetImageScaleType

`func (o *FieldInput) GetImageScaleType() ImageScaleType`

GetImageScaleType returns the ImageScaleType field if non-nil, zero value otherwise.

### GetImageScaleTypeOk

`func (o *FieldInput) GetImageScaleTypeOk() (*ImageScaleType, bool)`

GetImageScaleTypeOk returns a tuple with the ImageScaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScaleType

`func (o *FieldInput) SetImageScaleType(v ImageScaleType)`

SetImageScaleType sets ImageScaleType field to given value.

### HasImageScaleType

`func (o *FieldInput) HasImageScaleType() bool`

HasImageScaleType returns a boolean if a field has been set.

### GetImageGravity

`func (o *FieldInput) GetImageGravity() ImageGravity`

GetImageGravity returns the ImageGravity field if non-nil, zero value otherwise.

### GetImageGravityOk

`func (o *FieldInput) GetImageGravityOk() (*ImageGravity, bool)`

GetImageGravityOk returns a tuple with the ImageGravity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageGravity

`func (o *FieldInput) SetImageGravity(v ImageGravity)`

SetImageGravity sets ImageGravity field to given value.

### HasImageGravity

`func (o *FieldInput) HasImageGravity() bool`

HasImageGravity returns a boolean if a field has been set.

### GetDependencies

`func (o *FieldInput) GetDependencies() FieldDependencyInput`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *FieldInput) GetDependenciesOk() (*FieldDependencyInput, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *FieldInput) SetDependencies(v FieldDependencyInput)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *FieldInput) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### SetDependenciesNil

`func (o *FieldInput) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *FieldInput) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil
### GetId

`func (o *FieldInput) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldInput) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldInput) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FieldInput) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


