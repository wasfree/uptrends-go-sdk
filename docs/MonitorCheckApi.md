# \MonitorCheckApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitorCheckGetAccountMonitorChecks**](MonitorCheckApi.md#MonitorCheckGetAccountMonitorChecks) | **Get** /MonitorCheck | Returns all monitor check data.
[**MonitorCheckGetConcurrentMonitorPartialChecks**](MonitorCheckApi.md#MonitorCheckGetConcurrentMonitorPartialChecks) | **Get** /MonitorCheck/{monitorCheckId}/Concurrent | Gets all partial checks for a concurrent monitor check
[**MonitorCheckGetConsoleLogInfo**](MonitorCheckApi.md#MonitorCheckGetConsoleLogInfo) | **Get** /MonitorCheck/{monitorCheckId}/ConsoleLog | Returns console log information for a monitor check.
[**MonitorCheckGetHttpDetails**](MonitorCheckApi.md#MonitorCheckGetHttpDetails) | **Get** /MonitorCheck/{monitorCheckId}/Http | Returns HTTP details for a monitor check.
[**MonitorCheckGetMonitorCheck**](MonitorCheckApi.md#MonitorCheckGetMonitorCheck) | **Get** /MonitorCheck/Monitor/{monitorGuid} | Returns monitor check data for a specific monitor.
[**MonitorCheckGetMonitorGroupData**](MonitorCheckApi.md#MonitorCheckGetMonitorGroupData) | **Get** /MonitorCheck/MonitorGroup/{monitorGroupGuid} | Returns monitor check data for a specific monitor group.
[**MonitorCheckGetMultistepDetails**](MonitorCheckApi.md#MonitorCheckGetMultistepDetails) | **Get** /MonitorCheck/{monitorCheckId}/MultiStepAPI | Returns Multi-Step API details for a monitor check.
[**MonitorCheckGetPageSourceInfo**](MonitorCheckApi.md#MonitorCheckGetPageSourceInfo) | **Get** /MonitorCheck/{monitorCheckId}/PageSource | Returns page source information for a monitor check.
[**MonitorCheckGetScreenshots**](MonitorCheckApi.md#MonitorCheckGetScreenshots) | **Get** /MonitorCheck/{monitorCheckId}/Screenshot/{screenshotId} | Gets a specific screenshot for a specified monitor check
[**MonitorCheckGetSingleMonitorCheck**](MonitorCheckApi.md#MonitorCheckGetSingleMonitorCheck) | **Get** /MonitorCheck/{monitorCheckId} | Returns a single monitor check.
[**MonitorCheckGetTransactionDetails**](MonitorCheckApi.md#MonitorCheckGetTransactionDetails) | **Get** /MonitorCheck/{monitorCheckId}/Transaction | Returns transaction step details for a monitor check.
[**MonitorCheckGetWaterfallInfo**](MonitorCheckApi.md#MonitorCheckGetWaterfallInfo) | **Get** /MonitorCheck/{monitorCheckId}/Waterfall | Returns waterfall information for a monitor check.



## MonitorCheckGetAccountMonitorChecks

> MonitorCheckResponse MonitorCheckGetAccountMonitorChecks(ctx).ErrorLevel(errorLevel).ShowPartialMeasurements(showPartialMeasurements).Cursor(cursor).Sorting(sorting).Take(take).Start(start).End(end).PresetPeriod(presetPeriod).Execute()

Returns all monitor check data.

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
    errorLevel := "errorLevel_example" // string | Error level filter that should be applied. (default = NoError and above) (optional)
    showPartialMeasurements := true // bool | Show partial measurements from concurrent monitors (optional)
    cursor := "cursor_example" // string | A cursor value that should be used for traversing the dataset. (optional)
    sorting := "sorting_example" // string | Sorting direction based on timestamp. (optional) (default to "Descending")
    take := int32(56) // int32 | The number of records to return (Max value = 100) (optional) (default to 100)
    start := TODO // interface{} | The start of a custom period (can't be used together with the PresetPeriod parameter) (optional)
    end := TODO // interface{} | The end of a custom period (optional)
    presetPeriod := "presetPeriod_example" // string | The requested time period. (optional) (default to "Last24Hours")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetAccountMonitorChecks(context.Background()).ErrorLevel(errorLevel).ShowPartialMeasurements(showPartialMeasurements).Cursor(cursor).Sorting(sorting).Take(take).Start(start).End(end).PresetPeriod(presetPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetAccountMonitorChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetAccountMonitorChecks`: MonitorCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetAccountMonitorChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetAccountMonitorChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **errorLevel** | **string** | Error level filter that should be applied. (default &#x3D; NoError and above) | 
 **showPartialMeasurements** | **bool** | Show partial measurements from concurrent monitors | 
 **cursor** | **string** | A cursor value that should be used for traversing the dataset. | 
 **sorting** | **string** | Sorting direction based on timestamp. | [default to &quot;Descending&quot;]
 **take** | **int32** | The number of records to return (Max value &#x3D; 100) | [default to 100]
 **start** | [**interface{}**](interface{}.md) | The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | [**interface{}**](interface{}.md) | The end of a custom period | 
 **presetPeriod** | **string** | The requested time period. | [default to &quot;Last24Hours&quot;]

### Return type

[**MonitorCheckResponse**](MonitorCheckResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCheckGetConcurrentMonitorPartialChecks

> MonitorCheckResponse MonitorCheckGetConcurrentMonitorPartialChecks(ctx, monitorCheckId).Execute()

Gets all partial checks for a concurrent monitor check

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
    monitorCheckId := int64(789) // int64 | The monitor check Id to get the partial checks for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetConcurrentMonitorPartialChecks(context.Background(), monitorCheckId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetConcurrentMonitorPartialChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetConcurrentMonitorPartialChecks`: MonitorCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetConcurrentMonitorPartialChecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorCheckId** | **int64** | The monitor check Id to get the partial checks for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetConcurrentMonitorPartialChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorCheckResponse**](MonitorCheckResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCheckGetConsoleLogInfo

> WaterfallResponse MonitorCheckGetConsoleLogInfo(ctx, monitorCheckId).Step(step).Execute()

Returns console log information for a monitor check.

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
    monitorCheckId := int64(789) // int64 | The monitor check Id to get the detailed data for.
    step := int32(56) // int32 | For transactions only: the transaction step to get the console log for. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetConsoleLogInfo(context.Background(), monitorCheckId).Step(step).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetConsoleLogInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetConsoleLogInfo`: WaterfallResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetConsoleLogInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorCheckId** | **int64** | The monitor check Id to get the detailed data for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetConsoleLogInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **step** | **int32** | For transactions only: the transaction step to get the console log for. | 

### Return type

[**WaterfallResponse**](WaterfallResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCheckGetHttpDetails

> HttpDetailsResponse MonitorCheckGetHttpDetails(ctx, monitorCheckId).Execute()

Returns HTTP details for a monitor check.

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
    monitorCheckId := int64(789) // int64 | The monitor check Id to get the detailed data for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetHttpDetails(context.Background(), monitorCheckId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetHttpDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetHttpDetails`: HttpDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetHttpDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorCheckId** | **int64** | The monitor check Id to get the detailed data for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetHttpDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HttpDetailsResponse**](HttpDetailsResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCheckGetMonitorCheck

> MonitorCheckResponse MonitorCheckGetMonitorCheck(ctx, monitorGuid).ErrorLevel(errorLevel).ShowPartialMeasurements(showPartialMeasurements).Cursor(cursor).Sorting(sorting).Take(take).Start(start).End(end).PresetPeriod(presetPeriod).Execute()

Returns monitor check data for a specific monitor.

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
    monitorGuid := "monitorGuid_example" // string | The Guid of the monitor to get monitor checks for.
    errorLevel := "errorLevel_example" // string | Error level filter that should be applied. (default = NoError and above) (optional)
    showPartialMeasurements := true // bool | Show partial measurements from concurrent monitors (optional)
    cursor := "cursor_example" // string | A cursor value that should be used for traversing the dataset. (optional)
    sorting := "sorting_example" // string | Sorting direction based on timestamp. (optional) (default to "Descending")
    take := int32(56) // int32 | The number of records to return (Max value = 100) (optional) (default to 100)
    start := TODO // interface{} | The start of a custom period (can't be used together with the PresetPeriod parameter) (optional)
    end := TODO // interface{} | The end of a custom period (optional)
    presetPeriod := "presetPeriod_example" // string | The requested time period. (optional) (default to "Last24Hours")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetMonitorCheck(context.Background(), monitorGuid).ErrorLevel(errorLevel).ShowPartialMeasurements(showPartialMeasurements).Cursor(cursor).Sorting(sorting).Take(take).Start(start).End(end).PresetPeriod(presetPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetMonitorCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetMonitorCheck`: MonitorCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetMonitorCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The Guid of the monitor to get monitor checks for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetMonitorCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **errorLevel** | **string** | Error level filter that should be applied. (default &#x3D; NoError and above) | 
 **showPartialMeasurements** | **bool** | Show partial measurements from concurrent monitors | 
 **cursor** | **string** | A cursor value that should be used for traversing the dataset. | 
 **sorting** | **string** | Sorting direction based on timestamp. | [default to &quot;Descending&quot;]
 **take** | **int32** | The number of records to return (Max value &#x3D; 100) | [default to 100]
 **start** | [**interface{}**](interface{}.md) | The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | [**interface{}**](interface{}.md) | The end of a custom period | 
 **presetPeriod** | **string** | The requested time period. | [default to &quot;Last24Hours&quot;]

### Return type

[**MonitorCheckResponse**](MonitorCheckResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCheckGetMonitorGroupData

> MonitorCheckResponse MonitorCheckGetMonitorGroupData(ctx, monitorGroupGuid).ErrorLevel(errorLevel).ShowPartialMeasurements(showPartialMeasurements).Cursor(cursor).Sorting(sorting).Take(take).Start(start).End(end).PresetPeriod(presetPeriod).Execute()

Returns monitor check data for a specific monitor group.

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to get monitor checks for.
    errorLevel := "errorLevel_example" // string | Error level filter that should be applied. (default = NoError and above) (optional)
    showPartialMeasurements := true // bool | Show partial measurements from concurrent monitors (optional)
    cursor := "cursor_example" // string | A cursor value that should be used for traversing the dataset. (optional)
    sorting := "sorting_example" // string | Sorting direction based on timestamp. (optional) (default to "Descending")
    take := int32(56) // int32 | The number of records to return (Max value = 100) (optional) (default to 100)
    start := TODO // interface{} | The start of a custom period (can't be used together with the PresetPeriod parameter) (optional)
    end := TODO // interface{} | The end of a custom period (optional)
    presetPeriod := "presetPeriod_example" // string | The requested time period. (optional) (default to "Last24Hours")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetMonitorGroupData(context.Background(), monitorGroupGuid).ErrorLevel(errorLevel).ShowPartialMeasurements(showPartialMeasurements).Cursor(cursor).Sorting(sorting).Take(take).Start(start).End(end).PresetPeriod(presetPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetMonitorGroupData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetMonitorGroupData`: MonitorCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetMonitorGroupData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to get monitor checks for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetMonitorGroupDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **errorLevel** | **string** | Error level filter that should be applied. (default &#x3D; NoError and above) | 
 **showPartialMeasurements** | **bool** | Show partial measurements from concurrent monitors | 
 **cursor** | **string** | A cursor value that should be used for traversing the dataset. | 
 **sorting** | **string** | Sorting direction based on timestamp. | [default to &quot;Descending&quot;]
 **take** | **int32** | The number of records to return (Max value &#x3D; 100) | [default to 100]
 **start** | [**interface{}**](interface{}.md) | The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | [**interface{}**](interface{}.md) | The end of a custom period | 
 **presetPeriod** | **string** | The requested time period. | [default to &quot;Last24Hours&quot;]

### Return type

[**MonitorCheckResponse**](MonitorCheckResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCheckGetMultistepDetails

> MsaDetailsResponse MonitorCheckGetMultistepDetails(ctx, monitorCheckId).Execute()

Returns Multi-Step API details for a monitor check.

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
    monitorCheckId := int64(789) // int64 | The monitor check Id to get the detailed data for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetMultistepDetails(context.Background(), monitorCheckId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetMultistepDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetMultistepDetails`: MsaDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetMultistepDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorCheckId** | **int64** | The monitor check Id to get the detailed data for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetMultistepDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MsaDetailsResponse**](MsaDetailsResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCheckGetPageSourceInfo

> WaterfallResponse MonitorCheckGetPageSourceInfo(ctx, monitorCheckId).Step(step).Execute()

Returns page source information for a monitor check.

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
    monitorCheckId := int64(789) // int64 | The monitor check Id to get the detailed data for.
    step := int32(56) // int32 | For transactions only: the transaction step to get the page source for. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetPageSourceInfo(context.Background(), monitorCheckId).Step(step).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetPageSourceInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetPageSourceInfo`: WaterfallResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetPageSourceInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorCheckId** | **int64** | The monitor check Id to get the detailed data for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetPageSourceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **step** | **int32** | For transactions only: the transaction step to get the page source for. | 

### Return type

[**WaterfallResponse**](WaterfallResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCheckGetScreenshots

> ScreenshotResponse MonitorCheckGetScreenshots(ctx, monitorCheckId, screenshotId).Execute()

Gets a specific screenshot for a specified monitor check

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
    monitorCheckId := int64(789) // int64 | The monitor check Id to get the screenshot data for.
    screenshotId := "screenshotId_example" // string | The screenshot Id of the screenshot to get.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetScreenshots(context.Background(), monitorCheckId, screenshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetScreenshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetScreenshots`: ScreenshotResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetScreenshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorCheckId** | **int64** | The monitor check Id to get the screenshot data for. | 
**screenshotId** | **string** | The screenshot Id of the screenshot to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetScreenshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScreenshotResponse**](ScreenshotResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCheckGetSingleMonitorCheck

> SingleMonitorCheckResponse MonitorCheckGetSingleMonitorCheck(ctx, monitorCheckId).Execute()

Returns a single monitor check.

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
    monitorCheckId := int64(789) // int64 | The Id of the monitor check to get the data for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetSingleMonitorCheck(context.Background(), monitorCheckId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetSingleMonitorCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetSingleMonitorCheck`: SingleMonitorCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetSingleMonitorCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorCheckId** | **int64** | The Id of the monitor check to get the data for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetSingleMonitorCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SingleMonitorCheckResponse**](SingleMonitorCheckResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCheckGetTransactionDetails

> TransactionDetailsResponse MonitorCheckGetTransactionDetails(ctx, monitorCheckId).Execute()

Returns transaction step details for a monitor check.

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
    monitorCheckId := int64(789) // int64 | The monitor check Id to get the detailed data for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetTransactionDetails(context.Background(), monitorCheckId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetTransactionDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetTransactionDetails`: TransactionDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetTransactionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorCheckId** | **int64** | The monitor check Id to get the detailed data for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetTransactionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionDetailsResponse**](TransactionDetailsResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCheckGetWaterfallInfo

> WaterfallResponse MonitorCheckGetWaterfallInfo(ctx, monitorCheckId).Step(step).Execute()

Returns waterfall information for a monitor check.

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
    monitorCheckId := int64(789) // int64 | The monitor check Id to get the detailed data for.
    step := int32(56) // int32 | For transaction waterfalls only: the transaction step to get the waterfall for. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckApi.MonitorCheckGetWaterfallInfo(context.Background(), monitorCheckId).Step(step).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckApi.MonitorCheckGetWaterfallInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCheckGetWaterfallInfo`: WaterfallResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckApi.MonitorCheckGetWaterfallInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorCheckId** | **int64** | The monitor check Id to get the detailed data for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCheckGetWaterfallInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **step** | **int32** | For transaction waterfalls only: the transaction step to get the waterfall for. | 

### Return type

[**WaterfallResponse**](WaterfallResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

