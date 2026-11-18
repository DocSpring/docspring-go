# SupportTicketReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ticket** | [**SupportTicket**](SupportTicket.md) |  | 
**Message** | [**SupportMessage**](SupportMessage.md) |  | 

## Methods

### NewSupportTicketReply

`func NewSupportTicketReply(ticket SupportTicket, message SupportMessage, ) *SupportTicketReply`

NewSupportTicketReply instantiates a new SupportTicketReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketReplyWithDefaults

`func NewSupportTicketReplyWithDefaults() *SupportTicketReply`

NewSupportTicketReplyWithDefaults instantiates a new SupportTicketReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicket

`func (o *SupportTicketReply) GetTicket() SupportTicket`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *SupportTicketReply) GetTicketOk() (*SupportTicket, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *SupportTicketReply) SetTicket(v SupportTicket)`

SetTicket sets Ticket field to given value.


### GetMessage

`func (o *SupportTicketReply) GetMessage() SupportMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SupportTicketReply) GetMessageOk() (*SupportMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SupportTicketReply) SetMessage(v SupportMessage)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


