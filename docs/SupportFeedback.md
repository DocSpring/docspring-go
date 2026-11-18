# SupportFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The feedback's ID. | 
**ReceivedAt** | **time.Time** | When DocSpring received it. | 
**FollowUp** | **string** | What happens next, in a sentence for an AI agent to act on. Nobody replies to feedback. | 

## Methods

### NewSupportFeedback

`func NewSupportFeedback(id string, receivedAt time.Time, followUp string, ) *SupportFeedback`

NewSupportFeedback instantiates a new SupportFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportFeedbackWithDefaults

`func NewSupportFeedbackWithDefaults() *SupportFeedback`

NewSupportFeedbackWithDefaults instantiates a new SupportFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportFeedback) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportFeedback) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportFeedback) SetId(v string)`

SetId sets Id field to given value.


### GetReceivedAt

`func (o *SupportFeedback) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *SupportFeedback) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *SupportFeedback) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.


### GetFollowUp

`func (o *SupportFeedback) GetFollowUp() string`

GetFollowUp returns the FollowUp field if non-nil, zero value otherwise.

### GetFollowUpOk

`func (o *SupportFeedback) GetFollowUpOk() (*string, bool)`

GetFollowUpOk returns a tuple with the FollowUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUp

`func (o *SupportFeedback) SetFollowUp(v string)`

SetFollowUp sets FollowUp field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


