# VaultItemOneTimePasswordInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** |  | [optional] 
**Digits** | **int32** |  | 
**Period** | **int32** |  | 
**HashAlgorithm** | [**HashAlgorithm**](HashAlgorithm.md) |  | 

## Methods

### NewVaultItemOneTimePasswordInfo

`func NewVaultItemOneTimePasswordInfo(digits int32, period int32, hashAlgorithm HashAlgorithm, ) *VaultItemOneTimePasswordInfo`

NewVaultItemOneTimePasswordInfo instantiates a new VaultItemOneTimePasswordInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultItemOneTimePasswordInfoWithDefaults

`func NewVaultItemOneTimePasswordInfoWithDefaults() *VaultItemOneTimePasswordInfo`

NewVaultItemOneTimePasswordInfoWithDefaults instantiates a new VaultItemOneTimePasswordInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *VaultItemOneTimePasswordInfo) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *VaultItemOneTimePasswordInfo) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *VaultItemOneTimePasswordInfo) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *VaultItemOneTimePasswordInfo) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetDigits

`func (o *VaultItemOneTimePasswordInfo) GetDigits() int32`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *VaultItemOneTimePasswordInfo) GetDigitsOk() (*int32, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *VaultItemOneTimePasswordInfo) SetDigits(v int32)`

SetDigits sets Digits field to given value.


### GetPeriod

`func (o *VaultItemOneTimePasswordInfo) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *VaultItemOneTimePasswordInfo) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *VaultItemOneTimePasswordInfo) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetHashAlgorithm

`func (o *VaultItemOneTimePasswordInfo) GetHashAlgorithm() HashAlgorithm`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *VaultItemOneTimePasswordInfo) GetHashAlgorithmOk() (*HashAlgorithm, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *VaultItemOneTimePasswordInfo) SetHashAlgorithm(v HashAlgorithm)`

SetHashAlgorithm sets HashAlgorithm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


