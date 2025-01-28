# RandafEventState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**RemainReports** | Pointer to **int32** |  | [optional] 
**RemainDuration** | Pointer to **int32** | indicating a time in seconds. | [optional] 

## Methods

### NewRandafEventState

`func NewRandafEventState(active bool, ) *RandafEventState`

NewRandafEventState instantiates a new RandafEventState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandafEventStateWithDefaults

`func NewRandafEventStateWithDefaults() *RandafEventState`

NewRandafEventStateWithDefaults instantiates a new RandafEventState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RandafEventState) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RandafEventState) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RandafEventState) SetActive(v bool)`

SetActive sets Active field to given value.


### GetRemainReports

`func (o *RandafEventState) GetRemainReports() int32`

GetRemainReports returns the RemainReports field if non-nil, zero value otherwise.

### GetRemainReportsOk

`func (o *RandafEventState) GetRemainReportsOk() (*int32, bool)`

GetRemainReportsOk returns a tuple with the RemainReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainReports

`func (o *RandafEventState) SetRemainReports(v int32)`

SetRemainReports sets RemainReports field to given value.

### HasRemainReports

`func (o *RandafEventState) HasRemainReports() bool`

HasRemainReports returns a boolean if a field has been set.

### GetRemainDuration

`func (o *RandafEventState) GetRemainDuration() int32`

GetRemainDuration returns the RemainDuration field if non-nil, zero value otherwise.

### GetRemainDurationOk

`func (o *RandafEventState) GetRemainDurationOk() (*int32, bool)`

GetRemainDurationOk returns a tuple with the RemainDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainDuration

`func (o *RandafEventState) SetRemainDuration(v int32)`

SetRemainDuration sets RemainDuration field to given value.

### HasRemainDuration

`func (o *RandafEventState) HasRemainDuration() bool`

HasRemainDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


