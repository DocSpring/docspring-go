# TemplateFieldDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**TemplateFieldDependencyAction**](TemplateFieldDependencyAction.md) | What the rule does to the fields it applies to when its conditions are met. | 
**Match** | Pointer to [**TemplateFieldDependencyMatch**](TemplateFieldDependencyMatch.md) | Whether all of the conditions (the default) or any of them must be met. | [optional] 
**Conditions** | [**[]TemplateFieldDependencyCondition**](TemplateFieldDependencyCondition.md) | The conditions on other fields' values. | 
**Targets** | **[]string** | The names of the fields the rule shows or makes required. | 
**Group** | Pointer to **string** | Marks a rule made for a field group in the template editor. Its targets outside the group are removed when the template is saved, and the rule is removed once the group has no fields. | [optional] 

## Methods

### NewTemplateFieldDependency

`func NewTemplateFieldDependency(action TemplateFieldDependencyAction, conditions []TemplateFieldDependencyCondition, targets []string, ) *TemplateFieldDependency`

NewTemplateFieldDependency instantiates a new TemplateFieldDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldDependencyWithDefaults

`func NewTemplateFieldDependencyWithDefaults() *TemplateFieldDependency`

NewTemplateFieldDependencyWithDefaults instantiates a new TemplateFieldDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TemplateFieldDependency) GetAction() TemplateFieldDependencyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TemplateFieldDependency) GetActionOk() (*TemplateFieldDependencyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TemplateFieldDependency) SetAction(v TemplateFieldDependencyAction)`

SetAction sets Action field to given value.


### GetMatch

`func (o *TemplateFieldDependency) GetMatch() TemplateFieldDependencyMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *TemplateFieldDependency) GetMatchOk() (*TemplateFieldDependencyMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *TemplateFieldDependency) SetMatch(v TemplateFieldDependencyMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *TemplateFieldDependency) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetConditions

`func (o *TemplateFieldDependency) GetConditions() []TemplateFieldDependencyCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *TemplateFieldDependency) GetConditionsOk() (*[]TemplateFieldDependencyCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *TemplateFieldDependency) SetConditions(v []TemplateFieldDependencyCondition)`

SetConditions sets Conditions field to given value.


### GetTargets

`func (o *TemplateFieldDependency) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *TemplateFieldDependency) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *TemplateFieldDependency) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetGroup

`func (o *TemplateFieldDependency) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *TemplateFieldDependency) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *TemplateFieldDependency) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *TemplateFieldDependency) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


