# \TeamEnvVarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTeamEnvVarsGetCollection**](TeamEnvVarApi.md#ApiTeamEnvVarsGetCollection) | **Get** /api/team_env_vars | Retrieves the collection of TeamEnvVar resources.
[**ApiTeamEnvVarsIdDelete**](TeamEnvVarApi.md#ApiTeamEnvVarsIdDelete) | **Delete** /api/team_env_vars/{id} | Removes the TeamEnvVar resource.
[**ApiTeamEnvVarsIdGet**](TeamEnvVarApi.md#ApiTeamEnvVarsIdGet) | **Get** /api/team_env_vars/{id} | Retrieves a TeamEnvVar resource.
[**ApiTeamEnvVarsIdPatch**](TeamEnvVarApi.md#ApiTeamEnvVarsIdPatch) | **Patch** /api/team_env_vars/{id} | Updates the TeamEnvVar resource.
[**ApiTeamEnvVarsIdPut**](TeamEnvVarApi.md#ApiTeamEnvVarsIdPut) | **Put** /api/team_env_vars/{id} | Replaces the TeamEnvVar resource.
[**ApiTeamEnvVarsPost**](TeamEnvVarApi.md#ApiTeamEnvVarsPost) | **Post** /api/team_env_vars | Creates a TeamEnvVar resource.



## ApiTeamEnvVarsGetCollection

> []TeamEnvVar ApiTeamEnvVarsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Team(team).Team2(team2).EnvVarType(envVarType).EnvVarType2(envVarType2).Execute()

Retrieves the collection of TeamEnvVar resources.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    itemsPerPage := int32(56) // int32 | The number of items per page (optional) (default to 30)
    team := "team_example" // string |  (optional)
    team2 := []string{"Inner_example"} // []string |  (optional)
    envVarType := "envVarType_example" // string |  (optional)
    envVarType2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamEnvVarApi.ApiTeamEnvVarsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Team(team).Team2(team2).EnvVarType(envVarType).EnvVarType2(envVarType2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamEnvVarApi.ApiTeamEnvVarsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamEnvVarsGetCollection`: []TeamEnvVar
    fmt.Fprintf(os.Stdout, "Response from `TeamEnvVarApi.ApiTeamEnvVarsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamEnvVarsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **team** | **string** |  | 
 **team2** | **[]string** |  | 
 **envVarType** | **string** |  | 
 **envVarType2** | **[]string** |  | 

### Return type

[**[]TeamEnvVar**](TeamEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTeamEnvVarsIdDelete

> ApiTeamEnvVarsIdDelete(ctx, id).Execute()

Removes the TeamEnvVar resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | TeamEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamEnvVarApi.ApiTeamEnvVarsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamEnvVarApi.ApiTeamEnvVarsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TeamEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamEnvVarsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTeamEnvVarsIdGet

> TeamEnvVar ApiTeamEnvVarsIdGet(ctx, id).Execute()

Retrieves a TeamEnvVar resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | TeamEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamEnvVarApi.ApiTeamEnvVarsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamEnvVarApi.ApiTeamEnvVarsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamEnvVarsIdGet`: TeamEnvVar
    fmt.Fprintf(os.Stdout, "Response from `TeamEnvVarApi.ApiTeamEnvVarsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TeamEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamEnvVarsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamEnvVar**](TeamEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTeamEnvVarsIdPatch

> TeamEnvVar ApiTeamEnvVarsIdPatch(ctx, id).TeamEnvVar(teamEnvVar).Execute()

Updates the TeamEnvVar resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | TeamEnvVar identifier
    teamEnvVar := *openapiclient.NewTeamEnvVar() // TeamEnvVar | The updated TeamEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamEnvVarApi.ApiTeamEnvVarsIdPatch(context.Background(), id).TeamEnvVar(teamEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamEnvVarApi.ApiTeamEnvVarsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamEnvVarsIdPatch`: TeamEnvVar
    fmt.Fprintf(os.Stdout, "Response from `TeamEnvVarApi.ApiTeamEnvVarsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TeamEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamEnvVarsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamEnvVar** | [**TeamEnvVar**](TeamEnvVar.md) | The updated TeamEnvVar resource | 

### Return type

[**TeamEnvVar**](TeamEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTeamEnvVarsIdPut

> TeamEnvVar ApiTeamEnvVarsIdPut(ctx, id).TeamEnvVar(teamEnvVar).Execute()

Replaces the TeamEnvVar resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | TeamEnvVar identifier
    teamEnvVar := *openapiclient.NewTeamEnvVar() // TeamEnvVar | The updated TeamEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamEnvVarApi.ApiTeamEnvVarsIdPut(context.Background(), id).TeamEnvVar(teamEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamEnvVarApi.ApiTeamEnvVarsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamEnvVarsIdPut`: TeamEnvVar
    fmt.Fprintf(os.Stdout, "Response from `TeamEnvVarApi.ApiTeamEnvVarsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TeamEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamEnvVarsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamEnvVar** | [**TeamEnvVar**](TeamEnvVar.md) | The updated TeamEnvVar resource | 

### Return type

[**TeamEnvVar**](TeamEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTeamEnvVarsPost

> TeamEnvVar ApiTeamEnvVarsPost(ctx).TeamEnvVar(teamEnvVar).Execute()

Creates a TeamEnvVar resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamEnvVar := *openapiclient.NewTeamEnvVar() // TeamEnvVar | The new TeamEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamEnvVarApi.ApiTeamEnvVarsPost(context.Background()).TeamEnvVar(teamEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamEnvVarApi.ApiTeamEnvVarsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamEnvVarsPost`: TeamEnvVar
    fmt.Fprintf(os.Stdout, "Response from `TeamEnvVarApi.ApiTeamEnvVarsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamEnvVarsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamEnvVar** | [**TeamEnvVar**](TeamEnvVar.md) | The new TeamEnvVar resource | 

### Return type

[**TeamEnvVar**](TeamEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

