# \AccountApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountGetAccountStatistics**](AccountApi.md#AccountGetAccountStatistics) | **Get** /Account | Returns the account statistics.



## AccountGetAccountStatistics

> AccountStatistics AccountGetAccountStatistics(ctx).Execute()

Returns the account statistics.

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
    resp, r, err := api_client.AccountApi.AccountGetAccountStatistics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountGetAccountStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGetAccountStatistics`: AccountStatistics
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountGetAccountStatistics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountGetAccountStatisticsRequest struct via the builder pattern


### Return type

[**AccountStatistics**](AccountStatistics.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

