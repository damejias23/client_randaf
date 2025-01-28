# RandafEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportEvent** | Pointer to [**RandafEventReport**](RandafEventReport.md) |  | [optional] 
**EventSubsSyncInfo** | Pointer to [**RandafEventSubsSyncInfo**](RandafEventSubsSyncInfo.md) |  | [optional] 

## Methods

### NewRandafEventNotification

`func NewRandafEventNotification() *RandafEventNotification`

NewRandafEventNotification instantiates a new RandafEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandafEventNotificationWithDefaults

`func NewRandafEventNotificationWithDefaults() *RandafEventNotification`

NewRandafEventNotificationWithDefaults instantiates a new RandafEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportEvent

`func (o *RandafEventNotification) GetReportEvent() RandafEventReport`

GetReportEvent returns the ReportEvent field if non-nil, zero value otherwise.

### GetReportEventOk

`func (o *RandafEventNotification) GetReportEventOk() (*RandafEventReport, bool)`

GetReportEventOk returns a tuple with the ReportEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportEvent

`func (o *RandafEventNotification) SetReportEvent(v RandafEventReport)`

SetReportEvent sets ReportEvent field to given value.

### HasReportEvent

`func (o *RandafEventNotification) HasReportEvent() bool`

HasReportEvent returns a boolean if a field has been set.

### GetEventSubsSyncInfo

`func (o *RandafEventNotification) GetEventSubsSyncInfo() RandafEventSubsSyncInfo`

GetEventSubsSyncInfo returns the EventSubsSyncInfo field if non-nil, zero value otherwise.

### GetEventSubsSyncInfoOk

`func (o *RandafEventNotification) GetEventSubsSyncInfoOk() (*RandafEventSubsSyncInfo, bool)`

GetEventSubsSyncInfoOk returns a tuple with the EventSubsSyncInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubsSyncInfo

`func (o *RandafEventNotification) SetEventSubsSyncInfo(v RandafEventSubsSyncInfo)`

SetEventSubsSyncInfo sets EventSubsSyncInfo field to given value.

### HasEventSubsSyncInfo

`func (o *RandafEventNotification) HasEventSubsSyncInfo() bool`

HasEventSubsSyncInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


