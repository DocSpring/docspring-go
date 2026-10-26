# CombinedPdfList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CombinedPdf**](CombinedPdf.md) | The items on this page. | 
**NextCursor** | **NullableString** | Pass it as `cursor` to get the next page. Null on the last page. | 

## Methods

### NewCombinedPdfList

`func NewCombinedPdfList(data []CombinedPdf, nextCursor NullableString, ) *CombinedPdfList`

NewCombinedPdfList instantiates a new CombinedPdfList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinedPdfListWithDefaults

`func NewCombinedPdfListWithDefaults() *CombinedPdfList`

NewCombinedPdfListWithDefaults instantiates a new CombinedPdfList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CombinedPdfList) GetData() []CombinedPdf`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CombinedPdfList) GetDataOk() (*[]CombinedPdf, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CombinedPdfList) SetData(v []CombinedPdf)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *CombinedPdfList) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *CombinedPdfList) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *CombinedPdfList) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *CombinedPdfList) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *CombinedPdfList) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


