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

// checks if the CombinedSubmissionAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CombinedSubmissionAction{}

// CombinedSubmissionAction struct for CombinedSubmissionAction
type CombinedSubmissionAction struct {
	Id NullableString `json:"id"`
	IntegrationId NullableString `json:"integration_id"`
	State string `json:"state"`
	ActionType string `json:"action_type"`
	ActionCategory string `json:"action_category"`
	ResultData map[string]interface{} `json:"result_data"`
}

type _CombinedSubmissionAction CombinedSubmissionAction

// NewCombinedSubmissionAction instantiates a new CombinedSubmissionAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCombinedSubmissionAction(id NullableString, integrationId NullableString, state string, actionType string, actionCategory string, resultData map[string]interface{}) *CombinedSubmissionAction {
	this := CombinedSubmissionAction{}
	this.Id = id
	this.IntegrationId = integrationId
	this.State = state
	this.ActionType = actionType
	this.ActionCategory = actionCategory
	this.ResultData = resultData
	return &this
}

// NewCombinedSubmissionActionWithDefaults instantiates a new CombinedSubmissionAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCombinedSubmissionActionWithDefaults() *CombinedSubmissionAction {
	this := CombinedSubmissionAction{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CombinedSubmissionAction) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CombinedSubmissionAction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *CombinedSubmissionAction) SetId(v string) {
	o.Id.Set(&v)
}

// GetIntegrationId returns the IntegrationId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CombinedSubmissionAction) GetIntegrationId() string {
	if o == nil || o.IntegrationId.Get() == nil {
		var ret string
		return ret
	}

	return *o.IntegrationId.Get()
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CombinedSubmissionAction) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntegrationId.Get(), o.IntegrationId.IsSet()
}

// SetIntegrationId sets field value
func (o *CombinedSubmissionAction) SetIntegrationId(v string) {
	o.IntegrationId.Set(&v)
}

// GetState returns the State field value
func (o *CombinedSubmissionAction) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CombinedSubmissionAction) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CombinedSubmissionAction) SetState(v string) {
	o.State = v
}

// GetActionType returns the ActionType field value
func (o *CombinedSubmissionAction) GetActionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *CombinedSubmissionAction) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *CombinedSubmissionAction) SetActionType(v string) {
	o.ActionType = v
}

// GetActionCategory returns the ActionCategory field value
func (o *CombinedSubmissionAction) GetActionCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionCategory
}

// GetActionCategoryOk returns a tuple with the ActionCategory field value
// and a boolean to check if the value has been set.
func (o *CombinedSubmissionAction) GetActionCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionCategory, true
}

// SetActionCategory sets field value
func (o *CombinedSubmissionAction) SetActionCategory(v string) {
	o.ActionCategory = v
}

// GetResultData returns the ResultData field value
func (o *CombinedSubmissionAction) GetResultData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ResultData
}

// GetResultDataOk returns a tuple with the ResultData field value
// and a boolean to check if the value has been set.
func (o *CombinedSubmissionAction) GetResultDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ResultData, true
}

// SetResultData sets field value
func (o *CombinedSubmissionAction) SetResultData(v map[string]interface{}) {
	o.ResultData = v
}

func (o CombinedSubmissionAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CombinedSubmissionAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["integration_id"] = o.IntegrationId.Get()
	toSerialize["state"] = o.State
	toSerialize["action_type"] = o.ActionType
	toSerialize["action_category"] = o.ActionCategory
	toSerialize["result_data"] = o.ResultData
	return toSerialize, nil
}

func (o *CombinedSubmissionAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"integration_id",
		"state",
		"action_type",
		"action_category",
		"result_data",
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

	varCombinedSubmissionAction := _CombinedSubmissionAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCombinedSubmissionAction)

	if err != nil {
		return err
	}

	*o = CombinedSubmissionAction(varCombinedSubmissionAction)

	return err
}

type NullableCombinedSubmissionAction struct {
	value *CombinedSubmissionAction
	isSet bool
}

func (v NullableCombinedSubmissionAction) Get() *CombinedSubmissionAction {
	return v.value
}

func (v *NullableCombinedSubmissionAction) Set(val *CombinedSubmissionAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCombinedSubmissionAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCombinedSubmissionAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCombinedSubmissionAction(val *CombinedSubmissionAction) *NullableCombinedSubmissionAction {
	return &NullableCombinedSubmissionAction{value: val, isSet: true}
}

func (v NullableCombinedSubmissionAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCombinedSubmissionAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


