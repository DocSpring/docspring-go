# CreateSubmissionDataRequestEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Event** | [**SubmissionDataRequestEvent**](SubmissionDataRequestEvent.md) |  | 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateSubmissionDataRequestEventResponse

`func NewCreateSubmissionDataRequestEventResponse(status string, event SubmissionDataRequestEvent, ) *CreateSubmissionDataRequestEventResponse`

NewCreateSubmissionDataRequestEventResponse instantiates a new CreateSubmissionDataRequestEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubmissionDataRequestEventResponseWithDefaults

`func NewCreateSubmissionDataRequestEventResponseWithDefaults() *CreateSubmissionDataRequestEventResponse`

NewCreateSubmissionDataRequestEventResponseWithDefaults instantiates a new CreateSubmissionDataRequestEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateSubmissionDataRequestEventResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateSubmissionDataRequestEventResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateSubmissionDataRequestEventResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEvent

`func (o *CreateSubmissionDataRequestEventResponse) GetEvent() SubmissionDataRequestEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CreateSubmissionDataRequestEventResponse) GetEventOk() (*SubmissionDataRequestEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CreateSubmissionDataRequestEventResponse) SetEvent(v SubmissionDataRequestEvent)`

SetEvent sets Event field to given value.


### GetErrors

`func (o *CreateSubmissionDataRequestEventResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateSubmissionDataRequestEventResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateSubmissionDataRequestEventResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateSubmissionDataRequestEventResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


