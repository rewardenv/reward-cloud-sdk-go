# \ProjectTypeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectTypesGetCollection**](ProjectTypeApi.md#ApiProjectTypesGetCollection) | **Get** /api/project_types | Retrieves the collection of ProjectType resources.
[**ApiProjectTypesIdDelete**](ProjectTypeApi.md#ApiProjectTypesIdDelete) | **Delete** /api/project_types/{id} | Removes the ProjectType resource.
[**ApiProjectTypesIdGet**](ProjectTypeApi.md#ApiProjectTypesIdGet) | **Get** /api/project_types/{id} | Retrieves a ProjectType resource.
[**ApiProjectTypesIdPatch**](ProjectTypeApi.md#ApiProjectTypesIdPatch) | **Patch** /api/project_types/{id} | Updates the ProjectType resource.
[**ApiProjectTypesIdPut**](ProjectTypeApi.md#ApiProjectTypesIdPut) | **Put** /api/project_types/{id} | Replaces the ProjectType resource.
[**ApiProjectTypesPost**](ProjectTypeApi.md#ApiProjectTypesPost) | **Post** /api/project_types | Creates a ProjectType resource.



## ApiProjectTypesGetCollection

> []ProjectType ApiProjectTypesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).OrderName(orderName).Execute()

Retrieves the collection of ProjectType resources.



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
    orderName := "orderName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeApi.ApiProjectTypesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).OrderName(orderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeApi.ApiProjectTypesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypesGetCollection`: []ProjectType
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeApi.ApiProjectTypesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **orderName** | **string** |  | 

### Return type

[**[]ProjectType**](ProjectType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypesIdDelete

> ApiProjectTypesIdDelete(ctx, id).Execute()

Removes the ProjectType resource.



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
    id := "id_example" // string | ProjectType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectTypeApi.ApiProjectTypesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeApi.ApiProjectTypesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypesIdDeleteRequest struct via the builder pattern


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


## ApiProjectTypesIdGet

> ProjectType ApiProjectTypesIdGet(ctx, id).Execute()

Retrieves a ProjectType resource.



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
    id := "id_example" // string | ProjectType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeApi.ApiProjectTypesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeApi.ApiProjectTypesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypesIdGet`: ProjectType
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeApi.ApiProjectTypesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectType**](ProjectType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypesIdPatch

> ProjectType ApiProjectTypesIdPatch(ctx, id).ProjectType(projectType).Execute()

Updates the ProjectType resource.



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
    projectType := *openapiclient.NewProjectType() // ProjectType | The updated ProjectType resource
    id := "id_example" // string | ProjectType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeApi.ApiProjectTypesIdPatch(context.Background(), id).ProjectType(projectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeApi.ApiProjectTypesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypesIdPatch`: ProjectType
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeApi.ApiProjectTypesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectType** | [**ProjectType**](ProjectType.md) | The updated ProjectType resource | 


### Return type

[**ProjectType**](ProjectType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypesIdPut

> ProjectType ApiProjectTypesIdPut(ctx, id).ProjectType(projectType).Execute()

Replaces the ProjectType resource.



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
    projectType := *openapiclient.NewProjectType() // ProjectType | The updated ProjectType resource
    id := "id_example" // string | ProjectType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeApi.ApiProjectTypesIdPut(context.Background(), id).ProjectType(projectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeApi.ApiProjectTypesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypesIdPut`: ProjectType
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeApi.ApiProjectTypesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ProjectType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectType** | [**ProjectType**](ProjectType.md) | The updated ProjectType resource | 


### Return type

[**ProjectType**](ProjectType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectTypesPost

> ProjectType ApiProjectTypesPost(ctx).ProjectType(projectType).Execute()

Creates a ProjectType resource.



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
    projectType := *openapiclient.NewProjectType() // ProjectType | The new ProjectType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTypeApi.ApiProjectTypesPost(context.Background()).ProjectType(projectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypeApi.ApiProjectTypesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectTypesPost`: ProjectType
    fmt.Fprintf(os.Stdout, "Response from `ProjectTypeApi.ApiProjectTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectType** | [**ProjectType**](ProjectType.md) | The new ProjectType resource | 

### Return type

[**ProjectType**](ProjectType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

