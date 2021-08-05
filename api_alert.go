/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uptrends

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// AlertApiService AlertApi service
type AlertApiService service

type ApiAlertGetMonitorAlertsRequest struct {
	ctx _context.Context
	ApiService *AlertApiService
	monitorGuid string
	includeReminders *bool
	cursor *string
	sorting *string
	take *int32
	start *time.Time
	end *time.Time
	presetPeriod *string
}

func (r ApiAlertGetMonitorAlertsRequest) IncludeReminders(includeReminders bool) ApiAlertGetMonitorAlertsRequest {
	r.includeReminders = &includeReminders
	return r
}
func (r ApiAlertGetMonitorAlertsRequest) Cursor(cursor string) ApiAlertGetMonitorAlertsRequest {
	r.cursor = &cursor
	return r
}
func (r ApiAlertGetMonitorAlertsRequest) Sorting(sorting string) ApiAlertGetMonitorAlertsRequest {
	r.sorting = &sorting
	return r
}
func (r ApiAlertGetMonitorAlertsRequest) Take(take int32) ApiAlertGetMonitorAlertsRequest {
	r.take = &take
	return r
}
func (r ApiAlertGetMonitorAlertsRequest) Start(start time.Time) ApiAlertGetMonitorAlertsRequest {
	r.start = &start
	return r
}
func (r ApiAlertGetMonitorAlertsRequest) End(end time.Time) ApiAlertGetMonitorAlertsRequest {
	r.end = &end
	return r
}
func (r ApiAlertGetMonitorAlertsRequest) PresetPeriod(presetPeriod string) ApiAlertGetMonitorAlertsRequest {
	r.presetPeriod = &presetPeriod
	return r
}

func (r ApiAlertGetMonitorAlertsRequest) Execute() (AlertResponse, *_nethttp.Response, error) {
	return r.ApiService.AlertGetMonitorAlertsExecute(r)
}

/*
 * AlertGetMonitorAlerts Returns alerts for a specific monitor.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param monitorGuid The Guid of the monitor to get alerts for.
 * @return ApiAlertGetMonitorAlertsRequest
 */
func (a *AlertApiService) AlertGetMonitorAlerts(ctx _context.Context, monitorGuid string) ApiAlertGetMonitorAlertsRequest {
	return ApiAlertGetMonitorAlertsRequest{
		ApiService: a,
		ctx: ctx,
		monitorGuid: monitorGuid,
	}
}

/*
 * Execute executes the request
 * @return AlertResponse
 */
func (a *AlertApiService) AlertGetMonitorAlertsExecute(r ApiAlertGetMonitorAlertsRequest) (AlertResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AlertResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertApiService.AlertGetMonitorAlerts")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Alert/Monitor/{monitorGuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorGuid"+"}", _neturl.PathEscape(parameterToString(r.monitorGuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.includeReminders != nil {
		localVarQueryParams.Add("IncludeReminders", parameterToString(*r.includeReminders, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("Cursor", parameterToString(*r.cursor, ""))
	}
	if r.sorting != nil {
		localVarQueryParams.Add("Sorting", parameterToString(*r.sorting, ""))
	}
	if r.take != nil {
		localVarQueryParams.Add("Take", parameterToString(*r.take, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("Start", parameterToString(*r.start, ""))
	}
	if r.end != nil {
		localVarQueryParams.Add("End", parameterToString(*r.end, ""))
	}
	if r.presetPeriod != nil {
		localVarQueryParams.Add("PresetPeriod", parameterToString(*r.presetPeriod, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v MessageList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v MessageList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAlertGetMonitorGroupAlertsRequest struct {
	ctx _context.Context
	ApiService *AlertApiService
	monitorGroupGuid string
	includeReminders *bool
	cursor *string
	sorting *string
	take *int32
	start *time.Time
	end *time.Time
	presetPeriod *string
}

func (r ApiAlertGetMonitorGroupAlertsRequest) IncludeReminders(includeReminders bool) ApiAlertGetMonitorGroupAlertsRequest {
	r.includeReminders = &includeReminders
	return r
}
func (r ApiAlertGetMonitorGroupAlertsRequest) Cursor(cursor string) ApiAlertGetMonitorGroupAlertsRequest {
	r.cursor = &cursor
	return r
}
func (r ApiAlertGetMonitorGroupAlertsRequest) Sorting(sorting string) ApiAlertGetMonitorGroupAlertsRequest {
	r.sorting = &sorting
	return r
}
func (r ApiAlertGetMonitorGroupAlertsRequest) Take(take int32) ApiAlertGetMonitorGroupAlertsRequest {
	r.take = &take
	return r
}
func (r ApiAlertGetMonitorGroupAlertsRequest) Start(start time.Time) ApiAlertGetMonitorGroupAlertsRequest {
	r.start = &start
	return r
}
func (r ApiAlertGetMonitorGroupAlertsRequest) End(end time.Time) ApiAlertGetMonitorGroupAlertsRequest {
	r.end = &end
	return r
}
func (r ApiAlertGetMonitorGroupAlertsRequest) PresetPeriod(presetPeriod string) ApiAlertGetMonitorGroupAlertsRequest {
	r.presetPeriod = &presetPeriod
	return r
}

func (r ApiAlertGetMonitorGroupAlertsRequest) Execute() (AlertResponse, *_nethttp.Response, error) {
	return r.ApiService.AlertGetMonitorGroupAlertsExecute(r)
}

/*
 * AlertGetMonitorGroupAlerts Returns alerts for a specific monitor group.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param monitorGroupGuid The Guid of the monitor group to get alerts for.
 * @return ApiAlertGetMonitorGroupAlertsRequest
 */
func (a *AlertApiService) AlertGetMonitorGroupAlerts(ctx _context.Context, monitorGroupGuid string) ApiAlertGetMonitorGroupAlertsRequest {
	return ApiAlertGetMonitorGroupAlertsRequest{
		ApiService: a,
		ctx: ctx,
		monitorGroupGuid: monitorGroupGuid,
	}
}

/*
 * Execute executes the request
 * @return AlertResponse
 */
func (a *AlertApiService) AlertGetMonitorGroupAlertsExecute(r ApiAlertGetMonitorGroupAlertsRequest) (AlertResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AlertResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertApiService.AlertGetMonitorGroupAlerts")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Alert/MonitorGroup/{monitorGroupGuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorGroupGuid"+"}", _neturl.PathEscape(parameterToString(r.monitorGroupGuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.includeReminders != nil {
		localVarQueryParams.Add("IncludeReminders", parameterToString(*r.includeReminders, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("Cursor", parameterToString(*r.cursor, ""))
	}
	if r.sorting != nil {
		localVarQueryParams.Add("Sorting", parameterToString(*r.sorting, ""))
	}
	if r.take != nil {
		localVarQueryParams.Add("Take", parameterToString(*r.take, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("Start", parameterToString(*r.start, ""))
	}
	if r.end != nil {
		localVarQueryParams.Add("End", parameterToString(*r.end, ""))
	}
	if r.presetPeriod != nil {
		localVarQueryParams.Add("PresetPeriod", parameterToString(*r.presetPeriod, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v MessageList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v MessageList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}