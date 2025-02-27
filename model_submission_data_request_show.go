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

// checks if the SubmissionDataRequestShow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmissionDataRequestShow{}

// SubmissionDataRequestShow struct for SubmissionDataRequestShow
type SubmissionDataRequestShow struct {
	Id NullableString `json:"id"`
	Email NullableString `json:"email"`
	Name NullableString `json:"name"`
	Order NullableInt32 `json:"order"`
	SortOrder int32 `json:"sort_order"`
	Fields []string `json:"fields"`
	Metadata map[string]interface{} `json:"metadata"`
	State string `json:"state"`
	ViewedAt NullableString `json:"viewed_at"`
	CompletedAt NullableString `json:"completed_at"`
	Data map[string]interface{} `json:"data"`
	AuthType string `json:"auth_type"`
	AuthSecondFactorType string `json:"auth_second_factor_type"`
	AuthProvider NullableString `json:"auth_provider"`
	AuthSessionStartedAt NullableString `json:"auth_session_started_at"`
	AuthSessionIdHash NullableString `json:"auth_session_id_hash"`
	AuthUserIdHash NullableString `json:"auth_user_id_hash"`
	AuthUsernameHash NullableString `json:"auth_username_hash"`
	AuthPhoneNumberHash NullableString `json:"auth_phone_number_hash"`
	IpAddress NullableString `json:"ip_address"`
	UserAgent NullableString `json:"user_agent"`
	SubmissionId NullableString `json:"submission_id"`
}

type _SubmissionDataRequestShow SubmissionDataRequestShow

// NewSubmissionDataRequestShow instantiates a new SubmissionDataRequestShow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmissionDataRequestShow(id NullableString, email NullableString, name NullableString, order NullableInt32, sortOrder int32, fields []string, metadata map[string]interface{}, state string, viewedAt NullableString, completedAt NullableString, data map[string]interface{}, authType string, authSecondFactorType string, authProvider NullableString, authSessionStartedAt NullableString, authSessionIdHash NullableString, authUserIdHash NullableString, authUsernameHash NullableString, authPhoneNumberHash NullableString, ipAddress NullableString, userAgent NullableString, submissionId NullableString) *SubmissionDataRequestShow {
	this := SubmissionDataRequestShow{}
	this.Id = id
	this.Email = email
	this.Name = name
	this.Order = order
	this.SortOrder = sortOrder
	this.Fields = fields
	this.Metadata = metadata
	this.State = state
	this.ViewedAt = viewedAt
	this.CompletedAt = completedAt
	this.Data = data
	this.AuthType = authType
	this.AuthSecondFactorType = authSecondFactorType
	this.AuthProvider = authProvider
	this.AuthSessionStartedAt = authSessionStartedAt
	this.AuthSessionIdHash = authSessionIdHash
	this.AuthUserIdHash = authUserIdHash
	this.AuthUsernameHash = authUsernameHash
	this.AuthPhoneNumberHash = authPhoneNumberHash
	this.IpAddress = ipAddress
	this.UserAgent = userAgent
	this.SubmissionId = submissionId
	return &this
}

// NewSubmissionDataRequestShowWithDefaults instantiates a new SubmissionDataRequestShow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmissionDataRequestShowWithDefaults() *SubmissionDataRequestShow {
	this := SubmissionDataRequestShow{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *SubmissionDataRequestShow) SetId(v string) {
	o.Id.Set(&v)
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *SubmissionDataRequestShow) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *SubmissionDataRequestShow) SetName(v string) {
	o.Name.Set(&v)
}

// GetOrder returns the Order field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *SubmissionDataRequestShow) GetOrder() int32 {
	if o == nil || o.Order.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// SetOrder sets field value
func (o *SubmissionDataRequestShow) SetOrder(v int32) {
	o.Order.Set(&v)
}

// GetSortOrder returns the SortOrder field value
func (o *SubmissionDataRequestShow) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *SubmissionDataRequestShow) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *SubmissionDataRequestShow) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetFields returns the Fields field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *SubmissionDataRequestShow) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *SubmissionDataRequestShow) SetFields(v []string) {
	o.Fields = v
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *SubmissionDataRequestShow) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *SubmissionDataRequestShow) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetState returns the State field value
func (o *SubmissionDataRequestShow) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SubmissionDataRequestShow) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SubmissionDataRequestShow) SetState(v string) {
	o.State = v
}

// GetViewedAt returns the ViewedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetViewedAt() string {
	if o == nil || o.ViewedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ViewedAt.Get()
}

// GetViewedAtOk returns a tuple with the ViewedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetViewedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewedAt.Get(), o.ViewedAt.IsSet()
}

// SetViewedAt sets field value
func (o *SubmissionDataRequestShow) SetViewedAt(v string) {
	o.ViewedAt.Set(&v)
}

// GetCompletedAt returns the CompletedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetCompletedAt() string {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetCompletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// SetCompletedAt sets field value
func (o *SubmissionDataRequestShow) SetCompletedAt(v string) {
	o.CompletedAt.Set(&v)
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *SubmissionDataRequestShow) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubmissionDataRequestShow) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetAuthType returns the AuthType field value
func (o *SubmissionDataRequestShow) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *SubmissionDataRequestShow) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *SubmissionDataRequestShow) SetAuthType(v string) {
	o.AuthType = v
}

// GetAuthSecondFactorType returns the AuthSecondFactorType field value
func (o *SubmissionDataRequestShow) GetAuthSecondFactorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthSecondFactorType
}

// GetAuthSecondFactorTypeOk returns a tuple with the AuthSecondFactorType field value
// and a boolean to check if the value has been set.
func (o *SubmissionDataRequestShow) GetAuthSecondFactorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthSecondFactorType, true
}

// SetAuthSecondFactorType sets field value
func (o *SubmissionDataRequestShow) SetAuthSecondFactorType(v string) {
	o.AuthSecondFactorType = v
}

// GetAuthProvider returns the AuthProvider field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetAuthProvider() string {
	if o == nil || o.AuthProvider.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthProvider.Get()
}

// GetAuthProviderOk returns a tuple with the AuthProvider field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetAuthProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthProvider.Get(), o.AuthProvider.IsSet()
}

// SetAuthProvider sets field value
func (o *SubmissionDataRequestShow) SetAuthProvider(v string) {
	o.AuthProvider.Set(&v)
}

// GetAuthSessionStartedAt returns the AuthSessionStartedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetAuthSessionStartedAt() string {
	if o == nil || o.AuthSessionStartedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthSessionStartedAt.Get()
}

// GetAuthSessionStartedAtOk returns a tuple with the AuthSessionStartedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetAuthSessionStartedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthSessionStartedAt.Get(), o.AuthSessionStartedAt.IsSet()
}

// SetAuthSessionStartedAt sets field value
func (o *SubmissionDataRequestShow) SetAuthSessionStartedAt(v string) {
	o.AuthSessionStartedAt.Set(&v)
}

// GetAuthSessionIdHash returns the AuthSessionIdHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetAuthSessionIdHash() string {
	if o == nil || o.AuthSessionIdHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthSessionIdHash.Get()
}

// GetAuthSessionIdHashOk returns a tuple with the AuthSessionIdHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetAuthSessionIdHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthSessionIdHash.Get(), o.AuthSessionIdHash.IsSet()
}

// SetAuthSessionIdHash sets field value
func (o *SubmissionDataRequestShow) SetAuthSessionIdHash(v string) {
	o.AuthSessionIdHash.Set(&v)
}

// GetAuthUserIdHash returns the AuthUserIdHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetAuthUserIdHash() string {
	if o == nil || o.AuthUserIdHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthUserIdHash.Get()
}

// GetAuthUserIdHashOk returns a tuple with the AuthUserIdHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetAuthUserIdHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUserIdHash.Get(), o.AuthUserIdHash.IsSet()
}

// SetAuthUserIdHash sets field value
func (o *SubmissionDataRequestShow) SetAuthUserIdHash(v string) {
	o.AuthUserIdHash.Set(&v)
}

// GetAuthUsernameHash returns the AuthUsernameHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetAuthUsernameHash() string {
	if o == nil || o.AuthUsernameHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthUsernameHash.Get()
}

// GetAuthUsernameHashOk returns a tuple with the AuthUsernameHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetAuthUsernameHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUsernameHash.Get(), o.AuthUsernameHash.IsSet()
}

// SetAuthUsernameHash sets field value
func (o *SubmissionDataRequestShow) SetAuthUsernameHash(v string) {
	o.AuthUsernameHash.Set(&v)
}

// GetAuthPhoneNumberHash returns the AuthPhoneNumberHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetAuthPhoneNumberHash() string {
	if o == nil || o.AuthPhoneNumberHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthPhoneNumberHash.Get()
}

// GetAuthPhoneNumberHashOk returns a tuple with the AuthPhoneNumberHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetAuthPhoneNumberHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthPhoneNumberHash.Get(), o.AuthPhoneNumberHash.IsSet()
}

// SetAuthPhoneNumberHash sets field value
func (o *SubmissionDataRequestShow) SetAuthPhoneNumberHash(v string) {
	o.AuthPhoneNumberHash.Set(&v)
}

// GetIpAddress returns the IpAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetIpAddress() string {
	if o == nil || o.IpAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// SetIpAddress sets field value
func (o *SubmissionDataRequestShow) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}

// GetUserAgent returns the UserAgent field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetUserAgent() string {
	if o == nil || o.UserAgent.Get() == nil {
		var ret string
		return ret
	}

	return *o.UserAgent.Get()
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserAgent.Get(), o.UserAgent.IsSet()
}

// SetUserAgent sets field value
func (o *SubmissionDataRequestShow) SetUserAgent(v string) {
	o.UserAgent.Set(&v)
}

// GetSubmissionId returns the SubmissionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestShow) GetSubmissionId() string {
	if o == nil || o.SubmissionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubmissionId.Get()
}

// GetSubmissionIdOk returns a tuple with the SubmissionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestShow) GetSubmissionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubmissionId.Get(), o.SubmissionId.IsSet()
}

// SetSubmissionId sets field value
func (o *SubmissionDataRequestShow) SetSubmissionId(v string) {
	o.SubmissionId.Set(&v)
}

func (o SubmissionDataRequestShow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmissionDataRequestShow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["email"] = o.Email.Get()
	toSerialize["name"] = o.Name.Get()
	toSerialize["order"] = o.Order.Get()
	toSerialize["sort_order"] = o.SortOrder
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["state"] = o.State
	toSerialize["viewed_at"] = o.ViewedAt.Get()
	toSerialize["completed_at"] = o.CompletedAt.Get()
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["auth_type"] = o.AuthType
	toSerialize["auth_second_factor_type"] = o.AuthSecondFactorType
	toSerialize["auth_provider"] = o.AuthProvider.Get()
	toSerialize["auth_session_started_at"] = o.AuthSessionStartedAt.Get()
	toSerialize["auth_session_id_hash"] = o.AuthSessionIdHash.Get()
	toSerialize["auth_user_id_hash"] = o.AuthUserIdHash.Get()
	toSerialize["auth_username_hash"] = o.AuthUsernameHash.Get()
	toSerialize["auth_phone_number_hash"] = o.AuthPhoneNumberHash.Get()
	toSerialize["ip_address"] = o.IpAddress.Get()
	toSerialize["user_agent"] = o.UserAgent.Get()
	toSerialize["submission_id"] = o.SubmissionId.Get()
	return toSerialize, nil
}

func (o *SubmissionDataRequestShow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"email",
		"name",
		"order",
		"sort_order",
		"fields",
		"metadata",
		"state",
		"viewed_at",
		"completed_at",
		"data",
		"auth_type",
		"auth_second_factor_type",
		"auth_provider",
		"auth_session_started_at",
		"auth_session_id_hash",
		"auth_user_id_hash",
		"auth_username_hash",
		"auth_phone_number_hash",
		"ip_address",
		"user_agent",
		"submission_id",
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

	varSubmissionDataRequestShow := _SubmissionDataRequestShow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubmissionDataRequestShow)

	if err != nil {
		return err
	}

	*o = SubmissionDataRequestShow(varSubmissionDataRequestShow)

	return err
}

type NullableSubmissionDataRequestShow struct {
	value *SubmissionDataRequestShow
	isSet bool
}

func (v NullableSubmissionDataRequestShow) Get() *SubmissionDataRequestShow {
	return v.value
}

func (v *NullableSubmissionDataRequestShow) Set(val *SubmissionDataRequestShow) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmissionDataRequestShow) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmissionDataRequestShow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmissionDataRequestShow(val *SubmissionDataRequestShow) *NullableSubmissionDataRequestShow {
	return &NullableSubmissionDataRequestShow{value: val, isSet: true}
}

func (v NullableSubmissionDataRequestShow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmissionDataRequestShow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


