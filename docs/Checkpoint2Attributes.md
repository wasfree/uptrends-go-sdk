# Checkpoint2Attributes

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

### NewCheckpoint2Attributes

`func NewCheckpoint2Attributes(isPrimaryCheckpoint bool, supportsIpv6 bool, hasHighAvailability bool, ) *Checkpoint2Attributes`

NewCheckpoint2Attributes instantiates a new Checkpoint2Attributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpoint2AttributesWithDefaults

`func NewCheckpoint2AttributesWithDefaults() *Checkpoint2Attributes`

NewCheckpoint2AttributesWithDefaults instantiates a new Checkpoint2Attributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckpointName

`func (o *Checkpoint2Attributes) GetCheckpointName() string`

GetCheckpointName returns the CheckpointName field if non-nil, zero value otherwise.

### GetCheckpointNameOk

`func (o *Checkpoint2Attributes) GetCheckpointNameOk() (*string, bool)`

GetCheckpointNameOk returns a tuple with the CheckpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointName

`func (o *Checkpoint2Attributes) SetCheckpointName(v string)`

SetCheckpointName sets CheckpointName field to given value.

### HasCheckpointName

`func (o *Checkpoint2Attributes) HasCheckpointName() bool`

HasCheckpointName returns a boolean if a field has been set.

### GetCode

`func (o *Checkpoint2Attributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Checkpoint2Attributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Checkpoint2Attributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Checkpoint2Attributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *Checkpoint2Attributes) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *Checkpoint2Attributes) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *Checkpoint2Attributes) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *Checkpoint2Attributes) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpV6Addresses

`func (o *Checkpoint2Attributes) GetIpV6Addresses() []string`

GetIpV6Addresses returns the IpV6Addresses field if non-nil, zero value otherwise.

### GetIpV6AddressesOk

`func (o *Checkpoint2Attributes) GetIpV6AddressesOk() (*[]string, bool)`

GetIpV6AddressesOk returns a tuple with the IpV6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Addresses

`func (o *Checkpoint2Attributes) SetIpV6Addresses(v []string)`

SetIpV6Addresses sets IpV6Addresses field to given value.

### HasIpV6Addresses

`func (o *Checkpoint2Attributes) HasIpV6Addresses() bool`

HasIpV6Addresses returns a boolean if a field has been set.

### GetIsPrimaryCheckpoint

`func (o *Checkpoint2Attributes) GetIsPrimaryCheckpoint() bool`

GetIsPrimaryCheckpoint returns the IsPrimaryCheckpoint field if non-nil, zero value otherwise.

### GetIsPrimaryCheckpointOk

`func (o *Checkpoint2Attributes) GetIsPrimaryCheckpointOk() (*bool, bool)`

GetIsPrimaryCheckpointOk returns a tuple with the IsPrimaryCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryCheckpoint

`func (o *Checkpoint2Attributes) SetIsPrimaryCheckpoint(v bool)`

SetIsPrimaryCheckpoint sets IsPrimaryCheckpoint field to given value.


### GetSupportsIpv6

`func (o *Checkpoint2Attributes) GetSupportsIpv6() bool`

GetSupportsIpv6 returns the SupportsIpv6 field if non-nil, zero value otherwise.

### GetSupportsIpv6Ok

`func (o *Checkpoint2Attributes) GetSupportsIpv6Ok() (*bool, bool)`

GetSupportsIpv6Ok returns a tuple with the SupportsIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsIpv6

`func (o *Checkpoint2Attributes) SetSupportsIpv6(v bool)`

SetSupportsIpv6 sets SupportsIpv6 field to given value.


### GetHasHighAvailability

`func (o *Checkpoint2Attributes) GetHasHighAvailability() bool`

GetHasHighAvailability returns the HasHighAvailability field if non-nil, zero value otherwise.

### GetHasHighAvailabilityOk

`func (o *Checkpoint2Attributes) GetHasHighAvailabilityOk() (*bool, bool)`

GetHasHighAvailabilityOk returns a tuple with the HasHighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHighAvailability

`func (o *Checkpoint2Attributes) SetHasHighAvailability(v bool)`

SetHasHighAvailability sets HasHighAvailability field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


