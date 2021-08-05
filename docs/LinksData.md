# LinksData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | Url for next data set | [optional] 
**Self** | Pointer to **string** | Url for this data set (data might get updated, depending on your parameters) | [optional] 

## Methods

### NewLinksData

`func NewLinksData() *LinksData`

NewLinksData instantiates a new LinksData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksDataWithDefaults

`func NewLinksDataWithDefaults() *LinksData`

NewLinksDataWithDefaults instantiates a new LinksData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *LinksData) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *LinksData) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *LinksData) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *LinksData) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetSelf

`func (o *LinksData) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *LinksData) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *LinksData) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *LinksData) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


