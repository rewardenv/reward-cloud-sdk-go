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
	"io"
	"net/http"
	"net/url"
	"strings"
)

// EnvironmentAccessRabbitApiService EnvironmentAccessRabbitApi service
type EnvironmentAccessRabbitApiService service

type EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsGetCollectionRequest struct {
	ctx          context.Context
	ApiService   *EnvironmentAccessRabbitApiService
	page         *int32
	itemsPerPage *int32
}

// The collection page number
func (r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsGetCollectionRequest) Page(page int32) EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsGetCollectionRequest) ItemsPerPage(itemsPerPage int32) EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

func (r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsGetCollectionRequest) Execute() ([]EnvironmentAccessRabbit, *http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessRabbitsGetCollectionExecute(r)
}

/*
ApiEnvironmentAccessRabbitsGetCollection Retrieves the collection of EnvironmentAccessRabbit resources.

Retrieves the collection of EnvironmentAccessRabbit resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsGetCollectionRequest
*/
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsGetCollection(ctx context.Context) EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsGetCollectionRequest {
	return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []EnvironmentAccessRabbit
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsGetCollectionExecute(r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsGetCollectionRequest) ([]EnvironmentAccessRabbit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []EnvironmentAccessRabbit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessRabbitApiService.ApiEnvironmentAccessRabbitsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_rabbits"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdDeleteRequest struct {
	ctx        context.Context
	ApiService *EnvironmentAccessRabbitApiService
	id         string
}

func (r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessRabbitsIdDeleteExecute(r)
}

/*
ApiEnvironmentAccessRabbitsIdDelete Removes the EnvironmentAccessRabbit resource.

Removes the EnvironmentAccessRabbit resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id EnvironmentAccessRabbit identifier
	@return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdDeleteRequest
*/
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsIdDelete(ctx context.Context, id string) EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdDeleteRequest {
	return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsIdDeleteExecute(r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessRabbitApiService.ApiEnvironmentAccessRabbitsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_rabbits/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdGetRequest struct {
	ctx        context.Context
	ApiService *EnvironmentAccessRabbitApiService
	id         string
}

func (r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdGetRequest) Execute() (*EnvironmentAccessRabbit, *http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessRabbitsIdGetExecute(r)
}

/*
ApiEnvironmentAccessRabbitsIdGet Retrieves a EnvironmentAccessRabbit resource.

Retrieves a EnvironmentAccessRabbit resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id EnvironmentAccessRabbit identifier
	@return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdGetRequest
*/
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsIdGet(ctx context.Context, id string) EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdGetRequest {
	return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return EnvironmentAccessRabbit
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsIdGetExecute(r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdGetRequest) (*EnvironmentAccessRabbit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentAccessRabbit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessRabbitApiService.ApiEnvironmentAccessRabbitsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_rabbits/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPatchRequest struct {
	ctx                     context.Context
	ApiService              *EnvironmentAccessRabbitApiService
	environmentAccessRabbit *EnvironmentAccessRabbit
	id                      string
}

// The updated EnvironmentAccessRabbit resource
func (r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPatchRequest) EnvironmentAccessRabbit(environmentAccessRabbit EnvironmentAccessRabbit) EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPatchRequest {
	r.environmentAccessRabbit = &environmentAccessRabbit
	return r
}

func (r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPatchRequest) Execute() (*EnvironmentAccessRabbit, *http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessRabbitsIdPatchExecute(r)
}

/*
ApiEnvironmentAccessRabbitsIdPatch Updates the EnvironmentAccessRabbit resource.

Updates the EnvironmentAccessRabbit resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id EnvironmentAccessRabbit identifier
	@return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPatchRequest
*/
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsIdPatch(ctx context.Context, id string) EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPatchRequest {
	return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPatchRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return EnvironmentAccessRabbit
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsIdPatchExecute(r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPatchRequest) (*EnvironmentAccessRabbit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentAccessRabbit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessRabbitApiService.ApiEnvironmentAccessRabbitsIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_rabbits/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.environmentAccessRabbit == nil {
		return localVarReturnValue, nil, reportError("environmentAccessRabbit is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/merge-patch+json", "application/vnd.api+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.environmentAccessRabbit
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPutRequest struct {
	ctx                     context.Context
	ApiService              *EnvironmentAccessRabbitApiService
	environmentAccessRabbit *EnvironmentAccessRabbit
	id                      string
}

// The updated EnvironmentAccessRabbit resource
func (r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPutRequest) EnvironmentAccessRabbit(environmentAccessRabbit EnvironmentAccessRabbit) EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPutRequest {
	r.environmentAccessRabbit = &environmentAccessRabbit
	return r
}

func (r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPutRequest) Execute() (*EnvironmentAccessRabbit, *http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessRabbitsIdPutExecute(r)
}

/*
ApiEnvironmentAccessRabbitsIdPut Replaces the EnvironmentAccessRabbit resource.

Replaces the EnvironmentAccessRabbit resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id EnvironmentAccessRabbit identifier
	@return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPutRequest
*/
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsIdPut(ctx context.Context, id string) EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPutRequest {
	return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return EnvironmentAccessRabbit
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsIdPutExecute(r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsIdPutRequest) (*EnvironmentAccessRabbit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentAccessRabbit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessRabbitApiService.ApiEnvironmentAccessRabbitsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_rabbits/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.environmentAccessRabbit == nil {
		return localVarReturnValue, nil, reportError("environmentAccessRabbit is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.environmentAccessRabbit
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsPostRequest struct {
	ctx                     context.Context
	ApiService              *EnvironmentAccessRabbitApiService
	environmentAccessRabbit *EnvironmentAccessRabbit
}

// The new EnvironmentAccessRabbit resource
func (r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsPostRequest) EnvironmentAccessRabbit(environmentAccessRabbit EnvironmentAccessRabbit) EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsPostRequest {
	r.environmentAccessRabbit = &environmentAccessRabbit
	return r
}

func (r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsPostRequest) Execute() (*EnvironmentAccessRabbit, *http.Response, error) {
	return r.ApiService.ApiEnvironmentAccessRabbitsPostExecute(r)
}

/*
ApiEnvironmentAccessRabbitsPost Creates a EnvironmentAccessRabbit resource.

Creates a EnvironmentAccessRabbit resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsPostRequest
*/
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsPost(ctx context.Context) EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsPostRequest {
	return EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EnvironmentAccessRabbit
func (a *EnvironmentAccessRabbitApiService) ApiEnvironmentAccessRabbitsPostExecute(r EnvironmentAccessRabbitApiApiEnvironmentAccessRabbitsPostRequest) (*EnvironmentAccessRabbit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentAccessRabbit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAccessRabbitApiService.ApiEnvironmentAccessRabbitsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/environment_access_rabbits"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.environmentAccessRabbit == nil {
		return localVarReturnValue, nil, reportError("environmentAccessRabbit is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json", "application/json", "application/xml", "text/xml", "application/x-yaml", "text/csv", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.environmentAccessRabbit
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
