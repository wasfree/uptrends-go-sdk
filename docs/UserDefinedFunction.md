# UserDefinedFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | [**UserDefinedFunctionType**](UserDefinedFunctionType.md) |  | 
**Mappings** | Pointer to [**[]UserDefinedFunctionMapping**](UserDefinedFunctionMapping.md) |  | [optional] 
**Regex** | Pointer to **string** |  | [optional] 
**JwtSigningKey** | Pointer to **string** | This property is not supported yet | [optional] 
**JwtAlgorithm** | Pointer to [**JwtAlgorithm**](JwtAlgorithm.md) | This property is not supported yet | [optional] 

## Methods

### NewUserDefinedFunction

`func NewUserDefinedFunction(type_ UserDefinedFunctionType, ) *UserDefinedFunction`

NewUserDefinedFunction instantiates a new UserDefinedFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedFunctionWithDefaults

`func NewUserDefinedFunctionWithDefaults() *UserDefinedFunction`

NewUserDefinedFunctionWithDefaults instantiates a new UserDefinedFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserDefinedFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDefinedFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDefinedFunction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserDefinedFunction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UserDefinedFunction) GetType() UserDefinedFunctionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDefinedFunction) GetTypeOk() (*UserDefinedFunctionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDefinedFunction) SetType(v UserDefinedFunctionType)`

SetType sets Type field to given value.


### GetMappings

`func (o *UserDefinedFunction) GetMappings() []UserDefinedFunctionMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *UserDefinedFunction) GetMappingsOk() (*[]UserDefinedFunctionMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *UserDefinedFunction) SetMappings(v []UserDefinedFunctionMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *UserDefinedFunction) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetRegex

`func (o *UserDefinedFunction) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *UserDefinedFunction) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *UserDefinedFunction) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *UserDefinedFunction) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetJwtSigningKey

`func (o *UserDefinedFunction) GetJwtSigningKey() string`

GetJwtSigningKey returns the JwtSigningKey field if non-nil, zero value otherwise.

### GetJwtSigningKeyOk

`func (o *UserDefinedFunction) GetJwtSigningKeyOk() (*string, bool)`

GetJwtSigningKeyOk returns a tuple with the JwtSigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSigningKey

`func (o *UserDefinedFunction) SetJwtSigningKey(v string)`

SetJwtSigningKey sets JwtSigningKey field to given value.

### HasJwtSigningKey

`func (o *UserDefinedFunction) HasJwtSigningKey() bool`

HasJwtSigningKey returns a boolean if a field has been set.

### GetJwtAlgorithm

`func (o *UserDefinedFunction) GetJwtAlgorithm() JwtAlgorithm`

GetJwtAlgorithm returns the JwtAlgorithm field if non-nil, zero value otherwise.

### GetJwtAlgorithmOk

`func (o *UserDefinedFunction) GetJwtAlgorithmOk() (*JwtAlgorithm, bool)`

GetJwtAlgorithmOk returns a tuple with the JwtAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAlgorithm

`func (o *UserDefinedFunction) SetJwtAlgorithm(v JwtAlgorithm)`

SetJwtAlgorithm sets JwtAlgorithm field to given value.

### HasJwtAlgorithm

`func (o *UserDefinedFunction) HasJwtAlgorithm() bool`

HasJwtAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


