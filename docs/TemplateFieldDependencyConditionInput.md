# TemplateFieldDependencyConditionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The name of the field whose value is compared. | 
**Operator** | **string** | Must be `range`. | [default to "range"]
**Value** | **interface{}** | The lower end of the range, a number (or a number as a string, as the template editor saves it). Null or left out for no lower end. A range needs `value`, `range_max` or both. | 
**RangeMax** | Pointer to **interface{}** | The upper end of the range, a number (or a number as a string, as the template editor saves it). Null or left out for no upper end. | [optional] 

## Methods

### NewTemplateFieldDependencyConditionInput

`func NewTemplateFieldDependencyConditionInput(field string, value interface{}, ) *TemplateFieldDependencyConditionInput`

NewTemplateFieldDependencyConditionInput instantiates a new TemplateFieldDependencyConditionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldDependencyConditionInputWithDefaults

`func NewTemplateFieldDependencyConditionInputWithDefaults() *TemplateFieldDependencyConditionInput`

NewTemplateFieldDependencyConditionInputWithDefaults instantiates a new TemplateFieldDependencyConditionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *TemplateFieldDependencyConditionInput) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *TemplateFieldDependencyConditionInput) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *TemplateFieldDependencyConditionInput) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *TemplateFieldDependencyConditionInput) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TemplateFieldDependencyConditionInput) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TemplateFieldDependencyConditionInput) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *TemplateFieldDependencyConditionInput) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TemplateFieldDependencyConditionInput) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TemplateFieldDependencyConditionInput) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *TemplateFieldDependencyConditionInput) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TemplateFieldDependencyConditionInput) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRangeMax

`func (o *TemplateFieldDependencyConditionInput) GetRangeMax() interface{}`

GetRangeMax returns the RangeMax field if non-nil, zero value otherwise.

### GetRangeMaxOk

`func (o *TemplateFieldDependencyConditionInput) GetRangeMaxOk() (*interface{}, bool)`

GetRangeMaxOk returns a tuple with the RangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeMax

`func (o *TemplateFieldDependencyConditionInput) SetRangeMax(v interface{})`

SetRangeMax sets RangeMax field to given value.

### HasRangeMax

`func (o *TemplateFieldDependencyConditionInput) HasRangeMax() bool`

HasRangeMax returns a boolean if a field has been set.

### SetRangeMaxNil

`func (o *TemplateFieldDependencyConditionInput) SetRangeMaxNil(b bool)`

 SetRangeMaxNil sets the value for RangeMax to be an explicit nil

### UnsetRangeMax
`func (o *TemplateFieldDependencyConditionInput) UnsetRangeMax()`

UnsetRangeMax ensures that no value is present for RangeMax, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


