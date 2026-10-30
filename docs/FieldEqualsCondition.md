# FieldEqualsCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The name of the field whose value is compared. | 
**Operator** | [**FieldEqualsOperator**](FieldEqualsOperator.md) | How the field's value is compared with `value`. `in`: one of a list. `range`: from `value` to `range_max`. | 
**Value** | **interface{}** | The value the field's value is compared with. | 

## Methods

### NewFieldEqualsCondition

`func NewFieldEqualsCondition(field string, operator FieldEqualsOperator, value interface{}, ) *FieldEqualsCondition`

NewFieldEqualsCondition instantiates a new FieldEqualsCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldEqualsConditionWithDefaults

`func NewFieldEqualsConditionWithDefaults() *FieldEqualsCondition`

NewFieldEqualsConditionWithDefaults instantiates a new FieldEqualsCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldEqualsCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldEqualsCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldEqualsCondition) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *FieldEqualsCondition) GetOperator() FieldEqualsOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FieldEqualsCondition) GetOperatorOk() (*FieldEqualsOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FieldEqualsCondition) SetOperator(v FieldEqualsOperator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *FieldEqualsCondition) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldEqualsCondition) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldEqualsCondition) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *FieldEqualsCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FieldEqualsCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


