# TransactionStep2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**HasScreenshot** | **bool** |  | 
**HasWaterfall** | **bool** |  | 
**SubSteps** | Pointer to [**[]TransactionSubStep**](TransactionSubStep.md) |  | [optional] 

## Methods

### NewTransactionStep2

`func NewTransactionStep2(hasScreenshot bool, hasWaterfall bool, ) *TransactionStep2`

NewTransactionStep2 instantiates a new TransactionStep2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionStep2WithDefaults

`func NewTransactionStep2WithDefaults() *TransactionStep2`

NewTransactionStep2WithDefaults instantiates a new TransactionStep2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TransactionStep2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransactionStep2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransactionStep2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransactionStep2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHasScreenshot

`func (o *TransactionStep2) GetHasScreenshot() bool`

GetHasScreenshot returns the HasScreenshot field if non-nil, zero value otherwise.

### GetHasScreenshotOk

`func (o *TransactionStep2) GetHasScreenshotOk() (*bool, bool)`

GetHasScreenshotOk returns a tuple with the HasScreenshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasScreenshot

`func (o *TransactionStep2) SetHasScreenshot(v bool)`

SetHasScreenshot sets HasScreenshot field to given value.


### GetHasWaterfall

`func (o *TransactionStep2) GetHasWaterfall() bool`

GetHasWaterfall returns the HasWaterfall field if non-nil, zero value otherwise.

### GetHasWaterfallOk

`func (o *TransactionStep2) GetHasWaterfallOk() (*bool, bool)`

GetHasWaterfallOk returns a tuple with the HasWaterfall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWaterfall

`func (o *TransactionStep2) SetHasWaterfall(v bool)`

SetHasWaterfall sets HasWaterfall field to given value.


### GetSubSteps

`func (o *TransactionStep2) GetSubSteps() []TransactionSubStep`

GetSubSteps returns the SubSteps field if non-nil, zero value otherwise.

### GetSubStepsOk

`func (o *TransactionStep2) GetSubStepsOk() (*[]TransactionSubStep, bool)`

GetSubStepsOk returns a tuple with the SubSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubSteps

`func (o *TransactionStep2) SetSubSteps(v []TransactionSubStep)`

SetSubSteps sets SubSteps field to given value.

### HasSubSteps

`func (o *TransactionStep2) HasSubSteps() bool`

HasSubSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


