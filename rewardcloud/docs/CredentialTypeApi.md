# \CredentialTypeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCredentialTypesGetCollection**](CredentialTypeApi.md#ApiCredentialTypesGetCollection) | **Get** /api/credential_types | Retrieves the collection of CredentialType resources.
[**ApiCredentialTypesIdDelete**](CredentialTypeApi.md#ApiCredentialTypesIdDelete) | **Delete** /api/credential_types/{id} | Removes the CredentialType resource.
[**ApiCredentialTypesIdGet**](CredentialTypeApi.md#ApiCredentialTypesIdGet) | **Get** /api/credential_types/{id} | Retrieves a CredentialType resource.
[**ApiCredentialTypesIdPatch**](CredentialTypeApi.md#ApiCredentialTypesIdPatch) | **Patch** /api/credential_types/{id} | Updates the CredentialType resource.
[**ApiCredentialTypesIdPut**](CredentialTypeApi.md#ApiCredentialTypesIdPut) | **Put** /api/credential_types/{id} | Replaces the CredentialType resource.
[**ApiCredentialTypesPost**](CredentialTypeApi.md#ApiCredentialTypesPost) | **Post** /api/credential_types | Creates a CredentialType resource.



## ApiCredentialTypesGetCollection

> []CredentialType ApiCredentialTypesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of CredentialType resources.



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
    resp, r, err := apiClient.CredentialTypeApi.ApiCredentialTypesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypeApi.ApiCredentialTypesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiCredentialTypesGetCollection`: []CredentialType
    fmt.Fprintf(os.Stdout, "Response from `CredentialTypeApi.ApiCredentialTypesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCredentialTypesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]CredentialType**](CredentialType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCredentialTypesIdDelete

> ApiCredentialTypesIdDelete(ctx, id).Execute()

Removes the CredentialType resource.



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
    id := "id_example" // string | CredentialType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialTypeApi.ApiCredentialTypesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypeApi.ApiCredentialTypesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | CredentialType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCredentialTypesIdDeleteRequest struct via the builder pattern


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


## ApiCredentialTypesIdGet

> CredentialType ApiCredentialTypesIdGet(ctx, id).Execute()

Retrieves a CredentialType resource.



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
    id := "id_example" // string | CredentialType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialTypeApi.ApiCredentialTypesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypeApi.ApiCredentialTypesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiCredentialTypesIdGet`: CredentialType
    fmt.Fprintf(os.Stdout, "Response from `CredentialTypeApi.ApiCredentialTypesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | CredentialType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCredentialTypesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CredentialType**](CredentialType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCredentialTypesIdPatch

> CredentialType ApiCredentialTypesIdPatch(ctx, id).CredentialType(credentialType).Execute()

Updates the CredentialType resource.



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
    id := "id_example" // string | CredentialType identifier
    credentialType := *openapiclient.NewCredentialType() // CredentialType | The updated CredentialType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialTypeApi.ApiCredentialTypesIdPatch(context.Background(), id).CredentialType(credentialType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypeApi.ApiCredentialTypesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiCredentialTypesIdPatch`: CredentialType
    fmt.Fprintf(os.Stdout, "Response from `CredentialTypeApi.ApiCredentialTypesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | CredentialType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCredentialTypesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialType** | [**CredentialType**](CredentialType.md) | The updated CredentialType resource | 

### Return type

[**CredentialType**](CredentialType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCredentialTypesIdPut

> CredentialType ApiCredentialTypesIdPut(ctx, id).CredentialType(credentialType).Execute()

Replaces the CredentialType resource.



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
    id := "id_example" // string | CredentialType identifier
    credentialType := *openapiclient.NewCredentialType() // CredentialType | The updated CredentialType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialTypeApi.ApiCredentialTypesIdPut(context.Background(), id).CredentialType(credentialType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypeApi.ApiCredentialTypesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiCredentialTypesIdPut`: CredentialType
    fmt.Fprintf(os.Stdout, "Response from `CredentialTypeApi.ApiCredentialTypesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | CredentialType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCredentialTypesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialType** | [**CredentialType**](CredentialType.md) | The updated CredentialType resource | 

### Return type

[**CredentialType**](CredentialType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCredentialTypesPost

> CredentialType ApiCredentialTypesPost(ctx).CredentialType(credentialType).Execute()

Creates a CredentialType resource.



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
    credentialType := *openapiclient.NewCredentialType() // CredentialType | The new CredentialType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialTypeApi.ApiCredentialTypesPost(context.Background()).CredentialType(credentialType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypeApi.ApiCredentialTypesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiCredentialTypesPost`: CredentialType
    fmt.Fprintf(os.Stdout, "Response from `CredentialTypeApi.ApiCredentialTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCredentialTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialType** | [**CredentialType**](CredentialType.md) | The new CredentialType resource | 

### Return type

[**CredentialType**](CredentialType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

