# Checkpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckpointId** | **int32** | The Id of the checkpoint | 
**CheckpointName** | Pointer to **string** | The name of the checkpoint | [optional] 
**Code** | Pointer to **string** | The location code of the checkpoint | [optional] 
**Ipv4Addresses** | Pointer to **[]string** | The IPv4 addresses of the checkpoint  | [optional] 
**Ipv6Addresses** | Pointer to [**[]Ipv6Address**](Ipv6Address.md) | The IPv6 addresses of the checkpoint | [optional] 
**IsPrimaryCheckpoint** | **bool** | Checkpoint is primary | 
**SupportsIpv6** | **bool** | Checkpoint supports IPv6 | 
**HasHighAvailability** | **bool** | Checkpoint has high availability | 

## Methods

### NewCheckpoint

`func NewCheckpoint(checkpointId int32, isPrimaryCheckpoint bool, supportsIpv6 bool, hasHighAvailability bool, ) *Checkpoint`

NewCheckpoint instantiates a new Checkpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointWithDefaults

`func NewCheckpointWithDefaults() *Checkpoint`

NewCheckpointWithDefaults instantiates a new Checkpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckpointId

`func (o *Checkpoint) GetCheckpointId() int32`

GetCheckpointId returns the CheckpointId field if non-nil, zero value otherwise.

### GetCheckpointIdOk

`func (o *Checkpoint) GetCheckpointIdOk() (*int32, bool)`

GetCheckpointIdOk returns a tuple with the CheckpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointId

`func (o *Checkpoint) SetCheckpointId(v int32)`

SetCheckpointId sets CheckpointId field to given value.


### GetCheckpointName

`func (o *Checkpoint) GetCheckpointName() string`

GetCheckpointName returns the CheckpointName field if non-nil, zero value otherwise.

### GetCheckpointNameOk

`func (o *Checkpoint) GetCheckpointNameOk() (*string, bool)`

GetCheckpointNameOk returns a tuple with the CheckpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointName

`func (o *Checkpoint) SetCheckpointName(v string)`

SetCheckpointName sets CheckpointName field to given value.

### HasCheckpointName

`func (o *Checkpoint) HasCheckpointName() bool`

HasCheckpointName returns a boolean if a field has been set.

### GetCode

`func (o *Checkpoint) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Checkpoint) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Checkpoint) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Checkpoint) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *Checkpoint) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *Checkpoint) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *Checkpoint) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *Checkpoint) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *Checkpoint) GetIpv6Addresses() []Ipv6Address`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *Checkpoint) GetIpv6AddressesOk() (*[]Ipv6Address, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *Checkpoint) SetIpv6Addresses(v []Ipv6Address)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *Checkpoint) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetIsPrimaryCheckpoint

`func (o *Checkpoint) GetIsPrimaryCheckpoint() bool`

GetIsPrimaryCheckpoint returns the IsPrimaryCheckpoint field if non-nil, zero value otherwise.

### GetIsPrimaryCheckpointOk

`func (o *Checkpoint) GetIsPrimaryCheckpointOk() (*bool, bool)`

GetIsPrimaryCheckpointOk returns a tuple with the IsPrimaryCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryCheckpoint

`func (o *Checkpoint) SetIsPrimaryCheckpoint(v bool)`

SetIsPrimaryCheckpoint sets IsPrimaryCheckpoint field to given value.


### GetSupportsIpv6

`func (o *Checkpoint) GetSupportsIpv6() bool`

GetSupportsIpv6 returns the SupportsIpv6 field if non-nil, zero value otherwise.

### GetSupportsIpv6Ok

`func (o *Checkpoint) GetSupportsIpv6Ok() (*bool, bool)`

GetSupportsIpv6Ok returns a tuple with the SupportsIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsIpv6

`func (o *Checkpoint) SetSupportsIpv6(v bool)`

SetSupportsIpv6 sets SupportsIpv6 field to given value.


### GetHasHighAvailability

`func (o *Checkpoint) GetHasHighAvailability() bool`

GetHasHighAvailability returns the HasHighAvailability field if non-nil, zero value otherwise.

### GetHasHighAvailabilityOk

`func (o *Checkpoint) GetHasHighAvailabilityOk() (*bool, bool)`

GetHasHighAvailabilityOk returns a tuple with the HasHighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHighAvailability

`func (o *Checkpoint) SetHasHighAvailability(v bool)`

SetHasHighAvailability sets HasHighAvailability field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


