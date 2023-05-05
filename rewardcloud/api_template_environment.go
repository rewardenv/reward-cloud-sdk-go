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

// TemplateEnvironmentApiService TemplateEnvironmentApi service
type TemplateEnvironmentApiService service

type TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest struct {
	ctx              context.Context
	ApiService       *TemplateEnvironmentApiService
	page             *int32
	itemsPerPage     *int32
	id               *int32
	id2              *[]int32
	templateProject  *string
	templateProject2 *[]string
	orderUpdatedAt   *string
}

// The collection page number
func (r TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest) Page(page int32) TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest) ItemsPerPage(itemsPerPage int32) TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

func (r TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest) Id(id int32) TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest {
	r.id = &id
	return r
}

func (r TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest) Id2(id2 []int32) TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest {
	r.id2 = &id2
	return r
}

func (r TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest) TemplateProject(templateProject string) TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest {
	r.templateProject = &templateProject
	return r
}

func (r TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest) TemplateProject2(templateProject2 []string) TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest {
	r.templateProject2 = &templateProject2
	return r
}

func (r TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest) OrderUpdatedAt(orderUpdatedAt string) TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest {
	r.orderUpdatedAt = &orderUpdatedAt
	return r
}

func (r TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest) Execute() ([]TemplateEnvironmentTemplateEnvironmentOutput, *http.Response, error) {
	return r.ApiService.ApiTemplateEnvironmentsGetCollectionExecute(r)
}

/*
ApiTemplateEnvironmentsGetCollection Retrieves the collection of TemplateEnvironment resources.

Retrieves the collection of TemplateEnvironment resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest
*/
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsGetCollection(ctx context.Context) TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest {
	return TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []TemplateEnvironmentTemplateEnvironmentOutput
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsGetCollectionExecute(r TemplateEnvironmentApiApiTemplateEnvironmentsGetCollectionRequest) ([]TemplateEnvironmentTemplateEnvironmentOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []TemplateEnvironmentTemplateEnvironmentOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateEnvironmentApiService.ApiTemplateEnvironmentsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_environments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	}
	if r.id2 != nil {
		t := *r.id2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "id[]", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "id[]", t, "multi")
		}
	}
	if r.templateProject != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "templateProject", r.templateProject, "")
	}
	if r.templateProject2 != nil {
		t := *r.templateProject2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "templateProject[]", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "templateProject[]", t, "multi")
		}
	}
	if r.orderUpdatedAt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order[updatedAt]", r.orderUpdatedAt, "")
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

type TemplateEnvironmentApiApiTemplateEnvironmentsIdDeleteRequest struct {
	ctx        context.Context
	ApiService *TemplateEnvironmentApiService
	id         string
}

func (r TemplateEnvironmentApiApiTemplateEnvironmentsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiTemplateEnvironmentsIdDeleteExecute(r)
}

/*
ApiTemplateEnvironmentsIdDelete Removes the TemplateEnvironment resource.

Removes the TemplateEnvironment resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id TemplateEnvironment identifier
	@return TemplateEnvironmentApiApiTemplateEnvironmentsIdDeleteRequest
*/
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsIdDelete(ctx context.Context, id string) TemplateEnvironmentApiApiTemplateEnvironmentsIdDeleteRequest {
	return TemplateEnvironmentApiApiTemplateEnvironmentsIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsIdDeleteExecute(r TemplateEnvironmentApiApiTemplateEnvironmentsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateEnvironmentApiService.ApiTemplateEnvironmentsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_environments/{id}"
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

type TemplateEnvironmentApiApiTemplateEnvironmentsIdGetRequest struct {
	ctx        context.Context
	ApiService *TemplateEnvironmentApiService
	id         string
}

func (r TemplateEnvironmentApiApiTemplateEnvironmentsIdGetRequest) Execute() (*TemplateEnvironmentTemplateEnvironmentOutput, *http.Response, error) {
	return r.ApiService.ApiTemplateEnvironmentsIdGetExecute(r)
}

/*
ApiTemplateEnvironmentsIdGet Retrieves a TemplateEnvironment resource.

Retrieves a TemplateEnvironment resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id TemplateEnvironment identifier
	@return TemplateEnvironmentApiApiTemplateEnvironmentsIdGetRequest
*/
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsIdGet(ctx context.Context, id string) TemplateEnvironmentApiApiTemplateEnvironmentsIdGetRequest {
	return TemplateEnvironmentApiApiTemplateEnvironmentsIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return TemplateEnvironmentTemplateEnvironmentOutput
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsIdGetExecute(r TemplateEnvironmentApiApiTemplateEnvironmentsIdGetRequest) (*TemplateEnvironmentTemplateEnvironmentOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TemplateEnvironmentTemplateEnvironmentOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateEnvironmentApiService.ApiTemplateEnvironmentsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_environments/{id}"
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

type TemplateEnvironmentApiApiTemplateEnvironmentsIdPatchRequest struct {
	ctx                                         context.Context
	ApiService                                  *TemplateEnvironmentApiService
	templateEnvironmentTemplateEnvironmentInput *TemplateEnvironmentTemplateEnvironmentInput
	id                                          string
}

// The updated TemplateEnvironment resource
func (r TemplateEnvironmentApiApiTemplateEnvironmentsIdPatchRequest) TemplateEnvironmentTemplateEnvironmentInput(templateEnvironmentTemplateEnvironmentInput TemplateEnvironmentTemplateEnvironmentInput) TemplateEnvironmentApiApiTemplateEnvironmentsIdPatchRequest {
	r.templateEnvironmentTemplateEnvironmentInput = &templateEnvironmentTemplateEnvironmentInput
	return r
}

func (r TemplateEnvironmentApiApiTemplateEnvironmentsIdPatchRequest) Execute() (*TemplateEnvironmentTemplateEnvironmentOutput, *http.Response, error) {
	return r.ApiService.ApiTemplateEnvironmentsIdPatchExecute(r)
}

/*
ApiTemplateEnvironmentsIdPatch Updates the TemplateEnvironment resource.

Updates the TemplateEnvironment resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id TemplateEnvironment identifier
	@return TemplateEnvironmentApiApiTemplateEnvironmentsIdPatchRequest
*/
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsIdPatch(ctx context.Context, id string) TemplateEnvironmentApiApiTemplateEnvironmentsIdPatchRequest {
	return TemplateEnvironmentApiApiTemplateEnvironmentsIdPatchRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return TemplateEnvironmentTemplateEnvironmentOutput
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsIdPatchExecute(r TemplateEnvironmentApiApiTemplateEnvironmentsIdPatchRequest) (*TemplateEnvironmentTemplateEnvironmentOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TemplateEnvironmentTemplateEnvironmentOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateEnvironmentApiService.ApiTemplateEnvironmentsIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_environments/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.templateEnvironmentTemplateEnvironmentInput == nil {
		return localVarReturnValue, nil, reportError("templateEnvironmentTemplateEnvironmentInput is required and must be specified")
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
	localVarPostBody = r.templateEnvironmentTemplateEnvironmentInput
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

type TemplateEnvironmentApiApiTemplateEnvironmentsIdPutRequest struct {
	ctx                                         context.Context
	ApiService                                  *TemplateEnvironmentApiService
	templateEnvironmentTemplateEnvironmentInput *TemplateEnvironmentTemplateEnvironmentInput
	id                                          string
}

// The updated TemplateEnvironment resource
func (r TemplateEnvironmentApiApiTemplateEnvironmentsIdPutRequest) TemplateEnvironmentTemplateEnvironmentInput(templateEnvironmentTemplateEnvironmentInput TemplateEnvironmentTemplateEnvironmentInput) TemplateEnvironmentApiApiTemplateEnvironmentsIdPutRequest {
	r.templateEnvironmentTemplateEnvironmentInput = &templateEnvironmentTemplateEnvironmentInput
	return r
}

func (r TemplateEnvironmentApiApiTemplateEnvironmentsIdPutRequest) Execute() (*TemplateEnvironmentTemplateEnvironmentOutput, *http.Response, error) {
	return r.ApiService.ApiTemplateEnvironmentsIdPutExecute(r)
}

/*
ApiTemplateEnvironmentsIdPut Replaces the TemplateEnvironment resource.

Replaces the TemplateEnvironment resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id TemplateEnvironment identifier
	@return TemplateEnvironmentApiApiTemplateEnvironmentsIdPutRequest
*/
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsIdPut(ctx context.Context, id string) TemplateEnvironmentApiApiTemplateEnvironmentsIdPutRequest {
	return TemplateEnvironmentApiApiTemplateEnvironmentsIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return TemplateEnvironmentTemplateEnvironmentOutput
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsIdPutExecute(r TemplateEnvironmentApiApiTemplateEnvironmentsIdPutRequest) (*TemplateEnvironmentTemplateEnvironmentOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TemplateEnvironmentTemplateEnvironmentOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateEnvironmentApiService.ApiTemplateEnvironmentsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_environments/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.templateEnvironmentTemplateEnvironmentInput == nil {
		return localVarReturnValue, nil, reportError("templateEnvironmentTemplateEnvironmentInput is required and must be specified")
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
	localVarPostBody = r.templateEnvironmentTemplateEnvironmentInput
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

type TemplateEnvironmentApiApiTemplateEnvironmentsPostRequest struct {
	ctx                                         context.Context
	ApiService                                  *TemplateEnvironmentApiService
	templateEnvironmentTemplateEnvironmentInput *TemplateEnvironmentTemplateEnvironmentInput
}

// The new TemplateEnvironment resource
func (r TemplateEnvironmentApiApiTemplateEnvironmentsPostRequest) TemplateEnvironmentTemplateEnvironmentInput(templateEnvironmentTemplateEnvironmentInput TemplateEnvironmentTemplateEnvironmentInput) TemplateEnvironmentApiApiTemplateEnvironmentsPostRequest {
	r.templateEnvironmentTemplateEnvironmentInput = &templateEnvironmentTemplateEnvironmentInput
	return r
}

func (r TemplateEnvironmentApiApiTemplateEnvironmentsPostRequest) Execute() (*TemplateEnvironmentTemplateEnvironmentOutput, *http.Response, error) {
	return r.ApiService.ApiTemplateEnvironmentsPostExecute(r)
}

/*
ApiTemplateEnvironmentsPost Creates a TemplateEnvironment resource.

Creates a TemplateEnvironment resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TemplateEnvironmentApiApiTemplateEnvironmentsPostRequest
*/
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsPost(ctx context.Context) TemplateEnvironmentApiApiTemplateEnvironmentsPostRequest {
	return TemplateEnvironmentApiApiTemplateEnvironmentsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TemplateEnvironmentTemplateEnvironmentOutput
func (a *TemplateEnvironmentApiService) ApiTemplateEnvironmentsPostExecute(r TemplateEnvironmentApiApiTemplateEnvironmentsPostRequest) (*TemplateEnvironmentTemplateEnvironmentOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TemplateEnvironmentTemplateEnvironmentOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateEnvironmentApiService.ApiTemplateEnvironmentsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_environments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.templateEnvironmentTemplateEnvironmentInput == nil {
		return localVarReturnValue, nil, reportError("templateEnvironmentTemplateEnvironmentInput is required and must be specified")
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
	localVarPostBody = r.templateEnvironmentTemplateEnvironmentInput
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
