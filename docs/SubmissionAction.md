# SubmissionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**IntegrationId** | **NullableString** |  | 
**State** | **string** |  | 
**ActionType** | **string** |  | 
**ActionCategory** | **string** |  | 
**ResultData** | **map[string]interface{}** |  | 

## Methods

### NewSubmissionAction

`func NewSubmissionAction(id NullableString, integrationId NullableString, state string, actionType string, actionCategory string, resultData map[string]interface{}, ) *SubmissionAction`

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


### SetIdNil

`func (o *SubmissionAction) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubmissionAction) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
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

`func (o *SubmissionAction) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubmissionAction) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubmissionAction) SetState(v string)`

SetState sets State field to given value.


### GetActionType

`func (o *SubmissionAction) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *SubmissionAction) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *SubmissionAction) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetActionCategory

`func (o *SubmissionAction) GetActionCategory() string`

GetActionCategory returns the ActionCategory field if non-nil, zero value otherwise.

### GetActionCategoryOk

`func (o *SubmissionAction) GetActionCategoryOk() (*string, bool)`

GetActionCategoryOk returns a tuple with the ActionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCategory

`func (o *SubmissionAction) SetActionCategory(v string)`

SetActionCategory sets ActionCategory field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


