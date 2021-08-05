# TransactionCheckDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**TransactionAttributes**](TransactionAttributes.md) |  | [optional] 
**Id** | **int64** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) |  | [optional] 
**Links** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTransactionCheckDetails

`func NewTransactionCheckDetails(id int64, ) *TransactionCheckDetails`

NewTransactionCheckDetails instantiates a new TransactionCheckDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCheckDetailsWithDefaults

`func NewTransactionCheckDetailsWithDefaults() *TransactionCheckDetails`

NewTransactionCheckDetailsWithDefaults instantiates a new TransactionCheckDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *TransactionCheckDetails) GetAttributes() TransactionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TransactionCheckDetails) GetAttributesOk() (*TransactionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TransactionCheckDetails) SetAttributes(v TransactionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TransactionCheckDetails) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *TransactionCheckDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionCheckDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionCheckDetails) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *TransactionCheckDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionCheckDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionCheckDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionCheckDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *TransactionCheckDetails) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TransactionCheckDetails) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TransactionCheckDetails) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TransactionCheckDetails) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *TransactionCheckDetails) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransactionCheckDetails) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransactionCheckDetails) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TransactionCheckDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


