# CombinedPdf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The combined PDF's ID. | 
**State** | [**CombinedPdfState**](CombinedPdfState.md) | `pending` until the PDF has been generated, then `processed`. `insufficient_credits`: waiting for credits; `error_message` says whether the PDF is generated once you add them. `account_suspended`: waiting for the account to be reactivated, for up to 30 days. `error`: the PDFs couldn't be merged (see `error_message`). Waiting ends in `error` if it's discarded or takes over 30 days. | 
**SourcePdfs** | [**[]CombinedPdfSource**](CombinedPdfSource.md) | The PDFs to merge, in order. | 
**Metadata** | **map[string]interface{}** | Your own data about the combined PDF, such as an ID from your app. | 
**HasPassword** | **bool** | Whether the PDF needs a password to open (`password`). | 
**DeleteCustomFiles** | **bool** | Delete the custom files in `source_pdfs` once the combined PDF has been generated (they're kept if it fails). Other combined PDFs can't use them after that. A test API token can't delete custom files. | 
**ErrorMessage** | **NullableString** | Why the PDFs couldn't be merged, when `state` is `error`. | 
**Files** | [**CombinedPdfFiles**](CombinedPdfFiles.md) | The generated files, by format. Empty until the PDF has been generated, and once it has expired. | 
**Actions** | [**[]SubmissionAction**](SubmissionAction.md) | The uploads of the PDF to your storage (account integrations). | 
**Expired** | **bool** | An expired combined PDF's PDF, metadata, password and source URLs have been deleted. | 
**ExpiresAt** | **NullableTime** | When the combined PDF expires (`expires_in` after it's generated). Null if it's kept until you delete it. | 
**ProcessedAt** | **NullableTime** | When the PDF was generated. Null until then. | 
**CreatedAt** | **time.Time** | When the combined PDF was created. | 
**UpdatedAt** | **time.Time** | When the combined PDF was last updated. | 

## Methods

### NewCombinedPdf

`func NewCombinedPdf(id string, state CombinedPdfState, sourcePdfs []CombinedPdfSource, metadata map[string]interface{}, hasPassword bool, deleteCustomFiles bool, errorMessage NullableString, files CombinedPdfFiles, actions []SubmissionAction, expired bool, expiresAt NullableTime, processedAt NullableTime, createdAt time.Time, updatedAt time.Time, ) *CombinedPdf`

NewCombinedPdf instantiates a new CombinedPdf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinedPdfWithDefaults

`func NewCombinedPdfWithDefaults() *CombinedPdf`

NewCombinedPdfWithDefaults instantiates a new CombinedPdf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CombinedPdf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CombinedPdf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CombinedPdf) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *CombinedPdf) GetState() CombinedPdfState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CombinedPdf) GetStateOk() (*CombinedPdfState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CombinedPdf) SetState(v CombinedPdfState)`

SetState sets State field to given value.


### GetSourcePdfs

`func (o *CombinedPdf) GetSourcePdfs() []CombinedPdfSource`

GetSourcePdfs returns the SourcePdfs field if non-nil, zero value otherwise.

### GetSourcePdfsOk

`func (o *CombinedPdf) GetSourcePdfsOk() (*[]CombinedPdfSource, bool)`

GetSourcePdfsOk returns a tuple with the SourcePdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePdfs

`func (o *CombinedPdf) SetSourcePdfs(v []CombinedPdfSource)`

SetSourcePdfs sets SourcePdfs field to given value.


### GetMetadata

`func (o *CombinedPdf) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CombinedPdf) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CombinedPdf) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetHasPassword

`func (o *CombinedPdf) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *CombinedPdf) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *CombinedPdf) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.


### GetDeleteCustomFiles

`func (o *CombinedPdf) GetDeleteCustomFiles() bool`

GetDeleteCustomFiles returns the DeleteCustomFiles field if non-nil, zero value otherwise.

### GetDeleteCustomFilesOk

`func (o *CombinedPdf) GetDeleteCustomFilesOk() (*bool, bool)`

GetDeleteCustomFilesOk returns a tuple with the DeleteCustomFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteCustomFiles

`func (o *CombinedPdf) SetDeleteCustomFiles(v bool)`

SetDeleteCustomFiles sets DeleteCustomFiles field to given value.


### GetErrorMessage

`func (o *CombinedPdf) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CombinedPdf) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CombinedPdf) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *CombinedPdf) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *CombinedPdf) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetFiles

`func (o *CombinedPdf) GetFiles() CombinedPdfFiles`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CombinedPdf) GetFilesOk() (*CombinedPdfFiles, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CombinedPdf) SetFiles(v CombinedPdfFiles)`

SetFiles sets Files field to given value.


### GetActions

`func (o *CombinedPdf) GetActions() []SubmissionAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CombinedPdf) GetActionsOk() (*[]SubmissionAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CombinedPdf) SetActions(v []SubmissionAction)`

SetActions sets Actions field to given value.


### GetExpired

`func (o *CombinedPdf) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *CombinedPdf) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *CombinedPdf) SetExpired(v bool)`

SetExpired sets Expired field to given value.


### GetExpiresAt

`func (o *CombinedPdf) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CombinedPdf) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CombinedPdf) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *CombinedPdf) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CombinedPdf) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetProcessedAt

`func (o *CombinedPdf) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *CombinedPdf) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *CombinedPdf) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.


### SetProcessedAtNil

`func (o *CombinedPdf) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *CombinedPdf) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetCreatedAt

`func (o *CombinedPdf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CombinedPdf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CombinedPdf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CombinedPdf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CombinedPdf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CombinedPdf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


