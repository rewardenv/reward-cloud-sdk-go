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


// TeamEnvVarApiService TeamEnvVarApi service
type TeamEnvVarApiService service

type ApiApiTeamEnvVarsGetCollectionRequest struct {
	ctx context.Context
	ApiService *TeamEnvVarApiService
	page *int32
	itemsPerPage *int32
	team *string
	team2 *[]string
	envVarType *string
	envVarType2 *[]string
}

// The collection page number
func (r ApiApiTeamEnvVarsGetCollectionRequest) Page(page int32) ApiApiTeamEnvVarsGetCollectionRequest {
	r.page = &page
	return r
}

// The number of items per page
func (r ApiApiTeamEnvVarsGetCollectionRequest) ItemsPerPage(itemsPerPage int32) ApiApiTeamEnvVarsGetCollectionRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// 
func (r ApiApiTeamEnvVarsGetCollectionRequest) Team(team string) ApiApiTeamEnvVarsGetCollectionRequest {
	r.team = &team
	return r
}

// 
func (r ApiApiTeamEnvVarsGetCollectionRequest) Team2(team2 []string) ApiApiTeamEnvVarsGetCollectionRequest {
	r.team2 = &team2
	return r
}

// 
func (r ApiApiTeamEnvVarsGetCollectionRequest) EnvVarType(envVarType string) ApiApiTeamEnvVarsGetCollectionRequest {
	r.envVarType = &envVarType
	return r
}

// 
func (r ApiApiTeamEnvVarsGetCollectionRequest) EnvVarType2(envVarType2 []string) ApiApiTeamEnvVarsGetCollectionRequest {
	r.envVarType2 = &envVarType2
	return r
}

func (r ApiApiTeamEnvVarsGetCollectionRequest) Execute() ([]TeamEnvVar, *http.Response, error) {
	return r.ApiService.ApiTeamEnvVarsGetCollectionExecute(r)
}

/*
ApiTeamEnvVarsGetCollection Retrieves the collection of TeamEnvVar resources.

Retrieves the collection of TeamEnvVar resources.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiTeamEnvVarsGetCollectionRequest
*/
func (a *TeamEnvVarApiService) ApiTeamEnvVarsGetCollection(ctx context.Context) ApiApiTeamEnvVarsGetCollectionRequest {
	return ApiApiTeamEnvVarsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []TeamEnvVar
func (a *TeamEnvVarApiService) ApiTeamEnvVarsGetCollectionExecute(r ApiApiTeamEnvVarsGetCollectionRequest) ([]TeamEnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TeamEnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamEnvVarApiService.ApiTeamEnvVarsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team_env_vars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.itemsPerPage != nil {
		localVarQueryParams.Add("itemsPerPage", parameterToString(*r.itemsPerPage, ""))
	}
	if r.team != nil {
		localVarQueryParams.Add("team", parameterToString(*r.team, ""))
	}
	if r.team2 != nil {
		t := *r.team2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("team[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("team[]", parameterToString(t, "multi"))
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

type ApiApiTeamEnvVarsIdDeleteRequest struct {
	ctx context.Context
	ApiService *TeamEnvVarApiService
	id string
}

func (r ApiApiTeamEnvVarsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiTeamEnvVarsIdDeleteExecute(r)
}

/*
ApiTeamEnvVarsIdDelete Removes the TeamEnvVar resource.

Removes the TeamEnvVar resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id TeamEnvVar identifier
 @return ApiApiTeamEnvVarsIdDeleteRequest
*/
func (a *TeamEnvVarApiService) ApiTeamEnvVarsIdDelete(ctx context.Context, id string) ApiApiTeamEnvVarsIdDeleteRequest {
	return ApiApiTeamEnvVarsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *TeamEnvVarApiService) ApiTeamEnvVarsIdDeleteExecute(r ApiApiTeamEnvVarsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamEnvVarApiService.ApiTeamEnvVarsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team_env_vars/{id}"
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

type ApiApiTeamEnvVarsIdGetRequest struct {
	ctx context.Context
	ApiService *TeamEnvVarApiService
	id string
}

func (r ApiApiTeamEnvVarsIdGetRequest) Execute() (*TeamEnvVar, *http.Response, error) {
	return r.ApiService.ApiTeamEnvVarsIdGetExecute(r)
}

/*
ApiTeamEnvVarsIdGet Retrieves a TeamEnvVar resource.

Retrieves a TeamEnvVar resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id TeamEnvVar identifier
 @return ApiApiTeamEnvVarsIdGetRequest
*/
func (a *TeamEnvVarApiService) ApiTeamEnvVarsIdGet(ctx context.Context, id string) ApiApiTeamEnvVarsIdGetRequest {
	return ApiApiTeamEnvVarsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TeamEnvVar
func (a *TeamEnvVarApiService) ApiTeamEnvVarsIdGetExecute(r ApiApiTeamEnvVarsIdGetRequest) (*TeamEnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamEnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamEnvVarApiService.ApiTeamEnvVarsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team_env_vars/{id}"
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

type ApiApiTeamEnvVarsIdPatchRequest struct {
	ctx context.Context
	ApiService *TeamEnvVarApiService
	id string
	teamEnvVar *TeamEnvVar
}

// The updated TeamEnvVar resource
func (r ApiApiTeamEnvVarsIdPatchRequest) TeamEnvVar(teamEnvVar TeamEnvVar) ApiApiTeamEnvVarsIdPatchRequest {
	r.teamEnvVar = &teamEnvVar
	return r
}

func (r ApiApiTeamEnvVarsIdPatchRequest) Execute() (*TeamEnvVar, *http.Response, error) {
	return r.ApiService.ApiTeamEnvVarsIdPatchExecute(r)
}

/*
ApiTeamEnvVarsIdPatch Updates the TeamEnvVar resource.

Updates the TeamEnvVar resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id TeamEnvVar identifier
 @return ApiApiTeamEnvVarsIdPatchRequest
*/
func (a *TeamEnvVarApiService) ApiTeamEnvVarsIdPatch(ctx context.Context, id string) ApiApiTeamEnvVarsIdPatchRequest {
	return ApiApiTeamEnvVarsIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TeamEnvVar
func (a *TeamEnvVarApiService) ApiTeamEnvVarsIdPatchExecute(r ApiApiTeamEnvVarsIdPatchRequest) (*TeamEnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamEnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamEnvVarApiService.ApiTeamEnvVarsIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team_env_vars/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.teamEnvVar == nil {
		return localVarReturnValue, nil, reportError("teamEnvVar is required and must be specified")
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
	localVarPostBody = r.teamEnvVar
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

type ApiApiTeamEnvVarsIdPutRequest struct {
	ctx context.Context
	ApiService *TeamEnvVarApiService
	id string
	teamEnvVar *TeamEnvVar
}

// The updated TeamEnvVar resource
func (r ApiApiTeamEnvVarsIdPutRequest) TeamEnvVar(teamEnvVar TeamEnvVar) ApiApiTeamEnvVarsIdPutRequest {
	r.teamEnvVar = &teamEnvVar
	return r
}

func (r ApiApiTeamEnvVarsIdPutRequest) Execute() (*TeamEnvVar, *http.Response, error) {
	return r.ApiService.ApiTeamEnvVarsIdPutExecute(r)
}

/*
ApiTeamEnvVarsIdPut Replaces the TeamEnvVar resource.

Replaces the TeamEnvVar resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id TeamEnvVar identifier
 @return ApiApiTeamEnvVarsIdPutRequest
*/
func (a *TeamEnvVarApiService) ApiTeamEnvVarsIdPut(ctx context.Context, id string) ApiApiTeamEnvVarsIdPutRequest {
	return ApiApiTeamEnvVarsIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TeamEnvVar
func (a *TeamEnvVarApiService) ApiTeamEnvVarsIdPutExecute(r ApiApiTeamEnvVarsIdPutRequest) (*TeamEnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamEnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamEnvVarApiService.ApiTeamEnvVarsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team_env_vars/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.teamEnvVar == nil {
		return localVarReturnValue, nil, reportError("teamEnvVar is required and must be specified")
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
	localVarPostBody = r.teamEnvVar
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

type ApiApiTeamEnvVarsPostRequest struct {
	ctx context.Context
	ApiService *TeamEnvVarApiService
	teamEnvVar *TeamEnvVar
}

// The new TeamEnvVar resource
func (r ApiApiTeamEnvVarsPostRequest) TeamEnvVar(teamEnvVar TeamEnvVar) ApiApiTeamEnvVarsPostRequest {
	r.teamEnvVar = &teamEnvVar
	return r
}

func (r ApiApiTeamEnvVarsPostRequest) Execute() (*TeamEnvVar, *http.Response, error) {
	return r.ApiService.ApiTeamEnvVarsPostExecute(r)
}

/*
ApiTeamEnvVarsPost Creates a TeamEnvVar resource.

Creates a TeamEnvVar resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiTeamEnvVarsPostRequest
*/
func (a *TeamEnvVarApiService) ApiTeamEnvVarsPost(ctx context.Context) ApiApiTeamEnvVarsPostRequest {
	return ApiApiTeamEnvVarsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TeamEnvVar
func (a *TeamEnvVarApiService) ApiTeamEnvVarsPostExecute(r ApiApiTeamEnvVarsPostRequest) (*TeamEnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamEnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamEnvVarApiService.ApiTeamEnvVarsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team_env_vars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.teamEnvVar == nil {
		return localVarReturnValue, nil, reportError("teamEnvVar is required and must be specified")
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
	localVarPostBody = r.teamEnvVar
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
