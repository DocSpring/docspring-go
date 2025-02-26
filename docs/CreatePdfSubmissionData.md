# CreatePdfSubmissionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]interface{}** |  | 
**DataRequests** | Pointer to [**[]CreateSubmissionDataRequestData**](CreateSubmissionDataRequestData.md) |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**FieldOverrides** | Pointer to **map[string]interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Test** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreatePdfSubmissionData

`func NewCreatePdfSubmissionData(data map[string]interface{}, ) *CreatePdfSubmissionData`

NewCreatePdfSubmissionData instantiates a new CreatePdfSubmissionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePdfSubmissionDataWithDefaults

`func NewCreatePdfSubmissionDataWithDefaults() *CreatePdfSubmissionData`

NewCreatePdfSubmissionDataWithDefaults instantiates a new CreatePdfSubmissionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreatePdfSubmissionData) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreatePdfSubmissionData) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreatePdfSubmissionData) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetDataRequests

`func (o *CreatePdfSubmissionData) GetDataRequests() []CreateSubmissionDataRequestData`

GetDataRequests returns the DataRequests field if non-nil, zero value otherwise.

### GetDataRequestsOk

`func (o *CreatePdfSubmissionData) GetDataRequestsOk() (*[]CreateSubmissionDataRequestData, bool)`

GetDataRequestsOk returns a tuple with the DataRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRequests

`func (o *CreatePdfSubmissionData) SetDataRequests(v []CreateSubmissionDataRequestData)`

SetDataRequests sets DataRequests field to given value.

### HasDataRequests

`func (o *CreatePdfSubmissionData) HasDataRequests() bool`

HasDataRequests returns a boolean if a field has been set.

### GetEditable

`func (o *CreatePdfSubmissionData) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *CreatePdfSubmissionData) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *CreatePdfSubmissionData) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *CreatePdfSubmissionData) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetExpiresIn

`func (o *CreatePdfSubmissionData) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreatePdfSubmissionData) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreatePdfSubmissionData) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreatePdfSubmissionData) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetFieldOverrides

`func (o *CreatePdfSubmissionData) GetFieldOverrides() map[string]interface{}`

GetFieldOverrides returns the FieldOverrides field if non-nil, zero value otherwise.

### GetFieldOverridesOk

`func (o *CreatePdfSubmissionData) GetFieldOverridesOk() (*map[string]interface{}, bool)`

GetFieldOverridesOk returns a tuple with the FieldOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldOverrides

`func (o *CreatePdfSubmissionData) SetFieldOverrides(v map[string]interface{})`

SetFieldOverrides sets FieldOverrides field to given value.

### HasFieldOverrides

`func (o *CreatePdfSubmissionData) HasFieldOverrides() bool`

HasFieldOverrides returns a boolean if a field has been set.

### GetMetadata

`func (o *CreatePdfSubmissionData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePdfSubmissionData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePdfSubmissionData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreatePdfSubmissionData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPassword

`func (o *CreatePdfSubmissionData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreatePdfSubmissionData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreatePdfSubmissionData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreatePdfSubmissionData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTest

`func (o *CreatePdfSubmissionData) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CreatePdfSubmissionData) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CreatePdfSubmissionData) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CreatePdfSubmissionData) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


