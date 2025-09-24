# TemplateDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Errors** | Pointer to **[]string** |  | [optional] 
**LatestVersion** | Pointer to **NullableString** |  | [optional] 
**Versions** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewTemplateDeleteResponse

`func NewTemplateDeleteResponse(status string, ) *TemplateDeleteResponse`

NewTemplateDeleteResponse instantiates a new TemplateDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDeleteResponseWithDefaults

`func NewTemplateDeleteResponseWithDefaults() *TemplateDeleteResponse`

NewTemplateDeleteResponseWithDefaults instantiates a new TemplateDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TemplateDeleteResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateDeleteResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateDeleteResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrors

`func (o *TemplateDeleteResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TemplateDeleteResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TemplateDeleteResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *TemplateDeleteResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLatestVersion

`func (o *TemplateDeleteResponse) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *TemplateDeleteResponse) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *TemplateDeleteResponse) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *TemplateDeleteResponse) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### SetLatestVersionNil

`func (o *TemplateDeleteResponse) SetLatestVersionNil(b bool)`

 SetLatestVersionNil sets the value for LatestVersion to be an explicit nil

### UnsetLatestVersion
`func (o *TemplateDeleteResponse) UnsetLatestVersion()`

UnsetLatestVersion ensures that no value is present for LatestVersion, not even an explicit nil
### GetVersions

`func (o *TemplateDeleteResponse) GetVersions() []map[string]interface{}`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *TemplateDeleteResponse) GetVersionsOk() (*[]map[string]interface{}, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *TemplateDeleteResponse) SetVersions(v []map[string]interface{})`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *TemplateDeleteResponse) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


