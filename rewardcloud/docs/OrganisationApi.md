# \OrganisationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiOrganisationsGetCollection**](OrganisationApi.md#ApiOrganisationsGetCollection) | **Get** /api/organisations | Retrieves the collection of Organisation resources.
[**ApiOrganisationsIdDelete**](OrganisationApi.md#ApiOrganisationsIdDelete) | **Delete** /api/organisations/{id} | Removes the Organisation resource.
[**ApiOrganisationsIdGet**](OrganisationApi.md#ApiOrganisationsIdGet) | **Get** /api/organisations/{id} | Retrieves a Organisation resource.
[**ApiOrganisationsIdPatch**](OrganisationApi.md#ApiOrganisationsIdPatch) | **Patch** /api/organisations/{id} | Updates the Organisation resource.
[**ApiOrganisationsIdPut**](OrganisationApi.md#ApiOrganisationsIdPut) | **Put** /api/organisations/{id} | Replaces the Organisation resource.
[**ApiOrganisationsPost**](OrganisationApi.md#ApiOrganisationsPost) | **Post** /api/organisations | Creates a Organisation resource.



## ApiOrganisationsGetCollection

> []OrganisationOrganisationGet ApiOrganisationsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of Organisation resources.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationApi.ApiOrganisationsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.ApiOrganisationsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiOrganisationsGetCollection`: []OrganisationOrganisationGet
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.ApiOrganisationsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]OrganisationOrganisationGet**](OrganisationOrganisationGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOrganisationsIdDelete

> ApiOrganisationsIdDelete(ctx, id).Execute()

Removes the Organisation resource.



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
    id := "id_example" // string | Organisation identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationApi.ApiOrganisationsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.ApiOrganisationsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Organisation identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationsIdDeleteRequest struct via the builder pattern


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


## ApiOrganisationsIdGet

> OrganisationOrganisationGet ApiOrganisationsIdGet(ctx, id).Execute()

Retrieves a Organisation resource.



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
    id := "id_example" // string | Organisation identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationApi.ApiOrganisationsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.ApiOrganisationsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiOrganisationsIdGet`: OrganisationOrganisationGet
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.ApiOrganisationsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Organisation identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganisationOrganisationGet**](OrganisationOrganisationGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOrganisationsIdPatch

> OrganisationOrganisationGet ApiOrganisationsIdPatch(ctx, id).OrganisationOrganisationPost(organisationOrganisationPost).Execute()

Updates the Organisation resource.



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
    id := "id_example" // string | Organisation identifier
    organisationOrganisationPost := *openapiclient.NewOrganisationOrganisationPost() // OrganisationOrganisationPost | The updated Organisation resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationApi.ApiOrganisationsIdPatch(context.Background(), id).OrganisationOrganisationPost(organisationOrganisationPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.ApiOrganisationsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiOrganisationsIdPatch`: OrganisationOrganisationGet
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.ApiOrganisationsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Organisation identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organisationOrganisationPost** | [**OrganisationOrganisationPost**](OrganisationOrganisationPost.md) | The updated Organisation resource | 

### Return type

[**OrganisationOrganisationGet**](OrganisationOrganisationGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOrganisationsIdPut

> OrganisationOrganisationGet ApiOrganisationsIdPut(ctx, id).OrganisationOrganisationPost(organisationOrganisationPost).Execute()

Replaces the Organisation resource.



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
    id := "id_example" // string | Organisation identifier
    organisationOrganisationPost := *openapiclient.NewOrganisationOrganisationPost() // OrganisationOrganisationPost | The updated Organisation resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationApi.ApiOrganisationsIdPut(context.Background(), id).OrganisationOrganisationPost(organisationOrganisationPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.ApiOrganisationsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiOrganisationsIdPut`: OrganisationOrganisationGet
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.ApiOrganisationsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Organisation identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organisationOrganisationPost** | [**OrganisationOrganisationPost**](OrganisationOrganisationPost.md) | The updated Organisation resource | 

### Return type

[**OrganisationOrganisationGet**](OrganisationOrganisationGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOrganisationsPost

> OrganisationOrganisationGet ApiOrganisationsPost(ctx).OrganisationOrganisationPost(organisationOrganisationPost).Execute()

Creates a Organisation resource.



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
    organisationOrganisationPost := *openapiclient.NewOrganisationOrganisationPost() // OrganisationOrganisationPost | The new Organisation resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationApi.ApiOrganisationsPost(context.Background()).OrganisationOrganisationPost(organisationOrganisationPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.ApiOrganisationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiOrganisationsPost`: OrganisationOrganisationGet
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.ApiOrganisationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organisationOrganisationPost** | [**OrganisationOrganisationPost**](OrganisationOrganisationPost.md) | The new Organisation resource | 

### Return type

[**OrganisationOrganisationGet**](OrganisationOrganisationGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

