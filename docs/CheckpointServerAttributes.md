# CheckpointServerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip4Address** | Pointer to **string** | The ipv4 address | [optional] 
**IpV6Address** | Pointer to **string** | The ipv6 address | [optional] 
**Capabilities** | Pointer to [**[]CapabilityFilterEnum**](CapabilityFilterEnum.md) | List of server&#39;s capabilities | [optional] 

## Methods

### NewCheckpointServerAttributes

`func NewCheckpointServerAttributes() *CheckpointServerAttributes`

NewCheckpointServerAttributes instantiates a new CheckpointServerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointServerAttributesWithDefaults

`func NewCheckpointServerAttributesWithDefaults() *CheckpointServerAttributes`

NewCheckpointServerAttributesWithDefaults instantiates a new CheckpointServerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp4Address

`func (o *CheckpointServerAttributes) GetIp4Address() string`

GetIp4Address returns the Ip4Address field if non-nil, zero value otherwise.

### GetIp4AddressOk

`func (o *CheckpointServerAttributes) GetIp4AddressOk() (*string, bool)`

GetIp4AddressOk returns a tuple with the Ip4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4Address

`func (o *CheckpointServerAttributes) SetIp4Address(v string)`

SetIp4Address sets Ip4Address field to given value.

### HasIp4Address

`func (o *CheckpointServerAttributes) HasIp4Address() bool`

HasIp4Address returns a boolean if a field has been set.

### GetIpV6Address

`func (o *CheckpointServerAttributes) GetIpV6Address() string`

GetIpV6Address returns the IpV6Address field if non-nil, zero value otherwise.

### GetIpV6AddressOk

`func (o *CheckpointServerAttributes) GetIpV6AddressOk() (*string, bool)`

GetIpV6AddressOk returns a tuple with the IpV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Address

`func (o *CheckpointServerAttributes) SetIpV6Address(v string)`

SetIpV6Address sets IpV6Address field to given value.

### HasIpV6Address

`func (o *CheckpointServerAttributes) HasIpV6Address() bool`

HasIpV6Address returns a boolean if a field has been set.

### GetCapabilities

`func (o *CheckpointServerAttributes) GetCapabilities() []CapabilityFilterEnum`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CheckpointServerAttributes) GetCapabilitiesOk() (*[]CapabilityFilterEnum, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CheckpointServerAttributes) SetCapabilities(v []CapabilityFilterEnum)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CheckpointServerAttributes) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


