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

// RandafEventTrigger Describes how RANDAF should generate the report for the event
type RandafEventTrigger struct {
	RandafEventTriggerAnyOf *RandafEventTriggerAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RandafEventTrigger) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RandafEventTriggerAnyOf
	err = json.Unmarshal(data, &dst.RandafEventTriggerAnyOf);
	if err == nil {
		jsonRandafEventTriggerAnyOf, _ := json.Marshal(dst.RandafEventTriggerAnyOf)
		if string(jsonRandafEventTriggerAnyOf) == "{}" { // empty struct
			dst.RandafEventTriggerAnyOf = nil
		} else {
			return nil // data stored in dst.RandafEventTriggerAnyOf, return on the first match
		}
	} else {
		dst.RandafEventTriggerAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(RandafEventTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RandafEventTrigger) MarshalJSON() ([]byte, error) {
	if src.RandafEventTriggerAnyOf != nil {
		return json.Marshal(&src.RandafEventTriggerAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRandafEventTrigger struct {
	value *RandafEventTrigger
	isSet bool
}

func (v NullableRandafEventTrigger) Get() *RandafEventTrigger {
	return v.value
}

func (v *NullableRandafEventTrigger) Set(val *RandafEventTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableRandafEventTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableRandafEventTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRandafEventTrigger(val *RandafEventTrigger) *NullableRandafEventTrigger {
	return &NullableRandafEventTrigger{value: val, isSet: true}
}

func (v NullableRandafEventTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRandafEventTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


