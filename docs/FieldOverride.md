# FieldOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Static** | Pointer to **bool** | Print the field's `default` (a fixed value, or a formula for a number field) on the PDF instead of a value from `data`. | [optional] 
**Editable** | Pointer to **bool** | Keep this field fillable in the PDF instead of flattening it to text. The submission's `editable` overrides it. | [optional] 
**Default** | Pointer to **NullableString** | The value used when `data` has none for the field. | [optional] 
**Title** | Pointer to **string** | The field's label on the signing form. | [optional] 
**Required** | Pointer to **bool** | Whether `data` must have a value for the field. | [optional] 
**OptionList** | Pointer to **NullableString** | The values the field can take, separated by commas, such as `Yes,No`. Null allows any value. | [optional] 
**CombinedFieldNames** | Pointer to **string** | The names of the fields a combined field joins, separated by commas. | [optional] 
**CombinedFieldType** | Pointer to [**CombinedFieldType**](CombinedFieldType.md) | How a combined field joins its fields' values. | [optional] 
**CombinedFieldSeparator** | Pointer to **string** | The text between the values of a combined field, such as `, `. | [optional] 
**CombinedFieldFormat** | Pointer to **string** | The text a combined field prints, with `$0`, `$1` and so on for its fields' values, such as `$0, $1`. | [optional] 
**Integer** | Pointer to **bool** | Whether a number field takes only whole numbers. | [optional] 
**IncludeTime** | Pointer to **NullableBool** | Whether a date field takes a date and time (RFC 3339) rather than a date (YYYY-MM-DD). | [optional] 
**SignatureAllowDraw** | Pointer to **bool** | Let the person draw their signature or initials on the signing form. | [optional] 
**SignatureAllowType** | Pointer to **bool** | Let the person type their signature or initials on the signing form. | [optional] 
**BarcodeSymbology** | Pointer to [**BarcodeSymbology**](BarcodeSymbology.md) | The kind of barcode a field with display_type `barcode` prints. | [optional] 
**Hidden** | Pointer to **bool** | Leave the field off the PDF. Its value is still in the data, and combined fields can use it. | [optional] 
**X** | Pointer to **float32** | The distance from the left edge of the page, in the template editor's page units. | [optional] 
**Y** | Pointer to **float32** | The distance from the top of the page, in the template editor's page units. | [optional] 
**Width** | Pointer to **float32** | The field's width, in the template editor's page units. | [optional] 
**Height** | Pointer to **float32** | The field's height, in the template editor's page units. | [optional] 
**Rotation** | Pointer to **float32** | The field's rotation in degrees, clockwise. | [optional] 
**DisplayType** | Pointer to [**FieldDisplayType**](FieldDisplayType.md) | How the field's value is drawn on the PDF. | [optional] 
**Opacity** | Pointer to **float32** | The field's opacity, from 0 (invisible) to 1 (solid). | [optional] 
**DateTimeFormat** | Pointer to **NullableString** | How a date field's value is printed, as a strftime format such as `%m/%d/%Y`. Null uses the template's date format. | [optional] 
**Typeface** | Pointer to [**NullableTypeface**](Typeface.md) | The font of the field's text. Null uses the template's default. | [optional] 
**FontSize** | Pointer to **NullableInt32** | The font size in points, a whole number. Null uses the template's default. | [optional] 
**CharacterSpacing** | Pointer to **float32** | Extra space between characters, in points (0 to 50). | [optional] 
**Bold** | Pointer to **bool** | Print the text in bold. | [optional] 
**Uppercase** | Pointer to **bool** | Print the text in capitals. | [optional] 
**Strikethrough** | Pointer to **bool** | Strike through the text. | [optional] 
**KeepLeadingSpaces** | Pointer to **bool** | Keep the spaces at the start of the text and of each line, so you can line text up with spaces. | [optional] 
**Overflow** | Pointer to [**TextOverflow**](TextOverflow.md) | What happens to text that doesn't fit in the field. | [optional] 
**Multiline** | Pointer to **bool** | Wrap the text over several lines when it doesn't fit on one. | [optional] 
**MultilineLines** | Pointer to **NullableInt32** | The number of lines in a multiline field, for a PDF with ruled lines. Null for no fixed number of lines. | [optional] 
**MaxLength** | Pointer to **NullableInt32** | The most characters the field takes. Null for no limit. | [optional] 
**Color** | Pointer to **NullableString** | The text color, as 6 hex digits without `#`, such as `444444`. Null uses the template's default. | [optional] 
**BackgroundColor** | Pointer to **NullableString** | The field's background color, as 6 hex digits without `#`. Null for none. | [optional] 
**Alignment** | Pointer to [**TextAlignment**](TextAlignment.md) | How text is aligned in its field. `justify` lines up both edges of every line of multiline text except the last of each paragraph. | [optional] 
**VAlignment** | Pointer to [**TextVerticalAlignment**](TextVerticalAlignment.md) | How text is aligned vertically in its field. | [optional] 
**Comb** | Pointer to **bool** | Print one character in each of the field's equally spaced boxes (combs). | [optional] 
**CombNumberOfCells** | Pointer to **int32** | The number of boxes in a comb field. | [optional] 
**CombValueOffset** | Pointer to **int32** | The character of the value that the first box starts from, for a value split over several comb fields (such as the parts of a Social Security number). | [optional] 
**ShapeType** | Pointer to [**ShapeType**](ShapeType.md) | The shape a field with display_type `shape` draws. | [optional] 
**ShapeSides** | Pointer to **int32** | The number of sides of a `polygon` shape. | [optional] 
**ShapeBorderWidth** | Pointer to **int32** | The width of a shape's border, in points, a whole number. | [optional] 
**ShapeBorderColor** | Pointer to **string** | The color of a shape's border, as 6 hex digits without `#`. | [optional] 
**ShapeFillColor** | Pointer to **NullableString** | The color a shape is filled with, as 6 hex digits without `#`. Null for no fill. | [optional] 
**CheckCharacter** | Pointer to [**CheckCharacter**](CheckCharacter.md) | The character that a ticked check box prints. | [optional] 
**CheckColor** | Pointer to **NullableString** | The color of the check mark, as 6 hex digits without `#`. Null uses the template's default. | [optional] 
**QrcodeColor** | Pointer to **NullableString** | The color of a QR code, as 6 hex digits without `#`. Null for black. | [optional] 
**TrueText** | Pointer to **NullableString** | What a boolean field with display_type `text` prints for `true`. | [optional] 
**FalseText** | Pointer to **NullableString** | What a boolean field with display_type `text` prints for `false`. | [optional] 
**Condition** | Pointer to **NullableString** | The value a check mark or shape is drawn for: it's drawn when the field's value matches, as `string_condition_type` or `number_condition_type` says. | [optional] 
**StringConditionType** | Pointer to [**StringConditionType**](StringConditionType.md) | How a string field's value is compared with `condition`. | [optional] 
**NumberConditionType** | Pointer to [**NumberConditionType**](NumberConditionType.md) | How a number field's value is compared with `condition`, or with `number_condition_range_min` and `number_condition_range_max` for `range`. | [optional] 
**NumberConditionRangeMin** | Pointer to **NullableFloat32** | The lower end of a `range` condition. | [optional] 
**NumberConditionRangeMax** | Pointer to **NullableFloat32** | The upper end of a `range` condition. | [optional] 
**NumberConditionRangeExclusiveMin** | Pointer to **bool** | Leave the lower end itself out of a `range` condition. | [optional] 
**NumberConditionRangeExclusiveMax** | Pointer to **bool** | Leave the upper end itself out of a `range` condition. | [optional] 
**InvertBooleanCondition** | Pointer to **bool** | Draw a boolean field's check mark or shape for `false` instead of `true`. | [optional] 
**DecimalPlaces** | Pointer to **NullableInt32** | The number of decimal places a number field prints. Null prints the value as it was sent. | [optional] 
**ImageScaleType** | Pointer to [**ImageScaleType**](ImageScaleType.md) | How an image is resized to its field. | [optional] 
**ImageGravity** | Pointer to [**ImageGravity**](ImageGravity.md) | Where an image sits in its field when it doesn't fill it, or which part is kept when it's cropped. | [optional] 

## Methods

### NewFieldOverride

`func NewFieldOverride() *FieldOverride`

NewFieldOverride instantiates a new FieldOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldOverrideWithDefaults

`func NewFieldOverrideWithDefaults() *FieldOverride`

NewFieldOverrideWithDefaults instantiates a new FieldOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatic

`func (o *FieldOverride) GetStatic() bool`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *FieldOverride) GetStaticOk() (*bool, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *FieldOverride) SetStatic(v bool)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *FieldOverride) HasStatic() bool`

HasStatic returns a boolean if a field has been set.

### GetEditable

`func (o *FieldOverride) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *FieldOverride) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *FieldOverride) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *FieldOverride) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetDefault

`func (o *FieldOverride) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FieldOverride) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FieldOverride) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FieldOverride) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *FieldOverride) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *FieldOverride) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetTitle

`func (o *FieldOverride) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FieldOverride) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FieldOverride) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FieldOverride) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRequired

`func (o *FieldOverride) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldOverride) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldOverride) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FieldOverride) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetOptionList

`func (o *FieldOverride) GetOptionList() string`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *FieldOverride) GetOptionListOk() (*string, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *FieldOverride) SetOptionList(v string)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *FieldOverride) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### SetOptionListNil

`func (o *FieldOverride) SetOptionListNil(b bool)`

 SetOptionListNil sets the value for OptionList to be an explicit nil

### UnsetOptionList
`func (o *FieldOverride) UnsetOptionList()`

UnsetOptionList ensures that no value is present for OptionList, not even an explicit nil
### GetCombinedFieldNames

`func (o *FieldOverride) GetCombinedFieldNames() string`

GetCombinedFieldNames returns the CombinedFieldNames field if non-nil, zero value otherwise.

### GetCombinedFieldNamesOk

`func (o *FieldOverride) GetCombinedFieldNamesOk() (*string, bool)`

GetCombinedFieldNamesOk returns a tuple with the CombinedFieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldNames

`func (o *FieldOverride) SetCombinedFieldNames(v string)`

SetCombinedFieldNames sets CombinedFieldNames field to given value.

### HasCombinedFieldNames

`func (o *FieldOverride) HasCombinedFieldNames() bool`

HasCombinedFieldNames returns a boolean if a field has been set.

### GetCombinedFieldType

`func (o *FieldOverride) GetCombinedFieldType() CombinedFieldType`

GetCombinedFieldType returns the CombinedFieldType field if non-nil, zero value otherwise.

### GetCombinedFieldTypeOk

`func (o *FieldOverride) GetCombinedFieldTypeOk() (*CombinedFieldType, bool)`

GetCombinedFieldTypeOk returns a tuple with the CombinedFieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldType

`func (o *FieldOverride) SetCombinedFieldType(v CombinedFieldType)`

SetCombinedFieldType sets CombinedFieldType field to given value.

### HasCombinedFieldType

`func (o *FieldOverride) HasCombinedFieldType() bool`

HasCombinedFieldType returns a boolean if a field has been set.

### GetCombinedFieldSeparator

`func (o *FieldOverride) GetCombinedFieldSeparator() string`

GetCombinedFieldSeparator returns the CombinedFieldSeparator field if non-nil, zero value otherwise.

### GetCombinedFieldSeparatorOk

`func (o *FieldOverride) GetCombinedFieldSeparatorOk() (*string, bool)`

GetCombinedFieldSeparatorOk returns a tuple with the CombinedFieldSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldSeparator

`func (o *FieldOverride) SetCombinedFieldSeparator(v string)`

SetCombinedFieldSeparator sets CombinedFieldSeparator field to given value.

### HasCombinedFieldSeparator

`func (o *FieldOverride) HasCombinedFieldSeparator() bool`

HasCombinedFieldSeparator returns a boolean if a field has been set.

### GetCombinedFieldFormat

`func (o *FieldOverride) GetCombinedFieldFormat() string`

GetCombinedFieldFormat returns the CombinedFieldFormat field if non-nil, zero value otherwise.

### GetCombinedFieldFormatOk

`func (o *FieldOverride) GetCombinedFieldFormatOk() (*string, bool)`

GetCombinedFieldFormatOk returns a tuple with the CombinedFieldFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFieldFormat

`func (o *FieldOverride) SetCombinedFieldFormat(v string)`

SetCombinedFieldFormat sets CombinedFieldFormat field to given value.

### HasCombinedFieldFormat

`func (o *FieldOverride) HasCombinedFieldFormat() bool`

HasCombinedFieldFormat returns a boolean if a field has been set.

### GetInteger

`func (o *FieldOverride) GetInteger() bool`

GetInteger returns the Integer field if non-nil, zero value otherwise.

### GetIntegerOk

`func (o *FieldOverride) GetIntegerOk() (*bool, bool)`

GetIntegerOk returns a tuple with the Integer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteger

`func (o *FieldOverride) SetInteger(v bool)`

SetInteger sets Integer field to given value.

### HasInteger

`func (o *FieldOverride) HasInteger() bool`

HasInteger returns a boolean if a field has been set.

### GetIncludeTime

`func (o *FieldOverride) GetIncludeTime() bool`

GetIncludeTime returns the IncludeTime field if non-nil, zero value otherwise.

### GetIncludeTimeOk

`func (o *FieldOverride) GetIncludeTimeOk() (*bool, bool)`

GetIncludeTimeOk returns a tuple with the IncludeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTime

`func (o *FieldOverride) SetIncludeTime(v bool)`

SetIncludeTime sets IncludeTime field to given value.

### HasIncludeTime

`func (o *FieldOverride) HasIncludeTime() bool`

HasIncludeTime returns a boolean if a field has been set.

### SetIncludeTimeNil

`func (o *FieldOverride) SetIncludeTimeNil(b bool)`

 SetIncludeTimeNil sets the value for IncludeTime to be an explicit nil

### UnsetIncludeTime
`func (o *FieldOverride) UnsetIncludeTime()`

UnsetIncludeTime ensures that no value is present for IncludeTime, not even an explicit nil
### GetSignatureAllowDraw

`func (o *FieldOverride) GetSignatureAllowDraw() bool`

GetSignatureAllowDraw returns the SignatureAllowDraw field if non-nil, zero value otherwise.

### GetSignatureAllowDrawOk

`func (o *FieldOverride) GetSignatureAllowDrawOk() (*bool, bool)`

GetSignatureAllowDrawOk returns a tuple with the SignatureAllowDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAllowDraw

`func (o *FieldOverride) SetSignatureAllowDraw(v bool)`

SetSignatureAllowDraw sets SignatureAllowDraw field to given value.

### HasSignatureAllowDraw

`func (o *FieldOverride) HasSignatureAllowDraw() bool`

HasSignatureAllowDraw returns a boolean if a field has been set.

### GetSignatureAllowType

`func (o *FieldOverride) GetSignatureAllowType() bool`

GetSignatureAllowType returns the SignatureAllowType field if non-nil, zero value otherwise.

### GetSignatureAllowTypeOk

`func (o *FieldOverride) GetSignatureAllowTypeOk() (*bool, bool)`

GetSignatureAllowTypeOk returns a tuple with the SignatureAllowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAllowType

`func (o *FieldOverride) SetSignatureAllowType(v bool)`

SetSignatureAllowType sets SignatureAllowType field to given value.

### HasSignatureAllowType

`func (o *FieldOverride) HasSignatureAllowType() bool`

HasSignatureAllowType returns a boolean if a field has been set.

### GetBarcodeSymbology

`func (o *FieldOverride) GetBarcodeSymbology() BarcodeSymbology`

GetBarcodeSymbology returns the BarcodeSymbology field if non-nil, zero value otherwise.

### GetBarcodeSymbologyOk

`func (o *FieldOverride) GetBarcodeSymbologyOk() (*BarcodeSymbology, bool)`

GetBarcodeSymbologyOk returns a tuple with the BarcodeSymbology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeSymbology

`func (o *FieldOverride) SetBarcodeSymbology(v BarcodeSymbology)`

SetBarcodeSymbology sets BarcodeSymbology field to given value.

### HasBarcodeSymbology

`func (o *FieldOverride) HasBarcodeSymbology() bool`

HasBarcodeSymbology returns a boolean if a field has been set.

### GetHidden

`func (o *FieldOverride) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *FieldOverride) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *FieldOverride) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *FieldOverride) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetX

`func (o *FieldOverride) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *FieldOverride) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *FieldOverride) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *FieldOverride) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *FieldOverride) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *FieldOverride) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *FieldOverride) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *FieldOverride) HasY() bool`

HasY returns a boolean if a field has been set.

### GetWidth

`func (o *FieldOverride) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *FieldOverride) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *FieldOverride) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *FieldOverride) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *FieldOverride) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *FieldOverride) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *FieldOverride) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *FieldOverride) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetRotation

`func (o *FieldOverride) GetRotation() float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *FieldOverride) GetRotationOk() (*float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *FieldOverride) SetRotation(v float32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *FieldOverride) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetDisplayType

`func (o *FieldOverride) GetDisplayType() FieldDisplayType`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *FieldOverride) GetDisplayTypeOk() (*FieldDisplayType, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *FieldOverride) SetDisplayType(v FieldDisplayType)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *FieldOverride) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetOpacity

`func (o *FieldOverride) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *FieldOverride) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *FieldOverride) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *FieldOverride) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *FieldOverride) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *FieldOverride) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *FieldOverride) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *FieldOverride) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### SetDateTimeFormatNil

`func (o *FieldOverride) SetDateTimeFormatNil(b bool)`

 SetDateTimeFormatNil sets the value for DateTimeFormat to be an explicit nil

### UnsetDateTimeFormat
`func (o *FieldOverride) UnsetDateTimeFormat()`

UnsetDateTimeFormat ensures that no value is present for DateTimeFormat, not even an explicit nil
### GetTypeface

`func (o *FieldOverride) GetTypeface() Typeface`

GetTypeface returns the Typeface field if non-nil, zero value otherwise.

### GetTypefaceOk

`func (o *FieldOverride) GetTypefaceOk() (*Typeface, bool)`

GetTypefaceOk returns a tuple with the Typeface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeface

`func (o *FieldOverride) SetTypeface(v Typeface)`

SetTypeface sets Typeface field to given value.

### HasTypeface

`func (o *FieldOverride) HasTypeface() bool`

HasTypeface returns a boolean if a field has been set.

### SetTypefaceNil

`func (o *FieldOverride) SetTypefaceNil(b bool)`

 SetTypefaceNil sets the value for Typeface to be an explicit nil

### UnsetTypeface
`func (o *FieldOverride) UnsetTypeface()`

UnsetTypeface ensures that no value is present for Typeface, not even an explicit nil
### GetFontSize

`func (o *FieldOverride) GetFontSize() int32`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *FieldOverride) GetFontSizeOk() (*int32, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *FieldOverride) SetFontSize(v int32)`

SetFontSize sets FontSize field to given value.

### HasFontSize

`func (o *FieldOverride) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.

### SetFontSizeNil

`func (o *FieldOverride) SetFontSizeNil(b bool)`

 SetFontSizeNil sets the value for FontSize to be an explicit nil

### UnsetFontSize
`func (o *FieldOverride) UnsetFontSize()`

UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil
### GetCharacterSpacing

`func (o *FieldOverride) GetCharacterSpacing() float32`

GetCharacterSpacing returns the CharacterSpacing field if non-nil, zero value otherwise.

### GetCharacterSpacingOk

`func (o *FieldOverride) GetCharacterSpacingOk() (*float32, bool)`

GetCharacterSpacingOk returns a tuple with the CharacterSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterSpacing

`func (o *FieldOverride) SetCharacterSpacing(v float32)`

SetCharacterSpacing sets CharacterSpacing field to given value.

### HasCharacterSpacing

`func (o *FieldOverride) HasCharacterSpacing() bool`

HasCharacterSpacing returns a boolean if a field has been set.

### GetBold

`func (o *FieldOverride) GetBold() bool`

GetBold returns the Bold field if non-nil, zero value otherwise.

### GetBoldOk

`func (o *FieldOverride) GetBoldOk() (*bool, bool)`

GetBoldOk returns a tuple with the Bold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBold

`func (o *FieldOverride) SetBold(v bool)`

SetBold sets Bold field to given value.

### HasBold

`func (o *FieldOverride) HasBold() bool`

HasBold returns a boolean if a field has been set.

### GetUppercase

`func (o *FieldOverride) GetUppercase() bool`

GetUppercase returns the Uppercase field if non-nil, zero value otherwise.

### GetUppercaseOk

`func (o *FieldOverride) GetUppercaseOk() (*bool, bool)`

GetUppercaseOk returns a tuple with the Uppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUppercase

`func (o *FieldOverride) SetUppercase(v bool)`

SetUppercase sets Uppercase field to given value.

### HasUppercase

`func (o *FieldOverride) HasUppercase() bool`

HasUppercase returns a boolean if a field has been set.

### GetStrikethrough

`func (o *FieldOverride) GetStrikethrough() bool`

GetStrikethrough returns the Strikethrough field if non-nil, zero value otherwise.

### GetStrikethroughOk

`func (o *FieldOverride) GetStrikethroughOk() (*bool, bool)`

GetStrikethroughOk returns a tuple with the Strikethrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikethrough

`func (o *FieldOverride) SetStrikethrough(v bool)`

SetStrikethrough sets Strikethrough field to given value.

### HasStrikethrough

`func (o *FieldOverride) HasStrikethrough() bool`

HasStrikethrough returns a boolean if a field has been set.

### GetKeepLeadingSpaces

`func (o *FieldOverride) GetKeepLeadingSpaces() bool`

GetKeepLeadingSpaces returns the KeepLeadingSpaces field if non-nil, zero value otherwise.

### GetKeepLeadingSpacesOk

`func (o *FieldOverride) GetKeepLeadingSpacesOk() (*bool, bool)`

GetKeepLeadingSpacesOk returns a tuple with the KeepLeadingSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepLeadingSpaces

`func (o *FieldOverride) SetKeepLeadingSpaces(v bool)`

SetKeepLeadingSpaces sets KeepLeadingSpaces field to given value.

### HasKeepLeadingSpaces

`func (o *FieldOverride) HasKeepLeadingSpaces() bool`

HasKeepLeadingSpaces returns a boolean if a field has been set.

### GetOverflow

`func (o *FieldOverride) GetOverflow() TextOverflow`

GetOverflow returns the Overflow field if non-nil, zero value otherwise.

### GetOverflowOk

`func (o *FieldOverride) GetOverflowOk() (*TextOverflow, bool)`

GetOverflowOk returns a tuple with the Overflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflow

`func (o *FieldOverride) SetOverflow(v TextOverflow)`

SetOverflow sets Overflow field to given value.

### HasOverflow

`func (o *FieldOverride) HasOverflow() bool`

HasOverflow returns a boolean if a field has been set.

### GetMultiline

`func (o *FieldOverride) GetMultiline() bool`

GetMultiline returns the Multiline field if non-nil, zero value otherwise.

### GetMultilineOk

`func (o *FieldOverride) GetMultilineOk() (*bool, bool)`

GetMultilineOk returns a tuple with the Multiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiline

`func (o *FieldOverride) SetMultiline(v bool)`

SetMultiline sets Multiline field to given value.

### HasMultiline

`func (o *FieldOverride) HasMultiline() bool`

HasMultiline returns a boolean if a field has been set.

### GetMultilineLines

`func (o *FieldOverride) GetMultilineLines() int32`

GetMultilineLines returns the MultilineLines field if non-nil, zero value otherwise.

### GetMultilineLinesOk

`func (o *FieldOverride) GetMultilineLinesOk() (*int32, bool)`

GetMultilineLinesOk returns a tuple with the MultilineLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultilineLines

`func (o *FieldOverride) SetMultilineLines(v int32)`

SetMultilineLines sets MultilineLines field to given value.

### HasMultilineLines

`func (o *FieldOverride) HasMultilineLines() bool`

HasMultilineLines returns a boolean if a field has been set.

### SetMultilineLinesNil

`func (o *FieldOverride) SetMultilineLinesNil(b bool)`

 SetMultilineLinesNil sets the value for MultilineLines to be an explicit nil

### UnsetMultilineLines
`func (o *FieldOverride) UnsetMultilineLines()`

UnsetMultilineLines ensures that no value is present for MultilineLines, not even an explicit nil
### GetMaxLength

`func (o *FieldOverride) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *FieldOverride) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *FieldOverride) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *FieldOverride) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *FieldOverride) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *FieldOverride) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetColor

`func (o *FieldOverride) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *FieldOverride) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *FieldOverride) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *FieldOverride) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *FieldOverride) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *FieldOverride) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetBackgroundColor

`func (o *FieldOverride) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *FieldOverride) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *FieldOverride) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *FieldOverride) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *FieldOverride) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *FieldOverride) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetAlignment

`func (o *FieldOverride) GetAlignment() TextAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *FieldOverride) GetAlignmentOk() (*TextAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *FieldOverride) SetAlignment(v TextAlignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *FieldOverride) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetVAlignment

`func (o *FieldOverride) GetVAlignment() TextVerticalAlignment`

GetVAlignment returns the VAlignment field if non-nil, zero value otherwise.

### GetVAlignmentOk

`func (o *FieldOverride) GetVAlignmentOk() (*TextVerticalAlignment, bool)`

GetVAlignmentOk returns a tuple with the VAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAlignment

`func (o *FieldOverride) SetVAlignment(v TextVerticalAlignment)`

SetVAlignment sets VAlignment field to given value.

### HasVAlignment

`func (o *FieldOverride) HasVAlignment() bool`

HasVAlignment returns a boolean if a field has been set.

### GetComb

`func (o *FieldOverride) GetComb() bool`

GetComb returns the Comb field if non-nil, zero value otherwise.

### GetCombOk

`func (o *FieldOverride) GetCombOk() (*bool, bool)`

GetCombOk returns a tuple with the Comb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComb

`func (o *FieldOverride) SetComb(v bool)`

SetComb sets Comb field to given value.

### HasComb

`func (o *FieldOverride) HasComb() bool`

HasComb returns a boolean if a field has been set.

### GetCombNumberOfCells

`func (o *FieldOverride) GetCombNumberOfCells() int32`

GetCombNumberOfCells returns the CombNumberOfCells field if non-nil, zero value otherwise.

### GetCombNumberOfCellsOk

`func (o *FieldOverride) GetCombNumberOfCellsOk() (*int32, bool)`

GetCombNumberOfCellsOk returns a tuple with the CombNumberOfCells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombNumberOfCells

`func (o *FieldOverride) SetCombNumberOfCells(v int32)`

SetCombNumberOfCells sets CombNumberOfCells field to given value.

### HasCombNumberOfCells

`func (o *FieldOverride) HasCombNumberOfCells() bool`

HasCombNumberOfCells returns a boolean if a field has been set.

### GetCombValueOffset

`func (o *FieldOverride) GetCombValueOffset() int32`

GetCombValueOffset returns the CombValueOffset field if non-nil, zero value otherwise.

### GetCombValueOffsetOk

`func (o *FieldOverride) GetCombValueOffsetOk() (*int32, bool)`

GetCombValueOffsetOk returns a tuple with the CombValueOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombValueOffset

`func (o *FieldOverride) SetCombValueOffset(v int32)`

SetCombValueOffset sets CombValueOffset field to given value.

### HasCombValueOffset

`func (o *FieldOverride) HasCombValueOffset() bool`

HasCombValueOffset returns a boolean if a field has been set.

### GetShapeType

`func (o *FieldOverride) GetShapeType() ShapeType`

GetShapeType returns the ShapeType field if non-nil, zero value otherwise.

### GetShapeTypeOk

`func (o *FieldOverride) GetShapeTypeOk() (*ShapeType, bool)`

GetShapeTypeOk returns a tuple with the ShapeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeType

`func (o *FieldOverride) SetShapeType(v ShapeType)`

SetShapeType sets ShapeType field to given value.

### HasShapeType

`func (o *FieldOverride) HasShapeType() bool`

HasShapeType returns a boolean if a field has been set.

### GetShapeSides

`func (o *FieldOverride) GetShapeSides() int32`

GetShapeSides returns the ShapeSides field if non-nil, zero value otherwise.

### GetShapeSidesOk

`func (o *FieldOverride) GetShapeSidesOk() (*int32, bool)`

GetShapeSidesOk returns a tuple with the ShapeSides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeSides

`func (o *FieldOverride) SetShapeSides(v int32)`

SetShapeSides sets ShapeSides field to given value.

### HasShapeSides

`func (o *FieldOverride) HasShapeSides() bool`

HasShapeSides returns a boolean if a field has been set.

### GetShapeBorderWidth

`func (o *FieldOverride) GetShapeBorderWidth() int32`

GetShapeBorderWidth returns the ShapeBorderWidth field if non-nil, zero value otherwise.

### GetShapeBorderWidthOk

`func (o *FieldOverride) GetShapeBorderWidthOk() (*int32, bool)`

GetShapeBorderWidthOk returns a tuple with the ShapeBorderWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderWidth

`func (o *FieldOverride) SetShapeBorderWidth(v int32)`

SetShapeBorderWidth sets ShapeBorderWidth field to given value.

### HasShapeBorderWidth

`func (o *FieldOverride) HasShapeBorderWidth() bool`

HasShapeBorderWidth returns a boolean if a field has been set.

### GetShapeBorderColor

`func (o *FieldOverride) GetShapeBorderColor() string`

GetShapeBorderColor returns the ShapeBorderColor field if non-nil, zero value otherwise.

### GetShapeBorderColorOk

`func (o *FieldOverride) GetShapeBorderColorOk() (*string, bool)`

GetShapeBorderColorOk returns a tuple with the ShapeBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderColor

`func (o *FieldOverride) SetShapeBorderColor(v string)`

SetShapeBorderColor sets ShapeBorderColor field to given value.

### HasShapeBorderColor

`func (o *FieldOverride) HasShapeBorderColor() bool`

HasShapeBorderColor returns a boolean if a field has been set.

### GetShapeFillColor

`func (o *FieldOverride) GetShapeFillColor() string`

GetShapeFillColor returns the ShapeFillColor field if non-nil, zero value otherwise.

### GetShapeFillColorOk

`func (o *FieldOverride) GetShapeFillColorOk() (*string, bool)`

GetShapeFillColorOk returns a tuple with the ShapeFillColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeFillColor

`func (o *FieldOverride) SetShapeFillColor(v string)`

SetShapeFillColor sets ShapeFillColor field to given value.

### HasShapeFillColor

`func (o *FieldOverride) HasShapeFillColor() bool`

HasShapeFillColor returns a boolean if a field has been set.

### SetShapeFillColorNil

`func (o *FieldOverride) SetShapeFillColorNil(b bool)`

 SetShapeFillColorNil sets the value for ShapeFillColor to be an explicit nil

### UnsetShapeFillColor
`func (o *FieldOverride) UnsetShapeFillColor()`

UnsetShapeFillColor ensures that no value is present for ShapeFillColor, not even an explicit nil
### GetCheckCharacter

`func (o *FieldOverride) GetCheckCharacter() CheckCharacter`

GetCheckCharacter returns the CheckCharacter field if non-nil, zero value otherwise.

### GetCheckCharacterOk

`func (o *FieldOverride) GetCheckCharacterOk() (*CheckCharacter, bool)`

GetCheckCharacterOk returns a tuple with the CheckCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCharacter

`func (o *FieldOverride) SetCheckCharacter(v CheckCharacter)`

SetCheckCharacter sets CheckCharacter field to given value.

### HasCheckCharacter

`func (o *FieldOverride) HasCheckCharacter() bool`

HasCheckCharacter returns a boolean if a field has been set.

### GetCheckColor

`func (o *FieldOverride) GetCheckColor() string`

GetCheckColor returns the CheckColor field if non-nil, zero value otherwise.

### GetCheckColorOk

`func (o *FieldOverride) GetCheckColorOk() (*string, bool)`

GetCheckColorOk returns a tuple with the CheckColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckColor

`func (o *FieldOverride) SetCheckColor(v string)`

SetCheckColor sets CheckColor field to given value.

### HasCheckColor

`func (o *FieldOverride) HasCheckColor() bool`

HasCheckColor returns a boolean if a field has been set.

### SetCheckColorNil

`func (o *FieldOverride) SetCheckColorNil(b bool)`

 SetCheckColorNil sets the value for CheckColor to be an explicit nil

### UnsetCheckColor
`func (o *FieldOverride) UnsetCheckColor()`

UnsetCheckColor ensures that no value is present for CheckColor, not even an explicit nil
### GetQrcodeColor

`func (o *FieldOverride) GetQrcodeColor() string`

GetQrcodeColor returns the QrcodeColor field if non-nil, zero value otherwise.

### GetQrcodeColorOk

`func (o *FieldOverride) GetQrcodeColorOk() (*string, bool)`

GetQrcodeColorOk returns a tuple with the QrcodeColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrcodeColor

`func (o *FieldOverride) SetQrcodeColor(v string)`

SetQrcodeColor sets QrcodeColor field to given value.

### HasQrcodeColor

`func (o *FieldOverride) HasQrcodeColor() bool`

HasQrcodeColor returns a boolean if a field has been set.

### SetQrcodeColorNil

`func (o *FieldOverride) SetQrcodeColorNil(b bool)`

 SetQrcodeColorNil sets the value for QrcodeColor to be an explicit nil

### UnsetQrcodeColor
`func (o *FieldOverride) UnsetQrcodeColor()`

UnsetQrcodeColor ensures that no value is present for QrcodeColor, not even an explicit nil
### GetTrueText

`func (o *FieldOverride) GetTrueText() string`

GetTrueText returns the TrueText field if non-nil, zero value otherwise.

### GetTrueTextOk

`func (o *FieldOverride) GetTrueTextOk() (*string, bool)`

GetTrueTextOk returns a tuple with the TrueText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueText

`func (o *FieldOverride) SetTrueText(v string)`

SetTrueText sets TrueText field to given value.

### HasTrueText

`func (o *FieldOverride) HasTrueText() bool`

HasTrueText returns a boolean if a field has been set.

### SetTrueTextNil

`func (o *FieldOverride) SetTrueTextNil(b bool)`

 SetTrueTextNil sets the value for TrueText to be an explicit nil

### UnsetTrueText
`func (o *FieldOverride) UnsetTrueText()`

UnsetTrueText ensures that no value is present for TrueText, not even an explicit nil
### GetFalseText

`func (o *FieldOverride) GetFalseText() string`

GetFalseText returns the FalseText field if non-nil, zero value otherwise.

### GetFalseTextOk

`func (o *FieldOverride) GetFalseTextOk() (*string, bool)`

GetFalseTextOk returns a tuple with the FalseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseText

`func (o *FieldOverride) SetFalseText(v string)`

SetFalseText sets FalseText field to given value.

### HasFalseText

`func (o *FieldOverride) HasFalseText() bool`

HasFalseText returns a boolean if a field has been set.

### SetFalseTextNil

`func (o *FieldOverride) SetFalseTextNil(b bool)`

 SetFalseTextNil sets the value for FalseText to be an explicit nil

### UnsetFalseText
`func (o *FieldOverride) UnsetFalseText()`

UnsetFalseText ensures that no value is present for FalseText, not even an explicit nil
### GetCondition

`func (o *FieldOverride) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *FieldOverride) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *FieldOverride) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *FieldOverride) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *FieldOverride) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *FieldOverride) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetStringConditionType

`func (o *FieldOverride) GetStringConditionType() StringConditionType`

GetStringConditionType returns the StringConditionType field if non-nil, zero value otherwise.

### GetStringConditionTypeOk

`func (o *FieldOverride) GetStringConditionTypeOk() (*StringConditionType, bool)`

GetStringConditionTypeOk returns a tuple with the StringConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringConditionType

`func (o *FieldOverride) SetStringConditionType(v StringConditionType)`

SetStringConditionType sets StringConditionType field to given value.

### HasStringConditionType

`func (o *FieldOverride) HasStringConditionType() bool`

HasStringConditionType returns a boolean if a field has been set.

### GetNumberConditionType

`func (o *FieldOverride) GetNumberConditionType() NumberConditionType`

GetNumberConditionType returns the NumberConditionType field if non-nil, zero value otherwise.

### GetNumberConditionTypeOk

`func (o *FieldOverride) GetNumberConditionTypeOk() (*NumberConditionType, bool)`

GetNumberConditionTypeOk returns a tuple with the NumberConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionType

`func (o *FieldOverride) SetNumberConditionType(v NumberConditionType)`

SetNumberConditionType sets NumberConditionType field to given value.

### HasNumberConditionType

`func (o *FieldOverride) HasNumberConditionType() bool`

HasNumberConditionType returns a boolean if a field has been set.

### GetNumberConditionRangeMin

`func (o *FieldOverride) GetNumberConditionRangeMin() float32`

GetNumberConditionRangeMin returns the NumberConditionRangeMin field if non-nil, zero value otherwise.

### GetNumberConditionRangeMinOk

`func (o *FieldOverride) GetNumberConditionRangeMinOk() (*float32, bool)`

GetNumberConditionRangeMinOk returns a tuple with the NumberConditionRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeMin

`func (o *FieldOverride) SetNumberConditionRangeMin(v float32)`

SetNumberConditionRangeMin sets NumberConditionRangeMin field to given value.

### HasNumberConditionRangeMin

`func (o *FieldOverride) HasNumberConditionRangeMin() bool`

HasNumberConditionRangeMin returns a boolean if a field has been set.

### SetNumberConditionRangeMinNil

`func (o *FieldOverride) SetNumberConditionRangeMinNil(b bool)`

 SetNumberConditionRangeMinNil sets the value for NumberConditionRangeMin to be an explicit nil

### UnsetNumberConditionRangeMin
`func (o *FieldOverride) UnsetNumberConditionRangeMin()`

UnsetNumberConditionRangeMin ensures that no value is present for NumberConditionRangeMin, not even an explicit nil
### GetNumberConditionRangeMax

`func (o *FieldOverride) GetNumberConditionRangeMax() float32`

GetNumberConditionRangeMax returns the NumberConditionRangeMax field if non-nil, zero value otherwise.

### GetNumberConditionRangeMaxOk

`func (o *FieldOverride) GetNumberConditionRangeMaxOk() (*float32, bool)`

GetNumberConditionRangeMaxOk returns a tuple with the NumberConditionRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeMax

`func (o *FieldOverride) SetNumberConditionRangeMax(v float32)`

SetNumberConditionRangeMax sets NumberConditionRangeMax field to given value.

### HasNumberConditionRangeMax

`func (o *FieldOverride) HasNumberConditionRangeMax() bool`

HasNumberConditionRangeMax returns a boolean if a field has been set.

### SetNumberConditionRangeMaxNil

`func (o *FieldOverride) SetNumberConditionRangeMaxNil(b bool)`

 SetNumberConditionRangeMaxNil sets the value for NumberConditionRangeMax to be an explicit nil

### UnsetNumberConditionRangeMax
`func (o *FieldOverride) UnsetNumberConditionRangeMax()`

UnsetNumberConditionRangeMax ensures that no value is present for NumberConditionRangeMax, not even an explicit nil
### GetNumberConditionRangeExclusiveMin

`func (o *FieldOverride) GetNumberConditionRangeExclusiveMin() bool`

GetNumberConditionRangeExclusiveMin returns the NumberConditionRangeExclusiveMin field if non-nil, zero value otherwise.

### GetNumberConditionRangeExclusiveMinOk

`func (o *FieldOverride) GetNumberConditionRangeExclusiveMinOk() (*bool, bool)`

GetNumberConditionRangeExclusiveMinOk returns a tuple with the NumberConditionRangeExclusiveMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeExclusiveMin

`func (o *FieldOverride) SetNumberConditionRangeExclusiveMin(v bool)`

SetNumberConditionRangeExclusiveMin sets NumberConditionRangeExclusiveMin field to given value.

### HasNumberConditionRangeExclusiveMin

`func (o *FieldOverride) HasNumberConditionRangeExclusiveMin() bool`

HasNumberConditionRangeExclusiveMin returns a boolean if a field has been set.

### GetNumberConditionRangeExclusiveMax

`func (o *FieldOverride) GetNumberConditionRangeExclusiveMax() bool`

GetNumberConditionRangeExclusiveMax returns the NumberConditionRangeExclusiveMax field if non-nil, zero value otherwise.

### GetNumberConditionRangeExclusiveMaxOk

`func (o *FieldOverride) GetNumberConditionRangeExclusiveMaxOk() (*bool, bool)`

GetNumberConditionRangeExclusiveMaxOk returns a tuple with the NumberConditionRangeExclusiveMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConditionRangeExclusiveMax

`func (o *FieldOverride) SetNumberConditionRangeExclusiveMax(v bool)`

SetNumberConditionRangeExclusiveMax sets NumberConditionRangeExclusiveMax field to given value.

### HasNumberConditionRangeExclusiveMax

`func (o *FieldOverride) HasNumberConditionRangeExclusiveMax() bool`

HasNumberConditionRangeExclusiveMax returns a boolean if a field has been set.

### GetInvertBooleanCondition

`func (o *FieldOverride) GetInvertBooleanCondition() bool`

GetInvertBooleanCondition returns the InvertBooleanCondition field if non-nil, zero value otherwise.

### GetInvertBooleanConditionOk

`func (o *FieldOverride) GetInvertBooleanConditionOk() (*bool, bool)`

GetInvertBooleanConditionOk returns a tuple with the InvertBooleanCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertBooleanCondition

`func (o *FieldOverride) SetInvertBooleanCondition(v bool)`

SetInvertBooleanCondition sets InvertBooleanCondition field to given value.

### HasInvertBooleanCondition

`func (o *FieldOverride) HasInvertBooleanCondition() bool`

HasInvertBooleanCondition returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *FieldOverride) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *FieldOverride) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *FieldOverride) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *FieldOverride) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### SetDecimalPlacesNil

`func (o *FieldOverride) SetDecimalPlacesNil(b bool)`

 SetDecimalPlacesNil sets the value for DecimalPlaces to be an explicit nil

### UnsetDecimalPlaces
`func (o *FieldOverride) UnsetDecimalPlaces()`

UnsetDecimalPlaces ensures that no value is present for DecimalPlaces, not even an explicit nil
### GetImageScaleType

`func (o *FieldOverride) GetImageScaleType() ImageScaleType`

GetImageScaleType returns the ImageScaleType field if non-nil, zero value otherwise.

### GetImageScaleTypeOk

`func (o *FieldOverride) GetImageScaleTypeOk() (*ImageScaleType, bool)`

GetImageScaleTypeOk returns a tuple with the ImageScaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScaleType

`func (o *FieldOverride) SetImageScaleType(v ImageScaleType)`

SetImageScaleType sets ImageScaleType field to given value.

### HasImageScaleType

`func (o *FieldOverride) HasImageScaleType() bool`

HasImageScaleType returns a boolean if a field has been set.

### GetImageGravity

`func (o *FieldOverride) GetImageGravity() ImageGravity`

GetImageGravity returns the ImageGravity field if non-nil, zero value otherwise.

### GetImageGravityOk

`func (o *FieldOverride) GetImageGravityOk() (*ImageGravity, bool)`

GetImageGravityOk returns a tuple with the ImageGravity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageGravity

`func (o *FieldOverride) SetImageGravity(v ImageGravity)`

SetImageGravity sets ImageGravity field to given value.

### HasImageGravity

`func (o *FieldOverride) HasImageGravity() bool`

HasImageGravity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


