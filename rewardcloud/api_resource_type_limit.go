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


// ResourceTypeLimitApiService ResourceTypeLimitApi service
type ResourceTypeLimitApiService service

type ApiApiResourceTypeLimitsGetCollectionRequest struct {
	ctx context.Context
	ApiService *ResourceTypeLimitApiService
	page *int32
	itemsPerPage *int32
}

// The collection page number
func (r ApiApiResourceTypeLimitsGetCollectionRequest) Page(page int32) ApiApiResourceTypeLimitsGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r ApiApiResourceTypeLimitsGetCollectionRequest) ItemsPerPage(itemsPerPage int32) ApiApiResourceTypeLimitsGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

func (r ApiApiResourceTypeLimitsGetCollectionRequest) Execute() ([]ResourceTypeLimit, *http.Response, error) {
	return r.ApiService.ApiResourceTypeLimitsGetCollectionExecute(r)
}

/*
ApiResourceTypeLimitsGetCollection Retrieves the collection of ResourceTypeLimit resources.

Retrieves the collection of ResourceTypeLimit resources.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiResourceTypeLimitsGetCollectionRequest
*/
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsGetCollection(ctx context.Context) ApiApiResourceTypeLimitsGetCollectionRequest {
	return ApiApiResourceTypeLimitsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ResourceTypeLimit
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsGetCollectionExecute(r ApiApiResourceTypeLimitsGetCollectionRequest) ([]ResourceTypeLimit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ResourceTypeLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceTypeLimitApiService.ApiResourceTypeLimitsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resource_type_limits"

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

type ApiApiResourceTypeLimitsIdDeleteRequest struct {
	ctx context.Context
	ApiService *ResourceTypeLimitApiService
	id string
}

func (r ApiApiResourceTypeLimitsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiResourceTypeLimitsIdDeleteExecute(r)
}

/*
ApiResourceTypeLimitsIdDelete Removes the ResourceTypeLimit resource.

Removes the ResourceTypeLimit resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ResourceTypeLimit identifier
 @return ApiApiResourceTypeLimitsIdDeleteRequest
*/
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsIdDelete(ctx context.Context, id string) ApiApiResourceTypeLimitsIdDeleteRequest {
	return ApiApiResourceTypeLimitsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsIdDeleteExecute(r ApiApiResourceTypeLimitsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceTypeLimitApiService.ApiResourceTypeLimitsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resource_type_limits/{id}"
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

type ApiApiResourceTypeLimitsIdGetRequest struct {
	ctx context.Context
	ApiService *ResourceTypeLimitApiService
	id string
}

func (r ApiApiResourceTypeLimitsIdGetRequest) Execute() (*ResourceTypeLimit, *http.Response, error) {
	return r.ApiService.ApiResourceTypeLimitsIdGetExecute(r)
}

/*
ApiResourceTypeLimitsIdGet Retrieves a ResourceTypeLimit resource.

Retrieves a ResourceTypeLimit resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ResourceTypeLimit identifier
 @return ApiApiResourceTypeLimitsIdGetRequest
*/
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsIdGet(ctx context.Context, id string) ApiApiResourceTypeLimitsIdGetRequest {
	return ApiApiResourceTypeLimitsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ResourceTypeLimit
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsIdGetExecute(r ApiApiResourceTypeLimitsIdGetRequest) (*ResourceTypeLimit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourceTypeLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceTypeLimitApiService.ApiResourceTypeLimitsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resource_type_limits/{id}"
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

type ApiApiResourceTypeLimitsIdPatchRequest struct {
	ctx context.Context
	ApiService *ResourceTypeLimitApiService
	id string
	resourceTypeLimit *ResourceTypeLimit
}

// The updated ResourceTypeLimit resource
func (r ApiApiResourceTypeLimitsIdPatchRequest) ResourceTypeLimit(resourceTypeLimit ResourceTypeLimit) ApiApiResourceTypeLimitsIdPatchRequest {
	r.resourceTypeLimit = &resourceTypeLimit
	return r
}

func (r ApiApiResourceTypeLimitsIdPatchRequest) Execute() (*ResourceTypeLimit, *http.Response, error) {
	return r.ApiService.ApiResourceTypeLimitsIdPatchExecute(r)
}

/*
ApiResourceTypeLimitsIdPatch Updates the ResourceTypeLimit resource.

Updates the ResourceTypeLimit resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ResourceTypeLimit identifier
 @return ApiApiResourceTypeLimitsIdPatchRequest
*/
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsIdPatch(ctx context.Context, id string) ApiApiResourceTypeLimitsIdPatchRequest {
	return ApiApiResourceTypeLimitsIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ResourceTypeLimit
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsIdPatchExecute(r ApiApiResourceTypeLimitsIdPatchRequest) (*ResourceTypeLimit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourceTypeLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceTypeLimitApiService.ApiResourceTypeLimitsIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resource_type_limits/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.resourceTypeLimit == nil {
		return localVarReturnValue, nil, reportError("resourceTypeLimit is required and must be specified")
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
	localVarPostBody = r.resourceTypeLimit
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

type ApiApiResourceTypeLimitsIdPutRequest struct {
	ctx context.Context
	ApiService *ResourceTypeLimitApiService
	id string
	resourceTypeLimit *ResourceTypeLimit
}

// The updated ResourceTypeLimit resource
func (r ApiApiResourceTypeLimitsIdPutRequest) ResourceTypeLimit(resourceTypeLimit ResourceTypeLimit) ApiApiResourceTypeLimitsIdPutRequest {
	r.resourceTypeLimit = &resourceTypeLimit
	return r
}

func (r ApiApiResourceTypeLimitsIdPutRequest) Execute() (*ResourceTypeLimit, *http.Response, error) {
	return r.ApiService.ApiResourceTypeLimitsIdPutExecute(r)
}

/*
ApiResourceTypeLimitsIdPut Replaces the ResourceTypeLimit resource.

Replaces the ResourceTypeLimit resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ResourceTypeLimit identifier
 @return ApiApiResourceTypeLimitsIdPutRequest
*/
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsIdPut(ctx context.Context, id string) ApiApiResourceTypeLimitsIdPutRequest {
	return ApiApiResourceTypeLimitsIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ResourceTypeLimit
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsIdPutExecute(r ApiApiResourceTypeLimitsIdPutRequest) (*ResourceTypeLimit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourceTypeLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceTypeLimitApiService.ApiResourceTypeLimitsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resource_type_limits/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.resourceTypeLimit == nil {
		return localVarReturnValue, nil, reportError("resourceTypeLimit is required and must be specified")
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
	localVarPostBody = r.resourceTypeLimit
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

type ApiApiResourceTypeLimitsPostRequest struct {
	ctx context.Context
	ApiService *ResourceTypeLimitApiService
	resourceTypeLimit *ResourceTypeLimit
}

// The new ResourceTypeLimit resource
func (r ApiApiResourceTypeLimitsPostRequest) ResourceTypeLimit(resourceTypeLimit ResourceTypeLimit) ApiApiResourceTypeLimitsPostRequest {
	r.resourceTypeLimit = &resourceTypeLimit
	return r
}

func (r ApiApiResourceTypeLimitsPostRequest) Execute() (*ResourceTypeLimit, *http.Response, error) {
	return r.ApiService.ApiResourceTypeLimitsPostExecute(r)
}

/*
ApiResourceTypeLimitsPost Creates a ResourceTypeLimit resource.

Creates a ResourceTypeLimit resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiResourceTypeLimitsPostRequest
*/
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsPost(ctx context.Context) ApiApiResourceTypeLimitsPostRequest {
	return ApiApiResourceTypeLimitsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResourceTypeLimit
func (a *ResourceTypeLimitApiService) ApiResourceTypeLimitsPostExecute(r ApiApiResourceTypeLimitsPostRequest) (*ResourceTypeLimit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourceTypeLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceTypeLimitApiService.ApiResourceTypeLimitsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resource_type_limits"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.resourceTypeLimit == nil {
		return localVarReturnValue, nil, reportError("resourceTypeLimit is required and must be specified")
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
	localVarPostBody = r.resourceTypeLimit
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
