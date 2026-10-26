# UpdateTemplateDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheId** | Pointer to **string** | The `fields.key` of the new PDF, uploaded with the presigned upload (GET /uploads/presign). | [optional] 
**DocumentUrl** | Pointer to **string** | An https:// link to the template's PDF (up to 100 MB), instead of `cache_id`. DocSpring downloads it after responding, following redirects to other https:// links: the template's `document.state` is `downloading` until then (sync.api hosts wait for it). If it can't be downloaded, the state is `download_failed` and `document.error_message` says why. | [optional] 
**DocumentBase64** | Pointer to **string** | The template's PDF, encoded as base64 (a `data:` URI works too), instead of `cache_id`. Up to 7 MB before it's encoded: send a larger file with `cache_id` or `document_url`. | [optional] 
**Filename** | Pointer to **string** | The PDF's file name, for a `cache_id` (presigned uploads don't send one) or `document_base64`. A `document_url`'s is the last part of its path unless you set it. | [optional] 
**ReimportFields** | Pointer to **bool** | Replace all of the template's fields with the fillable fields in the new PDF, once DocSpring has processed it (`document.field_reimport_state`). As on a first upload, the default typeface, font size, bold and color are taken from the PDF's fields; the template's other settings are kept. If the PDF has no fillable fields, the template keeps its fields. A locked template can't re-import its fields. | [optional] 

## Methods

### NewUpdateTemplateDocumentRequest

`func NewUpdateTemplateDocumentRequest() *UpdateTemplateDocumentRequest`

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

### HasCacheId

`func (o *UpdateTemplateDocumentRequest) HasCacheId() bool`

HasCacheId returns a boolean if a field has been set.

### GetDocumentUrl

`func (o *UpdateTemplateDocumentRequest) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *UpdateTemplateDocumentRequest) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *UpdateTemplateDocumentRequest) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.

### HasDocumentUrl

`func (o *UpdateTemplateDocumentRequest) HasDocumentUrl() bool`

HasDocumentUrl returns a boolean if a field has been set.

### GetDocumentBase64

`func (o *UpdateTemplateDocumentRequest) GetDocumentBase64() string`

GetDocumentBase64 returns the DocumentBase64 field if non-nil, zero value otherwise.

### GetDocumentBase64Ok

`func (o *UpdateTemplateDocumentRequest) GetDocumentBase64Ok() (*string, bool)`

GetDocumentBase64Ok returns a tuple with the DocumentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentBase64

`func (o *UpdateTemplateDocumentRequest) SetDocumentBase64(v string)`

SetDocumentBase64 sets DocumentBase64 field to given value.

### HasDocumentBase64

`func (o *UpdateTemplateDocumentRequest) HasDocumentBase64() bool`

HasDocumentBase64 returns a boolean if a field has been set.

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


