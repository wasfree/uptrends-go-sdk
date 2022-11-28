# ServerHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckpointName** | Pointer to **string** | The name of the checkpoint | [optional] 
**Status** | Pointer to **string** | The status of the server | [optional] 
**StatusDetails** | Pointer to [**ServerHealthStatusDetails**](ServerHealthStatusDetails.md) |  | [optional] 

## Methods

### NewServerHealth

`func NewServerHealth() *ServerHealth`

NewServerHealth instantiates a new ServerHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerHealthWithDefaults

`func NewServerHealthWithDefaults() *ServerHealth`

NewServerHealthWithDefaults instantiates a new ServerHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckpointName

`func (o *ServerHealth) GetCheckpointName() string`

GetCheckpointName returns the CheckpointName field if non-nil, zero value otherwise.

### GetCheckpointNameOk

`func (o *ServerHealth) GetCheckpointNameOk() (*string, bool)`

GetCheckpointNameOk returns a tuple with the CheckpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointName

`func (o *ServerHealth) SetCheckpointName(v string)`

SetCheckpointName sets CheckpointName field to given value.

### HasCheckpointName

`func (o *ServerHealth) HasCheckpointName() bool`

HasCheckpointName returns a boolean if a field has been set.

### GetStatus

`func (o *ServerHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *ServerHealth) GetStatusDetails() ServerHealthStatusDetails`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *ServerHealth) GetStatusDetailsOk() (*ServerHealthStatusDetails, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *ServerHealth) SetStatusDetails(v ServerHealthStatusDetails)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *ServerHealth) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


