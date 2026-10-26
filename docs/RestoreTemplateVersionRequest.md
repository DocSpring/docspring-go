# RestoreTemplateVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The published version to restore to the draft (`1.2.0`, or `1` for the latest 1.x.x). | 

## Methods

### NewRestoreTemplateVersionRequest

`func NewRestoreTemplateVersionRequest(version string, ) *RestoreTemplateVersionRequest`

NewRestoreTemplateVersionRequest instantiates a new RestoreTemplateVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreTemplateVersionRequestWithDefaults

`func NewRestoreTemplateVersionRequestWithDefaults() *RestoreTemplateVersionRequest`

NewRestoreTemplateVersionRequestWithDefaults instantiates a new RestoreTemplateVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *RestoreTemplateVersionRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RestoreTemplateVersionRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RestoreTemplateVersionRequest) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


