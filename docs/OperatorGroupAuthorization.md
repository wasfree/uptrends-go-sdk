# OperatorGroupAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationId** | Pointer to **string** |  | [optional] 
**AuthorizationType** | [**OperatorGroupAuthorizationType**](OperatorGroupAuthorizationType.md) |  | 
**OperatorGuid** | Pointer to **string** |  | [optional] 
**OperatorGroupGuid** | Pointer to **string** |  | [optional] 

## Methods

### NewOperatorGroupAuthorization

`func NewOperatorGroupAuthorization(authorizationType OperatorGroupAuthorizationType, ) *OperatorGroupAuthorization`

NewOperatorGroupAuthorization instantiates a new OperatorGroupAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorGroupAuthorizationWithDefaults

`func NewOperatorGroupAuthorizationWithDefaults() *OperatorGroupAuthorization`

NewOperatorGroupAuthorizationWithDefaults instantiates a new OperatorGroupAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationId

`func (o *OperatorGroupAuthorization) GetAuthorizationId() string`

GetAuthorizationId returns the AuthorizationId field if non-nil, zero value otherwise.

### GetAuthorizationIdOk

`func (o *OperatorGroupAuthorization) GetAuthorizationIdOk() (*string, bool)`

GetAuthorizationIdOk returns a tuple with the AuthorizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationId

`func (o *OperatorGroupAuthorization) SetAuthorizationId(v string)`

SetAuthorizationId sets AuthorizationId field to given value.

### HasAuthorizationId

`func (o *OperatorGroupAuthorization) HasAuthorizationId() bool`

HasAuthorizationId returns a boolean if a field has been set.

### GetAuthorizationType

`func (o *OperatorGroupAuthorization) GetAuthorizationType() OperatorGroupAuthorizationType`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *OperatorGroupAuthorization) GetAuthorizationTypeOk() (*OperatorGroupAuthorizationType, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *OperatorGroupAuthorization) SetAuthorizationType(v OperatorGroupAuthorizationType)`

SetAuthorizationType sets AuthorizationType field to given value.


### GetOperatorGuid

`func (o *OperatorGroupAuthorization) GetOperatorGuid() string`

GetOperatorGuid returns the OperatorGuid field if non-nil, zero value otherwise.

### GetOperatorGuidOk

`func (o *OperatorGroupAuthorization) GetOperatorGuidOk() (*string, bool)`

GetOperatorGuidOk returns a tuple with the OperatorGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorGuid

`func (o *OperatorGroupAuthorization) SetOperatorGuid(v string)`

SetOperatorGuid sets OperatorGuid field to given value.

### HasOperatorGuid

`func (o *OperatorGroupAuthorization) HasOperatorGuid() bool`

HasOperatorGuid returns a boolean if a field has been set.

### GetOperatorGroupGuid

`func (o *OperatorGroupAuthorization) GetOperatorGroupGuid() string`

GetOperatorGroupGuid returns the OperatorGroupGuid field if non-nil, zero value otherwise.

### GetOperatorGroupGuidOk

`func (o *OperatorGroupAuthorization) GetOperatorGroupGuidOk() (*string, bool)`

GetOperatorGroupGuidOk returns a tuple with the OperatorGroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorGroupGuid

`func (o *OperatorGroupAuthorization) SetOperatorGroupGuid(v string)`

SetOperatorGroupGuid sets OperatorGroupGuid field to given value.

### HasOperatorGroupGuid

`func (o *OperatorGroupAuthorization) HasOperatorGroupGuid() bool`

HasOperatorGroupGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


