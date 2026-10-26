# TemplateFieldChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The field's ID. | 
**ChangedSettings** | **[]string** | The names of the settings that changed, such as `x` or `required`. | 
**Old** | [**Field**](Field.md) |  | 
**New** | [**Field**](Field.md) |  | 

## Methods

### NewTemplateFieldChange

`func NewTemplateFieldChange(id int32, changedSettings []string, old Field, new Field, ) *TemplateFieldChange`

NewTemplateFieldChange instantiates a new TemplateFieldChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldChangeWithDefaults

`func NewTemplateFieldChangeWithDefaults() *TemplateFieldChange`

NewTemplateFieldChangeWithDefaults instantiates a new TemplateFieldChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateFieldChange) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateFieldChange) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateFieldChange) SetId(v int32)`

SetId sets Id field to given value.


### GetChangedSettings

`func (o *TemplateFieldChange) GetChangedSettings() []string`

GetChangedSettings returns the ChangedSettings field if non-nil, zero value otherwise.

### GetChangedSettingsOk

`func (o *TemplateFieldChange) GetChangedSettingsOk() (*[]string, bool)`

GetChangedSettingsOk returns a tuple with the ChangedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedSettings

`func (o *TemplateFieldChange) SetChangedSettings(v []string)`

SetChangedSettings sets ChangedSettings field to given value.


### GetOld

`func (o *TemplateFieldChange) GetOld() Field`

GetOld returns the Old field if non-nil, zero value otherwise.

### GetOldOk

`func (o *TemplateFieldChange) GetOldOk() (*Field, bool)`

GetOldOk returns a tuple with the Old field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOld

`func (o *TemplateFieldChange) SetOld(v Field)`

SetOld sets Old field to given value.


### GetNew

`func (o *TemplateFieldChange) GetNew() Field`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *TemplateFieldChange) GetNewOk() (*Field, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *TemplateFieldChange) SetNew(v Field)`

SetNew sets New field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


