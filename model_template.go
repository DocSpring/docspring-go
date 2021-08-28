/*
 * API v1
 *
 * DocSpring is a service that helps you fill out and sign PDF templates.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package docspring

type Template struct {
	ExpirationInterval string `json:"expiration_interval,omitempty"`
	WebhookUrl string `json:"webhook_url,omitempty"`
	ParentFolderId string `json:"parent_folder_id,omitempty"`
	ExpireAfter float32 `json:"expire_after,omitempty"`
	AllowAdditionalProperties bool `json:"allow_additional_properties,omitempty"`
	Description string `json:"description,omitempty"`
	PublicSubmissions bool `json:"public_submissions,omitempty"`
	SlackWebhookUrl string `json:"slack_webhook_url,omitempty"`
	Path string `json:"path,omitempty"`
	PublicWebForm bool `json:"public_web_form,omitempty"`
	EditableSubmissions bool `json:"editable_submissions,omitempty"`
	ExpireSubmissions bool `json:"expire_submissions,omitempty"`
	Name string `json:"name,omitempty"`
	PermanentDocumentUrl string `json:"permanent_document_url,omitempty"`
	TemplateType string `json:"template_type,omitempty"`
	Id string `json:"id,omitempty"`
	PageDimensions [][]float32 `json:"page_dimensions,omitempty"`
	Locked bool `json:"locked,omitempty"`
	RedirectUrl string `json:"redirect_url,omitempty"`
	DocumentUrl string `json:"document_url,omitempty"`
}