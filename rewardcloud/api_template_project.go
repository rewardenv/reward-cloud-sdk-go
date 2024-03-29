/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.7.0-alpha
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

// TemplateProjectApiService TemplateProjectApi service
type TemplateProjectApiService service

type TemplateProjectApiApiTemplateProjectsGetCollectionRequest struct {
	ctx                 context.Context
	ApiService          *TemplateProjectApiService
	page                *int32
	itemsPerPage        *int32
	projectTypeVersion  *string
	projectTypeVersion2 *[]string
	orderUpdatedAt      *string
}

// The collection page number
func (r TemplateProjectApiApiTemplateProjectsGetCollectionRequest) Page(page int32) TemplateProjectApiApiTemplateProjectsGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r TemplateProjectApiApiTemplateProjectsGetCollectionRequest) ItemsPerPage(itemsPerPage int32) TemplateProjectApiApiTemplateProjectsGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

func (r TemplateProjectApiApiTemplateProjectsGetCollectionRequest) ProjectTypeVersion(projectTypeVersion string) TemplateProjectApiApiTemplateProjectsGetCollectionRequest {
	r.projectTypeVersion = &projectTypeVersion
	return r
}

func (r TemplateProjectApiApiTemplateProjectsGetCollectionRequest) ProjectTypeVersion2(projectTypeVersion2 []string) TemplateProjectApiApiTemplateProjectsGetCollectionRequest {
	r.projectTypeVersion2 = &projectTypeVersion2
	return r
}

func (r TemplateProjectApiApiTemplateProjectsGetCollectionRequest) OrderUpdatedAt(orderUpdatedAt string) TemplateProjectApiApiTemplateProjectsGetCollectionRequest {
	r.orderUpdatedAt = &orderUpdatedAt
	return r
}

func (r TemplateProjectApiApiTemplateProjectsGetCollectionRequest) Execute() ([]TemplateProjectTemplateProjectOutput, *http.Response, error) {
	return r.ApiService.ApiTemplateProjectsGetCollectionExecute(r)
}

/*
ApiTemplateProjectsGetCollection Retrieves the collection of TemplateProject resources.

Retrieves the collection of TemplateProject resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TemplateProjectApiApiTemplateProjectsGetCollectionRequest
*/
func (a *TemplateProjectApiService) ApiTemplateProjectsGetCollection(ctx context.Context) TemplateProjectApiApiTemplateProjectsGetCollectionRequest {
	return TemplateProjectApiApiTemplateProjectsGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []TemplateProjectTemplateProjectOutput
func (a *TemplateProjectApiService) ApiTemplateProjectsGetCollectionExecute(r TemplateProjectApiApiTemplateProjectsGetCollectionRequest) ([]TemplateProjectTemplateProjectOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []TemplateProjectTemplateProjectOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateProjectApiService.ApiTemplateProjectsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_projects"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.projectTypeVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "projectTypeVersion", r.projectTypeVersion, "")
	}
	if r.projectTypeVersion2 != nil {
		t := *r.projectTypeVersion2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "projectTypeVersion[]", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "projectTypeVersion[]", t, "multi")
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
	localVarHTTPHeaderAccepts := []string{"application/json", "text/html"}

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

type TemplateProjectApiApiTemplateProjectsIdDeleteRequest struct {
	ctx        context.Context
	ApiService *TemplateProjectApiService
	id         string
}

func (r TemplateProjectApiApiTemplateProjectsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiTemplateProjectsIdDeleteExecute(r)
}

/*
ApiTemplateProjectsIdDelete Removes the TemplateProject resource.

Removes the TemplateProject resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id TemplateProject identifier
	@return TemplateProjectApiApiTemplateProjectsIdDeleteRequest
*/
func (a *TemplateProjectApiService) ApiTemplateProjectsIdDelete(ctx context.Context, id string) TemplateProjectApiApiTemplateProjectsIdDeleteRequest {
	return TemplateProjectApiApiTemplateProjectsIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *TemplateProjectApiService) ApiTemplateProjectsIdDeleteExecute(r TemplateProjectApiApiTemplateProjectsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateProjectApiService.ApiTemplateProjectsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_projects/{id}"
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

type TemplateProjectApiApiTemplateProjectsIdGetRequest struct {
	ctx        context.Context
	ApiService *TemplateProjectApiService
	id         string
}

func (r TemplateProjectApiApiTemplateProjectsIdGetRequest) Execute() (*TemplateProjectTemplateProjectOutput, *http.Response, error) {
	return r.ApiService.ApiTemplateProjectsIdGetExecute(r)
}

/*
ApiTemplateProjectsIdGet Retrieves a TemplateProject resource.

Retrieves a TemplateProject resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id TemplateProject identifier
	@return TemplateProjectApiApiTemplateProjectsIdGetRequest
*/
func (a *TemplateProjectApiService) ApiTemplateProjectsIdGet(ctx context.Context, id string) TemplateProjectApiApiTemplateProjectsIdGetRequest {
	return TemplateProjectApiApiTemplateProjectsIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return TemplateProjectTemplateProjectOutput
func (a *TemplateProjectApiService) ApiTemplateProjectsIdGetExecute(r TemplateProjectApiApiTemplateProjectsIdGetRequest) (*TemplateProjectTemplateProjectOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TemplateProjectTemplateProjectOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateProjectApiService.ApiTemplateProjectsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_projects/{id}"
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
	localVarHTTPHeaderAccepts := []string{"application/json", "text/html"}

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

type TemplateProjectApiApiTemplateProjectsIdPatchRequest struct {
	ctx                                 context.Context
	ApiService                          *TemplateProjectApiService
	templateProjectTemplateProjectInput *TemplateProjectTemplateProjectInput
	id                                  string
}

// The updated TemplateProject resource
func (r TemplateProjectApiApiTemplateProjectsIdPatchRequest) TemplateProjectTemplateProjectInput(templateProjectTemplateProjectInput TemplateProjectTemplateProjectInput) TemplateProjectApiApiTemplateProjectsIdPatchRequest {
	r.templateProjectTemplateProjectInput = &templateProjectTemplateProjectInput
	return r
}

func (r TemplateProjectApiApiTemplateProjectsIdPatchRequest) Execute() (*TemplateProjectTemplateProjectOutput, *http.Response, error) {
	return r.ApiService.ApiTemplateProjectsIdPatchExecute(r)
}

/*
ApiTemplateProjectsIdPatch Updates the TemplateProject resource.

Updates the TemplateProject resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id TemplateProject identifier
	@return TemplateProjectApiApiTemplateProjectsIdPatchRequest
*/
func (a *TemplateProjectApiService) ApiTemplateProjectsIdPatch(ctx context.Context, id string) TemplateProjectApiApiTemplateProjectsIdPatchRequest {
	return TemplateProjectApiApiTemplateProjectsIdPatchRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return TemplateProjectTemplateProjectOutput
func (a *TemplateProjectApiService) ApiTemplateProjectsIdPatchExecute(r TemplateProjectApiApiTemplateProjectsIdPatchRequest) (*TemplateProjectTemplateProjectOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TemplateProjectTemplateProjectOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateProjectApiService.ApiTemplateProjectsIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_projects/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.templateProjectTemplateProjectInput == nil {
		return localVarReturnValue, nil, reportError("templateProjectTemplateProjectInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/merge-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.templateProjectTemplateProjectInput
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

type TemplateProjectApiApiTemplateProjectsIdPutRequest struct {
	ctx                                 context.Context
	ApiService                          *TemplateProjectApiService
	templateProjectTemplateProjectInput *TemplateProjectTemplateProjectInput
	id                                  string
}

// The updated TemplateProject resource
func (r TemplateProjectApiApiTemplateProjectsIdPutRequest) TemplateProjectTemplateProjectInput(templateProjectTemplateProjectInput TemplateProjectTemplateProjectInput) TemplateProjectApiApiTemplateProjectsIdPutRequest {
	r.templateProjectTemplateProjectInput = &templateProjectTemplateProjectInput
	return r
}

func (r TemplateProjectApiApiTemplateProjectsIdPutRequest) Execute() (*TemplateProjectTemplateProjectOutput, *http.Response, error) {
	return r.ApiService.ApiTemplateProjectsIdPutExecute(r)
}

/*
ApiTemplateProjectsIdPut Replaces the TemplateProject resource.

Replaces the TemplateProject resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id TemplateProject identifier
	@return TemplateProjectApiApiTemplateProjectsIdPutRequest
*/
func (a *TemplateProjectApiService) ApiTemplateProjectsIdPut(ctx context.Context, id string) TemplateProjectApiApiTemplateProjectsIdPutRequest {
	return TemplateProjectApiApiTemplateProjectsIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return TemplateProjectTemplateProjectOutput
func (a *TemplateProjectApiService) ApiTemplateProjectsIdPutExecute(r TemplateProjectApiApiTemplateProjectsIdPutRequest) (*TemplateProjectTemplateProjectOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TemplateProjectTemplateProjectOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateProjectApiService.ApiTemplateProjectsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_projects/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.templateProjectTemplateProjectInput == nil {
		return localVarReturnValue, nil, reportError("templateProjectTemplateProjectInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/html"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.templateProjectTemplateProjectInput
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

type TemplateProjectApiApiTemplateProjectsPostRequest struct {
	ctx                                 context.Context
	ApiService                          *TemplateProjectApiService
	templateProjectTemplateProjectInput *TemplateProjectTemplateProjectInput
}

// The new TemplateProject resource
func (r TemplateProjectApiApiTemplateProjectsPostRequest) TemplateProjectTemplateProjectInput(templateProjectTemplateProjectInput TemplateProjectTemplateProjectInput) TemplateProjectApiApiTemplateProjectsPostRequest {
	r.templateProjectTemplateProjectInput = &templateProjectTemplateProjectInput
	return r
}

func (r TemplateProjectApiApiTemplateProjectsPostRequest) Execute() (*TemplateProjectTemplateProjectOutput, *http.Response, error) {
	return r.ApiService.ApiTemplateProjectsPostExecute(r)
}

/*
ApiTemplateProjectsPost Creates a TemplateProject resource.

Creates a TemplateProject resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TemplateProjectApiApiTemplateProjectsPostRequest
*/
func (a *TemplateProjectApiService) ApiTemplateProjectsPost(ctx context.Context) TemplateProjectApiApiTemplateProjectsPostRequest {
	return TemplateProjectApiApiTemplateProjectsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TemplateProjectTemplateProjectOutput
func (a *TemplateProjectApiService) ApiTemplateProjectsPostExecute(r TemplateProjectApiApiTemplateProjectsPostRequest) (*TemplateProjectTemplateProjectOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TemplateProjectTemplateProjectOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateProjectApiService.ApiTemplateProjectsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/template_projects"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.templateProjectTemplateProjectInput == nil {
		return localVarReturnValue, nil, reportError("templateProjectTemplateProjectInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/html"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.templateProjectTemplateProjectInput
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
