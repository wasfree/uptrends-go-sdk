# PatternMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to **string** |  | [optional] 
**IsPositive** | **bool** |  | 
**DateTime** | Pointer to [**DateTimePatternMatch**](DateTimePatternMatch.md) |  | [optional] 

## Methods

### NewPatternMatch

`func NewPatternMatch(isPositive bool, ) *PatternMatch`

NewPatternMatch instantiates a new PatternMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatternMatchWithDefaults

`func NewPatternMatchWithDefaults() *PatternMatch`

NewPatternMatchWithDefaults instantiates a new PatternMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *PatternMatch) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *PatternMatch) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *PatternMatch) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *PatternMatch) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetIsPositive

`func (o *PatternMatch) GetIsPositive() bool`

GetIsPositive returns the IsPositive field if non-nil, zero value otherwise.

### GetIsPositiveOk

`func (o *PatternMatch) GetIsPositiveOk() (*bool, bool)`

GetIsPositiveOk returns a tuple with the IsPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPositive

`func (o *PatternMatch) SetIsPositive(v bool)`

SetIsPositive sets IsPositive field to given value.


### GetDateTime

`func (o *PatternMatch) GetDateTime() DateTimePatternMatch`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *PatternMatch) GetDateTimeOk() (*DateTimePatternMatch, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *PatternMatch) SetDateTime(v DateTimePatternMatch)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *PatternMatch) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


