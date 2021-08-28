/*
 * API v1
 *
 * DocSpring is a service that helps you fill out and sign PDF templates.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package docspring

type CreateCombinedSubmissionResponse struct {
	CombinedSubmission CombinedSubmission `json:"combined_submission,omitempty"`
	Errors []string `json:"errors,omitempty"`
	Status string `json:"status,omitempty"`
}
