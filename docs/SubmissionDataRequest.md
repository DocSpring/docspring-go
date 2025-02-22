# SubmissionDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**Email** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**Order** | **NullableInt32** |  | 
**SortOrder** | **int32** |  | 
**Fields** | **[]string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**State** | **string** |  | 
**ViewedAt** | **NullableString** |  | 
**CompletedAt** | **NullableString** |  | 
**Data** | **map[string]interface{}** |  | 
**AuthType** | **string** |  | 
**AuthSecondFactorType** | **string** |  | 
**AuthProvider** | **NullableString** |  | 
**AuthSessionStartedAt** | **NullableString** |  | 
**AuthSessionIdHash** | **NullableString** |  | 
**AuthUserIdHash** | **NullableString** |  | 
**AuthUsernameHash** | **NullableString** |  | 
**AuthPhoneNumberHash** | **NullableString** |  | 
**IpAddress** | **NullableString** |  | 
**UserAgent** | **NullableString** |  | 

## Methods

### NewSubmissionDataRequest

`func NewSubmissionDataRequest(id NullableString, email NullableString, name NullableString, order NullableInt32, sortOrder int32, fields []string, metadata map[string]interface{}, state string, viewedAt NullableString, completedAt NullableString, data map[string]interface{}, authType string, authSecondFactorType string, authProvider NullableString, authSessionStartedAt NullableString, authSessionIdHash NullableString, authUserIdHash NullableString, authUsernameHash NullableString, authPhoneNumberHash NullableString, ipAddress NullableString, userAgent NullableString, ) *SubmissionDataRequest`

NewSubmissionDataRequest instantiates a new SubmissionDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionDataRequestWithDefaults

`func NewSubmissionDataRequestWithDefaults() *SubmissionDataRequest`

NewSubmissionDataRequestWithDefaults instantiates a new SubmissionDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubmissionDataRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionDataRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionDataRequest) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *SubmissionDataRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubmissionDataRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetEmail

`func (o *SubmissionDataRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubmissionDataRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubmissionDataRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *SubmissionDataRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SubmissionDataRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *SubmissionDataRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubmissionDataRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubmissionDataRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *SubmissionDataRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SubmissionDataRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrder

`func (o *SubmissionDataRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SubmissionDataRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SubmissionDataRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.


### SetOrderNil

`func (o *SubmissionDataRequest) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *SubmissionDataRequest) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetSortOrder

`func (o *SubmissionDataRequest) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *SubmissionDataRequest) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *SubmissionDataRequest) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetFields

`func (o *SubmissionDataRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SubmissionDataRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SubmissionDataRequest) SetFields(v []string)`

SetFields sets Fields field to given value.


### SetFieldsNil

`func (o *SubmissionDataRequest) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *SubmissionDataRequest) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetMetadata

`func (o *SubmissionDataRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubmissionDataRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubmissionDataRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *SubmissionDataRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SubmissionDataRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetState

`func (o *SubmissionDataRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubmissionDataRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubmissionDataRequest) SetState(v string)`

SetState sets State field to given value.


### GetViewedAt

`func (o *SubmissionDataRequest) GetViewedAt() string`

GetViewedAt returns the ViewedAt field if non-nil, zero value otherwise.

### GetViewedAtOk

`func (o *SubmissionDataRequest) GetViewedAtOk() (*string, bool)`

GetViewedAtOk returns a tuple with the ViewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedAt

`func (o *SubmissionDataRequest) SetViewedAt(v string)`

SetViewedAt sets ViewedAt field to given value.


### SetViewedAtNil

`func (o *SubmissionDataRequest) SetViewedAtNil(b bool)`

 SetViewedAtNil sets the value for ViewedAt to be an explicit nil

### UnsetViewedAt
`func (o *SubmissionDataRequest) UnsetViewedAt()`

UnsetViewedAt ensures that no value is present for ViewedAt, not even an explicit nil
### GetCompletedAt

`func (o *SubmissionDataRequest) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *SubmissionDataRequest) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *SubmissionDataRequest) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *SubmissionDataRequest) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *SubmissionDataRequest) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetData

`func (o *SubmissionDataRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubmissionDataRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubmissionDataRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *SubmissionDataRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SubmissionDataRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetAuthType

`func (o *SubmissionDataRequest) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *SubmissionDataRequest) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *SubmissionDataRequest) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetAuthSecondFactorType

`func (o *SubmissionDataRequest) GetAuthSecondFactorType() string`

GetAuthSecondFactorType returns the AuthSecondFactorType field if non-nil, zero value otherwise.

### GetAuthSecondFactorTypeOk

`func (o *SubmissionDataRequest) GetAuthSecondFactorTypeOk() (*string, bool)`

GetAuthSecondFactorTypeOk returns a tuple with the AuthSecondFactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecondFactorType

`func (o *SubmissionDataRequest) SetAuthSecondFactorType(v string)`

SetAuthSecondFactorType sets AuthSecondFactorType field to given value.


### GetAuthProvider

`func (o *SubmissionDataRequest) GetAuthProvider() string`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *SubmissionDataRequest) GetAuthProviderOk() (*string, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *SubmissionDataRequest) SetAuthProvider(v string)`

SetAuthProvider sets AuthProvider field to given value.


### SetAuthProviderNil

`func (o *SubmissionDataRequest) SetAuthProviderNil(b bool)`

 SetAuthProviderNil sets the value for AuthProvider to be an explicit nil

### UnsetAuthProvider
`func (o *SubmissionDataRequest) UnsetAuthProvider()`

UnsetAuthProvider ensures that no value is present for AuthProvider, not even an explicit nil
### GetAuthSessionStartedAt

`func (o *SubmissionDataRequest) GetAuthSessionStartedAt() string`

GetAuthSessionStartedAt returns the AuthSessionStartedAt field if non-nil, zero value otherwise.

### GetAuthSessionStartedAtOk

`func (o *SubmissionDataRequest) GetAuthSessionStartedAtOk() (*string, bool)`

GetAuthSessionStartedAtOk returns a tuple with the AuthSessionStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionStartedAt

`func (o *SubmissionDataRequest) SetAuthSessionStartedAt(v string)`

SetAuthSessionStartedAt sets AuthSessionStartedAt field to given value.


### SetAuthSessionStartedAtNil

`func (o *SubmissionDataRequest) SetAuthSessionStartedAtNil(b bool)`

 SetAuthSessionStartedAtNil sets the value for AuthSessionStartedAt to be an explicit nil

### UnsetAuthSessionStartedAt
`func (o *SubmissionDataRequest) UnsetAuthSessionStartedAt()`

UnsetAuthSessionStartedAt ensures that no value is present for AuthSessionStartedAt, not even an explicit nil
### GetAuthSessionIdHash

`func (o *SubmissionDataRequest) GetAuthSessionIdHash() string`

GetAuthSessionIdHash returns the AuthSessionIdHash field if non-nil, zero value otherwise.

### GetAuthSessionIdHashOk

`func (o *SubmissionDataRequest) GetAuthSessionIdHashOk() (*string, bool)`

GetAuthSessionIdHashOk returns a tuple with the AuthSessionIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionIdHash

`func (o *SubmissionDataRequest) SetAuthSessionIdHash(v string)`

SetAuthSessionIdHash sets AuthSessionIdHash field to given value.


### SetAuthSessionIdHashNil

`func (o *SubmissionDataRequest) SetAuthSessionIdHashNil(b bool)`

 SetAuthSessionIdHashNil sets the value for AuthSessionIdHash to be an explicit nil

### UnsetAuthSessionIdHash
`func (o *SubmissionDataRequest) UnsetAuthSessionIdHash()`

UnsetAuthSessionIdHash ensures that no value is present for AuthSessionIdHash, not even an explicit nil
### GetAuthUserIdHash

`func (o *SubmissionDataRequest) GetAuthUserIdHash() string`

GetAuthUserIdHash returns the AuthUserIdHash field if non-nil, zero value otherwise.

### GetAuthUserIdHashOk

`func (o *SubmissionDataRequest) GetAuthUserIdHashOk() (*string, bool)`

GetAuthUserIdHashOk returns a tuple with the AuthUserIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserIdHash

`func (o *SubmissionDataRequest) SetAuthUserIdHash(v string)`

SetAuthUserIdHash sets AuthUserIdHash field to given value.


### SetAuthUserIdHashNil

`func (o *SubmissionDataRequest) SetAuthUserIdHashNil(b bool)`

 SetAuthUserIdHashNil sets the value for AuthUserIdHash to be an explicit nil

### UnsetAuthUserIdHash
`func (o *SubmissionDataRequest) UnsetAuthUserIdHash()`

UnsetAuthUserIdHash ensures that no value is present for AuthUserIdHash, not even an explicit nil
### GetAuthUsernameHash

`func (o *SubmissionDataRequest) GetAuthUsernameHash() string`

GetAuthUsernameHash returns the AuthUsernameHash field if non-nil, zero value otherwise.

### GetAuthUsernameHashOk

`func (o *SubmissionDataRequest) GetAuthUsernameHashOk() (*string, bool)`

GetAuthUsernameHashOk returns a tuple with the AuthUsernameHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsernameHash

`func (o *SubmissionDataRequest) SetAuthUsernameHash(v string)`

SetAuthUsernameHash sets AuthUsernameHash field to given value.


### SetAuthUsernameHashNil

`func (o *SubmissionDataRequest) SetAuthUsernameHashNil(b bool)`

 SetAuthUsernameHashNil sets the value for AuthUsernameHash to be an explicit nil

### UnsetAuthUsernameHash
`func (o *SubmissionDataRequest) UnsetAuthUsernameHash()`

UnsetAuthUsernameHash ensures that no value is present for AuthUsernameHash, not even an explicit nil
### GetAuthPhoneNumberHash

`func (o *SubmissionDataRequest) GetAuthPhoneNumberHash() string`

GetAuthPhoneNumberHash returns the AuthPhoneNumberHash field if non-nil, zero value otherwise.

### GetAuthPhoneNumberHashOk

`func (o *SubmissionDataRequest) GetAuthPhoneNumberHashOk() (*string, bool)`

GetAuthPhoneNumberHashOk returns a tuple with the AuthPhoneNumberHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPhoneNumberHash

`func (o *SubmissionDataRequest) SetAuthPhoneNumberHash(v string)`

SetAuthPhoneNumberHash sets AuthPhoneNumberHash field to given value.


### SetAuthPhoneNumberHashNil

`func (o *SubmissionDataRequest) SetAuthPhoneNumberHashNil(b bool)`

 SetAuthPhoneNumberHashNil sets the value for AuthPhoneNumberHash to be an explicit nil

### UnsetAuthPhoneNumberHash
`func (o *SubmissionDataRequest) UnsetAuthPhoneNumberHash()`

UnsetAuthPhoneNumberHash ensures that no value is present for AuthPhoneNumberHash, not even an explicit nil
### GetIpAddress

`func (o *SubmissionDataRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SubmissionDataRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SubmissionDataRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### SetIpAddressNil

`func (o *SubmissionDataRequest) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *SubmissionDataRequest) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetUserAgent

`func (o *SubmissionDataRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *SubmissionDataRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *SubmissionDataRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### SetUserAgentNil

`func (o *SubmissionDataRequest) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *SubmissionDataRequest) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


