# SensorValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Val** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSensorValue

`func NewSensorValue() *SensorValue`

NewSensorValue instantiates a new SensorValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorValueWithDefaults

`func NewSensorValueWithDefaults() *SensorValue`

NewSensorValueWithDefaults instantiates a new SensorValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVal

`func (o *SensorValue) GetVal() float32`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *SensorValue) GetValOk() (*float32, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *SensorValue) SetVal(v float32)`

SetVal sets Val field to given value.

### HasVal

`func (o *SensorValue) HasVal() bool`

HasVal returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SensorValue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SensorValue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SensorValue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SensorValue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


