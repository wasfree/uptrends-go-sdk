# \VaultApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VaultCreateAuthorizationForVaultSection**](VaultApi.md#VaultCreateAuthorizationForVaultSection) | **Post** /VaultSection/{vaultSectionGuid}/Authorization | Creates a new authorization for the specified vault section.
[**VaultCreateNewVaultItem**](VaultApi.md#VaultCreateNewVaultItem) | **Post** /VaultItem | Creates a new vault item.
[**VaultCreateNewVaultSection**](VaultApi.md#VaultCreateNewVaultSection) | **Post** /VaultSection | Creates a new vault section.
[**VaultDeleteAuthorizationForVaultSection**](VaultApi.md#VaultDeleteAuthorizationForVaultSection) | **Delete** /VaultSection/{vaultSectionGuid}/Authorization/{authorizationGuid} | Deletes the specified authorization for the specified vault section.
[**VaultDeleteVaultItem**](VaultApi.md#VaultDeleteVaultItem) | **Delete** /VaultItem/{vaultItemGuid} | Deletes the specified vault item.
[**VaultDeleteVaultSection**](VaultApi.md#VaultDeleteVaultSection) | **Delete** /VaultSection/{vaultSectionGuid} | Deletes the specified vault section.
[**VaultGetAllVaultItems**](VaultApi.md#VaultGetAllVaultItems) | **Get** /VaultItem | Returns all vault items.
[**VaultGetAllVaultSections**](VaultApi.md#VaultGetAllVaultSections) | **Get** /VaultSection | Returns all vault sections.
[**VaultGetAuthorizationsForVaultSection**](VaultApi.md#VaultGetAuthorizationsForVaultSection) | **Get** /VaultSection/{vaultSectionGuid}/Authorization | Returns all authorizations for the specified vault section.
[**VaultGetVaultItem**](VaultApi.md#VaultGetVaultItem) | **Get** /VaultItem/{vaultItemGuid} | Returns the specified vault item.
[**VaultGetVaultSection**](VaultApi.md#VaultGetVaultSection) | **Get** /VaultSection/{vaultSectionGuid} | Returns the specified vault section.
[**VaultPartiallyUpdateVaultItem**](VaultApi.md#VaultPartiallyUpdateVaultItem) | **Patch** /VaultItem/{vaultItemGuid} | Partially updates the specified vault item.
[**VaultUpdateVaultItem**](VaultApi.md#VaultUpdateVaultItem) | **Put** /VaultItem/{vaultItemGuid} | Updates the specified vault item.
[**VaultUpdateVaultSection**](VaultApi.md#VaultUpdateVaultSection) | **Put** /VaultSection/{vaultSectionGuid} | Updates the specified vault section.



## VaultCreateAuthorizationForVaultSection

> VaultSectionAuthorization VaultCreateAuthorizationForVaultSection(ctx, vaultSectionGuid).Authorization(authorization).Execute()

Creates a new authorization for the specified vault section.



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
    vaultSectionGuid := "vaultSectionGuid_example" // string | The Guid of the vault section for which to create the new authorization.
    authorization := *openapiclient.NewVaultSectionAuthorization(map[string][]openapiclient.VaultSectionAuthorizationType{ ... }) // VaultSectionAuthorization |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultCreateAuthorizationForVaultSection(context.Background(), vaultSectionGuid).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultCreateAuthorizationForVaultSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultCreateAuthorizationForVaultSection`: VaultSectionAuthorization
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.VaultCreateAuthorizationForVaultSection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultSectionGuid** | **string** | The Guid of the vault section for which to create the new authorization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVaultCreateAuthorizationForVaultSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | [**VaultSectionAuthorization**](VaultSectionAuthorization.md) |  | 

### Return type

[**VaultSectionAuthorization**](VaultSectionAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultCreateNewVaultItem

> VaultItem VaultCreateNewVaultItem(ctx).Item(item).Execute()

Creates a new vault item.



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
    item := *openapiclient.NewVaultItem() // VaultItem | The item to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultCreateNewVaultItem(context.Background()).Item(item).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultCreateNewVaultItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultCreateNewVaultItem`: VaultItem
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.VaultCreateNewVaultItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVaultCreateNewVaultItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **item** | [**VaultItem**](VaultItem.md) | The item to create | 

### Return type

[**VaultItem**](VaultItem.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultCreateNewVaultSection

> VaultSection VaultCreateNewVaultSection(ctx).Section(section).Execute()

Creates a new vault section.



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
    section := *openapiclient.NewVaultSection("VaultSectionGuid_example") // VaultSection | The attributes of the vault section that should be created. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultCreateNewVaultSection(context.Background()).Section(section).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultCreateNewVaultSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultCreateNewVaultSection`: VaultSection
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.VaultCreateNewVaultSection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVaultCreateNewVaultSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **section** | [**VaultSection**](VaultSection.md) | The attributes of the vault section that should be created. | 

### Return type

[**VaultSection**](VaultSection.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultDeleteAuthorizationForVaultSection

> VaultDeleteAuthorizationForVaultSection(ctx, vaultSectionGuid, authorizationGuid).Execute()

Deletes the specified authorization for the specified vault section.

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
    vaultSectionGuid := "vaultSectionGuid_example" // string | The Guid of the vault section for which the authorization should be deleted.
    authorizationGuid := "authorizationGuid_example" // string | The Guid of the authorization that should be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultDeleteAuthorizationForVaultSection(context.Background(), vaultSectionGuid, authorizationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultDeleteAuthorizationForVaultSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultSectionGuid** | **string** | The Guid of the vault section for which the authorization should be deleted. | 
**authorizationGuid** | **string** | The Guid of the authorization that should be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVaultDeleteAuthorizationForVaultSectionRequest struct via the builder pattern


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


## VaultDeleteVaultItem

> VaultDeleteVaultItem(ctx, vaultItemGuid).Execute()

Deletes the specified vault item.

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
    vaultItemGuid := "vaultItemGuid_example" // string | The Guid of the vault item that should be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultDeleteVaultItem(context.Background(), vaultItemGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultDeleteVaultItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultItemGuid** | **string** | The Guid of the vault item that should be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVaultDeleteVaultItemRequest struct via the builder pattern


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


## VaultDeleteVaultSection

> VaultDeleteVaultSection(ctx, vaultSectionGuid).Execute()

Deletes the specified vault section.

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
    vaultSectionGuid := "vaultSectionGuid_example" // string | The Guid of the vault section that should be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultDeleteVaultSection(context.Background(), vaultSectionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultDeleteVaultSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultSectionGuid** | **string** | The Guid of the vault section that should be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVaultDeleteVaultSectionRequest struct via the builder pattern


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


## VaultGetAllVaultItems

> []VaultItem VaultGetAllVaultItems(ctx).Execute()

Returns all vault items.

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
    resp, r, err := apiClient.VaultApi.VaultGetAllVaultItems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultGetAllVaultItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultGetAllVaultItems`: []VaultItem
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.VaultGetAllVaultItems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVaultGetAllVaultItemsRequest struct via the builder pattern


### Return type

[**[]VaultItem**](VaultItem.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultGetAllVaultSections

> []VaultSection VaultGetAllVaultSections(ctx).Execute()

Returns all vault sections.

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
    resp, r, err := apiClient.VaultApi.VaultGetAllVaultSections(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultGetAllVaultSections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultGetAllVaultSections`: []VaultSection
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.VaultGetAllVaultSections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVaultGetAllVaultSectionsRequest struct via the builder pattern


### Return type

[**[]VaultSection**](VaultSection.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultGetAuthorizationsForVaultSection

> VaultSectionAuthorization VaultGetAuthorizationsForVaultSection(ctx, vaultSectionGuid).Execute()

Returns all authorizations for the specified vault section.

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
    vaultSectionGuid := "vaultSectionGuid_example" // string | The Guid of the vault section for which to return authorizations.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultGetAuthorizationsForVaultSection(context.Background(), vaultSectionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultGetAuthorizationsForVaultSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultGetAuthorizationsForVaultSection`: VaultSectionAuthorization
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.VaultGetAuthorizationsForVaultSection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultSectionGuid** | **string** | The Guid of the vault section for which to return authorizations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVaultGetAuthorizationsForVaultSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VaultSectionAuthorization**](VaultSectionAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultGetVaultItem

> VaultItem VaultGetVaultItem(ctx, vaultItemGuid).Execute()

Returns the specified vault item.

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
    vaultItemGuid := "vaultItemGuid_example" // string | The Guid of the requested vault item.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultGetVaultItem(context.Background(), vaultItemGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultGetVaultItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultGetVaultItem`: VaultItem
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.VaultGetVaultItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultItemGuid** | **string** | The Guid of the requested vault item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVaultGetVaultItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VaultItem**](VaultItem.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultGetVaultSection

> VaultSection VaultGetVaultSection(ctx, vaultSectionGuid).Execute()

Returns the specified vault section.

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
    vaultSectionGuid := "vaultSectionGuid_example" // string | The Guid of the requested vault section.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultGetVaultSection(context.Background(), vaultSectionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultGetVaultSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VaultGetVaultSection`: VaultSection
    fmt.Fprintf(os.Stdout, "Response from `VaultApi.VaultGetVaultSection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultSectionGuid** | **string** | The Guid of the requested vault section. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVaultGetVaultSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VaultSection**](VaultSection.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultPartiallyUpdateVaultItem

> VaultPartiallyUpdateVaultItem(ctx, vaultItemGuid).Item(item).Execute()

Partially updates the specified vault item.



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
    vaultItemGuid := "vaultItemGuid_example" // string | The Guid of the vault item that should be updated.
    item := *openapiclient.NewVaultItem() // VaultItem | Part of the definition of the vault item that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultPartiallyUpdateVaultItem(context.Background(), vaultItemGuid).Item(item).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultPartiallyUpdateVaultItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultItemGuid** | **string** | The Guid of the vault item that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVaultPartiallyUpdateVaultItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **item** | [**VaultItem**](VaultItem.md) | Part of the definition of the vault item that should be updated. | 

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


## VaultUpdateVaultItem

> VaultUpdateVaultItem(ctx, vaultItemGuid).Item(item).Execute()

Updates the specified vault item.



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
    vaultItemGuid := "vaultItemGuid_example" // string | The Guid of the vault item that should be updated.
    item := *openapiclient.NewVaultItem() // VaultItem | The complete definition of the vault item that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultUpdateVaultItem(context.Background(), vaultItemGuid).Item(item).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultUpdateVaultItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultItemGuid** | **string** | The Guid of the vault item that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVaultUpdateVaultItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **item** | [**VaultItem**](VaultItem.md) | The complete definition of the vault item that should be updated. | 

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


## VaultUpdateVaultSection

> VaultUpdateVaultSection(ctx, vaultSectionGuid).Item(item).Execute()

Updates the specified vault section.

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
    vaultSectionGuid := "vaultSectionGuid_example" // string | The Guid of the vault section that should be updated.
    item := *openapiclient.NewVaultSection("VaultSectionGuid_example") // VaultSection |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultApi.VaultUpdateVaultSection(context.Background(), vaultSectionGuid).Item(item).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.VaultUpdateVaultSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultSectionGuid** | **string** | The Guid of the vault section that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVaultUpdateVaultSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **item** | [**VaultSection**](VaultSection.md) |  | 

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

