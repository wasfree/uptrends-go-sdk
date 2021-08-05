# IPv6Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | The Ipv6 address | [optional] 
**IsNative** | **bool** | Is this a native v6 address? | 

## Methods

### NewIPv6Address

`func NewIPv6Address(isNative bool, ) *IPv6Address`

NewIPv6Address instantiates a new IPv6Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv6AddressWithDefaults

`func NewIPv6AddressWithDefaults() *IPv6Address`

NewIPv6AddressWithDefaults instantiates a new IPv6Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *IPv6Address) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IPv6Address) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IPv6Address) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *IPv6Address) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsNative

`func (o *IPv6Address) GetIsNative() bool`

GetIsNative returns the IsNative field if non-nil, zero value otherwise.

### GetIsNativeOk

`func (o *IPv6Address) GetIsNativeOk() (*bool, bool)`

GetIsNativeOk returns a tuple with the IsNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNative

`func (o *IPv6Address) SetIsNative(v bool)`

SetIsNative sets IsNative field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


