# SubmissionValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** | True when creating the submission would succeed. | 
**TemplateId** | **string** | The template's ID. | 
**TemplateVersion** | **NullableString** | The published version the request was checked against, such as `1.2.0`. Null for the draft. | 
**Error** | Pointer to [**ErrorInfo**](ErrorInfo.md) |  | [optional] 

## Methods

### NewSubmissionValidation

`func NewSubmissionValidation(valid bool, templateId string, templateVersion NullableString, ) *SubmissionValidation`

NewSubmissionValidation instantiates a new SubmissionValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionValidationWithDefaults

`func NewSubmissionValidationWithDefaults() *SubmissionValidation`

NewSubmissionValidationWithDefaults instantiates a new SubmissionValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *SubmissionValidation) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *SubmissionValidation) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *SubmissionValidation) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetTemplateId

`func (o *SubmissionValidation) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SubmissionValidation) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SubmissionValidation) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetTemplateVersion

`func (o *SubmissionValidation) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *SubmissionValidation) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *SubmissionValidation) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.


### SetTemplateVersionNil

`func (o *SubmissionValidation) SetTemplateVersionNil(b bool)`

 SetTemplateVersionNil sets the value for TemplateVersion to be an explicit nil

### UnsetTemplateVersion
`func (o *SubmissionValidation) UnsetTemplateVersion()`

UnsetTemplateVersion ensures that no value is present for TemplateVersion, not even an explicit nil
### GetError

`func (o *SubmissionValidation) GetError() ErrorInfo`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SubmissionValidation) GetErrorOk() (*ErrorInfo, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SubmissionValidation) SetError(v ErrorInfo)`

SetError sets Error field to given value.

### HasError

`func (o *SubmissionValidation) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


