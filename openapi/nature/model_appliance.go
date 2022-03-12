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

// Appliance struct for Appliance
type Appliance struct {
	Id *string `json:"id,omitempty"`
	Device *DeviceCore `json:"device,omitempty"`
	Model *ApplianceModel `json:"model,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	// Basename of the image file included in the app. Ex: \"ico_ac_1\" 
	Image *string `json:"image,omitempty"`
	Type *ApplianceType `json:"type,omitempty"`
	Settings *AirConParams `json:"settings,omitempty"`
	Aircon *AirCon `json:"aircon,omitempty"`
	Signals []Signal `json:"signals,omitempty"`
	Tv *TV `json:"tv,omitempty"`
	Light *LIGHT `json:"light,omitempty"`
	SmartMeter *SmartMeter `json:"smart_meter,omitempty"`
}

// NewAppliance instantiates a new Appliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliance() *Appliance {
	this := Appliance{}
	var type_ ApplianceType = APPLIANCETYPE_AC
	this.Type = &type_
	return &this
}

// NewApplianceWithDefaults instantiates a new Appliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceWithDefaults() *Appliance {
	this := Appliance{}
	var type_ ApplianceType = APPLIANCETYPE_AC
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Appliance) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Appliance) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Appliance) SetId(v string) {
	o.Id = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *Appliance) GetDevice() DeviceCore {
	if o == nil || o.Device == nil {
		var ret DeviceCore
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetDeviceOk() (*DeviceCore, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *Appliance) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given DeviceCore and assigns it to the Device field.
func (o *Appliance) SetDevice(v DeviceCore) {
	o.Device = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *Appliance) GetModel() ApplianceModel {
	if o == nil || o.Model == nil {
		var ret ApplianceModel
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetModelOk() (*ApplianceModel, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *Appliance) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given ApplianceModel and assigns it to the Model field.
func (o *Appliance) SetModel(v ApplianceModel) {
	o.Model = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *Appliance) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *Appliance) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *Appliance) SetNickname(v string) {
	o.Nickname = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Appliance) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Appliance) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *Appliance) SetImage(v string) {
	o.Image = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Appliance) GetType() ApplianceType {
	if o == nil || o.Type == nil {
		var ret ApplianceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetTypeOk() (*ApplianceType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Appliance) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ApplianceType and assigns it to the Type field.
func (o *Appliance) SetType(v ApplianceType) {
	o.Type = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Appliance) GetSettings() AirConParams {
	if o == nil || o.Settings == nil {
		var ret AirConParams
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetSettingsOk() (*AirConParams, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Appliance) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AirConParams and assigns it to the Settings field.
func (o *Appliance) SetSettings(v AirConParams) {
	o.Settings = &v
}

// GetAircon returns the Aircon field value if set, zero value otherwise.
func (o *Appliance) GetAircon() AirCon {
	if o == nil || o.Aircon == nil {
		var ret AirCon
		return ret
	}
	return *o.Aircon
}

// GetAirconOk returns a tuple with the Aircon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetAirconOk() (*AirCon, bool) {
	if o == nil || o.Aircon == nil {
		return nil, false
	}
	return o.Aircon, true
}

// HasAircon returns a boolean if a field has been set.
func (o *Appliance) HasAircon() bool {
	if o != nil && o.Aircon != nil {
		return true
	}

	return false
}

// SetAircon gets a reference to the given AirCon and assigns it to the Aircon field.
func (o *Appliance) SetAircon(v AirCon) {
	o.Aircon = &v
}

// GetSignals returns the Signals field value if set, zero value otherwise.
func (o *Appliance) GetSignals() []Signal {
	if o == nil || o.Signals == nil {
		var ret []Signal
		return ret
	}
	return o.Signals
}

// GetSignalsOk returns a tuple with the Signals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetSignalsOk() ([]Signal, bool) {
	if o == nil || o.Signals == nil {
		return nil, false
	}
	return o.Signals, true
}

// HasSignals returns a boolean if a field has been set.
func (o *Appliance) HasSignals() bool {
	if o != nil && o.Signals != nil {
		return true
	}

	return false
}

// SetSignals gets a reference to the given []Signal and assigns it to the Signals field.
func (o *Appliance) SetSignals(v []Signal) {
	o.Signals = v
}

// GetTv returns the Tv field value if set, zero value otherwise.
func (o *Appliance) GetTv() TV {
	if o == nil || o.Tv == nil {
		var ret TV
		return ret
	}
	return *o.Tv
}

// GetTvOk returns a tuple with the Tv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetTvOk() (*TV, bool) {
	if o == nil || o.Tv == nil {
		return nil, false
	}
	return o.Tv, true
}

// HasTv returns a boolean if a field has been set.
func (o *Appliance) HasTv() bool {
	if o != nil && o.Tv != nil {
		return true
	}

	return false
}

// SetTv gets a reference to the given TV and assigns it to the Tv field.
func (o *Appliance) SetTv(v TV) {
	o.Tv = &v
}

// GetLight returns the Light field value if set, zero value otherwise.
func (o *Appliance) GetLight() LIGHT {
	if o == nil || o.Light == nil {
		var ret LIGHT
		return ret
	}
	return *o.Light
}

// GetLightOk returns a tuple with the Light field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetLightOk() (*LIGHT, bool) {
	if o == nil || o.Light == nil {
		return nil, false
	}
	return o.Light, true
}

// HasLight returns a boolean if a field has been set.
func (o *Appliance) HasLight() bool {
	if o != nil && o.Light != nil {
		return true
	}

	return false
}

// SetLight gets a reference to the given LIGHT and assigns it to the Light field.
func (o *Appliance) SetLight(v LIGHT) {
	o.Light = &v
}

// GetSmartMeter returns the SmartMeter field value if set, zero value otherwise.
func (o *Appliance) GetSmartMeter() SmartMeter {
	if o == nil || o.SmartMeter == nil {
		var ret SmartMeter
		return ret
	}
	return *o.SmartMeter
}

// GetSmartMeterOk returns a tuple with the SmartMeter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appliance) GetSmartMeterOk() (*SmartMeter, bool) {
	if o == nil || o.SmartMeter == nil {
		return nil, false
	}
	return o.SmartMeter, true
}

// HasSmartMeter returns a boolean if a field has been set.
func (o *Appliance) HasSmartMeter() bool {
	if o != nil && o.SmartMeter != nil {
		return true
	}

	return false
}

// SetSmartMeter gets a reference to the given SmartMeter and assigns it to the SmartMeter field.
func (o *Appliance) SetSmartMeter(v SmartMeter) {
	o.SmartMeter = &v
}

func (o Appliance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Aircon != nil {
		toSerialize["aircon"] = o.Aircon
	}
	if o.Signals != nil {
		toSerialize["signals"] = o.Signals
	}
	if o.Tv != nil {
		toSerialize["tv"] = o.Tv
	}
	if o.Light != nil {
		toSerialize["light"] = o.Light
	}
	if o.SmartMeter != nil {
		toSerialize["smart_meter"] = o.SmartMeter
	}
	return json.Marshal(toSerialize)
}

type NullableAppliance struct {
	value *Appliance
	isSet bool
}

func (v NullableAppliance) Get() *Appliance {
	return v.value
}

func (v *NullableAppliance) Set(val *Appliance) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliance) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliance(val *Appliance) *NullableAppliance {
	return &NullableAppliance{value: val, isSet: true}
}

func (v NullableAppliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

