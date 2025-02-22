# TemplateAddFieldsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Errors** | Pointer to **[]string** |  | [optional] 
**NewFieldIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewTemplateAddFieldsResponse

`func NewTemplateAddFieldsResponse(status string, ) *TemplateAddFieldsResponse`

NewTemplateAddFieldsResponse instantiates a new TemplateAddFieldsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateAddFieldsResponseWithDefaults

`func NewTemplateAddFieldsResponseWithDefaults() *TemplateAddFieldsResponse`

NewTemplateAddFieldsResponseWithDefaults instantiates a new TemplateAddFieldsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TemplateAddFieldsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateAddFieldsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateAddFieldsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrors

`func (o *TemplateAddFieldsResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TemplateAddFieldsResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TemplateAddFieldsResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *TemplateAddFieldsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetNewFieldIds

`func (o *TemplateAddFieldsResponse) GetNewFieldIds() []int32`

GetNewFieldIds returns the NewFieldIds field if non-nil, zero value otherwise.

### GetNewFieldIdsOk

`func (o *TemplateAddFieldsResponse) GetNewFieldIdsOk() (*[]int32, bool)`

GetNewFieldIdsOk returns a tuple with the NewFieldIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFieldIds

`func (o *TemplateAddFieldsResponse) SetNewFieldIds(v []int32)`

SetNewFieldIds sets NewFieldIds field to given value.

### HasNewFieldIds

`func (o *TemplateAddFieldsResponse) HasNewFieldIds() bool`

HasNewFieldIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


