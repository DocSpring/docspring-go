# UploadPresign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The S3 URL to upload the file to. | 
**Method** | [**UploadMethod**](UploadMethod.md) | The HTTP method for the upload. | 
**Fields** | **map[string]string** | Form fields to send with the file. Send them all, as they are. `key` is the uploaded file's ID: pass it as `cache_id`. | 
**Headers** | **map[string]string** | Headers to send with the upload. | 

## Methods

### NewUploadPresign

`func NewUploadPresign(url string, method UploadMethod, fields map[string]string, headers map[string]string, ) *UploadPresign`

NewUploadPresign instantiates a new UploadPresign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadPresignWithDefaults

`func NewUploadPresignWithDefaults() *UploadPresign`

NewUploadPresignWithDefaults instantiates a new UploadPresign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UploadPresign) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadPresign) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadPresign) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *UploadPresign) GetMethod() UploadMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *UploadPresign) GetMethodOk() (*UploadMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *UploadPresign) SetMethod(v UploadMethod)`

SetMethod sets Method field to given value.


### GetFields

`func (o *UploadPresign) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *UploadPresign) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *UploadPresign) SetFields(v map[string]string)`

SetFields sets Fields field to given value.


### GetHeaders

`func (o *UploadPresign) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *UploadPresign) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *UploadPresign) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


