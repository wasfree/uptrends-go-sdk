# PageElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Index of the element in the waterfall context | 
**StartTime** | **int32** | Starting time in milliseconds | 
**QueueTime** | **int32** | Number of milliseconds this element has been queueing, when appropriate. | 
**ResolveTime** | **int32** | Number of milliseconds needed to perform the DNS query for this element, when appropriate. | 
**ConnectTime** | **int32** | Number of milliseconds needed to establish a connection. | 
**StaleTime** | **int32** | Number of milliseconds the connection was stale | 
**HttpsHandshakeTime** | **int32** | Number of milliseconds needed for the HTTPS handshake | 
**SendTime** | **int32** | Number of milliseconds it took to send data | 
**WaitTime** | **int32** | Number of milliseconds the connection was in waiting state | 
**ReceiveTime** | **int32** | Number of milliseconds it took to retrieve the data | 
**TimeoutTime** | **int32** | Number of milliseconds the connection was timed-out  | 
**TotalTime** | **int32** | Total number of milliseconds it took for the connection to complete | 
**HttpStatusCode** | **int32** | The Http status code | 
**Url** | Pointer to **string** | The requested resource url | [optional] 
**TotalBytes** | **int32** | Total number of downloaded bytes | 
**ElementType** | Pointer to **string** | Requested resource element type, can be HTML, scripts, CSS, images, frames, objects, data and other | [optional] 
**RequestHeaders** | Pointer to **string** | The HTTP request headers used | [optional] 
**ResponseHeaders** | Pointer to **string** | The HTTP response headers retrieved | [optional] 
**ResolvedIpAddress** | Pointer to **map[string]interface{}** | The IP address that was found for the specified domain name as part of this monitor check. | [optional] 
**GroupIds** | Pointer to **[]int32** |  | [optional] 
**UrlIsBlocked** | **bool** | Was the Url excluded from waterfall (timing) data by the user? | 

## Methods

### NewPageElement

`func NewPageElement(index int32, startTime int32, queueTime int32, resolveTime int32, connectTime int32, staleTime int32, httpsHandshakeTime int32, sendTime int32, waitTime int32, receiveTime int32, timeoutTime int32, totalTime int32, httpStatusCode int32, totalBytes int32, urlIsBlocked bool, ) *PageElement`

NewPageElement instantiates a new PageElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageElementWithDefaults

`func NewPageElementWithDefaults() *PageElement`

NewPageElementWithDefaults instantiates a new PageElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *PageElement) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PageElement) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PageElement) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetStartTime

`func (o *PageElement) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PageElement) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PageElement) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetQueueTime

`func (o *PageElement) GetQueueTime() int32`

GetQueueTime returns the QueueTime field if non-nil, zero value otherwise.

### GetQueueTimeOk

`func (o *PageElement) GetQueueTimeOk() (*int32, bool)`

GetQueueTimeOk returns a tuple with the QueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTime

`func (o *PageElement) SetQueueTime(v int32)`

SetQueueTime sets QueueTime field to given value.


### GetResolveTime

`func (o *PageElement) GetResolveTime() int32`

GetResolveTime returns the ResolveTime field if non-nil, zero value otherwise.

### GetResolveTimeOk

`func (o *PageElement) GetResolveTimeOk() (*int32, bool)`

GetResolveTimeOk returns a tuple with the ResolveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveTime

`func (o *PageElement) SetResolveTime(v int32)`

SetResolveTime sets ResolveTime field to given value.


### GetConnectTime

`func (o *PageElement) GetConnectTime() int32`

GetConnectTime returns the ConnectTime field if non-nil, zero value otherwise.

### GetConnectTimeOk

`func (o *PageElement) GetConnectTimeOk() (*int32, bool)`

GetConnectTimeOk returns a tuple with the ConnectTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTime

`func (o *PageElement) SetConnectTime(v int32)`

SetConnectTime sets ConnectTime field to given value.


### GetStaleTime

`func (o *PageElement) GetStaleTime() int32`

GetStaleTime returns the StaleTime field if non-nil, zero value otherwise.

### GetStaleTimeOk

`func (o *PageElement) GetStaleTimeOk() (*int32, bool)`

GetStaleTimeOk returns a tuple with the StaleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleTime

`func (o *PageElement) SetStaleTime(v int32)`

SetStaleTime sets StaleTime field to given value.


### GetHttpsHandshakeTime

`func (o *PageElement) GetHttpsHandshakeTime() int32`

GetHttpsHandshakeTime returns the HttpsHandshakeTime field if non-nil, zero value otherwise.

### GetHttpsHandshakeTimeOk

`func (o *PageElement) GetHttpsHandshakeTimeOk() (*int32, bool)`

GetHttpsHandshakeTimeOk returns a tuple with the HttpsHandshakeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsHandshakeTime

`func (o *PageElement) SetHttpsHandshakeTime(v int32)`

SetHttpsHandshakeTime sets HttpsHandshakeTime field to given value.


### GetSendTime

`func (o *PageElement) GetSendTime() int32`

GetSendTime returns the SendTime field if non-nil, zero value otherwise.

### GetSendTimeOk

`func (o *PageElement) GetSendTimeOk() (*int32, bool)`

GetSendTimeOk returns a tuple with the SendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTime

`func (o *PageElement) SetSendTime(v int32)`

SetSendTime sets SendTime field to given value.


### GetWaitTime

`func (o *PageElement) GetWaitTime() int32`

GetWaitTime returns the WaitTime field if non-nil, zero value otherwise.

### GetWaitTimeOk

`func (o *PageElement) GetWaitTimeOk() (*int32, bool)`

GetWaitTimeOk returns a tuple with the WaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTime

`func (o *PageElement) SetWaitTime(v int32)`

SetWaitTime sets WaitTime field to given value.


### GetReceiveTime

`func (o *PageElement) GetReceiveTime() int32`

GetReceiveTime returns the ReceiveTime field if non-nil, zero value otherwise.

### GetReceiveTimeOk

`func (o *PageElement) GetReceiveTimeOk() (*int32, bool)`

GetReceiveTimeOk returns a tuple with the ReceiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveTime

`func (o *PageElement) SetReceiveTime(v int32)`

SetReceiveTime sets ReceiveTime field to given value.


### GetTimeoutTime

`func (o *PageElement) GetTimeoutTime() int32`

GetTimeoutTime returns the TimeoutTime field if non-nil, zero value otherwise.

### GetTimeoutTimeOk

`func (o *PageElement) GetTimeoutTimeOk() (*int32, bool)`

GetTimeoutTimeOk returns a tuple with the TimeoutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutTime

`func (o *PageElement) SetTimeoutTime(v int32)`

SetTimeoutTime sets TimeoutTime field to given value.


### GetTotalTime

`func (o *PageElement) GetTotalTime() int32`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *PageElement) GetTotalTimeOk() (*int32, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *PageElement) SetTotalTime(v int32)`

SetTotalTime sets TotalTime field to given value.


### GetHttpStatusCode

`func (o *PageElement) GetHttpStatusCode() int32`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *PageElement) GetHttpStatusCodeOk() (*int32, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *PageElement) SetHttpStatusCode(v int32)`

SetHttpStatusCode sets HttpStatusCode field to given value.


### GetUrl

`func (o *PageElement) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PageElement) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PageElement) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PageElement) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTotalBytes

`func (o *PageElement) GetTotalBytes() int32`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *PageElement) GetTotalBytesOk() (*int32, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *PageElement) SetTotalBytes(v int32)`

SetTotalBytes sets TotalBytes field to given value.


### GetElementType

`func (o *PageElement) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *PageElement) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *PageElement) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *PageElement) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *PageElement) GetRequestHeaders() string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *PageElement) GetRequestHeadersOk() (*string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *PageElement) SetRequestHeaders(v string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *PageElement) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *PageElement) GetResponseHeaders() string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *PageElement) GetResponseHeadersOk() (*string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *PageElement) SetResponseHeaders(v string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *PageElement) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetResolvedIpAddress

`func (o *PageElement) GetResolvedIpAddress() map[string]interface{}`

GetResolvedIpAddress returns the ResolvedIpAddress field if non-nil, zero value otherwise.

### GetResolvedIpAddressOk

`func (o *PageElement) GetResolvedIpAddressOk() (*map[string]interface{}, bool)`

GetResolvedIpAddressOk returns a tuple with the ResolvedIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedIpAddress

`func (o *PageElement) SetResolvedIpAddress(v map[string]interface{})`

SetResolvedIpAddress sets ResolvedIpAddress field to given value.

### HasResolvedIpAddress

`func (o *PageElement) HasResolvedIpAddress() bool`

HasResolvedIpAddress returns a boolean if a field has been set.

### GetGroupIds

`func (o *PageElement) GetGroupIds() []int32`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *PageElement) GetGroupIdsOk() (*[]int32, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *PageElement) SetGroupIds(v []int32)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *PageElement) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetUrlIsBlocked

`func (o *PageElement) GetUrlIsBlocked() bool`

GetUrlIsBlocked returns the UrlIsBlocked field if non-nil, zero value otherwise.

### GetUrlIsBlockedOk

`func (o *PageElement) GetUrlIsBlockedOk() (*bool, bool)`

GetUrlIsBlockedOk returns a tuple with the UrlIsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIsBlocked

`func (o *PageElement) SetUrlIsBlocked(v bool)`

SetUrlIsBlocked sets UrlIsBlocked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


