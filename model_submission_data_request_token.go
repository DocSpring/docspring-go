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

// checks if the SubmissionDataRequestToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmissionDataRequestToken{}

// SubmissionDataRequestToken struct for SubmissionDataRequestToken
type SubmissionDataRequestToken struct {
	Id NullableString `json:"id"`
	Secret string `json:"secret"`
	ExpiresAt NullableString `json:"expires_at"`
	DataRequestUrl NullableString `json:"data_request_url"`
}

type _SubmissionDataRequestToken SubmissionDataRequestToken

// NewSubmissionDataRequestToken instantiates a new SubmissionDataRequestToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmissionDataRequestToken(id NullableString, secret string, expiresAt NullableString, dataRequestUrl NullableString) *SubmissionDataRequestToken {
	this := SubmissionDataRequestToken{}
	this.Id = id
	this.Secret = secret
	this.ExpiresAt = expiresAt
	this.DataRequestUrl = dataRequestUrl
	return &this
}

// NewSubmissionDataRequestTokenWithDefaults instantiates a new SubmissionDataRequestToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmissionDataRequestTokenWithDefaults() *SubmissionDataRequestToken {
	this := SubmissionDataRequestToken{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestToken) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestToken) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *SubmissionDataRequestToken) SetId(v string) {
	o.Id.Set(&v)
}

// GetSecret returns the Secret field value
func (o *SubmissionDataRequestToken) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *SubmissionDataRequestToken) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *SubmissionDataRequestToken) SetSecret(v string) {
	o.Secret = v
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestToken) GetExpiresAt() string {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestToken) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *SubmissionDataRequestToken) SetExpiresAt(v string) {
	o.ExpiresAt.Set(&v)
}

// GetDataRequestUrl returns the DataRequestUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubmissionDataRequestToken) GetDataRequestUrl() string {
	if o == nil || o.DataRequestUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.DataRequestUrl.Get()
}

// GetDataRequestUrlOk returns a tuple with the DataRequestUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmissionDataRequestToken) GetDataRequestUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataRequestUrl.Get(), o.DataRequestUrl.IsSet()
}

// SetDataRequestUrl sets field value
func (o *SubmissionDataRequestToken) SetDataRequestUrl(v string) {
	o.DataRequestUrl.Set(&v)
}

func (o SubmissionDataRequestToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmissionDataRequestToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["secret"] = o.Secret
	toSerialize["expires_at"] = o.ExpiresAt.Get()
	toSerialize["data_request_url"] = o.DataRequestUrl.Get()
	return toSerialize, nil
}

func (o *SubmissionDataRequestToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"secret",
		"expires_at",
		"data_request_url",
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

	varSubmissionDataRequestToken := _SubmissionDataRequestToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubmissionDataRequestToken)

	if err != nil {
		return err
	}

	*o = SubmissionDataRequestToken(varSubmissionDataRequestToken)

	return err
}

type NullableSubmissionDataRequestToken struct {
	value *SubmissionDataRequestToken
	isSet bool
}

func (v NullableSubmissionDataRequestToken) Get() *SubmissionDataRequestToken {
	return v.value
}

func (v *NullableSubmissionDataRequestToken) Set(val *SubmissionDataRequestToken) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmissionDataRequestToken) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmissionDataRequestToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmissionDataRequestToken(val *SubmissionDataRequestToken) *NullableSubmissionDataRequestToken {
	return &NullableSubmissionDataRequestToken{value: val, isSet: true}
}

func (v NullableSubmissionDataRequestToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmissionDataRequestToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


