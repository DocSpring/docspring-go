# SignatureRequestToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The token's ID. | 
**SignatureRequestId** | **string** | The ID of the signature request the token opens. | 
**Type** | [**SignatureRequestTokenType**](SignatureRequestTokenType.md) | `api`: the link opens the signature request for 24 hours. `email`: 7 days. | 
**Secret** | **string** | The token's secret. It's only returned when the token is created. | 
**SignatureRequestUrl** | **string** | A link that opens the signature request with this token. | 
**ExpiresAt** | **time.Time** | When the link stops opening the signature request. | 
**CreatedAt** | **time.Time** | When the token was created. | 

## Methods

### NewSignatureRequestToken

`func NewSignatureRequestToken(id string, signatureRequestId string, type_ SignatureRequestTokenType, secret string, signatureRequestUrl string, expiresAt time.Time, createdAt time.Time, ) *SignatureRequestToken`

NewSignatureRequestToken instantiates a new SignatureRequestToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureRequestTokenWithDefaults

`func NewSignatureRequestTokenWithDefaults() *SignatureRequestToken`

NewSignatureRequestTokenWithDefaults instantiates a new SignatureRequestToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignatureRequestToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignatureRequestToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignatureRequestToken) SetId(v string)`

SetId sets Id field to given value.


### GetSignatureRequestId

`func (o *SignatureRequestToken) GetSignatureRequestId() string`

GetSignatureRequestId returns the SignatureRequestId field if non-nil, zero value otherwise.

### GetSignatureRequestIdOk

`func (o *SignatureRequestToken) GetSignatureRequestIdOk() (*string, bool)`

GetSignatureRequestIdOk returns a tuple with the SignatureRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureRequestId

`func (o *SignatureRequestToken) SetSignatureRequestId(v string)`

SetSignatureRequestId sets SignatureRequestId field to given value.


### GetType

`func (o *SignatureRequestToken) GetType() SignatureRequestTokenType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignatureRequestToken) GetTypeOk() (*SignatureRequestTokenType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignatureRequestToken) SetType(v SignatureRequestTokenType)`

SetType sets Type field to given value.


### GetSecret

`func (o *SignatureRequestToken) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SignatureRequestToken) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SignatureRequestToken) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetSignatureRequestUrl

`func (o *SignatureRequestToken) GetSignatureRequestUrl() string`

GetSignatureRequestUrl returns the SignatureRequestUrl field if non-nil, zero value otherwise.

### GetSignatureRequestUrlOk

`func (o *SignatureRequestToken) GetSignatureRequestUrlOk() (*string, bool)`

GetSignatureRequestUrlOk returns a tuple with the SignatureRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureRequestUrl

`func (o *SignatureRequestToken) SetSignatureRequestUrl(v string)`

SetSignatureRequestUrl sets SignatureRequestUrl field to given value.


### GetExpiresAt

`func (o *SignatureRequestToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SignatureRequestToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SignatureRequestToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetCreatedAt

`func (o *SignatureRequestToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SignatureRequestToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SignatureRequestToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


