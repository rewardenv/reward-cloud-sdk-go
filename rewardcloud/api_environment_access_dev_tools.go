/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.6.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// EnvironmentAccessDevToolsApiService EnvironmentAccessDevToolsApi service
type EnvironmentAccessDevToolsApiService service

type ApiApiEnvironmentAccessDevToolsGetCollectionRequest struct {
	ctx context.Context
	ApiService *EnvironmentAccessDevToolsApiService
	page *int32
	itemsPerPage *int32
}

// The collection page number
func (r ApiApiEnvironmentAccessDevToolsGetCollectionRequest) Page(page int32) ApiApiEnvironmentAccessDevToolsGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r ApiApiEnvironmentAccessDevToolsGetCollectionRequest) ItemsPerPage(itemsPerPage int32) ApiApiEnvironmentAccessDevToolsGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

func (r ApiApiEnvironmentAccessDevToolsGetCollectionRequest) Execute() ([]EnvironmentAccessDevTools, *http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessDevToolsGetCollectionExecute(r)
}

/*
ApiEnvironmentAccessDevToolsGetCollection Retrieves the collection of EnvironmentAccessDevTools resources.

Retrieves the collection of EnvironmentAccessDevTools resources.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiEnvironmentAccessDevToolsGetCollectionRequest
*/
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsGetCollection(ctx context.Context) ApiApiEnvironmentAccessDevToolsGetCollectionRequest {
	return ApiApiEnvironmentAccessDevToolsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []EnvironmentAccessDevTools
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsGetCollectionExecute(r ApiApiEnvironmentAccessDevToolsGetCollectionRequest) ([]EnvironmentAccessDevTools, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EnvironmentAccessDevTools
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessDevToolsApiService.ApiEnvironmentAccessDevToolsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_dev_tools"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.itemsPerPage != nil {
		localVarQueryParams.Add("itemsPerPage", parameterToString(*r.itemsPerPage, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/hal+json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiEnvironmentAccessDevToolsIdDeleteRequest struct {
	ctx context.Context
	ApiService *EnvironmentAccessDevToolsApiService
	id string
}

func (r ApiApiEnvironmentAccessDevToolsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessDevToolsIdDeleteExecute(r)
}

/*
ApiEnvironmentAccessDevToolsIdDelete Removes the EnvironmentAccessDevTools resource.

Removes the EnvironmentAccessDevTools resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id EnvironmentAccessDevTools identifier
 @return ApiApiEnvironmentAccessDevToolsIdDeleteRequest
*/
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsIdDelete(ctx context.Context, id string) ApiApiEnvironmentAccessDevToolsIdDeleteRequest {
	return ApiApiEnvironmentAccessDevToolsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsIdDeleteExecute(r ApiApiEnvironmentAccessDevToolsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessDevToolsApiService.ApiEnvironmentAccessDevToolsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_dev_tools/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiEnvironmentAccessDevToolsIdGetRequest struct {
	ctx context.Context
	ApiService *EnvironmentAccessDevToolsApiService
	id string
}

func (r ApiApiEnvironmentAccessDevToolsIdGetRequest) Execute() (*EnvironmentAccessDevTools, *http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessDevToolsIdGetExecute(r)
}

/*
ApiEnvironmentAccessDevToolsIdGet Retrieves a EnvironmentAccessDevTools resource.

Retrieves a EnvironmentAccessDevTools resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id EnvironmentAccessDevTools identifier
 @return ApiApiEnvironmentAccessDevToolsIdGetRequest
*/
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsIdGet(ctx context.Context, id string) ApiApiEnvironmentAccessDevToolsIdGetRequest {
	return ApiApiEnvironmentAccessDevToolsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EnvironmentAccessDevTools
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsIdGetExecute(r ApiApiEnvironmentAccessDevToolsIdGetRequest) (*EnvironmentAccessDevTools, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentAccessDevTools
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessDevToolsApiService.ApiEnvironmentAccessDevToolsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_dev_tools/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/hal+json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiEnvironmentAccessDevToolsIdPatchRequest struct {
	ctx context.Context
	ApiService *EnvironmentAccessDevToolsApiService
	id string
	environmentAccessDevTools *EnvironmentAccessDevTools
}

// The updated EnvironmentAccessDevTools resource
func (r ApiApiEnvironmentAccessDevToolsIdPatchRequest) EnvironmentAccessDevTools(environmentAccessDevTools EnvironmentAccessDevTools) ApiApiEnvironmentAccessDevToolsIdPatchRequest {
	r.environmentAccessDevTools = &environmentAccessDevTools
	return r
}

func (r ApiApiEnvironmentAccessDevToolsIdPatchRequest) Execute() (*EnvironmentAccessDevTools, *http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessDevToolsIdPatchExecute(r)
}

/*
ApiEnvironmentAccessDevToolsIdPatch Updates the EnvironmentAccessDevTools resource.

Updates the EnvironmentAccessDevTools resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id EnvironmentAccessDevTools identifier
 @return ApiApiEnvironmentAccessDevToolsIdPatchRequest
*/
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsIdPatch(ctx context.Context, id string) ApiApiEnvironmentAccessDevToolsIdPatchRequest {
	return ApiApiEnvironmentAccessDevToolsIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EnvironmentAccessDevTools
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsIdPatchExecute(r ApiApiEnvironmentAccessDevToolsIdPatchRequest) (*EnvironmentAccessDevTools, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentAccessDevTools
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessDevToolsApiService.ApiEnvironmentAccessDevToolsIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_dev_tools/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.environmentAccessDevTools == nil {
		return localVarReturnValue, nil, reportError("environmentAccessDevTools is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/hal+json", "text/html"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/hal+json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.environmentAccessDevTools
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiEnvironmentAccessDevToolsIdPutRequest struct {
	ctx context.Context
	ApiService *EnvironmentAccessDevToolsApiService
	id string
	environmentAccessDevTools *EnvironmentAccessDevTools
}

// The updated EnvironmentAccessDevTools resource
func (r ApiApiEnvironmentAccessDevToolsIdPutRequest) EnvironmentAccessDevTools(environmentAccessDevTools EnvironmentAccessDevTools) ApiApiEnvironmentAccessDevToolsIdPutRequest {
	r.environmentAccessDevTools = &environmentAccessDevTools
	return r
}

func (r ApiApiEnvironmentAccessDevToolsIdPutRequest) Execute() (*EnvironmentAccessDevTools, *http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessDevToolsIdPutExecute(r)
}

/*
ApiEnvironmentAccessDevToolsIdPut Replaces the EnvironmentAccessDevTools resource.

Replaces the EnvironmentAccessDevTools resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id EnvironmentAccessDevTools identifier
 @return ApiApiEnvironmentAccessDevToolsIdPutRequest
*/
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsIdPut(ctx context.Context, id string) ApiApiEnvironmentAccessDevToolsIdPutRequest {
	return ApiApiEnvironmentAccessDevToolsIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EnvironmentAccessDevTools
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsIdPutExecute(r ApiApiEnvironmentAccessDevToolsIdPutRequest) (*EnvironmentAccessDevTools, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentAccessDevTools
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessDevToolsApiService.ApiEnvironmentAccessDevToolsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_dev_tools/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.environmentAccessDevTools == nil {
		return localVarReturnValue, nil, reportError("environmentAccessDevTools is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/hal+json", "text/html"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/hal+json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.environmentAccessDevTools
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiEnvironmentAccessDevToolsPostRequest struct {
	ctx context.Context
	ApiService *EnvironmentAccessDevToolsApiService
	environmentAccessDevTools *EnvironmentAccessDevTools
}

// The new EnvironmentAccessDevTools resource
func (r ApiApiEnvironmentAccessDevToolsPostRequest) EnvironmentAccessDevTools(environmentAccessDevTools EnvironmentAccessDevTools) ApiApiEnvironmentAccessDevToolsPostRequest {
	r.environmentAccessDevTools = &environmentAccessDevTools
	return r
}

func (r ApiApiEnvironmentAccessDevToolsPostRequest) Execute() (*EnvironmentAccessDevTools, *http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessDevToolsPostExecute(r)
}

/*
ApiEnvironmentAccessDevToolsPost Creates a EnvironmentAccessDevTools resource.

Creates a EnvironmentAccessDevTools resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiEnvironmentAccessDevToolsPostRequest
*/
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsPost(ctx context.Context) ApiApiEnvironmentAccessDevToolsPostRequest {
	return ApiApiEnvironmentAccessDevToolsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EnvironmentAccessDevTools
func (a *EnvironmentAccessDevToolsApiService) ApiEnvironmentAccessDevToolsPostExecute(r ApiApiEnvironmentAccessDevToolsPostRequest) (*EnvironmentAccessDevTools, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentAccessDevTools
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessDevToolsApiService.ApiEnvironmentAccessDevToolsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_dev_tools"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.environmentAccessDevTools == nil {
		return localVarReturnValue, nil, reportError("environmentAccessDevTools is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/hal+json", "text/html"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/hal+json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.environmentAccessDevTools
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
