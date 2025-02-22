# SubmissionBatchWithSubmissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**State** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**ProcessedAt** | **NullableString** |  | 
**TotalCount** | **int32** |  | 
**PendingCount** | **int32** |  | 
**ErrorCount** | **int32** |  | 
**CompletionPercentage** | **float32** |  | 
**Submissions** | Pointer to [**[]SubmissionPreview**](SubmissionPreview.md) |  | [optional] 

## Methods

### NewSubmissionBatchWithSubmissions

`func NewSubmissionBatchWithSubmissions(id NullableString, state string, metadata map[string]interface{}, processedAt NullableString, totalCount int32, pendingCount int32, errorCount int32, completionPercentage float32, ) *SubmissionBatchWithSubmissions`

NewSubmissionBatchWithSubmissions instantiates a new SubmissionBatchWithSubmissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionBatchWithSubmissionsWithDefaults

`func NewSubmissionBatchWithSubmissionsWithDefaults() *SubmissionBatchWithSubmissions`

NewSubmissionBatchWithSubmissionsWithDefaults instantiates a new SubmissionBatchWithSubmissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubmissionBatchWithSubmissions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionBatchWithSubmissions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionBatchWithSubmissions) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *SubmissionBatchWithSubmissions) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubmissionBatchWithSubmissions) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetState

`func (o *SubmissionBatchWithSubmissions) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubmissionBatchWithSubmissions) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubmissionBatchWithSubmissions) SetState(v string)`

SetState sets State field to given value.


### GetMetadata

`func (o *SubmissionBatchWithSubmissions) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubmissionBatchWithSubmissions) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubmissionBatchWithSubmissions) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProcessedAt

`func (o *SubmissionBatchWithSubmissions) GetProcessedAt() string`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *SubmissionBatchWithSubmissions) GetProcessedAtOk() (*string, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *SubmissionBatchWithSubmissions) SetProcessedAt(v string)`

SetProcessedAt sets ProcessedAt field to given value.


### SetProcessedAtNil

`func (o *SubmissionBatchWithSubmissions) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *SubmissionBatchWithSubmissions) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetTotalCount

`func (o *SubmissionBatchWithSubmissions) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SubmissionBatchWithSubmissions) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SubmissionBatchWithSubmissions) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPendingCount

`func (o *SubmissionBatchWithSubmissions) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *SubmissionBatchWithSubmissions) GetPendingCountOk() (*int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *SubmissionBatchWithSubmissions) SetPendingCount(v int32)`

SetPendingCount sets PendingCount field to given value.


### GetErrorCount

`func (o *SubmissionBatchWithSubmissions) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *SubmissionBatchWithSubmissions) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *SubmissionBatchWithSubmissions) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.


### GetCompletionPercentage

`func (o *SubmissionBatchWithSubmissions) GetCompletionPercentage() float32`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *SubmissionBatchWithSubmissions) GetCompletionPercentageOk() (*float32, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *SubmissionBatchWithSubmissions) SetCompletionPercentage(v float32)`

SetCompletionPercentage sets CompletionPercentage field to given value.


### GetSubmissions

`func (o *SubmissionBatchWithSubmissions) GetSubmissions() []SubmissionPreview`

GetSubmissions returns the Submissions field if non-nil, zero value otherwise.

### GetSubmissionsOk

`func (o *SubmissionBatchWithSubmissions) GetSubmissionsOk() (*[]SubmissionPreview, bool)`

GetSubmissionsOk returns a tuple with the Submissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissions

`func (o *SubmissionBatchWithSubmissions) SetSubmissions(v []SubmissionPreview)`

SetSubmissions sets Submissions field to given value.

### HasSubmissions

`func (o *SubmissionBatchWithSubmissions) HasSubmissions() bool`

HasSubmissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


