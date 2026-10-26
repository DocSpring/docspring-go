# TemplateDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** | The template's ID. | 
**FromVersion** | **NullableString** | The earlier version, such as `1.1.0`. Null when no version was published before `to_version`: it's compared with a new template, so all of its fields are added. | 
**ToVersion** | **NullableString** | The later version, such as `1.2.0`. Null for the draft. | 
**Summary** | [**TemplateDiffSummary**](TemplateDiffSummary.md) | How many changes of each kind there are. | 
**Fields** | [**TemplateFieldChanges**](TemplateFieldChanges.md) | The fields that changed. | 
**ChangedSettings** | **[]string** | The names of the template's settings that changed, as they're named on the Template (`name`, `html`, `defaults`, `expire_after`, ...). Values aren't included, and secrets like `encrypt_pdfs_password` are named when they change. | 
**DocumentChanged** | **bool** | The template's PDF was replaced (or added, when there's no `from_version`). | 

## Methods

### NewTemplateDiff

`func NewTemplateDiff(templateId string, fromVersion NullableString, toVersion NullableString, summary TemplateDiffSummary, fields TemplateFieldChanges, changedSettings []string, documentChanged bool, ) *TemplateDiff`

NewTemplateDiff instantiates a new TemplateDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDiffWithDefaults

`func NewTemplateDiffWithDefaults() *TemplateDiff`

NewTemplateDiffWithDefaults instantiates a new TemplateDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *TemplateDiff) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateDiff) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateDiff) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetFromVersion

`func (o *TemplateDiff) GetFromVersion() string`

GetFromVersion returns the FromVersion field if non-nil, zero value otherwise.

### GetFromVersionOk

`func (o *TemplateDiff) GetFromVersionOk() (*string, bool)`

GetFromVersionOk returns a tuple with the FromVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromVersion

`func (o *TemplateDiff) SetFromVersion(v string)`

SetFromVersion sets FromVersion field to given value.


### SetFromVersionNil

`func (o *TemplateDiff) SetFromVersionNil(b bool)`

 SetFromVersionNil sets the value for FromVersion to be an explicit nil

### UnsetFromVersion
`func (o *TemplateDiff) UnsetFromVersion()`

UnsetFromVersion ensures that no value is present for FromVersion, not even an explicit nil
### GetToVersion

`func (o *TemplateDiff) GetToVersion() string`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *TemplateDiff) GetToVersionOk() (*string, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *TemplateDiff) SetToVersion(v string)`

SetToVersion sets ToVersion field to given value.


### SetToVersionNil

`func (o *TemplateDiff) SetToVersionNil(b bool)`

 SetToVersionNil sets the value for ToVersion to be an explicit nil

### UnsetToVersion
`func (o *TemplateDiff) UnsetToVersion()`

UnsetToVersion ensures that no value is present for ToVersion, not even an explicit nil
### GetSummary

`func (o *TemplateDiff) GetSummary() TemplateDiffSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TemplateDiff) GetSummaryOk() (*TemplateDiffSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TemplateDiff) SetSummary(v TemplateDiffSummary)`

SetSummary sets Summary field to given value.


### GetFields

`func (o *TemplateDiff) GetFields() TemplateFieldChanges`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TemplateDiff) GetFieldsOk() (*TemplateFieldChanges, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TemplateDiff) SetFields(v TemplateFieldChanges)`

SetFields sets Fields field to given value.


### GetChangedSettings

`func (o *TemplateDiff) GetChangedSettings() []string`

GetChangedSettings returns the ChangedSettings field if non-nil, zero value otherwise.

### GetChangedSettingsOk

`func (o *TemplateDiff) GetChangedSettingsOk() (*[]string, bool)`

GetChangedSettingsOk returns a tuple with the ChangedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedSettings

`func (o *TemplateDiff) SetChangedSettings(v []string)`

SetChangedSettings sets ChangedSettings field to given value.


### GetDocumentChanged

`func (o *TemplateDiff) GetDocumentChanged() bool`

GetDocumentChanged returns the DocumentChanged field if non-nil, zero value otherwise.

### GetDocumentChangedOk

`func (o *TemplateDiff) GetDocumentChangedOk() (*bool, bool)`

GetDocumentChangedOk returns a tuple with the DocumentChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentChanged

`func (o *TemplateDiff) SetDocumentChanged(v bool)`

SetDocumentChanged sets DocumentChanged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


