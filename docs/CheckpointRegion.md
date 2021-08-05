# CheckpointRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**ChildRegions** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewCheckpointRegion

`func NewCheckpointRegion(id int32, ) *CheckpointRegion`

NewCheckpointRegion instantiates a new CheckpointRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointRegionWithDefaults

`func NewCheckpointRegionWithDefaults() *CheckpointRegion`

NewCheckpointRegionWithDefaults instantiates a new CheckpointRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckpointRegion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckpointRegion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckpointRegion) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CheckpointRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckpointRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckpointRegion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckpointRegion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChildRegions

`func (o *CheckpointRegion) GetChildRegions() []int32`

GetChildRegions returns the ChildRegions field if non-nil, zero value otherwise.

### GetChildRegionsOk

`func (o *CheckpointRegion) GetChildRegionsOk() (*[]int32, bool)`

GetChildRegionsOk returns a tuple with the ChildRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildRegions

`func (o *CheckpointRegion) SetChildRegions(v []int32)`

SetChildRegions sets ChildRegions field to given value.

### HasChildRegions

`func (o *CheckpointRegion) HasChildRegions() bool`

HasChildRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


