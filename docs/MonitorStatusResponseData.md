# MonitorStatusResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**MonitorStatusAttributes**](MonitorStatusAttributes.md) |  | [optional] 
**Id** | **string** | Identifier  | 
**Type** | Pointer to **string** | Object type | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) | Relationships of the object | [optional] 
**Links** | Pointer to **map[string]string** | Links related to the object | [optional] 

## Methods

### NewMonitorStatusResponseData

`func NewMonitorStatusResponseData(id string, ) *MonitorStatusResponseData`

NewMonitorStatusResponseData instantiates a new MonitorStatusResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorStatusResponseDataWithDefaults

`func NewMonitorStatusResponseDataWithDefaults() *MonitorStatusResponseData`

NewMonitorStatusResponseDataWithDefaults instantiates a new MonitorStatusResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *MonitorStatusResponseData) GetAttributes() MonitorStatusAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MonitorStatusResponseData) GetAttributesOk() (*MonitorStatusAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MonitorStatusResponseData) SetAttributes(v MonitorStatusAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MonitorStatusResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MonitorStatusResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitorStatusResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitorStatusResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MonitorStatusResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorStatusResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorStatusResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonitorStatusResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *MonitorStatusResponseData) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MonitorStatusResponseData) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MonitorStatusResponseData) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MonitorStatusResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *MonitorStatusResponseData) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MonitorStatusResponseData) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MonitorStatusResponseData) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MonitorStatusResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


