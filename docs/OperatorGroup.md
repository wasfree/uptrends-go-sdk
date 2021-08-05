# OperatorGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatorGroupGuid** | Pointer to **string** | The unique identifier of this Operator group | [optional] 
**Description** | Pointer to **string** | The descriptive name of this operator group | [optional] 
**IsEveryone** | Pointer to **bool** | Indicates whether this is the default group for all operators | [optional] 
**IsAdministratorsGroup** | Pointer to **bool** |  | [optional] 

## Methods

### NewOperatorGroup

`func NewOperatorGroup() *OperatorGroup`

NewOperatorGroup instantiates a new OperatorGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorGroupWithDefaults

`func NewOperatorGroupWithDefaults() *OperatorGroup`

NewOperatorGroupWithDefaults instantiates a new OperatorGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatorGroupGuid

`func (o *OperatorGroup) GetOperatorGroupGuid() string`

GetOperatorGroupGuid returns the OperatorGroupGuid field if non-nil, zero value otherwise.

### GetOperatorGroupGuidOk

`func (o *OperatorGroup) GetOperatorGroupGuidOk() (*string, bool)`

GetOperatorGroupGuidOk returns a tuple with the OperatorGroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorGroupGuid

`func (o *OperatorGroup) SetOperatorGroupGuid(v string)`

SetOperatorGroupGuid sets OperatorGroupGuid field to given value.

### HasOperatorGroupGuid

`func (o *OperatorGroup) HasOperatorGroupGuid() bool`

HasOperatorGroupGuid returns a boolean if a field has been set.

### GetDescription

`func (o *OperatorGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OperatorGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OperatorGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OperatorGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEveryone

`func (o *OperatorGroup) GetIsEveryone() bool`

GetIsEveryone returns the IsEveryone field if non-nil, zero value otherwise.

### GetIsEveryoneOk

`func (o *OperatorGroup) GetIsEveryoneOk() (*bool, bool)`

GetIsEveryoneOk returns a tuple with the IsEveryone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEveryone

`func (o *OperatorGroup) SetIsEveryone(v bool)`

SetIsEveryone sets IsEveryone field to given value.

### HasIsEveryone

`func (o *OperatorGroup) HasIsEveryone() bool`

HasIsEveryone returns a boolean if a field has been set.

### GetIsAdministratorsGroup

`func (o *OperatorGroup) GetIsAdministratorsGroup() bool`

GetIsAdministratorsGroup returns the IsAdministratorsGroup field if non-nil, zero value otherwise.

### GetIsAdministratorsGroupOk

`func (o *OperatorGroup) GetIsAdministratorsGroupOk() (*bool, bool)`

GetIsAdministratorsGroupOk returns a tuple with the IsAdministratorsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdministratorsGroup

`func (o *OperatorGroup) SetIsAdministratorsGroup(v bool)`

SetIsAdministratorsGroup sets IsAdministratorsGroup field to given value.

### HasIsAdministratorsGroup

`func (o *OperatorGroup) HasIsAdministratorsGroup() bool`

HasIsAdministratorsGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


