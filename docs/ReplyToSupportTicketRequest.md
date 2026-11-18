# ReplyToSupportTicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | The message: an answer to the team's question, more detail, or whether their fix worked. Plain text. | 
**Outcome** | Pointer to [**SupportTicketReplyOutcome**](SupportTicketReplyOutcome.md) | `resolved`: the problem is fixed, so close the ticket. `still_broken`: it isn't, so reopen the ticket. Without it, the reply answers a ticket that was waiting for one and otherwise leaves the status as it is. | [optional] 

## Methods

### NewReplyToSupportTicketRequest

`func NewReplyToSupportTicketRequest(body string, ) *ReplyToSupportTicketRequest`

NewReplyToSupportTicketRequest instantiates a new ReplyToSupportTicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplyToSupportTicketRequestWithDefaults

`func NewReplyToSupportTicketRequestWithDefaults() *ReplyToSupportTicketRequest`

NewReplyToSupportTicketRequestWithDefaults instantiates a new ReplyToSupportTicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ReplyToSupportTicketRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ReplyToSupportTicketRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ReplyToSupportTicketRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetOutcome

`func (o *ReplyToSupportTicketRequest) GetOutcome() SupportTicketReplyOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *ReplyToSupportTicketRequest) GetOutcomeOk() (*SupportTicketReplyOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *ReplyToSupportTicketRequest) SetOutcome(v SupportTicketReplyOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *ReplyToSupportTicketRequest) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


