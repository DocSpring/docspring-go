# FieldComparisonCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The name of the field whose value is compared. | 
**Operator** | [**FieldComparisonOperator**](FieldComparisonOperator.md) | How the field's value is compared with `value`. `in`: one of a list. `range`: from `value` to `range_max`. | 
**Value** | **interface{}** | The number the field's value is compared with, or a number as a string (as the template editor saves it). The field must be a number field. | 

## Methods

### NewFieldComparisonCondition

`func NewFieldComparisonCondition(field string, operator FieldComparisonOperator, value interface{}, ) *FieldComparisonCondition`

NewFieldComparisonCondition instantiates a new FieldComparisonCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldComparisonConditionWithDefaults

`func NewFieldComparisonConditionWithDefaults() *FieldComparisonCondition`

NewFieldComparisonConditionWithDefaults instantiates a new FieldComparisonCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldComparisonCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldComparisonCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldComparisonCondition) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *FieldComparisonCondition) GetOperator() FieldComparisonOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FieldComparisonCondition) GetOperatorOk() (*FieldComparisonOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FieldComparisonCondition) SetOperator(v FieldComparisonOperator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *FieldComparisonCondition) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldComparisonCondition) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldComparisonCondition) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *FieldComparisonCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FieldComparisonCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


