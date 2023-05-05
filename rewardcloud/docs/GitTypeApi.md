# \GitTypeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiGitTypesGetCollection**](GitTypeApi.md#ApiGitTypesGetCollection) | **Get** /api/git_types | Retrieves the collection of GitType resources.
[**ApiGitTypesIdDelete**](GitTypeApi.md#ApiGitTypesIdDelete) | **Delete** /api/git_types/{id} | Removes the GitType resource.
[**ApiGitTypesIdGet**](GitTypeApi.md#ApiGitTypesIdGet) | **Get** /api/git_types/{id} | Retrieves a GitType resource.
[**ApiGitTypesIdPatch**](GitTypeApi.md#ApiGitTypesIdPatch) | **Patch** /api/git_types/{id} | Updates the GitType resource.
[**ApiGitTypesIdPut**](GitTypeApi.md#ApiGitTypesIdPut) | **Put** /api/git_types/{id} | Replaces the GitType resource.
[**ApiGitTypesPost**](GitTypeApi.md#ApiGitTypesPost) | **Post** /api/git_types | Creates a GitType resource.



## ApiGitTypesGetCollection

> []GitType ApiGitTypesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of GitType resources.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitTypeApi.ApiGitTypesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitTypeApi.ApiGitTypesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGitTypesGetCollection`: []GitType
    fmt.Fprintf(os.Stdout, "Response from `GitTypeApi.ApiGitTypesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiGitTypesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]GitType**](GitType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGitTypesIdDelete

> ApiGitTypesIdDelete(ctx, id).Execute()

Removes the GitType resource.



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
    id := "id_example" // string | GitType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GitTypeApi.ApiGitTypesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitTypeApi.ApiGitTypesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | GitType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiGitTypesIdDeleteRequest struct via the builder pattern


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


## ApiGitTypesIdGet

> GitType ApiGitTypesIdGet(ctx, id).Execute()

Retrieves a GitType resource.



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
    id := "id_example" // string | GitType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitTypeApi.ApiGitTypesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitTypeApi.ApiGitTypesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGitTypesIdGet`: GitType
    fmt.Fprintf(os.Stdout, "Response from `GitTypeApi.ApiGitTypesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | GitType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiGitTypesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitType**](GitType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGitTypesIdPatch

> GitType ApiGitTypesIdPatch(ctx, id).GitType(gitType).Execute()

Updates the GitType resource.



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
    gitType := *openapiclient.NewGitType() // GitType | The updated GitType resource
    id := "id_example" // string | GitType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitTypeApi.ApiGitTypesIdPatch(context.Background(), id).GitType(gitType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitTypeApi.ApiGitTypesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGitTypesIdPatch`: GitType
    fmt.Fprintf(os.Stdout, "Response from `GitTypeApi.ApiGitTypesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | GitType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiGitTypesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitType** | [**GitType**](GitType.md) | The updated GitType resource | 


### Return type

[**GitType**](GitType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGitTypesIdPut

> GitType ApiGitTypesIdPut(ctx, id).GitType(gitType).Execute()

Replaces the GitType resource.



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
    gitType := *openapiclient.NewGitType() // GitType | The updated GitType resource
    id := "id_example" // string | GitType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitTypeApi.ApiGitTypesIdPut(context.Background(), id).GitType(gitType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitTypeApi.ApiGitTypesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGitTypesIdPut`: GitType
    fmt.Fprintf(os.Stdout, "Response from `GitTypeApi.ApiGitTypesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | GitType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiGitTypesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitType** | [**GitType**](GitType.md) | The updated GitType resource | 


### Return type

[**GitType**](GitType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGitTypesPost

> GitType ApiGitTypesPost(ctx).GitType(gitType).Execute()

Creates a GitType resource.



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
    gitType := *openapiclient.NewGitType() // GitType | The new GitType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitTypeApi.ApiGitTypesPost(context.Background()).GitType(gitType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitTypeApi.ApiGitTypesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGitTypesPost`: GitType
    fmt.Fprintf(os.Stdout, "Response from `GitTypeApi.ApiGitTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiGitTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitType** | [**GitType**](GitType.md) | The new GitType resource | 

### Return type

[**GitType**](GitType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

