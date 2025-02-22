# CopyTemplateOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ParentFolderId** | **string** |  | 

## Methods

### NewCopyTemplateOptions

`func NewCopyTemplateOptions(parentFolderId string, ) *CopyTemplateOptions`

NewCopyTemplateOptions instantiates a new CopyTemplateOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyTemplateOptionsWithDefaults

`func NewCopyTemplateOptionsWithDefaults() *CopyTemplateOptions`

NewCopyTemplateOptionsWithDefaults instantiates a new CopyTemplateOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CopyTemplateOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CopyTemplateOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CopyTemplateOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CopyTemplateOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentFolderId

`func (o *CopyTemplateOptions) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *CopyTemplateOptions) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *CopyTemplateOptions) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


