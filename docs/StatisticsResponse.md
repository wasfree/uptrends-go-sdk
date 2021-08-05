# StatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Statistics**](Statistics.md) |  | [optional] 
**Links** | Pointer to [**LinksData**](LinksData.md) |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) |  | [optional] 
**Meta** | Pointer to [**MetaData**](MetaData.md) |  | [optional] 

## Methods

### NewStatisticsResponse

`func NewStatisticsResponse() *StatisticsResponse`

NewStatisticsResponse instantiates a new StatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsResponseWithDefaults

`func NewStatisticsResponseWithDefaults() *StatisticsResponse`

NewStatisticsResponseWithDefaults instantiates a new StatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *StatisticsResponse) GetData() []Statistics`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StatisticsResponse) GetDataOk() (*[]Statistics, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StatisticsResponse) SetData(v []Statistics)`

SetData sets Data field to given value.

### HasData

`func (o *StatisticsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *StatisticsResponse) GetLinks() LinksData`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StatisticsResponse) GetLinksOk() (*LinksData, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StatisticsResponse) SetLinks(v LinksData)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StatisticsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *StatisticsResponse) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StatisticsResponse) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StatisticsResponse) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StatisticsResponse) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetMeta

`func (o *StatisticsResponse) GetMeta() MetaData`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StatisticsResponse) GetMetaOk() (*MetaData, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StatisticsResponse) SetMeta(v MetaData)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StatisticsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


