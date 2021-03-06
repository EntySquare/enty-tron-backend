// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.request.proto

package message

import (
	fmt "fmt"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Topic case WITHDRAW:
// -
// WITHDRAW_COLLECT
type WithdrawCollectRequest struct {
	UserId               *string  `protobuf:"bytes,1,req,name=userId" json:"userId,omitempty"`
	RevenueAddress       *string  `protobuf:"bytes,2,req,name=revenueAddress" json:"revenueAddress,omitempty"`
	FromAddress          *string  `protobuf:"bytes,3,req,name=fromAddress" json:"fromAddress,omitempty"`
	CollectStatus        *string  `protobuf:"bytes,4,opt,name=collectStatus,def=0" json:"collectStatus,omitempty"`
	CollectBalance       *string  `protobuf:"bytes,5,opt,name=collectBalance" json:"collectBalance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawCollectRequest) Reset()         { *m = WithdrawCollectRequest{} }
func (m *WithdrawCollectRequest) String() string { return proto.CompactTextString(m) }
func (*WithdrawCollectRequest) ProtoMessage()    {}
func (*WithdrawCollectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f5d5e638555dd4a, []int{0}
}
func (m *WithdrawCollectRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WithdrawCollectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WithdrawCollectRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WithdrawCollectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawCollectRequest.Merge(m, src)
}
func (m *WithdrawCollectRequest) XXX_Size() int {
	return m.Size()
}
func (m *WithdrawCollectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawCollectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawCollectRequest proto.InternalMessageInfo

const Default_WithdrawCollectRequest_CollectStatus string = "0"

func (m *WithdrawCollectRequest) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *WithdrawCollectRequest) GetRevenueAddress() string {
	if m != nil && m.RevenueAddress != nil {
		return *m.RevenueAddress
	}
	return ""
}

func (m *WithdrawCollectRequest) GetFromAddress() string {
	if m != nil && m.FromAddress != nil {
		return *m.FromAddress
	}
	return ""
}

func (m *WithdrawCollectRequest) GetCollectStatus() string {
	if m != nil && m.CollectStatus != nil {
		return *m.CollectStatus
	}
	return Default_WithdrawCollectRequest_CollectStatus
}

func (m *WithdrawCollectRequest) GetCollectBalance() string {
	if m != nil && m.CollectBalance != nil {
		return *m.CollectBalance
	}
	return ""
}

type RiskRequest struct {
	UserId               *string  `protobuf:"bytes,1,req,name=userId" json:"userId,omitempty"`
	OrderId              *int64   `protobuf:"varint,2,req,name=orderId" json:"orderId,omitempty"`
	RevenueAddress       *string  `protobuf:"bytes,3,req,name=revenueAddress" json:"revenueAddress,omitempty"`
	WithdrawBalance      *string  `protobuf:"bytes,4,req,name=withdrawBalance" json:"withdrawBalance,omitempty"`
	Symbol               *string  `protobuf:"bytes,5,opt,name=symbol" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiskRequest) Reset()         { *m = RiskRequest{} }
func (m *RiskRequest) String() string { return proto.CompactTextString(m) }
func (*RiskRequest) ProtoMessage()    {}
func (*RiskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f5d5e638555dd4a, []int{1}
}
func (m *RiskRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RiskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RiskRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RiskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiskRequest.Merge(m, src)
}
func (m *RiskRequest) XXX_Size() int {
	return m.Size()
}
func (m *RiskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RiskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RiskRequest proto.InternalMessageInfo

func (m *RiskRequest) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *RiskRequest) GetOrderId() int64 {
	if m != nil && m.OrderId != nil {
		return *m.OrderId
	}
	return 0
}

func (m *RiskRequest) GetRevenueAddress() string {
	if m != nil && m.RevenueAddress != nil {
		return *m.RevenueAddress
	}
	return ""
}

func (m *RiskRequest) GetWithdrawBalance() string {
	if m != nil && m.WithdrawBalance != nil {
		return *m.WithdrawBalance
	}
	return ""
}

func (m *RiskRequest) GetSymbol() string {
	if m != nil && m.Symbol != nil {
		return *m.Symbol
	}
	return ""
}

type CollectCheck struct {
	OrderId              *int64   `protobuf:"varint,1,req,name=orderId" json:"orderId,omitempty"`
	RemainTxs            *int64   `protobuf:"varint,2,req,name=remainTxs" json:"remainTxs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectCheck) Reset()         { *m = CollectCheck{} }
func (m *CollectCheck) String() string { return proto.CompactTextString(m) }
func (*CollectCheck) ProtoMessage()    {}
func (*CollectCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f5d5e638555dd4a, []int{2}
}
func (m *CollectCheck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollectCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CollectCheck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CollectCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectCheck.Merge(m, src)
}
func (m *CollectCheck) XXX_Size() int {
	return m.Size()
}
func (m *CollectCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectCheck.DiscardUnknown(m)
}

var xxx_messageInfo_CollectCheck proto.InternalMessageInfo

func (m *CollectCheck) GetOrderId() int64 {
	if m != nil && m.OrderId != nil {
		return *m.OrderId
	}
	return 0
}

func (m *CollectCheck) GetRemainTxs() int64 {
	if m != nil && m.RemainTxs != nil {
		return *m.RemainTxs
	}
	return 0
}

type ChainHandlerPanicRequest struct {
	RhError              *string  `protobuf:"bytes,1,req,name=rhError" json:"rhError,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainHandlerPanicRequest) Reset()         { *m = ChainHandlerPanicRequest{} }
func (m *ChainHandlerPanicRequest) String() string { return proto.CompactTextString(m) }
func (*ChainHandlerPanicRequest) ProtoMessage()    {}
func (*ChainHandlerPanicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f5d5e638555dd4a, []int{3}
}
func (m *ChainHandlerPanicRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainHandlerPanicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainHandlerPanicRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainHandlerPanicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainHandlerPanicRequest.Merge(m, src)
}
func (m *ChainHandlerPanicRequest) XXX_Size() int {
	return m.Size()
}
func (m *ChainHandlerPanicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainHandlerPanicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChainHandlerPanicRequest proto.InternalMessageInfo

func (m *ChainHandlerPanicRequest) GetRhError() string {
	if m != nil && m.RhError != nil {
		return *m.RhError
	}
	return ""
}

func init() {
	proto.RegisterType((*WithdrawCollectRequest)(nil), "message.WithdrawCollectRequest")
	proto.RegisterType((*RiskRequest)(nil), "message.RiskRequest")
	proto.RegisterType((*CollectCheck)(nil), "message.CollectCheck")
	proto.RegisterType((*ChainHandlerPanicRequest)(nil), "message.ChainHandlerPanicRequest")
}

func init() { proto.RegisterFile("message.request.proto", fileDescriptor_9f5d5e638555dd4a) }

var fileDescriptor_9f5d5e638555dd4a = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x9d, 0xb6, 0x5a, 0x7a, 0xeb, 0x1f, 0x01, 0xcb, 0x2c, 0x24, 0x94, 0x2c, 0xb4, 0xab,
	0xe2, 0xc2, 0x95, 0x3b, 0x5b, 0x14, 0xdd, 0x49, 0x14, 0x5c, 0x8f, 0x99, 0xab, 0x09, 0x4d, 0x66,
	0xf4, 0xce, 0xc4, 0xea, 0x9b, 0xf8, 0x0a, 0xbe, 0x89, 0xe0, 0xc6, 0x47, 0x90, 0xfa, 0x22, 0x92,
	0x74, 0x82, 0x6d, 0x28, 0xb8, 0x3c, 0xdf, 0x1c, 0x98, 0xef, 0x5c, 0xd8, 0xcb, 0xd0, 0x18, 0xf1,
	0x80, 0x43, 0xc2, 0xa7, 0x1c, 0x8d, 0x1d, 0x3e, 0x92, 0xb6, 0xda, 0x6b, 0x3b, 0x1c, 0x7c, 0x32,
	0xe8, 0xdd, 0x26, 0x36, 0x96, 0x24, 0xa6, 0x63, 0x9d, 0xa6, 0x18, 0xd9, 0x70, 0xde, 0xf4, 0x7a,
	0xb0, 0x91, 0x1b, 0xa4, 0x4b, 0xc9, 0x59, 0xbf, 0x31, 0xe8, 0x84, 0x2e, 0x79, 0x07, 0xb0, 0x4d,
	0xf8, 0x8c, 0x2a, 0xc7, 0x53, 0x29, 0x09, 0x8d, 0xe1, 0x8d, 0xf2, 0xbd, 0x46, 0xbd, 0x3e, 0x74,
	0xef, 0x49, 0x67, 0x55, 0xa9, 0x59, 0x96, 0x16, 0x91, 0x77, 0x08, 0x5b, 0xd1, 0xfc, 0xcf, 0x6b,
	0x2b, 0x6c, 0x6e, 0x78, 0xab, 0xcf, 0x06, 0x9d, 0x13, 0x76, 0x14, 0x2e, 0xf3, 0xe2, 0x4b, 0x07,
	0x46, 0x22, 0x15, 0x2a, 0x42, 0xbe, 0x5e, 0x34, 0xc3, 0x1a, 0x0d, 0xde, 0x19, 0x74, 0xc3, 0xc4,
	0x4c, 0xfe, 0x9b, 0xc0, 0xa1, 0xad, 0x49, 0x96, 0x0f, 0x85, 0x7b, 0x33, 0xac, 0xe2, 0x8a, 0x71,
	0xcd, 0x95, 0xe3, 0x06, 0xb0, 0x33, 0x75, 0x67, 0xab, 0x94, 0x5a, 0x65, 0xb1, 0x8e, 0x0b, 0x07,
	0xf3, 0x9a, 0xdd, 0xe9, 0xd4, 0x39, 0xbb, 0x14, 0x9c, 0xc3, 0xa6, 0x3b, 0xf8, 0x38, 0xc6, 0x68,
	0xb2, 0xe8, 0xc4, 0x96, 0x9d, 0xf6, 0xa1, 0x43, 0x98, 0x89, 0x44, 0xdd, 0xbc, 0x18, 0xe7, 0xfb,
	0x07, 0x82, 0x63, 0xe0, 0xe3, 0x58, 0x24, 0xea, 0x42, 0x28, 0x99, 0x22, 0x5d, 0x09, 0x95, 0x44,
	0xd5, 0x7e, 0x0e, 0x6d, 0x8a, 0xcf, 0x88, 0x34, 0xb9, 0x03, 0x54, 0x71, 0xb4, 0xfb, 0x31, 0xf3,
	0xd9, 0xd7, 0xcc, 0x67, 0xdf, 0x33, 0x9f, 0xbd, 0xfd, 0xf8, 0x6b, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xec, 0x72, 0x83, 0x90, 0x2a, 0x02, 0x00, 0x00,
}

func (m *WithdrawCollectRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WithdrawCollectRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WithdrawCollectRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CollectBalance != nil {
		i -= len(*m.CollectBalance)
		copy(dAtA[i:], *m.CollectBalance)
		i = encodeVarintMessageRequest(dAtA, i, uint64(len(*m.CollectBalance)))
		i--
		dAtA[i] = 0x2a
	}
	if m.CollectStatus != nil {
		i -= len(*m.CollectStatus)
		copy(dAtA[i:], *m.CollectStatus)
		i = encodeVarintMessageRequest(dAtA, i, uint64(len(*m.CollectStatus)))
		i--
		dAtA[i] = 0x22
	}
	if m.FromAddress == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(*m.FromAddress)
		copy(dAtA[i:], *m.FromAddress)
		i = encodeVarintMessageRequest(dAtA, i, uint64(len(*m.FromAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.RevenueAddress == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(*m.RevenueAddress)
		copy(dAtA[i:], *m.RevenueAddress)
		i = encodeVarintMessageRequest(dAtA, i, uint64(len(*m.RevenueAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.UserId == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(*m.UserId)
		copy(dAtA[i:], *m.UserId)
		i = encodeVarintMessageRequest(dAtA, i, uint64(len(*m.UserId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RiskRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RiskRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RiskRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Symbol != nil {
		i -= len(*m.Symbol)
		copy(dAtA[i:], *m.Symbol)
		i = encodeVarintMessageRequest(dAtA, i, uint64(len(*m.Symbol)))
		i--
		dAtA[i] = 0x2a
	}
	if m.WithdrawBalance == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(*m.WithdrawBalance)
		copy(dAtA[i:], *m.WithdrawBalance)
		i = encodeVarintMessageRequest(dAtA, i, uint64(len(*m.WithdrawBalance)))
		i--
		dAtA[i] = 0x22
	}
	if m.RevenueAddress == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(*m.RevenueAddress)
		copy(dAtA[i:], *m.RevenueAddress)
		i = encodeVarintMessageRequest(dAtA, i, uint64(len(*m.RevenueAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.OrderId == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i = encodeVarintMessageRequest(dAtA, i, uint64(*m.OrderId))
		i--
		dAtA[i] = 0x10
	}
	if m.UserId == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(*m.UserId)
		copy(dAtA[i:], *m.UserId)
		i = encodeVarintMessageRequest(dAtA, i, uint64(len(*m.UserId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CollectCheck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectCheck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollectCheck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RemainTxs == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i = encodeVarintMessageRequest(dAtA, i, uint64(*m.RemainTxs))
		i--
		dAtA[i] = 0x10
	}
	if m.OrderId == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i = encodeVarintMessageRequest(dAtA, i, uint64(*m.OrderId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ChainHandlerPanicRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainHandlerPanicRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainHandlerPanicRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RhError == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(*m.RhError)
		copy(dAtA[i:], *m.RhError)
		i = encodeVarintMessageRequest(dAtA, i, uint64(len(*m.RhError)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessageRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessageRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WithdrawCollectRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserId != nil {
		l = len(*m.UserId)
		n += 1 + l + sovMessageRequest(uint64(l))
	}
	if m.RevenueAddress != nil {
		l = len(*m.RevenueAddress)
		n += 1 + l + sovMessageRequest(uint64(l))
	}
	if m.FromAddress != nil {
		l = len(*m.FromAddress)
		n += 1 + l + sovMessageRequest(uint64(l))
	}
	if m.CollectStatus != nil {
		l = len(*m.CollectStatus)
		n += 1 + l + sovMessageRequest(uint64(l))
	}
	if m.CollectBalance != nil {
		l = len(*m.CollectBalance)
		n += 1 + l + sovMessageRequest(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RiskRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserId != nil {
		l = len(*m.UserId)
		n += 1 + l + sovMessageRequest(uint64(l))
	}
	if m.OrderId != nil {
		n += 1 + sovMessageRequest(uint64(*m.OrderId))
	}
	if m.RevenueAddress != nil {
		l = len(*m.RevenueAddress)
		n += 1 + l + sovMessageRequest(uint64(l))
	}
	if m.WithdrawBalance != nil {
		l = len(*m.WithdrawBalance)
		n += 1 + l + sovMessageRequest(uint64(l))
	}
	if m.Symbol != nil {
		l = len(*m.Symbol)
		n += 1 + l + sovMessageRequest(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CollectCheck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderId != nil {
		n += 1 + sovMessageRequest(uint64(*m.OrderId))
	}
	if m.RemainTxs != nil {
		n += 1 + sovMessageRequest(uint64(*m.RemainTxs))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ChainHandlerPanicRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RhError != nil {
		l = len(*m.RhError)
		n += 1 + l + sovMessageRequest(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessageRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessageRequest(x uint64) (n int) {
	return sovMessageRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WithdrawCollectRequest) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageRequest
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
			return fmt.Errorf("proto: WithdrawCollectRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WithdrawCollectRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
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
				return ErrInvalidLengthMessageRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.UserId = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevenueAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
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
				return ErrInvalidLengthMessageRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.RevenueAddress = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
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
				return ErrInvalidLengthMessageRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.FromAddress = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
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
				return ErrInvalidLengthMessageRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.CollectStatus = &s
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
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
				return ErrInvalidLengthMessageRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.CollectBalance = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessageRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RiskRequest) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageRequest
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
			return fmt.Errorf("proto: RiskRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RiskRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
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
				return ErrInvalidLengthMessageRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.UserId = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OrderId = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevenueAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
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
				return ErrInvalidLengthMessageRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.RevenueAddress = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
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
				return ErrInvalidLengthMessageRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.WithdrawBalance = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
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
				return ErrInvalidLengthMessageRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Symbol = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessageRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CollectCheck) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageRequest
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
			return fmt.Errorf("proto: CollectCheck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectCheck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OrderId = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainTxs", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RemainTxs = &v
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipMessageRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChainHandlerPanicRequest) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageRequest
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
			return fmt.Errorf("proto: ChainHandlerPanicRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainHandlerPanicRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RhError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageRequest
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
				return ErrInvalidLengthMessageRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.RhError = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMessageRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessageRequest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessageRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessageRequest
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
					return 0, ErrIntOverflowMessageRequest
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
					return 0, ErrIntOverflowMessageRequest
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
				return 0, ErrInvalidLengthMessageRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessageRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessageRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessageRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessageRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessageRequest = fmt.Errorf("proto: unexpected end of group")
)
