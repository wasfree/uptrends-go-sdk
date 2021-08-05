# \StatusApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusGetMonitorGroupStatus**](StatusApi.md#StatusGetMonitorGroupStatus) | **Get** /Status/MonitorGroup/{monitorGroupGuid} | Gets a list of all monitor group status data.
[**StatusGetMonitorStatus**](StatusApi.md#StatusGetMonitorStatus) | **Get** /Status/Monitor/{monitorGuid} | Gets all monitor status data.



## StatusGetMonitorGroupStatus

> MonitorStatusListResponse StatusGetMonitorGroupStatus(ctx, monitorGroupGuid).Execute()

Gets a list of all monitor group status data.

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatusApi.StatusGetMonitorGroupStatus(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.StatusGetMonitorGroupStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusGetMonitorGroupStatus`: MonitorStatusListResponse
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.StatusGetMonitorGroupStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusGetMonitorGroupStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorStatusListResponse**](MonitorStatusListResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusGetMonitorStatus

> MonitorStatusResponse StatusGetMonitorStatus(ctx, monitorGuid).Execute()

Gets all monitor status data.

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
    monitorGuid := "monitorGuid_example" // string | The Guid of the monitor.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatusApi.StatusGetMonitorStatus(context.Background(), monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.StatusGetMonitorStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusGetMonitorStatus`: MonitorStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.StatusGetMonitorStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The Guid of the monitor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusGetMonitorStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorStatusResponse**](MonitorStatusResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

