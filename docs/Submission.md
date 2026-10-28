# Submission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The submission's ID. | 
**TemplateId** | **NullableString** | Null once the template has been permanently deleted. | 
**TemplateType** | [**TemplateType**](TemplateType.md) | The template's type: `pdf` or `html`. | 
**TemplateVersion** | **NullableString** | The published version of the template the submission used. Null for the draft. | 
**BatchId** | **NullableString** | The ID of the batch the submission was created in. | 
**Source** | [**SubmissionSource**](SubmissionSource.md) | `api`: created with the API. `web`: a web form. `reprocess`: a copy of another submission, made to generate its PDF again. | 
**State** | [**SubmissionState**](SubmissionState.md) | `pending` until the PDF has been generated, then `processed`. `waiting_for_signature_requests` until every signature request is completed. `insufficient_credits`: waiting for credits; `error_message` says whether the PDF is generated once you add them. `account_suspended`: waiting for the account to be reactivated, for up to 30 days. Waiting ends in `error` if it's discarded or takes over 30 days. Every other state is final, with no PDF: `invalid_data` (the data doesn't match the template's fields: see `data_errors`), `declined` (a signature request was declined), and the states where generating the PDF failed, which have an `error_message`: `syntax_error` (a mistake in an HTML template's Liquid or SCSS), `image_download_failed` or `image_processing_failed` (an image in the data or the template couldn't be downloaded or read) and `error` (anything else). Older submissions can also be `license_revoked` or `accidental`. | 
**Test** | **bool** | Test PDFs have a watermark and are free. Live PDFs use credits. Submissions made with a test API token are always test submissions. | 
**Editable** | **bool** | `true` keeps the PDF's text and check fields fillable, and `false` flattens every field. Left out, the template decides: its `editable_submissions` setting, or each field's own Editable option. A response shows the template's setting when the request left this out. Can't be `true` with `signature_requests`. | 
**Data** | **map[string]interface{}** | The submission's data, with the values entered in completed signature requests. Empty once the submission has expired. Null in a batch's submissions unless you pass `include_data=true`, and in the events of a webhook that leaves out submission data. | 
**DataErrors** | [**[]ErrorDetail**](ErrorDetail.md) | Why the data doesn't match the template's fields, when `state` is `invalid_data`: the same details as the 422 response that recorded it, with pointers into the request that created the submission (`/data/name`). Cleared once the submission has expired. | 
**ErrorMessage** | **NullableString** | Why the PDF couldn't be generated, when `state` is `error`. | 
**Metadata** | **map[string]interface{}** | Your own data about the submission, such as an ID from your app. It isn't printed on the PDF. Two keys also change the PDF: `pdf_filename` names its file (`files.pdf.filename`, and its download), and `pdf_title` is its title in the PDF's document properties (and an HTML template's `&lt;title&gt;`, unless its HTML has one). Without `pdf_title`, the title is the template's name. In `pdf_filename`, a `.pdf` ending is removed, each run of characters other than letters, digits, `-`, `_` and `.` becomes one `_`, and the first 128 characters are kept, then `.pdf` is added. Without it, the file is named after the submission's ID (`sub_1234567890abcdef01.pdf`). | 
**Files** | [**SubmissionFiles**](SubmissionFiles.md) | The generated files: `pdf` once the submission has been processed, `png` (the images of the PDF's pages, when `output_formats` has it: `files.png.pages[0]` is page 1), `audit_trail` once its signature requests are completed, and `preview` once a preview PDF has been generated. Empty once the submission has expired. | 
**Expired** | **bool** | An expired submission's data, metadata, password and PDF have been deleted. | 
**ExpiresAt** | **NullableTime** | When the submission expires. Null if it's kept until you delete it. | 
**ProcessedAt** | **NullableTime** | When the PDF was generated. Null until then. | 
**CreatedAt** | **time.Time** | When the submission was created. | 
**UpdatedAt** | **time.Time** | When the submission was last updated. | 
**HasPassword** | **bool** | Whether the PDF needs a password to open (`password`). | 
**SignatureRequests** | [**[]SignatureRequest**](SignatureRequest.md) | The people asked to sign the submission (and fill in their fields), in order. | 
**Actions** | [**[]SubmissionAction**](SubmissionAction.md) | The uploads of the submission's PDF to your storage (account integrations). | 
**TruncatedText** | **map[string][]string** | The text of each field that didn't fit in it, by field name. Left out (null) with `data`, and empty once the submission has expired. | 
**OutputFormats** | [**[]OutputFormat**](OutputFormat.md) | What the submission makes: `pdf`, and `png` (an image of each page of the PDF, up to 100 pages) when you ask for it. Default `["pdf"]`. POST /submissions/{submission_id}/generate_files adds `png` to a processed submission. | 
**PreviewGeneratedAt** | **NullableTime** | When the preview PDF was generated. Null if there is no preview. | 

## Methods

### NewSubmission

`func NewSubmission(id string, templateId NullableString, templateType TemplateType, templateVersion NullableString, batchId NullableString, source SubmissionSource, state SubmissionState, test bool, editable bool, data map[string]interface{}, dataErrors []ErrorDetail, errorMessage NullableString, metadata map[string]interface{}, files SubmissionFiles, expired bool, expiresAt NullableTime, processedAt NullableTime, createdAt time.Time, updatedAt time.Time, hasPassword bool, signatureRequests []SignatureRequest, actions []SubmissionAction, truncatedText map[string][]string, outputFormats []OutputFormat, previewGeneratedAt NullableTime, ) *Submission`

NewSubmission instantiates a new Submission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionWithDefaults

`func NewSubmissionWithDefaults() *Submission`

NewSubmissionWithDefaults instantiates a new Submission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *Submission) GetTemplateType() TemplateType`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *Submission) GetTemplateTypeOk() (*TemplateType, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *Submission) SetTemplateType(v TemplateType)`

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
### GetSource

`func (o *Submission) GetSource() SubmissionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Submission) GetSourceOk() (*SubmissionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Submission) SetSource(v SubmissionSource)`

SetSource sets Source field to given value.


### GetState

`func (o *Submission) GetState() SubmissionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Submission) GetStateOk() (*SubmissionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Submission) SetState(v SubmissionState)`

SetState sets State field to given value.


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
### GetDataErrors

`func (o *Submission) GetDataErrors() []ErrorDetail`

GetDataErrors returns the DataErrors field if non-nil, zero value otherwise.

### GetDataErrorsOk

`func (o *Submission) GetDataErrorsOk() (*[]ErrorDetail, bool)`

GetDataErrorsOk returns a tuple with the DataErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataErrors

`func (o *Submission) SetDataErrors(v []ErrorDetail)`

SetDataErrors sets DataErrors field to given value.


### GetErrorMessage

`func (o *Submission) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Submission) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Submission) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *Submission) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Submission) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
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


### GetFiles

`func (o *Submission) GetFiles() SubmissionFiles`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Submission) GetFilesOk() (*SubmissionFiles, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Submission) SetFiles(v SubmissionFiles)`

SetFiles sets Files field to given value.


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

`func (o *Submission) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Submission) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Submission) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *Submission) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Submission) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetProcessedAt

`func (o *Submission) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *Submission) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *Submission) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.


### SetProcessedAtNil

`func (o *Submission) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *Submission) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetCreatedAt

`func (o *Submission) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Submission) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Submission) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Submission) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Submission) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Submission) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetHasPassword

`func (o *Submission) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *Submission) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *Submission) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.


### GetSignatureRequests

`func (o *Submission) GetSignatureRequests() []SignatureRequest`

GetSignatureRequests returns the SignatureRequests field if non-nil, zero value otherwise.

### GetSignatureRequestsOk

`func (o *Submission) GetSignatureRequestsOk() (*[]SignatureRequest, bool)`

GetSignatureRequestsOk returns a tuple with the SignatureRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureRequests

`func (o *Submission) SetSignatureRequests(v []SignatureRequest)`

SetSignatureRequests sets SignatureRequests field to given value.


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


### GetTruncatedText

`func (o *Submission) GetTruncatedText() map[string][]string`

GetTruncatedText returns the TruncatedText field if non-nil, zero value otherwise.

### GetTruncatedTextOk

`func (o *Submission) GetTruncatedTextOk() (*map[string][]string, bool)`

GetTruncatedTextOk returns a tuple with the TruncatedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncatedText

`func (o *Submission) SetTruncatedText(v map[string][]string)`

SetTruncatedText sets TruncatedText field to given value.


### SetTruncatedTextNil

`func (o *Submission) SetTruncatedTextNil(b bool)`

 SetTruncatedTextNil sets the value for TruncatedText to be an explicit nil

### UnsetTruncatedText
`func (o *Submission) UnsetTruncatedText()`

UnsetTruncatedText ensures that no value is present for TruncatedText, not even an explicit nil
### GetOutputFormats

`func (o *Submission) GetOutputFormats() []OutputFormat`

GetOutputFormats returns the OutputFormats field if non-nil, zero value otherwise.

### GetOutputFormatsOk

`func (o *Submission) GetOutputFormatsOk() (*[]OutputFormat, bool)`

GetOutputFormatsOk returns a tuple with the OutputFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormats

`func (o *Submission) SetOutputFormats(v []OutputFormat)`

SetOutputFormats sets OutputFormats field to given value.


### GetPreviewGeneratedAt

`func (o *Submission) GetPreviewGeneratedAt() time.Time`

GetPreviewGeneratedAt returns the PreviewGeneratedAt field if non-nil, zero value otherwise.

### GetPreviewGeneratedAtOk

`func (o *Submission) GetPreviewGeneratedAtOk() (*time.Time, bool)`

GetPreviewGeneratedAtOk returns a tuple with the PreviewGeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewGeneratedAt

`func (o *Submission) SetPreviewGeneratedAt(v time.Time)`

SetPreviewGeneratedAt sets PreviewGeneratedAt field to given value.


### SetPreviewGeneratedAtNil

`func (o *Submission) SetPreviewGeneratedAtNil(b bool)`

 SetPreviewGeneratedAtNil sets the value for PreviewGeneratedAt to be an explicit nil

### UnsetPreviewGeneratedAt
`func (o *Submission) UnsetPreviewGeneratedAt()`

UnsetPreviewGeneratedAt ensures that no value is present for PreviewGeneratedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


