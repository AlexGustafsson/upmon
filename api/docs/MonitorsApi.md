# \MonitorsApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OriginsOriginIdServicesServiceIdMonitorsGet**](MonitorsApi.md#OriginsOriginIdServicesServiceIdMonitorsGet) | **Get** /origins/{originId}/services/{serviceId}/monitors | Retrieve all monitors for a service of an origin
[**OriginsOriginIdServicesServiceIdMonitorsMonitorIdGet**](MonitorsApi.md#OriginsOriginIdServicesServiceIdMonitorsMonitorIdGet) | **Get** /origins/{originId}/services/{serviceId}/monitors/{monitorId} | Retrieve a monitor of a service from an origin
[**OriginsOriginIdServicesServiceIdMonitorsMonitorIdStatusGet**](MonitorsApi.md#OriginsOriginIdServicesServiceIdMonitorsMonitorIdStatusGet) | **Get** /origins/{originId}/services/{serviceId}/monitors/{monitorId}/status | Retrieve the status of a monitor



## OriginsOriginIdServicesServiceIdMonitorsGet

> Monitors OriginsOriginIdServicesServiceIdMonitorsGet(ctx, originId, serviceId).Execute()

Retrieve all monitors for a service of an origin

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
    resp, r, err := api_client.MonitorsApi.OriginsOriginIdServicesServiceIdMonitorsGet(context.Background(), originId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.OriginsOriginIdServicesServiceIdMonitorsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OriginsOriginIdServicesServiceIdMonitorsGet`: Monitors
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.OriginsOriginIdServicesServiceIdMonitorsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originId** | **string** | The id of the target origin | 
**serviceId** | **string** | The id of the target service | 

### Other Parameters

Other parameters are passed through a pointer to a apiOriginsOriginIdServicesServiceIdMonitorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Monitors**](Monitors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OriginsOriginIdServicesServiceIdMonitorsMonitorIdGet

> Monitor OriginsOriginIdServicesServiceIdMonitorsMonitorIdGet(ctx, originId, serviceId, monitorId).Execute()

Retrieve a monitor of a service from an origin

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
    monitorId := "monitorId_example" // string | The id of the target monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.OriginsOriginIdServicesServiceIdMonitorsMonitorIdGet(context.Background(), originId, serviceId, monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.OriginsOriginIdServicesServiceIdMonitorsMonitorIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OriginsOriginIdServicesServiceIdMonitorsMonitorIdGet`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.OriginsOriginIdServicesServiceIdMonitorsMonitorIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originId** | **string** | The id of the target origin | 
**serviceId** | **string** | The id of the target service | 
**monitorId** | **string** | The id of the target monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiOriginsOriginIdServicesServiceIdMonitorsMonitorIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Monitor**](Monitor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OriginsOriginIdServicesServiceIdMonitorsMonitorIdStatusGet

> MonitorStatus OriginsOriginIdServicesServiceIdMonitorsMonitorIdStatusGet(ctx, originId, serviceId, monitorId).Execute()

Retrieve the status of a monitor

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
    monitorId := "monitorId_example" // string | The id of the target monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.OriginsOriginIdServicesServiceIdMonitorsMonitorIdStatusGet(context.Background(), originId, serviceId, monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.OriginsOriginIdServicesServiceIdMonitorsMonitorIdStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OriginsOriginIdServicesServiceIdMonitorsMonitorIdStatusGet`: MonitorStatus
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.OriginsOriginIdServicesServiceIdMonitorsMonitorIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originId** | **string** | The id of the target origin | 
**serviceId** | **string** | The id of the target service | 
**monitorId** | **string** | The id of the target monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiOriginsOriginIdServicesServiceIdMonitorsMonitorIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MonitorStatus**](MonitorStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

