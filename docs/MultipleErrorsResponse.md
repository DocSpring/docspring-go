# MultipleErrorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Errors** | **[]string** |  | 

## Methods

### NewMultipleErrorsResponse

`func NewMultipleErrorsResponse(status string, errors []string, ) *MultipleErrorsResponse`

NewMultipleErrorsResponse instantiates a new MultipleErrorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleErrorsResponseWithDefaults

`func NewMultipleErrorsResponseWithDefaults() *MultipleErrorsResponse`

NewMultipleErrorsResponseWithDefaults instantiates a new MultipleErrorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MultipleErrorsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MultipleErrorsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MultipleErrorsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrors

`func (o *MultipleErrorsResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MultipleErrorsResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MultipleErrorsResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


