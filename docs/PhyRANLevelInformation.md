# PhyRANLevelInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellId** | Pointer to **int32** | Cell identifier to which the UE is connected  | [optional] 
**CellName** | Pointer to **string** | Cell name indicator | [optional] 
**Plmn** | Pointer to **string** | PLMN indicator | [optional] 
**BitrateUplink** | Pointer to **int32** | Uplink bitrate  | [optional] 
**BitrateDownlink** | Pointer to **int32** | Downlink bitrate | [optional] 
**TransmissionUplink** | Pointer to **int32** | Uplink transmission | [optional] 
**TransmissionDownlink** | Pointer to **int32** | Downlink transmission | [optional] 
**RetransmissionUplink** | Pointer to **int32** | Uplink retransmissions | [optional] 
**RetransmissionDownlink** | Pointer to **int32** | Downlink retransmissions | [optional] 

## Methods

### NewPhyRANLevelInformation

`func NewPhyRANLevelInformation() *PhyRANLevelInformation`

NewPhyRANLevelInformation instantiates a new PhyRANLevelInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhyRANLevelInformationWithDefaults

`func NewPhyRANLevelInformationWithDefaults() *PhyRANLevelInformation`

NewPhyRANLevelInformationWithDefaults instantiates a new PhyRANLevelInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellId

`func (o *PhyRANLevelInformation) GetCellId() int32`

GetCellId returns the CellId field if non-nil, zero value otherwise.

### GetCellIdOk

`func (o *PhyRANLevelInformation) GetCellIdOk() (*int32, bool)`

GetCellIdOk returns a tuple with the CellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellId

`func (o *PhyRANLevelInformation) SetCellId(v int32)`

SetCellId sets CellId field to given value.

### HasCellId

`func (o *PhyRANLevelInformation) HasCellId() bool`

HasCellId returns a boolean if a field has been set.

### GetCellName

`func (o *PhyRANLevelInformation) GetCellName() string`

GetCellName returns the CellName field if non-nil, zero value otherwise.

### GetCellNameOk

`func (o *PhyRANLevelInformation) GetCellNameOk() (*string, bool)`

GetCellNameOk returns a tuple with the CellName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellName

`func (o *PhyRANLevelInformation) SetCellName(v string)`

SetCellName sets CellName field to given value.

### HasCellName

`func (o *PhyRANLevelInformation) HasCellName() bool`

HasCellName returns a boolean if a field has been set.

### GetPlmn

`func (o *PhyRANLevelInformation) GetPlmn() string`

GetPlmn returns the Plmn field if non-nil, zero value otherwise.

### GetPlmnOk

`func (o *PhyRANLevelInformation) GetPlmnOk() (*string, bool)`

GetPlmnOk returns a tuple with the Plmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmn

`func (o *PhyRANLevelInformation) SetPlmn(v string)`

SetPlmn sets Plmn field to given value.

### HasPlmn

`func (o *PhyRANLevelInformation) HasPlmn() bool`

HasPlmn returns a boolean if a field has been set.

### GetBitrateUplink

`func (o *PhyRANLevelInformation) GetBitrateUplink() int32`

GetBitrateUplink returns the BitrateUplink field if non-nil, zero value otherwise.

### GetBitrateUplinkOk

`func (o *PhyRANLevelInformation) GetBitrateUplinkOk() (*int32, bool)`

GetBitrateUplinkOk returns a tuple with the BitrateUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrateUplink

`func (o *PhyRANLevelInformation) SetBitrateUplink(v int32)`

SetBitrateUplink sets BitrateUplink field to given value.

### HasBitrateUplink

`func (o *PhyRANLevelInformation) HasBitrateUplink() bool`

HasBitrateUplink returns a boolean if a field has been set.

### GetBitrateDownlink

`func (o *PhyRANLevelInformation) GetBitrateDownlink() int32`

GetBitrateDownlink returns the BitrateDownlink field if non-nil, zero value otherwise.

### GetBitrateDownlinkOk

`func (o *PhyRANLevelInformation) GetBitrateDownlinkOk() (*int32, bool)`

GetBitrateDownlinkOk returns a tuple with the BitrateDownlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrateDownlink

`func (o *PhyRANLevelInformation) SetBitrateDownlink(v int32)`

SetBitrateDownlink sets BitrateDownlink field to given value.

### HasBitrateDownlink

`func (o *PhyRANLevelInformation) HasBitrateDownlink() bool`

HasBitrateDownlink returns a boolean if a field has been set.

### GetTransmissionUplink

`func (o *PhyRANLevelInformation) GetTransmissionUplink() int32`

GetTransmissionUplink returns the TransmissionUplink field if non-nil, zero value otherwise.

### GetTransmissionUplinkOk

`func (o *PhyRANLevelInformation) GetTransmissionUplinkOk() (*int32, bool)`

GetTransmissionUplinkOk returns a tuple with the TransmissionUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionUplink

`func (o *PhyRANLevelInformation) SetTransmissionUplink(v int32)`

SetTransmissionUplink sets TransmissionUplink field to given value.

### HasTransmissionUplink

`func (o *PhyRANLevelInformation) HasTransmissionUplink() bool`

HasTransmissionUplink returns a boolean if a field has been set.

### GetTransmissionDownlink

`func (o *PhyRANLevelInformation) GetTransmissionDownlink() int32`

GetTransmissionDownlink returns the TransmissionDownlink field if non-nil, zero value otherwise.

### GetTransmissionDownlinkOk

`func (o *PhyRANLevelInformation) GetTransmissionDownlinkOk() (*int32, bool)`

GetTransmissionDownlinkOk returns a tuple with the TransmissionDownlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionDownlink

`func (o *PhyRANLevelInformation) SetTransmissionDownlink(v int32)`

SetTransmissionDownlink sets TransmissionDownlink field to given value.

### HasTransmissionDownlink

`func (o *PhyRANLevelInformation) HasTransmissionDownlink() bool`

HasTransmissionDownlink returns a boolean if a field has been set.

### GetRetransmissionUplink

`func (o *PhyRANLevelInformation) GetRetransmissionUplink() int32`

GetRetransmissionUplink returns the RetransmissionUplink field if non-nil, zero value otherwise.

### GetRetransmissionUplinkOk

`func (o *PhyRANLevelInformation) GetRetransmissionUplinkOk() (*int32, bool)`

GetRetransmissionUplinkOk returns a tuple with the RetransmissionUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmissionUplink

`func (o *PhyRANLevelInformation) SetRetransmissionUplink(v int32)`

SetRetransmissionUplink sets RetransmissionUplink field to given value.

### HasRetransmissionUplink

`func (o *PhyRANLevelInformation) HasRetransmissionUplink() bool`

HasRetransmissionUplink returns a boolean if a field has been set.

### GetRetransmissionDownlink

`func (o *PhyRANLevelInformation) GetRetransmissionDownlink() int32`

GetRetransmissionDownlink returns the RetransmissionDownlink field if non-nil, zero value otherwise.

### GetRetransmissionDownlinkOk

`func (o *PhyRANLevelInformation) GetRetransmissionDownlinkOk() (*int32, bool)`

GetRetransmissionDownlinkOk returns a tuple with the RetransmissionDownlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmissionDownlink

`func (o *PhyRANLevelInformation) SetRetransmissionDownlink(v int32)`

SetRetransmissionDownlink sets RetransmissionDownlink field to given value.

### HasRetransmissionDownlink

`func (o *PhyRANLevelInformation) HasRetransmissionDownlink() bool`

HasRetransmissionDownlink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


