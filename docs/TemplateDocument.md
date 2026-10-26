# TemplateDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **NullableString** | The PDF's file name. Null until a `document_url` has been downloaded. | 
**Md5** | **NullableString** | The MD5 hash of the PDF. | 
**State** | [**TemplateDocumentState**](TemplateDocumentState.md) | `pending` while DocSpring processes the PDF, then `processed`. `password_required`, `malformed_pdf` and `unknown_error` mean the PDF couldn't be processed, and `error_message` says why: upload it again or upload another one. `downloading` while DocSpring downloads a `document_url`, then as for an upload. `download_failed` means it couldn't be downloaded, and `error_message` says why. A template keeps its PDF while a new one is downloaded, and when it couldn't be. | 
**ErrorMessage** | **NullableString** | Why the PDF couldn't be processed or downloaded, and how to fix it. Null unless `state` is an error. | 
**FieldReimportState** | [**FieldReimportState**](FieldReimportState.md) | Whether the fields are being re-imported from this PDF (`reimport_fields` when the PDF was replaced): `requested` while DocSpring processes the PDF, then `imported`, or `no_fields` when the PDF had no fillable fields and the template kept its fields. `none` when no re-import was asked for, or when it couldn't run (the PDF couldn't be processed, or the template was locked or published while it was). | 
**PageCount** | **int32** | 0 until the PDF has been processed. | 
**PageDimensions** | **[][]float32** | Each page's width and height in PDF points. Null until the PDF has been processed. | 
**HasFillableFields** | **NullableBool** | Whether the PDF has fillable form fields. Null until DocSpring has read the PDF. | 
**PageProcessingStatuses** | [**[]TemplatePageProcessingStatus**](TemplatePageProcessingStatus.md) | How processing is going on each page, in page order: its image, and its fields when DocSpring sets them up (`detect_fields` when the template was created). `state` becomes `processed` once every page is done. Empty until DocSpring has read the PDF, and for PDFs uploaded before DocSpring recorded page statuses. | 
**FieldLogicStatus** | [**NullableTemplateFieldLogicStatus**](TemplateFieldLogicStatus.md) | AI field setup's last step, which reads the whole form once every page is done and can still rename fields and add conditional logic and required fields: `pending` or `processing` until it finishes, then `processed`, `failed` or `skipped`. Null until every page is done, and when there's no logic step to run (`field_logic` is `none`, or the PDF has no fields). | 
**DownloadUrl** | **NullableString** | A temporary link to download the PDF. | 
**PermanentUrl** | **NullableString** | A link to download the PDF that doesn't expire. Null until the PDF has been processed. | 

## Methods

### NewTemplateDocument

`func NewTemplateDocument(filename NullableString, md5 NullableString, state TemplateDocumentState, errorMessage NullableString, fieldReimportState FieldReimportState, pageCount int32, pageDimensions [][]float32, hasFillableFields NullableBool, pageProcessingStatuses []TemplatePageProcessingStatus, fieldLogicStatus NullableTemplateFieldLogicStatus, downloadUrl NullableString, permanentUrl NullableString, ) *TemplateDocument`

NewTemplateDocument instantiates a new TemplateDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDocumentWithDefaults

`func NewTemplateDocumentWithDefaults() *TemplateDocument`

NewTemplateDocumentWithDefaults instantiates a new TemplateDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *TemplateDocument) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *TemplateDocument) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *TemplateDocument) SetFilename(v string)`

SetFilename sets Filename field to given value.


### SetFilenameNil

`func (o *TemplateDocument) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *TemplateDocument) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetMd5

`func (o *TemplateDocument) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *TemplateDocument) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *TemplateDocument) SetMd5(v string)`

SetMd5 sets Md5 field to given value.


### SetMd5Nil

`func (o *TemplateDocument) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *TemplateDocument) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetState

`func (o *TemplateDocument) GetState() TemplateDocumentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TemplateDocument) GetStateOk() (*TemplateDocumentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TemplateDocument) SetState(v TemplateDocumentState)`

SetState sets State field to given value.


### GetErrorMessage

`func (o *TemplateDocument) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TemplateDocument) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TemplateDocument) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *TemplateDocument) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TemplateDocument) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetFieldReimportState

`func (o *TemplateDocument) GetFieldReimportState() FieldReimportState`

GetFieldReimportState returns the FieldReimportState field if non-nil, zero value otherwise.

### GetFieldReimportStateOk

`func (o *TemplateDocument) GetFieldReimportStateOk() (*FieldReimportState, bool)`

GetFieldReimportStateOk returns a tuple with the FieldReimportState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldReimportState

`func (o *TemplateDocument) SetFieldReimportState(v FieldReimportState)`

SetFieldReimportState sets FieldReimportState field to given value.


### GetPageCount

`func (o *TemplateDocument) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *TemplateDocument) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *TemplateDocument) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.


### GetPageDimensions

`func (o *TemplateDocument) GetPageDimensions() [][]float32`

GetPageDimensions returns the PageDimensions field if non-nil, zero value otherwise.

### GetPageDimensionsOk

`func (o *TemplateDocument) GetPageDimensionsOk() (*[][]float32, bool)`

GetPageDimensionsOk returns a tuple with the PageDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageDimensions

`func (o *TemplateDocument) SetPageDimensions(v [][]float32)`

SetPageDimensions sets PageDimensions field to given value.


### SetPageDimensionsNil

`func (o *TemplateDocument) SetPageDimensionsNil(b bool)`

 SetPageDimensionsNil sets the value for PageDimensions to be an explicit nil

### UnsetPageDimensions
`func (o *TemplateDocument) UnsetPageDimensions()`

UnsetPageDimensions ensures that no value is present for PageDimensions, not even an explicit nil
### GetHasFillableFields

`func (o *TemplateDocument) GetHasFillableFields() bool`

GetHasFillableFields returns the HasFillableFields field if non-nil, zero value otherwise.

### GetHasFillableFieldsOk

`func (o *TemplateDocument) GetHasFillableFieldsOk() (*bool, bool)`

GetHasFillableFieldsOk returns a tuple with the HasFillableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFillableFields

`func (o *TemplateDocument) SetHasFillableFields(v bool)`

SetHasFillableFields sets HasFillableFields field to given value.


### SetHasFillableFieldsNil

`func (o *TemplateDocument) SetHasFillableFieldsNil(b bool)`

 SetHasFillableFieldsNil sets the value for HasFillableFields to be an explicit nil

### UnsetHasFillableFields
`func (o *TemplateDocument) UnsetHasFillableFields()`

UnsetHasFillableFields ensures that no value is present for HasFillableFields, not even an explicit nil
### GetPageProcessingStatuses

`func (o *TemplateDocument) GetPageProcessingStatuses() []TemplatePageProcessingStatus`

GetPageProcessingStatuses returns the PageProcessingStatuses field if non-nil, zero value otherwise.

### GetPageProcessingStatusesOk

`func (o *TemplateDocument) GetPageProcessingStatusesOk() (*[]TemplatePageProcessingStatus, bool)`

GetPageProcessingStatusesOk returns a tuple with the PageProcessingStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageProcessingStatuses

`func (o *TemplateDocument) SetPageProcessingStatuses(v []TemplatePageProcessingStatus)`

SetPageProcessingStatuses sets PageProcessingStatuses field to given value.


### GetFieldLogicStatus

`func (o *TemplateDocument) GetFieldLogicStatus() TemplateFieldLogicStatus`

GetFieldLogicStatus returns the FieldLogicStatus field if non-nil, zero value otherwise.

### GetFieldLogicStatusOk

`func (o *TemplateDocument) GetFieldLogicStatusOk() (*TemplateFieldLogicStatus, bool)`

GetFieldLogicStatusOk returns a tuple with the FieldLogicStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLogicStatus

`func (o *TemplateDocument) SetFieldLogicStatus(v TemplateFieldLogicStatus)`

SetFieldLogicStatus sets FieldLogicStatus field to given value.


### SetFieldLogicStatusNil

`func (o *TemplateDocument) SetFieldLogicStatusNil(b bool)`

 SetFieldLogicStatusNil sets the value for FieldLogicStatus to be an explicit nil

### UnsetFieldLogicStatus
`func (o *TemplateDocument) UnsetFieldLogicStatus()`

UnsetFieldLogicStatus ensures that no value is present for FieldLogicStatus, not even an explicit nil
### GetDownloadUrl

`func (o *TemplateDocument) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *TemplateDocument) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *TemplateDocument) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *TemplateDocument) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *TemplateDocument) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetPermanentUrl

`func (o *TemplateDocument) GetPermanentUrl() string`

GetPermanentUrl returns the PermanentUrl field if non-nil, zero value otherwise.

### GetPermanentUrlOk

`func (o *TemplateDocument) GetPermanentUrlOk() (*string, bool)`

GetPermanentUrlOk returns a tuple with the PermanentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentUrl

`func (o *TemplateDocument) SetPermanentUrl(v string)`

SetPermanentUrl sets PermanentUrl field to given value.


### SetPermanentUrlNil

`func (o *TemplateDocument) SetPermanentUrlNil(b bool)`

 SetPermanentUrlNil sets the value for PermanentUrl to be an explicit nil

### UnsetPermanentUrl
`func (o *TemplateDocument) UnsetPermanentUrl()`

UnsetPermanentUrl ensures that no value is present for PermanentUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


