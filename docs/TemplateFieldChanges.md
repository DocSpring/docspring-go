# TemplateFieldChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | [**[]Field**](Field.md) | The fields that `to_version` has and `from_version` doesn't. | 
**Removed** | [**[]Field**](Field.md) | The fields that `from_version` has and `to_version` doesn't, as they were. | 
**Changed** | [**[]TemplateFieldChange**](TemplateFieldChange.md) | The fields with changed settings. | 
**Reordered** | [**[]TemplateFieldMove**](TemplateFieldMove.md) | The fields that moved in the template's field order (the order of its fields across all its pages, as the fields endpoint lists them): fields that swapped places, or the field that moved the furthest. The fields that a move pushed along aren't listed. | 

## Methods

### NewTemplateFieldChanges

`func NewTemplateFieldChanges(added []Field, removed []Field, changed []TemplateFieldChange, reordered []TemplateFieldMove, ) *TemplateFieldChanges`

NewTemplateFieldChanges instantiates a new TemplateFieldChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldChangesWithDefaults

`func NewTemplateFieldChangesWithDefaults() *TemplateFieldChanges`

NewTemplateFieldChangesWithDefaults instantiates a new TemplateFieldChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *TemplateFieldChanges) GetAdded() []Field`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *TemplateFieldChanges) GetAddedOk() (*[]Field, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *TemplateFieldChanges) SetAdded(v []Field)`

SetAdded sets Added field to given value.


### GetRemoved

`func (o *TemplateFieldChanges) GetRemoved() []Field`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *TemplateFieldChanges) GetRemovedOk() (*[]Field, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *TemplateFieldChanges) SetRemoved(v []Field)`

SetRemoved sets Removed field to given value.


### GetChanged

`func (o *TemplateFieldChanges) GetChanged() []TemplateFieldChange`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *TemplateFieldChanges) GetChangedOk() (*[]TemplateFieldChange, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *TemplateFieldChanges) SetChanged(v []TemplateFieldChange)`

SetChanged sets Changed field to given value.


### GetReordered

`func (o *TemplateFieldChanges) GetReordered() []TemplateFieldMove`

GetReordered returns the Reordered field if non-nil, zero value otherwise.

### GetReorderedOk

`func (o *TemplateFieldChanges) GetReorderedOk() (*[]TemplateFieldMove, bool)`

GetReorderedOk returns a tuple with the Reordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReordered

`func (o *TemplateFieldChanges) SetReordered(v []TemplateFieldMove)`

SetReordered sets Reordered field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


