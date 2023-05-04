# \EnvironmentEnvVarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvironmentEnvVarsGetCollection**](EnvironmentEnvVarApi.md#ApiEnvironmentEnvVarsGetCollection) | **Get** /api/environment_env_vars | Retrieves the collection of EnvironmentEnvVar resources.
[**ApiEnvironmentEnvVarsIdDelete**](EnvironmentEnvVarApi.md#ApiEnvironmentEnvVarsIdDelete) | **Delete** /api/environment_env_vars/{id} | Removes the EnvironmentEnvVar resource.
[**ApiEnvironmentEnvVarsIdGet**](EnvironmentEnvVarApi.md#ApiEnvironmentEnvVarsIdGet) | **Get** /api/environment_env_vars/{id} | Retrieves a EnvironmentEnvVar resource.
[**ApiEnvironmentEnvVarsIdPatch**](EnvironmentEnvVarApi.md#ApiEnvironmentEnvVarsIdPatch) | **Patch** /api/environment_env_vars/{id} | Updates the EnvironmentEnvVar resource.
[**ApiEnvironmentEnvVarsIdPut**](EnvironmentEnvVarApi.md#ApiEnvironmentEnvVarsIdPut) | **Put** /api/environment_env_vars/{id} | Replaces the EnvironmentEnvVar resource.
[**ApiEnvironmentEnvVarsPost**](EnvironmentEnvVarApi.md#ApiEnvironmentEnvVarsPost) | **Post** /api/environment_env_vars | Creates a EnvironmentEnvVar resource.



## ApiEnvironmentEnvVarsGetCollection

> ApiEnvironmentEnvVarsGetCollection200Response ApiEnvironmentEnvVarsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Environment(environment).Environment2(environment2).EnvVarType(envVarType).EnvVarType2(envVarType2).Execute()

Retrieves the collection of EnvironmentEnvVar resources.



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
    environment := "environment_example" // string |  (optional)
    environment2 := []string{"Inner_example"} // []string |  (optional)
    envVarType := "envVarType_example" // string |  (optional)
    envVarType2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentEnvVarApi.ApiEnvironmentEnvVarsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Environment(environment).Environment2(environment2).EnvVarType(envVarType).EnvVarType2(envVarType2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentEnvVarApi.ApiEnvironmentEnvVarsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentEnvVarsGetCollection`: ApiEnvironmentEnvVarsGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentEnvVarApi.ApiEnvironmentEnvVarsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentEnvVarsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **environment** | **string** |  | 
 **environment2** | **[]string** |  | 
 **envVarType** | **string** |  | 
 **envVarType2** | **[]string** |  | 

### Return type

[**ApiEnvironmentEnvVarsGetCollection200Response**](ApiEnvironmentEnvVarsGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentEnvVarsIdDelete

> ApiEnvironmentEnvVarsIdDelete(ctx, id).Execute()

Removes the EnvironmentEnvVar resource.



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
    id := "id_example" // string | EnvironmentEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentEnvVarApi.ApiEnvironmentEnvVarsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentEnvVarApi.ApiEnvironmentEnvVarsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentEnvVarsIdDeleteRequest struct via the builder pattern


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


## ApiEnvironmentEnvVarsIdGet

> EnvironmentEnvVarJsonhal ApiEnvironmentEnvVarsIdGet(ctx, id).Execute()

Retrieves a EnvironmentEnvVar resource.



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
    id := "id_example" // string | EnvironmentEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentEnvVarApi.ApiEnvironmentEnvVarsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentEnvVarApi.ApiEnvironmentEnvVarsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentEnvVarsIdGet`: EnvironmentEnvVarJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentEnvVarApi.ApiEnvironmentEnvVarsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentEnvVarsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentEnvVarJsonhal**](EnvironmentEnvVarJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentEnvVarsIdPatch

> EnvironmentEnvVarJsonhal ApiEnvironmentEnvVarsIdPatch(ctx, id).EnvironmentEnvVar(environmentEnvVar).Execute()

Updates the EnvironmentEnvVar resource.



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
    id := "id_example" // string | EnvironmentEnvVar identifier
    environmentEnvVar := *openapiclient.NewEnvironmentEnvVar() // EnvironmentEnvVar | The updated EnvironmentEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentEnvVarApi.ApiEnvironmentEnvVarsIdPatch(context.Background(), id).EnvironmentEnvVar(environmentEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentEnvVarApi.ApiEnvironmentEnvVarsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentEnvVarsIdPatch`: EnvironmentEnvVarJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentEnvVarApi.ApiEnvironmentEnvVarsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentEnvVarsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEnvVar** | [**EnvironmentEnvVar**](EnvironmentEnvVar.md) | The updated EnvironmentEnvVar resource | 

### Return type

[**EnvironmentEnvVarJsonhal**](EnvironmentEnvVarJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentEnvVarsIdPut

> EnvironmentEnvVarJsonhal ApiEnvironmentEnvVarsIdPut(ctx, id).EnvironmentEnvVarJsonhal(environmentEnvVarJsonhal).Execute()

Replaces the EnvironmentEnvVar resource.



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
    id := "id_example" // string | EnvironmentEnvVar identifier
    environmentEnvVarJsonhal := *openapiclient.NewEnvironmentEnvVarJsonhal() // EnvironmentEnvVarJsonhal | The updated EnvironmentEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentEnvVarApi.ApiEnvironmentEnvVarsIdPut(context.Background(), id).EnvironmentEnvVarJsonhal(environmentEnvVarJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentEnvVarApi.ApiEnvironmentEnvVarsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentEnvVarsIdPut`: EnvironmentEnvVarJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentEnvVarApi.ApiEnvironmentEnvVarsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentEnvVarsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEnvVarJsonhal** | [**EnvironmentEnvVarJsonhal**](EnvironmentEnvVarJsonhal.md) | The updated EnvironmentEnvVar resource | 

### Return type

[**EnvironmentEnvVarJsonhal**](EnvironmentEnvVarJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentEnvVarsPost

> EnvironmentEnvVarJsonhal ApiEnvironmentEnvVarsPost(ctx).EnvironmentEnvVarJsonhal(environmentEnvVarJsonhal).Execute()

Creates a EnvironmentEnvVar resource.



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
    environmentEnvVarJsonhal := *openapiclient.NewEnvironmentEnvVarJsonhal() // EnvironmentEnvVarJsonhal | The new EnvironmentEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentEnvVarApi.ApiEnvironmentEnvVarsPost(context.Background()).EnvironmentEnvVarJsonhal(environmentEnvVarJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentEnvVarApi.ApiEnvironmentEnvVarsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentEnvVarsPost`: EnvironmentEnvVarJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentEnvVarApi.ApiEnvironmentEnvVarsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentEnvVarsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentEnvVarJsonhal** | [**EnvironmentEnvVarJsonhal**](EnvironmentEnvVarJsonhal.md) | The new EnvironmentEnvVar resource | 

### Return type

[**EnvironmentEnvVarJsonhal**](EnvironmentEnvVarJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

