# AlertAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertType** | [**AlertType**](AlertType.md) | Alert type indicating if this was an Error or Ok alert. | 
**MonitorGuid** | **string** | The monitor identifier. | 
**Timestamp** | **time.Time** | Date/time stamp of the alert. | 
**FirstError** | **time.Time** | Date/time stamp of the first monitor check. | 
**MonitorCheckId** | **int64** | The Id of the monitor check that triggered this alert. | 
**FirstErrorMonitorCheckId** | **int64** | The Id of the first monitor check error that led to this alert. | 
**ErrorDescription** | Pointer to **string** | A text value that describes the error that was found, or OK if no error was found. | [optional] 
**ErrorMessage** | Pointer to **string** | Any additional error information, if available. | [optional] 
**IncidentKey** | Pointer to **string** | The incident key of this alert. | [optional] 

## Methods

### NewAlertAttributes

`func NewAlertAttributes(alertType AlertType, monitorGuid string, timestamp time.Time, firstError time.Time, monitorCheckId int64, firstErrorMonitorCheckId int64, ) *AlertAttributes`

NewAlertAttributes instantiates a new AlertAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertAttributesWithDefaults

`func NewAlertAttributesWithDefaults() *AlertAttributes`

NewAlertAttributesWithDefaults instantiates a new AlertAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertType

`func (o *AlertAttributes) GetAlertType() AlertType`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *AlertAttributes) GetAlertTypeOk() (*AlertType, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *AlertAttributes) SetAlertType(v AlertType)`

SetAlertType sets AlertType field to given value.


### GetMonitorGuid

`func (o *AlertAttributes) GetMonitorGuid() string`

GetMonitorGuid returns the MonitorGuid field if non-nil, zero value otherwise.

### GetMonitorGuidOk

`func (o *AlertAttributes) GetMonitorGuidOk() (*string, bool)`

GetMonitorGuidOk returns a tuple with the MonitorGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGuid

`func (o *AlertAttributes) SetMonitorGuid(v string)`

SetMonitorGuid sets MonitorGuid field to given value.


### GetTimestamp

`func (o *AlertAttributes) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AlertAttributes) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AlertAttributes) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetFirstError

`func (o *AlertAttributes) GetFirstError() time.Time`

GetFirstError returns the FirstError field if non-nil, zero value otherwise.

### GetFirstErrorOk

`func (o *AlertAttributes) GetFirstErrorOk() (*time.Time, bool)`

GetFirstErrorOk returns a tuple with the FirstError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstError

`func (o *AlertAttributes) SetFirstError(v time.Time)`

SetFirstError sets FirstError field to given value.


### GetMonitorCheckId

`func (o *AlertAttributes) GetMonitorCheckId() int64`

GetMonitorCheckId returns the MonitorCheckId field if non-nil, zero value otherwise.

### GetMonitorCheckIdOk

`func (o *AlertAttributes) GetMonitorCheckIdOk() (*int64, bool)`

GetMonitorCheckIdOk returns a tuple with the MonitorCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorCheckId

`func (o *AlertAttributes) SetMonitorCheckId(v int64)`

SetMonitorCheckId sets MonitorCheckId field to given value.


### GetFirstErrorMonitorCheckId

`func (o *AlertAttributes) GetFirstErrorMonitorCheckId() int64`

GetFirstErrorMonitorCheckId returns the FirstErrorMonitorCheckId field if non-nil, zero value otherwise.

### GetFirstErrorMonitorCheckIdOk

`func (o *AlertAttributes) GetFirstErrorMonitorCheckIdOk() (*int64, bool)`

GetFirstErrorMonitorCheckIdOk returns a tuple with the FirstErrorMonitorCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstErrorMonitorCheckId

`func (o *AlertAttributes) SetFirstErrorMonitorCheckId(v int64)`

SetFirstErrorMonitorCheckId sets FirstErrorMonitorCheckId field to given value.


### GetErrorDescription

`func (o *AlertAttributes) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *AlertAttributes) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *AlertAttributes) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *AlertAttributes) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AlertAttributes) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AlertAttributes) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AlertAttributes) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AlertAttributes) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetIncidentKey

`func (o *AlertAttributes) GetIncidentKey() string`

GetIncidentKey returns the IncidentKey field if non-nil, zero value otherwise.

### GetIncidentKeyOk

`func (o *AlertAttributes) GetIncidentKeyOk() (*string, bool)`

GetIncidentKeyOk returns a tuple with the IncidentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentKey

`func (o *AlertAttributes) SetIncidentKey(v string)`

SetIncidentKey sets IncidentKey field to given value.

### HasIncidentKey

`func (o *AlertAttributes) HasIncidentKey() bool`

HasIncidentKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


