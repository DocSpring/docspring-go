# SignatureRequestInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the person the request is for. | [optional] 
**Email** | Pointer to **NullableString** | Their email address. | [optional] 
**Fields** | **[]string** | The fields this person fills in. Each field can only be in one signature request, and a field that `data` fills in needs `prefilled_values_editable`. | 
**Metadata** | Pointer to **map[string]interface{}** | Your own data about the signature request, such as an ID from your app. | [optional] 
**Role** | Pointer to **string** | Which of the template's signers the person is: `Signer 1`, `Signer 2` and so on, up to `Signer 100`, until templates have roles of their own. Either every signature request has a `role` or none does, and each role can only have one. | [optional] 
**AuthType** | [**SignatureRequestAuthType**](SignatureRequestAuthType.md) | How your app authenticated the person before it showed them the signature request. | 
**AuthSecondFactorType** | Pointer to [**SignatureRequestAuthSecondFactorType**](SignatureRequestAuthSecondFactorType.md) | The second factor your app authenticated them with, if any. | [optional] 
**AuthProvider** | Pointer to **NullableString** | The service that authenticated them, such as `google`. Required when `auth_type` is `oauth`. | [optional] 
**AuthSessionStartedAt** | Pointer to **NullableTime** | When your app authenticated them. Set for you with `email_link`. | [optional] 
**AuthSessionIdHash** | Pointer to **NullableString** | A hash of your app's session ID, for the audit trail. | [optional] 
**AuthUserIdHash** | Pointer to **NullableString** | Their user ID in your app, or a SHA-256 hash of a salt and the user ID. | [optional] 
**AuthUsernameHash** | Pointer to **NullableString** | Their username in your app, or a SHA-256 hash of a salt and the username. | [optional] 
**AuthPhoneNumberHash** | Pointer to **NullableString** | The phone number they authenticated with, or a SHA-256 hash of a salt and the phone number. | [optional] 
**IgnoreUnknownFields** | Pointer to **bool** | Leave out names in `fields` that the template doesn't have, for one field list that you send with several templates. Without it, an unknown name makes the submission `invalid_data`. | [optional] 
**PrefilledValuesEditable** | Pointer to **bool** | Let the person change the values that `data` already has for fields in `fields`. Without it, those fields are a 422. | [optional] 

## Methods

### NewSignatureRequestInput

`func NewSignatureRequestInput(fields []string, authType SignatureRequestAuthType, ) *SignatureRequestInput`

NewSignatureRequestInput instantiates a new SignatureRequestInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureRequestInputWithDefaults

`func NewSignatureRequestInputWithDefaults() *SignatureRequestInput`

NewSignatureRequestInputWithDefaults instantiates a new SignatureRequestInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SignatureRequestInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignatureRequestInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignatureRequestInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SignatureRequestInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SignatureRequestInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SignatureRequestInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *SignatureRequestInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignatureRequestInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignatureRequestInput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SignatureRequestInput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SignatureRequestInput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SignatureRequestInput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFields

`func (o *SignatureRequestInput) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SignatureRequestInput) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SignatureRequestInput) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetMetadata

`func (o *SignatureRequestInput) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SignatureRequestInput) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SignatureRequestInput) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SignatureRequestInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRole

`func (o *SignatureRequestInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SignatureRequestInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SignatureRequestInput) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SignatureRequestInput) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAuthType

`func (o *SignatureRequestInput) GetAuthType() SignatureRequestAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *SignatureRequestInput) GetAuthTypeOk() (*SignatureRequestAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *SignatureRequestInput) SetAuthType(v SignatureRequestAuthType)`

SetAuthType sets AuthType field to given value.


### GetAuthSecondFactorType

`func (o *SignatureRequestInput) GetAuthSecondFactorType() SignatureRequestAuthSecondFactorType`

GetAuthSecondFactorType returns the AuthSecondFactorType field if non-nil, zero value otherwise.

### GetAuthSecondFactorTypeOk

`func (o *SignatureRequestInput) GetAuthSecondFactorTypeOk() (*SignatureRequestAuthSecondFactorType, bool)`

GetAuthSecondFactorTypeOk returns a tuple with the AuthSecondFactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecondFactorType

`func (o *SignatureRequestInput) SetAuthSecondFactorType(v SignatureRequestAuthSecondFactorType)`

SetAuthSecondFactorType sets AuthSecondFactorType field to given value.

### HasAuthSecondFactorType

`func (o *SignatureRequestInput) HasAuthSecondFactorType() bool`

HasAuthSecondFactorType returns a boolean if a field has been set.

### GetAuthProvider

`func (o *SignatureRequestInput) GetAuthProvider() string`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *SignatureRequestInput) GetAuthProviderOk() (*string, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *SignatureRequestInput) SetAuthProvider(v string)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *SignatureRequestInput) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### SetAuthProviderNil

`func (o *SignatureRequestInput) SetAuthProviderNil(b bool)`

 SetAuthProviderNil sets the value for AuthProvider to be an explicit nil

### UnsetAuthProvider
`func (o *SignatureRequestInput) UnsetAuthProvider()`

UnsetAuthProvider ensures that no value is present for AuthProvider, not even an explicit nil
### GetAuthSessionStartedAt

`func (o *SignatureRequestInput) GetAuthSessionStartedAt() time.Time`

GetAuthSessionStartedAt returns the AuthSessionStartedAt field if non-nil, zero value otherwise.

### GetAuthSessionStartedAtOk

`func (o *SignatureRequestInput) GetAuthSessionStartedAtOk() (*time.Time, bool)`

GetAuthSessionStartedAtOk returns a tuple with the AuthSessionStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionStartedAt

`func (o *SignatureRequestInput) SetAuthSessionStartedAt(v time.Time)`

SetAuthSessionStartedAt sets AuthSessionStartedAt field to given value.

### HasAuthSessionStartedAt

`func (o *SignatureRequestInput) HasAuthSessionStartedAt() bool`

HasAuthSessionStartedAt returns a boolean if a field has been set.

### SetAuthSessionStartedAtNil

`func (o *SignatureRequestInput) SetAuthSessionStartedAtNil(b bool)`

 SetAuthSessionStartedAtNil sets the value for AuthSessionStartedAt to be an explicit nil

### UnsetAuthSessionStartedAt
`func (o *SignatureRequestInput) UnsetAuthSessionStartedAt()`

UnsetAuthSessionStartedAt ensures that no value is present for AuthSessionStartedAt, not even an explicit nil
### GetAuthSessionIdHash

`func (o *SignatureRequestInput) GetAuthSessionIdHash() string`

GetAuthSessionIdHash returns the AuthSessionIdHash field if non-nil, zero value otherwise.

### GetAuthSessionIdHashOk

`func (o *SignatureRequestInput) GetAuthSessionIdHashOk() (*string, bool)`

GetAuthSessionIdHashOk returns a tuple with the AuthSessionIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionIdHash

`func (o *SignatureRequestInput) SetAuthSessionIdHash(v string)`

SetAuthSessionIdHash sets AuthSessionIdHash field to given value.

### HasAuthSessionIdHash

`func (o *SignatureRequestInput) HasAuthSessionIdHash() bool`

HasAuthSessionIdHash returns a boolean if a field has been set.

### SetAuthSessionIdHashNil

`func (o *SignatureRequestInput) SetAuthSessionIdHashNil(b bool)`

 SetAuthSessionIdHashNil sets the value for AuthSessionIdHash to be an explicit nil

### UnsetAuthSessionIdHash
`func (o *SignatureRequestInput) UnsetAuthSessionIdHash()`

UnsetAuthSessionIdHash ensures that no value is present for AuthSessionIdHash, not even an explicit nil
### GetAuthUserIdHash

`func (o *SignatureRequestInput) GetAuthUserIdHash() string`

GetAuthUserIdHash returns the AuthUserIdHash field if non-nil, zero value otherwise.

### GetAuthUserIdHashOk

`func (o *SignatureRequestInput) GetAuthUserIdHashOk() (*string, bool)`

GetAuthUserIdHashOk returns a tuple with the AuthUserIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserIdHash

`func (o *SignatureRequestInput) SetAuthUserIdHash(v string)`

SetAuthUserIdHash sets AuthUserIdHash field to given value.

### HasAuthUserIdHash

`func (o *SignatureRequestInput) HasAuthUserIdHash() bool`

HasAuthUserIdHash returns a boolean if a field has been set.

### SetAuthUserIdHashNil

`func (o *SignatureRequestInput) SetAuthUserIdHashNil(b bool)`

 SetAuthUserIdHashNil sets the value for AuthUserIdHash to be an explicit nil

### UnsetAuthUserIdHash
`func (o *SignatureRequestInput) UnsetAuthUserIdHash()`

UnsetAuthUserIdHash ensures that no value is present for AuthUserIdHash, not even an explicit nil
### GetAuthUsernameHash

`func (o *SignatureRequestInput) GetAuthUsernameHash() string`

GetAuthUsernameHash returns the AuthUsernameHash field if non-nil, zero value otherwise.

### GetAuthUsernameHashOk

`func (o *SignatureRequestInput) GetAuthUsernameHashOk() (*string, bool)`

GetAuthUsernameHashOk returns a tuple with the AuthUsernameHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsernameHash

`func (o *SignatureRequestInput) SetAuthUsernameHash(v string)`

SetAuthUsernameHash sets AuthUsernameHash field to given value.

### HasAuthUsernameHash

`func (o *SignatureRequestInput) HasAuthUsernameHash() bool`

HasAuthUsernameHash returns a boolean if a field has been set.

### SetAuthUsernameHashNil

`func (o *SignatureRequestInput) SetAuthUsernameHashNil(b bool)`

 SetAuthUsernameHashNil sets the value for AuthUsernameHash to be an explicit nil

### UnsetAuthUsernameHash
`func (o *SignatureRequestInput) UnsetAuthUsernameHash()`

UnsetAuthUsernameHash ensures that no value is present for AuthUsernameHash, not even an explicit nil
### GetAuthPhoneNumberHash

`func (o *SignatureRequestInput) GetAuthPhoneNumberHash() string`

GetAuthPhoneNumberHash returns the AuthPhoneNumberHash field if non-nil, zero value otherwise.

### GetAuthPhoneNumberHashOk

`func (o *SignatureRequestInput) GetAuthPhoneNumberHashOk() (*string, bool)`

GetAuthPhoneNumberHashOk returns a tuple with the AuthPhoneNumberHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPhoneNumberHash

`func (o *SignatureRequestInput) SetAuthPhoneNumberHash(v string)`

SetAuthPhoneNumberHash sets AuthPhoneNumberHash field to given value.

### HasAuthPhoneNumberHash

`func (o *SignatureRequestInput) HasAuthPhoneNumberHash() bool`

HasAuthPhoneNumberHash returns a boolean if a field has been set.

### SetAuthPhoneNumberHashNil

`func (o *SignatureRequestInput) SetAuthPhoneNumberHashNil(b bool)`

 SetAuthPhoneNumberHashNil sets the value for AuthPhoneNumberHash to be an explicit nil

### UnsetAuthPhoneNumberHash
`func (o *SignatureRequestInput) UnsetAuthPhoneNumberHash()`

UnsetAuthPhoneNumberHash ensures that no value is present for AuthPhoneNumberHash, not even an explicit nil
### GetIgnoreUnknownFields

`func (o *SignatureRequestInput) GetIgnoreUnknownFields() bool`

GetIgnoreUnknownFields returns the IgnoreUnknownFields field if non-nil, zero value otherwise.

### GetIgnoreUnknownFieldsOk

`func (o *SignatureRequestInput) GetIgnoreUnknownFieldsOk() (*bool, bool)`

GetIgnoreUnknownFieldsOk returns a tuple with the IgnoreUnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUnknownFields

`func (o *SignatureRequestInput) SetIgnoreUnknownFields(v bool)`

SetIgnoreUnknownFields sets IgnoreUnknownFields field to given value.

### HasIgnoreUnknownFields

`func (o *SignatureRequestInput) HasIgnoreUnknownFields() bool`

HasIgnoreUnknownFields returns a boolean if a field has been set.

### GetPrefilledValuesEditable

`func (o *SignatureRequestInput) GetPrefilledValuesEditable() bool`

GetPrefilledValuesEditable returns the PrefilledValuesEditable field if non-nil, zero value otherwise.

### GetPrefilledValuesEditableOk

`func (o *SignatureRequestInput) GetPrefilledValuesEditableOk() (*bool, bool)`

GetPrefilledValuesEditableOk returns a tuple with the PrefilledValuesEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefilledValuesEditable

`func (o *SignatureRequestInput) SetPrefilledValuesEditable(v bool)`

SetPrefilledValuesEditable sets PrefilledValuesEditable field to given value.

### HasPrefilledValuesEditable

`func (o *SignatureRequestInput) HasPrefilledValuesEditable() bool`

HasPrefilledValuesEditable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


