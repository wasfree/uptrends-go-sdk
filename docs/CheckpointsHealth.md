# CheckpointsHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthyServers** | **int32** | The number of healthy servers | 
**Servers** | Pointer to [**[]ServerHealth**](ServerHealth.md) | The health of each server | [optional] 

## Methods

### NewCheckpointsHealth

`func NewCheckpointsHealth(healthyServers int32, ) *CheckpointsHealth`

NewCheckpointsHealth instantiates a new CheckpointsHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointsHealthWithDefaults

`func NewCheckpointsHealthWithDefaults() *CheckpointsHealth`

NewCheckpointsHealthWithDefaults instantiates a new CheckpointsHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthyServers

`func (o *CheckpointsHealth) GetHealthyServers() int32`

GetHealthyServers returns the HealthyServers field if non-nil, zero value otherwise.

### GetHealthyServersOk

`func (o *CheckpointsHealth) GetHealthyServersOk() (*int32, bool)`

GetHealthyServersOk returns a tuple with the HealthyServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyServers

`func (o *CheckpointsHealth) SetHealthyServers(v int32)`

SetHealthyServers sets HealthyServers field to given value.


### GetServers

`func (o *CheckpointsHealth) GetServers() []ServerHealth`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *CheckpointsHealth) GetServersOk() (*[]ServerHealth, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *CheckpointsHealth) SetServers(v []ServerHealth)`

SetServers sets Servers field to given value.

### HasServers

`func (o *CheckpointsHealth) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


