# SubmissionBatchData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Submissions** | **[]map[string]interface{}** |  | 
**Test** | Pointer to **bool** |  | [optional] 

## Methods

### NewSubmissionBatchData

`func NewSubmissionBatchData(submissions []map[string]interface{}, ) *SubmissionBatchData`

NewSubmissionBatchData instantiates a new SubmissionBatchData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionBatchDataWithDefaults

`func NewSubmissionBatchDataWithDefaults() *SubmissionBatchData`

NewSubmissionBatchDataWithDefaults instantiates a new SubmissionBatchData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *SubmissionBatchData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubmissionBatchData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubmissionBatchData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubmissionBatchData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSubmissions

`func (o *SubmissionBatchData) GetSubmissions() []map[string]interface{}`

GetSubmissions returns the Submissions field if non-nil, zero value otherwise.

### GetSubmissionsOk

`func (o *SubmissionBatchData) GetSubmissionsOk() (*[]map[string]interface{}, bool)`

GetSubmissionsOk returns a tuple with the Submissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissions

`func (o *SubmissionBatchData) SetSubmissions(v []map[string]interface{})`

SetSubmissions sets Submissions field to given value.


### GetTest

`func (o *SubmissionBatchData) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *SubmissionBatchData) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *SubmissionBatchData) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *SubmissionBatchData) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


