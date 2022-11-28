# CertificateArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **string** | The issuer of this certificate | [optional] 
**NotBefore** | **map[string]interface{}** | The start date of this certificate | 
**NotAfter** | **map[string]interface{}** | The certificate expiry date | 
**Password** | Pointer to **string** | Used to specify a new or changed password. The value will be encrypted when stored, and will not be displayed when a CertificateArchive is retrieved. | [optional] 
**ArchiveData** | Pointer to **string** | Used to specify the certificate content as a base64 string. The value will be encrypted when stored, and will not be displayed when a CertificateArchive is retrieved. | [optional] 

## Methods

### NewCertificateArchive

`func NewCertificateArchive(notBefore map[string]interface{}, notAfter map[string]interface{}, ) *CertificateArchive`

NewCertificateArchive instantiates a new CertificateArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateArchiveWithDefaults

`func NewCertificateArchiveWithDefaults() *CertificateArchive`

NewCertificateArchiveWithDefaults instantiates a new CertificateArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *CertificateArchive) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateArchive) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateArchive) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CertificateArchive) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertificateArchive) GetNotBefore() map[string]interface{}`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateArchive) GetNotBeforeOk() (*map[string]interface{}, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateArchive) SetNotBefore(v map[string]interface{})`

SetNotBefore sets NotBefore field to given value.


### GetNotAfter

`func (o *CertificateArchive) GetNotAfter() map[string]interface{}`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateArchive) GetNotAfterOk() (*map[string]interface{}, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateArchive) SetNotAfter(v map[string]interface{})`

SetNotAfter sets NotAfter field to given value.


### GetPassword

`func (o *CertificateArchive) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CertificateArchive) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CertificateArchive) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CertificateArchive) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetArchiveData

`func (o *CertificateArchive) GetArchiveData() string`

GetArchiveData returns the ArchiveData field if non-nil, zero value otherwise.

### GetArchiveDataOk

`func (o *CertificateArchive) GetArchiveDataOk() (*string, bool)`

GetArchiveDataOk returns a tuple with the ArchiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveData

`func (o *CertificateArchive) SetArchiveData(v string)`

SetArchiveData sets ArchiveData field to given value.

### HasArchiveData

`func (o *CertificateArchive) HasArchiveData() bool`

HasArchiveData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


