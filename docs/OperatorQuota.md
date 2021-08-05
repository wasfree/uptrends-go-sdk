# OperatorQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operators** | Pointer to **int32** |  | [optional] 
**OperatorsInUse** | Pointer to **int32** |  | [optional] 

## Methods

### NewOperatorQuota

`func NewOperatorQuota() *OperatorQuota`

NewOperatorQuota instantiates a new OperatorQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorQuotaWithDefaults

`func NewOperatorQuotaWithDefaults() *OperatorQuota`

NewOperatorQuotaWithDefaults instantiates a new OperatorQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperators

`func (o *OperatorQuota) GetOperators() int32`

GetOperators returns the Operators field if non-nil, zero value otherwise.

### GetOperatorsOk

`func (o *OperatorQuota) GetOperatorsOk() (*int32, bool)`

GetOperatorsOk returns a tuple with the Operators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperators

`func (o *OperatorQuota) SetOperators(v int32)`

SetOperators sets Operators field to given value.

### HasOperators

`func (o *OperatorQuota) HasOperators() bool`

HasOperators returns a boolean if a field has been set.

### GetOperatorsInUse

`func (o *OperatorQuota) GetOperatorsInUse() int32`

GetOperatorsInUse returns the OperatorsInUse field if non-nil, zero value otherwise.

### GetOperatorsInUseOk

`func (o *OperatorQuota) GetOperatorsInUseOk() (*int32, bool)`

GetOperatorsInUseOk returns a tuple with the OperatorsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorsInUse

`func (o *OperatorQuota) SetOperatorsInUse(v int32)`

SetOperatorsInUse sets OperatorsInUse field to given value.

### HasOperatorsInUse

`func (o *OperatorQuota) HasOperatorsInUse() bool`

HasOperatorsInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


