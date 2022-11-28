# VaultItemFileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the uploaded file. | [optional] 
**Size** | **int32** | Size of the uploaded file. | 

## Methods

### NewVaultItemFileInfo

`func NewVaultItemFileInfo(size int32, ) *VaultItemFileInfo`

NewVaultItemFileInfo instantiates a new VaultItemFileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultItemFileInfoWithDefaults

`func NewVaultItemFileInfoWithDefaults() *VaultItemFileInfo`

NewVaultItemFileInfoWithDefaults instantiates a new VaultItemFileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VaultItemFileInfo) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VaultItemFileInfo) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VaultItemFileInfo) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *VaultItemFileInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetName

`func (o *VaultItemFileInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VaultItemFileInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VaultItemFileInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VaultItemFileInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *VaultItemFileInfo) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VaultItemFileInfo) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VaultItemFileInfo) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


