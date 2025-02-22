# CreateCustomFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**CustomFile** | [**CustomFile**](CustomFile.md) |  | 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateCustomFileResponse

`func NewCreateCustomFileResponse(status string, customFile CustomFile, ) *CreateCustomFileResponse`

NewCreateCustomFileResponse instantiates a new CreateCustomFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomFileResponseWithDefaults

`func NewCreateCustomFileResponseWithDefaults() *CreateCustomFileResponse`

NewCreateCustomFileResponseWithDefaults instantiates a new CreateCustomFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateCustomFileResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCustomFileResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCustomFileResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCustomFile

`func (o *CreateCustomFileResponse) GetCustomFile() CustomFile`

GetCustomFile returns the CustomFile field if non-nil, zero value otherwise.

### GetCustomFileOk

`func (o *CreateCustomFileResponse) GetCustomFileOk() (*CustomFile, bool)`

GetCustomFileOk returns a tuple with the CustomFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFile

`func (o *CreateCustomFileResponse) SetCustomFile(v CustomFile)`

SetCustomFile sets CustomFile field to given value.


### GetErrors

`func (o *CreateCustomFileResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateCustomFileResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateCustomFileResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateCustomFileResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


