# AccountStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**MonitorQuota** | Pointer to [**MonitorQuota**](MonitorQuota.md) |  | [optional] 
**OperatorQuota** | Pointer to [**OperatorQuota**](OperatorQuota.md) |  | [optional] 
**RemainingMessageCredits** | Pointer to **int32** |  | [optional] 

## Methods

### NewAccountStatistics

`func NewAccountStatistics() *AccountStatistics`

NewAccountStatistics instantiates a new AccountStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStatisticsWithDefaults

`func NewAccountStatisticsWithDefaults() *AccountStatistics`

NewAccountStatisticsWithDefaults instantiates a new AccountStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountStatistics) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountStatistics) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountStatistics) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountStatistics) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetExpirationDate

`func (o *AccountStatistics) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *AccountStatistics) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *AccountStatistics) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *AccountStatistics) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetMonitorQuota

`func (o *AccountStatistics) GetMonitorQuota() MonitorQuota`

GetMonitorQuota returns the MonitorQuota field if non-nil, zero value otherwise.

### GetMonitorQuotaOk

`func (o *AccountStatistics) GetMonitorQuotaOk() (*MonitorQuota, bool)`

GetMonitorQuotaOk returns a tuple with the MonitorQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorQuota

`func (o *AccountStatistics) SetMonitorQuota(v MonitorQuota)`

SetMonitorQuota sets MonitorQuota field to given value.

### HasMonitorQuota

`func (o *AccountStatistics) HasMonitorQuota() bool`

HasMonitorQuota returns a boolean if a field has been set.

### GetOperatorQuota

`func (o *AccountStatistics) GetOperatorQuota() OperatorQuota`

GetOperatorQuota returns the OperatorQuota field if non-nil, zero value otherwise.

### GetOperatorQuotaOk

`func (o *AccountStatistics) GetOperatorQuotaOk() (*OperatorQuota, bool)`

GetOperatorQuotaOk returns a tuple with the OperatorQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorQuota

`func (o *AccountStatistics) SetOperatorQuota(v OperatorQuota)`

SetOperatorQuota sets OperatorQuota field to given value.

### HasOperatorQuota

`func (o *AccountStatistics) HasOperatorQuota() bool`

HasOperatorQuota returns a boolean if a field has been set.

### GetRemainingMessageCredits

`func (o *AccountStatistics) GetRemainingMessageCredits() int32`

GetRemainingMessageCredits returns the RemainingMessageCredits field if non-nil, zero value otherwise.

### GetRemainingMessageCreditsOk

`func (o *AccountStatistics) GetRemainingMessageCreditsOk() (*int32, bool)`

GetRemainingMessageCreditsOk returns a tuple with the RemainingMessageCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingMessageCredits

`func (o *AccountStatistics) SetRemainingMessageCredits(v int32)`

SetRemainingMessageCredits sets RemainingMessageCredits field to given value.

### HasRemainingMessageCredits

`func (o *AccountStatistics) HasRemainingMessageCredits() bool`

HasRemainingMessageCredits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


