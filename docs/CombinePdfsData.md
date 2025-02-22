# CombinePdfsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteCustomFiles** | Pointer to **bool** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SourcePdfs** | **[]map[string]interface{}** |  | 
**Test** | Pointer to **bool** |  | [optional] 

## Methods

### NewCombinePdfsData

`func NewCombinePdfsData(sourcePdfs []map[string]interface{}, ) *CombinePdfsData`

NewCombinePdfsData instantiates a new CombinePdfsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinePdfsDataWithDefaults

`func NewCombinePdfsDataWithDefaults() *CombinePdfsData`

NewCombinePdfsDataWithDefaults instantiates a new CombinePdfsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteCustomFiles

`func (o *CombinePdfsData) GetDeleteCustomFiles() bool`

GetDeleteCustomFiles returns the DeleteCustomFiles field if non-nil, zero value otherwise.

### GetDeleteCustomFilesOk

`func (o *CombinePdfsData) GetDeleteCustomFilesOk() (*bool, bool)`

GetDeleteCustomFilesOk returns a tuple with the DeleteCustomFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteCustomFiles

`func (o *CombinePdfsData) SetDeleteCustomFiles(v bool)`

SetDeleteCustomFiles sets DeleteCustomFiles field to given value.

### HasDeleteCustomFiles

`func (o *CombinePdfsData) HasDeleteCustomFiles() bool`

HasDeleteCustomFiles returns a boolean if a field has been set.

### GetExpiresIn

`func (o *CombinePdfsData) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CombinePdfsData) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CombinePdfsData) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CombinePdfsData) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetMetadata

`func (o *CombinePdfsData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CombinePdfsData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CombinePdfsData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CombinePdfsData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPassword

`func (o *CombinePdfsData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CombinePdfsData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CombinePdfsData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CombinePdfsData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSourcePdfs

`func (o *CombinePdfsData) GetSourcePdfs() []map[string]interface{}`

GetSourcePdfs returns the SourcePdfs field if non-nil, zero value otherwise.

### GetSourcePdfsOk

`func (o *CombinePdfsData) GetSourcePdfsOk() (*[]map[string]interface{}, bool)`

GetSourcePdfsOk returns a tuple with the SourcePdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePdfs

`func (o *CombinePdfsData) SetSourcePdfs(v []map[string]interface{})`

SetSourcePdfs sets SourcePdfs field to given value.


### GetTest

`func (o *CombinePdfsData) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CombinePdfsData) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CombinePdfsData) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CombinePdfsData) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


