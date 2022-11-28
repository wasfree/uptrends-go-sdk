# HttpEngineStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepName** | Pointer to **string** | The name of the step | [optional] 
**Url** | Pointer to **string** | Url the step was executed on | [optional] 
**HttpStatusCode** | Pointer to **string** | The HTTP status code returned | [optional] 
**HttpMethod** | Pointer to **string** | Http method used in this step | [optional] 
**HttpStatusDescription** | Pointer to **string** | Step description | [optional] 
**ResponseCompleted** | **bool** | Did the response complete? | 
**StepExecuted** | **bool** | Was this step executed? | 
**AssertionResultsInfo** | Pointer to [**HttpEngineStepAssertionResultsInfo**](HttpEngineStepAssertionResultsInfo.md) |  | [optional] 
**TotalTime** | **int64** | Number of milliseconds it took for this step to succeed | 
**ResponseHeaders** | Pointer to **string** | Response headers | [optional] 
**ResponseBody** | Pointer to **string** | Response body | [optional] 
**RequestHeaders** | Pointer to **string** | Request headers send | [optional] 
**RequestBody** | Pointer to **string** | Request body send | [optional] 
**ResolvedIpAddress** | Pointer to **string** | Resolved IP address | [optional] 

## Methods

### NewHttpEngineStep

`func NewHttpEngineStep(responseCompleted bool, stepExecuted bool, totalTime int64, ) *HttpEngineStep`

NewHttpEngineStep instantiates a new HttpEngineStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpEngineStepWithDefaults

`func NewHttpEngineStepWithDefaults() *HttpEngineStep`

NewHttpEngineStepWithDefaults instantiates a new HttpEngineStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepName

`func (o *HttpEngineStep) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *HttpEngineStep) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *HttpEngineStep) SetStepName(v string)`

SetStepName sets StepName field to given value.

### HasStepName

`func (o *HttpEngineStep) HasStepName() bool`

HasStepName returns a boolean if a field has been set.

### GetUrl

`func (o *HttpEngineStep) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpEngineStep) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpEngineStep) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HttpEngineStep) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHttpStatusCode

`func (o *HttpEngineStep) GetHttpStatusCode() string`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *HttpEngineStep) GetHttpStatusCodeOk() (*string, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *HttpEngineStep) SetHttpStatusCode(v string)`

SetHttpStatusCode sets HttpStatusCode field to given value.

### HasHttpStatusCode

`func (o *HttpEngineStep) HasHttpStatusCode() bool`

HasHttpStatusCode returns a boolean if a field has been set.

### GetHttpMethod

`func (o *HttpEngineStep) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *HttpEngineStep) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *HttpEngineStep) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *HttpEngineStep) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHttpStatusDescription

`func (o *HttpEngineStep) GetHttpStatusDescription() string`

GetHttpStatusDescription returns the HttpStatusDescription field if non-nil, zero value otherwise.

### GetHttpStatusDescriptionOk

`func (o *HttpEngineStep) GetHttpStatusDescriptionOk() (*string, bool)`

GetHttpStatusDescriptionOk returns a tuple with the HttpStatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusDescription

`func (o *HttpEngineStep) SetHttpStatusDescription(v string)`

SetHttpStatusDescription sets HttpStatusDescription field to given value.

### HasHttpStatusDescription

`func (o *HttpEngineStep) HasHttpStatusDescription() bool`

HasHttpStatusDescription returns a boolean if a field has been set.

### GetResponseCompleted

`func (o *HttpEngineStep) GetResponseCompleted() bool`

GetResponseCompleted returns the ResponseCompleted field if non-nil, zero value otherwise.

### GetResponseCompletedOk

`func (o *HttpEngineStep) GetResponseCompletedOk() (*bool, bool)`

GetResponseCompletedOk returns a tuple with the ResponseCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCompleted

`func (o *HttpEngineStep) SetResponseCompleted(v bool)`

SetResponseCompleted sets ResponseCompleted field to given value.


### GetStepExecuted

`func (o *HttpEngineStep) GetStepExecuted() bool`

GetStepExecuted returns the StepExecuted field if non-nil, zero value otherwise.

### GetStepExecutedOk

`func (o *HttpEngineStep) GetStepExecutedOk() (*bool, bool)`

GetStepExecutedOk returns a tuple with the StepExecuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepExecuted

`func (o *HttpEngineStep) SetStepExecuted(v bool)`

SetStepExecuted sets StepExecuted field to given value.


### GetAssertionResultsInfo

`func (o *HttpEngineStep) GetAssertionResultsInfo() HttpEngineStepAssertionResultsInfo`

GetAssertionResultsInfo returns the AssertionResultsInfo field if non-nil, zero value otherwise.

### GetAssertionResultsInfoOk

`func (o *HttpEngineStep) GetAssertionResultsInfoOk() (*HttpEngineStepAssertionResultsInfo, bool)`

GetAssertionResultsInfoOk returns a tuple with the AssertionResultsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionResultsInfo

`func (o *HttpEngineStep) SetAssertionResultsInfo(v HttpEngineStepAssertionResultsInfo)`

SetAssertionResultsInfo sets AssertionResultsInfo field to given value.

### HasAssertionResultsInfo

`func (o *HttpEngineStep) HasAssertionResultsInfo() bool`

HasAssertionResultsInfo returns a boolean if a field has been set.

### GetTotalTime

`func (o *HttpEngineStep) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *HttpEngineStep) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *HttpEngineStep) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.


### GetResponseHeaders

`func (o *HttpEngineStep) GetResponseHeaders() string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *HttpEngineStep) GetResponseHeadersOk() (*string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *HttpEngineStep) SetResponseHeaders(v string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *HttpEngineStep) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetResponseBody

`func (o *HttpEngineStep) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *HttpEngineStep) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *HttpEngineStep) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *HttpEngineStep) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *HttpEngineStep) GetRequestHeaders() string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *HttpEngineStep) GetRequestHeadersOk() (*string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *HttpEngineStep) SetRequestHeaders(v string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *HttpEngineStep) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetRequestBody

`func (o *HttpEngineStep) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *HttpEngineStep) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *HttpEngineStep) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *HttpEngineStep) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetResolvedIpAddress

`func (o *HttpEngineStep) GetResolvedIpAddress() string`

GetResolvedIpAddress returns the ResolvedIpAddress field if non-nil, zero value otherwise.

### GetResolvedIpAddressOk

`func (o *HttpEngineStep) GetResolvedIpAddressOk() (*string, bool)`

GetResolvedIpAddressOk returns a tuple with the ResolvedIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedIpAddress

`func (o *HttpEngineStep) SetResolvedIpAddress(v string)`

SetResolvedIpAddress sets ResolvedIpAddress field to given value.

### HasResolvedIpAddress

`func (o *HttpEngineStep) HasResolvedIpAddress() bool`

HasResolvedIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


