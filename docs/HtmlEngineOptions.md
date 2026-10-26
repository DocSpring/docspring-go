# HtmlEngineOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | [**PageFormat**](PageFormat.md) | The paper size. `Custom` uses `width` and `height`. | 
**Width** | Pointer to **string** | The page width for the `Custom` format, with a unit (`8.5in`, `210mm`). | [optional] 
**Height** | Pointer to **string** | The page height for the `Custom` format, with a unit. | [optional] 
**Scale** | **float32** | How much to scale the page, such as `0.8`. `1` prints it at its normal size. | 
**Landscape** | **bool** | Print pages in landscape orientation. | 
**PrintBackground** | **bool** | Print background colors and images. | 
**MarginTop** | **string** | The top margin, with a unit, such as `1in` or `20mm`. | 
**MarginBottom** | **string** | The bottom margin, with a unit. | 
**MarginLeft** | **string** | The left margin, with a unit. | 
**MarginRight** | **string** | The right margin, with a unit. | 

## Methods

### NewHtmlEngineOptions

`func NewHtmlEngineOptions(format PageFormat, scale float32, landscape bool, printBackground bool, marginTop string, marginBottom string, marginLeft string, marginRight string, ) *HtmlEngineOptions`

NewHtmlEngineOptions instantiates a new HtmlEngineOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHtmlEngineOptionsWithDefaults

`func NewHtmlEngineOptionsWithDefaults() *HtmlEngineOptions`

NewHtmlEngineOptionsWithDefaults instantiates a new HtmlEngineOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *HtmlEngineOptions) GetFormat() PageFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *HtmlEngineOptions) GetFormatOk() (*PageFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *HtmlEngineOptions) SetFormat(v PageFormat)`

SetFormat sets Format field to given value.


### GetWidth

`func (o *HtmlEngineOptions) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *HtmlEngineOptions) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *HtmlEngineOptions) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *HtmlEngineOptions) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *HtmlEngineOptions) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *HtmlEngineOptions) GetHeightOk() (*string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *HtmlEngineOptions) SetHeight(v string)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *HtmlEngineOptions) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetScale

`func (o *HtmlEngineOptions) GetScale() float32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *HtmlEngineOptions) GetScaleOk() (*float32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *HtmlEngineOptions) SetScale(v float32)`

SetScale sets Scale field to given value.


### GetLandscape

`func (o *HtmlEngineOptions) GetLandscape() bool`

GetLandscape returns the Landscape field if non-nil, zero value otherwise.

### GetLandscapeOk

`func (o *HtmlEngineOptions) GetLandscapeOk() (*bool, bool)`

GetLandscapeOk returns a tuple with the Landscape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandscape

`func (o *HtmlEngineOptions) SetLandscape(v bool)`

SetLandscape sets Landscape field to given value.


### GetPrintBackground

`func (o *HtmlEngineOptions) GetPrintBackground() bool`

GetPrintBackground returns the PrintBackground field if non-nil, zero value otherwise.

### GetPrintBackgroundOk

`func (o *HtmlEngineOptions) GetPrintBackgroundOk() (*bool, bool)`

GetPrintBackgroundOk returns a tuple with the PrintBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintBackground

`func (o *HtmlEngineOptions) SetPrintBackground(v bool)`

SetPrintBackground sets PrintBackground field to given value.


### GetMarginTop

`func (o *HtmlEngineOptions) GetMarginTop() string`

GetMarginTop returns the MarginTop field if non-nil, zero value otherwise.

### GetMarginTopOk

`func (o *HtmlEngineOptions) GetMarginTopOk() (*string, bool)`

GetMarginTopOk returns a tuple with the MarginTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginTop

`func (o *HtmlEngineOptions) SetMarginTop(v string)`

SetMarginTop sets MarginTop field to given value.


### GetMarginBottom

`func (o *HtmlEngineOptions) GetMarginBottom() string`

GetMarginBottom returns the MarginBottom field if non-nil, zero value otherwise.

### GetMarginBottomOk

`func (o *HtmlEngineOptions) GetMarginBottomOk() (*string, bool)`

GetMarginBottomOk returns a tuple with the MarginBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBottom

`func (o *HtmlEngineOptions) SetMarginBottom(v string)`

SetMarginBottom sets MarginBottom field to given value.


### GetMarginLeft

`func (o *HtmlEngineOptions) GetMarginLeft() string`

GetMarginLeft returns the MarginLeft field if non-nil, zero value otherwise.

### GetMarginLeftOk

`func (o *HtmlEngineOptions) GetMarginLeftOk() (*string, bool)`

GetMarginLeftOk returns a tuple with the MarginLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLeft

`func (o *HtmlEngineOptions) SetMarginLeft(v string)`

SetMarginLeft sets MarginLeft field to given value.


### GetMarginRight

`func (o *HtmlEngineOptions) GetMarginRight() string`

GetMarginRight returns the MarginRight field if non-nil, zero value otherwise.

### GetMarginRightOk

`func (o *HtmlEngineOptions) GetMarginRightOk() (*string, bool)`

GetMarginRightOk returns a tuple with the MarginRight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRight

`func (o *HtmlEngineOptions) SetMarginRight(v string)`

SetMarginRight sets MarginRight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


