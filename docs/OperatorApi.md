# \OperatorApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperatorAddDutyPeriodForOperator**](OperatorApi.md#OperatorAddDutyPeriodForOperator) | **Post** /Operator/{operatorGuid}/DutySchedule | Adds a duty schedule to the specified operator.
[**OperatorCreateOperator**](OperatorApi.md#OperatorCreateOperator) | **Post** /Operator | Creates a new operator.
[**OperatorDeleteAuthorizationForOperator**](OperatorApi.md#OperatorDeleteAuthorizationForOperator) | **Delete** /Operator/{operatorGuid}/Authorization/{authorizationType} | Removes the specified authorization of this operator.
[**OperatorDeleteDutyScheduleFromOperator**](OperatorApi.md#OperatorDeleteDutyScheduleFromOperator) | **Delete** /Operator/{operatorGuid}/DutySchedule/{dutyScheduleId} | Deletes the specified duty schedule of the specified operator.
[**OperatorDeleteOperator**](OperatorApi.md#OperatorDeleteOperator) | **Delete** /Operator/{operatorGuid} | Deletes an existing operator.
[**OperatorGetAllOperators**](OperatorApi.md#OperatorGetAllOperators) | **Get** /Operator | Gets a list of all operators.
[**OperatorGetAuthorizationsForOperator**](OperatorApi.md#OperatorGetAuthorizationsForOperator) | **Get** /Operator/{operatorGuid}/Authorization | Gets all authorizations for the specified operator.
[**OperatorGetDutyScheduleForOperator**](OperatorApi.md#OperatorGetDutyScheduleForOperator) | **Get** /Operator/{operatorGuid}/DutySchedule | Gets the duty schedules for an specified operator.
[**OperatorGetOperator**](OperatorApi.md#OperatorGetOperator) | **Get** /Operator/{operatorGuid} | Gets the details of the operator with the provided OperatorGuid.
[**OperatorGetOperatorGroupsForOperator**](OperatorApi.md#OperatorGetOperatorGroupsForOperator) | **Get** /Operator/{operatorGuid}/OperatorGroup | Gets a list of all operator groups for the specified operator.
[**OperatorPostAuthorizationForOperator**](OperatorApi.md#OperatorPostAuthorizationForOperator) | **Post** /Operator/{operatorGuid}/Authorization/{authorizationType} | Assigns the specified authorization to this operator.
[**OperatorUpdateDutyPeriodForOperator**](OperatorApi.md#OperatorUpdateDutyPeriodForOperator) | **Put** /Operator/{operatorGuid}/DutySchedule/{dutyScheduleId} | Updates the specified duty schedule of the specified operator.
[**OperatorUpdateOperator**](OperatorApi.md#OperatorUpdateOperator) | **Put** /Operator/{operatorGuid} | Updates an existing operator.
[**OperatorUpdateOperatorWithPatch**](OperatorApi.md#OperatorUpdateOperatorWithPatch) | **Patch** /Operator/{operatorGuid} | Updates an existing operator.



## OperatorAddDutyPeriodForOperator

> OperatorAddDutyPeriodForOperator(ctx, operatorGuid).Schedule(schedule).Execute()

Adds a duty schedule to the specified operator.

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
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator to add the duty schedule to
    schedule := *openapiclient.NewOperatorDutySchedule(int32(123), map[string][]openapiclient.OperatorScheduleMode{ ... }) // OperatorDutySchedule | The duty schedule to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorAddDutyPeriodForOperator(context.Background(), operatorGuid).Schedule(schedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorAddDutyPeriodForOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** | The Guid of the operator to add the duty schedule to | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorAddDutyPeriodForOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schedule** | [**OperatorDutySchedule**](OperatorDutySchedule.md) | The duty schedule to add | 

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


## OperatorCreateOperator

> Operator OperatorCreateOperator(ctx).UptrendsOperator(uptrendsOperator).Execute()

Creates a new operator.

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
    uptrendsOperator := *openapiclient.NewOperator() // Operator | The details of the operator to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorCreateOperator(context.Background()).UptrendsOperator(uptrendsOperator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorCreateOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorCreateOperator`: Operator
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.OperatorCreateOperator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperatorCreateOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uptrendsOperator** | [**Operator**](Operator.md) | The details of the operator to create | 

### Return type

[**Operator**](Operator.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorDeleteAuthorizationForOperator

> OperatorDeleteAuthorizationForOperator(ctx, operatorGuid, authorizationType).Execute()

Removes the specified authorization of this operator.

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
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator
    authorizationType := "authorizationType_example" // string | The type of authorization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorDeleteAuthorizationForOperator(context.Background(), operatorGuid, authorizationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorDeleteAuthorizationForOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** | The Guid of the operator | 
**authorizationType** | **string** | The type of authorization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorDeleteAuthorizationForOperatorRequest struct via the builder pattern


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


## OperatorDeleteDutyScheduleFromOperator

> OperatorDeleteDutyScheduleFromOperator(ctx, operatorGuid, dutyScheduleId).Execute()

Deletes the specified duty schedule of the specified operator.

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
    operatorGuid := "operatorGuid_example" // string | 
    dutyScheduleId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorDeleteDutyScheduleFromOperator(context.Background(), operatorGuid, dutyScheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorDeleteDutyScheduleFromOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** |  | 
**dutyScheduleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorDeleteDutyScheduleFromOperatorRequest struct via the builder pattern


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


## OperatorDeleteOperator

> OperatorDeleteOperator(ctx, operatorGuid).Execute()

Deletes an existing operator.

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
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorDeleteOperator(context.Background(), operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorDeleteOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** | The Guid of the operator to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorDeleteOperatorRequest struct via the builder pattern


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


## OperatorGetAllOperators

> []Operator OperatorGetAllOperators(ctx).Execute()

Gets a list of all operators.

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
    resp, r, err := apiClient.OperatorApi.OperatorGetAllOperators(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorGetAllOperators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorGetAllOperators`: []Operator
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.OperatorGetAllOperators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGetAllOperatorsRequest struct via the builder pattern


### Return type

[**[]Operator**](Operator.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorGetAuthorizationsForOperator

> []OperatorAuthorizationType OperatorGetAuthorizationsForOperator(ctx, operatorGuid).Execute()

Gets all authorizations for the specified operator.

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
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorGetAuthorizationsForOperator(context.Background(), operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorGetAuthorizationsForOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorGetAuthorizationsForOperator`: []OperatorAuthorizationType
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.OperatorGetAuthorizationsForOperator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** | The Guid of the operator | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGetAuthorizationsForOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OperatorAuthorizationType**](OperatorAuthorizationType.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorGetDutyScheduleForOperator

> []OperatorDutySchedule OperatorGetDutyScheduleForOperator(ctx, operatorGuid).Execute()

Gets the duty schedules for an specified operator.

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
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator to get the duty schedule for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorGetDutyScheduleForOperator(context.Background(), operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorGetDutyScheduleForOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorGetDutyScheduleForOperator`: []OperatorDutySchedule
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.OperatorGetDutyScheduleForOperator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** | The Guid of the operator to get the duty schedule for | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGetDutyScheduleForOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OperatorDutySchedule**](OperatorDutySchedule.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorGetOperator

> Operator OperatorGetOperator(ctx, operatorGuid).Execute()

Gets the details of the operator with the provided OperatorGuid.

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
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator for which to retrieve the details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorGetOperator(context.Background(), operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorGetOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorGetOperator`: Operator
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.OperatorGetOperator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** | The Guid of the operator for which to retrieve the details | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGetOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Operator**](Operator.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorGetOperatorGroupsForOperator

> []OperatorMember OperatorGetOperatorGroupsForOperator(ctx, operatorGuid).Execute()

Gets a list of all operator groups for the specified operator.

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
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator for which to retrieve the operator group guids

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorGetOperatorGroupsForOperator(context.Background(), operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorGetOperatorGroupsForOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorGetOperatorGroupsForOperator`: []OperatorMember
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.OperatorGetOperatorGroupsForOperator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** | The Guid of the operator for which to retrieve the operator group guids | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorGetOperatorGroupsForOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OperatorMember**](OperatorMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorPostAuthorizationForOperator

> OperatorPostAuthorizationForOperator(ctx, operatorGuid, authorizationType).Execute()

Assigns the specified authorization to this operator.

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
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator
    authorizationType := "authorizationType_example" // string | The type of authorization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorPostAuthorizationForOperator(context.Background(), operatorGuid, authorizationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorPostAuthorizationForOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** | The Guid of the operator | 
**authorizationType** | **string** | The type of authorization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorPostAuthorizationForOperatorRequest struct via the builder pattern


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


## OperatorUpdateDutyPeriodForOperator

> OperatorUpdateDutyPeriodForOperator(ctx, operatorGuid, dutyScheduleId).Schedule(schedule).Execute()

Updates the specified duty schedule of the specified operator.

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
    operatorGuid := "operatorGuid_example" // string | 
    dutyScheduleId := int32(56) // int32 | 
    schedule := *openapiclient.NewOperatorDutySchedule(int32(123), map[string][]openapiclient.OperatorScheduleMode{ ... }) // OperatorDutySchedule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorUpdateDutyPeriodForOperator(context.Background(), operatorGuid, dutyScheduleId).Schedule(schedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorUpdateDutyPeriodForOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** |  | 
**dutyScheduleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorUpdateDutyPeriodForOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **schedule** | [**OperatorDutySchedule**](OperatorDutySchedule.md) |  | 

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


## OperatorUpdateOperator

> OperatorUpdateOperator(ctx, operatorGuid).UptrendsOperator(uptrendsOperator).Execute()

Updates an existing operator.

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
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator to update
    uptrendsOperator := *openapiclient.NewOperator() // Operator | The updated details of the operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorUpdateOperator(context.Background(), operatorGuid).UptrendsOperator(uptrendsOperator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorUpdateOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** | The Guid of the operator to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorUpdateOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uptrendsOperator** | [**Operator**](Operator.md) | The updated details of the operator | 

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


## OperatorUpdateOperatorWithPatch

> OperatorUpdateOperatorWithPatch(ctx, operatorGuid).UptrendsOperator(uptrendsOperator).Execute()

Updates an existing operator.

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
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator to update
    uptrendsOperator := *openapiclient.NewOperator() // Operator | The updated details of the operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.OperatorUpdateOperatorWithPatch(context.Background(), operatorGuid).UptrendsOperator(uptrendsOperator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.OperatorUpdateOperatorWithPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorGuid** | **string** | The Guid of the operator to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorUpdateOperatorWithPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uptrendsOperator** | [**Operator**](Operator.md) | The updated details of the operator | 

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

