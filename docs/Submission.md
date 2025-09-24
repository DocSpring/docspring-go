# Submission

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
**TemplateType** | **string** |  | 
**TemplateVersion** | **NullableString** |  | 
**Test** | **bool** |  | 
**TruncatedText** | **map[string]interface{}** |  | 
**PdfHash** | **NullableString** |  | 
**DownloadUrl** | **NullableString** |  | 
**PermanentDownloadUrl** | **NullableString** |  | 
**PreviewDownloadUrl** | **NullableString** |  | 
**PreviewGeneratedAt** | **NullableString** |  | 
**AuditTrailDownloadUrl** | **NullableString** |  | 
**Actions** | [**[]SubmissionAction**](SubmissionAction.md) |  | 
**Source** | **string** |  | 
**Referrer** | **NullableString** |  | 
**Data** | **map[string]interface{}** |  | 

## Methods

### NewSubmission

`func NewSubmission(batchId NullableString, dataRequests []SubmissionDataRequest, editable NullableBool, expired bool, expiresAt NullableString, id NullableString, jsonSchemaErrors []string, metadata map[string]interface{}, password NullableString, processedAt NullableString, state string, templateId NullableString, templateType string, templateVersion NullableString, test bool, truncatedText map[string]interface{}, pdfHash NullableString, downloadUrl NullableString, permanentDownloadUrl NullableString, previewDownloadUrl NullableString, previewGeneratedAt NullableString, auditTrailDownloadUrl NullableString, actions []SubmissionAction, source string, referrer NullableString, data map[string]interface{}, ) *Submission`

NewSubmission instantiates a new Submission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionWithDefaults

`func NewSubmissionWithDefaults() *Submission`

NewSubmissionWithDefaults instantiates a new Submission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *Submission) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *Submission) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *Submission) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.


### SetBatchIdNil

`func (o *Submission) SetBatchIdNil(b bool)`

 SetBatchIdNil sets the value for BatchId to be an explicit nil

### UnsetBatchId
`func (o *Submission) UnsetBatchId()`

UnsetBatchId ensures that no value is present for BatchId, not even an explicit nil
### GetDataRequests

`func (o *Submission) GetDataRequests() []SubmissionDataRequest`

GetDataRequests returns the DataRequests field if non-nil, zero value otherwise.

### GetDataRequestsOk

`func (o *Submission) GetDataRequestsOk() (*[]SubmissionDataRequest, bool)`

GetDataRequestsOk returns a tuple with the DataRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRequests

`func (o *Submission) SetDataRequests(v []SubmissionDataRequest)`

SetDataRequests sets DataRequests field to given value.


### GetEditable

`func (o *Submission) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *Submission) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *Submission) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### SetEditableNil

`func (o *Submission) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *Submission) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetExpired

`func (o *Submission) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *Submission) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *Submission) SetExpired(v bool)`

SetExpired sets Expired field to given value.


### GetExpiresAt

`func (o *Submission) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Submission) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Submission) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *Submission) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Submission) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetId

`func (o *Submission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Submission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Submission) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *Submission) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Submission) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetJsonSchemaErrors

`func (o *Submission) GetJsonSchemaErrors() []string`

GetJsonSchemaErrors returns the JsonSchemaErrors field if non-nil, zero value otherwise.

### GetJsonSchemaErrorsOk

`func (o *Submission) GetJsonSchemaErrorsOk() (*[]string, bool)`

GetJsonSchemaErrorsOk returns a tuple with the JsonSchemaErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchemaErrors

`func (o *Submission) SetJsonSchemaErrors(v []string)`

SetJsonSchemaErrors sets JsonSchemaErrors field to given value.


### SetJsonSchemaErrorsNil

`func (o *Submission) SetJsonSchemaErrorsNil(b bool)`

 SetJsonSchemaErrorsNil sets the value for JsonSchemaErrors to be an explicit nil

### UnsetJsonSchemaErrors
`func (o *Submission) UnsetJsonSchemaErrors()`

UnsetJsonSchemaErrors ensures that no value is present for JsonSchemaErrors, not even an explicit nil
### GetMetadata

`func (o *Submission) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Submission) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Submission) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetPassword

`func (o *Submission) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Submission) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Submission) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *Submission) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *Submission) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetProcessedAt

`func (o *Submission) GetProcessedAt() string`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *Submission) GetProcessedAtOk() (*string, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *Submission) SetProcessedAt(v string)`

SetProcessedAt sets ProcessedAt field to given value.


### SetProcessedAtNil

`func (o *Submission) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *Submission) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetState

`func (o *Submission) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Submission) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Submission) SetState(v string)`

SetState sets State field to given value.


### GetTemplateId

`func (o *Submission) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Submission) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Submission) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### SetTemplateIdNil

`func (o *Submission) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *Submission) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetTemplateType

`func (o *Submission) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *Submission) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *Submission) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.


### GetTemplateVersion

`func (o *Submission) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *Submission) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *Submission) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.


### SetTemplateVersionNil

`func (o *Submission) SetTemplateVersionNil(b bool)`

 SetTemplateVersionNil sets the value for TemplateVersion to be an explicit nil

### UnsetTemplateVersion
`func (o *Submission) UnsetTemplateVersion()`

UnsetTemplateVersion ensures that no value is present for TemplateVersion, not even an explicit nil
### GetTest

`func (o *Submission) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *Submission) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *Submission) SetTest(v bool)`

SetTest sets Test field to given value.


### GetTruncatedText

`func (o *Submission) GetTruncatedText() map[string]interface{}`

GetTruncatedText returns the TruncatedText field if non-nil, zero value otherwise.

### GetTruncatedTextOk

`func (o *Submission) GetTruncatedTextOk() (*map[string]interface{}, bool)`

GetTruncatedTextOk returns a tuple with the TruncatedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncatedText

`func (o *Submission) SetTruncatedText(v map[string]interface{})`

SetTruncatedText sets TruncatedText field to given value.


### SetTruncatedTextNil

`func (o *Submission) SetTruncatedTextNil(b bool)`

 SetTruncatedTextNil sets the value for TruncatedText to be an explicit nil

### UnsetTruncatedText
`func (o *Submission) UnsetTruncatedText()`

UnsetTruncatedText ensures that no value is present for TruncatedText, not even an explicit nil
### GetPdfHash

`func (o *Submission) GetPdfHash() string`

GetPdfHash returns the PdfHash field if non-nil, zero value otherwise.

### GetPdfHashOk

`func (o *Submission) GetPdfHashOk() (*string, bool)`

GetPdfHashOk returns a tuple with the PdfHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfHash

`func (o *Submission) SetPdfHash(v string)`

SetPdfHash sets PdfHash field to given value.


### SetPdfHashNil

`func (o *Submission) SetPdfHashNil(b bool)`

 SetPdfHashNil sets the value for PdfHash to be an explicit nil

### UnsetPdfHash
`func (o *Submission) UnsetPdfHash()`

UnsetPdfHash ensures that no value is present for PdfHash, not even an explicit nil
### GetDownloadUrl

`func (o *Submission) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *Submission) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *Submission) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *Submission) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *Submission) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetPermanentDownloadUrl

`func (o *Submission) GetPermanentDownloadUrl() string`

GetPermanentDownloadUrl returns the PermanentDownloadUrl field if non-nil, zero value otherwise.

### GetPermanentDownloadUrlOk

`func (o *Submission) GetPermanentDownloadUrlOk() (*string, bool)`

GetPermanentDownloadUrlOk returns a tuple with the PermanentDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentDownloadUrl

`func (o *Submission) SetPermanentDownloadUrl(v string)`

SetPermanentDownloadUrl sets PermanentDownloadUrl field to given value.


### SetPermanentDownloadUrlNil

`func (o *Submission) SetPermanentDownloadUrlNil(b bool)`

 SetPermanentDownloadUrlNil sets the value for PermanentDownloadUrl to be an explicit nil

### UnsetPermanentDownloadUrl
`func (o *Submission) UnsetPermanentDownloadUrl()`

UnsetPermanentDownloadUrl ensures that no value is present for PermanentDownloadUrl, not even an explicit nil
### GetPreviewDownloadUrl

`func (o *Submission) GetPreviewDownloadUrl() string`

GetPreviewDownloadUrl returns the PreviewDownloadUrl field if non-nil, zero value otherwise.

### GetPreviewDownloadUrlOk

`func (o *Submission) GetPreviewDownloadUrlOk() (*string, bool)`

GetPreviewDownloadUrlOk returns a tuple with the PreviewDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewDownloadUrl

`func (o *Submission) SetPreviewDownloadUrl(v string)`

SetPreviewDownloadUrl sets PreviewDownloadUrl field to given value.


### SetPreviewDownloadUrlNil

`func (o *Submission) SetPreviewDownloadUrlNil(b bool)`

 SetPreviewDownloadUrlNil sets the value for PreviewDownloadUrl to be an explicit nil

### UnsetPreviewDownloadUrl
`func (o *Submission) UnsetPreviewDownloadUrl()`

UnsetPreviewDownloadUrl ensures that no value is present for PreviewDownloadUrl, not even an explicit nil
### GetPreviewGeneratedAt

`func (o *Submission) GetPreviewGeneratedAt() string`

GetPreviewGeneratedAt returns the PreviewGeneratedAt field if non-nil, zero value otherwise.

### GetPreviewGeneratedAtOk

`func (o *Submission) GetPreviewGeneratedAtOk() (*string, bool)`

GetPreviewGeneratedAtOk returns a tuple with the PreviewGeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewGeneratedAt

`func (o *Submission) SetPreviewGeneratedAt(v string)`

SetPreviewGeneratedAt sets PreviewGeneratedAt field to given value.


### SetPreviewGeneratedAtNil

`func (o *Submission) SetPreviewGeneratedAtNil(b bool)`

 SetPreviewGeneratedAtNil sets the value for PreviewGeneratedAt to be an explicit nil

### UnsetPreviewGeneratedAt
`func (o *Submission) UnsetPreviewGeneratedAt()`

UnsetPreviewGeneratedAt ensures that no value is present for PreviewGeneratedAt, not even an explicit nil
### GetAuditTrailDownloadUrl

`func (o *Submission) GetAuditTrailDownloadUrl() string`

GetAuditTrailDownloadUrl returns the AuditTrailDownloadUrl field if non-nil, zero value otherwise.

### GetAuditTrailDownloadUrlOk

`func (o *Submission) GetAuditTrailDownloadUrlOk() (*string, bool)`

GetAuditTrailDownloadUrlOk returns a tuple with the AuditTrailDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditTrailDownloadUrl

`func (o *Submission) SetAuditTrailDownloadUrl(v string)`

SetAuditTrailDownloadUrl sets AuditTrailDownloadUrl field to given value.


### SetAuditTrailDownloadUrlNil

`func (o *Submission) SetAuditTrailDownloadUrlNil(b bool)`

 SetAuditTrailDownloadUrlNil sets the value for AuditTrailDownloadUrl to be an explicit nil

### UnsetAuditTrailDownloadUrl
`func (o *Submission) UnsetAuditTrailDownloadUrl()`

UnsetAuditTrailDownloadUrl ensures that no value is present for AuditTrailDownloadUrl, not even an explicit nil
### GetActions

`func (o *Submission) GetActions() []SubmissionAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Submission) GetActionsOk() (*[]SubmissionAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Submission) SetActions(v []SubmissionAction)`

SetActions sets Actions field to given value.


### GetSource

`func (o *Submission) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Submission) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Submission) SetSource(v string)`

SetSource sets Source field to given value.


### GetReferrer

`func (o *Submission) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *Submission) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *Submission) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.


### SetReferrerNil

`func (o *Submission) SetReferrerNil(b bool)`

 SetReferrerNil sets the value for Referrer to be an explicit nil

### UnsetReferrer
`func (o *Submission) UnsetReferrer()`

UnsetReferrer ensures that no value is present for Referrer, not even an explicit nil
### GetData

`func (o *Submission) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Submission) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Submission) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *Submission) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Submission) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


