# ServerStatusDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastOnlineUtc** | Pointer to **map[string]interface{}** | The UTC time the server was last online | [optional] 

## Methods

### NewServerStatusDetails

`func NewServerStatusDetails() *ServerStatusDetails`

NewServerStatusDetails instantiates a new ServerStatusDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStatusDetailsWithDefaults

`func NewServerStatusDetailsWithDefaults() *ServerStatusDetails`

NewServerStatusDetailsWithDefaults instantiates a new ServerStatusDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastOnlineUtc

`func (o *ServerStatusDetails) GetLastOnlineUtc() map[string]interface{}`

GetLastOnlineUtc returns the LastOnlineUtc field if non-nil, zero value otherwise.

### GetLastOnlineUtcOk

`func (o *ServerStatusDetails) GetLastOnlineUtcOk() (*map[string]interface{}, bool)`

GetLastOnlineUtcOk returns a tuple with the LastOnlineUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnlineUtc

`func (o *ServerStatusDetails) SetLastOnlineUtc(v map[string]interface{})`

SetLastOnlineUtc sets LastOnlineUtc field to given value.

### HasLastOnlineUtc

`func (o *ServerStatusDetails) HasLastOnlineUtc() bool`

HasLastOnlineUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


