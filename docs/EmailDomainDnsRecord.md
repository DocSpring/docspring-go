# EmailDomainDnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purpose** | [**EmailDomainDnsRecordPurpose**](EmailDomainDnsRecordPurpose.md) | What the record is for. `dkim` signs the emails (required), `mail_from` lets SPF pass for your domain (recommended), and `dmarc` tells inboxes what to do with emails that fail these checks (required). | 
**Type** | [**DnsRecordType**](DnsRecordType.md) | The record's type. | 
**Name** | **string** | The record's name. Some DNS hosts add the domain to the end of a name by themselves. | 
**Value** | **string** | The record's value. | 

## Methods

### NewEmailDomainDnsRecord

`func NewEmailDomainDnsRecord(purpose EmailDomainDnsRecordPurpose, type_ DnsRecordType, name string, value string, ) *EmailDomainDnsRecord`

NewEmailDomainDnsRecord instantiates a new EmailDomainDnsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDomainDnsRecordWithDefaults

`func NewEmailDomainDnsRecordWithDefaults() *EmailDomainDnsRecord`

NewEmailDomainDnsRecordWithDefaults instantiates a new EmailDomainDnsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurpose

`func (o *EmailDomainDnsRecord) GetPurpose() EmailDomainDnsRecordPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *EmailDomainDnsRecord) GetPurposeOk() (*EmailDomainDnsRecordPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *EmailDomainDnsRecord) SetPurpose(v EmailDomainDnsRecordPurpose)`

SetPurpose sets Purpose field to given value.


### GetType

`func (o *EmailDomainDnsRecord) GetType() DnsRecordType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailDomainDnsRecord) GetTypeOk() (*DnsRecordType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailDomainDnsRecord) SetType(v DnsRecordType)`

SetType sets Type field to given value.


### GetName

`func (o *EmailDomainDnsRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailDomainDnsRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailDomainDnsRecord) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *EmailDomainDnsRecord) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EmailDomainDnsRecord) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EmailDomainDnsRecord) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


