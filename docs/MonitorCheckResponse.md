# MonitorCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MonitorCheck**](MonitorCheck.md) |  | [optional] 
**Links** | Pointer to [**LinksData**](LinksData.md) |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) |  | [optional] 
**Meta** | Pointer to [**MetaData**](MetaData.md) |  | [optional] 
**Cursors** | Pointer to [**CursorsData**](CursorsData.md) | Cursors can be used to navigate the dataset in a fixed manner | [optional] 

## Methods

### NewMonitorCheckResponse

`func NewMonitorCheckResponse() *MonitorCheckResponse`

NewMonitorCheckResponse instantiates a new MonitorCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorCheckResponseWithDefaults

`func NewMonitorCheckResponseWithDefaults() *MonitorCheckResponse`

NewMonitorCheckResponseWithDefaults instantiates a new MonitorCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MonitorCheckResponse) GetData() []MonitorCheck`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MonitorCheckResponse) GetDataOk() (*[]MonitorCheck, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MonitorCheckResponse) SetData(v []MonitorCheck)`

SetData sets Data field to given value.

### HasData

`func (o *MonitorCheckResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MonitorCheckResponse) GetLinks() LinksData`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MonitorCheckResponse) GetLinksOk() (*LinksData, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MonitorCheckResponse) SetLinks(v LinksData)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MonitorCheckResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *MonitorCheckResponse) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MonitorCheckResponse) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MonitorCheckResponse) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MonitorCheckResponse) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetMeta

`func (o *MonitorCheckResponse) GetMeta() MetaData`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MonitorCheckResponse) GetMetaOk() (*MetaData, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MonitorCheckResponse) SetMeta(v MetaData)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MonitorCheckResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetCursors

`func (o *MonitorCheckResponse) GetCursors() CursorsData`

GetCursors returns the Cursors field if non-nil, zero value otherwise.

### GetCursorsOk

`func (o *MonitorCheckResponse) GetCursorsOk() (*CursorsData, bool)`

GetCursorsOk returns a tuple with the Cursors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursors

`func (o *MonitorCheckResponse) SetCursors(v CursorsData)`

SetCursors sets Cursors field to given value.

### HasCursors

`func (o *MonitorCheckResponse) HasCursors() bool`

HasCursors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


