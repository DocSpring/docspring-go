# ErrorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**ErrorCode**](ErrorCode.md) | A machine-readable error code. Clients should handle codes they don't know: new codes can be added. `invalid_data`: a submission's data doesn't match its template's fields. The submission is still recorded. `insufficient_credits` (402): your credit balance can't cover a live PDF. The record was saved and waits for credits in the `insufficient_credits` state: its ID is in `submission_id`, `submission_batch_id` or `combined_pdf_id`. `upgrade_required` (403): your plan doesn't include this, such as more templates, another active integration (webhooks count) or the embedded template editor. `account_suspended` (403): the account is suspended, so it can't create submissions. `polling_timeout`: on a sync.api host, the files weren't ready in time. The record was created and is still being processed: its ID is in `submission_id`, `submission_batch_id`, `combined_pdf_id` or `template_id`. Without one of these IDs, the request itself got no answer in time, and the record may or may not have been created. `idempotency_key_in_use` (409): a request with the same Idempotency-Key is still being processed. Send the request again in a few seconds; when the first request has already created its record, its ID is in `submission_id`, `submission_batch_id` or `combined_pdf_id`. `idempotency_key_reused` (422): the Idempotency-Key was already used for a different request. | 
**Message** | **string** | A human-readable explanation, safe to show to developers. | 
**Details** | [**[]ErrorDetail**](ErrorDetail.md) | What was wrong with each part of the request. Empty when the error isn't about one part of it. | 
**RequestId** | **NullableString** | The ID of the request. Include it when you contact support. | 
**SubmissionId** | Pointer to **string** | When a submission's data fails validation, the submission is still recorded (state `invalid_data`) so you can see it in the dashboard. This is its ID. With `insufficient_credits`, the submission that waits for credits. With `polling_timeout` or `idempotency_key_in_use`, the submission that's still being processed. | [optional] 
**SubmissionBatchId** | Pointer to **string** | The ID of the recorded batch, when items in a batch fail validation. With `insufficient_credits`, the batch whose submissions wait for credits. With `polling_timeout` or `idempotency_key_in_use`, the batch that's still being processed. | [optional] 
**CombinedPdfId** | Pointer to **string** | With `insufficient_credits`, the combined PDF that waits for credits. With `polling_timeout` or `idempotency_key_in_use`, the combined PDF that's still being processed. | [optional] 
**TemplateId** | Pointer to **string** | With `polling_timeout`, the template whose PDF is still being processed. | [optional] 

## Methods

### NewErrorInfo

`func NewErrorInfo(code ErrorCode, message string, details []ErrorDetail, requestId NullableString, ) *ErrorInfo`

NewErrorInfo instantiates a new ErrorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorInfoWithDefaults

`func NewErrorInfoWithDefaults() *ErrorInfo`

NewErrorInfoWithDefaults instantiates a new ErrorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorInfo) GetCode() ErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorInfo) GetCodeOk() (*ErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorInfo) SetCode(v ErrorCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ErrorInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorInfo) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *ErrorInfo) GetDetails() []ErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorInfo) GetDetailsOk() (*[]ErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorInfo) SetDetails(v []ErrorDetail)`

SetDetails sets Details field to given value.


### GetRequestId

`func (o *ErrorInfo) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorInfo) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorInfo) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### SetRequestIdNil

`func (o *ErrorInfo) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *ErrorInfo) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetSubmissionId

`func (o *ErrorInfo) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *ErrorInfo) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *ErrorInfo) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.

### HasSubmissionId

`func (o *ErrorInfo) HasSubmissionId() bool`

HasSubmissionId returns a boolean if a field has been set.

### GetSubmissionBatchId

`func (o *ErrorInfo) GetSubmissionBatchId() string`

GetSubmissionBatchId returns the SubmissionBatchId field if non-nil, zero value otherwise.

### GetSubmissionBatchIdOk

`func (o *ErrorInfo) GetSubmissionBatchIdOk() (*string, bool)`

GetSubmissionBatchIdOk returns a tuple with the SubmissionBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionBatchId

`func (o *ErrorInfo) SetSubmissionBatchId(v string)`

SetSubmissionBatchId sets SubmissionBatchId field to given value.

### HasSubmissionBatchId

`func (o *ErrorInfo) HasSubmissionBatchId() bool`

HasSubmissionBatchId returns a boolean if a field has been set.

### GetCombinedPdfId

`func (o *ErrorInfo) GetCombinedPdfId() string`

GetCombinedPdfId returns the CombinedPdfId field if non-nil, zero value otherwise.

### GetCombinedPdfIdOk

`func (o *ErrorInfo) GetCombinedPdfIdOk() (*string, bool)`

GetCombinedPdfIdOk returns a tuple with the CombinedPdfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedPdfId

`func (o *ErrorInfo) SetCombinedPdfId(v string)`

SetCombinedPdfId sets CombinedPdfId field to given value.

### HasCombinedPdfId

`func (o *ErrorInfo) HasCombinedPdfId() bool`

HasCombinedPdfId returns a boolean if a field has been set.

### GetTemplateId

`func (o *ErrorInfo) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ErrorInfo) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ErrorInfo) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ErrorInfo) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


