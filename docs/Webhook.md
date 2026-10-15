# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **NullableString** |  | 
**Url** | **NullableString** |  | 
**Enabled** | **bool** |  | 
**EventTypes** | **[]string** |  | 
**Mode** | **NullableString** |  | 
**Version** | **int32** |  | 
**ApiTokenIds** | **[]string** |  | 
**TemplateIds** | **[]string** |  | 
**FolderIds** | **[]string** |  | 
**IncludeSubmissionData** | **bool** |  | 
**StripFields** | **[]string** |  | 
**Headers** | **map[string]string** |  | 
**ConsecutiveFailures** | **int32** |  | 
**LastFailureAt** | **NullableString** |  | 
**DisabledAt** | **NullableString** |  | 
**DisabledReason** | **NullableString** |  | 
**Secret** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 

## Methods

### NewWebhook

`func NewWebhook(id string, name NullableString, url NullableString, enabled bool, eventTypes []string, mode NullableString, version int32, apiTokenIds []string, templateIds []string, folderIds []string, includeSubmissionData bool, stripFields []string, headers map[string]string, consecutiveFailures int32, lastFailureAt NullableString, disabledAt NullableString, disabledReason NullableString, secret NullableString, createdAt NullableString, updatedAt NullableString, ) *Webhook`

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


### SetUrlNil

`func (o *Webhook) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Webhook) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
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

`func (o *Webhook) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *Webhook) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *Webhook) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetMode

`func (o *Webhook) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Webhook) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Webhook) SetMode(v string)`

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

`func (o *Webhook) GetLastFailureAt() string`

GetLastFailureAt returns the LastFailureAt field if non-nil, zero value otherwise.

### GetLastFailureAtOk

`func (o *Webhook) GetLastFailureAtOk() (*string, bool)`

GetLastFailureAtOk returns a tuple with the LastFailureAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureAt

`func (o *Webhook) SetLastFailureAt(v string)`

SetLastFailureAt sets LastFailureAt field to given value.


### SetLastFailureAtNil

`func (o *Webhook) SetLastFailureAtNil(b bool)`

 SetLastFailureAtNil sets the value for LastFailureAt to be an explicit nil

### UnsetLastFailureAt
`func (o *Webhook) UnsetLastFailureAt()`

UnsetLastFailureAt ensures that no value is present for LastFailureAt, not even an explicit nil
### GetDisabledAt

`func (o *Webhook) GetDisabledAt() string`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *Webhook) GetDisabledAtOk() (*string, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *Webhook) SetDisabledAt(v string)`

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

`func (o *Webhook) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Webhook) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Webhook) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Webhook) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Webhook) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Webhook) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Webhook) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Webhook) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Webhook) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Webhook) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


