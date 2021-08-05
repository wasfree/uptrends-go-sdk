# EscalationLevelIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationGuid** | Pointer to **string** | The unique key of this Integration. | [optional] 
**ExtraEmailAddresses** | Pointer to **[]string** | Extra email addresses | [optional] 
**ExtraEmailAddressesSpecified** | Pointer to **bool** | Specified Extra EmailAddresses For Patch request | [optional] 
**StatusHubServiceList** | Pointer to [**[]IntegrationServiceMap**](IntegrationServiceMap.md) | StatusHub Service Mapping | [optional] 
**StatusHubServiceListSpecified** | Pointer to **bool** | Specified StatusHubServiceList For Patch request | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this Integration is active. | [optional] 
**Hash** | Pointer to **string** |  | [optional] 

## Methods

### NewEscalationLevelIntegration

`func NewEscalationLevelIntegration() *EscalationLevelIntegration`

NewEscalationLevelIntegration instantiates a new EscalationLevelIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEscalationLevelIntegrationWithDefaults

`func NewEscalationLevelIntegrationWithDefaults() *EscalationLevelIntegration`

NewEscalationLevelIntegrationWithDefaults instantiates a new EscalationLevelIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationGuid

`func (o *EscalationLevelIntegration) GetIntegrationGuid() string`

GetIntegrationGuid returns the IntegrationGuid field if non-nil, zero value otherwise.

### GetIntegrationGuidOk

`func (o *EscalationLevelIntegration) GetIntegrationGuidOk() (*string, bool)`

GetIntegrationGuidOk returns a tuple with the IntegrationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationGuid

`func (o *EscalationLevelIntegration) SetIntegrationGuid(v string)`

SetIntegrationGuid sets IntegrationGuid field to given value.

### HasIntegrationGuid

`func (o *EscalationLevelIntegration) HasIntegrationGuid() bool`

HasIntegrationGuid returns a boolean if a field has been set.

### GetExtraEmailAddresses

`func (o *EscalationLevelIntegration) GetExtraEmailAddresses() []string`

GetExtraEmailAddresses returns the ExtraEmailAddresses field if non-nil, zero value otherwise.

### GetExtraEmailAddressesOk

`func (o *EscalationLevelIntegration) GetExtraEmailAddressesOk() (*[]string, bool)`

GetExtraEmailAddressesOk returns a tuple with the ExtraEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraEmailAddresses

`func (o *EscalationLevelIntegration) SetExtraEmailAddresses(v []string)`

SetExtraEmailAddresses sets ExtraEmailAddresses field to given value.

### HasExtraEmailAddresses

`func (o *EscalationLevelIntegration) HasExtraEmailAddresses() bool`

HasExtraEmailAddresses returns a boolean if a field has been set.

### GetExtraEmailAddressesSpecified

`func (o *EscalationLevelIntegration) GetExtraEmailAddressesSpecified() bool`

GetExtraEmailAddressesSpecified returns the ExtraEmailAddressesSpecified field if non-nil, zero value otherwise.

### GetExtraEmailAddressesSpecifiedOk

`func (o *EscalationLevelIntegration) GetExtraEmailAddressesSpecifiedOk() (*bool, bool)`

GetExtraEmailAddressesSpecifiedOk returns a tuple with the ExtraEmailAddressesSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraEmailAddressesSpecified

`func (o *EscalationLevelIntegration) SetExtraEmailAddressesSpecified(v bool)`

SetExtraEmailAddressesSpecified sets ExtraEmailAddressesSpecified field to given value.

### HasExtraEmailAddressesSpecified

`func (o *EscalationLevelIntegration) HasExtraEmailAddressesSpecified() bool`

HasExtraEmailAddressesSpecified returns a boolean if a field has been set.

### GetStatusHubServiceList

`func (o *EscalationLevelIntegration) GetStatusHubServiceList() []IntegrationServiceMap`

GetStatusHubServiceList returns the StatusHubServiceList field if non-nil, zero value otherwise.

### GetStatusHubServiceListOk

`func (o *EscalationLevelIntegration) GetStatusHubServiceListOk() (*[]IntegrationServiceMap, bool)`

GetStatusHubServiceListOk returns a tuple with the StatusHubServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHubServiceList

`func (o *EscalationLevelIntegration) SetStatusHubServiceList(v []IntegrationServiceMap)`

SetStatusHubServiceList sets StatusHubServiceList field to given value.

### HasStatusHubServiceList

`func (o *EscalationLevelIntegration) HasStatusHubServiceList() bool`

HasStatusHubServiceList returns a boolean if a field has been set.

### GetStatusHubServiceListSpecified

`func (o *EscalationLevelIntegration) GetStatusHubServiceListSpecified() bool`

GetStatusHubServiceListSpecified returns the StatusHubServiceListSpecified field if non-nil, zero value otherwise.

### GetStatusHubServiceListSpecifiedOk

`func (o *EscalationLevelIntegration) GetStatusHubServiceListSpecifiedOk() (*bool, bool)`

GetStatusHubServiceListSpecifiedOk returns a tuple with the StatusHubServiceListSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHubServiceListSpecified

`func (o *EscalationLevelIntegration) SetStatusHubServiceListSpecified(v bool)`

SetStatusHubServiceListSpecified sets StatusHubServiceListSpecified field to given value.

### HasStatusHubServiceListSpecified

`func (o *EscalationLevelIntegration) HasStatusHubServiceListSpecified() bool`

HasStatusHubServiceListSpecified returns a boolean if a field has been set.

### GetIsActive

`func (o *EscalationLevelIntegration) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *EscalationLevelIntegration) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *EscalationLevelIntegration) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *EscalationLevelIntegration) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetHash

`func (o *EscalationLevelIntegration) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *EscalationLevelIntegration) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *EscalationLevelIntegration) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *EscalationLevelIntegration) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


