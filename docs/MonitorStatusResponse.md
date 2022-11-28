# MonitorStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MonitorStatusResponseData**](MonitorStatusResponseData.md) |  | [optional] 
**Links** | Pointer to [**StatisticsResponseLinks**](StatisticsResponseLinks.md) |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) | Relationships of the object | [optional] 
**Meta** | Pointer to [**StatisticsResponseMeta**](StatisticsResponseMeta.md) |  | [optional] 

## Methods

### NewMonitorStatusResponse

`func NewMonitorStatusResponse() *MonitorStatusResponse`

NewMonitorStatusResponse instantiates a new MonitorStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorStatusResponseWithDefaults

`func NewMonitorStatusResponseWithDefaults() *MonitorStatusResponse`

NewMonitorStatusResponseWithDefaults instantiates a new MonitorStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MonitorStatusResponse) GetData() MonitorStatusResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MonitorStatusResponse) GetDataOk() (*MonitorStatusResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MonitorStatusResponse) SetData(v MonitorStatusResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *MonitorStatusResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MonitorStatusResponse) GetLinks() StatisticsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MonitorStatusResponse) GetLinksOk() (*StatisticsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MonitorStatusResponse) SetLinks(v StatisticsResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MonitorStatusResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *MonitorStatusResponse) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MonitorStatusResponse) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MonitorStatusResponse) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MonitorStatusResponse) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetMeta

`func (o *MonitorStatusResponse) GetMeta() StatisticsResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MonitorStatusResponse) GetMetaOk() (*StatisticsResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MonitorStatusResponse) SetMeta(v StatisticsResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MonitorStatusResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


