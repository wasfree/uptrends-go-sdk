# MsaDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsaDetailsResponseData**](MsaDetailsResponseData.md) |  | [optional] 
**Links** | Pointer to [**StatisticsResponseLinks**](StatisticsResponseLinks.md) |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) | Relationships of the object | [optional] 
**Meta** | Pointer to [**StatisticsResponseMeta**](StatisticsResponseMeta.md) |  | [optional] 

## Methods

### NewMsaDetailsResponse

`func NewMsaDetailsResponse() *MsaDetailsResponse`

NewMsaDetailsResponse instantiates a new MsaDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsaDetailsResponseWithDefaults

`func NewMsaDetailsResponseWithDefaults() *MsaDetailsResponse`

NewMsaDetailsResponseWithDefaults instantiates a new MsaDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsaDetailsResponse) GetData() MsaDetailsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsaDetailsResponse) GetDataOk() (*MsaDetailsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsaDetailsResponse) SetData(v MsaDetailsResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *MsaDetailsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsaDetailsResponse) GetLinks() StatisticsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsaDetailsResponse) GetLinksOk() (*StatisticsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsaDetailsResponse) SetLinks(v StatisticsResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsaDetailsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *MsaDetailsResponse) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MsaDetailsResponse) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MsaDetailsResponse) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MsaDetailsResponse) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetMeta

`func (o *MsaDetailsResponse) GetMeta() StatisticsResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsaDetailsResponse) GetMetaOk() (*StatisticsResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsaDetailsResponse) SetMeta(v StatisticsResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MsaDetailsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


