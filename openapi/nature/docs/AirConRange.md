# AirConRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modes** | Pointer to [**AirConRangeModes**](AirConRangeModes.md) |  | [optional] 
**FixedButtons** | Pointer to [**[]ACButton**](ACButton.md) |  | [optional] 

## Methods

### NewAirConRange

`func NewAirConRange() *AirConRange`

NewAirConRange instantiates a new AirConRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirConRangeWithDefaults

`func NewAirConRangeWithDefaults() *AirConRange`

NewAirConRangeWithDefaults instantiates a new AirConRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModes

`func (o *AirConRange) GetModes() AirConRangeModes`

GetModes returns the Modes field if non-nil, zero value otherwise.

### GetModesOk

`func (o *AirConRange) GetModesOk() (*AirConRangeModes, bool)`

GetModesOk returns a tuple with the Modes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModes

`func (o *AirConRange) SetModes(v AirConRangeModes)`

SetModes sets Modes field to given value.

### HasModes

`func (o *AirConRange) HasModes() bool`

HasModes returns a boolean if a field has been set.

### GetFixedButtons

`func (o *AirConRange) GetFixedButtons() []ACButton`

GetFixedButtons returns the FixedButtons field if non-nil, zero value otherwise.

### GetFixedButtonsOk

`func (o *AirConRange) GetFixedButtonsOk() (*[]ACButton, bool)`

GetFixedButtonsOk returns a tuple with the FixedButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedButtons

`func (o *AirConRange) SetFixedButtons(v []ACButton)`

SetFixedButtons sets FixedButtons field to given value.

### HasFixedButtons

`func (o *AirConRange) HasFixedButtons() bool`

HasFixedButtons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


