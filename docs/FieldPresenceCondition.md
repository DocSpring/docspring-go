# FieldPresenceCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The name of the field whose value is compared. | 
**Operator** | [**FieldPresenceOperator**](FieldPresenceOperator.md) | How the field's value is compared with `value`. `in`: one of a list. `range`: from `value` to `range_max`. | 

## Methods

### NewFieldPresenceCondition

`func NewFieldPresenceCondition(field string, operator FieldPresenceOperator, ) *FieldPresenceCondition`

NewFieldPresenceCondition instantiates a new FieldPresenceCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldPresenceConditionWithDefaults

`func NewFieldPresenceConditionWithDefaults() *FieldPresenceCondition`

NewFieldPresenceConditionWithDefaults instantiates a new FieldPresenceCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldPresenceCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldPresenceCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldPresenceCondition) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *FieldPresenceCondition) GetOperator() FieldPresenceOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FieldPresenceCondition) GetOperatorOk() (*FieldPresenceOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FieldPresenceCondition) SetOperator(v FieldPresenceOperator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


