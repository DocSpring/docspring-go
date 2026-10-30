# CreateSignatureRequestEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**SignatureRequestEventType**](SignatureRequestEventType.md) | `send_request`: you sent the signature request to the person (`message_type` and `message_recipient` are required). It's the only event the API records: DocSpring records everything that happens in the signing form itself, and people decline a signature request in its form. | 
**MessageType** | Pointer to [**NullableSignatureRequestMessageType**](SignatureRequestMessageType.md) | How you sent the request. Required for `send_request`. | [optional] 
**MessageRecipient** | Pointer to **NullableString** | Where you sent the request, such as an email address or a phone number. Required for `send_request`. | [optional] 
**OccurredAt** | Pointer to **time.Time** | When it happened. Default now. | [optional] 

## Methods

### NewCreateSignatureRequestEventRequest

`func NewCreateSignatureRequestEventRequest(eventType SignatureRequestEventType, ) *CreateSignatureRequestEventRequest`

NewCreateSignatureRequestEventRequest instantiates a new CreateSignatureRequestEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSignatureRequestEventRequestWithDefaults

`func NewCreateSignatureRequestEventRequestWithDefaults() *CreateSignatureRequestEventRequest`

NewCreateSignatureRequestEventRequestWithDefaults instantiates a new CreateSignatureRequestEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *CreateSignatureRequestEventRequest) GetEventType() SignatureRequestEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CreateSignatureRequestEventRequest) GetEventTypeOk() (*SignatureRequestEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CreateSignatureRequestEventRequest) SetEventType(v SignatureRequestEventType)`

SetEventType sets EventType field to given value.


### GetMessageType

`func (o *CreateSignatureRequestEventRequest) GetMessageType() SignatureRequestMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *CreateSignatureRequestEventRequest) GetMessageTypeOk() (*SignatureRequestMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *CreateSignatureRequestEventRequest) SetMessageType(v SignatureRequestMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *CreateSignatureRequestEventRequest) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### SetMessageTypeNil

`func (o *CreateSignatureRequestEventRequest) SetMessageTypeNil(b bool)`

 SetMessageTypeNil sets the value for MessageType to be an explicit nil

### UnsetMessageType
`func (o *CreateSignatureRequestEventRequest) UnsetMessageType()`

UnsetMessageType ensures that no value is present for MessageType, not even an explicit nil
### GetMessageRecipient

`func (o *CreateSignatureRequestEventRequest) GetMessageRecipient() string`

GetMessageRecipient returns the MessageRecipient field if non-nil, zero value otherwise.

### GetMessageRecipientOk

`func (o *CreateSignatureRequestEventRequest) GetMessageRecipientOk() (*string, bool)`

GetMessageRecipientOk returns a tuple with the MessageRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRecipient

`func (o *CreateSignatureRequestEventRequest) SetMessageRecipient(v string)`

SetMessageRecipient sets MessageRecipient field to given value.

### HasMessageRecipient

`func (o *CreateSignatureRequestEventRequest) HasMessageRecipient() bool`

HasMessageRecipient returns a boolean if a field has been set.

### SetMessageRecipientNil

`func (o *CreateSignatureRequestEventRequest) SetMessageRecipientNil(b bool)`

 SetMessageRecipientNil sets the value for MessageRecipient to be an explicit nil

### UnsetMessageRecipient
`func (o *CreateSignatureRequestEventRequest) UnsetMessageRecipient()`

UnsetMessageRecipient ensures that no value is present for MessageRecipient, not even an explicit nil
### GetOccurredAt

`func (o *CreateSignatureRequestEventRequest) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *CreateSignatureRequestEventRequest) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *CreateSignatureRequestEventRequest) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *CreateSignatureRequestEventRequest) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


