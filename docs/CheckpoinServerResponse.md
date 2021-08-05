# CheckpoinServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CheckpointServer**](CheckpointServer.md) |  | [optional] 
**Links** | Pointer to [**LinksData**](LinksData.md) |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) |  | [optional] 
**Meta** | Pointer to [**MetaData**](MetaData.md) |  | [optional] 

## Methods

### NewCheckpoinServerResponse

`func NewCheckpoinServerResponse() *CheckpoinServerResponse`

NewCheckpoinServerResponse instantiates a new CheckpoinServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpoinServerResponseWithDefaults

`func NewCheckpoinServerResponseWithDefaults() *CheckpoinServerResponse`

NewCheckpoinServerResponseWithDefaults instantiates a new CheckpoinServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CheckpoinServerResponse) GetData() CheckpointServer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CheckpoinServerResponse) GetDataOk() (*CheckpointServer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CheckpoinServerResponse) SetData(v CheckpointServer)`

SetData sets Data field to given value.

### HasData

`func (o *CheckpoinServerResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *CheckpoinServerResponse) GetLinks() LinksData`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CheckpoinServerResponse) GetLinksOk() (*LinksData, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CheckpoinServerResponse) SetLinks(v LinksData)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CheckpoinServerResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *CheckpoinServerResponse) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CheckpoinServerResponse) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CheckpoinServerResponse) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CheckpoinServerResponse) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetMeta

`func (o *CheckpoinServerResponse) GetMeta() MetaData`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CheckpoinServerResponse) GetMetaOk() (*MetaData, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CheckpoinServerResponse) SetMeta(v MetaData)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CheckpoinServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


