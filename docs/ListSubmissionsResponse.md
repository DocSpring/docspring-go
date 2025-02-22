# ListSubmissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Submissions** | [**[]Submission**](Submission.md) |  | 
**Limit** | **float32** |  | 
**NextCursor** | **NullableString** |  | 

## Methods

### NewListSubmissionsResponse

`func NewListSubmissionsResponse(submissions []Submission, limit float32, nextCursor NullableString, ) *ListSubmissionsResponse`

NewListSubmissionsResponse instantiates a new ListSubmissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubmissionsResponseWithDefaults

`func NewListSubmissionsResponseWithDefaults() *ListSubmissionsResponse`

NewListSubmissionsResponseWithDefaults instantiates a new ListSubmissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubmissions

`func (o *ListSubmissionsResponse) GetSubmissions() []Submission`

GetSubmissions returns the Submissions field if non-nil, zero value otherwise.

### GetSubmissionsOk

`func (o *ListSubmissionsResponse) GetSubmissionsOk() (*[]Submission, bool)`

GetSubmissionsOk returns a tuple with the Submissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissions

`func (o *ListSubmissionsResponse) SetSubmissions(v []Submission)`

SetSubmissions sets Submissions field to given value.


### GetLimit

`func (o *ListSubmissionsResponse) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListSubmissionsResponse) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListSubmissionsResponse) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetNextCursor

`func (o *ListSubmissionsResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListSubmissionsResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListSubmissionsResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *ListSubmissionsResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ListSubmissionsResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


