# Go API client for openapi

Read/Write Nature Remo

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.nature.global*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**Call1ApplianceOrdersPost**](docs/DefaultApi.md#call1applianceorderspost) | **Post** /1/appliance_orders | 
*DefaultApi* | [**Call1AppliancesApplianceAirconSettingsPost**](docs/DefaultApi.md#call1appliancesapplianceairconsettingspost) | **Post** /1/appliances/{appliance}/aircon_settings | 
*DefaultApi* | [**Call1AppliancesApplianceDeletePost**](docs/DefaultApi.md#call1appliancesappliancedeletepost) | **Post** /1/appliances/{appliance}/delete | 
*DefaultApi* | [**Call1AppliancesApplianceLightPost**](docs/DefaultApi.md#call1appliancesappliancelightpost) | **Post** /1/appliances/{appliance}/light | 
*DefaultApi* | [**Call1AppliancesAppliancePost**](docs/DefaultApi.md#call1appliancesappliancepost) | **Post** /1/appliances/{appliance} | 
*DefaultApi* | [**Call1AppliancesApplianceSignalOrdersPost**](docs/DefaultApi.md#call1appliancesappliancesignalorderspost) | **Post** /1/appliances/{appliance}/signal_orders | 
*DefaultApi* | [**Call1AppliancesApplianceSignalsGet**](docs/DefaultApi.md#call1appliancesappliancesignalsget) | **Get** /1/appliances/{appliance}/signals | 
*DefaultApi* | [**Call1AppliancesApplianceSignalsPost**](docs/DefaultApi.md#call1appliancesappliancesignalspost) | **Post** /1/appliances/{appliance}/signals | 
*DefaultApi* | [**Call1AppliancesApplianceTvPost**](docs/DefaultApi.md#call1appliancesappliancetvpost) | **Post** /1/appliances/{appliance}/tv | 
*DefaultApi* | [**Call1AppliancesGet**](docs/DefaultApi.md#call1appliancesget) | **Get** /1/appliances | 
*DefaultApi* | [**Call1AppliancesPost**](docs/DefaultApi.md#call1appliancespost) | **Post** /1/appliances | 
*DefaultApi* | [**Call1DetectappliancePost**](docs/DefaultApi.md#call1detectappliancepost) | **Post** /1/detectappliance | 
*DefaultApi* | [**Call1DevicesDeviceDeletePost**](docs/DefaultApi.md#call1devicesdevicedeletepost) | **Post** /1/devices/{device}/delete | 
*DefaultApi* | [**Call1DevicesDeviceHumidityOffsetPost**](docs/DefaultApi.md#call1devicesdevicehumidityoffsetpost) | **Post** /1/devices/{device}/humidity_offset | 
*DefaultApi* | [**Call1DevicesDevicePost**](docs/DefaultApi.md#call1devicesdevicepost) | **Post** /1/devices/{device} | 
*DefaultApi* | [**Call1DevicesDeviceTemperatureOffsetPost**](docs/DefaultApi.md#call1devicesdevicetemperatureoffsetpost) | **Post** /1/devices/{device}/temperature_offset | 
*DefaultApi* | [**Call1DevicesGet**](docs/DefaultApi.md#call1devicesget) | **Get** /1/devices | 
*DefaultApi* | [**Call1SignalsSignalDeletePost**](docs/DefaultApi.md#call1signalssignaldeletepost) | **Post** /1/signals/{signal}/delete | 
*DefaultApi* | [**Call1SignalsSignalPost**](docs/DefaultApi.md#call1signalssignalpost) | **Post** /1/signals/{signal} | 
*DefaultApi* | [**Call1SignalsSignalSendPost**](docs/DefaultApi.md#call1signalssignalsendpost) | **Post** /1/signals/{signal}/send | 
*DefaultApi* | [**Call1UsersMeGet**](docs/DefaultApi.md#call1usersmeget) | **Get** /1/users/me | 
*DefaultApi* | [**Call1UsersMePost**](docs/DefaultApi.md#call1usersmepost) | **Post** /1/users/me | 


## Documentation For Models

 - [ACButton](docs/ACButton.md)
 - [AirCon](docs/AirCon.md)
 - [AirConParams](docs/AirConParams.md)
 - [AirConRange](docs/AirConRange.md)
 - [AirConRangeMode](docs/AirConRangeMode.md)
 - [AirConRangeModes](docs/AirConRangeModes.md)
 - [AirDirection](docs/AirDirection.md)
 - [AirVolume](docs/AirVolume.md)
 - [Appliance](docs/Appliance.md)
 - [ApplianceModel](docs/ApplianceModel.md)
 - [ApplianceModelAndParam](docs/ApplianceModelAndParam.md)
 - [ApplianceType](docs/ApplianceType.md)
 - [Button](docs/Button.md)
 - [Device](docs/Device.md)
 - [DeviceCore](docs/DeviceCore.md)
 - [DeviceNewestEvents](docs/DeviceNewestEvents.md)
 - [EchonetLiteProperty](docs/EchonetLiteProperty.md)
 - [LIGHT](docs/LIGHT.md)
 - [LIGHTState](docs/LIGHTState.md)
 - [OperationMode](docs/OperationMode.md)
 - [SensorValue](docs/SensorValue.md)
 - [Signal](docs/Signal.md)
 - [SmartMeter](docs/SmartMeter.md)
 - [TV](docs/TV.md)
 - [TVState](docs/TVState.md)
 - [User](docs/User.md)


## Documentation For Authorization



### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://api.nature.global/oauth2/auth
- **Scopes**: 
 - **basic.read**: Read only access to user's profile (excluding email), Remos, appliances, signals.
 - **basic**: Read + write access to user's profile (excluding email), Remos, appliances, signals.
 - **sendir**: Send infrared signals through Remo.
 - **detectappliance**: Detect air conditioner model from infrared signals.

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



