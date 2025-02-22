# CreateSubmissionDataRequestTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Token** | [**SubmissionDataRequestToken**](SubmissionDataRequestToken.md) |  | 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateSubmissionDataRequestTokenResponse

`func NewCreateSubmissionDataRequestTokenResponse(status string, token SubmissionDataRequestToken, ) *CreateSubmissionDataRequestTokenResponse`

NewCreateSubmissionDataRequestTokenResponse instantiates a new CreateSubmissionDataRequestTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubmissionDataRequestTokenResponseWithDefaults

`func NewCreateSubmissionDataRequestTokenResponseWithDefaults() *CreateSubmissionDataRequestTokenResponse`

NewCreateSubmissionDataRequestTokenResponseWithDefaults instantiates a new CreateSubmissionDataRequestTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateSubmissionDataRequestTokenResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateSubmissionDataRequestTokenResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateSubmissionDataRequestTokenResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToken

`func (o *CreateSubmissionDataRequestTokenResponse) GetToken() SubmissionDataRequestToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateSubmissionDataRequestTokenResponse) GetTokenOk() (*SubmissionDataRequestToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateSubmissionDataRequestTokenResponse) SetToken(v SubmissionDataRequestToken)`

SetToken sets Token field to given value.


### GetErrors

`func (o *CreateSubmissionDataRequestTokenResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateSubmissionDataRequestTokenResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateSubmissionDataRequestTokenResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateSubmissionDataRequestTokenResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


