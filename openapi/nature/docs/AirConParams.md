# AirConParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Temp** | Pointer to **string** | The temperature in string format. The unit is described in Aircon object. The range of Temperatures which the air conditioner accepts depends on the air conditioner model and operation mode. Check the &#39;AirConRangeMode&#39; information in the response for the range of the particular air conditioner model and operation mode. | [optional] 
**Mode** | Pointer to [**OperationMode**](OperationMode.md) |  | [optional] 
**Vol** | Pointer to [**AirVolume**](AirVolume.md) |  | [optional] 
**Dir** | Pointer to [**AirDirection**](AirDirection.md) |  | [optional] 
**Button** | Pointer to [**ACButton**](ACButton.md) |  | [optional] 

## Methods

### NewAirConParams

`func NewAirConParams() *AirConParams`

NewAirConParams instantiates a new AirConParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirConParamsWithDefaults

`func NewAirConParamsWithDefaults() *AirConParams`

NewAirConParamsWithDefaults instantiates a new AirConParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemp

`func (o *AirConParams) GetTemp() string`

GetTemp returns the Temp field if non-nil, zero value otherwise.

### GetTempOk

`func (o *AirConParams) GetTempOk() (*string, bool)`

GetTempOk returns a tuple with the Temp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemp

`func (o *AirConParams) SetTemp(v string)`

SetTemp sets Temp field to given value.

### HasTemp

`func (o *AirConParams) HasTemp() bool`

HasTemp returns a boolean if a field has been set.

### GetMode

`func (o *AirConParams) GetMode() OperationMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AirConParams) GetModeOk() (*OperationMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AirConParams) SetMode(v OperationMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AirConParams) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetVol

`func (o *AirConParams) GetVol() AirVolume`

GetVol returns the Vol field if non-nil, zero value otherwise.

### GetVolOk

`func (o *AirConParams) GetVolOk() (*AirVolume, bool)`

GetVolOk returns a tuple with the Vol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol

`func (o *AirConParams) SetVol(v AirVolume)`

SetVol sets Vol field to given value.

### HasVol

`func (o *AirConParams) HasVol() bool`

HasVol returns a boolean if a field has been set.

### GetDir

`func (o *AirConParams) GetDir() AirDirection`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *AirConParams) GetDirOk() (*AirDirection, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *AirConParams) SetDir(v AirDirection)`

SetDir sets Dir field to given value.

### HasDir

`func (o *AirConParams) HasDir() bool`

HasDir returns a boolean if a field has been set.

### GetButton

`func (o *AirConParams) GetButton() ACButton`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *AirConParams) GetButtonOk() (*ACButton, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *AirConParams) SetButton(v ACButton)`

SetButton sets Button field to given value.

### HasButton

`func (o *AirConParams) HasButton() bool`

HasButton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


