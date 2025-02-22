# CreateSubmissionDataRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**DataRequest** | [**SubmissionDataRequestShow**](SubmissionDataRequestShow.md) |  | 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateSubmissionDataRequestResponse

`func NewCreateSubmissionDataRequestResponse(status string, dataRequest SubmissionDataRequestShow, ) *CreateSubmissionDataRequestResponse`

NewCreateSubmissionDataRequestResponse instantiates a new CreateSubmissionDataRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubmissionDataRequestResponseWithDefaults

`func NewCreateSubmissionDataRequestResponseWithDefaults() *CreateSubmissionDataRequestResponse`

NewCreateSubmissionDataRequestResponseWithDefaults instantiates a new CreateSubmissionDataRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateSubmissionDataRequestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateSubmissionDataRequestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateSubmissionDataRequestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDataRequest

`func (o *CreateSubmissionDataRequestResponse) GetDataRequest() SubmissionDataRequestShow`

GetDataRequest returns the DataRequest field if non-nil, zero value otherwise.

### GetDataRequestOk

`func (o *CreateSubmissionDataRequestResponse) GetDataRequestOk() (*SubmissionDataRequestShow, bool)`

GetDataRequestOk returns a tuple with the DataRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRequest

`func (o *CreateSubmissionDataRequestResponse) SetDataRequest(v SubmissionDataRequestShow)`

SetDataRequest sets DataRequest field to given value.


### GetErrors

`func (o *CreateSubmissionDataRequestResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateSubmissionDataRequestResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateSubmissionDataRequestResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateSubmissionDataRequestResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


