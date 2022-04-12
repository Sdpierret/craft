// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: craft/exp/v1beta1/exp_request.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RequestStatus int32

const (
	// REQUEST_STATUS_COMPLETE defines a request status completed.
	StatusCompleteRequest RequestStatus = 0
	// REQUEST_STATUS_PERIOD defines a request status current in process
	StatusOnGoingRequest RequestStatus = 1
	// REQUEST_STATUS_EXPIRED defines the expired request status.
	StatusExpiredRequest RequestStatus = 2
	// REQUEST_STATUS_EXPIRED defines the request status that member don't fund to exp module.
	StatusNoFundRequest RequestStatus = 3
)

var RequestStatus_name = map[int32]string{
	0: "REQUEST_STATUS_COMPLETE",
	1: "REQUEST_STATUS_ONGOING",
	2: "REQUEST_STATUS_EXPIRED",
	3: "REQUEST_STATUS_NOFUND",
}

var RequestStatus_value = map[string]int32{
	"REQUEST_STATUS_COMPLETE": 0,
	"REQUEST_STATUS_ONGOING":  1,
	"REQUEST_STATUS_EXPIRED":  2,
	"REQUEST_STATUS_NOFUND":   3,
}

func (x RequestStatus) String() string {
	return proto.EnumName(RequestStatus_name, int32(x))
}

func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94f58b9d8d17be18, []int{0}
}

type BurnRequest struct {
	Account       string        `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	BurnTokenLeft *types.Coin   `protobuf:"bytes,2,opt,name=burn_token_left,json=burnTokenLeft,proto3" json:"burn_token_left,omitempty"`
	RequestTime   time.Time     `protobuf:"bytes,3,opt,name=request_time,json=requestTime,proto3,stdtime" json:"request_time"`
	Status        RequestStatus `protobuf:"varint,4,opt,name=status,proto3,enum=craft.exp.v1beta1.RequestStatus" json:"status,omitempty" yaml:"request_status"`
}

func (m *BurnRequest) Reset()         { *m = BurnRequest{} }
func (m *BurnRequest) String() string { return proto.CompactTextString(m) }
func (*BurnRequest) ProtoMessage()    {}
func (*BurnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94f58b9d8d17be18, []int{0}
}
func (m *BurnRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BurnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BurnRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BurnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurnRequest.Merge(m, src)
}
func (m *BurnRequest) XXX_Size() int {
	return m.Size()
}
func (m *BurnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BurnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BurnRequest proto.InternalMessageInfo

func (m *BurnRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *BurnRequest) GetBurnTokenLeft() *types.Coin {
	if m != nil {
		return m.BurnTokenLeft
	}
	return nil
}

func (m *BurnRequest) GetRequestTime() time.Time {
	if m != nil {
		return m.RequestTime
	}
	return time.Time{}
}

func (m *BurnRequest) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return StatusCompleteRequest
}

type MintRequest struct {
	Account        string                                 `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	DaoTokenMinted github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=dao_token_minted,json=daoTokenMinted,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"dao_token_minted"`
	DaoTokenLeft   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=dao_token_left,json=daoTokenLeft,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"dao_token_left"`
	RequestTime    time.Time                              `protobuf:"bytes,4,opt,name=request_time,json=requestTime,proto3,stdtime" json:"request_time"`
	Status         RequestStatus                          `protobuf:"varint,5,opt,name=status,proto3,enum=craft.exp.v1beta1.RequestStatus" json:"status,omitempty" yaml:"request_status"`
}

func (m *MintRequest) Reset()         { *m = MintRequest{} }
func (m *MintRequest) String() string { return proto.CompactTextString(m) }
func (*MintRequest) ProtoMessage()    {}
func (*MintRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94f58b9d8d17be18, []int{1}
}
func (m *MintRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintRequest.Merge(m, src)
}
func (m *MintRequest) XXX_Size() int {
	return m.Size()
}
func (m *MintRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MintRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MintRequest proto.InternalMessageInfo

func (m *MintRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *MintRequest) GetRequestTime() time.Time {
	if m != nil {
		return m.RequestTime
	}
	return time.Time{}
}

func (m *MintRequest) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return StatusCompleteRequest
}

type MintRequestList struct {
	MintRequestList []*MintRequest `protobuf:"bytes,1,rep,name=mint_request_list,json=mintRequestList,proto3" json:"mint_request_list,omitempty" yaml:"mint_request_list"`
}

func (m *MintRequestList) Reset()         { *m = MintRequestList{} }
func (m *MintRequestList) String() string { return proto.CompactTextString(m) }
func (*MintRequestList) ProtoMessage()    {}
func (*MintRequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_94f58b9d8d17be18, []int{2}
}
func (m *MintRequestList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintRequestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintRequestList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintRequestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintRequestList.Merge(m, src)
}
func (m *MintRequestList) XXX_Size() int {
	return m.Size()
}
func (m *MintRequestList) XXX_DiscardUnknown() {
	xxx_messageInfo_MintRequestList.DiscardUnknown(m)
}

var xxx_messageInfo_MintRequestList proto.InternalMessageInfo

func (m *MintRequestList) GetMintRequestList() []*MintRequest {
	if m != nil {
		return m.MintRequestList
	}
	return nil
}

type BurnRequestList struct {
	BurnRequestList []*BurnRequest `protobuf:"bytes,2,rep,name=burn_request_list,json=burnRequestList,proto3" json:"burn_request_list,omitempty" yaml:"burn_request_list"`
}

func (m *BurnRequestList) Reset()         { *m = BurnRequestList{} }
func (m *BurnRequestList) String() string { return proto.CompactTextString(m) }
func (*BurnRequestList) ProtoMessage()    {}
func (*BurnRequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_94f58b9d8d17be18, []int{3}
}
func (m *BurnRequestList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BurnRequestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BurnRequestList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BurnRequestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurnRequestList.Merge(m, src)
}
func (m *BurnRequestList) XXX_Size() int {
	return m.Size()
}
func (m *BurnRequestList) XXX_DiscardUnknown() {
	xxx_messageInfo_BurnRequestList.DiscardUnknown(m)
}

var xxx_messageInfo_BurnRequestList proto.InternalMessageInfo

func (m *BurnRequestList) GetBurnRequestList() []*BurnRequest {
	if m != nil {
		return m.BurnRequestList
	}
	return nil
}

func init() {
	proto.RegisterEnum("craft.exp.v1beta1.RequestStatus", RequestStatus_name, RequestStatus_value)
	proto.RegisterType((*BurnRequest)(nil), "craft.exp.v1beta1.BurnRequest")
	proto.RegisterType((*MintRequest)(nil), "craft.exp.v1beta1.MintRequest")
	proto.RegisterType((*MintRequestList)(nil), "craft.exp.v1beta1.MintRequestList")
	proto.RegisterType((*BurnRequestList)(nil), "craft.exp.v1beta1.BurnRequestList")
}

func init() {
	proto.RegisterFile("craft/exp/v1beta1/exp_request.proto", fileDescriptor_94f58b9d8d17be18)
}

var fileDescriptor_94f58b9d8d17be18 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x4f, 0xd3, 0x50,
	0x1c, 0x5f, 0x37, 0x44, 0x78, 0xe3, 0xc7, 0xa8, 0x20, 0x5d, 0x63, 0xda, 0x66, 0x46, 0xb3, 0x98,
	0xf8, 0x1a, 0xa6, 0xf1, 0xc0, 0x8d, 0x41, 0x59, 0x48, 0x60, 0xc3, 0xae, 0x24, 0xc4, 0x4b, 0xd3,
	0x75, 0x6f, 0xa3, 0x61, 0xed, 0xab, 0xeb, 0xab, 0x19, 0xf1, 0x1f, 0x20, 0x9c, 0x88, 0x27, 0x2f,
	0x24, 0x24, 0xfe, 0x33, 0x1c, 0x39, 0x1a, 0x0f, 0xd3, 0xc0, 0xd9, 0x0b, 0xff, 0x80, 0xa6, 0xef,
	0xb5, 0x5a, 0x18, 0xc4, 0x44, 0xe3, 0x69, 0x7d, 0xed, 0xe7, 0xd7, 0xfb, 0x7c, 0xdf, 0x1b, 0x78,
	0x6c, 0xf7, 0xad, 0x0e, 0x51, 0xd1, 0xc0, 0x57, 0xdf, 0x2d, 0xb5, 0x10, 0xb1, 0x96, 0xa2, 0x67,
	0xb3, 0x8f, 0xde, 0x86, 0x28, 0x20, 0xd0, 0xef, 0x63, 0x82, 0xf9, 0x39, 0x0a, 0x82, 0x68, 0xe0,
	0xc3, 0x18, 0x24, 0x16, 0x6d, 0x1c, 0xb8, 0x38, 0x30, 0x29, 0x40, 0x65, 0x0b, 0x86, 0x16, 0xe7,
	0xbb, 0xb8, 0x8b, 0xd9, 0xfb, 0xe8, 0x29, 0x7e, 0x2b, 0x31, 0x8c, 0xda, 0xb2, 0x02, 0xf4, 0xcb,
	0xca, 0xc6, 0x8e, 0x17, 0x7f, 0x97, 0xbb, 0x18, 0x77, 0x7b, 0x48, 0xa5, 0xab, 0x56, 0xd8, 0x51,
	0x89, 0xe3, 0xa2, 0x80, 0x58, 0xae, 0xcf, 0x00, 0xa5, 0x0f, 0x59, 0x90, 0xaf, 0x86, 0x7d, 0x4f,
	0x67, 0xd1, 0x78, 0x01, 0xdc, 0xb7, 0x6c, 0x1b, 0x87, 0x1e, 0x11, 0x38, 0x85, 0x2b, 0x4f, 0xea,
	0xc9, 0x92, 0x5f, 0x01, 0xb3, 0xad, 0xb0, 0xef, 0x99, 0x04, 0xef, 0x23, 0xcf, 0xec, 0xa1, 0x0e,
	0x11, 0xb2, 0x0a, 0x57, 0xce, 0x57, 0x8a, 0x30, 0x0e, 0x1a, 0x85, 0x48, 0xb6, 0x02, 0x57, 0xb1,
	0xe3, 0xe9, 0xd3, 0x11, 0xc3, 0x88, 0x08, 0x9b, 0xa8, 0x43, 0xf8, 0x1a, 0x98, 0x8a, 0x2b, 0x30,
	0xa3, 0x1c, 0x42, 0x8e, 0xf2, 0x45, 0xc8, 0x42, 0xc2, 0x24, 0x24, 0x34, 0x92, 0x90, 0xd5, 0x89,
	0xb3, 0xa1, 0x9c, 0x39, 0xfe, 0x2a, 0x73, 0x7a, 0x3e, 0x66, 0x46, 0xdf, 0xf8, 0x26, 0x18, 0x0f,
	0x88, 0x45, 0xc2, 0x40, 0x18, 0x53, 0xb8, 0xf2, 0x4c, 0x45, 0x81, 0x23, 0x5d, 0xc2, 0x78, 0x47,
	0x4d, 0x8a, 0xab, 0x16, 0xaf, 0x86, 0xf2, 0xc2, 0x81, 0xe5, 0xf6, 0x96, 0x4b, 0x49, 0x04, 0xa6,
	0x50, 0xd2, 0x63, 0xa9, 0xe5, 0x89, 0xc3, 0x53, 0x99, 0xfb, 0x78, 0x2a, 0x73, 0xa5, 0x1f, 0x59,
	0x90, 0xdf, 0x72, 0x3c, 0xf2, 0xe7, 0x52, 0x76, 0x41, 0xa1, 0x6d, 0xe1, 0xb8, 0x13, 0xd7, 0xf1,
	0x08, 0x6a, 0xd3, 0x56, 0x26, 0xab, 0x30, 0x4a, 0xfe, 0x65, 0x28, 0x3f, 0xed, 0x3a, 0x64, 0x2f,
	0x6c, 0x41, 0x1b, 0xbb, 0xf1, 0x40, 0xe3, 0x9f, 0xe7, 0x41, 0x7b, 0x5f, 0x25, 0x07, 0x3e, 0x0a,
	0xe0, 0x1a, 0xb2, 0xf5, 0x99, 0xb6, 0x85, 0x69, 0x53, 0x5b, 0x54, 0x85, 0x37, 0xc0, 0xcc, 0x6f,
	0x65, 0xda, 0x76, 0xee, 0xaf, 0x74, 0xa7, 0x12, 0xdd, 0x5b, 0x27, 0x30, 0xf6, 0xef, 0x13, 0xb8,
	0xf7, 0x3f, 0x26, 0xf0, 0x1e, 0xcc, 0xa6, 0x06, 0xb0, 0xe9, 0x04, 0x84, 0xdf, 0x03, 0x73, 0x51,
	0xc1, 0xc9, 0x25, 0x32, 0x7b, 0x4e, 0x10, 0x8d, 0x23, 0x57, 0xce, 0x57, 0xa4, 0x5b, 0xcc, 0x53,
	0xf4, 0xea, 0xa3, 0xab, 0xa1, 0x2c, 0x30, 0xeb, 0x11, 0x89, 0x92, 0x3e, 0xeb, 0x5e, 0x77, 0x8a,
	0xcc, 0x53, 0x57, 0x22, 0x31, 0xa7, 0x87, 0xff, 0x9a, 0x79, 0xf6, 0x4e, 0xf3, 0x14, 0x3d, 0x6d,
	0x3e, 0x22, 0x51, 0xd2, 0xe9, 0x9d, 0x4a, 0x39, 0x3d, 0xfb, 0xce, 0x81, 0xe9, 0x6b, 0xc5, 0xf1,
	0xaf, 0xc0, 0xa2, 0xae, 0xbd, 0xde, 0xd1, 0x9a, 0x86, 0xd9, 0x34, 0x56, 0x8c, 0x9d, 0xa6, 0xb9,
	0xda, 0xd8, 0xda, 0xde, 0xd4, 0x0c, 0xad, 0x90, 0x11, 0x8b, 0x47, 0x27, 0xca, 0x02, 0x03, 0xae,
	0x62, 0xd7, 0xef, 0x21, 0x82, 0x92, 0x53, 0xfb, 0x12, 0x3c, 0xbc, 0xc1, 0x6b, 0xd4, 0x6b, 0x8d,
	0x8d, 0x7a, 0xad, 0xc0, 0x89, 0xc2, 0xd1, 0x89, 0x32, 0xcf, 0x68, 0x0d, 0xaf, 0x86, 0x1d, 0xaf,
	0x7b, 0x37, 0x4b, 0xdb, 0xdd, 0xde, 0xd0, 0xb5, 0xb5, 0x42, 0x36, 0xcd, 0xd2, 0x06, 0xbe, 0xd3,
	0x47, 0xed, 0x84, 0x55, 0x01, 0x0b, 0x37, 0x58, 0xf5, 0xc6, 0xfa, 0x4e, 0x7d, 0xad, 0x90, 0x13,
	0x17, 0x8f, 0x4e, 0x94, 0x07, 0x8c, 0x54, 0xc7, 0xeb, 0xa1, 0x97, 0x70, 0xc4, 0xb1, 0xc3, 0x4f,
	0x52, 0xa6, 0xfa, 0xe4, 0xec, 0x42, 0xe2, 0xce, 0x2f, 0x24, 0xee, 0xdb, 0x85, 0xc4, 0x1d, 0x5f,
	0x4a, 0x99, 0xf3, 0x4b, 0x29, 0xf3, 0xf9, 0x52, 0xca, 0xbc, 0xc9, 0x0f, 0xe8, 0x1f, 0x28, 0x3d,
	0xca, 0xad, 0x71, 0x7a, 0x34, 0x5f, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x9f, 0x94, 0x12,
	0x5a, 0x05, 0x00, 0x00,
}

func (m *BurnRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BurnRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BurnRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintExpRequest(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.RequestTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.RequestTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintExpRequest(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if m.BurnTokenLeft != nil {
		{
			size, err := m.BurnTokenLeft.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExpRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintExpRequest(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MintRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintExpRequest(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.RequestTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.RequestTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintExpRequest(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	{
		size := m.DaoTokenLeft.Size()
		i -= size
		if _, err := m.DaoTokenLeft.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExpRequest(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.DaoTokenMinted.Size()
		i -= size
		if _, err := m.DaoTokenMinted.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExpRequest(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintExpRequest(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MintRequestList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintRequestList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintRequestList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MintRequestList) > 0 {
		for iNdEx := len(m.MintRequestList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintRequestList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExpRequest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BurnRequestList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BurnRequestList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BurnRequestList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BurnRequestList) > 0 {
		for iNdEx := len(m.BurnRequestList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BurnRequestList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExpRequest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintExpRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovExpRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BurnRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovExpRequest(uint64(l))
	}
	if m.BurnTokenLeft != nil {
		l = m.BurnTokenLeft.Size()
		n += 1 + l + sovExpRequest(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.RequestTime)
	n += 1 + l + sovExpRequest(uint64(l))
	if m.Status != 0 {
		n += 1 + sovExpRequest(uint64(m.Status))
	}
	return n
}

func (m *MintRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovExpRequest(uint64(l))
	}
	l = m.DaoTokenMinted.Size()
	n += 1 + l + sovExpRequest(uint64(l))
	l = m.DaoTokenLeft.Size()
	n += 1 + l + sovExpRequest(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.RequestTime)
	n += 1 + l + sovExpRequest(uint64(l))
	if m.Status != 0 {
		n += 1 + sovExpRequest(uint64(m.Status))
	}
	return n
}

func (m *MintRequestList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MintRequestList) > 0 {
		for _, e := range m.MintRequestList {
			l = e.Size()
			n += 1 + l + sovExpRequest(uint64(l))
		}
	}
	return n
}

func (m *BurnRequestList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BurnRequestList) > 0 {
		for _, e := range m.BurnRequestList {
			l = e.Size()
			n += 1 + l + sovExpRequest(uint64(l))
		}
	}
	return n
}

func sovExpRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExpRequest(x uint64) (n int) {
	return sovExpRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BurnRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpRequest
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
			return fmt.Errorf("proto: BurnRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BurnRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpRequest
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
				return ErrInvalidLengthExpRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExpRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnTokenLeft", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExpRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExpRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BurnTokenLeft == nil {
				m.BurnTokenLeft = &types.Coin{}
			}
			if err := m.BurnTokenLeft.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExpRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExpRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.RequestTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= RequestStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExpRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExpRequest
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
func (m *MintRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpRequest
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
			return fmt.Errorf("proto: MintRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpRequest
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
				return ErrInvalidLengthExpRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExpRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaoTokenMinted", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpRequest
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
				return ErrInvalidLengthExpRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExpRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DaoTokenMinted.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaoTokenLeft", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpRequest
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
				return ErrInvalidLengthExpRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExpRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DaoTokenLeft.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExpRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExpRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.RequestTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= RequestStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExpRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExpRequest
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
func (m *MintRequestList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpRequest
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
			return fmt.Errorf("proto: MintRequestList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintRequestList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRequestList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExpRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExpRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintRequestList = append(m.MintRequestList, &MintRequest{})
			if err := m.MintRequestList[len(m.MintRequestList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExpRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExpRequest
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
func (m *BurnRequestList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpRequest
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
			return fmt.Errorf("proto: BurnRequestList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BurnRequestList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnRequestList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExpRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExpRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnRequestList = append(m.BurnRequestList, &BurnRequest{})
			if err := m.BurnRequestList[len(m.BurnRequestList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExpRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExpRequest
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
func skipExpRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExpRequest
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
					return 0, ErrIntOverflowExpRequest
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
					return 0, ErrIntOverflowExpRequest
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
				return 0, ErrInvalidLengthExpRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExpRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExpRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExpRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExpRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExpRequest = fmt.Errorf("proto: unexpected end of group")
)
