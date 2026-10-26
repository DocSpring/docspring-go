# ReorderTemplateFieldsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldIds** | **[]int32** | The ID of every field, in the new order. Each field stays on its page, and a page's fields take the order their IDs have here: the order they're printed in (a later field prints over an earlier one where they overlap), and the order of their values in the template's data schema and forms. | 

## Methods

### NewReorderTemplateFieldsRequest

`func NewReorderTemplateFieldsRequest(fieldIds []int32, ) *ReorderTemplateFieldsRequest`

NewReorderTemplateFieldsRequest instantiates a new ReorderTemplateFieldsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReorderTemplateFieldsRequestWithDefaults

`func NewReorderTemplateFieldsRequestWithDefaults() *ReorderTemplateFieldsRequest`

NewReorderTemplateFieldsRequestWithDefaults instantiates a new ReorderTemplateFieldsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldIds

`func (o *ReorderTemplateFieldsRequest) GetFieldIds() []int32`

GetFieldIds returns the FieldIds field if non-nil, zero value otherwise.

### GetFieldIdsOk

`func (o *ReorderTemplateFieldsRequest) GetFieldIdsOk() (*[]int32, bool)`

GetFieldIdsOk returns a tuple with the FieldIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIds

`func (o *ReorderTemplateFieldsRequest) SetFieldIds(v []int32)`

SetFieldIds sets FieldIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


