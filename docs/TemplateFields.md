# TemplateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** | The template's ID. | 
**Version** | **NullableString** | The published version the fields are of, such as `1.2.0`. Null for the draft. | 
**Fields** | [**[]Field**](Field.md) | The fields. | 

## Methods

### NewTemplateFields

`func NewTemplateFields(templateId string, version NullableString, fields []Field, ) *TemplateFields`

NewTemplateFields instantiates a new TemplateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldsWithDefaults

`func NewTemplateFieldsWithDefaults() *TemplateFields`

NewTemplateFieldsWithDefaults instantiates a new TemplateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *TemplateFields) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateFields) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateFields) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetVersion

`func (o *TemplateFields) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TemplateFields) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TemplateFields) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *TemplateFields) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TemplateFields) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetFields

`func (o *TemplateFields) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TemplateFields) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TemplateFields) SetFields(v []Field)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


