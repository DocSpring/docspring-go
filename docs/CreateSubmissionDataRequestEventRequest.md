# CreateSubmissionDataRequestEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** |  | 
**MessageType** | Pointer to **NullableString** |  | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


