# ApiAuthenticationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AuthenticationType** | [**ApiHttpAuthenticationType**](ApiHttpAuthenticationType.md) |  | 
**UserName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordSpecified** | **bool** |  | 

## Methods

### NewApiAuthenticationInfo

`func NewApiAuthenticationInfo(id string, authenticationType ApiHttpAuthenticationType, passwordSpecified bool, ) *ApiAuthenticationInfo`

NewApiAuthenticationInfo instantiates a new ApiAuthenticationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAuthenticationInfoWithDefaults

`func NewApiAuthenticationInfoWithDefaults() *ApiAuthenticationInfo`

NewApiAuthenticationInfoWithDefaults instantiates a new ApiAuthenticationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAuthenticationInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAuthenticationInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAuthenticationInfo) SetId(v string)`

SetId sets Id field to given value.


### GetAuthenticationType

`func (o *ApiAuthenticationInfo) GetAuthenticationType() ApiHttpAuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *ApiAuthenticationInfo) GetAuthenticationTypeOk() (*ApiHttpAuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *ApiAuthenticationInfo) SetAuthenticationType(v ApiHttpAuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetUserName

`func (o *ApiAuthenticationInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ApiAuthenticationInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ApiAuthenticationInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ApiAuthenticationInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *ApiAuthenticationInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiAuthenticationInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiAuthenticationInfo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiAuthenticationInfo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordSpecified

`func (o *ApiAuthenticationInfo) GetPasswordSpecified() bool`

GetPasswordSpecified returns the PasswordSpecified field if non-nil, zero value otherwise.

### GetPasswordSpecifiedOk

`func (o *ApiAuthenticationInfo) GetPasswordSpecifiedOk() (*bool, bool)`

GetPasswordSpecifiedOk returns a tuple with the PasswordSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSpecified

`func (o *ApiAuthenticationInfo) SetPasswordSpecified(v bool)`

SetPasswordSpecified sets PasswordSpecified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


