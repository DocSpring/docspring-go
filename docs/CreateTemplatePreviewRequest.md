# CreateTemplatePreviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]interface{}** | The values of the template's fields, by field name, as you'd send them to create a submission. They're rendered as they are: they aren't checked against the template's fields. A value that a field can't print, such as an image URL that doesn't download, makes the preview's `state` `error`, with an `error_message`. | 

## Methods

### NewCreateTemplatePreviewRequest

`func NewCreateTemplatePreviewRequest(data map[string]interface{}, ) *CreateTemplatePreviewRequest`

NewCreateTemplatePreviewRequest instantiates a new CreateTemplatePreviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTemplatePreviewRequestWithDefaults

`func NewCreateTemplatePreviewRequestWithDefaults() *CreateTemplatePreviewRequest`

NewCreateTemplatePreviewRequestWithDefaults instantiates a new CreateTemplatePreviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateTemplatePreviewRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTemplatePreviewRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTemplatePreviewRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


