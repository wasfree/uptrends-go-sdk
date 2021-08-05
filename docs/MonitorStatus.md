# MonitorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**MonitorStatusAttributes**](MonitorStatusAttributes.md) |  | [optional] 
**Id** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) |  | [optional] 
**Links** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMonitorStatus

`func NewMonitorStatus(id string, ) *MonitorStatus`

NewMonitorStatus instantiates a new MonitorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorStatusWithDefaults

`func NewMonitorStatusWithDefaults() *MonitorStatus`

NewMonitorStatusWithDefaults instantiates a new MonitorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *MonitorStatus) GetAttributes() MonitorStatusAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MonitorStatus) GetAttributesOk() (*MonitorStatusAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MonitorStatus) SetAttributes(v MonitorStatusAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MonitorStatus) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MonitorStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitorStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitorStatus) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MonitorStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonitorStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *MonitorStatus) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MonitorStatus) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MonitorStatus) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MonitorStatus) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *MonitorStatus) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MonitorStatus) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MonitorStatus) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MonitorStatus) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


