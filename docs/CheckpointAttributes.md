# CheckpointAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckpointName** | Pointer to **string** | The name of the checkpoint | [optional] 
**Code** | Pointer to **string** | Location code of the checkpoint | [optional] 
**Ipv4Addresses** | Pointer to **[]string** | Ipv4 addresses of the checkpoint  | [optional] 
**IpV6Addresses** | Pointer to **[]string** | Ipv6 addresses of the checkpoint | [optional] 
**IsPrimaryCheckpoint** | **bool** | Checkpoint is primary | 
**SupportsIpv6** | **bool** | Checkpoint supports IPv6 | 
**HasHighAvailability** | **bool** | Checkpoint has high availability | 

## Methods

### NewCheckpointAttributes

`func NewCheckpointAttributes(isPrimaryCheckpoint bool, supportsIpv6 bool, hasHighAvailability bool, ) *CheckpointAttributes`

NewCheckpointAttributes instantiates a new CheckpointAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointAttributesWithDefaults

`func NewCheckpointAttributesWithDefaults() *CheckpointAttributes`

NewCheckpointAttributesWithDefaults instantiates a new CheckpointAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckpointName

`func (o *CheckpointAttributes) GetCheckpointName() string`

GetCheckpointName returns the CheckpointName field if non-nil, zero value otherwise.

### GetCheckpointNameOk

`func (o *CheckpointAttributes) GetCheckpointNameOk() (*string, bool)`

GetCheckpointNameOk returns a tuple with the CheckpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointName

`func (o *CheckpointAttributes) SetCheckpointName(v string)`

SetCheckpointName sets CheckpointName field to given value.

### HasCheckpointName

`func (o *CheckpointAttributes) HasCheckpointName() bool`

HasCheckpointName returns a boolean if a field has been set.

### GetCode

`func (o *CheckpointAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CheckpointAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CheckpointAttributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CheckpointAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *CheckpointAttributes) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *CheckpointAttributes) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *CheckpointAttributes) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *CheckpointAttributes) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpV6Addresses

`func (o *CheckpointAttributes) GetIpV6Addresses() []string`

GetIpV6Addresses returns the IpV6Addresses field if non-nil, zero value otherwise.

### GetIpV6AddressesOk

`func (o *CheckpointAttributes) GetIpV6AddressesOk() (*[]string, bool)`

GetIpV6AddressesOk returns a tuple with the IpV6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Addresses

`func (o *CheckpointAttributes) SetIpV6Addresses(v []string)`

SetIpV6Addresses sets IpV6Addresses field to given value.

### HasIpV6Addresses

`func (o *CheckpointAttributes) HasIpV6Addresses() bool`

HasIpV6Addresses returns a boolean if a field has been set.

### GetIsPrimaryCheckpoint

`func (o *CheckpointAttributes) GetIsPrimaryCheckpoint() bool`

GetIsPrimaryCheckpoint returns the IsPrimaryCheckpoint field if non-nil, zero value otherwise.

### GetIsPrimaryCheckpointOk

`func (o *CheckpointAttributes) GetIsPrimaryCheckpointOk() (*bool, bool)`

GetIsPrimaryCheckpointOk returns a tuple with the IsPrimaryCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryCheckpoint

`func (o *CheckpointAttributes) SetIsPrimaryCheckpoint(v bool)`

SetIsPrimaryCheckpoint sets IsPrimaryCheckpoint field to given value.


### GetSupportsIpv6

`func (o *CheckpointAttributes) GetSupportsIpv6() bool`

GetSupportsIpv6 returns the SupportsIpv6 field if non-nil, zero value otherwise.

### GetSupportsIpv6Ok

`func (o *CheckpointAttributes) GetSupportsIpv6Ok() (*bool, bool)`

GetSupportsIpv6Ok returns a tuple with the SupportsIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsIpv6

`func (o *CheckpointAttributes) SetSupportsIpv6(v bool)`

SetSupportsIpv6 sets SupportsIpv6 field to given value.


### GetHasHighAvailability

`func (o *CheckpointAttributes) GetHasHighAvailability() bool`

GetHasHighAvailability returns the HasHighAvailability field if non-nil, zero value otherwise.

### GetHasHighAvailabilityOk

`func (o *CheckpointAttributes) GetHasHighAvailabilityOk() (*bool, bool)`

GetHasHighAvailabilityOk returns a tuple with the HasHighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHighAvailability

`func (o *CheckpointAttributes) SetHasHighAvailability(v bool)`

SetHasHighAvailability sets HasHighAvailability field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


