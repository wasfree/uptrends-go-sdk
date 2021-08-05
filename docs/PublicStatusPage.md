# PublicStatusPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicDashboardGuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsPublished** | Pointer to **bool** |  | [optional] 
**PresetPeriodType** | Pointer to [**PresetPeriodType**](PresetPeriodType.md) |  | [optional] 
**CustomizationInfo** | Pointer to [**CustomizationInfo**](CustomizationInfo.md) |  | [optional] 
**SlaGuid** | Pointer to **string** |  | [optional] 
**SlaGuidSpecified** | Pointer to **bool** |  | [optional] 
**MonitorGuids** | Pointer to **[]string** |  | [optional] 
**MonitorGroupGuids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPublicStatusPage

`func NewPublicStatusPage() *PublicStatusPage`

NewPublicStatusPage instantiates a new PublicStatusPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicStatusPageWithDefaults

`func NewPublicStatusPageWithDefaults() *PublicStatusPage`

NewPublicStatusPageWithDefaults instantiates a new PublicStatusPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicDashboardGuid

`func (o *PublicStatusPage) GetPublicDashboardGuid() string`

GetPublicDashboardGuid returns the PublicDashboardGuid field if non-nil, zero value otherwise.

### GetPublicDashboardGuidOk

`func (o *PublicStatusPage) GetPublicDashboardGuidOk() (*string, bool)`

GetPublicDashboardGuidOk returns a tuple with the PublicDashboardGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDashboardGuid

`func (o *PublicStatusPage) SetPublicDashboardGuid(v string)`

SetPublicDashboardGuid sets PublicDashboardGuid field to given value.

### HasPublicDashboardGuid

`func (o *PublicStatusPage) HasPublicDashboardGuid() bool`

HasPublicDashboardGuid returns a boolean if a field has been set.

### GetName

`func (o *PublicStatusPage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicStatusPage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicStatusPage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicStatusPage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsPublished

`func (o *PublicStatusPage) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *PublicStatusPage) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *PublicStatusPage) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *PublicStatusPage) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetPresetPeriodType

`func (o *PublicStatusPage) GetPresetPeriodType() PresetPeriodType`

GetPresetPeriodType returns the PresetPeriodType field if non-nil, zero value otherwise.

### GetPresetPeriodTypeOk

`func (o *PublicStatusPage) GetPresetPeriodTypeOk() (*PresetPeriodType, bool)`

GetPresetPeriodTypeOk returns a tuple with the PresetPeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetPeriodType

`func (o *PublicStatusPage) SetPresetPeriodType(v PresetPeriodType)`

SetPresetPeriodType sets PresetPeriodType field to given value.

### HasPresetPeriodType

`func (o *PublicStatusPage) HasPresetPeriodType() bool`

HasPresetPeriodType returns a boolean if a field has been set.

### GetCustomizationInfo

`func (o *PublicStatusPage) GetCustomizationInfo() CustomizationInfo`

GetCustomizationInfo returns the CustomizationInfo field if non-nil, zero value otherwise.

### GetCustomizationInfoOk

`func (o *PublicStatusPage) GetCustomizationInfoOk() (*CustomizationInfo, bool)`

GetCustomizationInfoOk returns a tuple with the CustomizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationInfo

`func (o *PublicStatusPage) SetCustomizationInfo(v CustomizationInfo)`

SetCustomizationInfo sets CustomizationInfo field to given value.

### HasCustomizationInfo

`func (o *PublicStatusPage) HasCustomizationInfo() bool`

HasCustomizationInfo returns a boolean if a field has been set.

### GetSlaGuid

`func (o *PublicStatusPage) GetSlaGuid() string`

GetSlaGuid returns the SlaGuid field if non-nil, zero value otherwise.

### GetSlaGuidOk

`func (o *PublicStatusPage) GetSlaGuidOk() (*string, bool)`

GetSlaGuidOk returns a tuple with the SlaGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaGuid

`func (o *PublicStatusPage) SetSlaGuid(v string)`

SetSlaGuid sets SlaGuid field to given value.

### HasSlaGuid

`func (o *PublicStatusPage) HasSlaGuid() bool`

HasSlaGuid returns a boolean if a field has been set.

### GetSlaGuidSpecified

`func (o *PublicStatusPage) GetSlaGuidSpecified() bool`

GetSlaGuidSpecified returns the SlaGuidSpecified field if non-nil, zero value otherwise.

### GetSlaGuidSpecifiedOk

`func (o *PublicStatusPage) GetSlaGuidSpecifiedOk() (*bool, bool)`

GetSlaGuidSpecifiedOk returns a tuple with the SlaGuidSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaGuidSpecified

`func (o *PublicStatusPage) SetSlaGuidSpecified(v bool)`

SetSlaGuidSpecified sets SlaGuidSpecified field to given value.

### HasSlaGuidSpecified

`func (o *PublicStatusPage) HasSlaGuidSpecified() bool`

HasSlaGuidSpecified returns a boolean if a field has been set.

### GetMonitorGuids

`func (o *PublicStatusPage) GetMonitorGuids() []string`

GetMonitorGuids returns the MonitorGuids field if non-nil, zero value otherwise.

### GetMonitorGuidsOk

`func (o *PublicStatusPage) GetMonitorGuidsOk() (*[]string, bool)`

GetMonitorGuidsOk returns a tuple with the MonitorGuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGuids

`func (o *PublicStatusPage) SetMonitorGuids(v []string)`

SetMonitorGuids sets MonitorGuids field to given value.

### HasMonitorGuids

`func (o *PublicStatusPage) HasMonitorGuids() bool`

HasMonitorGuids returns a boolean if a field has been set.

### GetMonitorGroupGuids

`func (o *PublicStatusPage) GetMonitorGroupGuids() []string`

GetMonitorGroupGuids returns the MonitorGroupGuids field if non-nil, zero value otherwise.

### GetMonitorGroupGuidsOk

`func (o *PublicStatusPage) GetMonitorGroupGuidsOk() (*[]string, bool)`

GetMonitorGroupGuidsOk returns a tuple with the MonitorGroupGuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGroupGuids

`func (o *PublicStatusPage) SetMonitorGroupGuids(v []string)`

SetMonitorGroupGuids sets MonitorGroupGuids field to given value.

### HasMonitorGroupGuids

`func (o *PublicStatusPage) HasMonitorGroupGuids() bool`

HasMonitorGroupGuids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


