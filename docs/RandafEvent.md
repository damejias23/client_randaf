# RandafEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RandafEventType**](RandafEventType.md) |  | 
**MaxReports** | Pointer to **int32** |  | [optional] 
**MaxResponseTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**MinInterval** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**NextReport** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**NextPeriodicReportTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewRandafEvent

`func NewRandafEvent(type_ RandafEventType, ) *RandafEvent`

NewRandafEvent instantiates a new RandafEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandafEventWithDefaults

`func NewRandafEventWithDefaults() *RandafEvent`

NewRandafEventWithDefaults instantiates a new RandafEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RandafEvent) GetType() RandafEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RandafEvent) GetTypeOk() (*RandafEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RandafEvent) SetType(v RandafEventType)`

SetType sets Type field to given value.


### GetMaxReports

`func (o *RandafEvent) GetMaxReports() int32`

GetMaxReports returns the MaxReports field if non-nil, zero value otherwise.

### GetMaxReportsOk

`func (o *RandafEvent) GetMaxReportsOk() (*int32, bool)`

GetMaxReportsOk returns a tuple with the MaxReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReports

`func (o *RandafEvent) SetMaxReports(v int32)`

SetMaxReports sets MaxReports field to given value.

### HasMaxReports

`func (o *RandafEvent) HasMaxReports() bool`

HasMaxReports returns a boolean if a field has been set.

### GetMaxResponseTime

`func (o *RandafEvent) GetMaxResponseTime() int32`

GetMaxResponseTime returns the MaxResponseTime field if non-nil, zero value otherwise.

### GetMaxResponseTimeOk

`func (o *RandafEvent) GetMaxResponseTimeOk() (*int32, bool)`

GetMaxResponseTimeOk returns a tuple with the MaxResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseTime

`func (o *RandafEvent) SetMaxResponseTime(v int32)`

SetMaxResponseTime sets MaxResponseTime field to given value.

### HasMaxResponseTime

`func (o *RandafEvent) HasMaxResponseTime() bool`

HasMaxResponseTime returns a boolean if a field has been set.

### GetMinInterval

`func (o *RandafEvent) GetMinInterval() int32`

GetMinInterval returns the MinInterval field if non-nil, zero value otherwise.

### GetMinIntervalOk

`func (o *RandafEvent) GetMinIntervalOk() (*int32, bool)`

GetMinIntervalOk returns a tuple with the MinInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInterval

`func (o *RandafEvent) SetMinInterval(v int32)`

SetMinInterval sets MinInterval field to given value.

### HasMinInterval

`func (o *RandafEvent) HasMinInterval() bool`

HasMinInterval returns a boolean if a field has been set.

### GetNextReport

`func (o *RandafEvent) GetNextReport() time.Time`

GetNextReport returns the NextReport field if non-nil, zero value otherwise.

### GetNextReportOk

`func (o *RandafEvent) GetNextReportOk() (*time.Time, bool)`

GetNextReportOk returns a tuple with the NextReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReport

`func (o *RandafEvent) SetNextReport(v time.Time)`

SetNextReport sets NextReport field to given value.

### HasNextReport

`func (o *RandafEvent) HasNextReport() bool`

HasNextReport returns a boolean if a field has been set.

### GetNextPeriodicReportTime

`func (o *RandafEvent) GetNextPeriodicReportTime() time.Time`

GetNextPeriodicReportTime returns the NextPeriodicReportTime field if non-nil, zero value otherwise.

### GetNextPeriodicReportTimeOk

`func (o *RandafEvent) GetNextPeriodicReportTimeOk() (*time.Time, bool)`

GetNextPeriodicReportTimeOk returns a tuple with the NextPeriodicReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPeriodicReportTime

`func (o *RandafEvent) SetNextPeriodicReportTime(v time.Time)`

SetNextPeriodicReportTime sets NextPeriodicReportTime field to given value.

### HasNextPeriodicReportTime

`func (o *RandafEvent) HasNextPeriodicReportTime() bool`

HasNextPeriodicReportTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


