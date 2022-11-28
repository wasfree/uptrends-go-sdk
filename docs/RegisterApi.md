# \RegisterApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegisterPost**](RegisterApi.md#RegisterPost) | **Post** /Register | Creates a new API account.



## RegisterPost

> RegistrationResponse RegisterPost(ctx).Description(description).OperatorGuid(operatorGuid).Execute()

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
    description := "description_example" // string | An optional description for the new API account, e.g. \"API\". If this is empty, it will be defaulted to \"API\" (optional) (default to "API")
    operatorGuid := "operatorGuid_example" // string | The operator guid for which the new API account needs to be created. Leave empty to create an API acount for your own operator. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegisterApi.RegisterPost(context.Background()).Description(description).OperatorGuid(operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegisterApi.RegisterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterPost`: RegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `RegisterApi.RegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | An optional description for the new API account, e.g. \&quot;API\&quot;. If this is empty, it will be defaulted to \&quot;API\&quot; | [default to &quot;API&quot;]
 **operatorGuid** | **string** | The operator guid for which the new API account needs to be created. Leave empty to create an API acount for your own operator. | 

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

