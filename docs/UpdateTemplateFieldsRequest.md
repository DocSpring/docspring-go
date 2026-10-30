# UpdateTemplateFieldsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]FieldInput**](FieldInput.md) | The changes, each with the `id` of the field it changes. Settings left out keep their values. Changing a shared setting changes it for every field with the same name. Renaming a field moves it to the shared settings of its new name, and conditional logic and formulas that use the old name follow it when no other field has that name. | 

## Methods

### NewUpdateTemplateFieldsRequest

`func NewUpdateTemplateFieldsRequest(fields []FieldInput, ) *UpdateTemplateFieldsRequest`

NewUpdateTemplateFieldsRequest instantiates a new UpdateTemplateFieldsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTemplateFieldsRequestWithDefaults

`func NewUpdateTemplateFieldsRequestWithDefaults() *UpdateTemplateFieldsRequest`

NewUpdateTemplateFieldsRequestWithDefaults instantiates a new UpdateTemplateFieldsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *UpdateTemplateFieldsRequest) GetFields() []FieldInput`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *UpdateTemplateFieldsRequest) GetFieldsOk() (*[]FieldInput, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *UpdateTemplateFieldsRequest) SetFields(v []FieldInput)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


