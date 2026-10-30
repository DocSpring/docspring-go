# TemplateDiffSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldsAdded** | **int32** | How many fields were added. | 
**FieldsRemoved** | **int32** | How many fields were removed. | 
**FieldsChanged** | **int32** | How many fields have changed settings. | 
**FieldsReordered** | **int32** | How many fields moved in the template's field order. | 
**SettingsChanged** | **int32** | How many of the template's settings changed. | 
**TotalChanges** | **int32** | All the changes, counting a new document as one. Zero when the versions are the same. | 

## Methods

### NewTemplateDiffSummary

`func NewTemplateDiffSummary(fieldsAdded int32, fieldsRemoved int32, fieldsChanged int32, fieldsReordered int32, settingsChanged int32, totalChanges int32, ) *TemplateDiffSummary`

NewTemplateDiffSummary instantiates a new TemplateDiffSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDiffSummaryWithDefaults

`func NewTemplateDiffSummaryWithDefaults() *TemplateDiffSummary`

NewTemplateDiffSummaryWithDefaults instantiates a new TemplateDiffSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldsAdded

`func (o *TemplateDiffSummary) GetFieldsAdded() int32`

GetFieldsAdded returns the FieldsAdded field if non-nil, zero value otherwise.

### GetFieldsAddedOk

`func (o *TemplateDiffSummary) GetFieldsAddedOk() (*int32, bool)`

GetFieldsAddedOk returns a tuple with the FieldsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsAdded

`func (o *TemplateDiffSummary) SetFieldsAdded(v int32)`

SetFieldsAdded sets FieldsAdded field to given value.


### GetFieldsRemoved

`func (o *TemplateDiffSummary) GetFieldsRemoved() int32`

GetFieldsRemoved returns the FieldsRemoved field if non-nil, zero value otherwise.

### GetFieldsRemovedOk

`func (o *TemplateDiffSummary) GetFieldsRemovedOk() (*int32, bool)`

GetFieldsRemovedOk returns a tuple with the FieldsRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsRemoved

`func (o *TemplateDiffSummary) SetFieldsRemoved(v int32)`

SetFieldsRemoved sets FieldsRemoved field to given value.


### GetFieldsChanged

`func (o *TemplateDiffSummary) GetFieldsChanged() int32`

GetFieldsChanged returns the FieldsChanged field if non-nil, zero value otherwise.

### GetFieldsChangedOk

`func (o *TemplateDiffSummary) GetFieldsChangedOk() (*int32, bool)`

GetFieldsChangedOk returns a tuple with the FieldsChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsChanged

`func (o *TemplateDiffSummary) SetFieldsChanged(v int32)`

SetFieldsChanged sets FieldsChanged field to given value.


### GetFieldsReordered

`func (o *TemplateDiffSummary) GetFieldsReordered() int32`

GetFieldsReordered returns the FieldsReordered field if non-nil, zero value otherwise.

### GetFieldsReorderedOk

`func (o *TemplateDiffSummary) GetFieldsReorderedOk() (*int32, bool)`

GetFieldsReorderedOk returns a tuple with the FieldsReordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsReordered

`func (o *TemplateDiffSummary) SetFieldsReordered(v int32)`

SetFieldsReordered sets FieldsReordered field to given value.


### GetSettingsChanged

`func (o *TemplateDiffSummary) GetSettingsChanged() int32`

GetSettingsChanged returns the SettingsChanged field if non-nil, zero value otherwise.

### GetSettingsChangedOk

`func (o *TemplateDiffSummary) GetSettingsChangedOk() (*int32, bool)`

GetSettingsChangedOk returns a tuple with the SettingsChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsChanged

`func (o *TemplateDiffSummary) SetSettingsChanged(v int32)`

SetSettingsChanged sets SettingsChanged field to given value.


### GetTotalChanges

`func (o *TemplateDiffSummary) GetTotalChanges() int32`

GetTotalChanges returns the TotalChanges field if non-nil, zero value otherwise.

### GetTotalChangesOk

`func (o *TemplateDiffSummary) GetTotalChangesOk() (*int32, bool)`

GetTotalChangesOk returns a tuple with the TotalChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChanges

`func (o *TemplateDiffSummary) SetTotalChanges(v int32)`

SetTotalChanges sets TotalChanges field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


