# \ProjectTypeVersionEnvVarExampleApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectTypeVersionEnvVarExamplesGetCollection**](ProjectTypeVersionEnvVarExampleApi.md#ApiProjectTypeVersionEnvVarExamplesGetCollection) | **Get** /api/project_type_version_env_var_examples | Retrieves the collection of ProjectTypeVersionEnvVarExample resources.
[**ApiProjectTypeVersionEnvVarExamplesIdDelete**](ProjectTypeVersionEnvVarExampleApi.md#ApiProjectTypeVersionEnvVarExamplesIdDelete) | **Delete** /api/project_type_version_env_var_examples/{id} | Removes the ProjectTypeVersionEnvVarExample resource.
[**ApiProjectTypeVersionEnvVarExamplesIdGet**](ProjectTypeVersionEnvVarExampleApi.md#ApiProjectTypeVersionEnvVarExamplesIdGet) | **Get** /api/project_type_version_env_var_examples/{id} | Retrieves a ProjectTypeVersionEnvVarExample resource.
[**ApiProjectTypeVersionEnvVarExamplesIdPatch**](ProjectTypeVersionEnvVarExampleApi.md#ApiProjectTypeVersionEnvVarExamplesIdPatch) | **Patch** /api/project_type_version_env_var_examples/{id} | Updates the ProjectTypeVersionEnvVarExample resource.
[**ApiProjectTypeVersionEnvVarExamplesIdPut**](ProjectTypeVersionEnvVarExampleApi.md#ApiProjectTypeVersionEnvVarExamplesIdPut) | **Put** /api/project_type_version_env_var_examples/{id} | Replaces the ProjectTypeVersionEnvVarExample resource.
[**ApiProjectTypeVersionEnvVarExamplesPost**](ProjectTypeVersionEnvVarExampleApi.md#ApiProjectTypeVersionEnvVarExamplesPost) | **Post** /api/project_type_version_env_var_examples | Creates a ProjectTypeVersionEnvVarExample resource.



## ApiProjectTypeVersionEnvVarExamplesGetCollection

> []ProjectTypeVersionEnvVarExample ApiProjectTypeVersionEnvVarExamplesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ProjectTypeVersionEnvVar(projectTypeVersionEnvVar).ProjectTypeVersionEnvVar2(projectTypeVersionEnvVar2).Execute()

Retrieves the collection of ProjectTypeVersionEnvVarExample resources.



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
    projectTypeVersionEnvVar := "projectTypeVersionEnvVar_example" // string |  (optional)
    projectTypeVersionEnvVar2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ProjectTypeVersionEnvVar(projectTypeVersionEnvVar).ProjectTypeVersionEnvVar2(projectTypeVersionEnvVar2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionEnvVarExamplesGetCollection`: []ProjectTypeVersionEnvVarExample
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarExamplesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **projectTypeVersionEnvVar** | **string** |  | 
 **projectTypeVersionEnvVar2** | **[]string** |  | 

### Return type

[**[]ProjectTypeVersionEnvVarExample**](ProjectTypeVersionEnvVarExample.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionEnvVarExamplesIdDelete

> ApiProjectTypeVersionEnvVarExamplesIdDelete(ctx, id).Execute()

Removes the ProjectTypeVersionEnvVarExample resource.



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
    id := "id_example" // string | ProjectTypeVersionEnvVarExample identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersionEnvVarExample identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarExamplesIdDeleteRequest struct via the builder pattern


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


## ApiProjectTypeVersionEnvVarExamplesIdGet

> ProjectTypeVersionEnvVarExample ApiProjectTypeVersionEnvVarExamplesIdGet(ctx, id).Execute()

Retrieves a ProjectTypeVersionEnvVarExample resource.



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
    id := "id_example" // string | ProjectTypeVersionEnvVarExample identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionEnvVarExamplesIdGet`: ProjectTypeVersionEnvVarExample
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersionEnvVarExample identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarExamplesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectTypeVersionEnvVarExample**](ProjectTypeVersionEnvVarExample.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionEnvVarExamplesIdPatch

> ProjectTypeVersionEnvVarExample ApiProjectTypeVersionEnvVarExamplesIdPatch(ctx, id).ProjectTypeVersionEnvVarExample(projectTypeVersionEnvVarExample).Execute()

Updates the ProjectTypeVersionEnvVarExample resource.



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
    id := "id_example" // string | ProjectTypeVersionEnvVarExample identifier
    projectTypeVersionEnvVarExample := *openapiclient.NewProjectTypeVersionEnvVarExample() // ProjectTypeVersionEnvVarExample | The updated ProjectTypeVersionEnvVarExample resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesIdPatch(context.Background(), id).ProjectTypeVersionEnvVarExample(projectTypeVersionEnvVarExample).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionEnvVarExamplesIdPatch`: ProjectTypeVersionEnvVarExample
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersionEnvVarExample identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarExamplesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectTypeVersionEnvVarExample** | [**ProjectTypeVersionEnvVarExample**](ProjectTypeVersionEnvVarExample.md) | The updated ProjectTypeVersionEnvVarExample resource | 

### Return type

[**ProjectTypeVersionEnvVarExample**](ProjectTypeVersionEnvVarExample.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionEnvVarExamplesIdPut

> ProjectTypeVersionEnvVarExample ApiProjectTypeVersionEnvVarExamplesIdPut(ctx, id).ProjectTypeVersionEnvVarExample(projectTypeVersionEnvVarExample).Execute()

Replaces the ProjectTypeVersionEnvVarExample resource.



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
    id := "id_example" // string | ProjectTypeVersionEnvVarExample identifier
    projectTypeVersionEnvVarExample := *openapiclient.NewProjectTypeVersionEnvVarExample() // ProjectTypeVersionEnvVarExample | The updated ProjectTypeVersionEnvVarExample resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesIdPut(context.Background(), id).ProjectTypeVersionEnvVarExample(projectTypeVersionEnvVarExample).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionEnvVarExamplesIdPut`: ProjectTypeVersionEnvVarExample
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersionEnvVarExample identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarExamplesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectTypeVersionEnvVarExample** | [**ProjectTypeVersionEnvVarExample**](ProjectTypeVersionEnvVarExample.md) | The updated ProjectTypeVersionEnvVarExample resource | 

### Return type

[**ProjectTypeVersionEnvVarExample**](ProjectTypeVersionEnvVarExample.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionEnvVarExamplesPost

> ProjectTypeVersionEnvVarExample ApiProjectTypeVersionEnvVarExamplesPost(ctx).ProjectTypeVersionEnvVarExample(projectTypeVersionEnvVarExample).Execute()

Creates a ProjectTypeVersionEnvVarExample resource.



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
    projectTypeVersionEnvVarExample := *openapiclient.NewProjectTypeVersionEnvVarExample() // ProjectTypeVersionEnvVarExample | The new ProjectTypeVersionEnvVarExample resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesPost(context.Background()).ProjectTypeVersionEnvVarExample(projectTypeVersionEnvVarExample).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionEnvVarExamplesPost`: ProjectTypeVersionEnvVarExample
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionEnvVarExampleApi.ApiProjectTypeVersionEnvVarExamplesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionEnvVarExamplesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectTypeVersionEnvVarExample** | [**ProjectTypeVersionEnvVarExample**](ProjectTypeVersionEnvVarExample.md) | The new ProjectTypeVersionEnvVarExample resource | 

### Return type

[**ProjectTypeVersionEnvVarExample**](ProjectTypeVersionEnvVarExample.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

