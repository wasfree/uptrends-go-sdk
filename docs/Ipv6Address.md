# Ipv6Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | The IPv6 address | [optional] 
**IsNative** | **bool** | This indicates whether this is a native IPv6 address | 

## Methods

### NewIpv6Address

`func NewIpv6Address(isNative bool, ) *Ipv6Address`

NewIpv6Address instantiates a new Ipv6Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6AddressWithDefaults

`func NewIpv6AddressWithDefaults() *Ipv6Address`

NewIpv6AddressWithDefaults instantiates a new Ipv6Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *Ipv6Address) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Ipv6Address) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Ipv6Address) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Ipv6Address) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsNative

`func (o *Ipv6Address) GetIsNative() bool`

GetIsNative returns the IsNative field if non-nil, zero value otherwise.

### GetIsNativeOk

`func (o *Ipv6Address) GetIsNativeOk() (*bool, bool)`

GetIsNativeOk returns a tuple with the IsNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNative

`func (o *Ipv6Address) SetIsNative(v bool)`

SetIsNative sets IsNative field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


