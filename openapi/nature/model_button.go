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

// Button struct for Button
type Button struct {
	// Name of button. It is used for \"POST /1/{applaince}/tv\" and \"POST /1/{appliance}/light\".
	Name *string `json:"name,omitempty"`
	// Basename of the image file included in the app. Ex: \"ico_ac_1\" 
	Image *string `json:"image,omitempty"`
	// Label of button.
	Label *string `json:"label,omitempty"`
}

// NewButton instantiates a new Button object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewButton() *Button {
	this := Button{}
	return &this
}

// NewButtonWithDefaults instantiates a new Button object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewButtonWithDefaults() *Button {
	this := Button{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Button) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Button) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Button) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Button) SetName(v string) {
	o.Name = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Button) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Button) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Button) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *Button) SetImage(v string) {
	o.Image = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Button) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Button) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Button) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Button) SetLabel(v string) {
	o.Label = &v
}

func (o Button) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableButton struct {
	value *Button
	isSet bool
}

func (v NullableButton) Get() *Button {
	return v.value
}

func (v *NullableButton) Set(val *Button) {
	v.value = val
	v.isSet = true
}

func (v NullableButton) IsSet() bool {
	return v.isSet
}

func (v *NullableButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableButton(val *Button) *NullableButton {
	return &NullableButton{value: val, isSet: true}
}

func (v NullableButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

