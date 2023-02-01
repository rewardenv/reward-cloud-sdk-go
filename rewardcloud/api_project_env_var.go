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


// ProjectEnvVarApiService ProjectEnvVarApi service
type ProjectEnvVarApiService service

type ApiApiProjectEnvVarsGetCollectionRequest struct {
	ctx context.Context
	ApiService *ProjectEnvVarApiService
	page *int32
	itemsPerPage *int32
	project *string
	project2 *[]string
	envVarType *string
	envVarType2 *[]string
}

// The collection page number
func (r ApiApiProjectEnvVarsGetCollectionRequest) Page(page int32) ApiApiProjectEnvVarsGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r ApiApiProjectEnvVarsGetCollectionRequest) ItemsPerPage(itemsPerPage int32) ApiApiProjectEnvVarsGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// 
func (r ApiApiProjectEnvVarsGetCollectionRequest) Project(project string) ApiApiProjectEnvVarsGetCollectionRequest {
	r.project = &project
	return r
}

// 
func (r ApiApiProjectEnvVarsGetCollectionRequest) Project2(project2 []string) ApiApiProjectEnvVarsGetCollectionRequest {
	r.project2 = &project2
	return r
}

// 
func (r ApiApiProjectEnvVarsGetCollectionRequest) EnvVarType(envVarType string) ApiApiProjectEnvVarsGetCollectionRequest {
	r.envVarType = &envVarType
	return r
}

// 
func (r ApiApiProjectEnvVarsGetCollectionRequest) EnvVarType2(envVarType2 []string) ApiApiProjectEnvVarsGetCollectionRequest {
	r.envVarType2 = &envVarType2
	return r
}

func (r ApiApiProjectEnvVarsGetCollectionRequest) Execute() ([]ProjectEnvVar, *http.Response, error) {
	return r.ApiService.ApiProjectEnvVarsGetCollectionExecute(r)
}

/*
ApiProjectEnvVarsGetCollection Retrieves the collection of ProjectEnvVar resources.

Retrieves the collection of ProjectEnvVar resources.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiProjectEnvVarsGetCollectionRequest
*/
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsGetCollection(ctx context.Context) ApiApiProjectEnvVarsGetCollectionRequest {
	return ApiApiProjectEnvVarsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ProjectEnvVar
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsGetCollectionExecute(r ApiApiProjectEnvVarsGetCollectionRequest) ([]ProjectEnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ProjectEnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectEnvVarApiService.ApiProjectEnvVarsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project_env_vars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.itemsPerPage != nil {
		localVarQueryParams.Add("itemsPerPage", parameterToString(*r.itemsPerPage, ""))
	}
	if r.project != nil {
		localVarQueryParams.Add("project", parameterToString(*r.project, ""))
	}
	if r.project2 != nil {
		t := *r.project2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("project[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("project[]", parameterToString(t, "multi"))
		}
	}
	if r.envVarType != nil {
		localVarQueryParams.Add("envVarType", parameterToString(*r.envVarType, ""))
	}
	if r.envVarType2 != nil {
		t := *r.envVarType2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("envVarType[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("envVarType[]", parameterToString(t, "multi"))
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

type ApiApiProjectEnvVarsIdDeleteRequest struct {
	ctx context.Context
	ApiService *ProjectEnvVarApiService
	id string
}

func (r ApiApiProjectEnvVarsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiProjectEnvVarsIdDeleteExecute(r)
}

/*
ApiProjectEnvVarsIdDelete Removes the ProjectEnvVar resource.

Removes the ProjectEnvVar resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ProjectEnvVar identifier
 @return ApiApiProjectEnvVarsIdDeleteRequest
*/
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsIdDelete(ctx context.Context, id string) ApiApiProjectEnvVarsIdDeleteRequest {
	return ApiApiProjectEnvVarsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsIdDeleteExecute(r ApiApiProjectEnvVarsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectEnvVarApiService.ApiProjectEnvVarsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project_env_vars/{id}"
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

type ApiApiProjectEnvVarsIdGetRequest struct {
	ctx context.Context
	ApiService *ProjectEnvVarApiService
	id string
}

func (r ApiApiProjectEnvVarsIdGetRequest) Execute() (*ProjectEnvVar, *http.Response, error) {
	return r.ApiService.ApiProjectEnvVarsIdGetExecute(r)
}

/*
ApiProjectEnvVarsIdGet Retrieves a ProjectEnvVar resource.

Retrieves a ProjectEnvVar resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ProjectEnvVar identifier
 @return ApiApiProjectEnvVarsIdGetRequest
*/
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsIdGet(ctx context.Context, id string) ApiApiProjectEnvVarsIdGetRequest {
	return ApiApiProjectEnvVarsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProjectEnvVar
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsIdGetExecute(r ApiApiProjectEnvVarsIdGetRequest) (*ProjectEnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProjectEnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectEnvVarApiService.ApiProjectEnvVarsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project_env_vars/{id}"
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

type ApiApiProjectEnvVarsIdPatchRequest struct {
	ctx context.Context
	ApiService *ProjectEnvVarApiService
	id string
	projectEnvVar *ProjectEnvVar
}

// The updated ProjectEnvVar resource
func (r ApiApiProjectEnvVarsIdPatchRequest) ProjectEnvVar(projectEnvVar ProjectEnvVar) ApiApiProjectEnvVarsIdPatchRequest {
	r.projectEnvVar = &projectEnvVar
	return r
}

func (r ApiApiProjectEnvVarsIdPatchRequest) Execute() (*ProjectEnvVar, *http.Response, error) {
	return r.ApiService.ApiProjectEnvVarsIdPatchExecute(r)
}

/*
ApiProjectEnvVarsIdPatch Updates the ProjectEnvVar resource.

Updates the ProjectEnvVar resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ProjectEnvVar identifier
 @return ApiApiProjectEnvVarsIdPatchRequest
*/
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsIdPatch(ctx context.Context, id string) ApiApiProjectEnvVarsIdPatchRequest {
	return ApiApiProjectEnvVarsIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProjectEnvVar
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsIdPatchExecute(r ApiApiProjectEnvVarsIdPatchRequest) (*ProjectEnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProjectEnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectEnvVarApiService.ApiProjectEnvVarsIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project_env_vars/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.projectEnvVar == nil {
		return localVarReturnValue, nil, reportError("projectEnvVar is required and must be specified")
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
	localVarPostBody = r.projectEnvVar
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

type ApiApiProjectEnvVarsIdPutRequest struct {
	ctx context.Context
	ApiService *ProjectEnvVarApiService
	id string
	projectEnvVar *ProjectEnvVar
}

// The updated ProjectEnvVar resource
func (r ApiApiProjectEnvVarsIdPutRequest) ProjectEnvVar(projectEnvVar ProjectEnvVar) ApiApiProjectEnvVarsIdPutRequest {
	r.projectEnvVar = &projectEnvVar
	return r
}

func (r ApiApiProjectEnvVarsIdPutRequest) Execute() (*ProjectEnvVar, *http.Response, error) {
	return r.ApiService.ApiProjectEnvVarsIdPutExecute(r)
}

/*
ApiProjectEnvVarsIdPut Replaces the ProjectEnvVar resource.

Replaces the ProjectEnvVar resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ProjectEnvVar identifier
 @return ApiApiProjectEnvVarsIdPutRequest
*/
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsIdPut(ctx context.Context, id string) ApiApiProjectEnvVarsIdPutRequest {
	return ApiApiProjectEnvVarsIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProjectEnvVar
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsIdPutExecute(r ApiApiProjectEnvVarsIdPutRequest) (*ProjectEnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProjectEnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectEnvVarApiService.ApiProjectEnvVarsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project_env_vars/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.projectEnvVar == nil {
		return localVarReturnValue, nil, reportError("projectEnvVar is required and must be specified")
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
	localVarPostBody = r.projectEnvVar
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

type ApiApiProjectEnvVarsPostRequest struct {
	ctx context.Context
	ApiService *ProjectEnvVarApiService
	projectEnvVar *ProjectEnvVar
}

// The new ProjectEnvVar resource
func (r ApiApiProjectEnvVarsPostRequest) ProjectEnvVar(projectEnvVar ProjectEnvVar) ApiApiProjectEnvVarsPostRequest {
	r.projectEnvVar = &projectEnvVar
	return r
}

func (r ApiApiProjectEnvVarsPostRequest) Execute() (*ProjectEnvVar, *http.Response, error) {
	return r.ApiService.ApiProjectEnvVarsPostExecute(r)
}

/*
ApiProjectEnvVarsPost Creates a ProjectEnvVar resource.

Creates a ProjectEnvVar resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiProjectEnvVarsPostRequest
*/
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsPost(ctx context.Context) ApiApiProjectEnvVarsPostRequest {
	return ApiApiProjectEnvVarsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProjectEnvVar
func (a *ProjectEnvVarApiService) ApiProjectEnvVarsPostExecute(r ApiApiProjectEnvVarsPostRequest) (*ProjectEnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProjectEnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectEnvVarApiService.ApiProjectEnvVarsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/project_env_vars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.projectEnvVar == nil {
		return localVarReturnValue, nil, reportError("projectEnvVar is required and must be specified")
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
	localVarPostBody = r.projectEnvVar
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