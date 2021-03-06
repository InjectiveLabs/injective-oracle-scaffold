// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/oracle/v1beta1/oracle.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	Relayers []string `protobuf:"bytes,1,rep,name=relayers,proto3" json:"relayers,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8fbf1e7a765423, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetRelayers() []string {
	if m != nil {
		return m.Relayers
	}
	return nil
}

type Ref struct {
	Rate uint64 `protobuf:"varint,1,opt,name=rate,proto3" json:"rate,omitempty"`
	// unix timestamp seconds
	ResolveTime uint64 `protobuf:"varint,2,opt,name=resolve_time,json=resolveTime,proto3" json:"resolve_time,omitempty"`
	Request_ID  uint64 `protobuf:"varint,3,opt,name=request_ID,json=requestID,proto3" json:"request_ID,omitempty"`
}

func (m *Ref) Reset()         { *m = Ref{} }
func (m *Ref) String() string { return proto.CompactTextString(m) }
func (*Ref) ProtoMessage()    {}
func (*Ref) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8fbf1e7a765423, []int{1}
}
func (m *Ref) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ref) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ref.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ref) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ref.Merge(m, src)
}
func (m *Ref) XXX_Size() int {
	return m.Size()
}
func (m *Ref) XXX_DiscardUnknown() {
	xxx_messageInfo_Ref.DiscardUnknown(m)
}

var xxx_messageInfo_Ref proto.InternalMessageInfo

func (m *Ref) GetRate() uint64 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *Ref) GetResolveTime() uint64 {
	if m != nil {
		return m.ResolveTime
	}
	return 0
}

func (m *Ref) GetRequest_ID() uint64 {
	if m != nil {
		return m.Request_ID
	}
	return 0
}

type ReferenceData struct {
	Rate            github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"rate"`
	LastUpdateBase  uint64                                 `protobuf:"varint,2,opt,name=last_update_base,json=lastUpdateBase,proto3" json:"last_update_base,omitempty"`
	LastUpdateQuote uint64                                 `protobuf:"varint,3,opt,name=last_update_quote,json=lastUpdateQuote,proto3" json:"last_update_quote,omitempty"`
}

func (m *ReferenceData) Reset()         { *m = ReferenceData{} }
func (m *ReferenceData) String() string { return proto.CompactTextString(m) }
func (*ReferenceData) ProtoMessage()    {}
func (*ReferenceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8fbf1e7a765423, []int{2}
}
func (m *ReferenceData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReferenceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReferenceData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReferenceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReferenceData.Merge(m, src)
}
func (m *ReferenceData) XXX_Size() int {
	return m.Size()
}
func (m *ReferenceData) XXX_DiscardUnknown() {
	xxx_messageInfo_ReferenceData.DiscardUnknown(m)
}

var xxx_messageInfo_ReferenceData proto.InternalMessageInfo

func (m *ReferenceData) GetLastUpdateBase() uint64 {
	if m != nil {
		return m.LastUpdateBase
	}
	return 0
}

func (m *ReferenceData) GetLastUpdateQuote() uint64 {
	if m != nil {
		return m.LastUpdateQuote
	}
	return 0
}

// Event type upon set ref
type SetRef struct {
	Relayer     string `protobuf:"bytes,1,opt,name=relayer,proto3" json:"relayer,omitempty"`
	Symbol      string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Rate        string `protobuf:"bytes,3,opt,name=rate,proto3" json:"rate,omitempty"`
	ResolveTime string `protobuf:"bytes,4,opt,name=resolve_time,json=resolveTime,proto3" json:"resolve_time,omitempty"`
	RequestId   string `protobuf:"bytes,5,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (m *SetRef) Reset()         { *m = SetRef{} }
func (m *SetRef) String() string { return proto.CompactTextString(m) }
func (*SetRef) ProtoMessage()    {}
func (*SetRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8fbf1e7a765423, []int{3}
}
func (m *SetRef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetRef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRef.Merge(m, src)
}
func (m *SetRef) XXX_Size() int {
	return m.Size()
}
func (m *SetRef) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRef.DiscardUnknown(m)
}

var xxx_messageInfo_SetRef proto.InternalMessageInfo

func (m *SetRef) GetRelayer() string {
	if m != nil {
		return m.Relayer
	}
	return ""
}

func (m *SetRef) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *SetRef) GetRate() string {
	if m != nil {
		return m.Rate
	}
	return ""
}

func (m *SetRef) GetResolveTime() string {
	if m != nil {
		return m.ResolveTime
	}
	return ""
}

func (m *SetRef) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

// Event type for getting reference data
type GetReferenceData struct {
	BaseQuoteSymbol string `protobuf:"bytes,1,opt,name=base_quote_symbol,json=baseQuoteSymbol,proto3" json:"base_quote_symbol,omitempty"`
	BaseQuoteRate   string `protobuf:"bytes,2,opt,name=base_quote_rate,json=baseQuoteRate,proto3" json:"base_quote_rate,omitempty"`
	LastUpdateBase  string `protobuf:"bytes,3,opt,name=last_update_base,json=lastUpdateBase,proto3" json:"last_update_base,omitempty"`
	LastUpdateQuote string `protobuf:"bytes,4,opt,name=last_update_quote,json=lastUpdateQuote,proto3" json:"last_update_quote,omitempty"`
}

func (m *GetReferenceData) Reset()         { *m = GetReferenceData{} }
func (m *GetReferenceData) String() string { return proto.CompactTextString(m) }
func (*GetReferenceData) ProtoMessage()    {}
func (*GetReferenceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8fbf1e7a765423, []int{4}
}
func (m *GetReferenceData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetReferenceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetReferenceData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetReferenceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReferenceData.Merge(m, src)
}
func (m *GetReferenceData) XXX_Size() int {
	return m.Size()
}
func (m *GetReferenceData) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReferenceData.DiscardUnknown(m)
}

var xxx_messageInfo_GetReferenceData proto.InternalMessageInfo

func (m *GetReferenceData) GetBaseQuoteSymbol() string {
	if m != nil {
		return m.BaseQuoteSymbol
	}
	return ""
}

func (m *GetReferenceData) GetBaseQuoteRate() string {
	if m != nil {
		return m.BaseQuoteRate
	}
	return ""
}

func (m *GetReferenceData) GetLastUpdateBase() string {
	if m != nil {
		return m.LastUpdateBase
	}
	return ""
}

func (m *GetReferenceData) GetLastUpdateQuote() string {
	if m != nil {
		return m.LastUpdateQuote
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "injective.oracle.v1beta1.Params")
	proto.RegisterType((*Ref)(nil), "injective.oracle.v1beta1.Ref")
	proto.RegisterType((*ReferenceData)(nil), "injective.oracle.v1beta1.ReferenceData")
	proto.RegisterType((*SetRef)(nil), "injective.oracle.v1beta1.SetRef")
	proto.RegisterType((*GetReferenceData)(nil), "injective.oracle.v1beta1.GetReferenceData")
}

func init() {
	proto.RegisterFile("injective/oracle/v1beta1/oracle.proto", fileDescriptor_1c8fbf1e7a765423)
}

var fileDescriptor_1c8fbf1e7a765423 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0xb6, 0x56, 0xfb, 0x74, 0xed, 0x6e, 0x10, 0x09, 0x0b, 0xa6, 0x6b, 0xc0, 0xa5,
	0x14, 0x9a, 0xb0, 0x78, 0xf3, 0x58, 0x0a, 0x52, 0xf0, 0xa0, 0x59, 0x45, 0xd0, 0x43, 0x98, 0x24,
	0xaf, 0xdd, 0x68, 0xd2, 0x69, 0x67, 0x26, 0x85, 0x7e, 0x8b, 0xfd, 0x08, 0x5e, 0xfd, 0x0e, 0x7e,
	0x80, 0x3d, 0xee, 0x51, 0x3c, 0x2c, 0xd2, 0x5e, 0xfc, 0x18, 0x92, 0x99, 0x69, 0x0c, 0x62, 0x61,
	0x4f, 0x99, 0xf7, 0x7f, 0x7f, 0x98, 0xff, 0xfb, 0xe5, 0x0d, 0x3c, 0x4f, 0xe7, 0x9f, 0x31, 0x96,
	0xe9, 0x0a, 0x7d, 0xc6, 0x69, 0x9c, 0xa1, 0xbf, 0x3a, 0x8b, 0x50, 0xd2, 0x33, 0x53, 0x7a, 0x0b,
	0xce, 0x24, 0xb3, 0xec, 0xca, 0xe6, 0x19, 0xdd, 0xd8, 0x8e, 0x1f, 0xcf, 0xd8, 0x8c, 0x29, 0x93,
	0x5f, 0x9e, 0xb4, 0xdf, 0x1d, 0x40, 0xfb, 0x0d, 0xe5, 0x34, 0x17, 0xd6, 0x31, 0xdc, 0xe7, 0x98,
	0xd1, 0x35, 0x72, 0x61, 0x93, 0x93, 0x66, 0xbf, 0x13, 0x54, 0xf5, 0xcb, 0xd6, 0xef, 0xaf, 0x3d,
	0xe2, 0x7e, 0x82, 0x66, 0x80, 0x53, 0xcb, 0x82, 0x16, 0xa7, 0x12, 0x6d, 0x72, 0x42, 0xfa, 0xad,
	0x40, 0x9d, 0xad, 0x67, 0xf0, 0x90, 0xa3, 0x60, 0xd9, 0x0a, 0x43, 0x99, 0xe6, 0x68, 0xdf, 0x51,
	0xbd, 0x07, 0x46, 0x7b, 0x97, 0xe6, 0x68, 0x3d, 0x05, 0xe0, 0xb8, 0x2c, 0x50, 0xc8, 0x70, 0x32,
	0xb6, 0x9b, 0xca, 0xd0, 0x31, 0xca, 0x64, 0xec, 0x7e, 0x23, 0x70, 0x10, 0xe0, 0x14, 0x39, 0xce,
	0x63, 0x1c, 0x53, 0x49, 0xad, 0x51, 0xed, 0x9e, 0xce, 0xc8, 0xbb, 0xba, 0xe9, 0x35, 0x7e, 0xde,
	0xf4, 0x4e, 0x67, 0xa9, 0xbc, 0x28, 0x22, 0x2f, 0x66, 0xb9, 0x1f, 0x33, 0x91, 0x33, 0x61, 0x3e,
	0x43, 0x91, 0x7c, 0xf1, 0xe5, 0x7a, 0x81, 0xc2, 0x9b, 0xcc, 0xa5, 0xc9, 0xd5, 0x87, 0xc3, 0x8c,
	0x0a, 0x19, 0x16, 0x8b, 0x84, 0x4a, 0x0c, 0x23, 0x2a, 0x76, 0xd9, 0x1e, 0x95, 0xfa, 0x7b, 0x25,
	0x8f, 0xa8, 0x40, 0x6b, 0x00, 0x47, 0x75, 0xe7, 0xb2, 0x60, 0x12, 0x4d, 0xca, 0xee, 0x5f, 0xeb,
	0xdb, 0x52, 0x76, 0x2f, 0x09, 0xb4, 0xcf, 0x51, 0x96, 0x30, 0x6c, 0xb8, 0x67, 0x28, 0xe9, 0x9c,
	0xc1, 0xae, 0xb4, 0x9e, 0x40, 0x5b, 0xac, 0xf3, 0x88, 0x65, 0xea, 0xc2, 0x4e, 0x60, 0xaa, 0x0a,
	0x5f, 0x53, 0xa9, 0xff, 0xc7, 0xd7, 0x52, 0xbd, 0x7d, 0xf8, 0xd2, 0xc4, 0xbe, 0xab, 0x0c, 0x15,
	0xbe, 0xc4, 0xfd, 0x4e, 0xe0, 0xf0, 0x95, 0x8a, 0x54, 0x23, 0x38, 0x80, 0xa3, 0x72, 0x62, 0x3d,
	0x4c, 0x68, 0xd2, 0xe8, 0x98, 0xdd, 0xb2, 0xa1, 0xa6, 0x39, 0xd7, 0xb1, 0x4e, 0xa1, 0x5b, 0xf3,
	0xaa, 0x84, 0x3a, 0xf7, 0x41, 0xe5, 0x0c, 0xf6, 0x11, 0xd5, 0xa3, 0xdc, 0x8a, 0xa8, 0x9e, 0xec,
	0x5f, 0xa2, 0xa3, 0xe5, 0xd5, 0xc6, 0x21, 0xd7, 0x1b, 0x87, 0xfc, 0xda, 0x38, 0xe4, 0x72, 0xeb,
	0x34, 0xae, 0xb7, 0x4e, 0xe3, 0xc7, 0xd6, 0x69, 0x7c, 0xfc, 0x50, 0xfb, 0xdf, 0x93, 0xdd, 0x6e,
	0xbf, 0xa6, 0x91, 0xf0, 0xab, 0x4d, 0x1f, 0xea, 0x4d, 0x1f, 0x8a, 0x98, 0x4e, 0xa7, 0x2c, 0x4b,
	0x6a, 0x9d, 0xf8, 0x82, 0xa6, 0x73, 0x3f, 0x67, 0x49, 0x91, 0xa1, 0xd8, 0x3d, 0x1c, 0xb5, 0x24,
	0x51, 0x5b, 0x3d, 0x80, 0x17, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x13, 0xa0, 0xd8, 0x16, 0x59,
	0x03, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Relayers) != len(that1.Relayers) {
		return false
	}
	for i := range this.Relayers {
		if this.Relayers[i] != that1.Relayers[i] {
			return false
		}
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relayers) > 0 {
		for iNdEx := len(m.Relayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Relayers[iNdEx])
			copy(dAtA[i:], m.Relayers[iNdEx])
			i = encodeVarintOracle(dAtA, i, uint64(len(m.Relayers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Ref) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ref) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ref) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Request_ID != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.Request_ID))
		i--
		dAtA[i] = 0x18
	}
	if m.ResolveTime != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.ResolveTime))
		i--
		dAtA[i] = 0x10
	}
	if m.Rate != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.Rate))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReferenceData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReferenceData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReferenceData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastUpdateQuote != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.LastUpdateQuote))
		i--
		dAtA[i] = 0x18
	}
	if m.LastUpdateBase != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.LastUpdateBase))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *SetRef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetRef) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetRef) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RequestId) > 0 {
		i -= len(m.RequestId)
		copy(dAtA[i:], m.RequestId)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.RequestId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ResolveTime) > 0 {
		i -= len(m.ResolveTime)
		copy(dAtA[i:], m.ResolveTime)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.ResolveTime)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Rate) > 0 {
		i -= len(m.Rate)
		copy(dAtA[i:], m.Rate)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Rate)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Relayer) > 0 {
		i -= len(m.Relayer)
		copy(dAtA[i:], m.Relayer)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Relayer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetReferenceData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetReferenceData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetReferenceData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LastUpdateQuote) > 0 {
		i -= len(m.LastUpdateQuote)
		copy(dAtA[i:], m.LastUpdateQuote)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.LastUpdateQuote)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LastUpdateBase) > 0 {
		i -= len(m.LastUpdateBase)
		copy(dAtA[i:], m.LastUpdateBase)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.LastUpdateBase)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseQuoteRate) > 0 {
		i -= len(m.BaseQuoteRate)
		copy(dAtA[i:], m.BaseQuoteRate)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.BaseQuoteRate)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BaseQuoteSymbol) > 0 {
		i -= len(m.BaseQuoteSymbol)
		copy(dAtA[i:], m.BaseQuoteSymbol)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.BaseQuoteSymbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Relayers) > 0 {
		for _, s := range m.Relayers {
			l = len(s)
			n += 1 + l + sovOracle(uint64(l))
		}
	}
	return n
}

func (m *Ref) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rate != 0 {
		n += 1 + sovOracle(uint64(m.Rate))
	}
	if m.ResolveTime != 0 {
		n += 1 + sovOracle(uint64(m.ResolveTime))
	}
	if m.Request_ID != 0 {
		n += 1 + sovOracle(uint64(m.Request_ID))
	}
	return n
}

func (m *ReferenceData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Rate.Size()
	n += 1 + l + sovOracle(uint64(l))
	if m.LastUpdateBase != 0 {
		n += 1 + sovOracle(uint64(m.LastUpdateBase))
	}
	if m.LastUpdateQuote != 0 {
		n += 1 + sovOracle(uint64(m.LastUpdateQuote))
	}
	return n
}

func (m *SetRef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Relayer)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.Rate)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.ResolveTime)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.RequestId)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func (m *GetReferenceData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BaseQuoteSymbol)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.BaseQuoteRate)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.LastUpdateBase)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.LastUpdateQuote)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayers = append(m.Relayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Ref) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ref: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ref: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			m.Rate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolveTime", wireType)
			}
			m.ResolveTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolveTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request_ID", wireType)
			}
			m.Request_ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Request_ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReferenceData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReferenceData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReferenceData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateBase", wireType)
			}
			m.LastUpdateBase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastUpdateBase |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateQuote", wireType)
			}
			m.LastUpdateQuote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastUpdateQuote |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetRef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetRef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetRef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolveTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResolveTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetReferenceData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetReferenceData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetReferenceData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseQuoteSymbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseQuoteSymbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseQuoteRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseQuoteRate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateBase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastUpdateBase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateQuote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastUpdateQuote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracle = fmt.Errorf("proto: unexpected end of group")
)
