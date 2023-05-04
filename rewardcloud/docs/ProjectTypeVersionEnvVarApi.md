# \ProjectTypeVersionEnvVarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectTypeVersionEnvVarsGetCollection**](ProjectTypeVersionEnvVarApi.md#ApiProjectTypeVersionEnvVarsGetCollection) | **Get** /api/project_type_version_env_vars | Retrieves the collection of ProjectTypeVersionEnvVar resources.
[**ApiProjectTypeVersionEnvVarsIdDelete**](ProjectTypeVersionEnvVarApi.md#ApiProjectTypeVersionEnvVarsIdDelete) | **Delete** /api/project_type_version_env_vars/{id} | Removes the ProjectTypeVersionEnvVar resource.
[**ApiProjectTypeVersionEnvVarsIdGet**](ProjectTypeVersionEnvVarApi.md#ApiProjectTypeVersionEnvVarsIdGet) | **Get** /api/project_type_version_env_vars/{id} | Retrieves a ProjectTypeVersionEnvVar resource.
[**ApiProjectTypeVersionEnvVarsIdPatch**](ProjectTypeVersionEnvVarApi.md#ApiProjectTypeVersionEnvVarsIdPatch) | **Patch** /api/project_type_version_env_vars/{id} | Updates the ProjectTypeVersionEnvVar resource.
[**ApiProjectTypeVersionEnvVarsIdPut**](ProjectTypeVersionEnvVarApi.md#ApiProjectTypeVersionEnvVarsIdPut) | **Put** /api/project_type_version_env_vars/{id} | Replaces the ProjectTypeVersionEnvVar resource.
[**ApiProjectTypeVersionEnvVarsPost**](ProjectTypeVersionEnvVarApi.md#ApiProjectTypeVersionEnvVarsPost) | **Post** /api/project_type_version_env_vars | Creates a ProjectTypeVersionEnvVar resource.



## ApiProjectTypeVersionEnvVarsGetCollection

> ApiProjectTypeVersionEnvVarsGetCollection200Response ApiProjectTypeVersionEnvVarsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ProjectTypeVersion(projectTypeVersion).ProjectTypeVersion2(projectTypeVersion2).Execute()

Retrieves the collection of ProjectTypeVersionEnvVar resources.



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
    projectTypeVersion := "projectTypeVersion_example" // string |  (optional)
    projectTypeVersion2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ProjectTypeVersion(projectTypeVersion).ProjectTypeVersion2(projectTypeVersion2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionEnvVarsGetCollection`: ApiProjectTypeVersionEnvVarsGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **projectTypeVersion** | **string** |  | 
 **projectTypeVersion2** | **[]string** |  | 

### Return type

[**ApiProjectTypeVersionEnvVarsGetCollection200Response**](ApiProjectTypeVersionEnvVarsGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionEnvVarsIdDelete

> ApiProjectTypeVersionEnvVarsIdDelete(ctx, id).Execute()

Removes the ProjectTypeVersionEnvVar resource.



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
    id := "id_example" // string | ProjectTypeVersionEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersionEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarsIdDeleteRequest struct via the builder pattern


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


## ApiProjectTypeVersionEnvVarsIdGet

> ProjectTypeVersionEnvVarJsonhal ApiProjectTypeVersionEnvVarsIdGet(ctx, id).Execute()

Retrieves a ProjectTypeVersionEnvVar resource.



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
    id := "id_example" // string | ProjectTypeVersionEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionEnvVarsIdGet`: ProjectTypeVersionEnvVarJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersionEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectTypeVersionEnvVarJsonhal**](ProjectTypeVersionEnvVarJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionEnvVarsIdPatch

> ProjectTypeVersionEnvVarJsonhal ApiProjectTypeVersionEnvVarsIdPatch(ctx, id).ProjectTypeVersionEnvVar(projectTypeVersionEnvVar).Execute()

Updates the ProjectTypeVersionEnvVar resource.



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
    id := "id_example" // string | ProjectTypeVersionEnvVar identifier
    projectTypeVersionEnvVar := *openapiclient.NewProjectTypeVersionEnvVar() // ProjectTypeVersionEnvVar | The updated ProjectTypeVersionEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsIdPatch(context.Background(), id).ProjectTypeVersionEnvVar(projectTypeVersionEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionEnvVarsIdPatch`: ProjectTypeVersionEnvVarJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersionEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectTypeVersionEnvVar** | [**ProjectTypeVersionEnvVar**](ProjectTypeVersionEnvVar.md) | The updated ProjectTypeVersionEnvVar resource | 

### Return type

[**ProjectTypeVersionEnvVarJsonhal**](ProjectTypeVersionEnvVarJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionEnvVarsIdPut

> ProjectTypeVersionEnvVarJsonhal ApiProjectTypeVersionEnvVarsIdPut(ctx, id).ProjectTypeVersionEnvVarJsonhal(projectTypeVersionEnvVarJsonhal).Execute()

Replaces the ProjectTypeVersionEnvVar resource.



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
    id := "id_example" // string | ProjectTypeVersionEnvVar identifier
    projectTypeVersionEnvVarJsonhal := *openapiclient.NewProjectTypeVersionEnvVarJsonhal() // ProjectTypeVersionEnvVarJsonhal | The updated ProjectTypeVersionEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsIdPut(context.Background(), id).ProjectTypeVersionEnvVarJsonhal(projectTypeVersionEnvVarJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionEnvVarsIdPut`: ProjectTypeVersionEnvVarJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersionEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectTypeVersionEnvVarJsonhal** | [**ProjectTypeVersionEnvVarJsonhal**](ProjectTypeVersionEnvVarJsonhal.md) | The updated ProjectTypeVersionEnvVar resource | 

### Return type

[**ProjectTypeVersionEnvVarJsonhal**](ProjectTypeVersionEnvVarJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionEnvVarsPost

> ProjectTypeVersionEnvVarJsonhal ApiProjectTypeVersionEnvVarsPost(ctx).ProjectTypeVersionEnvVarJsonhal(projectTypeVersionEnvVarJsonhal).Execute()

Creates a ProjectTypeVersionEnvVar resource.



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
    projectTypeVersionEnvVarJsonhal := *openapiclient.NewProjectTypeVersionEnvVarJsonhal() // ProjectTypeVersionEnvVarJsonhal | The new ProjectTypeVersionEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsPost(context.Background()).ProjectTypeVersionEnvVarJsonhal(projectTypeVersionEnvVarJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionEnvVarsPost`: ProjectTypeVersionEnvVarJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionEnvVarApi.ApiProjectTypeVersionEnvVarsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectTypeVersionEnvVarJsonhal** | [**ProjectTypeVersionEnvVarJsonhal**](ProjectTypeVersionEnvVarJsonhal.md) | The new ProjectTypeVersionEnvVar resource | 

### Return type

[**ProjectTypeVersionEnvVarJsonhal**](ProjectTypeVersionEnvVarJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

