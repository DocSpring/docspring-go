# EmailActionConditionsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | Pointer to [**TemplateFieldDependencyMatch**](TemplateFieldDependencyMatch.md) | Whether all of the conditions (the default) or any of them must be met. | [optional] 
**Conditions** | [**[]TemplateFieldDependencyConditionInput**](TemplateFieldDependencyConditionInput.md) | The conditions on other fields' values. | 

## Methods

### NewEmailActionConditionsInput

`func NewEmailActionConditionsInput(conditions []TemplateFieldDependencyConditionInput, ) *EmailActionConditionsInput`

NewEmailActionConditionsInput instantiates a new EmailActionConditionsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailActionConditionsInputWithDefaults

`func NewEmailActionConditionsInputWithDefaults() *EmailActionConditionsInput`

NewEmailActionConditionsInputWithDefaults instantiates a new EmailActionConditionsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *EmailActionConditionsInput) GetMatch() TemplateFieldDependencyMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *EmailActionConditionsInput) GetMatchOk() (*TemplateFieldDependencyMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *EmailActionConditionsInput) SetMatch(v TemplateFieldDependencyMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *EmailActionConditionsInput) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetConditions

`func (o *EmailActionConditionsInput) GetConditions() []TemplateFieldDependencyConditionInput`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *EmailActionConditionsInput) GetConditionsOk() (*[]TemplateFieldDependencyConditionInput, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *EmailActionConditionsInput) SetConditions(v []TemplateFieldDependencyConditionInput)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


