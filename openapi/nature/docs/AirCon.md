# AirCon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to [**AirConRange**](AirConRange.md) |  | [optional] 
**TempUnit** | Pointer to **string** |  | [optional] 

## Methods

### NewAirCon

`func NewAirCon() *AirCon`

NewAirCon instantiates a new AirCon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirConWithDefaults

`func NewAirConWithDefaults() *AirCon`

NewAirConWithDefaults instantiates a new AirCon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *AirCon) GetRange() AirConRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *AirCon) GetRangeOk() (*AirConRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *AirCon) SetRange(v AirConRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *AirCon) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetTempUnit

`func (o *AirCon) GetTempUnit() string`

GetTempUnit returns the TempUnit field if non-nil, zero value otherwise.

### GetTempUnitOk

`func (o *AirCon) GetTempUnitOk() (*string, bool)`

GetTempUnitOk returns a tuple with the TempUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempUnit

`func (o *AirCon) SetTempUnit(v string)`

SetTempUnit sets TempUnit field to given value.

### HasTempUnit

`func (o *AirCon) HasTempUnit() bool`

HasTempUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


