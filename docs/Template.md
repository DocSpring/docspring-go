# Template

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
**Defaults** | **map[string]interface{}** |  | 
**FieldOrder** | **[][]float32** |  | 
**Fields** | **map[string]interface{}** |  | 
**FooterHtml** | **NullableString** |  | 
**HeaderHtml** | **NullableString** |  | 
**HtmlEngineOptions** | **map[string]interface{}** |  | 
**Html** | **NullableString** |  | 
**PredefinedFields** | **[]map[string]interface{}** |  | 
**Scss** | **NullableString** |  | 
**SharedFieldData** | **map[string]interface{}** |  | 

## Methods

### NewTemplate

`func NewTemplate(addDataRequestSubmissionIdFooters bool, allowAdditionalProperties bool, description NullableString, documentFilename NullableString, documentMd5 NullableString, documentParseError bool, documentProcessed bool, documentState string, documentUrl NullableString, editableSubmissions bool, embedDomains NullableString, encryptPdfsPassword NullableString, encryptPdfs bool, expirationInterval string, expireAfter int32, expireSubmissions bool, externalPredefinedFieldsTemplateId NullableString, externalPredefinedFieldsTemplateName NullableString, firstTemplate bool, id NullableString, locked bool, mergeAuditTrailPdf bool, name NullableString, pageCount int32, pageDimensions [][]float32, parentFolderId NullableString, path NullableString, permanentDocumentUrl NullableString, publicSubmissions bool, publicWebForm bool, redirectUrl NullableString, slackWebhookUrl NullableString, templateType string, updatedAt NullableString, webhookUrl NullableString, demo bool, defaults map[string]interface{}, fieldOrder [][]float32, fields map[string]interface{}, footerHtml NullableString, headerHtml NullableString, htmlEngineOptions map[string]interface{}, html NullableString, predefinedFields []map[string]interface{}, scss NullableString, sharedFieldData map[string]interface{}, ) *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddDataRequestSubmissionIdFooters

`func (o *Template) GetAddDataRequestSubmissionIdFooters() bool`

GetAddDataRequestSubmissionIdFooters returns the AddDataRequestSubmissionIdFooters field if non-nil, zero value otherwise.

### GetAddDataRequestSubmissionIdFootersOk

`func (o *Template) GetAddDataRequestSubmissionIdFootersOk() (*bool, bool)`

GetAddDataRequestSubmissionIdFootersOk returns a tuple with the AddDataRequestSubmissionIdFooters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddDataRequestSubmissionIdFooters

`func (o *Template) SetAddDataRequestSubmissionIdFooters(v bool)`

SetAddDataRequestSubmissionIdFooters sets AddDataRequestSubmissionIdFooters field to given value.


### GetAllowAdditionalProperties

`func (o *Template) GetAllowAdditionalProperties() bool`

GetAllowAdditionalProperties returns the AllowAdditionalProperties field if non-nil, zero value otherwise.

### GetAllowAdditionalPropertiesOk

`func (o *Template) GetAllowAdditionalPropertiesOk() (*bool, bool)`

GetAllowAdditionalPropertiesOk returns a tuple with the AllowAdditionalProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdditionalProperties

`func (o *Template) SetAllowAdditionalProperties(v bool)`

SetAllowAdditionalProperties sets AllowAdditionalProperties field to given value.


### GetDescription

`func (o *Template) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Template) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Template) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Template) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Template) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDocumentFilename

`func (o *Template) GetDocumentFilename() string`

GetDocumentFilename returns the DocumentFilename field if non-nil, zero value otherwise.

### GetDocumentFilenameOk

`func (o *Template) GetDocumentFilenameOk() (*string, bool)`

GetDocumentFilenameOk returns a tuple with the DocumentFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFilename

`func (o *Template) SetDocumentFilename(v string)`

SetDocumentFilename sets DocumentFilename field to given value.


### SetDocumentFilenameNil

`func (o *Template) SetDocumentFilenameNil(b bool)`

 SetDocumentFilenameNil sets the value for DocumentFilename to be an explicit nil

### UnsetDocumentFilename
`func (o *Template) UnsetDocumentFilename()`

UnsetDocumentFilename ensures that no value is present for DocumentFilename, not even an explicit nil
### GetDocumentMd5

`func (o *Template) GetDocumentMd5() string`

GetDocumentMd5 returns the DocumentMd5 field if non-nil, zero value otherwise.

### GetDocumentMd5Ok

`func (o *Template) GetDocumentMd5Ok() (*string, bool)`

GetDocumentMd5Ok returns a tuple with the DocumentMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMd5

`func (o *Template) SetDocumentMd5(v string)`

SetDocumentMd5 sets DocumentMd5 field to given value.


### SetDocumentMd5Nil

`func (o *Template) SetDocumentMd5Nil(b bool)`

 SetDocumentMd5Nil sets the value for DocumentMd5 to be an explicit nil

### UnsetDocumentMd5
`func (o *Template) UnsetDocumentMd5()`

UnsetDocumentMd5 ensures that no value is present for DocumentMd5, not even an explicit nil
### GetDocumentParseError

`func (o *Template) GetDocumentParseError() bool`

GetDocumentParseError returns the DocumentParseError field if non-nil, zero value otherwise.

### GetDocumentParseErrorOk

`func (o *Template) GetDocumentParseErrorOk() (*bool, bool)`

GetDocumentParseErrorOk returns a tuple with the DocumentParseError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentParseError

`func (o *Template) SetDocumentParseError(v bool)`

SetDocumentParseError sets DocumentParseError field to given value.


### GetDocumentProcessed

`func (o *Template) GetDocumentProcessed() bool`

GetDocumentProcessed returns the DocumentProcessed field if non-nil, zero value otherwise.

### GetDocumentProcessedOk

`func (o *Template) GetDocumentProcessedOk() (*bool, bool)`

GetDocumentProcessedOk returns a tuple with the DocumentProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentProcessed

`func (o *Template) SetDocumentProcessed(v bool)`

SetDocumentProcessed sets DocumentProcessed field to given value.


### GetDocumentState

`func (o *Template) GetDocumentState() string`

GetDocumentState returns the DocumentState field if non-nil, zero value otherwise.

### GetDocumentStateOk

`func (o *Template) GetDocumentStateOk() (*string, bool)`

GetDocumentStateOk returns a tuple with the DocumentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentState

`func (o *Template) SetDocumentState(v string)`

SetDocumentState sets DocumentState field to given value.


### GetDocumentUrl

`func (o *Template) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *Template) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *Template) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.


### SetDocumentUrlNil

`func (o *Template) SetDocumentUrlNil(b bool)`

 SetDocumentUrlNil sets the value for DocumentUrl to be an explicit nil

### UnsetDocumentUrl
`func (o *Template) UnsetDocumentUrl()`

UnsetDocumentUrl ensures that no value is present for DocumentUrl, not even an explicit nil
### GetEditableSubmissions

`func (o *Template) GetEditableSubmissions() bool`

GetEditableSubmissions returns the EditableSubmissions field if non-nil, zero value otherwise.

### GetEditableSubmissionsOk

`func (o *Template) GetEditableSubmissionsOk() (*bool, bool)`

GetEditableSubmissionsOk returns a tuple with the EditableSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableSubmissions

`func (o *Template) SetEditableSubmissions(v bool)`

SetEditableSubmissions sets EditableSubmissions field to given value.


### GetEmbedDomains

`func (o *Template) GetEmbedDomains() string`

GetEmbedDomains returns the EmbedDomains field if non-nil, zero value otherwise.

### GetEmbedDomainsOk

`func (o *Template) GetEmbedDomainsOk() (*string, bool)`

GetEmbedDomainsOk returns a tuple with the EmbedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedDomains

`func (o *Template) SetEmbedDomains(v string)`

SetEmbedDomains sets EmbedDomains field to given value.


### SetEmbedDomainsNil

`func (o *Template) SetEmbedDomainsNil(b bool)`

 SetEmbedDomainsNil sets the value for EmbedDomains to be an explicit nil

### UnsetEmbedDomains
`func (o *Template) UnsetEmbedDomains()`

UnsetEmbedDomains ensures that no value is present for EmbedDomains, not even an explicit nil
### GetEncryptPdfsPassword

`func (o *Template) GetEncryptPdfsPassword() string`

GetEncryptPdfsPassword returns the EncryptPdfsPassword field if non-nil, zero value otherwise.

### GetEncryptPdfsPasswordOk

`func (o *Template) GetEncryptPdfsPasswordOk() (*string, bool)`

GetEncryptPdfsPasswordOk returns a tuple with the EncryptPdfsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptPdfsPassword

`func (o *Template) SetEncryptPdfsPassword(v string)`

SetEncryptPdfsPassword sets EncryptPdfsPassword field to given value.


### SetEncryptPdfsPasswordNil

`func (o *Template) SetEncryptPdfsPasswordNil(b bool)`

 SetEncryptPdfsPasswordNil sets the value for EncryptPdfsPassword to be an explicit nil

### UnsetEncryptPdfsPassword
`func (o *Template) UnsetEncryptPdfsPassword()`

UnsetEncryptPdfsPassword ensures that no value is present for EncryptPdfsPassword, not even an explicit nil
### GetEncryptPdfs

`func (o *Template) GetEncryptPdfs() bool`

GetEncryptPdfs returns the EncryptPdfs field if non-nil, zero value otherwise.

### GetEncryptPdfsOk

`func (o *Template) GetEncryptPdfsOk() (*bool, bool)`

GetEncryptPdfsOk returns a tuple with the EncryptPdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptPdfs

`func (o *Template) SetEncryptPdfs(v bool)`

SetEncryptPdfs sets EncryptPdfs field to given value.


### GetExpirationInterval

`func (o *Template) GetExpirationInterval() string`

GetExpirationInterval returns the ExpirationInterval field if non-nil, zero value otherwise.

### GetExpirationIntervalOk

`func (o *Template) GetExpirationIntervalOk() (*string, bool)`

GetExpirationIntervalOk returns a tuple with the ExpirationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationInterval

`func (o *Template) SetExpirationInterval(v string)`

SetExpirationInterval sets ExpirationInterval field to given value.


### GetExpireAfter

`func (o *Template) GetExpireAfter() int32`

GetExpireAfter returns the ExpireAfter field if non-nil, zero value otherwise.

### GetExpireAfterOk

`func (o *Template) GetExpireAfterOk() (*int32, bool)`

GetExpireAfterOk returns a tuple with the ExpireAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfter

`func (o *Template) SetExpireAfter(v int32)`

SetExpireAfter sets ExpireAfter field to given value.


### GetExpireSubmissions

`func (o *Template) GetExpireSubmissions() bool`

GetExpireSubmissions returns the ExpireSubmissions field if non-nil, zero value otherwise.

### GetExpireSubmissionsOk

`func (o *Template) GetExpireSubmissionsOk() (*bool, bool)`

GetExpireSubmissionsOk returns a tuple with the ExpireSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireSubmissions

`func (o *Template) SetExpireSubmissions(v bool)`

SetExpireSubmissions sets ExpireSubmissions field to given value.


### GetExternalPredefinedFieldsTemplateId

`func (o *Template) GetExternalPredefinedFieldsTemplateId() string`

GetExternalPredefinedFieldsTemplateId returns the ExternalPredefinedFieldsTemplateId field if non-nil, zero value otherwise.

### GetExternalPredefinedFieldsTemplateIdOk

`func (o *Template) GetExternalPredefinedFieldsTemplateIdOk() (*string, bool)`

GetExternalPredefinedFieldsTemplateIdOk returns a tuple with the ExternalPredefinedFieldsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPredefinedFieldsTemplateId

`func (o *Template) SetExternalPredefinedFieldsTemplateId(v string)`

SetExternalPredefinedFieldsTemplateId sets ExternalPredefinedFieldsTemplateId field to given value.


### SetExternalPredefinedFieldsTemplateIdNil

`func (o *Template) SetExternalPredefinedFieldsTemplateIdNil(b bool)`

 SetExternalPredefinedFieldsTemplateIdNil sets the value for ExternalPredefinedFieldsTemplateId to be an explicit nil

### UnsetExternalPredefinedFieldsTemplateId
`func (o *Template) UnsetExternalPredefinedFieldsTemplateId()`

UnsetExternalPredefinedFieldsTemplateId ensures that no value is present for ExternalPredefinedFieldsTemplateId, not even an explicit nil
### GetExternalPredefinedFieldsTemplateName

`func (o *Template) GetExternalPredefinedFieldsTemplateName() string`

GetExternalPredefinedFieldsTemplateName returns the ExternalPredefinedFieldsTemplateName field if non-nil, zero value otherwise.

### GetExternalPredefinedFieldsTemplateNameOk

`func (o *Template) GetExternalPredefinedFieldsTemplateNameOk() (*string, bool)`

GetExternalPredefinedFieldsTemplateNameOk returns a tuple with the ExternalPredefinedFieldsTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPredefinedFieldsTemplateName

`func (o *Template) SetExternalPredefinedFieldsTemplateName(v string)`

SetExternalPredefinedFieldsTemplateName sets ExternalPredefinedFieldsTemplateName field to given value.


### SetExternalPredefinedFieldsTemplateNameNil

`func (o *Template) SetExternalPredefinedFieldsTemplateNameNil(b bool)`

 SetExternalPredefinedFieldsTemplateNameNil sets the value for ExternalPredefinedFieldsTemplateName to be an explicit nil

### UnsetExternalPredefinedFieldsTemplateName
`func (o *Template) UnsetExternalPredefinedFieldsTemplateName()`

UnsetExternalPredefinedFieldsTemplateName ensures that no value is present for ExternalPredefinedFieldsTemplateName, not even an explicit nil
### GetFirstTemplate

`func (o *Template) GetFirstTemplate() bool`

GetFirstTemplate returns the FirstTemplate field if non-nil, zero value otherwise.

### GetFirstTemplateOk

`func (o *Template) GetFirstTemplateOk() (*bool, bool)`

GetFirstTemplateOk returns a tuple with the FirstTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTemplate

`func (o *Template) SetFirstTemplate(v bool)`

SetFirstTemplate sets FirstTemplate field to given value.


### GetId

`func (o *Template) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Template) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Template) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *Template) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Template) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLocked

`func (o *Template) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Template) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Template) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetMergeAuditTrailPdf

`func (o *Template) GetMergeAuditTrailPdf() bool`

GetMergeAuditTrailPdf returns the MergeAuditTrailPdf field if non-nil, zero value otherwise.

### GetMergeAuditTrailPdfOk

`func (o *Template) GetMergeAuditTrailPdfOk() (*bool, bool)`

GetMergeAuditTrailPdfOk returns a tuple with the MergeAuditTrailPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAuditTrailPdf

`func (o *Template) SetMergeAuditTrailPdf(v bool)`

SetMergeAuditTrailPdf sets MergeAuditTrailPdf field to given value.


### GetName

`func (o *Template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Template) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Template) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Template) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPageCount

`func (o *Template) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *Template) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *Template) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.


### GetPageDimensions

`func (o *Template) GetPageDimensions() [][]float32`

GetPageDimensions returns the PageDimensions field if non-nil, zero value otherwise.

### GetPageDimensionsOk

`func (o *Template) GetPageDimensionsOk() (*[][]float32, bool)`

GetPageDimensionsOk returns a tuple with the PageDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageDimensions

`func (o *Template) SetPageDimensions(v [][]float32)`

SetPageDimensions sets PageDimensions field to given value.


### SetPageDimensionsNil

`func (o *Template) SetPageDimensionsNil(b bool)`

 SetPageDimensionsNil sets the value for PageDimensions to be an explicit nil

### UnsetPageDimensions
`func (o *Template) UnsetPageDimensions()`

UnsetPageDimensions ensures that no value is present for PageDimensions, not even an explicit nil
### GetParentFolderId

`func (o *Template) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *Template) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *Template) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.


### SetParentFolderIdNil

`func (o *Template) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *Template) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetPath

`func (o *Template) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Template) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Template) SetPath(v string)`

SetPath sets Path field to given value.


### SetPathNil

`func (o *Template) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *Template) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPermanentDocumentUrl

`func (o *Template) GetPermanentDocumentUrl() string`

GetPermanentDocumentUrl returns the PermanentDocumentUrl field if non-nil, zero value otherwise.

### GetPermanentDocumentUrlOk

`func (o *Template) GetPermanentDocumentUrlOk() (*string, bool)`

GetPermanentDocumentUrlOk returns a tuple with the PermanentDocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentDocumentUrl

`func (o *Template) SetPermanentDocumentUrl(v string)`

SetPermanentDocumentUrl sets PermanentDocumentUrl field to given value.


### SetPermanentDocumentUrlNil

`func (o *Template) SetPermanentDocumentUrlNil(b bool)`

 SetPermanentDocumentUrlNil sets the value for PermanentDocumentUrl to be an explicit nil

### UnsetPermanentDocumentUrl
`func (o *Template) UnsetPermanentDocumentUrl()`

UnsetPermanentDocumentUrl ensures that no value is present for PermanentDocumentUrl, not even an explicit nil
### GetPublicSubmissions

`func (o *Template) GetPublicSubmissions() bool`

GetPublicSubmissions returns the PublicSubmissions field if non-nil, zero value otherwise.

### GetPublicSubmissionsOk

`func (o *Template) GetPublicSubmissionsOk() (*bool, bool)`

GetPublicSubmissionsOk returns a tuple with the PublicSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSubmissions

`func (o *Template) SetPublicSubmissions(v bool)`

SetPublicSubmissions sets PublicSubmissions field to given value.


### GetPublicWebForm

`func (o *Template) GetPublicWebForm() bool`

GetPublicWebForm returns the PublicWebForm field if non-nil, zero value otherwise.

### GetPublicWebFormOk

`func (o *Template) GetPublicWebFormOk() (*bool, bool)`

GetPublicWebFormOk returns a tuple with the PublicWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicWebForm

`func (o *Template) SetPublicWebForm(v bool)`

SetPublicWebForm sets PublicWebForm field to given value.


### GetRedirectUrl

`func (o *Template) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *Template) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *Template) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### SetRedirectUrlNil

`func (o *Template) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *Template) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetSlackWebhookUrl

`func (o *Template) GetSlackWebhookUrl() string`

GetSlackWebhookUrl returns the SlackWebhookUrl field if non-nil, zero value otherwise.

### GetSlackWebhookUrlOk

`func (o *Template) GetSlackWebhookUrlOk() (*string, bool)`

GetSlackWebhookUrlOk returns a tuple with the SlackWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhookUrl

`func (o *Template) SetSlackWebhookUrl(v string)`

SetSlackWebhookUrl sets SlackWebhookUrl field to given value.


### SetSlackWebhookUrlNil

`func (o *Template) SetSlackWebhookUrlNil(b bool)`

 SetSlackWebhookUrlNil sets the value for SlackWebhookUrl to be an explicit nil

### UnsetSlackWebhookUrl
`func (o *Template) UnsetSlackWebhookUrl()`

UnsetSlackWebhookUrl ensures that no value is present for SlackWebhookUrl, not even an explicit nil
### GetTemplateType

`func (o *Template) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *Template) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *Template) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.


### GetUpdatedAt

`func (o *Template) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Template) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Template) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Template) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Template) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetWebhookUrl

`func (o *Template) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *Template) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *Template) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### SetWebhookUrlNil

`func (o *Template) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *Template) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetDemo

`func (o *Template) GetDemo() bool`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *Template) GetDemoOk() (*bool, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *Template) SetDemo(v bool)`

SetDemo sets Demo field to given value.


### GetDefaults

`func (o *Template) GetDefaults() map[string]interface{}`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *Template) GetDefaultsOk() (*map[string]interface{}, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *Template) SetDefaults(v map[string]interface{})`

SetDefaults sets Defaults field to given value.


### GetFieldOrder

`func (o *Template) GetFieldOrder() [][]float32`

GetFieldOrder returns the FieldOrder field if non-nil, zero value otherwise.

### GetFieldOrderOk

`func (o *Template) GetFieldOrderOk() (*[][]float32, bool)`

GetFieldOrderOk returns a tuple with the FieldOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldOrder

`func (o *Template) SetFieldOrder(v [][]float32)`

SetFieldOrder sets FieldOrder field to given value.


### GetFields

`func (o *Template) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Template) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Template) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.


### GetFooterHtml

`func (o *Template) GetFooterHtml() string`

GetFooterHtml returns the FooterHtml field if non-nil, zero value otherwise.

### GetFooterHtmlOk

`func (o *Template) GetFooterHtmlOk() (*string, bool)`

GetFooterHtmlOk returns a tuple with the FooterHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterHtml

`func (o *Template) SetFooterHtml(v string)`

SetFooterHtml sets FooterHtml field to given value.


### SetFooterHtmlNil

`func (o *Template) SetFooterHtmlNil(b bool)`

 SetFooterHtmlNil sets the value for FooterHtml to be an explicit nil

### UnsetFooterHtml
`func (o *Template) UnsetFooterHtml()`

UnsetFooterHtml ensures that no value is present for FooterHtml, not even an explicit nil
### GetHeaderHtml

`func (o *Template) GetHeaderHtml() string`

GetHeaderHtml returns the HeaderHtml field if non-nil, zero value otherwise.

### GetHeaderHtmlOk

`func (o *Template) GetHeaderHtmlOk() (*string, bool)`

GetHeaderHtmlOk returns a tuple with the HeaderHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderHtml

`func (o *Template) SetHeaderHtml(v string)`

SetHeaderHtml sets HeaderHtml field to given value.


### SetHeaderHtmlNil

`func (o *Template) SetHeaderHtmlNil(b bool)`

 SetHeaderHtmlNil sets the value for HeaderHtml to be an explicit nil

### UnsetHeaderHtml
`func (o *Template) UnsetHeaderHtml()`

UnsetHeaderHtml ensures that no value is present for HeaderHtml, not even an explicit nil
### GetHtmlEngineOptions

`func (o *Template) GetHtmlEngineOptions() map[string]interface{}`

GetHtmlEngineOptions returns the HtmlEngineOptions field if non-nil, zero value otherwise.

### GetHtmlEngineOptionsOk

`func (o *Template) GetHtmlEngineOptionsOk() (*map[string]interface{}, bool)`

GetHtmlEngineOptionsOk returns a tuple with the HtmlEngineOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlEngineOptions

`func (o *Template) SetHtmlEngineOptions(v map[string]interface{})`

SetHtmlEngineOptions sets HtmlEngineOptions field to given value.


### GetHtml

`func (o *Template) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *Template) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *Template) SetHtml(v string)`

SetHtml sets Html field to given value.


### SetHtmlNil

`func (o *Template) SetHtmlNil(b bool)`

 SetHtmlNil sets the value for Html to be an explicit nil

### UnsetHtml
`func (o *Template) UnsetHtml()`

UnsetHtml ensures that no value is present for Html, not even an explicit nil
### GetPredefinedFields

`func (o *Template) GetPredefinedFields() []map[string]interface{}`

GetPredefinedFields returns the PredefinedFields field if non-nil, zero value otherwise.

### GetPredefinedFieldsOk

`func (o *Template) GetPredefinedFieldsOk() (*[]map[string]interface{}, bool)`

GetPredefinedFieldsOk returns a tuple with the PredefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedFields

`func (o *Template) SetPredefinedFields(v []map[string]interface{})`

SetPredefinedFields sets PredefinedFields field to given value.


### GetScss

`func (o *Template) GetScss() string`

GetScss returns the Scss field if non-nil, zero value otherwise.

### GetScssOk

`func (o *Template) GetScssOk() (*string, bool)`

GetScssOk returns a tuple with the Scss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScss

`func (o *Template) SetScss(v string)`

SetScss sets Scss field to given value.


### SetScssNil

`func (o *Template) SetScssNil(b bool)`

 SetScssNil sets the value for Scss to be an explicit nil

### UnsetScss
`func (o *Template) UnsetScss()`

UnsetScss ensures that no value is present for Scss, not even an explicit nil
### GetSharedFieldData

`func (o *Template) GetSharedFieldData() map[string]interface{}`

GetSharedFieldData returns the SharedFieldData field if non-nil, zero value otherwise.

### GetSharedFieldDataOk

`func (o *Template) GetSharedFieldDataOk() (*map[string]interface{}, bool)`

GetSharedFieldDataOk returns a tuple with the SharedFieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedFieldData

`func (o *Template) SetSharedFieldData(v map[string]interface{})`

SetSharedFieldData sets SharedFieldData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


