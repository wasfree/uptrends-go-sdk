# W3CNavigationTiming

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestStart** | **int32** | Equal to the requestStart as defined by the W3C.  It is a timestamp indicating when the browser starts requesting the resource from the webserver after the DNS lookup and TCP connection. | 
**TimeToFirstByte** | **int32** | Equal to the difference between requestStart and responseStart as defined by the W3C.  In short, it&#39;s the time between when the first request was sent from browser to server, and when the first bytes of the following response were received by the browser. | 
**DomInteractive** | **int32** | Equal to domInteractive as defined by W3C.  It is a timestamp, indicating the document readiness is set to &#39;interactive&#39;, to indicate that the browser has stopped parsing the page and the user can start interacting with it.  Resources such as scripts, images, stylesheets, or frames may still be loading. | 
**DomComplete** | **int32** | Equal to the domComplete as defined by W3C.  It is a timestamp, indicating when the main document has been parsed, the DOM has been fully loaded, and the page readiness is set to &#39;complete&#39;. | 
**LoadEvent** | **int32** | Equal to loadEventEnd as defined by W3C.  It is a timestamp, indicating when the load event of the current document has completed, including all dependent resources such as stylesheets and images. | 

## Methods

### NewW3CNavigationTiming

`func NewW3CNavigationTiming(requestStart int32, timeToFirstByte int32, domInteractive int32, domComplete int32, loadEvent int32, ) *W3CNavigationTiming`

NewW3CNavigationTiming instantiates a new W3CNavigationTiming object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewW3CNavigationTimingWithDefaults

`func NewW3CNavigationTimingWithDefaults() *W3CNavigationTiming`

NewW3CNavigationTimingWithDefaults instantiates a new W3CNavigationTiming object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestStart

`func (o *W3CNavigationTiming) GetRequestStart() int32`

GetRequestStart returns the RequestStart field if non-nil, zero value otherwise.

### GetRequestStartOk

`func (o *W3CNavigationTiming) GetRequestStartOk() (*int32, bool)`

GetRequestStartOk returns a tuple with the RequestStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStart

`func (o *W3CNavigationTiming) SetRequestStart(v int32)`

SetRequestStart sets RequestStart field to given value.


### GetTimeToFirstByte

`func (o *W3CNavigationTiming) GetTimeToFirstByte() int32`

GetTimeToFirstByte returns the TimeToFirstByte field if non-nil, zero value otherwise.

### GetTimeToFirstByteOk

`func (o *W3CNavigationTiming) GetTimeToFirstByteOk() (*int32, bool)`

GetTimeToFirstByteOk returns a tuple with the TimeToFirstByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToFirstByte

`func (o *W3CNavigationTiming) SetTimeToFirstByte(v int32)`

SetTimeToFirstByte sets TimeToFirstByte field to given value.


### GetDomInteractive

`func (o *W3CNavigationTiming) GetDomInteractive() int32`

GetDomInteractive returns the DomInteractive field if non-nil, zero value otherwise.

### GetDomInteractiveOk

`func (o *W3CNavigationTiming) GetDomInteractiveOk() (*int32, bool)`

GetDomInteractiveOk returns a tuple with the DomInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomInteractive

`func (o *W3CNavigationTiming) SetDomInteractive(v int32)`

SetDomInteractive sets DomInteractive field to given value.


### GetDomComplete

`func (o *W3CNavigationTiming) GetDomComplete() int32`

GetDomComplete returns the DomComplete field if non-nil, zero value otherwise.

### GetDomCompleteOk

`func (o *W3CNavigationTiming) GetDomCompleteOk() (*int32, bool)`

GetDomCompleteOk returns a tuple with the DomComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomComplete

`func (o *W3CNavigationTiming) SetDomComplete(v int32)`

SetDomComplete sets DomComplete field to given value.


### GetLoadEvent

`func (o *W3CNavigationTiming) GetLoadEvent() int32`

GetLoadEvent returns the LoadEvent field if non-nil, zero value otherwise.

### GetLoadEventOk

`func (o *W3CNavigationTiming) GetLoadEventOk() (*int32, bool)`

GetLoadEventOk returns a tuple with the LoadEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadEvent

`func (o *W3CNavigationTiming) SetLoadEvent(v int32)`

SetLoadEvent sets LoadEvent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


