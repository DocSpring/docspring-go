# SignatureRequestEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The event's ID. | 
**SubmissionId** | **string** | The ID of the signature request's submission. | 
**SignatureRequestId** | **string** | The ID of the signature request. | 
**EventType** | [**SignatureRequestEventType**](SignatureRequestEventType.md) | What happened. | 
**MessageType** | [**NullableSignatureRequestMessageType**](SignatureRequestMessageType.md) | How you sent the request. Required for `send_request`. | 
**MessageRecipient** | **NullableString** | Where you sent the request, such as an email address or a phone number. Required for `send_request`. | 
**OccurredAt** | **time.Time** | When it happened. Default now. | 
**CreatedAt** | **time.Time** | When the event was recorded. | 

## Methods

### NewSignatureRequestEvent

`func NewSignatureRequestEvent(id string, submissionId string, signatureRequestId string, eventType SignatureRequestEventType, messageType NullableSignatureRequestMessageType, messageRecipient NullableString, occurredAt time.Time, createdAt time.Time, ) *SignatureRequestEvent`

NewSignatureRequestEvent instantiates a new SignatureRequestEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureRequestEventWithDefaults

`func NewSignatureRequestEventWithDefaults() *SignatureRequestEvent`

NewSignatureRequestEventWithDefaults instantiates a new SignatureRequestEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignatureRequestEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignatureRequestEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignatureRequestEvent) SetId(v string)`

SetId sets Id field to given value.


### GetSubmissionId

`func (o *SignatureRequestEvent) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *SignatureRequestEvent) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *SignatureRequestEvent) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.


### GetSignatureRequestId

`func (o *SignatureRequestEvent) GetSignatureRequestId() string`

GetSignatureRequestId returns the SignatureRequestId field if non-nil, zero value otherwise.

### GetSignatureRequestIdOk

`func (o *SignatureRequestEvent) GetSignatureRequestIdOk() (*string, bool)`

GetSignatureRequestIdOk returns a tuple with the SignatureRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureRequestId

`func (o *SignatureRequestEvent) SetSignatureRequestId(v string)`

SetSignatureRequestId sets SignatureRequestId field to given value.


### GetEventType

`func (o *SignatureRequestEvent) GetEventType() SignatureRequestEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SignatureRequestEvent) GetEventTypeOk() (*SignatureRequestEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SignatureRequestEvent) SetEventType(v SignatureRequestEventType)`

SetEventType sets EventType field to given value.


### GetMessageType

`func (o *SignatureRequestEvent) GetMessageType() SignatureRequestMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *SignatureRequestEvent) GetMessageTypeOk() (*SignatureRequestMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *SignatureRequestEvent) SetMessageType(v SignatureRequestMessageType)`

SetMessageType sets MessageType field to given value.


### SetMessageTypeNil

`func (o *SignatureRequestEvent) SetMessageTypeNil(b bool)`

 SetMessageTypeNil sets the value for MessageType to be an explicit nil

### UnsetMessageType
`func (o *SignatureRequestEvent) UnsetMessageType()`

UnsetMessageType ensures that no value is present for MessageType, not even an explicit nil
### GetMessageRecipient

`func (o *SignatureRequestEvent) GetMessageRecipient() string`

GetMessageRecipient returns the MessageRecipient field if non-nil, zero value otherwise.

### GetMessageRecipientOk

`func (o *SignatureRequestEvent) GetMessageRecipientOk() (*string, bool)`

GetMessageRecipientOk returns a tuple with the MessageRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRecipient

`func (o *SignatureRequestEvent) SetMessageRecipient(v string)`

SetMessageRecipient sets MessageRecipient field to given value.


### SetMessageRecipientNil

`func (o *SignatureRequestEvent) SetMessageRecipientNil(b bool)`

 SetMessageRecipientNil sets the value for MessageRecipient to be an explicit nil

### UnsetMessageRecipient
`func (o *SignatureRequestEvent) UnsetMessageRecipient()`

UnsetMessageRecipient ensures that no value is present for MessageRecipient, not even an explicit nil
### GetOccurredAt

`func (o *SignatureRequestEvent) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *SignatureRequestEvent) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *SignatureRequestEvent) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetCreatedAt

`func (o *SignatureRequestEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SignatureRequestEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SignatureRequestEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


