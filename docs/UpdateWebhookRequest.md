# UpdateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | A name for the webhook, shown in the dashboard. | [optional] 
**Url** | Pointer to **string** | The endpoint that receives webhook POST requests. | [optional] 
**Enabled** | Pointer to **bool** | Whether DocSpring sends events to the webhook. Webhooks are disabled after repeated failures. | [optional] 
**EventTypes** | Pointer to [**[]WebhookEventType**](WebhookEventType.md) | The events the webhook subscribes to. A webhook made in the dashboard can also subscribe to API token, account integration and user events, which are listed with their own names. | [optional] 
**Mode** | Pointer to [**NullableWebhookMode**](WebhookMode.md) | Only send events for test or live submissions. Null sends both. | [optional] 
**ApiTokenIds** | Pointer to **[]string** | Only send events from these API tokens. Combined with the other scope lists using OR. Applies to submission, signature request, combined PDF, submission batch and custom file events. | [optional] 
**TemplateIds** | Pointer to **[]string** | Only send events for these templates. Combined with the other scope lists using OR. Applies to submission, signature request and template events. | [optional] 
**FolderIds** | Pointer to **[]string** | Only send events for templates in these folders, including their subfolders. Combined with the other scope lists using OR. Applies to submission, signature request, template and folder events. | [optional] 
**IncludeSubmissionData** | Pointer to **bool** | Include the submission's field data in the payload. When false, the submission's data and truncated_text, and each signature request's data, ip_address and user_agent, are null. | [optional] 
**StripFields** | Pointer to **[]string** | Dot-separated paths to remove from the event's data object. Form fields live under `data` (`data.ssn` removes the ssn field). A bare name removes a top-level key (`metadata` removes the whole metadata object). | [optional] 
**Headers** | Pointer to **map[string]string** | Headers to send with each request. | [optional] 

## Methods

### NewUpdateWebhookRequest

`func NewUpdateWebhookRequest() *UpdateWebhookRequest`

NewUpdateWebhookRequest instantiates a new UpdateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookRequestWithDefaults

`func NewUpdateWebhookRequestWithDefaults() *UpdateWebhookRequest`

NewUpdateWebhookRequestWithDefaults instantiates a new UpdateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateWebhookRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateWebhookRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *UpdateWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateWebhookRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateWebhookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateWebhookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateWebhookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateWebhookRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypes

`func (o *UpdateWebhookRequest) GetEventTypes() []WebhookEventType`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *UpdateWebhookRequest) GetEventTypesOk() (*[]WebhookEventType, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *UpdateWebhookRequest) SetEventTypes(v []WebhookEventType)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *UpdateWebhookRequest) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetMode

`func (o *UpdateWebhookRequest) GetMode() WebhookMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateWebhookRequest) GetModeOk() (*WebhookMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateWebhookRequest) SetMode(v WebhookMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpdateWebhookRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *UpdateWebhookRequest) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *UpdateWebhookRequest) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetApiTokenIds

`func (o *UpdateWebhookRequest) GetApiTokenIds() []string`

GetApiTokenIds returns the ApiTokenIds field if non-nil, zero value otherwise.

### GetApiTokenIdsOk

`func (o *UpdateWebhookRequest) GetApiTokenIdsOk() (*[]string, bool)`

GetApiTokenIdsOk returns a tuple with the ApiTokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenIds

`func (o *UpdateWebhookRequest) SetApiTokenIds(v []string)`

SetApiTokenIds sets ApiTokenIds field to given value.

### HasApiTokenIds

`func (o *UpdateWebhookRequest) HasApiTokenIds() bool`

HasApiTokenIds returns a boolean if a field has been set.

### GetTemplateIds

`func (o *UpdateWebhookRequest) GetTemplateIds() []string`

GetTemplateIds returns the TemplateIds field if non-nil, zero value otherwise.

### GetTemplateIdsOk

`func (o *UpdateWebhookRequest) GetTemplateIdsOk() (*[]string, bool)`

GetTemplateIdsOk returns a tuple with the TemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIds

`func (o *UpdateWebhookRequest) SetTemplateIds(v []string)`

SetTemplateIds sets TemplateIds field to given value.

### HasTemplateIds

`func (o *UpdateWebhookRequest) HasTemplateIds() bool`

HasTemplateIds returns a boolean if a field has been set.

### GetFolderIds

`func (o *UpdateWebhookRequest) GetFolderIds() []string`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *UpdateWebhookRequest) GetFolderIdsOk() (*[]string, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *UpdateWebhookRequest) SetFolderIds(v []string)`

SetFolderIds sets FolderIds field to given value.

### HasFolderIds

`func (o *UpdateWebhookRequest) HasFolderIds() bool`

HasFolderIds returns a boolean if a field has been set.

### GetIncludeSubmissionData

`func (o *UpdateWebhookRequest) GetIncludeSubmissionData() bool`

GetIncludeSubmissionData returns the IncludeSubmissionData field if non-nil, zero value otherwise.

### GetIncludeSubmissionDataOk

`func (o *UpdateWebhookRequest) GetIncludeSubmissionDataOk() (*bool, bool)`

GetIncludeSubmissionDataOk returns a tuple with the IncludeSubmissionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubmissionData

`func (o *UpdateWebhookRequest) SetIncludeSubmissionData(v bool)`

SetIncludeSubmissionData sets IncludeSubmissionData field to given value.

### HasIncludeSubmissionData

`func (o *UpdateWebhookRequest) HasIncludeSubmissionData() bool`

HasIncludeSubmissionData returns a boolean if a field has been set.

### GetStripFields

`func (o *UpdateWebhookRequest) GetStripFields() []string`

GetStripFields returns the StripFields field if non-nil, zero value otherwise.

### GetStripFieldsOk

`func (o *UpdateWebhookRequest) GetStripFieldsOk() (*[]string, bool)`

GetStripFieldsOk returns a tuple with the StripFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripFields

`func (o *UpdateWebhookRequest) SetStripFields(v []string)`

SetStripFields sets StripFields field to given value.

### HasStripFields

`func (o *UpdateWebhookRequest) HasStripFields() bool`

HasStripFields returns a boolean if a field has been set.

### GetHeaders

`func (o *UpdateWebhookRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *UpdateWebhookRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *UpdateWebhookRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *UpdateWebhookRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


