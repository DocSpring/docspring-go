# TemplateFieldDependencyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**TemplateFieldDependencyAction**](TemplateFieldDependencyAction.md) | What the rule does to the fields it applies to when its conditions are met. | 
**Match** | Pointer to [**TemplateFieldDependencyMatch**](TemplateFieldDependencyMatch.md) | Whether all of the conditions (the default) or any of them must be met. | [optional] 
**Conditions** | [**[]TemplateFieldDependencyConditionInput**](TemplateFieldDependencyConditionInput.md) | The conditions on other fields' values. | 
**Targets** | **[]string** | The names of the fields the rule shows or makes required. | 
**Group** | Pointer to **string** | Marks a rule made for a field group in the template editor. Its targets outside the group are removed when the template is saved, and the rule is removed once the group has no fields. | [optional] 

## Methods

### NewTemplateFieldDependencyInput

`func NewTemplateFieldDependencyInput(action TemplateFieldDependencyAction, conditions []TemplateFieldDependencyConditionInput, targets []string, ) *TemplateFieldDependencyInput`

NewTemplateFieldDependencyInput instantiates a new TemplateFieldDependencyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldDependencyInputWithDefaults

`func NewTemplateFieldDependencyInputWithDefaults() *TemplateFieldDependencyInput`

NewTemplateFieldDependencyInputWithDefaults instantiates a new TemplateFieldDependencyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TemplateFieldDependencyInput) GetAction() TemplateFieldDependencyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TemplateFieldDependencyInput) GetActionOk() (*TemplateFieldDependencyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TemplateFieldDependencyInput) SetAction(v TemplateFieldDependencyAction)`

SetAction sets Action field to given value.


### GetMatch

`func (o *TemplateFieldDependencyInput) GetMatch() TemplateFieldDependencyMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *TemplateFieldDependencyInput) GetMatchOk() (*TemplateFieldDependencyMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *TemplateFieldDependencyInput) SetMatch(v TemplateFieldDependencyMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *TemplateFieldDependencyInput) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetConditions

`func (o *TemplateFieldDependencyInput) GetConditions() []TemplateFieldDependencyConditionInput`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *TemplateFieldDependencyInput) GetConditionsOk() (*[]TemplateFieldDependencyConditionInput, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *TemplateFieldDependencyInput) SetConditions(v []TemplateFieldDependencyConditionInput)`

SetConditions sets Conditions field to given value.


### GetTargets

`func (o *TemplateFieldDependencyInput) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *TemplateFieldDependencyInput) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *TemplateFieldDependencyInput) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetGroup

`func (o *TemplateFieldDependencyInput) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *TemplateFieldDependencyInput) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *TemplateFieldDependencyInput) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *TemplateFieldDependencyInput) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


