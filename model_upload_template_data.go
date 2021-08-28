/*
 * API v1
 *
 * DocSpring is a service that helps you fill out and sign PDF templates.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package docspring

type UploadTemplateData struct {
	ExpirationInterval string `json:"expiration_interval,omitempty"`
	WebhookUrl string `json:"webhook_url,omitempty"`
	Scss string `json:"scss,omitempty"`
	AllowAdditionalProperties bool `json:"allow_additional_properties,omitempty"`
	Document UploadTemplateDataDocument `json:"document,omitempty"`
	ExpireAfter float32 `json:"expire_after,omitempty"`
	Description string `json:"description,omitempty"`
	PublicSubmissions bool `json:"public_submissions,omitempty"`
	SlackWebhookUrl string `json:"slack_webhook_url,omitempty"`
	HeaderHtml string `json:"header_html,omitempty"`
	PublicWebForm bool `json:"public_web_form,omitempty"`
	EditableSubmissions bool `json:"editable_submissions,omitempty"`
	ExpireSubmissions bool `json:"expire_submissions,omitempty"`
	Name string `json:"name,omitempty"`
	FooterHtml string `json:"footer_html,omitempty"`
	Html string `json:"html,omitempty"`
	TemplateType string `json:"template_type,omitempty"`
	RedirectUrl string `json:"redirect_url,omitempty"`
}
