# TemplateFormSectionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The section's ID, unique among the template's sections. Any string you choose. | 
**Title** | **string** | The section's heading on the web form. | 
**Description** | Pointer to **NullableString** | Text shown under the section's heading. Null for none. | [optional] 
**FieldIds** | **[]int32** | The IDs of the section's fields (Field's `id`), in the order the web form shows them. A field can be in one section only. | 

## Methods

### NewTemplateFormSectionInput

`func NewTemplateFormSectionInput(id string, title string, fieldIds []int32, ) *TemplateFormSectionInput`

NewTemplateFormSectionInput instantiates a new TemplateFormSectionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFormSectionInputWithDefaults

`func NewTemplateFormSectionInputWithDefaults() *TemplateFormSectionInput`

NewTemplateFormSectionInputWithDefaults instantiates a new TemplateFormSectionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateFormSectionInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateFormSectionInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateFormSectionInput) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *TemplateFormSectionInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplateFormSectionInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplateFormSectionInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *TemplateFormSectionInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateFormSectionInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateFormSectionInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateFormSectionInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateFormSectionInput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateFormSectionInput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFieldIds

`func (o *TemplateFormSectionInput) GetFieldIds() []int32`

GetFieldIds returns the FieldIds field if non-nil, zero value otherwise.

### GetFieldIdsOk

`func (o *TemplateFormSectionInput) GetFieldIdsOk() (*[]int32, bool)`

GetFieldIdsOk returns a tuple with the FieldIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIds

`func (o *TemplateFormSectionInput) SetFieldIds(v []int32)`

SetFieldIds sets FieldIds field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


