# CopyTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Default: the template's name followed by a number, such as "Contract (2)". For a published version, the name followed by the version, such as "Contract (v1.2.0)". | [optional] 
**ParentFolderId** | Pointer to **NullableString** | The folder to put the copy in. Null puts it at the root. Left out, the copy goes in the template's own folder. | [optional] 
**Version** | Pointer to **string** | Copy this published version (`1.2.0`, or `1` for the latest 1.x.x) instead of the draft. | [optional] 

## Methods

### NewCopyTemplateRequest

`func NewCopyTemplateRequest() *CopyTemplateRequest`

NewCopyTemplateRequest instantiates a new CopyTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyTemplateRequestWithDefaults

`func NewCopyTemplateRequestWithDefaults() *CopyTemplateRequest`

NewCopyTemplateRequestWithDefaults instantiates a new CopyTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CopyTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CopyTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CopyTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CopyTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentFolderId

`func (o *CopyTemplateRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *CopyTemplateRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *CopyTemplateRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *CopyTemplateRequest) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *CopyTemplateRequest) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *CopyTemplateRequest) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetVersion

`func (o *CopyTemplateRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CopyTemplateRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CopyTemplateRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CopyTemplateRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


