# RegistrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** | The username of the new API account | [optional] 
**Password** | Pointer to **string** | The password of the new API account | [optional] 
**AccountId** | Pointer to **string** | The Uptrends Account Id for which the new API account was created | [optional] 
**OperatorName** | Pointer to **string** | The Uptrends Operator on behalf of whom the new API account was created | [optional] 
**Status** | [**RegisterStatus**](RegisterStatus.md) |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewRegistrationResponse

`func NewRegistrationResponse(status RegisterStatus, ) *RegistrationResponse`

NewRegistrationResponse instantiates a new RegistrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationResponseWithDefaults

`func NewRegistrationResponseWithDefaults() *RegistrationResponse`

NewRegistrationResponseWithDefaults instantiates a new RegistrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *RegistrationResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *RegistrationResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *RegistrationResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *RegistrationResponse) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *RegistrationResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegistrationResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegistrationResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegistrationResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAccountId

`func (o *RegistrationResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RegistrationResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RegistrationResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *RegistrationResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetOperatorName

`func (o *RegistrationResponse) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *RegistrationResponse) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *RegistrationResponse) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.

### HasOperatorName

`func (o *RegistrationResponse) HasOperatorName() bool`

HasOperatorName returns a boolean if a field has been set.

### GetStatus

`func (o *RegistrationResponse) GetStatus() RegisterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegistrationResponse) GetStatusOk() (*RegisterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegistrationResponse) SetStatus(v RegisterStatus)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *RegistrationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RegistrationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RegistrationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RegistrationResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


