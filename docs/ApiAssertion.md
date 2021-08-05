# ApiAssertion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**ApiAssertionSourceType**](ApiAssertionSourceType.md) |  | 
**Property** | Pointer to **string** |  | [optional] 
**Comparison** | [**ApiComparisonType**](ApiComparisonType.md) |  | 
**TargetValue** | Pointer to **string** |  | [optional] 

## Methods

### NewApiAssertion

`func NewApiAssertion(source ApiAssertionSourceType, comparison ApiComparisonType, ) *ApiAssertion`

NewApiAssertion instantiates a new ApiAssertion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAssertionWithDefaults

`func NewApiAssertionWithDefaults() *ApiAssertion`

NewApiAssertionWithDefaults instantiates a new ApiAssertion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ApiAssertion) GetSource() ApiAssertionSourceType`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiAssertion) GetSourceOk() (*ApiAssertionSourceType, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiAssertion) SetSource(v ApiAssertionSourceType)`

SetSource sets Source field to given value.


### GetProperty

`func (o *ApiAssertion) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ApiAssertion) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ApiAssertion) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ApiAssertion) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetComparison

`func (o *ApiAssertion) GetComparison() ApiComparisonType`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *ApiAssertion) GetComparisonOk() (*ApiComparisonType, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *ApiAssertion) SetComparison(v ApiComparisonType)`

SetComparison sets Comparison field to given value.


### GetTargetValue

`func (o *ApiAssertion) GetTargetValue() string`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *ApiAssertion) GetTargetValueOk() (*string, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *ApiAssertion) SetTargetValue(v string)`

SetTargetValue sets TargetValue field to given value.

### HasTargetValue

`func (o *ApiAssertion) HasTargetValue() bool`

HasTargetValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


