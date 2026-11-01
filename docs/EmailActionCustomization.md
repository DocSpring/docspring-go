# EmailActionCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the account email action sends for this template. It also needs to be on itself (`enabled`). | 
**Customized** | [**[]EmailActionSetting**](EmailActionSetting.md) | The settings this template changes. The others follow the account email action. | 
**ConsecutiveFailures** | **int32** | How many of this template's emails in a row have failed because of a setting it changes. | 
**FailingSince** | **NullableTime** | When the current run of those failures started. | 
**LastFailureAt** | **NullableTime** | When one of those emails last failed. | 
**LastFailureMessage** | **NullableString** | Why it failed. | 
**DisabledAt** | **NullableTime** | When DocSpring turned the email action off for this template, such as when an address the template added bounced. | 
**DisabledReason** | **NullableString** | Why DocSpring turned it off for this template. | 
**CreatedAt** | **time.Time** | When the template first changed the email action. | 
**UpdatedAt** | **time.Time** | When the template's changes last changed. | 

## Methods

### NewEmailActionCustomization

`func NewEmailActionCustomization(enabled bool, customized []EmailActionSetting, consecutiveFailures int32, failingSince NullableTime, lastFailureAt NullableTime, lastFailureMessage NullableString, disabledAt NullableTime, disabledReason NullableString, createdAt time.Time, updatedAt time.Time, ) *EmailActionCustomization`

NewEmailActionCustomization instantiates a new EmailActionCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailActionCustomizationWithDefaults

`func NewEmailActionCustomizationWithDefaults() *EmailActionCustomization`

NewEmailActionCustomizationWithDefaults instantiates a new EmailActionCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *EmailActionCustomization) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EmailActionCustomization) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EmailActionCustomization) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCustomized

`func (o *EmailActionCustomization) GetCustomized() []EmailActionSetting`

GetCustomized returns the Customized field if non-nil, zero value otherwise.

### GetCustomizedOk

`func (o *EmailActionCustomization) GetCustomizedOk() (*[]EmailActionSetting, bool)`

GetCustomizedOk returns a tuple with the Customized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomized

`func (o *EmailActionCustomization) SetCustomized(v []EmailActionSetting)`

SetCustomized sets Customized field to given value.


### GetConsecutiveFailures

`func (o *EmailActionCustomization) GetConsecutiveFailures() int32`

GetConsecutiveFailures returns the ConsecutiveFailures field if non-nil, zero value otherwise.

### GetConsecutiveFailuresOk

`func (o *EmailActionCustomization) GetConsecutiveFailuresOk() (*int32, bool)`

GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveFailures

`func (o *EmailActionCustomization) SetConsecutiveFailures(v int32)`

SetConsecutiveFailures sets ConsecutiveFailures field to given value.


### GetFailingSince

`func (o *EmailActionCustomization) GetFailingSince() time.Time`

GetFailingSince returns the FailingSince field if non-nil, zero value otherwise.

### GetFailingSinceOk

`func (o *EmailActionCustomization) GetFailingSinceOk() (*time.Time, bool)`

GetFailingSinceOk returns a tuple with the FailingSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingSince

`func (o *EmailActionCustomization) SetFailingSince(v time.Time)`

SetFailingSince sets FailingSince field to given value.


### SetFailingSinceNil

`func (o *EmailActionCustomization) SetFailingSinceNil(b bool)`

 SetFailingSinceNil sets the value for FailingSince to be an explicit nil

### UnsetFailingSince
`func (o *EmailActionCustomization) UnsetFailingSince()`

UnsetFailingSince ensures that no value is present for FailingSince, not even an explicit nil
### GetLastFailureAt

`func (o *EmailActionCustomization) GetLastFailureAt() time.Time`

GetLastFailureAt returns the LastFailureAt field if non-nil, zero value otherwise.

### GetLastFailureAtOk

`func (o *EmailActionCustomization) GetLastFailureAtOk() (*time.Time, bool)`

GetLastFailureAtOk returns a tuple with the LastFailureAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureAt

`func (o *EmailActionCustomization) SetLastFailureAt(v time.Time)`

SetLastFailureAt sets LastFailureAt field to given value.


### SetLastFailureAtNil

`func (o *EmailActionCustomization) SetLastFailureAtNil(b bool)`

 SetLastFailureAtNil sets the value for LastFailureAt to be an explicit nil

### UnsetLastFailureAt
`func (o *EmailActionCustomization) UnsetLastFailureAt()`

UnsetLastFailureAt ensures that no value is present for LastFailureAt, not even an explicit nil
### GetLastFailureMessage

`func (o *EmailActionCustomization) GetLastFailureMessage() string`

GetLastFailureMessage returns the LastFailureMessage field if non-nil, zero value otherwise.

### GetLastFailureMessageOk

`func (o *EmailActionCustomization) GetLastFailureMessageOk() (*string, bool)`

GetLastFailureMessageOk returns a tuple with the LastFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureMessage

`func (o *EmailActionCustomization) SetLastFailureMessage(v string)`

SetLastFailureMessage sets LastFailureMessage field to given value.


### SetLastFailureMessageNil

`func (o *EmailActionCustomization) SetLastFailureMessageNil(b bool)`

 SetLastFailureMessageNil sets the value for LastFailureMessage to be an explicit nil

### UnsetLastFailureMessage
`func (o *EmailActionCustomization) UnsetLastFailureMessage()`

UnsetLastFailureMessage ensures that no value is present for LastFailureMessage, not even an explicit nil
### GetDisabledAt

`func (o *EmailActionCustomization) GetDisabledAt() time.Time`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *EmailActionCustomization) GetDisabledAtOk() (*time.Time, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *EmailActionCustomization) SetDisabledAt(v time.Time)`

SetDisabledAt sets DisabledAt field to given value.


### SetDisabledAtNil

`func (o *EmailActionCustomization) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *EmailActionCustomization) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetDisabledReason

`func (o *EmailActionCustomization) GetDisabledReason() string`

GetDisabledReason returns the DisabledReason field if non-nil, zero value otherwise.

### GetDisabledReasonOk

`func (o *EmailActionCustomization) GetDisabledReasonOk() (*string, bool)`

GetDisabledReasonOk returns a tuple with the DisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReason

`func (o *EmailActionCustomization) SetDisabledReason(v string)`

SetDisabledReason sets DisabledReason field to given value.


### SetDisabledReasonNil

`func (o *EmailActionCustomization) SetDisabledReasonNil(b bool)`

 SetDisabledReasonNil sets the value for DisabledReason to be an explicit nil

### UnsetDisabledReason
`func (o *EmailActionCustomization) UnsetDisabledReason()`

UnsetDisabledReason ensures that no value is present for DisabledReason, not even an explicit nil
### GetCreatedAt

`func (o *EmailActionCustomization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailActionCustomization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailActionCustomization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EmailActionCustomization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EmailActionCustomization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EmailActionCustomization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


