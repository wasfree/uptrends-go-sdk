# SelectedPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedPeriodType** | Pointer to [**SelectedPeriodType**](SelectedPeriodType.md) | The type of period | [optional] 
**Start** | Pointer to **time.Time** | The start of a custom period (can&#39;t be used together with the SelectedPeriodType parameter) | [optional] 
**End** | Pointer to **time.Time** | The end of a custom period | [optional] 
**PresetPeriod** | Pointer to [**PresetPeriodType**](PresetPeriodType.md) | The requested time period. | [optional] 

## Methods

### NewSelectedPeriod

`func NewSelectedPeriod() *SelectedPeriod`

NewSelectedPeriod instantiates a new SelectedPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedPeriodWithDefaults

`func NewSelectedPeriodWithDefaults() *SelectedPeriod`

NewSelectedPeriodWithDefaults instantiates a new SelectedPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedPeriodType

`func (o *SelectedPeriod) GetSelectedPeriodType() SelectedPeriodType`

GetSelectedPeriodType returns the SelectedPeriodType field if non-nil, zero value otherwise.

### GetSelectedPeriodTypeOk

`func (o *SelectedPeriod) GetSelectedPeriodTypeOk() (*SelectedPeriodType, bool)`

GetSelectedPeriodTypeOk returns a tuple with the SelectedPeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPeriodType

`func (o *SelectedPeriod) SetSelectedPeriodType(v SelectedPeriodType)`

SetSelectedPeriodType sets SelectedPeriodType field to given value.

### HasSelectedPeriodType

`func (o *SelectedPeriod) HasSelectedPeriodType() bool`

HasSelectedPeriodType returns a boolean if a field has been set.

### GetStart

`func (o *SelectedPeriod) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SelectedPeriod) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SelectedPeriod) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *SelectedPeriod) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *SelectedPeriod) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SelectedPeriod) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SelectedPeriod) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SelectedPeriod) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetPresetPeriod

`func (o *SelectedPeriod) GetPresetPeriod() PresetPeriodType`

GetPresetPeriod returns the PresetPeriod field if non-nil, zero value otherwise.

### GetPresetPeriodOk

`func (o *SelectedPeriod) GetPresetPeriodOk() (*PresetPeriodType, bool)`

GetPresetPeriodOk returns a tuple with the PresetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetPeriod

`func (o *SelectedPeriod) SetPresetPeriod(v PresetPeriodType)`

SetPresetPeriod sets PresetPeriod field to given value.

### HasPresetPeriod

`func (o *SelectedPeriod) HasPresetPeriod() bool`

HasPresetPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


