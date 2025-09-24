# ErrorOrMultipleErrorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Error** | Pointer to **string** | Single error message (when only one error occurred) | [optional] 
**Errors** | Pointer to **[]string** | Array of error messages (when multiple validation errors occurred) | [optional] 

## Methods

### NewErrorOrMultipleErrorsResponse

`func NewErrorOrMultipleErrorsResponse(status string, ) *ErrorOrMultipleErrorsResponse`

NewErrorOrMultipleErrorsResponse instantiates a new ErrorOrMultipleErrorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorOrMultipleErrorsResponseWithDefaults

`func NewErrorOrMultipleErrorsResponseWithDefaults() *ErrorOrMultipleErrorsResponse`

NewErrorOrMultipleErrorsResponseWithDefaults instantiates a new ErrorOrMultipleErrorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ErrorOrMultipleErrorsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorOrMultipleErrorsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorOrMultipleErrorsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *ErrorOrMultipleErrorsResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorOrMultipleErrorsResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorOrMultipleErrorsResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ErrorOrMultipleErrorsResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorOrMultipleErrorsResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorOrMultipleErrorsResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorOrMultipleErrorsResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ErrorOrMultipleErrorsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


