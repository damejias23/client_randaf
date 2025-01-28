<<<<<<< HEAD
# client_randaf
=======
# Go API client for randaf

RANDAF Event Exposure Service.  
© 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.2.3
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import randaf "github.com/damejias23/client_randaf"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), randaf.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), randaf.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), randaf.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), randaf.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/nrandaf-evts/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndividualSubscriptionDocumentApi* | [**DeleteSubscription**](docs/IndividualSubscriptionDocumentApi.md#deletesubscription) | **Delete** /subscriptions/{subscriptionId} | Nrandaf_EventExposure Unsubscribe service Operation
*SubscriptionsCollectionCollectionApi* | [**CreateSubscription**](docs/SubscriptionsCollectionCollectionApi.md#createsubscription) | **Post** /subscriptions | Nrandaf_EventExposure Subscribe service Operation


## Documentation For Models

 - [GnbInformation](docs/GnbInformation.md)
 - [PhyRANLevelInformation](docs/PhyRANLevelInformation.md)
 - [RandafCreateEventSubscription](docs/RandafCreateEventSubscription.md)
 - [RandafCreatedEventSubscription](docs/RandafCreatedEventSubscription.md)
 - [RandafEvent](docs/RandafEvent.md)
 - [RandafEventMode](docs/RandafEventMode.md)
 - [RandafEventNotification](docs/RandafEventNotification.md)
 - [RandafEventReport](docs/RandafEventReport.md)
 - [RandafEventState](docs/RandafEventState.md)
 - [RandafEventSubsSyncInfo](docs/RandafEventSubsSyncInfo.md)
 - [RandafEventSubscription](docs/RandafEventSubscription.md)
 - [RandafEventSubscriptionInfo](docs/RandafEventSubscriptionInfo.md)
 - [RandafEventTrigger](docs/RandafEventTrigger.md)
 - [RandafEventTriggerAnyOf](docs/RandafEventTriggerAnyOf.md)
 - [RandafEventType](docs/RandafEventType.md)
 - [RandafEventTypeAnyOf](docs/RandafEventTypeAnyOf.md)
 - [TargetUeInformation](docs/TargetUeInformation.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **nrandaf-evts**: Access to the Nrandaf_EventExposure API

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



>>>>>>> master
