# GnbInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellId** | Pointer to **int32** | Cell identifier to which the UE is connected  | [optional] 
**Plmn** | Pointer to **string** | PLMN indicator | [optional] 
**Status** | Pointer to **string** | Status indicator | [optional] 
**StatusTimestamp** | Pointer to **string** | Status timestamp indicator | [optional] 

## Methods

### NewGnbInformation

`func NewGnbInformation() *GnbInformation`

NewGnbInformation instantiates a new GnbInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGnbInformationWithDefaults

`func NewGnbInformationWithDefaults() *GnbInformation`

NewGnbInformationWithDefaults instantiates a new GnbInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellId

`func (o *GnbInformation) GetCellId() int32`

GetCellId returns the CellId field if non-nil, zero value otherwise.

### GetCellIdOk

`func (o *GnbInformation) GetCellIdOk() (*int32, bool)`

GetCellIdOk returns a tuple with the CellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellId

`func (o *GnbInformation) SetCellId(v int32)`

SetCellId sets CellId field to given value.

### HasCellId

`func (o *GnbInformation) HasCellId() bool`

HasCellId returns a boolean if a field has been set.

### GetPlmn

`func (o *GnbInformation) GetPlmn() string`

GetPlmn returns the Plmn field if non-nil, zero value otherwise.

### GetPlmnOk

`func (o *GnbInformation) GetPlmnOk() (*string, bool)`

GetPlmnOk returns a tuple with the Plmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmn

`func (o *GnbInformation) SetPlmn(v string)`

SetPlmn sets Plmn field to given value.

### HasPlmn

`func (o *GnbInformation) HasPlmn() bool`

HasPlmn returns a boolean if a field has been set.

### GetStatus

`func (o *GnbInformation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GnbInformation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GnbInformation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GnbInformation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTimestamp

`func (o *GnbInformation) GetStatusTimestamp() string`

GetStatusTimestamp returns the StatusTimestamp field if non-nil, zero value otherwise.

### GetStatusTimestampOk

`func (o *GnbInformation) GetStatusTimestampOk() (*string, bool)`

GetStatusTimestampOk returns a tuple with the StatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTimestamp

`func (o *GnbInformation) SetStatusTimestamp(v string)`

SetStatusTimestamp sets StatusTimestamp field to given value.

### HasStatusTimestamp

`func (o *GnbInformation) HasStatusTimestamp() bool`

HasStatusTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


