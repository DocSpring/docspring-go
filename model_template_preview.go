/*
DocSpring API

DocSpring provides an API that helps you fill out and sign PDF templates.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package docspring

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TemplatePreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplatePreview{}

// TemplatePreview struct for TemplatePreview
type TemplatePreview struct {
	AddDataRequestSubmissionIdFooters bool `json:"add_data_request_submission_id_footers"`
	AllowAdditionalProperties bool `json:"allow_additional_properties"`
	Description NullableString `json:"description"`
	DocumentFilename NullableString `json:"document_filename"`
	DocumentMd5 NullableString `json:"document_md5"`
	DocumentParseError bool `json:"document_parse_error"`
	DocumentProcessed bool `json:"document_processed"`
	DocumentState string `json:"document_state"`
	DocumentUrl NullableString `json:"document_url"`
	EditableSubmissions bool `json:"editable_submissions"`
	EmbedDomains NullableString `json:"embed_domains"`
	EncryptPdfsPassword NullableString `json:"encrypt_pdfs_password"`
	EncryptPdfs bool `json:"encrypt_pdfs"`
	ExpirationInterval string `json:"expiration_interval"`
	ExpireAfter int32 `json:"expire_after"`
	ExpireSubmissions bool `json:"expire_submissions"`
	ExternalPredefinedFieldsTemplateId NullableString `json:"external_predefined_fields_template_id"`
	ExternalPredefinedFieldsTemplateName NullableString `json:"external_predefined_fields_template_name"`
	FirstTemplate bool `json:"first_template"`
	Id NullableString `json:"id"`
	Locked bool `json:"locked"`
	MergeAuditTrailPdf bool `json:"merge_audit_trail_pdf"`
	Name NullableString `json:"name"`
	PageCount int32 `json:"page_count"`
	PageDimensions [][]float32 `json:"page_dimensions"`
	ParentFolderId NullableString `json:"parent_folder_id"`
	Path NullableString `json:"path"`
	PermanentDocumentUrl NullableString `json:"permanent_document_url"`
	PublicSubmissions bool `json:"public_submissions"`
	PublicWebForm bool `json:"public_web_form"`
	RedirectUrl NullableString `json:"redirect_url"`
	SlackWebhookUrl NullableString `json:"slack_webhook_url"`
	TemplateType string `json:"template_type"`
	UpdatedAt NullableString `json:"updated_at"`
	WebhookUrl NullableString `json:"webhook_url"`
	Demo bool `json:"demo"`
}

type _TemplatePreview TemplatePreview

// NewTemplatePreview instantiates a new TemplatePreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplatePreview(addDataRequestSubmissionIdFooters bool, allowAdditionalProperties bool, description NullableString, documentFilename NullableString, documentMd5 NullableString, documentParseError bool, documentProcessed bool, documentState string, documentUrl NullableString, editableSubmissions bool, embedDomains NullableString, encryptPdfsPassword NullableString, encryptPdfs bool, expirationInterval string, expireAfter int32, expireSubmissions bool, externalPredefinedFieldsTemplateId NullableString, externalPredefinedFieldsTemplateName NullableString, firstTemplate bool, id NullableString, locked bool, mergeAuditTrailPdf bool, name NullableString, pageCount int32, pageDimensions [][]float32, parentFolderId NullableString, path NullableString, permanentDocumentUrl NullableString, publicSubmissions bool, publicWebForm bool, redirectUrl NullableString, slackWebhookUrl NullableString, templateType string, updatedAt NullableString, webhookUrl NullableString, demo bool) *TemplatePreview {
	this := TemplatePreview{}
	this.AddDataRequestSubmissionIdFooters = addDataRequestSubmissionIdFooters
	this.AllowAdditionalProperties = allowAdditionalProperties
	this.Description = description
	this.DocumentFilename = documentFilename
	this.DocumentMd5 = documentMd5
	this.DocumentParseError = documentParseError
	this.DocumentProcessed = documentProcessed
	this.DocumentState = documentState
	this.DocumentUrl = documentUrl
	this.EditableSubmissions = editableSubmissions
	this.EmbedDomains = embedDomains
	this.EncryptPdfsPassword = encryptPdfsPassword
	this.EncryptPdfs = encryptPdfs
	this.ExpirationInterval = expirationInterval
	this.ExpireAfter = expireAfter
	this.ExpireSubmissions = expireSubmissions
	this.ExternalPredefinedFieldsTemplateId = externalPredefinedFieldsTemplateId
	this.ExternalPredefinedFieldsTemplateName = externalPredefinedFieldsTemplateName
	this.FirstTemplate = firstTemplate
	this.Id = id
	this.Locked = locked
	this.MergeAuditTrailPdf = mergeAuditTrailPdf
	this.Name = name
	this.PageCount = pageCount
	this.PageDimensions = pageDimensions
	this.ParentFolderId = parentFolderId
	this.Path = path
	this.PermanentDocumentUrl = permanentDocumentUrl
	this.PublicSubmissions = publicSubmissions
	this.PublicWebForm = publicWebForm
	this.RedirectUrl = redirectUrl
	this.SlackWebhookUrl = slackWebhookUrl
	this.TemplateType = templateType
	this.UpdatedAt = updatedAt
	this.WebhookUrl = webhookUrl
	this.Demo = demo
	return &this
}

// NewTemplatePreviewWithDefaults instantiates a new TemplatePreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatePreviewWithDefaults() *TemplatePreview {
	this := TemplatePreview{}
	return &this
}

// GetAddDataRequestSubmissionIdFooters returns the AddDataRequestSubmissionIdFooters field value
func (o *TemplatePreview) GetAddDataRequestSubmissionIdFooters() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AddDataRequestSubmissionIdFooters
}

// GetAddDataRequestSubmissionIdFootersOk returns a tuple with the AddDataRequestSubmissionIdFooters field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetAddDataRequestSubmissionIdFootersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddDataRequestSubmissionIdFooters, true
}

// SetAddDataRequestSubmissionIdFooters sets field value
func (o *TemplatePreview) SetAddDataRequestSubmissionIdFooters(v bool) {
	o.AddDataRequestSubmissionIdFooters = v
}

// GetAllowAdditionalProperties returns the AllowAdditionalProperties field value
func (o *TemplatePreview) GetAllowAdditionalProperties() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAdditionalProperties
}

// GetAllowAdditionalPropertiesOk returns a tuple with the AllowAdditionalProperties field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetAllowAdditionalPropertiesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAdditionalProperties, true
}

// SetAllowAdditionalProperties sets field value
func (o *TemplatePreview) SetAllowAdditionalProperties(v bool) {
	o.AllowAdditionalProperties = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *TemplatePreview) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetDocumentFilename returns the DocumentFilename field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetDocumentFilename() string {
	if o == nil || o.DocumentFilename.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentFilename.Get()
}

// GetDocumentFilenameOk returns a tuple with the DocumentFilename field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetDocumentFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentFilename.Get(), o.DocumentFilename.IsSet()
}

// SetDocumentFilename sets field value
func (o *TemplatePreview) SetDocumentFilename(v string) {
	o.DocumentFilename.Set(&v)
}

// GetDocumentMd5 returns the DocumentMd5 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetDocumentMd5() string {
	if o == nil || o.DocumentMd5.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentMd5.Get()
}

// GetDocumentMd5Ok returns a tuple with the DocumentMd5 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetDocumentMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentMd5.Get(), o.DocumentMd5.IsSet()
}

// SetDocumentMd5 sets field value
func (o *TemplatePreview) SetDocumentMd5(v string) {
	o.DocumentMd5.Set(&v)
}

// GetDocumentParseError returns the DocumentParseError field value
func (o *TemplatePreview) GetDocumentParseError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DocumentParseError
}

// GetDocumentParseErrorOk returns a tuple with the DocumentParseError field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetDocumentParseErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentParseError, true
}

// SetDocumentParseError sets field value
func (o *TemplatePreview) SetDocumentParseError(v bool) {
	o.DocumentParseError = v
}

// GetDocumentProcessed returns the DocumentProcessed field value
func (o *TemplatePreview) GetDocumentProcessed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DocumentProcessed
}

// GetDocumentProcessedOk returns a tuple with the DocumentProcessed field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetDocumentProcessedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentProcessed, true
}

// SetDocumentProcessed sets field value
func (o *TemplatePreview) SetDocumentProcessed(v bool) {
	o.DocumentProcessed = v
}

// GetDocumentState returns the DocumentState field value
func (o *TemplatePreview) GetDocumentState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentState
}

// GetDocumentStateOk returns a tuple with the DocumentState field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetDocumentStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentState, true
}

// SetDocumentState sets field value
func (o *TemplatePreview) SetDocumentState(v string) {
	o.DocumentState = v
}

// GetDocumentUrl returns the DocumentUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetDocumentUrl() string {
	if o == nil || o.DocumentUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentUrl.Get()
}

// GetDocumentUrlOk returns a tuple with the DocumentUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetDocumentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentUrl.Get(), o.DocumentUrl.IsSet()
}

// SetDocumentUrl sets field value
func (o *TemplatePreview) SetDocumentUrl(v string) {
	o.DocumentUrl.Set(&v)
}

// GetEditableSubmissions returns the EditableSubmissions field value
func (o *TemplatePreview) GetEditableSubmissions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EditableSubmissions
}

// GetEditableSubmissionsOk returns a tuple with the EditableSubmissions field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetEditableSubmissionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditableSubmissions, true
}

// SetEditableSubmissions sets field value
func (o *TemplatePreview) SetEditableSubmissions(v bool) {
	o.EditableSubmissions = v
}

// GetEmbedDomains returns the EmbedDomains field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetEmbedDomains() string {
	if o == nil || o.EmbedDomains.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmbedDomains.Get()
}

// GetEmbedDomainsOk returns a tuple with the EmbedDomains field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetEmbedDomainsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmbedDomains.Get(), o.EmbedDomains.IsSet()
}

// SetEmbedDomains sets field value
func (o *TemplatePreview) SetEmbedDomains(v string) {
	o.EmbedDomains.Set(&v)
}

// GetEncryptPdfsPassword returns the EncryptPdfsPassword field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetEncryptPdfsPassword() string {
	if o == nil || o.EncryptPdfsPassword.Get() == nil {
		var ret string
		return ret
	}

	return *o.EncryptPdfsPassword.Get()
}

// GetEncryptPdfsPasswordOk returns a tuple with the EncryptPdfsPassword field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetEncryptPdfsPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptPdfsPassword.Get(), o.EncryptPdfsPassword.IsSet()
}

// SetEncryptPdfsPassword sets field value
func (o *TemplatePreview) SetEncryptPdfsPassword(v string) {
	o.EncryptPdfsPassword.Set(&v)
}

// GetEncryptPdfs returns the EncryptPdfs field value
func (o *TemplatePreview) GetEncryptPdfs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EncryptPdfs
}

// GetEncryptPdfsOk returns a tuple with the EncryptPdfs field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetEncryptPdfsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptPdfs, true
}

// SetEncryptPdfs sets field value
func (o *TemplatePreview) SetEncryptPdfs(v bool) {
	o.EncryptPdfs = v
}

// GetExpirationInterval returns the ExpirationInterval field value
func (o *TemplatePreview) GetExpirationInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationInterval
}

// GetExpirationIntervalOk returns a tuple with the ExpirationInterval field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetExpirationIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationInterval, true
}

// SetExpirationInterval sets field value
func (o *TemplatePreview) SetExpirationInterval(v string) {
	o.ExpirationInterval = v
}

// GetExpireAfter returns the ExpireAfter field value
func (o *TemplatePreview) GetExpireAfter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpireAfter
}

// GetExpireAfterOk returns a tuple with the ExpireAfter field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetExpireAfterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpireAfter, true
}

// SetExpireAfter sets field value
func (o *TemplatePreview) SetExpireAfter(v int32) {
	o.ExpireAfter = v
}

// GetExpireSubmissions returns the ExpireSubmissions field value
func (o *TemplatePreview) GetExpireSubmissions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExpireSubmissions
}

// GetExpireSubmissionsOk returns a tuple with the ExpireSubmissions field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetExpireSubmissionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpireSubmissions, true
}

// SetExpireSubmissions sets field value
func (o *TemplatePreview) SetExpireSubmissions(v bool) {
	o.ExpireSubmissions = v
}

// GetExternalPredefinedFieldsTemplateId returns the ExternalPredefinedFieldsTemplateId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetExternalPredefinedFieldsTemplateId() string {
	if o == nil || o.ExternalPredefinedFieldsTemplateId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalPredefinedFieldsTemplateId.Get()
}

// GetExternalPredefinedFieldsTemplateIdOk returns a tuple with the ExternalPredefinedFieldsTemplateId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetExternalPredefinedFieldsTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalPredefinedFieldsTemplateId.Get(), o.ExternalPredefinedFieldsTemplateId.IsSet()
}

// SetExternalPredefinedFieldsTemplateId sets field value
func (o *TemplatePreview) SetExternalPredefinedFieldsTemplateId(v string) {
	o.ExternalPredefinedFieldsTemplateId.Set(&v)
}

// GetExternalPredefinedFieldsTemplateName returns the ExternalPredefinedFieldsTemplateName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetExternalPredefinedFieldsTemplateName() string {
	if o == nil || o.ExternalPredefinedFieldsTemplateName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalPredefinedFieldsTemplateName.Get()
}

// GetExternalPredefinedFieldsTemplateNameOk returns a tuple with the ExternalPredefinedFieldsTemplateName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetExternalPredefinedFieldsTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalPredefinedFieldsTemplateName.Get(), o.ExternalPredefinedFieldsTemplateName.IsSet()
}

// SetExternalPredefinedFieldsTemplateName sets field value
func (o *TemplatePreview) SetExternalPredefinedFieldsTemplateName(v string) {
	o.ExternalPredefinedFieldsTemplateName.Set(&v)
}

// GetFirstTemplate returns the FirstTemplate field value
func (o *TemplatePreview) GetFirstTemplate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FirstTemplate
}

// GetFirstTemplateOk returns a tuple with the FirstTemplate field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetFirstTemplateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstTemplate, true
}

// SetFirstTemplate sets field value
func (o *TemplatePreview) SetFirstTemplate(v bool) {
	o.FirstTemplate = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *TemplatePreview) SetId(v string) {
	o.Id.Set(&v)
}

// GetLocked returns the Locked field value
func (o *TemplatePreview) GetLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *TemplatePreview) SetLocked(v bool) {
	o.Locked = v
}

// GetMergeAuditTrailPdf returns the MergeAuditTrailPdf field value
func (o *TemplatePreview) GetMergeAuditTrailPdf() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MergeAuditTrailPdf
}

// GetMergeAuditTrailPdfOk returns a tuple with the MergeAuditTrailPdf field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetMergeAuditTrailPdfOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MergeAuditTrailPdf, true
}

// SetMergeAuditTrailPdf sets field value
func (o *TemplatePreview) SetMergeAuditTrailPdf(v bool) {
	o.MergeAuditTrailPdf = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *TemplatePreview) SetName(v string) {
	o.Name.Set(&v)
}

// GetPageCount returns the PageCount field value
func (o *TemplatePreview) GetPageCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageCount
}

// GetPageCountOk returns a tuple with the PageCount field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetPageCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageCount, true
}

// SetPageCount sets field value
func (o *TemplatePreview) SetPageCount(v int32) {
	o.PageCount = v
}

// GetPageDimensions returns the PageDimensions field value
// If the value is explicit nil, the zero value for [][]float32 will be returned
func (o *TemplatePreview) GetPageDimensions() [][]float32 {
	if o == nil {
		var ret [][]float32
		return ret
	}

	return o.PageDimensions
}

// GetPageDimensionsOk returns a tuple with the PageDimensions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetPageDimensionsOk() ([][]float32, bool) {
	if o == nil || IsNil(o.PageDimensions) {
		return nil, false
	}
	return o.PageDimensions, true
}

// SetPageDimensions sets field value
func (o *TemplatePreview) SetPageDimensions(v [][]float32) {
	o.PageDimensions = v
}

// GetParentFolderId returns the ParentFolderId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetParentFolderId() string {
	if o == nil || o.ParentFolderId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentFolderId.Get()
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetParentFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentFolderId.Get(), o.ParentFolderId.IsSet()
}

// SetParentFolderId sets field value
func (o *TemplatePreview) SetParentFolderId(v string) {
	o.ParentFolderId.Set(&v)
}

// GetPath returns the Path field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}

	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// SetPath sets field value
func (o *TemplatePreview) SetPath(v string) {
	o.Path.Set(&v)
}

// GetPermanentDocumentUrl returns the PermanentDocumentUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetPermanentDocumentUrl() string {
	if o == nil || o.PermanentDocumentUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.PermanentDocumentUrl.Get()
}

// GetPermanentDocumentUrlOk returns a tuple with the PermanentDocumentUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetPermanentDocumentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermanentDocumentUrl.Get(), o.PermanentDocumentUrl.IsSet()
}

// SetPermanentDocumentUrl sets field value
func (o *TemplatePreview) SetPermanentDocumentUrl(v string) {
	o.PermanentDocumentUrl.Set(&v)
}

// GetPublicSubmissions returns the PublicSubmissions field value
func (o *TemplatePreview) GetPublicSubmissions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PublicSubmissions
}

// GetPublicSubmissionsOk returns a tuple with the PublicSubmissions field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetPublicSubmissionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicSubmissions, true
}

// SetPublicSubmissions sets field value
func (o *TemplatePreview) SetPublicSubmissions(v bool) {
	o.PublicSubmissions = v
}

// GetPublicWebForm returns the PublicWebForm field value
func (o *TemplatePreview) GetPublicWebForm() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PublicWebForm
}

// GetPublicWebFormOk returns a tuple with the PublicWebForm field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetPublicWebFormOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicWebForm, true
}

// SetPublicWebForm sets field value
func (o *TemplatePreview) SetPublicWebForm(v bool) {
	o.PublicWebForm = v
}

// GetRedirectUrl returns the RedirectUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.RedirectUrl.Get()
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUrl.Get(), o.RedirectUrl.IsSet()
}

// SetRedirectUrl sets field value
func (o *TemplatePreview) SetRedirectUrl(v string) {
	o.RedirectUrl.Set(&v)
}

// GetSlackWebhookUrl returns the SlackWebhookUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetSlackWebhookUrl() string {
	if o == nil || o.SlackWebhookUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.SlackWebhookUrl.Get()
}

// GetSlackWebhookUrlOk returns a tuple with the SlackWebhookUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetSlackWebhookUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlackWebhookUrl.Get(), o.SlackWebhookUrl.IsSet()
}

// SetSlackWebhookUrl sets field value
func (o *TemplatePreview) SetSlackWebhookUrl(v string) {
	o.SlackWebhookUrl.Set(&v)
}

// GetTemplateType returns the TemplateType field value
func (o *TemplatePreview) GetTemplateType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetTemplateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateType, true
}

// SetTemplateType sets field value
func (o *TemplatePreview) SetTemplateType(v string) {
	o.TemplateType = v
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *TemplatePreview) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}

// GetWebhookUrl returns the WebhookUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplatePreview) GetWebhookUrl() string {
	if o == nil || o.WebhookUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.WebhookUrl.Get()
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePreview) GetWebhookUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebhookUrl.Get(), o.WebhookUrl.IsSet()
}

// SetWebhookUrl sets field value
func (o *TemplatePreview) SetWebhookUrl(v string) {
	o.WebhookUrl.Set(&v)
}

// GetDemo returns the Demo field value
func (o *TemplatePreview) GetDemo() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Demo
}

// GetDemoOk returns a tuple with the Demo field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetDemoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Demo, true
}

// SetDemo sets field value
func (o *TemplatePreview) SetDemo(v bool) {
	o.Demo = v
}

func (o TemplatePreview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplatePreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["add_data_request_submission_id_footers"] = o.AddDataRequestSubmissionIdFooters
	toSerialize["allow_additional_properties"] = o.AllowAdditionalProperties
	toSerialize["description"] = o.Description.Get()
	toSerialize["document_filename"] = o.DocumentFilename.Get()
	toSerialize["document_md5"] = o.DocumentMd5.Get()
	toSerialize["document_parse_error"] = o.DocumentParseError
	toSerialize["document_processed"] = o.DocumentProcessed
	toSerialize["document_state"] = o.DocumentState
	toSerialize["document_url"] = o.DocumentUrl.Get()
	toSerialize["editable_submissions"] = o.EditableSubmissions
	toSerialize["embed_domains"] = o.EmbedDomains.Get()
	toSerialize["encrypt_pdfs_password"] = o.EncryptPdfsPassword.Get()
	toSerialize["encrypt_pdfs"] = o.EncryptPdfs
	toSerialize["expiration_interval"] = o.ExpirationInterval
	toSerialize["expire_after"] = o.ExpireAfter
	toSerialize["expire_submissions"] = o.ExpireSubmissions
	toSerialize["external_predefined_fields_template_id"] = o.ExternalPredefinedFieldsTemplateId.Get()
	toSerialize["external_predefined_fields_template_name"] = o.ExternalPredefinedFieldsTemplateName.Get()
	toSerialize["first_template"] = o.FirstTemplate
	toSerialize["id"] = o.Id.Get()
	toSerialize["locked"] = o.Locked
	toSerialize["merge_audit_trail_pdf"] = o.MergeAuditTrailPdf
	toSerialize["name"] = o.Name.Get()
	toSerialize["page_count"] = o.PageCount
	if o.PageDimensions != nil {
		toSerialize["page_dimensions"] = o.PageDimensions
	}
	toSerialize["parent_folder_id"] = o.ParentFolderId.Get()
	toSerialize["path"] = o.Path.Get()
	toSerialize["permanent_document_url"] = o.PermanentDocumentUrl.Get()
	toSerialize["public_submissions"] = o.PublicSubmissions
	toSerialize["public_web_form"] = o.PublicWebForm
	toSerialize["redirect_url"] = o.RedirectUrl.Get()
	toSerialize["slack_webhook_url"] = o.SlackWebhookUrl.Get()
	toSerialize["template_type"] = o.TemplateType
	toSerialize["updated_at"] = o.UpdatedAt.Get()
	toSerialize["webhook_url"] = o.WebhookUrl.Get()
	toSerialize["demo"] = o.Demo
	return toSerialize, nil
}

func (o *TemplatePreview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"add_data_request_submission_id_footers",
		"allow_additional_properties",
		"description",
		"document_filename",
		"document_md5",
		"document_parse_error",
		"document_processed",
		"document_state",
		"document_url",
		"editable_submissions",
		"embed_domains",
		"encrypt_pdfs_password",
		"encrypt_pdfs",
		"expiration_interval",
		"expire_after",
		"expire_submissions",
		"external_predefined_fields_template_id",
		"external_predefined_fields_template_name",
		"first_template",
		"id",
		"locked",
		"merge_audit_trail_pdf",
		"name",
		"page_count",
		"page_dimensions",
		"parent_folder_id",
		"path",
		"permanent_document_url",
		"public_submissions",
		"public_web_form",
		"redirect_url",
		"slack_webhook_url",
		"template_type",
		"updated_at",
		"webhook_url",
		"demo",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTemplatePreview := _TemplatePreview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTemplatePreview)

	if err != nil {
		return err
	}

	*o = TemplatePreview(varTemplatePreview)

	return err
}

type NullableTemplatePreview struct {
	value *TemplatePreview
	isSet bool
}

func (v NullableTemplatePreview) Get() *TemplatePreview {
	return v.value
}

func (v *NullableTemplatePreview) Set(val *TemplatePreview) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplatePreview) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplatePreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplatePreview(val *TemplatePreview) *NullableTemplatePreview {
	return &NullableTemplatePreview{value: val, isSet: true}
}

func (v NullableTemplatePreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplatePreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


