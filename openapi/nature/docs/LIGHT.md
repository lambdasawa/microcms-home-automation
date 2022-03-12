# LIGHT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**LIGHTState**](LIGHTState.md) |  | [optional] 
**Buttons** | Pointer to [**[]Button**](Button.md) |  | [optional] 

## Methods

### NewLIGHT

`func NewLIGHT() *LIGHT`

NewLIGHT instantiates a new LIGHT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLIGHTWithDefaults

`func NewLIGHTWithDefaults() *LIGHT`

NewLIGHTWithDefaults instantiates a new LIGHT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *LIGHT) GetState() LIGHTState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LIGHT) GetStateOk() (*LIGHTState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LIGHT) SetState(v LIGHTState)`

SetState sets State field to given value.

### HasState

`func (o *LIGHT) HasState() bool`

HasState returns a boolean if a field has been set.

### GetButtons

`func (o *LIGHT) GetButtons() []Button`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *LIGHT) GetButtonsOk() (*[]Button, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *LIGHT) SetButtons(v []Button)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *LIGHT) HasButtons() bool`

HasButtons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


