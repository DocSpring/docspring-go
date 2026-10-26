# TemplatePageImages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** | The template's ID. | 
**Version** | **NullableString** | The published version the images are of, such as `1.2.0`. Null for the draft. | 
**PageCount** | **int32** | How many pages the PDF has. | 
**Pages** | [**[]PageImage**](PageImage.md) | An image of each page, in order: `pages[0]` is page 1. Their `download_url`s expire, so get them again when you need them. | 

## Methods

### NewTemplatePageImages

`func NewTemplatePageImages(templateId string, version NullableString, pageCount int32, pages []PageImage, ) *TemplatePageImages`

NewTemplatePageImages instantiates a new TemplatePageImages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePageImagesWithDefaults

`func NewTemplatePageImagesWithDefaults() *TemplatePageImages`

NewTemplatePageImagesWithDefaults instantiates a new TemplatePageImages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *TemplatePageImages) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplatePageImages) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplatePageImages) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetVersion

`func (o *TemplatePageImages) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TemplatePageImages) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TemplatePageImages) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *TemplatePageImages) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TemplatePageImages) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetPageCount

`func (o *TemplatePageImages) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *TemplatePageImages) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *TemplatePageImages) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.


### GetPages

`func (o *TemplatePageImages) GetPages() []PageImage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *TemplatePageImages) GetPagesOk() (*[]PageImage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *TemplatePageImages) SetPages(v []PageImage)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


