# PublishVersionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**VersionType** | **string** |  | 

## Methods

### NewPublishVersionData

`func NewPublishVersionData(versionType string, ) *PublishVersionData`

NewPublishVersionData instantiates a new PublishVersionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishVersionDataWithDefaults

`func NewPublishVersionDataWithDefaults() *PublishVersionData`

NewPublishVersionDataWithDefaults instantiates a new PublishVersionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PublishVersionData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublishVersionData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublishVersionData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublishVersionData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersionType

`func (o *PublishVersionData) GetVersionType() string`

GetVersionType returns the VersionType field if non-nil, zero value otherwise.

### GetVersionTypeOk

`func (o *PublishVersionData) GetVersionTypeOk() (*string, bool)`

GetVersionTypeOk returns a tuple with the VersionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionType

`func (o *PublishVersionData) SetVersionType(v string)`

SetVersionType sets VersionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


