# SubmissionBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The batch's ID. | 
**State** | [**SubmissionBatchState**](SubmissionBatchState.md) | `pending` while the PDFs are being generated, then `processed`. `error` as soon as a submission fails, while the others can still be pending: the batch has finished once `processed_at` is set. | 
**Metadata** | **map[string]interface{}** | Your own data about the batch, such as an ID from your app. | 
**TotalCount** | **int32** | The number of submissions in the request. | 
**PendingCount** | **int32** | The submissions whose PDFs haven't been generated yet. | 
**ErrorCount** | **int32** | The submissions that were invalid or whose PDFs couldn't be generated. | 
**CompletionPercentage** | **int32** | The percentage of the batch's submissions that have finished, from 0 to 100. | 
**Submissions** | [**[]Submission**](Submission.md) | The batch's submissions, in the order of the request. Their `data` is null unless you pass `include_data=true`. A batch that failed validation only has the submissions that were recorded as `invalid_data`. | 
**ProcessedAt** | **NullableTime** | When every submission had finished. | 
**CreatedAt** | **time.Time** | When the batch was created. | 
**UpdatedAt** | **time.Time** | When the batch was last updated. | 

## Methods

### NewSubmissionBatch

`func NewSubmissionBatch(id string, state SubmissionBatchState, metadata map[string]interface{}, totalCount int32, pendingCount int32, errorCount int32, completionPercentage int32, submissions []Submission, processedAt NullableTime, createdAt time.Time, updatedAt time.Time, ) *SubmissionBatch`

NewSubmissionBatch instantiates a new SubmissionBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionBatchWithDefaults

`func NewSubmissionBatchWithDefaults() *SubmissionBatch`

NewSubmissionBatchWithDefaults instantiates a new SubmissionBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubmissionBatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionBatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionBatch) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *SubmissionBatch) GetState() SubmissionBatchState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubmissionBatch) GetStateOk() (*SubmissionBatchState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubmissionBatch) SetState(v SubmissionBatchState)`

SetState sets State field to given value.


### GetMetadata

`func (o *SubmissionBatch) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubmissionBatch) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubmissionBatch) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetTotalCount

`func (o *SubmissionBatch) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SubmissionBatch) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SubmissionBatch) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPendingCount

`func (o *SubmissionBatch) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *SubmissionBatch) GetPendingCountOk() (*int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *SubmissionBatch) SetPendingCount(v int32)`

SetPendingCount sets PendingCount field to given value.


### GetErrorCount

`func (o *SubmissionBatch) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *SubmissionBatch) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *SubmissionBatch) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.


### GetCompletionPercentage

`func (o *SubmissionBatch) GetCompletionPercentage() int32`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *SubmissionBatch) GetCompletionPercentageOk() (*int32, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *SubmissionBatch) SetCompletionPercentage(v int32)`

SetCompletionPercentage sets CompletionPercentage field to given value.


### GetSubmissions

`func (o *SubmissionBatch) GetSubmissions() []Submission`

GetSubmissions returns the Submissions field if non-nil, zero value otherwise.

### GetSubmissionsOk

`func (o *SubmissionBatch) GetSubmissionsOk() (*[]Submission, bool)`

GetSubmissionsOk returns a tuple with the Submissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissions

`func (o *SubmissionBatch) SetSubmissions(v []Submission)`

SetSubmissions sets Submissions field to given value.


### GetProcessedAt

`func (o *SubmissionBatch) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *SubmissionBatch) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *SubmissionBatch) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.


### SetProcessedAtNil

`func (o *SubmissionBatch) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *SubmissionBatch) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetCreatedAt

`func (o *SubmissionBatch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubmissionBatch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubmissionBatch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SubmissionBatch) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubmissionBatch) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubmissionBatch) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


