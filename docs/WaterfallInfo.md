# WaterfallInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | Pointer to [**[]PageElement**](PageElement.md) | Elements in the waterfall | [optional] 
**DomainGroupNames** | Pointer to **[]string** | List of domain groups used in the waterfall elements | [optional] 

## Methods

### NewWaterfallInfo

`func NewWaterfallInfo() *WaterfallInfo`

NewWaterfallInfo instantiates a new WaterfallInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaterfallInfoWithDefaults

`func NewWaterfallInfoWithDefaults() *WaterfallInfo`

NewWaterfallInfoWithDefaults instantiates a new WaterfallInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *WaterfallInfo) GetElements() []PageElement`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *WaterfallInfo) GetElementsOk() (*[]PageElement, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *WaterfallInfo) SetElements(v []PageElement)`

SetElements sets Elements field to given value.

### HasElements

`func (o *WaterfallInfo) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetDomainGroupNames

`func (o *WaterfallInfo) GetDomainGroupNames() []string`

GetDomainGroupNames returns the DomainGroupNames field if non-nil, zero value otherwise.

### GetDomainGroupNamesOk

`func (o *WaterfallInfo) GetDomainGroupNamesOk() (*[]string, bool)`

GetDomainGroupNamesOk returns a tuple with the DomainGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupNames

`func (o *WaterfallInfo) SetDomainGroupNames(v []string)`

SetDomainGroupNames sets DomainGroupNames field to given value.

### HasDomainGroupNames

`func (o *WaterfallInfo) HasDomainGroupNames() bool`

HasDomainGroupNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


