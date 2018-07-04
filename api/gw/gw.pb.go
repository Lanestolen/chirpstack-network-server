// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gw.proto

package gw // import "github.com/brocaar/loraserver/api/gw"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/brocaar/loraserver/api/common"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UplinkTXInfo struct {
	// Frequency (Hz).
	Frequency uint32 `protobuf:"varint,1,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Modulation.
	Modulation common.Modulation `protobuf:"varint,2,opt,name=modulation,proto3,enum=common.Modulation" json:"modulation,omitempty"`
	// Types that are valid to be assigned to ModulationInfo:
	//	*UplinkTXInfo_LoraModulationInfo
	//	*UplinkTXInfo_FskModulationInfo
	ModulationInfo       isUplinkTXInfo_ModulationInfo `protobuf_oneof:"modulation_info"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *UplinkTXInfo) Reset()         { *m = UplinkTXInfo{} }
func (m *UplinkTXInfo) String() string { return proto.CompactTextString(m) }
func (*UplinkTXInfo) ProtoMessage()    {}
func (*UplinkTXInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_90ea3f7651ea710d, []int{0}
}
func (m *UplinkTXInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkTXInfo.Unmarshal(m, b)
}
func (m *UplinkTXInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkTXInfo.Marshal(b, m, deterministic)
}
func (dst *UplinkTXInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkTXInfo.Merge(dst, src)
}
func (m *UplinkTXInfo) XXX_Size() int {
	return xxx_messageInfo_UplinkTXInfo.Size(m)
}
func (m *UplinkTXInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkTXInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkTXInfo proto.InternalMessageInfo

type isUplinkTXInfo_ModulationInfo interface {
	isUplinkTXInfo_ModulationInfo()
}

type UplinkTXInfo_LoraModulationInfo struct {
	LoraModulationInfo *LoRaModulationInfo `protobuf:"bytes,3,opt,name=lora_modulation_info,json=loraModulationInfo,proto3,oneof"`
}
type UplinkTXInfo_FskModulationInfo struct {
	FskModulationInfo *FSKModulationInfo `protobuf:"bytes,4,opt,name=fsk_modulation_info,json=fskModulationInfo,proto3,oneof"`
}

func (*UplinkTXInfo_LoraModulationInfo) isUplinkTXInfo_ModulationInfo() {}
func (*UplinkTXInfo_FskModulationInfo) isUplinkTXInfo_ModulationInfo()  {}

func (m *UplinkTXInfo) GetModulationInfo() isUplinkTXInfo_ModulationInfo {
	if m != nil {
		return m.ModulationInfo
	}
	return nil
}

func (m *UplinkTXInfo) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *UplinkTXInfo) GetModulation() common.Modulation {
	if m != nil {
		return m.Modulation
	}
	return common.Modulation_LORA
}

func (m *UplinkTXInfo) GetLoraModulationInfo() *LoRaModulationInfo {
	if x, ok := m.GetModulationInfo().(*UplinkTXInfo_LoraModulationInfo); ok {
		return x.LoraModulationInfo
	}
	return nil
}

func (m *UplinkTXInfo) GetFskModulationInfo() *FSKModulationInfo {
	if x, ok := m.GetModulationInfo().(*UplinkTXInfo_FskModulationInfo); ok {
		return x.FskModulationInfo
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UplinkTXInfo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UplinkTXInfo_OneofMarshaler, _UplinkTXInfo_OneofUnmarshaler, _UplinkTXInfo_OneofSizer, []interface{}{
		(*UplinkTXInfo_LoraModulationInfo)(nil),
		(*UplinkTXInfo_FskModulationInfo)(nil),
	}
}

func _UplinkTXInfo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UplinkTXInfo)
	// modulation_info
	switch x := m.ModulationInfo.(type) {
	case *UplinkTXInfo_LoraModulationInfo:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LoraModulationInfo); err != nil {
			return err
		}
	case *UplinkTXInfo_FskModulationInfo:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FskModulationInfo); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UplinkTXInfo.ModulationInfo has unexpected type %T", x)
	}
	return nil
}

func _UplinkTXInfo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UplinkTXInfo)
	switch tag {
	case 3: // modulation_info.lora_modulation_info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoRaModulationInfo)
		err := b.DecodeMessage(msg)
		m.ModulationInfo = &UplinkTXInfo_LoraModulationInfo{msg}
		return true, err
	case 4: // modulation_info.fsk_modulation_info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FSKModulationInfo)
		err := b.DecodeMessage(msg)
		m.ModulationInfo = &UplinkTXInfo_FskModulationInfo{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UplinkTXInfo_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UplinkTXInfo)
	// modulation_info
	switch x := m.ModulationInfo.(type) {
	case *UplinkTXInfo_LoraModulationInfo:
		s := proto.Size(x.LoraModulationInfo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UplinkTXInfo_FskModulationInfo:
		s := proto.Size(x.FskModulationInfo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LoRaModulationInfo struct {
	// Bandwidth.
	Bandwidth uint32 `protobuf:"varint,1,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	// Speading-factor.
	SpreadingFactor uint32 `protobuf:"varint,2,opt,name=spreading_factor,json=spreadingFactor,proto3" json:"spreading_factor,omitempty"`
	// Code-rate.
	CodeRate string `protobuf:"bytes,3,opt,name=code_rate,json=codeRate,proto3" json:"code_rate,omitempty"`
	// Polarization inversion.
	PolarizationInversion bool     `protobuf:"varint,4,opt,name=polarization_inversion,json=polarizationInversion,proto3" json:"polarization_inversion,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *LoRaModulationInfo) Reset()         { *m = LoRaModulationInfo{} }
func (m *LoRaModulationInfo) String() string { return proto.CompactTextString(m) }
func (*LoRaModulationInfo) ProtoMessage()    {}
func (*LoRaModulationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_90ea3f7651ea710d, []int{1}
}
func (m *LoRaModulationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoRaModulationInfo.Unmarshal(m, b)
}
func (m *LoRaModulationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoRaModulationInfo.Marshal(b, m, deterministic)
}
func (dst *LoRaModulationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoRaModulationInfo.Merge(dst, src)
}
func (m *LoRaModulationInfo) XXX_Size() int {
	return xxx_messageInfo_LoRaModulationInfo.Size(m)
}
func (m *LoRaModulationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LoRaModulationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LoRaModulationInfo proto.InternalMessageInfo

func (m *LoRaModulationInfo) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *LoRaModulationInfo) GetSpreadingFactor() uint32 {
	if m != nil {
		return m.SpreadingFactor
	}
	return 0
}

func (m *LoRaModulationInfo) GetCodeRate() string {
	if m != nil {
		return m.CodeRate
	}
	return ""
}

func (m *LoRaModulationInfo) GetPolarizationInversion() bool {
	if m != nil {
		return m.PolarizationInversion
	}
	return false
}

type FSKModulationInfo struct {
	// Bandwidth.
	Bandwidth uint32 `protobuf:"varint,1,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	// Bitrate.
	Bitrate              uint32   `protobuf:"varint,2,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FSKModulationInfo) Reset()         { *m = FSKModulationInfo{} }
func (m *FSKModulationInfo) String() string { return proto.CompactTextString(m) }
func (*FSKModulationInfo) ProtoMessage()    {}
func (*FSKModulationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_90ea3f7651ea710d, []int{2}
}
func (m *FSKModulationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FSKModulationInfo.Unmarshal(m, b)
}
func (m *FSKModulationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FSKModulationInfo.Marshal(b, m, deterministic)
}
func (dst *FSKModulationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FSKModulationInfo.Merge(dst, src)
}
func (m *FSKModulationInfo) XXX_Size() int {
	return xxx_messageInfo_FSKModulationInfo.Size(m)
}
func (m *FSKModulationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FSKModulationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FSKModulationInfo proto.InternalMessageInfo

func (m *FSKModulationInfo) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *FSKModulationInfo) GetBitrate() uint32 {
	if m != nil {
		return m.Bitrate
	}
	return 0
}

type Location struct {
	// Latitude.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude.
	Altitude             float64  `protobuf:"fixed64,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_90ea3f7651ea710d, []int{3}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

type UplinkRXInfo struct {
	// Gateway ID.
	GatewayId []byte `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	// RX time (only set when the gateway has a GPS module).
	Time *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// RX time since GPS epoch (only set when the gateway has a GPS module).
	TimeSinceGpsEpoch *duration.Duration `protobuf:"bytes,3,opt,name=time_since_gps_epoch,json=timeSinceGpsEpoch,proto3" json:"time_since_gps_epoch,omitempty"`
	// Gateway internal timestamp.
	Timestamp uint32 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// RSSI.
	Rssi int32 `protobuf:"varint,5,opt,name=rssi,proto3" json:"rssi,omitempty"`
	// LoRa SNR.
	LoraSnr float64 `protobuf:"fixed64,6,opt,name=lora_snr,json=loraSnr,proto3" json:"lora_snr,omitempty"`
	// Channel.
	Channel uint32 `protobuf:"varint,7,opt,name=channel,proto3" json:"channel,omitempty"`
	// RF Chain.
	RfChain uint32 `protobuf:"varint,8,opt,name=rf_chain,json=rfChain,proto3" json:"rf_chain,omitempty"`
	// Board.
	Board uint32 `protobuf:"varint,9,opt,name=board,proto3" json:"board,omitempty"`
	// Antenna.
	Antenna uint32 `protobuf:"varint,10,opt,name=antenna,proto3" json:"antenna,omitempty"`
	// Location.
	Location             *Location `protobuf:"bytes,11,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UplinkRXInfo) Reset()         { *m = UplinkRXInfo{} }
func (m *UplinkRXInfo) String() string { return proto.CompactTextString(m) }
func (*UplinkRXInfo) ProtoMessage()    {}
func (*UplinkRXInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_90ea3f7651ea710d, []int{4}
}
func (m *UplinkRXInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkRXInfo.Unmarshal(m, b)
}
func (m *UplinkRXInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkRXInfo.Marshal(b, m, deterministic)
}
func (dst *UplinkRXInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkRXInfo.Merge(dst, src)
}
func (m *UplinkRXInfo) XXX_Size() int {
	return xxx_messageInfo_UplinkRXInfo.Size(m)
}
func (m *UplinkRXInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkRXInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkRXInfo proto.InternalMessageInfo

func (m *UplinkRXInfo) GetGatewayId() []byte {
	if m != nil {
		return m.GatewayId
	}
	return nil
}

func (m *UplinkRXInfo) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *UplinkRXInfo) GetTimeSinceGpsEpoch() *duration.Duration {
	if m != nil {
		return m.TimeSinceGpsEpoch
	}
	return nil
}

func (m *UplinkRXInfo) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *UplinkRXInfo) GetRssi() int32 {
	if m != nil {
		return m.Rssi
	}
	return 0
}

func (m *UplinkRXInfo) GetLoraSnr() float64 {
	if m != nil {
		return m.LoraSnr
	}
	return 0
}

func (m *UplinkRXInfo) GetChannel() uint32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *UplinkRXInfo) GetRfChain() uint32 {
	if m != nil {
		return m.RfChain
	}
	return 0
}

func (m *UplinkRXInfo) GetBoard() uint32 {
	if m != nil {
		return m.Board
	}
	return 0
}

func (m *UplinkRXInfo) GetAntenna() uint32 {
	if m != nil {
		return m.Antenna
	}
	return 0
}

func (m *UplinkRXInfo) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type DownlinkTXInfo struct {
	// Gateway ID.
	GatewayId []byte `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	// Frame must be sent immediately.
	Immediately bool `protobuf:"varint,2,opt,name=immediately,proto3" json:"immediately,omitempty"`
	// Emit frame at the given time since GPS epoch.
	TimeSinceGpsEpoch *duration.Duration `protobuf:"bytes,3,opt,name=time_since_gps_epoch,json=timeSinceGpsEpoch,proto3" json:"time_since_gps_epoch,omitempty"`
	// Emit the frame at the given gateway internal timestamp.
	Timestamp uint32 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// TX frequency (in Hz).
	Frequency uint32 `protobuf:"varint,5,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// TX power (in dBm).
	Power int32 `protobuf:"varint,6,opt,name=power,proto3" json:"power,omitempty"`
	// Modulation.
	Modulation common.Modulation `protobuf:"varint,7,opt,name=modulation,proto3,enum=common.Modulation" json:"modulation,omitempty"`
	// Types that are valid to be assigned to ModulationInfo:
	//	*DownlinkTXInfo_LoraModulationInfo
	//	*DownlinkTXInfo_FskModulationInfo
	ModulationInfo isDownlinkTXInfo_ModulationInfo `protobuf_oneof:"modulation_info"`
	// The board identifier for emitting the frame.
	Board uint32 `protobuf:"varint,10,opt,name=board,proto3" json:"board,omitempty"`
	// The antenna identifier for emitting the frame.
	Antenna              uint32   `protobuf:"varint,11,opt,name=antenna,proto3" json:"antenna,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownlinkTXInfo) Reset()         { *m = DownlinkTXInfo{} }
func (m *DownlinkTXInfo) String() string { return proto.CompactTextString(m) }
func (*DownlinkTXInfo) ProtoMessage()    {}
func (*DownlinkTXInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_90ea3f7651ea710d, []int{5}
}
func (m *DownlinkTXInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkTXInfo.Unmarshal(m, b)
}
func (m *DownlinkTXInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkTXInfo.Marshal(b, m, deterministic)
}
func (dst *DownlinkTXInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkTXInfo.Merge(dst, src)
}
func (m *DownlinkTXInfo) XXX_Size() int {
	return xxx_messageInfo_DownlinkTXInfo.Size(m)
}
func (m *DownlinkTXInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkTXInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkTXInfo proto.InternalMessageInfo

type isDownlinkTXInfo_ModulationInfo interface {
	isDownlinkTXInfo_ModulationInfo()
}

type DownlinkTXInfo_LoraModulationInfo struct {
	LoraModulationInfo *LoRaModulationInfo `protobuf:"bytes,8,opt,name=lora_modulation_info,json=loraModulationInfo,proto3,oneof"`
}
type DownlinkTXInfo_FskModulationInfo struct {
	FskModulationInfo *FSKModulationInfo `protobuf:"bytes,9,opt,name=fsk_modulation_info,json=fskModulationInfo,proto3,oneof"`
}

func (*DownlinkTXInfo_LoraModulationInfo) isDownlinkTXInfo_ModulationInfo() {}
func (*DownlinkTXInfo_FskModulationInfo) isDownlinkTXInfo_ModulationInfo()  {}

func (m *DownlinkTXInfo) GetModulationInfo() isDownlinkTXInfo_ModulationInfo {
	if m != nil {
		return m.ModulationInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetGatewayId() []byte {
	if m != nil {
		return m.GatewayId
	}
	return nil
}

func (m *DownlinkTXInfo) GetImmediately() bool {
	if m != nil {
		return m.Immediately
	}
	return false
}

func (m *DownlinkTXInfo) GetTimeSinceGpsEpoch() *duration.Duration {
	if m != nil {
		return m.TimeSinceGpsEpoch
	}
	return nil
}

func (m *DownlinkTXInfo) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DownlinkTXInfo) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *DownlinkTXInfo) GetPower() int32 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *DownlinkTXInfo) GetModulation() common.Modulation {
	if m != nil {
		return m.Modulation
	}
	return common.Modulation_LORA
}

func (m *DownlinkTXInfo) GetLoraModulationInfo() *LoRaModulationInfo {
	if x, ok := m.GetModulationInfo().(*DownlinkTXInfo_LoraModulationInfo); ok {
		return x.LoraModulationInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetFskModulationInfo() *FSKModulationInfo {
	if x, ok := m.GetModulationInfo().(*DownlinkTXInfo_FskModulationInfo); ok {
		return x.FskModulationInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetBoard() uint32 {
	if m != nil {
		return m.Board
	}
	return 0
}

func (m *DownlinkTXInfo) GetAntenna() uint32 {
	if m != nil {
		return m.Antenna
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DownlinkTXInfo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DownlinkTXInfo_OneofMarshaler, _DownlinkTXInfo_OneofUnmarshaler, _DownlinkTXInfo_OneofSizer, []interface{}{
		(*DownlinkTXInfo_LoraModulationInfo)(nil),
		(*DownlinkTXInfo_FskModulationInfo)(nil),
	}
}

func _DownlinkTXInfo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DownlinkTXInfo)
	// modulation_info
	switch x := m.ModulationInfo.(type) {
	case *DownlinkTXInfo_LoraModulationInfo:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LoraModulationInfo); err != nil {
			return err
		}
	case *DownlinkTXInfo_FskModulationInfo:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FskModulationInfo); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DownlinkTXInfo.ModulationInfo has unexpected type %T", x)
	}
	return nil
}

func _DownlinkTXInfo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DownlinkTXInfo)
	switch tag {
	case 8: // modulation_info.lora_modulation_info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoRaModulationInfo)
		err := b.DecodeMessage(msg)
		m.ModulationInfo = &DownlinkTXInfo_LoraModulationInfo{msg}
		return true, err
	case 9: // modulation_info.fsk_modulation_info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FSKModulationInfo)
		err := b.DecodeMessage(msg)
		m.ModulationInfo = &DownlinkTXInfo_FskModulationInfo{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DownlinkTXInfo_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DownlinkTXInfo)
	// modulation_info
	switch x := m.ModulationInfo.(type) {
	case *DownlinkTXInfo_LoraModulationInfo:
		s := proto.Size(x.LoraModulationInfo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DownlinkTXInfo_FskModulationInfo:
		s := proto.Size(x.FskModulationInfo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type UplinkFrame struct {
	// PHYPayload.
	PhyPayload []byte `protobuf:"bytes,1,opt,name=phy_payload,json=phyPayload,proto3" json:"phy_payload,omitempty"`
	// TX meta-data.
	TxInfo *UplinkTXInfo `protobuf:"bytes,2,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// RX meta-data.
	RxInfo               *UplinkRXInfo `protobuf:"bytes,3,opt,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UplinkFrame) Reset()         { *m = UplinkFrame{} }
func (m *UplinkFrame) String() string { return proto.CompactTextString(m) }
func (*UplinkFrame) ProtoMessage()    {}
func (*UplinkFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_90ea3f7651ea710d, []int{6}
}
func (m *UplinkFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkFrame.Unmarshal(m, b)
}
func (m *UplinkFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkFrame.Marshal(b, m, deterministic)
}
func (dst *UplinkFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkFrame.Merge(dst, src)
}
func (m *UplinkFrame) XXX_Size() int {
	return xxx_messageInfo_UplinkFrame.Size(m)
}
func (m *UplinkFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkFrame.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkFrame proto.InternalMessageInfo

func (m *UplinkFrame) GetPhyPayload() []byte {
	if m != nil {
		return m.PhyPayload
	}
	return nil
}

func (m *UplinkFrame) GetTxInfo() *UplinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *UplinkFrame) GetRxInfo() *UplinkRXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

type UplinkFrameSet struct {
	// PHYPayload.
	PhyPayload []byte `protobuf:"bytes,1,opt,name=phy_payload,json=phyPayload,proto3" json:"phy_payload,omitempty"`
	// TX meta-data.
	TxInfo *UplinkTXInfo `protobuf:"bytes,2,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// RX meta-data set.
	RxInfo               []*UplinkRXInfo `protobuf:"bytes,3,rep,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UplinkFrameSet) Reset()         { *m = UplinkFrameSet{} }
func (m *UplinkFrameSet) String() string { return proto.CompactTextString(m) }
func (*UplinkFrameSet) ProtoMessage()    {}
func (*UplinkFrameSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_90ea3f7651ea710d, []int{7}
}
func (m *UplinkFrameSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkFrameSet.Unmarshal(m, b)
}
func (m *UplinkFrameSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkFrameSet.Marshal(b, m, deterministic)
}
func (dst *UplinkFrameSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkFrameSet.Merge(dst, src)
}
func (m *UplinkFrameSet) XXX_Size() int {
	return xxx_messageInfo_UplinkFrameSet.Size(m)
}
func (m *UplinkFrameSet) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkFrameSet.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkFrameSet proto.InternalMessageInfo

func (m *UplinkFrameSet) GetPhyPayload() []byte {
	if m != nil {
		return m.PhyPayload
	}
	return nil
}

func (m *UplinkFrameSet) GetTxInfo() *UplinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *UplinkFrameSet) GetRxInfo() []*UplinkRXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

type DownlinkFrame struct {
	// PHYPayload.
	PhyPayload []byte `protobuf:"bytes,1,opt,name=phy_payload,json=phyPayload,proto3" json:"phy_payload,omitempty"`
	// TX meta-data.
	TxInfo *DownlinkTXInfo `protobuf:"bytes,2,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// Token (uint16 value).
	Token                uint32   `protobuf:"varint,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownlinkFrame) Reset()         { *m = DownlinkFrame{} }
func (m *DownlinkFrame) String() string { return proto.CompactTextString(m) }
func (*DownlinkFrame) ProtoMessage()    {}
func (*DownlinkFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_90ea3f7651ea710d, []int{8}
}
func (m *DownlinkFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkFrame.Unmarshal(m, b)
}
func (m *DownlinkFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkFrame.Marshal(b, m, deterministic)
}
func (dst *DownlinkFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkFrame.Merge(dst, src)
}
func (m *DownlinkFrame) XXX_Size() int {
	return xxx_messageInfo_DownlinkFrame.Size(m)
}
func (m *DownlinkFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkFrame.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkFrame proto.InternalMessageInfo

func (m *DownlinkFrame) GetPhyPayload() []byte {
	if m != nil {
		return m.PhyPayload
	}
	return nil
}

func (m *DownlinkFrame) GetTxInfo() *DownlinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *DownlinkFrame) GetToken() uint32 {
	if m != nil {
		return m.Token
	}
	return 0
}

func init() {
	proto.RegisterType((*UplinkTXInfo)(nil), "gw.UplinkTXInfo")
	proto.RegisterType((*LoRaModulationInfo)(nil), "gw.LoRaModulationInfo")
	proto.RegisterType((*FSKModulationInfo)(nil), "gw.FSKModulationInfo")
	proto.RegisterType((*Location)(nil), "gw.Location")
	proto.RegisterType((*UplinkRXInfo)(nil), "gw.UplinkRXInfo")
	proto.RegisterType((*DownlinkTXInfo)(nil), "gw.DownlinkTXInfo")
	proto.RegisterType((*UplinkFrame)(nil), "gw.UplinkFrame")
	proto.RegisterType((*UplinkFrameSet)(nil), "gw.UplinkFrameSet")
	proto.RegisterType((*DownlinkFrame)(nil), "gw.DownlinkFrame")
}

func init() { proto.RegisterFile("gw.proto", fileDescriptor_gw_90ea3f7651ea710d) }

var fileDescriptor_gw_90ea3f7651ea710d = []byte{
	// 802 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x5b, 0x8f, 0xdb, 0x44,
	0x14, 0xc6, 0x9b, 0xcd, 0xc6, 0x39, 0xd9, 0x6c, 0xbb, 0x43, 0x5a, 0xb9, 0xe1, 0xd2, 0x28, 0x42,
	0x28, 0x2b, 0x24, 0x47, 0x0a, 0xea, 0x1f, 0x28, 0x65, 0xcb, 0xd2, 0x22, 0xa1, 0x49, 0x91, 0x10,
	0x2f, 0x66, 0x62, 0x8f, 0x9d, 0x51, 0xec, 0x19, 0x33, 0x9e, 0xac, 0x37, 0x3c, 0x83, 0xe0, 0xef,
	0xf0, 0xfb, 0x78, 0x41, 0x33, 0x63, 0x27, 0x4e, 0x02, 0x2c, 0x88, 0xcb, 0x53, 0x72, 0xbe, 0x73,
	0x99, 0xef, 0x5c, 0x0d, 0x6e, 0x52, 0xfa, 0xb9, 0x14, 0x4a, 0xa0, 0x93, 0xa4, 0x1c, 0x3e, 0x4b,
	0x98, 0x5a, 0xae, 0x17, 0x7e, 0x28, 0xb2, 0xe9, 0x42, 0x8a, 0x90, 0x10, 0x39, 0x4d, 0x85, 0x24,
	0x05, 0x95, 0xb7, 0x54, 0x4e, 0x49, 0xce, 0xa6, 0xa1, 0xc8, 0x32, 0xc1, 0xab, 0x1f, 0xeb, 0x3a,
	0x7c, 0x9a, 0x08, 0x91, 0xa4, 0x74, 0x6a, 0xa4, 0xc5, 0x3a, 0x9e, 0x2a, 0x96, 0xd1, 0x42, 0x91,
	0x2c, 0xaf, 0x0c, 0xde, 0x3f, 0x34, 0x88, 0xd6, 0x92, 0x28, 0x56, 0x07, 0x18, 0xff, 0x7c, 0x02,
	0xe7, 0x5f, 0xe5, 0x29, 0xe3, 0xab, 0x37, 0x5f, 0xdf, 0xf0, 0x58, 0xa0, 0x77, 0xa1, 0x1b, 0x4b,
	0xfa, 0xdd, 0x9a, 0xf2, 0x70, 0xe3, 0x39, 0x23, 0x67, 0xd2, 0xc7, 0x3b, 0x00, 0xcd, 0x00, 0x32,
	0x11, 0xad, 0x53, 0x13, 0xc2, 0x3b, 0x19, 0x39, 0x93, 0x8b, 0x19, 0xf2, 0x2b, 0x4a, 0x5f, 0x6c,
	0x35, 0xb8, 0x61, 0x85, 0x3e, 0x87, 0x81, 0xce, 0x24, 0xd8, 0x41, 0x01, 0xe3, 0xb1, 0xf0, 0x5a,
	0x23, 0x67, 0xd2, 0x9b, 0x3d, 0xf6, 0x93, 0xd2, 0x7f, 0x2d, 0x30, 0xd9, 0x79, 0x6b, 0x1e, 0x9f,
	0xbd, 0x85, 0x91, 0xf6, 0xda, 0x47, 0xd1, 0x4b, 0x78, 0x3b, 0x2e, 0x56, 0x47, 0xa1, 0x4e, 0x4d,
	0xa8, 0x47, 0x3a, 0xd4, 0xf5, 0xfc, 0xd5, 0x51, 0xa4, 0xcb, 0xb8, 0x58, 0xed, 0x83, 0xcf, 0x2f,
	0xe1, 0xc1, 0x41, 0x90, 0xf1, 0x2f, 0x0e, 0xa0, 0x63, 0x22, 0xba, 0x20, 0x0b, 0xc2, 0xa3, 0x92,
	0x45, 0x6a, 0x59, 0x17, 0x64, 0x0b, 0xa0, 0x2b, 0x78, 0x58, 0xe4, 0x92, 0x92, 0x88, 0xf1, 0x24,
	0x88, 0x49, 0xa8, 0x84, 0x34, 0x65, 0xe9, 0xe3, 0x07, 0x5b, 0xfc, 0xda, 0xc0, 0xe8, 0x1d, 0xe8,
	0x86, 0x22, 0xa2, 0x81, 0x24, 0x8a, 0x9a, 0xe4, 0xbb, 0xd8, 0xd5, 0x00, 0x26, 0x8a, 0xa2, 0x67,
	0xf0, 0x38, 0x17, 0x29, 0x91, 0xec, 0xfb, 0x9a, 0xd1, 0x2d, 0x95, 0x85, 0x2e, 0xb2, 0xce, 0xcd,
	0xc5, 0x8f, 0x9a, 0xda, 0x9b, 0x5a, 0x39, 0x7e, 0x05, 0x97, 0x47, 0x09, 0xdf, 0xc3, 0xd8, 0x83,
	0xce, 0x82, 0x29, 0x43, 0xc2, 0x12, 0xad, 0xc5, 0xf1, 0xb7, 0xe0, 0xbe, 0x16, 0xa1, 0x6d, 0xda,
	0x10, 0x5c, 0x1d, 0x51, 0xad, 0x23, 0x6a, 0x42, 0x38, 0x78, 0x2b, 0xeb, 0xf8, 0xa9, 0xe0, 0x89,
	0x55, 0x9e, 0x18, 0xe5, 0x0e, 0xd0, 0x9e, 0x24, 0xad, 0x3c, 0x5b, 0xd6, 0xb3, 0x96, 0xc7, 0x3f,
	0xb6, 0xea, 0x69, 0xc3, 0x76, 0xda, 0xde, 0x03, 0x48, 0x88, 0xa2, 0x25, 0xd9, 0x04, 0x2c, 0x32,
	0x0f, 0x9d, 0xe3, 0x6e, 0x85, 0xdc, 0x44, 0xc8, 0x87, 0x53, 0x3d, 0xd0, 0xe6, 0x91, 0xde, 0x6c,
	0xe8, 0xdb, 0x61, 0xf6, 0xeb, 0x61, 0xf6, 0xdf, 0xd4, 0xd3, 0x8e, 0x8d, 0x9d, 0x1e, 0x35, 0xfd,
	0x1b, 0x14, 0x8c, 0x87, 0x34, 0x48, 0xf2, 0x22, 0xa0, 0xb9, 0x08, 0x97, 0xd5, 0xa8, 0x3d, 0x39,
	0xf2, 0x7f, 0x51, 0x2d, 0x03, 0xbe, 0xd4, 0x6e, 0x73, 0xed, 0xf5, 0x32, 0x2f, 0x3e, 0xd5, 0x3e,
	0x3a, 0xcb, 0xed, 0x32, 0x99, 0x26, 0xf4, 0xf1, 0x0e, 0x40, 0x08, 0x4e, 0x65, 0x51, 0x30, 0xaf,
	0x3d, 0x72, 0x26, 0x6d, 0x6c, 0xfe, 0xa3, 0x27, 0xe0, 0x9a, 0x41, 0x2f, 0xb8, 0xf4, 0xce, 0x4c,
	0xe6, 0x1d, 0x2d, 0xcf, 0xb9, 0xd4, 0x45, 0x0f, 0x97, 0x84, 0x73, 0x9a, 0x7a, 0x1d, 0x5b, 0xf4,
	0x4a, 0xd4, 0x4e, 0x32, 0x0e, 0xc2, 0x25, 0x61, 0xdc, 0x73, 0xad, 0x4a, 0xc6, 0x9f, 0x68, 0x11,
	0x0d, 0xa0, 0xbd, 0x10, 0x44, 0x46, 0x5e, 0xd7, 0xe0, 0x56, 0xd0, 0xa1, 0x08, 0x57, 0x94, 0x73,
	0xe2, 0x81, 0xb5, 0xaf, 0x44, 0x34, 0xd1, 0xef, 0xdb, 0xfe, 0x79, 0x3d, 0x93, 0xf1, 0xb9, 0x5d,
	0x2e, 0x8b, 0xe1, 0xad, 0x76, 0xfc, 0x6b, 0x0b, 0x2e, 0x5e, 0x88, 0x92, 0x37, 0xf6, 0xfe, 0x9e,
	0x4e, 0x8c, 0xa0, 0xc7, 0xb2, 0x8c, 0x46, 0x8c, 0x28, 0x9a, 0x6e, 0x4c, 0x43, 0x5c, 0xdc, 0x84,
	0xfe, 0xc7, 0xda, 0xef, 0x9d, 0xa8, 0xf6, 0xe1, 0x89, 0x1a, 0x40, 0x3b, 0x17, 0x25, 0xb5, 0x2d,
	0x68, 0x63, 0x2b, 0x1c, 0x1c, 0xae, 0xce, 0x3f, 0x3a, 0x5c, 0xee, 0xbf, 0x77, 0xb8, 0xba, 0x7f,
	0xf7, 0x70, 0xed, 0x86, 0x02, 0xfe, 0x60, 0x28, 0x7a, 0x7b, 0x43, 0xf1, 0x7b, 0x87, 0xee, 0x07,
	0x07, 0x7a, 0x76, 0x0b, 0xaf, 0x25, 0xc9, 0x28, 0x7a, 0x0a, 0xbd, 0x7c, 0xb9, 0x09, 0x72, 0xb2,
	0x49, 0x05, 0xa9, 0x7b, 0x0f, 0xf9, 0x72, 0xf3, 0xa5, 0x45, 0xd0, 0x15, 0x74, 0xd4, 0x9d, 0x25,
	0x6c, 0x37, 0xf1, 0xa1, 0x26, 0xdc, 0xfc, 0x6c, 0xe0, 0x33, 0x75, 0x67, 0xe8, 0x5d, 0x41, 0x47,
	0xde, 0x35, 0xef, 0x7b, 0xc3, 0x14, 0x57, 0xa6, 0xd2, 0x98, 0x8e, 0x7f, 0x72, 0xe0, 0xa2, 0x41,
	0x63, 0x4e, 0xd5, 0x7f, 0xc7, 0xa4, 0xf5, 0xa7, 0x4c, 0x0a, 0xe8, 0xd7, 0xdb, 0xf0, 0x17, 0x2b,
	0xf2, 0xd1, 0x21, 0x0f, 0xa4, 0x83, 0xef, 0xaf, 0xd4, 0x96, 0xc9, 0x00, 0xda, 0x4a, 0xac, 0x28,
	0x37, 0x15, 0xe9, 0x63, 0x2b, 0x3c, 0xff, 0xf0, 0x9b, 0x0f, 0xee, 0xff, 0xe6, 0x27, 0xe5, 0xe2,
	0xcc, 0x6c, 0xcc, 0xc7, 0xbf, 0x05, 0x00, 0x00, 0xff, 0xff, 0x96, 0xa2, 0x8d, 0x51, 0x30, 0x08,
	0x00, 0x00,
}