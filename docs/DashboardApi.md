# \DashboardApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardCloneDashboard**](DashboardApi.md#DashboardCloneDashboard) | **Post** /Dashboard/{dashboardGuid}/Clone | Clone the specified dashboard.
[**DashboardDeleteDashboard**](DashboardApi.md#DashboardDeleteDashboard) | **Delete** /Dashboard/{dashboardGuid} | Delete the specified dashboard.
[**DashboardGetAllDashboards**](DashboardApi.md#DashboardGetAllDashboards) | **Get** /Dashboard | Returns data for all dashboards.
[**DashboardGetOneDashboard**](DashboardApi.md#DashboardGetOneDashboard) | **Get** /Dashboard/{dashboardGuid} | Returns data for the specified dashboard.
[**DashboardPartiallyUpdateDashboard**](DashboardApi.md#DashboardPartiallyUpdateDashboard) | **Patch** /Dashboard/{dashboardGuid} | Partially update the specified dashboard.
[**DashboardUpdateDashboard**](DashboardApi.md#DashboardUpdateDashboard) | **Put** /Dashboard/{dashboardGuid} | Update the specified dashboard.



## DashboardCloneDashboard

> Dashboard DashboardCloneDashboard(ctx, dashboardGuid).Execute()

Clone the specified dashboard.

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
    dashboardGuid := "dashboardGuid_example" // string | The guid of the specified dashboard.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardApi.DashboardCloneDashboard(context.Background(), dashboardGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.DashboardCloneDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DashboardCloneDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardApi.DashboardCloneDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardGuid** | **string** | The guid of the specified dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardCloneDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardDeleteDashboard

> DashboardDeleteDashboard(ctx, dashboardGuid).Execute()

Delete the specified dashboard.

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
    dashboardGuid := "dashboardGuid_example" // string | The guid of the specified dashboard.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardApi.DashboardDeleteDashboard(context.Background(), dashboardGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.DashboardDeleteDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardGuid** | **string** | The guid of the specified dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardDeleteDashboardRequest struct via the builder pattern


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


## DashboardGetAllDashboards

> []Dashboard DashboardGetAllDashboards(ctx).Execute()

Returns data for all dashboards.

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
    resp, r, err := api_client.DashboardApi.DashboardGetAllDashboards(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.DashboardGetAllDashboards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DashboardGetAllDashboards`: []Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardApi.DashboardGetAllDashboards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardGetAllDashboardsRequest struct via the builder pattern


### Return type

[**[]Dashboard**](Dashboard.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardGetOneDashboard

> Dashboard DashboardGetOneDashboard(ctx, dashboardGuid).Execute()

Returns data for the specified dashboard.

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
    dashboardGuid := "dashboardGuid_example" // string | The guid of the specified dashboard.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardApi.DashboardGetOneDashboard(context.Background(), dashboardGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.DashboardGetOneDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DashboardGetOneDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardApi.DashboardGetOneDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardGuid** | **string** | The guid of the specified dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardGetOneDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardPartiallyUpdateDashboard

> DashboardPartiallyUpdateDashboard(ctx, dashboardGuid).Dashboard(dashboard).Execute()

Partially update the specified dashboard.

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
    dashboardGuid := "dashboardGuid_example" // string | The guid of the specified dashboard.
    dashboard := *openapiclient.NewDashboard() // Dashboard | The details for the dashboard.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardApi.DashboardPartiallyUpdateDashboard(context.Background(), dashboardGuid).Dashboard(dashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.DashboardPartiallyUpdateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardGuid** | **string** | The guid of the specified dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardPartiallyUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboard** | [**Dashboard**](Dashboard.md) | The details for the dashboard. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardUpdateDashboard

> DashboardUpdateDashboard(ctx, dashboardGuid).Dashboard(dashboard).Execute()

Update the specified dashboard.

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
    dashboardGuid := "dashboardGuid_example" // string | The guid of the specified dashboard.
    dashboard := *openapiclient.NewDashboard() // Dashboard | The details for the dashboard.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardApi.DashboardUpdateDashboard(context.Background(), dashboardGuid).Dashboard(dashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.DashboardUpdateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardGuid** | **string** | The guid of the specified dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboard** | [**Dashboard**](Dashboard.md) | The details for the dashboard. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

