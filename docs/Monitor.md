# Monitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorGuid** | Pointer to **string** | The unique key of this monitor. | [optional] 
**Name** | Pointer to **string** | The name of this monitor. | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this monitor is actively running according to the monitoring interval. | [optional] [default to true]
**Hash** | Pointer to **string** | Hash corresponding with this monitor. | [optional] 
**GenerateAlert** | Pointer to **bool** | Indicates whether this monitor should generate alerts. | [optional] [default to true]
**IsLocked** | Pointer to **bool** | Indicates whether this monitor is locked. | [optional] 
**CheckInterval** | Pointer to **int32** | Indicates the interval in seconds | [optional] 
**Credits** | Pointer to **int32** |  | [optional] 
**MonitorMode** | Pointer to [**MonitorMode**](MonitorMode.md) |  | [optional] 
**PredefinedVariables** | Pointer to [**[]PredefinedVariable**](PredefinedVariable.md) |  | [optional] 
**MsaSteps** | Pointer to [**[]MsaStep**](MsaStep.md) |  | [optional] 
**UserDefinedFunctions** | Pointer to [**[]UserDefinedFunction**](UserDefinedFunction.md) |  | [optional] 
**CustomMetrics** | Pointer to [**[]CustomMetric**](CustomMetric.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomField**](CustomField.md) |  | [optional] 
**SelectedCheckpoints** | Pointer to [**SelectedCheckpoints**](SelectedCheckpoints.md) |  | [optional] 
**UsePrimaryCheckpointsOnly** | Pointer to **bool** |  | [optional] 
**SelfServiceTransactionScript** | Pointer to **string** |  | [optional] 
**MonitorType** | Pointer to [**MonitorType**](MonitorType.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**AlertOnLoadTimeLimit1** | Pointer to **bool** |  | [optional] 
**LoadTimeLimit1** | Pointer to **int32** |  | [optional] 
**AlertOnLoadTimeLimit2** | Pointer to **bool** |  | [optional] 
**LoadTimeLimit2** | Pointer to **int32** |  | [optional] 
**BlockGoogleAnalytics** | Pointer to **bool** |  | [optional] 
**BlockUptrendsRum** | Pointer to **bool** |  | [optional] 
**BlockUrls** | Pointer to **[]string** |  | [optional] 
**RequestHeaders** | Pointer to [**[]RequestHeader**](RequestHeader.md) |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**NameForPhoneAlerts** | Pointer to **string** |  | [optional] 
**AuthenticationType** | Pointer to [**ApiHttpAuthenticationType**](ApiHttpAuthenticationType.md) |  | [optional] 
**ThrottlingOptions** | Pointer to [**ThrottlingOptions**](ThrottlingOptions.md) |  | [optional] 
**TransactionStepDefinition** | Pointer to [**TransactionStepDefinition**](TransactionStepDefinition.md) | Only valid for Transaction monitors: the data structure that specifies the transaction steps (and sub steps) to execute. | [optional] 
**CertificateName** | Pointer to **string** |  | [optional] 
**CertificateOrganization** | Pointer to **string** |  | [optional] 
**CertificateOrganizationalUnit** | Pointer to **string** |  | [optional] 
**CertificateSerialNumber** | Pointer to **string** |  | [optional] 
**CertificateFingerprint** | Pointer to **string** |  | [optional] 
**CertificateIssuerName** | Pointer to **string** |  | [optional] 
**CertificateIssuerCompanyName** | Pointer to **string** |  | [optional] 
**CertificateIssuerOrganizationalUnit** | Pointer to **string** |  | [optional] 
**CertificateExpirationWarningDays** | Pointer to **int32** |  | [optional] 
**CheckCertificateErrors** | Pointer to **bool** |  | [optional] 
**AlertOnMaximumBytes** | Pointer to **bool** |  | [optional] 
**MaximumBytes** | Pointer to **int32** |  | [optional] 
**AlertOnMaximumSize** | Pointer to **bool** |  | [optional] 
**ElementMaximumSize** | Pointer to **int32** |  | [optional] 
**IgnoreExternalElements** | Pointer to **bool** |  | [optional] 
**AlertOnPercentageFail** | Pointer to **bool** |  | [optional] 
**FailedObjectPercentage** | Pointer to **int32** |  | [optional] 
**DomainGroupGuid** | Pointer to **string** |  | [optional] 
**DomainGroupGuidSpecified** | Pointer to **bool** |  | [optional] 
**DnsServer** | Pointer to **string** |  | [optional] 
**DnsQuery** | Pointer to [**DnsQuery**](DnsQuery.md) |  | [optional] 
**DnsExpectedResult** | Pointer to **string** |  | [optional] 
**DnsTestValue** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**IpVersion** | Pointer to [**IpVersion**](IpVersion.md) |  | [optional] 
**NativeIPv6Only** | Pointer to **bool** |  | [optional] 
**AlertOnMinimumBytes** | Pointer to **bool** |  | [optional] 
**MinimumBytes** | Pointer to **int32** |  | [optional] 
**DatabaseName** | Pointer to **string** |  | [optional] 
**NetworkAddress** | Pointer to **string** |  | [optional] 
**ImapSecureConnection** | Pointer to **bool** |  | [optional] 
**SftpAction** | Pointer to [**SftpAction**](SftpAction.md) |  | [optional] 
**SftpActionPath** | Pointer to **string** |  | [optional] 
**HttpMethod** | Pointer to [**HttpMethod**](HttpMethod.md) |  | [optional] 
**ExpectedHttpStatusCode** | Pointer to **int32** |  | [optional] 
**ExpectedHttpStatusCodeSpecified** | Pointer to **bool** |  | [optional] 
**TlsVersion** | Pointer to [**TlsVersion**](TlsVersion.md) |  | [optional] 
**RequestBody** | Pointer to **string** |  | [optional] 
**MatchPatterns** | Pointer to [**[]PatternMatch**](PatternMatch.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**BrowserType** | Pointer to [**BrowserType**](BrowserType.md) |  | [optional] 
**BrowserWindowDimensions** | Pointer to [**BrowserWindowDimensions**](BrowserWindowDimensions.md) |  | [optional] 
**UseConcurrentMonitoring** | Pointer to **bool** |  | [optional] 
**ConcurrentUnconfirmedErrorThreshold** | Pointer to **int32** |  | [optional] 
**ConcurrentConfirmedErrorThreshold** | Pointer to **int32** |  | [optional] 

## Methods

### NewMonitor

`func NewMonitor() *Monitor`

NewMonitor instantiates a new Monitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorWithDefaults

`func NewMonitorWithDefaults() *Monitor`

NewMonitorWithDefaults instantiates a new Monitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorGuid

`func (o *Monitor) GetMonitorGuid() string`

GetMonitorGuid returns the MonitorGuid field if non-nil, zero value otherwise.

### GetMonitorGuidOk

`func (o *Monitor) GetMonitorGuidOk() (*string, bool)`

GetMonitorGuidOk returns a tuple with the MonitorGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGuid

`func (o *Monitor) SetMonitorGuid(v string)`

SetMonitorGuid sets MonitorGuid field to given value.

### HasMonitorGuid

`func (o *Monitor) HasMonitorGuid() bool`

HasMonitorGuid returns a boolean if a field has been set.

### GetName

`func (o *Monitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Monitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Monitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Monitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *Monitor) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Monitor) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Monitor) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Monitor) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetHash

`func (o *Monitor) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Monitor) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Monitor) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Monitor) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetGenerateAlert

`func (o *Monitor) GetGenerateAlert() bool`

GetGenerateAlert returns the GenerateAlert field if non-nil, zero value otherwise.

### GetGenerateAlertOk

`func (o *Monitor) GetGenerateAlertOk() (*bool, bool)`

GetGenerateAlertOk returns a tuple with the GenerateAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateAlert

`func (o *Monitor) SetGenerateAlert(v bool)`

SetGenerateAlert sets GenerateAlert field to given value.

### HasGenerateAlert

`func (o *Monitor) HasGenerateAlert() bool`

HasGenerateAlert returns a boolean if a field has been set.

### GetIsLocked

`func (o *Monitor) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *Monitor) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *Monitor) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *Monitor) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetCheckInterval

`func (o *Monitor) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *Monitor) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *Monitor) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *Monitor) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetCredits

`func (o *Monitor) GetCredits() int32`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *Monitor) GetCreditsOk() (*int32, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *Monitor) SetCredits(v int32)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *Monitor) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetMonitorMode

`func (o *Monitor) GetMonitorMode() MonitorMode`

GetMonitorMode returns the MonitorMode field if non-nil, zero value otherwise.

### GetMonitorModeOk

`func (o *Monitor) GetMonitorModeOk() (*MonitorMode, bool)`

GetMonitorModeOk returns a tuple with the MonitorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorMode

`func (o *Monitor) SetMonitorMode(v MonitorMode)`

SetMonitorMode sets MonitorMode field to given value.

### HasMonitorMode

`func (o *Monitor) HasMonitorMode() bool`

HasMonitorMode returns a boolean if a field has been set.

### GetPredefinedVariables

`func (o *Monitor) GetPredefinedVariables() []PredefinedVariable`

GetPredefinedVariables returns the PredefinedVariables field if non-nil, zero value otherwise.

### GetPredefinedVariablesOk

`func (o *Monitor) GetPredefinedVariablesOk() (*[]PredefinedVariable, bool)`

GetPredefinedVariablesOk returns a tuple with the PredefinedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedVariables

`func (o *Monitor) SetPredefinedVariables(v []PredefinedVariable)`

SetPredefinedVariables sets PredefinedVariables field to given value.

### HasPredefinedVariables

`func (o *Monitor) HasPredefinedVariables() bool`

HasPredefinedVariables returns a boolean if a field has been set.

### GetMsaSteps

`func (o *Monitor) GetMsaSteps() []MsaStep`

GetMsaSteps returns the MsaSteps field if non-nil, zero value otherwise.

### GetMsaStepsOk

`func (o *Monitor) GetMsaStepsOk() (*[]MsaStep, bool)`

GetMsaStepsOk returns a tuple with the MsaSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsaSteps

`func (o *Monitor) SetMsaSteps(v []MsaStep)`

SetMsaSteps sets MsaSteps field to given value.

### HasMsaSteps

`func (o *Monitor) HasMsaSteps() bool`

HasMsaSteps returns a boolean if a field has been set.

### GetUserDefinedFunctions

`func (o *Monitor) GetUserDefinedFunctions() []UserDefinedFunction`

GetUserDefinedFunctions returns the UserDefinedFunctions field if non-nil, zero value otherwise.

### GetUserDefinedFunctionsOk

`func (o *Monitor) GetUserDefinedFunctionsOk() (*[]UserDefinedFunction, bool)`

GetUserDefinedFunctionsOk returns a tuple with the UserDefinedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFunctions

`func (o *Monitor) SetUserDefinedFunctions(v []UserDefinedFunction)`

SetUserDefinedFunctions sets UserDefinedFunctions field to given value.

### HasUserDefinedFunctions

`func (o *Monitor) HasUserDefinedFunctions() bool`

HasUserDefinedFunctions returns a boolean if a field has been set.

### GetCustomMetrics

`func (o *Monitor) GetCustomMetrics() []CustomMetric`

GetCustomMetrics returns the CustomMetrics field if non-nil, zero value otherwise.

### GetCustomMetricsOk

`func (o *Monitor) GetCustomMetricsOk() (*[]CustomMetric, bool)`

GetCustomMetricsOk returns a tuple with the CustomMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetrics

`func (o *Monitor) SetCustomMetrics(v []CustomMetric)`

SetCustomMetrics sets CustomMetrics field to given value.

### HasCustomMetrics

`func (o *Monitor) HasCustomMetrics() bool`

HasCustomMetrics returns a boolean if a field has been set.

### GetCustomFields

`func (o *Monitor) GetCustomFields() []CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Monitor) GetCustomFieldsOk() (*[]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Monitor) SetCustomFields(v []CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Monitor) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetSelectedCheckpoints

`func (o *Monitor) GetSelectedCheckpoints() SelectedCheckpoints`

GetSelectedCheckpoints returns the SelectedCheckpoints field if non-nil, zero value otherwise.

### GetSelectedCheckpointsOk

`func (o *Monitor) GetSelectedCheckpointsOk() (*SelectedCheckpoints, bool)`

GetSelectedCheckpointsOk returns a tuple with the SelectedCheckpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedCheckpoints

`func (o *Monitor) SetSelectedCheckpoints(v SelectedCheckpoints)`

SetSelectedCheckpoints sets SelectedCheckpoints field to given value.

### HasSelectedCheckpoints

`func (o *Monitor) HasSelectedCheckpoints() bool`

HasSelectedCheckpoints returns a boolean if a field has been set.

### GetUsePrimaryCheckpointsOnly

`func (o *Monitor) GetUsePrimaryCheckpointsOnly() bool`

GetUsePrimaryCheckpointsOnly returns the UsePrimaryCheckpointsOnly field if non-nil, zero value otherwise.

### GetUsePrimaryCheckpointsOnlyOk

`func (o *Monitor) GetUsePrimaryCheckpointsOnlyOk() (*bool, bool)`

GetUsePrimaryCheckpointsOnlyOk returns a tuple with the UsePrimaryCheckpointsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePrimaryCheckpointsOnly

`func (o *Monitor) SetUsePrimaryCheckpointsOnly(v bool)`

SetUsePrimaryCheckpointsOnly sets UsePrimaryCheckpointsOnly field to given value.

### HasUsePrimaryCheckpointsOnly

`func (o *Monitor) HasUsePrimaryCheckpointsOnly() bool`

HasUsePrimaryCheckpointsOnly returns a boolean if a field has been set.

### GetSelfServiceTransactionScript

`func (o *Monitor) GetSelfServiceTransactionScript() string`

GetSelfServiceTransactionScript returns the SelfServiceTransactionScript field if non-nil, zero value otherwise.

### GetSelfServiceTransactionScriptOk

`func (o *Monitor) GetSelfServiceTransactionScriptOk() (*string, bool)`

GetSelfServiceTransactionScriptOk returns a tuple with the SelfServiceTransactionScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceTransactionScript

`func (o *Monitor) SetSelfServiceTransactionScript(v string)`

SetSelfServiceTransactionScript sets SelfServiceTransactionScript field to given value.

### HasSelfServiceTransactionScript

`func (o *Monitor) HasSelfServiceTransactionScript() bool`

HasSelfServiceTransactionScript returns a boolean if a field has been set.

### GetMonitorType

`func (o *Monitor) GetMonitorType() MonitorType`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *Monitor) GetMonitorTypeOk() (*MonitorType, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *Monitor) SetMonitorType(v MonitorType)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *Monitor) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetNotes

`func (o *Monitor) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Monitor) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Monitor) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Monitor) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAlertOnLoadTimeLimit1

`func (o *Monitor) GetAlertOnLoadTimeLimit1() bool`

GetAlertOnLoadTimeLimit1 returns the AlertOnLoadTimeLimit1 field if non-nil, zero value otherwise.

### GetAlertOnLoadTimeLimit1Ok

`func (o *Monitor) GetAlertOnLoadTimeLimit1Ok() (*bool, bool)`

GetAlertOnLoadTimeLimit1Ok returns a tuple with the AlertOnLoadTimeLimit1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnLoadTimeLimit1

`func (o *Monitor) SetAlertOnLoadTimeLimit1(v bool)`

SetAlertOnLoadTimeLimit1 sets AlertOnLoadTimeLimit1 field to given value.

### HasAlertOnLoadTimeLimit1

`func (o *Monitor) HasAlertOnLoadTimeLimit1() bool`

HasAlertOnLoadTimeLimit1 returns a boolean if a field has been set.

### GetLoadTimeLimit1

`func (o *Monitor) GetLoadTimeLimit1() int32`

GetLoadTimeLimit1 returns the LoadTimeLimit1 field if non-nil, zero value otherwise.

### GetLoadTimeLimit1Ok

`func (o *Monitor) GetLoadTimeLimit1Ok() (*int32, bool)`

GetLoadTimeLimit1Ok returns a tuple with the LoadTimeLimit1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadTimeLimit1

`func (o *Monitor) SetLoadTimeLimit1(v int32)`

SetLoadTimeLimit1 sets LoadTimeLimit1 field to given value.

### HasLoadTimeLimit1

`func (o *Monitor) HasLoadTimeLimit1() bool`

HasLoadTimeLimit1 returns a boolean if a field has been set.

### GetAlertOnLoadTimeLimit2

`func (o *Monitor) GetAlertOnLoadTimeLimit2() bool`

GetAlertOnLoadTimeLimit2 returns the AlertOnLoadTimeLimit2 field if non-nil, zero value otherwise.

### GetAlertOnLoadTimeLimit2Ok

`func (o *Monitor) GetAlertOnLoadTimeLimit2Ok() (*bool, bool)`

GetAlertOnLoadTimeLimit2Ok returns a tuple with the AlertOnLoadTimeLimit2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnLoadTimeLimit2

`func (o *Monitor) SetAlertOnLoadTimeLimit2(v bool)`

SetAlertOnLoadTimeLimit2 sets AlertOnLoadTimeLimit2 field to given value.

### HasAlertOnLoadTimeLimit2

`func (o *Monitor) HasAlertOnLoadTimeLimit2() bool`

HasAlertOnLoadTimeLimit2 returns a boolean if a field has been set.

### GetLoadTimeLimit2

`func (o *Monitor) GetLoadTimeLimit2() int32`

GetLoadTimeLimit2 returns the LoadTimeLimit2 field if non-nil, zero value otherwise.

### GetLoadTimeLimit2Ok

`func (o *Monitor) GetLoadTimeLimit2Ok() (*int32, bool)`

GetLoadTimeLimit2Ok returns a tuple with the LoadTimeLimit2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadTimeLimit2

`func (o *Monitor) SetLoadTimeLimit2(v int32)`

SetLoadTimeLimit2 sets LoadTimeLimit2 field to given value.

### HasLoadTimeLimit2

`func (o *Monitor) HasLoadTimeLimit2() bool`

HasLoadTimeLimit2 returns a boolean if a field has been set.

### GetBlockGoogleAnalytics

`func (o *Monitor) GetBlockGoogleAnalytics() bool`

GetBlockGoogleAnalytics returns the BlockGoogleAnalytics field if non-nil, zero value otherwise.

### GetBlockGoogleAnalyticsOk

`func (o *Monitor) GetBlockGoogleAnalyticsOk() (*bool, bool)`

GetBlockGoogleAnalyticsOk returns a tuple with the BlockGoogleAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockGoogleAnalytics

`func (o *Monitor) SetBlockGoogleAnalytics(v bool)`

SetBlockGoogleAnalytics sets BlockGoogleAnalytics field to given value.

### HasBlockGoogleAnalytics

`func (o *Monitor) HasBlockGoogleAnalytics() bool`

HasBlockGoogleAnalytics returns a boolean if a field has been set.

### GetBlockUptrendsRum

`func (o *Monitor) GetBlockUptrendsRum() bool`

GetBlockUptrendsRum returns the BlockUptrendsRum field if non-nil, zero value otherwise.

### GetBlockUptrendsRumOk

`func (o *Monitor) GetBlockUptrendsRumOk() (*bool, bool)`

GetBlockUptrendsRumOk returns a tuple with the BlockUptrendsRum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUptrendsRum

`func (o *Monitor) SetBlockUptrendsRum(v bool)`

SetBlockUptrendsRum sets BlockUptrendsRum field to given value.

### HasBlockUptrendsRum

`func (o *Monitor) HasBlockUptrendsRum() bool`

HasBlockUptrendsRum returns a boolean if a field has been set.

### GetBlockUrls

`func (o *Monitor) GetBlockUrls() []string`

GetBlockUrls returns the BlockUrls field if non-nil, zero value otherwise.

### GetBlockUrlsOk

`func (o *Monitor) GetBlockUrlsOk() (*[]string, bool)`

GetBlockUrlsOk returns a tuple with the BlockUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUrls

`func (o *Monitor) SetBlockUrls(v []string)`

SetBlockUrls sets BlockUrls field to given value.

### HasBlockUrls

`func (o *Monitor) HasBlockUrls() bool`

HasBlockUrls returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *Monitor) GetRequestHeaders() []RequestHeader`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *Monitor) GetRequestHeadersOk() (*[]RequestHeader, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *Monitor) SetRequestHeaders(v []RequestHeader)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *Monitor) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetUserAgent

`func (o *Monitor) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *Monitor) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *Monitor) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *Monitor) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUsername

`func (o *Monitor) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Monitor) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Monitor) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Monitor) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *Monitor) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Monitor) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Monitor) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Monitor) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetNameForPhoneAlerts

`func (o *Monitor) GetNameForPhoneAlerts() string`

GetNameForPhoneAlerts returns the NameForPhoneAlerts field if non-nil, zero value otherwise.

### GetNameForPhoneAlertsOk

`func (o *Monitor) GetNameForPhoneAlertsOk() (*string, bool)`

GetNameForPhoneAlertsOk returns a tuple with the NameForPhoneAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameForPhoneAlerts

`func (o *Monitor) SetNameForPhoneAlerts(v string)`

SetNameForPhoneAlerts sets NameForPhoneAlerts field to given value.

### HasNameForPhoneAlerts

`func (o *Monitor) HasNameForPhoneAlerts() bool`

HasNameForPhoneAlerts returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *Monitor) GetAuthenticationType() ApiHttpAuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *Monitor) GetAuthenticationTypeOk() (*ApiHttpAuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *Monitor) SetAuthenticationType(v ApiHttpAuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *Monitor) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetThrottlingOptions

`func (o *Monitor) GetThrottlingOptions() ThrottlingOptions`

GetThrottlingOptions returns the ThrottlingOptions field if non-nil, zero value otherwise.

### GetThrottlingOptionsOk

`func (o *Monitor) GetThrottlingOptionsOk() (*ThrottlingOptions, bool)`

GetThrottlingOptionsOk returns a tuple with the ThrottlingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingOptions

`func (o *Monitor) SetThrottlingOptions(v ThrottlingOptions)`

SetThrottlingOptions sets ThrottlingOptions field to given value.

### HasThrottlingOptions

`func (o *Monitor) HasThrottlingOptions() bool`

HasThrottlingOptions returns a boolean if a field has been set.

### GetTransactionStepDefinition

`func (o *Monitor) GetTransactionStepDefinition() TransactionStepDefinition`

GetTransactionStepDefinition returns the TransactionStepDefinition field if non-nil, zero value otherwise.

### GetTransactionStepDefinitionOk

`func (o *Monitor) GetTransactionStepDefinitionOk() (*TransactionStepDefinition, bool)`

GetTransactionStepDefinitionOk returns a tuple with the TransactionStepDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStepDefinition

`func (o *Monitor) SetTransactionStepDefinition(v TransactionStepDefinition)`

SetTransactionStepDefinition sets TransactionStepDefinition field to given value.

### HasTransactionStepDefinition

`func (o *Monitor) HasTransactionStepDefinition() bool`

HasTransactionStepDefinition returns a boolean if a field has been set.

### GetCertificateName

`func (o *Monitor) GetCertificateName() string`

GetCertificateName returns the CertificateName field if non-nil, zero value otherwise.

### GetCertificateNameOk

`func (o *Monitor) GetCertificateNameOk() (*string, bool)`

GetCertificateNameOk returns a tuple with the CertificateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateName

`func (o *Monitor) SetCertificateName(v string)`

SetCertificateName sets CertificateName field to given value.

### HasCertificateName

`func (o *Monitor) HasCertificateName() bool`

HasCertificateName returns a boolean if a field has been set.

### GetCertificateOrganization

`func (o *Monitor) GetCertificateOrganization() string`

GetCertificateOrganization returns the CertificateOrganization field if non-nil, zero value otherwise.

### GetCertificateOrganizationOk

`func (o *Monitor) GetCertificateOrganizationOk() (*string, bool)`

GetCertificateOrganizationOk returns a tuple with the CertificateOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOrganization

`func (o *Monitor) SetCertificateOrganization(v string)`

SetCertificateOrganization sets CertificateOrganization field to given value.

### HasCertificateOrganization

`func (o *Monitor) HasCertificateOrganization() bool`

HasCertificateOrganization returns a boolean if a field has been set.

### GetCertificateOrganizationalUnit

`func (o *Monitor) GetCertificateOrganizationalUnit() string`

GetCertificateOrganizationalUnit returns the CertificateOrganizationalUnit field if non-nil, zero value otherwise.

### GetCertificateOrganizationalUnitOk

`func (o *Monitor) GetCertificateOrganizationalUnitOk() (*string, bool)`

GetCertificateOrganizationalUnitOk returns a tuple with the CertificateOrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOrganizationalUnit

`func (o *Monitor) SetCertificateOrganizationalUnit(v string)`

SetCertificateOrganizationalUnit sets CertificateOrganizationalUnit field to given value.

### HasCertificateOrganizationalUnit

`func (o *Monitor) HasCertificateOrganizationalUnit() bool`

HasCertificateOrganizationalUnit returns a boolean if a field has been set.

### GetCertificateSerialNumber

`func (o *Monitor) GetCertificateSerialNumber() string`

GetCertificateSerialNumber returns the CertificateSerialNumber field if non-nil, zero value otherwise.

### GetCertificateSerialNumberOk

`func (o *Monitor) GetCertificateSerialNumberOk() (*string, bool)`

GetCertificateSerialNumberOk returns a tuple with the CertificateSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSerialNumber

`func (o *Monitor) SetCertificateSerialNumber(v string)`

SetCertificateSerialNumber sets CertificateSerialNumber field to given value.

### HasCertificateSerialNumber

`func (o *Monitor) HasCertificateSerialNumber() bool`

HasCertificateSerialNumber returns a boolean if a field has been set.

### GetCertificateFingerprint

`func (o *Monitor) GetCertificateFingerprint() string`

GetCertificateFingerprint returns the CertificateFingerprint field if non-nil, zero value otherwise.

### GetCertificateFingerprintOk

`func (o *Monitor) GetCertificateFingerprintOk() (*string, bool)`

GetCertificateFingerprintOk returns a tuple with the CertificateFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFingerprint

`func (o *Monitor) SetCertificateFingerprint(v string)`

SetCertificateFingerprint sets CertificateFingerprint field to given value.

### HasCertificateFingerprint

`func (o *Monitor) HasCertificateFingerprint() bool`

HasCertificateFingerprint returns a boolean if a field has been set.

### GetCertificateIssuerName

`func (o *Monitor) GetCertificateIssuerName() string`

GetCertificateIssuerName returns the CertificateIssuerName field if non-nil, zero value otherwise.

### GetCertificateIssuerNameOk

`func (o *Monitor) GetCertificateIssuerNameOk() (*string, bool)`

GetCertificateIssuerNameOk returns a tuple with the CertificateIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIssuerName

`func (o *Monitor) SetCertificateIssuerName(v string)`

SetCertificateIssuerName sets CertificateIssuerName field to given value.

### HasCertificateIssuerName

`func (o *Monitor) HasCertificateIssuerName() bool`

HasCertificateIssuerName returns a boolean if a field has been set.

### GetCertificateIssuerCompanyName

`func (o *Monitor) GetCertificateIssuerCompanyName() string`

GetCertificateIssuerCompanyName returns the CertificateIssuerCompanyName field if non-nil, zero value otherwise.

### GetCertificateIssuerCompanyNameOk

`func (o *Monitor) GetCertificateIssuerCompanyNameOk() (*string, bool)`

GetCertificateIssuerCompanyNameOk returns a tuple with the CertificateIssuerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIssuerCompanyName

`func (o *Monitor) SetCertificateIssuerCompanyName(v string)`

SetCertificateIssuerCompanyName sets CertificateIssuerCompanyName field to given value.

### HasCertificateIssuerCompanyName

`func (o *Monitor) HasCertificateIssuerCompanyName() bool`

HasCertificateIssuerCompanyName returns a boolean if a field has been set.

### GetCertificateIssuerOrganizationalUnit

`func (o *Monitor) GetCertificateIssuerOrganizationalUnit() string`

GetCertificateIssuerOrganizationalUnit returns the CertificateIssuerOrganizationalUnit field if non-nil, zero value otherwise.

### GetCertificateIssuerOrganizationalUnitOk

`func (o *Monitor) GetCertificateIssuerOrganizationalUnitOk() (*string, bool)`

GetCertificateIssuerOrganizationalUnitOk returns a tuple with the CertificateIssuerOrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIssuerOrganizationalUnit

`func (o *Monitor) SetCertificateIssuerOrganizationalUnit(v string)`

SetCertificateIssuerOrganizationalUnit sets CertificateIssuerOrganizationalUnit field to given value.

### HasCertificateIssuerOrganizationalUnit

`func (o *Monitor) HasCertificateIssuerOrganizationalUnit() bool`

HasCertificateIssuerOrganizationalUnit returns a boolean if a field has been set.

### GetCertificateExpirationWarningDays

`func (o *Monitor) GetCertificateExpirationWarningDays() int32`

GetCertificateExpirationWarningDays returns the CertificateExpirationWarningDays field if non-nil, zero value otherwise.

### GetCertificateExpirationWarningDaysOk

`func (o *Monitor) GetCertificateExpirationWarningDaysOk() (*int32, bool)`

GetCertificateExpirationWarningDaysOk returns a tuple with the CertificateExpirationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExpirationWarningDays

`func (o *Monitor) SetCertificateExpirationWarningDays(v int32)`

SetCertificateExpirationWarningDays sets CertificateExpirationWarningDays field to given value.

### HasCertificateExpirationWarningDays

`func (o *Monitor) HasCertificateExpirationWarningDays() bool`

HasCertificateExpirationWarningDays returns a boolean if a field has been set.

### GetCheckCertificateErrors

`func (o *Monitor) GetCheckCertificateErrors() bool`

GetCheckCertificateErrors returns the CheckCertificateErrors field if non-nil, zero value otherwise.

### GetCheckCertificateErrorsOk

`func (o *Monitor) GetCheckCertificateErrorsOk() (*bool, bool)`

GetCheckCertificateErrorsOk returns a tuple with the CheckCertificateErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCertificateErrors

`func (o *Monitor) SetCheckCertificateErrors(v bool)`

SetCheckCertificateErrors sets CheckCertificateErrors field to given value.

### HasCheckCertificateErrors

`func (o *Monitor) HasCheckCertificateErrors() bool`

HasCheckCertificateErrors returns a boolean if a field has been set.

### GetAlertOnMaximumBytes

`func (o *Monitor) GetAlertOnMaximumBytes() bool`

GetAlertOnMaximumBytes returns the AlertOnMaximumBytes field if non-nil, zero value otherwise.

### GetAlertOnMaximumBytesOk

`func (o *Monitor) GetAlertOnMaximumBytesOk() (*bool, bool)`

GetAlertOnMaximumBytesOk returns a tuple with the AlertOnMaximumBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnMaximumBytes

`func (o *Monitor) SetAlertOnMaximumBytes(v bool)`

SetAlertOnMaximumBytes sets AlertOnMaximumBytes field to given value.

### HasAlertOnMaximumBytes

`func (o *Monitor) HasAlertOnMaximumBytes() bool`

HasAlertOnMaximumBytes returns a boolean if a field has been set.

### GetMaximumBytes

`func (o *Monitor) GetMaximumBytes() int32`

GetMaximumBytes returns the MaximumBytes field if non-nil, zero value otherwise.

### GetMaximumBytesOk

`func (o *Monitor) GetMaximumBytesOk() (*int32, bool)`

GetMaximumBytesOk returns a tuple with the MaximumBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBytes

`func (o *Monitor) SetMaximumBytes(v int32)`

SetMaximumBytes sets MaximumBytes field to given value.

### HasMaximumBytes

`func (o *Monitor) HasMaximumBytes() bool`

HasMaximumBytes returns a boolean if a field has been set.

### GetAlertOnMaximumSize

`func (o *Monitor) GetAlertOnMaximumSize() bool`

GetAlertOnMaximumSize returns the AlertOnMaximumSize field if non-nil, zero value otherwise.

### GetAlertOnMaximumSizeOk

`func (o *Monitor) GetAlertOnMaximumSizeOk() (*bool, bool)`

GetAlertOnMaximumSizeOk returns a tuple with the AlertOnMaximumSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnMaximumSize

`func (o *Monitor) SetAlertOnMaximumSize(v bool)`

SetAlertOnMaximumSize sets AlertOnMaximumSize field to given value.

### HasAlertOnMaximumSize

`func (o *Monitor) HasAlertOnMaximumSize() bool`

HasAlertOnMaximumSize returns a boolean if a field has been set.

### GetElementMaximumSize

`func (o *Monitor) GetElementMaximumSize() int32`

GetElementMaximumSize returns the ElementMaximumSize field if non-nil, zero value otherwise.

### GetElementMaximumSizeOk

`func (o *Monitor) GetElementMaximumSizeOk() (*int32, bool)`

GetElementMaximumSizeOk returns a tuple with the ElementMaximumSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementMaximumSize

`func (o *Monitor) SetElementMaximumSize(v int32)`

SetElementMaximumSize sets ElementMaximumSize field to given value.

### HasElementMaximumSize

`func (o *Monitor) HasElementMaximumSize() bool`

HasElementMaximumSize returns a boolean if a field has been set.

### GetIgnoreExternalElements

`func (o *Monitor) GetIgnoreExternalElements() bool`

GetIgnoreExternalElements returns the IgnoreExternalElements field if non-nil, zero value otherwise.

### GetIgnoreExternalElementsOk

`func (o *Monitor) GetIgnoreExternalElementsOk() (*bool, bool)`

GetIgnoreExternalElementsOk returns a tuple with the IgnoreExternalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreExternalElements

`func (o *Monitor) SetIgnoreExternalElements(v bool)`

SetIgnoreExternalElements sets IgnoreExternalElements field to given value.

### HasIgnoreExternalElements

`func (o *Monitor) HasIgnoreExternalElements() bool`

HasIgnoreExternalElements returns a boolean if a field has been set.

### GetAlertOnPercentageFail

`func (o *Monitor) GetAlertOnPercentageFail() bool`

GetAlertOnPercentageFail returns the AlertOnPercentageFail field if non-nil, zero value otherwise.

### GetAlertOnPercentageFailOk

`func (o *Monitor) GetAlertOnPercentageFailOk() (*bool, bool)`

GetAlertOnPercentageFailOk returns a tuple with the AlertOnPercentageFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnPercentageFail

`func (o *Monitor) SetAlertOnPercentageFail(v bool)`

SetAlertOnPercentageFail sets AlertOnPercentageFail field to given value.

### HasAlertOnPercentageFail

`func (o *Monitor) HasAlertOnPercentageFail() bool`

HasAlertOnPercentageFail returns a boolean if a field has been set.

### GetFailedObjectPercentage

`func (o *Monitor) GetFailedObjectPercentage() int32`

GetFailedObjectPercentage returns the FailedObjectPercentage field if non-nil, zero value otherwise.

### GetFailedObjectPercentageOk

`func (o *Monitor) GetFailedObjectPercentageOk() (*int32, bool)`

GetFailedObjectPercentageOk returns a tuple with the FailedObjectPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedObjectPercentage

`func (o *Monitor) SetFailedObjectPercentage(v int32)`

SetFailedObjectPercentage sets FailedObjectPercentage field to given value.

### HasFailedObjectPercentage

`func (o *Monitor) HasFailedObjectPercentage() bool`

HasFailedObjectPercentage returns a boolean if a field has been set.

### GetDomainGroupGuid

`func (o *Monitor) GetDomainGroupGuid() string`

GetDomainGroupGuid returns the DomainGroupGuid field if non-nil, zero value otherwise.

### GetDomainGroupGuidOk

`func (o *Monitor) GetDomainGroupGuidOk() (*string, bool)`

GetDomainGroupGuidOk returns a tuple with the DomainGroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupGuid

`func (o *Monitor) SetDomainGroupGuid(v string)`

SetDomainGroupGuid sets DomainGroupGuid field to given value.

### HasDomainGroupGuid

`func (o *Monitor) HasDomainGroupGuid() bool`

HasDomainGroupGuid returns a boolean if a field has been set.

### GetDomainGroupGuidSpecified

`func (o *Monitor) GetDomainGroupGuidSpecified() bool`

GetDomainGroupGuidSpecified returns the DomainGroupGuidSpecified field if non-nil, zero value otherwise.

### GetDomainGroupGuidSpecifiedOk

`func (o *Monitor) GetDomainGroupGuidSpecifiedOk() (*bool, bool)`

GetDomainGroupGuidSpecifiedOk returns a tuple with the DomainGroupGuidSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupGuidSpecified

`func (o *Monitor) SetDomainGroupGuidSpecified(v bool)`

SetDomainGroupGuidSpecified sets DomainGroupGuidSpecified field to given value.

### HasDomainGroupGuidSpecified

`func (o *Monitor) HasDomainGroupGuidSpecified() bool`

HasDomainGroupGuidSpecified returns a boolean if a field has been set.

### GetDnsServer

`func (o *Monitor) GetDnsServer() string`

GetDnsServer returns the DnsServer field if non-nil, zero value otherwise.

### GetDnsServerOk

`func (o *Monitor) GetDnsServerOk() (*string, bool)`

GetDnsServerOk returns a tuple with the DnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServer

`func (o *Monitor) SetDnsServer(v string)`

SetDnsServer sets DnsServer field to given value.

### HasDnsServer

`func (o *Monitor) HasDnsServer() bool`

HasDnsServer returns a boolean if a field has been set.

### GetDnsQuery

`func (o *Monitor) GetDnsQuery() DnsQuery`

GetDnsQuery returns the DnsQuery field if non-nil, zero value otherwise.

### GetDnsQueryOk

`func (o *Monitor) GetDnsQueryOk() (*DnsQuery, bool)`

GetDnsQueryOk returns a tuple with the DnsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQuery

`func (o *Monitor) SetDnsQuery(v DnsQuery)`

SetDnsQuery sets DnsQuery field to given value.

### HasDnsQuery

`func (o *Monitor) HasDnsQuery() bool`

HasDnsQuery returns a boolean if a field has been set.

### GetDnsExpectedResult

`func (o *Monitor) GetDnsExpectedResult() string`

GetDnsExpectedResult returns the DnsExpectedResult field if non-nil, zero value otherwise.

### GetDnsExpectedResultOk

`func (o *Monitor) GetDnsExpectedResultOk() (*string, bool)`

GetDnsExpectedResultOk returns a tuple with the DnsExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsExpectedResult

`func (o *Monitor) SetDnsExpectedResult(v string)`

SetDnsExpectedResult sets DnsExpectedResult field to given value.

### HasDnsExpectedResult

`func (o *Monitor) HasDnsExpectedResult() bool`

HasDnsExpectedResult returns a boolean if a field has been set.

### GetDnsTestValue

`func (o *Monitor) GetDnsTestValue() string`

GetDnsTestValue returns the DnsTestValue field if non-nil, zero value otherwise.

### GetDnsTestValueOk

`func (o *Monitor) GetDnsTestValueOk() (*string, bool)`

GetDnsTestValueOk returns a tuple with the DnsTestValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTestValue

`func (o *Monitor) SetDnsTestValue(v string)`

SetDnsTestValue sets DnsTestValue field to given value.

### HasDnsTestValue

`func (o *Monitor) HasDnsTestValue() bool`

HasDnsTestValue returns a boolean if a field has been set.

### GetPort

`func (o *Monitor) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Monitor) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Monitor) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Monitor) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetIpVersion

`func (o *Monitor) GetIpVersion() IpVersion`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *Monitor) GetIpVersionOk() (*IpVersion, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *Monitor) SetIpVersion(v IpVersion)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *Monitor) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetNativeIPv6Only

`func (o *Monitor) GetNativeIPv6Only() bool`

GetNativeIPv6Only returns the NativeIPv6Only field if non-nil, zero value otherwise.

### GetNativeIPv6OnlyOk

`func (o *Monitor) GetNativeIPv6OnlyOk() (*bool, bool)`

GetNativeIPv6OnlyOk returns a tuple with the NativeIPv6Only field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIPv6Only

`func (o *Monitor) SetNativeIPv6Only(v bool)`

SetNativeIPv6Only sets NativeIPv6Only field to given value.

### HasNativeIPv6Only

`func (o *Monitor) HasNativeIPv6Only() bool`

HasNativeIPv6Only returns a boolean if a field has been set.

### GetAlertOnMinimumBytes

`func (o *Monitor) GetAlertOnMinimumBytes() bool`

GetAlertOnMinimumBytes returns the AlertOnMinimumBytes field if non-nil, zero value otherwise.

### GetAlertOnMinimumBytesOk

`func (o *Monitor) GetAlertOnMinimumBytesOk() (*bool, bool)`

GetAlertOnMinimumBytesOk returns a tuple with the AlertOnMinimumBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnMinimumBytes

`func (o *Monitor) SetAlertOnMinimumBytes(v bool)`

SetAlertOnMinimumBytes sets AlertOnMinimumBytes field to given value.

### HasAlertOnMinimumBytes

`func (o *Monitor) HasAlertOnMinimumBytes() bool`

HasAlertOnMinimumBytes returns a boolean if a field has been set.

### GetMinimumBytes

`func (o *Monitor) GetMinimumBytes() int32`

GetMinimumBytes returns the MinimumBytes field if non-nil, zero value otherwise.

### GetMinimumBytesOk

`func (o *Monitor) GetMinimumBytesOk() (*int32, bool)`

GetMinimumBytesOk returns a tuple with the MinimumBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBytes

`func (o *Monitor) SetMinimumBytes(v int32)`

SetMinimumBytes sets MinimumBytes field to given value.

### HasMinimumBytes

`func (o *Monitor) HasMinimumBytes() bool`

HasMinimumBytes returns a boolean if a field has been set.

### GetDatabaseName

`func (o *Monitor) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *Monitor) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *Monitor) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *Monitor) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetNetworkAddress

`func (o *Monitor) GetNetworkAddress() string`

GetNetworkAddress returns the NetworkAddress field if non-nil, zero value otherwise.

### GetNetworkAddressOk

`func (o *Monitor) GetNetworkAddressOk() (*string, bool)`

GetNetworkAddressOk returns a tuple with the NetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddress

`func (o *Monitor) SetNetworkAddress(v string)`

SetNetworkAddress sets NetworkAddress field to given value.

### HasNetworkAddress

`func (o *Monitor) HasNetworkAddress() bool`

HasNetworkAddress returns a boolean if a field has been set.

### GetImapSecureConnection

`func (o *Monitor) GetImapSecureConnection() bool`

GetImapSecureConnection returns the ImapSecureConnection field if non-nil, zero value otherwise.

### GetImapSecureConnectionOk

`func (o *Monitor) GetImapSecureConnectionOk() (*bool, bool)`

GetImapSecureConnectionOk returns a tuple with the ImapSecureConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapSecureConnection

`func (o *Monitor) SetImapSecureConnection(v bool)`

SetImapSecureConnection sets ImapSecureConnection field to given value.

### HasImapSecureConnection

`func (o *Monitor) HasImapSecureConnection() bool`

HasImapSecureConnection returns a boolean if a field has been set.

### GetSftpAction

`func (o *Monitor) GetSftpAction() SftpAction`

GetSftpAction returns the SftpAction field if non-nil, zero value otherwise.

### GetSftpActionOk

`func (o *Monitor) GetSftpActionOk() (*SftpAction, bool)`

GetSftpActionOk returns a tuple with the SftpAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSftpAction

`func (o *Monitor) SetSftpAction(v SftpAction)`

SetSftpAction sets SftpAction field to given value.

### HasSftpAction

`func (o *Monitor) HasSftpAction() bool`

HasSftpAction returns a boolean if a field has been set.

### GetSftpActionPath

`func (o *Monitor) GetSftpActionPath() string`

GetSftpActionPath returns the SftpActionPath field if non-nil, zero value otherwise.

### GetSftpActionPathOk

`func (o *Monitor) GetSftpActionPathOk() (*string, bool)`

GetSftpActionPathOk returns a tuple with the SftpActionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSftpActionPath

`func (o *Monitor) SetSftpActionPath(v string)`

SetSftpActionPath sets SftpActionPath field to given value.

### HasSftpActionPath

`func (o *Monitor) HasSftpActionPath() bool`

HasSftpActionPath returns a boolean if a field has been set.

### GetHttpMethod

`func (o *Monitor) GetHttpMethod() HttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *Monitor) GetHttpMethodOk() (*HttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *Monitor) SetHttpMethod(v HttpMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *Monitor) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetExpectedHttpStatusCode

`func (o *Monitor) GetExpectedHttpStatusCode() int32`

GetExpectedHttpStatusCode returns the ExpectedHttpStatusCode field if non-nil, zero value otherwise.

### GetExpectedHttpStatusCodeOk

`func (o *Monitor) GetExpectedHttpStatusCodeOk() (*int32, bool)`

GetExpectedHttpStatusCodeOk returns a tuple with the ExpectedHttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHttpStatusCode

`func (o *Monitor) SetExpectedHttpStatusCode(v int32)`

SetExpectedHttpStatusCode sets ExpectedHttpStatusCode field to given value.

### HasExpectedHttpStatusCode

`func (o *Monitor) HasExpectedHttpStatusCode() bool`

HasExpectedHttpStatusCode returns a boolean if a field has been set.

### GetExpectedHttpStatusCodeSpecified

`func (o *Monitor) GetExpectedHttpStatusCodeSpecified() bool`

GetExpectedHttpStatusCodeSpecified returns the ExpectedHttpStatusCodeSpecified field if non-nil, zero value otherwise.

### GetExpectedHttpStatusCodeSpecifiedOk

`func (o *Monitor) GetExpectedHttpStatusCodeSpecifiedOk() (*bool, bool)`

GetExpectedHttpStatusCodeSpecifiedOk returns a tuple with the ExpectedHttpStatusCodeSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHttpStatusCodeSpecified

`func (o *Monitor) SetExpectedHttpStatusCodeSpecified(v bool)`

SetExpectedHttpStatusCodeSpecified sets ExpectedHttpStatusCodeSpecified field to given value.

### HasExpectedHttpStatusCodeSpecified

`func (o *Monitor) HasExpectedHttpStatusCodeSpecified() bool`

HasExpectedHttpStatusCodeSpecified returns a boolean if a field has been set.

### GetTlsVersion

`func (o *Monitor) GetTlsVersion() TlsVersion`

GetTlsVersion returns the TlsVersion field if non-nil, zero value otherwise.

### GetTlsVersionOk

`func (o *Monitor) GetTlsVersionOk() (*TlsVersion, bool)`

GetTlsVersionOk returns a tuple with the TlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersion

`func (o *Monitor) SetTlsVersion(v TlsVersion)`

SetTlsVersion sets TlsVersion field to given value.

### HasTlsVersion

`func (o *Monitor) HasTlsVersion() bool`

HasTlsVersion returns a boolean if a field has been set.

### GetRequestBody

`func (o *Monitor) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *Monitor) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *Monitor) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *Monitor) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetMatchPatterns

`func (o *Monitor) GetMatchPatterns() []PatternMatch`

GetMatchPatterns returns the MatchPatterns field if non-nil, zero value otherwise.

### GetMatchPatternsOk

`func (o *Monitor) GetMatchPatternsOk() (*[]PatternMatch, bool)`

GetMatchPatternsOk returns a tuple with the MatchPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPatterns

`func (o *Monitor) SetMatchPatterns(v []PatternMatch)`

SetMatchPatterns sets MatchPatterns field to given value.

### HasMatchPatterns

`func (o *Monitor) HasMatchPatterns() bool`

HasMatchPatterns returns a boolean if a field has been set.

### GetUrl

`func (o *Monitor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Monitor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Monitor) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Monitor) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBrowserType

`func (o *Monitor) GetBrowserType() BrowserType`

GetBrowserType returns the BrowserType field if non-nil, zero value otherwise.

### GetBrowserTypeOk

`func (o *Monitor) GetBrowserTypeOk() (*BrowserType, bool)`

GetBrowserTypeOk returns a tuple with the BrowserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserType

`func (o *Monitor) SetBrowserType(v BrowserType)`

SetBrowserType sets BrowserType field to given value.

### HasBrowserType

`func (o *Monitor) HasBrowserType() bool`

HasBrowserType returns a boolean if a field has been set.

### GetBrowserWindowDimensions

`func (o *Monitor) GetBrowserWindowDimensions() BrowserWindowDimensions`

GetBrowserWindowDimensions returns the BrowserWindowDimensions field if non-nil, zero value otherwise.

### GetBrowserWindowDimensionsOk

`func (o *Monitor) GetBrowserWindowDimensionsOk() (*BrowserWindowDimensions, bool)`

GetBrowserWindowDimensionsOk returns a tuple with the BrowserWindowDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserWindowDimensions

`func (o *Monitor) SetBrowserWindowDimensions(v BrowserWindowDimensions)`

SetBrowserWindowDimensions sets BrowserWindowDimensions field to given value.

### HasBrowserWindowDimensions

`func (o *Monitor) HasBrowserWindowDimensions() bool`

HasBrowserWindowDimensions returns a boolean if a field has been set.

### GetUseConcurrentMonitoring

`func (o *Monitor) GetUseConcurrentMonitoring() bool`

GetUseConcurrentMonitoring returns the UseConcurrentMonitoring field if non-nil, zero value otherwise.

### GetUseConcurrentMonitoringOk

`func (o *Monitor) GetUseConcurrentMonitoringOk() (*bool, bool)`

GetUseConcurrentMonitoringOk returns a tuple with the UseConcurrentMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseConcurrentMonitoring

`func (o *Monitor) SetUseConcurrentMonitoring(v bool)`

SetUseConcurrentMonitoring sets UseConcurrentMonitoring field to given value.

### HasUseConcurrentMonitoring

`func (o *Monitor) HasUseConcurrentMonitoring() bool`

HasUseConcurrentMonitoring returns a boolean if a field has been set.

### GetConcurrentUnconfirmedErrorThreshold

`func (o *Monitor) GetConcurrentUnconfirmedErrorThreshold() int32`

GetConcurrentUnconfirmedErrorThreshold returns the ConcurrentUnconfirmedErrorThreshold field if non-nil, zero value otherwise.

### GetConcurrentUnconfirmedErrorThresholdOk

`func (o *Monitor) GetConcurrentUnconfirmedErrorThresholdOk() (*int32, bool)`

GetConcurrentUnconfirmedErrorThresholdOk returns a tuple with the ConcurrentUnconfirmedErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentUnconfirmedErrorThreshold

`func (o *Monitor) SetConcurrentUnconfirmedErrorThreshold(v int32)`

SetConcurrentUnconfirmedErrorThreshold sets ConcurrentUnconfirmedErrorThreshold field to given value.

### HasConcurrentUnconfirmedErrorThreshold

`func (o *Monitor) HasConcurrentUnconfirmedErrorThreshold() bool`

HasConcurrentUnconfirmedErrorThreshold returns a boolean if a field has been set.

### GetConcurrentConfirmedErrorThreshold

`func (o *Monitor) GetConcurrentConfirmedErrorThreshold() int32`

GetConcurrentConfirmedErrorThreshold returns the ConcurrentConfirmedErrorThreshold field if non-nil, zero value otherwise.

### GetConcurrentConfirmedErrorThresholdOk

`func (o *Monitor) GetConcurrentConfirmedErrorThresholdOk() (*int32, bool)`

GetConcurrentConfirmedErrorThresholdOk returns a tuple with the ConcurrentConfirmedErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentConfirmedErrorThreshold

`func (o *Monitor) SetConcurrentConfirmedErrorThreshold(v int32)`

SetConcurrentConfirmedErrorThreshold sets ConcurrentConfirmedErrorThreshold field to given value.

### HasConcurrentConfirmedErrorThreshold

`func (o *Monitor) HasConcurrentConfirmedErrorThreshold() bool`

HasConcurrentConfirmedErrorThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


