# \ProjectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectsGetCollection**](ProjectApi.md#ApiProjectsGetCollection) | **Get** /api/projects | Retrieves the collection of Project resources.
[**ApiProjectsIdDelete**](ProjectApi.md#ApiProjectsIdDelete) | **Delete** /api/projects/{id} | Removes the Project resource.
[**ApiProjectsIdGet**](ProjectApi.md#ApiProjectsIdGet) | **Get** /api/projects/{id} | Retrieves a Project resource.
[**ApiProjectsIdPatch**](ProjectApi.md#ApiProjectsIdPatch) | **Patch** /api/projects/{id} | Updates the Project resource.
[**ApiProjectsIdPut**](ProjectApi.md#ApiProjectsIdPut) | **Put** /api/projects/{id} | Replaces the Project resource.
[**ApiProjectsIdcostsGet**](ProjectApi.md#ApiProjectsIdcostsGet) | **Get** /api/projects/{id}/costs | Retrieves a Project resource.
[**ApiProjectsPost**](ProjectApi.md#ApiProjectsPost) | **Post** /api/projects | Creates a Project resource.



## ApiProjectsGetCollection

> ApiProjectsGetCollection200Response ApiProjectsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).TeamOrganisationId(teamOrganisationId).TeamOrganisationId2(teamOrganisationId2).State(state).State2(state2).Git(git).Git2(git2).ServiceAccount(serviceAccount).ServiceAccount2(serviceAccount2).Team(team).Team2(team2).ProjectTypeVersion(projectTypeVersion).ProjectTypeVersion2(projectTypeVersion2).OrderUpdatedAt(orderUpdatedAt).Execute()

Retrieves the collection of Project resources.



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
    teamOrganisationId := int32(56) // int32 |  (optional)
    teamOrganisationId2 := []int32{int32(123)} // []int32 |  (optional)
    state := "state_example" // string |  (optional)
    state2 := []string{"Inner_example"} // []string |  (optional)
    git := "git_example" // string |  (optional)
    git2 := []string{"Inner_example"} // []string |  (optional)
    serviceAccount := "serviceAccount_example" // string |  (optional)
    serviceAccount2 := []string{"Inner_example"} // []string |  (optional)
    team := "team_example" // string |  (optional)
    team2 := []string{"Inner_example"} // []string |  (optional)
    projectTypeVersion := "projectTypeVersion_example" // string |  (optional)
    projectTypeVersion2 := []string{"Inner_example"} // []string |  (optional)
    orderUpdatedAt := "orderUpdatedAt_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ApiProjectsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).TeamOrganisationId(teamOrganisationId).TeamOrganisationId2(teamOrganisationId2).State(state).State2(state2).Git(git).Git2(git2).ServiceAccount(serviceAccount).ServiceAccount2(serviceAccount2).Team(team).Team2(team2).ProjectTypeVersion(projectTypeVersion).ProjectTypeVersion2(projectTypeVersion2).OrderUpdatedAt(orderUpdatedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ApiProjectsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsGetCollection`: ApiProjectsGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ApiProjectsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **teamOrganisationId** | **int32** |  | 
 **teamOrganisationId2** | **[]int32** |  | 
 **state** | **string** |  | 
 **state2** | **[]string** |  | 
 **git** | **string** |  | 
 **git2** | **[]string** |  | 
 **serviceAccount** | **string** |  | 
 **serviceAccount2** | **[]string** |  | 
 **team** | **string** |  | 
 **team2** | **[]string** |  | 
 **projectTypeVersion** | **string** |  | 
 **projectTypeVersion2** | **[]string** |  | 
 **orderUpdatedAt** | **string** |  | 

### Return type

[**ApiProjectsGetCollection200Response**](ApiProjectsGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsIdDelete

> ApiProjectsIdDelete(ctx, id).Execute()

Removes the Project resource.



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
    id := "id_example" // string | Project identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ApiProjectsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ApiProjectsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdDeleteRequest struct via the builder pattern


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


## ApiProjectsIdGet

> ProjectJsonhalProjectGet ApiProjectsIdGet(ctx, id).Execute()

Retrieves a Project resource.



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
    id := "id_example" // string | Project identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ApiProjectsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ApiProjectsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdGet`: ProjectJsonhalProjectGet
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ApiProjectsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectJsonhalProjectGet**](ProjectJsonhalProjectGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsIdPatch

> ProjectJsonhalProjectGet ApiProjectsIdPatch(ctx, id).ProjectProjectPost(projectProjectPost).Execute()

Updates the Project resource.



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
    id := "id_example" // string | Project identifier
    projectProjectPost := *openapiclient.NewProjectProjectPost() // ProjectProjectPost | The updated Project resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ApiProjectsIdPatch(context.Background(), id).ProjectProjectPost(projectProjectPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ApiProjectsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdPatch`: ProjectJsonhalProjectGet
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ApiProjectsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectProjectPost** | [**ProjectProjectPost**](ProjectProjectPost.md) | The updated Project resource | 

### Return type

[**ProjectJsonhalProjectGet**](ProjectJsonhalProjectGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsIdPut

> ProjectJsonhalProjectGet ApiProjectsIdPut(ctx, id).ProjectJsonhalProjectPost(projectJsonhalProjectPost).Execute()

Replaces the Project resource.



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
    id := "id_example" // string | Project identifier
    projectJsonhalProjectPost := *openapiclient.NewProjectJsonhalProjectPost() // ProjectJsonhalProjectPost | The updated Project resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ApiProjectsIdPut(context.Background(), id).ProjectJsonhalProjectPost(projectJsonhalProjectPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ApiProjectsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdPut`: ProjectJsonhalProjectGet
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ApiProjectsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectJsonhalProjectPost** | [**ProjectJsonhalProjectPost**](ProjectJsonhalProjectPost.md) | The updated Project resource | 

### Return type

[**ProjectJsonhalProjectGet**](ProjectJsonhalProjectGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsIdcostsGet

> ProjectJsonhalProjectGet ApiProjectsIdcostsGet(ctx, id).Execute()

Retrieves a Project resource.



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
    id := "id_example" // string | Project identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ApiProjectsIdcostsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ApiProjectsIdcostsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdcostsGet`: ProjectJsonhalProjectGet
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ApiProjectsIdcostsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdcostsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectJsonhalProjectGet**](ProjectJsonhalProjectGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsPost

> ProjectJsonhalProjectGet ApiProjectsPost(ctx).ProjectJsonhalProjectPost(projectJsonhalProjectPost).Execute()

Creates a Project resource.



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
    projectJsonhalProjectPost := *openapiclient.NewProjectJsonhalProjectPost() // ProjectJsonhalProjectPost | The new Project resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ApiProjectsPost(context.Background()).ProjectJsonhalProjectPost(projectJsonhalProjectPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ApiProjectsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsPost`: ProjectJsonhalProjectGet
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ApiProjectsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectJsonhalProjectPost** | [**ProjectJsonhalProjectPost**](ProjectJsonhalProjectPost.md) | The new Project resource | 

### Return type

[**ProjectJsonhalProjectGet**](ProjectJsonhalProjectGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

