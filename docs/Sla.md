# Sla

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SlaGuid** | Pointer to **string** | The unique key of this sla. | [optional] 
**Hash** | Pointer to **string** | Hash corresponding with this sla. | [optional] 
**Description** | Pointer to **string** | The description name of this sla. | [optional] 
**UptimeErrorThreshold** | Pointer to **float64** |  | [optional] 
**UptimeWarningThreshold** | Pointer to **float64** |  | [optional] 
**LoadTimeErrorThreshold** | Pointer to **float64** |  | [optional] 
**OperatorReponseTimeThreshold** | Pointer to **int32** |  | [optional] 

## Methods

### NewSla

`func NewSla() *Sla`

NewSla instantiates a new Sla object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlaWithDefaults

`func NewSlaWithDefaults() *Sla`

NewSlaWithDefaults instantiates a new Sla object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlaGuid

`func (o *Sla) GetSlaGuid() string`

GetSlaGuid returns the SlaGuid field if non-nil, zero value otherwise.

### GetSlaGuidOk

`func (o *Sla) GetSlaGuidOk() (*string, bool)`

GetSlaGuidOk returns a tuple with the SlaGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaGuid

`func (o *Sla) SetSlaGuid(v string)`

SetSlaGuid sets SlaGuid field to given value.

### HasSlaGuid

`func (o *Sla) HasSlaGuid() bool`

HasSlaGuid returns a boolean if a field has been set.

### GetHash

`func (o *Sla) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Sla) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Sla) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Sla) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetDescription

`func (o *Sla) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Sla) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Sla) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Sla) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUptimeErrorThreshold

`func (o *Sla) GetUptimeErrorThreshold() float64`

GetUptimeErrorThreshold returns the UptimeErrorThreshold field if non-nil, zero value otherwise.

### GetUptimeErrorThresholdOk

`func (o *Sla) GetUptimeErrorThresholdOk() (*float64, bool)`

GetUptimeErrorThresholdOk returns a tuple with the UptimeErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeErrorThreshold

`func (o *Sla) SetUptimeErrorThreshold(v float64)`

SetUptimeErrorThreshold sets UptimeErrorThreshold field to given value.

### HasUptimeErrorThreshold

`func (o *Sla) HasUptimeErrorThreshold() bool`

HasUptimeErrorThreshold returns a boolean if a field has been set.

### GetUptimeWarningThreshold

`func (o *Sla) GetUptimeWarningThreshold() float64`

GetUptimeWarningThreshold returns the UptimeWarningThreshold field if non-nil, zero value otherwise.

### GetUptimeWarningThresholdOk

`func (o *Sla) GetUptimeWarningThresholdOk() (*float64, bool)`

GetUptimeWarningThresholdOk returns a tuple with the UptimeWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeWarningThreshold

`func (o *Sla) SetUptimeWarningThreshold(v float64)`

SetUptimeWarningThreshold sets UptimeWarningThreshold field to given value.

### HasUptimeWarningThreshold

`func (o *Sla) HasUptimeWarningThreshold() bool`

HasUptimeWarningThreshold returns a boolean if a field has been set.

### GetLoadTimeErrorThreshold

`func (o *Sla) GetLoadTimeErrorThreshold() float64`

GetLoadTimeErrorThreshold returns the LoadTimeErrorThreshold field if non-nil, zero value otherwise.

### GetLoadTimeErrorThresholdOk

`func (o *Sla) GetLoadTimeErrorThresholdOk() (*float64, bool)`

GetLoadTimeErrorThresholdOk returns a tuple with the LoadTimeErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadTimeErrorThreshold

`func (o *Sla) SetLoadTimeErrorThreshold(v float64)`

SetLoadTimeErrorThreshold sets LoadTimeErrorThreshold field to given value.

### HasLoadTimeErrorThreshold

`func (o *Sla) HasLoadTimeErrorThreshold() bool`

HasLoadTimeErrorThreshold returns a boolean if a field has been set.

### GetOperatorReponseTimeThreshold

`func (o *Sla) GetOperatorReponseTimeThreshold() int32`

GetOperatorReponseTimeThreshold returns the OperatorReponseTimeThreshold field if non-nil, zero value otherwise.

### GetOperatorReponseTimeThresholdOk

`func (o *Sla) GetOperatorReponseTimeThresholdOk() (*int32, bool)`

GetOperatorReponseTimeThresholdOk returns a tuple with the OperatorReponseTimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorReponseTimeThreshold

`func (o *Sla) SetOperatorReponseTimeThreshold(v int32)`

SetOperatorReponseTimeThreshold sets OperatorReponseTimeThreshold field to given value.

### HasOperatorReponseTimeThreshold

`func (o *Sla) HasOperatorReponseTimeThreshold() bool`

HasOperatorReponseTimeThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


