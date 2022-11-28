# CustomizationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MainColor** | Pointer to **string** |  | [optional] 
**BackgroundColor** | Pointer to **string** |  | [optional] 
**TextColor** | Pointer to **string** |  | [optional] 
**TitleText** | Pointer to **string** |  | [optional] 
**FooterText** | Pointer to **string** |  | [optional] 
**MonitorNameOverrideFieldName** | Pointer to **string** |  | [optional] 
**SortColumnsNewToOld** | **bool** |  | 
**SortRowsProperty** | [**SortOrderEnum**](SortOrderEnum.md) |  | 
**CommentTitle** | Pointer to **string** |  | [optional] 
**CommentText** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomizationInfo

`func NewCustomizationInfo(sortColumnsNewToOld bool, sortRowsProperty SortOrderEnum, ) *CustomizationInfo`

NewCustomizationInfo instantiates a new CustomizationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomizationInfoWithDefaults

`func NewCustomizationInfoWithDefaults() *CustomizationInfo`

NewCustomizationInfoWithDefaults instantiates a new CustomizationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMainColor

`func (o *CustomizationInfo) GetMainColor() string`

GetMainColor returns the MainColor field if non-nil, zero value otherwise.

### GetMainColorOk

`func (o *CustomizationInfo) GetMainColorOk() (*string, bool)`

GetMainColorOk returns a tuple with the MainColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainColor

`func (o *CustomizationInfo) SetMainColor(v string)`

SetMainColor sets MainColor field to given value.

### HasMainColor

`func (o *CustomizationInfo) HasMainColor() bool`

HasMainColor returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *CustomizationInfo) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *CustomizationInfo) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *CustomizationInfo) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *CustomizationInfo) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetTextColor

`func (o *CustomizationInfo) GetTextColor() string`

GetTextColor returns the TextColor field if non-nil, zero value otherwise.

### GetTextColorOk

`func (o *CustomizationInfo) GetTextColorOk() (*string, bool)`

GetTextColorOk returns a tuple with the TextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextColor

`func (o *CustomizationInfo) SetTextColor(v string)`

SetTextColor sets TextColor field to given value.

### HasTextColor

`func (o *CustomizationInfo) HasTextColor() bool`

HasTextColor returns a boolean if a field has been set.

### GetTitleText

`func (o *CustomizationInfo) GetTitleText() string`

GetTitleText returns the TitleText field if non-nil, zero value otherwise.

### GetTitleTextOk

`func (o *CustomizationInfo) GetTitleTextOk() (*string, bool)`

GetTitleTextOk returns a tuple with the TitleText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleText

`func (o *CustomizationInfo) SetTitleText(v string)`

SetTitleText sets TitleText field to given value.

### HasTitleText

`func (o *CustomizationInfo) HasTitleText() bool`

HasTitleText returns a boolean if a field has been set.

### GetFooterText

`func (o *CustomizationInfo) GetFooterText() string`

GetFooterText returns the FooterText field if non-nil, zero value otherwise.

### GetFooterTextOk

`func (o *CustomizationInfo) GetFooterTextOk() (*string, bool)`

GetFooterTextOk returns a tuple with the FooterText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterText

`func (o *CustomizationInfo) SetFooterText(v string)`

SetFooterText sets FooterText field to given value.

### HasFooterText

`func (o *CustomizationInfo) HasFooterText() bool`

HasFooterText returns a boolean if a field has been set.

### GetMonitorNameOverrideFieldName

`func (o *CustomizationInfo) GetMonitorNameOverrideFieldName() string`

GetMonitorNameOverrideFieldName returns the MonitorNameOverrideFieldName field if non-nil, zero value otherwise.

### GetMonitorNameOverrideFieldNameOk

`func (o *CustomizationInfo) GetMonitorNameOverrideFieldNameOk() (*string, bool)`

GetMonitorNameOverrideFieldNameOk returns a tuple with the MonitorNameOverrideFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorNameOverrideFieldName

`func (o *CustomizationInfo) SetMonitorNameOverrideFieldName(v string)`

SetMonitorNameOverrideFieldName sets MonitorNameOverrideFieldName field to given value.

### HasMonitorNameOverrideFieldName

`func (o *CustomizationInfo) HasMonitorNameOverrideFieldName() bool`

HasMonitorNameOverrideFieldName returns a boolean if a field has been set.

### GetSortColumnsNewToOld

`func (o *CustomizationInfo) GetSortColumnsNewToOld() bool`

GetSortColumnsNewToOld returns the SortColumnsNewToOld field if non-nil, zero value otherwise.

### GetSortColumnsNewToOldOk

`func (o *CustomizationInfo) GetSortColumnsNewToOldOk() (*bool, bool)`

GetSortColumnsNewToOldOk returns a tuple with the SortColumnsNewToOld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortColumnsNewToOld

`func (o *CustomizationInfo) SetSortColumnsNewToOld(v bool)`

SetSortColumnsNewToOld sets SortColumnsNewToOld field to given value.


### GetSortRowsProperty

`func (o *CustomizationInfo) GetSortRowsProperty() SortOrderEnum`

GetSortRowsProperty returns the SortRowsProperty field if non-nil, zero value otherwise.

### GetSortRowsPropertyOk

`func (o *CustomizationInfo) GetSortRowsPropertyOk() (*SortOrderEnum, bool)`

GetSortRowsPropertyOk returns a tuple with the SortRowsProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortRowsProperty

`func (o *CustomizationInfo) SetSortRowsProperty(v SortOrderEnum)`

SetSortRowsProperty sets SortRowsProperty field to given value.


### GetCommentTitle

`func (o *CustomizationInfo) GetCommentTitle() string`

GetCommentTitle returns the CommentTitle field if non-nil, zero value otherwise.

### GetCommentTitleOk

`func (o *CustomizationInfo) GetCommentTitleOk() (*string, bool)`

GetCommentTitleOk returns a tuple with the CommentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentTitle

`func (o *CustomizationInfo) SetCommentTitle(v string)`

SetCommentTitle sets CommentTitle field to given value.

### HasCommentTitle

`func (o *CustomizationInfo) HasCommentTitle() bool`

HasCommentTitle returns a boolean if a field has been set.

### GetCommentText

`func (o *CustomizationInfo) GetCommentText() string`

GetCommentText returns the CommentText field if non-nil, zero value otherwise.

### GetCommentTextOk

`func (o *CustomizationInfo) GetCommentTextOk() (*string, bool)`

GetCommentTextOk returns a tuple with the CommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentText

`func (o *CustomizationInfo) SetCommentText(v string)`

SetCommentText sets CommentText field to given value.

### HasCommentText

`func (o *CustomizationInfo) HasCommentText() bool`

HasCommentText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


