# AlertResponseCursors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | Cursor for next data set | [optional] 
**Self** | Pointer to **string** | Cursor for this data set (data might get updated, depending on your parameters) | [optional] 

## Methods

### NewAlertResponseCursors

`func NewAlertResponseCursors() *AlertResponseCursors`

NewAlertResponseCursors instantiates a new AlertResponseCursors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertResponseCursorsWithDefaults

`func NewAlertResponseCursorsWithDefaults() *AlertResponseCursors`

NewAlertResponseCursorsWithDefaults instantiates a new AlertResponseCursors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *AlertResponseCursors) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AlertResponseCursors) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AlertResponseCursors) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *AlertResponseCursors) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetSelf

`func (o *AlertResponseCursors) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AlertResponseCursors) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AlertResponseCursors) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AlertResponseCursors) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


