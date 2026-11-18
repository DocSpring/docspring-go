# SupportTicketList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SupportTicket**](SupportTicket.md) | The items on this page. | 
**NextCursor** | **NullableString** | Pass it as `cursor` to get the next page. Null on the last page. | 

## Methods

### NewSupportTicketList

`func NewSupportTicketList(data []SupportTicket, nextCursor NullableString, ) *SupportTicketList`

NewSupportTicketList instantiates a new SupportTicketList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketListWithDefaults

`func NewSupportTicketListWithDefaults() *SupportTicketList`

NewSupportTicketListWithDefaults instantiates a new SupportTicketList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SupportTicketList) GetData() []SupportTicket`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SupportTicketList) GetDataOk() (*[]SupportTicket, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SupportTicketList) SetData(v []SupportTicket)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *SupportTicketList) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *SupportTicketList) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *SupportTicketList) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *SupportTicketList) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *SupportTicketList) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


