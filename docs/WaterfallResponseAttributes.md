# WaterfallResponseAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageLoadMetrics** | Pointer to [**WaterfallInfoPageLoadMetrics**](WaterfallInfoPageLoadMetrics.md) |  | [optional] 
**W3CNavigationTiming** | Pointer to [**WaterfallInfoW3CNavigationTiming**](WaterfallInfoW3CNavigationTiming.md) |  | [optional] 
**Elements** | Pointer to [**[]PageElement**](PageElement.md) | Elements in the waterfall | [optional] 
**DomainGroupNames** | Pointer to **[]string** | List of domain groups used in the waterfall elements | [optional] 

## Methods

### NewWaterfallResponseAttributes

`func NewWaterfallResponseAttributes() *WaterfallResponseAttributes`

NewWaterfallResponseAttributes instantiates a new WaterfallResponseAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaterfallResponseAttributesWithDefaults

`func NewWaterfallResponseAttributesWithDefaults() *WaterfallResponseAttributes`

NewWaterfallResponseAttributesWithDefaults instantiates a new WaterfallResponseAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageLoadMetrics

`func (o *WaterfallResponseAttributes) GetPageLoadMetrics() WaterfallInfoPageLoadMetrics`

GetPageLoadMetrics returns the PageLoadMetrics field if non-nil, zero value otherwise.

### GetPageLoadMetricsOk

`func (o *WaterfallResponseAttributes) GetPageLoadMetricsOk() (*WaterfallInfoPageLoadMetrics, bool)`

GetPageLoadMetricsOk returns a tuple with the PageLoadMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLoadMetrics

`func (o *WaterfallResponseAttributes) SetPageLoadMetrics(v WaterfallInfoPageLoadMetrics)`

SetPageLoadMetrics sets PageLoadMetrics field to given value.

### HasPageLoadMetrics

`func (o *WaterfallResponseAttributes) HasPageLoadMetrics() bool`

HasPageLoadMetrics returns a boolean if a field has been set.

### GetW3CNavigationTiming

`func (o *WaterfallResponseAttributes) GetW3CNavigationTiming() WaterfallInfoW3CNavigationTiming`

GetW3CNavigationTiming returns the W3CNavigationTiming field if non-nil, zero value otherwise.

### GetW3CNavigationTimingOk

`func (o *WaterfallResponseAttributes) GetW3CNavigationTimingOk() (*WaterfallInfoW3CNavigationTiming, bool)`

GetW3CNavigationTimingOk returns a tuple with the W3CNavigationTiming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW3CNavigationTiming

`func (o *WaterfallResponseAttributes) SetW3CNavigationTiming(v WaterfallInfoW3CNavigationTiming)`

SetW3CNavigationTiming sets W3CNavigationTiming field to given value.

### HasW3CNavigationTiming

`func (o *WaterfallResponseAttributes) HasW3CNavigationTiming() bool`

HasW3CNavigationTiming returns a boolean if a field has been set.

### GetElements

`func (o *WaterfallResponseAttributes) GetElements() []PageElement`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *WaterfallResponseAttributes) GetElementsOk() (*[]PageElement, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *WaterfallResponseAttributes) SetElements(v []PageElement)`

SetElements sets Elements field to given value.

### HasElements

`func (o *WaterfallResponseAttributes) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetDomainGroupNames

`func (o *WaterfallResponseAttributes) GetDomainGroupNames() []string`

GetDomainGroupNames returns the DomainGroupNames field if non-nil, zero value otherwise.

### GetDomainGroupNamesOk

`func (o *WaterfallResponseAttributes) GetDomainGroupNamesOk() (*[]string, bool)`

GetDomainGroupNamesOk returns a tuple with the DomainGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupNames

`func (o *WaterfallResponseAttributes) SetDomainGroupNames(v []string)`

SetDomainGroupNames sets DomainGroupNames field to given value.

### HasDomainGroupNames

`func (o *WaterfallResponseAttributes) HasDomainGroupNames() bool`

HasDomainGroupNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


