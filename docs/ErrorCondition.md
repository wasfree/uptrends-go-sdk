# ErrorCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorConditionType** | [**ErrorConditionType**](ErrorConditionType.md) | Error condition type | 
**Value** | Pointer to **string** | The value of this error conditions (not used in all types) | [optional] 
**Percentage** | Pointer to **string** | The percentage of this error conditions (not used in all types) | [optional] 
**Level** | Pointer to [**ErrorConditionConsoleLevel**](ErrorConditionConsoleLevel.md) | The level of this error conditions (not used in all types) | [optional] 
**MatchType** | Pointer to [**ErrorConditionMatchType**](ErrorConditionMatchType.md) | The match type of this error conditions (not used in all types) | [optional] 
**Effect** | Pointer to [**ErrorConditionEffect**](ErrorConditionEffect.md) | The effect of this error conditions (not used in all types) | [optional] 

## Methods

### NewErrorCondition

`func NewErrorCondition(errorConditionType ErrorConditionType, ) *ErrorCondition`

NewErrorCondition instantiates a new ErrorCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorConditionWithDefaults

`func NewErrorConditionWithDefaults() *ErrorCondition`

NewErrorConditionWithDefaults instantiates a new ErrorCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorConditionType

`func (o *ErrorCondition) GetErrorConditionType() ErrorConditionType`

GetErrorConditionType returns the ErrorConditionType field if non-nil, zero value otherwise.

### GetErrorConditionTypeOk

`func (o *ErrorCondition) GetErrorConditionTypeOk() (*ErrorConditionType, bool)`

GetErrorConditionTypeOk returns a tuple with the ErrorConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorConditionType

`func (o *ErrorCondition) SetErrorConditionType(v ErrorConditionType)`

SetErrorConditionType sets ErrorConditionType field to given value.


### GetValue

`func (o *ErrorCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ErrorCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercentage

`func (o *ErrorCondition) GetPercentage() string`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ErrorCondition) GetPercentageOk() (*string, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ErrorCondition) SetPercentage(v string)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ErrorCondition) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetLevel

`func (o *ErrorCondition) GetLevel() ErrorConditionConsoleLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ErrorCondition) GetLevelOk() (*ErrorConditionConsoleLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ErrorCondition) SetLevel(v ErrorConditionConsoleLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ErrorCondition) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMatchType

`func (o *ErrorCondition) GetMatchType() ErrorConditionMatchType`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *ErrorCondition) GetMatchTypeOk() (*ErrorConditionMatchType, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *ErrorCondition) SetMatchType(v ErrorConditionMatchType)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *ErrorCondition) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetEffect

`func (o *ErrorCondition) GetEffect() ErrorConditionEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *ErrorCondition) GetEffectOk() (*ErrorConditionEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *ErrorCondition) SetEffect(v ErrorConditionEffect)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *ErrorCondition) HasEffect() bool`

HasEffect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


