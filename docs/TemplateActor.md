# TemplateActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TemplateActorType**](TemplateActorType.md) | `user`: a user in your account, also when they make the change with an app they connected (an OAuth access token). `api`: an API token. `external_user`: someone using the embedded template editor, as identified by your app. | 
**UserId** | **NullableString** | The user's ID, for `user`. | 
**ExternalId** | **NullableString** | The ID your app gave the embedded editor's user, for `external_user`. | 
**Name** | **NullableString** | The user's name, or the name your app gave the embedded editor's user. | 
**Email** | **NullableString** | The email address your app gave the embedded editor's user, for `external_user`. | 

## Methods

### NewTemplateActor

`func NewTemplateActor(type_ TemplateActorType, userId NullableString, externalId NullableString, name NullableString, email NullableString, ) *TemplateActor`

NewTemplateActor instantiates a new TemplateActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateActorWithDefaults

`func NewTemplateActorWithDefaults() *TemplateActor`

NewTemplateActorWithDefaults instantiates a new TemplateActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateActor) GetType() TemplateActorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateActor) GetTypeOk() (*TemplateActorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateActor) SetType(v TemplateActorType)`

SetType sets Type field to given value.


### GetUserId

`func (o *TemplateActor) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TemplateActor) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TemplateActor) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *TemplateActor) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *TemplateActor) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetExternalId

`func (o *TemplateActor) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TemplateActor) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TemplateActor) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *TemplateActor) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *TemplateActor) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetName

`func (o *TemplateActor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateActor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateActor) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *TemplateActor) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplateActor) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *TemplateActor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TemplateActor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TemplateActor) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *TemplateActor) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *TemplateActor) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


