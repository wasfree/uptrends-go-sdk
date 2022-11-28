# MsaStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Method** | [**HttpMethods**](HttpMethods.md) |  | 
**Body** | Pointer to **string** | The body that will be send in the request. Only used if BodyType equals Raw | [optional] 
**BodyType** | Pointer to [**MsaBodyType**](MsaBodyType.md) | Determines what kind of body the request will have. | [optional] 
**VaultFileId** | Pointer to **string** | The guid of the vaultfile that will be send in the request. Only used if BodyType equals VaultFiles | [optional] 
**RequestHeaders** | Pointer to [**[]ApiHttpHeaderValue**](ApiHttpHeaderValue.md) |  | [optional] 
**Variables** | Pointer to [**[]ApiVariableDefinition**](ApiVariableDefinition.md) |  | [optional] 
**Assertions** | Pointer to [**[]ApiAssertion**](ApiAssertion.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UseFixedClientCertificate** | **bool** |  | 
**Authentication** | Pointer to [**ApiAuthenticationInfo**](ApiAuthenticationInfo.md) |  | [optional] 
**IgnoreCertificateErrors** | **bool** |  | 
**ClientCertificateVaultItem** | Pointer to **string** |  | [optional] 
**Delay** | **int32** |  | 
**StepType** | [**MsaStepType**](MsaStepType.md) |  | 
**RetryUntilSuccessful** | **bool** |  | 
**MaxAttempts** | **int32** |  | 
**RetryWaitMilliseconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewMsaStep

`func NewMsaStep(method HttpMethods, useFixedClientCertificate bool, ignoreCertificateErrors bool, delay int32, stepType MsaStepType, retryUntilSuccessful bool, maxAttempts int32, ) *MsaStep`

NewMsaStep instantiates a new MsaStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsaStepWithDefaults

`func NewMsaStepWithDefaults() *MsaStep`

NewMsaStepWithDefaults instantiates a new MsaStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *MsaStep) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MsaStep) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MsaStep) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MsaStep) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *MsaStep) GetMethod() HttpMethods`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MsaStep) GetMethodOk() (*HttpMethods, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MsaStep) SetMethod(v HttpMethods)`

SetMethod sets Method field to given value.


### GetBody

`func (o *MsaStep) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MsaStep) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MsaStep) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *MsaStep) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBodyType

`func (o *MsaStep) GetBodyType() MsaBodyType`

GetBodyType returns the BodyType field if non-nil, zero value otherwise.

### GetBodyTypeOk

`func (o *MsaStep) GetBodyTypeOk() (*MsaBodyType, bool)`

GetBodyTypeOk returns a tuple with the BodyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyType

`func (o *MsaStep) SetBodyType(v MsaBodyType)`

SetBodyType sets BodyType field to given value.

### HasBodyType

`func (o *MsaStep) HasBodyType() bool`

HasBodyType returns a boolean if a field has been set.

### GetVaultFileId

`func (o *MsaStep) GetVaultFileId() string`

GetVaultFileId returns the VaultFileId field if non-nil, zero value otherwise.

### GetVaultFileIdOk

`func (o *MsaStep) GetVaultFileIdOk() (*string, bool)`

GetVaultFileIdOk returns a tuple with the VaultFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultFileId

`func (o *MsaStep) SetVaultFileId(v string)`

SetVaultFileId sets VaultFileId field to given value.

### HasVaultFileId

`func (o *MsaStep) HasVaultFileId() bool`

HasVaultFileId returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *MsaStep) GetRequestHeaders() []ApiHttpHeaderValue`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *MsaStep) GetRequestHeadersOk() (*[]ApiHttpHeaderValue, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *MsaStep) SetRequestHeaders(v []ApiHttpHeaderValue)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *MsaStep) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetVariables

`func (o *MsaStep) GetVariables() []ApiVariableDefinition`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *MsaStep) GetVariablesOk() (*[]ApiVariableDefinition, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *MsaStep) SetVariables(v []ApiVariableDefinition)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *MsaStep) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetAssertions

`func (o *MsaStep) GetAssertions() []ApiAssertion`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *MsaStep) GetAssertionsOk() (*[]ApiAssertion, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *MsaStep) SetAssertions(v []ApiAssertion)`

SetAssertions sets Assertions field to given value.

### HasAssertions

`func (o *MsaStep) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.

### GetName

`func (o *MsaStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MsaStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MsaStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MsaStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUseFixedClientCertificate

`func (o *MsaStep) GetUseFixedClientCertificate() bool`

GetUseFixedClientCertificate returns the UseFixedClientCertificate field if non-nil, zero value otherwise.

### GetUseFixedClientCertificateOk

`func (o *MsaStep) GetUseFixedClientCertificateOk() (*bool, bool)`

GetUseFixedClientCertificateOk returns a tuple with the UseFixedClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFixedClientCertificate

`func (o *MsaStep) SetUseFixedClientCertificate(v bool)`

SetUseFixedClientCertificate sets UseFixedClientCertificate field to given value.


### GetAuthentication

`func (o *MsaStep) GetAuthentication() ApiAuthenticationInfo`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *MsaStep) GetAuthenticationOk() (*ApiAuthenticationInfo, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *MsaStep) SetAuthentication(v ApiAuthenticationInfo)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *MsaStep) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetIgnoreCertificateErrors

`func (o *MsaStep) GetIgnoreCertificateErrors() bool`

GetIgnoreCertificateErrors returns the IgnoreCertificateErrors field if non-nil, zero value otherwise.

### GetIgnoreCertificateErrorsOk

`func (o *MsaStep) GetIgnoreCertificateErrorsOk() (*bool, bool)`

GetIgnoreCertificateErrorsOk returns a tuple with the IgnoreCertificateErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCertificateErrors

`func (o *MsaStep) SetIgnoreCertificateErrors(v bool)`

SetIgnoreCertificateErrors sets IgnoreCertificateErrors field to given value.


### GetClientCertificateVaultItem

`func (o *MsaStep) GetClientCertificateVaultItem() string`

GetClientCertificateVaultItem returns the ClientCertificateVaultItem field if non-nil, zero value otherwise.

### GetClientCertificateVaultItemOk

`func (o *MsaStep) GetClientCertificateVaultItemOk() (*string, bool)`

GetClientCertificateVaultItemOk returns a tuple with the ClientCertificateVaultItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateVaultItem

`func (o *MsaStep) SetClientCertificateVaultItem(v string)`

SetClientCertificateVaultItem sets ClientCertificateVaultItem field to given value.

### HasClientCertificateVaultItem

`func (o *MsaStep) HasClientCertificateVaultItem() bool`

HasClientCertificateVaultItem returns a boolean if a field has been set.

### GetDelay

`func (o *MsaStep) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *MsaStep) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *MsaStep) SetDelay(v int32)`

SetDelay sets Delay field to given value.


### GetStepType

`func (o *MsaStep) GetStepType() MsaStepType`

GetStepType returns the StepType field if non-nil, zero value otherwise.

### GetStepTypeOk

`func (o *MsaStep) GetStepTypeOk() (*MsaStepType, bool)`

GetStepTypeOk returns a tuple with the StepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepType

`func (o *MsaStep) SetStepType(v MsaStepType)`

SetStepType sets StepType field to given value.


### GetRetryUntilSuccessful

`func (o *MsaStep) GetRetryUntilSuccessful() bool`

GetRetryUntilSuccessful returns the RetryUntilSuccessful field if non-nil, zero value otherwise.

### GetRetryUntilSuccessfulOk

`func (o *MsaStep) GetRetryUntilSuccessfulOk() (*bool, bool)`

GetRetryUntilSuccessfulOk returns a tuple with the RetryUntilSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUntilSuccessful

`func (o *MsaStep) SetRetryUntilSuccessful(v bool)`

SetRetryUntilSuccessful sets RetryUntilSuccessful field to given value.


### GetMaxAttempts

`func (o *MsaStep) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *MsaStep) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *MsaStep) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.


### GetRetryWaitMilliseconds

`func (o *MsaStep) GetRetryWaitMilliseconds() int32`

GetRetryWaitMilliseconds returns the RetryWaitMilliseconds field if non-nil, zero value otherwise.

### GetRetryWaitMillisecondsOk

`func (o *MsaStep) GetRetryWaitMillisecondsOk() (*int32, bool)`

GetRetryWaitMillisecondsOk returns a tuple with the RetryWaitMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryWaitMilliseconds

`func (o *MsaStep) SetRetryWaitMilliseconds(v int32)`

SetRetryWaitMilliseconds sets RetryWaitMilliseconds field to given value.

### HasRetryWaitMilliseconds

`func (o *MsaStep) HasRetryWaitMilliseconds() bool`

HasRetryWaitMilliseconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


