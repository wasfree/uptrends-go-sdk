# \OperatorGroupApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperatorGroupAddDutyScheduleToAllMembers**](OperatorGroupApi.md#OperatorGroupAddDutyScheduleToAllMembers) | **Post** /OperatorGroup/{operatorGroupGuid}/DutySchedule/AddDutyScheduleForAllMembers | Adds the provided duty schedule to all operators in the group specified
[**OperatorGroupAddOperatorToOperatorGroup**](OperatorGroupApi.md#OperatorGroupAddOperatorToOperatorGroup) | **Post** /OperatorGroup/{operatorGroupGuid}/Member/{operatorGuid} | Adds the specified operator to the operator group
[**OperatorGroupAllOperatorsInGroupOffDuty**](OperatorGroupApi.md#OperatorGroupAllOperatorsInGroupOffDuty) | **Post** /OperatorGroup/{operatorGroupGuid}/AllOperatorsOffDuty | Set the OnDuty flag to off for all operators that are a member of the operator group specified by the operator group GUID
[**OperatorGroupAllOperatorsInGroupOnDuty**](OperatorGroupApi.md#OperatorGroupAllOperatorsInGroupOnDuty) | **Post** /OperatorGroup/{operatorGroupGuid}/AllOperatorsOnDuty | Set the OnDuty flag to on for all operators that are a member of the operator group specified by the operator group GUID
[**OperatorGroupCreateOperatorGroup**](OperatorGroupApi.md#OperatorGroupCreateOperatorGroup) | **Post** /OperatorGroup | Creates a new operator group
[**OperatorGroupDeleteAuthorizationForOperatorGroup**](OperatorGroupApi.md#OperatorGroupDeleteAuthorizationForOperatorGroup) | **Delete** /OperatorGroup/{operatorGroupGuid}/Authorization/{authorizationType} | Removes the specified authorization of the operator group.
[**OperatorGroupDeleteOperatorGroup**](OperatorGroupApi.md#OperatorGroupDeleteOperatorGroup) | **Delete** /OperatorGroup/{operatorGroupGuid} | Deletes the specified operator group
[**OperatorGroupGetAllOperatorGroups**](OperatorGroupApi.md#OperatorGroupGetAllOperatorGroups) | **Get** /OperatorGroup | Gets all operator groups
[**OperatorGroupGetAuthorizationsForOperatorGroup**](OperatorGroupApi.md#OperatorGroupGetAuthorizationsForOperatorGroup) | **Get** /OperatorGroup/{operatorGroupGuid}/Authorization | Gets all authorizations for the specified operator group.
[**OperatorGroupGetOperatorGroup**](OperatorGroupApi.md#OperatorGroupGetOperatorGroup) | **Get** /OperatorGroup/{operatorGroupGuid} | Gets the details of a operator group
[**OperatorGroupGetOperatorGroupMembers**](OperatorGroupApi.md#OperatorGroupGetOperatorGroupMembers) | **Get** /OperatorGroup/{operatorGroupGuid}/Member | Gets a list of all members of an operator group
[**OperatorGroupPostAuthorizationForOperatorGroup**](OperatorGroupApi.md#OperatorGroupPostAuthorizationForOperatorGroup) | **Post** /OperatorGroup/{operatorGroupGuid}/Authorization/{authorizationType} | Assigns the specified authorization to the operator group.
[**OperatorGroupRemoveOperatorFromOperatorGroup**](OperatorGroupApi.md#OperatorGroupRemoveOperatorFromOperatorGroup) | **Delete** /OperatorGroup/{operatorGroupGuid}/Member/{operatorGuid} | Removes the specified operator from the operator group
[**OperatorGroupUpdateOperatorGroup**](OperatorGroupApi.md#OperatorGroupUpdateOperatorGroup) | **Put** /OperatorGroup/{operatorGroupGuid} | Updates the operator group with the Guid specified



## OperatorGroupAddDutyScheduleToAllMembers

> OperatorGroupAddDutyScheduleToAllMembers(ctx, operatorGroupGuid).DutySchedule(dutySchedule).Execute()

Adds the provided duty schedule to all operators in the group specified

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | 
    dutySchedule := *openapiclient.NewOperatorDutySchedule(int32(123), map[string][]openapiclient.OperatorScheduleMode{ ... }) // OperatorDutySchedule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupAddDutyScheduleToAllMembers(context.Background(), operatorGroupGuid).DutySchedule(dutySchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupAddDutyScheduleToAllMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupAddDutyScheduleToAllMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dutySchedule** | [**OperatorDutySchedule**](OperatorDutySchedule.md) |  | 

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


## OperatorGroupAddOperatorToOperatorGroup

> OperatorGroupAddOperatorToOperatorGroup(ctx, operatorGroupGuid, operatorGuid).Execute()

Adds the specified operator to the operator group

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group to add the operator to
    operatorGuid := "operatorGuid_example" // string | The operator Guid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupAddOperatorToOperatorGroup(context.Background(), operatorGroupGuid, operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupAddOperatorToOperatorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** | The Guid of the operator group to add the operator to | 
**operatorGuid** | **string** | The operator Guid | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupAddOperatorToOperatorGroupRequest struct via the builder pattern


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


## OperatorGroupAllOperatorsInGroupOffDuty

> OperatorGroupAllOperatorsInGroupOffDuty(ctx, operatorGroupGuid).Execute()

Set the OnDuty flag to off for all operators that are a member of the operator group specified by the operator group GUID

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | The operator group GUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupAllOperatorsInGroupOffDuty(context.Background(), operatorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupAllOperatorsInGroupOffDuty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** | The operator group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupAllOperatorsInGroupOffDutyRequest struct via the builder pattern


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


## OperatorGroupAllOperatorsInGroupOnDuty

> OperatorGroupAllOperatorsInGroupOnDuty(ctx, operatorGroupGuid).Execute()

Set the OnDuty flag to on for all operators that are a member of the operator group specified by the operator group GUID

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | The operator group GUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupAllOperatorsInGroupOnDuty(context.Background(), operatorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupAllOperatorsInGroupOnDuty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** | The operator group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupAllOperatorsInGroupOnDutyRequest struct via the builder pattern


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


## OperatorGroupCreateOperatorGroup

> OperatorGroup OperatorGroupCreateOperatorGroup(ctx).OperatorGroup(operatorGroup).Execute()

Creates a new operator group

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
    operatorGroup := *openapiclient.NewOperatorGroup() // OperatorGroup | The operatorGroup object to be created (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupCreateOperatorGroup(context.Background()).OperatorGroup(operatorGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupCreateOperatorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorGroupCreateOperatorGroup`: OperatorGroup
    fmt.Fprintf(os.Stdout, "Response from `OperatorGroupApi.OperatorGroupCreateOperatorGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupCreateOperatorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operatorGroup** | [**OperatorGroup**](OperatorGroup.md) | The operatorGroup object to be created | 

### Return type

[**OperatorGroup**](OperatorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorGroupDeleteAuthorizationForOperatorGroup

> OperatorGroupDeleteAuthorizationForOperatorGroup(ctx, operatorGroupGuid, authorizationType).Execute()

Removes the specified authorization of the operator group.

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group
    authorizationType := "authorizationType_example" // string | The type of authorization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupDeleteAuthorizationForOperatorGroup(context.Background(), operatorGroupGuid, authorizationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupDeleteAuthorizationForOperatorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** | The Guid of the operator group | 
**authorizationType** | **string** | The type of authorization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupDeleteAuthorizationForOperatorGroupRequest struct via the builder pattern


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


## OperatorGroupDeleteOperatorGroup

> OperatorGroupDeleteOperatorGroup(ctx, operatorGroupGuid).Execute()

Deletes the specified operator group

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupDeleteOperatorGroup(context.Background(), operatorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupDeleteOperatorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** | The Guid of the operator group to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupDeleteOperatorGroupRequest struct via the builder pattern


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


## OperatorGroupGetAllOperatorGroups

> []OperatorGroup OperatorGroupGetAllOperatorGroups(ctx).Execute()

Gets all operator groups

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
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupGetAllOperatorGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupGetAllOperatorGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorGroupGetAllOperatorGroups`: []OperatorGroup
    fmt.Fprintf(os.Stdout, "Response from `OperatorGroupApi.OperatorGroupGetAllOperatorGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupGetAllOperatorGroupsRequest struct via the builder pattern


### Return type

[**[]OperatorGroup**](OperatorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorGroupGetAuthorizationsForOperatorGroup

> []OperatorGroupAuthorizationType OperatorGroupGetAuthorizationsForOperatorGroup(ctx, operatorGroupGuid).Execute()

Gets all authorizations for the specified operator group.

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupGetAuthorizationsForOperatorGroup(context.Background(), operatorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupGetAuthorizationsForOperatorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorGroupGetAuthorizationsForOperatorGroup`: []OperatorGroupAuthorizationType
    fmt.Fprintf(os.Stdout, "Response from `OperatorGroupApi.OperatorGroupGetAuthorizationsForOperatorGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** | The Guid of the operator group | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupGetAuthorizationsForOperatorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OperatorGroupAuthorizationType**](OperatorGroupAuthorizationType.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorGroupGetOperatorGroup

> OperatorGroup OperatorGroupGetOperatorGroup(ctx, operatorGroupGuid).Execute()

Gets the details of a operator group

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupGetOperatorGroup(context.Background(), operatorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupGetOperatorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorGroupGetOperatorGroup`: OperatorGroup
    fmt.Fprintf(os.Stdout, "Response from `OperatorGroupApi.OperatorGroupGetOperatorGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** | The Guid of the operator group to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupGetOperatorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperatorGroup**](OperatorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorGroupGetOperatorGroupMembers

> OperatorGroupMember OperatorGroupGetOperatorGroupMembers(ctx, operatorGroupGuid).Execute()

Gets a list of all members of an operator group

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group to retrieve the members for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupGetOperatorGroupMembers(context.Background(), operatorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupGetOperatorGroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorGroupGetOperatorGroupMembers`: OperatorGroupMember
    fmt.Fprintf(os.Stdout, "Response from `OperatorGroupApi.OperatorGroupGetOperatorGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** | The Guid of the operator group to retrieve the members for | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupGetOperatorGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperatorGroupMember**](OperatorGroupMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorGroupPostAuthorizationForOperatorGroup

> OperatorGroupPostAuthorizationForOperatorGroup(ctx, operatorGroupGuid, authorizationType).Execute()

Assigns the specified authorization to the operator group.

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group
    authorizationType := "authorizationType_example" // string | The type of authorization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupPostAuthorizationForOperatorGroup(context.Background(), operatorGroupGuid, authorizationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupPostAuthorizationForOperatorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** | The Guid of the operator group | 
**authorizationType** | **string** | The type of authorization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupPostAuthorizationForOperatorGroupRequest struct via the builder pattern


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


## OperatorGroupRemoveOperatorFromOperatorGroup

> OperatorGroupRemoveOperatorFromOperatorGroup(ctx, operatorGroupGuid, operatorGuid).Execute()

Removes the specified operator from the operator group

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group to remove the operator from
    operatorGuid := "operatorGuid_example" // string | The operator Guid to be removed

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupRemoveOperatorFromOperatorGroup(context.Background(), operatorGroupGuid, operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupRemoveOperatorFromOperatorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** | The Guid of the operator group to remove the operator from | 
**operatorGuid** | **string** | The operator Guid to be removed | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupRemoveOperatorFromOperatorGroupRequest struct via the builder pattern


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


## OperatorGroupUpdateOperatorGroup

> OperatorGroupUpdateOperatorGroup(ctx, operatorGroupGuid).OperatorGroup(operatorGroup).Execute()

Updates the operator group with the Guid specified

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
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group to be updated
    operatorGroup := *openapiclient.NewOperatorGroup() // OperatorGroup | The operator group to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorGroupApi.OperatorGroupUpdateOperatorGroup(context.Background(), operatorGroupGuid).OperatorGroup(operatorGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorGroupApi.OperatorGroupUpdateOperatorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGroupGuid** | **string** | The Guid of the operator group to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGroupUpdateOperatorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operatorGroup** | [**OperatorGroup**](OperatorGroup.md) | The operator group to be updated | 

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

