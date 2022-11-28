# MonitorGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorGroupGuid** | Pointer to **string** | The unique identifier of this monitor group | [optional] 
**Description** | Pointer to **string** | The descriptive name of this probe group | [optional] 
**IsAll** | **bool** | Indicates whether this is the default group for all probes | 
**IsQuotaUnlimited** | Pointer to **bool** | Indicates whether the monitor quota is unlimited Only administrators can change this | [optional] 
**BasicMonitorQuota** | Pointer to **int32** | The basic monitor quota for the monitor group Only administrators can change this | [optional] 
**BrowserMonitorQuota** | Pointer to **int32** | The browser monitor quota for the monitor group Only administrators can change this | [optional] 
**TransactionMonitorQuota** | Pointer to **int32** | The transaction monitor quota for the monitor group Only administrators can change this | [optional] 
**ApiMonitorQuota** | Pointer to **int32** | The api monitor quota for the monitor group Only administrators can change this | [optional] 
**ClassicQuota** | Pointer to **int32** | The classic quota for the monitor group Only administrators can change this | [optional] 

## Methods

### NewMonitorGroup

`func NewMonitorGroup(isAll bool, ) *MonitorGroup`

NewMonitorGroup instantiates a new MonitorGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorGroupWithDefaults

`func NewMonitorGroupWithDefaults() *MonitorGroup`

NewMonitorGroupWithDefaults instantiates a new MonitorGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorGroupGuid

`func (o *MonitorGroup) GetMonitorGroupGuid() string`

GetMonitorGroupGuid returns the MonitorGroupGuid field if non-nil, zero value otherwise.

### GetMonitorGroupGuidOk

`func (o *MonitorGroup) GetMonitorGroupGuidOk() (*string, bool)`

GetMonitorGroupGuidOk returns a tuple with the MonitorGroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGroupGuid

`func (o *MonitorGroup) SetMonitorGroupGuid(v string)`

SetMonitorGroupGuid sets MonitorGroupGuid field to given value.

### HasMonitorGroupGuid

`func (o *MonitorGroup) HasMonitorGroupGuid() bool`

HasMonitorGroupGuid returns a boolean if a field has been set.

### GetDescription

`func (o *MonitorGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitorGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitorGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MonitorGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsAll

`func (o *MonitorGroup) GetIsAll() bool`

GetIsAll returns the IsAll field if non-nil, zero value otherwise.

### GetIsAllOk

`func (o *MonitorGroup) GetIsAllOk() (*bool, bool)`

GetIsAllOk returns a tuple with the IsAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAll

`func (o *MonitorGroup) SetIsAll(v bool)`

SetIsAll sets IsAll field to given value.


### GetIsQuotaUnlimited

`func (o *MonitorGroup) GetIsQuotaUnlimited() bool`

GetIsQuotaUnlimited returns the IsQuotaUnlimited field if non-nil, zero value otherwise.

### GetIsQuotaUnlimitedOk

`func (o *MonitorGroup) GetIsQuotaUnlimitedOk() (*bool, bool)`

GetIsQuotaUnlimitedOk returns a tuple with the IsQuotaUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuotaUnlimited

`func (o *MonitorGroup) SetIsQuotaUnlimited(v bool)`

SetIsQuotaUnlimited sets IsQuotaUnlimited field to given value.

### HasIsQuotaUnlimited

`func (o *MonitorGroup) HasIsQuotaUnlimited() bool`

HasIsQuotaUnlimited returns a boolean if a field has been set.

### GetBasicMonitorQuota

`func (o *MonitorGroup) GetBasicMonitorQuota() int32`

GetBasicMonitorQuota returns the BasicMonitorQuota field if non-nil, zero value otherwise.

### GetBasicMonitorQuotaOk

`func (o *MonitorGroup) GetBasicMonitorQuotaOk() (*int32, bool)`

GetBasicMonitorQuotaOk returns a tuple with the BasicMonitorQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicMonitorQuota

`func (o *MonitorGroup) SetBasicMonitorQuota(v int32)`

SetBasicMonitorQuota sets BasicMonitorQuota field to given value.

### HasBasicMonitorQuota

`func (o *MonitorGroup) HasBasicMonitorQuota() bool`

HasBasicMonitorQuota returns a boolean if a field has been set.

### GetBrowserMonitorQuota

`func (o *MonitorGroup) GetBrowserMonitorQuota() int32`

GetBrowserMonitorQuota returns the BrowserMonitorQuota field if non-nil, zero value otherwise.

### GetBrowserMonitorQuotaOk

`func (o *MonitorGroup) GetBrowserMonitorQuotaOk() (*int32, bool)`

GetBrowserMonitorQuotaOk returns a tuple with the BrowserMonitorQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMonitorQuota

`func (o *MonitorGroup) SetBrowserMonitorQuota(v int32)`

SetBrowserMonitorQuota sets BrowserMonitorQuota field to given value.

### HasBrowserMonitorQuota

`func (o *MonitorGroup) HasBrowserMonitorQuota() bool`

HasBrowserMonitorQuota returns a boolean if a field has been set.

### GetTransactionMonitorQuota

`func (o *MonitorGroup) GetTransactionMonitorQuota() int32`

GetTransactionMonitorQuota returns the TransactionMonitorQuota field if non-nil, zero value otherwise.

### GetTransactionMonitorQuotaOk

`func (o *MonitorGroup) GetTransactionMonitorQuotaOk() (*int32, bool)`

GetTransactionMonitorQuotaOk returns a tuple with the TransactionMonitorQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionMonitorQuota

`func (o *MonitorGroup) SetTransactionMonitorQuota(v int32)`

SetTransactionMonitorQuota sets TransactionMonitorQuota field to given value.

### HasTransactionMonitorQuota

`func (o *MonitorGroup) HasTransactionMonitorQuota() bool`

HasTransactionMonitorQuota returns a boolean if a field has been set.

### GetApiMonitorQuota

`func (o *MonitorGroup) GetApiMonitorQuota() int32`

GetApiMonitorQuota returns the ApiMonitorQuota field if non-nil, zero value otherwise.

### GetApiMonitorQuotaOk

`func (o *MonitorGroup) GetApiMonitorQuotaOk() (*int32, bool)`

GetApiMonitorQuotaOk returns a tuple with the ApiMonitorQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMonitorQuota

`func (o *MonitorGroup) SetApiMonitorQuota(v int32)`

SetApiMonitorQuota sets ApiMonitorQuota field to given value.

### HasApiMonitorQuota

`func (o *MonitorGroup) HasApiMonitorQuota() bool`

HasApiMonitorQuota returns a boolean if a field has been set.

### GetClassicQuota

`func (o *MonitorGroup) GetClassicQuota() int32`

GetClassicQuota returns the ClassicQuota field if non-nil, zero value otherwise.

### GetClassicQuotaOk

`func (o *MonitorGroup) GetClassicQuotaOk() (*int32, bool)`

GetClassicQuotaOk returns a tuple with the ClassicQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassicQuota

`func (o *MonitorGroup) SetClassicQuota(v int32)`

SetClassicQuota sets ClassicQuota field to given value.

### HasClassicQuota

`func (o *MonitorGroup) HasClassicQuota() bool`

HasClassicQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


