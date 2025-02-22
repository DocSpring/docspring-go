# SuccessErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewSuccessErrorResponse

`func NewSuccessErrorResponse(status string, ) *SuccessErrorResponse`

NewSuccessErrorResponse instantiates a new SuccessErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessErrorResponseWithDefaults

`func NewSuccessErrorResponseWithDefaults() *SuccessErrorResponse`

NewSuccessErrorResponseWithDefaults instantiates a new SuccessErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SuccessErrorResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SuccessErrorResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SuccessErrorResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *SuccessErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SuccessErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SuccessErrorResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SuccessErrorResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


