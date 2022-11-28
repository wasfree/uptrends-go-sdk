# VaultItemCertificateArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **string** | The issuer of this certificate | [optional] 
**NotBefore** | **map[string]interface{}** | The start date of this certificate | 
**NotAfter** | **map[string]interface{}** | The certificate expiry date | 
**Password** | Pointer to **string** | Used to specify a new or changed password. The value will be encrypted when stored, and will not be displayed when a CertificateArchive is retrieved. | [optional] 
**ArchiveData** | Pointer to **string** | Used to specify the certificate content as a base64 string. The value will be encrypted when stored, and will not be displayed when a CertificateArchive is retrieved. | [optional] 

## Methods

### NewVaultItemCertificateArchive

`func NewVaultItemCertificateArchive(notBefore map[string]interface{}, notAfter map[string]interface{}, ) *VaultItemCertificateArchive`

NewVaultItemCertificateArchive instantiates a new VaultItemCertificateArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultItemCertificateArchiveWithDefaults

`func NewVaultItemCertificateArchiveWithDefaults() *VaultItemCertificateArchive`

NewVaultItemCertificateArchiveWithDefaults instantiates a new VaultItemCertificateArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *VaultItemCertificateArchive) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *VaultItemCertificateArchive) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *VaultItemCertificateArchive) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *VaultItemCertificateArchive) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetNotBefore

`func (o *VaultItemCertificateArchive) GetNotBefore() map[string]interface{}`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *VaultItemCertificateArchive) GetNotBeforeOk() (*map[string]interface{}, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *VaultItemCertificateArchive) SetNotBefore(v map[string]interface{})`

SetNotBefore sets NotBefore field to given value.


### GetNotAfter

`func (o *VaultItemCertificateArchive) GetNotAfter() map[string]interface{}`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *VaultItemCertificateArchive) GetNotAfterOk() (*map[string]interface{}, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *VaultItemCertificateArchive) SetNotAfter(v map[string]interface{})`

SetNotAfter sets NotAfter field to given value.


### GetPassword

`func (o *VaultItemCertificateArchive) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VaultItemCertificateArchive) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VaultItemCertificateArchive) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VaultItemCertificateArchive) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetArchiveData

`func (o *VaultItemCertificateArchive) GetArchiveData() string`

GetArchiveData returns the ArchiveData field if non-nil, zero value otherwise.

### GetArchiveDataOk

`func (o *VaultItemCertificateArchive) GetArchiveDataOk() (*string, bool)`

GetArchiveDataOk returns a tuple with the ArchiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveData

`func (o *VaultItemCertificateArchive) SetArchiveData(v string)`

SetArchiveData sets ArchiveData field to given value.

### HasArchiveData

`func (o *VaultItemCertificateArchive) HasArchiveData() bool`

HasArchiveData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


