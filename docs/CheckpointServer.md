# CheckpointServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**CheckpointServerAttributes**](CheckpointServerAttributes.md) |  | [optional] 
**Id** | **int32** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) |  | [optional] 
**Links** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCheckpointServer

`func NewCheckpointServer(id int32, ) *CheckpointServer`

NewCheckpointServer instantiates a new CheckpointServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointServerWithDefaults

`func NewCheckpointServerWithDefaults() *CheckpointServer`

NewCheckpointServerWithDefaults instantiates a new CheckpointServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CheckpointServer) GetAttributes() CheckpointServerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CheckpointServer) GetAttributesOk() (*CheckpointServerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CheckpointServer) SetAttributes(v CheckpointServerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CheckpointServer) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *CheckpointServer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckpointServer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckpointServer) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *CheckpointServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckpointServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckpointServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckpointServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *CheckpointServer) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CheckpointServer) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CheckpointServer) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CheckpointServer) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *CheckpointServer) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CheckpointServer) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CheckpointServer) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CheckpointServer) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


