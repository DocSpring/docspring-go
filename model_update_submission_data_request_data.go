/*
 * API v1
 *
 * DocSpring is a service that helps you fill out and sign PDF templates.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package docspring

type UpdateSubmissionDataRequestData struct {
	AuthPhoneNumberHash string `json:"auth_phone_number_hash,omitempty"`
	AuthProvider string `json:"auth_provider,omitempty"`
	AuthSecondFactorType string `json:"auth_second_factor_type,omitempty"`
	AuthSessionIdHash string `json:"auth_session_id_hash,omitempty"`
	AuthSessionStartedAt string `json:"auth_session_started_at,omitempty"`
	AuthType string `json:"auth_type,omitempty"`
	AuthUserIdHash string `json:"auth_user_id_hash,omitempty"`
	AuthUsernameHash string `json:"auth_username_hash,omitempty"`
	Email string `json:"email,omitempty"`
	Fields []string `json:"fields,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name,omitempty"`
	Order int32 `json:"order,omitempty"`
}
