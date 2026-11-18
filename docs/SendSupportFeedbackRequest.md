# SendSupportFeedbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | One line, written like an issue title. | 
**Kind** | [**SupportReportKind**](SupportReportKind.md) | What kind of problem this is. | 
**Category** | Pointer to [**SupportReportCategory**](SupportReportCategory.md) | The part of DocSpring this is about. | [optional] 
**Description** | **string** | What was tried, what was expected and what happened. | 
**Impact** | Pointer to [**SupportReportImpact**](SupportReportImpact.md) | `blocked`: the task could not be done. `worked_around`: it was done another way. `friction`: it was only slower. | [optional] 
**Task** | Pointer to **string** | What the user asked for, in a short phrase. | [optional] 
**RequestIds** | Pointer to **[]string** | The request IDs of the requests that went wrong: `error.request_id` in an error response, or any response's `X-Request-Id` header. | [optional] 
**Tool** | Pointer to **string** | The API operation, MCP tool or CLI command involved. | [optional] 
**Error** | Pointer to **string** | The exact error text. | [optional] 
**Suggestion** | Pointer to **string** | What would have helped. | [optional] 
**Source** | Pointer to [**SupportReportSource**](SupportReportSource.md) | `user`: the user asked for this to be sent. `agent` (the default): an AI agent noticed the problem itself. | [optional] 
**Session** | Pointer to **string** | A label the sender chooses for its work, such as a session ID, a branch or a task name. When an account runs several agents at once, each can list only the tickets opened with its own label. | [optional] 
**Agent** | Pointer to [**SupportReportAgent**](SupportReportAgent.md) | The AI agent sending this, as it describes itself. Stored as unverified. | [optional] 

## Methods

### NewSendSupportFeedbackRequest

`func NewSendSupportFeedbackRequest(title string, kind SupportReportKind, description string, ) *SendSupportFeedbackRequest`

NewSendSupportFeedbackRequest instantiates a new SendSupportFeedbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendSupportFeedbackRequestWithDefaults

`func NewSendSupportFeedbackRequestWithDefaults() *SendSupportFeedbackRequest`

NewSendSupportFeedbackRequestWithDefaults instantiates a new SendSupportFeedbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SendSupportFeedbackRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendSupportFeedbackRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendSupportFeedbackRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetKind

`func (o *SendSupportFeedbackRequest) GetKind() SupportReportKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SendSupportFeedbackRequest) GetKindOk() (*SupportReportKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SendSupportFeedbackRequest) SetKind(v SupportReportKind)`

SetKind sets Kind field to given value.


### GetCategory

`func (o *SendSupportFeedbackRequest) GetCategory() SupportReportCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SendSupportFeedbackRequest) GetCategoryOk() (*SupportReportCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SendSupportFeedbackRequest) SetCategory(v SupportReportCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SendSupportFeedbackRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *SendSupportFeedbackRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SendSupportFeedbackRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SendSupportFeedbackRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImpact

`func (o *SendSupportFeedbackRequest) GetImpact() SupportReportImpact`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *SendSupportFeedbackRequest) GetImpactOk() (*SupportReportImpact, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *SendSupportFeedbackRequest) SetImpact(v SupportReportImpact)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *SendSupportFeedbackRequest) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetTask

`func (o *SendSupportFeedbackRequest) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *SendSupportFeedbackRequest) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *SendSupportFeedbackRequest) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *SendSupportFeedbackRequest) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetRequestIds

`func (o *SendSupportFeedbackRequest) GetRequestIds() []string`

GetRequestIds returns the RequestIds field if non-nil, zero value otherwise.

### GetRequestIdsOk

`func (o *SendSupportFeedbackRequest) GetRequestIdsOk() (*[]string, bool)`

GetRequestIdsOk returns a tuple with the RequestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIds

`func (o *SendSupportFeedbackRequest) SetRequestIds(v []string)`

SetRequestIds sets RequestIds field to given value.

### HasRequestIds

`func (o *SendSupportFeedbackRequest) HasRequestIds() bool`

HasRequestIds returns a boolean if a field has been set.

### GetTool

`func (o *SendSupportFeedbackRequest) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *SendSupportFeedbackRequest) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *SendSupportFeedbackRequest) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *SendSupportFeedbackRequest) HasTool() bool`

HasTool returns a boolean if a field has been set.

### GetError

`func (o *SendSupportFeedbackRequest) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SendSupportFeedbackRequest) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SendSupportFeedbackRequest) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SendSupportFeedbackRequest) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSuggestion

`func (o *SendSupportFeedbackRequest) GetSuggestion() string`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *SendSupportFeedbackRequest) GetSuggestionOk() (*string, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *SendSupportFeedbackRequest) SetSuggestion(v string)`

SetSuggestion sets Suggestion field to given value.

### HasSuggestion

`func (o *SendSupportFeedbackRequest) HasSuggestion() bool`

HasSuggestion returns a boolean if a field has been set.

### GetSource

`func (o *SendSupportFeedbackRequest) GetSource() SupportReportSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SendSupportFeedbackRequest) GetSourceOk() (*SupportReportSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SendSupportFeedbackRequest) SetSource(v SupportReportSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *SendSupportFeedbackRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSession

`func (o *SendSupportFeedbackRequest) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SendSupportFeedbackRequest) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SendSupportFeedbackRequest) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *SendSupportFeedbackRequest) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetAgent

`func (o *SendSupportFeedbackRequest) GetAgent() SupportReportAgent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *SendSupportFeedbackRequest) GetAgentOk() (*SupportReportAgent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *SendSupportFeedbackRequest) SetAgent(v SupportReportAgent)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *SendSupportFeedbackRequest) HasAgent() bool`

HasAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


