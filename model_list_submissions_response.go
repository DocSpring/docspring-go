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

// checks if the ListSubmissionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSubmissionsResponse{}

// ListSubmissionsResponse struct for ListSubmissionsResponse
type ListSubmissionsResponse struct {
	Submissions []Submission `json:"submissions"`
	Limit float32 `json:"limit"`
	NextCursor NullableString `json:"next_cursor"`
}

type _ListSubmissionsResponse ListSubmissionsResponse

// NewListSubmissionsResponse instantiates a new ListSubmissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSubmissionsResponse(submissions []Submission, limit float32, nextCursor NullableString) *ListSubmissionsResponse {
	this := ListSubmissionsResponse{}
	this.Submissions = submissions
	this.Limit = limit
	this.NextCursor = nextCursor
	return &this
}

// NewListSubmissionsResponseWithDefaults instantiates a new ListSubmissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSubmissionsResponseWithDefaults() *ListSubmissionsResponse {
	this := ListSubmissionsResponse{}
	return &this
}

// GetSubmissions returns the Submissions field value
func (o *ListSubmissionsResponse) GetSubmissions() []Submission {
	if o == nil {
		var ret []Submission
		return ret
	}

	return o.Submissions
}

// GetSubmissionsOk returns a tuple with the Submissions field value
// and a boolean to check if the value has been set.
func (o *ListSubmissionsResponse) GetSubmissionsOk() ([]Submission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Submissions, true
}

// SetSubmissions sets field value
func (o *ListSubmissionsResponse) SetSubmissions(v []Submission) {
	o.Submissions = v
}

// GetLimit returns the Limit field value
func (o *ListSubmissionsResponse) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListSubmissionsResponse) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListSubmissionsResponse) SetLimit(v float32) {
	o.Limit = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListSubmissionsResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListSubmissionsResponse) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *ListSubmissionsResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

func (o ListSubmissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSubmissionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["submissions"] = o.Submissions
	toSerialize["limit"] = o.Limit
	toSerialize["next_cursor"] = o.NextCursor.Get()
	return toSerialize, nil
}

func (o *ListSubmissionsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"submissions",
		"limit",
		"next_cursor",
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

	varListSubmissionsResponse := _ListSubmissionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListSubmissionsResponse)

	if err != nil {
		return err
	}

	*o = ListSubmissionsResponse(varListSubmissionsResponse)

	return err
}

type NullableListSubmissionsResponse struct {
	value *ListSubmissionsResponse
	isSet bool
}

func (v NullableListSubmissionsResponse) Get() *ListSubmissionsResponse {
	return v.value
}

func (v *NullableListSubmissionsResponse) Set(val *ListSubmissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListSubmissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListSubmissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSubmissionsResponse(val *ListSubmissionsResponse) *NullableListSubmissionsResponse {
	return &NullableListSubmissionsResponse{value: val, isSet: true}
}

func (v NullableListSubmissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSubmissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


