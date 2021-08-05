# \ScheduledReportApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScheduledReportCreateScheduledReport**](ScheduledReportApi.md#ScheduledReportCreateScheduledReport) | **Post** /ScheduledReport | Creates a new scheduled report.
[**ScheduledReportDeleteSpecifiedScheduledReport**](ScheduledReportApi.md#ScheduledReportDeleteSpecifiedScheduledReport) | **Delete** /ScheduledReport/{scheduledReportGuid} | Delete the specified scheduled report.
[**ScheduledReportGetAllScheduledReports**](ScheduledReportApi.md#ScheduledReportGetAllScheduledReports) | **Get** /ScheduledReport | Returns data for all scheduled reports.
[**ScheduledReportGetSpecifiedScheduledReport**](ScheduledReportApi.md#ScheduledReportGetSpecifiedScheduledReport) | **Get** /ScheduledReport/{scheduledReportGuid} | Returns data for the specified scheduled report.
[**ScheduledReportPartiallyUpdateScheduledReport**](ScheduledReportApi.md#ScheduledReportPartiallyUpdateScheduledReport) | **Patch** /ScheduledReport/{scheduledReportGuid} | Partially update the specified scheduled report.
[**ScheduledReportUpdateScheduledReport**](ScheduledReportApi.md#ScheduledReportUpdateScheduledReport) | **Put** /ScheduledReport/{scheduledReportGuid} | Update the specified scheduled report.



## ScheduledReportCreateScheduledReport

> ScheduledReport ScheduledReportCreateScheduledReport(ctx).ScheduledReport(scheduledReport).Execute()

Creates a new scheduled report.

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
    scheduledReport := *openapiclient.NewScheduledReport() // ScheduledReport | The details for the scheduled report.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledReportApi.ScheduledReportCreateScheduledReport(context.Background()).ScheduledReport(scheduledReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledReportApi.ScheduledReportCreateScheduledReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduledReportCreateScheduledReport`: ScheduledReport
    fmt.Fprintf(os.Stdout, "Response from `ScheduledReportApi.ScheduledReportCreateScheduledReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduledReportCreateScheduledReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduledReport** | [**ScheduledReport**](ScheduledReport.md) | The details for the scheduled report. | 

### Return type

[**ScheduledReport**](ScheduledReport.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledReportDeleteSpecifiedScheduledReport

> ScheduledReportDeleteSpecifiedScheduledReport(ctx, scheduledReportGuid).Execute()

Delete the specified scheduled report.

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
    scheduledReportGuid := "scheduledReportGuid_example" // string | The guid of the specified scheduled report.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledReportApi.ScheduledReportDeleteSpecifiedScheduledReport(context.Background(), scheduledReportGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledReportApi.ScheduledReportDeleteSpecifiedScheduledReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduledReportGuid** | **string** | The guid of the specified scheduled report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduledReportDeleteSpecifiedScheduledReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledReportGetAllScheduledReports

> []ScheduledReport ScheduledReportGetAllScheduledReports(ctx).Execute()

Returns data for all scheduled reports.

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledReportApi.ScheduledReportGetAllScheduledReports(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledReportApi.ScheduledReportGetAllScheduledReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduledReportGetAllScheduledReports`: []ScheduledReport
    fmt.Fprintf(os.Stdout, "Response from `ScheduledReportApi.ScheduledReportGetAllScheduledReports`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScheduledReportGetAllScheduledReportsRequest struct via the builder pattern


### Return type

[**[]ScheduledReport**](ScheduledReport.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledReportGetSpecifiedScheduledReport

> ScheduledReport ScheduledReportGetSpecifiedScheduledReport(ctx, scheduledReportGuid).Execute()

Returns data for the specified scheduled report.

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
    scheduledReportGuid := "scheduledReportGuid_example" // string | The guid of the specified scheduled report.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledReportApi.ScheduledReportGetSpecifiedScheduledReport(context.Background(), scheduledReportGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledReportApi.ScheduledReportGetSpecifiedScheduledReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduledReportGetSpecifiedScheduledReport`: ScheduledReport
    fmt.Fprintf(os.Stdout, "Response from `ScheduledReportApi.ScheduledReportGetSpecifiedScheduledReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduledReportGuid** | **string** | The guid of the specified scheduled report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduledReportGetSpecifiedScheduledReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduledReport**](ScheduledReport.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledReportPartiallyUpdateScheduledReport

> ScheduledReport ScheduledReportPartiallyUpdateScheduledReport(ctx, scheduledReportGuid).ScheduledReport(scheduledReport).Execute()

Partially update the specified scheduled report.

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
    scheduledReportGuid := "scheduledReportGuid_example" // string | The guid of the specified scheduled report.
    scheduledReport := *openapiclient.NewScheduledReport() // ScheduledReport | The details for the scheduled report.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledReportApi.ScheduledReportPartiallyUpdateScheduledReport(context.Background(), scheduledReportGuid).ScheduledReport(scheduledReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledReportApi.ScheduledReportPartiallyUpdateScheduledReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduledReportPartiallyUpdateScheduledReport`: ScheduledReport
    fmt.Fprintf(os.Stdout, "Response from `ScheduledReportApi.ScheduledReportPartiallyUpdateScheduledReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduledReportGuid** | **string** | The guid of the specified scheduled report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduledReportPartiallyUpdateScheduledReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduledReport** | [**ScheduledReport**](ScheduledReport.md) | The details for the scheduled report. | 

### Return type

[**ScheduledReport**](ScheduledReport.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledReportUpdateScheduledReport

> ScheduledReport ScheduledReportUpdateScheduledReport(ctx, scheduledReportGuid).ScheduledReport(scheduledReport).Execute()

Update the specified scheduled report.

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
    scheduledReportGuid := "scheduledReportGuid_example" // string | The guid of the specified scheduled report.
    scheduledReport := *openapiclient.NewScheduledReport() // ScheduledReport | The details for the scheduled report.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledReportApi.ScheduledReportUpdateScheduledReport(context.Background(), scheduledReportGuid).ScheduledReport(scheduledReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledReportApi.ScheduledReportUpdateScheduledReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduledReportUpdateScheduledReport`: ScheduledReport
    fmt.Fprintf(os.Stdout, "Response from `ScheduledReportApi.ScheduledReportUpdateScheduledReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduledReportGuid** | **string** | The guid of the specified scheduled report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduledReportUpdateScheduledReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduledReport** | [**ScheduledReport**](ScheduledReport.md) | The details for the scheduled report. | 

### Return type

[**ScheduledReport**](ScheduledReport.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

