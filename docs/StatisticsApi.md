# \StatisticsApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatisticsGetMonitorGroupStatistics**](StatisticsApi.md#StatisticsGetMonitorGroupStatistics) | **Get** /Statistics/MonitorGroup/{monitorGroupGuid} | Gets the monitor group statistics.
[**StatisticsGetMonitorStatistics**](StatisticsApi.md#StatisticsGetMonitorStatistics) | **Get** /Statistics/Monitor/{monitorGuid} | Gets the monitor statistics.



## StatisticsGetMonitorGroupStatistics

> StatisticsResponse StatisticsGetMonitorGroupStatistics(ctx, monitorGroupGuid).Filter(filter).Start(start).End(end).PresetPeriod(presetPeriod).Execute()

Gets the monitor group statistics.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group.
    filter := "filter_example" // string | The filter for the requested response fields. E.g. \"Alerts,SlaTarget\". (optional)
    start := time.Now() // time.Time | The start of a custom period (can't be used together with the PresetPeriod parameter) (optional)
    end := time.Now() // time.Time | The end of a custom period (optional)
    presetPeriod := "presetPeriod_example" // string | The requested time period. (optional) (default to "Last24Hours")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatisticsApi.StatisticsGetMonitorGroupStatistics(context.Background(), monitorGroupGuid).Filter(filter).Start(start).End(end).PresetPeriod(presetPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.StatisticsGetMonitorGroupStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatisticsGetMonitorGroupStatistics`: StatisticsResponse
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.StatisticsGetMonitorGroupStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticsGetMonitorGroupStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | The filter for the requested response fields. E.g. \&quot;Alerts,SlaTarget\&quot;. | 
 **start** | **time.Time** | The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | **time.Time** | The end of a custom period | 
 **presetPeriod** | **string** | The requested time period. | [default to &quot;Last24Hours&quot;]

### Return type

[**StatisticsResponse**](StatisticsResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticsGetMonitorStatistics

> StatisticsResponse StatisticsGetMonitorStatistics(ctx, monitorGuid).Filter(filter).Start(start).End(end).PresetPeriod(presetPeriod).Execute()

Gets the monitor statistics.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | The Guid of the monitor.
    filter := "filter_example" // string | The filter for the requested response fields. E.g. \"Alerts,SlaTarget\". (optional)
    start := time.Now() // time.Time | The start of a custom period (can't be used together with the PresetPeriod parameter) (optional)
    end := time.Now() // time.Time | The end of a custom period (optional)
    presetPeriod := "presetPeriod_example" // string | The requested time period. (optional) (default to "Last24Hours")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatisticsApi.StatisticsGetMonitorStatistics(context.Background(), monitorGuid).Filter(filter).Start(start).End(end).PresetPeriod(presetPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.StatisticsGetMonitorStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatisticsGetMonitorStatistics`: StatisticsResponse
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.StatisticsGetMonitorStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The Guid of the monitor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticsGetMonitorStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | The filter for the requested response fields. E.g. \&quot;Alerts,SlaTarget\&quot;. | 
 **start** | **time.Time** | The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | **time.Time** | The end of a custom period | 
 **presetPeriod** | **string** | The requested time period. | [default to &quot;Last24Hours&quot;]

### Return type

[**StatisticsResponse**](StatisticsResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

