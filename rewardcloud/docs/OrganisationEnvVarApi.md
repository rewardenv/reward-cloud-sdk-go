# \OrganisationEnvVarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiOrganisationEnvVarsGetCollection**](OrganisationEnvVarApi.md#ApiOrganisationEnvVarsGetCollection) | **Get** /api/organisation_env_vars | Retrieves the collection of OrganisationEnvVar resources.
[**ApiOrganisationEnvVarsIdDelete**](OrganisationEnvVarApi.md#ApiOrganisationEnvVarsIdDelete) | **Delete** /api/organisation_env_vars/{id} | Removes the OrganisationEnvVar resource.
[**ApiOrganisationEnvVarsIdGet**](OrganisationEnvVarApi.md#ApiOrganisationEnvVarsIdGet) | **Get** /api/organisation_env_vars/{id} | Retrieves a OrganisationEnvVar resource.
[**ApiOrganisationEnvVarsIdPatch**](OrganisationEnvVarApi.md#ApiOrganisationEnvVarsIdPatch) | **Patch** /api/organisation_env_vars/{id} | Updates the OrganisationEnvVar resource.
[**ApiOrganisationEnvVarsIdPut**](OrganisationEnvVarApi.md#ApiOrganisationEnvVarsIdPut) | **Put** /api/organisation_env_vars/{id} | Replaces the OrganisationEnvVar resource.
[**ApiOrganisationEnvVarsPost**](OrganisationEnvVarApi.md#ApiOrganisationEnvVarsPost) | **Post** /api/organisation_env_vars | Creates a OrganisationEnvVar resource.



## ApiOrganisationEnvVarsGetCollection

> []OrganisationEnvVar ApiOrganisationEnvVarsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Organisation(organisation).Organisation2(organisation2).EnvVarType(envVarType).EnvVarType2(envVarType2).Execute()

Retrieves the collection of OrganisationEnvVar resources.



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
    organisation := "organisation_example" // string |  (optional)
    organisation2 := []string{"Inner_example"} // []string |  (optional)
    envVarType := "envVarType_example" // string |  (optional)
    envVarType2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationEnvVarApi.ApiOrganisationEnvVarsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Organisation(organisation).Organisation2(organisation2).EnvVarType(envVarType).EnvVarType2(envVarType2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationEnvVarApi.ApiOrganisationEnvVarsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiOrganisationEnvVarsGetCollection`: []OrganisationEnvVar
    fmt.Fprintf(os.Stdout, "Response from `OrganisationEnvVarApi.ApiOrganisationEnvVarsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationEnvVarsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **organisation** | **string** |  | 
 **organisation2** | **[]string** |  | 
 **envVarType** | **string** |  | 
 **envVarType2** | **[]string** |  | 

### Return type

[**[]OrganisationEnvVar**](OrganisationEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOrganisationEnvVarsIdDelete

> ApiOrganisationEnvVarsIdDelete(ctx, id).Execute()

Removes the OrganisationEnvVar resource.



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
    id := "id_example" // string | OrganisationEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganisationEnvVarApi.ApiOrganisationEnvVarsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationEnvVarApi.ApiOrganisationEnvVarsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | OrganisationEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationEnvVarsIdDeleteRequest struct via the builder pattern


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


## ApiOrganisationEnvVarsIdGet

> OrganisationEnvVar ApiOrganisationEnvVarsIdGet(ctx, id).Execute()

Retrieves a OrganisationEnvVar resource.



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
    id := "id_example" // string | OrganisationEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationEnvVarApi.ApiOrganisationEnvVarsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationEnvVarApi.ApiOrganisationEnvVarsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiOrganisationEnvVarsIdGet`: OrganisationEnvVar
    fmt.Fprintf(os.Stdout, "Response from `OrganisationEnvVarApi.ApiOrganisationEnvVarsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | OrganisationEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationEnvVarsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganisationEnvVar**](OrganisationEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOrganisationEnvVarsIdPatch

> OrganisationEnvVar ApiOrganisationEnvVarsIdPatch(ctx, id).OrganisationEnvVar(organisationEnvVar).Execute()

Updates the OrganisationEnvVar resource.



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
    organisationEnvVar := *openapiclient.NewOrganisationEnvVar() // OrganisationEnvVar | The updated OrganisationEnvVar resource
    id := "id_example" // string | OrganisationEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationEnvVarApi.ApiOrganisationEnvVarsIdPatch(context.Background(), id).OrganisationEnvVar(organisationEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationEnvVarApi.ApiOrganisationEnvVarsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiOrganisationEnvVarsIdPatch`: OrganisationEnvVar
    fmt.Fprintf(os.Stdout, "Response from `OrganisationEnvVarApi.ApiOrganisationEnvVarsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | OrganisationEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationEnvVarsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organisationEnvVar** | [**OrganisationEnvVar**](OrganisationEnvVar.md) | The updated OrganisationEnvVar resource | 


### Return type

[**OrganisationEnvVar**](OrganisationEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOrganisationEnvVarsIdPut

> OrganisationEnvVar ApiOrganisationEnvVarsIdPut(ctx, id).OrganisationEnvVar(organisationEnvVar).Execute()

Replaces the OrganisationEnvVar resource.



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
    organisationEnvVar := *openapiclient.NewOrganisationEnvVar() // OrganisationEnvVar | The updated OrganisationEnvVar resource
    id := "id_example" // string | OrganisationEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationEnvVarApi.ApiOrganisationEnvVarsIdPut(context.Background(), id).OrganisationEnvVar(organisationEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationEnvVarApi.ApiOrganisationEnvVarsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiOrganisationEnvVarsIdPut`: OrganisationEnvVar
    fmt.Fprintf(os.Stdout, "Response from `OrganisationEnvVarApi.ApiOrganisationEnvVarsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | OrganisationEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationEnvVarsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organisationEnvVar** | [**OrganisationEnvVar**](OrganisationEnvVar.md) | The updated OrganisationEnvVar resource | 


### Return type

[**OrganisationEnvVar**](OrganisationEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOrganisationEnvVarsPost

> OrganisationEnvVar ApiOrganisationEnvVarsPost(ctx).OrganisationEnvVar(organisationEnvVar).Execute()

Creates a OrganisationEnvVar resource.



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
    organisationEnvVar := *openapiclient.NewOrganisationEnvVar() // OrganisationEnvVar | The new OrganisationEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationEnvVarApi.ApiOrganisationEnvVarsPost(context.Background()).OrganisationEnvVar(organisationEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationEnvVarApi.ApiOrganisationEnvVarsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiOrganisationEnvVarsPost`: OrganisationEnvVar
    fmt.Fprintf(os.Stdout, "Response from `OrganisationEnvVarApi.ApiOrganisationEnvVarsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiOrganisationEnvVarsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organisationEnvVar** | [**OrganisationEnvVar**](OrganisationEnvVar.md) | The new OrganisationEnvVar resource | 

### Return type

[**OrganisationEnvVar**](OrganisationEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

