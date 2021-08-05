# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**AlertAttributes**](AlertAttributes.md) |  | [optional] 
**Id** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) |  | [optional] 
**Links** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAlert

`func NewAlert(id string, ) *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Alert) GetAttributes() AlertAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Alert) GetAttributesOk() (*AlertAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Alert) SetAttributes(v AlertAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Alert) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *Alert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alert) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Alert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Alert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Alert) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Alert) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *Alert) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Alert) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Alert) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Alert) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *Alert) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Alert) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Alert) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Alert) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


