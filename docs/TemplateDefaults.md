# TemplateDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typeface** | Pointer to [**NullableTypeface**](Typeface.md) | The font of text fields. | [optional] 
**FontSize** | Pointer to **NullableFloat32** | The font size in points, a whole number. | [optional] 
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
**ShapeBorderWidth** | Pointer to **float32** | The border width of shapes, in whole points (at least 1). | [optional] 

## Methods

### NewTemplateDefaults

`func NewTemplateDefaults() *TemplateDefaults`

NewTemplateDefaults instantiates a new TemplateDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDefaultsWithDefaults

`func NewTemplateDefaultsWithDefaults() *TemplateDefaults`

NewTemplateDefaultsWithDefaults instantiates a new TemplateDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeface

`func (o *TemplateDefaults) GetTypeface() Typeface`

GetTypeface returns the Typeface field if non-nil, zero value otherwise.

### GetTypefaceOk

`func (o *TemplateDefaults) GetTypefaceOk() (*Typeface, bool)`

GetTypefaceOk returns a tuple with the Typeface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeface

`func (o *TemplateDefaults) SetTypeface(v Typeface)`

SetTypeface sets Typeface field to given value.

### HasTypeface

`func (o *TemplateDefaults) HasTypeface() bool`

HasTypeface returns a boolean if a field has been set.

### SetTypefaceNil

`func (o *TemplateDefaults) SetTypefaceNil(b bool)`

 SetTypefaceNil sets the value for Typeface to be an explicit nil

### UnsetTypeface
`func (o *TemplateDefaults) UnsetTypeface()`

UnsetTypeface ensures that no value is present for Typeface, not even an explicit nil
### GetFontSize

`func (o *TemplateDefaults) GetFontSize() float32`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *TemplateDefaults) GetFontSizeOk() (*float32, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *TemplateDefaults) SetFontSize(v float32)`

SetFontSize sets FontSize field to given value.

### HasFontSize

`func (o *TemplateDefaults) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.

### SetFontSizeNil

`func (o *TemplateDefaults) SetFontSizeNil(b bool)`

 SetFontSizeNil sets the value for FontSize to be an explicit nil

### UnsetFontSize
`func (o *TemplateDefaults) UnsetFontSize()`

UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil
### GetCharacterSpacing

`func (o *TemplateDefaults) GetCharacterSpacing() float32`

GetCharacterSpacing returns the CharacterSpacing field if non-nil, zero value otherwise.

### GetCharacterSpacingOk

`func (o *TemplateDefaults) GetCharacterSpacingOk() (*float32, bool)`

GetCharacterSpacingOk returns a tuple with the CharacterSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterSpacing

`func (o *TemplateDefaults) SetCharacterSpacing(v float32)`

SetCharacterSpacing sets CharacterSpacing field to given value.

### HasCharacterSpacing

`func (o *TemplateDefaults) HasCharacterSpacing() bool`

HasCharacterSpacing returns a boolean if a field has been set.

### SetCharacterSpacingNil

`func (o *TemplateDefaults) SetCharacterSpacingNil(b bool)`

 SetCharacterSpacingNil sets the value for CharacterSpacing to be an explicit nil

### UnsetCharacterSpacing
`func (o *TemplateDefaults) UnsetCharacterSpacing()`

UnsetCharacterSpacing ensures that no value is present for CharacterSpacing, not even an explicit nil
### GetColor

`func (o *TemplateDefaults) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TemplateDefaults) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TemplateDefaults) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TemplateDefaults) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *TemplateDefaults) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *TemplateDefaults) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetBackgroundColor

`func (o *TemplateDefaults) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *TemplateDefaults) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *TemplateDefaults) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *TemplateDefaults) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *TemplateDefaults) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *TemplateDefaults) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetBold

`func (o *TemplateDefaults) GetBold() bool`

GetBold returns the Bold field if non-nil, zero value otherwise.

### GetBoldOk

`func (o *TemplateDefaults) GetBoldOk() (*bool, bool)`

GetBoldOk returns a tuple with the Bold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBold

`func (o *TemplateDefaults) SetBold(v bool)`

SetBold sets Bold field to given value.

### HasBold

`func (o *TemplateDefaults) HasBold() bool`

HasBold returns a boolean if a field has been set.

### GetUppercase

`func (o *TemplateDefaults) GetUppercase() bool`

GetUppercase returns the Uppercase field if non-nil, zero value otherwise.

### GetUppercaseOk

`func (o *TemplateDefaults) GetUppercaseOk() (*bool, bool)`

GetUppercaseOk returns a tuple with the Uppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUppercase

`func (o *TemplateDefaults) SetUppercase(v bool)`

SetUppercase sets Uppercase field to given value.

### HasUppercase

`func (o *TemplateDefaults) HasUppercase() bool`

HasUppercase returns a boolean if a field has been set.

### GetAlignment

`func (o *TemplateDefaults) GetAlignment() TextAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *TemplateDefaults) GetAlignmentOk() (*TextAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *TemplateDefaults) SetAlignment(v TextAlignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *TemplateDefaults) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetOpacity

`func (o *TemplateDefaults) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *TemplateDefaults) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *TemplateDefaults) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *TemplateDefaults) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetDateFormat

`func (o *TemplateDefaults) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *TemplateDefaults) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *TemplateDefaults) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *TemplateDefaults) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### SetDateFormatNil

`func (o *TemplateDefaults) SetDateFormatNil(b bool)`

 SetDateFormatNil sets the value for DateFormat to be an explicit nil

### UnsetDateFormat
`func (o *TemplateDefaults) UnsetDateFormat()`

UnsetDateFormat ensures that no value is present for DateFormat, not even an explicit nil
### GetDateTimeFormat

`func (o *TemplateDefaults) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *TemplateDefaults) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *TemplateDefaults) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *TemplateDefaults) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### SetDateTimeFormatNil

`func (o *TemplateDefaults) SetDateTimeFormatNil(b bool)`

 SetDateTimeFormatNil sets the value for DateTimeFormat to be an explicit nil

### UnsetDateTimeFormat
`func (o *TemplateDefaults) UnsetDateTimeFormat()`

UnsetDateTimeFormat ensures that no value is present for DateTimeFormat, not even an explicit nil
### GetCheckColor

`func (o *TemplateDefaults) GetCheckColor() string`

GetCheckColor returns the CheckColor field if non-nil, zero value otherwise.

### GetCheckColorOk

`func (o *TemplateDefaults) GetCheckColorOk() (*string, bool)`

GetCheckColorOk returns a tuple with the CheckColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckColor

`func (o *TemplateDefaults) SetCheckColor(v string)`

SetCheckColor sets CheckColor field to given value.

### HasCheckColor

`func (o *TemplateDefaults) HasCheckColor() bool`

HasCheckColor returns a boolean if a field has been set.

### SetCheckColorNil

`func (o *TemplateDefaults) SetCheckColorNil(b bool)`

 SetCheckColorNil sets the value for CheckColor to be an explicit nil

### UnsetCheckColor
`func (o *TemplateDefaults) UnsetCheckColor()`

UnsetCheckColor ensures that no value is present for CheckColor, not even an explicit nil
### GetCheckCharacter

`func (o *TemplateDefaults) GetCheckCharacter() CheckCharacter`

GetCheckCharacter returns the CheckCharacter field if non-nil, zero value otherwise.

### GetCheckCharacterOk

`func (o *TemplateDefaults) GetCheckCharacterOk() (*CheckCharacter, bool)`

GetCheckCharacterOk returns a tuple with the CheckCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCharacter

`func (o *TemplateDefaults) SetCheckCharacter(v CheckCharacter)`

SetCheckCharacter sets CheckCharacter field to given value.

### HasCheckCharacter

`func (o *TemplateDefaults) HasCheckCharacter() bool`

HasCheckCharacter returns a boolean if a field has been set.

### SetCheckCharacterNil

`func (o *TemplateDefaults) SetCheckCharacterNil(b bool)`

 SetCheckCharacterNil sets the value for CheckCharacter to be an explicit nil

### UnsetCheckCharacter
`func (o *TemplateDefaults) UnsetCheckCharacter()`

UnsetCheckCharacter ensures that no value is present for CheckCharacter, not even an explicit nil
### GetCheckBackgroundColor

`func (o *TemplateDefaults) GetCheckBackgroundColor() string`

GetCheckBackgroundColor returns the CheckBackgroundColor field if non-nil, zero value otherwise.

### GetCheckBackgroundColorOk

`func (o *TemplateDefaults) GetCheckBackgroundColorOk() (*string, bool)`

GetCheckBackgroundColorOk returns a tuple with the CheckBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckBackgroundColor

`func (o *TemplateDefaults) SetCheckBackgroundColor(v string)`

SetCheckBackgroundColor sets CheckBackgroundColor field to given value.

### HasCheckBackgroundColor

`func (o *TemplateDefaults) HasCheckBackgroundColor() bool`

HasCheckBackgroundColor returns a boolean if a field has been set.

### SetCheckBackgroundColorNil

`func (o *TemplateDefaults) SetCheckBackgroundColorNil(b bool)`

 SetCheckBackgroundColorNil sets the value for CheckBackgroundColor to be an explicit nil

### UnsetCheckBackgroundColor
`func (o *TemplateDefaults) UnsetCheckBackgroundColor()`

UnsetCheckBackgroundColor ensures that no value is present for CheckBackgroundColor, not even an explicit nil
### GetRequiredFieldBackgroundColor

`func (o *TemplateDefaults) GetRequiredFieldBackgroundColor() string`

GetRequiredFieldBackgroundColor returns the RequiredFieldBackgroundColor field if non-nil, zero value otherwise.

### GetRequiredFieldBackgroundColorOk

`func (o *TemplateDefaults) GetRequiredFieldBackgroundColorOk() (*string, bool)`

GetRequiredFieldBackgroundColorOk returns a tuple with the RequiredFieldBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFieldBackgroundColor

`func (o *TemplateDefaults) SetRequiredFieldBackgroundColor(v string)`

SetRequiredFieldBackgroundColor sets RequiredFieldBackgroundColor field to given value.

### HasRequiredFieldBackgroundColor

`func (o *TemplateDefaults) HasRequiredFieldBackgroundColor() bool`

HasRequiredFieldBackgroundColor returns a boolean if a field has been set.

### SetRequiredFieldBackgroundColorNil

`func (o *TemplateDefaults) SetRequiredFieldBackgroundColorNil(b bool)`

 SetRequiredFieldBackgroundColorNil sets the value for RequiredFieldBackgroundColor to be an explicit nil

### UnsetRequiredFieldBackgroundColor
`func (o *TemplateDefaults) UnsetRequiredFieldBackgroundColor()`

UnsetRequiredFieldBackgroundColor ensures that no value is present for RequiredFieldBackgroundColor, not even an explicit nil
### GetTrueText

`func (o *TemplateDefaults) GetTrueText() string`

GetTrueText returns the TrueText field if non-nil, zero value otherwise.

### GetTrueTextOk

`func (o *TemplateDefaults) GetTrueTextOk() (*string, bool)`

GetTrueTextOk returns a tuple with the TrueText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueText

`func (o *TemplateDefaults) SetTrueText(v string)`

SetTrueText sets TrueText field to given value.

### HasTrueText

`func (o *TemplateDefaults) HasTrueText() bool`

HasTrueText returns a boolean if a field has been set.

### SetTrueTextNil

`func (o *TemplateDefaults) SetTrueTextNil(b bool)`

 SetTrueTextNil sets the value for TrueText to be an explicit nil

### UnsetTrueText
`func (o *TemplateDefaults) UnsetTrueText()`

UnsetTrueText ensures that no value is present for TrueText, not even an explicit nil
### GetFalseText

`func (o *TemplateDefaults) GetFalseText() string`

GetFalseText returns the FalseText field if non-nil, zero value otherwise.

### GetFalseTextOk

`func (o *TemplateDefaults) GetFalseTextOk() (*string, bool)`

GetFalseTextOk returns a tuple with the FalseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseText

`func (o *TemplateDefaults) SetFalseText(v string)`

SetFalseText sets FalseText field to given value.

### HasFalseText

`func (o *TemplateDefaults) HasFalseText() bool`

HasFalseText returns a boolean if a field has been set.

### SetFalseTextNil

`func (o *TemplateDefaults) SetFalseTextNil(b bool)`

 SetFalseTextNil sets the value for FalseText to be an explicit nil

### UnsetFalseText
`func (o *TemplateDefaults) UnsetFalseText()`

UnsetFalseText ensures that no value is present for FalseText, not even an explicit nil
### GetShapeFillColor

`func (o *TemplateDefaults) GetShapeFillColor() string`

GetShapeFillColor returns the ShapeFillColor field if non-nil, zero value otherwise.

### GetShapeFillColorOk

`func (o *TemplateDefaults) GetShapeFillColorOk() (*string, bool)`

GetShapeFillColorOk returns a tuple with the ShapeFillColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeFillColor

`func (o *TemplateDefaults) SetShapeFillColor(v string)`

SetShapeFillColor sets ShapeFillColor field to given value.

### HasShapeFillColor

`func (o *TemplateDefaults) HasShapeFillColor() bool`

HasShapeFillColor returns a boolean if a field has been set.

### SetShapeFillColorNil

`func (o *TemplateDefaults) SetShapeFillColorNil(b bool)`

 SetShapeFillColorNil sets the value for ShapeFillColor to be an explicit nil

### UnsetShapeFillColor
`func (o *TemplateDefaults) UnsetShapeFillColor()`

UnsetShapeFillColor ensures that no value is present for ShapeFillColor, not even an explicit nil
### GetShapeBorderColor

`func (o *TemplateDefaults) GetShapeBorderColor() string`

GetShapeBorderColor returns the ShapeBorderColor field if non-nil, zero value otherwise.

### GetShapeBorderColorOk

`func (o *TemplateDefaults) GetShapeBorderColorOk() (*string, bool)`

GetShapeBorderColorOk returns a tuple with the ShapeBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderColor

`func (o *TemplateDefaults) SetShapeBorderColor(v string)`

SetShapeBorderColor sets ShapeBorderColor field to given value.

### HasShapeBorderColor

`func (o *TemplateDefaults) HasShapeBorderColor() bool`

HasShapeBorderColor returns a boolean if a field has been set.

### SetShapeBorderColorNil

`func (o *TemplateDefaults) SetShapeBorderColorNil(b bool)`

 SetShapeBorderColorNil sets the value for ShapeBorderColor to be an explicit nil

### UnsetShapeBorderColor
`func (o *TemplateDefaults) UnsetShapeBorderColor()`

UnsetShapeBorderColor ensures that no value is present for ShapeBorderColor, not even an explicit nil
### GetShapeBorderWidth

`func (o *TemplateDefaults) GetShapeBorderWidth() float32`

GetShapeBorderWidth returns the ShapeBorderWidth field if non-nil, zero value otherwise.

### GetShapeBorderWidthOk

`func (o *TemplateDefaults) GetShapeBorderWidthOk() (*float32, bool)`

GetShapeBorderWidthOk returns a tuple with the ShapeBorderWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBorderWidth

`func (o *TemplateDefaults) SetShapeBorderWidth(v float32)`

SetShapeBorderWidth sets ShapeBorderWidth field to given value.

### HasShapeBorderWidth

`func (o *TemplateDefaults) HasShapeBorderWidth() bool`

HasShapeBorderWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


