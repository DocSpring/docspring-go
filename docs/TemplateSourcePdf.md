# TemplateSourcePdf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Must be `template`. | [default to "template"]
**Id** | **string** | The ID of the submission, combined PDF, template or custom file. | 
**TemplateVersion** | Pointer to **string** | For a `template` source, the version to use: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published). | [optional] 
**Processed** | Pointer to **bool** | For a `template` source: true (the default) merges the template's processed PDF, which is the PDF submissions are filled in on (decrypted, with its form fields removed). false merges the PDF as it was uploaded. | [optional] 

## Methods

### NewTemplateSourcePdf

`func NewTemplateSourcePdf(id string, ) *TemplateSourcePdf`

NewTemplateSourcePdf instantiates a new TemplateSourcePdf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateSourcePdfWithDefaults

`func NewTemplateSourcePdfWithDefaults() *TemplateSourcePdf`

NewTemplateSourcePdfWithDefaults instantiates a new TemplateSourcePdf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateSourcePdf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateSourcePdf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateSourcePdf) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *TemplateSourcePdf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateSourcePdf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateSourcePdf) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateVersion

`func (o *TemplateSourcePdf) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *TemplateSourcePdf) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *TemplateSourcePdf) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *TemplateSourcePdf) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.

### GetProcessed

`func (o *TemplateSourcePdf) GetProcessed() bool`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *TemplateSourcePdf) GetProcessedOk() (*bool, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *TemplateSourcePdf) SetProcessed(v bool)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *TemplateSourcePdf) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


