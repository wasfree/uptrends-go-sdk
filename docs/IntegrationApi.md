# \IntegrationApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationGetAllIntegrations**](IntegrationApi.md#IntegrationGetAllIntegrations) | **Get** /Integration | Gets a list of all integrations.



## IntegrationGetAllIntegrations

> []Integration IntegrationGetAllIntegrations(ctx).Execute()

Gets a list of all integrations.

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
    resp, r, err := api_client.IntegrationApi.IntegrationGetAllIntegrations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationApi.IntegrationGetAllIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationGetAllIntegrations`: []Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationApi.IntegrationGetAllIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationGetAllIntegrationsRequest struct via the builder pattern


### Return type

[**[]Integration**](Integration.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

