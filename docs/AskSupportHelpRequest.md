# AskSupportHelpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | **string** | What you want to know, in a sentence or two. For a follow-up question, include what was asked and answered before that matters to it: each request is answered by itself. | 
**Tool** | Pointer to **string** | The API operation, MCP tool or CLI command involved, if any. | [optional] 
**Error** | Pointer to **string** | The exact error text, if any. | [optional] 
**Conversation** | Pointer to [**[]SupportHelpTurn**](SupportHelpTurn.md) | For a chat: the questions and answers so far, oldest first, when `question` follows on from them. At most 40 turns: start a new conversation after that. | [optional] 

## Methods

### NewAskSupportHelpRequest

`func NewAskSupportHelpRequest(question string, ) *AskSupportHelpRequest`

NewAskSupportHelpRequest instantiates a new AskSupportHelpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskSupportHelpRequestWithDefaults

`func NewAskSupportHelpRequestWithDefaults() *AskSupportHelpRequest`

NewAskSupportHelpRequestWithDefaults instantiates a new AskSupportHelpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *AskSupportHelpRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *AskSupportHelpRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *AskSupportHelpRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetTool

`func (o *AskSupportHelpRequest) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *AskSupportHelpRequest) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *AskSupportHelpRequest) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *AskSupportHelpRequest) HasTool() bool`

HasTool returns a boolean if a field has been set.

### GetError

`func (o *AskSupportHelpRequest) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AskSupportHelpRequest) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AskSupportHelpRequest) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AskSupportHelpRequest) HasError() bool`

HasError returns a boolean if a field has been set.

### GetConversation

`func (o *AskSupportHelpRequest) GetConversation() []SupportHelpTurn`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *AskSupportHelpRequest) GetConversationOk() (*[]SupportHelpTurn, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *AskSupportHelpRequest) SetConversation(v []SupportHelpTurn)`

SetConversation sets Conversation field to given value.

### HasConversation

`func (o *AskSupportHelpRequest) HasConversation() bool`

HasConversation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


