# \OutgoingPhoneNumberApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OutgoingPhoneNumberGetAllOutgoingPhoneNumbers**](OutgoingPhoneNumberApi.md#OutgoingPhoneNumberGetAllOutgoingPhoneNumbers) | **Get** /OutgoingPhoneNumber | Gets a list of all outgoing phone numbers available.



## OutgoingPhoneNumberGetAllOutgoingPhoneNumbers

> []OutgoingPhoneNumberDetails OutgoingPhoneNumberGetAllOutgoingPhoneNumbers(ctx).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OutgoingPhoneNumberApi.OutgoingPhoneNumberGetAllOutgoingPhoneNumbers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutgoingPhoneNumberApi.OutgoingPhoneNumberGetAllOutgoingPhoneNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutgoingPhoneNumberGetAllOutgoingPhoneNumbers`: []OutgoingPhoneNumberDetails
    fmt.Fprintf(os.Stdout, "Response from `OutgoingPhoneNumberApi.OutgoingPhoneNumberGetAllOutgoingPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOutgoingPhoneNumberGetAllOutgoingPhoneNumbersRequest struct via the builder pattern


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

