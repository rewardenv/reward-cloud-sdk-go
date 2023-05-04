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
)

// DataTransferTypeApiService DataTransferTypeApi service
type DataTransferTypeApiService service

type ApiApiDataTransferTypesGetCollectionRequest struct {
	ctx          context.Context
	ApiService   *DataTransferTypeApiService
	page         *int32
	itemsPerPage *int32
}

// The collection page number
func (r ApiApiDataTransferTypesGetCollectionRequest) Page(page int32) ApiApiDataTransferTypesGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r ApiApiDataTransferTypesGetCollectionRequest) ItemsPerPage(itemsPerPage int32) ApiApiDataTransferTypesGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

func (r ApiApiDataTransferTypesGetCollectionRequest) Execute() (*ApiDataTransferTypesGetCollection200Response, *http.Response, error) {
	return r.ApiService.ApiDataTransferTypesGetCollectionExecute(r)
}

/*
ApiDataTransferTypesGetCollection Retrieves the collection of DataTransferType resources.

Retrieves the collection of DataTransferType resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiDataTransferTypesGetCollectionRequest
*/
func (a *DataTransferTypeApiService) ApiDataTransferTypesGetCollection(ctx context.Context) ApiApiDataTransferTypesGetCollectionRequest {
	return ApiApiDataTransferTypesGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiDataTransferTypesGetCollection200Response
func (a *DataTransferTypeApiService) ApiDataTransferTypesGetCollectionExecute(r ApiApiDataTransferTypesGetCollectionRequest) (*ApiDataTransferTypesGetCollection200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiDataTransferTypesGetCollection200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferTypeApiService.ApiDataTransferTypesGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.itemsPerPage != nil {
		localVarQueryParams.Add("itemsPerPage", parameterToString(*r.itemsPerPage, ""))
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

type ApiApiDataTransferTypesIdDeleteRequest struct {
	ctx        context.Context
	ApiService *DataTransferTypeApiService
	id         string
}

func (r ApiApiDataTransferTypesIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiDataTransferTypesIdDeleteExecute(r)
}

/*
ApiDataTransferTypesIdDelete Removes the DataTransferType resource.

Removes the DataTransferType resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id DataTransferType identifier
	@return ApiApiDataTransferTypesIdDeleteRequest
*/
func (a *DataTransferTypeApiService) ApiDataTransferTypesIdDelete(ctx context.Context, id string) ApiApiDataTransferTypesIdDeleteRequest {
	return ApiApiDataTransferTypesIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *DataTransferTypeApiService) ApiDataTransferTypesIdDeleteExecute(r ApiApiDataTransferTypesIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferTypeApiService.ApiDataTransferTypesIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_types/{id}"
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

type ApiApiDataTransferTypesIdGetRequest struct {
	ctx        context.Context
	ApiService *DataTransferTypeApiService
	id         string
}

func (r ApiApiDataTransferTypesIdGetRequest) Execute() (*DataTransferTypeJsonhal, *http.Response, error) {
	return r.ApiService.ApiDataTransferTypesIdGetExecute(r)
}

/*
ApiDataTransferTypesIdGet Retrieves a DataTransferType resource.

Retrieves a DataTransferType resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id DataTransferType identifier
	@return ApiApiDataTransferTypesIdGetRequest
*/
func (a *DataTransferTypeApiService) ApiDataTransferTypesIdGet(ctx context.Context, id string) ApiApiDataTransferTypesIdGetRequest {
	return ApiApiDataTransferTypesIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DataTransferTypeJsonhal
func (a *DataTransferTypeApiService) ApiDataTransferTypesIdGetExecute(r ApiApiDataTransferTypesIdGetRequest) (*DataTransferTypeJsonhal, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataTransferTypeJsonhal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferTypeApiService.ApiDataTransferTypesIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_types/{id}"
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

type ApiApiDataTransferTypesIdPatchRequest struct {
	ctx              context.Context
	ApiService       *DataTransferTypeApiService
	id               string
	dataTransferType *DataTransferType
}

// The updated DataTransferType resource
func (r ApiApiDataTransferTypesIdPatchRequest) DataTransferType(dataTransferType DataTransferType) ApiApiDataTransferTypesIdPatchRequest {
	r.dataTransferType = &dataTransferType
	return r
}

func (r ApiApiDataTransferTypesIdPatchRequest) Execute() (*DataTransferTypeJsonhal, *http.Response, error) {
	return r.ApiService.ApiDataTransferTypesIdPatchExecute(r)
}

/*
ApiDataTransferTypesIdPatch Updates the DataTransferType resource.

Updates the DataTransferType resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id DataTransferType identifier
	@return ApiApiDataTransferTypesIdPatchRequest
*/
func (a *DataTransferTypeApiService) ApiDataTransferTypesIdPatch(ctx context.Context, id string) ApiApiDataTransferTypesIdPatchRequest {
	return ApiApiDataTransferTypesIdPatchRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DataTransferTypeJsonhal
func (a *DataTransferTypeApiService) ApiDataTransferTypesIdPatchExecute(r ApiApiDataTransferTypesIdPatchRequest) (*DataTransferTypeJsonhal, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataTransferTypeJsonhal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferTypeApiService.ApiDataTransferTypesIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_types/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataTransferType == nil {
		return localVarReturnValue, nil, reportError("dataTransferType is required and must be specified")
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
	localVarPostBody = r.dataTransferType
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

type ApiApiDataTransferTypesIdPutRequest struct {
	ctx                     context.Context
	ApiService              *DataTransferTypeApiService
	id                      string
	dataTransferTypeJsonhal *DataTransferTypeJsonhal
}

// The updated DataTransferType resource
func (r ApiApiDataTransferTypesIdPutRequest) DataTransferTypeJsonhal(dataTransferTypeJsonhal DataTransferTypeJsonhal) ApiApiDataTransferTypesIdPutRequest {
	r.dataTransferTypeJsonhal = &dataTransferTypeJsonhal
	return r
}

func (r ApiApiDataTransferTypesIdPutRequest) Execute() (*DataTransferTypeJsonhal, *http.Response, error) {
	return r.ApiService.ApiDataTransferTypesIdPutExecute(r)
}

/*
ApiDataTransferTypesIdPut Replaces the DataTransferType resource.

Replaces the DataTransferType resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id DataTransferType identifier
	@return ApiApiDataTransferTypesIdPutRequest
*/
func (a *DataTransferTypeApiService) ApiDataTransferTypesIdPut(ctx context.Context, id string) ApiApiDataTransferTypesIdPutRequest {
	return ApiApiDataTransferTypesIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DataTransferTypeJsonhal
func (a *DataTransferTypeApiService) ApiDataTransferTypesIdPutExecute(r ApiApiDataTransferTypesIdPutRequest) (*DataTransferTypeJsonhal, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataTransferTypeJsonhal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferTypeApiService.ApiDataTransferTypesIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_types/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataTransferTypeJsonhal == nil {
		return localVarReturnValue, nil, reportError("dataTransferTypeJsonhal is required and must be specified")
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
	localVarPostBody = r.dataTransferTypeJsonhal
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

type ApiApiDataTransferTypesPostRequest struct {
	ctx                     context.Context
	ApiService              *DataTransferTypeApiService
	dataTransferTypeJsonhal *DataTransferTypeJsonhal
}

// The new DataTransferType resource
func (r ApiApiDataTransferTypesPostRequest) DataTransferTypeJsonhal(dataTransferTypeJsonhal DataTransferTypeJsonhal) ApiApiDataTransferTypesPostRequest {
	r.dataTransferTypeJsonhal = &dataTransferTypeJsonhal
	return r
}

func (r ApiApiDataTransferTypesPostRequest) Execute() (*DataTransferTypeJsonhal, *http.Response, error) {
	return r.ApiService.ApiDataTransferTypesPostExecute(r)
}

/*
ApiDataTransferTypesPost Creates a DataTransferType resource.

Creates a DataTransferType resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiDataTransferTypesPostRequest
*/
func (a *DataTransferTypeApiService) ApiDataTransferTypesPost(ctx context.Context) ApiApiDataTransferTypesPostRequest {
	return ApiApiDataTransferTypesPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DataTransferTypeJsonhal
func (a *DataTransferTypeApiService) ApiDataTransferTypesPostExecute(r ApiApiDataTransferTypesPostRequest) (*DataTransferTypeJsonhal, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataTransferTypeJsonhal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataTransferTypeApiService.ApiDataTransferTypesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data_transfer_types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataTransferTypeJsonhal == nil {
		return localVarReturnValue, nil, reportError("dataTransferTypeJsonhal is required and must be specified")
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
	localVarPostBody = r.dataTransferTypeJsonhal
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
