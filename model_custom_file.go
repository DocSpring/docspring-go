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

// checks if the CustomFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFile{}

// CustomFile struct for CustomFile
type CustomFile struct {
	Id NullableString `json:"id"`
	Url NullableString `json:"url"`
}

type _CustomFile CustomFile

// NewCustomFile instantiates a new CustomFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFile(id NullableString, url NullableString) *CustomFile {
	this := CustomFile{}
	this.Id = id
	this.Url = url
	return &this
}

// NewCustomFileWithDefaults instantiates a new CustomFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFileWithDefaults() *CustomFile {
	this := CustomFile{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomFile) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFile) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *CustomFile) SetId(v string) {
	o.Id.Set(&v)
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomFile) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFile) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *CustomFile) SetUrl(v string) {
	o.Url.Set(&v)
}

func (o CustomFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["url"] = o.Url.Get()
	return toSerialize, nil
}

func (o *CustomFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
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

	varCustomFile := _CustomFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomFile)

	if err != nil {
		return err
	}

	*o = CustomFile(varCustomFile)

	return err
}

type NullableCustomFile struct {
	value *CustomFile
	isSet bool
}

func (v NullableCustomFile) Get() *CustomFile {
	return v.value
}

func (v *NullableCustomFile) Set(val *CustomFile) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFile) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFile(val *CustomFile) *NullableCustomFile {
	return &NullableCustomFile{value: val, isSet: true}
}

func (v NullableCustomFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


