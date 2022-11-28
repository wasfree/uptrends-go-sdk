# HttpCheckDetailsAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseBody** | Pointer to **string** | The HTML code retrieved from the target | [optional] 
**ResponseHeaders** | Pointer to **string** | The HTTP response headers retrieved from the target  | [optional] 
**Url** | Pointer to **string** | The URL of the HTTP Check.  | [optional] 

## Methods

### NewHttpCheckDetailsAttributes

`func NewHttpCheckDetailsAttributes() *HttpCheckDetailsAttributes`

NewHttpCheckDetailsAttributes instantiates a new HttpCheckDetailsAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpCheckDetailsAttributesWithDefaults

`func NewHttpCheckDetailsAttributesWithDefaults() *HttpCheckDetailsAttributes`

NewHttpCheckDetailsAttributesWithDefaults instantiates a new HttpCheckDetailsAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseBody

`func (o *HttpCheckDetailsAttributes) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *HttpCheckDetailsAttributes) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *HttpCheckDetailsAttributes) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *HttpCheckDetailsAttributes) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *HttpCheckDetailsAttributes) GetResponseHeaders() string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *HttpCheckDetailsAttributes) GetResponseHeadersOk() (*string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *HttpCheckDetailsAttributes) SetResponseHeaders(v string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *HttpCheckDetailsAttributes) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetUrl

`func (o *HttpCheckDetailsAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpCheckDetailsAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpCheckDetailsAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HttpCheckDetailsAttributes) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


