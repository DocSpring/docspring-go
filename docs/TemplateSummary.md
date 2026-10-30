# TemplateSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The template's ID. | 
**Name** | **NullableString** | The template's name. | 
**Description** | **NullableString** | Shown under the template's name at the top of its hosted forms, and used as their meta description. Not shown on embedded forms or signature requests. | 
**TemplateType** | [**TemplateType**](TemplateType.md) | `pdf`: fields placed on an uploaded PDF. `html`: a PDF rendered from HTML and SCSS. It can't be changed after the template is created. | 
**Ready** | **bool** | Whether the template can generate PDFs. HTML templates are always ready. A PDF template is ready once DocSpring has stored its PDF, made an image of each page and finished looking for fields on each page, which can be after `document.state` becomes `processed`. A PDF template with no PDF (`document` is null) isn't ready. | 
**Document** | [**NullableTemplateDocumentSummary**](TemplateDocumentSummary.md) | The PDF of a PDF template: its file name, page count and processing state. Null for HTML templates, and for a PDF template that has no PDF yet. | 
**ParentFolderId** | **NullableString** | The ID of the folder the template is in. Null for a template at the root. Setting it moves the template and all of its published versions. | 
**Path** | **string** | The path of the folder the template is in, such as `/Contracts`. `/` for a template at the root. | 
**LatestVersion** | **NullableString** | The template's latest published version. Null if no version has been published. | 
**VersionPublishedAt** | **NullableTime** | When this version was published. Null for the draft. | 
**Locked** | **bool** | A locked template can't be changed or deleted until it's unlocked (`locked` can always be changed). | 
**LastChangedAt** | **NullableTime** | When the draft was last changed. | 
**CreatedAt** | **time.Time** | When the template was created. | 
**UpdatedAt** | **time.Time** | When the template was last updated. | 

## Methods

### NewTemplateSummary

`func NewTemplateSummary(id string, name NullableString, description NullableString, templateType TemplateType, ready bool, document NullableTemplateDocumentSummary, parentFolderId NullableString, path string, latestVersion NullableString, versionPublishedAt NullableTime, locked bool, lastChangedAt NullableTime, createdAt time.Time, updatedAt time.Time, ) *TemplateSummary`

NewTemplateSummary instantiates a new TemplateSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateSummaryWithDefaults

`func NewTemplateSummaryWithDefaults() *TemplateSummary`

NewTemplateSummaryWithDefaults instantiates a new TemplateSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateSummary) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TemplateSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateSummary) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *TemplateSummary) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplateSummary) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TemplateSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateSummary) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *TemplateSummary) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateSummary) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTemplateType

`func (o *TemplateSummary) GetTemplateType() TemplateType`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *TemplateSummary) GetTemplateTypeOk() (*TemplateType, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *TemplateSummary) SetTemplateType(v TemplateType)`

SetTemplateType sets TemplateType field to given value.


### GetReady

`func (o *TemplateSummary) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *TemplateSummary) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *TemplateSummary) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetDocument

`func (o *TemplateSummary) GetDocument() TemplateDocumentSummary`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *TemplateSummary) GetDocumentOk() (*TemplateDocumentSummary, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *TemplateSummary) SetDocument(v TemplateDocumentSummary)`

SetDocument sets Document field to given value.


### SetDocumentNil

`func (o *TemplateSummary) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *TemplateSummary) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetParentFolderId

`func (o *TemplateSummary) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *TemplateSummary) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *TemplateSummary) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.


### SetParentFolderIdNil

`func (o *TemplateSummary) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *TemplateSummary) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetPath

`func (o *TemplateSummary) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TemplateSummary) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TemplateSummary) SetPath(v string)`

SetPath sets Path field to given value.


### GetLatestVersion

`func (o *TemplateSummary) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *TemplateSummary) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *TemplateSummary) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.


### SetLatestVersionNil

`func (o *TemplateSummary) SetLatestVersionNil(b bool)`

 SetLatestVersionNil sets the value for LatestVersion to be an explicit nil

### UnsetLatestVersion
`func (o *TemplateSummary) UnsetLatestVersion()`

UnsetLatestVersion ensures that no value is present for LatestVersion, not even an explicit nil
### GetVersionPublishedAt

`func (o *TemplateSummary) GetVersionPublishedAt() time.Time`

GetVersionPublishedAt returns the VersionPublishedAt field if non-nil, zero value otherwise.

### GetVersionPublishedAtOk

`func (o *TemplateSummary) GetVersionPublishedAtOk() (*time.Time, bool)`

GetVersionPublishedAtOk returns a tuple with the VersionPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPublishedAt

`func (o *TemplateSummary) SetVersionPublishedAt(v time.Time)`

SetVersionPublishedAt sets VersionPublishedAt field to given value.


### SetVersionPublishedAtNil

`func (o *TemplateSummary) SetVersionPublishedAtNil(b bool)`

 SetVersionPublishedAtNil sets the value for VersionPublishedAt to be an explicit nil

### UnsetVersionPublishedAt
`func (o *TemplateSummary) UnsetVersionPublishedAt()`

UnsetVersionPublishedAt ensures that no value is present for VersionPublishedAt, not even an explicit nil
### GetLocked

`func (o *TemplateSummary) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *TemplateSummary) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *TemplateSummary) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetLastChangedAt

`func (o *TemplateSummary) GetLastChangedAt() time.Time`

GetLastChangedAt returns the LastChangedAt field if non-nil, zero value otherwise.

### GetLastChangedAtOk

`func (o *TemplateSummary) GetLastChangedAtOk() (*time.Time, bool)`

GetLastChangedAtOk returns a tuple with the LastChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedAt

`func (o *TemplateSummary) SetLastChangedAt(v time.Time)`

SetLastChangedAt sets LastChangedAt field to given value.


### SetLastChangedAtNil

`func (o *TemplateSummary) SetLastChangedAtNil(b bool)`

 SetLastChangedAtNil sets the value for LastChangedAt to be an explicit nil

### UnsetLastChangedAt
`func (o *TemplateSummary) UnsetLastChangedAt()`

UnsetLastChangedAt ensures that no value is present for LastChangedAt, not even an explicit nil
### GetCreatedAt

`func (o *TemplateSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TemplateSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


