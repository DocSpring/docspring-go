# SubmissionDataRequestEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**SubmissionId** | **NullableString** |  | 
**SubmissionDataRequestId** | **NullableString** |  | 
**EventType** | **string** |  | 
**MessageType** | **NullableString** |  | 
**MessageRecipient** | **NullableString** |  | 
**OccurredAt** | **NullableString** |  | 

## Methods

### NewSubmissionDataRequestEvent

`func NewSubmissionDataRequestEvent(id NullableString, submissionId NullableString, submissionDataRequestId NullableString, eventType string, messageType NullableString, messageRecipient NullableString, occurredAt NullableString, ) *SubmissionDataRequestEvent`

NewSubmissionDataRequestEvent instantiates a new SubmissionDataRequestEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionDataRequestEventWithDefaults

`func NewSubmissionDataRequestEventWithDefaults() *SubmissionDataRequestEvent`

NewSubmissionDataRequestEventWithDefaults instantiates a new SubmissionDataRequestEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubmissionDataRequestEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionDataRequestEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionDataRequestEvent) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *SubmissionDataRequestEvent) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubmissionDataRequestEvent) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSubmissionId

`func (o *SubmissionDataRequestEvent) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *SubmissionDataRequestEvent) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *SubmissionDataRequestEvent) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.


### SetSubmissionIdNil

`func (o *SubmissionDataRequestEvent) SetSubmissionIdNil(b bool)`

 SetSubmissionIdNil sets the value for SubmissionId to be an explicit nil

### UnsetSubmissionId
`func (o *SubmissionDataRequestEvent) UnsetSubmissionId()`

UnsetSubmissionId ensures that no value is present for SubmissionId, not even an explicit nil
### GetSubmissionDataRequestId

`func (o *SubmissionDataRequestEvent) GetSubmissionDataRequestId() string`

GetSubmissionDataRequestId returns the SubmissionDataRequestId field if non-nil, zero value otherwise.

### GetSubmissionDataRequestIdOk

`func (o *SubmissionDataRequestEvent) GetSubmissionDataRequestIdOk() (*string, bool)`

GetSubmissionDataRequestIdOk returns a tuple with the SubmissionDataRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDataRequestId

`func (o *SubmissionDataRequestEvent) SetSubmissionDataRequestId(v string)`

SetSubmissionDataRequestId sets SubmissionDataRequestId field to given value.


### SetSubmissionDataRequestIdNil

`func (o *SubmissionDataRequestEvent) SetSubmissionDataRequestIdNil(b bool)`

 SetSubmissionDataRequestIdNil sets the value for SubmissionDataRequestId to be an explicit nil

### UnsetSubmissionDataRequestId
`func (o *SubmissionDataRequestEvent) UnsetSubmissionDataRequestId()`

UnsetSubmissionDataRequestId ensures that no value is present for SubmissionDataRequestId, not even an explicit nil
### GetEventType

`func (o *SubmissionDataRequestEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SubmissionDataRequestEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SubmissionDataRequestEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetMessageType

`func (o *SubmissionDataRequestEvent) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *SubmissionDataRequestEvent) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *SubmissionDataRequestEvent) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.


### SetMessageTypeNil

`func (o *SubmissionDataRequestEvent) SetMessageTypeNil(b bool)`

 SetMessageTypeNil sets the value for MessageType to be an explicit nil

### UnsetMessageType
`func (o *SubmissionDataRequestEvent) UnsetMessageType()`

UnsetMessageType ensures that no value is present for MessageType, not even an explicit nil
### GetMessageRecipient

`func (o *SubmissionDataRequestEvent) GetMessageRecipient() string`

GetMessageRecipient returns the MessageRecipient field if non-nil, zero value otherwise.

### GetMessageRecipientOk

`func (o *SubmissionDataRequestEvent) GetMessageRecipientOk() (*string, bool)`

GetMessageRecipientOk returns a tuple with the MessageRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRecipient

`func (o *SubmissionDataRequestEvent) SetMessageRecipient(v string)`

SetMessageRecipient sets MessageRecipient field to given value.


### SetMessageRecipientNil

`func (o *SubmissionDataRequestEvent) SetMessageRecipientNil(b bool)`

 SetMessageRecipientNil sets the value for MessageRecipient to be an explicit nil

### UnsetMessageRecipient
`func (o *SubmissionDataRequestEvent) UnsetMessageRecipient()`

UnsetMessageRecipient ensures that no value is present for MessageRecipient, not even an explicit nil
### GetOccurredAt

`func (o *SubmissionDataRequestEvent) GetOccurredAt() string`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *SubmissionDataRequestEvent) GetOccurredAtOk() (*string, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *SubmissionDataRequestEvent) SetOccurredAt(v string)`

SetOccurredAt sets OccurredAt field to given value.


### SetOccurredAtNil

`func (o *SubmissionDataRequestEvent) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *SubmissionDataRequestEvent) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


