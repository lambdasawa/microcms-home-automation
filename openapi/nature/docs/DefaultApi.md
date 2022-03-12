# \DefaultApi

All URIs are relative to *https://api.nature.global*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Call1ApplianceOrdersPost**](DefaultApi.md#Call1ApplianceOrdersPost) | **Post** /1/appliance_orders | 
[**Call1AppliancesApplianceAirconSettingsPost**](DefaultApi.md#Call1AppliancesApplianceAirconSettingsPost) | **Post** /1/appliances/{appliance}/aircon_settings | 
[**Call1AppliancesApplianceDeletePost**](DefaultApi.md#Call1AppliancesApplianceDeletePost) | **Post** /1/appliances/{appliance}/delete | 
[**Call1AppliancesApplianceLightPost**](DefaultApi.md#Call1AppliancesApplianceLightPost) | **Post** /1/appliances/{appliance}/light | 
[**Call1AppliancesAppliancePost**](DefaultApi.md#Call1AppliancesAppliancePost) | **Post** /1/appliances/{appliance} | 
[**Call1AppliancesApplianceSignalOrdersPost**](DefaultApi.md#Call1AppliancesApplianceSignalOrdersPost) | **Post** /1/appliances/{appliance}/signal_orders | 
[**Call1AppliancesApplianceSignalsGet**](DefaultApi.md#Call1AppliancesApplianceSignalsGet) | **Get** /1/appliances/{appliance}/signals | 
[**Call1AppliancesApplianceSignalsPost**](DefaultApi.md#Call1AppliancesApplianceSignalsPost) | **Post** /1/appliances/{appliance}/signals | 
[**Call1AppliancesApplianceTvPost**](DefaultApi.md#Call1AppliancesApplianceTvPost) | **Post** /1/appliances/{appliance}/tv | 
[**Call1AppliancesGet**](DefaultApi.md#Call1AppliancesGet) | **Get** /1/appliances | 
[**Call1AppliancesPost**](DefaultApi.md#Call1AppliancesPost) | **Post** /1/appliances | 
[**Call1DetectappliancePost**](DefaultApi.md#Call1DetectappliancePost) | **Post** /1/detectappliance | 
[**Call1DevicesDeviceDeletePost**](DefaultApi.md#Call1DevicesDeviceDeletePost) | **Post** /1/devices/{device}/delete | 
[**Call1DevicesDeviceHumidityOffsetPost**](DefaultApi.md#Call1DevicesDeviceHumidityOffsetPost) | **Post** /1/devices/{device}/humidity_offset | 
[**Call1DevicesDevicePost**](DefaultApi.md#Call1DevicesDevicePost) | **Post** /1/devices/{device} | 
[**Call1DevicesDeviceTemperatureOffsetPost**](DefaultApi.md#Call1DevicesDeviceTemperatureOffsetPost) | **Post** /1/devices/{device}/temperature_offset | 
[**Call1DevicesGet**](DefaultApi.md#Call1DevicesGet) | **Get** /1/devices | 
[**Call1SignalsSignalDeletePost**](DefaultApi.md#Call1SignalsSignalDeletePost) | **Post** /1/signals/{signal}/delete | 
[**Call1SignalsSignalPost**](DefaultApi.md#Call1SignalsSignalPost) | **Post** /1/signals/{signal} | 
[**Call1SignalsSignalSendPost**](DefaultApi.md#Call1SignalsSignalSendPost) | **Post** /1/signals/{signal}/send | 
[**Call1UsersMeGet**](DefaultApi.md#Call1UsersMeGet) | **Get** /1/users/me | 
[**Call1UsersMePost**](DefaultApi.md#Call1UsersMePost) | **Post** /1/users/me | 



## Call1ApplianceOrdersPost

> Call1ApplianceOrdersPost(ctx).Appliances(appliances).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appliances := "appliances_example" // string | List of all appliances' IDs comma separated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1ApplianceOrdersPost(context.Background()).Appliances(appliances).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1ApplianceOrdersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCall1ApplianceOrdersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appliances** | **string** | List of all appliances&#39; IDs comma separated | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1AppliancesApplianceAirconSettingsPost

> Call1AppliancesApplianceAirconSettingsPost(ctx, appliance).Temperature(temperature).OperationMode(operationMode).AirVolume(airVolume).AirDirection(airDirection).Button(button).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appliance := "appliance_example" // string | Appliance ID.
    temperature := "temperature_example" // string | Temperature (optional)
    operationMode := "operationMode_example" // string | AC operation mode (optional)
    airVolume := "airVolume_example" // string | AC air volume (optional)
    airDirection := "airDirection_example" // string | AC air direction (optional)
    button := "button_example" // string | Button (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1AppliancesApplianceAirconSettingsPost(context.Background(), appliance).Temperature(temperature).OperationMode(operationMode).AirVolume(airVolume).AirDirection(airDirection).Button(button).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1AppliancesApplianceAirconSettingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appliance** | **string** | Appliance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1AppliancesApplianceAirconSettingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **temperature** | **string** | Temperature | 
 **operationMode** | **string** | AC operation mode | 
 **airVolume** | **string** | AC air volume | 
 **airDirection** | **string** | AC air direction | 
 **button** | **string** | Button | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1AppliancesApplianceDeletePost

> Call1AppliancesApplianceDeletePost(ctx, appliance).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appliance := "appliance_example" // string | Appliance ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1AppliancesApplianceDeletePost(context.Background(), appliance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1AppliancesApplianceDeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appliance** | **string** | Appliance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1AppliancesApplianceDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1AppliancesApplianceLightPost

> Call1AppliancesApplianceLightPost(ctx, appliance).Button(button).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appliance := "appliance_example" // string | Appliance ID.
    button := "button_example" // string | Button name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1AppliancesApplianceLightPost(context.Background(), appliance).Button(button).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1AppliancesApplianceLightPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appliance** | **string** | Appliance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1AppliancesApplianceLightPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **button** | **string** | Button name | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1AppliancesAppliancePost

> Appliance Call1AppliancesAppliancePost(ctx, appliance).Image(image).Nickname(nickname).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appliance := "appliance_example" // string | Appliance ID.
    image := "image_example" // string | Basename of the image file included in the app. Ex: \\\"ico_ac_1\\\" 
    nickname := "nickname_example" // string | Appliance name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1AppliancesAppliancePost(context.Background(), appliance).Image(image).Nickname(nickname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1AppliancesAppliancePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call1AppliancesAppliancePost`: Appliance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Call1AppliancesAppliancePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appliance** | **string** | Appliance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1AppliancesAppliancePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **image** | **string** | Basename of the image file included in the app. Ex: \\\&quot;ico_ac_1\\\&quot;  | 
 **nickname** | **string** | Appliance name | 

### Return type

[**Appliance**](Appliance.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1AppliancesApplianceSignalOrdersPost

> Call1AppliancesApplianceSignalOrdersPost(ctx, appliance).Signals(signals).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appliance := "appliance_example" // string | Appliance ID.
    signals := "signals_example" // string | List of all signals' IDs comma separated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1AppliancesApplianceSignalOrdersPost(context.Background(), appliance).Signals(signals).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1AppliancesApplianceSignalOrdersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appliance** | **string** | Appliance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1AppliancesApplianceSignalOrdersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signals** | **string** | List of all signals&#39; IDs comma separated | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1AppliancesApplianceSignalsGet

> []Signal Call1AppliancesApplianceSignalsGet(ctx, appliance).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appliance := "appliance_example" // string | Appliance ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1AppliancesApplianceSignalsGet(context.Background(), appliance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1AppliancesApplianceSignalsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call1AppliancesApplianceSignalsGet`: []Signal
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Call1AppliancesApplianceSignalsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appliance** | **string** | Appliance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1AppliancesApplianceSignalsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Signal**](Signal.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1AppliancesApplianceSignalsPost

> Signal Call1AppliancesApplianceSignalsPost(ctx, appliance).Message(message).Image(image).Name(name).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appliance := "appliance_example" // string | Appliance ID.
    message := "message_example" // string | JSON serialized object describing infrared signals. Includes \\\"data\\\", \\\"freq\\\" and \\\"format\\\" keys.
    image := "image_example" // string | Basename of the image file included in the app. Ex: \\\"ico_io\\\" 
    name := "name_example" // string | Signal name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1AppliancesApplianceSignalsPost(context.Background(), appliance).Message(message).Image(image).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1AppliancesApplianceSignalsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call1AppliancesApplianceSignalsPost`: Signal
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Call1AppliancesApplianceSignalsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appliance** | **string** | Appliance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1AppliancesApplianceSignalsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | **string** | JSON serialized object describing infrared signals. Includes \\\&quot;data\\\&quot;, \\\&quot;freq\\\&quot; and \\\&quot;format\\\&quot; keys. | 
 **image** | **string** | Basename of the image file included in the app. Ex: \\\&quot;ico_io\\\&quot;  | 
 **name** | **string** | Signal name | 

### Return type

[**Signal**](Signal.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1AppliancesApplianceTvPost

> Call1AppliancesApplianceTvPost(ctx, appliance).Button(button).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appliance := "appliance_example" // string | Appliance ID.
    button := "button_example" // string | Button name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1AppliancesApplianceTvPost(context.Background(), appliance).Button(button).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1AppliancesApplianceTvPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appliance** | **string** | Appliance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1AppliancesApplianceTvPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **button** | **string** | Button name | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1AppliancesGet

> []Appliance Call1AppliancesGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1AppliancesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1AppliancesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call1AppliancesGet`: []Appliance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Call1AppliancesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCall1AppliancesGetRequest struct via the builder pattern


### Return type

[**[]Appliance**](Appliance.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1AppliancesPost

> Appliance Call1AppliancesPost(ctx).Nickname(nickname).Device(device).Image(image).Model(model).ModelType(modelType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nickname := "nickname_example" // string | Appliance name
    device := "device_example" // string | Device ID
    image := "image_example" // string | Basename of the image file included in the app. Ex: \\\"ico_ac_1\\\" 
    model := "model_example" // string | ApplianceModel ID if the appliance we're trying to create is included in IRDB. (optional)
    modelType := "modelType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1AppliancesPost(context.Background()).Nickname(nickname).Device(device).Image(image).Model(model).ModelType(modelType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1AppliancesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call1AppliancesPost`: Appliance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Call1AppliancesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCall1AppliancesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nickname** | **string** | Appliance name | 
 **device** | **string** | Device ID | 
 **image** | **string** | Basename of the image file included in the app. Ex: \\\&quot;ico_ac_1\\\&quot;  | 
 **model** | **string** | ApplianceModel ID if the appliance we&#39;re trying to create is included in IRDB. | 
 **modelType** | **string** |  | 

### Return type

[**Appliance**](Appliance.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1DetectappliancePost

> []ApplianceModelAndParam Call1DetectappliancePost(ctx).Message(message).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    message := "message_example" // string | JSON serialized object describing infrared signals. Includes \\\"data\\\", \\\"freq\\\" and \\\"format\\\" keys.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1DetectappliancePost(context.Background()).Message(message).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1DetectappliancePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call1DetectappliancePost`: []ApplianceModelAndParam
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Call1DetectappliancePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCall1DetectappliancePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **string** | JSON serialized object describing infrared signals. Includes \\\&quot;data\\\&quot;, \\\&quot;freq\\\&quot; and \\\&quot;format\\\&quot; keys. | 

### Return type

[**[]ApplianceModelAndParam**](ApplianceModelAndParam.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1DevicesDeviceDeletePost

> Call1DevicesDeviceDeletePost(ctx, device).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    device := "device_example" // string | Device ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1DevicesDeviceDeletePost(context.Background(), device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1DevicesDeviceDeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**device** | **string** | Device ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1DevicesDeviceDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1DevicesDeviceHumidityOffsetPost

> Call1DevicesDeviceHumidityOffsetPost(ctx, device).Offset(offset).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    device := "device_example" // string | Device ID.
    offset := int32(56) // int32 | Humidity offset value added to the measured humidity.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1DevicesDeviceHumidityOffsetPost(context.Background(), device).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1DevicesDeviceHumidityOffsetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**device** | **string** | Device ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1DevicesDeviceHumidityOffsetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Humidity offset value added to the measured humidity. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1DevicesDevicePost

> Call1DevicesDevicePost(ctx, device).Name(name).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    device := "device_example" // string | Device ID.
    name := "name_example" // string | Signal name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1DevicesDevicePost(context.Background(), device).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1DevicesDevicePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**device** | **string** | Device ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1DevicesDevicePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Signal name | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1DevicesDeviceTemperatureOffsetPost

> Call1DevicesDeviceTemperatureOffsetPost(ctx, device).Offset(offset).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    device := "device_example" // string | Device ID.
    offset := int32(56) // int32 | Temperature offset value added to the measured temperature.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1DevicesDeviceTemperatureOffsetPost(context.Background(), device).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1DevicesDeviceTemperatureOffsetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**device** | **string** | Device ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1DevicesDeviceTemperatureOffsetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Temperature offset value added to the measured temperature. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1DevicesGet

> []Device Call1DevicesGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1DevicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1DevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call1DevicesGet`: []Device
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Call1DevicesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCall1DevicesGetRequest struct via the builder pattern


### Return type

[**[]Device**](Device.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1SignalsSignalDeletePost

> Call1SignalsSignalDeletePost(ctx, signal).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    signal := "signal_example" // string | Signal ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1SignalsSignalDeletePost(context.Background(), signal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1SignalsSignalDeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signal** | **string** | Signal ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1SignalsSignalDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1SignalsSignalPost

> Call1SignalsSignalPost(ctx, signal).Image(image).Name(name).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    signal := "signal_example" // string | Signal ID.
    image := "image_example" // string | Basename of the image file included in the app. Ex: \\\"ico_io\\\" 
    name := "name_example" // string | Signal name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1SignalsSignalPost(context.Background(), signal).Image(image).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1SignalsSignalPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signal** | **string** | Signal ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1SignalsSignalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **image** | **string** | Basename of the image file included in the app. Ex: \\\&quot;ico_io\\\&quot;  | 
 **name** | **string** | Signal name | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1SignalsSignalSendPost

> Call1SignalsSignalSendPost(ctx, signal).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    signal := "signal_example" // string | Signal ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1SignalsSignalSendPost(context.Background(), signal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1SignalsSignalSendPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signal** | **string** | Signal ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall1SignalsSignalSendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1UsersMeGet

> User Call1UsersMeGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1UsersMeGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1UsersMeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call1UsersMeGet`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Call1UsersMeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCall1UsersMeGetRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Call1UsersMePost

> User Call1UsersMePost(ctx).Nickname(nickname).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nickname := "nickname_example" // string | User's nickname. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Call1UsersMePost(context.Background()).Nickname(nickname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Call1UsersMePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call1UsersMePost`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Call1UsersMePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCall1UsersMePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nickname** | **string** | User&#39;s nickname.  | 

### Return type

[**User**](User.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

