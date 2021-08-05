# AssertionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | **bool** | Was the assertion completed? | 
**Passed** | **bool** | Did the assertion pass? | 
**Description** | Pointer to **string** | Description | [optional] 
**ExceptionDescription** | Pointer to **string** | Discription of the exception trown if applicable | [optional] 

## Methods

### NewAssertionInfo

`func NewAssertionInfo(completed bool, passed bool, ) *AssertionInfo`

NewAssertionInfo instantiates a new AssertionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssertionInfoWithDefaults

`func NewAssertionInfoWithDefaults() *AssertionInfo`

NewAssertionInfoWithDefaults instantiates a new AssertionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *AssertionInfo) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *AssertionInfo) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *AssertionInfo) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetPassed

`func (o *AssertionInfo) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *AssertionInfo) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *AssertionInfo) SetPassed(v bool)`

SetPassed sets Passed field to given value.


### GetDescription

`func (o *AssertionInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssertionInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssertionInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssertionInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExceptionDescription

`func (o *AssertionInfo) GetExceptionDescription() string`

GetExceptionDescription returns the ExceptionDescription field if non-nil, zero value otherwise.

### GetExceptionDescriptionOk

`func (o *AssertionInfo) GetExceptionDescriptionOk() (*string, bool)`

GetExceptionDescriptionOk returns a tuple with the ExceptionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionDescription

`func (o *AssertionInfo) SetExceptionDescription(v string)`

SetExceptionDescription sets ExceptionDescription field to given value.

### HasExceptionDescription

`func (o *AssertionInfo) HasExceptionDescription() bool`

HasExceptionDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


