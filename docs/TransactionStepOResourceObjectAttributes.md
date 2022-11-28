# TransactionStepOResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepNumber** | **int32** | Step (index) number | 
**StepName** | Pointer to **string** | The name of the step | [optional] 
**TotalTime** | **float64** | Number of milliseconds it took for this step to succeed | 
**Elements** | Pointer to **[]string** | Text representation of the step element results | [optional] 
**HasError** | **bool** | Did this step result in an error? | 

## Methods

### NewTransactionStepOResourceObjectAttributes

`func NewTransactionStepOResourceObjectAttributes(stepNumber int32, totalTime float64, hasError bool, ) *TransactionStepOResourceObjectAttributes`

NewTransactionStepOResourceObjectAttributes instantiates a new TransactionStepOResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionStepOResourceObjectAttributesWithDefaults

`func NewTransactionStepOResourceObjectAttributesWithDefaults() *TransactionStepOResourceObjectAttributes`

NewTransactionStepOResourceObjectAttributesWithDefaults instantiates a new TransactionStepOResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepNumber

`func (o *TransactionStepOResourceObjectAttributes) GetStepNumber() int32`

GetStepNumber returns the StepNumber field if non-nil, zero value otherwise.

### GetStepNumberOk

`func (o *TransactionStepOResourceObjectAttributes) GetStepNumberOk() (*int32, bool)`

GetStepNumberOk returns a tuple with the StepNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepNumber

`func (o *TransactionStepOResourceObjectAttributes) SetStepNumber(v int32)`

SetStepNumber sets StepNumber field to given value.


### GetStepName

`func (o *TransactionStepOResourceObjectAttributes) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *TransactionStepOResourceObjectAttributes) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *TransactionStepOResourceObjectAttributes) SetStepName(v string)`

SetStepName sets StepName field to given value.

### HasStepName

`func (o *TransactionStepOResourceObjectAttributes) HasStepName() bool`

HasStepName returns a boolean if a field has been set.

### GetTotalTime

`func (o *TransactionStepOResourceObjectAttributes) GetTotalTime() float64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *TransactionStepOResourceObjectAttributes) GetTotalTimeOk() (*float64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *TransactionStepOResourceObjectAttributes) SetTotalTime(v float64)`

SetTotalTime sets TotalTime field to given value.


### GetElements

`func (o *TransactionStepOResourceObjectAttributes) GetElements() []string`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *TransactionStepOResourceObjectAttributes) GetElementsOk() (*[]string, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *TransactionStepOResourceObjectAttributes) SetElements(v []string)`

SetElements sets Elements field to given value.

### HasElements

`func (o *TransactionStepOResourceObjectAttributes) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetHasError

`func (o *TransactionStepOResourceObjectAttributes) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *TransactionStepOResourceObjectAttributes) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *TransactionStepOResourceObjectAttributes) SetHasError(v bool)`

SetHasError sets HasError field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


