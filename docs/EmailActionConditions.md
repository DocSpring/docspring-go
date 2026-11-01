# EmailActionConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | Pointer to [**TemplateFieldDependencyMatch**](TemplateFieldDependencyMatch.md) | Whether all of the conditions (the default) or any of them must be met. | [optional] 
**Conditions** | [**[]TemplateFieldDependencyCondition**](TemplateFieldDependencyCondition.md) | The conditions on other fields' values. | 

## Methods

### NewEmailActionConditions

`func NewEmailActionConditions(conditions []TemplateFieldDependencyCondition, ) *EmailActionConditions`

NewEmailActionConditions instantiates a new EmailActionConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailActionConditionsWithDefaults

`func NewEmailActionConditionsWithDefaults() *EmailActionConditions`

NewEmailActionConditionsWithDefaults instantiates a new EmailActionConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *EmailActionConditions) GetMatch() TemplateFieldDependencyMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *EmailActionConditions) GetMatchOk() (*TemplateFieldDependencyMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *EmailActionConditions) SetMatch(v TemplateFieldDependencyMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *EmailActionConditions) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetConditions

`func (o *EmailActionConditions) GetConditions() []TemplateFieldDependencyCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *EmailActionConditions) GetConditionsOk() (*[]TemplateFieldDependencyCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *EmailActionConditions) SetConditions(v []TemplateFieldDependencyCondition)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


