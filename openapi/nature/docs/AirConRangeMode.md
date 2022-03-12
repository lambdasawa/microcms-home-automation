# AirConRangeMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Temp** | Pointer to **[]string** |  | [optional] 
**Vol** | Pointer to [**[]AirVolume**](AirVolume.md) |  | [optional] 
**Dir** | Pointer to [**[]AirDirection**](AirDirection.md) |  | [optional] 

## Methods

### NewAirConRangeMode

`func NewAirConRangeMode() *AirConRangeMode`

NewAirConRangeMode instantiates a new AirConRangeMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirConRangeModeWithDefaults

`func NewAirConRangeModeWithDefaults() *AirConRangeMode`

NewAirConRangeModeWithDefaults instantiates a new AirConRangeMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemp

`func (o *AirConRangeMode) GetTemp() []string`

GetTemp returns the Temp field if non-nil, zero value otherwise.

### GetTempOk

`func (o *AirConRangeMode) GetTempOk() (*[]string, bool)`

GetTempOk returns a tuple with the Temp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemp

`func (o *AirConRangeMode) SetTemp(v []string)`

SetTemp sets Temp field to given value.

### HasTemp

`func (o *AirConRangeMode) HasTemp() bool`

HasTemp returns a boolean if a field has been set.

### GetVol

`func (o *AirConRangeMode) GetVol() []AirVolume`

GetVol returns the Vol field if non-nil, zero value otherwise.

### GetVolOk

`func (o *AirConRangeMode) GetVolOk() (*[]AirVolume, bool)`

GetVolOk returns a tuple with the Vol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol

`func (o *AirConRangeMode) SetVol(v []AirVolume)`

SetVol sets Vol field to given value.

### HasVol

`func (o *AirConRangeMode) HasVol() bool`

HasVol returns a boolean if a field has been set.

### GetDir

`func (o *AirConRangeMode) GetDir() []AirDirection`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *AirConRangeMode) GetDirOk() (*[]AirDirection, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *AirConRangeMode) SetDir(v []AirDirection)`

SetDir sets Dir field to given value.

### HasDir

`func (o *AirConRangeMode) HasDir() bool`

HasDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


