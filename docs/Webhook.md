# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The webhook's ID. | 
**Name** | **NullableString** | A name for the webhook, shown in the dashboard. | 
**Url** | **string** | The endpoint that receives webhook POST requests. | 
**Enabled** | **bool** | Whether DocSpring sends events to the webhook. Webhooks are disabled after repeated failures. | 
**EventTypes** | [**[]WebhookEventType**](WebhookEventType.md) | The events the webhook subscribes to. A webhook made in the dashboard can also subscribe to API token, account integration and user events, which are listed with their own names. | 
**Mode** | [**NullableWebhookMode**](WebhookMode.md) | Only send events for test or live submissions. Null sends both. | 
**Version** | **int32** | The format of the requests DocSpring sends. Webhooks created with API v2 use 4: each event's `data` is the resource as API v2 returns it. Webhooks made in the dashboard or with API v1 use 3, which has API v1's resources and names (`combined_submission`, `submission_data_request`). Versions 1 and 2 are legacy formats for submission events. | 
**ApiTokenIds** | **[]string** | Only send events from these API tokens. Combined with the other scope lists using OR. Applies to submission, signature request, combined PDF, submission batch and custom file events. | 
**TemplateIds** | **[]string** | Only send events for these templates. Combined with the other scope lists using OR. Applies to submission, signature request and template events. | 
**FolderIds** | **[]string** | Only send events for templates in these folders, including their subfolders. Combined with the other scope lists using OR. Applies to submission, signature request, template and folder events. | 
**IncludeSubmissionData** | **bool** | Include the submission's field data in the payload. When false, the submission's data and truncated_text, and each signature request's data, ip_address and user_agent, are null. | 
**StripFields** | **[]string** | Dot-separated paths to remove from the event's data object. Form fields live under `data` (`data.ssn` removes the ssn field). A bare name removes a top-level key (`metadata` removes the whole metadata object). | 
**Headers** | **map[string]string** | Headers to send with each request. | 
**ConsecutiveFailures** | **int32** | How many deliveries in a row have failed. | 
**LastFailureAt** | **NullableTime** | When a delivery last failed. | 
**DisabledAt** | **NullableTime** | When the webhook was disabled. Null while it is enabled. | 
**DisabledReason** | **NullableString** | Why the webhook was disabled, such as after repeated failures. | 
**Secret** | **NullableString** | The key that signs each request (the X-DocSpring-Signature header), so you can check that a request came from DocSpring. Null for some webhooks created before signing existed. | 
**CreatedAt** | **time.Time** | When the webhook was created. | 
**UpdatedAt** | **time.Time** | When the webhook was last updated. | 

## Methods

### NewWebhook

`func NewWebhook(id string, name NullableString, url string, enabled bool, eventTypes []WebhookEventType, mode NullableWebhookMode, version int32, apiTokenIds []string, templateIds []string, folderIds []string, includeSubmissionData bool, stripFields []string, headers map[string]string, consecutiveFailures int32, lastFailureAt NullableTime, disabledAt NullableTime, disabledReason NullableString, secret NullableString, createdAt time.Time, updatedAt time.Time, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Webhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Webhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Webhook) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Webhook) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Webhook) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabled

`func (o *Webhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Webhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Webhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEventTypes

`func (o *Webhook) GetEventTypes() []WebhookEventType`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *Webhook) GetEventTypesOk() (*[]WebhookEventType, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *Webhook) SetEventTypes(v []WebhookEventType)`

SetEventTypes sets EventTypes field to given value.


### GetMode

`func (o *Webhook) GetMode() WebhookMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Webhook) GetModeOk() (*WebhookMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Webhook) SetMode(v WebhookMode)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *Webhook) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *Webhook) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetVersion

`func (o *Webhook) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Webhook) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Webhook) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetApiTokenIds

`func (o *Webhook) GetApiTokenIds() []string`

GetApiTokenIds returns the ApiTokenIds field if non-nil, zero value otherwise.

### GetApiTokenIdsOk

`func (o *Webhook) GetApiTokenIdsOk() (*[]string, bool)`

GetApiTokenIdsOk returns a tuple with the ApiTokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenIds

`func (o *Webhook) SetApiTokenIds(v []string)`

SetApiTokenIds sets ApiTokenIds field to given value.


### GetTemplateIds

`func (o *Webhook) GetTemplateIds() []string`

GetTemplateIds returns the TemplateIds field if non-nil, zero value otherwise.

### GetTemplateIdsOk

`func (o *Webhook) GetTemplateIdsOk() (*[]string, bool)`

GetTemplateIdsOk returns a tuple with the TemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIds

`func (o *Webhook) SetTemplateIds(v []string)`

SetTemplateIds sets TemplateIds field to given value.


### GetFolderIds

`func (o *Webhook) GetFolderIds() []string`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *Webhook) GetFolderIdsOk() (*[]string, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *Webhook) SetFolderIds(v []string)`

SetFolderIds sets FolderIds field to given value.


### GetIncludeSubmissionData

`func (o *Webhook) GetIncludeSubmissionData() bool`

GetIncludeSubmissionData returns the IncludeSubmissionData field if non-nil, zero value otherwise.

### GetIncludeSubmissionDataOk

`func (o *Webhook) GetIncludeSubmissionDataOk() (*bool, bool)`

GetIncludeSubmissionDataOk returns a tuple with the IncludeSubmissionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubmissionData

`func (o *Webhook) SetIncludeSubmissionData(v bool)`

SetIncludeSubmissionData sets IncludeSubmissionData field to given value.


### GetStripFields

`func (o *Webhook) GetStripFields() []string`

GetStripFields returns the StripFields field if non-nil, zero value otherwise.

### GetStripFieldsOk

`func (o *Webhook) GetStripFieldsOk() (*[]string, bool)`

GetStripFieldsOk returns a tuple with the StripFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripFields

`func (o *Webhook) SetStripFields(v []string)`

SetStripFields sets StripFields field to given value.


### GetHeaders

`func (o *Webhook) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Webhook) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Webhook) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.


### GetConsecutiveFailures

`func (o *Webhook) GetConsecutiveFailures() int32`

GetConsecutiveFailures returns the ConsecutiveFailures field if non-nil, zero value otherwise.

### GetConsecutiveFailuresOk

`func (o *Webhook) GetConsecutiveFailuresOk() (*int32, bool)`

GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveFailures

`func (o *Webhook) SetConsecutiveFailures(v int32)`

SetConsecutiveFailures sets ConsecutiveFailures field to given value.


### GetLastFailureAt

`func (o *Webhook) GetLastFailureAt() time.Time`

GetLastFailureAt returns the LastFailureAt field if non-nil, zero value otherwise.

### GetLastFailureAtOk

`func (o *Webhook) GetLastFailureAtOk() (*time.Time, bool)`

GetLastFailureAtOk returns a tuple with the LastFailureAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureAt

`func (o *Webhook) SetLastFailureAt(v time.Time)`

SetLastFailureAt sets LastFailureAt field to given value.


### SetLastFailureAtNil

`func (o *Webhook) SetLastFailureAtNil(b bool)`

 SetLastFailureAtNil sets the value for LastFailureAt to be an explicit nil

### UnsetLastFailureAt
`func (o *Webhook) UnsetLastFailureAt()`

UnsetLastFailureAt ensures that no value is present for LastFailureAt, not even an explicit nil
### GetDisabledAt

`func (o *Webhook) GetDisabledAt() time.Time`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *Webhook) GetDisabledAtOk() (*time.Time, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *Webhook) SetDisabledAt(v time.Time)`

SetDisabledAt sets DisabledAt field to given value.


### SetDisabledAtNil

`func (o *Webhook) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *Webhook) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetDisabledReason

`func (o *Webhook) GetDisabledReason() string`

GetDisabledReason returns the DisabledReason field if non-nil, zero value otherwise.

### GetDisabledReasonOk

`func (o *Webhook) GetDisabledReasonOk() (*string, bool)`

GetDisabledReasonOk returns a tuple with the DisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReason

`func (o *Webhook) SetDisabledReason(v string)`

SetDisabledReason sets DisabledReason field to given value.


### SetDisabledReasonNil

`func (o *Webhook) SetDisabledReasonNil(b bool)`

 SetDisabledReasonNil sets the value for DisabledReason to be an explicit nil

### UnsetDisabledReason
`func (o *Webhook) UnsetDisabledReason()`

UnsetDisabledReason ensures that no value is present for DisabledReason, not even an explicit nil
### GetSecret

`func (o *Webhook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Webhook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Webhook) SetSecret(v string)`

SetSecret sets Secret field to given value.


### SetSecretNil

`func (o *Webhook) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *Webhook) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetCreatedAt

`func (o *Webhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Webhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Webhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Webhook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Webhook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Webhook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


