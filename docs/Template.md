# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The template's ID. | 
**Name** | **NullableString** | The template's name. | 
**Description** | **NullableString** | Shown under the template's name at the top of its hosted forms, and used as their meta description. Not shown on embedded forms or signature requests. | 
**TemplateType** | [**TemplateType**](TemplateType.md) | `pdf`: fields placed on an uploaded PDF. `html`: a PDF rendered from HTML and SCSS. It can't be changed after the template is created. | 
**Ready** | **bool** | Whether the template can generate PDFs. HTML templates are always ready. A PDF template is ready once DocSpring has stored its PDF, made an image of each page and finished looking for fields on each page, which can be after `document.state` becomes `processed`. A PDF template with no PDF (`document` is null) isn't ready. | 
**Document** | [**NullableTemplateDocument**](TemplateDocument.md) | The PDF of a PDF template. Null for HTML templates, and for a PDF template that has no PDF yet. | 
**ParentFolderId** | **NullableString** | The ID of the folder the template is in. Null for a template at the root. Setting it moves the template and all of its published versions. | 
**Path** | **string** | The path of the folder the template is in, such as `/Contracts`. `/` for a template at the root. | 
**LatestVersion** | **NullableString** | The template's latest published version. Null if no version has been published. | 
**VersionPublishedAt** | **NullableTime** | When this version was published. Null for the draft. | 
**Locked** | **bool** | A locked template can't be changed or deleted until it's unlocked (`locked` can always be changed). | 
**LastChangedAt** | **NullableTime** | When the draft was last changed. | 
**CreatedAt** | **time.Time** | When the template was created. | 
**UpdatedAt** | **time.Time** | When the template was last updated. | 
**Version** | **NullableString** | The published version this is, such as `1.2.0`. Null for the draft. | 
**AddSignatureRequestSubmissionIdFooters** | **bool** | Print the submission ID in the footer of each page of PDFs signed through signature requests. | 
**AiFieldNamingLanguage** | [**NullableAiFieldNamingLanguage**](AiFieldNamingLanguage.md) | The language AI field setup names fields in: `auto` uses the form's language when it's written in the Latin alphabet without accents, and English otherwise; `english` always uses English; `form_language` uses the form's language and script. New templates take it from your account's template settings. Null for HTML templates. | 
**AllowAdditionalProperties** | **bool** | Accept submission data with properties that aren't fields of the template. | 
**AllowSkippingOptionalFields** | **bool** | Visual forms and signature request forms show the Continue button as soon as the required fields are filled in, without visiting each optional field. The `allowSkippingOptionalFields` embed option overrides it. | 
**DefaultFormMode** | [**TemplateFormMode**](TemplateFormMode.md) | The hosted form that /templates/{template_id}/form shows when its URL has no `mode` parameter: `web` (the web form), `visual` (the visual form, where fields are filled in on the PDF) or `dual` (both, in tabs). HTML templates have no visual form, so `web` is the only value they accept. | 
**EditableSubmissions** | **bool** | Generate PDFs whose text and check fields stay fillable ("Generate Editable PDFs"). A submission's `editable` overrides it. | 
**EmbedDomains** | **[]string** | The sites that can embed the template's live signature requests and its embedded template editor: a host such as `app.example.com`, with an optional leading wildcard (`*.example.com`), port (`localhost:3000`) or scheme (`https://app.example.com`; without one, `http` and `https` both work). Browsers open a live signature request in a frame only on these sites (and DocSpring's), so with an empty list it can't be embedded anywhere; a test signature request can be embedded on any site. The embedded template editor loads only for a page on one of these sites (by its Referer or Origin header), and on any site when the list is empty. Forms embedded on a page (Visual Forms JS and the simple form) aren't restricted. | 
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
**LastChangedBy** | [**NullableTemplateActor**](TemplateActor.md) | Who last changed the draft: a user in your account, an API token, or someone using the embedded template editor. Null if no change has been recorded. | 
**Defaults** | [**TemplateDefaults**](TemplateDefaults.md) |  | 
**FieldDependencies** | [**[]TemplateFieldDependency**](TemplateFieldDependency.md) | Conditional logic: each rule shows or hides its target fields, or makes them required, depending on the values of other fields. | 
**FieldGroups** | [**[]FieldGroup**](FieldGroup.md) | The groups of the template's fields, with their settings. A field's name puts it in groups: `address/city` is in the group `address`, and `items/0/price` is in the list `items` and its items' group `items/~`. | 
**PredefinedFields** | [**[]PredefinedField**](PredefinedField.md) | The field names that people using the embedded template editor choose from, each with the data type and settings a field gets when it's given that name. Setting it replaces the whole list; an empty list turns them off. | 
**PredefinedFieldsSampleData** | **map[string]interface{}** | Sample submission data that sets `predefined_fields`: each value becomes a predefined field, with a type that fits it (`string`, `number`, `boolean`, `date`, `email` or `url`) and `required_in_template` true. Nested objects and arrays become names with `/`, such as `customer/email`. Without `predefined_fields` in the same request, it replaces them; with them, it adds to them and sets the types of the ones it has values for. | 
**Html** | **NullableString** | The HTML of an HTML template, as a Liquid template. Null for PDF templates. | 
**Scss** | **NullableString** | The SCSS of an HTML template. | 
**HeaderHtml** | **NullableString** | The HTML printed at the top of each page of an HTML template. | 
**FooterHtml** | **NullableString** | The HTML printed at the bottom of each page of an HTML template. | 
**HtmlEngineOptions** | [**NullableHtmlEngineOptions**](HtmlEngineOptions.md) | How an HTML template's PDFs are printed. Null for PDF templates. Setting it replaces the whole object: options left out use their defaults. | 
**HtmlEscapeData** | **NullableBool** | Whether an HTML template HTML-escapes the values it prints. When true, `{{ value }}` prints a value as text, so a `&lt;` or `&amp;` in it is printed as that character, and `{{ value | raw }}` prints a value as HTML. New HTML templates default to true. Templates created before October 2026, or with API v1, default to false, which prints every value as HTML. Null for PDF templates. | 
**EmbedEditToken** | **NullableString** | The token that opens the template in the embedded template editor. | 
**Versions** | [**[]TemplateVersion**](TemplateVersion.md) | The template's published versions, newest first (up to 100). | 

## Methods

### NewTemplate

`func NewTemplate(id string, name NullableString, description NullableString, templateType TemplateType, ready bool, document NullableTemplateDocument, parentFolderId NullableString, path string, latestVersion NullableString, versionPublishedAt NullableTime, locked bool, lastChangedAt NullableTime, createdAt time.Time, updatedAt time.Time, version NullableString, addSignatureRequestSubmissionIdFooters bool, aiFieldNamingLanguage NullableAiFieldNamingLanguage, allowAdditionalProperties bool, allowSkippingOptionalFields bool, defaultFormMode TemplateFormMode, editableSubmissions bool, embedDomains []string, embedEditAllowSettings bool, embedEditAllowVersioning bool, embedEditAllowDocumentReplacement bool, encryptPdfs bool, hasEncryptPdfsPassword bool, expireSubmissions bool, expireAfter int32, expirationInterval ExpirationInterval, mergeAuditTrailPdf bool, parseDatesUsingTemplateDateFormat bool, publicSubmissions bool, publicWebForm bool, redirectUrl NullableString, hasSlackWebhookUrl bool, lastChangedBy NullableTemplateActor, defaults TemplateDefaults, fieldDependencies []TemplateFieldDependency, fieldGroups []FieldGroup, predefinedFields []PredefinedField, predefinedFieldsSampleData map[string]interface{}, html NullableString, scss NullableString, headerHtml NullableString, footerHtml NullableString, htmlEngineOptions NullableHtmlEngineOptions, htmlEscapeData NullableBool, embedEditToken NullableString, versions []TemplateVersion, ) *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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
### GetTemplateType

`func (o *Template) GetTemplateType() TemplateType`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *Template) GetTemplateTypeOk() (*TemplateType, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *Template) SetTemplateType(v TemplateType)`

SetTemplateType sets TemplateType field to given value.


### GetReady

`func (o *Template) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *Template) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *Template) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetDocument

`func (o *Template) GetDocument() TemplateDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *Template) GetDocumentOk() (*TemplateDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *Template) SetDocument(v TemplateDocument)`

SetDocument sets Document field to given value.


### SetDocumentNil

`func (o *Template) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *Template) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
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


### GetLatestVersion

`func (o *Template) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *Template) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *Template) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.


### SetLatestVersionNil

`func (o *Template) SetLatestVersionNil(b bool)`

 SetLatestVersionNil sets the value for LatestVersion to be an explicit nil

### UnsetLatestVersion
`func (o *Template) UnsetLatestVersion()`

UnsetLatestVersion ensures that no value is present for LatestVersion, not even an explicit nil
### GetVersionPublishedAt

`func (o *Template) GetVersionPublishedAt() time.Time`

GetVersionPublishedAt returns the VersionPublishedAt field if non-nil, zero value otherwise.

### GetVersionPublishedAtOk

`func (o *Template) GetVersionPublishedAtOk() (*time.Time, bool)`

GetVersionPublishedAtOk returns a tuple with the VersionPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPublishedAt

`func (o *Template) SetVersionPublishedAt(v time.Time)`

SetVersionPublishedAt sets VersionPublishedAt field to given value.


### SetVersionPublishedAtNil

`func (o *Template) SetVersionPublishedAtNil(b bool)`

 SetVersionPublishedAtNil sets the value for VersionPublishedAt to be an explicit nil

### UnsetVersionPublishedAt
`func (o *Template) UnsetVersionPublishedAt()`

UnsetVersionPublishedAt ensures that no value is present for VersionPublishedAt, not even an explicit nil
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


### GetLastChangedAt

`func (o *Template) GetLastChangedAt() time.Time`

GetLastChangedAt returns the LastChangedAt field if non-nil, zero value otherwise.

### GetLastChangedAtOk

`func (o *Template) GetLastChangedAtOk() (*time.Time, bool)`

GetLastChangedAtOk returns a tuple with the LastChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedAt

`func (o *Template) SetLastChangedAt(v time.Time)`

SetLastChangedAt sets LastChangedAt field to given value.


### SetLastChangedAtNil

`func (o *Template) SetLastChangedAtNil(b bool)`

 SetLastChangedAtNil sets the value for LastChangedAt to be an explicit nil

### UnsetLastChangedAt
`func (o *Template) UnsetLastChangedAt()`

UnsetLastChangedAt ensures that no value is present for LastChangedAt, not even an explicit nil
### GetCreatedAt

`func (o *Template) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Template) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Template) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Template) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Template) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Template) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *Template) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Template) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Template) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *Template) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Template) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAddSignatureRequestSubmissionIdFooters

`func (o *Template) GetAddSignatureRequestSubmissionIdFooters() bool`

GetAddSignatureRequestSubmissionIdFooters returns the AddSignatureRequestSubmissionIdFooters field if non-nil, zero value otherwise.

### GetAddSignatureRequestSubmissionIdFootersOk

`func (o *Template) GetAddSignatureRequestSubmissionIdFootersOk() (*bool, bool)`

GetAddSignatureRequestSubmissionIdFootersOk returns a tuple with the AddSignatureRequestSubmissionIdFooters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSignatureRequestSubmissionIdFooters

`func (o *Template) SetAddSignatureRequestSubmissionIdFooters(v bool)`

SetAddSignatureRequestSubmissionIdFooters sets AddSignatureRequestSubmissionIdFooters field to given value.


### GetAiFieldNamingLanguage

`func (o *Template) GetAiFieldNamingLanguage() AiFieldNamingLanguage`

GetAiFieldNamingLanguage returns the AiFieldNamingLanguage field if non-nil, zero value otherwise.

### GetAiFieldNamingLanguageOk

`func (o *Template) GetAiFieldNamingLanguageOk() (*AiFieldNamingLanguage, bool)`

GetAiFieldNamingLanguageOk returns a tuple with the AiFieldNamingLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiFieldNamingLanguage

`func (o *Template) SetAiFieldNamingLanguage(v AiFieldNamingLanguage)`

SetAiFieldNamingLanguage sets AiFieldNamingLanguage field to given value.


### SetAiFieldNamingLanguageNil

`func (o *Template) SetAiFieldNamingLanguageNil(b bool)`

 SetAiFieldNamingLanguageNil sets the value for AiFieldNamingLanguage to be an explicit nil

### UnsetAiFieldNamingLanguage
`func (o *Template) UnsetAiFieldNamingLanguage()`

UnsetAiFieldNamingLanguage ensures that no value is present for AiFieldNamingLanguage, not even an explicit nil
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


### GetAllowSkippingOptionalFields

`func (o *Template) GetAllowSkippingOptionalFields() bool`

GetAllowSkippingOptionalFields returns the AllowSkippingOptionalFields field if non-nil, zero value otherwise.

### GetAllowSkippingOptionalFieldsOk

`func (o *Template) GetAllowSkippingOptionalFieldsOk() (*bool, bool)`

GetAllowSkippingOptionalFieldsOk returns a tuple with the AllowSkippingOptionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSkippingOptionalFields

`func (o *Template) SetAllowSkippingOptionalFields(v bool)`

SetAllowSkippingOptionalFields sets AllowSkippingOptionalFields field to given value.


### GetDefaultFormMode

`func (o *Template) GetDefaultFormMode() TemplateFormMode`

GetDefaultFormMode returns the DefaultFormMode field if non-nil, zero value otherwise.

### GetDefaultFormModeOk

`func (o *Template) GetDefaultFormModeOk() (*TemplateFormMode, bool)`

GetDefaultFormModeOk returns a tuple with the DefaultFormMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFormMode

`func (o *Template) SetDefaultFormMode(v TemplateFormMode)`

SetDefaultFormMode sets DefaultFormMode field to given value.


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

`func (o *Template) GetEmbedDomains() []string`

GetEmbedDomains returns the EmbedDomains field if non-nil, zero value otherwise.

### GetEmbedDomainsOk

`func (o *Template) GetEmbedDomainsOk() (*[]string, bool)`

GetEmbedDomainsOk returns a tuple with the EmbedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedDomains

`func (o *Template) SetEmbedDomains(v []string)`

SetEmbedDomains sets EmbedDomains field to given value.


### GetEmbedEditAllowSettings

`func (o *Template) GetEmbedEditAllowSettings() bool`

GetEmbedEditAllowSettings returns the EmbedEditAllowSettings field if non-nil, zero value otherwise.

### GetEmbedEditAllowSettingsOk

`func (o *Template) GetEmbedEditAllowSettingsOk() (*bool, bool)`

GetEmbedEditAllowSettingsOk returns a tuple with the EmbedEditAllowSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowSettings

`func (o *Template) SetEmbedEditAllowSettings(v bool)`

SetEmbedEditAllowSettings sets EmbedEditAllowSettings field to given value.


### GetEmbedEditAllowVersioning

`func (o *Template) GetEmbedEditAllowVersioning() bool`

GetEmbedEditAllowVersioning returns the EmbedEditAllowVersioning field if non-nil, zero value otherwise.

### GetEmbedEditAllowVersioningOk

`func (o *Template) GetEmbedEditAllowVersioningOk() (*bool, bool)`

GetEmbedEditAllowVersioningOk returns a tuple with the EmbedEditAllowVersioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowVersioning

`func (o *Template) SetEmbedEditAllowVersioning(v bool)`

SetEmbedEditAllowVersioning sets EmbedEditAllowVersioning field to given value.


### GetEmbedEditAllowDocumentReplacement

`func (o *Template) GetEmbedEditAllowDocumentReplacement() bool`

GetEmbedEditAllowDocumentReplacement returns the EmbedEditAllowDocumentReplacement field if non-nil, zero value otherwise.

### GetEmbedEditAllowDocumentReplacementOk

`func (o *Template) GetEmbedEditAllowDocumentReplacementOk() (*bool, bool)`

GetEmbedEditAllowDocumentReplacementOk returns a tuple with the EmbedEditAllowDocumentReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowDocumentReplacement

`func (o *Template) SetEmbedEditAllowDocumentReplacement(v bool)`

SetEmbedEditAllowDocumentReplacement sets EmbedEditAllowDocumentReplacement field to given value.


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


### GetHasEncryptPdfsPassword

`func (o *Template) GetHasEncryptPdfsPassword() bool`

GetHasEncryptPdfsPassword returns the HasEncryptPdfsPassword field if non-nil, zero value otherwise.

### GetHasEncryptPdfsPasswordOk

`func (o *Template) GetHasEncryptPdfsPasswordOk() (*bool, bool)`

GetHasEncryptPdfsPasswordOk returns a tuple with the HasEncryptPdfsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEncryptPdfsPassword

`func (o *Template) SetHasEncryptPdfsPassword(v bool)`

SetHasEncryptPdfsPassword sets HasEncryptPdfsPassword field to given value.


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


### GetExpirationInterval

`func (o *Template) GetExpirationInterval() ExpirationInterval`

GetExpirationInterval returns the ExpirationInterval field if non-nil, zero value otherwise.

### GetExpirationIntervalOk

`func (o *Template) GetExpirationIntervalOk() (*ExpirationInterval, bool)`

GetExpirationIntervalOk returns a tuple with the ExpirationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationInterval

`func (o *Template) SetExpirationInterval(v ExpirationInterval)`

SetExpirationInterval sets ExpirationInterval field to given value.


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


### GetParseDatesUsingTemplateDateFormat

`func (o *Template) GetParseDatesUsingTemplateDateFormat() bool`

GetParseDatesUsingTemplateDateFormat returns the ParseDatesUsingTemplateDateFormat field if non-nil, zero value otherwise.

### GetParseDatesUsingTemplateDateFormatOk

`func (o *Template) GetParseDatesUsingTemplateDateFormatOk() (*bool, bool)`

GetParseDatesUsingTemplateDateFormatOk returns a tuple with the ParseDatesUsingTemplateDateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseDatesUsingTemplateDateFormat

`func (o *Template) SetParseDatesUsingTemplateDateFormat(v bool)`

SetParseDatesUsingTemplateDateFormat sets ParseDatesUsingTemplateDateFormat field to given value.


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
### GetHasSlackWebhookUrl

`func (o *Template) GetHasSlackWebhookUrl() bool`

GetHasSlackWebhookUrl returns the HasSlackWebhookUrl field if non-nil, zero value otherwise.

### GetHasSlackWebhookUrlOk

`func (o *Template) GetHasSlackWebhookUrlOk() (*bool, bool)`

GetHasSlackWebhookUrlOk returns a tuple with the HasSlackWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSlackWebhookUrl

`func (o *Template) SetHasSlackWebhookUrl(v bool)`

SetHasSlackWebhookUrl sets HasSlackWebhookUrl field to given value.


### GetLastChangedBy

`func (o *Template) GetLastChangedBy() TemplateActor`

GetLastChangedBy returns the LastChangedBy field if non-nil, zero value otherwise.

### GetLastChangedByOk

`func (o *Template) GetLastChangedByOk() (*TemplateActor, bool)`

GetLastChangedByOk returns a tuple with the LastChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedBy

`func (o *Template) SetLastChangedBy(v TemplateActor)`

SetLastChangedBy sets LastChangedBy field to given value.


### SetLastChangedByNil

`func (o *Template) SetLastChangedByNil(b bool)`

 SetLastChangedByNil sets the value for LastChangedBy to be an explicit nil

### UnsetLastChangedBy
`func (o *Template) UnsetLastChangedBy()`

UnsetLastChangedBy ensures that no value is present for LastChangedBy, not even an explicit nil
### GetDefaults

`func (o *Template) GetDefaults() TemplateDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *Template) GetDefaultsOk() (*TemplateDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *Template) SetDefaults(v TemplateDefaults)`

SetDefaults sets Defaults field to given value.


### GetFieldDependencies

`func (o *Template) GetFieldDependencies() []TemplateFieldDependency`

GetFieldDependencies returns the FieldDependencies field if non-nil, zero value otherwise.

### GetFieldDependenciesOk

`func (o *Template) GetFieldDependenciesOk() (*[]TemplateFieldDependency, bool)`

GetFieldDependenciesOk returns a tuple with the FieldDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDependencies

`func (o *Template) SetFieldDependencies(v []TemplateFieldDependency)`

SetFieldDependencies sets FieldDependencies field to given value.


### GetFieldGroups

`func (o *Template) GetFieldGroups() []FieldGroup`

GetFieldGroups returns the FieldGroups field if non-nil, zero value otherwise.

### GetFieldGroupsOk

`func (o *Template) GetFieldGroupsOk() (*[]FieldGroup, bool)`

GetFieldGroupsOk returns a tuple with the FieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroups

`func (o *Template) SetFieldGroups(v []FieldGroup)`

SetFieldGroups sets FieldGroups field to given value.


### GetPredefinedFields

`func (o *Template) GetPredefinedFields() []PredefinedField`

GetPredefinedFields returns the PredefinedFields field if non-nil, zero value otherwise.

### GetPredefinedFieldsOk

`func (o *Template) GetPredefinedFieldsOk() (*[]PredefinedField, bool)`

GetPredefinedFieldsOk returns a tuple with the PredefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedFields

`func (o *Template) SetPredefinedFields(v []PredefinedField)`

SetPredefinedFields sets PredefinedFields field to given value.


### GetPredefinedFieldsSampleData

`func (o *Template) GetPredefinedFieldsSampleData() map[string]interface{}`

GetPredefinedFieldsSampleData returns the PredefinedFieldsSampleData field if non-nil, zero value otherwise.

### GetPredefinedFieldsSampleDataOk

`func (o *Template) GetPredefinedFieldsSampleDataOk() (*map[string]interface{}, bool)`

GetPredefinedFieldsSampleDataOk returns a tuple with the PredefinedFieldsSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedFieldsSampleData

`func (o *Template) SetPredefinedFieldsSampleData(v map[string]interface{})`

SetPredefinedFieldsSampleData sets PredefinedFieldsSampleData field to given value.


### SetPredefinedFieldsSampleDataNil

`func (o *Template) SetPredefinedFieldsSampleDataNil(b bool)`

 SetPredefinedFieldsSampleDataNil sets the value for PredefinedFieldsSampleData to be an explicit nil

### UnsetPredefinedFieldsSampleData
`func (o *Template) UnsetPredefinedFieldsSampleData()`

UnsetPredefinedFieldsSampleData ensures that no value is present for PredefinedFieldsSampleData, not even an explicit nil
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
### GetHtmlEngineOptions

`func (o *Template) GetHtmlEngineOptions() HtmlEngineOptions`

GetHtmlEngineOptions returns the HtmlEngineOptions field if non-nil, zero value otherwise.

### GetHtmlEngineOptionsOk

`func (o *Template) GetHtmlEngineOptionsOk() (*HtmlEngineOptions, bool)`

GetHtmlEngineOptionsOk returns a tuple with the HtmlEngineOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlEngineOptions

`func (o *Template) SetHtmlEngineOptions(v HtmlEngineOptions)`

SetHtmlEngineOptions sets HtmlEngineOptions field to given value.


### SetHtmlEngineOptionsNil

`func (o *Template) SetHtmlEngineOptionsNil(b bool)`

 SetHtmlEngineOptionsNil sets the value for HtmlEngineOptions to be an explicit nil

### UnsetHtmlEngineOptions
`func (o *Template) UnsetHtmlEngineOptions()`

UnsetHtmlEngineOptions ensures that no value is present for HtmlEngineOptions, not even an explicit nil
### GetHtmlEscapeData

`func (o *Template) GetHtmlEscapeData() bool`

GetHtmlEscapeData returns the HtmlEscapeData field if non-nil, zero value otherwise.

### GetHtmlEscapeDataOk

`func (o *Template) GetHtmlEscapeDataOk() (*bool, bool)`

GetHtmlEscapeDataOk returns a tuple with the HtmlEscapeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlEscapeData

`func (o *Template) SetHtmlEscapeData(v bool)`

SetHtmlEscapeData sets HtmlEscapeData field to given value.


### SetHtmlEscapeDataNil

`func (o *Template) SetHtmlEscapeDataNil(b bool)`

 SetHtmlEscapeDataNil sets the value for HtmlEscapeData to be an explicit nil

### UnsetHtmlEscapeData
`func (o *Template) UnsetHtmlEscapeData()`

UnsetHtmlEscapeData ensures that no value is present for HtmlEscapeData, not even an explicit nil
### GetEmbedEditToken

`func (o *Template) GetEmbedEditToken() string`

GetEmbedEditToken returns the EmbedEditToken field if non-nil, zero value otherwise.

### GetEmbedEditTokenOk

`func (o *Template) GetEmbedEditTokenOk() (*string, bool)`

GetEmbedEditTokenOk returns a tuple with the EmbedEditToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditToken

`func (o *Template) SetEmbedEditToken(v string)`

SetEmbedEditToken sets EmbedEditToken field to given value.


### SetEmbedEditTokenNil

`func (o *Template) SetEmbedEditTokenNil(b bool)`

 SetEmbedEditTokenNil sets the value for EmbedEditToken to be an explicit nil

### UnsetEmbedEditToken
`func (o *Template) UnsetEmbedEditToken()`

UnsetEmbedEditToken ensures that no value is present for EmbedEditToken, not even an explicit nil
### GetVersions

`func (o *Template) GetVersions() []TemplateVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Template) GetVersionsOk() (*[]TemplateVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Template) SetVersions(v []TemplateVersion)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


