# SendAccountEmailActionTestEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | Pointer to **string** | A template the email action sends for: the test uses its PDF and name. Left out, it uses the one that changed most recently. | [optional] 

## Methods

### NewSendAccountEmailActionTestEmailRequest

`func NewSendAccountEmailActionTestEmailRequest() *SendAccountEmailActionTestEmailRequest`

NewSendAccountEmailActionTestEmailRequest instantiates a new SendAccountEmailActionTestEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendAccountEmailActionTestEmailRequestWithDefaults

`func NewSendAccountEmailActionTestEmailRequestWithDefaults() *SendAccountEmailActionTestEmailRequest`

NewSendAccountEmailActionTestEmailRequestWithDefaults instantiates a new SendAccountEmailActionTestEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *SendAccountEmailActionTestEmailRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SendAccountEmailActionTestEmailRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SendAccountEmailActionTestEmailRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *SendAccountEmailActionTestEmailRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


