# \MiscellaneousApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MiscellaneousGetAllOutgoingPhoneNumbers**](MiscellaneousApi.md#MiscellaneousGetAllOutgoingPhoneNumbers) | **Get** /OutgoingPhoneNumber | Gets a list of all outgoing phone numbers available.
[**MiscellaneousGetAllTimezones**](MiscellaneousApi.md#MiscellaneousGetAllTimezones) | **Get** /Timezone | Gets all timezones available.
[**MiscellaneousGetTimezoneById**](MiscellaneousApi.md#MiscellaneousGetTimezoneById) | **Get** /Timezone/{timezoneId} | Gets the timezone with the specified Id.



## MiscellaneousGetAllOutgoingPhoneNumbers

> []OutgoingPhoneNumberDetails MiscellaneousGetAllOutgoingPhoneNumbers(ctx).Execute()

Gets a list of all outgoing phone numbers available.

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
    resp, r, err := api_client.MiscellaneousApi.MiscellaneousGetAllOutgoingPhoneNumbers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.MiscellaneousGetAllOutgoingPhoneNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiscellaneousGetAllOutgoingPhoneNumbers`: []OutgoingPhoneNumberDetails
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.MiscellaneousGetAllOutgoingPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMiscellaneousGetAllOutgoingPhoneNumbersRequest struct via the builder pattern


### Return type

[**[]OutgoingPhoneNumberDetails**](OutgoingPhoneNumberDetails.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MiscellaneousGetAllTimezones

> []Timezone MiscellaneousGetAllTimezones(ctx).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiscellaneousApi.MiscellaneousGetAllTimezones(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.MiscellaneousGetAllTimezones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiscellaneousGetAllTimezones`: []Timezone
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.MiscellaneousGetAllTimezones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMiscellaneousGetAllTimezonesRequest struct via the builder pattern


### Return type

[**[]Timezone**](Timezone.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MiscellaneousGetTimezoneById

> Timezone MiscellaneousGetTimezoneById(ctx, timezoneId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiscellaneousApi.MiscellaneousGetTimezoneById(context.Background(), timezoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.MiscellaneousGetTimezoneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiscellaneousGetTimezoneById`: Timezone
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.MiscellaneousGetTimezoneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timezoneId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMiscellaneousGetTimezoneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Timezone**](Timezone.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

