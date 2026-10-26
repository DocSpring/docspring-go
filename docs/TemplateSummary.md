# TemplateSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The template's ID. | 
**Name** | **NullableString** | The template's name. | 
**Description** | **NullableString** | Shown under the template's name at the top of its hosted forms, and used as their meta description. Not shown on embedded forms or signature requests. | 
**TemplateType** | [**TemplateType**](TemplateType.md) | `pdf`: fields placed on an uploaded PDF. `html`: a PDF rendered from HTML and SCSS. It can't be changed after the template is created. | 
**Ready** | **bool** | Whether the template can generate PDFs. HTML templates are always ready. A PDF template is ready once its document has been processed (`document.state` is `processed`). | 
**Document** | [**NullableTemplateDocument**](TemplateDocument.md) | The PDF of a PDF template. Null for HTML templates, and for a PDF template that has no PDF yet. | 
**ParentFolderId** | **NullableString** | The ID of the folder the template is in. Null for a template at the root. Setting it moves the template and all of its published versions. | 
**Path** | **string** | The path of the folder the template is in, such as `/Contracts`. `/` for a template at the root. | 
**Version** | **NullableString** | The published version this is, such as `1.2.0`. Null for the draft. | 
**LatestVersion** | **NullableString** | The template's latest published version. Null if no version has been published. | 
**VersionPublishedAt** | **NullableTime** | When this version was published. Null for the draft. | 
**Locked** | **bool** | A locked template can't be changed or deleted until it's unlocked (`locked` can always be changed). | 
**AddSignatureRequestSubmissionIdFooters** | **bool** | Print the submission ID in the footer of each page of PDFs signed through signature requests. | 
**AiFieldNamingLanguage** | [**NullableAiFieldNamingLanguage**](AiFieldNamingLanguage.md) | The language AI field setup names fields in: `auto` uses the form's language when it's written in the Latin alphabet without accents, and English otherwise; `english` always uses English; `form_language` uses the form's language and script. New templates take it from your account's template settings. Null for HTML templates. | 
**AllowAdditionalProperties** | **bool** | Accept submission data with properties that aren't fields of the template. | 
**AllowSkippingOptionalFields** | **bool** | Visual forms and signature request forms show the Continue button as soon as the required fields are filled in, without visiting each optional field. The `allowSkippingOptionalFields` embed option overrides it. | 
**DefaultFormMode** | [**TemplateFormMode**](TemplateFormMode.md) | The hosted form that /templates/{template_id}/form shows when its URL has no `mode` parameter: `web` (the web form), `visual` (the visual form, where fields are filled in on the PDF) or `dual` (both, in tabs). HTML templates have no visual form, so `web` is the only value they accept. | 
**EditableSubmissions** | **bool** | Generate PDFs whose text and check fields stay fillable ("Generate Editable PDFs"). A submission's `editable` overrides it. | 
**EmbedDomains** | **[]string** | The sites where the template's signature requests, forms and embedded template editor can be loaded: a host such as `app.example.com`, with an optional leading wildcard (`*.example.com`), port (`localhost:3000`) or scheme (`https://app.example.com`; without one, `http` and `https` both work). When it's empty, live signature requests don't load on any site, and the embedded template editor loads on every site. | 
**EmbedEditAllowSettings** | **bool** | Allow the embedded template editor to change sensitive template settings (webhooks, redirect URL, encryption, public access, expiration and so on). | 
**EmbedEditAllowVersioning** | **bool** | Allow the embedded template editor to publish, restore, compare and delete published versions. | 
**EmbedEditAllowDocumentReplacement** | **bool** | Allow the embedded template editor to replace the template's PDF. | 
**EncryptPdfs** | **bool** | Encrypt generated PDFs with `encrypt_pdfs_password`. | 
**HasEncryptPdfsPassword** | **bool** | Whether `encrypt_pdfs_password` is set. | 
**ExpireSubmissions** | **bool** | Delete each submission's data and PDF after `expire_after` `expiration_interval`. | 
**ExpireAfter** | **int32** | How many `expiration_interval`s a submission's data and PDF are kept for, when `expire_submissions` is on. At most 5256000 (10 years in minutes). | 
**ExpirationInterval** | [**ExpirationInterval**](ExpirationInterval.md) | The unit of `expire_after`. | 
**MergeAuditTrailPdf** | **bool** | Add the audit trail to the end of PDFs signed through signature requests. | 
**ParseDatesUsingTemplateDateFormat** | **bool** | Date fields also accept NN/NN/YYYY dates (with slashes, dashes or dots), read in the order of the template's default date format: month first for %m/%d/%Y, day first for %d/%m/%Y. When false, date fields only accept YYYY-MM-DD. Applies to PDF templates: an HTML template's date fields only accept YYYY-MM-DD. Date fields that include a time take RFC 3339 date-times either way. | 
**PublicSubmissions** | **bool** | Allow anyone with a submission's ID to download its PDF. | 
**PublicWebForm** | **bool** | Allow anyone to fill in the template's web form. | 
**RedirectUrl** | **NullableString** | Where the web form sends people after they submit it. | 
**HasSlackWebhookUrl** | **bool** | Whether `slack_webhook_url` is set. | 
**LastChangedAt** | **NullableTime** | When the draft was last changed. | 
**LastChangedBy** | [**NullableTemplateActor**](TemplateActor.md) | Who last changed the draft: a user in your account, an API token, or someone using the embedded template editor. Null if no change has been recorded. | 
**CreatedAt** | **time.Time** | When the template was created. | 
**UpdatedAt** | **time.Time** | When the template was last updated. | 

## Methods

### NewTemplateSummary

`func NewTemplateSummary(id string, name NullableString, description NullableString, templateType TemplateType, ready bool, document NullableTemplateDocument, parentFolderId NullableString, path string, version NullableString, latestVersion NullableString, versionPublishedAt NullableTime, locked bool, addSignatureRequestSubmissionIdFooters bool, aiFieldNamingLanguage NullableAiFieldNamingLanguage, allowAdditionalProperties bool, allowSkippingOptionalFields bool, defaultFormMode TemplateFormMode, editableSubmissions bool, embedDomains []string, embedEditAllowSettings bool, embedEditAllowVersioning bool, embedEditAllowDocumentReplacement bool, encryptPdfs bool, hasEncryptPdfsPassword bool, expireSubmissions bool, expireAfter int32, expirationInterval ExpirationInterval, mergeAuditTrailPdf bool, parseDatesUsingTemplateDateFormat bool, publicSubmissions bool, publicWebForm bool, redirectUrl NullableString, hasSlackWebhookUrl bool, lastChangedAt NullableTime, lastChangedBy NullableTemplateActor, createdAt time.Time, updatedAt time.Time, ) *TemplateSummary`

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

`func (o *TemplateSummary) GetDocument() TemplateDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *TemplateSummary) GetDocumentOk() (*TemplateDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *TemplateSummary) SetDocument(v TemplateDocument)`

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


### GetVersion

`func (o *TemplateSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TemplateSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TemplateSummary) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *TemplateSummary) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TemplateSummary) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
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


### GetAddSignatureRequestSubmissionIdFooters

`func (o *TemplateSummary) GetAddSignatureRequestSubmissionIdFooters() bool`

GetAddSignatureRequestSubmissionIdFooters returns the AddSignatureRequestSubmissionIdFooters field if non-nil, zero value otherwise.

### GetAddSignatureRequestSubmissionIdFootersOk

`func (o *TemplateSummary) GetAddSignatureRequestSubmissionIdFootersOk() (*bool, bool)`

GetAddSignatureRequestSubmissionIdFootersOk returns a tuple with the AddSignatureRequestSubmissionIdFooters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSignatureRequestSubmissionIdFooters

`func (o *TemplateSummary) SetAddSignatureRequestSubmissionIdFooters(v bool)`

SetAddSignatureRequestSubmissionIdFooters sets AddSignatureRequestSubmissionIdFooters field to given value.


### GetAiFieldNamingLanguage

`func (o *TemplateSummary) GetAiFieldNamingLanguage() AiFieldNamingLanguage`

GetAiFieldNamingLanguage returns the AiFieldNamingLanguage field if non-nil, zero value otherwise.

### GetAiFieldNamingLanguageOk

`func (o *TemplateSummary) GetAiFieldNamingLanguageOk() (*AiFieldNamingLanguage, bool)`

GetAiFieldNamingLanguageOk returns a tuple with the AiFieldNamingLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiFieldNamingLanguage

`func (o *TemplateSummary) SetAiFieldNamingLanguage(v AiFieldNamingLanguage)`

SetAiFieldNamingLanguage sets AiFieldNamingLanguage field to given value.


### SetAiFieldNamingLanguageNil

`func (o *TemplateSummary) SetAiFieldNamingLanguageNil(b bool)`

 SetAiFieldNamingLanguageNil sets the value for AiFieldNamingLanguage to be an explicit nil

### UnsetAiFieldNamingLanguage
`func (o *TemplateSummary) UnsetAiFieldNamingLanguage()`

UnsetAiFieldNamingLanguage ensures that no value is present for AiFieldNamingLanguage, not even an explicit nil
### GetAllowAdditionalProperties

`func (o *TemplateSummary) GetAllowAdditionalProperties() bool`

GetAllowAdditionalProperties returns the AllowAdditionalProperties field if non-nil, zero value otherwise.

### GetAllowAdditionalPropertiesOk

`func (o *TemplateSummary) GetAllowAdditionalPropertiesOk() (*bool, bool)`

GetAllowAdditionalPropertiesOk returns a tuple with the AllowAdditionalProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdditionalProperties

`func (o *TemplateSummary) SetAllowAdditionalProperties(v bool)`

SetAllowAdditionalProperties sets AllowAdditionalProperties field to given value.


### GetAllowSkippingOptionalFields

`func (o *TemplateSummary) GetAllowSkippingOptionalFields() bool`

GetAllowSkippingOptionalFields returns the AllowSkippingOptionalFields field if non-nil, zero value otherwise.

### GetAllowSkippingOptionalFieldsOk

`func (o *TemplateSummary) GetAllowSkippingOptionalFieldsOk() (*bool, bool)`

GetAllowSkippingOptionalFieldsOk returns a tuple with the AllowSkippingOptionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSkippingOptionalFields

`func (o *TemplateSummary) SetAllowSkippingOptionalFields(v bool)`

SetAllowSkippingOptionalFields sets AllowSkippingOptionalFields field to given value.


### GetDefaultFormMode

`func (o *TemplateSummary) GetDefaultFormMode() TemplateFormMode`

GetDefaultFormMode returns the DefaultFormMode field if non-nil, zero value otherwise.

### GetDefaultFormModeOk

`func (o *TemplateSummary) GetDefaultFormModeOk() (*TemplateFormMode, bool)`

GetDefaultFormModeOk returns a tuple with the DefaultFormMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFormMode

`func (o *TemplateSummary) SetDefaultFormMode(v TemplateFormMode)`

SetDefaultFormMode sets DefaultFormMode field to given value.


### GetEditableSubmissions

`func (o *TemplateSummary) GetEditableSubmissions() bool`

GetEditableSubmissions returns the EditableSubmissions field if non-nil, zero value otherwise.

### GetEditableSubmissionsOk

`func (o *TemplateSummary) GetEditableSubmissionsOk() (*bool, bool)`

GetEditableSubmissionsOk returns a tuple with the EditableSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableSubmissions

`func (o *TemplateSummary) SetEditableSubmissions(v bool)`

SetEditableSubmissions sets EditableSubmissions field to given value.


### GetEmbedDomains

`func (o *TemplateSummary) GetEmbedDomains() []string`

GetEmbedDomains returns the EmbedDomains field if non-nil, zero value otherwise.

### GetEmbedDomainsOk

`func (o *TemplateSummary) GetEmbedDomainsOk() (*[]string, bool)`

GetEmbedDomainsOk returns a tuple with the EmbedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedDomains

`func (o *TemplateSummary) SetEmbedDomains(v []string)`

SetEmbedDomains sets EmbedDomains field to given value.


### GetEmbedEditAllowSettings

`func (o *TemplateSummary) GetEmbedEditAllowSettings() bool`

GetEmbedEditAllowSettings returns the EmbedEditAllowSettings field if non-nil, zero value otherwise.

### GetEmbedEditAllowSettingsOk

`func (o *TemplateSummary) GetEmbedEditAllowSettingsOk() (*bool, bool)`

GetEmbedEditAllowSettingsOk returns a tuple with the EmbedEditAllowSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowSettings

`func (o *TemplateSummary) SetEmbedEditAllowSettings(v bool)`

SetEmbedEditAllowSettings sets EmbedEditAllowSettings field to given value.


### GetEmbedEditAllowVersioning

`func (o *TemplateSummary) GetEmbedEditAllowVersioning() bool`

GetEmbedEditAllowVersioning returns the EmbedEditAllowVersioning field if non-nil, zero value otherwise.

### GetEmbedEditAllowVersioningOk

`func (o *TemplateSummary) GetEmbedEditAllowVersioningOk() (*bool, bool)`

GetEmbedEditAllowVersioningOk returns a tuple with the EmbedEditAllowVersioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowVersioning

`func (o *TemplateSummary) SetEmbedEditAllowVersioning(v bool)`

SetEmbedEditAllowVersioning sets EmbedEditAllowVersioning field to given value.


### GetEmbedEditAllowDocumentReplacement

`func (o *TemplateSummary) GetEmbedEditAllowDocumentReplacement() bool`

GetEmbedEditAllowDocumentReplacement returns the EmbedEditAllowDocumentReplacement field if non-nil, zero value otherwise.

### GetEmbedEditAllowDocumentReplacementOk

`func (o *TemplateSummary) GetEmbedEditAllowDocumentReplacementOk() (*bool, bool)`

GetEmbedEditAllowDocumentReplacementOk returns a tuple with the EmbedEditAllowDocumentReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowDocumentReplacement

`func (o *TemplateSummary) SetEmbedEditAllowDocumentReplacement(v bool)`

SetEmbedEditAllowDocumentReplacement sets EmbedEditAllowDocumentReplacement field to given value.


### GetEncryptPdfs

`func (o *TemplateSummary) GetEncryptPdfs() bool`

GetEncryptPdfs returns the EncryptPdfs field if non-nil, zero value otherwise.

### GetEncryptPdfsOk

`func (o *TemplateSummary) GetEncryptPdfsOk() (*bool, bool)`

GetEncryptPdfsOk returns a tuple with the EncryptPdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptPdfs

`func (o *TemplateSummary) SetEncryptPdfs(v bool)`

SetEncryptPdfs sets EncryptPdfs field to given value.


### GetHasEncryptPdfsPassword

`func (o *TemplateSummary) GetHasEncryptPdfsPassword() bool`

GetHasEncryptPdfsPassword returns the HasEncryptPdfsPassword field if non-nil, zero value otherwise.

### GetHasEncryptPdfsPasswordOk

`func (o *TemplateSummary) GetHasEncryptPdfsPasswordOk() (*bool, bool)`

GetHasEncryptPdfsPasswordOk returns a tuple with the HasEncryptPdfsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEncryptPdfsPassword

`func (o *TemplateSummary) SetHasEncryptPdfsPassword(v bool)`

SetHasEncryptPdfsPassword sets HasEncryptPdfsPassword field to given value.


### GetExpireSubmissions

`func (o *TemplateSummary) GetExpireSubmissions() bool`

GetExpireSubmissions returns the ExpireSubmissions field if non-nil, zero value otherwise.

### GetExpireSubmissionsOk

`func (o *TemplateSummary) GetExpireSubmissionsOk() (*bool, bool)`

GetExpireSubmissionsOk returns a tuple with the ExpireSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireSubmissions

`func (o *TemplateSummary) SetExpireSubmissions(v bool)`

SetExpireSubmissions sets ExpireSubmissions field to given value.


### GetExpireAfter

`func (o *TemplateSummary) GetExpireAfter() int32`

GetExpireAfter returns the ExpireAfter field if non-nil, zero value otherwise.

### GetExpireAfterOk

`func (o *TemplateSummary) GetExpireAfterOk() (*int32, bool)`

GetExpireAfterOk returns a tuple with the ExpireAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfter

`func (o *TemplateSummary) SetExpireAfter(v int32)`

SetExpireAfter sets ExpireAfter field to given value.


### GetExpirationInterval

`func (o *TemplateSummary) GetExpirationInterval() ExpirationInterval`

GetExpirationInterval returns the ExpirationInterval field if non-nil, zero value otherwise.

### GetExpirationIntervalOk

`func (o *TemplateSummary) GetExpirationIntervalOk() (*ExpirationInterval, bool)`

GetExpirationIntervalOk returns a tuple with the ExpirationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationInterval

`func (o *TemplateSummary) SetExpirationInterval(v ExpirationInterval)`

SetExpirationInterval sets ExpirationInterval field to given value.


### GetMergeAuditTrailPdf

`func (o *TemplateSummary) GetMergeAuditTrailPdf() bool`

GetMergeAuditTrailPdf returns the MergeAuditTrailPdf field if non-nil, zero value otherwise.

### GetMergeAuditTrailPdfOk

`func (o *TemplateSummary) GetMergeAuditTrailPdfOk() (*bool, bool)`

GetMergeAuditTrailPdfOk returns a tuple with the MergeAuditTrailPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAuditTrailPdf

`func (o *TemplateSummary) SetMergeAuditTrailPdf(v bool)`

SetMergeAuditTrailPdf sets MergeAuditTrailPdf field to given value.


### GetParseDatesUsingTemplateDateFormat

`func (o *TemplateSummary) GetParseDatesUsingTemplateDateFormat() bool`

GetParseDatesUsingTemplateDateFormat returns the ParseDatesUsingTemplateDateFormat field if non-nil, zero value otherwise.

### GetParseDatesUsingTemplateDateFormatOk

`func (o *TemplateSummary) GetParseDatesUsingTemplateDateFormatOk() (*bool, bool)`

GetParseDatesUsingTemplateDateFormatOk returns a tuple with the ParseDatesUsingTemplateDateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseDatesUsingTemplateDateFormat

`func (o *TemplateSummary) SetParseDatesUsingTemplateDateFormat(v bool)`

SetParseDatesUsingTemplateDateFormat sets ParseDatesUsingTemplateDateFormat field to given value.


### GetPublicSubmissions

`func (o *TemplateSummary) GetPublicSubmissions() bool`

GetPublicSubmissions returns the PublicSubmissions field if non-nil, zero value otherwise.

### GetPublicSubmissionsOk

`func (o *TemplateSummary) GetPublicSubmissionsOk() (*bool, bool)`

GetPublicSubmissionsOk returns a tuple with the PublicSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSubmissions

`func (o *TemplateSummary) SetPublicSubmissions(v bool)`

SetPublicSubmissions sets PublicSubmissions field to given value.


### GetPublicWebForm

`func (o *TemplateSummary) GetPublicWebForm() bool`

GetPublicWebForm returns the PublicWebForm field if non-nil, zero value otherwise.

### GetPublicWebFormOk

`func (o *TemplateSummary) GetPublicWebFormOk() (*bool, bool)`

GetPublicWebFormOk returns a tuple with the PublicWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicWebForm

`func (o *TemplateSummary) SetPublicWebForm(v bool)`

SetPublicWebForm sets PublicWebForm field to given value.


### GetRedirectUrl

`func (o *TemplateSummary) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *TemplateSummary) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *TemplateSummary) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### SetRedirectUrlNil

`func (o *TemplateSummary) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *TemplateSummary) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetHasSlackWebhookUrl

`func (o *TemplateSummary) GetHasSlackWebhookUrl() bool`

GetHasSlackWebhookUrl returns the HasSlackWebhookUrl field if non-nil, zero value otherwise.

### GetHasSlackWebhookUrlOk

`func (o *TemplateSummary) GetHasSlackWebhookUrlOk() (*bool, bool)`

GetHasSlackWebhookUrlOk returns a tuple with the HasSlackWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSlackWebhookUrl

`func (o *TemplateSummary) SetHasSlackWebhookUrl(v bool)`

SetHasSlackWebhookUrl sets HasSlackWebhookUrl field to given value.


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
### GetLastChangedBy

`func (o *TemplateSummary) GetLastChangedBy() TemplateActor`

GetLastChangedBy returns the LastChangedBy field if non-nil, zero value otherwise.

### GetLastChangedByOk

`func (o *TemplateSummary) GetLastChangedByOk() (*TemplateActor, bool)`

GetLastChangedByOk returns a tuple with the LastChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedBy

`func (o *TemplateSummary) SetLastChangedBy(v TemplateActor)`

SetLastChangedBy sets LastChangedBy field to given value.


### SetLastChangedByNil

`func (o *TemplateSummary) SetLastChangedByNil(b bool)`

 SetLastChangedByNil sets the value for LastChangedBy to be an explicit nil

### UnsetLastChangedBy
`func (o *TemplateSummary) UnsetLastChangedBy()`

UnsetLastChangedBy ensures that no value is present for LastChangedBy, not even an explicit nil
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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


