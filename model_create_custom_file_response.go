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

// checks if the CreateCustomFileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomFileResponse{}

// CreateCustomFileResponse struct for CreateCustomFileResponse
type CreateCustomFileResponse struct {
	Status string `json:"status"`
	CustomFile CustomFile `json:"custom_file"`
	Errors []string `json:"errors,omitempty"`
}

type _CreateCustomFileResponse CreateCustomFileResponse

// NewCreateCustomFileResponse instantiates a new CreateCustomFileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomFileResponse(status string, customFile CustomFile) *CreateCustomFileResponse {
	this := CreateCustomFileResponse{}
	this.Status = status
	this.CustomFile = customFile
	return &this
}

// NewCreateCustomFileResponseWithDefaults instantiates a new CreateCustomFileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomFileResponseWithDefaults() *CreateCustomFileResponse {
	this := CreateCustomFileResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *CreateCustomFileResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateCustomFileResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateCustomFileResponse) SetStatus(v string) {
	o.Status = v
}

// GetCustomFile returns the CustomFile field value
func (o *CreateCustomFileResponse) GetCustomFile() CustomFile {
	if o == nil {
		var ret CustomFile
		return ret
	}

	return o.CustomFile
}

// GetCustomFileOk returns a tuple with the CustomFile field value
// and a boolean to check if the value has been set.
func (o *CreateCustomFileResponse) GetCustomFileOk() (*CustomFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomFile, true
}

// SetCustomFile sets field value
func (o *CreateCustomFileResponse) SetCustomFile(v CustomFile) {
	o.CustomFile = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateCustomFileResponse) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomFileResponse) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateCustomFileResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *CreateCustomFileResponse) SetErrors(v []string) {
	o.Errors = v
}

func (o CreateCustomFileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomFileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["custom_file"] = o.CustomFile
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

func (o *CreateCustomFileResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"custom_file",
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

	varCreateCustomFileResponse := _CreateCustomFileResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCustomFileResponse)

	if err != nil {
		return err
	}

	*o = CreateCustomFileResponse(varCreateCustomFileResponse)

	return err
}

type NullableCreateCustomFileResponse struct {
	value *CreateCustomFileResponse
	isSet bool
}

func (v NullableCreateCustomFileResponse) Get() *CreateCustomFileResponse {
	return v.value
}

func (v *NullableCreateCustomFileResponse) Set(val *CreateCustomFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomFileResponse(val *CreateCustomFileResponse) *NullableCreateCustomFileResponse {
	return &NullableCreateCustomFileResponse{value: val, isSet: true}
}

func (v NullableCreateCustomFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


