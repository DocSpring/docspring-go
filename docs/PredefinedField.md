# PredefinedField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The field's name, the key of its value in submission data. A `/` nests the key in objects and arrays (`customer/email`, `signers/0/name`). | 
**DisplayName** | Pointer to **NullableString** | The name the template editor shows instead of `name`, where a `/` puts it in a folder, such as `Defendant/Name`. Null shows `name`. Two names can't share a display name. | [optional] 
**Title** | Pointer to **NullableString** | The Title setting a field gets when it's given this name. Null leaves the field's title as it is. | [optional] 
**Type** | [**FieldType**](FieldType.md) | The data type a field gets when it's given this name, as a field's `type`. | 
**OptionList** | Pointer to **NullableString** | String fields only: the values allowed, separated by commas, such as `Yes,No`. Null allows any value. | [optional] 
**RequiredInTemplate** | Pointer to **bool** | Whether the template must have a field with this name. The embedded editor shows the names that are missing, and with `disableDoneButtonIfPredefinedFieldsMissing` keeps its Done button disabled until the template has them. Default false. | [optional] 
**FieldRequired** | Pointer to **bool** | The Required setting a field gets when it's given this name. Default true. | [optional] 

## Methods

### NewPredefinedField

`func NewPredefinedField(name string, type_ FieldType, ) *PredefinedField`

NewPredefinedField instantiates a new PredefinedField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredefinedFieldWithDefaults

`func NewPredefinedFieldWithDefaults() *PredefinedField`

NewPredefinedFieldWithDefaults instantiates a new PredefinedField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PredefinedField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PredefinedField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PredefinedField) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *PredefinedField) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PredefinedField) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PredefinedField) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PredefinedField) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *PredefinedField) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *PredefinedField) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetTitle

`func (o *PredefinedField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PredefinedField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PredefinedField) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PredefinedField) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PredefinedField) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PredefinedField) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *PredefinedField) GetType() FieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PredefinedField) GetTypeOk() (*FieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PredefinedField) SetType(v FieldType)`

SetType sets Type field to given value.


### GetOptionList

`func (o *PredefinedField) GetOptionList() string`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *PredefinedField) GetOptionListOk() (*string, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *PredefinedField) SetOptionList(v string)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *PredefinedField) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### SetOptionListNil

`func (o *PredefinedField) SetOptionListNil(b bool)`

 SetOptionListNil sets the value for OptionList to be an explicit nil

### UnsetOptionList
`func (o *PredefinedField) UnsetOptionList()`

UnsetOptionList ensures that no value is present for OptionList, not even an explicit nil
### GetRequiredInTemplate

`func (o *PredefinedField) GetRequiredInTemplate() bool`

GetRequiredInTemplate returns the RequiredInTemplate field if non-nil, zero value otherwise.

### GetRequiredInTemplateOk

`func (o *PredefinedField) GetRequiredInTemplateOk() (*bool, bool)`

GetRequiredInTemplateOk returns a tuple with the RequiredInTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInTemplate

`func (o *PredefinedField) SetRequiredInTemplate(v bool)`

SetRequiredInTemplate sets RequiredInTemplate field to given value.

### HasRequiredInTemplate

`func (o *PredefinedField) HasRequiredInTemplate() bool`

HasRequiredInTemplate returns a boolean if a field has been set.

### GetFieldRequired

`func (o *PredefinedField) GetFieldRequired() bool`

GetFieldRequired returns the FieldRequired field if non-nil, zero value otherwise.

### GetFieldRequiredOk

`func (o *PredefinedField) GetFieldRequiredOk() (*bool, bool)`

GetFieldRequiredOk returns a tuple with the FieldRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldRequired

`func (o *PredefinedField) SetFieldRequired(v bool)`

SetFieldRequired sets FieldRequired field to given value.

### HasFieldRequired

`func (o *PredefinedField) HasFieldRequired() bool`

HasFieldRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


