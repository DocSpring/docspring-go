# FieldGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The group's name: `address` for the field `address/city`. A list's item numbers are written as `~`, so the fields `items/0/price` and `items/1/price` are in the groups `items` and `items/~`, the list's items. | 
**Type** | [**FieldGroupType**](FieldGroupType.md) | `array` for a list, whose fields' names number its items after the group's name (`items/0/price`, or `items/~/price` in an HTML template), and `object` for the other groups. | 
**Title** | **string** | The group's title in the template's data schema and on forms. When it's empty, a title is made from the group's name. | 
**Description** | **string** | Help text for the group, added to the template's data schema and shown on forms. | 
**Metadata** | **string** | Text for your own use, added to the template's data schema as `x-metadata`. DocSpring doesn't use it. | 
**MinItems** | **NullableInt32** | Lists only (`type` `array`): the fewest items submission data must have. Null for no minimum. | 
**MaxItems** | **NullableInt32** | HTML templates' lists only: the most items submission data can have. Null for no limit. A PDF template's list takes as many items as its fields are numbered for, or more with `allow_additional_items`. | 
**AllowAdditionalItems** | **bool** | PDF templates' lists only: submission data can have more items than the list's fields are numbered for. Default false. | 

## Methods

### NewFieldGroup

`func NewFieldGroup(name string, type_ FieldGroupType, title string, description string, metadata string, minItems NullableInt32, maxItems NullableInt32, allowAdditionalItems bool, ) *FieldGroup`

NewFieldGroup instantiates a new FieldGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldGroupWithDefaults

`func NewFieldGroupWithDefaults() *FieldGroup`

NewFieldGroupWithDefaults instantiates a new FieldGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FieldGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldGroup) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FieldGroup) GetType() FieldGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldGroup) GetTypeOk() (*FieldGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldGroup) SetType(v FieldGroupType)`

SetType sets Type field to given value.


### GetTitle

`func (o *FieldGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FieldGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FieldGroup) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *FieldGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMetadata

`func (o *FieldGroup) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FieldGroup) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FieldGroup) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetMinItems

`func (o *FieldGroup) GetMinItems() int32`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *FieldGroup) GetMinItemsOk() (*int32, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *FieldGroup) SetMinItems(v int32)`

SetMinItems sets MinItems field to given value.


### SetMinItemsNil

`func (o *FieldGroup) SetMinItemsNil(b bool)`

 SetMinItemsNil sets the value for MinItems to be an explicit nil

### UnsetMinItems
`func (o *FieldGroup) UnsetMinItems()`

UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
### GetMaxItems

`func (o *FieldGroup) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *FieldGroup) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *FieldGroup) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.


### SetMaxItemsNil

`func (o *FieldGroup) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *FieldGroup) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetAllowAdditionalItems

`func (o *FieldGroup) GetAllowAdditionalItems() bool`

GetAllowAdditionalItems returns the AllowAdditionalItems field if non-nil, zero value otherwise.

### GetAllowAdditionalItemsOk

`func (o *FieldGroup) GetAllowAdditionalItemsOk() (*bool, bool)`

GetAllowAdditionalItemsOk returns a tuple with the AllowAdditionalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdditionalItems

`func (o *FieldGroup) SetAllowAdditionalItems(v bool)`

SetAllowAdditionalItems sets AllowAdditionalItems field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


