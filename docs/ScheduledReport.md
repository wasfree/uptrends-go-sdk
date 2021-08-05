# ScheduledReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledReportGuid** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**DashboardGuid** | Pointer to **string** |  | [optional] 
**DashboardGuidSpecified** | Pointer to **bool** |  | [optional] 
**FileType** | Pointer to [**ScheduledReportFileType**](ScheduledReportFileType.md) |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**Schedule**](Schedule.md) |  | [optional] 
**SelectedPeriod** | Pointer to [**PresetPeriodTypeWithExclusive**](PresetPeriodTypeWithExclusive.md) |  | [optional] 
**InternalNotes** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Recipients** | Pointer to [**Recipients**](Recipients.md) |  | [optional] 

## Methods

### NewScheduledReport

`func NewScheduledReport() *ScheduledReport`

NewScheduledReport instantiates a new ScheduledReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledReportWithDefaults

`func NewScheduledReportWithDefaults() *ScheduledReport`

NewScheduledReportWithDefaults instantiates a new ScheduledReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledReportGuid

`func (o *ScheduledReport) GetScheduledReportGuid() string`

GetScheduledReportGuid returns the ScheduledReportGuid field if non-nil, zero value otherwise.

### GetScheduledReportGuidOk

`func (o *ScheduledReport) GetScheduledReportGuidOk() (*string, bool)`

GetScheduledReportGuidOk returns a tuple with the ScheduledReportGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReportGuid

`func (o *ScheduledReport) SetScheduledReportGuid(v string)`

SetScheduledReportGuid sets ScheduledReportGuid field to given value.

### HasScheduledReportGuid

`func (o *ScheduledReport) HasScheduledReportGuid() bool`

HasScheduledReportGuid returns a boolean if a field has been set.

### GetHash

`func (o *ScheduledReport) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ScheduledReport) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ScheduledReport) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ScheduledReport) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetDashboardGuid

`func (o *ScheduledReport) GetDashboardGuid() string`

GetDashboardGuid returns the DashboardGuid field if non-nil, zero value otherwise.

### GetDashboardGuidOk

`func (o *ScheduledReport) GetDashboardGuidOk() (*string, bool)`

GetDashboardGuidOk returns a tuple with the DashboardGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardGuid

`func (o *ScheduledReport) SetDashboardGuid(v string)`

SetDashboardGuid sets DashboardGuid field to given value.

### HasDashboardGuid

`func (o *ScheduledReport) HasDashboardGuid() bool`

HasDashboardGuid returns a boolean if a field has been set.

### GetDashboardGuidSpecified

`func (o *ScheduledReport) GetDashboardGuidSpecified() bool`

GetDashboardGuidSpecified returns the DashboardGuidSpecified field if non-nil, zero value otherwise.

### GetDashboardGuidSpecifiedOk

`func (o *ScheduledReport) GetDashboardGuidSpecifiedOk() (*bool, bool)`

GetDashboardGuidSpecifiedOk returns a tuple with the DashboardGuidSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardGuidSpecified

`func (o *ScheduledReport) SetDashboardGuidSpecified(v bool)`

SetDashboardGuidSpecified sets DashboardGuidSpecified field to given value.

### HasDashboardGuidSpecified

`func (o *ScheduledReport) HasDashboardGuidSpecified() bool`

HasDashboardGuidSpecified returns a boolean if a field has been set.

### GetFileType

`func (o *ScheduledReport) GetFileType() ScheduledReportFileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *ScheduledReport) GetFileTypeOk() (*ScheduledReportFileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *ScheduledReport) SetFileType(v ScheduledReportFileType)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *ScheduledReport) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetIsActive

`func (o *ScheduledReport) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ScheduledReport) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ScheduledReport) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ScheduledReport) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetSchedule

`func (o *ScheduledReport) GetSchedule() Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScheduledReport) GetScheduleOk() (*Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScheduledReport) SetSchedule(v Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ScheduledReport) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSelectedPeriod

`func (o *ScheduledReport) GetSelectedPeriod() PresetPeriodTypeWithExclusive`

GetSelectedPeriod returns the SelectedPeriod field if non-nil, zero value otherwise.

### GetSelectedPeriodOk

`func (o *ScheduledReport) GetSelectedPeriodOk() (*PresetPeriodTypeWithExclusive, bool)`

GetSelectedPeriodOk returns a tuple with the SelectedPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPeriod

`func (o *ScheduledReport) SetSelectedPeriod(v PresetPeriodTypeWithExclusive)`

SetSelectedPeriod sets SelectedPeriod field to given value.

### HasSelectedPeriod

`func (o *ScheduledReport) HasSelectedPeriod() bool`

HasSelectedPeriod returns a boolean if a field has been set.

### GetInternalNotes

`func (o *ScheduledReport) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *ScheduledReport) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *ScheduledReport) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *ScheduledReport) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetNotes

`func (o *ScheduledReport) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ScheduledReport) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ScheduledReport) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ScheduledReport) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRecipients

`func (o *ScheduledReport) GetRecipients() Recipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ScheduledReport) GetRecipientsOk() (*Recipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ScheduledReport) SetRecipients(v Recipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *ScheduledReport) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


