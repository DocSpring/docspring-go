# FieldGroupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The group's name: `address` for the field `address/city`. A list's item numbers are written as `~`, so the fields `items/0/price` and `items/1/price` are in the groups `items` and `items/~`, the list's items. | 
**Title** | Pointer to **string** | The group's title in the template's data schema and on forms. When it's empty, a title is made from the group's name. | [optional] 
**Description** | Pointer to **string** | Help text for the group, added to the template's data schema and shown on forms. | [optional] 
**Metadata** | Pointer to **string** | Text for your own use, added to the template's data schema as `x-metadata`. DocSpring doesn't use it. | [optional] 
**MinItems** | Pointer to **NullableInt32** | Lists only (`type` `array`): the fewest items submission data must have. Null for no minimum. | [optional] 
**MaxItems** | Pointer to **NullableInt32** | HTML templates' lists only: the most items submission data can have. Null for no limit. A PDF template's list takes as many items as its fields are numbered for, or more with `allow_additional_items`. | [optional] 
**AllowAdditionalItems** | Pointer to **bool** | PDF templates' lists only: submission data can have more items than the list's fields are numbered for. Default false. | [optional] 

## Methods

### NewFieldGroupInput

`func NewFieldGroupInput(name string, ) *FieldGroupInput`

NewFieldGroupInput instantiates a new FieldGroupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldGroupInputWithDefaults

`func NewFieldGroupInputWithDefaults() *FieldGroupInput`

NewFieldGroupInputWithDefaults instantiates a new FieldGroupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FieldGroupInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldGroupInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldGroupInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *FieldGroupInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FieldGroupInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FieldGroupInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FieldGroupInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *FieldGroupInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldGroupInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldGroupInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldGroupInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *FieldGroupInput) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FieldGroupInput) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FieldGroupInput) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FieldGroupInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMinItems

`func (o *FieldGroupInput) GetMinItems() int32`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *FieldGroupInput) GetMinItemsOk() (*int32, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *FieldGroupInput) SetMinItems(v int32)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *FieldGroupInput) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### SetMinItemsNil

`func (o *FieldGroupInput) SetMinItemsNil(b bool)`

 SetMinItemsNil sets the value for MinItems to be an explicit nil

### UnsetMinItems
`func (o *FieldGroupInput) UnsetMinItems()`

UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
### GetMaxItems

`func (o *FieldGroupInput) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *FieldGroupInput) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *FieldGroupInput) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *FieldGroupInput) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *FieldGroupInput) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *FieldGroupInput) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetAllowAdditionalItems

`func (o *FieldGroupInput) GetAllowAdditionalItems() bool`

GetAllowAdditionalItems returns the AllowAdditionalItems field if non-nil, zero value otherwise.

### GetAllowAdditionalItemsOk

`func (o *FieldGroupInput) GetAllowAdditionalItemsOk() (*bool, bool)`

GetAllowAdditionalItemsOk returns a tuple with the AllowAdditionalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdditionalItems

`func (o *FieldGroupInput) SetAllowAdditionalItems(v bool)`

SetAllowAdditionalItems sets AllowAdditionalItems field to given value.

### HasAllowAdditionalItems

`func (o *FieldGroupInput) HasAllowAdditionalItems() bool`

HasAllowAdditionalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


