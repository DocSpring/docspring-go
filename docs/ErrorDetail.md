# ErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | What's wrong: `required`, `unknown_property`, `type`, `enum`, `format`, a JSON Schema keyword such as `minLength`, or a DocSpring code. New codes can be added. | 
**Message** | **string** | A human-readable explanation of the problem. | 
**Pointer** | Pointer to **string** | A JSON pointer to the value in the request body (RFC 6901). An empty string means the whole body. In a submission's `data_errors`, it points into the request that created the submission. | [optional] 
**Parameter** | Pointer to **string** | The query parameter or header the detail is about. | [optional] 
**SubmissionId** | Pointer to **string** | For a batch item that failed validation, the ID of the submission recorded for it. | [optional] 

## Methods

### NewErrorDetail

`func NewErrorDetail(code string, message string, ) *ErrorDetail`

NewErrorDetail instantiates a new ErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailWithDefaults

`func NewErrorDetailWithDefaults() *ErrorDetail`

NewErrorDetailWithDefaults instantiates a new ErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorDetail) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPointer

`func (o *ErrorDetail) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *ErrorDetail) GetPointerOk() (*string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *ErrorDetail) SetPointer(v string)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *ErrorDetail) HasPointer() bool`

HasPointer returns a boolean if a field has been set.

### GetParameter

`func (o *ErrorDetail) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ErrorDetail) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ErrorDetail) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *ErrorDetail) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetSubmissionId

`func (o *ErrorDetail) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *ErrorDetail) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *ErrorDetail) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.

### HasSubmissionId

`func (o *ErrorDetail) HasSubmissionId() bool`

HasSubmissionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


