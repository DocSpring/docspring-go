# BatchGeneratePdfs201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**SubmissionBatch** | [**SubmissionBatch**](SubmissionBatch.md) |  | 
**Submissions** | **[]map[string]interface{}** |  | 

## Methods

### NewBatchGeneratePdfs201Response

`func NewBatchGeneratePdfs201Response(status string, submissionBatch SubmissionBatch, submissions []map[string]interface{}, ) *BatchGeneratePdfs201Response`

NewBatchGeneratePdfs201Response instantiates a new BatchGeneratePdfs201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchGeneratePdfs201ResponseWithDefaults

`func NewBatchGeneratePdfs201ResponseWithDefaults() *BatchGeneratePdfs201Response`

NewBatchGeneratePdfs201ResponseWithDefaults instantiates a new BatchGeneratePdfs201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BatchGeneratePdfs201Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchGeneratePdfs201Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchGeneratePdfs201Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *BatchGeneratePdfs201Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BatchGeneratePdfs201Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BatchGeneratePdfs201Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BatchGeneratePdfs201Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrors

`func (o *BatchGeneratePdfs201Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchGeneratePdfs201Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchGeneratePdfs201Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchGeneratePdfs201Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSubmissionBatch

`func (o *BatchGeneratePdfs201Response) GetSubmissionBatch() SubmissionBatch`

GetSubmissionBatch returns the SubmissionBatch field if non-nil, zero value otherwise.

### GetSubmissionBatchOk

`func (o *BatchGeneratePdfs201Response) GetSubmissionBatchOk() (*SubmissionBatch, bool)`

GetSubmissionBatchOk returns a tuple with the SubmissionBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionBatch

`func (o *BatchGeneratePdfs201Response) SetSubmissionBatch(v SubmissionBatch)`

SetSubmissionBatch sets SubmissionBatch field to given value.


### GetSubmissions

`func (o *BatchGeneratePdfs201Response) GetSubmissions() []map[string]interface{}`

GetSubmissions returns the Submissions field if non-nil, zero value otherwise.

### GetSubmissionsOk

`func (o *BatchGeneratePdfs201Response) GetSubmissionsOk() (*[]map[string]interface{}, bool)`

GetSubmissionsOk returns a tuple with the Submissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissions

`func (o *BatchGeneratePdfs201Response) SetSubmissions(v []map[string]interface{})`

SetSubmissions sets Submissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


