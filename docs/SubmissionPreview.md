# SubmissionPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | **NullableString** |  | 
**DataRequests** | [**[]SubmissionDataRequest**](SubmissionDataRequest.md) |  | 
**Editable** | **NullableBool** |  | 
**Expired** | **bool** |  | 
**ExpiresAt** | **NullableString** |  | 
**Id** | **NullableString** |  | 
**JsonSchemaErrors** | **[]string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**Password** | **NullableString** |  | 
**ProcessedAt** | **NullableString** |  | 
**State** | **string** |  | 
**TemplateId** | **NullableString** |  | 
**Test** | **bool** |  | 
**TruncatedText** | **map[string]interface{}** |  | 
**PdfHash** | **NullableString** |  | 
**DownloadUrl** | **NullableString** |  | 
**PermanentDownloadUrl** | **NullableString** |  | 
**PreviewDownloadUrl** | **NullableString** |  | 
**PreviewGeneratedAt** | **NullableString** |  | 
**AuditTrailDownloadUrl** | **NullableString** |  | 
**Actions** | [**[]SubmissionAction**](SubmissionAction.md) |  | 

## Methods

### NewSubmissionPreview

`func NewSubmissionPreview(batchId NullableString, dataRequests []SubmissionDataRequest, editable NullableBool, expired bool, expiresAt NullableString, id NullableString, jsonSchemaErrors []string, metadata map[string]interface{}, password NullableString, processedAt NullableString, state string, templateId NullableString, test bool, truncatedText map[string]interface{}, pdfHash NullableString, downloadUrl NullableString, permanentDownloadUrl NullableString, previewDownloadUrl NullableString, previewGeneratedAt NullableString, auditTrailDownloadUrl NullableString, actions []SubmissionAction, ) *SubmissionPreview`

NewSubmissionPreview instantiates a new SubmissionPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionPreviewWithDefaults

`func NewSubmissionPreviewWithDefaults() *SubmissionPreview`

NewSubmissionPreviewWithDefaults instantiates a new SubmissionPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *SubmissionPreview) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *SubmissionPreview) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *SubmissionPreview) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.


### SetBatchIdNil

`func (o *SubmissionPreview) SetBatchIdNil(b bool)`

 SetBatchIdNil sets the value for BatchId to be an explicit nil

### UnsetBatchId
`func (o *SubmissionPreview) UnsetBatchId()`

UnsetBatchId ensures that no value is present for BatchId, not even an explicit nil
### GetDataRequests

`func (o *SubmissionPreview) GetDataRequests() []SubmissionDataRequest`

GetDataRequests returns the DataRequests field if non-nil, zero value otherwise.

### GetDataRequestsOk

`func (o *SubmissionPreview) GetDataRequestsOk() (*[]SubmissionDataRequest, bool)`

GetDataRequestsOk returns a tuple with the DataRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRequests

`func (o *SubmissionPreview) SetDataRequests(v []SubmissionDataRequest)`

SetDataRequests sets DataRequests field to given value.


### GetEditable

`func (o *SubmissionPreview) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *SubmissionPreview) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *SubmissionPreview) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### SetEditableNil

`func (o *SubmissionPreview) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *SubmissionPreview) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetExpired

`func (o *SubmissionPreview) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *SubmissionPreview) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *SubmissionPreview) SetExpired(v bool)`

SetExpired sets Expired field to given value.


### GetExpiresAt

`func (o *SubmissionPreview) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SubmissionPreview) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SubmissionPreview) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *SubmissionPreview) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *SubmissionPreview) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetId

`func (o *SubmissionPreview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionPreview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionPreview) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *SubmissionPreview) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubmissionPreview) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetJsonSchemaErrors

`func (o *SubmissionPreview) GetJsonSchemaErrors() []string`

GetJsonSchemaErrors returns the JsonSchemaErrors field if non-nil, zero value otherwise.

### GetJsonSchemaErrorsOk

`func (o *SubmissionPreview) GetJsonSchemaErrorsOk() (*[]string, bool)`

GetJsonSchemaErrorsOk returns a tuple with the JsonSchemaErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchemaErrors

`func (o *SubmissionPreview) SetJsonSchemaErrors(v []string)`

SetJsonSchemaErrors sets JsonSchemaErrors field to given value.


### SetJsonSchemaErrorsNil

`func (o *SubmissionPreview) SetJsonSchemaErrorsNil(b bool)`

 SetJsonSchemaErrorsNil sets the value for JsonSchemaErrors to be an explicit nil

### UnsetJsonSchemaErrors
`func (o *SubmissionPreview) UnsetJsonSchemaErrors()`

UnsetJsonSchemaErrors ensures that no value is present for JsonSchemaErrors, not even an explicit nil
### GetMetadata

`func (o *SubmissionPreview) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubmissionPreview) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubmissionPreview) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetPassword

`func (o *SubmissionPreview) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SubmissionPreview) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SubmissionPreview) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *SubmissionPreview) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SubmissionPreview) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetProcessedAt

`func (o *SubmissionPreview) GetProcessedAt() string`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *SubmissionPreview) GetProcessedAtOk() (*string, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *SubmissionPreview) SetProcessedAt(v string)`

SetProcessedAt sets ProcessedAt field to given value.


### SetProcessedAtNil

`func (o *SubmissionPreview) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *SubmissionPreview) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetState

`func (o *SubmissionPreview) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubmissionPreview) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubmissionPreview) SetState(v string)`

SetState sets State field to given value.


### GetTemplateId

`func (o *SubmissionPreview) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SubmissionPreview) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SubmissionPreview) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### SetTemplateIdNil

`func (o *SubmissionPreview) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *SubmissionPreview) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetTest

`func (o *SubmissionPreview) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *SubmissionPreview) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *SubmissionPreview) SetTest(v bool)`

SetTest sets Test field to given value.


### GetTruncatedText

`func (o *SubmissionPreview) GetTruncatedText() map[string]interface{}`

GetTruncatedText returns the TruncatedText field if non-nil, zero value otherwise.

### GetTruncatedTextOk

`func (o *SubmissionPreview) GetTruncatedTextOk() (*map[string]interface{}, bool)`

GetTruncatedTextOk returns a tuple with the TruncatedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncatedText

`func (o *SubmissionPreview) SetTruncatedText(v map[string]interface{})`

SetTruncatedText sets TruncatedText field to given value.


### SetTruncatedTextNil

`func (o *SubmissionPreview) SetTruncatedTextNil(b bool)`

 SetTruncatedTextNil sets the value for TruncatedText to be an explicit nil

### UnsetTruncatedText
`func (o *SubmissionPreview) UnsetTruncatedText()`

UnsetTruncatedText ensures that no value is present for TruncatedText, not even an explicit nil
### GetPdfHash

`func (o *SubmissionPreview) GetPdfHash() string`

GetPdfHash returns the PdfHash field if non-nil, zero value otherwise.

### GetPdfHashOk

`func (o *SubmissionPreview) GetPdfHashOk() (*string, bool)`

GetPdfHashOk returns a tuple with the PdfHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfHash

`func (o *SubmissionPreview) SetPdfHash(v string)`

SetPdfHash sets PdfHash field to given value.


### SetPdfHashNil

`func (o *SubmissionPreview) SetPdfHashNil(b bool)`

 SetPdfHashNil sets the value for PdfHash to be an explicit nil

### UnsetPdfHash
`func (o *SubmissionPreview) UnsetPdfHash()`

UnsetPdfHash ensures that no value is present for PdfHash, not even an explicit nil
### GetDownloadUrl

`func (o *SubmissionPreview) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *SubmissionPreview) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *SubmissionPreview) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *SubmissionPreview) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *SubmissionPreview) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetPermanentDownloadUrl

`func (o *SubmissionPreview) GetPermanentDownloadUrl() string`

GetPermanentDownloadUrl returns the PermanentDownloadUrl field if non-nil, zero value otherwise.

### GetPermanentDownloadUrlOk

`func (o *SubmissionPreview) GetPermanentDownloadUrlOk() (*string, bool)`

GetPermanentDownloadUrlOk returns a tuple with the PermanentDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentDownloadUrl

`func (o *SubmissionPreview) SetPermanentDownloadUrl(v string)`

SetPermanentDownloadUrl sets PermanentDownloadUrl field to given value.


### SetPermanentDownloadUrlNil

`func (o *SubmissionPreview) SetPermanentDownloadUrlNil(b bool)`

 SetPermanentDownloadUrlNil sets the value for PermanentDownloadUrl to be an explicit nil

### UnsetPermanentDownloadUrl
`func (o *SubmissionPreview) UnsetPermanentDownloadUrl()`

UnsetPermanentDownloadUrl ensures that no value is present for PermanentDownloadUrl, not even an explicit nil
### GetPreviewDownloadUrl

`func (o *SubmissionPreview) GetPreviewDownloadUrl() string`

GetPreviewDownloadUrl returns the PreviewDownloadUrl field if non-nil, zero value otherwise.

### GetPreviewDownloadUrlOk

`func (o *SubmissionPreview) GetPreviewDownloadUrlOk() (*string, bool)`

GetPreviewDownloadUrlOk returns a tuple with the PreviewDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewDownloadUrl

`func (o *SubmissionPreview) SetPreviewDownloadUrl(v string)`

SetPreviewDownloadUrl sets PreviewDownloadUrl field to given value.


### SetPreviewDownloadUrlNil

`func (o *SubmissionPreview) SetPreviewDownloadUrlNil(b bool)`

 SetPreviewDownloadUrlNil sets the value for PreviewDownloadUrl to be an explicit nil

### UnsetPreviewDownloadUrl
`func (o *SubmissionPreview) UnsetPreviewDownloadUrl()`

UnsetPreviewDownloadUrl ensures that no value is present for PreviewDownloadUrl, not even an explicit nil
### GetPreviewGeneratedAt

`func (o *SubmissionPreview) GetPreviewGeneratedAt() string`

GetPreviewGeneratedAt returns the PreviewGeneratedAt field if non-nil, zero value otherwise.

### GetPreviewGeneratedAtOk

`func (o *SubmissionPreview) GetPreviewGeneratedAtOk() (*string, bool)`

GetPreviewGeneratedAtOk returns a tuple with the PreviewGeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewGeneratedAt

`func (o *SubmissionPreview) SetPreviewGeneratedAt(v string)`

SetPreviewGeneratedAt sets PreviewGeneratedAt field to given value.


### SetPreviewGeneratedAtNil

`func (o *SubmissionPreview) SetPreviewGeneratedAtNil(b bool)`

 SetPreviewGeneratedAtNil sets the value for PreviewGeneratedAt to be an explicit nil

### UnsetPreviewGeneratedAt
`func (o *SubmissionPreview) UnsetPreviewGeneratedAt()`

UnsetPreviewGeneratedAt ensures that no value is present for PreviewGeneratedAt, not even an explicit nil
### GetAuditTrailDownloadUrl

`func (o *SubmissionPreview) GetAuditTrailDownloadUrl() string`

GetAuditTrailDownloadUrl returns the AuditTrailDownloadUrl field if non-nil, zero value otherwise.

### GetAuditTrailDownloadUrlOk

`func (o *SubmissionPreview) GetAuditTrailDownloadUrlOk() (*string, bool)`

GetAuditTrailDownloadUrlOk returns a tuple with the AuditTrailDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditTrailDownloadUrl

`func (o *SubmissionPreview) SetAuditTrailDownloadUrl(v string)`

SetAuditTrailDownloadUrl sets AuditTrailDownloadUrl field to given value.


### SetAuditTrailDownloadUrlNil

`func (o *SubmissionPreview) SetAuditTrailDownloadUrlNil(b bool)`

 SetAuditTrailDownloadUrlNil sets the value for AuditTrailDownloadUrl to be an explicit nil

### UnsetAuditTrailDownloadUrl
`func (o *SubmissionPreview) UnsetAuditTrailDownloadUrl()`

UnsetAuditTrailDownloadUrl ensures that no value is present for AuditTrailDownloadUrl, not even an explicit nil
### GetActions

`func (o *SubmissionPreview) GetActions() []SubmissionAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *SubmissionPreview) GetActionsOk() (*[]SubmissionAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *SubmissionPreview) SetActions(v []SubmissionAction)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


