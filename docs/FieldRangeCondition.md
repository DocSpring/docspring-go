# FieldRangeCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The name of the field whose value is compared. | 
**Operator** | **string** | Must be `range`. | [default to "range"]
**Value** | Pointer to **interface{}** | The lower end of the range, a number (or a number as a string, as the template editor saves it). Null or left out for no lower end. A range needs `value`, `range_max` or both. | [optional] 
**RangeMax** | Pointer to **interface{}** | The upper end of the range, a number (or a number as a string, as the template editor saves it). Null or left out for no upper end. | [optional] 

## Methods

### NewFieldRangeCondition

`func NewFieldRangeCondition(field string, ) *FieldRangeCondition`

NewFieldRangeCondition instantiates a new FieldRangeCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldRangeConditionWithDefaults

`func NewFieldRangeConditionWithDefaults() *FieldRangeCondition`

NewFieldRangeConditionWithDefaults instantiates a new FieldRangeCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldRangeCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldRangeCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldRangeCondition) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *FieldRangeCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FieldRangeCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FieldRangeCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *FieldRangeCondition) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldRangeCondition) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldRangeCondition) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *FieldRangeCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *FieldRangeCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FieldRangeCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRangeMax

`func (o *FieldRangeCondition) GetRangeMax() interface{}`

GetRangeMax returns the RangeMax field if non-nil, zero value otherwise.

### GetRangeMaxOk

`func (o *FieldRangeCondition) GetRangeMaxOk() (*interface{}, bool)`

GetRangeMaxOk returns a tuple with the RangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeMax

`func (o *FieldRangeCondition) SetRangeMax(v interface{})`

SetRangeMax sets RangeMax field to given value.

### HasRangeMax

`func (o *FieldRangeCondition) HasRangeMax() bool`

HasRangeMax returns a boolean if a field has been set.

### SetRangeMaxNil

`func (o *FieldRangeCondition) SetRangeMaxNil(b bool)`

 SetRangeMaxNil sets the value for RangeMax to be an explicit nil

### UnsetRangeMax
`func (o *FieldRangeCondition) UnsetRangeMax()`

UnsetRangeMax ensures that no value is present for RangeMax, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


