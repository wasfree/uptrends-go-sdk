# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardGuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DashboardFilter** | Pointer to [**DashboardFilter**](DashboardFilter.md) |  | [optional] 
**AutoRefresh** | Pointer to **bool** |  | [optional] 

## Methods

### NewDashboard

`func NewDashboard() *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardGuid

`func (o *Dashboard) GetDashboardGuid() string`

GetDashboardGuid returns the DashboardGuid field if non-nil, zero value otherwise.

### GetDashboardGuidOk

`func (o *Dashboard) GetDashboardGuidOk() (*string, bool)`

GetDashboardGuidOk returns a tuple with the DashboardGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardGuid

`func (o *Dashboard) SetDashboardGuid(v string)`

SetDashboardGuid sets DashboardGuid field to given value.

### HasDashboardGuid

`func (o *Dashboard) HasDashboardGuid() bool`

HasDashboardGuid returns a boolean if a field has been set.

### GetName

`func (o *Dashboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dashboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dashboard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dashboard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDashboardFilter

`func (o *Dashboard) GetDashboardFilter() DashboardFilter`

GetDashboardFilter returns the DashboardFilter field if non-nil, zero value otherwise.

### GetDashboardFilterOk

`func (o *Dashboard) GetDashboardFilterOk() (*DashboardFilter, bool)`

GetDashboardFilterOk returns a tuple with the DashboardFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardFilter

`func (o *Dashboard) SetDashboardFilter(v DashboardFilter)`

SetDashboardFilter sets DashboardFilter field to given value.

### HasDashboardFilter

`func (o *Dashboard) HasDashboardFilter() bool`

HasDashboardFilter returns a boolean if a field has been set.

### GetAutoRefresh

`func (o *Dashboard) GetAutoRefresh() bool`

GetAutoRefresh returns the AutoRefresh field if non-nil, zero value otherwise.

### GetAutoRefreshOk

`func (o *Dashboard) GetAutoRefreshOk() (*bool, bool)`

GetAutoRefreshOk returns a tuple with the AutoRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefresh

`func (o *Dashboard) SetAutoRefresh(v bool)`

SetAutoRefresh sets AutoRefresh field to given value.

### HasAutoRefresh

`func (o *Dashboard) HasAutoRefresh() bool`

HasAutoRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


