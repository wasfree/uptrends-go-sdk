# ApiVariableDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**ApiVariableSourceType**](ApiVariableSourceType.md) |  | 
**Property** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Arguments** | Pointer to [**[]ApiVariableDefinition**](ApiVariableDefinition.md) |  | [optional] 

## Methods

### NewApiVariableDefinition

`func NewApiVariableDefinition(source ApiVariableSourceType, ) *ApiVariableDefinition`

NewApiVariableDefinition instantiates a new ApiVariableDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiVariableDefinitionWithDefaults

`func NewApiVariableDefinitionWithDefaults() *ApiVariableDefinition`

NewApiVariableDefinitionWithDefaults instantiates a new ApiVariableDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ApiVariableDefinition) GetSource() ApiVariableSourceType`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiVariableDefinition) GetSourceOk() (*ApiVariableSourceType, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiVariableDefinition) SetSource(v ApiVariableSourceType)`

SetSource sets Source field to given value.


### GetProperty

`func (o *ApiVariableDefinition) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ApiVariableDefinition) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ApiVariableDefinition) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ApiVariableDefinition) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetName

`func (o *ApiVariableDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiVariableDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiVariableDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiVariableDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArguments

`func (o *ApiVariableDefinition) GetArguments() []ApiVariableDefinition`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ApiVariableDefinition) GetArgumentsOk() (*[]ApiVariableDefinition, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ApiVariableDefinition) SetArguments(v []ApiVariableDefinition)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ApiVariableDefinition) HasArguments() bool`

HasArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


