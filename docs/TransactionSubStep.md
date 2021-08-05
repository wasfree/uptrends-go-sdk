# TransactionSubStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | [**SubStepType**](SubStepType.md) |  | 
**Url** | Pointer to **string** |  | [optional] 
**SetValue** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionSubStep

`func NewTransactionSubStep(type_ SubStepType, ) *TransactionSubStep`

NewTransactionSubStep instantiates a new TransactionSubStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSubStepWithDefaults

`func NewTransactionSubStepWithDefaults() *TransactionSubStep`

NewTransactionSubStepWithDefaults instantiates a new TransactionSubStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TransactionSubStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransactionSubStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransactionSubStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransactionSubStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TransactionSubStep) GetType() SubStepType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionSubStep) GetTypeOk() (*SubStepType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionSubStep) SetType(v SubStepType)`

SetType sets Type field to given value.


### GetUrl

`func (o *TransactionSubStep) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TransactionSubStep) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TransactionSubStep) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TransactionSubStep) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSetValue

`func (o *TransactionSubStep) GetSetValue() string`

GetSetValue returns the SetValue field if non-nil, zero value otherwise.

### GetSetValueOk

`func (o *TransactionSubStep) GetSetValueOk() (*string, bool)`

GetSetValueOk returns a tuple with the SetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetValue

`func (o *TransactionSubStep) SetSetValue(v string)`

SetSetValue sets SetValue field to given value.

### HasSetValue

`func (o *TransactionSubStep) HasSetValue() bool`

HasSetValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


