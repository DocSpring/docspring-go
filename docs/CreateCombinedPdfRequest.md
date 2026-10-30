# CreateCombinedPdfRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcePdfs** | [**[]CombinedPdfSourceInput**](CombinedPdfSourceInput.md) | The PDFs to merge, in order. | 
**Metadata** | Pointer to **map[string]interface{}** | Your own data about the combined PDF, such as an ID from your app. Its `pdf_filename` key names the PDF's file (`files.pdf.filename`, and its download): a `.pdf` ending is removed, each run of characters other than letters, digits, `-`, `_` and `.` becomes one `_`, and the first 128 characters are kept, then `.pdf` is added. Without it, the file is named after the combined PDF's ID (`com_1234567890abcdef01.pdf`). | [optional] 
**Password** | Pointer to **string** | A password that the PDF needs to be opened with. It's never returned; `has_password` says if it's set. | [optional] 
**DeleteCustomFiles** | Pointer to **bool** | Delete the custom files in `source_pdfs` once the combined PDF has been generated (they're kept if it fails). Other combined PDFs can't use them after that. A test API token can't delete custom files. | [optional] 
**ExpiresIn** | Pointer to **int32** | Seconds after the PDF is generated until the combined PDF expires, up to 10 years (315360000). One that fails is kept for at least 48 hours. Without it, the combined PDF is kept until you delete it. | [optional] 

## Methods

### NewCreateCombinedPdfRequest

`func NewCreateCombinedPdfRequest(sourcePdfs []CombinedPdfSourceInput, ) *CreateCombinedPdfRequest`

NewCreateCombinedPdfRequest instantiates a new CreateCombinedPdfRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCombinedPdfRequestWithDefaults

`func NewCreateCombinedPdfRequestWithDefaults() *CreateCombinedPdfRequest`

NewCreateCombinedPdfRequestWithDefaults instantiates a new CreateCombinedPdfRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcePdfs

`func (o *CreateCombinedPdfRequest) GetSourcePdfs() []CombinedPdfSourceInput`

GetSourcePdfs returns the SourcePdfs field if non-nil, zero value otherwise.

### GetSourcePdfsOk

`func (o *CreateCombinedPdfRequest) GetSourcePdfsOk() (*[]CombinedPdfSourceInput, bool)`

GetSourcePdfsOk returns a tuple with the SourcePdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePdfs

`func (o *CreateCombinedPdfRequest) SetSourcePdfs(v []CombinedPdfSourceInput)`

SetSourcePdfs sets SourcePdfs field to given value.


### GetMetadata

`func (o *CreateCombinedPdfRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateCombinedPdfRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateCombinedPdfRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateCombinedPdfRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPassword

`func (o *CreateCombinedPdfRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateCombinedPdfRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateCombinedPdfRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateCombinedPdfRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDeleteCustomFiles

`func (o *CreateCombinedPdfRequest) GetDeleteCustomFiles() bool`

GetDeleteCustomFiles returns the DeleteCustomFiles field if non-nil, zero value otherwise.

### GetDeleteCustomFilesOk

`func (o *CreateCombinedPdfRequest) GetDeleteCustomFilesOk() (*bool, bool)`

GetDeleteCustomFilesOk returns a tuple with the DeleteCustomFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteCustomFiles

`func (o *CreateCombinedPdfRequest) SetDeleteCustomFiles(v bool)`

SetDeleteCustomFiles sets DeleteCustomFiles field to given value.

### HasDeleteCustomFiles

`func (o *CreateCombinedPdfRequest) HasDeleteCustomFiles() bool`

HasDeleteCustomFiles returns a boolean if a field has been set.

### GetExpiresIn

`func (o *CreateCombinedPdfRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreateCombinedPdfRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreateCombinedPdfRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreateCombinedPdfRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


