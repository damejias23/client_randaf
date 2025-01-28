# RandafEventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RandafEventType**](RandafEventType.md) |  | 
**State** | [**RandafEventState**](RandafEventState.md) |  | 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**SubscriptionId** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**PhyRANLevelInfos** | Pointer to [**[]PhyRANLevelInformation**](PhyRANLevelInformation.md) |  | [optional] 
**GnbInfos** | Pointer to [**[]GnbInformation**](GnbInformation.md) |  | [optional] 

## Methods

### NewRandafEventReport

`func NewRandafEventReport(type_ RandafEventType, state RandafEventState, timeStamp time.Time, ) *RandafEventReport`

NewRandafEventReport instantiates a new RandafEventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandafEventReportWithDefaults

`func NewRandafEventReportWithDefaults() *RandafEventReport`

NewRandafEventReportWithDefaults instantiates a new RandafEventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RandafEventReport) GetType() RandafEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RandafEventReport) GetTypeOk() (*RandafEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RandafEventReport) SetType(v RandafEventType)`

SetType sets Type field to given value.


### GetState

`func (o *RandafEventReport) GetState() RandafEventState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RandafEventReport) GetStateOk() (*RandafEventState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RandafEventReport) SetState(v RandafEventState)`

SetState sets State field to given value.


### GetTimeStamp

`func (o *RandafEventReport) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *RandafEventReport) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *RandafEventReport) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetSubscriptionId

`func (o *RandafEventReport) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RandafEventReport) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RandafEventReport) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *RandafEventReport) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetPhyRANLevelInfos

`func (o *RandafEventReport) GetPhyRANLevelInfos() []PhyRANLevelInformation`

GetPhyRANLevelInfos returns the PhyRANLevelInfos field if non-nil, zero value otherwise.

### GetPhyRANLevelInfosOk

`func (o *RandafEventReport) GetPhyRANLevelInfosOk() (*[]PhyRANLevelInformation, bool)`

GetPhyRANLevelInfosOk returns a tuple with the PhyRANLevelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhyRANLevelInfos

`func (o *RandafEventReport) SetPhyRANLevelInfos(v []PhyRANLevelInformation)`

SetPhyRANLevelInfos sets PhyRANLevelInfos field to given value.

### HasPhyRANLevelInfos

`func (o *RandafEventReport) HasPhyRANLevelInfos() bool`

HasPhyRANLevelInfos returns a boolean if a field has been set.

### GetGnbInfos

`func (o *RandafEventReport) GetGnbInfos() []GnbInformation`

GetGnbInfos returns the GnbInfos field if non-nil, zero value otherwise.

### GetGnbInfosOk

`func (o *RandafEventReport) GetGnbInfosOk() (*[]GnbInformation, bool)`

GetGnbInfosOk returns a tuple with the GnbInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbInfos

`func (o *RandafEventReport) SetGnbInfos(v []GnbInformation)`

SetGnbInfos sets GnbInfos field to given value.

### HasGnbInfos

`func (o *RandafEventReport) HasGnbInfos() bool`

HasGnbInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


