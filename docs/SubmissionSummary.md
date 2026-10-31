# SubmissionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The submission's ID. | 
**TemplateId** | **NullableString** | Null once the template has been permanently deleted. | 
**TemplateType** | [**TemplateType**](TemplateType.md) | The template's type: `pdf`, `html` or `docx`. | 
**TemplateVersion** | **NullableString** | The published version of the template the submission used. Null for the draft. | 
**BatchId** | **NullableString** | The ID of the batch the submission was created in. | 
**Source** | [**SubmissionSource**](SubmissionSource.md) | `api`: created with the API. `web`: a web form. `reprocess`: a copy of another submission, made to generate its PDF again. | 
**State** | [**SubmissionState**](SubmissionState.md) | `pending` until the files have been generated, then `processed`. `waiting_for_signature_requests` until every signature request is completed. `insufficient_credits`: waiting for credits; `error_message` says whether the files are generated once you add them. `account_suspended`: waiting for the account to be reactivated, for up to 30 days. Waiting ends in `error` if it's discarded or takes over 30 days. Every other state is final, with no files: `invalid_data` (the data doesn't match the template's fields: see `data_errors`), `declined` (a signature request was declined), and the states where generating the files failed, which have an `error_message`: `syntax_error` (a mistake in an HTML template's Liquid or SCSS, or in a Word template's tags), `image_download_failed` or `image_processing_failed` (an image in the data or the template couldn't be downloaded or read) and `error` (anything else). Older submissions can also be `license_revoked` or `accidental`. | 
**Test** | **bool** | Test PDFs have a watermark and are free. Live PDFs use credits. Submissions made with a test API token are always test submissions. | 
**Editable** | **bool** | `true` keeps the PDF's text and check fields fillable, and `false` flattens every field. Left out, the template decides: its `editable_submissions` setting, or each field's own Editable option. A response shows the template's setting when the request left this out. Can't be `true` with `signature_requests`. | 
**Data** | **map[string]interface{}** | The submission's data, with the values entered in completed signature requests. Null unless you pass `include_data=true`. Empty once the submission has expired. | 
**DataErrors** | [**[]ErrorDetail**](ErrorDetail.md) | Why the data doesn't match the template's fields, when `state` is `invalid_data`: the same details as the 422 response that recorded it, with pointers into the request that created the submission (`/data/name`). Cleared once the submission has expired. | 
**ErrorMessage** | **NullableString** | Why the PDF couldn't be generated, when `state` is `error`. | 
**Metadata** | **map[string]interface{}** | Your own data about the submission, such as an ID from your app. It isn't printed on the PDF. Two keys also change the PDF: `pdf_filename` names its file (`files.pdf.filename`, and its download), and `pdf_title` is its title in the PDF's document properties (and an HTML template's `&lt;title&gt;`, unless its HTML has one). Without `pdf_title`, the title is the template's name. In `pdf_filename`, a `.pdf` ending is removed, each run of characters other than letters, digits, `-`, `_` and `.` becomes one `_`, and the first 128 characters are kept, then `.pdf` is added. Without it, the file is named after the submission's ID (`sub_1234567890abcdef01.pdf`). | 
**Files** | [**SubmissionFiles**](SubmissionFiles.md) | Empty unless you pass `include_files=true`: then the generated files, as when you get the submission, without the `png` pages' links (`files.png.pages` is null). Get the submission for its page images. | 
**Expired** | **bool** | An expired submission's data, metadata, password and PDF have been deleted. | 
**ExpiresAt** | **NullableTime** | When the submission expires. Null if it's kept until you delete it. | 
**ProcessedAt** | **NullableTime** | When the PDF was generated. Null until then. | 
**CreatedAt** | **time.Time** | When the submission was created. | 
**UpdatedAt** | **time.Time** | When the submission was last updated. | 

## Methods

### NewSubmissionSummary

`func NewSubmissionSummary(id string, templateId NullableString, templateType TemplateType, templateVersion NullableString, batchId NullableString, source SubmissionSource, state SubmissionState, test bool, editable bool, data map[string]interface{}, dataErrors []ErrorDetail, errorMessage NullableString, metadata map[string]interface{}, files SubmissionFiles, expired bool, expiresAt NullableTime, processedAt NullableTime, createdAt time.Time, updatedAt time.Time, ) *SubmissionSummary`

NewSubmissionSummary instantiates a new SubmissionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionSummaryWithDefaults

`func NewSubmissionSummaryWithDefaults() *SubmissionSummary`

NewSubmissionSummaryWithDefaults instantiates a new SubmissionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubmissionSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionSummary) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *SubmissionSummary) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SubmissionSummary) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SubmissionSummary) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### SetTemplateIdNil

`func (o *SubmissionSummary) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *SubmissionSummary) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetTemplateType

`func (o *SubmissionSummary) GetTemplateType() TemplateType`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *SubmissionSummary) GetTemplateTypeOk() (*TemplateType, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *SubmissionSummary) SetTemplateType(v TemplateType)`

SetTemplateType sets TemplateType field to given value.


### GetTemplateVersion

`func (o *SubmissionSummary) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *SubmissionSummary) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *SubmissionSummary) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.


### SetTemplateVersionNil

`func (o *SubmissionSummary) SetTemplateVersionNil(b bool)`

 SetTemplateVersionNil sets the value for TemplateVersion to be an explicit nil

### UnsetTemplateVersion
`func (o *SubmissionSummary) UnsetTemplateVersion()`

UnsetTemplateVersion ensures that no value is present for TemplateVersion, not even an explicit nil
### GetBatchId

`func (o *SubmissionSummary) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *SubmissionSummary) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *SubmissionSummary) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.


### SetBatchIdNil

`func (o *SubmissionSummary) SetBatchIdNil(b bool)`

 SetBatchIdNil sets the value for BatchId to be an explicit nil

### UnsetBatchId
`func (o *SubmissionSummary) UnsetBatchId()`

UnsetBatchId ensures that no value is present for BatchId, not even an explicit nil
### GetSource

`func (o *SubmissionSummary) GetSource() SubmissionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SubmissionSummary) GetSourceOk() (*SubmissionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SubmissionSummary) SetSource(v SubmissionSource)`

SetSource sets Source field to given value.


### GetState

`func (o *SubmissionSummary) GetState() SubmissionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubmissionSummary) GetStateOk() (*SubmissionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubmissionSummary) SetState(v SubmissionState)`

SetState sets State field to given value.


### GetTest

`func (o *SubmissionSummary) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *SubmissionSummary) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *SubmissionSummary) SetTest(v bool)`

SetTest sets Test field to given value.


### GetEditable

`func (o *SubmissionSummary) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *SubmissionSummary) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *SubmissionSummary) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetData

`func (o *SubmissionSummary) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubmissionSummary) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubmissionSummary) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *SubmissionSummary) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SubmissionSummary) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataErrors

`func (o *SubmissionSummary) GetDataErrors() []ErrorDetail`

GetDataErrors returns the DataErrors field if non-nil, zero value otherwise.

### GetDataErrorsOk

`func (o *SubmissionSummary) GetDataErrorsOk() (*[]ErrorDetail, bool)`

GetDataErrorsOk returns a tuple with the DataErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataErrors

`func (o *SubmissionSummary) SetDataErrors(v []ErrorDetail)`

SetDataErrors sets DataErrors field to given value.


### GetErrorMessage

`func (o *SubmissionSummary) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SubmissionSummary) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SubmissionSummary) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *SubmissionSummary) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *SubmissionSummary) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetMetadata

`func (o *SubmissionSummary) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubmissionSummary) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubmissionSummary) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetFiles

`func (o *SubmissionSummary) GetFiles() SubmissionFiles`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SubmissionSummary) GetFilesOk() (*SubmissionFiles, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SubmissionSummary) SetFiles(v SubmissionFiles)`

SetFiles sets Files field to given value.


### GetExpired

`func (o *SubmissionSummary) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *SubmissionSummary) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *SubmissionSummary) SetExpired(v bool)`

SetExpired sets Expired field to given value.


### GetExpiresAt

`func (o *SubmissionSummary) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SubmissionSummary) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SubmissionSummary) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *SubmissionSummary) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *SubmissionSummary) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetProcessedAt

`func (o *SubmissionSummary) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *SubmissionSummary) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *SubmissionSummary) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.


### SetProcessedAtNil

`func (o *SubmissionSummary) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *SubmissionSummary) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetCreatedAt

`func (o *SubmissionSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubmissionSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubmissionSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SubmissionSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubmissionSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubmissionSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


