# Appliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Device** | Pointer to [**DeviceCore**](DeviceCore.md) |  | [optional] 
**Model** | Pointer to [**ApplianceModel**](ApplianceModel.md) |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** | Basename of the image file included in the app. Ex: \&quot;ico_ac_1\&quot;  | [optional] 
**Type** | Pointer to [**ApplianceType**](ApplianceType.md) |  | [optional] [default to APPLIANCETYPE_AC]
**Settings** | Pointer to [**AirConParams**](AirConParams.md) |  | [optional] 
**Aircon** | Pointer to [**AirCon**](AirCon.md) |  | [optional] 
**Signals** | Pointer to [**[]Signal**](Signal.md) |  | [optional] 
**Tv** | Pointer to [**TV**](TV.md) |  | [optional] 
**Light** | Pointer to [**LIGHT**](LIGHT.md) |  | [optional] 
**SmartMeter** | Pointer to [**SmartMeter**](SmartMeter.md) |  | [optional] 

## Methods

### NewAppliance

`func NewAppliance() *Appliance`

NewAppliance instantiates a new Appliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceWithDefaults

`func NewApplianceWithDefaults() *Appliance`

NewApplianceWithDefaults instantiates a new Appliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Appliance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Appliance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Appliance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Appliance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDevice

`func (o *Appliance) GetDevice() DeviceCore`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Appliance) GetDeviceOk() (*DeviceCore, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Appliance) SetDevice(v DeviceCore)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *Appliance) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetModel

`func (o *Appliance) GetModel() ApplianceModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Appliance) GetModelOk() (*ApplianceModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Appliance) SetModel(v ApplianceModel)`

SetModel sets Model field to given value.

### HasModel

`func (o *Appliance) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNickname

`func (o *Appliance) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *Appliance) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *Appliance) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *Appliance) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetImage

`func (o *Appliance) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Appliance) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Appliance) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Appliance) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *Appliance) GetType() ApplianceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Appliance) GetTypeOk() (*ApplianceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Appliance) SetType(v ApplianceType)`

SetType sets Type field to given value.

### HasType

`func (o *Appliance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSettings

`func (o *Appliance) GetSettings() AirConParams`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Appliance) GetSettingsOk() (*AirConParams, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Appliance) SetSettings(v AirConParams)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Appliance) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetAircon

`func (o *Appliance) GetAircon() AirCon`

GetAircon returns the Aircon field if non-nil, zero value otherwise.

### GetAirconOk

`func (o *Appliance) GetAirconOk() (*AirCon, bool)`

GetAirconOk returns a tuple with the Aircon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAircon

`func (o *Appliance) SetAircon(v AirCon)`

SetAircon sets Aircon field to given value.

### HasAircon

`func (o *Appliance) HasAircon() bool`

HasAircon returns a boolean if a field has been set.

### GetSignals

`func (o *Appliance) GetSignals() []Signal`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *Appliance) GetSignalsOk() (*[]Signal, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *Appliance) SetSignals(v []Signal)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *Appliance) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### GetTv

`func (o *Appliance) GetTv() TV`

GetTv returns the Tv field if non-nil, zero value otherwise.

### GetTvOk

`func (o *Appliance) GetTvOk() (*TV, bool)`

GetTvOk returns a tuple with the Tv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTv

`func (o *Appliance) SetTv(v TV)`

SetTv sets Tv field to given value.

### HasTv

`func (o *Appliance) HasTv() bool`

HasTv returns a boolean if a field has been set.

### GetLight

`func (o *Appliance) GetLight() LIGHT`

GetLight returns the Light field if non-nil, zero value otherwise.

### GetLightOk

`func (o *Appliance) GetLightOk() (*LIGHT, bool)`

GetLightOk returns a tuple with the Light field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLight

`func (o *Appliance) SetLight(v LIGHT)`

SetLight sets Light field to given value.

### HasLight

`func (o *Appliance) HasLight() bool`

HasLight returns a boolean if a field has been set.

### GetSmartMeter

`func (o *Appliance) GetSmartMeter() SmartMeter`

GetSmartMeter returns the SmartMeter field if non-nil, zero value otherwise.

### GetSmartMeterOk

`func (o *Appliance) GetSmartMeterOk() (*SmartMeter, bool)`

GetSmartMeterOk returns a tuple with the SmartMeter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartMeter

`func (o *Appliance) SetSmartMeter(v SmartMeter)`

SetSmartMeter sets SmartMeter field to given value.

### HasSmartMeter

`func (o *Appliance) HasSmartMeter() bool`

HasSmartMeter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


