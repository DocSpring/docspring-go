# SuccessMultipleErrorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSuccessMultipleErrorsResponse

`func NewSuccessMultipleErrorsResponse(status string, ) *SuccessMultipleErrorsResponse`

NewSuccessMultipleErrorsResponse instantiates a new SuccessMultipleErrorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessMultipleErrorsResponseWithDefaults

`func NewSuccessMultipleErrorsResponseWithDefaults() *SuccessMultipleErrorsResponse`

NewSuccessMultipleErrorsResponseWithDefaults instantiates a new SuccessMultipleErrorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SuccessMultipleErrorsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SuccessMultipleErrorsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SuccessMultipleErrorsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrors

`func (o *SuccessMultipleErrorsResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SuccessMultipleErrorsResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SuccessMultipleErrorsResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SuccessMultipleErrorsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


