# CreateTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The template's name. | [optional] 
**Description** | Pointer to **NullableString** | Shown under the template's name at the top of its hosted forms, and used as their meta description. Not shown on embedded forms or signature requests. | [optional] 
**TemplateType** | [**TemplateType**](TemplateType.md) | `pdf`: fields placed on an uploaded PDF. `html`: a PDF rendered from HTML and SCSS. It can't be changed after the template is created. | 
**ParentFolderId** | Pointer to **NullableString** | The ID of the folder the template is in. Null for a template at the root. Setting it moves the template and all of its published versions. | [optional] 
**Locked** | Pointer to **bool** | A locked template can't be changed or deleted until it's unlocked (`locked` can always be changed). | [optional] 
**AddSignatureRequestSubmissionIdFooters** | Pointer to **bool** | Print the submission ID in the footer of each page of PDFs signed through signature requests. | [optional] 
**AiFieldNamingLanguage** | Pointer to [**AiFieldNamingLanguage**](AiFieldNamingLanguage.md) | The language AI field setup names fields in: `auto` uses the form's language when it's written in the Latin alphabet without accents, and English otherwise; `english` always uses English; `form_language` uses the form's language and script. New templates take it from your account's template settings. Null for HTML templates. | [optional] 
**AllowAdditionalProperties** | Pointer to **bool** | Accept submission data with properties that aren't fields of the template. | [optional] 
**AllowSkippingOptionalFields** | Pointer to **bool** | Visual forms and signature request forms show the Continue button as soon as the required fields are filled in, without visiting each optional field. The `allowSkippingOptionalFields` embed option overrides it. | [optional] 
**DefaultFormMode** | Pointer to [**TemplateFormMode**](TemplateFormMode.md) | The hosted form that /templates/{template_id}/form shows when its URL has no `mode` parameter: `web` (the web form), `visual` (the visual form, where fields are filled in on the PDF) or `dual` (both, in tabs). HTML templates have no visual form, so `web` is the only value they accept. | [optional] 
**EditableSubmissions** | Pointer to **bool** | Generate PDFs whose text and check fields stay fillable ("Generate Editable PDFs"). A submission's `editable` overrides it. | [optional] 
**EmbedDomains** | Pointer to **[]string** | The sites that can embed the template's live signature requests and its embedded template editor: a host such as `app.example.com`, with an optional leading wildcard (`*.example.com`), port (`localhost:3000`) or scheme (`https://app.example.com`; without one, `http` and `https` both work). Browsers open a live signature request in a frame only on these sites (and DocSpring's), so with an empty list it can't be embedded anywhere; a test signature request can be embedded on any site. The embedded template editor loads only for a page on one of these sites (by its Referer or Origin header), and on any site when the list is empty. Forms embedded on a page (Visual Forms JS and the simple form) aren't restricted. | [optional] 
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
**PredefinedFields** | Pointer to [**[]PredefinedFieldInput**](PredefinedFieldInput.md) | The field names that people using the embedded template editor choose from, each with the data type and settings a field gets when it's given that name. Setting it replaces the whole list; an empty list turns them off. | [optional] 
**PredefinedFieldsSampleData** | Pointer to **map[string]interface{}** | Sample submission data that sets `predefined_fields`: each value becomes a predefined field, with a type that fits it (`string`, `number`, `boolean`, `date`, `email` or `url`) and `required_in_template` true. Nested objects and arrays become names with `/`, such as `customer/email`. Without `predefined_fields` in the same request, it replaces them; with them, it adds to them and sets the types of the ones it has values for. | [optional] 
**Html** | Pointer to **NullableString** | The HTML of an HTML template, as a Liquid template. Null for PDF templates. | [optional] 
**Scss** | Pointer to **NullableString** | The SCSS of an HTML template. | [optional] 
**HeaderHtml** | Pointer to **NullableString** | The HTML printed at the top of each page of an HTML template. | [optional] 
**FooterHtml** | Pointer to **NullableString** | The HTML printed at the bottom of each page of an HTML template. | [optional] 
**HtmlEngineOptions** | Pointer to [**NullableHtmlEngineOptionsInput**](HtmlEngineOptionsInput.md) | Null resets the options to their defaults. | [optional] 
**HtmlEscapeData** | Pointer to **bool** | Whether an HTML template HTML-escapes the values it prints. When true, `{{ value }}` prints a value as text, so a `&lt;` or `&amp;` in it is printed as that character, and `{{ value | raw }}` prints a value as HTML. New HTML templates default to true. Templates created before October 2026, or with API v1, default to false, which prints every value as HTML. Null for PDF templates. | [optional] 
**CacheId** | Pointer to **string** | A PDF template's document: the `fields.key` of a file you uploaded with the presigned upload (GET /uploads/presign). A PDF template needs one of `cache_id`, `document_url` and `document_base64`. To upload the PDF with the request, use POST /templates/upload. | [optional] 
**DocumentUrl** | Pointer to **string** | An https:// link to the template's PDF (up to 100 MB), instead of `cache_id`. DocSpring downloads it after responding, following redirects to other https:// links: the template's `document.state` is `downloading` until then (sync.api hosts wait for it). If it can't be downloaded, the state is `download_failed` and `document.error_message` says why. | [optional] 
**DocumentBase64** | Pointer to **string** | The template's PDF, encoded as base64 (a `data:` URI works too), instead of `cache_id`. Up to 7 MB before it's encoded: send a larger file with `cache_id` or `document_url`. | [optional] 
**Filename** | Pointer to **string** | The PDF's file name, for a `cache_id` (presigned uploads don't send one) or `document_base64`. A `document_url`'s is the last part of its path unless you set it. | [optional] 
**DetectFields** | Pointer to **bool** | Set up the PDF's fields as DocSpring's "Automatically detect fields" upload option does: import or detect the fields, name them with AI, and set up conditional logic. `field_source`, `field_naming` and `field_logic` choose how. Default false. `document.page_processing_statuses` and `document.field_logic_status` show how it's going, and `document.state` becomes `processed` once every page's fields are set up. | [optional] 
**FieldSource** | Pointer to [**FieldSetupSource**](FieldSetupSource.md) | With `detect_fields`, where the fields come from. `auto` (the default) imports the form fields of a fillable PDF and detects the fields of a PDF without them; `pdf_detection` detects the fields from the page layout, even on a fillable PDF; `pdf_form_fields` only imports the PDF's form fields. | [optional] 
**FieldNaming** | Pointer to [**FieldSetupNaming**](FieldSetupNaming.md) | With `detect_fields`, how the fields are named. `ai` (the default) names them and sets their types and settings with AI; `preserve` keeps the names of the PDF's form fields, and AI only names the fields it detected. `preserve` needs `field_source` `auto` or `pdf_form_fields`. | [optional] 
**FieldLogic** | Pointer to [**FieldSetupLogic**](FieldSetupLogic.md) | With `detect_fields`, whether AI reads the whole form once the fields are named and sets up conditional logic and required fields (`ai`) or not (`none`). Default: your account's template settings. | [optional] 

## Methods

### NewCreateTemplateRequest

`func NewCreateTemplateRequest(templateType TemplateType, ) *CreateTemplateRequest`

NewCreateTemplateRequest instantiates a new CreateTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTemplateRequestWithDefaults

`func NewCreateTemplateRequestWithDefaults() *CreateTemplateRequest`

NewCreateTemplateRequestWithDefaults instantiates a new CreateTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTemplateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTemplateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTemplateType

`func (o *CreateTemplateRequest) GetTemplateType() TemplateType`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *CreateTemplateRequest) GetTemplateTypeOk() (*TemplateType, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *CreateTemplateRequest) SetTemplateType(v TemplateType)`

SetTemplateType sets TemplateType field to given value.


### GetParentFolderId

`func (o *CreateTemplateRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *CreateTemplateRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *CreateTemplateRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *CreateTemplateRequest) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *CreateTemplateRequest) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *CreateTemplateRequest) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetLocked

`func (o *CreateTemplateRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CreateTemplateRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CreateTemplateRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *CreateTemplateRequest) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetAddSignatureRequestSubmissionIdFooters

`func (o *CreateTemplateRequest) GetAddSignatureRequestSubmissionIdFooters() bool`

GetAddSignatureRequestSubmissionIdFooters returns the AddSignatureRequestSubmissionIdFooters field if non-nil, zero value otherwise.

### GetAddSignatureRequestSubmissionIdFootersOk

`func (o *CreateTemplateRequest) GetAddSignatureRequestSubmissionIdFootersOk() (*bool, bool)`

GetAddSignatureRequestSubmissionIdFootersOk returns a tuple with the AddSignatureRequestSubmissionIdFooters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSignatureRequestSubmissionIdFooters

`func (o *CreateTemplateRequest) SetAddSignatureRequestSubmissionIdFooters(v bool)`

SetAddSignatureRequestSubmissionIdFooters sets AddSignatureRequestSubmissionIdFooters field to given value.

### HasAddSignatureRequestSubmissionIdFooters

`func (o *CreateTemplateRequest) HasAddSignatureRequestSubmissionIdFooters() bool`

HasAddSignatureRequestSubmissionIdFooters returns a boolean if a field has been set.

### GetAiFieldNamingLanguage

`func (o *CreateTemplateRequest) GetAiFieldNamingLanguage() AiFieldNamingLanguage`

GetAiFieldNamingLanguage returns the AiFieldNamingLanguage field if non-nil, zero value otherwise.

### GetAiFieldNamingLanguageOk

`func (o *CreateTemplateRequest) GetAiFieldNamingLanguageOk() (*AiFieldNamingLanguage, bool)`

GetAiFieldNamingLanguageOk returns a tuple with the AiFieldNamingLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiFieldNamingLanguage

`func (o *CreateTemplateRequest) SetAiFieldNamingLanguage(v AiFieldNamingLanguage)`

SetAiFieldNamingLanguage sets AiFieldNamingLanguage field to given value.

### HasAiFieldNamingLanguage

`func (o *CreateTemplateRequest) HasAiFieldNamingLanguage() bool`

HasAiFieldNamingLanguage returns a boolean if a field has been set.

### GetAllowAdditionalProperties

`func (o *CreateTemplateRequest) GetAllowAdditionalProperties() bool`

GetAllowAdditionalProperties returns the AllowAdditionalProperties field if non-nil, zero value otherwise.

### GetAllowAdditionalPropertiesOk

`func (o *CreateTemplateRequest) GetAllowAdditionalPropertiesOk() (*bool, bool)`

GetAllowAdditionalPropertiesOk returns a tuple with the AllowAdditionalProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdditionalProperties

`func (o *CreateTemplateRequest) SetAllowAdditionalProperties(v bool)`

SetAllowAdditionalProperties sets AllowAdditionalProperties field to given value.

### HasAllowAdditionalProperties

`func (o *CreateTemplateRequest) HasAllowAdditionalProperties() bool`

HasAllowAdditionalProperties returns a boolean if a field has been set.

### GetAllowSkippingOptionalFields

`func (o *CreateTemplateRequest) GetAllowSkippingOptionalFields() bool`

GetAllowSkippingOptionalFields returns the AllowSkippingOptionalFields field if non-nil, zero value otherwise.

### GetAllowSkippingOptionalFieldsOk

`func (o *CreateTemplateRequest) GetAllowSkippingOptionalFieldsOk() (*bool, bool)`

GetAllowSkippingOptionalFieldsOk returns a tuple with the AllowSkippingOptionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSkippingOptionalFields

`func (o *CreateTemplateRequest) SetAllowSkippingOptionalFields(v bool)`

SetAllowSkippingOptionalFields sets AllowSkippingOptionalFields field to given value.

### HasAllowSkippingOptionalFields

`func (o *CreateTemplateRequest) HasAllowSkippingOptionalFields() bool`

HasAllowSkippingOptionalFields returns a boolean if a field has been set.

### GetDefaultFormMode

`func (o *CreateTemplateRequest) GetDefaultFormMode() TemplateFormMode`

GetDefaultFormMode returns the DefaultFormMode field if non-nil, zero value otherwise.

### GetDefaultFormModeOk

`func (o *CreateTemplateRequest) GetDefaultFormModeOk() (*TemplateFormMode, bool)`

GetDefaultFormModeOk returns a tuple with the DefaultFormMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFormMode

`func (o *CreateTemplateRequest) SetDefaultFormMode(v TemplateFormMode)`

SetDefaultFormMode sets DefaultFormMode field to given value.

### HasDefaultFormMode

`func (o *CreateTemplateRequest) HasDefaultFormMode() bool`

HasDefaultFormMode returns a boolean if a field has been set.

### GetEditableSubmissions

`func (o *CreateTemplateRequest) GetEditableSubmissions() bool`

GetEditableSubmissions returns the EditableSubmissions field if non-nil, zero value otherwise.

### GetEditableSubmissionsOk

`func (o *CreateTemplateRequest) GetEditableSubmissionsOk() (*bool, bool)`

GetEditableSubmissionsOk returns a tuple with the EditableSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableSubmissions

`func (o *CreateTemplateRequest) SetEditableSubmissions(v bool)`

SetEditableSubmissions sets EditableSubmissions field to given value.

### HasEditableSubmissions

`func (o *CreateTemplateRequest) HasEditableSubmissions() bool`

HasEditableSubmissions returns a boolean if a field has been set.

### GetEmbedDomains

`func (o *CreateTemplateRequest) GetEmbedDomains() []string`

GetEmbedDomains returns the EmbedDomains field if non-nil, zero value otherwise.

### GetEmbedDomainsOk

`func (o *CreateTemplateRequest) GetEmbedDomainsOk() (*[]string, bool)`

GetEmbedDomainsOk returns a tuple with the EmbedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedDomains

`func (o *CreateTemplateRequest) SetEmbedDomains(v []string)`

SetEmbedDomains sets EmbedDomains field to given value.

### HasEmbedDomains

`func (o *CreateTemplateRequest) HasEmbedDomains() bool`

HasEmbedDomains returns a boolean if a field has been set.

### GetEmbedEditAllowSettings

`func (o *CreateTemplateRequest) GetEmbedEditAllowSettings() bool`

GetEmbedEditAllowSettings returns the EmbedEditAllowSettings field if non-nil, zero value otherwise.

### GetEmbedEditAllowSettingsOk

`func (o *CreateTemplateRequest) GetEmbedEditAllowSettingsOk() (*bool, bool)`

GetEmbedEditAllowSettingsOk returns a tuple with the EmbedEditAllowSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowSettings

`func (o *CreateTemplateRequest) SetEmbedEditAllowSettings(v bool)`

SetEmbedEditAllowSettings sets EmbedEditAllowSettings field to given value.

### HasEmbedEditAllowSettings

`func (o *CreateTemplateRequest) HasEmbedEditAllowSettings() bool`

HasEmbedEditAllowSettings returns a boolean if a field has been set.

### GetEmbedEditAllowVersioning

`func (o *CreateTemplateRequest) GetEmbedEditAllowVersioning() bool`

GetEmbedEditAllowVersioning returns the EmbedEditAllowVersioning field if non-nil, zero value otherwise.

### GetEmbedEditAllowVersioningOk

`func (o *CreateTemplateRequest) GetEmbedEditAllowVersioningOk() (*bool, bool)`

GetEmbedEditAllowVersioningOk returns a tuple with the EmbedEditAllowVersioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowVersioning

`func (o *CreateTemplateRequest) SetEmbedEditAllowVersioning(v bool)`

SetEmbedEditAllowVersioning sets EmbedEditAllowVersioning field to given value.

### HasEmbedEditAllowVersioning

`func (o *CreateTemplateRequest) HasEmbedEditAllowVersioning() bool`

HasEmbedEditAllowVersioning returns a boolean if a field has been set.

### GetEmbedEditAllowDocumentReplacement

`func (o *CreateTemplateRequest) GetEmbedEditAllowDocumentReplacement() bool`

GetEmbedEditAllowDocumentReplacement returns the EmbedEditAllowDocumentReplacement field if non-nil, zero value otherwise.

### GetEmbedEditAllowDocumentReplacementOk

`func (o *CreateTemplateRequest) GetEmbedEditAllowDocumentReplacementOk() (*bool, bool)`

GetEmbedEditAllowDocumentReplacementOk returns a tuple with the EmbedEditAllowDocumentReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedEditAllowDocumentReplacement

`func (o *CreateTemplateRequest) SetEmbedEditAllowDocumentReplacement(v bool)`

SetEmbedEditAllowDocumentReplacement sets EmbedEditAllowDocumentReplacement field to given value.

### HasEmbedEditAllowDocumentReplacement

`func (o *CreateTemplateRequest) HasEmbedEditAllowDocumentReplacement() bool`

HasEmbedEditAllowDocumentReplacement returns a boolean if a field has been set.

### GetEncryptPdfs

`func (o *CreateTemplateRequest) GetEncryptPdfs() bool`

GetEncryptPdfs returns the EncryptPdfs field if non-nil, zero value otherwise.

### GetEncryptPdfsOk

`func (o *CreateTemplateRequest) GetEncryptPdfsOk() (*bool, bool)`

GetEncryptPdfsOk returns a tuple with the EncryptPdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptPdfs

`func (o *CreateTemplateRequest) SetEncryptPdfs(v bool)`

SetEncryptPdfs sets EncryptPdfs field to given value.

### HasEncryptPdfs

`func (o *CreateTemplateRequest) HasEncryptPdfs() bool`

HasEncryptPdfs returns a boolean if a field has been set.

### GetEncryptPdfsPassword

`func (o *CreateTemplateRequest) GetEncryptPdfsPassword() string`

GetEncryptPdfsPassword returns the EncryptPdfsPassword field if non-nil, zero value otherwise.

### GetEncryptPdfsPasswordOk

`func (o *CreateTemplateRequest) GetEncryptPdfsPasswordOk() (*string, bool)`

GetEncryptPdfsPasswordOk returns a tuple with the EncryptPdfsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptPdfsPassword

`func (o *CreateTemplateRequest) SetEncryptPdfsPassword(v string)`

SetEncryptPdfsPassword sets EncryptPdfsPassword field to given value.

### HasEncryptPdfsPassword

`func (o *CreateTemplateRequest) HasEncryptPdfsPassword() bool`

HasEncryptPdfsPassword returns a boolean if a field has been set.

### SetEncryptPdfsPasswordNil

`func (o *CreateTemplateRequest) SetEncryptPdfsPasswordNil(b bool)`

 SetEncryptPdfsPasswordNil sets the value for EncryptPdfsPassword to be an explicit nil

### UnsetEncryptPdfsPassword
`func (o *CreateTemplateRequest) UnsetEncryptPdfsPassword()`

UnsetEncryptPdfsPassword ensures that no value is present for EncryptPdfsPassword, not even an explicit nil
### GetExpireSubmissions

`func (o *CreateTemplateRequest) GetExpireSubmissions() bool`

GetExpireSubmissions returns the ExpireSubmissions field if non-nil, zero value otherwise.

### GetExpireSubmissionsOk

`func (o *CreateTemplateRequest) GetExpireSubmissionsOk() (*bool, bool)`

GetExpireSubmissionsOk returns a tuple with the ExpireSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireSubmissions

`func (o *CreateTemplateRequest) SetExpireSubmissions(v bool)`

SetExpireSubmissions sets ExpireSubmissions field to given value.

### HasExpireSubmissions

`func (o *CreateTemplateRequest) HasExpireSubmissions() bool`

HasExpireSubmissions returns a boolean if a field has been set.

### GetExpireAfter

`func (o *CreateTemplateRequest) GetExpireAfter() int32`

GetExpireAfter returns the ExpireAfter field if non-nil, zero value otherwise.

### GetExpireAfterOk

`func (o *CreateTemplateRequest) GetExpireAfterOk() (*int32, bool)`

GetExpireAfterOk returns a tuple with the ExpireAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfter

`func (o *CreateTemplateRequest) SetExpireAfter(v int32)`

SetExpireAfter sets ExpireAfter field to given value.

### HasExpireAfter

`func (o *CreateTemplateRequest) HasExpireAfter() bool`

HasExpireAfter returns a boolean if a field has been set.

### GetExpirationInterval

`func (o *CreateTemplateRequest) GetExpirationInterval() ExpirationInterval`

GetExpirationInterval returns the ExpirationInterval field if non-nil, zero value otherwise.

### GetExpirationIntervalOk

`func (o *CreateTemplateRequest) GetExpirationIntervalOk() (*ExpirationInterval, bool)`

GetExpirationIntervalOk returns a tuple with the ExpirationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationInterval

`func (o *CreateTemplateRequest) SetExpirationInterval(v ExpirationInterval)`

SetExpirationInterval sets ExpirationInterval field to given value.

### HasExpirationInterval

`func (o *CreateTemplateRequest) HasExpirationInterval() bool`

HasExpirationInterval returns a boolean if a field has been set.

### GetMergeAuditTrailPdf

`func (o *CreateTemplateRequest) GetMergeAuditTrailPdf() bool`

GetMergeAuditTrailPdf returns the MergeAuditTrailPdf field if non-nil, zero value otherwise.

### GetMergeAuditTrailPdfOk

`func (o *CreateTemplateRequest) GetMergeAuditTrailPdfOk() (*bool, bool)`

GetMergeAuditTrailPdfOk returns a tuple with the MergeAuditTrailPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAuditTrailPdf

`func (o *CreateTemplateRequest) SetMergeAuditTrailPdf(v bool)`

SetMergeAuditTrailPdf sets MergeAuditTrailPdf field to given value.

### HasMergeAuditTrailPdf

`func (o *CreateTemplateRequest) HasMergeAuditTrailPdf() bool`

HasMergeAuditTrailPdf returns a boolean if a field has been set.

### GetParseDatesUsingTemplateDateFormat

`func (o *CreateTemplateRequest) GetParseDatesUsingTemplateDateFormat() bool`

GetParseDatesUsingTemplateDateFormat returns the ParseDatesUsingTemplateDateFormat field if non-nil, zero value otherwise.

### GetParseDatesUsingTemplateDateFormatOk

`func (o *CreateTemplateRequest) GetParseDatesUsingTemplateDateFormatOk() (*bool, bool)`

GetParseDatesUsingTemplateDateFormatOk returns a tuple with the ParseDatesUsingTemplateDateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseDatesUsingTemplateDateFormat

`func (o *CreateTemplateRequest) SetParseDatesUsingTemplateDateFormat(v bool)`

SetParseDatesUsingTemplateDateFormat sets ParseDatesUsingTemplateDateFormat field to given value.

### HasParseDatesUsingTemplateDateFormat

`func (o *CreateTemplateRequest) HasParseDatesUsingTemplateDateFormat() bool`

HasParseDatesUsingTemplateDateFormat returns a boolean if a field has been set.

### GetPublicSubmissions

`func (o *CreateTemplateRequest) GetPublicSubmissions() bool`

GetPublicSubmissions returns the PublicSubmissions field if non-nil, zero value otherwise.

### GetPublicSubmissionsOk

`func (o *CreateTemplateRequest) GetPublicSubmissionsOk() (*bool, bool)`

GetPublicSubmissionsOk returns a tuple with the PublicSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSubmissions

`func (o *CreateTemplateRequest) SetPublicSubmissions(v bool)`

SetPublicSubmissions sets PublicSubmissions field to given value.

### HasPublicSubmissions

`func (o *CreateTemplateRequest) HasPublicSubmissions() bool`

HasPublicSubmissions returns a boolean if a field has been set.

### GetPublicWebForm

`func (o *CreateTemplateRequest) GetPublicWebForm() bool`

GetPublicWebForm returns the PublicWebForm field if non-nil, zero value otherwise.

### GetPublicWebFormOk

`func (o *CreateTemplateRequest) GetPublicWebFormOk() (*bool, bool)`

GetPublicWebFormOk returns a tuple with the PublicWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicWebForm

`func (o *CreateTemplateRequest) SetPublicWebForm(v bool)`

SetPublicWebForm sets PublicWebForm field to given value.

### HasPublicWebForm

`func (o *CreateTemplateRequest) HasPublicWebForm() bool`

HasPublicWebForm returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *CreateTemplateRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreateTemplateRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreateTemplateRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CreateTemplateRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *CreateTemplateRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *CreateTemplateRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetSlackWebhookUrl

`func (o *CreateTemplateRequest) GetSlackWebhookUrl() string`

GetSlackWebhookUrl returns the SlackWebhookUrl field if non-nil, zero value otherwise.

### GetSlackWebhookUrlOk

`func (o *CreateTemplateRequest) GetSlackWebhookUrlOk() (*string, bool)`

GetSlackWebhookUrlOk returns a tuple with the SlackWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhookUrl

`func (o *CreateTemplateRequest) SetSlackWebhookUrl(v string)`

SetSlackWebhookUrl sets SlackWebhookUrl field to given value.

### HasSlackWebhookUrl

`func (o *CreateTemplateRequest) HasSlackWebhookUrl() bool`

HasSlackWebhookUrl returns a boolean if a field has been set.

### SetSlackWebhookUrlNil

`func (o *CreateTemplateRequest) SetSlackWebhookUrlNil(b bool)`

 SetSlackWebhookUrlNil sets the value for SlackWebhookUrl to be an explicit nil

### UnsetSlackWebhookUrl
`func (o *CreateTemplateRequest) UnsetSlackWebhookUrl()`

UnsetSlackWebhookUrl ensures that no value is present for SlackWebhookUrl, not even an explicit nil
### GetDefaults

`func (o *CreateTemplateRequest) GetDefaults() TemplateDefaultsInput`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *CreateTemplateRequest) GetDefaultsOk() (*TemplateDefaultsInput, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *CreateTemplateRequest) SetDefaults(v TemplateDefaultsInput)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *CreateTemplateRequest) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetFieldDependencies

`func (o *CreateTemplateRequest) GetFieldDependencies() []TemplateFieldDependencyInput`

GetFieldDependencies returns the FieldDependencies field if non-nil, zero value otherwise.

### GetFieldDependenciesOk

`func (o *CreateTemplateRequest) GetFieldDependenciesOk() (*[]TemplateFieldDependencyInput, bool)`

GetFieldDependenciesOk returns a tuple with the FieldDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDependencies

`func (o *CreateTemplateRequest) SetFieldDependencies(v []TemplateFieldDependencyInput)`

SetFieldDependencies sets FieldDependencies field to given value.

### HasFieldDependencies

`func (o *CreateTemplateRequest) HasFieldDependencies() bool`

HasFieldDependencies returns a boolean if a field has been set.

### GetPredefinedFields

`func (o *CreateTemplateRequest) GetPredefinedFields() []PredefinedFieldInput`

GetPredefinedFields returns the PredefinedFields field if non-nil, zero value otherwise.

### GetPredefinedFieldsOk

`func (o *CreateTemplateRequest) GetPredefinedFieldsOk() (*[]PredefinedFieldInput, bool)`

GetPredefinedFieldsOk returns a tuple with the PredefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedFields

`func (o *CreateTemplateRequest) SetPredefinedFields(v []PredefinedFieldInput)`

SetPredefinedFields sets PredefinedFields field to given value.

### HasPredefinedFields

`func (o *CreateTemplateRequest) HasPredefinedFields() bool`

HasPredefinedFields returns a boolean if a field has been set.

### GetPredefinedFieldsSampleData

`func (o *CreateTemplateRequest) GetPredefinedFieldsSampleData() map[string]interface{}`

GetPredefinedFieldsSampleData returns the PredefinedFieldsSampleData field if non-nil, zero value otherwise.

### GetPredefinedFieldsSampleDataOk

`func (o *CreateTemplateRequest) GetPredefinedFieldsSampleDataOk() (*map[string]interface{}, bool)`

GetPredefinedFieldsSampleDataOk returns a tuple with the PredefinedFieldsSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedFieldsSampleData

`func (o *CreateTemplateRequest) SetPredefinedFieldsSampleData(v map[string]interface{})`

SetPredefinedFieldsSampleData sets PredefinedFieldsSampleData field to given value.

### HasPredefinedFieldsSampleData

`func (o *CreateTemplateRequest) HasPredefinedFieldsSampleData() bool`

HasPredefinedFieldsSampleData returns a boolean if a field has been set.

### SetPredefinedFieldsSampleDataNil

`func (o *CreateTemplateRequest) SetPredefinedFieldsSampleDataNil(b bool)`

 SetPredefinedFieldsSampleDataNil sets the value for PredefinedFieldsSampleData to be an explicit nil

### UnsetPredefinedFieldsSampleData
`func (o *CreateTemplateRequest) UnsetPredefinedFieldsSampleData()`

UnsetPredefinedFieldsSampleData ensures that no value is present for PredefinedFieldsSampleData, not even an explicit nil
### GetHtml

`func (o *CreateTemplateRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *CreateTemplateRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *CreateTemplateRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *CreateTemplateRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### SetHtmlNil

`func (o *CreateTemplateRequest) SetHtmlNil(b bool)`

 SetHtmlNil sets the value for Html to be an explicit nil

### UnsetHtml
`func (o *CreateTemplateRequest) UnsetHtml()`

UnsetHtml ensures that no value is present for Html, not even an explicit nil
### GetScss

`func (o *CreateTemplateRequest) GetScss() string`

GetScss returns the Scss field if non-nil, zero value otherwise.

### GetScssOk

`func (o *CreateTemplateRequest) GetScssOk() (*string, bool)`

GetScssOk returns a tuple with the Scss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScss

`func (o *CreateTemplateRequest) SetScss(v string)`

SetScss sets Scss field to given value.

### HasScss

`func (o *CreateTemplateRequest) HasScss() bool`

HasScss returns a boolean if a field has been set.

### SetScssNil

`func (o *CreateTemplateRequest) SetScssNil(b bool)`

 SetScssNil sets the value for Scss to be an explicit nil

### UnsetScss
`func (o *CreateTemplateRequest) UnsetScss()`

UnsetScss ensures that no value is present for Scss, not even an explicit nil
### GetHeaderHtml

`func (o *CreateTemplateRequest) GetHeaderHtml() string`

GetHeaderHtml returns the HeaderHtml field if non-nil, zero value otherwise.

### GetHeaderHtmlOk

`func (o *CreateTemplateRequest) GetHeaderHtmlOk() (*string, bool)`

GetHeaderHtmlOk returns a tuple with the HeaderHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderHtml

`func (o *CreateTemplateRequest) SetHeaderHtml(v string)`

SetHeaderHtml sets HeaderHtml field to given value.

### HasHeaderHtml

`func (o *CreateTemplateRequest) HasHeaderHtml() bool`

HasHeaderHtml returns a boolean if a field has been set.

### SetHeaderHtmlNil

`func (o *CreateTemplateRequest) SetHeaderHtmlNil(b bool)`

 SetHeaderHtmlNil sets the value for HeaderHtml to be an explicit nil

### UnsetHeaderHtml
`func (o *CreateTemplateRequest) UnsetHeaderHtml()`

UnsetHeaderHtml ensures that no value is present for HeaderHtml, not even an explicit nil
### GetFooterHtml

`func (o *CreateTemplateRequest) GetFooterHtml() string`

GetFooterHtml returns the FooterHtml field if non-nil, zero value otherwise.

### GetFooterHtmlOk

`func (o *CreateTemplateRequest) GetFooterHtmlOk() (*string, bool)`

GetFooterHtmlOk returns a tuple with the FooterHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterHtml

`func (o *CreateTemplateRequest) SetFooterHtml(v string)`

SetFooterHtml sets FooterHtml field to given value.

### HasFooterHtml

`func (o *CreateTemplateRequest) HasFooterHtml() bool`

HasFooterHtml returns a boolean if a field has been set.

### SetFooterHtmlNil

`func (o *CreateTemplateRequest) SetFooterHtmlNil(b bool)`

 SetFooterHtmlNil sets the value for FooterHtml to be an explicit nil

### UnsetFooterHtml
`func (o *CreateTemplateRequest) UnsetFooterHtml()`

UnsetFooterHtml ensures that no value is present for FooterHtml, not even an explicit nil
### GetHtmlEngineOptions

`func (o *CreateTemplateRequest) GetHtmlEngineOptions() HtmlEngineOptionsInput`

GetHtmlEngineOptions returns the HtmlEngineOptions field if non-nil, zero value otherwise.

### GetHtmlEngineOptionsOk

`func (o *CreateTemplateRequest) GetHtmlEngineOptionsOk() (*HtmlEngineOptionsInput, bool)`

GetHtmlEngineOptionsOk returns a tuple with the HtmlEngineOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlEngineOptions

`func (o *CreateTemplateRequest) SetHtmlEngineOptions(v HtmlEngineOptionsInput)`

SetHtmlEngineOptions sets HtmlEngineOptions field to given value.

### HasHtmlEngineOptions

`func (o *CreateTemplateRequest) HasHtmlEngineOptions() bool`

HasHtmlEngineOptions returns a boolean if a field has been set.

### SetHtmlEngineOptionsNil

`func (o *CreateTemplateRequest) SetHtmlEngineOptionsNil(b bool)`

 SetHtmlEngineOptionsNil sets the value for HtmlEngineOptions to be an explicit nil

### UnsetHtmlEngineOptions
`func (o *CreateTemplateRequest) UnsetHtmlEngineOptions()`

UnsetHtmlEngineOptions ensures that no value is present for HtmlEngineOptions, not even an explicit nil
### GetHtmlEscapeData

`func (o *CreateTemplateRequest) GetHtmlEscapeData() bool`

GetHtmlEscapeData returns the HtmlEscapeData field if non-nil, zero value otherwise.

### GetHtmlEscapeDataOk

`func (o *CreateTemplateRequest) GetHtmlEscapeDataOk() (*bool, bool)`

GetHtmlEscapeDataOk returns a tuple with the HtmlEscapeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlEscapeData

`func (o *CreateTemplateRequest) SetHtmlEscapeData(v bool)`

SetHtmlEscapeData sets HtmlEscapeData field to given value.

### HasHtmlEscapeData

`func (o *CreateTemplateRequest) HasHtmlEscapeData() bool`

HasHtmlEscapeData returns a boolean if a field has been set.

### GetCacheId

`func (o *CreateTemplateRequest) GetCacheId() string`

GetCacheId returns the CacheId field if non-nil, zero value otherwise.

### GetCacheIdOk

`func (o *CreateTemplateRequest) GetCacheIdOk() (*string, bool)`

GetCacheIdOk returns a tuple with the CacheId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheId

`func (o *CreateTemplateRequest) SetCacheId(v string)`

SetCacheId sets CacheId field to given value.

### HasCacheId

`func (o *CreateTemplateRequest) HasCacheId() bool`

HasCacheId returns a boolean if a field has been set.

### GetDocumentUrl

`func (o *CreateTemplateRequest) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *CreateTemplateRequest) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *CreateTemplateRequest) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.

### HasDocumentUrl

`func (o *CreateTemplateRequest) HasDocumentUrl() bool`

HasDocumentUrl returns a boolean if a field has been set.

### GetDocumentBase64

`func (o *CreateTemplateRequest) GetDocumentBase64() string`

GetDocumentBase64 returns the DocumentBase64 field if non-nil, zero value otherwise.

### GetDocumentBase64Ok

`func (o *CreateTemplateRequest) GetDocumentBase64Ok() (*string, bool)`

GetDocumentBase64Ok returns a tuple with the DocumentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentBase64

`func (o *CreateTemplateRequest) SetDocumentBase64(v string)`

SetDocumentBase64 sets DocumentBase64 field to given value.

### HasDocumentBase64

`func (o *CreateTemplateRequest) HasDocumentBase64() bool`

HasDocumentBase64 returns a boolean if a field has been set.

### GetFilename

`func (o *CreateTemplateRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CreateTemplateRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CreateTemplateRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CreateTemplateRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetDetectFields

`func (o *CreateTemplateRequest) GetDetectFields() bool`

GetDetectFields returns the DetectFields field if non-nil, zero value otherwise.

### GetDetectFieldsOk

`func (o *CreateTemplateRequest) GetDetectFieldsOk() (*bool, bool)`

GetDetectFieldsOk returns a tuple with the DetectFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectFields

`func (o *CreateTemplateRequest) SetDetectFields(v bool)`

SetDetectFields sets DetectFields field to given value.

### HasDetectFields

`func (o *CreateTemplateRequest) HasDetectFields() bool`

HasDetectFields returns a boolean if a field has been set.

### GetFieldSource

`func (o *CreateTemplateRequest) GetFieldSource() FieldSetupSource`

GetFieldSource returns the FieldSource field if non-nil, zero value otherwise.

### GetFieldSourceOk

`func (o *CreateTemplateRequest) GetFieldSourceOk() (*FieldSetupSource, bool)`

GetFieldSourceOk returns a tuple with the FieldSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSource

`func (o *CreateTemplateRequest) SetFieldSource(v FieldSetupSource)`

SetFieldSource sets FieldSource field to given value.

### HasFieldSource

`func (o *CreateTemplateRequest) HasFieldSource() bool`

HasFieldSource returns a boolean if a field has been set.

### GetFieldNaming

`func (o *CreateTemplateRequest) GetFieldNaming() FieldSetupNaming`

GetFieldNaming returns the FieldNaming field if non-nil, zero value otherwise.

### GetFieldNamingOk

`func (o *CreateTemplateRequest) GetFieldNamingOk() (*FieldSetupNaming, bool)`

GetFieldNamingOk returns a tuple with the FieldNaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNaming

`func (o *CreateTemplateRequest) SetFieldNaming(v FieldSetupNaming)`

SetFieldNaming sets FieldNaming field to given value.

### HasFieldNaming

`func (o *CreateTemplateRequest) HasFieldNaming() bool`

HasFieldNaming returns a boolean if a field has been set.

### GetFieldLogic

`func (o *CreateTemplateRequest) GetFieldLogic() FieldSetupLogic`

GetFieldLogic returns the FieldLogic field if non-nil, zero value otherwise.

### GetFieldLogicOk

`func (o *CreateTemplateRequest) GetFieldLogicOk() (*FieldSetupLogic, bool)`

GetFieldLogicOk returns a tuple with the FieldLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLogic

`func (o *CreateTemplateRequest) SetFieldLogic(v FieldSetupLogic)`

SetFieldLogic sets FieldLogic field to given value.

### HasFieldLogic

`func (o *CreateTemplateRequest) HasFieldLogic() bool`

HasFieldLogic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


