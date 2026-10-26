# CreateSignatureRequestTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**SignatureRequestTokenType**](SignatureRequestTokenType.md) | `api` (the default): the link opens the signature request for 24 hours, for a form your app opens or a link it gives the person right away. `email`: 7 days, for a link you send by email. | [optional] 

## Methods

### NewCreateSignatureRequestTokenRequest

`func NewCreateSignatureRequestTokenRequest() *CreateSignatureRequestTokenRequest`

NewCreateSignatureRequestTokenRequest instantiates a new CreateSignatureRequestTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSignatureRequestTokenRequestWithDefaults

`func NewCreateSignatureRequestTokenRequestWithDefaults() *CreateSignatureRequestTokenRequest`

NewCreateSignatureRequestTokenRequestWithDefaults instantiates a new CreateSignatureRequestTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateSignatureRequestTokenRequest) GetType() SignatureRequestTokenType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSignatureRequestTokenRequest) GetTypeOk() (*SignatureRequestTokenType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSignatureRequestTokenRequest) SetType(v SignatureRequestTokenType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSignatureRequestTokenRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


