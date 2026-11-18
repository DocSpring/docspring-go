# SupportHelp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The help request's ID. | 
**State** | [**SupportHelpState**](SupportHelpState.md) | `pending` until there's an answer, then `completed`. `failed`: the question couldn't be answered just now. | 
**Answered** | **NullableBool** | Whether the documentation had an answer. Null while `state` is `pending`. | 
**Answer** | **NullableString** | The answer, as plain text. When `answered` is false, what the documentation doesn't cover. Null while `state` is `pending`. | 
**Sources** | [**[]SupportHelpSource**](SupportHelpSource.md) | The documentation pages the answer came from. | 
**Next** | **string** | What to do when the answer doesn't solve the problem, in a sentence for an AI agent to act on. | 

## Methods

### NewSupportHelp

`func NewSupportHelp(id string, state SupportHelpState, answered NullableBool, answer NullableString, sources []SupportHelpSource, next string, ) *SupportHelp`

NewSupportHelp instantiates a new SupportHelp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportHelpWithDefaults

`func NewSupportHelpWithDefaults() *SupportHelp`

NewSupportHelpWithDefaults instantiates a new SupportHelp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportHelp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportHelp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportHelp) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *SupportHelp) GetState() SupportHelpState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SupportHelp) GetStateOk() (*SupportHelpState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SupportHelp) SetState(v SupportHelpState)`

SetState sets State field to given value.


### GetAnswered

`func (o *SupportHelp) GetAnswered() bool`

GetAnswered returns the Answered field if non-nil, zero value otherwise.

### GetAnsweredOk

`func (o *SupportHelp) GetAnsweredOk() (*bool, bool)`

GetAnsweredOk returns a tuple with the Answered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswered

`func (o *SupportHelp) SetAnswered(v bool)`

SetAnswered sets Answered field to given value.


### SetAnsweredNil

`func (o *SupportHelp) SetAnsweredNil(b bool)`

 SetAnsweredNil sets the value for Answered to be an explicit nil

### UnsetAnswered
`func (o *SupportHelp) UnsetAnswered()`

UnsetAnswered ensures that no value is present for Answered, not even an explicit nil
### GetAnswer

`func (o *SupportHelp) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *SupportHelp) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *SupportHelp) SetAnswer(v string)`

SetAnswer sets Answer field to given value.


### SetAnswerNil

`func (o *SupportHelp) SetAnswerNil(b bool)`

 SetAnswerNil sets the value for Answer to be an explicit nil

### UnsetAnswer
`func (o *SupportHelp) UnsetAnswer()`

UnsetAnswer ensures that no value is present for Answer, not even an explicit nil
### GetSources

`func (o *SupportHelp) GetSources() []SupportHelpSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SupportHelp) GetSourcesOk() (*[]SupportHelpSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SupportHelp) SetSources(v []SupportHelpSource)`

SetSources sets Sources field to given value.


### GetNext

`func (o *SupportHelp) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SupportHelp) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SupportHelp) SetNext(v string)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


