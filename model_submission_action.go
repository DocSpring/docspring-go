/*
 * API v1
 *
 * DocSpring is a service that helps you fill out and sign PDF templates.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package docspring

type SubmissionAction struct {
	Id string `json:"id"`
	IntegrationId string `json:"integration_id"`
	State string `json:"state"`
	ActionCategory string `json:"action_category"`
	ActionType string `json:"action_type"`
	ResultData map[string]interface{} `json:"result_data"`
}
