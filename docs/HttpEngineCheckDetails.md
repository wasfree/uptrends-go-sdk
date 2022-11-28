# HttpEngineCheckDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**HttpEngineCheckDetailsAttributes**](HttpEngineCheckDetailsAttributes.md) |  | [optional] 
**Id** | **int64** | Identifier  | 
**Type** | Pointer to **string** | Object type | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) | Relationships of the object | [optional] 
**Links** | Pointer to **map[string]string** | Links related to the object | [optional] 

## Methods

### NewHttpEngineCheckDetails

`func NewHttpEngineCheckDetails(id int64, ) *HttpEngineCheckDetails`

NewHttpEngineCheckDetails instantiates a new HttpEngineCheckDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpEngineCheckDetailsWithDefaults

`func NewHttpEngineCheckDetailsWithDefaults() *HttpEngineCheckDetails`

NewHttpEngineCheckDetailsWithDefaults instantiates a new HttpEngineCheckDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *HttpEngineCheckDetails) GetAttributes() HttpEngineCheckDetailsAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *HttpEngineCheckDetails) GetAttributesOk() (*HttpEngineCheckDetailsAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *HttpEngineCheckDetails) SetAttributes(v HttpEngineCheckDetailsAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *HttpEngineCheckDetails) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *HttpEngineCheckDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HttpEngineCheckDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HttpEngineCheckDetails) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *HttpEngineCheckDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HttpEngineCheckDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HttpEngineCheckDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HttpEngineCheckDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *HttpEngineCheckDetails) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *HttpEngineCheckDetails) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *HttpEngineCheckDetails) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *HttpEngineCheckDetails) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *HttpEngineCheckDetails) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HttpEngineCheckDetails) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HttpEngineCheckDetails) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *HttpEngineCheckDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


