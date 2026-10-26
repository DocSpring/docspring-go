# GeneratedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**GeneratedFileState**](GeneratedFileState.md) | `processed` once the file has been made. `pending` while it's being made, and `error` if it couldn't be made (see `error_message`). | 
**ErrorMessage** | **NullableString** | Why the file couldn't be made, when `state` is `error`. | 
**DownloadUrl** | **NullableString** | A temporary link to download the file. Null until the file has been made, and when your account turns off download links for submissions. | 
**PermanentUrl** | **NullableString** | A link to download the file that doesn't expire: it redirects to a new temporary link. It needs your API token or a signed-in session, except for a submission of a template that makes its submissions public. Null until the file has been made. | 
**Filename** | **string** | The file's name, such as `sub_123.pdf`. | 
**ContentType** | **string** | The file's media type, such as `application/pdf`. | 
**Size** | **NullableInt32** | The file's size in bytes. Null until the file has been made. | 
**Sha256** | **NullableString** | The SHA-256 hash of the file. Null until the file has been made. | 

## Methods

### NewGeneratedFile

`func NewGeneratedFile(state GeneratedFileState, errorMessage NullableString, downloadUrl NullableString, permanentUrl NullableString, filename string, contentType string, size NullableInt32, sha256 NullableString, ) *GeneratedFile`

NewGeneratedFile instantiates a new GeneratedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratedFileWithDefaults

`func NewGeneratedFileWithDefaults() *GeneratedFile`

NewGeneratedFileWithDefaults instantiates a new GeneratedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *GeneratedFile) GetState() GeneratedFileState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GeneratedFile) GetStateOk() (*GeneratedFileState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GeneratedFile) SetState(v GeneratedFileState)`

SetState sets State field to given value.


### GetErrorMessage

`func (o *GeneratedFile) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GeneratedFile) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GeneratedFile) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *GeneratedFile) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GeneratedFile) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetDownloadUrl

`func (o *GeneratedFile) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *GeneratedFile) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *GeneratedFile) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *GeneratedFile) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *GeneratedFile) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetPermanentUrl

`func (o *GeneratedFile) GetPermanentUrl() string`

GetPermanentUrl returns the PermanentUrl field if non-nil, zero value otherwise.

### GetPermanentUrlOk

`func (o *GeneratedFile) GetPermanentUrlOk() (*string, bool)`

GetPermanentUrlOk returns a tuple with the PermanentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentUrl

`func (o *GeneratedFile) SetPermanentUrl(v string)`

SetPermanentUrl sets PermanentUrl field to given value.


### SetPermanentUrlNil

`func (o *GeneratedFile) SetPermanentUrlNil(b bool)`

 SetPermanentUrlNil sets the value for PermanentUrl to be an explicit nil

### UnsetPermanentUrl
`func (o *GeneratedFile) UnsetPermanentUrl()`

UnsetPermanentUrl ensures that no value is present for PermanentUrl, not even an explicit nil
### GetFilename

`func (o *GeneratedFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *GeneratedFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *GeneratedFile) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *GeneratedFile) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GeneratedFile) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GeneratedFile) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetSize

`func (o *GeneratedFile) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GeneratedFile) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GeneratedFile) SetSize(v int32)`

SetSize sets Size field to given value.


### SetSizeNil

`func (o *GeneratedFile) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *GeneratedFile) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSha256

`func (o *GeneratedFile) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *GeneratedFile) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *GeneratedFile) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### SetSha256Nil

`func (o *GeneratedFile) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *GeneratedFile) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


