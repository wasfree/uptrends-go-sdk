# \AlertDefinitionApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertDefinitionAddIntegrationToEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionAddIntegrationToEscalationLevel) | **Post** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration | Adds an integration to a specified escalation level.
[**AlertDefinitionAddMonitorGroupToAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionAddMonitorGroupToAlertDefinition) | **Post** /AlertDefinition/{alertDefinitionGuid}/Member/MonitorGroup/{monitorGroupGuid} | Adds a monitor group to the specified alert definition.
[**AlertDefinitionAddMonitorToAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionAddMonitorToAlertDefinition) | **Post** /AlertDefinition/{alertDefinitionGuid}/Member/Monitor/{monitorGuid} | Adds a monitor to the specified alert definition.
[**AlertDefinitionAddOperatorGroupToEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionAddOperatorGroupToEscalationLevel) | **Post** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/OperatorGroup/{operatorGroupGuid} | Adds an operator group to the specified escalation level.
[**AlertDefinitionAddOperatorToEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionAddOperatorToEscalationLevel) | **Post** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/Operator/{operatorGuid} | Adds an operator to the specified escalation level.
[**AlertDefinitionCreateAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionCreateAlertDefinition) | **Post** /AlertDefinition | Creates a new alert definition.
[**AlertDefinitionDeleteAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionDeleteAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid} | Deletes an existing alert definition.
[**AlertDefinitionGetAllAlertDefinitions**](AlertDefinitionApi.md#AlertDefinitionGetAllAlertDefinitions) | **Get** /AlertDefinition | Gets a list of all alert definitions.
[**AlertDefinitionGetAllEscalationLevelIntegrations**](AlertDefinitionApi.md#AlertDefinitionGetAllEscalationLevelIntegrations) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration | Gets all integrations for a specified escalation level.
[**AlertDefinitionGetAllEscalationLevels**](AlertDefinitionApi.md#AlertDefinitionGetAllEscalationLevels) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel | Gets all escalation level information for the specified alert definition.
[**AlertDefinitionGetAllMembers**](AlertDefinitionApi.md#AlertDefinitionGetAllMembers) | **Get** /AlertDefinition/{alertDefinitionGuid}/Member | Gets a list of all monitor and monitor group guids for the specified alert definition.
[**AlertDefinitionGetEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionGetEscalationLevel) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId} | Gets the escalation level information of the specified alert definition.
[**AlertDefinitionGetEscalationLevelIntegration**](AlertDefinitionApi.md#AlertDefinitionGetEscalationLevelIntegration) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Gets a single integration for a specified escalation level.
[**AlertDefinitionGetEscalationLevelOperator**](AlertDefinitionApi.md#AlertDefinitionGetEscalationLevelOperator) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member | Gets the operator and operator group guids for the specified escalation level.
[**AlertDefinitionGetSpecifiedAlertDefinitions**](AlertDefinitionApi.md#AlertDefinitionGetSpecifiedAlertDefinitions) | **Get** /AlertDefinition/{alertDefinitionGuid} | Gets the specified alert definition.
[**AlertDefinitionPatchAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionPatchAlertDefinition) | **Patch** /AlertDefinition/{alertDefinitionGuid} | Partially updates the definition for the specified alert definition.
[**AlertDefinitionPatchAlertDefinitionEscalation**](AlertDefinitionApi.md#AlertDefinitionPatchAlertDefinitionEscalation) | **Patch** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId} | Partially updates the escalation level for the specified alert definition.
[**AlertDefinitionPutAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionPutAlertDefinition) | **Put** /AlertDefinition/{alertDefinitionGuid} | Updates the definition for the specified alert definition.
[**AlertDefinitionPutAlertDefinitionEscalation**](AlertDefinitionApi.md#AlertDefinitionPutAlertDefinitionEscalation) | **Put** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId} | Updates the escalation level for the specified alert definition.
[**AlertDefinitionRemoveIntegrationFromEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionRemoveIntegrationFromEscalationLevel) | **Delete** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Removes an integration from a specified escalation level.
[**AlertDefinitionRemoveMonitorFromAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionRemoveMonitorFromAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid}/Member/Monitor/{monitorGuid} | Removes a monitor for the specified alert definition.
[**AlertDefinitionRemoveMonitorGroupFromAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionRemoveMonitorGroupFromAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid}/Member/MonitorGroup/{monitorGroupGuid} | Removes a monitor group for the specified alert definition.
[**AlertDefinitionRemoveOperatorFromEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionRemoveOperatorFromEscalationLevel) | **Delete** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/Operator/{operatorGuid} | Removes an operator for the specified escalation level.
[**AlertDefinitionRemoveOperatorGroupFromEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionRemoveOperatorGroupFromEscalationLevel) | **Delete** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/OperatorGroup/{operatorGroupGuid} | Removes an operator group for the specified escalation level.
[**AlertDefinitionUpdateIntegrationForEscalationWithPatch**](AlertDefinitionApi.md#AlertDefinitionUpdateIntegrationForEscalationWithPatch) | **Patch** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Partially updates an integration for a specified escalation level.
[**AlertDefinitionUpdateIntegrationForEscalationWithPut**](AlertDefinitionApi.md#AlertDefinitionUpdateIntegrationForEscalationWithPut) | **Put** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Updates an integration for a specified escalation level.



## AlertDefinitionAddIntegrationToEscalationLevel

> AlertDefinitionMonitorGroup AlertDefinitionAddIntegrationToEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId).EscalationLevelIntegration(escalationLevelIntegration).Execute()

Adds an integration to a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    escalationLevelIntegration := *openapiclient.NewEscalationLevelIntegration() // EscalationLevelIntegration | The integration to add.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionAddIntegrationToEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId).EscalationLevelIntegration(escalationLevelIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionAddIntegrationToEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionAddIntegrationToEscalationLevel`: AlertDefinitionMonitorGroup
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionAddIntegrationToEscalationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionAddIntegrationToEscalationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **escalationLevelIntegration** | [**EscalationLevelIntegration**](EscalationLevelIntegration.md) | The integration to add. | 

### Return type

[**AlertDefinitionMonitorGroup**](AlertDefinitionMonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionAddMonitorGroupToAlertDefinition

> AlertDefinitionMonitorGroup AlertDefinitionAddMonitorGroupToAlertDefinition(ctx, alertDefinitionGuid, monitorGroupGuid).Execute()

Adds a monitor group to the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition to modify.
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to add.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionAddMonitorGroupToAlertDefinition(context.Background(), alertDefinitionGuid, monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionAddMonitorGroupToAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionAddMonitorGroupToAlertDefinition`: AlertDefinitionMonitorGroup
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionAddMonitorGroupToAlertDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition to modify. | 
**monitorGroupGuid** | **string** | The Guid of the monitor group to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionAddMonitorGroupToAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertDefinitionMonitorGroup**](AlertDefinitionMonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionAddMonitorToAlertDefinition

> AlertDefinitionMonitor AlertDefinitionAddMonitorToAlertDefinition(ctx, alertDefinitionGuid, monitorGuid).Execute()

Adds a monitor to the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition to modify.
    monitorGuid := "monitorGuid_example" // string | The Guid of the monitor to add.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionAddMonitorToAlertDefinition(context.Background(), alertDefinitionGuid, monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionAddMonitorToAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionAddMonitorToAlertDefinition`: AlertDefinitionMonitor
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionAddMonitorToAlertDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition to modify. | 
**monitorGuid** | **string** | The Guid of the monitor to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionAddMonitorToAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertDefinitionMonitor**](AlertDefinitionMonitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionAddOperatorGroupToEscalationLevel

> AlertDefinitionOperatorGroup AlertDefinitionAddOperatorGroupToEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGroupGuid).Execute()

Adds an operator group to the specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group to add.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionAddOperatorGroupToEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId, operatorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionAddOperatorGroupToEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionAddOperatorGroupToEscalationLevel`: AlertDefinitionOperatorGroup
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionAddOperatorGroupToEscalationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**operatorGroupGuid** | **string** | The Guid of the operator group to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionAddOperatorGroupToEscalationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AlertDefinitionOperatorGroup**](AlertDefinitionOperatorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionAddOperatorToEscalationLevel

> AlertDefinitionOperator AlertDefinitionAddOperatorToEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGuid).Execute()

Adds an operator to the specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator to add.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionAddOperatorToEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId, operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionAddOperatorToEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionAddOperatorToEscalationLevel`: AlertDefinitionOperator
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionAddOperatorToEscalationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**operatorGuid** | **string** | The Guid of the operator to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionAddOperatorToEscalationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AlertDefinitionOperator**](AlertDefinitionOperator.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionCreateAlertDefinition

> AlertDefinition AlertDefinitionCreateAlertDefinition(ctx).AlertDefinition(alertDefinition).Execute()

Creates a new alert definition.

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
    alertDefinition := *openapiclient.NewAlertDefinition() // AlertDefinition | The details of the alert definition to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionCreateAlertDefinition(context.Background()).AlertDefinition(alertDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionCreateAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionCreateAlertDefinition`: AlertDefinition
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionCreateAlertDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionCreateAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertDefinition** | [**AlertDefinition**](AlertDefinition.md) | The details of the alert definition to create. | 

### Return type

[**AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionDeleteAlertDefinition

> AlertDefinitionDeleteAlertDefinition(ctx, alertDefinitionGuid).Execute()

Deletes an existing alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition to remove.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionDeleteAlertDefinition(context.Background(), alertDefinitionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionDeleteAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionDeleteAlertDefinitionRequest struct via the builder pattern


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


## AlertDefinitionGetAllAlertDefinitions

> []AlertDefinition AlertDefinitionGetAllAlertDefinitions(ctx).Execute()

Gets a list of all alert definitions.

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
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionGetAllAlertDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionGetAllAlertDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetAllAlertDefinitions`: []AlertDefinition
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionGetAllAlertDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetAllAlertDefinitionsRequest struct via the builder pattern


### Return type

[**[]AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetAllEscalationLevelIntegrations

> []Integration AlertDefinitionGetAllEscalationLevelIntegrations(ctx, alertDefinitionGuid, escalationLevelId).Execute()

Gets all integrations for a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionGetAllEscalationLevelIntegrations(context.Background(), alertDefinitionGuid, escalationLevelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionGetAllEscalationLevelIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetAllEscalationLevelIntegrations`: []Integration
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionGetAllEscalationLevelIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetAllEscalationLevelIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AlertDefinitionGetAllEscalationLevels

> []EscalationLevel AlertDefinitionGetAllEscalationLevels(ctx, alertDefinitionGuid).Execute()

Gets all escalation level information for the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition for which to return all escalation levels.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionGetAllEscalationLevels(context.Background(), alertDefinitionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionGetAllEscalationLevels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetAllEscalationLevels`: []EscalationLevel
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionGetAllEscalationLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition for which to return all escalation levels. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetAllEscalationLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EscalationLevel**](EscalationLevel.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetAllMembers

> []AlertDefinitionMember AlertDefinitionGetAllMembers(ctx, alertDefinitionGuid).Execute()

Gets a list of all monitor and monitor group guids for the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition for which to return the members.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionGetAllMembers(context.Background(), alertDefinitionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionGetAllMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetAllMembers`: []AlertDefinitionMember
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionGetAllMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition for which to return the members. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetAllMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AlertDefinitionMember**](AlertDefinitionMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetEscalationLevel

> EscalationLevel AlertDefinitionGetEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId).Execute()

Gets the escalation level information of the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionGetEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionGetEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetEscalationLevel`: EscalationLevel
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionGetEscalationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetEscalationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EscalationLevel**](EscalationLevel.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetEscalationLevelIntegration

> Integration AlertDefinitionGetEscalationLevelIntegration(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid).Execute()

Gets a single integration for a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    integrationGuid := "integrationGuid_example" // string | The Guid of the integration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionGetEscalationLevelIntegration(context.Background(), alertDefinitionGuid, escalationLevelId, integrationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionGetEscalationLevelIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetEscalationLevelIntegration`: Integration
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionGetEscalationLevelIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**integrationGuid** | **string** | The Guid of the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetEscalationLevelIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Integration**](Integration.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetEscalationLevelOperator

> []AlertEscalationLevelMember AlertDefinitionGetEscalationLevelOperator(ctx, alertDefinitionGuid, escalationLevelId).Execute()

Gets the operator and operator group guids for the specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionGetEscalationLevelOperator(context.Background(), alertDefinitionGuid, escalationLevelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionGetEscalationLevelOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetEscalationLevelOperator`: []AlertEscalationLevelMember
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionGetEscalationLevelOperator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetEscalationLevelOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AlertEscalationLevelMember**](AlertEscalationLevelMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetSpecifiedAlertDefinitions

> AlertDefinition AlertDefinitionGetSpecifiedAlertDefinitions(ctx, alertDefinitionGuid).Execute()

Gets the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionGetSpecifiedAlertDefinitions(context.Background(), alertDefinitionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionGetSpecifiedAlertDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetSpecifiedAlertDefinitions`: AlertDefinition
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionApi.AlertDefinitionGetSpecifiedAlertDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetSpecifiedAlertDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionPatchAlertDefinition

> AlertDefinitionPatchAlertDefinition(ctx, alertDefinitionGuid).AlertDefinition(alertDefinition).Execute()

Partially updates the definition for the specified alert definition.



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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition that should be updated.
    alertDefinition := *openapiclient.NewAlertDefinition() // AlertDefinition | The partial definition for the alert definition that should be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionPatchAlertDefinition(context.Background(), alertDefinitionGuid).AlertDefinition(alertDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionPatchAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionPatchAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertDefinition** | [**AlertDefinition**](AlertDefinition.md) | The partial definition for the alert definition that should be updated. | 

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


## AlertDefinitionPatchAlertDefinitionEscalation

> AlertDefinitionPatchAlertDefinitionEscalation(ctx, alertDefinitionGuid, escalationLevelId).EscalationLevel(escalationLevel).Execute()

Partially updates the escalation level for the specified alert definition.



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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition that should be updated.
    escalationLevelId := int32(56) // int32 | The level number of the escalation that should be updated.
    escalationLevel := *openapiclient.NewEscalationLevel() // EscalationLevel | The escalation level for the alert definition that should be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionPatchAlertDefinitionEscalation(context.Background(), alertDefinitionGuid, escalationLevelId).EscalationLevel(escalationLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionPatchAlertDefinitionEscalation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition that should be updated. | 
**escalationLevelId** | **int32** | The level number of the escalation that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionPatchAlertDefinitionEscalationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **escalationLevel** | [**EscalationLevel**](EscalationLevel.md) | The escalation level for the alert definition that should be updated. | 

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


## AlertDefinitionPutAlertDefinition

> AlertDefinitionPutAlertDefinition(ctx, alertDefinitionGuid).AlertDefinition(alertDefinition).Execute()

Updates the definition for the specified alert definition.



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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition that should be updated.
    alertDefinition := *openapiclient.NewAlertDefinition() // AlertDefinition | The partial definition for the alert definition that should be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionPutAlertDefinition(context.Background(), alertDefinitionGuid).AlertDefinition(alertDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionPutAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionPutAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertDefinition** | [**AlertDefinition**](AlertDefinition.md) | The partial definition for the alert definition that should be updated. | 

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


## AlertDefinitionPutAlertDefinitionEscalation

> AlertDefinitionPutAlertDefinitionEscalation(ctx, alertDefinitionGuid, escalationLevelId).EscalationLevel(escalationLevel).Execute()

Updates the escalation level for the specified alert definition.



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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition that should be updated.
    escalationLevelId := int32(56) // int32 | The level number of the escalation that should be updated.
    escalationLevel := *openapiclient.NewEscalationLevel() // EscalationLevel | The escalation level for the alert definition that should be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionPutAlertDefinitionEscalation(context.Background(), alertDefinitionGuid, escalationLevelId).EscalationLevel(escalationLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionPutAlertDefinitionEscalation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition that should be updated. | 
**escalationLevelId** | **int32** | The level number of the escalation that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionPutAlertDefinitionEscalationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **escalationLevel** | [**EscalationLevel**](EscalationLevel.md) | The escalation level for the alert definition that should be updated. | 

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


## AlertDefinitionRemoveIntegrationFromEscalationLevel

> AlertDefinitionRemoveIntegrationFromEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid).Execute()

Removes an integration from a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    integrationGuid := "integrationGuid_example" // string | The Guid of the integration to remove.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionRemoveIntegrationFromEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId, integrationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionRemoveIntegrationFromEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**integrationGuid** | **string** | The Guid of the integration to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionRemoveIntegrationFromEscalationLevelRequest struct via the builder pattern


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


## AlertDefinitionRemoveMonitorFromAlertDefinition

> AlertDefinitionRemoveMonitorFromAlertDefinition(ctx, alertDefinitionGuid, monitorGuid).Execute()

Removes a monitor for the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition to modify.
    monitorGuid := "monitorGuid_example" // string | The Guid of the monitor to remove.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionRemoveMonitorFromAlertDefinition(context.Background(), alertDefinitionGuid, monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionRemoveMonitorFromAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition to modify. | 
**monitorGuid** | **string** | The Guid of the monitor to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionRemoveMonitorFromAlertDefinitionRequest struct via the builder pattern


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


## AlertDefinitionRemoveMonitorGroupFromAlertDefinition

> AlertDefinitionRemoveMonitorGroupFromAlertDefinition(ctx, alertDefinitionGuid, monitorGroupGuid).Execute()

Removes a monitor group for the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition to modify.
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to remove.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionRemoveMonitorGroupFromAlertDefinition(context.Background(), alertDefinitionGuid, monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionRemoveMonitorGroupFromAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition to modify. | 
**monitorGroupGuid** | **string** | The Guid of the monitor group to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionRemoveMonitorGroupFromAlertDefinitionRequest struct via the builder pattern


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


## AlertDefinitionRemoveOperatorFromEscalationLevel

> AlertDefinitionRemoveOperatorFromEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGuid).Execute()

Removes an operator for the specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator to remove.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionRemoveOperatorFromEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId, operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionRemoveOperatorFromEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**operatorGuid** | **string** | The Guid of the operator to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionRemoveOperatorFromEscalationLevelRequest struct via the builder pattern


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


## AlertDefinitionRemoveOperatorGroupFromEscalationLevel

> AlertDefinitionRemoveOperatorGroupFromEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGroupGuid).Execute()

Removes an operator group for the specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group to remove.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionRemoveOperatorGroupFromEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId, operatorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionRemoveOperatorGroupFromEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**operatorGroupGuid** | **string** | The Guid of the operator group to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionRemoveOperatorGroupFromEscalationLevelRequest struct via the builder pattern


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


## AlertDefinitionUpdateIntegrationForEscalationWithPatch

> AlertDefinitionUpdateIntegrationForEscalationWithPatch(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid).EscalationLevelIntegration(escalationLevelIntegration).Execute()

Partially updates an integration for a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    integrationGuid := "integrationGuid_example" // string | The Guid of the integration to update.
    escalationLevelIntegration := *openapiclient.NewEscalationLevelIntegration() // EscalationLevelIntegration | The partial definition for the integration that should be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionUpdateIntegrationForEscalationWithPatch(context.Background(), alertDefinitionGuid, escalationLevelId, integrationGuid).EscalationLevelIntegration(escalationLevelIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionUpdateIntegrationForEscalationWithPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**integrationGuid** | **string** | The Guid of the integration to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionUpdateIntegrationForEscalationWithPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **escalationLevelIntegration** | [**EscalationLevelIntegration**](EscalationLevelIntegration.md) | The partial definition for the integration that should be updated. | 

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


## AlertDefinitionUpdateIntegrationForEscalationWithPut

> AlertDefinitionUpdateIntegrationForEscalationWithPut(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid).EscalationLevelIntegration(escalationLevelIntegration).Execute()

Updates an integration for a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    integrationGuid := "integrationGuid_example" // string | The Guid of the integration to update.
    escalationLevelIntegration := *openapiclient.NewEscalationLevelIntegration() // EscalationLevelIntegration | The definition for the integration that should be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertDefinitionApi.AlertDefinitionUpdateIntegrationForEscalationWithPut(context.Background(), alertDefinitionGuid, escalationLevelId, integrationGuid).EscalationLevelIntegration(escalationLevelIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionApi.AlertDefinitionUpdateIntegrationForEscalationWithPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**integrationGuid** | **string** | The Guid of the integration to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionUpdateIntegrationForEscalationWithPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **escalationLevelIntegration** | [**EscalationLevelIntegration**](EscalationLevelIntegration.md) | The definition for the integration that should be updated. | 

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

