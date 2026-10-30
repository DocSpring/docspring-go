# RenameFolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The folder's name. It can't contain `/`, and is unique within its parent folder. | 

## Methods

### NewRenameFolderRequest

`func NewRenameFolderRequest(name string, ) *RenameFolderRequest`

NewRenameFolderRequest instantiates a new RenameFolderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameFolderRequestWithDefaults

`func NewRenameFolderRequestWithDefaults() *RenameFolderRequest`

NewRenameFolderRequestWithDefaults instantiates a new RenameFolderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RenameFolderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RenameFolderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RenameFolderRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


