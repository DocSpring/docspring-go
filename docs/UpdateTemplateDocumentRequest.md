# UpdateTemplateDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheId** | **string** | The `fields.key` of the new PDF, uploaded with the presigned upload (GET /uploads/presign). | 
**Filename** | Pointer to **string** | The PDF's file name, for a `cache_id` (presigned uploads don't send one). | [optional] 
**ReimportFields** | Pointer to **bool** | Replace all of the template's fields with the fillable fields in the new PDF, once DocSpring has processed it (`document.field_reimport_state`). As on a first upload, the default typeface, font size, bold and color are taken from the PDF's fields; the template's other settings are kept. If the PDF has no fillable fields, the template keeps its fields. A locked template can't re-import its fields. | [optional] 

## Methods

### NewUpdateTemplateDocumentRequest

`func NewUpdateTemplateDocumentRequest(cacheId string, ) *UpdateTemplateDocumentRequest`

NewUpdateTemplateDocumentRequest instantiates a new UpdateTemplateDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTemplateDocumentRequestWithDefaults

`func NewUpdateTemplateDocumentRequestWithDefaults() *UpdateTemplateDocumentRequest`

NewUpdateTemplateDocumentRequestWithDefaults instantiates a new UpdateTemplateDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheId

`func (o *UpdateTemplateDocumentRequest) GetCacheId() string`

GetCacheId returns the CacheId field if non-nil, zero value otherwise.

### GetCacheIdOk

`func (o *UpdateTemplateDocumentRequest) GetCacheIdOk() (*string, bool)`

GetCacheIdOk returns a tuple with the CacheId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheId

`func (o *UpdateTemplateDocumentRequest) SetCacheId(v string)`

SetCacheId sets CacheId field to given value.


### GetFilename

`func (o *UpdateTemplateDocumentRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *UpdateTemplateDocumentRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *UpdateTemplateDocumentRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *UpdateTemplateDocumentRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetReimportFields

`func (o *UpdateTemplateDocumentRequest) GetReimportFields() bool`

GetReimportFields returns the ReimportFields field if non-nil, zero value otherwise.

### GetReimportFieldsOk

`func (o *UpdateTemplateDocumentRequest) GetReimportFieldsOk() (*bool, bool)`

GetReimportFieldsOk returns a tuple with the ReimportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReimportFields

`func (o *UpdateTemplateDocumentRequest) SetReimportFields(v bool)`

SetReimportFields sets ReimportFields field to given value.

### HasReimportFields

`func (o *UpdateTemplateDocumentRequest) HasReimportFields() bool`

HasReimportFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


