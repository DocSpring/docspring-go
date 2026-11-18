# SupportMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The message's ID. | 
**Author** | [**SupportMessageAuthor**](SupportMessageAuthor.md) | `reporter`: your side. `team`: a person at DocSpring. `system`: DocSpring itself. | 
**Body** | **string** | The message, as plain text. | 
**CreatedAt** | **time.Time** | When the message was sent. | 
**FirstReadAt** | **NullableTime** | When one of the team's messages was first returned by the API. It shows that the message was fetched, not that a person read it. Null for your own messages. | 

## Methods

### NewSupportMessage

`func NewSupportMessage(id string, author SupportMessageAuthor, body string, createdAt time.Time, firstReadAt NullableTime, ) *SupportMessage`

NewSupportMessage instantiates a new SupportMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportMessageWithDefaults

`func NewSupportMessageWithDefaults() *SupportMessage`

NewSupportMessageWithDefaults instantiates a new SupportMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportMessage) SetId(v string)`

SetId sets Id field to given value.


### GetAuthor

`func (o *SupportMessage) GetAuthor() SupportMessageAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SupportMessage) GetAuthorOk() (*SupportMessageAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SupportMessage) SetAuthor(v SupportMessageAuthor)`

SetAuthor sets Author field to given value.


### GetBody

`func (o *SupportMessage) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SupportMessage) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SupportMessage) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *SupportMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SupportMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SupportMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFirstReadAt

`func (o *SupportMessage) GetFirstReadAt() time.Time`

GetFirstReadAt returns the FirstReadAt field if non-nil, zero value otherwise.

### GetFirstReadAtOk

`func (o *SupportMessage) GetFirstReadAtOk() (*time.Time, bool)`

GetFirstReadAtOk returns a tuple with the FirstReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstReadAt

`func (o *SupportMessage) SetFirstReadAt(v time.Time)`

SetFirstReadAt sets FirstReadAt field to given value.


### SetFirstReadAtNil

`func (o *SupportMessage) SetFirstReadAtNil(b bool)`

 SetFirstReadAtNil sets the value for FirstReadAt to be an explicit nil

### UnsetFirstReadAt
`func (o *SupportMessage) UnsetFirstReadAt()`

UnsetFirstReadAt ensures that no value is present for FirstReadAt, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


