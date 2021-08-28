/*
 * API v1
 *
 * DocSpring is a service that helps you fill out and sign PDF templates.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package docspring

type UploadTemplateDataDocumentMetadata struct {
	Filename string `json:"filename,omitempty"`
	Size int32 `json:"size,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
}
