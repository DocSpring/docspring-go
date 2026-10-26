# TemplateSummaryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TemplateSummary**](TemplateSummary.md) | The items on this page. | 
**NextCursor** | **NullableString** | Pass it as `cursor` to get the next page. Null on the last page. | 

## Methods

### NewTemplateSummaryList

`func NewTemplateSummaryList(data []TemplateSummary, nextCursor NullableString, ) *TemplateSummaryList`

NewTemplateSummaryList instantiates a new TemplateSummaryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateSummaryListWithDefaults

`func NewTemplateSummaryListWithDefaults() *TemplateSummaryList`

NewTemplateSummaryListWithDefaults instantiates a new TemplateSummaryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TemplateSummaryList) GetData() []TemplateSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TemplateSummaryList) GetDataOk() (*[]TemplateSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TemplateSummaryList) SetData(v []TemplateSummary)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *TemplateSummaryList) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *TemplateSummaryList) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *TemplateSummaryList) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *TemplateSummaryList) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *TemplateSummaryList) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


