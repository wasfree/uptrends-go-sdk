# \PublicStatusPageApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicStatusPageAddAuthorizationToPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPageAddAuthorizationToPublicStatusPage) | **Post** /PublicStatusPage/{publicStatusPageGuid}/Authorization | Creates a new authorization for the specified public status page.
[**PublicStatusPageDeletePublicStatusPage**](PublicStatusPageApi.md#PublicStatusPageDeletePublicStatusPage) | **Delete** /PublicStatusPage/{publicStatusPageGuid} | Deletes the definition of the specified public status page.
[**PublicStatusPageGetAuthorizationsForPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPageGetAuthorizationsForPublicStatusPage) | **Get** /PublicStatusPage/{publicStatusPageGuid}/Authorization | Returns all authorizations for the specified public status page.
[**PublicStatusPageGetPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPageGetPublicStatusPage) | **Get** /PublicStatusPage/{publicStatusPageGuid} | Returns the definition of the specified public status page.
[**PublicStatusPageGetPublicStatusPages**](PublicStatusPageApi.md#PublicStatusPageGetPublicStatusPages) | **Get** /PublicStatusPage | Returns the definition of all public status pages available in the account.
[**PublicStatusPagePatchPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPagePatchPublicStatusPage) | **Patch** /PublicStatusPage/{publicStatusPageGuid} | Partially updates the definition of the specified public status page.
[**PublicStatusPagePostPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPagePostPublicStatusPage) | **Post** /PublicStatusPage | Creates a new public status page.
[**PublicStatusPagePutPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPagePutPublicStatusPage) | **Put** /PublicStatusPage/{publicStatusPageGuid} | Updates the definition of the specified public status page.
[**PublicStatusPageRemoveAuthorizationFromPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPageRemoveAuthorizationFromPublicStatusPage) | **Delete** /PublicStatusPage/{publicStatusPageGuid}/Authorization/{authorizationGuid} | Removes an authorization from a public status page.



## PublicStatusPageAddAuthorizationToPublicStatusPage

> PSPAuthorization PublicStatusPageAddAuthorizationToPublicStatusPage(ctx, publicStatusPageGuid).Authorization(authorization).Execute()

Creates a new authorization for the specified public status page.



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
    publicStatusPageGuid := "publicStatusPageGuid_example" // string | The Guid of the public status page.
    authorization := *openapiclient.NewPSPAuthorization(openapiclient.PSPAuthorizationType("ViewPublicDashboard")) // PSPAuthorization | The complete definition of the authorization that should be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicStatusPageApi.PublicStatusPageAddAuthorizationToPublicStatusPage(context.Background(), publicStatusPageGuid).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicStatusPageApi.PublicStatusPageAddAuthorizationToPublicStatusPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicStatusPageAddAuthorizationToPublicStatusPage`: PSPAuthorization
    fmt.Fprintf(os.Stdout, "Response from `PublicStatusPageApi.PublicStatusPageAddAuthorizationToPublicStatusPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicStatusPageGuid** | **string** | The Guid of the public status page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicStatusPageAddAuthorizationToPublicStatusPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | [**PSPAuthorization**](PSPAuthorization.md) | The complete definition of the authorization that should be created. | 

### Return type

[**PSPAuthorization**](PSPAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicStatusPageDeletePublicStatusPage

> PublicStatusPageDeletePublicStatusPage(ctx, publicStatusPageGuid).Execute()

Deletes the definition of the specified public status page.

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
    publicStatusPageGuid := "publicStatusPageGuid_example" // string | The Guid of the public status page that should be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicStatusPageApi.PublicStatusPageDeletePublicStatusPage(context.Background(), publicStatusPageGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicStatusPageApi.PublicStatusPageDeletePublicStatusPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicStatusPageGuid** | **string** | The Guid of the public status page that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicStatusPageDeletePublicStatusPageRequest struct via the builder pattern


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


## PublicStatusPageGetAuthorizationsForPublicStatusPage

> []PSPAuthorization PublicStatusPageGetAuthorizationsForPublicStatusPage(ctx, publicStatusPageGuid).Execute()

Returns all authorizations for the specified public status page.

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
    publicStatusPageGuid := "publicStatusPageGuid_example" // string | The Guid of the public status page.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicStatusPageApi.PublicStatusPageGetAuthorizationsForPublicStatusPage(context.Background(), publicStatusPageGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicStatusPageApi.PublicStatusPageGetAuthorizationsForPublicStatusPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicStatusPageGetAuthorizationsForPublicStatusPage`: []PSPAuthorization
    fmt.Fprintf(os.Stdout, "Response from `PublicStatusPageApi.PublicStatusPageGetAuthorizationsForPublicStatusPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicStatusPageGuid** | **string** | The Guid of the public status page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicStatusPageGetAuthorizationsForPublicStatusPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PSPAuthorization**](PSPAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicStatusPageGetPublicStatusPage

> PublicStatusPage PublicStatusPageGetPublicStatusPage(ctx, publicStatusPageGuid).Execute()

Returns the definition of the specified public status page.

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
    publicStatusPageGuid := "publicStatusPageGuid_example" // string | The Guid of the requested public status page.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicStatusPageApi.PublicStatusPageGetPublicStatusPage(context.Background(), publicStatusPageGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicStatusPageApi.PublicStatusPageGetPublicStatusPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicStatusPageGetPublicStatusPage`: PublicStatusPage
    fmt.Fprintf(os.Stdout, "Response from `PublicStatusPageApi.PublicStatusPageGetPublicStatusPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicStatusPageGuid** | **string** | The Guid of the requested public status page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicStatusPageGetPublicStatusPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicStatusPage**](PublicStatusPage.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicStatusPageGetPublicStatusPages

> []PublicStatusPage PublicStatusPageGetPublicStatusPages(ctx).Execute()

Returns the definition of all public status pages available in the account.

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
    resp, r, err := api_client.PublicStatusPageApi.PublicStatusPageGetPublicStatusPages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicStatusPageApi.PublicStatusPageGetPublicStatusPages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicStatusPageGetPublicStatusPages`: []PublicStatusPage
    fmt.Fprintf(os.Stdout, "Response from `PublicStatusPageApi.PublicStatusPageGetPublicStatusPages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPublicStatusPageGetPublicStatusPagesRequest struct via the builder pattern


### Return type

[**[]PublicStatusPage**](PublicStatusPage.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicStatusPagePatchPublicStatusPage

> PublicStatusPagePatchPublicStatusPage(ctx, publicStatusPageGuid).PublicStatusPage(publicStatusPage).Execute()

Partially updates the definition of the specified public status page.



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
    publicStatusPageGuid := "publicStatusPageGuid_example" // string | The Guid of the public status page that should be updated.
    publicStatusPage := *openapiclient.NewPublicStatusPage() // PublicStatusPage | The partial definition for the public status page that should be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicStatusPageApi.PublicStatusPagePatchPublicStatusPage(context.Background(), publicStatusPageGuid).PublicStatusPage(publicStatusPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicStatusPageApi.PublicStatusPagePatchPublicStatusPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicStatusPageGuid** | **string** | The Guid of the public status page that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicStatusPagePatchPublicStatusPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicStatusPage** | [**PublicStatusPage**](PublicStatusPage.md) | The partial definition for the public status page that should be updated. | 

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


## PublicStatusPagePostPublicStatusPage

> PublicStatusPage PublicStatusPagePostPublicStatusPage(ctx).PublicStatusPage(publicStatusPage).Execute()

Creates a new public status page.

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
    publicStatusPage := *openapiclient.NewPublicStatusPage() // PublicStatusPage | The complete definition for the public status page that should be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicStatusPageApi.PublicStatusPagePostPublicStatusPage(context.Background()).PublicStatusPage(publicStatusPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicStatusPageApi.PublicStatusPagePostPublicStatusPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicStatusPagePostPublicStatusPage`: PublicStatusPage
    fmt.Fprintf(os.Stdout, "Response from `PublicStatusPageApi.PublicStatusPagePostPublicStatusPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicStatusPagePostPublicStatusPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicStatusPage** | [**PublicStatusPage**](PublicStatusPage.md) | The complete definition for the public status page that should be updated. | 

### Return type

[**PublicStatusPage**](PublicStatusPage.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicStatusPagePutPublicStatusPage

> PublicStatusPagePutPublicStatusPage(ctx, publicStatusPageGuid).PublicStatusPage(publicStatusPage).Execute()

Updates the definition of the specified public status page.



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
    publicStatusPageGuid := "publicStatusPageGuid_example" // string | The Guid of the public status page that should be updated.
    publicStatusPage := *openapiclient.NewPublicStatusPage() // PublicStatusPage | The complete definition for the public status page that should be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicStatusPageApi.PublicStatusPagePutPublicStatusPage(context.Background(), publicStatusPageGuid).PublicStatusPage(publicStatusPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicStatusPageApi.PublicStatusPagePutPublicStatusPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicStatusPageGuid** | **string** | The Guid of the public status page that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicStatusPagePutPublicStatusPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicStatusPage** | [**PublicStatusPage**](PublicStatusPage.md) | The complete definition for the public status page that should be updated. | 

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


## PublicStatusPageRemoveAuthorizationFromPublicStatusPage

> PublicStatusPageRemoveAuthorizationFromPublicStatusPage(ctx, publicStatusPageGuid, authorizationGuid).Execute()

Removes an authorization from a public status page.

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
    publicStatusPageGuid := "publicStatusPageGuid_example" // string | The Guid of the public status page.
    authorizationGuid := "authorizationGuid_example" // string | The Guid of the authorization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicStatusPageApi.PublicStatusPageRemoveAuthorizationFromPublicStatusPage(context.Background(), publicStatusPageGuid, authorizationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicStatusPageApi.PublicStatusPageRemoveAuthorizationFromPublicStatusPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicStatusPageGuid** | **string** | The Guid of the public status page. | 
**authorizationGuid** | **string** | The Guid of the authorization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicStatusPageRemoveAuthorizationFromPublicStatusPageRequest struct via the builder pattern


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

