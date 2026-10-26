# PageImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** | The page number, from 1. | 
**Width** | **int32** | The image's width in pixels. | 
**Height** | **int32** | The image's height in pixels. | 
**DownloadUrl** | **NullableString** | A temporary link to download the image. A submission's is null when your account turns off download links for submissions. | 
**PermanentUrl** | **NullableString** | A link to download a submission's image that doesn't expire: it redirects to a new temporary link. It needs your API token or a signed-in session, except for a submission of a template that makes its submissions public. Null for a template's page images. | 
**Filename** | **string** | The image's file name, such as `sub_123_page_1.png`. | 
**ContentType** | **string** | The image's media type, `image/png` for a submission's pages or `image/jpeg` for a template's. | 
**Size** | **int32** | The image's size in bytes. | 
**Sha256** | **NullableString** | The SHA-256 hash of the image. Null for a template's page images. | 

## Methods

### NewPageImage

`func NewPageImage(page int32, width int32, height int32, downloadUrl NullableString, permanentUrl NullableString, filename string, contentType string, size int32, sha256 NullableString, ) *PageImage`

NewPageImage instantiates a new PageImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageImageWithDefaults

`func NewPageImageWithDefaults() *PageImage`

NewPageImageWithDefaults instantiates a new PageImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PageImage) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageImage) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageImage) SetPage(v int32)`

SetPage sets Page field to given value.


### GetWidth

`func (o *PageImage) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PageImage) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PageImage) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *PageImage) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PageImage) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PageImage) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetDownloadUrl

`func (o *PageImage) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *PageImage) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *PageImage) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *PageImage) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *PageImage) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetPermanentUrl

`func (o *PageImage) GetPermanentUrl() string`

GetPermanentUrl returns the PermanentUrl field if non-nil, zero value otherwise.

### GetPermanentUrlOk

`func (o *PageImage) GetPermanentUrlOk() (*string, bool)`

GetPermanentUrlOk returns a tuple with the PermanentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentUrl

`func (o *PageImage) SetPermanentUrl(v string)`

SetPermanentUrl sets PermanentUrl field to given value.


### SetPermanentUrlNil

`func (o *PageImage) SetPermanentUrlNil(b bool)`

 SetPermanentUrlNil sets the value for PermanentUrl to be an explicit nil

### UnsetPermanentUrl
`func (o *PageImage) UnsetPermanentUrl()`

UnsetPermanentUrl ensures that no value is present for PermanentUrl, not even an explicit nil
### GetFilename

`func (o *PageImage) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PageImage) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PageImage) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *PageImage) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PageImage) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PageImage) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetSize

`func (o *PageImage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageImage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageImage) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSha256

`func (o *PageImage) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *PageImage) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *PageImage) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### SetSha256Nil

`func (o *PageImage) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *PageImage) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


