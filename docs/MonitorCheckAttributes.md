# MonitorCheckAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorGuid** | **string** | Monitor identifier | 
**Timestamp** | **time.Time** | Date/time stamp of the check | 
**ErrorCode** | **int32** | The numeric Uptrends error code in case of an error result, or 0 in case of an OK result. | 
**TotalTime** | **float64** | The number of milliseconds needed to complete the monitor check. | 
**ResolveTime** | **float64** | The number of milliseconds needed to perform the DNS query for this check, when appropriate. | 
**ConnectionTime** | **float64** | The number of milliseconds needed to establish a connection. | 
**DownloadTime** | **float64** | The number of milliseconds needed to download the response data. | 
**TotalBytes** | Pointer to **int32** | The number of downloaded bytes for this check. | [optional] 
**ResolvedIpAddress** | Pointer to **string** | The IP address that was found for the specified domain name as part of this monitor check. | [optional] 
**ErrorLevel** | [**ErrorLevel**](ErrorLevel.md) | A value that represents the OK/Error state for this check: NoError if the result was OK, Unconfirmed if an error was found, Confirmed if an error was found as a double check, right after an Unconfirmed error. | 
**ErrorDescription** | Pointer to **string** | A text value that describes the error that was found, or OK if no error was found. | [optional] 
**ErrorMessage** | Pointer to **string** | Any additional error information, if available. | [optional] 
**StagingMode** | **bool** | Did the check come from a staging monitor? | 
**ServerId** | Pointer to **int32** | The Id of the Uptrends checkpoint server that performed this check. | [optional] 
**HttpStatusCode** | Pointer to **int32** | The HTTP status code returned (if applicable) | [optional] 
**IsPartialCheck** | **bool** | This is a partial concurrent measurement, part of a concurrent check | 
**IsConcurrentCheck** | Pointer to **bool** | Is this a master concurrent check record | [optional] 

## Methods

### NewMonitorCheckAttributes

`func NewMonitorCheckAttributes(monitorGuid string, timestamp time.Time, errorCode int32, totalTime float64, resolveTime float64, connectionTime float64, downloadTime float64, errorLevel ErrorLevel, stagingMode bool, isPartialCheck bool, ) *MonitorCheckAttributes`

NewMonitorCheckAttributes instantiates a new MonitorCheckAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorCheckAttributesWithDefaults

`func NewMonitorCheckAttributesWithDefaults() *MonitorCheckAttributes`

NewMonitorCheckAttributesWithDefaults instantiates a new MonitorCheckAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorGuid

`func (o *MonitorCheckAttributes) GetMonitorGuid() string`

GetMonitorGuid returns the MonitorGuid field if non-nil, zero value otherwise.

### GetMonitorGuidOk

`func (o *MonitorCheckAttributes) GetMonitorGuidOk() (*string, bool)`

GetMonitorGuidOk returns a tuple with the MonitorGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGuid

`func (o *MonitorCheckAttributes) SetMonitorGuid(v string)`

SetMonitorGuid sets MonitorGuid field to given value.


### GetTimestamp

`func (o *MonitorCheckAttributes) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MonitorCheckAttributes) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MonitorCheckAttributes) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetErrorCode

`func (o *MonitorCheckAttributes) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MonitorCheckAttributes) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MonitorCheckAttributes) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.


### GetTotalTime

`func (o *MonitorCheckAttributes) GetTotalTime() float64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *MonitorCheckAttributes) GetTotalTimeOk() (*float64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *MonitorCheckAttributes) SetTotalTime(v float64)`

SetTotalTime sets TotalTime field to given value.


### GetResolveTime

`func (o *MonitorCheckAttributes) GetResolveTime() float64`

GetResolveTime returns the ResolveTime field if non-nil, zero value otherwise.

### GetResolveTimeOk

`func (o *MonitorCheckAttributes) GetResolveTimeOk() (*float64, bool)`

GetResolveTimeOk returns a tuple with the ResolveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveTime

`func (o *MonitorCheckAttributes) SetResolveTime(v float64)`

SetResolveTime sets ResolveTime field to given value.


### GetConnectionTime

`func (o *MonitorCheckAttributes) GetConnectionTime() float64`

GetConnectionTime returns the ConnectionTime field if non-nil, zero value otherwise.

### GetConnectionTimeOk

`func (o *MonitorCheckAttributes) GetConnectionTimeOk() (*float64, bool)`

GetConnectionTimeOk returns a tuple with the ConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTime

`func (o *MonitorCheckAttributes) SetConnectionTime(v float64)`

SetConnectionTime sets ConnectionTime field to given value.


### GetDownloadTime

`func (o *MonitorCheckAttributes) GetDownloadTime() float64`

GetDownloadTime returns the DownloadTime field if non-nil, zero value otherwise.

### GetDownloadTimeOk

`func (o *MonitorCheckAttributes) GetDownloadTimeOk() (*float64, bool)`

GetDownloadTimeOk returns a tuple with the DownloadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadTime

`func (o *MonitorCheckAttributes) SetDownloadTime(v float64)`

SetDownloadTime sets DownloadTime field to given value.


### GetTotalBytes

`func (o *MonitorCheckAttributes) GetTotalBytes() int32`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *MonitorCheckAttributes) GetTotalBytesOk() (*int32, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *MonitorCheckAttributes) SetTotalBytes(v int32)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *MonitorCheckAttributes) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetResolvedIpAddress

`func (o *MonitorCheckAttributes) GetResolvedIpAddress() string`

GetResolvedIpAddress returns the ResolvedIpAddress field if non-nil, zero value otherwise.

### GetResolvedIpAddressOk

`func (o *MonitorCheckAttributes) GetResolvedIpAddressOk() (*string, bool)`

GetResolvedIpAddressOk returns a tuple with the ResolvedIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedIpAddress

`func (o *MonitorCheckAttributes) SetResolvedIpAddress(v string)`

SetResolvedIpAddress sets ResolvedIpAddress field to given value.

### HasResolvedIpAddress

`func (o *MonitorCheckAttributes) HasResolvedIpAddress() bool`

HasResolvedIpAddress returns a boolean if a field has been set.

### GetErrorLevel

`func (o *MonitorCheckAttributes) GetErrorLevel() ErrorLevel`

GetErrorLevel returns the ErrorLevel field if non-nil, zero value otherwise.

### GetErrorLevelOk

`func (o *MonitorCheckAttributes) GetErrorLevelOk() (*ErrorLevel, bool)`

GetErrorLevelOk returns a tuple with the ErrorLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLevel

`func (o *MonitorCheckAttributes) SetErrorLevel(v ErrorLevel)`

SetErrorLevel sets ErrorLevel field to given value.


### GetErrorDescription

`func (o *MonitorCheckAttributes) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *MonitorCheckAttributes) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *MonitorCheckAttributes) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *MonitorCheckAttributes) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorMessage

`func (o *MonitorCheckAttributes) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *MonitorCheckAttributes) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *MonitorCheckAttributes) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *MonitorCheckAttributes) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStagingMode

`func (o *MonitorCheckAttributes) GetStagingMode() bool`

GetStagingMode returns the StagingMode field if non-nil, zero value otherwise.

### GetStagingModeOk

`func (o *MonitorCheckAttributes) GetStagingModeOk() (*bool, bool)`

GetStagingModeOk returns a tuple with the StagingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingMode

`func (o *MonitorCheckAttributes) SetStagingMode(v bool)`

SetStagingMode sets StagingMode field to given value.


### GetServerId

`func (o *MonitorCheckAttributes) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *MonitorCheckAttributes) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *MonitorCheckAttributes) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *MonitorCheckAttributes) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetHttpStatusCode

`func (o *MonitorCheckAttributes) GetHttpStatusCode() int32`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *MonitorCheckAttributes) GetHttpStatusCodeOk() (*int32, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *MonitorCheckAttributes) SetHttpStatusCode(v int32)`

SetHttpStatusCode sets HttpStatusCode field to given value.

### HasHttpStatusCode

`func (o *MonitorCheckAttributes) HasHttpStatusCode() bool`

HasHttpStatusCode returns a boolean if a field has been set.

### GetIsPartialCheck

`func (o *MonitorCheckAttributes) GetIsPartialCheck() bool`

GetIsPartialCheck returns the IsPartialCheck field if non-nil, zero value otherwise.

### GetIsPartialCheckOk

`func (o *MonitorCheckAttributes) GetIsPartialCheckOk() (*bool, bool)`

GetIsPartialCheckOk returns a tuple with the IsPartialCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartialCheck

`func (o *MonitorCheckAttributes) SetIsPartialCheck(v bool)`

SetIsPartialCheck sets IsPartialCheck field to given value.


### GetIsConcurrentCheck

`func (o *MonitorCheckAttributes) GetIsConcurrentCheck() bool`

GetIsConcurrentCheck returns the IsConcurrentCheck field if non-nil, zero value otherwise.

### GetIsConcurrentCheckOk

`func (o *MonitorCheckAttributes) GetIsConcurrentCheckOk() (*bool, bool)`

GetIsConcurrentCheckOk returns a tuple with the IsConcurrentCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConcurrentCheck

`func (o *MonitorCheckAttributes) SetIsConcurrentCheck(v bool)`

SetIsConcurrentCheck sets IsConcurrentCheck field to given value.

### HasIsConcurrentCheck

`func (o *MonitorCheckAttributes) HasIsConcurrentCheck() bool`

HasIsConcurrentCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


