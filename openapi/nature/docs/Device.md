# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TemperatureOffset** | Pointer to **float32** |  | [optional] 
**HumidityOffset** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**FirmwareVersion** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**NewestEvents** | Pointer to [**DeviceNewestEvents**](DeviceNewestEvents.md) |  | [optional] 

## Methods

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Device) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Device) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemperatureOffset

`func (o *Device) GetTemperatureOffset() float32`

GetTemperatureOffset returns the TemperatureOffset field if non-nil, zero value otherwise.

### GetTemperatureOffsetOk

`func (o *Device) GetTemperatureOffsetOk() (*float32, bool)`

GetTemperatureOffsetOk returns a tuple with the TemperatureOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureOffset

`func (o *Device) SetTemperatureOffset(v float32)`

SetTemperatureOffset sets TemperatureOffset field to given value.

### HasTemperatureOffset

`func (o *Device) HasTemperatureOffset() bool`

HasTemperatureOffset returns a boolean if a field has been set.

### GetHumidityOffset

`func (o *Device) GetHumidityOffset() float32`

GetHumidityOffset returns the HumidityOffset field if non-nil, zero value otherwise.

### GetHumidityOffsetOk

`func (o *Device) GetHumidityOffsetOk() (*float32, bool)`

GetHumidityOffsetOk returns a tuple with the HumidityOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidityOffset

`func (o *Device) SetHumidityOffset(v float32)`

SetHumidityOffset sets HumidityOffset field to given value.

### HasHumidityOffset

`func (o *Device) HasHumidityOffset() bool`

HasHumidityOffset returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Device) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Device) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Device) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Device) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Device) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Device) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Device) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Device) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *Device) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *Device) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *Device) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *Device) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetMacAddress

`func (o *Device) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Device) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Device) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *Device) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Device) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Device) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Device) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Device) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetNewestEvents

`func (o *Device) GetNewestEvents() DeviceNewestEvents`

GetNewestEvents returns the NewestEvents field if non-nil, zero value otherwise.

### GetNewestEventsOk

`func (o *Device) GetNewestEventsOk() (*DeviceNewestEvents, bool)`

GetNewestEventsOk returns a tuple with the NewestEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewestEvents

`func (o *Device) SetNewestEvents(v DeviceNewestEvents)`

SetNewestEvents sets NewestEvents field to given value.

### HasNewestEvents

`func (o *Device) HasNewestEvents() bool`

HasNewestEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


