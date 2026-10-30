# CreateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | A name for the webhook, shown in the dashboard. | [optional] 
**Url** | **string** | The endpoint that receives webhook POST requests. | 
**Enabled** | Pointer to **bool** | Whether DocSpring sends events to the webhook. Webhooks are disabled after repeated failures. | [optional] 
**EventTypes** | [**[]WebhookEventType**](WebhookEventType.md) | The events the webhook subscribes to. A webhook made in the dashboard can also subscribe to API token, account integration and user events, which are listed with their own names. | 
**Mode** | Pointer to [**NullableWebhookMode**](WebhookMode.md) | Only send events for test or live submissions. Null sends both. | [optional] 
**ApiTokenIds** | Pointer to **[]string** | Only send events from these API tokens. Combined with the other scope lists using OR. Applies to submission, signature request, combined PDF, submission batch and custom file events. | [optional] 
**TemplateIds** | Pointer to **[]string** | Only send events for these templates. Combined with the other scope lists using OR. Applies to submission, signature request and template events. | [optional] 
**FolderIds** | Pointer to **[]string** | Only send events for templates in these folders, including their subfolders. Combined with the other scope lists using OR. Applies to submission, signature request, template and folder events. | [optional] 
**IncludeSubmissionData** | Pointer to **bool** | Include the submission's field data in the payload. When false, the submission's data and truncated_text, and each signature request's data, ip_address and user_agent, are null. | [optional] 
**StripFields** | Pointer to **[]string** | Dot-separated paths to remove from the event's data object. Form fields live under `data` (`data.ssn` removes the ssn field). A bare name removes a top-level key (`metadata` removes the whole metadata object). | [optional] 
**Headers** | Pointer to **map[string]string** | Headers to send with each request. | [optional] 

## Methods

### NewCreateWebhookRequest

`func NewCreateWebhookRequest(url string, eventTypes []WebhookEventType, ) *CreateWebhookRequest`

NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookRequestWithDefaults

`func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest`

NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateWebhookRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateWebhookRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *CreateWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabled

`func (o *CreateWebhookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateWebhookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateWebhookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateWebhookRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypes

`func (o *CreateWebhookRequest) GetEventTypes() []WebhookEventType`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *CreateWebhookRequest) GetEventTypesOk() (*[]WebhookEventType, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *CreateWebhookRequest) SetEventTypes(v []WebhookEventType)`

SetEventTypes sets EventTypes field to given value.


### GetMode

`func (o *CreateWebhookRequest) GetMode() WebhookMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateWebhookRequest) GetModeOk() (*WebhookMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateWebhookRequest) SetMode(v WebhookMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateWebhookRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *CreateWebhookRequest) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *CreateWebhookRequest) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetApiTokenIds

`func (o *CreateWebhookRequest) GetApiTokenIds() []string`

GetApiTokenIds returns the ApiTokenIds field if non-nil, zero value otherwise.

### GetApiTokenIdsOk

`func (o *CreateWebhookRequest) GetApiTokenIdsOk() (*[]string, bool)`

GetApiTokenIdsOk returns a tuple with the ApiTokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenIds

`func (o *CreateWebhookRequest) SetApiTokenIds(v []string)`

SetApiTokenIds sets ApiTokenIds field to given value.

### HasApiTokenIds

`func (o *CreateWebhookRequest) HasApiTokenIds() bool`

HasApiTokenIds returns a boolean if a field has been set.

### GetTemplateIds

`func (o *CreateWebhookRequest) GetTemplateIds() []string`

GetTemplateIds returns the TemplateIds field if non-nil, zero value otherwise.

### GetTemplateIdsOk

`func (o *CreateWebhookRequest) GetTemplateIdsOk() (*[]string, bool)`

GetTemplateIdsOk returns a tuple with the TemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIds

`func (o *CreateWebhookRequest) SetTemplateIds(v []string)`

SetTemplateIds sets TemplateIds field to given value.

### HasTemplateIds

`func (o *CreateWebhookRequest) HasTemplateIds() bool`

HasTemplateIds returns a boolean if a field has been set.

### GetFolderIds

`func (o *CreateWebhookRequest) GetFolderIds() []string`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *CreateWebhookRequest) GetFolderIdsOk() (*[]string, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *CreateWebhookRequest) SetFolderIds(v []string)`

SetFolderIds sets FolderIds field to given value.

### HasFolderIds

`func (o *CreateWebhookRequest) HasFolderIds() bool`

HasFolderIds returns a boolean if a field has been set.

### GetIncludeSubmissionData

`func (o *CreateWebhookRequest) GetIncludeSubmissionData() bool`

GetIncludeSubmissionData returns the IncludeSubmissionData field if non-nil, zero value otherwise.

### GetIncludeSubmissionDataOk

`func (o *CreateWebhookRequest) GetIncludeSubmissionDataOk() (*bool, bool)`

GetIncludeSubmissionDataOk returns a tuple with the IncludeSubmissionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubmissionData

`func (o *CreateWebhookRequest) SetIncludeSubmissionData(v bool)`

SetIncludeSubmissionData sets IncludeSubmissionData field to given value.

### HasIncludeSubmissionData

`func (o *CreateWebhookRequest) HasIncludeSubmissionData() bool`

HasIncludeSubmissionData returns a boolean if a field has been set.

### GetStripFields

`func (o *CreateWebhookRequest) GetStripFields() []string`

GetStripFields returns the StripFields field if non-nil, zero value otherwise.

### GetStripFieldsOk

`func (o *CreateWebhookRequest) GetStripFieldsOk() (*[]string, bool)`

GetStripFieldsOk returns a tuple with the StripFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripFields

`func (o *CreateWebhookRequest) SetStripFields(v []string)`

SetStripFields sets StripFields field to given value.

### HasStripFields

`func (o *CreateWebhookRequest) HasStripFields() bool`

HasStripFields returns a boolean if a field has been set.

### GetHeaders

`func (o *CreateWebhookRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CreateWebhookRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CreateWebhookRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CreateWebhookRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


