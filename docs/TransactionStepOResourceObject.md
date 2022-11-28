# TransactionStepOResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**TransactionStepOResourceObjectAttributes**](TransactionStepOResourceObjectAttributes.md) |  | [optional] 
**Id** | **int32** | Identifier  | 
**Type** | Pointer to **string** | Object type | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) | Relationships of the object | [optional] 
**Links** | Pointer to **map[string]string** | Links related to the object | [optional] 

## Methods

### NewTransactionStepOResourceObject

`func NewTransactionStepOResourceObject(id int32, ) *TransactionStepOResourceObject`

NewTransactionStepOResourceObject instantiates a new TransactionStepOResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionStepOResourceObjectWithDefaults

`func NewTransactionStepOResourceObjectWithDefaults() *TransactionStepOResourceObject`

NewTransactionStepOResourceObjectWithDefaults instantiates a new TransactionStepOResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *TransactionStepOResourceObject) GetAttributes() TransactionStepOResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TransactionStepOResourceObject) GetAttributesOk() (*TransactionStepOResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TransactionStepOResourceObject) SetAttributes(v TransactionStepOResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TransactionStepOResourceObject) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *TransactionStepOResourceObject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionStepOResourceObject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionStepOResourceObject) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *TransactionStepOResourceObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionStepOResourceObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionStepOResourceObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionStepOResourceObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *TransactionStepOResourceObject) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TransactionStepOResourceObject) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TransactionStepOResourceObject) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TransactionStepOResourceObject) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *TransactionStepOResourceObject) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransactionStepOResourceObject) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransactionStepOResourceObject) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TransactionStepOResourceObject) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


