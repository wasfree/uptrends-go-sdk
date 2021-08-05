# MonitorGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorGroupGuid** | Pointer to **string** | The unique identifier of this monitor group | [optional] 
**Description** | Pointer to **string** | The descriptive name of this probe group | [optional] 
**IsAll** | **bool** | Indicates whether this is the default group for all probes | 

## Methods

### NewMonitorGroup

`func NewMonitorGroup(isAll bool, ) *MonitorGroup`

NewMonitorGroup instantiates a new MonitorGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorGroupWithDefaults

`func NewMonitorGroupWithDefaults() *MonitorGroup`

NewMonitorGroupWithDefaults instantiates a new MonitorGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorGroupGuid

`func (o *MonitorGroup) GetMonitorGroupGuid() string`

GetMonitorGroupGuid returns the MonitorGroupGuid field if non-nil, zero value otherwise.

### GetMonitorGroupGuidOk

`func (o *MonitorGroup) GetMonitorGroupGuidOk() (*string, bool)`

GetMonitorGroupGuidOk returns a tuple with the MonitorGroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGroupGuid

`func (o *MonitorGroup) SetMonitorGroupGuid(v string)`

SetMonitorGroupGuid sets MonitorGroupGuid field to given value.

### HasMonitorGroupGuid

`func (o *MonitorGroup) HasMonitorGroupGuid() bool`

HasMonitorGroupGuid returns a boolean if a field has been set.

### GetDescription

`func (o *MonitorGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitorGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitorGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MonitorGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsAll

`func (o *MonitorGroup) GetIsAll() bool`

GetIsAll returns the IsAll field if non-nil, zero value otherwise.

### GetIsAllOk

`func (o *MonitorGroup) GetIsAllOk() (*bool, bool)`

GetIsAllOk returns a tuple with the IsAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAll

`func (o *MonitorGroup) SetIsAll(v bool)`

SetIsAll sets IsAll field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


