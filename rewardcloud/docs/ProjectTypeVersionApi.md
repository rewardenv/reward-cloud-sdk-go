# \ProjectTypeVersionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectTypeVersionsGetCollection**](ProjectTypeVersionApi.md#ApiProjectTypeVersionsGetCollection) | **Get** /api/project_type_versions | Retrieves the collection of ProjectTypeVersion resources.
[**ApiProjectTypeVersionsIdDelete**](ProjectTypeVersionApi.md#ApiProjectTypeVersionsIdDelete) | **Delete** /api/project_type_versions/{id} | Removes the ProjectTypeVersion resource.
[**ApiProjectTypeVersionsIdGet**](ProjectTypeVersionApi.md#ApiProjectTypeVersionsIdGet) | **Get** /api/project_type_versions/{id} | Retrieves a ProjectTypeVersion resource.
[**ApiProjectTypeVersionsIdPatch**](ProjectTypeVersionApi.md#ApiProjectTypeVersionsIdPatch) | **Patch** /api/project_type_versions/{id} | Updates the ProjectTypeVersion resource.
[**ApiProjectTypeVersionsIdPut**](ProjectTypeVersionApi.md#ApiProjectTypeVersionsIdPut) | **Put** /api/project_type_versions/{id} | Replaces the ProjectTypeVersion resource.
[**ApiProjectTypeVersionsPost**](ProjectTypeVersionApi.md#ApiProjectTypeVersionsPost) | **Post** /api/project_type_versions | Creates a ProjectTypeVersion resource.



## ApiProjectTypeVersionsGetCollection

> []ProjectTypeVersion ApiProjectTypeVersionsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ProjectType(projectType).ProjectType2(projectType2).Execute()

Retrieves the collection of ProjectTypeVersion resources.



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
    projectType := "projectType_example" // string |  (optional)
    projectType2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionApi.ApiProjectTypeVersionsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ProjectType(projectType).ProjectType2(projectType2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionApi.ApiProjectTypeVersionsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionsGetCollection`: []ProjectTypeVersion
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionApi.ApiProjectTypeVersionsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **projectType** | **string** |  | 
 **projectType2** | **[]string** |  | 

### Return type

[**[]ProjectTypeVersion**](ProjectTypeVersion.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionsIdDelete

> ApiProjectTypeVersionsIdDelete(ctx, id).Execute()

Removes the ProjectTypeVersion resource.



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
    id := "id_example" // string | ProjectTypeVersion identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectTypeVersionApi.ApiProjectTypeVersionsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionApi.ApiProjectTypeVersionsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersion identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionsIdDeleteRequest struct via the builder pattern


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


## ApiProjectTypeVersionsIdGet

> ProjectTypeVersion ApiProjectTypeVersionsIdGet(ctx, id).Execute()

Retrieves a ProjectTypeVersion resource.



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
    id := "id_example" // string | ProjectTypeVersion identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionApi.ApiProjectTypeVersionsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionApi.ApiProjectTypeVersionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionsIdGet`: ProjectTypeVersion
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionApi.ApiProjectTypeVersionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersion identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectTypeVersion**](ProjectTypeVersion.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionsIdPatch

> ProjectTypeVersion ApiProjectTypeVersionsIdPatch(ctx, id).ProjectTypeVersion(projectTypeVersion).Execute()

Updates the ProjectTypeVersion resource.



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
    projectTypeVersion := *openapiclient.NewProjectTypeVersion() // ProjectTypeVersion | The updated ProjectTypeVersion resource
    id := "id_example" // string | ProjectTypeVersion identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionApi.ApiProjectTypeVersionsIdPatch(context.Background(), id).ProjectTypeVersion(projectTypeVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionApi.ApiProjectTypeVersionsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionsIdPatch`: ProjectTypeVersion
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionApi.ApiProjectTypeVersionsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersion identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectTypeVersion** | [**ProjectTypeVersion**](ProjectTypeVersion.md) | The updated ProjectTypeVersion resource | 


### Return type

[**ProjectTypeVersion**](ProjectTypeVersion.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionsIdPut

> ProjectTypeVersion ApiProjectTypeVersionsIdPut(ctx, id).ProjectTypeVersion(projectTypeVersion).Execute()

Replaces the ProjectTypeVersion resource.



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
    projectTypeVersion := *openapiclient.NewProjectTypeVersion() // ProjectTypeVersion | The updated ProjectTypeVersion resource
    id := "id_example" // string | ProjectTypeVersion identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionApi.ApiProjectTypeVersionsIdPut(context.Background(), id).ProjectTypeVersion(projectTypeVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionApi.ApiProjectTypeVersionsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionsIdPut`: ProjectTypeVersion
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionApi.ApiProjectTypeVersionsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectTypeVersion identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectTypeVersion** | [**ProjectTypeVersion**](ProjectTypeVersion.md) | The updated ProjectTypeVersion resource | 


### Return type

[**ProjectTypeVersion**](ProjectTypeVersion.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypeVersionsPost

> ProjectTypeVersion ApiProjectTypeVersionsPost(ctx).ProjectTypeVersion(projectTypeVersion).Execute()

Creates a ProjectTypeVersion resource.



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
    projectTypeVersion := *openapiclient.NewProjectTypeVersion() // ProjectTypeVersion | The new ProjectTypeVersion resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeVersionApi.ApiProjectTypeVersionsPost(context.Background()).ProjectTypeVersion(projectTypeVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeVersionApi.ApiProjectTypeVersionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypeVersionsPost`: ProjectTypeVersion
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeVersionApi.ApiProjectTypeVersionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypeVersionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectTypeVersion** | [**ProjectTypeVersion**](ProjectTypeVersion.md) | The new ProjectTypeVersion resource | 

### Return type

[**ProjectTypeVersion**](ProjectTypeVersion.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

