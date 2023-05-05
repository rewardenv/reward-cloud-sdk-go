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

// ExportedDataApiService ExportedDataApi service
type ExportedDataApiService service

type ExportedDataApiApiExportedDatasGetCollectionRequest struct {
	ctx                   context.Context
	ApiService            *ExportedDataApiService
	page                  *int32
	itemsPerPage          *int32
	dataTransferDataType  *string
	dataTransferDataType2 *[]string
	environment           *string
	environment2          *[]string
	orderCreatedAt        *string
	orderUpdatedAt        *string
}

// The collection page number
func (r ExportedDataApiApiExportedDatasGetCollectionRequest) Page(page int32) ExportedDataApiApiExportedDatasGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r ExportedDataApiApiExportedDatasGetCollectionRequest) ItemsPerPage(itemsPerPage int32) ExportedDataApiApiExportedDatasGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

func (r ExportedDataApiApiExportedDatasGetCollectionRequest) DataTransferDataType(dataTransferDataType string) ExportedDataApiApiExportedDatasGetCollectionRequest {
	r.dataTransferDataType = &dataTransferDataType
	return r
}

func (r ExportedDataApiApiExportedDatasGetCollectionRequest) DataTransferDataType2(dataTransferDataType2 []string) ExportedDataApiApiExportedDatasGetCollectionRequest {
	r.dataTransferDataType2 = &dataTransferDataType2
	return r
}

func (r ExportedDataApiApiExportedDatasGetCollectionRequest) Environment(environment string) ExportedDataApiApiExportedDatasGetCollectionRequest {
	r.environment = &environment
	return r
}

func (r ExportedDataApiApiExportedDatasGetCollectionRequest) Environment2(environment2 []string) ExportedDataApiApiExportedDatasGetCollectionRequest {
	r.environment2 = &environment2
	return r
}

func (r ExportedDataApiApiExportedDatasGetCollectionRequest) OrderCreatedAt(orderCreatedAt string) ExportedDataApiApiExportedDatasGetCollectionRequest {
	r.orderCreatedAt = &orderCreatedAt
	return r
}

func (r ExportedDataApiApiExportedDatasGetCollectionRequest) OrderUpdatedAt(orderUpdatedAt string) ExportedDataApiApiExportedDatasGetCollectionRequest {
	r.orderUpdatedAt = &orderUpdatedAt
	return r
}

func (r ExportedDataApiApiExportedDatasGetCollectionRequest) Execute() ([]ExportedData, *http.Response, error) {
	return r.ApiService.ApiExportedDatasGetCollectionExecute(r)
}

/*
ApiExportedDatasGetCollection Retrieves the collection of ExportedData resources.

Retrieves the collection of ExportedData resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ExportedDataApiApiExportedDatasGetCollectionRequest
*/
func (a *ExportedDataApiService) ApiExportedDatasGetCollection(ctx context.Context) ExportedDataApiApiExportedDatasGetCollectionRequest {
	return ExportedDataApiApiExportedDatasGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ExportedData
func (a *ExportedDataApiService) ApiExportedDatasGetCollectionExecute(r ExportedDataApiApiExportedDatasGetCollectionRequest) ([]ExportedData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ExportedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportedDataApiService.ApiExportedDatasGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/exported_datas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
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

type ExportedDataApiApiExportedDatasIdDeleteRequest struct {
	ctx        context.Context
	ApiService *ExportedDataApiService
	id         string
}

func (r ExportedDataApiApiExportedDatasIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiExportedDatasIdDeleteExecute(r)
}

/*
ApiExportedDatasIdDelete Removes the ExportedData resource.

Removes the ExportedData resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ExportedData identifier
	@return ExportedDataApiApiExportedDatasIdDeleteRequest
*/
func (a *ExportedDataApiService) ApiExportedDatasIdDelete(ctx context.Context, id string) ExportedDataApiApiExportedDatasIdDeleteRequest {
	return ExportedDataApiApiExportedDatasIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *ExportedDataApiService) ApiExportedDatasIdDeleteExecute(r ExportedDataApiApiExportedDatasIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportedDataApiService.ApiExportedDatasIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/exported_datas/{id}"
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

type ExportedDataApiApiExportedDatasIdGetRequest struct {
	ctx        context.Context
	ApiService *ExportedDataApiService
	id         string
}

func (r ExportedDataApiApiExportedDatasIdGetRequest) Execute() (*ExportedData, *http.Response, error) {
	return r.ApiService.ApiExportedDatasIdGetExecute(r)
}

/*
ApiExportedDatasIdGet Retrieves a ExportedData resource.

Retrieves a ExportedData resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ExportedData identifier
	@return ExportedDataApiApiExportedDatasIdGetRequest
*/
func (a *ExportedDataApiService) ApiExportedDatasIdGet(ctx context.Context, id string) ExportedDataApiApiExportedDatasIdGetRequest {
	return ExportedDataApiApiExportedDatasIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ExportedData
func (a *ExportedDataApiService) ApiExportedDatasIdGetExecute(r ExportedDataApiApiExportedDatasIdGetRequest) (*ExportedData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExportedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportedDataApiService.ApiExportedDatasIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/exported_datas/{id}"
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

type ExportedDataApiApiExportedDatasIdPatchRequest struct {
	ctx          context.Context
	ApiService   *ExportedDataApiService
	exportedData *ExportedData
	id           string
}

// The updated ExportedData resource
func (r ExportedDataApiApiExportedDatasIdPatchRequest) ExportedData(exportedData ExportedData) ExportedDataApiApiExportedDatasIdPatchRequest {
	r.exportedData = &exportedData
	return r
}

func (r ExportedDataApiApiExportedDatasIdPatchRequest) Execute() (*ExportedData, *http.Response, error) {
	return r.ApiService.ApiExportedDatasIdPatchExecute(r)
}

/*
ApiExportedDatasIdPatch Updates the ExportedData resource.

Updates the ExportedData resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ExportedData identifier
	@return ExportedDataApiApiExportedDatasIdPatchRequest
*/
func (a *ExportedDataApiService) ApiExportedDatasIdPatch(ctx context.Context, id string) ExportedDataApiApiExportedDatasIdPatchRequest {
	return ExportedDataApiApiExportedDatasIdPatchRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ExportedData
func (a *ExportedDataApiService) ApiExportedDatasIdPatchExecute(r ExportedDataApiApiExportedDatasIdPatchRequest) (*ExportedData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExportedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportedDataApiService.ApiExportedDatasIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/exported_datas/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.exportedData == nil {
		return localVarReturnValue, nil, reportError("exportedData is required and must be specified")
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
	localVarPostBody = r.exportedData
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

type ExportedDataApiApiExportedDatasIdPutRequest struct {
	ctx          context.Context
	ApiService   *ExportedDataApiService
	exportedData *ExportedData
	id           string
}

// The updated ExportedData resource
func (r ExportedDataApiApiExportedDatasIdPutRequest) ExportedData(exportedData ExportedData) ExportedDataApiApiExportedDatasIdPutRequest {
	r.exportedData = &exportedData
	return r
}

func (r ExportedDataApiApiExportedDatasIdPutRequest) Execute() (*ExportedData, *http.Response, error) {
	return r.ApiService.ApiExportedDatasIdPutExecute(r)
}

/*
ApiExportedDatasIdPut Replaces the ExportedData resource.

Replaces the ExportedData resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ExportedData identifier
	@return ExportedDataApiApiExportedDatasIdPutRequest
*/
func (a *ExportedDataApiService) ApiExportedDatasIdPut(ctx context.Context, id string) ExportedDataApiApiExportedDatasIdPutRequest {
	return ExportedDataApiApiExportedDatasIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ExportedData
func (a *ExportedDataApiService) ApiExportedDatasIdPutExecute(r ExportedDataApiApiExportedDatasIdPutRequest) (*ExportedData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExportedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportedDataApiService.ApiExportedDatasIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/exported_datas/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.exportedData == nil {
		return localVarReturnValue, nil, reportError("exportedData is required and must be specified")
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
	localVarPostBody = r.exportedData
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

type ExportedDataApiApiExportedDatasPostRequest struct {
	ctx          context.Context
	ApiService   *ExportedDataApiService
	exportedData *ExportedData
}

// The new ExportedData resource
func (r ExportedDataApiApiExportedDatasPostRequest) ExportedData(exportedData ExportedData) ExportedDataApiApiExportedDatasPostRequest {
	r.exportedData = &exportedData
	return r
}

func (r ExportedDataApiApiExportedDatasPostRequest) Execute() (*ExportedData, *http.Response, error) {
	return r.ApiService.ApiExportedDatasPostExecute(r)
}

/*
ApiExportedDatasPost Creates a ExportedData resource.

Creates a ExportedData resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ExportedDataApiApiExportedDatasPostRequest
*/
func (a *ExportedDataApiService) ApiExportedDatasPost(ctx context.Context) ExportedDataApiApiExportedDatasPostRequest {
	return ExportedDataApiApiExportedDatasPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExportedData
func (a *ExportedDataApiService) ApiExportedDatasPostExecute(r ExportedDataApiApiExportedDatasPostRequest) (*ExportedData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExportedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportedDataApiService.ApiExportedDatasPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/exported_datas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.exportedData == nil {
		return localVarReturnValue, nil, reportError("exportedData is required and must be specified")
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
	localVarPostBody = r.exportedData
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
