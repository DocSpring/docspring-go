# CombinedPdfSourceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Must be `url`. | [default to "url"]
**Id** | **string** | The ID of the submission, combined PDF, template or custom file. | 
**TemplateVersion** | Pointer to **string** | For a `template` source, the version to use: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published). | [optional] 
**Processed** | Pointer to **bool** | For a `template` source: true (the default) merges the template's processed PDF, which is the PDF submissions are filled in on (decrypted, with its form fields removed). false merges the PDF as it was uploaded. | [optional] 
**Url** | **string** | For a `url` source, the URL to download the PDF from. Removed once the combined PDF has expired. | 

## Methods

### NewCombinedPdfSourceInput

`func NewCombinedPdfSourceInput(id string, url string, ) *CombinedPdfSourceInput`

NewCombinedPdfSourceInput instantiates a new CombinedPdfSourceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinedPdfSourceInputWithDefaults

`func NewCombinedPdfSourceInputWithDefaults() *CombinedPdfSourceInput`

NewCombinedPdfSourceInputWithDefaults instantiates a new CombinedPdfSourceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CombinedPdfSourceInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CombinedPdfSourceInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CombinedPdfSourceInput) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CombinedPdfSourceInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CombinedPdfSourceInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CombinedPdfSourceInput) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateVersion

`func (o *CombinedPdfSourceInput) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *CombinedPdfSourceInput) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *CombinedPdfSourceInput) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *CombinedPdfSourceInput) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.

### GetProcessed

`func (o *CombinedPdfSourceInput) GetProcessed() bool`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *CombinedPdfSourceInput) GetProcessedOk() (*bool, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *CombinedPdfSourceInput) SetProcessed(v bool)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *CombinedPdfSourceInput) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetUrl

`func (o *CombinedPdfSourceInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CombinedPdfSourceInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CombinedPdfSourceInput) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


