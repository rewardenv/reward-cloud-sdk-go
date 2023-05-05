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

// ImportedDataApiService ImportedDataApi service
type ImportedDataApiService service

type ImportedDataApiApiImportedDatasGetCollectionRequest struct {
	ctx                   context.Context
	ApiService            *ImportedDataApiService
	page                  *int32
	itemsPerPage          *int32
	environment           *string
	environment2          *[]string
	dataTransferDataType  *string
	dataTransferDataType2 *[]string
	state                 *string
	state2                *[]string
	orderCreatedAt        *string
	orderUpdatedAt        *string
}

// The collection page number
func (r ImportedDataApiApiImportedDatasGetCollectionRequest) Page(page int32) ImportedDataApiApiImportedDatasGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r ImportedDataApiApiImportedDatasGetCollectionRequest) ItemsPerPage(itemsPerPage int32) ImportedDataApiApiImportedDatasGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

func (r ImportedDataApiApiImportedDatasGetCollectionRequest) Environment(environment string) ImportedDataApiApiImportedDatasGetCollectionRequest {
	r.environment = &environment
	return r
}

func (r ImportedDataApiApiImportedDatasGetCollectionRequest) Environment2(environment2 []string) ImportedDataApiApiImportedDatasGetCollectionRequest {
	r.environment2 = &environment2
	return r
}

func (r ImportedDataApiApiImportedDatasGetCollectionRequest) DataTransferDataType(dataTransferDataType string) ImportedDataApiApiImportedDatasGetCollectionRequest {
	r.dataTransferDataType = &dataTransferDataType
	return r
}

func (r ImportedDataApiApiImportedDatasGetCollectionRequest) DataTransferDataType2(dataTransferDataType2 []string) ImportedDataApiApiImportedDatasGetCollectionRequest {
	r.dataTransferDataType2 = &dataTransferDataType2
	return r
}

func (r ImportedDataApiApiImportedDatasGetCollectionRequest) State(state string) ImportedDataApiApiImportedDatasGetCollectionRequest {
	r.state = &state
	return r
}

func (r ImportedDataApiApiImportedDatasGetCollectionRequest) State2(state2 []string) ImportedDataApiApiImportedDatasGetCollectionRequest {
	r.state2 = &state2
	return r
}

func (r ImportedDataApiApiImportedDatasGetCollectionRequest) OrderCreatedAt(orderCreatedAt string) ImportedDataApiApiImportedDatasGetCollectionRequest {
	r.orderCreatedAt = &orderCreatedAt
	return r
}

func (r ImportedDataApiApiImportedDatasGetCollectionRequest) OrderUpdatedAt(orderUpdatedAt string) ImportedDataApiApiImportedDatasGetCollectionRequest {
	r.orderUpdatedAt = &orderUpdatedAt
	return r
}

func (r ImportedDataApiApiImportedDatasGetCollectionRequest) Execute() ([]ImportedData, *http.Response, error) {
	return r.ApiService.ApiImportedDatasGetCollectionExecute(r)
}

/*
ApiImportedDatasGetCollection Retrieves the collection of ImportedData resources.

Retrieves the collection of ImportedData resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ImportedDataApiApiImportedDatasGetCollectionRequest
*/
func (a *ImportedDataApiService) ApiImportedDatasGetCollection(ctx context.Context) ImportedDataApiApiImportedDatasGetCollectionRequest {
	return ImportedDataApiApiImportedDatasGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ImportedData
func (a *ImportedDataApiService) ApiImportedDatasGetCollectionExecute(r ImportedDataApiApiImportedDatasGetCollectionRequest) ([]ImportedData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ImportedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportedDataApiService.ApiImportedDatasGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/imported_datas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.environment != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "environment", r.environment, "")
	}
	if r.environment2 != nil {
		t := *r.environment2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "environment[]", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "environment[]", t, "multi")
		}
	}
	if r.dataTransferDataType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dataTransferDataType", r.dataTransferDataType, "")
	}
	if r.dataTransferDataType2 != nil {
		t := *r.dataTransferDataType2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "dataTransferDataType[]", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "dataTransferDataType[]", t, "multi")
		}
	}
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "")
	}
	if r.state2 != nil {
		t := *r.state2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "state[]", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "state[]", t, "multi")
		}
	}
	if r.orderCreatedAt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order[createdAt]", r.orderCreatedAt, "")
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ImportedDataApiApiImportedDatasIdDeleteRequest struct {
	ctx        context.Context
	ApiService *ImportedDataApiService
	id         string
}

func (r ImportedDataApiApiImportedDatasIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiImportedDatasIdDeleteExecute(r)
}

/*
ApiImportedDatasIdDelete Removes the ImportedData resource.

Removes the ImportedData resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ImportedData identifier
	@return ImportedDataApiApiImportedDatasIdDeleteRequest
*/
func (a *ImportedDataApiService) ApiImportedDatasIdDelete(ctx context.Context, id string) ImportedDataApiApiImportedDatasIdDeleteRequest {
	return ImportedDataApiApiImportedDatasIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *ImportedDataApiService) ApiImportedDatasIdDeleteExecute(r ImportedDataApiApiImportedDatasIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportedDataApiService.ApiImportedDatasIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/imported_datas/{id}"
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

type ImportedDataApiApiImportedDatasIdGetRequest struct {
	ctx        context.Context
	ApiService *ImportedDataApiService
	id         string
}

func (r ImportedDataApiApiImportedDatasIdGetRequest) Execute() (*ImportedData, *http.Response, error) {
	return r.ApiService.ApiImportedDatasIdGetExecute(r)
}

/*
ApiImportedDatasIdGet Retrieves a ImportedData resource.

Retrieves a ImportedData resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ImportedData identifier
	@return ImportedDataApiApiImportedDatasIdGetRequest
*/
func (a *ImportedDataApiService) ApiImportedDatasIdGet(ctx context.Context, id string) ImportedDataApiApiImportedDatasIdGetRequest {
	return ImportedDataApiApiImportedDatasIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ImportedData
func (a *ImportedDataApiService) ApiImportedDatasIdGetExecute(r ImportedDataApiApiImportedDatasIdGetRequest) (*ImportedData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ImportedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportedDataApiService.ApiImportedDatasIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/imported_datas/{id}"
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ImportedDataApiApiImportedDatasIdPatchRequest struct {
	ctx          context.Context
	ApiService   *ImportedDataApiService
	importedData *ImportedData
	id           string
}

// The updated ImportedData resource
func (r ImportedDataApiApiImportedDatasIdPatchRequest) ImportedData(importedData ImportedData) ImportedDataApiApiImportedDatasIdPatchRequest {
	r.importedData = &importedData
	return r
}

func (r ImportedDataApiApiImportedDatasIdPatchRequest) Execute() (*ImportedData, *http.Response, error) {
	return r.ApiService.ApiImportedDatasIdPatchExecute(r)
}

/*
ApiImportedDatasIdPatch Updates the ImportedData resource.

Updates the ImportedData resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ImportedData identifier
	@return ImportedDataApiApiImportedDatasIdPatchRequest
*/
func (a *ImportedDataApiService) ApiImportedDatasIdPatch(ctx context.Context, id string) ImportedDataApiApiImportedDatasIdPatchRequest {
	return ImportedDataApiApiImportedDatasIdPatchRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ImportedData
func (a *ImportedDataApiService) ApiImportedDatasIdPatchExecute(r ImportedDataApiApiImportedDatasIdPatchRequest) (*ImportedData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ImportedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportedDataApiService.ApiImportedDatasIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/imported_datas/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.importedData == nil {
		return localVarReturnValue, nil, reportError("importedData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/merge-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.importedData
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

type ImportedDataApiApiImportedDatasIdPutRequest struct {
	ctx          context.Context
	ApiService   *ImportedDataApiService
	importedData *ImportedData
	id           string
}

// The updated ImportedData resource
func (r ImportedDataApiApiImportedDatasIdPutRequest) ImportedData(importedData ImportedData) ImportedDataApiApiImportedDatasIdPutRequest {
	r.importedData = &importedData
	return r
}

func (r ImportedDataApiApiImportedDatasIdPutRequest) Execute() (*ImportedData, *http.Response, error) {
	return r.ApiService.ApiImportedDatasIdPutExecute(r)
}

/*
ApiImportedDatasIdPut Replaces the ImportedData resource.

Replaces the ImportedData resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ImportedData identifier
	@return ImportedDataApiApiImportedDatasIdPutRequest
*/
func (a *ImportedDataApiService) ApiImportedDatasIdPut(ctx context.Context, id string) ImportedDataApiApiImportedDatasIdPutRequest {
	return ImportedDataApiApiImportedDatasIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ImportedData
func (a *ImportedDataApiService) ApiImportedDatasIdPutExecute(r ImportedDataApiApiImportedDatasIdPutRequest) (*ImportedData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ImportedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportedDataApiService.ApiImportedDatasIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/imported_datas/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.importedData == nil {
		return localVarReturnValue, nil, reportError("importedData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.importedData
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

type ImportedDataApiApiImportedDatasPostRequest struct {
	ctx          context.Context
	ApiService   *ImportedDataApiService
	importedData *ImportedData
}

// The new ImportedData resource
func (r ImportedDataApiApiImportedDatasPostRequest) ImportedData(importedData ImportedData) ImportedDataApiApiImportedDatasPostRequest {
	r.importedData = &importedData
	return r
}

func (r ImportedDataApiApiImportedDatasPostRequest) Execute() (*ImportedData, *http.Response, error) {
	return r.ApiService.ApiImportedDatasPostExecute(r)
}

/*
ApiImportedDatasPost Creates a ImportedData resource.

Creates a ImportedData resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ImportedDataApiApiImportedDatasPostRequest
*/
func (a *ImportedDataApiService) ApiImportedDatasPost(ctx context.Context) ImportedDataApiApiImportedDatasPostRequest {
	return ImportedDataApiApiImportedDatasPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ImportedData
func (a *ImportedDataApiService) ApiImportedDatasPostExecute(r ImportedDataApiApiImportedDatasPostRequest) (*ImportedData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ImportedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportedDataApiService.ApiImportedDatasPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/imported_datas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.importedData == nil {
		return localVarReturnValue, nil, reportError("importedData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.importedData
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
