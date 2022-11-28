# TransactionCheckDetailsAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepResults** | Pointer to [**[]TransactionStepOResourceObject**](TransactionStepOResourceObject.md) | Results of the transaction steps | [optional] 
**ResponseBody** | Pointer to **string** | The HTML code returned in case of errors | [optional] 

## Methods

### NewTransactionCheckDetailsAttributes

`func NewTransactionCheckDetailsAttributes() *TransactionCheckDetailsAttributes`

NewTransactionCheckDetailsAttributes instantiates a new TransactionCheckDetailsAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCheckDetailsAttributesWithDefaults

`func NewTransactionCheckDetailsAttributesWithDefaults() *TransactionCheckDetailsAttributes`

NewTransactionCheckDetailsAttributesWithDefaults instantiates a new TransactionCheckDetailsAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepResults

`func (o *TransactionCheckDetailsAttributes) GetStepResults() []TransactionStepOResourceObject`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *TransactionCheckDetailsAttributes) GetStepResultsOk() (*[]TransactionStepOResourceObject, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *TransactionCheckDetailsAttributes) SetStepResults(v []TransactionStepOResourceObject)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *TransactionCheckDetailsAttributes) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.

### GetResponseBody

`func (o *TransactionCheckDetailsAttributes) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *TransactionCheckDetailsAttributes) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *TransactionCheckDetailsAttributes) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *TransactionCheckDetailsAttributes) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


