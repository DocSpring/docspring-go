# SupportTicketNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** | The masked address of the user the reply is emailed to: the user the API token or connected app belongs to. Null when that user was removed. | 
**Watch** | [**SupportTicketWatch**](SupportTicketWatch.md) |  | 

## Methods

### NewSupportTicketNotifications

`func NewSupportTicketNotifications(email NullableString, watch SupportTicketWatch, ) *SupportTicketNotifications`

NewSupportTicketNotifications instantiates a new SupportTicketNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketNotificationsWithDefaults

`func NewSupportTicketNotificationsWithDefaults() *SupportTicketNotifications`

NewSupportTicketNotificationsWithDefaults instantiates a new SupportTicketNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SupportTicketNotifications) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SupportTicketNotifications) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SupportTicketNotifications) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *SupportTicketNotifications) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SupportTicketNotifications) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetWatch

`func (o *SupportTicketNotifications) GetWatch() SupportTicketWatch`

GetWatch returns the Watch field if non-nil, zero value otherwise.

### GetWatchOk

`func (o *SupportTicketNotifications) GetWatchOk() (*SupportTicketWatch, bool)`

GetWatchOk returns a tuple with the Watch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatch

`func (o *SupportTicketNotifications) SetWatch(v SupportTicketWatch)`

SetWatch sets Watch field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


