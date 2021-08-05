# AlertDefinitionMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorGuid** | Pointer to **string** | Guid of Monitor in Alert Definition | [optional] 
**MonitorGroupGuid** | Pointer to **string** | Guid of GroupMonitor in Alert Definition | [optional] 

## Methods

### NewAlertDefinitionMember

`func NewAlertDefinitionMember() *AlertDefinitionMember`

NewAlertDefinitionMember instantiates a new AlertDefinitionMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertDefinitionMemberWithDefaults

`func NewAlertDefinitionMemberWithDefaults() *AlertDefinitionMember`

NewAlertDefinitionMemberWithDefaults instantiates a new AlertDefinitionMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorGuid

`func (o *AlertDefinitionMember) GetMonitorGuid() string`

GetMonitorGuid returns the MonitorGuid field if non-nil, zero value otherwise.

### GetMonitorGuidOk

`func (o *AlertDefinitionMember) GetMonitorGuidOk() (*string, bool)`

GetMonitorGuidOk returns a tuple with the MonitorGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGuid

`func (o *AlertDefinitionMember) SetMonitorGuid(v string)`

SetMonitorGuid sets MonitorGuid field to given value.

### HasMonitorGuid

`func (o *AlertDefinitionMember) HasMonitorGuid() bool`

HasMonitorGuid returns a boolean if a field has been set.

### GetMonitorGroupGuid

`func (o *AlertDefinitionMember) GetMonitorGroupGuid() string`

GetMonitorGroupGuid returns the MonitorGroupGuid field if non-nil, zero value otherwise.

### GetMonitorGroupGuidOk

`func (o *AlertDefinitionMember) GetMonitorGroupGuidOk() (*string, bool)`

GetMonitorGroupGuidOk returns a tuple with the MonitorGroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGroupGuid

`func (o *AlertDefinitionMember) SetMonitorGroupGuid(v string)`

SetMonitorGroupGuid sets MonitorGroupGuid field to given value.

### HasMonitorGroupGuid

`func (o *AlertDefinitionMember) HasMonitorGroupGuid() bool`

HasMonitorGroupGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


