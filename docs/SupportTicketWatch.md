# SupportTicketWatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The ticket's `watch_url`. | 
**MinIntervalSeconds** | **int32** | Poll no more often than this. | 
**How** | **string** | What to do with the URL, in a sentence for an AI agent to act on. | 

## Methods

### NewSupportTicketWatch

`func NewSupportTicketWatch(url string, minIntervalSeconds int32, how string, ) *SupportTicketWatch`

NewSupportTicketWatch instantiates a new SupportTicketWatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketWatchWithDefaults

`func NewSupportTicketWatchWithDefaults() *SupportTicketWatch`

NewSupportTicketWatchWithDefaults instantiates a new SupportTicketWatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SupportTicketWatch) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SupportTicketWatch) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SupportTicketWatch) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMinIntervalSeconds

`func (o *SupportTicketWatch) GetMinIntervalSeconds() int32`

GetMinIntervalSeconds returns the MinIntervalSeconds field if non-nil, zero value otherwise.

### GetMinIntervalSecondsOk

`func (o *SupportTicketWatch) GetMinIntervalSecondsOk() (*int32, bool)`

GetMinIntervalSecondsOk returns a tuple with the MinIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIntervalSeconds

`func (o *SupportTicketWatch) SetMinIntervalSeconds(v int32)`

SetMinIntervalSeconds sets MinIntervalSeconds field to given value.


### GetHow

`func (o *SupportTicketWatch) GetHow() string`

GetHow returns the How field if non-nil, zero value otherwise.

### GetHowOk

`func (o *SupportTicketWatch) GetHowOk() (*string, bool)`

GetHowOk returns a tuple with the How field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHow

`func (o *SupportTicketWatch) SetHow(v string)`

SetHow sets How field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


