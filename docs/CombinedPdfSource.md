# CombinedPdfSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CombinedPdfSourceType**](CombinedPdfSourceType.md) | What the PDF is: a submission's PDF, another combined PDF, a template's PDF, a custom file (POST /custom_files) or a PDF to download from a URL. | 
**Id** | Pointer to **string** | The ID of the submission, combined PDF, template or custom file. | [optional] 
**Url** | Pointer to **string** | For a `url` source, the URL to download the PDF from. Removed once the combined PDF has expired. | [optional] 
**TemplateVersion** | Pointer to **string** | For a `template` source, the version to use: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published). | [optional] 
**Processed** | Pointer to **bool** | For a `template` source: true (the default) merges the template's processed PDF, which is the PDF submissions are filled in on (decrypted, with its form fields removed). false merges the PDF as it was uploaded. | [optional] 

## Methods

### NewCombinedPdfSource

`func NewCombinedPdfSource(type_ CombinedPdfSourceType, ) *CombinedPdfSource`

NewCombinedPdfSource instantiates a new CombinedPdfSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinedPdfSourceWithDefaults

`func NewCombinedPdfSourceWithDefaults() *CombinedPdfSource`

NewCombinedPdfSourceWithDefaults instantiates a new CombinedPdfSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CombinedPdfSource) GetType() CombinedPdfSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CombinedPdfSource) GetTypeOk() (*CombinedPdfSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CombinedPdfSource) SetType(v CombinedPdfSourceType)`

SetType sets Type field to given value.


### GetId

`func (o *CombinedPdfSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CombinedPdfSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CombinedPdfSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CombinedPdfSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *CombinedPdfSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CombinedPdfSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CombinedPdfSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CombinedPdfSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTemplateVersion

`func (o *CombinedPdfSource) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *CombinedPdfSource) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *CombinedPdfSource) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *CombinedPdfSource) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.

### GetProcessed

`func (o *CombinedPdfSource) GetProcessed() bool`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *CombinedPdfSource) GetProcessedOk() (*bool, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *CombinedPdfSource) SetProcessed(v bool)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *CombinedPdfSource) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


