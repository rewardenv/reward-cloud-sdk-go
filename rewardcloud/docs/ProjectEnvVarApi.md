# \ProjectEnvVarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectEnvVarsGetCollection**](ProjectEnvVarApi.md#ApiProjectEnvVarsGetCollection) | **Get** /api/project_env_vars | Retrieves the collection of ProjectEnvVar resources.
[**ApiProjectEnvVarsIdDelete**](ProjectEnvVarApi.md#ApiProjectEnvVarsIdDelete) | **Delete** /api/project_env_vars/{id} | Removes the ProjectEnvVar resource.
[**ApiProjectEnvVarsIdGet**](ProjectEnvVarApi.md#ApiProjectEnvVarsIdGet) | **Get** /api/project_env_vars/{id} | Retrieves a ProjectEnvVar resource.
[**ApiProjectEnvVarsIdPatch**](ProjectEnvVarApi.md#ApiProjectEnvVarsIdPatch) | **Patch** /api/project_env_vars/{id} | Updates the ProjectEnvVar resource.
[**ApiProjectEnvVarsIdPut**](ProjectEnvVarApi.md#ApiProjectEnvVarsIdPut) | **Put** /api/project_env_vars/{id} | Replaces the ProjectEnvVar resource.
[**ApiProjectEnvVarsPost**](ProjectEnvVarApi.md#ApiProjectEnvVarsPost) | **Post** /api/project_env_vars | Creates a ProjectEnvVar resource.



## ApiProjectEnvVarsGetCollection

> []ProjectEnvVar ApiProjectEnvVarsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Project(project).Project2(project2).EnvVarType(envVarType).EnvVarType2(envVarType2).Execute()

Retrieves the collection of ProjectEnvVar resources.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    itemsPerPage := int32(56) // int32 | The number of items per page (optional) (default to 30)
    project := "project_example" // string |  (optional)
    project2 := []string{"Inner_example"} // []string |  (optional)
    envVarType := "envVarType_example" // string |  (optional)
    envVarType2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectEnvVarApi.ApiProjectEnvVarsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Project(project).Project2(project2).EnvVarType(envVarType).EnvVarType2(envVarType2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvVarApi.ApiProjectEnvVarsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectEnvVarsGetCollection`: []ProjectEnvVar
    fmt.Fprintf(os.Stdout, "Response from `ProjectEnvVarApi.ApiProjectEnvVarsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectEnvVarsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **project** | **string** |  | 
 **project2** | **[]string** |  | 
 **envVarType** | **string** |  | 
 **envVarType2** | **[]string** |  | 

### Return type

[**[]ProjectEnvVar**](ProjectEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectEnvVarsIdDelete

> ApiProjectEnvVarsIdDelete(ctx, id).Execute()

Removes the ProjectEnvVar resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | ProjectEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectEnvVarApi.ApiProjectEnvVarsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvVarApi.ApiProjectEnvVarsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectEnvVarsIdDeleteRequest struct via the builder pattern


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


## ApiProjectEnvVarsIdGet

> ProjectEnvVar ApiProjectEnvVarsIdGet(ctx, id).Execute()

Retrieves a ProjectEnvVar resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | ProjectEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectEnvVarApi.ApiProjectEnvVarsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvVarApi.ApiProjectEnvVarsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectEnvVarsIdGet`: ProjectEnvVar
    fmt.Fprintf(os.Stdout, "Response from `ProjectEnvVarApi.ApiProjectEnvVarsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectEnvVarsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectEnvVar**](ProjectEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectEnvVarsIdPatch

> ProjectEnvVar ApiProjectEnvVarsIdPatch(ctx, id).ProjectEnvVar(projectEnvVar).Execute()

Updates the ProjectEnvVar resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectEnvVar := *openapiclient.NewProjectEnvVar() // ProjectEnvVar | The updated ProjectEnvVar resource
    id := "id_example" // string | ProjectEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectEnvVarApi.ApiProjectEnvVarsIdPatch(context.Background(), id).ProjectEnvVar(projectEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvVarApi.ApiProjectEnvVarsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectEnvVarsIdPatch`: ProjectEnvVar
    fmt.Fprintf(os.Stdout, "Response from `ProjectEnvVarApi.ApiProjectEnvVarsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectEnvVarsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectEnvVar** | [**ProjectEnvVar**](ProjectEnvVar.md) | The updated ProjectEnvVar resource | 


### Return type

[**ProjectEnvVar**](ProjectEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectEnvVarsIdPut

> ProjectEnvVar ApiProjectEnvVarsIdPut(ctx, id).ProjectEnvVar(projectEnvVar).Execute()

Replaces the ProjectEnvVar resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectEnvVar := *openapiclient.NewProjectEnvVar() // ProjectEnvVar | The updated ProjectEnvVar resource
    id := "id_example" // string | ProjectEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectEnvVarApi.ApiProjectEnvVarsIdPut(context.Background(), id).ProjectEnvVar(projectEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvVarApi.ApiProjectEnvVarsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectEnvVarsIdPut`: ProjectEnvVar
    fmt.Fprintf(os.Stdout, "Response from `ProjectEnvVarApi.ApiProjectEnvVarsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectEnvVarsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectEnvVar** | [**ProjectEnvVar**](ProjectEnvVar.md) | The updated ProjectEnvVar resource | 


### Return type

[**ProjectEnvVar**](ProjectEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectEnvVarsPost

> ProjectEnvVar ApiProjectEnvVarsPost(ctx).ProjectEnvVar(projectEnvVar).Execute()

Creates a ProjectEnvVar resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectEnvVar := *openapiclient.NewProjectEnvVar() // ProjectEnvVar | The new ProjectEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectEnvVarApi.ApiProjectEnvVarsPost(context.Background()).ProjectEnvVar(projectEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvVarApi.ApiProjectEnvVarsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectEnvVarsPost`: ProjectEnvVar
    fmt.Fprintf(os.Stdout, "Response from `ProjectEnvVarApi.ApiProjectEnvVarsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectEnvVarsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectEnvVar** | [**ProjectEnvVar**](ProjectEnvVar.md) | The new ProjectEnvVar resource | 

### Return type

[**ProjectEnvVar**](ProjectEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

