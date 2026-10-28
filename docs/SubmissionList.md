# SubmissionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SubmissionSummary**](SubmissionSummary.md) | The items on this page. | 
**NextCursor** | **NullableString** | Pass it as `cursor` to get the next page. Null on the last page. | 

## Methods

### NewSubmissionList

`func NewSubmissionList(data []SubmissionSummary, nextCursor NullableString, ) *SubmissionList`

NewSubmissionList instantiates a new SubmissionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionListWithDefaults

`func NewSubmissionListWithDefaults() *SubmissionList`

NewSubmissionListWithDefaults instantiates a new SubmissionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubmissionList) GetData() []SubmissionSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubmissionList) GetDataOk() (*[]SubmissionSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubmissionList) SetData(v []SubmissionSummary)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *SubmissionList) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *SubmissionList) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *SubmissionList) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *SubmissionList) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *SubmissionList) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


