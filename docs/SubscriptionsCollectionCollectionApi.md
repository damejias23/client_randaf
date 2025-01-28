# \SubscriptionsCollectionCollectionApi

All URIs are relative to *https://example.com/nrandaf-evts/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionsCollectionCollectionApi.md#CreateSubscription) | **Post** /subscriptions | Nrandaf_EventExposure Subscribe service Operation



## CreateSubscription

> RandafCreatedEventSubscription CreateSubscription(ctx).RandafCreateEventSubscription(randafCreateEventSubscription).Execute()

Nrandaf_EventExposure Subscribe service Operation

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
    randafCreateEventSubscription := *openapiclient.NewRandafCreateEventSubscription(*openapiclient.NewRandafEventSubscription(*openapiclient.NewRandafEvent(*openapiclient.NewRandafEventType()), "EventNotifyUri_example")) // RandafCreateEventSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionCollectionApi.CreateSubscription(context.Background()).RandafCreateEventSubscription(randafCreateEventSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionCollectionApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: RandafCreatedEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionCollectionApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **randafCreateEventSubscription** | [**RandafCreateEventSubscription**](RandafCreateEventSubscription.md) |  | 

### Return type

[**RandafCreatedEventSubscription**](RandafCreatedEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

