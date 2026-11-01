# SignatureRequestReminders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalDays** | **int32** | Days between reminders. | 
**MaxCount** | **int32** | The most reminders each person gets. | 

## Methods

### NewSignatureRequestReminders

`func NewSignatureRequestReminders(intervalDays int32, maxCount int32, ) *SignatureRequestReminders`

NewSignatureRequestReminders instantiates a new SignatureRequestReminders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureRequestRemindersWithDefaults

`func NewSignatureRequestRemindersWithDefaults() *SignatureRequestReminders`

NewSignatureRequestRemindersWithDefaults instantiates a new SignatureRequestReminders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalDays

`func (o *SignatureRequestReminders) GetIntervalDays() int32`

GetIntervalDays returns the IntervalDays field if non-nil, zero value otherwise.

### GetIntervalDaysOk

`func (o *SignatureRequestReminders) GetIntervalDaysOk() (*int32, bool)`

GetIntervalDaysOk returns a tuple with the IntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalDays

`func (o *SignatureRequestReminders) SetIntervalDays(v int32)`

SetIntervalDays sets IntervalDays field to given value.


### GetMaxCount

`func (o *SignatureRequestReminders) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *SignatureRequestReminders) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *SignatureRequestReminders) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


