# \RegisterApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegisterPost**](RegisterApi.md#RegisterPost) | **Post** /Register | Creates a new API account.



## RegisterPost

> RegistrationResponse RegisterPost(ctx).Execute()

Creates a new API account.



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
    resp, r, err := api_client.RegisterApi.RegisterPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegisterApi.RegisterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterPost`: RegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `RegisterApi.RegisterPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPostRequest struct via the builder pattern


### Return type

[**RegistrationResponse**](RegistrationResponse.md)

### Authorization

[user-basicauth](../README.md#user-basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

