# UpdateSignatureRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Your own data about the signature request, such as an ID from your app. | [optional] 
**AuthType** | Pointer to [**SignatureRequestAuthType**](SignatureRequestAuthType.md) | How your app authenticated the person before it showed them the signature request. | [optional] 
**AuthSecondFactorType** | Pointer to [**SignatureRequestAuthSecondFactorType**](SignatureRequestAuthSecondFactorType.md) | The second factor your app authenticated them with, if any. | [optional] 
**AuthProvider** | Pointer to **NullableString** | The service that authenticated them, such as `google`. Required when `auth_type` is `oauth`. | [optional] 
**AuthSessionStartedAt** | Pointer to **NullableTime** | When your app authenticated them. Set for you with `email_link`. | [optional] 
**AuthSessionIdHash** | Pointer to **NullableString** | A hash of your app's session ID, for the audit trail. | [optional] 
**AuthUserIdHash** | Pointer to **NullableString** | Their user ID in your app, or a SHA-256 hash of a salt and the user ID. | [optional] 
**AuthUsernameHash** | Pointer to **NullableString** | Their username in your app, or a SHA-256 hash of a salt and the username. | [optional] 
**AuthPhoneNumberHash** | Pointer to **NullableString** | The phone number they authenticated with, or a SHA-256 hash of a salt and the phone number. | [optional] 

## Methods

### NewUpdateSignatureRequestRequest

`func NewUpdateSignatureRequestRequest() *UpdateSignatureRequestRequest`

NewUpdateSignatureRequestRequest instantiates a new UpdateSignatureRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSignatureRequestRequestWithDefaults

`func NewUpdateSignatureRequestRequestWithDefaults() *UpdateSignatureRequestRequest`

NewUpdateSignatureRequestRequestWithDefaults instantiates a new UpdateSignatureRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *UpdateSignatureRequestRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateSignatureRequestRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateSignatureRequestRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateSignatureRequestRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAuthType

`func (o *UpdateSignatureRequestRequest) GetAuthType() SignatureRequestAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UpdateSignatureRequestRequest) GetAuthTypeOk() (*SignatureRequestAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UpdateSignatureRequestRequest) SetAuthType(v SignatureRequestAuthType)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *UpdateSignatureRequestRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthSecondFactorType

`func (o *UpdateSignatureRequestRequest) GetAuthSecondFactorType() SignatureRequestAuthSecondFactorType`

GetAuthSecondFactorType returns the AuthSecondFactorType field if non-nil, zero value otherwise.

### GetAuthSecondFactorTypeOk

`func (o *UpdateSignatureRequestRequest) GetAuthSecondFactorTypeOk() (*SignatureRequestAuthSecondFactorType, bool)`

GetAuthSecondFactorTypeOk returns a tuple with the AuthSecondFactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecondFactorType

`func (o *UpdateSignatureRequestRequest) SetAuthSecondFactorType(v SignatureRequestAuthSecondFactorType)`

SetAuthSecondFactorType sets AuthSecondFactorType field to given value.

### HasAuthSecondFactorType

`func (o *UpdateSignatureRequestRequest) HasAuthSecondFactorType() bool`

HasAuthSecondFactorType returns a boolean if a field has been set.

### GetAuthProvider

`func (o *UpdateSignatureRequestRequest) GetAuthProvider() string`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *UpdateSignatureRequestRequest) GetAuthProviderOk() (*string, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *UpdateSignatureRequestRequest) SetAuthProvider(v string)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *UpdateSignatureRequestRequest) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### SetAuthProviderNil

`func (o *UpdateSignatureRequestRequest) SetAuthProviderNil(b bool)`

 SetAuthProviderNil sets the value for AuthProvider to be an explicit nil

### UnsetAuthProvider
`func (o *UpdateSignatureRequestRequest) UnsetAuthProvider()`

UnsetAuthProvider ensures that no value is present for AuthProvider, not even an explicit nil
### GetAuthSessionStartedAt

`func (o *UpdateSignatureRequestRequest) GetAuthSessionStartedAt() time.Time`

GetAuthSessionStartedAt returns the AuthSessionStartedAt field if non-nil, zero value otherwise.

### GetAuthSessionStartedAtOk

`func (o *UpdateSignatureRequestRequest) GetAuthSessionStartedAtOk() (*time.Time, bool)`

GetAuthSessionStartedAtOk returns a tuple with the AuthSessionStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionStartedAt

`func (o *UpdateSignatureRequestRequest) SetAuthSessionStartedAt(v time.Time)`

SetAuthSessionStartedAt sets AuthSessionStartedAt field to given value.

### HasAuthSessionStartedAt

`func (o *UpdateSignatureRequestRequest) HasAuthSessionStartedAt() bool`

HasAuthSessionStartedAt returns a boolean if a field has been set.

### SetAuthSessionStartedAtNil

`func (o *UpdateSignatureRequestRequest) SetAuthSessionStartedAtNil(b bool)`

 SetAuthSessionStartedAtNil sets the value for AuthSessionStartedAt to be an explicit nil

### UnsetAuthSessionStartedAt
`func (o *UpdateSignatureRequestRequest) UnsetAuthSessionStartedAt()`

UnsetAuthSessionStartedAt ensures that no value is present for AuthSessionStartedAt, not even an explicit nil
### GetAuthSessionIdHash

`func (o *UpdateSignatureRequestRequest) GetAuthSessionIdHash() string`

GetAuthSessionIdHash returns the AuthSessionIdHash field if non-nil, zero value otherwise.

### GetAuthSessionIdHashOk

`func (o *UpdateSignatureRequestRequest) GetAuthSessionIdHashOk() (*string, bool)`

GetAuthSessionIdHashOk returns a tuple with the AuthSessionIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionIdHash

`func (o *UpdateSignatureRequestRequest) SetAuthSessionIdHash(v string)`

SetAuthSessionIdHash sets AuthSessionIdHash field to given value.

### HasAuthSessionIdHash

`func (o *UpdateSignatureRequestRequest) HasAuthSessionIdHash() bool`

HasAuthSessionIdHash returns a boolean if a field has been set.

### SetAuthSessionIdHashNil

`func (o *UpdateSignatureRequestRequest) SetAuthSessionIdHashNil(b bool)`

 SetAuthSessionIdHashNil sets the value for AuthSessionIdHash to be an explicit nil

### UnsetAuthSessionIdHash
`func (o *UpdateSignatureRequestRequest) UnsetAuthSessionIdHash()`

UnsetAuthSessionIdHash ensures that no value is present for AuthSessionIdHash, not even an explicit nil
### GetAuthUserIdHash

`func (o *UpdateSignatureRequestRequest) GetAuthUserIdHash() string`

GetAuthUserIdHash returns the AuthUserIdHash field if non-nil, zero value otherwise.

### GetAuthUserIdHashOk

`func (o *UpdateSignatureRequestRequest) GetAuthUserIdHashOk() (*string, bool)`

GetAuthUserIdHashOk returns a tuple with the AuthUserIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserIdHash

`func (o *UpdateSignatureRequestRequest) SetAuthUserIdHash(v string)`

SetAuthUserIdHash sets AuthUserIdHash field to given value.

### HasAuthUserIdHash

`func (o *UpdateSignatureRequestRequest) HasAuthUserIdHash() bool`

HasAuthUserIdHash returns a boolean if a field has been set.

### SetAuthUserIdHashNil

`func (o *UpdateSignatureRequestRequest) SetAuthUserIdHashNil(b bool)`

 SetAuthUserIdHashNil sets the value for AuthUserIdHash to be an explicit nil

### UnsetAuthUserIdHash
`func (o *UpdateSignatureRequestRequest) UnsetAuthUserIdHash()`

UnsetAuthUserIdHash ensures that no value is present for AuthUserIdHash, not even an explicit nil
### GetAuthUsernameHash

`func (o *UpdateSignatureRequestRequest) GetAuthUsernameHash() string`

GetAuthUsernameHash returns the AuthUsernameHash field if non-nil, zero value otherwise.

### GetAuthUsernameHashOk

`func (o *UpdateSignatureRequestRequest) GetAuthUsernameHashOk() (*string, bool)`

GetAuthUsernameHashOk returns a tuple with the AuthUsernameHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsernameHash

`func (o *UpdateSignatureRequestRequest) SetAuthUsernameHash(v string)`

SetAuthUsernameHash sets AuthUsernameHash field to given value.

### HasAuthUsernameHash

`func (o *UpdateSignatureRequestRequest) HasAuthUsernameHash() bool`

HasAuthUsernameHash returns a boolean if a field has been set.

### SetAuthUsernameHashNil

`func (o *UpdateSignatureRequestRequest) SetAuthUsernameHashNil(b bool)`

 SetAuthUsernameHashNil sets the value for AuthUsernameHash to be an explicit nil

### UnsetAuthUsernameHash
`func (o *UpdateSignatureRequestRequest) UnsetAuthUsernameHash()`

UnsetAuthUsernameHash ensures that no value is present for AuthUsernameHash, not even an explicit nil
### GetAuthPhoneNumberHash

`func (o *UpdateSignatureRequestRequest) GetAuthPhoneNumberHash() string`

GetAuthPhoneNumberHash returns the AuthPhoneNumberHash field if non-nil, zero value otherwise.

### GetAuthPhoneNumberHashOk

`func (o *UpdateSignatureRequestRequest) GetAuthPhoneNumberHashOk() (*string, bool)`

GetAuthPhoneNumberHashOk returns a tuple with the AuthPhoneNumberHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPhoneNumberHash

`func (o *UpdateSignatureRequestRequest) SetAuthPhoneNumberHash(v string)`

SetAuthPhoneNumberHash sets AuthPhoneNumberHash field to given value.

### HasAuthPhoneNumberHash

`func (o *UpdateSignatureRequestRequest) HasAuthPhoneNumberHash() bool`

HasAuthPhoneNumberHash returns a boolean if a field has been set.

### SetAuthPhoneNumberHashNil

`func (o *UpdateSignatureRequestRequest) SetAuthPhoneNumberHashNil(b bool)`

 SetAuthPhoneNumberHashNil sets the value for AuthPhoneNumberHash to be an explicit nil

### UnsetAuthPhoneNumberHash
`func (o *UpdateSignatureRequestRequest) UnsetAuthPhoneNumberHash()`

UnsetAuthPhoneNumberHash ensures that no value is present for AuthPhoneNumberHash, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


