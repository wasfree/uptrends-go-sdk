# AssertionResultsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalAssertions** | **int32** |  | 
**PassedAssertions** | **int32** |  | 
**AssertionInfos** | Pointer to [**[]AssertionInfo**](AssertionInfo.md) |  | [optional] 

## Methods

### NewAssertionResultsInfo

`func NewAssertionResultsInfo(totalAssertions int32, passedAssertions int32, ) *AssertionResultsInfo`

NewAssertionResultsInfo instantiates a new AssertionResultsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssertionResultsInfoWithDefaults

`func NewAssertionResultsInfoWithDefaults() *AssertionResultsInfo`

NewAssertionResultsInfoWithDefaults instantiates a new AssertionResultsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalAssertions

`func (o *AssertionResultsInfo) GetTotalAssertions() int32`

GetTotalAssertions returns the TotalAssertions field if non-nil, zero value otherwise.

### GetTotalAssertionsOk

`func (o *AssertionResultsInfo) GetTotalAssertionsOk() (*int32, bool)`

GetTotalAssertionsOk returns a tuple with the TotalAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssertions

`func (o *AssertionResultsInfo) SetTotalAssertions(v int32)`

SetTotalAssertions sets TotalAssertions field to given value.


### GetPassedAssertions

`func (o *AssertionResultsInfo) GetPassedAssertions() int32`

GetPassedAssertions returns the PassedAssertions field if non-nil, zero value otherwise.

### GetPassedAssertionsOk

`func (o *AssertionResultsInfo) GetPassedAssertionsOk() (*int32, bool)`

GetPassedAssertionsOk returns a tuple with the PassedAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassedAssertions

`func (o *AssertionResultsInfo) SetPassedAssertions(v int32)`

SetPassedAssertions sets PassedAssertions field to given value.


### GetAssertionInfos

`func (o *AssertionResultsInfo) GetAssertionInfos() []AssertionInfo`

GetAssertionInfos returns the AssertionInfos field if non-nil, zero value otherwise.

### GetAssertionInfosOk

`func (o *AssertionResultsInfo) GetAssertionInfosOk() (*[]AssertionInfo, bool)`

GetAssertionInfosOk returns a tuple with the AssertionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionInfos

`func (o *AssertionResultsInfo) SetAssertionInfos(v []AssertionInfo)`

SetAssertionInfos sets AssertionInfos field to given value.

### HasAssertionInfos

`func (o *AssertionResultsInfo) HasAssertionInfos() bool`

HasAssertionInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


