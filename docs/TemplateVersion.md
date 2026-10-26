# TemplateVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The version number, such as `1.2.0`. | 
**Description** | **NullableString** | What changed in the version, as written when it was published. | 
**PublishedAt** | **time.Time** | When the version was published. | 
**PublishedBy** | [**NullableTemplateActor**](TemplateActor.md) | Who published the version. Null when that wasn't recorded. | 

## Methods

### NewTemplateVersion

`func NewTemplateVersion(version string, description NullableString, publishedAt time.Time, publishedBy NullableTemplateActor, ) *TemplateVersion`

NewTemplateVersion instantiates a new TemplateVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateVersionWithDefaults

`func NewTemplateVersionWithDefaults() *TemplateVersion`

NewTemplateVersionWithDefaults instantiates a new TemplateVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *TemplateVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TemplateVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TemplateVersion) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *TemplateVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateVersion) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *TemplateVersion) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateVersion) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPublishedAt

`func (o *TemplateVersion) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *TemplateVersion) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *TemplateVersion) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### GetPublishedBy

`func (o *TemplateVersion) GetPublishedBy() TemplateActor`

GetPublishedBy returns the PublishedBy field if non-nil, zero value otherwise.

### GetPublishedByOk

`func (o *TemplateVersion) GetPublishedByOk() (*TemplateActor, bool)`

GetPublishedByOk returns a tuple with the PublishedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedBy

`func (o *TemplateVersion) SetPublishedBy(v TemplateActor)`

SetPublishedBy sets PublishedBy field to given value.


### SetPublishedByNil

`func (o *TemplateVersion) SetPublishedByNil(b bool)`

 SetPublishedByNil sets the value for PublishedBy to be an explicit nil

### UnsetPublishedBy
`func (o *TemplateVersion) UnsetPublishedBy()`

UnsetPublishedBy ensures that no value is present for PublishedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


