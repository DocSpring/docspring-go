# SupportHelpTurn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | [**SupportHelpTurnRole**](SupportHelpTurnRole.md) | `user`: a question. `assistant`: an answer. | 
**Text** | **string** | What was said. | 

## Methods

### NewSupportHelpTurn

`func NewSupportHelpTurn(role SupportHelpTurnRole, text string, ) *SupportHelpTurn`

NewSupportHelpTurn instantiates a new SupportHelpTurn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportHelpTurnWithDefaults

`func NewSupportHelpTurnWithDefaults() *SupportHelpTurn`

NewSupportHelpTurnWithDefaults instantiates a new SupportHelpTurn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *SupportHelpTurn) GetRole() SupportHelpTurnRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SupportHelpTurn) GetRoleOk() (*SupportHelpTurnRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SupportHelpTurn) SetRole(v SupportHelpTurnRole)`

SetRole sets Role field to given value.


### GetText

`func (o *SupportHelpTurn) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SupportHelpTurn) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SupportHelpTurn) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


