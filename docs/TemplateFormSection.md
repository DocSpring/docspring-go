# TemplateFormSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The section's ID, unique among the template's sections. Any string you choose. | 
**Title** | **string** | The section's heading on the web form. | 
**Description** | **NullableString** | Text shown under the section's heading. Null for none. | 
**FieldIds** | **[]int32** | The IDs of the section's fields (Field's `id`), in the order the web form shows them. A field can be in one section only. | 

## Methods

### NewTemplateFormSection

`func NewTemplateFormSection(id string, title string, description NullableString, fieldIds []int32, ) *TemplateFormSection`

NewTemplateFormSection instantiates a new TemplateFormSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFormSectionWithDefaults

`func NewTemplateFormSectionWithDefaults() *TemplateFormSection`

NewTemplateFormSectionWithDefaults instantiates a new TemplateFormSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateFormSection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateFormSection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateFormSection) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *TemplateFormSection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplateFormSection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplateFormSection) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *TemplateFormSection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateFormSection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateFormSection) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *TemplateFormSection) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateFormSection) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFieldIds

`func (o *TemplateFormSection) GetFieldIds() []int32`

GetFieldIds returns the FieldIds field if non-nil, zero value otherwise.

### GetFieldIdsOk

`func (o *TemplateFormSection) GetFieldIdsOk() (*[]int32, bool)`

GetFieldIdsOk returns a tuple with the FieldIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIds

`func (o *TemplateFormSection) SetFieldIds(v []int32)`

SetFieldIds sets FieldIds field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


