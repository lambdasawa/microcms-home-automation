/*
Nature API

Read/Write Nature Remo

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AirVolume Empty means automatic. Numbers express the amount of volume. The range of AirVolumes which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
type AirVolume string

// List of AirVolume
const (
	AIRVOLUME_EMPTY AirVolume = ""
	AIRVOLUME_AUTO AirVolume = "auto"
	AIRVOLUME__1 AirVolume = "1"
	AIRVOLUME__2 AirVolume = "2"
	AIRVOLUME__3 AirVolume = "3"
	AIRVOLUME__4 AirVolume = "4"
	AIRVOLUME__5 AirVolume = "5"
	AIRVOLUME__6 AirVolume = "6"
	AIRVOLUME__7 AirVolume = "7"
	AIRVOLUME__8 AirVolume = "8"
	AIRVOLUME__9 AirVolume = "9"
	AIRVOLUME__10 AirVolume = "10"
)

// All allowed values of AirVolume enum
var AllowedAirVolumeEnumValues = []AirVolume{
	"",
	"auto",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
}

func (v *AirVolume) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AirVolume(value)
	for _, existing := range AllowedAirVolumeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AirVolume", value)
}

// NewAirVolumeFromValue returns a pointer to a valid AirVolume
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAirVolumeFromValue(v string) (*AirVolume, error) {
	ev := AirVolume(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AirVolume: valid values are %v", v, AllowedAirVolumeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AirVolume) IsValid() bool {
	for _, existing := range AllowedAirVolumeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AirVolume value
func (v AirVolume) Ptr() *AirVolume {
	return &v
}

type NullableAirVolume struct {
	value *AirVolume
	isSet bool
}

func (v NullableAirVolume) Get() *AirVolume {
	return v.value
}

func (v *NullableAirVolume) Set(val *AirVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableAirVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableAirVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAirVolume(val *AirVolume) *NullableAirVolume {
	return &NullableAirVolume{value: val, isSet: true}
}

func (v NullableAirVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAirVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

