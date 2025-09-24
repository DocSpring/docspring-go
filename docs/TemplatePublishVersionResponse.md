# TemplatePublishVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Result** | **map[string]interface{}** |  | 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTemplatePublishVersionResponse

`func NewTemplatePublishVersionResponse(status string, result map[string]interface{}, ) *TemplatePublishVersionResponse`

NewTemplatePublishVersionResponse instantiates a new TemplatePublishVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePublishVersionResponseWithDefaults

`func NewTemplatePublishVersionResponseWithDefaults() *TemplatePublishVersionResponse`

NewTemplatePublishVersionResponseWithDefaults instantiates a new TemplatePublishVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TemplatePublishVersionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplatePublishVersionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplatePublishVersionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResult

`func (o *TemplatePublishVersionResponse) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TemplatePublishVersionResponse) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TemplatePublishVersionResponse) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.


### GetErrors

`func (o *TemplatePublishVersionResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TemplatePublishVersionResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TemplatePublishVersionResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *TemplatePublishVersionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


