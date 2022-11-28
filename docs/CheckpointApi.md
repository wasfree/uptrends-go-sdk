# \CheckpointApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckpointGetAllCheckpoints**](CheckpointApi.md#CheckpointGetAllCheckpoints) | **Get** /Checkpoint | Returns all the checkpoints that are not deleted
[**CheckpointGetCheckpoint**](CheckpointApi.md#CheckpointGetCheckpoint) | **Get** /Checkpoint/{checkpointId} | Returns the specified checkpoint, deleted or not
[**CheckpointRegionGetAllCheckpointRegions**](CheckpointApi.md#CheckpointRegionGetAllCheckpointRegions) | **Get** /CheckpointRegion | Returns all the checkpoint regions.
[**CheckpointRegionGetCheckpointRegionCheckpoints**](CheckpointApi.md#CheckpointRegionGetCheckpointRegionCheckpoints) | **Get** /CheckpointRegion/{checkpointRegionId}/Checkpoint | Returns the checkpoints for the specified checkpoint region.
[**CheckpointRegionGetSpecifiedCheckpointRegion**](CheckpointApi.md#CheckpointRegionGetSpecifiedCheckpointRegion) | **Get** /CheckpointRegion/{checkpointRegionId} | Returns the specified checkpoint region.
[**CheckpointServerGetAllServerIpv4Addresses**](CheckpointApi.md#CheckpointServerGetAllServerIpv4Addresses) | **Get** /Checkpoint/Server/Ipv4 | Anonymous call that returns all the IPv4 addresses of all the checkpoint servers.
[**CheckpointServerGetAllServerIpv6Addresses**](CheckpointApi.md#CheckpointServerGetAllServerIpv6Addresses) | **Get** /Checkpoint/Server/Ipv6 | Anonymous call that returns all the IPv6 addresses of all the checkpoint servers.
[**CheckpointServerGetServer**](CheckpointApi.md#CheckpointServerGetServer) | **Get** /Checkpoint/Server/{serverId} | Returns the requested checkpoint server.
[**PrivateCheckpointHealthGetPrivateCheckpointHealthForRegion**](CheckpointApi.md#PrivateCheckpointHealthGetPrivateCheckpointHealthForRegion) | **Get** /PrivateCheckpointHealthForRegion | Returns the status of the private checkpoints in the given region.
[**PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth**](CheckpointApi.md#PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth) | **Get** /PrivateCheckpointHealth | Returns the status of the specified private checkpoints.



## CheckpointGetAllCheckpoints

> CheckpointListResponse CheckpointGetAllCheckpoints(ctx).Execute()

Returns all the checkpoints that are not deleted

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
    resp, r, err := apiClient.CheckpointApi.CheckpointGetAllCheckpoints(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckpointApi.CheckpointGetAllCheckpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckpointGetAllCheckpoints`: CheckpointListResponse
    fmt.Fprintf(os.Stdout, "Response from `CheckpointApi.CheckpointGetAllCheckpoints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckpointGetAllCheckpointsRequest struct via the builder pattern


### Return type

[**CheckpointListResponse**](CheckpointListResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckpointGetCheckpoint

> CheckpointResponse CheckpointGetCheckpoint(ctx, checkpointId).Execute()

Returns the specified checkpoint, deleted or not

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
    checkpointId := int32(56) // int32 | The Id of the requested checkpoint.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckpointApi.CheckpointGetCheckpoint(context.Background(), checkpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckpointApi.CheckpointGetCheckpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckpointGetCheckpoint`: CheckpointResponse
    fmt.Fprintf(os.Stdout, "Response from `CheckpointApi.CheckpointGetCheckpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkpointId** | **int32** | The Id of the requested checkpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckpointGetCheckpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckpointResponse**](CheckpointResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckpointRegionGetAllCheckpointRegions

> []CheckpointRegion CheckpointRegionGetAllCheckpointRegions(ctx).Execute()

Returns all the checkpoint regions.

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
    resp, r, err := apiClient.CheckpointApi.CheckpointRegionGetAllCheckpointRegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckpointApi.CheckpointRegionGetAllCheckpointRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckpointRegionGetAllCheckpointRegions`: []CheckpointRegion
    fmt.Fprintf(os.Stdout, "Response from `CheckpointApi.CheckpointRegionGetAllCheckpointRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckpointRegionGetAllCheckpointRegionsRequest struct via the builder pattern


### Return type

[**[]CheckpointRegion**](CheckpointRegion.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckpointRegionGetCheckpointRegionCheckpoints

> []Checkpoint CheckpointRegionGetCheckpointRegionCheckpoints(ctx, checkpointRegionId).Execute()

Returns the checkpoints for the specified checkpoint region.

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
    checkpointRegionId := int32(56) // int32 | The id for the specified checkpoint region.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckpointApi.CheckpointRegionGetCheckpointRegionCheckpoints(context.Background(), checkpointRegionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckpointApi.CheckpointRegionGetCheckpointRegionCheckpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckpointRegionGetCheckpointRegionCheckpoints`: []Checkpoint
    fmt.Fprintf(os.Stdout, "Response from `CheckpointApi.CheckpointRegionGetCheckpointRegionCheckpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkpointRegionId** | **int32** | The id for the specified checkpoint region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckpointRegionGetCheckpointRegionCheckpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Checkpoint**](Checkpoint.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckpointRegionGetSpecifiedCheckpointRegion

> CheckpointRegion CheckpointRegionGetSpecifiedCheckpointRegion(ctx, checkpointRegionId).Execute()

Returns the specified checkpoint region.

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
    checkpointRegionId := int32(56) // int32 | The id for the specified checkpoint region.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckpointApi.CheckpointRegionGetSpecifiedCheckpointRegion(context.Background(), checkpointRegionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckpointApi.CheckpointRegionGetSpecifiedCheckpointRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckpointRegionGetSpecifiedCheckpointRegion`: CheckpointRegion
    fmt.Fprintf(os.Stdout, "Response from `CheckpointApi.CheckpointRegionGetSpecifiedCheckpointRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkpointRegionId** | **int32** | The id for the specified checkpoint region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckpointRegionGetSpecifiedCheckpointRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckpointRegion**](CheckpointRegion.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckpointServerGetAllServerIpv4Addresses

> ListStringResponse CheckpointServerGetAllServerIpv4Addresses(ctx).Execute()

Anonymous call that returns all the IPv4 addresses of all the checkpoint servers.

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
    resp, r, err := apiClient.CheckpointApi.CheckpointServerGetAllServerIpv4Addresses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckpointApi.CheckpointServerGetAllServerIpv4Addresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckpointServerGetAllServerIpv4Addresses`: ListStringResponse
    fmt.Fprintf(os.Stdout, "Response from `CheckpointApi.CheckpointServerGetAllServerIpv4Addresses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckpointServerGetAllServerIpv4AddressesRequest struct via the builder pattern


### Return type

[**ListStringResponse**](ListStringResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckpointServerGetAllServerIpv6Addresses

> ListStringResponse CheckpointServerGetAllServerIpv6Addresses(ctx).Execute()

Anonymous call that returns all the IPv6 addresses of all the checkpoint servers.

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
    resp, r, err := apiClient.CheckpointApi.CheckpointServerGetAllServerIpv6Addresses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckpointApi.CheckpointServerGetAllServerIpv6Addresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckpointServerGetAllServerIpv6Addresses`: ListStringResponse
    fmt.Fprintf(os.Stdout, "Response from `CheckpointApi.CheckpointServerGetAllServerIpv6Addresses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckpointServerGetAllServerIpv6AddressesRequest struct via the builder pattern


### Return type

[**ListStringResponse**](ListStringResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckpointServerGetServer

> CheckpoinServerResponse CheckpointServerGetServer(ctx, serverId).Execute()

Returns the requested checkpoint server.

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
    serverId := int32(56) // int32 | The Id of the requested server.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckpointApi.CheckpointServerGetServer(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckpointApi.CheckpointServerGetServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckpointServerGetServer`: CheckpoinServerResponse
    fmt.Fprintf(os.Stdout, "Response from `CheckpointApi.CheckpointServerGetServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | The Id of the requested server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckpointServerGetServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckpoinServerResponse**](CheckpoinServerResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCheckpointHealthGetPrivateCheckpointHealthForRegion

> CheckpointsHealth PrivateCheckpointHealthGetPrivateCheckpointHealthForRegion(ctx).RegionId(regionId).Execute()

Returns the status of the private checkpoints in the given region.

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
    regionId := int32(56) // int32 | The id of the region.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckpointApi.PrivateCheckpointHealthGetPrivateCheckpointHealthForRegion(context.Background()).RegionId(regionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckpointApi.PrivateCheckpointHealthGetPrivateCheckpointHealthForRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateCheckpointHealthGetPrivateCheckpointHealthForRegion`: CheckpointsHealth
    fmt.Fprintf(os.Stdout, "Response from `CheckpointApi.PrivateCheckpointHealthGetPrivateCheckpointHealthForRegion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCheckpointHealthGetPrivateCheckpointHealthForRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **int32** | The id of the region. | 

### Return type

[**CheckpointsHealth**](CheckpointsHealth.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth

> CheckpointsHealth PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth(ctx).Filter(filter).Execute()

Returns the status of the specified private checkpoints.

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
    filter := "filter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckpointApi.PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckpointApi.PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth`: CheckpointsHealth
    fmt.Fprintf(os.Stdout, "Response from `CheckpointApi.PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |  | 

### Return type

[**CheckpointsHealth**](CheckpointsHealth.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

