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
	"strings"
)

// DataTransferDataTypeApiService DataTransferDataTypeApi service
type DataTransferDataTypeApiService service

type DataTransferDataTypeApiApiDataTransferDataTypesGetCollectionRequest struct {
	ctx          context.Context
	ApiService   *DataTransferDataTypeApiService
	page         *int32
	itemsPerPage *int32
}

// The collection page number
func (r DataTransferDataTypeApiApiDataTransferDataTypesGetCollectionRequest) Page(page int32) DataTransferDataTypeApiApiDataTransferDataTypesGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r DataTransferDataTypeApiApiDataTransferDataTypesGetCollectionRequest) ItemsPerPage(itemsPerPage int32) DataTransferDataTypeApiApiDataTransferDataTypesGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

func (r DataTransferDataTypeApiApiDataTransferDataTypesGetCollectionRequest) Execute() ([]DataTransferDataType, *http.Response, error) {
	return r.ApiService.ApiDataTransferDataTypesGetCollectionExecute(r)
}

/*
ApiDataTransferDataTypesGetCollection Retrieves the collection of DataTransferDataType resources.

Retrieves the collection of DataTransferDataType resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DataTransferDataTypeApiApiDataTransferDataTypesGetCollectionRequest
*/
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesGetCollection(ctx context.Context) DataTransferDataTypeApiApiDataTransferDataTypesGetCollectionRequest {
	return DataTransferDataTypeApiApiDataTransferDataTypesGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []DataTransferDataType
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesGetCollectionExecute(r DataTransferDataTypeApiApiDataTransferDataTypesGetCollectionRequest) ([]DataTransferDataType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []DataTransferDataType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferDataTypeApiService.ApiDataTransferDataTypesGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_data_types"

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

type DataTransferDataTypeApiApiDataTransferDataTypesIdDeleteRequest struct {
	ctx        context.Context
	ApiService *DataTransferDataTypeApiService
	id         string
}

func (r DataTransferDataTypeApiApiDataTransferDataTypesIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiDataTransferDataTypesIdDeleteExecute(r)
}

/*
ApiDataTransferDataTypesIdDelete Removes the DataTransferDataType resource.

Removes the DataTransferDataType resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id DataTransferDataType identifier
	@return DataTransferDataTypeApiApiDataTransferDataTypesIdDeleteRequest
*/
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesIdDelete(ctx context.Context, id string) DataTransferDataTypeApiApiDataTransferDataTypesIdDeleteRequest {
	return DataTransferDataTypeApiApiDataTransferDataTypesIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesIdDeleteExecute(r DataTransferDataTypeApiApiDataTransferDataTypesIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferDataTypeApiService.ApiDataTransferDataTypesIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_data_types/{id}"
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

type DataTransferDataTypeApiApiDataTransferDataTypesIdGetRequest struct {
	ctx        context.Context
	ApiService *DataTransferDataTypeApiService
	id         string
}

func (r DataTransferDataTypeApiApiDataTransferDataTypesIdGetRequest) Execute() (*DataTransferDataType, *http.Response, error) {
	return r.ApiService.ApiDataTransferDataTypesIdGetExecute(r)
}

/*
ApiDataTransferDataTypesIdGet Retrieves a DataTransferDataType resource.

Retrieves a DataTransferDataType resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id DataTransferDataType identifier
	@return DataTransferDataTypeApiApiDataTransferDataTypesIdGetRequest
*/
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesIdGet(ctx context.Context, id string) DataTransferDataTypeApiApiDataTransferDataTypesIdGetRequest {
	return DataTransferDataTypeApiApiDataTransferDataTypesIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DataTransferDataType
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesIdGetExecute(r DataTransferDataTypeApiApiDataTransferDataTypesIdGetRequest) (*DataTransferDataType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataTransferDataType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferDataTypeApiService.ApiDataTransferDataTypesIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_data_types/{id}"
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

type DataTransferDataTypeApiApiDataTransferDataTypesIdPatchRequest struct {
	ctx                  context.Context
	ApiService           *DataTransferDataTypeApiService
	dataTransferDataType *DataTransferDataType
	id                   string
}

// The updated DataTransferDataType resource
func (r DataTransferDataTypeApiApiDataTransferDataTypesIdPatchRequest) DataTransferDataType(dataTransferDataType DataTransferDataType) DataTransferDataTypeApiApiDataTransferDataTypesIdPatchRequest {
	r.dataTransferDataType = &dataTransferDataType
	return r
}

func (r DataTransferDataTypeApiApiDataTransferDataTypesIdPatchRequest) Execute() (*DataTransferDataType, *http.Response, error) {
	return r.ApiService.ApiDataTransferDataTypesIdPatchExecute(r)
}

/*
ApiDataTransferDataTypesIdPatch Updates the DataTransferDataType resource.

Updates the DataTransferDataType resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id DataTransferDataType identifier
	@return DataTransferDataTypeApiApiDataTransferDataTypesIdPatchRequest
*/
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesIdPatch(ctx context.Context, id string) DataTransferDataTypeApiApiDataTransferDataTypesIdPatchRequest {
	return DataTransferDataTypeApiApiDataTransferDataTypesIdPatchRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DataTransferDataType
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesIdPatchExecute(r DataTransferDataTypeApiApiDataTransferDataTypesIdPatchRequest) (*DataTransferDataType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataTransferDataType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferDataTypeApiService.ApiDataTransferDataTypesIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_data_types/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataTransferDataType == nil {
		return localVarReturnValue, nil, reportError("dataTransferDataType is required and must be specified")
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
	localVarPostBody = r.dataTransferDataType
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

type DataTransferDataTypeApiApiDataTransferDataTypesIdPutRequest struct {
	ctx                  context.Context
	ApiService           *DataTransferDataTypeApiService
	dataTransferDataType *DataTransferDataType
	id                   string
}

// The updated DataTransferDataType resource
func (r DataTransferDataTypeApiApiDataTransferDataTypesIdPutRequest) DataTransferDataType(dataTransferDataType DataTransferDataType) DataTransferDataTypeApiApiDataTransferDataTypesIdPutRequest {
	r.dataTransferDataType = &dataTransferDataType
	return r
}

func (r DataTransferDataTypeApiApiDataTransferDataTypesIdPutRequest) Execute() (*DataTransferDataType, *http.Response, error) {
	return r.ApiService.ApiDataTransferDataTypesIdPutExecute(r)
}

/*
ApiDataTransferDataTypesIdPut Replaces the DataTransferDataType resource.

Replaces the DataTransferDataType resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id DataTransferDataType identifier
	@return DataTransferDataTypeApiApiDataTransferDataTypesIdPutRequest
*/
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesIdPut(ctx context.Context, id string) DataTransferDataTypeApiApiDataTransferDataTypesIdPutRequest {
	return DataTransferDataTypeApiApiDataTransferDataTypesIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DataTransferDataType
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesIdPutExecute(r DataTransferDataTypeApiApiDataTransferDataTypesIdPutRequest) (*DataTransferDataType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataTransferDataType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferDataTypeApiService.ApiDataTransferDataTypesIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_data_types/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataTransferDataType == nil {
		return localVarReturnValue, nil, reportError("dataTransferDataType is required and must be specified")
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
	localVarPostBody = r.dataTransferDataType
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

type DataTransferDataTypeApiApiDataTransferDataTypesPostRequest struct {
	ctx                  context.Context
	ApiService           *DataTransferDataTypeApiService
	dataTransferDataType *DataTransferDataType
}

// The new DataTransferDataType resource
func (r DataTransferDataTypeApiApiDataTransferDataTypesPostRequest) DataTransferDataType(dataTransferDataType DataTransferDataType) DataTransferDataTypeApiApiDataTransferDataTypesPostRequest {
	r.dataTransferDataType = &dataTransferDataType
	return r
}

func (r DataTransferDataTypeApiApiDataTransferDataTypesPostRequest) Execute() (*DataTransferDataType, *http.Response, error) {
	return r.ApiService.ApiDataTransferDataTypesPostExecute(r)
}

/*
ApiDataTransferDataTypesPost Creates a DataTransferDataType resource.

Creates a DataTransferDataType resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DataTransferDataTypeApiApiDataTransferDataTypesPostRequest
*/
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesPost(ctx context.Context) DataTransferDataTypeApiApiDataTransferDataTypesPostRequest {
	return DataTransferDataTypeApiApiDataTransferDataTypesPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DataTransferDataType
func (a *DataTransferDataTypeApiService) ApiDataTransferDataTypesPostExecute(r DataTransferDataTypeApiApiDataTransferDataTypesPostRequest) (*DataTransferDataType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataTransferDataType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferDataTypeApiService.ApiDataTransferDataTypesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_data_types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataTransferDataType == nil {
		return localVarReturnValue, nil, reportError("dataTransferDataType is required and must be specified")
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
	localVarPostBody = r.dataTransferDataType
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
