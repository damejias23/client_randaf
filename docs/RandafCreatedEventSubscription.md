# RandafCreatedEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**RandafEventSubscription**](RandafEventSubscription.md) |  | 
**SubscriptionId** | **string** | String providing an URI formatted according to RFC 3986. | 
**ReportEvent** | Pointer to [**RandafEventReport**](RandafEventReport.md) |  | [optional] 

## Methods

### NewRandafCreatedEventSubscription

`func NewRandafCreatedEventSubscription(subscription RandafEventSubscription, subscriptionId string, ) *RandafCreatedEventSubscription`

NewRandafCreatedEventSubscription instantiates a new RandafCreatedEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandafCreatedEventSubscriptionWithDefaults

`func NewRandafCreatedEventSubscriptionWithDefaults() *RandafCreatedEventSubscription`

NewRandafCreatedEventSubscriptionWithDefaults instantiates a new RandafCreatedEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *RandafCreatedEventSubscription) GetSubscription() RandafEventSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *RandafCreatedEventSubscription) GetSubscriptionOk() (*RandafEventSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *RandafCreatedEventSubscription) SetSubscription(v RandafEventSubscription)`

SetSubscription sets Subscription field to given value.


### GetSubscriptionId

`func (o *RandafCreatedEventSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RandafCreatedEventSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RandafCreatedEventSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetReportEvent

`func (o *RandafCreatedEventSubscription) GetReportEvent() RandafEventReport`

GetReportEvent returns the ReportEvent field if non-nil, zero value otherwise.

### GetReportEventOk

`func (o *RandafCreatedEventSubscription) GetReportEventOk() (*RandafEventReport, bool)`

GetReportEventOk returns a tuple with the ReportEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportEvent

`func (o *RandafCreatedEventSubscription) SetReportEvent(v RandafEventReport)`

SetReportEvent sets ReportEvent field to given value.

### HasReportEvent

`func (o *RandafCreatedEventSubscription) HasReportEvent() bool`

HasReportEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


