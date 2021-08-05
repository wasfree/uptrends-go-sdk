# Recipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operators** | Pointer to **[]string** |  | [optional] 
**OperatorGroups** | Pointer to **[]string** |  | [optional] 
**ExtraEmailAddresses** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRecipients

`func NewRecipients() *Recipients`

NewRecipients instantiates a new Recipients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientsWithDefaults

`func NewRecipientsWithDefaults() *Recipients`

NewRecipientsWithDefaults instantiates a new Recipients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperators

`func (o *Recipients) GetOperators() []string`

GetOperators returns the Operators field if non-nil, zero value otherwise.

### GetOperatorsOk

`func (o *Recipients) GetOperatorsOk() (*[]string, bool)`

GetOperatorsOk returns a tuple with the Operators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperators

`func (o *Recipients) SetOperators(v []string)`

SetOperators sets Operators field to given value.

### HasOperators

`func (o *Recipients) HasOperators() bool`

HasOperators returns a boolean if a field has been set.

### GetOperatorGroups

`func (o *Recipients) GetOperatorGroups() []string`

GetOperatorGroups returns the OperatorGroups field if non-nil, zero value otherwise.

### GetOperatorGroupsOk

`func (o *Recipients) GetOperatorGroupsOk() (*[]string, bool)`

GetOperatorGroupsOk returns a tuple with the OperatorGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorGroups

`func (o *Recipients) SetOperatorGroups(v []string)`

SetOperatorGroups sets OperatorGroups field to given value.

### HasOperatorGroups

`func (o *Recipients) HasOperatorGroups() bool`

HasOperatorGroups returns a boolean if a field has been set.

### GetExtraEmailAddresses

`func (o *Recipients) GetExtraEmailAddresses() []string`

GetExtraEmailAddresses returns the ExtraEmailAddresses field if non-nil, zero value otherwise.

### GetExtraEmailAddressesOk

`func (o *Recipients) GetExtraEmailAddressesOk() (*[]string, bool)`

GetExtraEmailAddressesOk returns a tuple with the ExtraEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraEmailAddresses

`func (o *Recipients) SetExtraEmailAddresses(v []string)`

SetExtraEmailAddresses sets ExtraEmailAddresses field to given value.

### HasExtraEmailAddresses

`func (o *Recipients) HasExtraEmailAddresses() bool`

HasExtraEmailAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


