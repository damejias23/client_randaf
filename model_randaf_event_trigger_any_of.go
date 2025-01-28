/*
Nrandaf_EventExposure

RANDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package randaf

import (
	"encoding/json"
	"fmt"
)

// RandafEventTriggerAnyOf the model 'RandafEventTriggerAnyOf'
type RandafEventTriggerAnyOf string

// List of RandafEventTrigger_anyOf
const (
	ONE_TIME RandafEventTriggerAnyOf = "ONE_TIME"
	CONTINUOUS RandafEventTriggerAnyOf = "CONTINUOUS"
	PERIODIC RandafEventTriggerAnyOf = "PERIODIC"
)

// All allowed values of RandafEventTriggerAnyOf enum
var AllowedRandafEventTriggerAnyOfEnumValues = []RandafEventTriggerAnyOf{
	"ONE_TIME",
	"CONTINUOUS",
	"PERIODIC",
}

func (v *RandafEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RandafEventTriggerAnyOf(value)
	for _, existing := range AllowedRandafEventTriggerAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RandafEventTriggerAnyOf", value)
}

// NewRandafEventTriggerAnyOfFromValue returns a pointer to a valid RandafEventTriggerAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRandafEventTriggerAnyOfFromValue(v string) (*RandafEventTriggerAnyOf, error) {
	ev := RandafEventTriggerAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RandafEventTriggerAnyOf: valid values are %v", v, AllowedRandafEventTriggerAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RandafEventTriggerAnyOf) IsValid() bool {
	for _, existing := range AllowedRandafEventTriggerAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RandafEventTrigger_anyOf value
func (v RandafEventTriggerAnyOf) Ptr() *RandafEventTriggerAnyOf {
	return &v
}

type NullableRandafEventTriggerAnyOf struct {
	value *RandafEventTriggerAnyOf
	isSet bool
}

func (v NullableRandafEventTriggerAnyOf) Get() *RandafEventTriggerAnyOf {
	return v.value
}

func (v *NullableRandafEventTriggerAnyOf) Set(val *RandafEventTriggerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRandafEventTriggerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRandafEventTriggerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRandafEventTriggerAnyOf(val *RandafEventTriggerAnyOf) *NullableRandafEventTriggerAnyOf {
	return &NullableRandafEventTriggerAnyOf{value: val, isSet: true}
}

func (v NullableRandafEventTriggerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRandafEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

