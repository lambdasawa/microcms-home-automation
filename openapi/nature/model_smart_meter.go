/*
Nature API

Read/Write Nature Remo

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SmartMeter struct for SmartMeter
type SmartMeter struct {
	EchonetliteProperties []EchonetLiteProperty `json:"echonetlite_properties,omitempty"`
}

// NewSmartMeter instantiates a new SmartMeter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartMeter() *SmartMeter {
	this := SmartMeter{}
	return &this
}

// NewSmartMeterWithDefaults instantiates a new SmartMeter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartMeterWithDefaults() *SmartMeter {
	this := SmartMeter{}
	return &this
}

// GetEchonetliteProperties returns the EchonetliteProperties field value if set, zero value otherwise.
func (o *SmartMeter) GetEchonetliteProperties() []EchonetLiteProperty {
	if o == nil || o.EchonetliteProperties == nil {
		var ret []EchonetLiteProperty
		return ret
	}
	return o.EchonetliteProperties
}

// GetEchonetlitePropertiesOk returns a tuple with the EchonetliteProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartMeter) GetEchonetlitePropertiesOk() ([]EchonetLiteProperty, bool) {
	if o == nil || o.EchonetliteProperties == nil {
		return nil, false
	}
	return o.EchonetliteProperties, true
}

// HasEchonetliteProperties returns a boolean if a field has been set.
func (o *SmartMeter) HasEchonetliteProperties() bool {
	if o != nil && o.EchonetliteProperties != nil {
		return true
	}

	return false
}

// SetEchonetliteProperties gets a reference to the given []EchonetLiteProperty and assigns it to the EchonetliteProperties field.
func (o *SmartMeter) SetEchonetliteProperties(v []EchonetLiteProperty) {
	o.EchonetliteProperties = v
}

func (o SmartMeter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EchonetliteProperties != nil {
		toSerialize["echonetlite_properties"] = o.EchonetliteProperties
	}
	return json.Marshal(toSerialize)
}

type NullableSmartMeter struct {
	value *SmartMeter
	isSet bool
}

func (v NullableSmartMeter) Get() *SmartMeter {
	return v.value
}

func (v *NullableSmartMeter) Set(val *SmartMeter) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartMeter) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartMeter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartMeter(val *SmartMeter) *NullableSmartMeter {
	return &NullableSmartMeter{value: val, isSet: true}
}

func (v NullableSmartMeter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartMeter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


