# CreateSubmissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Submission** | [**SubmissionPreview**](SubmissionPreview.md) |  | 
**Error** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**Code** | Pointer to **string** | &#x60;insufficient_credits&#x60; when the response is a 402. | [optional] 

## Methods

### NewCreateSubmissionResponse

`func NewCreateSubmissionResponse(status string, submission SubmissionPreview, ) *CreateSubmissionResponse`

NewCreateSubmissionResponse instantiates a new CreateSubmissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubmissionResponseWithDefaults

`func NewCreateSubmissionResponseWithDefaults() *CreateSubmissionResponse`

NewCreateSubmissionResponseWithDefaults instantiates a new CreateSubmissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateSubmissionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateSubmissionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateSubmissionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmission

`func (o *CreateSubmissionResponse) GetSubmission() SubmissionPreview`

GetSubmission returns the Submission field if non-nil, zero value otherwise.

### GetSubmissionOk

`func (o *CreateSubmissionResponse) GetSubmissionOk() (*SubmissionPreview, bool)`

GetSubmissionOk returns a tuple with the Submission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmission

`func (o *CreateSubmissionResponse) SetSubmission(v SubmissionPreview)`

SetSubmission sets Submission field to given value.


### GetError

`func (o *CreateSubmissionResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateSubmissionResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateSubmissionResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CreateSubmissionResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrors

`func (o *CreateSubmissionResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateSubmissionResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateSubmissionResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateSubmissionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetCode

`func (o *CreateSubmissionResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateSubmissionResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateSubmissionResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateSubmissionResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


