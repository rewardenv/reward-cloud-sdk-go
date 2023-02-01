# \GitApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiGitsGetCollection**](GitApi.md#ApiGitsGetCollection) | **Get** /api/gits | Retrieves the collection of Git resources.
[**ApiGitsIdDelete**](GitApi.md#ApiGitsIdDelete) | **Delete** /api/gits/{id} | Removes the Git resource.
[**ApiGitsIdGet**](GitApi.md#ApiGitsIdGet) | **Get** /api/gits/{id} | Retrieves a Git resource.
[**ApiGitsIdPatch**](GitApi.md#ApiGitsIdPatch) | **Patch** /api/gits/{id} | Updates the Git resource.
[**ApiGitsIdPut**](GitApi.md#ApiGitsIdPut) | **Put** /api/gits/{id} | Replaces the Git resource.
[**ApiGitsPost**](GitApi.md#ApiGitsPost) | **Post** /api/gits | Creates a Git resource.



## ApiGitsGetCollection

> []Git ApiGitsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Project(project).Project2(project2).GitType(gitType).GitType2(gitType2).CredentialType(credentialType).CredentialType2(credentialType2).Execute()

Retrieves the collection of Git resources.



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
    project := "project_example" // string |  (optional)
    project2 := []string{"Inner_example"} // []string |  (optional)
    gitType := "gitType_example" // string |  (optional)
    gitType2 := []string{"Inner_example"} // []string |  (optional)
    credentialType := "credentialType_example" // string |  (optional)
    credentialType2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.ApiGitsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Project(project).Project2(project2).GitType(gitType).GitType2(gitType2).CredentialType(credentialType).CredentialType2(credentialType2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.ApiGitsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGitsGetCollection`: []Git
    fmt.Fprintf(os.Stdout, "Response from `GitApi.ApiGitsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiGitsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **project** | **string** |  | 
 **project2** | **[]string** |  | 
 **gitType** | **string** |  | 
 **gitType2** | **[]string** |  | 
 **credentialType** | **string** |  | 
 **credentialType2** | **[]string** |  | 

### Return type

[**[]Git**](Git.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGitsIdDelete

> ApiGitsIdDelete(ctx, id).Execute()

Removes the Git resource.



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
    id := "id_example" // string | Git identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.ApiGitsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.ApiGitsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Git identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiGitsIdDeleteRequest struct via the builder pattern


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


## ApiGitsIdGet

> Git ApiGitsIdGet(ctx, id).Execute()

Retrieves a Git resource.



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
    id := "id_example" // string | Git identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.ApiGitsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.ApiGitsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGitsIdGet`: Git
    fmt.Fprintf(os.Stdout, "Response from `GitApi.ApiGitsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Git identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiGitsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Git**](Git.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGitsIdPatch

> Git ApiGitsIdPatch(ctx, id).Git(git).Execute()

Updates the Git resource.



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
    id := "id_example" // string | Git identifier
    git := *openapiclient.NewGit() // Git | The updated Git resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.ApiGitsIdPatch(context.Background(), id).Git(git).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.ApiGitsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGitsIdPatch`: Git
    fmt.Fprintf(os.Stdout, "Response from `GitApi.ApiGitsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Git identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiGitsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **git** | [**Git**](Git.md) | The updated Git resource | 

### Return type

[**Git**](Git.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGitsIdPut

> Git ApiGitsIdPut(ctx, id).Git(git).Execute()

Replaces the Git resource.



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
    id := "id_example" // string | Git identifier
    git := *openapiclient.NewGit() // Git | The updated Git resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.ApiGitsIdPut(context.Background(), id).Git(git).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.ApiGitsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGitsIdPut`: Git
    fmt.Fprintf(os.Stdout, "Response from `GitApi.ApiGitsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Git identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiGitsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **git** | [**Git**](Git.md) | The updated Git resource | 

### Return type

[**Git**](Git.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGitsPost

> Git ApiGitsPost(ctx).Git(git).Execute()

Creates a Git resource.



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
    git := *openapiclient.NewGit() // Git | The new Git resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.ApiGitsPost(context.Background()).Git(git).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.ApiGitsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGitsPost`: Git
    fmt.Fprintf(os.Stdout, "Response from `GitApi.ApiGitsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiGitsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **git** | [**Git**](Git.md) | The new Git resource | 

### Return type

[**Git**](Git.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

