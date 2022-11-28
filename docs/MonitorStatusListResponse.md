# MonitorStatusListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MonitorStatus**](MonitorStatus.md) | The resposne data/monitor checks | [optional] 
**Links** | Pointer to [**StatisticsResponseLinks**](StatisticsResponseLinks.md) |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) | Relationships of the object | [optional] 
**Meta** | Pointer to [**StatisticsResponseMeta**](StatisticsResponseMeta.md) |  | [optional] 

## Methods

### NewMonitorStatusListResponse

`func NewMonitorStatusListResponse() *MonitorStatusListResponse`

NewMonitorStatusListResponse instantiates a new MonitorStatusListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorStatusListResponseWithDefaults

`func NewMonitorStatusListResponseWithDefaults() *MonitorStatusListResponse`

NewMonitorStatusListResponseWithDefaults instantiates a new MonitorStatusListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MonitorStatusListResponse) GetData() []MonitorStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MonitorStatusListResponse) GetDataOk() (*[]MonitorStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MonitorStatusListResponse) SetData(v []MonitorStatus)`

SetData sets Data field to given value.

### HasData

`func (o *MonitorStatusListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MonitorStatusListResponse) GetLinks() StatisticsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MonitorStatusListResponse) GetLinksOk() (*StatisticsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MonitorStatusListResponse) SetLinks(v StatisticsResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MonitorStatusListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *MonitorStatusListResponse) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MonitorStatusListResponse) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MonitorStatusListResponse) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MonitorStatusListResponse) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetMeta

`func (o *MonitorStatusListResponse) GetMeta() StatisticsResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MonitorStatusListResponse) GetMetaOk() (*StatisticsResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MonitorStatusListResponse) SetMeta(v StatisticsResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MonitorStatusListResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


