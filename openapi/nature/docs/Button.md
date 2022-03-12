# Button

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of button. It is used for \&quot;POST /1/{applaince}/tv\&quot; and \&quot;POST /1/{appliance}/light\&quot;. | [optional] 
**Image** | Pointer to **string** | Basename of the image file included in the app. Ex: \&quot;ico_ac_1\&quot;  | [optional] 
**Label** | Pointer to **string** | Label of button. | [optional] 

## Methods

### NewButton

`func NewButton() *Button`

NewButton instantiates a new Button object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewButtonWithDefaults

`func NewButtonWithDefaults() *Button`

NewButtonWithDefaults instantiates a new Button object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Button) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Button) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Button) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Button) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *Button) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Button) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Button) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Button) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLabel

`func (o *Button) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Button) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Button) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Button) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


