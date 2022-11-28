# IntegrationAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationId** | Pointer to **string** | The unique ID of this authorization | [optional] 
**AuthorizationType** | [**IntegrationAuthorizationType**](IntegrationAuthorizationType.md) | The authorization type | 
**OperatorGuid** | Pointer to **string** | The GUID of the operator (NULL if this authorization is for an operator group) | [optional] 
**OperatorGroupGuid** | Pointer to **string** | The GUID of the operator group (NULL if this authorization is for an operator) | [optional] 

## Methods

### NewIntegrationAuthorization

`func NewIntegrationAuthorization(authorizationType IntegrationAuthorizationType, ) *IntegrationAuthorization`

NewIntegrationAuthorization instantiates a new IntegrationAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationAuthorizationWithDefaults

`func NewIntegrationAuthorizationWithDefaults() *IntegrationAuthorization`

NewIntegrationAuthorizationWithDefaults instantiates a new IntegrationAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationId

`func (o *IntegrationAuthorization) GetAuthorizationId() string`

GetAuthorizationId returns the AuthorizationId field if non-nil, zero value otherwise.

### GetAuthorizationIdOk

`func (o *IntegrationAuthorization) GetAuthorizationIdOk() (*string, bool)`

GetAuthorizationIdOk returns a tuple with the AuthorizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationId

`func (o *IntegrationAuthorization) SetAuthorizationId(v string)`

SetAuthorizationId sets AuthorizationId field to given value.

### HasAuthorizationId

`func (o *IntegrationAuthorization) HasAuthorizationId() bool`

HasAuthorizationId returns a boolean if a field has been set.

### GetAuthorizationType

`func (o *IntegrationAuthorization) GetAuthorizationType() IntegrationAuthorizationType`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *IntegrationAuthorization) GetAuthorizationTypeOk() (*IntegrationAuthorizationType, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *IntegrationAuthorization) SetAuthorizationType(v IntegrationAuthorizationType)`

SetAuthorizationType sets AuthorizationType field to given value.


### GetOperatorGuid

`func (o *IntegrationAuthorization) GetOperatorGuid() string`

GetOperatorGuid returns the OperatorGuid field if non-nil, zero value otherwise.

### GetOperatorGuidOk

`func (o *IntegrationAuthorization) GetOperatorGuidOk() (*string, bool)`

GetOperatorGuidOk returns a tuple with the OperatorGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorGuid

`func (o *IntegrationAuthorization) SetOperatorGuid(v string)`

SetOperatorGuid sets OperatorGuid field to given value.

### HasOperatorGuid

`func (o *IntegrationAuthorization) HasOperatorGuid() bool`

HasOperatorGuid returns a boolean if a field has been set.

### GetOperatorGroupGuid

`func (o *IntegrationAuthorization) GetOperatorGroupGuid() string`

GetOperatorGroupGuid returns the OperatorGroupGuid field if non-nil, zero value otherwise.

### GetOperatorGroupGuidOk

`func (o *IntegrationAuthorization) GetOperatorGroupGuidOk() (*string, bool)`

GetOperatorGroupGuidOk returns a tuple with the OperatorGroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorGroupGuid

`func (o *IntegrationAuthorization) SetOperatorGroupGuid(v string)`

SetOperatorGroupGuid sets OperatorGroupGuid field to given value.

### HasOperatorGroupGuid

`func (o *IntegrationAuthorization) HasOperatorGroupGuid() bool`

HasOperatorGroupGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


