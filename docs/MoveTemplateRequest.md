# MoveTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentFolderId** | **NullableString** | The folder to move the template and all of its published versions to. Null moves them to the root. | 

## Methods

### NewMoveTemplateRequest

`func NewMoveTemplateRequest(parentFolderId NullableString, ) *MoveTemplateRequest`

NewMoveTemplateRequest instantiates a new MoveTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveTemplateRequestWithDefaults

`func NewMoveTemplateRequestWithDefaults() *MoveTemplateRequest`

NewMoveTemplateRequestWithDefaults instantiates a new MoveTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentFolderId

`func (o *MoveTemplateRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MoveTemplateRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *MoveTemplateRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.


### SetParentFolderIdNil

`func (o *MoveTemplateRequest) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *MoveTemplateRequest) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


