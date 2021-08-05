# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleType** | Pointer to [**ScheduleType**](ScheduleType.md) |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**WeekDay** | Pointer to [**DayOfWeek**](DayOfWeek.md) |  | [optional] 
**MonthDay** | Pointer to **int32** |  | [optional] 

## Methods

### NewSchedule

`func NewSchedule() *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleType

`func (o *Schedule) GetScheduleType() ScheduleType`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *Schedule) GetScheduleTypeOk() (*ScheduleType, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *Schedule) SetScheduleType(v ScheduleType)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *Schedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetTime

`func (o *Schedule) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Schedule) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Schedule) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *Schedule) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetWeekDay

`func (o *Schedule) GetWeekDay() DayOfWeek`

GetWeekDay returns the WeekDay field if non-nil, zero value otherwise.

### GetWeekDayOk

`func (o *Schedule) GetWeekDayOk() (*DayOfWeek, bool)`

GetWeekDayOk returns a tuple with the WeekDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekDay

`func (o *Schedule) SetWeekDay(v DayOfWeek)`

SetWeekDay sets WeekDay field to given value.

### HasWeekDay

`func (o *Schedule) HasWeekDay() bool`

HasWeekDay returns a boolean if a field has been set.

### GetMonthDay

`func (o *Schedule) GetMonthDay() int32`

GetMonthDay returns the MonthDay field if non-nil, zero value otherwise.

### GetMonthDayOk

`func (o *Schedule) GetMonthDayOk() (*int32, bool)`

GetMonthDayOk returns a tuple with the MonthDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthDay

`func (o *Schedule) SetMonthDay(v int32)`

SetMonthDay sets MonthDay field to given value.

### HasMonthDay

`func (o *Schedule) HasMonthDay() bool`

HasMonthDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


