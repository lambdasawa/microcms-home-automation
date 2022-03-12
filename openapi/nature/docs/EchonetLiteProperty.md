# EchonetLiteProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Epc** | Pointer to **int32** | ECHONET Property | [optional] 
**Val** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEchonetLiteProperty

`func NewEchonetLiteProperty() *EchonetLiteProperty`

NewEchonetLiteProperty instantiates a new EchonetLiteProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEchonetLitePropertyWithDefaults

`func NewEchonetLitePropertyWithDefaults() *EchonetLiteProperty`

NewEchonetLitePropertyWithDefaults instantiates a new EchonetLiteProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EchonetLiteProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EchonetLiteProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EchonetLiteProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EchonetLiteProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEpc

`func (o *EchonetLiteProperty) GetEpc() int32`

GetEpc returns the Epc field if non-nil, zero value otherwise.

### GetEpcOk

`func (o *EchonetLiteProperty) GetEpcOk() (*int32, bool)`

GetEpcOk returns a tuple with the Epc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpc

`func (o *EchonetLiteProperty) SetEpc(v int32)`

SetEpc sets Epc field to given value.

### HasEpc

`func (o *EchonetLiteProperty) HasEpc() bool`

HasEpc returns a boolean if a field has been set.

### GetVal

`func (o *EchonetLiteProperty) GetVal() string`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *EchonetLiteProperty) GetValOk() (*string, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *EchonetLiteProperty) SetVal(v string)`

SetVal sets Val field to given value.

### HasVal

`func (o *EchonetLiteProperty) HasVal() bool`

HasVal returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EchonetLiteProperty) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EchonetLiteProperty) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EchonetLiteProperty) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EchonetLiteProperty) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


