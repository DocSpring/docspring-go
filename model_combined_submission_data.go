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

// checks if the CombinedSubmissionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CombinedSubmissionData{}

// CombinedSubmissionData struct for CombinedSubmissionData
type CombinedSubmissionData struct {
	ExpiresIn *int32 `json:"expires_in,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Password *string `json:"password,omitempty"`
	SubmissionIds []string `json:"submission_ids"`
	Test *bool `json:"test,omitempty"`
}

type _CombinedSubmissionData CombinedSubmissionData

// NewCombinedSubmissionData instantiates a new CombinedSubmissionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCombinedSubmissionData(submissionIds []string) *CombinedSubmissionData {
	this := CombinedSubmissionData{}
	this.SubmissionIds = submissionIds
	return &this
}

// NewCombinedSubmissionDataWithDefaults instantiates a new CombinedSubmissionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCombinedSubmissionDataWithDefaults() *CombinedSubmissionData {
	this := CombinedSubmissionData{}
	return &this
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *CombinedSubmissionData) GetExpiresIn() int32 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedSubmissionData) GetExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *CombinedSubmissionData) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *CombinedSubmissionData) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CombinedSubmissionData) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedSubmissionData) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CombinedSubmissionData) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CombinedSubmissionData) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CombinedSubmissionData) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedSubmissionData) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CombinedSubmissionData) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CombinedSubmissionData) SetPassword(v string) {
	o.Password = &v
}

// GetSubmissionIds returns the SubmissionIds field value
func (o *CombinedSubmissionData) GetSubmissionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubmissionIds
}

// GetSubmissionIdsOk returns a tuple with the SubmissionIds field value
// and a boolean to check if the value has been set.
func (o *CombinedSubmissionData) GetSubmissionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubmissionIds, true
}

// SetSubmissionIds sets field value
func (o *CombinedSubmissionData) SetSubmissionIds(v []string) {
	o.SubmissionIds = v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *CombinedSubmissionData) GetTest() bool {
	if o == nil || IsNil(o.Test) {
		var ret bool
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedSubmissionData) GetTestOk() (*bool, bool) {
	if o == nil || IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *CombinedSubmissionData) HasTest() bool {
	if o != nil && !IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given bool and assigns it to the Test field.
func (o *CombinedSubmissionData) SetTest(v bool) {
	o.Test = &v
}

func (o CombinedSubmissionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CombinedSubmissionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["submission_ids"] = o.SubmissionIds
	if !IsNil(o.Test) {
		toSerialize["test"] = o.Test
	}
	return toSerialize, nil
}

func (o *CombinedSubmissionData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"submission_ids",
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

	varCombinedSubmissionData := _CombinedSubmissionData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCombinedSubmissionData)

	if err != nil {
		return err
	}

	*o = CombinedSubmissionData(varCombinedSubmissionData)

	return err
}

type NullableCombinedSubmissionData struct {
	value *CombinedSubmissionData
	isSet bool
}

func (v NullableCombinedSubmissionData) Get() *CombinedSubmissionData {
	return v.value
}

func (v *NullableCombinedSubmissionData) Set(val *CombinedSubmissionData) {
	v.value = val
	v.isSet = true
}

func (v NullableCombinedSubmissionData) IsSet() bool {
	return v.isSet
}

func (v *NullableCombinedSubmissionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCombinedSubmissionData(val *CombinedSubmissionData) *NullableCombinedSubmissionData {
	return &NullableCombinedSubmissionData{value: val, isSet: true}
}

func (v NullableCombinedSubmissionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCombinedSubmissionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


