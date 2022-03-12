# Signal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** | Basename of the image file included in the app. Ex: \&quot;ico_ac_1\&quot;  | [optional] 

## Methods

### NewSignal

`func NewSignal() *Signal`

NewSignal instantiates a new Signal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalWithDefaults

`func NewSignalWithDefaults() *Signal`

NewSignalWithDefaults instantiates a new Signal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Signal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Signal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Signal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Signal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Signal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Signal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Signal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Signal) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *Signal) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Signal) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Signal) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Signal) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


