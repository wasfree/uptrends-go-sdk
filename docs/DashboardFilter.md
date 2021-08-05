# DashboardFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedPeriod** | Pointer to [**SelectedPeriod**](SelectedPeriod.md) |  | [optional] 
**Monitors** | Pointer to **[]string** |  | [optional] 
**MonitorGroups** | Pointer to **[]string** |  | [optional] 
**Checkpoints** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewDashboardFilter

`func NewDashboardFilter() *DashboardFilter`

NewDashboardFilter instantiates a new DashboardFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardFilterWithDefaults

`func NewDashboardFilterWithDefaults() *DashboardFilter`

NewDashboardFilterWithDefaults instantiates a new DashboardFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedPeriod

`func (o *DashboardFilter) GetSelectedPeriod() SelectedPeriod`

GetSelectedPeriod returns the SelectedPeriod field if non-nil, zero value otherwise.

### GetSelectedPeriodOk

`func (o *DashboardFilter) GetSelectedPeriodOk() (*SelectedPeriod, bool)`

GetSelectedPeriodOk returns a tuple with the SelectedPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPeriod

`func (o *DashboardFilter) SetSelectedPeriod(v SelectedPeriod)`

SetSelectedPeriod sets SelectedPeriod field to given value.

### HasSelectedPeriod

`func (o *DashboardFilter) HasSelectedPeriod() bool`

HasSelectedPeriod returns a boolean if a field has been set.

### GetMonitors

`func (o *DashboardFilter) GetMonitors() []string`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *DashboardFilter) GetMonitorsOk() (*[]string, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *DashboardFilter) SetMonitors(v []string)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *DashboardFilter) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetMonitorGroups

`func (o *DashboardFilter) GetMonitorGroups() []string`

GetMonitorGroups returns the MonitorGroups field if non-nil, zero value otherwise.

### GetMonitorGroupsOk

`func (o *DashboardFilter) GetMonitorGroupsOk() (*[]string, bool)`

GetMonitorGroupsOk returns a tuple with the MonitorGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGroups

`func (o *DashboardFilter) SetMonitorGroups(v []string)`

SetMonitorGroups sets MonitorGroups field to given value.

### HasMonitorGroups

`func (o *DashboardFilter) HasMonitorGroups() bool`

HasMonitorGroups returns a boolean if a field has been set.

### GetCheckpoints

`func (o *DashboardFilter) GetCheckpoints() []int32`

GetCheckpoints returns the Checkpoints field if non-nil, zero value otherwise.

### GetCheckpointsOk

`func (o *DashboardFilter) GetCheckpointsOk() (*[]int32, bool)`

GetCheckpointsOk returns a tuple with the Checkpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoints

`func (o *DashboardFilter) SetCheckpoints(v []int32)`

SetCheckpoints sets Checkpoints field to given value.

### HasCheckpoints

`func (o *DashboardFilter) HasCheckpoints() bool`

HasCheckpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


