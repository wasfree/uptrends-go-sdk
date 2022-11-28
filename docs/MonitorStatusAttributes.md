# MonitorStatusAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorLevel** | [**LastErrorLevel**](LastErrorLevel.md) | The error level for the monitor status | 
**LastCheck** | Pointer to **map[string]interface{}** | Last checked timeStamp for this monitor | [optional] 
**CheckpointId** | Pointer to **int32** | Checkpoint id for the monitor status | [optional] 
**CheckpointName** | Pointer to **string** | Checkpoint name for the monitor status | [optional] 
**ErrorDescription** | Pointer to **string** | Error description for the monitor status | [optional] 
**UptimePercentage** | **float64** | Uptime percentage for the monitor status | 
**ErrorCode** | **int32** | Error code for the monitor status | 
**LastMonitorCheckId** | Pointer to **int64** | Last monitor check id | [optional] 
**TotalTime** | Pointer to **int32** | Total time for the monitor status | [optional] 

## Methods

### NewMonitorStatusAttributes

`func NewMonitorStatusAttributes(errorLevel LastErrorLevel, uptimePercentage float64, errorCode int32, ) *MonitorStatusAttributes`

NewMonitorStatusAttributes instantiates a new MonitorStatusAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorStatusAttributesWithDefaults

`func NewMonitorStatusAttributesWithDefaults() *MonitorStatusAttributes`

NewMonitorStatusAttributesWithDefaults instantiates a new MonitorStatusAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorLevel

`func (o *MonitorStatusAttributes) GetErrorLevel() LastErrorLevel`

GetErrorLevel returns the ErrorLevel field if non-nil, zero value otherwise.

### GetErrorLevelOk

`func (o *MonitorStatusAttributes) GetErrorLevelOk() (*LastErrorLevel, bool)`

GetErrorLevelOk returns a tuple with the ErrorLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLevel

`func (o *MonitorStatusAttributes) SetErrorLevel(v LastErrorLevel)`

SetErrorLevel sets ErrorLevel field to given value.


### GetLastCheck

`func (o *MonitorStatusAttributes) GetLastCheck() map[string]interface{}`

GetLastCheck returns the LastCheck field if non-nil, zero value otherwise.

### GetLastCheckOk

`func (o *MonitorStatusAttributes) GetLastCheckOk() (*map[string]interface{}, bool)`

GetLastCheckOk returns a tuple with the LastCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheck

`func (o *MonitorStatusAttributes) SetLastCheck(v map[string]interface{})`

SetLastCheck sets LastCheck field to given value.

### HasLastCheck

`func (o *MonitorStatusAttributes) HasLastCheck() bool`

HasLastCheck returns a boolean if a field has been set.

### GetCheckpointId

`func (o *MonitorStatusAttributes) GetCheckpointId() int32`

GetCheckpointId returns the CheckpointId field if non-nil, zero value otherwise.

### GetCheckpointIdOk

`func (o *MonitorStatusAttributes) GetCheckpointIdOk() (*int32, bool)`

GetCheckpointIdOk returns a tuple with the CheckpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointId

`func (o *MonitorStatusAttributes) SetCheckpointId(v int32)`

SetCheckpointId sets CheckpointId field to given value.

### HasCheckpointId

`func (o *MonitorStatusAttributes) HasCheckpointId() bool`

HasCheckpointId returns a boolean if a field has been set.

### GetCheckpointName

`func (o *MonitorStatusAttributes) GetCheckpointName() string`

GetCheckpointName returns the CheckpointName field if non-nil, zero value otherwise.

### GetCheckpointNameOk

`func (o *MonitorStatusAttributes) GetCheckpointNameOk() (*string, bool)`

GetCheckpointNameOk returns a tuple with the CheckpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointName

`func (o *MonitorStatusAttributes) SetCheckpointName(v string)`

SetCheckpointName sets CheckpointName field to given value.

### HasCheckpointName

`func (o *MonitorStatusAttributes) HasCheckpointName() bool`

HasCheckpointName returns a boolean if a field has been set.

### GetErrorDescription

`func (o *MonitorStatusAttributes) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *MonitorStatusAttributes) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *MonitorStatusAttributes) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *MonitorStatusAttributes) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetUptimePercentage

`func (o *MonitorStatusAttributes) GetUptimePercentage() float64`

GetUptimePercentage returns the UptimePercentage field if non-nil, zero value otherwise.

### GetUptimePercentageOk

`func (o *MonitorStatusAttributes) GetUptimePercentageOk() (*float64, bool)`

GetUptimePercentageOk returns a tuple with the UptimePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimePercentage

`func (o *MonitorStatusAttributes) SetUptimePercentage(v float64)`

SetUptimePercentage sets UptimePercentage field to given value.


### GetErrorCode

`func (o *MonitorStatusAttributes) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MonitorStatusAttributes) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MonitorStatusAttributes) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.


### GetLastMonitorCheckId

`func (o *MonitorStatusAttributes) GetLastMonitorCheckId() int64`

GetLastMonitorCheckId returns the LastMonitorCheckId field if non-nil, zero value otherwise.

### GetLastMonitorCheckIdOk

`func (o *MonitorStatusAttributes) GetLastMonitorCheckIdOk() (*int64, bool)`

GetLastMonitorCheckIdOk returns a tuple with the LastMonitorCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMonitorCheckId

`func (o *MonitorStatusAttributes) SetLastMonitorCheckId(v int64)`

SetLastMonitorCheckId sets LastMonitorCheckId field to given value.

### HasLastMonitorCheckId

`func (o *MonitorStatusAttributes) HasLastMonitorCheckId() bool`

HasLastMonitorCheckId returns a boolean if a field has been set.

### GetTotalTime

`func (o *MonitorStatusAttributes) GetTotalTime() int32`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *MonitorStatusAttributes) GetTotalTimeOk() (*int32, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *MonitorStatusAttributes) SetTotalTime(v int32)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *MonitorStatusAttributes) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


