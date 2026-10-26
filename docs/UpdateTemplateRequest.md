# UpdateTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The template's name. | [optional] 
**Description** | Pointer to **NullableString** | Shown under the template's name at the top of its hosted forms, and used as their meta description. Not shown on embedded forms or signature requests. | [optional] 
**ParentFolderId** | Pointer to **NullableString** | The ID of the folder the template is in. Null for a template at the root. Setting it moves the template and all of its published versions. | [optional] 
**Locked** | Pointer to **bool** | A locked template can't be changed or deleted until it's unlocked (`locked` can always be changed). | [optional] 
**AddSignatureRequestSubmissionIdFooters** | Pointer to **bool** | Print the submission ID in the footer of each page of PDFs signed through signature requests. | [optional] 
**AiFieldNamingLanguage** | Pointer to [**AiFieldNamingLanguage**](AiFieldNamingLanguage.md) | The language AI field setup names fields in: `auto` uses the form's language when it's written in the Latin alphabet without accents, and English otherwise; `english` always uses English; `form_language` uses the form's language and script. New templates take it from your account's template settings. Null for HTML templates. | [optional] 
**AllowAdditionalProperties** | Pointer to **bool** | Accept submission data with properties that aren't fields of the template. | [optional] 
**AllowSkippingOptionalFields** | Pointer to **bool** | Visual forms and signature request forms show the Continue button as soon as the required fields are filled in, without visiting each optional field. The `allowSkippingOptionalFields` embed option overrides it. | [optional] 
**DefaultFormMode** | Pointer to [**TemplateFormMode**](TemplateFormMode.md) | The hosted form that /templates/{template_id}/form shows when its URL has no `mode` parameter: `web` (the web form), `visual` (the visual form, where fields are filled in on the PDF) or `dual` (both, in tabs). HTML templates have no visual form, so `web` is the only value they accept. | [optional] 
**EditableSubmissions** | Pointer to **bool** | Generate PDFs whose text and check fields stay fillable ("Generate Editable PDFs"). A submission's `editable` overrides it. | [optional] 
**EmbedDomains** | Pointer to **[]string** | The sites where the template's signature requests, forms and embedded template editor can be loaded: a host such as `app.example.com`, with an optional leading wildcard (`*.example.com`), port (`localhost:3000`) or scheme (`https://app.example.com`; without one, `http` and `https` both work). When it's empty, live signature requests don't load on any site, and the embedded template editor loads on every site. | [optional] 
**EmbedEditAllowSettings** | Pointer to **bool** | Allow the embedded template editor to change sensitive template settings (webhooks, redirect URL, encryption, public access, expiration and so on). | [optional] 
**EmbedEditAllowVersioning** | Pointer to **bool** | Allow the embedded template editor to publish, restore, compare and delete published versions. | [optional] 
**EmbedEditAllowDocumentReplacement** | Pointer to **bool** | Allow the embedded template editor to replace the template's PDF. | [optional] 
**EncryptPdfs** | Pointer to **bool** | Encrypt generated PDFs with `encrypt_pdfs_password`. | [optional] 
**EncryptPdfsPassword** | Pointer to **NullableString** | The password that opens encrypted PDFs. It's never returned; `has_encrypt_pdfs_password` says if it's set. | [optional] 
**ExpireSubmissions** | Pointer to **bool** | Delete each submission's data and PDF after `expire_after` `expiration_interval`. | [optional] 
**ExpireAfter** | Pointer to **int32** | How many `expiration_interval`s a submission's data and PDF are kept for, when `expire_submissions` is on. At most 5256000 (10 years in minutes). | [optional] 
**ExpirationInterval** | Pointer to [**ExpirationInterval**](ExpirationInterval.md) | The unit of `expire_after`. | [optional] 
**MergeAuditTrailPdf** | Pointer to **bool** | Add the audit trail to the end of PDFs signed through signature requests. | [optional] 
**ParseDatesUsingTemplateDateFormat** | Pointer to **bool** | Date fields also accept NN/NN/YYYY dates (with slashes, dashes or dots), read in the order of the template's default date format: month first for %m/%d/%Y, day first for %d/%m/%Y. When false, date fields only accept YYYY-MM-DD. Applies to PDF templates: an HTML template's date fields only accept YYYY-MM-DD. Date fields that include a time take RFC 3339 date-times either way. | [optional] 
**PublicSubmissions** | Pointer to **bool** | Allow anyone with a submission's ID to download its PDF. | [optional] 
**PublicWebForm** | Pointer to **bool** | Allow anyone to fill in the template's web form. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | Where the web form sends people after they submit it. | [optional] 
**SlackWebhookUrl** | Pointer to **NullableString** | A Slack incoming webhook URL that is notified about the template's submissions. It's never returned; `has_slack_webhook_url` says if it's set. | [optional] 
**Defaults** | Pointer to [**TemplateDefaultsInput**](TemplateDefaultsInput.md) | The default style of the template's fields. Setting it replaces the whole object. | [optional] 
**FieldDependencies** | Pointer to [**[]TemplateFieldDependencyInput**](TemplateFieldDependencyInput.md) | The template's conditional logic rules. Setting them replaces all of the rules. | [optional] 
**FieldGroups** | Pointer to [**[]FieldGroupInput**](FieldGroupInput.md) | Changes to the settings of the template's field groups, each with the `name` of the group it changes. Settings left out keep their values, and groups left out are unchanged. | [optional] 
**PredefinedFields** | Pointer to [**[]PredefinedFieldInput**](PredefinedFieldInput.md) | The field names that people using the embedded template editor choose from, each with the data type and settings a field gets when it's given that name. Setting it replaces the whole list; an empty list turns them off. | [optional] 
**PredefinedFieldsSampleData** | Pointer to **map[string]interface{}** | Sample submission data that sets `predefined_fields`: each value becomes a predefined field, with a type that fits it (`string`, `number`, `boolean`, `date`, `email` or `url`) and `required_in_template` true. Nested objects and arrays become names with `/`, such as `customer/email`. Without `predefined_fields` in the same request, it replaces them; with them, it adds to them and sets the types of the ones it has values for. | [optional] 
**Html** | Pointer to **NullableString** | The HTML of an HTML template, as a Liquid template. Null for PDF templates. | [optional] 
**Scss** | Pointer to **NullableString** | The SCSS of an HTML template. | [optional] 
**HeaderHtml** | Pointer to **NullableString** | The HTML printed at the top of each page of an HTML template. | [optional] 
**FooterHtml** | Pointer to **NullableString** | The HTML printed at the bottom of each page of an HTML template. | [optional] 
**HtmlEngineOptions** | Pointer to [**NullableHtmlEngineOptionsInput**](HtmlEngineOptionsInput.md) | Null resets the options to their defaults. | [optional] 

## Methods

### NewUpdateTemplateRequest

`func NewUpdateTemplateRequest() *UpdateTemplateRequest`

NewUpdateTemplateRequest instantiates a new UpdateTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTemplateRequestWithDefaults

`func NewUpdateTemplateRequestWithDefaults() *UpdateTemplateRequest`

NewUpdateTemplateRequestWithDefaults instantiates a new UpdateTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateTemplateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateTemplateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentFolderId

`func (o *UpdateTemplateRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *UpdateTemplateRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *UpdateTemplateRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *UpdateTemplateRequest) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *UpdateTemplateRequest) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *UpdateTemplateRequest) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetLocked

`func (o *UpdateTemplateRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UpdateTemplateRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UpdateTemplateRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UpdateTemplateRequest) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetAddSignatureRequestSubmissionIdFooters

`func (o *UpdateTemplateRequest) GetAddSignatureRequestSubmissionIdFooters() bool`

GetAddSignatureRequestSubmissionIdFooters returns the AddSignatureRequestSubmissionIdFooters field if non-nil, zero value otherwise.

### GetAddSignatureRequestSubmissionIdFootersOk

`func (o *UpdateTemplateRequest) GetAddSignatureRequestSubmissionIdFootersOk() (*bool, bool)`

GetAddSignatureRequestSubmissionIdFootersOk returns a tuple with the AddSignatureRequestSubmissionIdFooters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSignatureRequestSubmissionIdFooters

`func (o *UpdateTemplateRequest) SetAddSignatureRequestSubmissionIdFooters(v bool)`

SetAddSignatureRequestSubmissionIdFooters sets AddSignatureRequestSubmissionIdFooters field to given value.

### HasAddSignatureRequestSubmissionIdFooters

`func (o *UpdateTemplateRequest) HasAddSignatureRequestSubmissionIdFooters() bool`

HasAddSignatureRequestSubmissionIdFooters returns a boolean if a field has been set.

### GetAiFieldNamingLanguage

`func (o *UpdateTemplateRequest) GetAiFieldNamingLanguage() AiFieldNamingLanguage`

GetAiFieldNamingLanguage returns the AiFieldNamingLanguage field if non-nil, zero value otherwise.

### GetAiFieldNamingLanguageOk

`func (o *UpdateTemplateRequest) GetAiFieldNamingLanguageOk() (*AiFieldNamingLanguage, bool)`

GetAiFieldNamingLanguageOk returns a tuple with the AiFieldNamingLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiFieldNamingLanguage

`func (o *UpdateTemplateRequest) SetAiFieldNamingLanguage(v AiFieldNamingLanguage)`

SetAiFieldNamingLanguage sets AiFieldNamingLanguage field to given value.

### HasAiFieldNamingLanguage

`func (o *UpdateTemplateRequest) HasAiFieldNamingLanguage() bool`

HasAiFieldNamingLanguage returns a boolean if a field has been set.

### GetAllowAdditionalProperties

`func (o *UpdateTemplateRequest) GetAllowAdditionalProperties() bool`

GetAllowAdditionalProperties returns the AllowAdditionalProperties field if non-nil, zero value otherwise.

### GetAllowAdditionalPropertiesOk

`func (o *UpdateTemplateRequest) GetAllowAdditionalPropertiesOk() (*bool, bool)`

GetAllowAdditionalPropertiesOk returns a tuple with the AllowAdditionalProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdditionalProperties

`func (o *UpdateTemplateRequest) SetAllowAdditionalProperties(v bool)`

SetAllowAdditionalProperties sets AllowAdditionalProperties field to given value.

### HasAllowAdditionalProperties

`func (o *UpdateTemplateRequest) HasAllowAdditionalProperties() bool`

HasAllowAdditionalProperties returns a boolean if a field has been set.

### GetAllowSkippingOptionalFields

`func (o *UpdateTemplateRequest) GetAllowSkippingOptionalFields() bool`

GetAllowSkippingOptionalFields returns the AllowSkippingOptionalFields field if non-nil, zero value otherwise.

### GetAllowSkippingOptionalFieldsOk

`func (o *UpdateTemplateRequest) GetAllowSkippingOptionalFieldsOk() (*bool, bool)`

GetAllowSkippingOptionalFieldsOk returns a tuple with the AllowSkippingOptionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSkippingOptionalFields

`func (o *UpdateTemplateRequest) SetAllowSkippingOptionalFields(v bool)`

SetAllowSkippingOptionalFields sets AllowSkippingOptionalFields field to given value.

### HasAllowSkippingOptionalFields

`func (o *UpdateTemplateRequest) HasAllowSkippingOptionalFields() bool`

HasAllowSkippingOptionalFields returns a boolean if a field has been set.

### GetDefaultFormMode

`func (o *UpdateTemplateRequest) GetDefaultFormMode() TemplateFormMode`

GetDefaultFormMode returns the DefaultFormMode field if non-nil, zero value otherwise.

### GetDefaultFormModeOk

`func (o *UpdateTemplateRequest) GetDefaultFormModeOk() (*TemplateFormMode, bool)`

GetDefaultFormModeOk returns a tuple with the DefaultFormMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFormMode

`func (o *UpdateTemplateRequest) SetDefaultFormMode(v TemplateFormMode)`

SetDefaultFormMode sets DefaultFormMode field to given value.

### HasDefaultFormMode

`func (o *UpdateTemplateRequest) HasDefaultFormMode() bool`

HasDefaultFormMode returns a boolean if a field has been set.

### GetEditableSubmissions

`func (o *UpdateTemplateRequest) GetEditableSubmissions() bool`

GetEditableSubmissions returns the EditableSubmissions field if non-nil, zero value otherwise.

### GetEditableSubmissionsOk

`func (o *UpdateTemplateRequest) GetEditableSubmissionsOk() (*bool, bool)`

GetEditableSubmissionsOk returns a tuple with the EditableSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableSubmissions

`func (o *UpdateTemplateRequest) SetEditableSubmissions(v bool)`

SetEditableSubmissions sets EditableSubmissions field to given value.

### HasEditableSubmissions

`func (o *UpdateTemplateRequest) HasEditableSubmissions() bool`

HasEditableSubmissions returns a boolean if a field has been set.

### GetEmbedDomains

`func (o *UpdateTemplateRequest) GetEmbedDomains() []string`

GetEmbedDomains returns the EmbedDomains field if non-nil, zero value otherwise.

### GetEmbedDomainsOk

`func (o *UpdateTemplateRequest) GetEmbedDomainsOk() (*[]string, bool)`

GetEmbedDomainsOk returns a tuple with the EmbedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedDomains

`func (o *UpdateTemplateRequest) SetEmbedDomains(v []string)`

SetEmbedDomains sets EmbedDomains field to given value.

### HasEmbedDomains

`func (o *UpdateTemplateRequest) HasEmbedDomains() bool`

HasEmbedDomains returns a boolean if a field has been set.

### GetEmbedEditAllowSettings

`func (o *UpdateTemplateRequest) GetEmbedEditAllowSettings() bool`

GetEmbedEditAllowSettings returns the EmbedEditAllowSettings field if non-nil, zero value otherwise.

### GetEmbedEditAllowSettingsOk

`func (o *UpdateTemplateRequest) GetEmbedEditAllowSettingsOk() (*bool, bool)`

GetEmbedEditAllowSettingsOk returns a tuple with the EmbedEditAllowSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowSettings

`func (o *UpdateTemplateRequest) SetEmbedEditAllowSettings(v bool)`

SetEmbedEditAllowSettings sets EmbedEditAllowSettings field to given value.

### HasEmbedEditAllowSettings

`func (o *UpdateTemplateRequest) HasEmbedEditAllowSettings() bool`

HasEmbedEditAllowSettings returns a boolean if a field has been set.

### GetEmbedEditAllowVersioning

`func (o *UpdateTemplateRequest) GetEmbedEditAllowVersioning() bool`

GetEmbedEditAllowVersioning returns the EmbedEditAllowVersioning field if non-nil, zero value otherwise.

### GetEmbedEditAllowVersioningOk

`func (o *UpdateTemplateRequest) GetEmbedEditAllowVersioningOk() (*bool, bool)`

GetEmbedEditAllowVersioningOk returns a tuple with the EmbedEditAllowVersioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowVersioning

`func (o *UpdateTemplateRequest) SetEmbedEditAllowVersioning(v bool)`

SetEmbedEditAllowVersioning sets EmbedEditAllowVersioning field to given value.

### HasEmbedEditAllowVersioning

`func (o *UpdateTemplateRequest) HasEmbedEditAllowVersioning() bool`

HasEmbedEditAllowVersioning returns a boolean if a field has been set.

### GetEmbedEditAllowDocumentReplacement

`func (o *UpdateTemplateRequest) GetEmbedEditAllowDocumentReplacement() bool`

GetEmbedEditAllowDocumentReplacement returns the EmbedEditAllowDocumentReplacement field if non-nil, zero value otherwise.

### GetEmbedEditAllowDocumentReplacementOk

`func (o *UpdateTemplateRequest) GetEmbedEditAllowDocumentReplacementOk() (*bool, bool)`

GetEmbedEditAllowDocumentReplacementOk returns a tuple with the EmbedEditAllowDocumentReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowDocumentReplacement

`func (o *UpdateTemplateRequest) SetEmbedEditAllowDocumentReplacement(v bool)`

SetEmbedEditAllowDocumentReplacement sets EmbedEditAllowDocumentReplacement field to given value.

### HasEmbedEditAllowDocumentReplacement

`func (o *UpdateTemplateRequest) HasEmbedEditAllowDocumentReplacement() bool`

HasEmbedEditAllowDocumentReplacement returns a boolean if a field has been set.

### GetEncryptPdfs

`func (o *UpdateTemplateRequest) GetEncryptPdfs() bool`

GetEncryptPdfs returns the EncryptPdfs field if non-nil, zero value otherwise.

### GetEncryptPdfsOk

`func (o *UpdateTemplateRequest) GetEncryptPdfsOk() (*bool, bool)`

GetEncryptPdfsOk returns a tuple with the EncryptPdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptPdfs

`func (o *UpdateTemplateRequest) SetEncryptPdfs(v bool)`

SetEncryptPdfs sets EncryptPdfs field to given value.

### HasEncryptPdfs

`func (o *UpdateTemplateRequest) HasEncryptPdfs() bool`

HasEncryptPdfs returns a boolean if a field has been set.

### GetEncryptPdfsPassword

`func (o *UpdateTemplateRequest) GetEncryptPdfsPassword() string`

GetEncryptPdfsPassword returns the EncryptPdfsPassword field if non-nil, zero value otherwise.

### GetEncryptPdfsPasswordOk

`func (o *UpdateTemplateRequest) GetEncryptPdfsPasswordOk() (*string, bool)`

GetEncryptPdfsPasswordOk returns a tuple with the EncryptPdfsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptPdfsPassword

`func (o *UpdateTemplateRequest) SetEncryptPdfsPassword(v string)`

SetEncryptPdfsPassword sets EncryptPdfsPassword field to given value.

### HasEncryptPdfsPassword

`func (o *UpdateTemplateRequest) HasEncryptPdfsPassword() bool`

HasEncryptPdfsPassword returns a boolean if a field has been set.

### SetEncryptPdfsPasswordNil

`func (o *UpdateTemplateRequest) SetEncryptPdfsPasswordNil(b bool)`

 SetEncryptPdfsPasswordNil sets the value for EncryptPdfsPassword to be an explicit nil

### UnsetEncryptPdfsPassword
`func (o *UpdateTemplateRequest) UnsetEncryptPdfsPassword()`

UnsetEncryptPdfsPassword ensures that no value is present for EncryptPdfsPassword, not even an explicit nil
### GetExpireSubmissions

`func (o *UpdateTemplateRequest) GetExpireSubmissions() bool`

GetExpireSubmissions returns the ExpireSubmissions field if non-nil, zero value otherwise.

### GetExpireSubmissionsOk

`func (o *UpdateTemplateRequest) GetExpireSubmissionsOk() (*bool, bool)`

GetExpireSubmissionsOk returns a tuple with the ExpireSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireSubmissions

`func (o *UpdateTemplateRequest) SetExpireSubmissions(v bool)`

SetExpireSubmissions sets ExpireSubmissions field to given value.

### HasExpireSubmissions

`func (o *UpdateTemplateRequest) HasExpireSubmissions() bool`

HasExpireSubmissions returns a boolean if a field has been set.

### GetExpireAfter

`func (o *UpdateTemplateRequest) GetExpireAfter() int32`

GetExpireAfter returns the ExpireAfter field if non-nil, zero value otherwise.

### GetExpireAfterOk

`func (o *UpdateTemplateRequest) GetExpireAfterOk() (*int32, bool)`

GetExpireAfterOk returns a tuple with the ExpireAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfter

`func (o *UpdateTemplateRequest) SetExpireAfter(v int32)`

SetExpireAfter sets ExpireAfter field to given value.

### HasExpireAfter

`func (o *UpdateTemplateRequest) HasExpireAfter() bool`

HasExpireAfter returns a boolean if a field has been set.

### GetExpirationInterval

`func (o *UpdateTemplateRequest) GetExpirationInterval() ExpirationInterval`

GetExpirationInterval returns the ExpirationInterval field if non-nil, zero value otherwise.

### GetExpirationIntervalOk

`func (o *UpdateTemplateRequest) GetExpirationIntervalOk() (*ExpirationInterval, bool)`

GetExpirationIntervalOk returns a tuple with the ExpirationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationInterval

`func (o *UpdateTemplateRequest) SetExpirationInterval(v ExpirationInterval)`

SetExpirationInterval sets ExpirationInterval field to given value.

### HasExpirationInterval

`func (o *UpdateTemplateRequest) HasExpirationInterval() bool`

HasExpirationInterval returns a boolean if a field has been set.

### GetMergeAuditTrailPdf

`func (o *UpdateTemplateRequest) GetMergeAuditTrailPdf() bool`

GetMergeAuditTrailPdf returns the MergeAuditTrailPdf field if non-nil, zero value otherwise.

### GetMergeAuditTrailPdfOk

`func (o *UpdateTemplateRequest) GetMergeAuditTrailPdfOk() (*bool, bool)`

GetMergeAuditTrailPdfOk returns a tuple with the MergeAuditTrailPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAuditTrailPdf

`func (o *UpdateTemplateRequest) SetMergeAuditTrailPdf(v bool)`

SetMergeAuditTrailPdf sets MergeAuditTrailPdf field to given value.

### HasMergeAuditTrailPdf

`func (o *UpdateTemplateRequest) HasMergeAuditTrailPdf() bool`

HasMergeAuditTrailPdf returns a boolean if a field has been set.

### GetParseDatesUsingTemplateDateFormat

`func (o *UpdateTemplateRequest) GetParseDatesUsingTemplateDateFormat() bool`

GetParseDatesUsingTemplateDateFormat returns the ParseDatesUsingTemplateDateFormat field if non-nil, zero value otherwise.

### GetParseDatesUsingTemplateDateFormatOk

`func (o *UpdateTemplateRequest) GetParseDatesUsingTemplateDateFormatOk() (*bool, bool)`

GetParseDatesUsingTemplateDateFormatOk returns a tuple with the ParseDatesUsingTemplateDateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseDatesUsingTemplateDateFormat

`func (o *UpdateTemplateRequest) SetParseDatesUsingTemplateDateFormat(v bool)`

SetParseDatesUsingTemplateDateFormat sets ParseDatesUsingTemplateDateFormat field to given value.

### HasParseDatesUsingTemplateDateFormat

`func (o *UpdateTemplateRequest) HasParseDatesUsingTemplateDateFormat() bool`

HasParseDatesUsingTemplateDateFormat returns a boolean if a field has been set.

### GetPublicSubmissions

`func (o *UpdateTemplateRequest) GetPublicSubmissions() bool`

GetPublicSubmissions returns the PublicSubmissions field if non-nil, zero value otherwise.

### GetPublicSubmissionsOk

`func (o *UpdateTemplateRequest) GetPublicSubmissionsOk() (*bool, bool)`

GetPublicSubmissionsOk returns a tuple with the PublicSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSubmissions

`func (o *UpdateTemplateRequest) SetPublicSubmissions(v bool)`

SetPublicSubmissions sets PublicSubmissions field to given value.

### HasPublicSubmissions

`func (o *UpdateTemplateRequest) HasPublicSubmissions() bool`

HasPublicSubmissions returns a boolean if a field has been set.

### GetPublicWebForm

`func (o *UpdateTemplateRequest) GetPublicWebForm() bool`

GetPublicWebForm returns the PublicWebForm field if non-nil, zero value otherwise.

### GetPublicWebFormOk

`func (o *UpdateTemplateRequest) GetPublicWebFormOk() (*bool, bool)`

GetPublicWebFormOk returns a tuple with the PublicWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicWebForm

`func (o *UpdateTemplateRequest) SetPublicWebForm(v bool)`

SetPublicWebForm sets PublicWebForm field to given value.

### HasPublicWebForm

`func (o *UpdateTemplateRequest) HasPublicWebForm() bool`

HasPublicWebForm returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *UpdateTemplateRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *UpdateTemplateRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *UpdateTemplateRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *UpdateTemplateRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *UpdateTemplateRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *UpdateTemplateRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetSlackWebhookUrl

`func (o *UpdateTemplateRequest) GetSlackWebhookUrl() string`

GetSlackWebhookUrl returns the SlackWebhookUrl field if non-nil, zero value otherwise.

### GetSlackWebhookUrlOk

`func (o *UpdateTemplateRequest) GetSlackWebhookUrlOk() (*string, bool)`

GetSlackWebhookUrlOk returns a tuple with the SlackWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhookUrl

`func (o *UpdateTemplateRequest) SetSlackWebhookUrl(v string)`

SetSlackWebhookUrl sets SlackWebhookUrl field to given value.

### HasSlackWebhookUrl

`func (o *UpdateTemplateRequest) HasSlackWebhookUrl() bool`

HasSlackWebhookUrl returns a boolean if a field has been set.

### SetSlackWebhookUrlNil

`func (o *UpdateTemplateRequest) SetSlackWebhookUrlNil(b bool)`

 SetSlackWebhookUrlNil sets the value for SlackWebhookUrl to be an explicit nil

### UnsetSlackWebhookUrl
`func (o *UpdateTemplateRequest) UnsetSlackWebhookUrl()`

UnsetSlackWebhookUrl ensures that no value is present for SlackWebhookUrl, not even an explicit nil
### GetDefaults

`func (o *UpdateTemplateRequest) GetDefaults() TemplateDefaultsInput`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *UpdateTemplateRequest) GetDefaultsOk() (*TemplateDefaultsInput, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *UpdateTemplateRequest) SetDefaults(v TemplateDefaultsInput)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *UpdateTemplateRequest) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetFieldDependencies

`func (o *UpdateTemplateRequest) GetFieldDependencies() []TemplateFieldDependencyInput`

GetFieldDependencies returns the FieldDependencies field if non-nil, zero value otherwise.

### GetFieldDependenciesOk

`func (o *UpdateTemplateRequest) GetFieldDependenciesOk() (*[]TemplateFieldDependencyInput, bool)`

GetFieldDependenciesOk returns a tuple with the FieldDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDependencies

`func (o *UpdateTemplateRequest) SetFieldDependencies(v []TemplateFieldDependencyInput)`

SetFieldDependencies sets FieldDependencies field to given value.

### HasFieldDependencies

`func (o *UpdateTemplateRequest) HasFieldDependencies() bool`

HasFieldDependencies returns a boolean if a field has been set.

### GetFieldGroups

`func (o *UpdateTemplateRequest) GetFieldGroups() []FieldGroupInput`

GetFieldGroups returns the FieldGroups field if non-nil, zero value otherwise.

### GetFieldGroupsOk

`func (o *UpdateTemplateRequest) GetFieldGroupsOk() (*[]FieldGroupInput, bool)`

GetFieldGroupsOk returns a tuple with the FieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroups

`func (o *UpdateTemplateRequest) SetFieldGroups(v []FieldGroupInput)`

SetFieldGroups sets FieldGroups field to given value.

### HasFieldGroups

`func (o *UpdateTemplateRequest) HasFieldGroups() bool`

HasFieldGroups returns a boolean if a field has been set.

### GetPredefinedFields

`func (o *UpdateTemplateRequest) GetPredefinedFields() []PredefinedFieldInput`

GetPredefinedFields returns the PredefinedFields field if non-nil, zero value otherwise.

### GetPredefinedFieldsOk

`func (o *UpdateTemplateRequest) GetPredefinedFieldsOk() (*[]PredefinedFieldInput, bool)`

GetPredefinedFieldsOk returns a tuple with the PredefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedFields

`func (o *UpdateTemplateRequest) SetPredefinedFields(v []PredefinedFieldInput)`

SetPredefinedFields sets PredefinedFields field to given value.

### HasPredefinedFields

`func (o *UpdateTemplateRequest) HasPredefinedFields() bool`

HasPredefinedFields returns a boolean if a field has been set.

### GetPredefinedFieldsSampleData

`func (o *UpdateTemplateRequest) GetPredefinedFieldsSampleData() map[string]interface{}`

GetPredefinedFieldsSampleData returns the PredefinedFieldsSampleData field if non-nil, zero value otherwise.

### GetPredefinedFieldsSampleDataOk

`func (o *UpdateTemplateRequest) GetPredefinedFieldsSampleDataOk() (*map[string]interface{}, bool)`

GetPredefinedFieldsSampleDataOk returns a tuple with the PredefinedFieldsSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedFieldsSampleData

`func (o *UpdateTemplateRequest) SetPredefinedFieldsSampleData(v map[string]interface{})`

SetPredefinedFieldsSampleData sets PredefinedFieldsSampleData field to given value.

### HasPredefinedFieldsSampleData

`func (o *UpdateTemplateRequest) HasPredefinedFieldsSampleData() bool`

HasPredefinedFieldsSampleData returns a boolean if a field has been set.

### SetPredefinedFieldsSampleDataNil

`func (o *UpdateTemplateRequest) SetPredefinedFieldsSampleDataNil(b bool)`

 SetPredefinedFieldsSampleDataNil sets the value for PredefinedFieldsSampleData to be an explicit nil

### UnsetPredefinedFieldsSampleData
`func (o *UpdateTemplateRequest) UnsetPredefinedFieldsSampleData()`

UnsetPredefinedFieldsSampleData ensures that no value is present for PredefinedFieldsSampleData, not even an explicit nil
### GetHtml

`func (o *UpdateTemplateRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *UpdateTemplateRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *UpdateTemplateRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *UpdateTemplateRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### SetHtmlNil

`func (o *UpdateTemplateRequest) SetHtmlNil(b bool)`

 SetHtmlNil sets the value for Html to be an explicit nil

### UnsetHtml
`func (o *UpdateTemplateRequest) UnsetHtml()`

UnsetHtml ensures that no value is present for Html, not even an explicit nil
### GetScss

`func (o *UpdateTemplateRequest) GetScss() string`

GetScss returns the Scss field if non-nil, zero value otherwise.

### GetScssOk

`func (o *UpdateTemplateRequest) GetScssOk() (*string, bool)`

GetScssOk returns a tuple with the Scss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScss

`func (o *UpdateTemplateRequest) SetScss(v string)`

SetScss sets Scss field to given value.

### HasScss

`func (o *UpdateTemplateRequest) HasScss() bool`

HasScss returns a boolean if a field has been set.

### SetScssNil

`func (o *UpdateTemplateRequest) SetScssNil(b bool)`

 SetScssNil sets the value for Scss to be an explicit nil

### UnsetScss
`func (o *UpdateTemplateRequest) UnsetScss()`

UnsetScss ensures that no value is present for Scss, not even an explicit nil
### GetHeaderHtml

`func (o *UpdateTemplateRequest) GetHeaderHtml() string`

GetHeaderHtml returns the HeaderHtml field if non-nil, zero value otherwise.

### GetHeaderHtmlOk

`func (o *UpdateTemplateRequest) GetHeaderHtmlOk() (*string, bool)`

GetHeaderHtmlOk returns a tuple with the HeaderHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderHtml

`func (o *UpdateTemplateRequest) SetHeaderHtml(v string)`

SetHeaderHtml sets HeaderHtml field to given value.

### HasHeaderHtml

`func (o *UpdateTemplateRequest) HasHeaderHtml() bool`

HasHeaderHtml returns a boolean if a field has been set.

### SetHeaderHtmlNil

`func (o *UpdateTemplateRequest) SetHeaderHtmlNil(b bool)`

 SetHeaderHtmlNil sets the value for HeaderHtml to be an explicit nil

### UnsetHeaderHtml
`func (o *UpdateTemplateRequest) UnsetHeaderHtml()`

UnsetHeaderHtml ensures that no value is present for HeaderHtml, not even an explicit nil
### GetFooterHtml

`func (o *UpdateTemplateRequest) GetFooterHtml() string`

GetFooterHtml returns the FooterHtml field if non-nil, zero value otherwise.

### GetFooterHtmlOk

`func (o *UpdateTemplateRequest) GetFooterHtmlOk() (*string, bool)`

GetFooterHtmlOk returns a tuple with the FooterHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterHtml

`func (o *UpdateTemplateRequest) SetFooterHtml(v string)`

SetFooterHtml sets FooterHtml field to given value.

### HasFooterHtml

`func (o *UpdateTemplateRequest) HasFooterHtml() bool`

HasFooterHtml returns a boolean if a field has been set.

### SetFooterHtmlNil

`func (o *UpdateTemplateRequest) SetFooterHtmlNil(b bool)`

 SetFooterHtmlNil sets the value for FooterHtml to be an explicit nil

### UnsetFooterHtml
`func (o *UpdateTemplateRequest) UnsetFooterHtml()`

UnsetFooterHtml ensures that no value is present for FooterHtml, not even an explicit nil
### GetHtmlEngineOptions

`func (o *UpdateTemplateRequest) GetHtmlEngineOptions() HtmlEngineOptionsInput`

GetHtmlEngineOptions returns the HtmlEngineOptions field if non-nil, zero value otherwise.

### GetHtmlEngineOptionsOk

`func (o *UpdateTemplateRequest) GetHtmlEngineOptionsOk() (*HtmlEngineOptionsInput, bool)`

GetHtmlEngineOptionsOk returns a tuple with the HtmlEngineOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlEngineOptions

`func (o *UpdateTemplateRequest) SetHtmlEngineOptions(v HtmlEngineOptionsInput)`

SetHtmlEngineOptions sets HtmlEngineOptions field to given value.

### HasHtmlEngineOptions

`func (o *UpdateTemplateRequest) HasHtmlEngineOptions() bool`

HasHtmlEngineOptions returns a boolean if a field has been set.

### SetHtmlEngineOptionsNil

`func (o *UpdateTemplateRequest) SetHtmlEngineOptionsNil(b bool)`

 SetHtmlEngineOptionsNil sets the value for HtmlEngineOptions to be an explicit nil

### UnsetHtmlEngineOptions
`func (o *UpdateTemplateRequest) UnsetHtmlEngineOptions()`

UnsetHtmlEngineOptions ensures that no value is present for HtmlEngineOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


