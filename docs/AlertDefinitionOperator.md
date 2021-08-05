# AlertDefinitionOperator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertDefinition** | **string** |  | 
**Escalationlevel** | **int32** |  | 
**Operator** | **string** |  | 

## Methods

### NewAlertDefinitionOperator

`func NewAlertDefinitionOperator(alertDefinition string, escalationlevel int32, operator string, ) *AlertDefinitionOperator`

NewAlertDefinitionOperator instantiates a new AlertDefinitionOperator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertDefinitionOperatorWithDefaults

`func NewAlertDefinitionOperatorWithDefaults() *AlertDefinitionOperator`

NewAlertDefinitionOperatorWithDefaults instantiates a new AlertDefinitionOperator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertDefinition

`func (o *AlertDefinitionOperator) GetAlertDefinition() string`

GetAlertDefinition returns the AlertDefinition field if non-nil, zero value otherwise.

### GetAlertDefinitionOk

`func (o *AlertDefinitionOperator) GetAlertDefinitionOk() (*string, bool)`

GetAlertDefinitionOk returns a tuple with the AlertDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDefinition

`func (o *AlertDefinitionOperator) SetAlertDefinition(v string)`

SetAlertDefinition sets AlertDefinition field to given value.


### GetEscalationlevel

`func (o *AlertDefinitionOperator) GetEscalationlevel() int32`

GetEscalationlevel returns the Escalationlevel field if non-nil, zero value otherwise.

### GetEscalationlevelOk

`func (o *AlertDefinitionOperator) GetEscalationlevelOk() (*int32, bool)`

GetEscalationlevelOk returns a tuple with the Escalationlevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalationlevel

`func (o *AlertDefinitionOperator) SetEscalationlevel(v int32)`

SetEscalationlevel sets Escalationlevel field to given value.


### GetOperator

`func (o *AlertDefinitionOperator) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AlertDefinitionOperator) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AlertDefinitionOperator) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


