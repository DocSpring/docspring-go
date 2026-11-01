# EmailDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The email domain's ID. | 
**Domain** | **string** | The domain name, such as `example.com` or `mail.example.com`. Email actions can send from any address at it. | 
**State** | [**EmailDomainState**](EmailDomainState.md) | `pending` until its DNS records are found, then `verified`. `failed`: its records weren't found within 3 days, or were removed. Verify it again after fixing them. | 
**DkimState** | [**EmailDomainRecordState**](EmailDomainRecordState.md) | Whether the DKIM records were found. | 
**MailFromState** | [**EmailDomainRecordState**](EmailDomainRecordState.md) | Whether the MAIL FROM records were found. They're optional: they let SPF pass for your domain too. | 
**DnsRecords** | [**[]EmailDomainDnsRecord**](EmailDomainDnsRecord.md) | The DNS records to add to the domain. Empty while DocSpring is still setting the domain up: verify it again in a few minutes to get them. | 
**DmarcPolicy** | **NullableString** | The policy (`p=`) of the domain's DMARC record, once it has been found. | 
**LastError** | **NullableString** | Why the domain isn't verified yet, or what to fix. | 
**VerifiedAt** | **NullableTime** | When the domain was verified. | 
**CheckedAt** | **NullableTime** | When DocSpring last checked the domain's DNS records. | 
**CreatedAt** | **time.Time** | When the domain was added. | 

## Methods

### NewEmailDomain

`func NewEmailDomain(id string, domain string, state EmailDomainState, dkimState EmailDomainRecordState, mailFromState EmailDomainRecordState, dnsRecords []EmailDomainDnsRecord, dmarcPolicy NullableString, lastError NullableString, verifiedAt NullableTime, checkedAt NullableTime, createdAt time.Time, ) *EmailDomain`

NewEmailDomain instantiates a new EmailDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDomainWithDefaults

`func NewEmailDomainWithDefaults() *EmailDomain`

NewEmailDomainWithDefaults instantiates a new EmailDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailDomain) SetId(v string)`

SetId sets Id field to given value.


### GetDomain

`func (o *EmailDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetState

`func (o *EmailDomain) GetState() EmailDomainState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EmailDomain) GetStateOk() (*EmailDomainState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EmailDomain) SetState(v EmailDomainState)`

SetState sets State field to given value.


### GetDkimState

`func (o *EmailDomain) GetDkimState() EmailDomainRecordState`

GetDkimState returns the DkimState field if non-nil, zero value otherwise.

### GetDkimStateOk

`func (o *EmailDomain) GetDkimStateOk() (*EmailDomainRecordState, bool)`

GetDkimStateOk returns a tuple with the DkimState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimState

`func (o *EmailDomain) SetDkimState(v EmailDomainRecordState)`

SetDkimState sets DkimState field to given value.


### GetMailFromState

`func (o *EmailDomain) GetMailFromState() EmailDomainRecordState`

GetMailFromState returns the MailFromState field if non-nil, zero value otherwise.

### GetMailFromStateOk

`func (o *EmailDomain) GetMailFromStateOk() (*EmailDomainRecordState, bool)`

GetMailFromStateOk returns a tuple with the MailFromState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailFromState

`func (o *EmailDomain) SetMailFromState(v EmailDomainRecordState)`

SetMailFromState sets MailFromState field to given value.


### GetDnsRecords

`func (o *EmailDomain) GetDnsRecords() []EmailDomainDnsRecord`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *EmailDomain) GetDnsRecordsOk() (*[]EmailDomainDnsRecord, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *EmailDomain) SetDnsRecords(v []EmailDomainDnsRecord)`

SetDnsRecords sets DnsRecords field to given value.


### GetDmarcPolicy

`func (o *EmailDomain) GetDmarcPolicy() string`

GetDmarcPolicy returns the DmarcPolicy field if non-nil, zero value otherwise.

### GetDmarcPolicyOk

`func (o *EmailDomain) GetDmarcPolicyOk() (*string, bool)`

GetDmarcPolicyOk returns a tuple with the DmarcPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarcPolicy

`func (o *EmailDomain) SetDmarcPolicy(v string)`

SetDmarcPolicy sets DmarcPolicy field to given value.


### SetDmarcPolicyNil

`func (o *EmailDomain) SetDmarcPolicyNil(b bool)`

 SetDmarcPolicyNil sets the value for DmarcPolicy to be an explicit nil

### UnsetDmarcPolicy
`func (o *EmailDomain) UnsetDmarcPolicy()`

UnsetDmarcPolicy ensures that no value is present for DmarcPolicy, not even an explicit nil
### GetLastError

`func (o *EmailDomain) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *EmailDomain) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *EmailDomain) SetLastError(v string)`

SetLastError sets LastError field to given value.


### SetLastErrorNil

`func (o *EmailDomain) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *EmailDomain) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetVerifiedAt

`func (o *EmailDomain) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *EmailDomain) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *EmailDomain) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.


### SetVerifiedAtNil

`func (o *EmailDomain) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *EmailDomain) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil
### GetCheckedAt

`func (o *EmailDomain) GetCheckedAt() time.Time`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *EmailDomain) GetCheckedAtOk() (*time.Time, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *EmailDomain) SetCheckedAt(v time.Time)`

SetCheckedAt sets CheckedAt field to given value.


### SetCheckedAtNil

`func (o *EmailDomain) SetCheckedAtNil(b bool)`

 SetCheckedAtNil sets the value for CheckedAt to be an explicit nil

### UnsetCheckedAt
`func (o *EmailDomain) UnsetCheckedAt()`

UnsetCheckedAt ensures that no value is present for CheckedAt, not even an explicit nil
### GetCreatedAt

`func (o *EmailDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


