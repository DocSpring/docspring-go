# Submission422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Error** | Pointer to **string** | Single error message (for non-validation errors) | [optional] 
**Submission** | Pointer to [**SubmissionPreview**](SubmissionPreview.md) |  | [optional] 
**Errors** | Pointer to **[]string** | Array of validation error messages (when submission data is invalid) | [optional] 

## Methods

### NewSubmission422Response

`func NewSubmission422Response(status string, ) *Submission422Response`

NewSubmission422Response instantiates a new Submission422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmission422ResponseWithDefaults

`func NewSubmission422ResponseWithDefaults() *Submission422Response`

NewSubmission422ResponseWithDefaults instantiates a new Submission422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Submission422Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Submission422Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Submission422Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *Submission422Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Submission422Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Submission422Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Submission422Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSubmission

`func (o *Submission422Response) GetSubmission() SubmissionPreview`

GetSubmission returns the Submission field if non-nil, zero value otherwise.

### GetSubmissionOk

`func (o *Submission422Response) GetSubmissionOk() (*SubmissionPreview, bool)`

GetSubmissionOk returns a tuple with the Submission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmission

`func (o *Submission422Response) SetSubmission(v SubmissionPreview)`

SetSubmission sets Submission field to given value.

### HasSubmission

`func (o *Submission422Response) HasSubmission() bool`

HasSubmission returns a boolean if a field has been set.

### GetErrors

`func (o *Submission422Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Submission422Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Submission422Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Submission422Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


