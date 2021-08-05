# HttpAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseBody** | Pointer to **string** | The HTML code retrieved from the target | [optional] 
**ResponseHeaders** | Pointer to **string** | The HTTP response headers retrieved from the target  | [optional] 

## Methods

### NewHttpAttributes

`func NewHttpAttributes() *HttpAttributes`

NewHttpAttributes instantiates a new HttpAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpAttributesWithDefaults

`func NewHttpAttributesWithDefaults() *HttpAttributes`

NewHttpAttributesWithDefaults instantiates a new HttpAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseBody

`func (o *HttpAttributes) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *HttpAttributes) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *HttpAttributes) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *HttpAttributes) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *HttpAttributes) GetResponseHeaders() string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *HttpAttributes) GetResponseHeadersOk() (*string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *HttpAttributes) SetResponseHeaders(v string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *HttpAttributes) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


