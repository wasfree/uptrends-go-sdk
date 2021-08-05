# OperatorDutySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique ID of this maintenance period | 
**ScheduleMode** | [**OperatorScheduleMode**](OperatorScheduleMode.md) | The schedule mode (one time, daily, weekly, monthly) | 
**StartDateTime** | Pointer to **time.Time** | The start date/time for this schedule (for one-time schedules only) | [optional] 
**EndDateTime** | Pointer to **time.Time** | The end date/time for this maintenance period (for one-time maintenance periods only) | [optional] 
**WeekDay** | Pointer to [**DayOfWeek**](DayOfWeek.md) | The weekday for this maintenance period (for weekly maintenance periods only) | [optional] 
**MonthDay** | Pointer to **int32** | the month day for this maintenance period (for montly maintenance periods only) | [optional] 
**StartTime** | Pointer to **string** | The start time of this maintenance period | [optional] 
**EndTime** | Pointer to **string** | The end time of this maintenance period | [optional] 

## Methods

### NewOperatorDutySchedule

`func NewOperatorDutySchedule(id int32, scheduleMode OperatorScheduleMode, ) *OperatorDutySchedule`

NewOperatorDutySchedule instantiates a new OperatorDutySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorDutyScheduleWithDefaults

`func NewOperatorDutyScheduleWithDefaults() *OperatorDutySchedule`

NewOperatorDutyScheduleWithDefaults instantiates a new OperatorDutySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OperatorDutySchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperatorDutySchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperatorDutySchedule) SetId(v int32)`

SetId sets Id field to given value.


### GetScheduleMode

`func (o *OperatorDutySchedule) GetScheduleMode() OperatorScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *OperatorDutySchedule) GetScheduleModeOk() (*OperatorScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *OperatorDutySchedule) SetScheduleMode(v OperatorScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.


### GetStartDateTime

`func (o *OperatorDutySchedule) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *OperatorDutySchedule) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *OperatorDutySchedule) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *OperatorDutySchedule) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *OperatorDutySchedule) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *OperatorDutySchedule) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *OperatorDutySchedule) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *OperatorDutySchedule) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetWeekDay

`func (o *OperatorDutySchedule) GetWeekDay() DayOfWeek`

GetWeekDay returns the WeekDay field if non-nil, zero value otherwise.

### GetWeekDayOk

`func (o *OperatorDutySchedule) GetWeekDayOk() (*DayOfWeek, bool)`

GetWeekDayOk returns a tuple with the WeekDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekDay

`func (o *OperatorDutySchedule) SetWeekDay(v DayOfWeek)`

SetWeekDay sets WeekDay field to given value.

### HasWeekDay

`func (o *OperatorDutySchedule) HasWeekDay() bool`

HasWeekDay returns a boolean if a field has been set.

### GetMonthDay

`func (o *OperatorDutySchedule) GetMonthDay() int32`

GetMonthDay returns the MonthDay field if non-nil, zero value otherwise.

### GetMonthDayOk

`func (o *OperatorDutySchedule) GetMonthDayOk() (*int32, bool)`

GetMonthDayOk returns a tuple with the MonthDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthDay

`func (o *OperatorDutySchedule) SetMonthDay(v int32)`

SetMonthDay sets MonthDay field to given value.

### HasMonthDay

`func (o *OperatorDutySchedule) HasMonthDay() bool`

HasMonthDay returns a boolean if a field has been set.

### GetStartTime

`func (o *OperatorDutySchedule) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *OperatorDutySchedule) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *OperatorDutySchedule) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *OperatorDutySchedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *OperatorDutySchedule) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *OperatorDutySchedule) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *OperatorDutySchedule) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *OperatorDutySchedule) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


