# HtmlEngineOptionsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to [**PageFormat**](PageFormat.md) | The paper size. `Custom` uses `width` and `height`. | [optional] 
**Width** | Pointer to **string** | The page width for the `Custom` format, with a unit (`8.5in`, `210mm`). | [optional] 
**Height** | Pointer to **string** | The page height for the `Custom` format, with a unit. | [optional] 
**Scale** | Pointer to **float32** | How much to scale the page, such as `0.8`. `1` prints it at its normal size. | [optional] 
**Landscape** | Pointer to **bool** | Print pages in landscape orientation. | [optional] 
**PrintBackground** | Pointer to **bool** | Print background colors and images. | [optional] 
**MarginTop** | Pointer to **string** | The top margin, with a unit, such as `1in` or `20mm`. | [optional] 
**MarginBottom** | Pointer to **string** | The bottom margin, with a unit. | [optional] 
**MarginLeft** | Pointer to **string** | The left margin, with a unit. | [optional] 
**MarginRight** | Pointer to **string** | The right margin, with a unit. | [optional] 

## Methods

### NewHtmlEngineOptionsInput

`func NewHtmlEngineOptionsInput() *HtmlEngineOptionsInput`

NewHtmlEngineOptionsInput instantiates a new HtmlEngineOptionsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHtmlEngineOptionsInputWithDefaults

`func NewHtmlEngineOptionsInputWithDefaults() *HtmlEngineOptionsInput`

NewHtmlEngineOptionsInputWithDefaults instantiates a new HtmlEngineOptionsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *HtmlEngineOptionsInput) GetFormat() PageFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *HtmlEngineOptionsInput) GetFormatOk() (*PageFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *HtmlEngineOptionsInput) SetFormat(v PageFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *HtmlEngineOptionsInput) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetWidth

`func (o *HtmlEngineOptionsInput) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *HtmlEngineOptionsInput) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *HtmlEngineOptionsInput) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *HtmlEngineOptionsInput) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *HtmlEngineOptionsInput) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *HtmlEngineOptionsInput) GetHeightOk() (*string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *HtmlEngineOptionsInput) SetHeight(v string)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *HtmlEngineOptionsInput) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetScale

`func (o *HtmlEngineOptionsInput) GetScale() float32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *HtmlEngineOptionsInput) GetScaleOk() (*float32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *HtmlEngineOptionsInput) SetScale(v float32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *HtmlEngineOptionsInput) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetLandscape

`func (o *HtmlEngineOptionsInput) GetLandscape() bool`

GetLandscape returns the Landscape field if non-nil, zero value otherwise.

### GetLandscapeOk

`func (o *HtmlEngineOptionsInput) GetLandscapeOk() (*bool, bool)`

GetLandscapeOk returns a tuple with the Landscape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandscape

`func (o *HtmlEngineOptionsInput) SetLandscape(v bool)`

SetLandscape sets Landscape field to given value.

### HasLandscape

`func (o *HtmlEngineOptionsInput) HasLandscape() bool`

HasLandscape returns a boolean if a field has been set.

### GetPrintBackground

`func (o *HtmlEngineOptionsInput) GetPrintBackground() bool`

GetPrintBackground returns the PrintBackground field if non-nil, zero value otherwise.

### GetPrintBackgroundOk

`func (o *HtmlEngineOptionsInput) GetPrintBackgroundOk() (*bool, bool)`

GetPrintBackgroundOk returns a tuple with the PrintBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintBackground

`func (o *HtmlEngineOptionsInput) SetPrintBackground(v bool)`

SetPrintBackground sets PrintBackground field to given value.

### HasPrintBackground

`func (o *HtmlEngineOptionsInput) HasPrintBackground() bool`

HasPrintBackground returns a boolean if a field has been set.

### GetMarginTop

`func (o *HtmlEngineOptionsInput) GetMarginTop() string`

GetMarginTop returns the MarginTop field if non-nil, zero value otherwise.

### GetMarginTopOk

`func (o *HtmlEngineOptionsInput) GetMarginTopOk() (*string, bool)`

GetMarginTopOk returns a tuple with the MarginTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginTop

`func (o *HtmlEngineOptionsInput) SetMarginTop(v string)`

SetMarginTop sets MarginTop field to given value.

### HasMarginTop

`func (o *HtmlEngineOptionsInput) HasMarginTop() bool`

HasMarginTop returns a boolean if a field has been set.

### GetMarginBottom

`func (o *HtmlEngineOptionsInput) GetMarginBottom() string`

GetMarginBottom returns the MarginBottom field if non-nil, zero value otherwise.

### GetMarginBottomOk

`func (o *HtmlEngineOptionsInput) GetMarginBottomOk() (*string, bool)`

GetMarginBottomOk returns a tuple with the MarginBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBottom

`func (o *HtmlEngineOptionsInput) SetMarginBottom(v string)`

SetMarginBottom sets MarginBottom field to given value.

### HasMarginBottom

`func (o *HtmlEngineOptionsInput) HasMarginBottom() bool`

HasMarginBottom returns a boolean if a field has been set.

### GetMarginLeft

`func (o *HtmlEngineOptionsInput) GetMarginLeft() string`

GetMarginLeft returns the MarginLeft field if non-nil, zero value otherwise.

### GetMarginLeftOk

`func (o *HtmlEngineOptionsInput) GetMarginLeftOk() (*string, bool)`

GetMarginLeftOk returns a tuple with the MarginLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLeft

`func (o *HtmlEngineOptionsInput) SetMarginLeft(v string)`

SetMarginLeft sets MarginLeft field to given value.

### HasMarginLeft

`func (o *HtmlEngineOptionsInput) HasMarginLeft() bool`

HasMarginLeft returns a boolean if a field has been set.

### GetMarginRight

`func (o *HtmlEngineOptionsInput) GetMarginRight() string`

GetMarginRight returns the MarginRight field if non-nil, zero value otherwise.

### GetMarginRightOk

`func (o *HtmlEngineOptionsInput) GetMarginRightOk() (*string, bool)`

GetMarginRightOk returns a tuple with the MarginRight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRight

`func (o *HtmlEngineOptionsInput) SetMarginRight(v string)`

SetMarginRight sets MarginRight field to given value.

### HasMarginRight

`func (o *HtmlEngineOptionsInput) HasMarginRight() bool`

HasMarginRight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


