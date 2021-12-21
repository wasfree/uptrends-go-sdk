# PageLoadMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CumulativeLayoutShift** | Pointer to **float32** | Measures the sum total of all individual layout shift scores for every unexpected layout shift that occurs during the entire lifespan of the page. A score lower than 0.1 is considered a good score. Between 0.1 and 0.25 could use some improvement. 0.25 or higher is a poor score. https://web.dev/cls/?utm_source&#x3D;devtools  In case the data was inconclusive null is returned. | [optional] 
**FirstContentfulPaint** | Pointer to **int32** | Measures the time from when the page starts loading to when any of the page&#39;s content is rendered on screen. A score lower than 1000ms is considered a good score. https://web.dev/fcp/  In case the data was inconclusive null is returned. | [optional] 
**LargestContentfulPaint** | Pointer to **int32** | The render timestamp of the largest image or text block visible within the viewport. A score lower than 2500ms is good. Between 2500ms and 4000ms needs improvement. 4000ms or higher is poor https://web.dev/lcp/  In case the data was inconclusive null is returned. | [optional] 
**TotalBlockingTime** | Pointer to **int32** | The Total Blocking Time (TBT) metric measures the total amount of time between First Contentful Paint (FCP) and Time to Interactive (TTI) where the main thread was blocked for long enough to prevent input responsiveness.  https://web.dev/tbt/  In case the data was inconclusive null is returned. | [optional] 
**TimeToInteractive** | Pointer to **int32** | TTI measures how long it takes a page to become fully interactive.   https://web.dev/interactive/  In case the data was inconclusive null is returned. | [optional] 

## Methods

### NewPageLoadMetrics

`func NewPageLoadMetrics() *PageLoadMetrics`

NewPageLoadMetrics instantiates a new PageLoadMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageLoadMetricsWithDefaults

`func NewPageLoadMetricsWithDefaults() *PageLoadMetrics`

NewPageLoadMetricsWithDefaults instantiates a new PageLoadMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCumulativeLayoutShift

`func (o *PageLoadMetrics) GetCumulativeLayoutShift() float32`

GetCumulativeLayoutShift returns the CumulativeLayoutShift field if non-nil, zero value otherwise.

### GetCumulativeLayoutShiftOk

`func (o *PageLoadMetrics) GetCumulativeLayoutShiftOk() (*float32, bool)`

GetCumulativeLayoutShiftOk returns a tuple with the CumulativeLayoutShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeLayoutShift

`func (o *PageLoadMetrics) SetCumulativeLayoutShift(v float32)`

SetCumulativeLayoutShift sets CumulativeLayoutShift field to given value.

### HasCumulativeLayoutShift

`func (o *PageLoadMetrics) HasCumulativeLayoutShift() bool`

HasCumulativeLayoutShift returns a boolean if a field has been set.

### GetFirstContentfulPaint

`func (o *PageLoadMetrics) GetFirstContentfulPaint() int32`

GetFirstContentfulPaint returns the FirstContentfulPaint field if non-nil, zero value otherwise.

### GetFirstContentfulPaintOk

`func (o *PageLoadMetrics) GetFirstContentfulPaintOk() (*int32, bool)`

GetFirstContentfulPaintOk returns a tuple with the FirstContentfulPaint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstContentfulPaint

`func (o *PageLoadMetrics) SetFirstContentfulPaint(v int32)`

SetFirstContentfulPaint sets FirstContentfulPaint field to given value.

### HasFirstContentfulPaint

`func (o *PageLoadMetrics) HasFirstContentfulPaint() bool`

HasFirstContentfulPaint returns a boolean if a field has been set.

### GetLargestContentfulPaint

`func (o *PageLoadMetrics) GetLargestContentfulPaint() int32`

GetLargestContentfulPaint returns the LargestContentfulPaint field if non-nil, zero value otherwise.

### GetLargestContentfulPaintOk

`func (o *PageLoadMetrics) GetLargestContentfulPaintOk() (*int32, bool)`

GetLargestContentfulPaintOk returns a tuple with the LargestContentfulPaint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestContentfulPaint

`func (o *PageLoadMetrics) SetLargestContentfulPaint(v int32)`

SetLargestContentfulPaint sets LargestContentfulPaint field to given value.

### HasLargestContentfulPaint

`func (o *PageLoadMetrics) HasLargestContentfulPaint() bool`

HasLargestContentfulPaint returns a boolean if a field has been set.

### GetTotalBlockingTime

`func (o *PageLoadMetrics) GetTotalBlockingTime() int32`

GetTotalBlockingTime returns the TotalBlockingTime field if non-nil, zero value otherwise.

### GetTotalBlockingTimeOk

`func (o *PageLoadMetrics) GetTotalBlockingTimeOk() (*int32, bool)`

GetTotalBlockingTimeOk returns a tuple with the TotalBlockingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBlockingTime

`func (o *PageLoadMetrics) SetTotalBlockingTime(v int32)`

SetTotalBlockingTime sets TotalBlockingTime field to given value.

### HasTotalBlockingTime

`func (o *PageLoadMetrics) HasTotalBlockingTime() bool`

HasTotalBlockingTime returns a boolean if a field has been set.

### GetTimeToInteractive

`func (o *PageLoadMetrics) GetTimeToInteractive() int32`

GetTimeToInteractive returns the TimeToInteractive field if non-nil, zero value otherwise.

### GetTimeToInteractiveOk

`func (o *PageLoadMetrics) GetTimeToInteractiveOk() (*int32, bool)`

GetTimeToInteractiveOk returns a tuple with the TimeToInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToInteractive

`func (o *PageLoadMetrics) SetTimeToInteractive(v int32)`

SetTimeToInteractive sets TimeToInteractive field to given value.

### HasTimeToInteractive

`func (o *PageLoadMetrics) HasTimeToInteractive() bool`

HasTimeToInteractive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


