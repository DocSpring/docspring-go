# FieldInCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The name of the field whose value is compared. | 
**Operator** | **string** | Must be `in`. | [default to "in"]
**Value** | **[]interface{}** | The values the field's value must be one of. | 

## Methods

### NewFieldInCondition

`func NewFieldInCondition(field string, value []interface{}, ) *FieldInCondition`

NewFieldInCondition instantiates a new FieldInCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldInConditionWithDefaults

`func NewFieldInConditionWithDefaults() *FieldInCondition`

NewFieldInConditionWithDefaults instantiates a new FieldInCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldInCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldInCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldInCondition) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *FieldInCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FieldInCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FieldInCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *FieldInCondition) GetValue() []interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldInCondition) GetValueOk() (*[]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldInCondition) SetValue(v []interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


