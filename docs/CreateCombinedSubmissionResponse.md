# CreateCombinedSubmissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**CombinedSubmission** | [**CombinedSubmission**](CombinedSubmission.md) |  | 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateCombinedSubmissionResponse

`func NewCreateCombinedSubmissionResponse(status string, combinedSubmission CombinedSubmission, ) *CreateCombinedSubmissionResponse`

NewCreateCombinedSubmissionResponse instantiates a new CreateCombinedSubmissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCombinedSubmissionResponseWithDefaults

`func NewCreateCombinedSubmissionResponseWithDefaults() *CreateCombinedSubmissionResponse`

NewCreateCombinedSubmissionResponseWithDefaults instantiates a new CreateCombinedSubmissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateCombinedSubmissionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCombinedSubmissionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCombinedSubmissionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCombinedSubmission

`func (o *CreateCombinedSubmissionResponse) GetCombinedSubmission() CombinedSubmission`

GetCombinedSubmission returns the CombinedSubmission field if non-nil, zero value otherwise.

### GetCombinedSubmissionOk

`func (o *CreateCombinedSubmissionResponse) GetCombinedSubmissionOk() (*CombinedSubmission, bool)`

GetCombinedSubmissionOk returns a tuple with the CombinedSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedSubmission

`func (o *CreateCombinedSubmissionResponse) SetCombinedSubmission(v CombinedSubmission)`

SetCombinedSubmission sets CombinedSubmission field to given value.


### GetErrors

`func (o *CreateCombinedSubmissionResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateCombinedSubmissionResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateCombinedSubmissionResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateCombinedSubmissionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


