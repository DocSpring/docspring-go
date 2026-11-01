# EmailDomainList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EmailDomain**](EmailDomain.md) | The items on this page. | 
**NextCursor** | **NullableString** | Pass it as `cursor` to get the next page. Null on the last page. | 

## Methods

### NewEmailDomainList

`func NewEmailDomainList(data []EmailDomain, nextCursor NullableString, ) *EmailDomainList`

NewEmailDomainList instantiates a new EmailDomainList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDomainListWithDefaults

`func NewEmailDomainListWithDefaults() *EmailDomainList`

NewEmailDomainListWithDefaults instantiates a new EmailDomainList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EmailDomainList) GetData() []EmailDomain`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmailDomainList) GetDataOk() (*[]EmailDomain, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmailDomainList) SetData(v []EmailDomain)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *EmailDomainList) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *EmailDomainList) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *EmailDomainList) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *EmailDomainList) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *EmailDomainList) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


