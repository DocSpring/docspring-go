# MoveTemplateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentFolderId** | **NullableString** |  | 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewMoveTemplateData

`func NewMoveTemplateData(parentFolderId NullableString, ) *MoveTemplateData`

NewMoveTemplateData instantiates a new MoveTemplateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveTemplateDataWithDefaults

`func NewMoveTemplateDataWithDefaults() *MoveTemplateData`

NewMoveTemplateDataWithDefaults instantiates a new MoveTemplateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentFolderId

`func (o *MoveTemplateData) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MoveTemplateData) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *MoveTemplateData) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.


### SetParentFolderIdNil

`func (o *MoveTemplateData) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *MoveTemplateData) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetVersion

`func (o *MoveTemplateData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MoveTemplateData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MoveTemplateData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MoveTemplateData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


