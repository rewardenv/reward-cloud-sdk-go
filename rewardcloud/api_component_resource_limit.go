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
	"reflect"
)


// ComponentResourceLimitApiService ComponentResourceLimitApi service
type ComponentResourceLimitApiService service

type ApiApiComponentResourceLimitsGetCollectionRequest struct {
	ctx context.Context
	ApiService *ComponentResourceLimitApiService
	page *int32
	itemsPerPage *int32
	projectTypeVersion *string
	projectTypeVersion2 *[]string
	resourceType *string
	resourceType2 *[]string
	componentVersion *string
	componentVersion2 *[]string
}

// The collection page number
func (r ApiApiComponentResourceLimitsGetCollectionRequest) Page(page int32) ApiApiComponentResourceLimitsGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r ApiApiComponentResourceLimitsGetCollectionRequest) ItemsPerPage(itemsPerPage int32) ApiApiComponentResourceLimitsGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// 
func (r ApiApiComponentResourceLimitsGetCollectionRequest) ProjectTypeVersion(projectTypeVersion string) ApiApiComponentResourceLimitsGetCollectionRequest {
	r.projectTypeVersion = &projectTypeVersion
	return r
}

// 
func (r ApiApiComponentResourceLimitsGetCollectionRequest) ProjectTypeVersion2(projectTypeVersion2 []string) ApiApiComponentResourceLimitsGetCollectionRequest {
	r.projectTypeVersion2 = &projectTypeVersion2
	return r
}

// 
func (r ApiApiComponentResourceLimitsGetCollectionRequest) ResourceType(resourceType string) ApiApiComponentResourceLimitsGetCollectionRequest {
	r.resourceType = &resourceType
	return r
}

// 
func (r ApiApiComponentResourceLimitsGetCollectionRequest) ResourceType2(resourceType2 []string) ApiApiComponentResourceLimitsGetCollectionRequest {
	r.resourceType2 = &resourceType2
	return r
}

// 
func (r ApiApiComponentResourceLimitsGetCollectionRequest) ComponentVersion(componentVersion string) ApiApiComponentResourceLimitsGetCollectionRequest {
	r.componentVersion = &componentVersion
	return r
}

// 
func (r ApiApiComponentResourceLimitsGetCollectionRequest) ComponentVersion2(componentVersion2 []string) ApiApiComponentResourceLimitsGetCollectionRequest {
	r.componentVersion2 = &componentVersion2
	return r
}

func (r ApiApiComponentResourceLimitsGetCollectionRequest) Execute() ([]ComponentResourceLimit, *http.Response, error) {
	return r.ApiService.ApiComponentResourceLimitsGetCollectionExecute(r)
}

/*
ApiComponentResourceLimitsGetCollection Retrieves the collection of ComponentResourceLimit resources.

Retrieves the collection of ComponentResourceLimit resources.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiComponentResourceLimitsGetCollectionRequest
*/
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsGetCollection(ctx context.Context) ApiApiComponentResourceLimitsGetCollectionRequest {
	return ApiApiComponentResourceLimitsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ComponentResourceLimit
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsGetCollectionExecute(r ApiApiComponentResourceLimitsGetCollectionRequest) ([]ComponentResourceLimit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ComponentResourceLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentResourceLimitApiService.ApiComponentResourceLimitsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_resource_limits"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.itemsPerPage != nil {
		localVarQueryParams.Add("itemsPerPage", parameterToString(*r.itemsPerPage, ""))
	}
	if r.projectTypeVersion != nil {
		localVarQueryParams.Add("projectTypeVersion", parameterToString(*r.projectTypeVersion, ""))
	}
	if r.projectTypeVersion2 != nil {
		t := *r.projectTypeVersion2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("projectTypeVersion[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("projectTypeVersion[]", parameterToString(t, "multi"))
		}
	}
	if r.resourceType != nil {
		localVarQueryParams.Add("resourceType", parameterToString(*r.resourceType, ""))
	}
	if r.resourceType2 != nil {
		t := *r.resourceType2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("resourceType[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("resourceType[]", parameterToString(t, "multi"))
		}
	}
	if r.componentVersion != nil {
		localVarQueryParams.Add("componentVersion", parameterToString(*r.componentVersion, ""))
	}
	if r.componentVersion2 != nil {
		t := *r.componentVersion2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("componentVersion[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("componentVersion[]", parameterToString(t, "multi"))
		}
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

type ApiApiComponentResourceLimitsIdDeleteRequest struct {
	ctx context.Context
	ApiService *ComponentResourceLimitApiService
	id string
}

func (r ApiApiComponentResourceLimitsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiComponentResourceLimitsIdDeleteExecute(r)
}

/*
ApiComponentResourceLimitsIdDelete Removes the ComponentResourceLimit resource.

Removes the ComponentResourceLimit resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ComponentResourceLimit identifier
 @return ApiApiComponentResourceLimitsIdDeleteRequest
*/
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsIdDelete(ctx context.Context, id string) ApiApiComponentResourceLimitsIdDeleteRequest {
	return ApiApiComponentResourceLimitsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsIdDeleteExecute(r ApiApiComponentResourceLimitsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentResourceLimitApiService.ApiComponentResourceLimitsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_resource_limits/{id}"
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

type ApiApiComponentResourceLimitsIdGetRequest struct {
	ctx context.Context
	ApiService *ComponentResourceLimitApiService
	id string
}

func (r ApiApiComponentResourceLimitsIdGetRequest) Execute() (*ComponentResourceLimit, *http.Response, error) {
	return r.ApiService.ApiComponentResourceLimitsIdGetExecute(r)
}

/*
ApiComponentResourceLimitsIdGet Retrieves a ComponentResourceLimit resource.

Retrieves a ComponentResourceLimit resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ComponentResourceLimit identifier
 @return ApiApiComponentResourceLimitsIdGetRequest
*/
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsIdGet(ctx context.Context, id string) ApiApiComponentResourceLimitsIdGetRequest {
	return ApiApiComponentResourceLimitsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ComponentResourceLimit
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsIdGetExecute(r ApiApiComponentResourceLimitsIdGetRequest) (*ComponentResourceLimit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ComponentResourceLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentResourceLimitApiService.ApiComponentResourceLimitsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_resource_limits/{id}"
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

type ApiApiComponentResourceLimitsIdPatchRequest struct {
	ctx context.Context
	ApiService *ComponentResourceLimitApiService
	id string
	componentResourceLimit *ComponentResourceLimit
}

// The updated ComponentResourceLimit resource
func (r ApiApiComponentResourceLimitsIdPatchRequest) ComponentResourceLimit(componentResourceLimit ComponentResourceLimit) ApiApiComponentResourceLimitsIdPatchRequest {
	r.componentResourceLimit = &componentResourceLimit
	return r
}

func (r ApiApiComponentResourceLimitsIdPatchRequest) Execute() (*ComponentResourceLimit, *http.Response, error) {
	return r.ApiService.ApiComponentResourceLimitsIdPatchExecute(r)
}

/*
ApiComponentResourceLimitsIdPatch Updates the ComponentResourceLimit resource.

Updates the ComponentResourceLimit resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ComponentResourceLimit identifier
 @return ApiApiComponentResourceLimitsIdPatchRequest
*/
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsIdPatch(ctx context.Context, id string) ApiApiComponentResourceLimitsIdPatchRequest {
	return ApiApiComponentResourceLimitsIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ComponentResourceLimit
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsIdPatchExecute(r ApiApiComponentResourceLimitsIdPatchRequest) (*ComponentResourceLimit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ComponentResourceLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentResourceLimitApiService.ApiComponentResourceLimitsIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_resource_limits/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.componentResourceLimit == nil {
		return localVarReturnValue, nil, reportError("componentResourceLimit is required and must be specified")
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
	localVarPostBody = r.componentResourceLimit
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

type ApiApiComponentResourceLimitsIdPutRequest struct {
	ctx context.Context
	ApiService *ComponentResourceLimitApiService
	id string
	componentResourceLimit *ComponentResourceLimit
}

// The updated ComponentResourceLimit resource
func (r ApiApiComponentResourceLimitsIdPutRequest) ComponentResourceLimit(componentResourceLimit ComponentResourceLimit) ApiApiComponentResourceLimitsIdPutRequest {
	r.componentResourceLimit = &componentResourceLimit
	return r
}

func (r ApiApiComponentResourceLimitsIdPutRequest) Execute() (*ComponentResourceLimit, *http.Response, error) {
	return r.ApiService.ApiComponentResourceLimitsIdPutExecute(r)
}

/*
ApiComponentResourceLimitsIdPut Replaces the ComponentResourceLimit resource.

Replaces the ComponentResourceLimit resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ComponentResourceLimit identifier
 @return ApiApiComponentResourceLimitsIdPutRequest
*/
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsIdPut(ctx context.Context, id string) ApiApiComponentResourceLimitsIdPutRequest {
	return ApiApiComponentResourceLimitsIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ComponentResourceLimit
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsIdPutExecute(r ApiApiComponentResourceLimitsIdPutRequest) (*ComponentResourceLimit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ComponentResourceLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentResourceLimitApiService.ApiComponentResourceLimitsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_resource_limits/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.componentResourceLimit == nil {
		return localVarReturnValue, nil, reportError("componentResourceLimit is required and must be specified")
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
	localVarPostBody = r.componentResourceLimit
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

type ApiApiComponentResourceLimitsPostRequest struct {
	ctx context.Context
	ApiService *ComponentResourceLimitApiService
	componentResourceLimit *ComponentResourceLimit
}

// The new ComponentResourceLimit resource
func (r ApiApiComponentResourceLimitsPostRequest) ComponentResourceLimit(componentResourceLimit ComponentResourceLimit) ApiApiComponentResourceLimitsPostRequest {
	r.componentResourceLimit = &componentResourceLimit
	return r
}

func (r ApiApiComponentResourceLimitsPostRequest) Execute() (*ComponentResourceLimit, *http.Response, error) {
	return r.ApiService.ApiComponentResourceLimitsPostExecute(r)
}

/*
ApiComponentResourceLimitsPost Creates a ComponentResourceLimit resource.

Creates a ComponentResourceLimit resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiComponentResourceLimitsPostRequest
*/
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsPost(ctx context.Context) ApiApiComponentResourceLimitsPostRequest {
	return ApiApiComponentResourceLimitsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ComponentResourceLimit
func (a *ComponentResourceLimitApiService) ApiComponentResourceLimitsPostExecute(r ApiApiComponentResourceLimitsPostRequest) (*ComponentResourceLimit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ComponentResourceLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentResourceLimitApiService.ApiComponentResourceLimitsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_resource_limits"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.componentResourceLimit == nil {
		return localVarReturnValue, nil, reportError("componentResourceLimit is required and must be specified")
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
	localVarPostBody = r.componentResourceLimit
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
