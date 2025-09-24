# CreateSubmissionDataRequestEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** |  | 
**MessageRecipient** | Pointer to **NullableString** |  | [optional] 
**MessageType** | Pointer to **NullableString** |  | [optional] 
**OccurredAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateSubmissionDataRequestEventRequest

`func NewCreateSubmissionDataRequestEventRequest(eventType string, ) *CreateSubmissionDataRequestEventRequest`

NewCreateSubmissionDataRequestEventRequest instantiates a new CreateSubmissionDataRequestEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubmissionDataRequestEventRequestWithDefaults

`func NewCreateSubmissionDataRequestEventRequestWithDefaults() *CreateSubmissionDataRequestEventRequest`

NewCreateSubmissionDataRequestEventRequestWithDefaults instantiates a new CreateSubmissionDataRequestEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *CreateSubmissionDataRequestEventRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CreateSubmissionDataRequestEventRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CreateSubmissionDataRequestEventRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetMessageRecipient

`func (o *CreateSubmissionDataRequestEventRequest) GetMessageRecipient() string`

GetMessageRecipient returns the MessageRecipient field if non-nil, zero value otherwise.

### GetMessageRecipientOk

`func (o *CreateSubmissionDataRequestEventRequest) GetMessageRecipientOk() (*string, bool)`

GetMessageRecipientOk returns a tuple with the MessageRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRecipient

`func (o *CreateSubmissionDataRequestEventRequest) SetMessageRecipient(v string)`

SetMessageRecipient sets MessageRecipient field to given value.

### HasMessageRecipient

`func (o *CreateSubmissionDataRequestEventRequest) HasMessageRecipient() bool`

HasMessageRecipient returns a boolean if a field has been set.

### SetMessageRecipientNil

`func (o *CreateSubmissionDataRequestEventRequest) SetMessageRecipientNil(b bool)`

 SetMessageRecipientNil sets the value for MessageRecipient to be an explicit nil

### UnsetMessageRecipient
`func (o *CreateSubmissionDataRequestEventRequest) UnsetMessageRecipient()`

UnsetMessageRecipient ensures that no value is present for MessageRecipient, not even an explicit nil
### GetMessageType

`func (o *CreateSubmissionDataRequestEventRequest) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *CreateSubmissionDataRequestEventRequest) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *CreateSubmissionDataRequestEventRequest) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *CreateSubmissionDataRequestEventRequest) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### SetMessageTypeNil

`func (o *CreateSubmissionDataRequestEventRequest) SetMessageTypeNil(b bool)`

 SetMessageTypeNil sets the value for MessageType to be an explicit nil

### UnsetMessageType
`func (o *CreateSubmissionDataRequestEventRequest) UnsetMessageType()`

UnsetMessageType ensures that no value is present for MessageType, not even an explicit nil
### GetOccurredAt

`func (o *CreateSubmissionDataRequestEventRequest) GetOccurredAt() string`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *CreateSubmissionDataRequestEventRequest) GetOccurredAtOk() (*string, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *CreateSubmissionDataRequestEventRequest) SetOccurredAt(v string)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *CreateSubmissionDataRequestEventRequest) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *CreateSubmissionDataRequestEventRequest) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *CreateSubmissionDataRequestEventRequest) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


