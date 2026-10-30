# TemplateFieldDependencyCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The name of the field whose value is compared. | 
**Operator** | [**TemplateFieldDependencyOperator**](TemplateFieldDependencyOperator.md) | How the field's value is compared with `value`. `in`: one of a list. `range`: from `value` to `range_max`. | 
**Value** | Pointer to **interface{}** | The value to compare the field's value with. | [optional] 
**RangeMax** | Pointer to **interface{}** | The upper end of a `range`, a number (or a number as a string, as the template editor saves it). | [optional] 

## Methods

### NewTemplateFieldDependencyCondition

`func NewTemplateFieldDependencyCondition(field string, operator TemplateFieldDependencyOperator, ) *TemplateFieldDependencyCondition`

NewTemplateFieldDependencyCondition instantiates a new TemplateFieldDependencyCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldDependencyConditionWithDefaults

`func NewTemplateFieldDependencyConditionWithDefaults() *TemplateFieldDependencyCondition`

NewTemplateFieldDependencyConditionWithDefaults instantiates a new TemplateFieldDependencyCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *TemplateFieldDependencyCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *TemplateFieldDependencyCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *TemplateFieldDependencyCondition) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *TemplateFieldDependencyCondition) GetOperator() TemplateFieldDependencyOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TemplateFieldDependencyCondition) GetOperatorOk() (*TemplateFieldDependencyOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TemplateFieldDependencyCondition) SetOperator(v TemplateFieldDependencyOperator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *TemplateFieldDependencyCondition) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TemplateFieldDependencyCondition) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TemplateFieldDependencyCondition) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *TemplateFieldDependencyCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TemplateFieldDependencyCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TemplateFieldDependencyCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRangeMax

`func (o *TemplateFieldDependencyCondition) GetRangeMax() interface{}`

GetRangeMax returns the RangeMax field if non-nil, zero value otherwise.

### GetRangeMaxOk

`func (o *TemplateFieldDependencyCondition) GetRangeMaxOk() (*interface{}, bool)`

GetRangeMaxOk returns a tuple with the RangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeMax

`func (o *TemplateFieldDependencyCondition) SetRangeMax(v interface{})`

SetRangeMax sets RangeMax field to given value.

### HasRangeMax

`func (o *TemplateFieldDependencyCondition) HasRangeMax() bool`

HasRangeMax returns a boolean if a field has been set.

### SetRangeMaxNil

`func (o *TemplateFieldDependencyCondition) SetRangeMaxNil(b bool)`

 SetRangeMaxNil sets the value for RangeMax to be an explicit nil

### UnsetRangeMax
`func (o *TemplateFieldDependencyCondition) UnsetRangeMax()`

UnsetRangeMax ensures that no value is present for RangeMax, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


