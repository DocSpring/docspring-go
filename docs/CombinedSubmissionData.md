# CombinedSubmissionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SubmissionIds** | **[]string** |  | 
**Test** | Pointer to **bool** |  | [optional] 

## Methods

### NewCombinedSubmissionData

`func NewCombinedSubmissionData(submissionIds []string, ) *CombinedSubmissionData`

NewCombinedSubmissionData instantiates a new CombinedSubmissionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinedSubmissionDataWithDefaults

`func NewCombinedSubmissionDataWithDefaults() *CombinedSubmissionData`

NewCombinedSubmissionDataWithDefaults instantiates a new CombinedSubmissionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *CombinedSubmissionData) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CombinedSubmissionData) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CombinedSubmissionData) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CombinedSubmissionData) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetMetadata

`func (o *CombinedSubmissionData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CombinedSubmissionData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CombinedSubmissionData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CombinedSubmissionData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPassword

`func (o *CombinedSubmissionData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CombinedSubmissionData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CombinedSubmissionData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CombinedSubmissionData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSubmissionIds

`func (o *CombinedSubmissionData) GetSubmissionIds() []string`

GetSubmissionIds returns the SubmissionIds field if non-nil, zero value otherwise.

### GetSubmissionIdsOk

`func (o *CombinedSubmissionData) GetSubmissionIdsOk() (*[]string, bool)`

GetSubmissionIdsOk returns a tuple with the SubmissionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionIds

`func (o *CombinedSubmissionData) SetSubmissionIds(v []string)`

SetSubmissionIds sets SubmissionIds field to given value.


### GetTest

`func (o *CombinedSubmissionData) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CombinedSubmissionData) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CombinedSubmissionData) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CombinedSubmissionData) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


