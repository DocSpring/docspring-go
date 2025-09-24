# UpdateSubmissionDataRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPhoneNumberHash** | Pointer to **NullableString** |  | [optional] 
**AuthProvider** | Pointer to **NullableString** |  | [optional] 
**AuthSecondFactorType** | Pointer to **NullableString** |  | [optional] 
**AuthSessionIdHash** | Pointer to **NullableString** |  | [optional] 
**AuthSessionStartedAt** | Pointer to **NullableString** |  | [optional] 
**AuthType** | Pointer to **string** |  | [optional] 
**AuthUserIdHash** | Pointer to **NullableString** |  | [optional] 
**AuthUsernameHash** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateSubmissionDataRequestData

`func NewUpdateSubmissionDataRequestData() *UpdateSubmissionDataRequestData`

NewUpdateSubmissionDataRequestData instantiates a new UpdateSubmissionDataRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubmissionDataRequestDataWithDefaults

`func NewUpdateSubmissionDataRequestDataWithDefaults() *UpdateSubmissionDataRequestData`

NewUpdateSubmissionDataRequestDataWithDefaults instantiates a new UpdateSubmissionDataRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthPhoneNumberHash

`func (o *UpdateSubmissionDataRequestData) GetAuthPhoneNumberHash() string`

GetAuthPhoneNumberHash returns the AuthPhoneNumberHash field if non-nil, zero value otherwise.

### GetAuthPhoneNumberHashOk

`func (o *UpdateSubmissionDataRequestData) GetAuthPhoneNumberHashOk() (*string, bool)`

GetAuthPhoneNumberHashOk returns a tuple with the AuthPhoneNumberHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPhoneNumberHash

`func (o *UpdateSubmissionDataRequestData) SetAuthPhoneNumberHash(v string)`

SetAuthPhoneNumberHash sets AuthPhoneNumberHash field to given value.

### HasAuthPhoneNumberHash

`func (o *UpdateSubmissionDataRequestData) HasAuthPhoneNumberHash() bool`

HasAuthPhoneNumberHash returns a boolean if a field has been set.

### SetAuthPhoneNumberHashNil

`func (o *UpdateSubmissionDataRequestData) SetAuthPhoneNumberHashNil(b bool)`

 SetAuthPhoneNumberHashNil sets the value for AuthPhoneNumberHash to be an explicit nil

### UnsetAuthPhoneNumberHash
`func (o *UpdateSubmissionDataRequestData) UnsetAuthPhoneNumberHash()`

UnsetAuthPhoneNumberHash ensures that no value is present for AuthPhoneNumberHash, not even an explicit nil
### GetAuthProvider

`func (o *UpdateSubmissionDataRequestData) GetAuthProvider() string`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *UpdateSubmissionDataRequestData) GetAuthProviderOk() (*string, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *UpdateSubmissionDataRequestData) SetAuthProvider(v string)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *UpdateSubmissionDataRequestData) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### SetAuthProviderNil

`func (o *UpdateSubmissionDataRequestData) SetAuthProviderNil(b bool)`

 SetAuthProviderNil sets the value for AuthProvider to be an explicit nil

### UnsetAuthProvider
`func (o *UpdateSubmissionDataRequestData) UnsetAuthProvider()`

UnsetAuthProvider ensures that no value is present for AuthProvider, not even an explicit nil
### GetAuthSecondFactorType

`func (o *UpdateSubmissionDataRequestData) GetAuthSecondFactorType() string`

GetAuthSecondFactorType returns the AuthSecondFactorType field if non-nil, zero value otherwise.

### GetAuthSecondFactorTypeOk

`func (o *UpdateSubmissionDataRequestData) GetAuthSecondFactorTypeOk() (*string, bool)`

GetAuthSecondFactorTypeOk returns a tuple with the AuthSecondFactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecondFactorType

`func (o *UpdateSubmissionDataRequestData) SetAuthSecondFactorType(v string)`

SetAuthSecondFactorType sets AuthSecondFactorType field to given value.

### HasAuthSecondFactorType

`func (o *UpdateSubmissionDataRequestData) HasAuthSecondFactorType() bool`

HasAuthSecondFactorType returns a boolean if a field has been set.

### SetAuthSecondFactorTypeNil

`func (o *UpdateSubmissionDataRequestData) SetAuthSecondFactorTypeNil(b bool)`

 SetAuthSecondFactorTypeNil sets the value for AuthSecondFactorType to be an explicit nil

### UnsetAuthSecondFactorType
`func (o *UpdateSubmissionDataRequestData) UnsetAuthSecondFactorType()`

UnsetAuthSecondFactorType ensures that no value is present for AuthSecondFactorType, not even an explicit nil
### GetAuthSessionIdHash

`func (o *UpdateSubmissionDataRequestData) GetAuthSessionIdHash() string`

GetAuthSessionIdHash returns the AuthSessionIdHash field if non-nil, zero value otherwise.

### GetAuthSessionIdHashOk

`func (o *UpdateSubmissionDataRequestData) GetAuthSessionIdHashOk() (*string, bool)`

GetAuthSessionIdHashOk returns a tuple with the AuthSessionIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionIdHash

`func (o *UpdateSubmissionDataRequestData) SetAuthSessionIdHash(v string)`

SetAuthSessionIdHash sets AuthSessionIdHash field to given value.

### HasAuthSessionIdHash

`func (o *UpdateSubmissionDataRequestData) HasAuthSessionIdHash() bool`

HasAuthSessionIdHash returns a boolean if a field has been set.

### SetAuthSessionIdHashNil

`func (o *UpdateSubmissionDataRequestData) SetAuthSessionIdHashNil(b bool)`

 SetAuthSessionIdHashNil sets the value for AuthSessionIdHash to be an explicit nil

### UnsetAuthSessionIdHash
`func (o *UpdateSubmissionDataRequestData) UnsetAuthSessionIdHash()`

UnsetAuthSessionIdHash ensures that no value is present for AuthSessionIdHash, not even an explicit nil
### GetAuthSessionStartedAt

`func (o *UpdateSubmissionDataRequestData) GetAuthSessionStartedAt() string`

GetAuthSessionStartedAt returns the AuthSessionStartedAt field if non-nil, zero value otherwise.

### GetAuthSessionStartedAtOk

`func (o *UpdateSubmissionDataRequestData) GetAuthSessionStartedAtOk() (*string, bool)`

GetAuthSessionStartedAtOk returns a tuple with the AuthSessionStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionStartedAt

`func (o *UpdateSubmissionDataRequestData) SetAuthSessionStartedAt(v string)`

SetAuthSessionStartedAt sets AuthSessionStartedAt field to given value.

### HasAuthSessionStartedAt

`func (o *UpdateSubmissionDataRequestData) HasAuthSessionStartedAt() bool`

HasAuthSessionStartedAt returns a boolean if a field has been set.

### SetAuthSessionStartedAtNil

`func (o *UpdateSubmissionDataRequestData) SetAuthSessionStartedAtNil(b bool)`

 SetAuthSessionStartedAtNil sets the value for AuthSessionStartedAt to be an explicit nil

### UnsetAuthSessionStartedAt
`func (o *UpdateSubmissionDataRequestData) UnsetAuthSessionStartedAt()`

UnsetAuthSessionStartedAt ensures that no value is present for AuthSessionStartedAt, not even an explicit nil
### GetAuthType

`func (o *UpdateSubmissionDataRequestData) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UpdateSubmissionDataRequestData) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UpdateSubmissionDataRequestData) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *UpdateSubmissionDataRequestData) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthUserIdHash

`func (o *UpdateSubmissionDataRequestData) GetAuthUserIdHash() string`

GetAuthUserIdHash returns the AuthUserIdHash field if non-nil, zero value otherwise.

### GetAuthUserIdHashOk

`func (o *UpdateSubmissionDataRequestData) GetAuthUserIdHashOk() (*string, bool)`

GetAuthUserIdHashOk returns a tuple with the AuthUserIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserIdHash

`func (o *UpdateSubmissionDataRequestData) SetAuthUserIdHash(v string)`

SetAuthUserIdHash sets AuthUserIdHash field to given value.

### HasAuthUserIdHash

`func (o *UpdateSubmissionDataRequestData) HasAuthUserIdHash() bool`

HasAuthUserIdHash returns a boolean if a field has been set.

### SetAuthUserIdHashNil

`func (o *UpdateSubmissionDataRequestData) SetAuthUserIdHashNil(b bool)`

 SetAuthUserIdHashNil sets the value for AuthUserIdHash to be an explicit nil

### UnsetAuthUserIdHash
`func (o *UpdateSubmissionDataRequestData) UnsetAuthUserIdHash()`

UnsetAuthUserIdHash ensures that no value is present for AuthUserIdHash, not even an explicit nil
### GetAuthUsernameHash

`func (o *UpdateSubmissionDataRequestData) GetAuthUsernameHash() string`

GetAuthUsernameHash returns the AuthUsernameHash field if non-nil, zero value otherwise.

### GetAuthUsernameHashOk

`func (o *UpdateSubmissionDataRequestData) GetAuthUsernameHashOk() (*string, bool)`

GetAuthUsernameHashOk returns a tuple with the AuthUsernameHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsernameHash

`func (o *UpdateSubmissionDataRequestData) SetAuthUsernameHash(v string)`

SetAuthUsernameHash sets AuthUsernameHash field to given value.

### HasAuthUsernameHash

`func (o *UpdateSubmissionDataRequestData) HasAuthUsernameHash() bool`

HasAuthUsernameHash returns a boolean if a field has been set.

### SetAuthUsernameHashNil

`func (o *UpdateSubmissionDataRequestData) SetAuthUsernameHashNil(b bool)`

 SetAuthUsernameHashNil sets the value for AuthUsernameHash to be an explicit nil

### UnsetAuthUsernameHash
`func (o *UpdateSubmissionDataRequestData) UnsetAuthUsernameHash()`

UnsetAuthUsernameHash ensures that no value is present for AuthUsernameHash, not even an explicit nil
### GetEmail

`func (o *UpdateSubmissionDataRequestData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateSubmissionDataRequestData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateSubmissionDataRequestData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateSubmissionDataRequestData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdateSubmissionDataRequestData) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdateSubmissionDataRequestData) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFields

`func (o *UpdateSubmissionDataRequestData) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *UpdateSubmissionDataRequestData) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *UpdateSubmissionDataRequestData) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *UpdateSubmissionDataRequestData) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateSubmissionDataRequestData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateSubmissionDataRequestData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateSubmissionDataRequestData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateSubmissionDataRequestData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UpdateSubmissionDataRequestData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSubmissionDataRequestData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSubmissionDataRequestData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSubmissionDataRequestData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateSubmissionDataRequestData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateSubmissionDataRequestData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrder

`func (o *UpdateSubmissionDataRequestData) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UpdateSubmissionDataRequestData) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UpdateSubmissionDataRequestData) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *UpdateSubmissionDataRequestData) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


