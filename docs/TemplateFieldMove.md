# TemplateFieldMove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The field's ID. | 
**OldIndex** | **int32** | Its place in `from_version`'s field order, starting at 0. | 
**NewIndex** | **int32** | Its place in `to_version`'s field order, starting at 0. | 

## Methods

### NewTemplateFieldMove

`func NewTemplateFieldMove(id int32, oldIndex int32, newIndex int32, ) *TemplateFieldMove`

NewTemplateFieldMove instantiates a new TemplateFieldMove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldMoveWithDefaults

`func NewTemplateFieldMoveWithDefaults() *TemplateFieldMove`

NewTemplateFieldMoveWithDefaults instantiates a new TemplateFieldMove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateFieldMove) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateFieldMove) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateFieldMove) SetId(v int32)`

SetId sets Id field to given value.


### GetOldIndex

`func (o *TemplateFieldMove) GetOldIndex() int32`

GetOldIndex returns the OldIndex field if non-nil, zero value otherwise.

### GetOldIndexOk

`func (o *TemplateFieldMove) GetOldIndexOk() (*int32, bool)`

GetOldIndexOk returns a tuple with the OldIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldIndex

`func (o *TemplateFieldMove) SetOldIndex(v int32)`

SetOldIndex sets OldIndex field to given value.


### GetNewIndex

`func (o *TemplateFieldMove) GetNewIndex() int32`

GetNewIndex returns the NewIndex field if non-nil, zero value otherwise.

### GetNewIndexOk

`func (o *TemplateFieldMove) GetNewIndexOk() (*int32, bool)`

GetNewIndexOk returns a tuple with the NewIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIndex

`func (o *TemplateFieldMove) SetNewIndex(v int32)`

SetNewIndex sets NewIndex field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


