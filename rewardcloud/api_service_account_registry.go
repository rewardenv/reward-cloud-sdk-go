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
	"reflect"
	"strings"
)

// ServiceAccountRegistryApiService ServiceAccountRegistryApi service
type ServiceAccountRegistryApiService service

type ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest struct {
	ctx             context.Context
	ApiService      *ServiceAccountRegistryApiService
	page            *int32
	itemsPerPage    *int32
	serviceAccount  *string
	serviceAccount2 *[]string
}

// The collection page number
func (r ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest) Page(page int32) ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest) ItemsPerPage(itemsPerPage int32) ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

func (r ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest) ServiceAccount(serviceAccount string) ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest {
	r.serviceAccount = &serviceAccount
	return r
}

func (r ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest) ServiceAccount2(serviceAccount2 []string) ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest {
	r.serviceAccount2 = &serviceAccount2
	return r
}

func (r ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest) Execute() ([]ServiceAccountRegistry, *http.Response, error) {
	return r.ApiService.ApiServiceAccountRegistriesGetCollectionExecute(r)
}

/*
ApiServiceAccountRegistriesGetCollection Retrieves the collection of ServiceAccountRegistry resources.

Retrieves the collection of ServiceAccountRegistry resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest
*/
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesGetCollection(ctx context.Context) ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest {
	return ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ServiceAccountRegistry
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesGetCollectionExecute(r ServiceAccountRegistryApiApiServiceAccountRegistriesGetCollectionRequest) ([]ServiceAccountRegistry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ServiceAccountRegistry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountRegistryApiService.ApiServiceAccountRegistriesGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service_account_registries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.serviceAccount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "serviceAccount", r.serviceAccount, "")
	}
	if r.serviceAccount2 != nil {
		t := *r.serviceAccount2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "serviceAccount[]", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "serviceAccount[]", t, "multi")
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

type ServiceAccountRegistryApiApiServiceAccountRegistriesIdDeleteRequest struct {
	ctx        context.Context
	ApiService *ServiceAccountRegistryApiService
	id         string
}

func (r ServiceAccountRegistryApiApiServiceAccountRegistriesIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiServiceAccountRegistriesIdDeleteExecute(r)
}

/*
ApiServiceAccountRegistriesIdDelete Removes the ServiceAccountRegistry resource.

Removes the ServiceAccountRegistry resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ServiceAccountRegistry identifier
	@return ServiceAccountRegistryApiApiServiceAccountRegistriesIdDeleteRequest
*/
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesIdDelete(ctx context.Context, id string) ServiceAccountRegistryApiApiServiceAccountRegistriesIdDeleteRequest {
	return ServiceAccountRegistryApiApiServiceAccountRegistriesIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesIdDeleteExecute(r ServiceAccountRegistryApiApiServiceAccountRegistriesIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountRegistryApiService.ApiServiceAccountRegistriesIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service_account_registries/{id}"
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

type ServiceAccountRegistryApiApiServiceAccountRegistriesIdGetRequest struct {
	ctx        context.Context
	ApiService *ServiceAccountRegistryApiService
	id         string
}

func (r ServiceAccountRegistryApiApiServiceAccountRegistriesIdGetRequest) Execute() (*ServiceAccountRegistry, *http.Response, error) {
	return r.ApiService.ApiServiceAccountRegistriesIdGetExecute(r)
}

/*
ApiServiceAccountRegistriesIdGet Retrieves a ServiceAccountRegistry resource.

Retrieves a ServiceAccountRegistry resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ServiceAccountRegistry identifier
	@return ServiceAccountRegistryApiApiServiceAccountRegistriesIdGetRequest
*/
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesIdGet(ctx context.Context, id string) ServiceAccountRegistryApiApiServiceAccountRegistriesIdGetRequest {
	return ServiceAccountRegistryApiApiServiceAccountRegistriesIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ServiceAccountRegistry
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesIdGetExecute(r ServiceAccountRegistryApiApiServiceAccountRegistriesIdGetRequest) (*ServiceAccountRegistry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServiceAccountRegistry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountRegistryApiService.ApiServiceAccountRegistriesIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service_account_registries/{id}"
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

type ServiceAccountRegistryApiApiServiceAccountRegistriesIdPatchRequest struct {
	ctx                    context.Context
	ApiService             *ServiceAccountRegistryApiService
	serviceAccountRegistry *ServiceAccountRegistry
	id                     string
}

// The updated ServiceAccountRegistry resource
func (r ServiceAccountRegistryApiApiServiceAccountRegistriesIdPatchRequest) ServiceAccountRegistry(serviceAccountRegistry ServiceAccountRegistry) ServiceAccountRegistryApiApiServiceAccountRegistriesIdPatchRequest {
	r.serviceAccountRegistry = &serviceAccountRegistry
	return r
}

func (r ServiceAccountRegistryApiApiServiceAccountRegistriesIdPatchRequest) Execute() (*ServiceAccountRegistry, *http.Response, error) {
	return r.ApiService.ApiServiceAccountRegistriesIdPatchExecute(r)
}

/*
ApiServiceAccountRegistriesIdPatch Updates the ServiceAccountRegistry resource.

Updates the ServiceAccountRegistry resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ServiceAccountRegistry identifier
	@return ServiceAccountRegistryApiApiServiceAccountRegistriesIdPatchRequest
*/
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesIdPatch(ctx context.Context, id string) ServiceAccountRegistryApiApiServiceAccountRegistriesIdPatchRequest {
	return ServiceAccountRegistryApiApiServiceAccountRegistriesIdPatchRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ServiceAccountRegistry
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesIdPatchExecute(r ServiceAccountRegistryApiApiServiceAccountRegistriesIdPatchRequest) (*ServiceAccountRegistry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServiceAccountRegistry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountRegistryApiService.ApiServiceAccountRegistriesIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service_account_registries/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceAccountRegistry == nil {
		return localVarReturnValue, nil, reportError("serviceAccountRegistry is required and must be specified")
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
	localVarPostBody = r.serviceAccountRegistry
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

type ServiceAccountRegistryApiApiServiceAccountRegistriesIdPutRequest struct {
	ctx                    context.Context
	ApiService             *ServiceAccountRegistryApiService
	serviceAccountRegistry *ServiceAccountRegistry
	id                     string
}

// The updated ServiceAccountRegistry resource
func (r ServiceAccountRegistryApiApiServiceAccountRegistriesIdPutRequest) ServiceAccountRegistry(serviceAccountRegistry ServiceAccountRegistry) ServiceAccountRegistryApiApiServiceAccountRegistriesIdPutRequest {
	r.serviceAccountRegistry = &serviceAccountRegistry
	return r
}

func (r ServiceAccountRegistryApiApiServiceAccountRegistriesIdPutRequest) Execute() (*ServiceAccountRegistry, *http.Response, error) {
	return r.ApiService.ApiServiceAccountRegistriesIdPutExecute(r)
}

/*
ApiServiceAccountRegistriesIdPut Replaces the ServiceAccountRegistry resource.

Replaces the ServiceAccountRegistry resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ServiceAccountRegistry identifier
	@return ServiceAccountRegistryApiApiServiceAccountRegistriesIdPutRequest
*/
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesIdPut(ctx context.Context, id string) ServiceAccountRegistryApiApiServiceAccountRegistriesIdPutRequest {
	return ServiceAccountRegistryApiApiServiceAccountRegistriesIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ServiceAccountRegistry
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesIdPutExecute(r ServiceAccountRegistryApiApiServiceAccountRegistriesIdPutRequest) (*ServiceAccountRegistry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServiceAccountRegistry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountRegistryApiService.ApiServiceAccountRegistriesIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service_account_registries/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceAccountRegistry == nil {
		return localVarReturnValue, nil, reportError("serviceAccountRegistry is required and must be specified")
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
	localVarPostBody = r.serviceAccountRegistry
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

type ServiceAccountRegistryApiApiServiceAccountRegistriesPostRequest struct {
	ctx                    context.Context
	ApiService             *ServiceAccountRegistryApiService
	serviceAccountRegistry *ServiceAccountRegistry
}

// The new ServiceAccountRegistry resource
func (r ServiceAccountRegistryApiApiServiceAccountRegistriesPostRequest) ServiceAccountRegistry(serviceAccountRegistry ServiceAccountRegistry) ServiceAccountRegistryApiApiServiceAccountRegistriesPostRequest {
	r.serviceAccountRegistry = &serviceAccountRegistry
	return r
}

func (r ServiceAccountRegistryApiApiServiceAccountRegistriesPostRequest) Execute() (*ServiceAccountRegistry, *http.Response, error) {
	return r.ApiService.ApiServiceAccountRegistriesPostExecute(r)
}

/*
ApiServiceAccountRegistriesPost Creates a ServiceAccountRegistry resource.

Creates a ServiceAccountRegistry resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ServiceAccountRegistryApiApiServiceAccountRegistriesPostRequest
*/
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesPost(ctx context.Context) ServiceAccountRegistryApiApiServiceAccountRegistriesPostRequest {
	return ServiceAccountRegistryApiApiServiceAccountRegistriesPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ServiceAccountRegistry
func (a *ServiceAccountRegistryApiService) ApiServiceAccountRegistriesPostExecute(r ServiceAccountRegistryApiApiServiceAccountRegistriesPostRequest) (*ServiceAccountRegistry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServiceAccountRegistry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountRegistryApiService.ApiServiceAccountRegistriesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service_account_registries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceAccountRegistry == nil {
		return localVarReturnValue, nil, reportError("serviceAccountRegistry is required and must be specified")
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
	localVarPostBody = r.serviceAccountRegistry
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
