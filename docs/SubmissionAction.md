# SubmissionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The upload's ID. | 
**IntegrationId** | **NullableString** | The ID of the account integration it used. | 
**State** | [**SubmissionActionState**](SubmissionActionState.md) | `pending` until the upload has run, then `processed`. `failed`: your storage rejected it. `error`: it failed in DocSpring. | 
**ActionType** | [**SubmissionActionType**](SubmissionActionType.md) | The kind of storage, such as `aws_s3_upload` for an S3 bucket. | 
**ResultData** | **map[string]interface{}** | Where the PDF was uploaded, such as the S3 bucket and key. | 
**ErrorMessage** | **NullableString** | Why the upload failed, when `state` is `failed` or `error`. | 

## Methods

### NewSubmissionAction

`func NewSubmissionAction(id string, integrationId NullableString, state SubmissionActionState, actionType SubmissionActionType, resultData map[string]interface{}, errorMessage NullableString, ) *SubmissionAction`

NewSubmissionAction instantiates a new SubmissionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionActionWithDefaults

`func NewSubmissionActionWithDefaults() *SubmissionAction`

NewSubmissionActionWithDefaults instantiates a new SubmissionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubmissionAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionAction) SetId(v string)`

SetId sets Id field to given value.


### GetIntegrationId

`func (o *SubmissionAction) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *SubmissionAction) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *SubmissionAction) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### SetIntegrationIdNil

`func (o *SubmissionAction) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *SubmissionAction) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetState

`func (o *SubmissionAction) GetState() SubmissionActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubmissionAction) GetStateOk() (*SubmissionActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubmissionAction) SetState(v SubmissionActionState)`

SetState sets State field to given value.


### GetActionType

`func (o *SubmissionAction) GetActionType() SubmissionActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *SubmissionAction) GetActionTypeOk() (*SubmissionActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *SubmissionAction) SetActionType(v SubmissionActionType)`

SetActionType sets ActionType field to given value.


### GetResultData

`func (o *SubmissionAction) GetResultData() map[string]interface{}`

GetResultData returns the ResultData field if non-nil, zero value otherwise.

### GetResultDataOk

`func (o *SubmissionAction) GetResultDataOk() (*map[string]interface{}, bool)`

GetResultDataOk returns a tuple with the ResultData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultData

`func (o *SubmissionAction) SetResultData(v map[string]interface{})`

SetResultData sets ResultData field to given value.


### GetErrorMessage

`func (o *SubmissionAction) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SubmissionAction) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SubmissionAction) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *SubmissionAction) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *SubmissionAction) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


