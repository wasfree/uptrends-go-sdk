# Statistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**StatisticsAttributes**](StatisticsAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | Identifier  | [optional] 
**Type** | Pointer to **string** | Object type | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) | Relationships of the object | [optional] 
**Links** | Pointer to **map[string]string** | Links related to the object | [optional] 

## Methods

### NewStatistics

`func NewStatistics() *Statistics`

NewStatistics instantiates a new Statistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsWithDefaults

`func NewStatisticsWithDefaults() *Statistics`

NewStatisticsWithDefaults instantiates a new Statistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Statistics) GetAttributes() StatisticsAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Statistics) GetAttributesOk() (*StatisticsAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Statistics) SetAttributes(v StatisticsAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Statistics) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *Statistics) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Statistics) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Statistics) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Statistics) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Statistics) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Statistics) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Statistics) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Statistics) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *Statistics) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Statistics) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Statistics) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Statistics) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *Statistics) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Statistics) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Statistics) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Statistics) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


