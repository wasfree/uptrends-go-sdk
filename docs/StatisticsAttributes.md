# StatisticsAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTime** | **map[string]interface{}** |  | 
**EndDateTime** | **map[string]interface{}** |  | 
**Alerts** | Pointer to **int32** |  | [optional] 
**SlaTarget** | Pointer to **float64** |  | [optional] 
**SlaTargetTotalTime** | Pointer to **float64** |  | [optional] 
**OperatorResponseTarget** | Pointer to **float64** |  | [optional] 
**OperatorResponseTime** | Pointer to **float64** |  | [optional] 
**DowntimePercentage** | Pointer to **float64** |  | [optional] 
**UptimePercentage** | Pointer to **float64** |  | [optional] 
**Checks** | Pointer to **int64** |  | [optional] 
**ConfirmedErrors** | Pointer to **int32** |  | [optional] 
**UnconfirmedErrors** | Pointer to **int32** |  | [optional] 
**Uptime** | Pointer to **int64** |  | [optional] 
**Downtime** | Pointer to **int64** |  | [optional] 
**TotalTime** | Pointer to **float32** |  | [optional] 
**ResolveTime** | Pointer to **float32** |  | [optional] 
**ConnectionTime** | Pointer to **float32** |  | [optional] 
**DownloadTime** | Pointer to **float32** |  | [optional] 
**TotalBytes** | Pointer to **int64** |  | [optional] 

## Methods

### NewStatisticsAttributes

`func NewStatisticsAttributes(startDateTime map[string]interface{}, endDateTime map[string]interface{}, ) *StatisticsAttributes`

NewStatisticsAttributes instantiates a new StatisticsAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsAttributesWithDefaults

`func NewStatisticsAttributesWithDefaults() *StatisticsAttributes`

NewStatisticsAttributesWithDefaults instantiates a new StatisticsAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTime

`func (o *StatisticsAttributes) GetStartDateTime() map[string]interface{}`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *StatisticsAttributes) GetStartDateTimeOk() (*map[string]interface{}, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *StatisticsAttributes) SetStartDateTime(v map[string]interface{})`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *StatisticsAttributes) GetEndDateTime() map[string]interface{}`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *StatisticsAttributes) GetEndDateTimeOk() (*map[string]interface{}, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *StatisticsAttributes) SetEndDateTime(v map[string]interface{})`

SetEndDateTime sets EndDateTime field to given value.


### GetAlerts

`func (o *StatisticsAttributes) GetAlerts() int32`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *StatisticsAttributes) GetAlertsOk() (*int32, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *StatisticsAttributes) SetAlerts(v int32)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *StatisticsAttributes) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetSlaTarget

`func (o *StatisticsAttributes) GetSlaTarget() float64`

GetSlaTarget returns the SlaTarget field if non-nil, zero value otherwise.

### GetSlaTargetOk

`func (o *StatisticsAttributes) GetSlaTargetOk() (*float64, bool)`

GetSlaTargetOk returns a tuple with the SlaTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaTarget

`func (o *StatisticsAttributes) SetSlaTarget(v float64)`

SetSlaTarget sets SlaTarget field to given value.

### HasSlaTarget

`func (o *StatisticsAttributes) HasSlaTarget() bool`

HasSlaTarget returns a boolean if a field has been set.

### GetSlaTargetTotalTime

`func (o *StatisticsAttributes) GetSlaTargetTotalTime() float64`

GetSlaTargetTotalTime returns the SlaTargetTotalTime field if non-nil, zero value otherwise.

### GetSlaTargetTotalTimeOk

`func (o *StatisticsAttributes) GetSlaTargetTotalTimeOk() (*float64, bool)`

GetSlaTargetTotalTimeOk returns a tuple with the SlaTargetTotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaTargetTotalTime

`func (o *StatisticsAttributes) SetSlaTargetTotalTime(v float64)`

SetSlaTargetTotalTime sets SlaTargetTotalTime field to given value.

### HasSlaTargetTotalTime

`func (o *StatisticsAttributes) HasSlaTargetTotalTime() bool`

HasSlaTargetTotalTime returns a boolean if a field has been set.

### GetOperatorResponseTarget

`func (o *StatisticsAttributes) GetOperatorResponseTarget() float64`

GetOperatorResponseTarget returns the OperatorResponseTarget field if non-nil, zero value otherwise.

### GetOperatorResponseTargetOk

`func (o *StatisticsAttributes) GetOperatorResponseTargetOk() (*float64, bool)`

GetOperatorResponseTargetOk returns a tuple with the OperatorResponseTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorResponseTarget

`func (o *StatisticsAttributes) SetOperatorResponseTarget(v float64)`

SetOperatorResponseTarget sets OperatorResponseTarget field to given value.

### HasOperatorResponseTarget

`func (o *StatisticsAttributes) HasOperatorResponseTarget() bool`

HasOperatorResponseTarget returns a boolean if a field has been set.

### GetOperatorResponseTime

`func (o *StatisticsAttributes) GetOperatorResponseTime() float64`

GetOperatorResponseTime returns the OperatorResponseTime field if non-nil, zero value otherwise.

### GetOperatorResponseTimeOk

`func (o *StatisticsAttributes) GetOperatorResponseTimeOk() (*float64, bool)`

GetOperatorResponseTimeOk returns a tuple with the OperatorResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorResponseTime

`func (o *StatisticsAttributes) SetOperatorResponseTime(v float64)`

SetOperatorResponseTime sets OperatorResponseTime field to given value.

### HasOperatorResponseTime

`func (o *StatisticsAttributes) HasOperatorResponseTime() bool`

HasOperatorResponseTime returns a boolean if a field has been set.

### GetDowntimePercentage

`func (o *StatisticsAttributes) GetDowntimePercentage() float64`

GetDowntimePercentage returns the DowntimePercentage field if non-nil, zero value otherwise.

### GetDowntimePercentageOk

`func (o *StatisticsAttributes) GetDowntimePercentageOk() (*float64, bool)`

GetDowntimePercentageOk returns a tuple with the DowntimePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimePercentage

`func (o *StatisticsAttributes) SetDowntimePercentage(v float64)`

SetDowntimePercentage sets DowntimePercentage field to given value.

### HasDowntimePercentage

`func (o *StatisticsAttributes) HasDowntimePercentage() bool`

HasDowntimePercentage returns a boolean if a field has been set.

### GetUptimePercentage

`func (o *StatisticsAttributes) GetUptimePercentage() float64`

GetUptimePercentage returns the UptimePercentage field if non-nil, zero value otherwise.

### GetUptimePercentageOk

`func (o *StatisticsAttributes) GetUptimePercentageOk() (*float64, bool)`

GetUptimePercentageOk returns a tuple with the UptimePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimePercentage

`func (o *StatisticsAttributes) SetUptimePercentage(v float64)`

SetUptimePercentage sets UptimePercentage field to given value.

### HasUptimePercentage

`func (o *StatisticsAttributes) HasUptimePercentage() bool`

HasUptimePercentage returns a boolean if a field has been set.

### GetChecks

`func (o *StatisticsAttributes) GetChecks() int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *StatisticsAttributes) GetChecksOk() (*int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *StatisticsAttributes) SetChecks(v int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *StatisticsAttributes) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetConfirmedErrors

`func (o *StatisticsAttributes) GetConfirmedErrors() int32`

GetConfirmedErrors returns the ConfirmedErrors field if non-nil, zero value otherwise.

### GetConfirmedErrorsOk

`func (o *StatisticsAttributes) GetConfirmedErrorsOk() (*int32, bool)`

GetConfirmedErrorsOk returns a tuple with the ConfirmedErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedErrors

`func (o *StatisticsAttributes) SetConfirmedErrors(v int32)`

SetConfirmedErrors sets ConfirmedErrors field to given value.

### HasConfirmedErrors

`func (o *StatisticsAttributes) HasConfirmedErrors() bool`

HasConfirmedErrors returns a boolean if a field has been set.

### GetUnconfirmedErrors

`func (o *StatisticsAttributes) GetUnconfirmedErrors() int32`

GetUnconfirmedErrors returns the UnconfirmedErrors field if non-nil, zero value otherwise.

### GetUnconfirmedErrorsOk

`func (o *StatisticsAttributes) GetUnconfirmedErrorsOk() (*int32, bool)`

GetUnconfirmedErrorsOk returns a tuple with the UnconfirmedErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfirmedErrors

`func (o *StatisticsAttributes) SetUnconfirmedErrors(v int32)`

SetUnconfirmedErrors sets UnconfirmedErrors field to given value.

### HasUnconfirmedErrors

`func (o *StatisticsAttributes) HasUnconfirmedErrors() bool`

HasUnconfirmedErrors returns a boolean if a field has been set.

### GetUptime

`func (o *StatisticsAttributes) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *StatisticsAttributes) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *StatisticsAttributes) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *StatisticsAttributes) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetDowntime

`func (o *StatisticsAttributes) GetDowntime() int64`

GetDowntime returns the Downtime field if non-nil, zero value otherwise.

### GetDowntimeOk

`func (o *StatisticsAttributes) GetDowntimeOk() (*int64, bool)`

GetDowntimeOk returns a tuple with the Downtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntime

`func (o *StatisticsAttributes) SetDowntime(v int64)`

SetDowntime sets Downtime field to given value.

### HasDowntime

`func (o *StatisticsAttributes) HasDowntime() bool`

HasDowntime returns a boolean if a field has been set.

### GetTotalTime

`func (o *StatisticsAttributes) GetTotalTime() float32`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *StatisticsAttributes) GetTotalTimeOk() (*float32, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *StatisticsAttributes) SetTotalTime(v float32)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *StatisticsAttributes) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.

### GetResolveTime

`func (o *StatisticsAttributes) GetResolveTime() float32`

GetResolveTime returns the ResolveTime field if non-nil, zero value otherwise.

### GetResolveTimeOk

`func (o *StatisticsAttributes) GetResolveTimeOk() (*float32, bool)`

GetResolveTimeOk returns a tuple with the ResolveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveTime

`func (o *StatisticsAttributes) SetResolveTime(v float32)`

SetResolveTime sets ResolveTime field to given value.

### HasResolveTime

`func (o *StatisticsAttributes) HasResolveTime() bool`

HasResolveTime returns a boolean if a field has been set.

### GetConnectionTime

`func (o *StatisticsAttributes) GetConnectionTime() float32`

GetConnectionTime returns the ConnectionTime field if non-nil, zero value otherwise.

### GetConnectionTimeOk

`func (o *StatisticsAttributes) GetConnectionTimeOk() (*float32, bool)`

GetConnectionTimeOk returns a tuple with the ConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTime

`func (o *StatisticsAttributes) SetConnectionTime(v float32)`

SetConnectionTime sets ConnectionTime field to given value.

### HasConnectionTime

`func (o *StatisticsAttributes) HasConnectionTime() bool`

HasConnectionTime returns a boolean if a field has been set.

### GetDownloadTime

`func (o *StatisticsAttributes) GetDownloadTime() float32`

GetDownloadTime returns the DownloadTime field if non-nil, zero value otherwise.

### GetDownloadTimeOk

`func (o *StatisticsAttributes) GetDownloadTimeOk() (*float32, bool)`

GetDownloadTimeOk returns a tuple with the DownloadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadTime

`func (o *StatisticsAttributes) SetDownloadTime(v float32)`

SetDownloadTime sets DownloadTime field to given value.

### HasDownloadTime

`func (o *StatisticsAttributes) HasDownloadTime() bool`

HasDownloadTime returns a boolean if a field has been set.

### GetTotalBytes

`func (o *StatisticsAttributes) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *StatisticsAttributes) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *StatisticsAttributes) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *StatisticsAttributes) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


