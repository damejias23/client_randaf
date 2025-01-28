# RandafEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventRequest** | [**RandafEvent**](RandafEvent.md) |  | 
**EventNotifyUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**Options** | Pointer to [**RandafEventMode**](RandafEventMode.md) |  | [optional] 
**TgtUe** | Pointer to [**TargetUeInformation**](TargetUeInformation.md) |  | [optional] 

## Methods

### NewRandafEventSubscription

`func NewRandafEventSubscription(eventRequest RandafEvent, eventNotifyUri string, ) *RandafEventSubscription`

NewRandafEventSubscription instantiates a new RandafEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandafEventSubscriptionWithDefaults

`func NewRandafEventSubscriptionWithDefaults() *RandafEventSubscription`

NewRandafEventSubscriptionWithDefaults instantiates a new RandafEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventRequest

`func (o *RandafEventSubscription) GetEventRequest() RandafEvent`

GetEventRequest returns the EventRequest field if non-nil, zero value otherwise.

### GetEventRequestOk

`func (o *RandafEventSubscription) GetEventRequestOk() (*RandafEvent, bool)`

GetEventRequestOk returns a tuple with the EventRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRequest

`func (o *RandafEventSubscription) SetEventRequest(v RandafEvent)`

SetEventRequest sets EventRequest field to given value.


### GetEventNotifyUri

`func (o *RandafEventSubscription) GetEventNotifyUri() string`

GetEventNotifyUri returns the EventNotifyUri field if non-nil, zero value otherwise.

### GetEventNotifyUriOk

`func (o *RandafEventSubscription) GetEventNotifyUriOk() (*string, bool)`

GetEventNotifyUriOk returns a tuple with the EventNotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifyUri

`func (o *RandafEventSubscription) SetEventNotifyUri(v string)`

SetEventNotifyUri sets EventNotifyUri field to given value.


### GetOptions

`func (o *RandafEventSubscription) GetOptions() RandafEventMode`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RandafEventSubscription) GetOptionsOk() (*RandafEventMode, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RandafEventSubscription) SetOptions(v RandafEventMode)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RandafEventSubscription) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetTgtUe

`func (o *RandafEventSubscription) GetTgtUe() TargetUeInformation`

GetTgtUe returns the TgtUe field if non-nil, zero value otherwise.

### GetTgtUeOk

`func (o *RandafEventSubscription) GetTgtUeOk() (*TargetUeInformation, bool)`

GetTgtUeOk returns a tuple with the TgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUe

`func (o *RandafEventSubscription) SetTgtUe(v TargetUeInformation)`

SetTgtUe sets TgtUe field to given value.

### HasTgtUe

`func (o *RandafEventSubscription) HasTgtUe() bool`

HasTgtUe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


