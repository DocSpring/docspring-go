# TemplatePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The preview's ID. | 
**TemplateId** | **string** | The template's ID. | 
**Version** | **NullableString** | The published version that was rendered, such as `1.2.0`. Null for the draft. | 
**State** | [**TemplatePreviewState**](TemplatePreviewState.md) | `pending` while it renders (usually a few seconds), then `processed`, or `error` with an `error_message`. | 
**ErrorMessage** | **NullableString** | Why the preview couldn't be rendered, such as an error in an HTML template's Liquid, or a value in the data that a field can't print. | 
**PageCount** | **NullableInt32** | How many pages the template printed, once it's processed. | 
**Pages** | [**[]PageImage**](PageImage.md) | A PNG image of each page once it's processed, in order: `pages[0]` is page 1. A preview has no PDF. Their `download_url`s expire, so get the preview again when you need them. | 
**ExpiresAt** | **time.Time** | When the preview and its images are deleted. | 

## Methods

### NewTemplatePreview

`func NewTemplatePreview(id string, templateId string, version NullableString, state TemplatePreviewState, errorMessage NullableString, pageCount NullableInt32, pages []PageImage, expiresAt time.Time, ) *TemplatePreview`

NewTemplatePreview instantiates a new TemplatePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePreviewWithDefaults

`func NewTemplatePreviewWithDefaults() *TemplatePreview`

NewTemplatePreviewWithDefaults instantiates a new TemplatePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplatePreview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplatePreview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplatePreview) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *TemplatePreview) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplatePreview) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplatePreview) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetVersion

`func (o *TemplatePreview) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TemplatePreview) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TemplatePreview) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *TemplatePreview) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TemplatePreview) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetState

`func (o *TemplatePreview) GetState() TemplatePreviewState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TemplatePreview) GetStateOk() (*TemplatePreviewState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TemplatePreview) SetState(v TemplatePreviewState)`

SetState sets State field to given value.


### GetErrorMessage

`func (o *TemplatePreview) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TemplatePreview) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TemplatePreview) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *TemplatePreview) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TemplatePreview) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetPageCount

`func (o *TemplatePreview) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *TemplatePreview) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *TemplatePreview) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.


### SetPageCountNil

`func (o *TemplatePreview) SetPageCountNil(b bool)`

 SetPageCountNil sets the value for PageCount to be an explicit nil

### UnsetPageCount
`func (o *TemplatePreview) UnsetPageCount()`

UnsetPageCount ensures that no value is present for PageCount, not even an explicit nil
### GetPages

`func (o *TemplatePreview) GetPages() []PageImage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *TemplatePreview) GetPagesOk() (*[]PageImage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *TemplatePreview) SetPages(v []PageImage)`

SetPages sets Pages field to given value.


### GetExpiresAt

`func (o *TemplatePreview) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TemplatePreview) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TemplatePreview) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


