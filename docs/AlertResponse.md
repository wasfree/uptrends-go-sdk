# AlertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Alert**](Alert.md) |  | [optional] 
**Links** | Pointer to [**LinksData**](LinksData.md) |  | [optional] 
**Relationships** | Pointer to [**[]RelationObject**](RelationObject.md) |  | [optional] 
**Meta** | Pointer to [**MetaData**](MetaData.md) |  | [optional] 
**Cursors** | Pointer to [**CursorsData**](CursorsData.md) | Cursors can be used to navigate the dataset in a fixed manner | [optional] 

## Methods

### NewAlertResponse

`func NewAlertResponse() *AlertResponse`

NewAlertResponse instantiates a new AlertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertResponseWithDefaults

`func NewAlertResponseWithDefaults() *AlertResponse`

NewAlertResponseWithDefaults instantiates a new AlertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AlertResponse) GetData() []Alert`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlertResponse) GetDataOk() (*[]Alert, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlertResponse) SetData(v []Alert)`

SetData sets Data field to given value.

### HasData

`func (o *AlertResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *AlertResponse) GetLinks() LinksData`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlertResponse) GetLinksOk() (*LinksData, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlertResponse) SetLinks(v LinksData)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AlertResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *AlertResponse) GetRelationships() []RelationObject`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AlertResponse) GetRelationshipsOk() (*[]RelationObject, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AlertResponse) SetRelationships(v []RelationObject)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AlertResponse) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetMeta

`func (o *AlertResponse) GetMeta() MetaData`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AlertResponse) GetMetaOk() (*MetaData, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AlertResponse) SetMeta(v MetaData)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AlertResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetCursors

`func (o *AlertResponse) GetCursors() CursorsData`

GetCursors returns the Cursors field if non-nil, zero value otherwise.

### GetCursorsOk

`func (o *AlertResponse) GetCursorsOk() (*CursorsData, bool)`

GetCursorsOk returns a tuple with the Cursors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursors

`func (o *AlertResponse) SetCursors(v CursorsData)`

SetCursors sets Cursors field to given value.

### HasCursors

`func (o *AlertResponse) HasCursors() bool`

HasCursors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


