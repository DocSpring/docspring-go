# SubmissionFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pdf** | Pointer to [**GeneratedFile**](GeneratedFile.md) | The submission's PDF. | [optional] 
**AuditTrail** | Pointer to [**GeneratedFile**](GeneratedFile.md) | The audit trail of a submission with signature requests: who signed, when and how. It's also added to the end of the PDF when the template's `merge_audit_trail_pdf` is on. | [optional] 
**Preview** | Pointer to [**GeneratedFile**](GeneratedFile.md) | The preview PDF from POST /submissions/{submission_id}/generate_preview, such as for a submission still waiting for its signature requests. | [optional] 

## Methods

### NewSubmissionFiles

`func NewSubmissionFiles() *SubmissionFiles`

NewSubmissionFiles instantiates a new SubmissionFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionFilesWithDefaults

`func NewSubmissionFilesWithDefaults() *SubmissionFiles`

NewSubmissionFilesWithDefaults instantiates a new SubmissionFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPdf

`func (o *SubmissionFiles) GetPdf() GeneratedFile`

GetPdf returns the Pdf field if non-nil, zero value otherwise.

### GetPdfOk

`func (o *SubmissionFiles) GetPdfOk() (*GeneratedFile, bool)`

GetPdfOk returns a tuple with the Pdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdf

`func (o *SubmissionFiles) SetPdf(v GeneratedFile)`

SetPdf sets Pdf field to given value.

### HasPdf

`func (o *SubmissionFiles) HasPdf() bool`

HasPdf returns a boolean if a field has been set.

### GetAuditTrail

`func (o *SubmissionFiles) GetAuditTrail() GeneratedFile`

GetAuditTrail returns the AuditTrail field if non-nil, zero value otherwise.

### GetAuditTrailOk

`func (o *SubmissionFiles) GetAuditTrailOk() (*GeneratedFile, bool)`

GetAuditTrailOk returns a tuple with the AuditTrail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditTrail

`func (o *SubmissionFiles) SetAuditTrail(v GeneratedFile)`

SetAuditTrail sets AuditTrail field to given value.

### HasAuditTrail

`func (o *SubmissionFiles) HasAuditTrail() bool`

HasAuditTrail returns a boolean if a field has been set.

### GetPreview

`func (o *SubmissionFiles) GetPreview() GeneratedFile`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *SubmissionFiles) GetPreviewOk() (*GeneratedFile, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *SubmissionFiles) SetPreview(v GeneratedFile)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *SubmissionFiles) HasPreview() bool`

HasPreview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


