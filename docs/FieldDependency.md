# FieldDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**TemplateFieldDependencyAction**](TemplateFieldDependencyAction.md) | What the rule does to the fields it applies to when its conditions are met. | 
**Match** | Pointer to [**TemplateFieldDependencyMatch**](TemplateFieldDependencyMatch.md) | Whether all of the conditions (the default) or any of them must be met. | [optional] 
**Conditions** | [**[]TemplateFieldDependencyCondition**](TemplateFieldDependencyCondition.md) | The conditions on other fields' values. | 

## Methods

### NewFieldDependency

`func NewFieldDependency(action TemplateFieldDependencyAction, conditions []TemplateFieldDependencyCondition, ) *FieldDependency`

NewFieldDependency instantiates a new FieldDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDependencyWithDefaults

`func NewFieldDependencyWithDefaults() *FieldDependency`

NewFieldDependencyWithDefaults instantiates a new FieldDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FieldDependency) GetAction() TemplateFieldDependencyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FieldDependency) GetActionOk() (*TemplateFieldDependencyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FieldDependency) SetAction(v TemplateFieldDependencyAction)`

SetAction sets Action field to given value.


### GetMatch

`func (o *FieldDependency) GetMatch() TemplateFieldDependencyMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *FieldDependency) GetMatchOk() (*TemplateFieldDependencyMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *FieldDependency) SetMatch(v TemplateFieldDependencyMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *FieldDependency) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetConditions

`func (o *FieldDependency) GetConditions() []TemplateFieldDependencyCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *FieldDependency) GetConditionsOk() (*[]TemplateFieldDependencyCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *FieldDependency) SetConditions(v []TemplateFieldDependencyCondition)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


