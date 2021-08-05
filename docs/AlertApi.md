# \AlertApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertGetMonitorAlerts**](AlertApi.md#AlertGetMonitorAlerts) | **Get** /Alert/Monitor/{monitorGuid} | Returns alerts for a specific monitor.
[**AlertGetMonitorGroupAlerts**](AlertApi.md#AlertGetMonitorGroupAlerts) | **Get** /Alert/MonitorGroup/{monitorGroupGuid} | Returns alerts for a specific monitor group.



## AlertGetMonitorAlerts

> AlertResponse AlertGetMonitorAlerts(ctx, monitorGuid).IncludeReminders(includeReminders).Cursor(cursor).Sorting(sorting).Take(take).Start(start).End(end).PresetPeriod(presetPeriod).Execute()

Returns alerts for a specific monitor.

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
    monitorGuid := "monitorGuid_example" // string | The Guid of the monitor to get alerts for.
    includeReminders := true // bool | When true reminder alerts will be included in the results (optional) (default to false)
    cursor := "cursor_example" // string | A cursor value that should be used for traversing the dataset. (optional)
    sorting := "sorting_example" // string | Sorting direction based on timestamp. (optional) (default to "Descending")
    take := int32(56) // int32 | The number of records to return (Max value = 100) (optional) (default to 100)
    start := time.Now() // time.Time | The start of a custom period (can't be used together with the PresetPeriod parameter) (optional)
    end := time.Now() // time.Time | The end of a custom period (optional)
    presetPeriod := "presetPeriod_example" // string | The requested time period. (optional) (default to "Last24Hours")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertApi.AlertGetMonitorAlerts(context.Background(), monitorGuid).IncludeReminders(includeReminders).Cursor(cursor).Sorting(sorting).Take(take).Start(start).End(end).PresetPeriod(presetPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertApi.AlertGetMonitorAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertGetMonitorAlerts`: AlertResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertApi.AlertGetMonitorAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The Guid of the monitor to get alerts for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertGetMonitorAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeReminders** | **bool** | When true reminder alerts will be included in the results | [default to false]
 **cursor** | **string** | A cursor value that should be used for traversing the dataset. | 
 **sorting** | **string** | Sorting direction based on timestamp. | [default to &quot;Descending&quot;]
 **take** | **int32** | The number of records to return (Max value &#x3D; 100) | [default to 100]
 **start** | **time.Time** | The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | **time.Time** | The end of a custom period | 
 **presetPeriod** | **string** | The requested time period. | [default to &quot;Last24Hours&quot;]

### Return type

[**AlertResponse**](AlertResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertGetMonitorGroupAlerts

> AlertResponse AlertGetMonitorGroupAlerts(ctx, monitorGroupGuid).IncludeReminders(includeReminders).Cursor(cursor).Sorting(sorting).Take(take).Start(start).End(end).PresetPeriod(presetPeriod).Execute()

Returns alerts for a specific monitor group.

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to get alerts for.
    includeReminders := true // bool | When true reminder alerts will be included in the results (optional) (default to false)
    cursor := "cursor_example" // string | A cursor value that should be used for traversing the dataset. (optional)
    sorting := "sorting_example" // string | Sorting direction based on timestamp. (optional) (default to "Descending")
    take := int32(56) // int32 | The number of records to return (Max value = 100) (optional) (default to 100)
    start := time.Now() // time.Time | The start of a custom period (can't be used together with the PresetPeriod parameter) (optional)
    end := time.Now() // time.Time | The end of a custom period (optional)
    presetPeriod := "presetPeriod_example" // string | The requested time period. (optional) (default to "Last24Hours")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertApi.AlertGetMonitorGroupAlerts(context.Background(), monitorGroupGuid).IncludeReminders(includeReminders).Cursor(cursor).Sorting(sorting).Take(take).Start(start).End(end).PresetPeriod(presetPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertApi.AlertGetMonitorGroupAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertGetMonitorGroupAlerts`: AlertResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertApi.AlertGetMonitorGroupAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to get alerts for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertGetMonitorGroupAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeReminders** | **bool** | When true reminder alerts will be included in the results | [default to false]
 **cursor** | **string** | A cursor value that should be used for traversing the dataset. | 
 **sorting** | **string** | Sorting direction based on timestamp. | [default to &quot;Descending&quot;]
 **take** | **int32** | The number of records to return (Max value &#x3D; 100) | [default to 100]
 **start** | **time.Time** | The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | **time.Time** | The end of a custom period | 
 **presetPeriod** | **string** | The requested time period. | [default to &quot;Last24Hours&quot;]

### Return type

[**AlertResponse**](AlertResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

