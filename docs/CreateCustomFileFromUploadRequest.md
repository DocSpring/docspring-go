# CreateCustomFileFromUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheId** | **string** | The `fields.key` from GET /uploads/presign, after you upload the file to its URL. | 

## Methods

### NewCreateCustomFileFromUploadRequest

`func NewCreateCustomFileFromUploadRequest(cacheId string, ) *CreateCustomFileFromUploadRequest`

NewCreateCustomFileFromUploadRequest instantiates a new CreateCustomFileFromUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomFileFromUploadRequestWithDefaults

`func NewCreateCustomFileFromUploadRequestWithDefaults() *CreateCustomFileFromUploadRequest`

NewCreateCustomFileFromUploadRequestWithDefaults instantiates a new CreateCustomFileFromUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheId

`func (o *CreateCustomFileFromUploadRequest) GetCacheId() string`

GetCacheId returns the CacheId field if non-nil, zero value otherwise.

### GetCacheIdOk

`func (o *CreateCustomFileFromUploadRequest) GetCacheIdOk() (*string, bool)`

GetCacheIdOk returns a tuple with the CacheId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheId

`func (o *CreateCustomFileFromUploadRequest) SetCacheId(v string)`

SetCacheId sets CacheId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


