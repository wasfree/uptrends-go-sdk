# HttpCheckDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**HttpCheckDetailsAttributes**](HttpCheckDetailsAttributes.md) |  | [optional] 
**Id** | **int64** | Identifier  | 
**Type** | Pointer to **string** | Object type | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) | Relationships of the object | [optional] 
**Links** | Pointer to **map[string]string** | Links related to the object | [optional] 

## Methods

### NewHttpCheckDetails

`func NewHttpCheckDetails(id int64, ) *HttpCheckDetails`

NewHttpCheckDetails instantiates a new HttpCheckDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpCheckDetailsWithDefaults

`func NewHttpCheckDetailsWithDefaults() *HttpCheckDetails`

NewHttpCheckDetailsWithDefaults instantiates a new HttpCheckDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *HttpCheckDetails) GetAttributes() HttpCheckDetailsAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *HttpCheckDetails) GetAttributesOk() (*HttpCheckDetailsAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *HttpCheckDetails) SetAttributes(v HttpCheckDetailsAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *HttpCheckDetails) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *HttpCheckDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HttpCheckDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HttpCheckDetails) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *HttpCheckDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HttpCheckDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HttpCheckDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HttpCheckDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *HttpCheckDetails) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *HttpCheckDetails) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *HttpCheckDetails) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *HttpCheckDetails) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *HttpCheckDetails) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HttpCheckDetails) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HttpCheckDetails) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *HttpCheckDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


