# CreateSubmissionBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Your own data about the batch, such as an ID from your app. | [optional] 
**TemplateId** | Pointer to **string** | The template for each submission that doesn't name its own `template_id`. | [optional] 
**TemplateVersion** | Pointer to **string** | The template version to use: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published). | [optional] 
**Test** | Pointer to **bool** | The default `test` for each submission that doesn't set its own. | [optional] 
**Submissions** | [**[]SubmissionInput**](SubmissionInput.md) | Up to 50 submissions. Each one takes the properties of POST /templates/{template_id}/submissions except `signature_requests`, plus `template_id` and `template_version` when it uses a different template from the batch. | 

## Methods

### NewCreateSubmissionBatchRequest

`func NewCreateSubmissionBatchRequest(submissions []SubmissionInput, ) *CreateSubmissionBatchRequest`

NewCreateSubmissionBatchRequest instantiates a new CreateSubmissionBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubmissionBatchRequestWithDefaults

`func NewCreateSubmissionBatchRequestWithDefaults() *CreateSubmissionBatchRequest`

NewCreateSubmissionBatchRequestWithDefaults instantiates a new CreateSubmissionBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CreateSubmissionBatchRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSubmissionBatchRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSubmissionBatchRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSubmissionBatchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTemplateId

`func (o *CreateSubmissionBatchRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateSubmissionBatchRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CreateSubmissionBatchRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CreateSubmissionBatchRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateVersion

`func (o *CreateSubmissionBatchRequest) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *CreateSubmissionBatchRequest) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *CreateSubmissionBatchRequest) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *CreateSubmissionBatchRequest) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.

### GetTest

`func (o *CreateSubmissionBatchRequest) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CreateSubmissionBatchRequest) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CreateSubmissionBatchRequest) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CreateSubmissionBatchRequest) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetSubmissions

`func (o *CreateSubmissionBatchRequest) GetSubmissions() []SubmissionInput`

GetSubmissions returns the Submissions field if non-nil, zero value otherwise.

### GetSubmissionsOk

`func (o *CreateSubmissionBatchRequest) GetSubmissionsOk() (*[]SubmissionInput, bool)`

GetSubmissionsOk returns a tuple with the Submissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissions

`func (o *CreateSubmissionBatchRequest) SetSubmissions(v []SubmissionInput)`

SetSubmissions sets Submissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


