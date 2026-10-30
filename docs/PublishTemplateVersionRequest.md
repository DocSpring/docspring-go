# PublishTemplateVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionType** | [**TemplateVersionType**](TemplateVersionType.md) | Which part of the latest version number to increase. | 
**Description** | Pointer to **string** | What changed in the version, shown with it in DocSpring. | [optional] 

## Methods

### NewPublishTemplateVersionRequest

`func NewPublishTemplateVersionRequest(versionType TemplateVersionType, ) *PublishTemplateVersionRequest`

NewPublishTemplateVersionRequest instantiates a new PublishTemplateVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishTemplateVersionRequestWithDefaults

`func NewPublishTemplateVersionRequestWithDefaults() *PublishTemplateVersionRequest`

NewPublishTemplateVersionRequestWithDefaults instantiates a new PublishTemplateVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionType

`func (o *PublishTemplateVersionRequest) GetVersionType() TemplateVersionType`

GetVersionType returns the VersionType field if non-nil, zero value otherwise.

### GetVersionTypeOk

`func (o *PublishTemplateVersionRequest) GetVersionTypeOk() (*TemplateVersionType, bool)`

GetVersionTypeOk returns a tuple with the VersionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionType

`func (o *PublishTemplateVersionRequest) SetVersionType(v TemplateVersionType)`

SetVersionType sets VersionType field to given value.


### GetDescription

`func (o *PublishTemplateVersionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublishTemplateVersionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublishTemplateVersionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublishTemplateVersionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


