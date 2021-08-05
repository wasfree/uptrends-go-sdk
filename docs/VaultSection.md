# VaultSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaultSectionGuid** | **string** | The Guid of this vault section | 
**Name** | Pointer to **string** | The name for this vault section | [optional] 

## Methods

### NewVaultSection

`func NewVaultSection(vaultSectionGuid string, ) *VaultSection`

NewVaultSection instantiates a new VaultSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultSectionWithDefaults

`func NewVaultSectionWithDefaults() *VaultSection`

NewVaultSectionWithDefaults instantiates a new VaultSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaultSectionGuid

`func (o *VaultSection) GetVaultSectionGuid() string`

GetVaultSectionGuid returns the VaultSectionGuid field if non-nil, zero value otherwise.

### GetVaultSectionGuidOk

`func (o *VaultSection) GetVaultSectionGuidOk() (*string, bool)`

GetVaultSectionGuidOk returns a tuple with the VaultSectionGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSectionGuid

`func (o *VaultSection) SetVaultSectionGuid(v string)`

SetVaultSectionGuid sets VaultSectionGuid field to given value.


### GetName

`func (o *VaultSection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VaultSection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VaultSection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VaultSection) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


