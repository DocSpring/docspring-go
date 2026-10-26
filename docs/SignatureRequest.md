# SignatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The signature request's ID. | 
**SubmissionId** | **string** | The ID of the submission the signature request belongs to. | 
**Name** | **NullableString** | The name of the person the request is for. | 
**Email** | **NullableString** | Their email address. | 
**Fields** | **[]string** | The fields this person fills in. Null for a signature request created with API v1 without a field list: it asks for every field the submission's data doesn't have. | 
**Metadata** | **map[string]interface{}** | Your own data about the signature request, such as an ID from your app. | 
**Role** | **string** | Which of the template's signers the person is. Until templates have roles of their own, a template's signers are `Signer 1`, `Signer 2` and so on, up to `Signer 100`: `Signer 1` is the one whose details the template's `@@signature_requests/0` fields print. | 
**SortOrder** | **int32** | The signer's position, from 0: `Signer 1` is 0. The template's `@@signature_requests/&lt;sort_order&gt;` fields print this signature request's details. | 
**State** | [**SignatureRequestState**](SignatureRequestState.md) | `pending` until the person completes the signature request (`completed`) or declines it (`declined`). | 
**Data** | **map[string]interface{}** | The values this person entered for their own fields (`fields`). Null when their submission's data is left out: in lists without `include_data=true`, in the events of a webhook that leaves out submission data, and once the submission has expired. | 
**ViewedAt** | **NullableTime** | When the person first opened the signature request. | 
**CompletedAt** | **NullableTime** | When they completed it. | 
**DeclinedAt** | **NullableTime** | When they declined it. | 
**AuthType** | [**SignatureRequestAuthType**](SignatureRequestAuthType.md) | How your app authenticated the person before it showed them the signature request. | 
**AuthSecondFactorType** | [**SignatureRequestAuthSecondFactorType**](SignatureRequestAuthSecondFactorType.md) | The second factor your app authenticated them with, if any. | 
**AuthProvider** | **NullableString** | The service that authenticated them, such as `google`. Required when `auth_type` is `oauth`. | 
**AuthSessionStartedAt** | **NullableTime** | When your app authenticated them. Set for you with `email_link`. | 
**AuthSessionIdHash** | **NullableString** | A hash of your app's session ID, for the audit trail. | 
**AuthUserIdHash** | **NullableString** | Their user ID in your app, or a SHA-256 hash of a salt and the user ID. | 
**AuthUsernameHash** | **NullableString** | Their username in your app, or a SHA-256 hash of a salt and the username. | 
**AuthPhoneNumberHash** | **NullableString** | The phone number they authenticated with, or a SHA-256 hash of a salt and the phone number. | 
**IpAddress** | **NullableString** | The IP address they completed the signature request from. Null in the events of a webhook that leaves out submission data. | 
**UserAgent** | **NullableString** | The User-Agent header of the browser they completed the signature request with. Null in the events of a webhook that leaves out submission data. | 
**CreatedAt** | **time.Time** | When the signature request was created. | 
**UpdatedAt** | **time.Time** | When the signature request was last updated. | 

## Methods

### NewSignatureRequest

`func NewSignatureRequest(id string, submissionId string, name NullableString, email NullableString, fields []string, metadata map[string]interface{}, role string, sortOrder int32, state SignatureRequestState, data map[string]interface{}, viewedAt NullableTime, completedAt NullableTime, declinedAt NullableTime, authType SignatureRequestAuthType, authSecondFactorType SignatureRequestAuthSecondFactorType, authProvider NullableString, authSessionStartedAt NullableTime, authSessionIdHash NullableString, authUserIdHash NullableString, authUsernameHash NullableString, authPhoneNumberHash NullableString, ipAddress NullableString, userAgent NullableString, createdAt time.Time, updatedAt time.Time, ) *SignatureRequest`

NewSignatureRequest instantiates a new SignatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureRequestWithDefaults

`func NewSignatureRequestWithDefaults() *SignatureRequest`

NewSignatureRequestWithDefaults instantiates a new SignatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignatureRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignatureRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignatureRequest) SetId(v string)`

SetId sets Id field to given value.


### GetSubmissionId

`func (o *SignatureRequest) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *SignatureRequest) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *SignatureRequest) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.


### GetName

`func (o *SignatureRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignatureRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignatureRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *SignatureRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SignatureRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *SignatureRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignatureRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignatureRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *SignatureRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SignatureRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFields

`func (o *SignatureRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SignatureRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SignatureRequest) SetFields(v []string)`

SetFields sets Fields field to given value.


### SetFieldsNil

`func (o *SignatureRequest) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *SignatureRequest) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetMetadata

`func (o *SignatureRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SignatureRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SignatureRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetRole

`func (o *SignatureRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SignatureRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SignatureRequest) SetRole(v string)`

SetRole sets Role field to given value.


### GetSortOrder

`func (o *SignatureRequest) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *SignatureRequest) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *SignatureRequest) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetState

`func (o *SignatureRequest) GetState() SignatureRequestState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SignatureRequest) GetStateOk() (*SignatureRequestState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SignatureRequest) SetState(v SignatureRequestState)`

SetState sets State field to given value.


### GetData

`func (o *SignatureRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SignatureRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SignatureRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *SignatureRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SignatureRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetViewedAt

`func (o *SignatureRequest) GetViewedAt() time.Time`

GetViewedAt returns the ViewedAt field if non-nil, zero value otherwise.

### GetViewedAtOk

`func (o *SignatureRequest) GetViewedAtOk() (*time.Time, bool)`

GetViewedAtOk returns a tuple with the ViewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedAt

`func (o *SignatureRequest) SetViewedAt(v time.Time)`

SetViewedAt sets ViewedAt field to given value.


### SetViewedAtNil

`func (o *SignatureRequest) SetViewedAtNil(b bool)`

 SetViewedAtNil sets the value for ViewedAt to be an explicit nil

### UnsetViewedAt
`func (o *SignatureRequest) UnsetViewedAt()`

UnsetViewedAt ensures that no value is present for ViewedAt, not even an explicit nil
### GetCompletedAt

`func (o *SignatureRequest) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *SignatureRequest) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *SignatureRequest) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *SignatureRequest) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *SignatureRequest) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetDeclinedAt

`func (o *SignatureRequest) GetDeclinedAt() time.Time`

GetDeclinedAt returns the DeclinedAt field if non-nil, zero value otherwise.

### GetDeclinedAtOk

`func (o *SignatureRequest) GetDeclinedAtOk() (*time.Time, bool)`

GetDeclinedAtOk returns a tuple with the DeclinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclinedAt

`func (o *SignatureRequest) SetDeclinedAt(v time.Time)`

SetDeclinedAt sets DeclinedAt field to given value.


### SetDeclinedAtNil

`func (o *SignatureRequest) SetDeclinedAtNil(b bool)`

 SetDeclinedAtNil sets the value for DeclinedAt to be an explicit nil

### UnsetDeclinedAt
`func (o *SignatureRequest) UnsetDeclinedAt()`

UnsetDeclinedAt ensures that no value is present for DeclinedAt, not even an explicit nil
### GetAuthType

`func (o *SignatureRequest) GetAuthType() SignatureRequestAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *SignatureRequest) GetAuthTypeOk() (*SignatureRequestAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *SignatureRequest) SetAuthType(v SignatureRequestAuthType)`

SetAuthType sets AuthType field to given value.


### GetAuthSecondFactorType

`func (o *SignatureRequest) GetAuthSecondFactorType() SignatureRequestAuthSecondFactorType`

GetAuthSecondFactorType returns the AuthSecondFactorType field if non-nil, zero value otherwise.

### GetAuthSecondFactorTypeOk

`func (o *SignatureRequest) GetAuthSecondFactorTypeOk() (*SignatureRequestAuthSecondFactorType, bool)`

GetAuthSecondFactorTypeOk returns a tuple with the AuthSecondFactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecondFactorType

`func (o *SignatureRequest) SetAuthSecondFactorType(v SignatureRequestAuthSecondFactorType)`

SetAuthSecondFactorType sets AuthSecondFactorType field to given value.


### GetAuthProvider

`func (o *SignatureRequest) GetAuthProvider() string`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *SignatureRequest) GetAuthProviderOk() (*string, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *SignatureRequest) SetAuthProvider(v string)`

SetAuthProvider sets AuthProvider field to given value.


### SetAuthProviderNil

`func (o *SignatureRequest) SetAuthProviderNil(b bool)`

 SetAuthProviderNil sets the value for AuthProvider to be an explicit nil

### UnsetAuthProvider
`func (o *SignatureRequest) UnsetAuthProvider()`

UnsetAuthProvider ensures that no value is present for AuthProvider, not even an explicit nil
### GetAuthSessionStartedAt

`func (o *SignatureRequest) GetAuthSessionStartedAt() time.Time`

GetAuthSessionStartedAt returns the AuthSessionStartedAt field if non-nil, zero value otherwise.

### GetAuthSessionStartedAtOk

`func (o *SignatureRequest) GetAuthSessionStartedAtOk() (*time.Time, bool)`

GetAuthSessionStartedAtOk returns a tuple with the AuthSessionStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionStartedAt

`func (o *SignatureRequest) SetAuthSessionStartedAt(v time.Time)`

SetAuthSessionStartedAt sets AuthSessionStartedAt field to given value.


### SetAuthSessionStartedAtNil

`func (o *SignatureRequest) SetAuthSessionStartedAtNil(b bool)`

 SetAuthSessionStartedAtNil sets the value for AuthSessionStartedAt to be an explicit nil

### UnsetAuthSessionStartedAt
`func (o *SignatureRequest) UnsetAuthSessionStartedAt()`

UnsetAuthSessionStartedAt ensures that no value is present for AuthSessionStartedAt, not even an explicit nil
### GetAuthSessionIdHash

`func (o *SignatureRequest) GetAuthSessionIdHash() string`

GetAuthSessionIdHash returns the AuthSessionIdHash field if non-nil, zero value otherwise.

### GetAuthSessionIdHashOk

`func (o *SignatureRequest) GetAuthSessionIdHashOk() (*string, bool)`

GetAuthSessionIdHashOk returns a tuple with the AuthSessionIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionIdHash

`func (o *SignatureRequest) SetAuthSessionIdHash(v string)`

SetAuthSessionIdHash sets AuthSessionIdHash field to given value.


### SetAuthSessionIdHashNil

`func (o *SignatureRequest) SetAuthSessionIdHashNil(b bool)`

 SetAuthSessionIdHashNil sets the value for AuthSessionIdHash to be an explicit nil

### UnsetAuthSessionIdHash
`func (o *SignatureRequest) UnsetAuthSessionIdHash()`

UnsetAuthSessionIdHash ensures that no value is present for AuthSessionIdHash, not even an explicit nil
### GetAuthUserIdHash

`func (o *SignatureRequest) GetAuthUserIdHash() string`

GetAuthUserIdHash returns the AuthUserIdHash field if non-nil, zero value otherwise.

### GetAuthUserIdHashOk

`func (o *SignatureRequest) GetAuthUserIdHashOk() (*string, bool)`

GetAuthUserIdHashOk returns a tuple with the AuthUserIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserIdHash

`func (o *SignatureRequest) SetAuthUserIdHash(v string)`

SetAuthUserIdHash sets AuthUserIdHash field to given value.


### SetAuthUserIdHashNil

`func (o *SignatureRequest) SetAuthUserIdHashNil(b bool)`

 SetAuthUserIdHashNil sets the value for AuthUserIdHash to be an explicit nil

### UnsetAuthUserIdHash
`func (o *SignatureRequest) UnsetAuthUserIdHash()`

UnsetAuthUserIdHash ensures that no value is present for AuthUserIdHash, not even an explicit nil
### GetAuthUsernameHash

`func (o *SignatureRequest) GetAuthUsernameHash() string`

GetAuthUsernameHash returns the AuthUsernameHash field if non-nil, zero value otherwise.

### GetAuthUsernameHashOk

`func (o *SignatureRequest) GetAuthUsernameHashOk() (*string, bool)`

GetAuthUsernameHashOk returns a tuple with the AuthUsernameHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsernameHash

`func (o *SignatureRequest) SetAuthUsernameHash(v string)`

SetAuthUsernameHash sets AuthUsernameHash field to given value.


### SetAuthUsernameHashNil

`func (o *SignatureRequest) SetAuthUsernameHashNil(b bool)`

 SetAuthUsernameHashNil sets the value for AuthUsernameHash to be an explicit nil

### UnsetAuthUsernameHash
`func (o *SignatureRequest) UnsetAuthUsernameHash()`

UnsetAuthUsernameHash ensures that no value is present for AuthUsernameHash, not even an explicit nil
### GetAuthPhoneNumberHash

`func (o *SignatureRequest) GetAuthPhoneNumberHash() string`

GetAuthPhoneNumberHash returns the AuthPhoneNumberHash field if non-nil, zero value otherwise.

### GetAuthPhoneNumberHashOk

`func (o *SignatureRequest) GetAuthPhoneNumberHashOk() (*string, bool)`

GetAuthPhoneNumberHashOk returns a tuple with the AuthPhoneNumberHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPhoneNumberHash

`func (o *SignatureRequest) SetAuthPhoneNumberHash(v string)`

SetAuthPhoneNumberHash sets AuthPhoneNumberHash field to given value.


### SetAuthPhoneNumberHashNil

`func (o *SignatureRequest) SetAuthPhoneNumberHashNil(b bool)`

 SetAuthPhoneNumberHashNil sets the value for AuthPhoneNumberHash to be an explicit nil

### UnsetAuthPhoneNumberHash
`func (o *SignatureRequest) UnsetAuthPhoneNumberHash()`

UnsetAuthPhoneNumberHash ensures that no value is present for AuthPhoneNumberHash, not even an explicit nil
### GetIpAddress

`func (o *SignatureRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SignatureRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SignatureRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### SetIpAddressNil

`func (o *SignatureRequest) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *SignatureRequest) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetUserAgent

`func (o *SignatureRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *SignatureRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *SignatureRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### SetUserAgentNil

`func (o *SignatureRequest) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *SignatureRequest) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetCreatedAt

`func (o *SignatureRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SignatureRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SignatureRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SignatureRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SignatureRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SignatureRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


