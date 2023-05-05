# \TeamApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTeamsGetCollection**](TeamApi.md#ApiTeamsGetCollection) | **Get** /api/teams | Retrieves the collection of Team resources.
[**ApiTeamsIdDelete**](TeamApi.md#ApiTeamsIdDelete) | **Delete** /api/teams/{id} | Removes the Team resource.
[**ApiTeamsIdGet**](TeamApi.md#ApiTeamsIdGet) | **Get** /api/teams/{id} | Retrieves a Team resource.
[**ApiTeamsIdPatch**](TeamApi.md#ApiTeamsIdPatch) | **Patch** /api/teams/{id} | Updates the Team resource.
[**ApiTeamsIdPut**](TeamApi.md#ApiTeamsIdPut) | **Put** /api/teams/{id} | Replaces the Team resource.
[**ApiTeamsPost**](TeamApi.md#ApiTeamsPost) | **Post** /api/teams | Creates a Team resource.



## ApiTeamsGetCollection

> []TeamTeamGet ApiTeamsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ExistsOrganisation(existsOrganisation).Name(name).User(user).User2(user2).Organisation(organisation).Organisation2(organisation2).Execute()

Retrieves the collection of Team resources.



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
    existsOrganisation := true // bool |  (optional)
    name := "name_example" // string |  (optional)
    user := "user_example" // string |  (optional)
    user2 := []string{"Inner_example"} // []string |  (optional)
    organisation := "organisation_example" // string |  (optional)
    organisation2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamApi.ApiTeamsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ExistsOrganisation(existsOrganisation).Name(name).User(user).User2(user2).Organisation(organisation).Organisation2(organisation2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.ApiTeamsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamsGetCollection`: []TeamTeamGet
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.ApiTeamsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **existsOrganisation** | **bool** |  | 
 **name** | **string** |  | 
 **user** | **string** |  | 
 **user2** | **[]string** |  | 
 **organisation** | **string** |  | 
 **organisation2** | **[]string** |  | 

### Return type

[**[]TeamTeamGet**](TeamTeamGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTeamsIdDelete

> ApiTeamsIdDelete(ctx, id).Execute()

Removes the Team resource.



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
    id := "id_example" // string | Team identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TeamApi.ApiTeamsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.ApiTeamsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Team identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamsIdDeleteRequest struct via the builder pattern


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


## ApiTeamsIdGet

> TeamTeamGet ApiTeamsIdGet(ctx, id).Execute()

Retrieves a Team resource.



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
    id := "id_example" // string | Team identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamApi.ApiTeamsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.ApiTeamsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamsIdGet`: TeamTeamGet
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.ApiTeamsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Team identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamTeamGet**](TeamTeamGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTeamsIdPatch

> TeamTeamGet ApiTeamsIdPatch(ctx, id).TeamTeamPost(teamTeamPost).Execute()

Updates the Team resource.



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
    teamTeamPost := *openapiclient.NewTeamTeamPost() // TeamTeamPost | The updated Team resource
    id := "id_example" // string | Team identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamApi.ApiTeamsIdPatch(context.Background(), id).TeamTeamPost(teamTeamPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.ApiTeamsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamsIdPatch`: TeamTeamGet
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.ApiTeamsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Team identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamTeamPost** | [**TeamTeamPost**](TeamTeamPost.md) | The updated Team resource | 


### Return type

[**TeamTeamGet**](TeamTeamGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTeamsIdPut

> TeamTeamGet ApiTeamsIdPut(ctx, id).TeamTeamPost(teamTeamPost).Execute()

Replaces the Team resource.



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
    teamTeamPost := *openapiclient.NewTeamTeamPost() // TeamTeamPost | The updated Team resource
    id := "id_example" // string | Team identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamApi.ApiTeamsIdPut(context.Background(), id).TeamTeamPost(teamTeamPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.ApiTeamsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamsIdPut`: TeamTeamGet
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.ApiTeamsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Team identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamTeamPost** | [**TeamTeamPost**](TeamTeamPost.md) | The updated Team resource | 


### Return type

[**TeamTeamGet**](TeamTeamGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTeamsPost

> TeamTeamGet ApiTeamsPost(ctx).TeamTeamPost(teamTeamPost).Execute()

Creates a Team resource.



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
    teamTeamPost := *openapiclient.NewTeamTeamPost() // TeamTeamPost | The new Team resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamApi.ApiTeamsPost(context.Background()).TeamTeamPost(teamTeamPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.ApiTeamsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamsPost`: TeamTeamGet
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.ApiTeamsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamTeamPost** | [**TeamTeamPost**](TeamTeamPost.md) | The new Team resource | 

### Return type

[**TeamTeamGet**](TeamTeamGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

