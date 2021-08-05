# VaultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaultItemGuid** | Pointer to **string** | The unique key of this vault item | [optional] 
**Hash** | Pointer to **string** | The hash of this vault item | [optional] 
**Name** | Pointer to **string** | The name of this vault item | [optional] 
**Value** | Pointer to **string** | The value that is stored in this vault item. Not used for Certificate Archives | [optional] 
**VaultSectionGuid** | Pointer to **string** | The unique identifier of the vault section that this vault item belongs to | [optional] 
**VaultItemType** | Pointer to [**VaultItemTypes**](VaultItemTypes.md) | The vault item type | [optional] 
**IsSensitive** | Pointer to **bool** | Whether or not the vault item is considered sensitive.  | [optional] 
**Notes** | Pointer to **string** | Notes about this vault item | [optional] 
**UserName** | Pointer to **string** | The UserName of a credentialset | [optional] 
**Password** | Pointer to **string** | The password associated with a credentialset | [optional] 
**CertificateArchive** | Pointer to [**CertificateArchive**](CertificateArchive.md) | The certificate archive that is stored in this vault item, if applicable | [optional] 
**FileInfo** | Pointer to [**FileInfo**](FileInfo.md) | The file info that is stored in this vault item, if applicable | [optional] 

## Methods

### NewVaultItem

`func NewVaultItem() *VaultItem`

NewVaultItem instantiates a new VaultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultItemWithDefaults

`func NewVaultItemWithDefaults() *VaultItem`

NewVaultItemWithDefaults instantiates a new VaultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaultItemGuid

`func (o *VaultItem) GetVaultItemGuid() string`

GetVaultItemGuid returns the VaultItemGuid field if non-nil, zero value otherwise.

### GetVaultItemGuidOk

`func (o *VaultItem) GetVaultItemGuidOk() (*string, bool)`

GetVaultItemGuidOk returns a tuple with the VaultItemGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultItemGuid

`func (o *VaultItem) SetVaultItemGuid(v string)`

SetVaultItemGuid sets VaultItemGuid field to given value.

### HasVaultItemGuid

`func (o *VaultItem) HasVaultItemGuid() bool`

HasVaultItemGuid returns a boolean if a field has been set.

### GetHash

`func (o *VaultItem) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *VaultItem) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *VaultItem) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *VaultItem) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetName

`func (o *VaultItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VaultItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VaultItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VaultItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *VaultItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VaultItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VaultItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VaultItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVaultSectionGuid

`func (o *VaultItem) GetVaultSectionGuid() string`

GetVaultSectionGuid returns the VaultSectionGuid field if non-nil, zero value otherwise.

### GetVaultSectionGuidOk

`func (o *VaultItem) GetVaultSectionGuidOk() (*string, bool)`

GetVaultSectionGuidOk returns a tuple with the VaultSectionGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSectionGuid

`func (o *VaultItem) SetVaultSectionGuid(v string)`

SetVaultSectionGuid sets VaultSectionGuid field to given value.

### HasVaultSectionGuid

`func (o *VaultItem) HasVaultSectionGuid() bool`

HasVaultSectionGuid returns a boolean if a field has been set.

### GetVaultItemType

`func (o *VaultItem) GetVaultItemType() VaultItemTypes`

GetVaultItemType returns the VaultItemType field if non-nil, zero value otherwise.

### GetVaultItemTypeOk

`func (o *VaultItem) GetVaultItemTypeOk() (*VaultItemTypes, bool)`

GetVaultItemTypeOk returns a tuple with the VaultItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultItemType

`func (o *VaultItem) SetVaultItemType(v VaultItemTypes)`

SetVaultItemType sets VaultItemType field to given value.

### HasVaultItemType

`func (o *VaultItem) HasVaultItemType() bool`

HasVaultItemType returns a boolean if a field has been set.

### GetIsSensitive

`func (o *VaultItem) GetIsSensitive() bool`

GetIsSensitive returns the IsSensitive field if non-nil, zero value otherwise.

### GetIsSensitiveOk

`func (o *VaultItem) GetIsSensitiveOk() (*bool, bool)`

GetIsSensitiveOk returns a tuple with the IsSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSensitive

`func (o *VaultItem) SetIsSensitive(v bool)`

SetIsSensitive sets IsSensitive field to given value.

### HasIsSensitive

`func (o *VaultItem) HasIsSensitive() bool`

HasIsSensitive returns a boolean if a field has been set.

### GetNotes

`func (o *VaultItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *VaultItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *VaultItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *VaultItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetUserName

`func (o *VaultItem) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VaultItem) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VaultItem) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *VaultItem) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *VaultItem) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VaultItem) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VaultItem) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VaultItem) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCertificateArchive

`func (o *VaultItem) GetCertificateArchive() CertificateArchive`

GetCertificateArchive returns the CertificateArchive field if non-nil, zero value otherwise.

### GetCertificateArchiveOk

`func (o *VaultItem) GetCertificateArchiveOk() (*CertificateArchive, bool)`

GetCertificateArchiveOk returns a tuple with the CertificateArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateArchive

`func (o *VaultItem) SetCertificateArchive(v CertificateArchive)`

SetCertificateArchive sets CertificateArchive field to given value.

### HasCertificateArchive

`func (o *VaultItem) HasCertificateArchive() bool`

HasCertificateArchive returns a boolean if a field has been set.

### GetFileInfo

`func (o *VaultItem) GetFileInfo() FileInfo`

GetFileInfo returns the FileInfo field if non-nil, zero value otherwise.

### GetFileInfoOk

`func (o *VaultItem) GetFileInfoOk() (*FileInfo, bool)`

GetFileInfoOk returns a tuple with the FileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileInfo

`func (o *VaultItem) SetFileInfo(v FileInfo)`

SetFileInfo sets FileInfo field to given value.

### HasFileInfo

`func (o *VaultItem) HasFileInfo() bool`

HasFileInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


