# PageImages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**GeneratedFileState**](GeneratedFileState.md) | `processed` once the images have been made, with the PDF or by POST /submissions/{submission_id}/generate_files. `pending` until then, and `error` if they couldn't be made (see `error_message`). | 
**ErrorMessage** | **NullableString** | Why the images couldn't be made, when `state` is `error`. | 
**PageCount** | **NullableInt32** | How many pages have an image. Null until the images have been made. | 
**Pages** | [**[]PageImage**](PageImage.md) | An image of each page, in order: `pages[0]` is page 1. Lists of submissions leave it out (null): get the submission to see them. Empty until the images have been made. | 

## Methods

### NewPageImages

`func NewPageImages(state GeneratedFileState, errorMessage NullableString, pageCount NullableInt32, pages []PageImage, ) *PageImages`

NewPageImages instantiates a new PageImages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageImagesWithDefaults

`func NewPageImagesWithDefaults() *PageImages`

NewPageImagesWithDefaults instantiates a new PageImages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *PageImages) GetState() GeneratedFileState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PageImages) GetStateOk() (*GeneratedFileState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PageImages) SetState(v GeneratedFileState)`

SetState sets State field to given value.


### GetErrorMessage

`func (o *PageImages) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PageImages) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PageImages) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *PageImages) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *PageImages) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetPageCount

`func (o *PageImages) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *PageImages) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *PageImages) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.


### SetPageCountNil

`func (o *PageImages) SetPageCountNil(b bool)`

 SetPageCountNil sets the value for PageCount to be an explicit nil

### UnsetPageCount
`func (o *PageImages) UnsetPageCount()`

UnsetPageCount ensures that no value is present for PageCount, not even an explicit nil
### GetPages

`func (o *PageImages) GetPages() []PageImage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PageImages) GetPagesOk() (*[]PageImage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PageImages) SetPages(v []PageImage)`

SetPages sets Pages field to given value.


### SetPagesNil

`func (o *PageImages) SetPagesNil(b bool)`

 SetPagesNil sets the value for Pages to be an explicit nil

### UnsetPages
`func (o *PageImages) UnsetPages()`

UnsetPages ensures that no value is present for Pages, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


