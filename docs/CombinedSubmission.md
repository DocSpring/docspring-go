# CombinedSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**State** | **string** |  | 
**Expired** | **bool** |  | 
**ExpiresIn** | **NullableInt32** |  | 
**ExpiresAt** | **NullableString** |  | 
**ProcessedAt** | **NullableString** |  | 
**ErrorMessage** | **NullableString** |  | 
**SubmissionIds** | **[]string** |  | 
**SourcePdfs** | **[]map[string]interface{}** |  | 
**Metadata** | **map[string]interface{}** |  | 
**Password** | **NullableString** |  | 
**PdfHash** | **NullableString** |  | 
**DownloadUrl** | **NullableString** |  | 
**Actions** | [**[]CombinedSubmissionAction**](CombinedSubmissionAction.md) |  | 

## Methods

### NewCombinedSubmission

`func NewCombinedSubmission(id NullableString, state string, expired bool, expiresIn NullableInt32, expiresAt NullableString, processedAt NullableString, errorMessage NullableString, submissionIds []string, sourcePdfs []map[string]interface{}, metadata map[string]interface{}, password NullableString, pdfHash NullableString, downloadUrl NullableString, actions []CombinedSubmissionAction, ) *CombinedSubmission`

NewCombinedSubmission instantiates a new CombinedSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinedSubmissionWithDefaults

`func NewCombinedSubmissionWithDefaults() *CombinedSubmission`

NewCombinedSubmissionWithDefaults instantiates a new CombinedSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CombinedSubmission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CombinedSubmission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CombinedSubmission) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *CombinedSubmission) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CombinedSubmission) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetState

`func (o *CombinedSubmission) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CombinedSubmission) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CombinedSubmission) SetState(v string)`

SetState sets State field to given value.


### GetExpired

`func (o *CombinedSubmission) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *CombinedSubmission) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *CombinedSubmission) SetExpired(v bool)`

SetExpired sets Expired field to given value.


### GetExpiresIn

`func (o *CombinedSubmission) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CombinedSubmission) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CombinedSubmission) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.


### SetExpiresInNil

`func (o *CombinedSubmission) SetExpiresInNil(b bool)`

 SetExpiresInNil sets the value for ExpiresIn to be an explicit nil

### UnsetExpiresIn
`func (o *CombinedSubmission) UnsetExpiresIn()`

UnsetExpiresIn ensures that no value is present for ExpiresIn, not even an explicit nil
### GetExpiresAt

`func (o *CombinedSubmission) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CombinedSubmission) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CombinedSubmission) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *CombinedSubmission) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CombinedSubmission) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetProcessedAt

`func (o *CombinedSubmission) GetProcessedAt() string`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *CombinedSubmission) GetProcessedAtOk() (*string, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *CombinedSubmission) SetProcessedAt(v string)`

SetProcessedAt sets ProcessedAt field to given value.


### SetProcessedAtNil

`func (o *CombinedSubmission) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *CombinedSubmission) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetErrorMessage

`func (o *CombinedSubmission) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CombinedSubmission) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CombinedSubmission) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *CombinedSubmission) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *CombinedSubmission) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetSubmissionIds

`func (o *CombinedSubmission) GetSubmissionIds() []string`

GetSubmissionIds returns the SubmissionIds field if non-nil, zero value otherwise.

### GetSubmissionIdsOk

`func (o *CombinedSubmission) GetSubmissionIdsOk() (*[]string, bool)`

GetSubmissionIdsOk returns a tuple with the SubmissionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionIds

`func (o *CombinedSubmission) SetSubmissionIds(v []string)`

SetSubmissionIds sets SubmissionIds field to given value.


### GetSourcePdfs

`func (o *CombinedSubmission) GetSourcePdfs() []map[string]interface{}`

GetSourcePdfs returns the SourcePdfs field if non-nil, zero value otherwise.

### GetSourcePdfsOk

`func (o *CombinedSubmission) GetSourcePdfsOk() (*[]map[string]interface{}, bool)`

GetSourcePdfsOk returns a tuple with the SourcePdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePdfs

`func (o *CombinedSubmission) SetSourcePdfs(v []map[string]interface{})`

SetSourcePdfs sets SourcePdfs field to given value.


### GetMetadata

`func (o *CombinedSubmission) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CombinedSubmission) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CombinedSubmission) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetPassword

`func (o *CombinedSubmission) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CombinedSubmission) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CombinedSubmission) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *CombinedSubmission) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CombinedSubmission) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPdfHash

`func (o *CombinedSubmission) GetPdfHash() string`

GetPdfHash returns the PdfHash field if non-nil, zero value otherwise.

### GetPdfHashOk

`func (o *CombinedSubmission) GetPdfHashOk() (*string, bool)`

GetPdfHashOk returns a tuple with the PdfHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfHash

`func (o *CombinedSubmission) SetPdfHash(v string)`

SetPdfHash sets PdfHash field to given value.


### SetPdfHashNil

`func (o *CombinedSubmission) SetPdfHashNil(b bool)`

 SetPdfHashNil sets the value for PdfHash to be an explicit nil

### UnsetPdfHash
`func (o *CombinedSubmission) UnsetPdfHash()`

UnsetPdfHash ensures that no value is present for PdfHash, not even an explicit nil
### GetDownloadUrl

`func (o *CombinedSubmission) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *CombinedSubmission) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *CombinedSubmission) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *CombinedSubmission) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *CombinedSubmission) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetActions

`func (o *CombinedSubmission) GetActions() []CombinedSubmissionAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CombinedSubmission) GetActionsOk() (*[]CombinedSubmissionAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CombinedSubmission) SetActions(v []CombinedSubmissionAction)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


