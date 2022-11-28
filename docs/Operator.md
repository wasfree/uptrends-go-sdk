# Operator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatorGuid** | Pointer to **string** | The unique identifier of this operator | [optional] 
**Hash** | Pointer to **string** | The hash of this operator. | [optional] 
**Password** | Pointer to **string** | The password is a required field if AllowNativeLogin is set to True | [optional] 
**FullName** | Pointer to **string** | The full name of this operator | [optional] 
**Email** | Pointer to **string** | The email address of this operator. This also serves as the username | [optional] 
**MobilePhone** | Pointer to **string** | The phone number of this operator to which SMS and phone alerts can be sent. Start with a plus (+) sign and your country code | [optional] 
**OutgoingPhoneNumberId** | Pointer to **int32** | The id of the phone number that will be used to send phone alerts (See /OutgoingPhoneNumber API under Miscellaneous for available ids) | [optional] 
**OutgoingPhoneNumberIdSpecified** | Pointer to **bool** |  | [optional] 
**IsAccountAdministrator** | Pointer to **bool** | This indicates if the operator is the account administrator. | [optional] 
**BackupEmail** | Pointer to **string** | The backup email address of this operator | [optional] 
**IsOnDuty** | Pointer to **bool** | Indicates whether the operator is currently active | [optional] 
**CultureName** | Pointer to **string** | If ommitted the operator will use the account culture. If set it will override the account default | [optional] 
**CultureNameSpecified** | Pointer to **bool** |  | [optional] 
**TimeZoneId** | Pointer to **int32** | The id of the timezone of this operator (See /Timezone API under Miscellaneous for available timezones) | [optional] 
**TimeZoneIdSpecified** | Pointer to **bool** |  | [optional] 
**SmsProvider** | Pointer to [**SmsProvider**](SmsProvider.md) | The SMS provider used to send out SMS alerts | [optional] 
**UseNumericSender** | Pointer to **bool** | Set to True to override the default behavior of sending SMS alerts with textual sender ID | [optional] 
**UseNumericSenderSpecified** | Pointer to **bool** |  | [optional] 
**AllowNativeLogin** | Pointer to **bool** | This can only be set to false if the account has SSO enabled. Ommitting or providing null will use the account default | [optional] 
**AllowNativeLoginSpecified** | Pointer to **bool** |  | [optional] 
**AllowSingleSignon** | Pointer to **bool** | This can only be set to true if the account has SSO enabled. Ommitting or providing null will use the account default | [optional] 
**AllowSingleSignonSpecified** | Pointer to **bool** |  | [optional] 
**DefaultDashboard** | Pointer to **string** | This is used to set the default dashboard for the operator.  Valid options are: UseAccountSpecifiedDashboard (This will use the dashboard specified for the account) Any built-in dashboard: e.g. AccountOverview Any custom dashboard to which the operator has access to, defined by the guid of this dashboard | [optional] 
**SetupMode** | Pointer to [**OperatorSetupMode**](OperatorSetupMode.md) |  | [optional] 

## Methods

### NewOperator

`func NewOperator() *Operator`

NewOperator instantiates a new Operator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorWithDefaults

`func NewOperatorWithDefaults() *Operator`

NewOperatorWithDefaults instantiates a new Operator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatorGuid

`func (o *Operator) GetOperatorGuid() string`

GetOperatorGuid returns the OperatorGuid field if non-nil, zero value otherwise.

### GetOperatorGuidOk

`func (o *Operator) GetOperatorGuidOk() (*string, bool)`

GetOperatorGuidOk returns a tuple with the OperatorGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorGuid

`func (o *Operator) SetOperatorGuid(v string)`

SetOperatorGuid sets OperatorGuid field to given value.

### HasOperatorGuid

`func (o *Operator) HasOperatorGuid() bool`

HasOperatorGuid returns a boolean if a field has been set.

### GetHash

`func (o *Operator) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Operator) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Operator) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Operator) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetPassword

`func (o *Operator) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Operator) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Operator) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Operator) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetFullName

`func (o *Operator) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Operator) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Operator) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *Operator) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetEmail

`func (o *Operator) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Operator) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Operator) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Operator) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMobilePhone

`func (o *Operator) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *Operator) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *Operator) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *Operator) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetOutgoingPhoneNumberId

`func (o *Operator) GetOutgoingPhoneNumberId() int32`

GetOutgoingPhoneNumberId returns the OutgoingPhoneNumberId field if non-nil, zero value otherwise.

### GetOutgoingPhoneNumberIdOk

`func (o *Operator) GetOutgoingPhoneNumberIdOk() (*int32, bool)`

GetOutgoingPhoneNumberIdOk returns a tuple with the OutgoingPhoneNumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingPhoneNumberId

`func (o *Operator) SetOutgoingPhoneNumberId(v int32)`

SetOutgoingPhoneNumberId sets OutgoingPhoneNumberId field to given value.

### HasOutgoingPhoneNumberId

`func (o *Operator) HasOutgoingPhoneNumberId() bool`

HasOutgoingPhoneNumberId returns a boolean if a field has been set.

### GetOutgoingPhoneNumberIdSpecified

`func (o *Operator) GetOutgoingPhoneNumberIdSpecified() bool`

GetOutgoingPhoneNumberIdSpecified returns the OutgoingPhoneNumberIdSpecified field if non-nil, zero value otherwise.

### GetOutgoingPhoneNumberIdSpecifiedOk

`func (o *Operator) GetOutgoingPhoneNumberIdSpecifiedOk() (*bool, bool)`

GetOutgoingPhoneNumberIdSpecifiedOk returns a tuple with the OutgoingPhoneNumberIdSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingPhoneNumberIdSpecified

`func (o *Operator) SetOutgoingPhoneNumberIdSpecified(v bool)`

SetOutgoingPhoneNumberIdSpecified sets OutgoingPhoneNumberIdSpecified field to given value.

### HasOutgoingPhoneNumberIdSpecified

`func (o *Operator) HasOutgoingPhoneNumberIdSpecified() bool`

HasOutgoingPhoneNumberIdSpecified returns a boolean if a field has been set.

### GetIsAccountAdministrator

`func (o *Operator) GetIsAccountAdministrator() bool`

GetIsAccountAdministrator returns the IsAccountAdministrator field if non-nil, zero value otherwise.

### GetIsAccountAdministratorOk

`func (o *Operator) GetIsAccountAdministratorOk() (*bool, bool)`

GetIsAccountAdministratorOk returns a tuple with the IsAccountAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountAdministrator

`func (o *Operator) SetIsAccountAdministrator(v bool)`

SetIsAccountAdministrator sets IsAccountAdministrator field to given value.

### HasIsAccountAdministrator

`func (o *Operator) HasIsAccountAdministrator() bool`

HasIsAccountAdministrator returns a boolean if a field has been set.

### GetBackupEmail

`func (o *Operator) GetBackupEmail() string`

GetBackupEmail returns the BackupEmail field if non-nil, zero value otherwise.

### GetBackupEmailOk

`func (o *Operator) GetBackupEmailOk() (*string, bool)`

GetBackupEmailOk returns a tuple with the BackupEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEmail

`func (o *Operator) SetBackupEmail(v string)`

SetBackupEmail sets BackupEmail field to given value.

### HasBackupEmail

`func (o *Operator) HasBackupEmail() bool`

HasBackupEmail returns a boolean if a field has been set.

### GetIsOnDuty

`func (o *Operator) GetIsOnDuty() bool`

GetIsOnDuty returns the IsOnDuty field if non-nil, zero value otherwise.

### GetIsOnDutyOk

`func (o *Operator) GetIsOnDutyOk() (*bool, bool)`

GetIsOnDutyOk returns a tuple with the IsOnDuty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnDuty

`func (o *Operator) SetIsOnDuty(v bool)`

SetIsOnDuty sets IsOnDuty field to given value.

### HasIsOnDuty

`func (o *Operator) HasIsOnDuty() bool`

HasIsOnDuty returns a boolean if a field has been set.

### GetCultureName

`func (o *Operator) GetCultureName() string`

GetCultureName returns the CultureName field if non-nil, zero value otherwise.

### GetCultureNameOk

`func (o *Operator) GetCultureNameOk() (*string, bool)`

GetCultureNameOk returns a tuple with the CultureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultureName

`func (o *Operator) SetCultureName(v string)`

SetCultureName sets CultureName field to given value.

### HasCultureName

`func (o *Operator) HasCultureName() bool`

HasCultureName returns a boolean if a field has been set.

### GetCultureNameSpecified

`func (o *Operator) GetCultureNameSpecified() bool`

GetCultureNameSpecified returns the CultureNameSpecified field if non-nil, zero value otherwise.

### GetCultureNameSpecifiedOk

`func (o *Operator) GetCultureNameSpecifiedOk() (*bool, bool)`

GetCultureNameSpecifiedOk returns a tuple with the CultureNameSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultureNameSpecified

`func (o *Operator) SetCultureNameSpecified(v bool)`

SetCultureNameSpecified sets CultureNameSpecified field to given value.

### HasCultureNameSpecified

`func (o *Operator) HasCultureNameSpecified() bool`

HasCultureNameSpecified returns a boolean if a field has been set.

### GetTimeZoneId

`func (o *Operator) GetTimeZoneId() int32`

GetTimeZoneId returns the TimeZoneId field if non-nil, zero value otherwise.

### GetTimeZoneIdOk

`func (o *Operator) GetTimeZoneIdOk() (*int32, bool)`

GetTimeZoneIdOk returns a tuple with the TimeZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneId

`func (o *Operator) SetTimeZoneId(v int32)`

SetTimeZoneId sets TimeZoneId field to given value.

### HasTimeZoneId

`func (o *Operator) HasTimeZoneId() bool`

HasTimeZoneId returns a boolean if a field has been set.

### GetTimeZoneIdSpecified

`func (o *Operator) GetTimeZoneIdSpecified() bool`

GetTimeZoneIdSpecified returns the TimeZoneIdSpecified field if non-nil, zero value otherwise.

### GetTimeZoneIdSpecifiedOk

`func (o *Operator) GetTimeZoneIdSpecifiedOk() (*bool, bool)`

GetTimeZoneIdSpecifiedOk returns a tuple with the TimeZoneIdSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneIdSpecified

`func (o *Operator) SetTimeZoneIdSpecified(v bool)`

SetTimeZoneIdSpecified sets TimeZoneIdSpecified field to given value.

### HasTimeZoneIdSpecified

`func (o *Operator) HasTimeZoneIdSpecified() bool`

HasTimeZoneIdSpecified returns a boolean if a field has been set.

### GetSmsProvider

`func (o *Operator) GetSmsProvider() SmsProvider`

GetSmsProvider returns the SmsProvider field if non-nil, zero value otherwise.

### GetSmsProviderOk

`func (o *Operator) GetSmsProviderOk() (*SmsProvider, bool)`

GetSmsProviderOk returns a tuple with the SmsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsProvider

`func (o *Operator) SetSmsProvider(v SmsProvider)`

SetSmsProvider sets SmsProvider field to given value.

### HasSmsProvider

`func (o *Operator) HasSmsProvider() bool`

HasSmsProvider returns a boolean if a field has been set.

### GetUseNumericSender

`func (o *Operator) GetUseNumericSender() bool`

GetUseNumericSender returns the UseNumericSender field if non-nil, zero value otherwise.

### GetUseNumericSenderOk

`func (o *Operator) GetUseNumericSenderOk() (*bool, bool)`

GetUseNumericSenderOk returns a tuple with the UseNumericSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumericSender

`func (o *Operator) SetUseNumericSender(v bool)`

SetUseNumericSender sets UseNumericSender field to given value.

### HasUseNumericSender

`func (o *Operator) HasUseNumericSender() bool`

HasUseNumericSender returns a boolean if a field has been set.

### GetUseNumericSenderSpecified

`func (o *Operator) GetUseNumericSenderSpecified() bool`

GetUseNumericSenderSpecified returns the UseNumericSenderSpecified field if non-nil, zero value otherwise.

### GetUseNumericSenderSpecifiedOk

`func (o *Operator) GetUseNumericSenderSpecifiedOk() (*bool, bool)`

GetUseNumericSenderSpecifiedOk returns a tuple with the UseNumericSenderSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumericSenderSpecified

`func (o *Operator) SetUseNumericSenderSpecified(v bool)`

SetUseNumericSenderSpecified sets UseNumericSenderSpecified field to given value.

### HasUseNumericSenderSpecified

`func (o *Operator) HasUseNumericSenderSpecified() bool`

HasUseNumericSenderSpecified returns a boolean if a field has been set.

### GetAllowNativeLogin

`func (o *Operator) GetAllowNativeLogin() bool`

GetAllowNativeLogin returns the AllowNativeLogin field if non-nil, zero value otherwise.

### GetAllowNativeLoginOk

`func (o *Operator) GetAllowNativeLoginOk() (*bool, bool)`

GetAllowNativeLoginOk returns a tuple with the AllowNativeLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNativeLogin

`func (o *Operator) SetAllowNativeLogin(v bool)`

SetAllowNativeLogin sets AllowNativeLogin field to given value.

### HasAllowNativeLogin

`func (o *Operator) HasAllowNativeLogin() bool`

HasAllowNativeLogin returns a boolean if a field has been set.

### GetAllowNativeLoginSpecified

`func (o *Operator) GetAllowNativeLoginSpecified() bool`

GetAllowNativeLoginSpecified returns the AllowNativeLoginSpecified field if non-nil, zero value otherwise.

### GetAllowNativeLoginSpecifiedOk

`func (o *Operator) GetAllowNativeLoginSpecifiedOk() (*bool, bool)`

GetAllowNativeLoginSpecifiedOk returns a tuple with the AllowNativeLoginSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNativeLoginSpecified

`func (o *Operator) SetAllowNativeLoginSpecified(v bool)`

SetAllowNativeLoginSpecified sets AllowNativeLoginSpecified field to given value.

### HasAllowNativeLoginSpecified

`func (o *Operator) HasAllowNativeLoginSpecified() bool`

HasAllowNativeLoginSpecified returns a boolean if a field has been set.

### GetAllowSingleSignon

`func (o *Operator) GetAllowSingleSignon() bool`

GetAllowSingleSignon returns the AllowSingleSignon field if non-nil, zero value otherwise.

### GetAllowSingleSignonOk

`func (o *Operator) GetAllowSingleSignonOk() (*bool, bool)`

GetAllowSingleSignonOk returns a tuple with the AllowSingleSignon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSingleSignon

`func (o *Operator) SetAllowSingleSignon(v bool)`

SetAllowSingleSignon sets AllowSingleSignon field to given value.

### HasAllowSingleSignon

`func (o *Operator) HasAllowSingleSignon() bool`

HasAllowSingleSignon returns a boolean if a field has been set.

### GetAllowSingleSignonSpecified

`func (o *Operator) GetAllowSingleSignonSpecified() bool`

GetAllowSingleSignonSpecified returns the AllowSingleSignonSpecified field if non-nil, zero value otherwise.

### GetAllowSingleSignonSpecifiedOk

`func (o *Operator) GetAllowSingleSignonSpecifiedOk() (*bool, bool)`

GetAllowSingleSignonSpecifiedOk returns a tuple with the AllowSingleSignonSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSingleSignonSpecified

`func (o *Operator) SetAllowSingleSignonSpecified(v bool)`

SetAllowSingleSignonSpecified sets AllowSingleSignonSpecified field to given value.

### HasAllowSingleSignonSpecified

`func (o *Operator) HasAllowSingleSignonSpecified() bool`

HasAllowSingleSignonSpecified returns a boolean if a field has been set.

### GetDefaultDashboard

`func (o *Operator) GetDefaultDashboard() string`

GetDefaultDashboard returns the DefaultDashboard field if non-nil, zero value otherwise.

### GetDefaultDashboardOk

`func (o *Operator) GetDefaultDashboardOk() (*string, bool)`

GetDefaultDashboardOk returns a tuple with the DefaultDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDashboard

`func (o *Operator) SetDefaultDashboard(v string)`

SetDefaultDashboard sets DefaultDashboard field to given value.

### HasDefaultDashboard

`func (o *Operator) HasDefaultDashboard() bool`

HasDefaultDashboard returns a boolean if a field has been set.

### GetSetupMode

`func (o *Operator) GetSetupMode() OperatorSetupMode`

GetSetupMode returns the SetupMode field if non-nil, zero value otherwise.

### GetSetupModeOk

`func (o *Operator) GetSetupModeOk() (*OperatorSetupMode, bool)`

GetSetupModeOk returns a tuple with the SetupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupMode

`func (o *Operator) SetSetupMode(v OperatorSetupMode)`

SetSetupMode sets SetupMode field to given value.

### HasSetupMode

`func (o *Operator) HasSetupMode() bool`

HasSetupMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


