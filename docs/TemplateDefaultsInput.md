# TemplateDefaultsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typeface** | Pointer to [**NullableTypeface**](Typeface.md) | The font of text fields. | [optional] 
**FontSize** | Pointer to **NullableInt32** | The font size in points, a whole number. | [optional] 
**CharacterSpacing** | Pointer to **NullableFloat32** | Extra space between characters, in points (0 to 50). | [optional] 
**Color** | Pointer to **NullableString** | The text color, as 6 hex digits without `#`, such as `444444`. | [optional] 
**BackgroundColor** | Pointer to **NullableString** | The field's background color, as 6 hex digits. | [optional] 
**Bold** | Pointer to **bool** | Print text in bold. | [optional] 
**Uppercase** | Pointer to **bool** | Print text in capital letters. | [optional] 
**Alignment** | Pointer to [**TextAlignment**](TextAlignment.md) | How text is aligned in its field. `justify` lines up both edges of every line of multiline text except the last of each paragraph. | [optional] 
**Opacity** | Pointer to **float32** | The opacity of the field's content, from 0 (invisible) to 1. | [optional] 
**DateFormat** | Pointer to **NullableString** | How date fields are printed, as a strftime format such as `%m/%d/%Y`. | [optional] 
**DateTimeFormat** | Pointer to **NullableString** | How date-time fields are printed, as a strftime format such as `%m/%d/%Y %H:%M:%S`. | [optional] 
**CheckColor** | Pointer to **NullableString** | The color of check marks, as 6 hex digits. | [optional] 
**CheckCharacter** | Pointer to [**NullableCheckCharacter**](CheckCharacter.md) | The character that a ticked check box prints. | [optional] 
**CheckBackgroundColor** | Pointer to **NullableString** | The check box's background color, as 6 hex digits. | [optional] 
**RequiredFieldBackgroundColor** | Pointer to **NullableString** | The background color that highlights required fields on visual forms and signature request forms, as 6 hex digits. PDFs don't print it. | [optional] 
**TrueText** | Pointer to **NullableString** | The text a boolean field prints for `true`. | [optional] 
**FalseText** | Pointer to **NullableString** | The text a boolean field prints for `false`. | [optional] 
**ShapeFillColor** | Pointer to **NullableString** | The fill color of shapes, as 6 hex digits. | [optional] 
**ShapeBorderColor** | Pointer to **NullableString** | The border color of shapes, as 6 hex digits. | [optional] 
**ShapeBorderWidth** | Pointer to **int32** | The border width of shapes, in whole points (at least 1). | [optional] 

## Methods

### NewTemplateDefaultsInput

`func NewTemplateDefaultsInput() *TemplateDefaultsInput`

NewTemplateDefaultsInput instantiates a new TemplateDefaultsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDefaultsInputWithDefaults

`func NewTemplateDefaultsInputWithDefaults() *TemplateDefaultsInput`

NewTemplateDefaultsInputWithDefaults instantiates a new TemplateDefaultsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeface

`func (o *TemplateDefaultsInput) GetTypeface() Typeface`

GetTypeface returns the Typeface field if non-nil, zero value otherwise.

### GetTypefaceOk

`func (o *TemplateDefaultsInput) GetTypefaceOk() (*Typeface, bool)`

GetTypefaceOk returns a tuple with the Typeface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeface

`func (o *TemplateDefaultsInput) SetTypeface(v Typeface)`

SetTypeface sets Typeface field to given value.

### HasTypeface

`func (o *TemplateDefaultsInput) HasTypeface() bool`

HasTypeface returns a boolean if a field has been set.

### SetTypefaceNil

`func (o *TemplateDefaultsInput) SetTypefaceNil(b bool)`

 SetTypefaceNil sets the value for Typeface to be an explicit nil

### UnsetTypeface
`func (o *TemplateDefaultsInput) UnsetTypeface()`

UnsetTypeface ensures that no value is present for Typeface, not even an explicit nil
### GetFontSize

`func (o *TemplateDefaultsInput) GetFontSize() int32`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *TemplateDefaultsInput) GetFontSizeOk() (*int32, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *TemplateDefaultsInput) SetFontSize(v int32)`

SetFontSize sets FontSize field to given value.

### HasFontSize

`func (o *TemplateDefaultsInput) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.

### SetFontSizeNil

`func (o *TemplateDefaultsInput) SetFontSizeNil(b bool)`

 SetFontSizeNil sets the value for FontSize to be an explicit nil

### UnsetFontSize
`func (o *TemplateDefaultsInput) UnsetFontSize()`

UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil
### GetCharacterSpacing

`func (o *TemplateDefaultsInput) GetCharacterSpacing() float32`

GetCharacterSpacing returns the CharacterSpacing field if non-nil, zero value otherwise.

### GetCharacterSpacingOk

`func (o *TemplateDefaultsInput) GetCharacterSpacingOk() (*float32, bool)`

GetCharacterSpacingOk returns a tuple with the CharacterSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterSpacing

`func (o *TemplateDefaultsInput) SetCharacterSpacing(v float32)`

SetCharacterSpacing sets CharacterSpacing field to given value.

### HasCharacterSpacing

`func (o *TemplateDefaultsInput) HasCharacterSpacing() bool`

HasCharacterSpacing returns a boolean if a field has been set.

### SetCharacterSpacingNil

`func (o *TemplateDefaultsInput) SetCharacterSpacingNil(b bool)`

 SetCharacterSpacingNil sets the value for CharacterSpacing to be an explicit nil

### UnsetCharacterSpacing
`func (o *TemplateDefaultsInput) UnsetCharacterSpacing()`

UnsetCharacterSpacing ensures that no value is present for CharacterSpacing, not even an explicit nil
### GetColor

`func (o *TemplateDefaultsInput) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TemplateDefaultsInput) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TemplateDefaultsInput) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TemplateDefaultsInput) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *TemplateDefaultsInput) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *TemplateDefaultsInput) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetBackgroundColor

`func (o *TemplateDefaultsInput) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *TemplateDefaultsInput) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *TemplateDefaultsInput) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *TemplateDefaultsInput) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *TemplateDefaultsInput) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *TemplateDefaultsInput) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetBold

`func (o *TemplateDefaultsInput) GetBold() bool`

GetBold returns the Bold field if non-nil, zero value otherwise.

### GetBoldOk

`func (o *TemplateDefaultsInput) GetBoldOk() (*bool, bool)`

GetBoldOk returns a tuple with the Bold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBold

`func (o *TemplateDefaultsInput) SetBold(v bool)`

SetBold sets Bold field to given value.

### HasBold

`func (o *TemplateDefaultsInput) HasBold() bool`

HasBold returns a boolean if a field has been set.

### GetUppercase

`func (o *TemplateDefaultsInput) GetUppercase() bool`

GetUppercase returns the Uppercase field if non-nil, zero value otherwise.

### GetUppercaseOk

`func (o *TemplateDefaultsInput) GetUppercaseOk() (*bool, bool)`

GetUppercaseOk returns a tuple with the Uppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUppercase

`func (o *TemplateDefaultsInput) SetUppercase(v bool)`

SetUppercase sets Uppercase field to given value.

### HasUppercase

`func (o *TemplateDefaultsInput) HasUppercase() bool`

HasUppercase returns a boolean if a field has been set.

### GetAlignment

`func (o *TemplateDefaultsInput) GetAlignment() TextAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *TemplateDefaultsInput) GetAlignmentOk() (*TextAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *TemplateDefaultsInput) SetAlignment(v TextAlignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *TemplateDefaultsInput) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetOpacity

`func (o *TemplateDefaultsInput) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *TemplateDefaultsInput) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *TemplateDefaultsInput) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *TemplateDefaultsInput) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetDateFormat

`func (o *TemplateDefaultsInput) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *TemplateDefaultsInput) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *TemplateDefaultsInput) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *TemplateDefaultsInput) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### SetDateFormatNil

`func (o *TemplateDefaultsInput) SetDateFormatNil(b bool)`

 SetDateFormatNil sets the value for DateFormat to be an explicit nil

### UnsetDateFormat
`func (o *TemplateDefaultsInput) UnsetDateFormat()`

UnsetDateFormat ensures that no value is present for DateFormat, not even an explicit nil
### GetDateTimeFormat

`func (o *TemplateDefaultsInput) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *TemplateDefaultsInput) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *TemplateDefaultsInput) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *TemplateDefaultsInput) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### SetDateTimeFormatNil

`func (o *TemplateDefaultsInput) SetDateTimeFormatNil(b bool)`

 SetDateTimeFormatNil sets the value for DateTimeFormat to be an explicit nil

### UnsetDateTimeFormat
`func (o *TemplateDefaultsInput) UnsetDateTimeFormat()`

UnsetDateTimeFormat ensures that no value is present for DateTimeFormat, not even an explicit nil
### GetCheckColor

`func (o *TemplateDefaultsInput) GetCheckColor() string`

GetCheckColor returns the CheckColor field if non-nil, zero value otherwise.

### GetCheckColorOk

`func (o *TemplateDefaultsInput) GetCheckColorOk() (*string, bool)`

GetCheckColorOk returns a tuple with the CheckColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckColor

`func (o *TemplateDefaultsInput) SetCheckColor(v string)`

SetCheckColor sets CheckColor field to given value.

### HasCheckColor

`func (o *TemplateDefaultsInput) HasCheckColor() bool`

HasCheckColor returns a boolean if a field has been set.

### SetCheckColorNil

`func (o *TemplateDefaultsInput) SetCheckColorNil(b bool)`

 SetCheckColorNil sets the value for CheckColor to be an explicit nil

### UnsetCheckColor
`func (o *TemplateDefaultsInput) UnsetCheckColor()`

UnsetCheckColor ensures that no value is present for CheckColor, not even an explicit nil
### GetCheckCharacter

`func (o *TemplateDefaultsInput) GetCheckCharacter() CheckCharacter`

GetCheckCharacter returns the CheckCharacter field if non-nil, zero value otherwise.

### GetCheckCharacterOk

`func (o *TemplateDefaultsInput) GetCheckCharacterOk() (*CheckCharacter, bool)`

GetCheckCharacterOk returns a tuple with the CheckCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCharacter

`func (o *TemplateDefaultsInput) SetCheckCharacter(v CheckCharacter)`

SetCheckCharacter sets CheckCharacter field to given value.

### HasCheckCharacter

`func (o *TemplateDefaultsInput) HasCheckCharacter() bool`

HasCheckCharacter returns a boolean if a field has been set.

### SetCheckCharacterNil

`func (o *TemplateDefaultsInput) SetCheckCharacterNil(b bool)`

 SetCheckCharacterNil sets the value for CheckCharacter to be an explicit nil

### UnsetCheckCharacter
`func (o *TemplateDefaultsInput) UnsetCheckCharacter()`

UnsetCheckCharacter ensures that no value is present for CheckCharacter, not even an explicit nil
### GetCheckBackgroundColor

`func (o *TemplateDefaultsInput) GetCheckBackgroundColor() string`

GetCheckBackgroundColor returns the CheckBackgroundColor field if non-nil, zero value otherwise.

### GetCheckBackgroundColorOk

`func (o *TemplateDefaultsInput) GetCheckBackgroundColorOk() (*string, bool)`

GetCheckBackgroundColorOk returns a tuple with the CheckBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckBackgroundColor

`func (o *TemplateDefaultsInput) SetCheckBackgroundColor(v string)`

SetCheckBackgroundColor sets CheckBackgroundColor field to given value.

### HasCheckBackgroundColor

`func (o *TemplateDefaultsInput) HasCheckBackgroundColor() bool`

HasCheckBackgroundColor returns a boolean if a field has been set.

### SetCheckBackgroundColorNil

`func (o *TemplateDefaultsInput) SetCheckBackgroundColorNil(b bool)`

 SetCheckBackgroundColorNil sets the value for CheckBackgroundColor to be an explicit nil

### UnsetCheckBackgroundColor
`func (o *TemplateDefaultsInput) UnsetCheckBackgroundColor()`

UnsetCheckBackgroundColor ensures that no value is present for CheckBackgroundColor, not even an explicit nil
### GetRequiredFieldBackgroundColor

`func (o *TemplateDefaultsInput) GetRequiredFieldBackgroundColor() string`

GetRequiredFieldBackgroundColor returns the RequiredFieldBackgroundColor field if non-nil, zero value otherwise.

### GetRequiredFieldBackgroundColorOk

`func (o *TemplateDefaultsInput) GetRequiredFieldBackgroundColorOk() (*string, bool)`

GetRequiredFieldBackgroundColorOk returns a tuple with the RequiredFieldBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFieldBackgroundColor

`func (o *TemplateDefaultsInput) SetRequiredFieldBackgroundColor(v string)`

SetRequiredFieldBackgroundColor sets RequiredFieldBackgroundColor field to given value.

### HasRequiredFieldBackgroundColor

`func (o *TemplateDefaultsInput) HasRequiredFieldBackgroundColor() bool`

HasRequiredFieldBackgroundColor returns a boolean if a field has been set.

### SetRequiredFieldBackgroundColorNil

`func (o *TemplateDefaultsInput) SetRequiredFieldBackgroundColorNil(b bool)`

 SetRequiredFieldBackgroundColorNil sets the value for RequiredFieldBackgroundColor to be an explicit nil

### UnsetRequiredFieldBackgroundColor
`func (o *TemplateDefaultsInput) UnsetRequiredFieldBackgroundColor()`

UnsetRequiredFieldBackgroundColor ensures that no value is present for RequiredFieldBackgroundColor, not even an explicit nil
### GetTrueText

`func (o *TemplateDefaultsInput) GetTrueText() string`

GetTrueText returns the TrueText field if non-nil, zero value otherwise.

### GetTrueTextOk

`func (o *TemplateDefaultsInput) GetTrueTextOk() (*string, bool)`

GetTrueTextOk returns a tuple with the TrueText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueText

`func (o *TemplateDefaultsInput) SetTrueText(v string)`

SetTrueText sets TrueText field to given value.

### HasTrueText

`func (o *TemplateDefaultsInput) HasTrueText() bool`

HasTrueText returns a boolean if a field has been set.

### SetTrueTextNil

`func (o *TemplateDefaultsInput) SetTrueTextNil(b bool)`

 SetTrueTextNil sets the value for TrueText to be an explicit nil

### UnsetTrueText
`func (o *TemplateDefaultsInput) UnsetTrueText()`

UnsetTrueText ensures that no value is present for TrueText, not even an explicit nil
### GetFalseText

`func (o *TemplateDefaultsInput) GetFalseText() string`

GetFalseText returns the FalseText field if non-nil, zero value otherwise.

### GetFalseTextOk

`func (o *TemplateDefaultsInput) GetFalseTextOk() (*string, bool)`

GetFalseTextOk returns a tuple with the FalseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseText

`func (o *TemplateDefaultsInput) SetFalseText(v string)`

SetFalseText sets FalseText field to given value.

### HasFalseText

`func (o *TemplateDefaultsInput) HasFalseText() bool`

HasFalseText returns a boolean if a field has been set.

### SetFalseTextNil

`func (o *TemplateDefaultsInput) SetFalseTextNil(b bool)`

 SetFalseTextNil sets the value for FalseText to be an explicit nil

### UnsetFalseText
`func (o *TemplateDefaultsInput) UnsetFalseText()`

UnsetFalseText ensures that no value is present for FalseText, not even an explicit nil
### GetShapeFillColor

`func (o *TemplateDefaultsInput) GetShapeFillColor() string`

GetShapeFillColor returns the ShapeFillColor field if non-nil, zero value otherwise.

### GetShapeFillColorOk

`func (o *TemplateDefaultsInput) GetShapeFillColorOk() (*string, bool)`

GetShapeFillColorOk returns a tuple with the ShapeFillColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeFillColor

`func (o *TemplateDefaultsInput) SetShapeFillColor(v string)`

SetShapeFillColor sets ShapeFillColor field to given value.

### HasShapeFillColor

`func (o *TemplateDefaultsInput) HasShapeFillColor() bool`

HasShapeFillColor returns a boolean if a field has been set.

### SetShapeFillColorNil

`func (o *TemplateDefaultsInput) SetShapeFillColorNil(b bool)`

 SetShapeFillColorNil sets the value for ShapeFillColor to be an explicit nil

### UnsetShapeFillColor
`func (o *TemplateDefaultsInput) UnsetShapeFillColor()`

UnsetShapeFillColor ensures that no value is present for ShapeFillColor, not even an explicit nil
### GetShapeBorderColor

`func (o *TemplateDefaultsInput) GetShapeBorderColor() string`

GetShapeBorderColor returns the ShapeBorderColor field if non-nil, zero value otherwise.

### GetShapeBorderColorOk

`func (o *TemplateDefaultsInput) GetShapeBorderColorOk() (*string, bool)`

GetShapeBorderColorOk returns a tuple with the ShapeBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderColor

`func (o *TemplateDefaultsInput) SetShapeBorderColor(v string)`

SetShapeBorderColor sets ShapeBorderColor field to given value.

### HasShapeBorderColor

`func (o *TemplateDefaultsInput) HasShapeBorderColor() bool`

HasShapeBorderColor returns a boolean if a field has been set.

### SetShapeBorderColorNil

`func (o *TemplateDefaultsInput) SetShapeBorderColorNil(b bool)`

 SetShapeBorderColorNil sets the value for ShapeBorderColor to be an explicit nil

### UnsetShapeBorderColor
`func (o *TemplateDefaultsInput) UnsetShapeBorderColor()`

UnsetShapeBorderColor ensures that no value is present for ShapeBorderColor, not even an explicit nil
### GetShapeBorderWidth

`func (o *TemplateDefaultsInput) GetShapeBorderWidth() int32`

GetShapeBorderWidth returns the ShapeBorderWidth field if non-nil, zero value otherwise.

### GetShapeBorderWidthOk

`func (o *TemplateDefaultsInput) GetShapeBorderWidthOk() (*int32, bool)`

GetShapeBorderWidthOk returns a tuple with the ShapeBorderWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderWidth

`func (o *TemplateDefaultsInput) SetShapeBorderWidth(v int32)`

SetShapeBorderWidth sets ShapeBorderWidth field to given value.

### HasShapeBorderWidth

`func (o *TemplateDefaultsInput) HasShapeBorderWidth() bool`

HasShapeBorderWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


