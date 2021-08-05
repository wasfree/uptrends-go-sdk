# AlertDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertDefinitionGuid** | Pointer to **string** | The unique key of this Alert Definition. | [optional] 
**Hash** | Pointer to **string** | Hash corresponding with this alert definition. | [optional] 
**AlertName** | Pointer to **string** | Name of this Alert Definition. | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this Alert Definition is active. | [optional] 

## Methods

### NewAlertDefinition

`func NewAlertDefinition() *AlertDefinition`

NewAlertDefinition instantiates a new AlertDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertDefinitionWithDefaults

`func NewAlertDefinitionWithDefaults() *AlertDefinition`

NewAlertDefinitionWithDefaults instantiates a new AlertDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertDefinitionGuid

`func (o *AlertDefinition) GetAlertDefinitionGuid() string`

GetAlertDefinitionGuid returns the AlertDefinitionGuid field if non-nil, zero value otherwise.

### GetAlertDefinitionGuidOk

`func (o *AlertDefinition) GetAlertDefinitionGuidOk() (*string, bool)`

GetAlertDefinitionGuidOk returns a tuple with the AlertDefinitionGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDefinitionGuid

`func (o *AlertDefinition) SetAlertDefinitionGuid(v string)`

SetAlertDefinitionGuid sets AlertDefinitionGuid field to given value.

### HasAlertDefinitionGuid

`func (o *AlertDefinition) HasAlertDefinitionGuid() bool`

HasAlertDefinitionGuid returns a boolean if a field has been set.

### GetHash

`func (o *AlertDefinition) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AlertDefinition) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AlertDefinition) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *AlertDefinition) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetAlertName

`func (o *AlertDefinition) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *AlertDefinition) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *AlertDefinition) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.

### HasAlertName

`func (o *AlertDefinition) HasAlertName() bool`

HasAlertName returns a boolean if a field has been set.

### GetIsActive

`func (o *AlertDefinition) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AlertDefinition) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AlertDefinition) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AlertDefinition) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


