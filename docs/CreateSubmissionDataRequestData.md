# CreateSubmissionDataRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPhoneNumberHash** | Pointer to **NullableString** |  | [optional] 
**AuthProvider** | Pointer to **NullableString** |  | [optional] 
**AuthSecondFactorType** | Pointer to **string** |  | [optional] 
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
**Skipped** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreateSubmissionDataRequestData

`func NewCreateSubmissionDataRequestData() *CreateSubmissionDataRequestData`

NewCreateSubmissionDataRequestData instantiates a new CreateSubmissionDataRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubmissionDataRequestDataWithDefaults

`func NewCreateSubmissionDataRequestDataWithDefaults() *CreateSubmissionDataRequestData`

NewCreateSubmissionDataRequestDataWithDefaults instantiates a new CreateSubmissionDataRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthPhoneNumberHash

`func (o *CreateSubmissionDataRequestData) GetAuthPhoneNumberHash() string`

GetAuthPhoneNumberHash returns the AuthPhoneNumberHash field if non-nil, zero value otherwise.

### GetAuthPhoneNumberHashOk

`func (o *CreateSubmissionDataRequestData) GetAuthPhoneNumberHashOk() (*string, bool)`

GetAuthPhoneNumberHashOk returns a tuple with the AuthPhoneNumberHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPhoneNumberHash

`func (o *CreateSubmissionDataRequestData) SetAuthPhoneNumberHash(v string)`

SetAuthPhoneNumberHash sets AuthPhoneNumberHash field to given value.

### HasAuthPhoneNumberHash

`func (o *CreateSubmissionDataRequestData) HasAuthPhoneNumberHash() bool`

HasAuthPhoneNumberHash returns a boolean if a field has been set.

### SetAuthPhoneNumberHashNil

`func (o *CreateSubmissionDataRequestData) SetAuthPhoneNumberHashNil(b bool)`

 SetAuthPhoneNumberHashNil sets the value for AuthPhoneNumberHash to be an explicit nil

### UnsetAuthPhoneNumberHash
`func (o *CreateSubmissionDataRequestData) UnsetAuthPhoneNumberHash()`

UnsetAuthPhoneNumberHash ensures that no value is present for AuthPhoneNumberHash, not even an explicit nil
### GetAuthProvider

`func (o *CreateSubmissionDataRequestData) GetAuthProvider() string`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *CreateSubmissionDataRequestData) GetAuthProviderOk() (*string, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *CreateSubmissionDataRequestData) SetAuthProvider(v string)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *CreateSubmissionDataRequestData) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### SetAuthProviderNil

`func (o *CreateSubmissionDataRequestData) SetAuthProviderNil(b bool)`

 SetAuthProviderNil sets the value for AuthProvider to be an explicit nil

### UnsetAuthProvider
`func (o *CreateSubmissionDataRequestData) UnsetAuthProvider()`

UnsetAuthProvider ensures that no value is present for AuthProvider, not even an explicit nil
### GetAuthSecondFactorType

`func (o *CreateSubmissionDataRequestData) GetAuthSecondFactorType() string`

GetAuthSecondFactorType returns the AuthSecondFactorType field if non-nil, zero value otherwise.

### GetAuthSecondFactorTypeOk

`func (o *CreateSubmissionDataRequestData) GetAuthSecondFactorTypeOk() (*string, bool)`

GetAuthSecondFactorTypeOk returns a tuple with the AuthSecondFactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecondFactorType

`func (o *CreateSubmissionDataRequestData) SetAuthSecondFactorType(v string)`

SetAuthSecondFactorType sets AuthSecondFactorType field to given value.

### HasAuthSecondFactorType

`func (o *CreateSubmissionDataRequestData) HasAuthSecondFactorType() bool`

HasAuthSecondFactorType returns a boolean if a field has been set.

### GetAuthSessionIdHash

`func (o *CreateSubmissionDataRequestData) GetAuthSessionIdHash() string`

GetAuthSessionIdHash returns the AuthSessionIdHash field if non-nil, zero value otherwise.

### GetAuthSessionIdHashOk

`func (o *CreateSubmissionDataRequestData) GetAuthSessionIdHashOk() (*string, bool)`

GetAuthSessionIdHashOk returns a tuple with the AuthSessionIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionIdHash

`func (o *CreateSubmissionDataRequestData) SetAuthSessionIdHash(v string)`

SetAuthSessionIdHash sets AuthSessionIdHash field to given value.

### HasAuthSessionIdHash

`func (o *CreateSubmissionDataRequestData) HasAuthSessionIdHash() bool`

HasAuthSessionIdHash returns a boolean if a field has been set.

### SetAuthSessionIdHashNil

`func (o *CreateSubmissionDataRequestData) SetAuthSessionIdHashNil(b bool)`

 SetAuthSessionIdHashNil sets the value for AuthSessionIdHash to be an explicit nil

### UnsetAuthSessionIdHash
`func (o *CreateSubmissionDataRequestData) UnsetAuthSessionIdHash()`

UnsetAuthSessionIdHash ensures that no value is present for AuthSessionIdHash, not even an explicit nil
### GetAuthSessionStartedAt

`func (o *CreateSubmissionDataRequestData) GetAuthSessionStartedAt() string`

GetAuthSessionStartedAt returns the AuthSessionStartedAt field if non-nil, zero value otherwise.

### GetAuthSessionStartedAtOk

`func (o *CreateSubmissionDataRequestData) GetAuthSessionStartedAtOk() (*string, bool)`

GetAuthSessionStartedAtOk returns a tuple with the AuthSessionStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionStartedAt

`func (o *CreateSubmissionDataRequestData) SetAuthSessionStartedAt(v string)`

SetAuthSessionStartedAt sets AuthSessionStartedAt field to given value.

### HasAuthSessionStartedAt

`func (o *CreateSubmissionDataRequestData) HasAuthSessionStartedAt() bool`

HasAuthSessionStartedAt returns a boolean if a field has been set.

### SetAuthSessionStartedAtNil

`func (o *CreateSubmissionDataRequestData) SetAuthSessionStartedAtNil(b bool)`

 SetAuthSessionStartedAtNil sets the value for AuthSessionStartedAt to be an explicit nil

### UnsetAuthSessionStartedAt
`func (o *CreateSubmissionDataRequestData) UnsetAuthSessionStartedAt()`

UnsetAuthSessionStartedAt ensures that no value is present for AuthSessionStartedAt, not even an explicit nil
### GetAuthType

`func (o *CreateSubmissionDataRequestData) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *CreateSubmissionDataRequestData) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *CreateSubmissionDataRequestData) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *CreateSubmissionDataRequestData) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthUserIdHash

`func (o *CreateSubmissionDataRequestData) GetAuthUserIdHash() string`

GetAuthUserIdHash returns the AuthUserIdHash field if non-nil, zero value otherwise.

### GetAuthUserIdHashOk

`func (o *CreateSubmissionDataRequestData) GetAuthUserIdHashOk() (*string, bool)`

GetAuthUserIdHashOk returns a tuple with the AuthUserIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserIdHash

`func (o *CreateSubmissionDataRequestData) SetAuthUserIdHash(v string)`

SetAuthUserIdHash sets AuthUserIdHash field to given value.

### HasAuthUserIdHash

`func (o *CreateSubmissionDataRequestData) HasAuthUserIdHash() bool`

HasAuthUserIdHash returns a boolean if a field has been set.

### SetAuthUserIdHashNil

`func (o *CreateSubmissionDataRequestData) SetAuthUserIdHashNil(b bool)`

 SetAuthUserIdHashNil sets the value for AuthUserIdHash to be an explicit nil

### UnsetAuthUserIdHash
`func (o *CreateSubmissionDataRequestData) UnsetAuthUserIdHash()`

UnsetAuthUserIdHash ensures that no value is present for AuthUserIdHash, not even an explicit nil
### GetAuthUsernameHash

`func (o *CreateSubmissionDataRequestData) GetAuthUsernameHash() string`

GetAuthUsernameHash returns the AuthUsernameHash field if non-nil, zero value otherwise.

### GetAuthUsernameHashOk

`func (o *CreateSubmissionDataRequestData) GetAuthUsernameHashOk() (*string, bool)`

GetAuthUsernameHashOk returns a tuple with the AuthUsernameHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsernameHash

`func (o *CreateSubmissionDataRequestData) SetAuthUsernameHash(v string)`

SetAuthUsernameHash sets AuthUsernameHash field to given value.

### HasAuthUsernameHash

`func (o *CreateSubmissionDataRequestData) HasAuthUsernameHash() bool`

HasAuthUsernameHash returns a boolean if a field has been set.

### SetAuthUsernameHashNil

`func (o *CreateSubmissionDataRequestData) SetAuthUsernameHashNil(b bool)`

 SetAuthUsernameHashNil sets the value for AuthUsernameHash to be an explicit nil

### UnsetAuthUsernameHash
`func (o *CreateSubmissionDataRequestData) UnsetAuthUsernameHash()`

UnsetAuthUsernameHash ensures that no value is present for AuthUsernameHash, not even an explicit nil
### GetEmail

`func (o *CreateSubmissionDataRequestData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateSubmissionDataRequestData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateSubmissionDataRequestData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateSubmissionDataRequestData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CreateSubmissionDataRequestData) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateSubmissionDataRequestData) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFields

`func (o *CreateSubmissionDataRequestData) GetFields() []*string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CreateSubmissionDataRequestData) GetFieldsOk() (*[]*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CreateSubmissionDataRequestData) SetFields(v []*string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CreateSubmissionDataRequestData) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateSubmissionDataRequestData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSubmissionDataRequestData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSubmissionDataRequestData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSubmissionDataRequestData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateSubmissionDataRequestData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSubmissionDataRequestData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSubmissionDataRequestData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSubmissionDataRequestData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateSubmissionDataRequestData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateSubmissionDataRequestData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrder

`func (o *CreateSubmissionDataRequestData) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CreateSubmissionDataRequestData) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CreateSubmissionDataRequestData) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CreateSubmissionDataRequestData) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSkipped

`func (o *CreateSubmissionDataRequestData) GetSkipped() bool`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *CreateSubmissionDataRequestData) GetSkippedOk() (*bool, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *CreateSubmissionDataRequestData) SetSkipped(v bool)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *CreateSubmissionDataRequestData) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### SetSkippedNil

`func (o *CreateSubmissionDataRequestData) SetSkippedNil(b bool)`

 SetSkippedNil sets the value for Skipped to be an explicit nil

### UnsetSkipped
`func (o *CreateSubmissionDataRequestData) UnsetSkipped()`

UnsetSkipped ensures that no value is present for Skipped, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


