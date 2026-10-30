# TemplateDocumentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **NullableString** | The PDF's file name. Null until a `document_url` has been downloaded. | 
**State** | [**TemplateDocumentState**](TemplateDocumentState.md) | `pending` while DocSpring processes the PDF, then `processed` (`ready` says when its pages are done too). `password_required`, `malformed_pdf` and `unknown_error` mean the PDF couldn't be processed, and `error_message` says why: upload it again or upload another one. `downloading` while DocSpring downloads a `document_url`, then as for an upload. `download_failed` means it couldn't be downloaded, and `error_message` says why. A template keeps its PDF while a new one is downloaded, and when it couldn't be. | 
**ErrorMessage** | **NullableString** | Why the PDF couldn't be processed or downloaded, and how to fix it. Null unless `state` is an error. | 
**PageCount** | **int32** | 0 until the PDF has been processed. | 

## Methods

### NewTemplateDocumentSummary

`func NewTemplateDocumentSummary(filename NullableString, state TemplateDocumentState, errorMessage NullableString, pageCount int32, ) *TemplateDocumentSummary`

NewTemplateDocumentSummary instantiates a new TemplateDocumentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDocumentSummaryWithDefaults

`func NewTemplateDocumentSummaryWithDefaults() *TemplateDocumentSummary`

NewTemplateDocumentSummaryWithDefaults instantiates a new TemplateDocumentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *TemplateDocumentSummary) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *TemplateDocumentSummary) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *TemplateDocumentSummary) SetFilename(v string)`

SetFilename sets Filename field to given value.


### SetFilenameNil

`func (o *TemplateDocumentSummary) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *TemplateDocumentSummary) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetState

`func (o *TemplateDocumentSummary) GetState() TemplateDocumentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TemplateDocumentSummary) GetStateOk() (*TemplateDocumentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TemplateDocumentSummary) SetState(v TemplateDocumentState)`

SetState sets State field to given value.


### GetErrorMessage

`func (o *TemplateDocumentSummary) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TemplateDocumentSummary) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TemplateDocumentSummary) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *TemplateDocumentSummary) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TemplateDocumentSummary) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetPageCount

`func (o *TemplateDocumentSummary) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *TemplateDocumentSummary) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *TemplateDocumentSummary) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


