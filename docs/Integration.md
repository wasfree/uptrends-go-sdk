# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationGuid** | **string** | Guid of Integration in Alert Definition Escalation Level | 
**Name** | Pointer to **string** | Name of Integration in Alert Definition Escalation Level | [optional] 
**Type** | [**IntegrationTypeEnum**](IntegrationTypeEnum.md) | Type of Integration in Alert Definition Escalation Level | 
**ExtraEmailAddresses** | Pointer to **string** | Extra emailadresses for this integration (if type &#x3D;&#x3D; email) | [optional] 
**StatusHubServiceList** | Pointer to [**[]IntegrationServiceMap**](IntegrationServiceMap.md) | All statushubs for this integration (if type &#x3D;&#x3D; statushub) | [optional] 
**IntegrationServices** | Pointer to **[]string** | All integrations services. | [optional] 
**Hash** | Pointer to **string** |  | [optional] 

## Methods

### NewIntegration

`func NewIntegration(integrationGuid string, type_ IntegrationTypeEnum, ) *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationGuid

`func (o *Integration) GetIntegrationGuid() string`

GetIntegrationGuid returns the IntegrationGuid field if non-nil, zero value otherwise.

### GetIntegrationGuidOk

`func (o *Integration) GetIntegrationGuidOk() (*string, bool)`

GetIntegrationGuidOk returns a tuple with the IntegrationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationGuid

`func (o *Integration) SetIntegrationGuid(v string)`

SetIntegrationGuid sets IntegrationGuid field to given value.


### GetName

`func (o *Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Integration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Integration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Integration) GetType() IntegrationTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Integration) GetTypeOk() (*IntegrationTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Integration) SetType(v IntegrationTypeEnum)`

SetType sets Type field to given value.


### GetExtraEmailAddresses

`func (o *Integration) GetExtraEmailAddresses() string`

GetExtraEmailAddresses returns the ExtraEmailAddresses field if non-nil, zero value otherwise.

### GetExtraEmailAddressesOk

`func (o *Integration) GetExtraEmailAddressesOk() (*string, bool)`

GetExtraEmailAddressesOk returns a tuple with the ExtraEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraEmailAddresses

`func (o *Integration) SetExtraEmailAddresses(v string)`

SetExtraEmailAddresses sets ExtraEmailAddresses field to given value.

### HasExtraEmailAddresses

`func (o *Integration) HasExtraEmailAddresses() bool`

HasExtraEmailAddresses returns a boolean if a field has been set.

### GetStatusHubServiceList

`func (o *Integration) GetStatusHubServiceList() []IntegrationServiceMap`

GetStatusHubServiceList returns the StatusHubServiceList field if non-nil, zero value otherwise.

### GetStatusHubServiceListOk

`func (o *Integration) GetStatusHubServiceListOk() (*[]IntegrationServiceMap, bool)`

GetStatusHubServiceListOk returns a tuple with the StatusHubServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHubServiceList

`func (o *Integration) SetStatusHubServiceList(v []IntegrationServiceMap)`

SetStatusHubServiceList sets StatusHubServiceList field to given value.

### HasStatusHubServiceList

`func (o *Integration) HasStatusHubServiceList() bool`

HasStatusHubServiceList returns a boolean if a field has been set.

### GetIntegrationServices

`func (o *Integration) GetIntegrationServices() []string`

GetIntegrationServices returns the IntegrationServices field if non-nil, zero value otherwise.

### GetIntegrationServicesOk

`func (o *Integration) GetIntegrationServicesOk() (*[]string, bool)`

GetIntegrationServicesOk returns a tuple with the IntegrationServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationServices

`func (o *Integration) SetIntegrationServices(v []string)`

SetIntegrationServices sets IntegrationServices field to given value.

### HasIntegrationServices

`func (o *Integration) HasIntegrationServices() bool`

HasIntegrationServices returns a boolean if a field has been set.

### GetHash

`func (o *Integration) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Integration) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Integration) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Integration) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


