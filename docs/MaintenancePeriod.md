# MaintenancePeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique ID of this maintenance period | 
**ScheduleMode** | [**ScheduleMode**](ScheduleMode.md) | The schedule mode (one time, daily, weekly, monthly) | 
**StartDateTime** | Pointer to **time.Time** | The start date/time for this schedule (for one-time schedules only) | [optional] 
**EndDateTime** | Pointer to **time.Time** | The end date/time for this maintenance period (for one-time maintenance periods only) | [optional] 
**WeekDay** | Pointer to [**DayOfWeek**](DayOfWeek.md) | The weekday for this maintenance period (for weekly maintenance periods only) | [optional] 
**MonthDay** | Pointer to **int32** | the month day for this maintenance period (for montly maintenance periods only) | [optional] 
**StartTime** | Pointer to **string** | The start time of this maintenance period | [optional] 
**EndTime** | Pointer to **string** | The end time of this maintenance period | [optional] 
**MaintenanceType** | [**MaintenanceTypes**](MaintenanceTypes.md) | Indicates whether, during the maintenance periods, only alerting will be disabled, or if the entire monitor will be stopped | 
**Description** | Pointer to **string** | The description for this maintenance period | [optional] 

## Methods

### NewMaintenancePeriod

`func NewMaintenancePeriod(id int32, scheduleMode ScheduleMode, maintenanceType MaintenanceTypes, ) *MaintenancePeriod`

NewMaintenancePeriod instantiates a new MaintenancePeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenancePeriodWithDefaults

`func NewMaintenancePeriodWithDefaults() *MaintenancePeriod`

NewMaintenancePeriodWithDefaults instantiates a new MaintenancePeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaintenancePeriod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintenancePeriod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintenancePeriod) SetId(v int32)`

SetId sets Id field to given value.


### GetScheduleMode

`func (o *MaintenancePeriod) GetScheduleMode() ScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *MaintenancePeriod) GetScheduleModeOk() (*ScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *MaintenancePeriod) SetScheduleMode(v ScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.


### GetStartDateTime

`func (o *MaintenancePeriod) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MaintenancePeriod) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MaintenancePeriod) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MaintenancePeriod) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MaintenancePeriod) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MaintenancePeriod) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MaintenancePeriod) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MaintenancePeriod) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetWeekDay

`func (o *MaintenancePeriod) GetWeekDay() DayOfWeek`

GetWeekDay returns the WeekDay field if non-nil, zero value otherwise.

### GetWeekDayOk

`func (o *MaintenancePeriod) GetWeekDayOk() (*DayOfWeek, bool)`

GetWeekDayOk returns a tuple with the WeekDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekDay

`func (o *MaintenancePeriod) SetWeekDay(v DayOfWeek)`

SetWeekDay sets WeekDay field to given value.

### HasWeekDay

`func (o *MaintenancePeriod) HasWeekDay() bool`

HasWeekDay returns a boolean if a field has been set.

### GetMonthDay

`func (o *MaintenancePeriod) GetMonthDay() int32`

GetMonthDay returns the MonthDay field if non-nil, zero value otherwise.

### GetMonthDayOk

`func (o *MaintenancePeriod) GetMonthDayOk() (*int32, bool)`

GetMonthDayOk returns a tuple with the MonthDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthDay

`func (o *MaintenancePeriod) SetMonthDay(v int32)`

SetMonthDay sets MonthDay field to given value.

### HasMonthDay

`func (o *MaintenancePeriod) HasMonthDay() bool`

HasMonthDay returns a boolean if a field has been set.

### GetStartTime

`func (o *MaintenancePeriod) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MaintenancePeriod) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MaintenancePeriod) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MaintenancePeriod) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *MaintenancePeriod) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MaintenancePeriod) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MaintenancePeriod) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MaintenancePeriod) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMaintenanceType

`func (o *MaintenancePeriod) GetMaintenanceType() MaintenanceTypes`

GetMaintenanceType returns the MaintenanceType field if non-nil, zero value otherwise.

### GetMaintenanceTypeOk

`func (o *MaintenancePeriod) GetMaintenanceTypeOk() (*MaintenanceTypes, bool)`

GetMaintenanceTypeOk returns a tuple with the MaintenanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceType

`func (o *MaintenancePeriod) SetMaintenanceType(v MaintenanceTypes)`

SetMaintenanceType sets MaintenanceType field to given value.


### GetDescription

`func (o *MaintenancePeriod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MaintenancePeriod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MaintenancePeriod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MaintenancePeriod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


