# \MonitorGroupApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitorGroupAddMaintenancePeriodToAllMembers**](MonitorGroupApi.md#MonitorGroupAddMaintenancePeriodToAllMembers) | **Post** /MonitorGroup/{monitorGroupGuid}/AddMaintenancePeriodToAllMembers | Adds the provided maintenance period to all monitors in the group specified 
[**MonitorGroupAddMonitorToMonitorGroup**](MonitorGroupApi.md#MonitorGroupAddMonitorToMonitorGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/Member/{monitorGuid} | Adds the specified monitor to the monitor group 
[**MonitorGroupCreateMonitorGroup**](MonitorGroupApi.md#MonitorGroupCreateMonitorGroup) | **Post** /MonitorGroup | Creates a new monitor group
[**MonitorGroupDeleteMonitorGroup**](MonitorGroupApi.md#MonitorGroupDeleteMonitorGroup) | **Delete** /MonitorGroup/{monitorGroupGuid} | Deletes the specified monitor group
[**MonitorGroupGetAllMonitorGroups**](MonitorGroupApi.md#MonitorGroupGetAllMonitorGroups) | **Get** /MonitorGroup | Gets all monitor groups
[**MonitorGroupGetMonitorGroup**](MonitorGroupApi.md#MonitorGroupGetMonitorGroup) | **Get** /MonitorGroup/{monitorGroupGuid} | Gets the details of a monitor group
[**MonitorGroupGetMonitorGroupMembers**](MonitorGroupApi.md#MonitorGroupGetMonitorGroupMembers) | **Get** /MonitorGroup/{monitorGroupGuid}/Member | Gets a list of all members of a monitor group
[**MonitorGroupRemoveMonitorFromMonitorGroup**](MonitorGroupApi.md#MonitorGroupRemoveMonitorFromMonitorGroup) | **Delete** /MonitorGroup/{monitorGroupGuid}/Member/{monitorGuid} | Removes the specified monitor from the monitor group
[**MonitorGroupStartAllMonitorAlertsInGroup**](MonitorGroupApi.md#MonitorGroupStartAllMonitorAlertsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StartAllMonitorAlerts | Starts alerting for all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupStartAllMonitorsInGroup**](MonitorGroupApi.md#MonitorGroupStartAllMonitorsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StartAllMonitors | Starts all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupStopAllMonitorAlertsInGroup**](MonitorGroupApi.md#MonitorGroupStopAllMonitorAlertsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StopAllMonitorAlerts | Stops alerting for all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupStopAllMonitorsInGroup**](MonitorGroupApi.md#MonitorGroupStopAllMonitorsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StopAllMonitors | Stops all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupUpdateMonitorGroup**](MonitorGroupApi.md#MonitorGroupUpdateMonitorGroup) | **Put** /MonitorGroup/{monitorGroupGuid} | Updates the monitor group with the Guid specified



## MonitorGroupAddMaintenancePeriodToAllMembers

> MonitorGroupAddMaintenancePeriodToAllMembers(ctx, monitorGroupGuid).MaintenancePeriod(maintenancePeriod).Execute()

Adds the provided maintenance period to all monitors in the group specified 

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | 
    maintenancePeriod := *openapiclient.NewMaintenancePeriod(int32(123), openapiclient.ScheduleMode("OneTime"), openapiclient.MaintenanceTypes("DisableMonitoring")) // MaintenancePeriod | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupAddMaintenancePeriodToAllMembers(context.Background(), monitorGroupGuid).MaintenancePeriod(maintenancePeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupAddMaintenancePeriodToAllMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupAddMaintenancePeriodToAllMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenancePeriod** | [**MaintenancePeriod**](MaintenancePeriod.md) |  | 

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


## MonitorGroupAddMonitorToMonitorGroup

> MonitorGroupAddMonitorToMonitorGroup(ctx, monitorGroupGuid, monitorGuid).Execute()

Adds the specified monitor to the monitor group 

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to add the monitor to
    monitorGuid := "monitorGuid_example" // string | The monitor Guid

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupAddMonitorToMonitorGroup(context.Background(), monitorGroupGuid, monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupAddMonitorToMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to add the monitor to | 
**monitorGuid** | **string** | The monitor Guid | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupAddMonitorToMonitorGroupRequest struct via the builder pattern


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


## MonitorGroupCreateMonitorGroup

> MonitorGroup MonitorGroupCreateMonitorGroup(ctx).MonitorGroup(monitorGroup).Execute()

Creates a new monitor group

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
    monitorGroup := *openapiclient.NewMonitorGroup(false) // MonitorGroup | The MonitorGroup object to be created

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupCreateMonitorGroup(context.Background()).MonitorGroup(monitorGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupCreateMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGroupCreateMonitorGroup`: MonitorGroup
    fmt.Fprintf(os.Stdout, "Response from `MonitorGroupApi.MonitorGroupCreateMonitorGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupCreateMonitorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitorGroup** | [**MonitorGroup**](MonitorGroup.md) | The MonitorGroup object to be created | 

### Return type

[**MonitorGroup**](MonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGroupDeleteMonitorGroup

> MonitorGroupDeleteMonitorGroup(ctx, monitorGroupGuid).Execute()

Deletes the specified monitor group

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to be deleted

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupDeleteMonitorGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupDeleteMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupDeleteMonitorGroupRequest struct via the builder pattern


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


## MonitorGroupGetAllMonitorGroups

> []MonitorGroup MonitorGroupGetAllMonitorGroups(ctx).Execute()

Gets all monitor groups

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
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupGetAllMonitorGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupGetAllMonitorGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGroupGetAllMonitorGroups`: []MonitorGroup
    fmt.Fprintf(os.Stdout, "Response from `MonitorGroupApi.MonitorGroupGetAllMonitorGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupGetAllMonitorGroupsRequest struct via the builder pattern


### Return type

[**[]MonitorGroup**](MonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGroupGetMonitorGroup

> MonitorGroup MonitorGroupGetMonitorGroup(ctx, monitorGroupGuid).Execute()

Gets the details of a monitor group

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to be retrieved

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupGetMonitorGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupGetMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGroupGetMonitorGroup`: MonitorGroup
    fmt.Fprintf(os.Stdout, "Response from `MonitorGroupApi.MonitorGroupGetMonitorGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupGetMonitorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorGroup**](MonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGroupGetMonitorGroupMembers

> MonitorGroupMember MonitorGroupGetMonitorGroupMembers(ctx, monitorGroupGuid).Execute()

Gets a list of all members of a monitor group

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to retrieve the members for

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupGetMonitorGroupMembers(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupGetMonitorGroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGroupGetMonitorGroupMembers`: MonitorGroupMember
    fmt.Fprintf(os.Stdout, "Response from `MonitorGroupApi.MonitorGroupGetMonitorGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to retrieve the members for | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupGetMonitorGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorGroupMember**](MonitorGroupMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGroupRemoveMonitorFromMonitorGroup

> MonitorGroupRemoveMonitorFromMonitorGroup(ctx, monitorGroupGuid, monitorGuid).Execute()

Removes the specified monitor from the monitor group

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to remove the monitor from
    monitorGuid := "monitorGuid_example" // string | The monitor Guid to be removed

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupRemoveMonitorFromMonitorGroup(context.Background(), monitorGroupGuid, monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupRemoveMonitorFromMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to remove the monitor from | 
**monitorGuid** | **string** | The monitor Guid to be removed | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupRemoveMonitorFromMonitorGroupRequest struct via the builder pattern


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


## MonitorGroupStartAllMonitorAlertsInGroup

> MonitorGroupStartAllMonitorAlertsInGroup(ctx, monitorGroupGuid).Execute()

Starts alerting for all monitors that are a member of the monitor group specified by the monitor group GUID

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The monitor group GUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupStartAllMonitorAlertsInGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupStartAllMonitorAlertsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The monitor group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupStartAllMonitorAlertsInGroupRequest struct via the builder pattern


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


## MonitorGroupStartAllMonitorsInGroup

> MonitorGroupStartAllMonitorsInGroup(ctx, monitorGroupGuid).Execute()

Starts all monitors that are a member of the monitor group specified by the monitor group GUID

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The monitor group GUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupStartAllMonitorsInGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupStartAllMonitorsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The monitor group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupStartAllMonitorsInGroupRequest struct via the builder pattern


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


## MonitorGroupStopAllMonitorAlertsInGroup

> MonitorGroupStopAllMonitorAlertsInGroup(ctx, monitorGroupGuid).Execute()

Stops alerting for all monitors that are a member of the monitor group specified by the monitor group GUID

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The monitor group GUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupStopAllMonitorAlertsInGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupStopAllMonitorAlertsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The monitor group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupStopAllMonitorAlertsInGroupRequest struct via the builder pattern


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


## MonitorGroupStopAllMonitorsInGroup

> MonitorGroupStopAllMonitorsInGroup(ctx, monitorGroupGuid).Execute()

Stops all monitors that are a member of the monitor group specified by the monitor group GUID

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The monitor group GUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupStopAllMonitorsInGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupStopAllMonitorsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The monitor group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupStopAllMonitorsInGroupRequest struct via the builder pattern


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


## MonitorGroupUpdateMonitorGroup

> MonitorGroupUpdateMonitorGroup(ctx, monitorGroupGuid).Item(item).Execute()

Updates the monitor group with the Guid specified

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to be updated
    item := *openapiclient.NewMonitorGroup(false) // MonitorGroup | The monitor group to be updated

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorGroupApi.MonitorGroupUpdateMonitorGroup(context.Background(), monitorGroupGuid).Item(item).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupApi.MonitorGroupUpdateMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupUpdateMonitorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **item** | [**MonitorGroup**](MonitorGroup.md) | The monitor group to be updated | 

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

