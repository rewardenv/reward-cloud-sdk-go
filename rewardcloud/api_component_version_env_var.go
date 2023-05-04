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
	"reflect"
	"strings"
)

// ComponentVersionEnvVarApiService ComponentVersionEnvVarApi service
type ComponentVersionEnvVarApiService service

type ApiApiComponentVersionEnvVarsGetCollectionRequest struct {
	ctx               context.Context
	ApiService        *ComponentVersionEnvVarApiService
	page              *int32
	itemsPerPage      *int32
	componentVersion  *string
	componentVersion2 *[]string
	envVarType        *string
	envVarType2       *[]string
}

// The collection page number
func (r ApiApiComponentVersionEnvVarsGetCollectionRequest) Page(page int32) ApiApiComponentVersionEnvVarsGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r ApiApiComponentVersionEnvVarsGetCollectionRequest) ItemsPerPage(itemsPerPage int32) ApiApiComponentVersionEnvVarsGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

func (r ApiApiComponentVersionEnvVarsGetCollectionRequest) ComponentVersion(componentVersion string) ApiApiComponentVersionEnvVarsGetCollectionRequest {
	r.componentVersion = &componentVersion
	return r
}

func (r ApiApiComponentVersionEnvVarsGetCollectionRequest) ComponentVersion2(componentVersion2 []string) ApiApiComponentVersionEnvVarsGetCollectionRequest {
	r.componentVersion2 = &componentVersion2
	return r
}

func (r ApiApiComponentVersionEnvVarsGetCollectionRequest) EnvVarType(envVarType string) ApiApiComponentVersionEnvVarsGetCollectionRequest {
	r.envVarType = &envVarType
	return r
}

func (r ApiApiComponentVersionEnvVarsGetCollectionRequest) EnvVarType2(envVarType2 []string) ApiApiComponentVersionEnvVarsGetCollectionRequest {
	r.envVarType2 = &envVarType2
	return r
}

func (r ApiApiComponentVersionEnvVarsGetCollectionRequest) Execute() (*ApiComponentVersionEnvVarsGetCollection200Response, *http.Response, error) {
	return r.ApiService.ApiComponentVersionEnvVarsGetCollectionExecute(r)
}

/*
ApiComponentVersionEnvVarsGetCollection Retrieves the collection of ComponentVersionEnvVar resources.

Retrieves the collection of ComponentVersionEnvVar resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiComponentVersionEnvVarsGetCollectionRequest
*/
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsGetCollection(ctx context.Context) ApiApiComponentVersionEnvVarsGetCollectionRequest {
	return ApiApiComponentVersionEnvVarsGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiComponentVersionEnvVarsGetCollection200Response
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsGetCollectionExecute(r ApiApiComponentVersionEnvVarsGetCollectionRequest) (*ApiComponentVersionEnvVarsGetCollection200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiComponentVersionEnvVarsGetCollection200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentVersionEnvVarApiService.ApiComponentVersionEnvVarsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_version_env_vars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.itemsPerPage != nil {
		localVarQueryParams.Add("itemsPerPage", parameterToString(*r.itemsPerPage, ""))
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
	localVarHTTPHeaderAccepts := []string{"application/hal+json", "application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

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

type ApiApiComponentVersionEnvVarsIdDeleteRequest struct {
	ctx        context.Context
	ApiService *ComponentVersionEnvVarApiService
	id         string
}

func (r ApiApiComponentVersionEnvVarsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiComponentVersionEnvVarsIdDeleteExecute(r)
}

/*
ApiComponentVersionEnvVarsIdDelete Removes the ComponentVersionEnvVar resource.

Removes the ComponentVersionEnvVar resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ComponentVersionEnvVar identifier
	@return ApiApiComponentVersionEnvVarsIdDeleteRequest
*/
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsIdDelete(ctx context.Context, id string) ApiApiComponentVersionEnvVarsIdDeleteRequest {
	return ApiApiComponentVersionEnvVarsIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsIdDeleteExecute(r ApiApiComponentVersionEnvVarsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentVersionEnvVarApiService.ApiComponentVersionEnvVarsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_version_env_vars/{id}"
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

type ApiApiComponentVersionEnvVarsIdGetRequest struct {
	ctx        context.Context
	ApiService *ComponentVersionEnvVarApiService
	id         string
}

func (r ApiApiComponentVersionEnvVarsIdGetRequest) Execute() (*ComponentVersionEnvVarJsonhal, *http.Response, error) {
	return r.ApiService.ApiComponentVersionEnvVarsIdGetExecute(r)
}

/*
ApiComponentVersionEnvVarsIdGet Retrieves a ComponentVersionEnvVar resource.

Retrieves a ComponentVersionEnvVar resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ComponentVersionEnvVar identifier
	@return ApiApiComponentVersionEnvVarsIdGetRequest
*/
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsIdGet(ctx context.Context, id string) ApiApiComponentVersionEnvVarsIdGetRequest {
	return ApiApiComponentVersionEnvVarsIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ComponentVersionEnvVarJsonhal
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsIdGetExecute(r ApiApiComponentVersionEnvVarsIdGetRequest) (*ComponentVersionEnvVarJsonhal, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ComponentVersionEnvVarJsonhal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentVersionEnvVarApiService.ApiComponentVersionEnvVarsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_version_env_vars/{id}"
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
	localVarHTTPHeaderAccepts := []string{"application/hal+json", "application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

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

type ApiApiComponentVersionEnvVarsIdPatchRequest struct {
	ctx                    context.Context
	ApiService             *ComponentVersionEnvVarApiService
	id                     string
	componentVersionEnvVar *ComponentVersionEnvVar
}

// The updated ComponentVersionEnvVar resource
func (r ApiApiComponentVersionEnvVarsIdPatchRequest) ComponentVersionEnvVar(componentVersionEnvVar ComponentVersionEnvVar) ApiApiComponentVersionEnvVarsIdPatchRequest {
	r.componentVersionEnvVar = &componentVersionEnvVar
	return r
}

func (r ApiApiComponentVersionEnvVarsIdPatchRequest) Execute() (*ComponentVersionEnvVarJsonhal, *http.Response, error) {
	return r.ApiService.ApiComponentVersionEnvVarsIdPatchExecute(r)
}

/*
ApiComponentVersionEnvVarsIdPatch Updates the ComponentVersionEnvVar resource.

Updates the ComponentVersionEnvVar resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ComponentVersionEnvVar identifier
	@return ApiApiComponentVersionEnvVarsIdPatchRequest
*/
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsIdPatch(ctx context.Context, id string) ApiApiComponentVersionEnvVarsIdPatchRequest {
	return ApiApiComponentVersionEnvVarsIdPatchRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ComponentVersionEnvVarJsonhal
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsIdPatchExecute(r ApiApiComponentVersionEnvVarsIdPatchRequest) (*ComponentVersionEnvVarJsonhal, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ComponentVersionEnvVarJsonhal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentVersionEnvVarApiService.ApiComponentVersionEnvVarsIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_version_env_vars/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.componentVersionEnvVar == nil {
		return localVarReturnValue, nil, reportError("componentVersionEnvVar is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json", "application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.componentVersionEnvVar
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

type ApiApiComponentVersionEnvVarsIdPutRequest struct {
	ctx                           context.Context
	ApiService                    *ComponentVersionEnvVarApiService
	id                            string
	componentVersionEnvVarJsonhal *ComponentVersionEnvVarJsonhal
}

// The updated ComponentVersionEnvVar resource
func (r ApiApiComponentVersionEnvVarsIdPutRequest) ComponentVersionEnvVarJsonhal(componentVersionEnvVarJsonhal ComponentVersionEnvVarJsonhal) ApiApiComponentVersionEnvVarsIdPutRequest {
	r.componentVersionEnvVarJsonhal = &componentVersionEnvVarJsonhal
	return r
}

func (r ApiApiComponentVersionEnvVarsIdPutRequest) Execute() (*ComponentVersionEnvVarJsonhal, *http.Response, error) {
	return r.ApiService.ApiComponentVersionEnvVarsIdPutExecute(r)
}

/*
ApiComponentVersionEnvVarsIdPut Replaces the ComponentVersionEnvVar resource.

Replaces the ComponentVersionEnvVar resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ComponentVersionEnvVar identifier
	@return ApiApiComponentVersionEnvVarsIdPutRequest
*/
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsIdPut(ctx context.Context, id string) ApiApiComponentVersionEnvVarsIdPutRequest {
	return ApiApiComponentVersionEnvVarsIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ComponentVersionEnvVarJsonhal
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsIdPutExecute(r ApiApiComponentVersionEnvVarsIdPutRequest) (*ComponentVersionEnvVarJsonhal, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ComponentVersionEnvVarJsonhal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentVersionEnvVarApiService.ApiComponentVersionEnvVarsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_version_env_vars/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.componentVersionEnvVarJsonhal == nil {
		return localVarReturnValue, nil, reportError("componentVersionEnvVarJsonhal is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/hal+json", "application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json", "application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.componentVersionEnvVarJsonhal
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

type ApiApiComponentVersionEnvVarsPostRequest struct {
	ctx                           context.Context
	ApiService                    *ComponentVersionEnvVarApiService
	componentVersionEnvVarJsonhal *ComponentVersionEnvVarJsonhal
}

// The new ComponentVersionEnvVar resource
func (r ApiApiComponentVersionEnvVarsPostRequest) ComponentVersionEnvVarJsonhal(componentVersionEnvVarJsonhal ComponentVersionEnvVarJsonhal) ApiApiComponentVersionEnvVarsPostRequest {
	r.componentVersionEnvVarJsonhal = &componentVersionEnvVarJsonhal
	return r
}

func (r ApiApiComponentVersionEnvVarsPostRequest) Execute() (*ComponentVersionEnvVarJsonhal, *http.Response, error) {
	return r.ApiService.ApiComponentVersionEnvVarsPostExecute(r)
}

/*
ApiComponentVersionEnvVarsPost Creates a ComponentVersionEnvVar resource.

Creates a ComponentVersionEnvVar resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiComponentVersionEnvVarsPostRequest
*/
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsPost(ctx context.Context) ApiApiComponentVersionEnvVarsPostRequest {
	return ApiApiComponentVersionEnvVarsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ComponentVersionEnvVarJsonhal
func (a *ComponentVersionEnvVarApiService) ApiComponentVersionEnvVarsPostExecute(r ApiApiComponentVersionEnvVarsPostRequest) (*ComponentVersionEnvVarJsonhal, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ComponentVersionEnvVarJsonhal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentVersionEnvVarApiService.ApiComponentVersionEnvVarsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/component_version_env_vars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.componentVersionEnvVarJsonhal == nil {
		return localVarReturnValue, nil, reportError("componentVersionEnvVarJsonhal is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/hal+json", "application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json", "application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.componentVersionEnvVarJsonhal
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
