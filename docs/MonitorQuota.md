# MonitorQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicMonitors** | Pointer to **int32** |  | [optional] 
**BasicMonitorsInUse** | Pointer to **int32** |  | [optional] 
**BrowserMonitors** | Pointer to **int32** |  | [optional] 
**BrowserMonitorsInUse** | Pointer to **int32** |  | [optional] 
**ApiMonitoringCredits** | Pointer to **int32** |  | [optional] 
**ApiMonitoringCreditsInUse** | Pointer to **int32** |  | [optional] 
**TransactionCredits** | Pointer to **int32** |  | [optional] 
**TransactionCreditsInUse** | Pointer to **int32** |  | [optional] 
**Monitors** | Pointer to **int32** |  | [optional] 
**MonitorsInUse** | Pointer to **int32** |  | [optional] 

## Methods

### NewMonitorQuota

`func NewMonitorQuota() *MonitorQuota`

NewMonitorQuota instantiates a new MonitorQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorQuotaWithDefaults

`func NewMonitorQuotaWithDefaults() *MonitorQuota`

NewMonitorQuotaWithDefaults instantiates a new MonitorQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicMonitors

`func (o *MonitorQuota) GetBasicMonitors() int32`

GetBasicMonitors returns the BasicMonitors field if non-nil, zero value otherwise.

### GetBasicMonitorsOk

`func (o *MonitorQuota) GetBasicMonitorsOk() (*int32, bool)`

GetBasicMonitorsOk returns a tuple with the BasicMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicMonitors

`func (o *MonitorQuota) SetBasicMonitors(v int32)`

SetBasicMonitors sets BasicMonitors field to given value.

### HasBasicMonitors

`func (o *MonitorQuota) HasBasicMonitors() bool`

HasBasicMonitors returns a boolean if a field has been set.

### GetBasicMonitorsInUse

`func (o *MonitorQuota) GetBasicMonitorsInUse() int32`

GetBasicMonitorsInUse returns the BasicMonitorsInUse field if non-nil, zero value otherwise.

### GetBasicMonitorsInUseOk

`func (o *MonitorQuota) GetBasicMonitorsInUseOk() (*int32, bool)`

GetBasicMonitorsInUseOk returns a tuple with the BasicMonitorsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicMonitorsInUse

`func (o *MonitorQuota) SetBasicMonitorsInUse(v int32)`

SetBasicMonitorsInUse sets BasicMonitorsInUse field to given value.

### HasBasicMonitorsInUse

`func (o *MonitorQuota) HasBasicMonitorsInUse() bool`

HasBasicMonitorsInUse returns a boolean if a field has been set.

### GetBrowserMonitors

`func (o *MonitorQuota) GetBrowserMonitors() int32`

GetBrowserMonitors returns the BrowserMonitors field if non-nil, zero value otherwise.

### GetBrowserMonitorsOk

`func (o *MonitorQuota) GetBrowserMonitorsOk() (*int32, bool)`

GetBrowserMonitorsOk returns a tuple with the BrowserMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMonitors

`func (o *MonitorQuota) SetBrowserMonitors(v int32)`

SetBrowserMonitors sets BrowserMonitors field to given value.

### HasBrowserMonitors

`func (o *MonitorQuota) HasBrowserMonitors() bool`

HasBrowserMonitors returns a boolean if a field has been set.

### GetBrowserMonitorsInUse

`func (o *MonitorQuota) GetBrowserMonitorsInUse() int32`

GetBrowserMonitorsInUse returns the BrowserMonitorsInUse field if non-nil, zero value otherwise.

### GetBrowserMonitorsInUseOk

`func (o *MonitorQuota) GetBrowserMonitorsInUseOk() (*int32, bool)`

GetBrowserMonitorsInUseOk returns a tuple with the BrowserMonitorsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMonitorsInUse

`func (o *MonitorQuota) SetBrowserMonitorsInUse(v int32)`

SetBrowserMonitorsInUse sets BrowserMonitorsInUse field to given value.

### HasBrowserMonitorsInUse

`func (o *MonitorQuota) HasBrowserMonitorsInUse() bool`

HasBrowserMonitorsInUse returns a boolean if a field has been set.

### GetApiMonitoringCredits

`func (o *MonitorQuota) GetApiMonitoringCredits() int32`

GetApiMonitoringCredits returns the ApiMonitoringCredits field if non-nil, zero value otherwise.

### GetApiMonitoringCreditsOk

`func (o *MonitorQuota) GetApiMonitoringCreditsOk() (*int32, bool)`

GetApiMonitoringCreditsOk returns a tuple with the ApiMonitoringCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMonitoringCredits

`func (o *MonitorQuota) SetApiMonitoringCredits(v int32)`

SetApiMonitoringCredits sets ApiMonitoringCredits field to given value.

### HasApiMonitoringCredits

`func (o *MonitorQuota) HasApiMonitoringCredits() bool`

HasApiMonitoringCredits returns a boolean if a field has been set.

### GetApiMonitoringCreditsInUse

`func (o *MonitorQuota) GetApiMonitoringCreditsInUse() int32`

GetApiMonitoringCreditsInUse returns the ApiMonitoringCreditsInUse field if non-nil, zero value otherwise.

### GetApiMonitoringCreditsInUseOk

`func (o *MonitorQuota) GetApiMonitoringCreditsInUseOk() (*int32, bool)`

GetApiMonitoringCreditsInUseOk returns a tuple with the ApiMonitoringCreditsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMonitoringCreditsInUse

`func (o *MonitorQuota) SetApiMonitoringCreditsInUse(v int32)`

SetApiMonitoringCreditsInUse sets ApiMonitoringCreditsInUse field to given value.

### HasApiMonitoringCreditsInUse

`func (o *MonitorQuota) HasApiMonitoringCreditsInUse() bool`

HasApiMonitoringCreditsInUse returns a boolean if a field has been set.

### GetTransactionCredits

`func (o *MonitorQuota) GetTransactionCredits() int32`

GetTransactionCredits returns the TransactionCredits field if non-nil, zero value otherwise.

### GetTransactionCreditsOk

`func (o *MonitorQuota) GetTransactionCreditsOk() (*int32, bool)`

GetTransactionCreditsOk returns a tuple with the TransactionCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCredits

`func (o *MonitorQuota) SetTransactionCredits(v int32)`

SetTransactionCredits sets TransactionCredits field to given value.

### HasTransactionCredits

`func (o *MonitorQuota) HasTransactionCredits() bool`

HasTransactionCredits returns a boolean if a field has been set.

### GetTransactionCreditsInUse

`func (o *MonitorQuota) GetTransactionCreditsInUse() int32`

GetTransactionCreditsInUse returns the TransactionCreditsInUse field if non-nil, zero value otherwise.

### GetTransactionCreditsInUseOk

`func (o *MonitorQuota) GetTransactionCreditsInUseOk() (*int32, bool)`

GetTransactionCreditsInUseOk returns a tuple with the TransactionCreditsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCreditsInUse

`func (o *MonitorQuota) SetTransactionCreditsInUse(v int32)`

SetTransactionCreditsInUse sets TransactionCreditsInUse field to given value.

### HasTransactionCreditsInUse

`func (o *MonitorQuota) HasTransactionCreditsInUse() bool`

HasTransactionCreditsInUse returns a boolean if a field has been set.

### GetMonitors

`func (o *MonitorQuota) GetMonitors() int32`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *MonitorQuota) GetMonitorsOk() (*int32, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *MonitorQuota) SetMonitors(v int32)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *MonitorQuota) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetMonitorsInUse

`func (o *MonitorQuota) GetMonitorsInUse() int32`

GetMonitorsInUse returns the MonitorsInUse field if non-nil, zero value otherwise.

### GetMonitorsInUseOk

`func (o *MonitorQuota) GetMonitorsInUseOk() (*int32, bool)`

GetMonitorsInUseOk returns a tuple with the MonitorsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorsInUse

`func (o *MonitorQuota) SetMonitorsInUse(v int32)`

SetMonitorsInUse sets MonitorsInUse field to given value.

### HasMonitorsInUse

`func (o *MonitorQuota) HasMonitorsInUse() bool`

HasMonitorsInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


