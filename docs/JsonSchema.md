# JsonSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Definitions** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**AdditionalPropertiesField** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJsonSchema

`func NewJsonSchema() *JsonSchema`

NewJsonSchema instantiates a new JsonSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonSchemaWithDefaults

`func NewJsonSchemaWithDefaults() *JsonSchema`

NewJsonSchemaWithDefaults instantiates a new JsonSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *JsonSchema) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *JsonSchema) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *JsonSchema) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *JsonSchema) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetId

`func (o *JsonSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JsonSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *JsonSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *JsonSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *JsonSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *JsonSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *JsonSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JsonSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JsonSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JsonSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefinitions

`func (o *JsonSchema) GetDefinitions() map[string]interface{}`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *JsonSchema) GetDefinitionsOk() (*map[string]interface{}, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *JsonSchema) SetDefinitions(v map[string]interface{})`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *JsonSchema) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.

### GetType

`func (o *JsonSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JsonSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JsonSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JsonSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProperties

`func (o *JsonSchema) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *JsonSchema) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *JsonSchema) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *JsonSchema) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetAdditionalPropertiesField

`func (o *JsonSchema) GetAdditionalPropertiesField() bool`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *JsonSchema) GetAdditionalPropertiesFieldOk() (*bool, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *JsonSchema) SetAdditionalPropertiesField(v bool)`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.

### HasAdditionalPropertiesField

`func (o *JsonSchema) HasAdditionalPropertiesField() bool`

HasAdditionalPropertiesField returns a boolean if a field has been set.

### GetRequired

`func (o *JsonSchema) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *JsonSchema) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *JsonSchema) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *JsonSchema) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


