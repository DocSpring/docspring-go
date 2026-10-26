# TemplateModelCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lang** | [**ModelCodeLanguage**](ModelCodeLanguage.md) | The language of `code`. | 
**RootTypeName** | **string** | The name of the type for the whole submission data object. | 
**Code** | **string** | The source code of the types, starting with `root_type_name`. | 

## Methods

### NewTemplateModelCode

`func NewTemplateModelCode(lang ModelCodeLanguage, rootTypeName string, code string, ) *TemplateModelCode`

NewTemplateModelCode instantiates a new TemplateModelCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateModelCodeWithDefaults

`func NewTemplateModelCodeWithDefaults() *TemplateModelCode`

NewTemplateModelCodeWithDefaults instantiates a new TemplateModelCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLang

`func (o *TemplateModelCode) GetLang() ModelCodeLanguage`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *TemplateModelCode) GetLangOk() (*ModelCodeLanguage, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *TemplateModelCode) SetLang(v ModelCodeLanguage)`

SetLang sets Lang field to given value.


### GetRootTypeName

`func (o *TemplateModelCode) GetRootTypeName() string`

GetRootTypeName returns the RootTypeName field if non-nil, zero value otherwise.

### GetRootTypeNameOk

`func (o *TemplateModelCode) GetRootTypeNameOk() (*string, bool)`

GetRootTypeNameOk returns a tuple with the RootTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootTypeName

`func (o *TemplateModelCode) SetRootTypeName(v string)`

SetRootTypeName sets RootTypeName field to given value.


### GetCode

`func (o *TemplateModelCode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TemplateModelCode) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TemplateModelCode) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


