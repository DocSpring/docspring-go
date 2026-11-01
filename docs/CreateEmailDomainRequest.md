# CreateEmailDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain name, such as `example.com` or `mail.example.com`. Email actions can send from any address at it. | 

## Methods

### NewCreateEmailDomainRequest

`func NewCreateEmailDomainRequest(domain string, ) *CreateEmailDomainRequest`

NewCreateEmailDomainRequest instantiates a new CreateEmailDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmailDomainRequestWithDefaults

`func NewCreateEmailDomainRequestWithDefaults() *CreateEmailDomainRequest`

NewCreateEmailDomainRequestWithDefaults instantiates a new CreateEmailDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CreateEmailDomainRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateEmailDomainRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateEmailDomainRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


