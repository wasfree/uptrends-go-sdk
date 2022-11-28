# CheckpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CheckpointResponseData**](CheckpointResponseData.md) |  | [optional] 
**Links** | Pointer to [**StatisticsResponseLinks**](StatisticsResponseLinks.md) |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) | Relationships of the object | [optional] 
**Meta** | Pointer to [**StatisticsResponseMeta**](StatisticsResponseMeta.md) |  | [optional] 

## Methods

### NewCheckpointResponse

`func NewCheckpointResponse() *CheckpointResponse`

NewCheckpointResponse instantiates a new CheckpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointResponseWithDefaults

`func NewCheckpointResponseWithDefaults() *CheckpointResponse`

NewCheckpointResponseWithDefaults instantiates a new CheckpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CheckpointResponse) GetData() CheckpointResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CheckpointResponse) GetDataOk() (*CheckpointResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CheckpointResponse) SetData(v CheckpointResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CheckpointResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *CheckpointResponse) GetLinks() StatisticsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CheckpointResponse) GetLinksOk() (*StatisticsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CheckpointResponse) SetLinks(v StatisticsResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CheckpointResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *CheckpointResponse) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CheckpointResponse) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CheckpointResponse) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CheckpointResponse) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetMeta

`func (o *CheckpointResponse) GetMeta() StatisticsResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CheckpointResponse) GetMetaOk() (*StatisticsResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CheckpointResponse) SetMeta(v StatisticsResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CheckpointResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


