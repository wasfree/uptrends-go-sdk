# \TimezoneApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TimezoneGetAllTimezones**](TimezoneApi.md#TimezoneGetAllTimezones) | **Get** /Timezone | Gets all timezones available.
[**TimezoneGetTimezoneById**](TimezoneApi.md#TimezoneGetTimezoneById) | **Get** /Timezone/{timezoneId} | Gets the timezone with the specified Id.



## TimezoneGetAllTimezones

> []Timezone TimezoneGetAllTimezones(ctx).Execute()

Gets all timezones available.

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
    resp, r, err := apiClient.TimezoneApi.TimezoneGetAllTimezones(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimezoneApi.TimezoneGetAllTimezones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TimezoneGetAllTimezones`: []Timezone
    fmt.Fprintf(os.Stdout, "Response from `TimezoneApi.TimezoneGetAllTimezones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTimezoneGetAllTimezonesRequest struct via the builder pattern


### Return type

[**[]Timezone**](Timezone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimezoneGetTimezoneById

> Timezone TimezoneGetTimezoneById(ctx, timezoneId).Execute()

Gets the timezone with the specified Id.

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
    timezoneId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimezoneApi.TimezoneGetTimezoneById(context.Background(), timezoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimezoneApi.TimezoneGetTimezoneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TimezoneGetTimezoneById`: Timezone
    fmt.Fprintf(os.Stdout, "Response from `TimezoneApi.TimezoneGetTimezoneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timezoneId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTimezoneGetTimezoneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Timezone**](Timezone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

