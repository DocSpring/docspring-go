# AddTemplateFieldsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]FieldInput**](FieldInput.md) | The fields to add. Settings left out have their defaults, except that a field with the name of other fields takes their shared settings. Shared settings it sets change them for every field with the name. | 

## Methods

### NewAddTemplateFieldsRequest

`func NewAddTemplateFieldsRequest(fields []FieldInput, ) *AddTemplateFieldsRequest`

NewAddTemplateFieldsRequest instantiates a new AddTemplateFieldsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTemplateFieldsRequestWithDefaults

`func NewAddTemplateFieldsRequestWithDefaults() *AddTemplateFieldsRequest`

NewAddTemplateFieldsRequestWithDefaults instantiates a new AddTemplateFieldsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *AddTemplateFieldsRequest) GetFields() []FieldInput`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AddTemplateFieldsRequest) GetFieldsOk() (*[]FieldInput, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AddTemplateFieldsRequest) SetFields(v []FieldInput)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


