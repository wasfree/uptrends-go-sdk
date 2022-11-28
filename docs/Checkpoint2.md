# Checkpoint2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**Checkpoint2Attributes**](Checkpoint2Attributes.md) |  | [optional] 
**Id** | **int32** | Identifier  | 
**Type** | Pointer to **string** | Object type | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) | Relationships of the object | [optional] 
**Links** | Pointer to **map[string]string** | Links related to the object | [optional] 

## Methods

### NewCheckpoint2

`func NewCheckpoint2(id int32, ) *Checkpoint2`

NewCheckpoint2 instantiates a new Checkpoint2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpoint2WithDefaults

`func NewCheckpoint2WithDefaults() *Checkpoint2`

NewCheckpoint2WithDefaults instantiates a new Checkpoint2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Checkpoint2) GetAttributes() Checkpoint2Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Checkpoint2) GetAttributesOk() (*Checkpoint2Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Checkpoint2) SetAttributes(v Checkpoint2Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Checkpoint2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *Checkpoint2) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Checkpoint2) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Checkpoint2) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *Checkpoint2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Checkpoint2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Checkpoint2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Checkpoint2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *Checkpoint2) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Checkpoint2) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Checkpoint2) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Checkpoint2) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *Checkpoint2) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Checkpoint2) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Checkpoint2) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Checkpoint2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


