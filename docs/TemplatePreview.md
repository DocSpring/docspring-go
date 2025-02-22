# TemplatePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddDataRequestSubmissionIdFooters** | **bool** |  | 
**AllowAdditionalProperties** | **bool** |  | 
**Description** | **NullableString** |  | 
**DocumentFilename** | **NullableString** |  | 
**DocumentMd5** | **NullableString** |  | 
**DocumentParseError** | **bool** |  | 
**DocumentProcessed** | **bool** |  | 
**DocumentState** | **string** |  | 
**DocumentUrl** | **NullableString** |  | 
**EditableSubmissions** | **bool** |  | 
**EmbedDomains** | **NullableString** |  | 
**EncryptPdfsPassword** | **NullableString** |  | 
**EncryptPdfs** | **bool** |  | 
**ExpirationInterval** | **string** |  | 
**ExpireAfter** | **int32** |  | 
**ExpireSubmissions** | **bool** |  | 
**ExternalPredefinedFieldsTemplateId** | **NullableString** |  | 
**ExternalPredefinedFieldsTemplateName** | **NullableString** |  | 
**FirstTemplate** | **bool** |  | 
**Id** | **NullableString** |  | 
**Locked** | **bool** |  | 
**MergeAuditTrailPdf** | **bool** |  | 
**Name** | **NullableString** |  | 
**PageCount** | **int32** |  | 
**PageDimensions** | **[][]float32** |  | 
**ParentFolderId** | **NullableString** |  | 
**Path** | **NullableString** |  | 
**PermanentDocumentUrl** | **NullableString** |  | 
**PublicSubmissions** | **bool** |  | 
**PublicWebForm** | **bool** |  | 
**RedirectUrl** | **NullableString** |  | 
**SlackWebhookUrl** | **NullableString** |  | 
**TemplateType** | **string** |  | 
**UpdatedAt** | **NullableString** |  | 
**WebhookUrl** | **NullableString** |  | 
**Demo** | **bool** |  | 

## Methods

### NewTemplatePreview

`func NewTemplatePreview(addDataRequestSubmissionIdFooters bool, allowAdditionalProperties bool, description NullableString, documentFilename NullableString, documentMd5 NullableString, documentParseError bool, documentProcessed bool, documentState string, documentUrl NullableString, editableSubmissions bool, embedDomains NullableString, encryptPdfsPassword NullableString, encryptPdfs bool, expirationInterval string, expireAfter int32, expireSubmissions bool, externalPredefinedFieldsTemplateId NullableString, externalPredefinedFieldsTemplateName NullableString, firstTemplate bool, id NullableString, locked bool, mergeAuditTrailPdf bool, name NullableString, pageCount int32, pageDimensions [][]float32, parentFolderId NullableString, path NullableString, permanentDocumentUrl NullableString, publicSubmissions bool, publicWebForm bool, redirectUrl NullableString, slackWebhookUrl NullableString, templateType string, updatedAt NullableString, webhookUrl NullableString, demo bool, ) *TemplatePreview`

NewTemplatePreview instantiates a new TemplatePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePreviewWithDefaults

`func NewTemplatePreviewWithDefaults() *TemplatePreview`

NewTemplatePreviewWithDefaults instantiates a new TemplatePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddDataRequestSubmissionIdFooters

`func (o *TemplatePreview) GetAddDataRequestSubmissionIdFooters() bool`

GetAddDataRequestSubmissionIdFooters returns the AddDataRequestSubmissionIdFooters field if non-nil, zero value otherwise.

### GetAddDataRequestSubmissionIdFootersOk

`func (o *TemplatePreview) GetAddDataRequestSubmissionIdFootersOk() (*bool, bool)`

GetAddDataRequestSubmissionIdFootersOk returns a tuple with the AddDataRequestSubmissionIdFooters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddDataRequestSubmissionIdFooters

`func (o *TemplatePreview) SetAddDataRequestSubmissionIdFooters(v bool)`

SetAddDataRequestSubmissionIdFooters sets AddDataRequestSubmissionIdFooters field to given value.


### GetAllowAdditionalProperties

`func (o *TemplatePreview) GetAllowAdditionalProperties() bool`

GetAllowAdditionalProperties returns the AllowAdditionalProperties field if non-nil, zero value otherwise.

### GetAllowAdditionalPropertiesOk

`func (o *TemplatePreview) GetAllowAdditionalPropertiesOk() (*bool, bool)`

GetAllowAdditionalPropertiesOk returns a tuple with the AllowAdditionalProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdditionalProperties

`func (o *TemplatePreview) SetAllowAdditionalProperties(v bool)`

SetAllowAdditionalProperties sets AllowAdditionalProperties field to given value.


### GetDescription

`func (o *TemplatePreview) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplatePreview) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplatePreview) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *TemplatePreview) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplatePreview) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDocumentFilename

`func (o *TemplatePreview) GetDocumentFilename() string`

GetDocumentFilename returns the DocumentFilename field if non-nil, zero value otherwise.

### GetDocumentFilenameOk

`func (o *TemplatePreview) GetDocumentFilenameOk() (*string, bool)`

GetDocumentFilenameOk returns a tuple with the DocumentFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFilename

`func (o *TemplatePreview) SetDocumentFilename(v string)`

SetDocumentFilename sets DocumentFilename field to given value.


### SetDocumentFilenameNil

`func (o *TemplatePreview) SetDocumentFilenameNil(b bool)`

 SetDocumentFilenameNil sets the value for DocumentFilename to be an explicit nil

### UnsetDocumentFilename
`func (o *TemplatePreview) UnsetDocumentFilename()`

UnsetDocumentFilename ensures that no value is present for DocumentFilename, not even an explicit nil
### GetDocumentMd5

`func (o *TemplatePreview) GetDocumentMd5() string`

GetDocumentMd5 returns the DocumentMd5 field if non-nil, zero value otherwise.

### GetDocumentMd5Ok

`func (o *TemplatePreview) GetDocumentMd5Ok() (*string, bool)`

GetDocumentMd5Ok returns a tuple with the DocumentMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMd5

`func (o *TemplatePreview) SetDocumentMd5(v string)`

SetDocumentMd5 sets DocumentMd5 field to given value.


### SetDocumentMd5Nil

`func (o *TemplatePreview) SetDocumentMd5Nil(b bool)`

 SetDocumentMd5Nil sets the value for DocumentMd5 to be an explicit nil

### UnsetDocumentMd5
`func (o *TemplatePreview) UnsetDocumentMd5()`

UnsetDocumentMd5 ensures that no value is present for DocumentMd5, not even an explicit nil
### GetDocumentParseError

`func (o *TemplatePreview) GetDocumentParseError() bool`

GetDocumentParseError returns the DocumentParseError field if non-nil, zero value otherwise.

### GetDocumentParseErrorOk

`func (o *TemplatePreview) GetDocumentParseErrorOk() (*bool, bool)`

GetDocumentParseErrorOk returns a tuple with the DocumentParseError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentParseError

`func (o *TemplatePreview) SetDocumentParseError(v bool)`

SetDocumentParseError sets DocumentParseError field to given value.


### GetDocumentProcessed

`func (o *TemplatePreview) GetDocumentProcessed() bool`

GetDocumentProcessed returns the DocumentProcessed field if non-nil, zero value otherwise.

### GetDocumentProcessedOk

`func (o *TemplatePreview) GetDocumentProcessedOk() (*bool, bool)`

GetDocumentProcessedOk returns a tuple with the DocumentProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentProcessed

`func (o *TemplatePreview) SetDocumentProcessed(v bool)`

SetDocumentProcessed sets DocumentProcessed field to given value.


### GetDocumentState

`func (o *TemplatePreview) GetDocumentState() string`

GetDocumentState returns the DocumentState field if non-nil, zero value otherwise.

### GetDocumentStateOk

`func (o *TemplatePreview) GetDocumentStateOk() (*string, bool)`

GetDocumentStateOk returns a tuple with the DocumentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentState

`func (o *TemplatePreview) SetDocumentState(v string)`

SetDocumentState sets DocumentState field to given value.


### GetDocumentUrl

`func (o *TemplatePreview) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *TemplatePreview) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *TemplatePreview) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.


### SetDocumentUrlNil

`func (o *TemplatePreview) SetDocumentUrlNil(b bool)`

 SetDocumentUrlNil sets the value for DocumentUrl to be an explicit nil

### UnsetDocumentUrl
`func (o *TemplatePreview) UnsetDocumentUrl()`

UnsetDocumentUrl ensures that no value is present for DocumentUrl, not even an explicit nil
### GetEditableSubmissions

`func (o *TemplatePreview) GetEditableSubmissions() bool`

GetEditableSubmissions returns the EditableSubmissions field if non-nil, zero value otherwise.

### GetEditableSubmissionsOk

`func (o *TemplatePreview) GetEditableSubmissionsOk() (*bool, bool)`

GetEditableSubmissionsOk returns a tuple with the EditableSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableSubmissions

`func (o *TemplatePreview) SetEditableSubmissions(v bool)`

SetEditableSubmissions sets EditableSubmissions field to given value.


### GetEmbedDomains

`func (o *TemplatePreview) GetEmbedDomains() string`

GetEmbedDomains returns the EmbedDomains field if non-nil, zero value otherwise.

### GetEmbedDomainsOk

`func (o *TemplatePreview) GetEmbedDomainsOk() (*string, bool)`

GetEmbedDomainsOk returns a tuple with the EmbedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedDomains

`func (o *TemplatePreview) SetEmbedDomains(v string)`

SetEmbedDomains sets EmbedDomains field to given value.


### SetEmbedDomainsNil

`func (o *TemplatePreview) SetEmbedDomainsNil(b bool)`

 SetEmbedDomainsNil sets the value for EmbedDomains to be an explicit nil

### UnsetEmbedDomains
`func (o *TemplatePreview) UnsetEmbedDomains()`

UnsetEmbedDomains ensures that no value is present for EmbedDomains, not even an explicit nil
### GetEncryptPdfsPassword

`func (o *TemplatePreview) GetEncryptPdfsPassword() string`

GetEncryptPdfsPassword returns the EncryptPdfsPassword field if non-nil, zero value otherwise.

### GetEncryptPdfsPasswordOk

`func (o *TemplatePreview) GetEncryptPdfsPasswordOk() (*string, bool)`

GetEncryptPdfsPasswordOk returns a tuple with the EncryptPdfsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptPdfsPassword

`func (o *TemplatePreview) SetEncryptPdfsPassword(v string)`

SetEncryptPdfsPassword sets EncryptPdfsPassword field to given value.


### SetEncryptPdfsPasswordNil

`func (o *TemplatePreview) SetEncryptPdfsPasswordNil(b bool)`

 SetEncryptPdfsPasswordNil sets the value for EncryptPdfsPassword to be an explicit nil

### UnsetEncryptPdfsPassword
`func (o *TemplatePreview) UnsetEncryptPdfsPassword()`

UnsetEncryptPdfsPassword ensures that no value is present for EncryptPdfsPassword, not even an explicit nil
### GetEncryptPdfs

`func (o *TemplatePreview) GetEncryptPdfs() bool`

GetEncryptPdfs returns the EncryptPdfs field if non-nil, zero value otherwise.

### GetEncryptPdfsOk

`func (o *TemplatePreview) GetEncryptPdfsOk() (*bool, bool)`

GetEncryptPdfsOk returns a tuple with the EncryptPdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptPdfs

`func (o *TemplatePreview) SetEncryptPdfs(v bool)`

SetEncryptPdfs sets EncryptPdfs field to given value.


### GetExpirationInterval

`func (o *TemplatePreview) GetExpirationInterval() string`

GetExpirationInterval returns the ExpirationInterval field if non-nil, zero value otherwise.

### GetExpirationIntervalOk

`func (o *TemplatePreview) GetExpirationIntervalOk() (*string, bool)`

GetExpirationIntervalOk returns a tuple with the ExpirationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationInterval

`func (o *TemplatePreview) SetExpirationInterval(v string)`

SetExpirationInterval sets ExpirationInterval field to given value.


### GetExpireAfter

`func (o *TemplatePreview) GetExpireAfter() int32`

GetExpireAfter returns the ExpireAfter field if non-nil, zero value otherwise.

### GetExpireAfterOk

`func (o *TemplatePreview) GetExpireAfterOk() (*int32, bool)`

GetExpireAfterOk returns a tuple with the ExpireAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfter

`func (o *TemplatePreview) SetExpireAfter(v int32)`

SetExpireAfter sets ExpireAfter field to given value.


### GetExpireSubmissions

`func (o *TemplatePreview) GetExpireSubmissions() bool`

GetExpireSubmissions returns the ExpireSubmissions field if non-nil, zero value otherwise.

### GetExpireSubmissionsOk

`func (o *TemplatePreview) GetExpireSubmissionsOk() (*bool, bool)`

GetExpireSubmissionsOk returns a tuple with the ExpireSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireSubmissions

`func (o *TemplatePreview) SetExpireSubmissions(v bool)`

SetExpireSubmissions sets ExpireSubmissions field to given value.


### GetExternalPredefinedFieldsTemplateId

`func (o *TemplatePreview) GetExternalPredefinedFieldsTemplateId() string`

GetExternalPredefinedFieldsTemplateId returns the ExternalPredefinedFieldsTemplateId field if non-nil, zero value otherwise.

### GetExternalPredefinedFieldsTemplateIdOk

`func (o *TemplatePreview) GetExternalPredefinedFieldsTemplateIdOk() (*string, bool)`

GetExternalPredefinedFieldsTemplateIdOk returns a tuple with the ExternalPredefinedFieldsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPredefinedFieldsTemplateId

`func (o *TemplatePreview) SetExternalPredefinedFieldsTemplateId(v string)`

SetExternalPredefinedFieldsTemplateId sets ExternalPredefinedFieldsTemplateId field to given value.


### SetExternalPredefinedFieldsTemplateIdNil

`func (o *TemplatePreview) SetExternalPredefinedFieldsTemplateIdNil(b bool)`

 SetExternalPredefinedFieldsTemplateIdNil sets the value for ExternalPredefinedFieldsTemplateId to be an explicit nil

### UnsetExternalPredefinedFieldsTemplateId
`func (o *TemplatePreview) UnsetExternalPredefinedFieldsTemplateId()`

UnsetExternalPredefinedFieldsTemplateId ensures that no value is present for ExternalPredefinedFieldsTemplateId, not even an explicit nil
### GetExternalPredefinedFieldsTemplateName

`func (o *TemplatePreview) GetExternalPredefinedFieldsTemplateName() string`

GetExternalPredefinedFieldsTemplateName returns the ExternalPredefinedFieldsTemplateName field if non-nil, zero value otherwise.

### GetExternalPredefinedFieldsTemplateNameOk

`func (o *TemplatePreview) GetExternalPredefinedFieldsTemplateNameOk() (*string, bool)`

GetExternalPredefinedFieldsTemplateNameOk returns a tuple with the ExternalPredefinedFieldsTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPredefinedFieldsTemplateName

`func (o *TemplatePreview) SetExternalPredefinedFieldsTemplateName(v string)`

SetExternalPredefinedFieldsTemplateName sets ExternalPredefinedFieldsTemplateName field to given value.


### SetExternalPredefinedFieldsTemplateNameNil

`func (o *TemplatePreview) SetExternalPredefinedFieldsTemplateNameNil(b bool)`

 SetExternalPredefinedFieldsTemplateNameNil sets the value for ExternalPredefinedFieldsTemplateName to be an explicit nil

### UnsetExternalPredefinedFieldsTemplateName
`func (o *TemplatePreview) UnsetExternalPredefinedFieldsTemplateName()`

UnsetExternalPredefinedFieldsTemplateName ensures that no value is present for ExternalPredefinedFieldsTemplateName, not even an explicit nil
### GetFirstTemplate

`func (o *TemplatePreview) GetFirstTemplate() bool`

GetFirstTemplate returns the FirstTemplate field if non-nil, zero value otherwise.

### GetFirstTemplateOk

`func (o *TemplatePreview) GetFirstTemplateOk() (*bool, bool)`

GetFirstTemplateOk returns a tuple with the FirstTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTemplate

`func (o *TemplatePreview) SetFirstTemplate(v bool)`

SetFirstTemplate sets FirstTemplate field to given value.


### GetId

`func (o *TemplatePreview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplatePreview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplatePreview) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *TemplatePreview) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TemplatePreview) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLocked

`func (o *TemplatePreview) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *TemplatePreview) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *TemplatePreview) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetMergeAuditTrailPdf

`func (o *TemplatePreview) GetMergeAuditTrailPdf() bool`

GetMergeAuditTrailPdf returns the MergeAuditTrailPdf field if non-nil, zero value otherwise.

### GetMergeAuditTrailPdfOk

`func (o *TemplatePreview) GetMergeAuditTrailPdfOk() (*bool, bool)`

GetMergeAuditTrailPdfOk returns a tuple with the MergeAuditTrailPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAuditTrailPdf

`func (o *TemplatePreview) SetMergeAuditTrailPdf(v bool)`

SetMergeAuditTrailPdf sets MergeAuditTrailPdf field to given value.


### GetName

`func (o *TemplatePreview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplatePreview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplatePreview) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *TemplatePreview) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplatePreview) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPageCount

`func (o *TemplatePreview) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *TemplatePreview) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *TemplatePreview) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.


### GetPageDimensions

`func (o *TemplatePreview) GetPageDimensions() [][]float32`

GetPageDimensions returns the PageDimensions field if non-nil, zero value otherwise.

### GetPageDimensionsOk

`func (o *TemplatePreview) GetPageDimensionsOk() (*[][]float32, bool)`

GetPageDimensionsOk returns a tuple with the PageDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageDimensions

`func (o *TemplatePreview) SetPageDimensions(v [][]float32)`

SetPageDimensions sets PageDimensions field to given value.


### SetPageDimensionsNil

`func (o *TemplatePreview) SetPageDimensionsNil(b bool)`

 SetPageDimensionsNil sets the value for PageDimensions to be an explicit nil

### UnsetPageDimensions
`func (o *TemplatePreview) UnsetPageDimensions()`

UnsetPageDimensions ensures that no value is present for PageDimensions, not even an explicit nil
### GetParentFolderId

`func (o *TemplatePreview) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *TemplatePreview) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *TemplatePreview) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.


### SetParentFolderIdNil

`func (o *TemplatePreview) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *TemplatePreview) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetPath

`func (o *TemplatePreview) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TemplatePreview) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TemplatePreview) SetPath(v string)`

SetPath sets Path field to given value.


### SetPathNil

`func (o *TemplatePreview) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *TemplatePreview) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPermanentDocumentUrl

`func (o *TemplatePreview) GetPermanentDocumentUrl() string`

GetPermanentDocumentUrl returns the PermanentDocumentUrl field if non-nil, zero value otherwise.

### GetPermanentDocumentUrlOk

`func (o *TemplatePreview) GetPermanentDocumentUrlOk() (*string, bool)`

GetPermanentDocumentUrlOk returns a tuple with the PermanentDocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentDocumentUrl

`func (o *TemplatePreview) SetPermanentDocumentUrl(v string)`

SetPermanentDocumentUrl sets PermanentDocumentUrl field to given value.


### SetPermanentDocumentUrlNil

`func (o *TemplatePreview) SetPermanentDocumentUrlNil(b bool)`

 SetPermanentDocumentUrlNil sets the value for PermanentDocumentUrl to be an explicit nil

### UnsetPermanentDocumentUrl
`func (o *TemplatePreview) UnsetPermanentDocumentUrl()`

UnsetPermanentDocumentUrl ensures that no value is present for PermanentDocumentUrl, not even an explicit nil
### GetPublicSubmissions

`func (o *TemplatePreview) GetPublicSubmissions() bool`

GetPublicSubmissions returns the PublicSubmissions field if non-nil, zero value otherwise.

### GetPublicSubmissionsOk

`func (o *TemplatePreview) GetPublicSubmissionsOk() (*bool, bool)`

GetPublicSubmissionsOk returns a tuple with the PublicSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSubmissions

`func (o *TemplatePreview) SetPublicSubmissions(v bool)`

SetPublicSubmissions sets PublicSubmissions field to given value.


### GetPublicWebForm

`func (o *TemplatePreview) GetPublicWebForm() bool`

GetPublicWebForm returns the PublicWebForm field if non-nil, zero value otherwise.

### GetPublicWebFormOk

`func (o *TemplatePreview) GetPublicWebFormOk() (*bool, bool)`

GetPublicWebFormOk returns a tuple with the PublicWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicWebForm

`func (o *TemplatePreview) SetPublicWebForm(v bool)`

SetPublicWebForm sets PublicWebForm field to given value.


### GetRedirectUrl

`func (o *TemplatePreview) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *TemplatePreview) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *TemplatePreview) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### SetRedirectUrlNil

`func (o *TemplatePreview) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *TemplatePreview) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetSlackWebhookUrl

`func (o *TemplatePreview) GetSlackWebhookUrl() string`

GetSlackWebhookUrl returns the SlackWebhookUrl field if non-nil, zero value otherwise.

### GetSlackWebhookUrlOk

`func (o *TemplatePreview) GetSlackWebhookUrlOk() (*string, bool)`

GetSlackWebhookUrlOk returns a tuple with the SlackWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhookUrl

`func (o *TemplatePreview) SetSlackWebhookUrl(v string)`

SetSlackWebhookUrl sets SlackWebhookUrl field to given value.


### SetSlackWebhookUrlNil

`func (o *TemplatePreview) SetSlackWebhookUrlNil(b bool)`

 SetSlackWebhookUrlNil sets the value for SlackWebhookUrl to be an explicit nil

### UnsetSlackWebhookUrl
`func (o *TemplatePreview) UnsetSlackWebhookUrl()`

UnsetSlackWebhookUrl ensures that no value is present for SlackWebhookUrl, not even an explicit nil
### GetTemplateType

`func (o *TemplatePreview) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *TemplatePreview) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *TemplatePreview) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.


### GetUpdatedAt

`func (o *TemplatePreview) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplatePreview) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplatePreview) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *TemplatePreview) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TemplatePreview) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetWebhookUrl

`func (o *TemplatePreview) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *TemplatePreview) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *TemplatePreview) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### SetWebhookUrlNil

`func (o *TemplatePreview) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *TemplatePreview) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetDemo

`func (o *TemplatePreview) GetDemo() bool`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *TemplatePreview) GetDemoOk() (*bool, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *TemplatePreview) SetDemo(v bool)`

SetDemo sets Demo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


