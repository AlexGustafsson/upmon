# \ServicesApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OriginsOriginIdServicesGet**](ServicesApi.md#OriginsOriginIdServicesGet) | **Get** /origins/{originId}/services | Retrieve all monitored services for an origin
[**OriginsOriginIdServicesServiceIdGet**](ServicesApi.md#OriginsOriginIdServicesServiceIdGet) | **Get** /origins/{originId}/services/{serviceId} | Retrieve a service
[**OriginsOriginIdServicesServiceIdStatusGet**](ServicesApi.md#OriginsOriginIdServicesServiceIdStatusGet) | **Get** /origins/{originId}/services/{serviceId}/status | Retrieve the status of a service
[**ServicesGet**](ServicesApi.md#ServicesGet) | **Get** /services | Retrieve all monitored services



## OriginsOriginIdServicesGet

> Services OriginsOriginIdServicesGet(ctx, originId).Execute()

Retrieve all monitored services for an origin

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
    originId := "originId_example" // string | The id of the target origin

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.OriginsOriginIdServicesGet(context.Background(), originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.OriginsOriginIdServicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OriginsOriginIdServicesGet`: Services
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.OriginsOriginIdServicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originId** | **string** | The id of the target origin | 

### Other Parameters

Other parameters are passed through a pointer to a apiOriginsOriginIdServicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Services**](Services.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OriginsOriginIdServicesServiceIdGet

> Service OriginsOriginIdServicesServiceIdGet(ctx, originId, serviceId).Execute()

Retrieve a service

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
    originId := "originId_example" // string | The id of the target origin
    serviceId := "serviceId_example" // string | The id of the target service

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.OriginsOriginIdServicesServiceIdGet(context.Background(), originId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.OriginsOriginIdServicesServiceIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OriginsOriginIdServicesServiceIdGet`: Service
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.OriginsOriginIdServicesServiceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originId** | **string** | The id of the target origin | 
**serviceId** | **string** | The id of the target service | 

### Other Parameters

Other parameters are passed through a pointer to a apiOriginsOriginIdServicesServiceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OriginsOriginIdServicesServiceIdStatusGet

> ServiceStatus OriginsOriginIdServicesServiceIdStatusGet(ctx, originId, serviceId).Execute()

Retrieve the status of a service

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
    originId := "originId_example" // string | The id of the target origin
    serviceId := "serviceId_example" // string | The id of the target service

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.OriginsOriginIdServicesServiceIdStatusGet(context.Background(), originId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.OriginsOriginIdServicesServiceIdStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OriginsOriginIdServicesServiceIdStatusGet`: ServiceStatus
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.OriginsOriginIdServicesServiceIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originId** | **string** | The id of the target origin | 
**serviceId** | **string** | The id of the target service | 

### Other Parameters

Other parameters are passed through a pointer to a apiOriginsOriginIdServicesServiceIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceStatus**](ServiceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesGet

> Services ServicesGet(ctx).Execute()

Retrieve all monitored services

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesGet`: Services
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesGetRequest struct via the builder pattern


### Return type

[**Services**](Services.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

