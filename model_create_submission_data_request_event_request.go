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

// checks if the CreateSubmissionDataRequestEventRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubmissionDataRequestEventRequest{}

// CreateSubmissionDataRequestEventRequest struct for CreateSubmissionDataRequestEventRequest
type CreateSubmissionDataRequestEventRequest struct {
	EventType string `json:"event_type"`
	MessageType NullableString `json:"message_type,omitempty"`
}

type _CreateSubmissionDataRequestEventRequest CreateSubmissionDataRequestEventRequest

// NewCreateSubmissionDataRequestEventRequest instantiates a new CreateSubmissionDataRequestEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubmissionDataRequestEventRequest(eventType string) *CreateSubmissionDataRequestEventRequest {
	this := CreateSubmissionDataRequestEventRequest{}
	this.EventType = eventType
	return &this
}

// NewCreateSubmissionDataRequestEventRequestWithDefaults instantiates a new CreateSubmissionDataRequestEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubmissionDataRequestEventRequestWithDefaults() *CreateSubmissionDataRequestEventRequest {
	this := CreateSubmissionDataRequestEventRequest{}
	return &this
}

// GetEventType returns the EventType field value
func (o *CreateSubmissionDataRequestEventRequest) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *CreateSubmissionDataRequestEventRequest) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *CreateSubmissionDataRequestEventRequest) SetEventType(v string) {
	o.EventType = v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSubmissionDataRequestEventRequest) GetMessageType() string {
	if o == nil || IsNil(o.MessageType.Get()) {
		var ret string
		return ret
	}
	return *o.MessageType.Get()
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSubmissionDataRequestEventRequest) GetMessageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageType.Get(), o.MessageType.IsSet()
}

// HasMessageType returns a boolean if a field has been set.
func (o *CreateSubmissionDataRequestEventRequest) HasMessageType() bool {
	if o != nil && o.MessageType.IsSet() {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given NullableString and assigns it to the MessageType field.
func (o *CreateSubmissionDataRequestEventRequest) SetMessageType(v string) {
	o.MessageType.Set(&v)
}
// SetMessageTypeNil sets the value for MessageType to be an explicit nil
func (o *CreateSubmissionDataRequestEventRequest) SetMessageTypeNil() {
	o.MessageType.Set(nil)
}

// UnsetMessageType ensures that no value is present for MessageType, not even an explicit nil
func (o *CreateSubmissionDataRequestEventRequest) UnsetMessageType() {
	o.MessageType.Unset()
}

func (o CreateSubmissionDataRequestEventRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubmissionDataRequestEventRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_type"] = o.EventType
	if o.MessageType.IsSet() {
		toSerialize["message_type"] = o.MessageType.Get()
	}
	return toSerialize, nil
}

func (o *CreateSubmissionDataRequestEventRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_type",
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

	varCreateSubmissionDataRequestEventRequest := _CreateSubmissionDataRequestEventRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSubmissionDataRequestEventRequest)

	if err != nil {
		return err
	}

	*o = CreateSubmissionDataRequestEventRequest(varCreateSubmissionDataRequestEventRequest)

	return err
}

type NullableCreateSubmissionDataRequestEventRequest struct {
	value *CreateSubmissionDataRequestEventRequest
	isSet bool
}

func (v NullableCreateSubmissionDataRequestEventRequest) Get() *CreateSubmissionDataRequestEventRequest {
	return v.value
}

func (v *NullableCreateSubmissionDataRequestEventRequest) Set(val *CreateSubmissionDataRequestEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubmissionDataRequestEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubmissionDataRequestEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubmissionDataRequestEventRequest(val *CreateSubmissionDataRequestEventRequest) *NullableCreateSubmissionDataRequestEventRequest {
	return &NullableCreateSubmissionDataRequestEventRequest{value: val, isSet: true}
}

func (v NullableCreateSubmissionDataRequestEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubmissionDataRequestEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


