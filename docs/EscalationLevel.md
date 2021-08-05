# EscalationLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EscalationMode** | Pointer to [**EscalationMode**](EscalationMode.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ThresholdErrorCount** | Pointer to **int32** |  | [optional] 
**ThresholdMinutes** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**NumberOfReminders** | Pointer to **int32** |  | [optional] 
**ReminderDelay** | Pointer to **int32** |  | [optional] 
**IncludeTraceRoute** | Pointer to **bool** |  | [optional] 
**Hash** | Pointer to **string** | Hash corresponding with this escalation level. | [optional] 

## Methods

### NewEscalationLevel

`func NewEscalationLevel() *EscalationLevel`

NewEscalationLevel instantiates a new EscalationLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEscalationLevelWithDefaults

`func NewEscalationLevelWithDefaults() *EscalationLevel`

NewEscalationLevelWithDefaults instantiates a new EscalationLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEscalationMode

`func (o *EscalationLevel) GetEscalationMode() EscalationMode`

GetEscalationMode returns the EscalationMode field if non-nil, zero value otherwise.

### GetEscalationModeOk

`func (o *EscalationLevel) GetEscalationModeOk() (*EscalationMode, bool)`

GetEscalationModeOk returns a tuple with the EscalationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalationMode

`func (o *EscalationLevel) SetEscalationMode(v EscalationMode)`

SetEscalationMode sets EscalationMode field to given value.

### HasEscalationMode

`func (o *EscalationLevel) HasEscalationMode() bool`

HasEscalationMode returns a boolean if a field has been set.

### GetId

`func (o *EscalationLevel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EscalationLevel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EscalationLevel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EscalationLevel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetThresholdErrorCount

`func (o *EscalationLevel) GetThresholdErrorCount() int32`

GetThresholdErrorCount returns the ThresholdErrorCount field if non-nil, zero value otherwise.

### GetThresholdErrorCountOk

`func (o *EscalationLevel) GetThresholdErrorCountOk() (*int32, bool)`

GetThresholdErrorCountOk returns a tuple with the ThresholdErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdErrorCount

`func (o *EscalationLevel) SetThresholdErrorCount(v int32)`

SetThresholdErrorCount sets ThresholdErrorCount field to given value.

### HasThresholdErrorCount

`func (o *EscalationLevel) HasThresholdErrorCount() bool`

HasThresholdErrorCount returns a boolean if a field has been set.

### GetThresholdMinutes

`func (o *EscalationLevel) GetThresholdMinutes() int32`

GetThresholdMinutes returns the ThresholdMinutes field if non-nil, zero value otherwise.

### GetThresholdMinutesOk

`func (o *EscalationLevel) GetThresholdMinutesOk() (*int32, bool)`

GetThresholdMinutesOk returns a tuple with the ThresholdMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMinutes

`func (o *EscalationLevel) SetThresholdMinutes(v int32)`

SetThresholdMinutes sets ThresholdMinutes field to given value.

### HasThresholdMinutes

`func (o *EscalationLevel) HasThresholdMinutes() bool`

HasThresholdMinutes returns a boolean if a field has been set.

### GetIsActive

`func (o *EscalationLevel) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *EscalationLevel) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *EscalationLevel) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *EscalationLevel) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetMessage

`func (o *EscalationLevel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EscalationLevel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EscalationLevel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EscalationLevel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNumberOfReminders

`func (o *EscalationLevel) GetNumberOfReminders() int32`

GetNumberOfReminders returns the NumberOfReminders field if non-nil, zero value otherwise.

### GetNumberOfRemindersOk

`func (o *EscalationLevel) GetNumberOfRemindersOk() (*int32, bool)`

GetNumberOfRemindersOk returns a tuple with the NumberOfReminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfReminders

`func (o *EscalationLevel) SetNumberOfReminders(v int32)`

SetNumberOfReminders sets NumberOfReminders field to given value.

### HasNumberOfReminders

`func (o *EscalationLevel) HasNumberOfReminders() bool`

HasNumberOfReminders returns a boolean if a field has been set.

### GetReminderDelay

`func (o *EscalationLevel) GetReminderDelay() int32`

GetReminderDelay returns the ReminderDelay field if non-nil, zero value otherwise.

### GetReminderDelayOk

`func (o *EscalationLevel) GetReminderDelayOk() (*int32, bool)`

GetReminderDelayOk returns a tuple with the ReminderDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderDelay

`func (o *EscalationLevel) SetReminderDelay(v int32)`

SetReminderDelay sets ReminderDelay field to given value.

### HasReminderDelay

`func (o *EscalationLevel) HasReminderDelay() bool`

HasReminderDelay returns a boolean if a field has been set.

### GetIncludeTraceRoute

`func (o *EscalationLevel) GetIncludeTraceRoute() bool`

GetIncludeTraceRoute returns the IncludeTraceRoute field if non-nil, zero value otherwise.

### GetIncludeTraceRouteOk

`func (o *EscalationLevel) GetIncludeTraceRouteOk() (*bool, bool)`

GetIncludeTraceRouteOk returns a tuple with the IncludeTraceRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTraceRoute

`func (o *EscalationLevel) SetIncludeTraceRoute(v bool)`

SetIncludeTraceRoute sets IncludeTraceRoute field to given value.

### HasIncludeTraceRoute

`func (o *EscalationLevel) HasIncludeTraceRoute() bool`

HasIncludeTraceRoute returns a boolean if a field has been set.

### GetHash

`func (o *EscalationLevel) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *EscalationLevel) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *EscalationLevel) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *EscalationLevel) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


