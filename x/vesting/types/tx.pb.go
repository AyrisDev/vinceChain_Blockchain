// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vince/vesting/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgCreateClawbackVestingAccount defines a message that enables creating a
// ClawbackVestingAccount.
type MsgCreateClawbackVestingAccount struct {
	// from_address specifies the account to provide the funds and sign the
	// clawback request
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	// to_address specifies the account to receive the funds
	ToAddress string `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// start_time defines the time at which the vesting period begins
	StartTime time.Time `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// lockup_periods defines the unlocking schedule relative to the start_time
	LockupPeriods []types.Period `protobuf:"bytes,4,rep,name=lockup_periods,json=lockupPeriods,proto3" json:"lockup_periods"`
	// vesting_periods defines thevesting schedule relative to the start_time
	VestingPeriods []types.Period `protobuf:"bytes,5,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods"`
	// merge specifies a the creation mechanism for existing
	// ClawbackVestingAccounts. If true, merge this new grant into an existing
	// ClawbackVestingAccount, or create it if it does not exist. If false,
	// creates a new account. New grants to an existing account must be from the
	// same from_address.
	Merge bool `protobuf:"varint,6,opt,name=merge,proto3" json:"merge,omitempty"`
}

func (m *MsgCreateClawbackVestingAccount) Reset()         { *m = MsgCreateClawbackVestingAccount{} }
func (m *MsgCreateClawbackVestingAccount) String() string { return proto.CompactTextString(m) }
func (*MsgCreateClawbackVestingAccount) ProtoMessage()    {}
func (*MsgCreateClawbackVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e70a71da0ce69ab, []int{0}
}
func (m *MsgCreateClawbackVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateClawbackVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateClawbackVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateClawbackVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateClawbackVestingAccount.Merge(m, src)
}
func (m *MsgCreateClawbackVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateClawbackVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateClawbackVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateClawbackVestingAccount proto.InternalMessageInfo

func (m *MsgCreateClawbackVestingAccount) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgCreateClawbackVestingAccount) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgCreateClawbackVestingAccount) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *MsgCreateClawbackVestingAccount) GetLockupPeriods() []types.Period {
	if m != nil {
		return m.LockupPeriods
	}
	return nil
}

func (m *MsgCreateClawbackVestingAccount) GetVestingPeriods() []types.Period {
	if m != nil {
		return m.VestingPeriods
	}
	return nil
}

func (m *MsgCreateClawbackVestingAccount) GetMerge() bool {
	if m != nil {
		return m.Merge
	}
	return false
}

// MsgCreateClawbackVestingAccountResponse defines the
// MsgCreateClawbackVestingAccount response type.
type MsgCreateClawbackVestingAccountResponse struct {
}

func (m *MsgCreateClawbackVestingAccountResponse) Reset() {
	*m = MsgCreateClawbackVestingAccountResponse{}
}
func (m *MsgCreateClawbackVestingAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateClawbackVestingAccountResponse) ProtoMessage()    {}
func (*MsgCreateClawbackVestingAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e70a71da0ce69ab, []int{1}
}
func (m *MsgCreateClawbackVestingAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateClawbackVestingAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateClawbackVestingAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateClawbackVestingAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateClawbackVestingAccountResponse.Merge(m, src)
}
func (m *MsgCreateClawbackVestingAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateClawbackVestingAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateClawbackVestingAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateClawbackVestingAccountResponse proto.InternalMessageInfo

// MsgClawback defines a message that removes unvested tokens from a
// ClawbackVestingAccount.
type MsgClawback struct {
	// funder_address is the address which funded the account
	FunderAddress string `protobuf:"bytes,1,opt,name=funder_address,json=funderAddress,proto3" json:"funder_address,omitempty"`
	// account_address is the address of the ClawbackVestingAccount to claw back
	// from.
	AccountAddress string `protobuf:"bytes,2,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
	// dest_address specifies where the clawed-back tokens should be transferred
	// to. If empty, the tokens will be transferred back to the original funder of
	// the account.
	DestAddress string `protobuf:"bytes,3,opt,name=dest_address,json=destAddress,proto3" json:"dest_address,omitempty"`
}

func (m *MsgClawback) Reset()         { *m = MsgClawback{} }
func (m *MsgClawback) String() string { return proto.CompactTextString(m) }
func (*MsgClawback) ProtoMessage()    {}
func (*MsgClawback) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e70a71da0ce69ab, []int{2}
}
func (m *MsgClawback) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClawback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClawback.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClawback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClawback.Merge(m, src)
}
func (m *MsgClawback) XXX_Size() int {
	return m.Size()
}
func (m *MsgClawback) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClawback.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClawback proto.InternalMessageInfo

func (m *MsgClawback) GetFunderAddress() string {
	if m != nil {
		return m.FunderAddress
	}
	return ""
}

func (m *MsgClawback) GetAccountAddress() string {
	if m != nil {
		return m.AccountAddress
	}
	return ""
}

func (m *MsgClawback) GetDestAddress() string {
	if m != nil {
		return m.DestAddress
	}
	return ""
}

// MsgClawbackResponse defines the MsgClawback response type.
type MsgClawbackResponse struct {
}

func (m *MsgClawbackResponse) Reset()         { *m = MsgClawbackResponse{} }
func (m *MsgClawbackResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClawbackResponse) ProtoMessage()    {}
func (*MsgClawbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e70a71da0ce69ab, []int{3}
}
func (m *MsgClawbackResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClawbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClawbackResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClawbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClawbackResponse.Merge(m, src)
}
func (m *MsgClawbackResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClawbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClawbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClawbackResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateClawbackVestingAccount)(nil), "vince.vesting.v1.MsgCreateClawbackVestingAccount")
	proto.RegisterType((*MsgCreateClawbackVestingAccountResponse)(nil), "vince.vesting.v1.MsgCreateClawbackVestingAccountResponse")
	proto.RegisterType((*MsgClawback)(nil), "vince.vesting.v1.MsgClawback")
	proto.RegisterType((*MsgClawbackResponse)(nil), "vince.vesting.v1.MsgClawbackResponse")
}

func init() { proto.RegisterFile("vince/vesting/v1/tx.proto", fileDescriptor_0e70a71da0ce69ab) }

var fileDescriptor_0e70a71da0ce69ab = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x8b, 0xd3, 0x4e,
	0x18, 0xc6, 0x9b, 0xb6, 0xbb, 0xb4, 0xd3, 0xdd, 0x2e, 0xcc, 0x7f, 0xff, 0x50, 0xea, 0x9a, 0xb4,
	0x45, 0x69, 0xbd, 0x4c, 0x68, 0x0b, 0x2a, 0x0a, 0xc2, 0xb6, 0x47, 0x29, 0x68, 0x11, 0x0f, 0x5e,
	0xca, 0x34, 0x99, 0x66, 0xc3, 0x36, 0x79, 0x43, 0x66, 0x52, 0xd7, 0x83, 0x20, 0x7e, 0x82, 0x05,
	0xbf, 0x80, 0x1f, 0x67, 0x8f, 0x05, 0x0f, 0x7a, 0xd2, 0xa5, 0xf5, 0xe0, 0xc7, 0x90, 0x64, 0x26,
	0xd9, 0xb2, 0xba, 0x16, 0xbd, 0x65, 0xde, 0xf7, 0xf7, 0x3e, 0x3c, 0xef, 0x33, 0x49, 0xd0, 0x2d,
	0x66, 0x9d, 0xb0, 0x39, 0xf8, 0xe6, 0x82, 0x71, 0xe1, 0xfa, 0x8e, 0xb9, 0xe8, 0x9a, 0xe2, 0x8c,
	0x04, 0x21, 0x08, 0xc0, 0x58, 0x35, 0x89, 0x6a, 0x92, 0x45, 0xb7, 0x7e, 0xe4, 0x00, 0x38, 0x73,
	0x66, 0xd2, 0xc0, 0x35, 0xa9, 0xef, 0x83, 0xa0, 0xc2, 0x05, 0x9f, 0xcb, 0x89, 0xfa, 0xa1, 0x03,
	0x0e, 0x24, 0x8f, 0x66, 0xfc, 0xa4, 0xaa, 0x86, 0x9a, 0x49, 0x4e, 0xd3, 0x68, 0x66, 0x0a, 0xd7,
	0x63, 0x5c, 0x50, 0x2f, 0x50, 0xc0, 0x1d, 0x0b, 0xb8, 0x07, 0x7c, 0xc3, 0xc4, 0x94, 0x09, 0xda,
	0x4d, 0xcf, 0x92, 0x6a, 0x5d, 0xe6, 0x91, 0x31, 0xe2, 0xce, 0x30, 0x64, 0x54, 0xb0, 0xe1, 0x9c,
	0xbe, 0x9e, 0x52, 0xeb, 0xf4, 0xa5, 0x44, 0x8e, 0x2d, 0x0b, 0x22, 0x5f, 0xe0, 0x26, 0xda, 0x9b,
	0x85, 0xe0, 0x4d, 0xa8, 0x6d, 0x87, 0x8c, 0xf3, 0x9a, 0xd6, 0xd0, 0x3a, 0xe5, 0x71, 0x25, 0xae,
	0x1d, 0xcb, 0x12, 0xbe, 0x8d, 0x90, 0x80, 0x0c, 0xc8, 0x27, 0x40, 0x59, 0x40, 0xda, 0x1e, 0x22,
	0xc4, 0x05, 0x0d, 0xc5, 0x24, 0x36, 0x59, 0x2b, 0x34, 0xb4, 0x4e, 0xa5, 0x57, 0x27, 0x72, 0x03,
	0x92, 0x6e, 0x40, 0x5e, 0xa4, 0x1b, 0x0c, 0x4a, 0x17, 0x5f, 0x8d, 0xdc, 0xf9, 0x37, 0x43, 0x1b,
	0x97, 0x93, 0xb9, 0xb8, 0x83, 0x9f, 0xa2, 0xea, 0x1c, 0xac, 0xd3, 0x28, 0x98, 0x04, 0x2c, 0x74,
	0xc1, 0xe6, 0xb5, 0x62, 0xa3, 0xd0, 0xa9, 0xf4, 0x74, 0x22, 0x37, 0xdd, 0x48, 0x34, 0xd9, 0x94,
	0x3c, 0x4b, 0xb0, 0x41, 0x31, 0x16, 0x1b, 0xef, 0xcb, 0x59, 0x59, 0xe3, 0x78, 0x84, 0x0e, 0x14,
	0x9e, 0xa9, 0xed, 0xfc, 0x85, 0x5a, 0x55, 0x75, 0x53, 0xb9, 0x43, 0xb4, 0xe3, 0xb1, 0xd0, 0x61,
	0xb5, 0xdd, 0x86, 0xd6, 0x29, 0x8d, 0xe5, 0xe1, 0x51, 0xf1, 0xc7, 0x47, 0x23, 0xd7, 0xba, 0x87,
	0xda, 0x5b, 0x12, 0x1e, 0x33, 0x1e, 0x80, 0xcf, 0x59, 0xeb, 0x9d, 0x86, 0x2a, 0x31, 0xab, 0x28,
	0x7c, 0x17, 0x55, 0x67, 0x91, 0x6f, 0xb3, 0xf0, 0x5a, 0xf6, 0xfb, 0xb2, 0x9a, 0xc6, 0xdb, 0x46,
	0x07, 0x54, 0x2a, 0x5d, 0xbb, 0x82, 0xaa, 0x2a, 0xa7, 0x60, 0x13, 0xed, 0xd9, 0x8c, 0x5f, 0x51,
	0x05, 0x79, 0x93, 0x71, 0x4d, 0x21, 0xad, 0xff, 0xd1, 0x7f, 0x1b, 0x0e, 0x52, 0x67, 0xbd, 0xcf,
	0x79, 0x54, 0x18, 0x71, 0x07, 0x2f, 0x35, 0x74, 0xf4, 0xc7, 0x97, 0xa5, 0x4f, 0x7e, 0x7d, 0xc1,
	0xc9, 0x96, 0xfd, 0xeb, 0x8f, 0xff, 0x61, 0x28, 0x0b, 0xed, 0xc9, 0xfb, 0x4f, 0xdf, 0x3f, 0xe4,
	0x1f, 0xe2, 0xfb, 0xe6, 0x6f, 0xbf, 0x3b, 0xd3, 0x4a, 0x44, 0x26, 0x96, 0x52, 0x99, 0xa4, 0xf7,
	0xae, 0xb2, 0xc1, 0x6f, 0x51, 0x29, 0x0b, 0xdc, 0xb8, 0xc9, 0x88, 0x02, 0xea, 0xed, 0x2d, 0x40,
	0xe6, 0xaa, 0x9d, 0xb8, 0x6a, 0x62, 0xe3, 0x26, 0x57, 0x6a, 0x60, 0xf0, 0xfc, 0x62, 0xa5, 0x6b,
	0xcb, 0x95, 0xae, 0x5d, 0xae, 0x74, 0xed, 0x7c, 0xad, 0xe7, 0x96, 0x6b, 0x3d, 0xf7, 0x65, 0xad,
	0xe7, 0x5e, 0x3d, 0x70, 0x5c, 0x71, 0x12, 0x4d, 0x89, 0x05, 0x5e, 0x2a, 0x32, 0x83, 0xc8, 0xb7,
	0x93, 0xbf, 0xc3, 0x95, 0x6c, 0xdf, 0x3c, 0xcb, 0xb4, 0xc5, 0x9b, 0x80, 0xf1, 0xe9, 0x6e, 0xf2,
	0x49, 0xf5, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x24, 0x59, 0xb6, 0x65, 0x89, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateClawbackVestingAccount creats a vesting account that is subject to
	// clawback and the configuration of vesting and lockup schedules.
	CreateClawbackVestingAccount(ctx context.Context, in *MsgCreateClawbackVestingAccount, opts ...grpc.CallOption) (*MsgCreateClawbackVestingAccountResponse, error)
	// Clawback removes the unvested tokens from a ClawbackVestingAccount.
	Clawback(ctx context.Context, in *MsgClawback, opts ...grpc.CallOption) (*MsgClawbackResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateClawbackVestingAccount(ctx context.Context, in *MsgCreateClawbackVestingAccount, opts ...grpc.CallOption) (*MsgCreateClawbackVestingAccountResponse, error) {
	out := new(MsgCreateClawbackVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/vince.vesting.v1.Msg/CreateClawbackVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Clawback(ctx context.Context, in *MsgClawback, opts ...grpc.CallOption) (*MsgClawbackResponse, error) {
	out := new(MsgClawbackResponse)
	err := c.cc.Invoke(ctx, "/vince.vesting.v1.Msg/Clawback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateClawbackVestingAccount creats a vesting account that is subject to
	// clawback and the configuration of vesting and lockup schedules.
	CreateClawbackVestingAccount(context.Context, *MsgCreateClawbackVestingAccount) (*MsgCreateClawbackVestingAccountResponse, error)
	// Clawback removes the unvested tokens from a ClawbackVestingAccount.
	Clawback(context.Context, *MsgClawback) (*MsgClawbackResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateClawbackVestingAccount(ctx context.Context, req *MsgCreateClawbackVestingAccount) (*MsgCreateClawbackVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClawbackVestingAccount not implemented")
}
func (*UnimplementedMsgServer) Clawback(ctx context.Context, req *MsgClawback) (*MsgClawbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clawback not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateClawbackVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateClawbackVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateClawbackVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vince.vesting.v1.Msg/CreateClawbackVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateClawbackVestingAccount(ctx, req.(*MsgCreateClawbackVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Clawback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClawback)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Clawback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vince.vesting.v1.Msg/Clawback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Clawback(ctx, req.(*MsgClawback))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vince.vesting.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClawbackVestingAccount",
			Handler:    _Msg_CreateClawbackVestingAccount_Handler,
		},
		{
			MethodName: "Clawback",
			Handler:    _Msg_Clawback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vince/vesting/v1/tx.proto",
}

func (m *MsgCreateClawbackVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateClawbackVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateClawbackVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Merge {
		i--
		if m.Merge {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LockupPeriods) > 0 {
		for iNdEx := len(m.LockupPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LockupPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTx(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateClawbackVestingAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateClawbackVestingAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateClawbackVestingAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgClawback) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClawback) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClawback) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DestAddress) > 0 {
		i -= len(m.DestAddress)
		copy(dAtA[i:], m.DestAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DestAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccountAddress) > 0 {
		i -= len(m.AccountAddress)
		copy(dAtA[i:], m.AccountAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AccountAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FunderAddress) > 0 {
		i -= len(m.FunderAddress)
		copy(dAtA[i:], m.FunderAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FunderAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClawbackResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClawbackResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClawbackResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateClawbackVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovTx(uint64(l))
	if len(m.LockupPeriods) > 0 {
		for _, e := range m.LockupPeriods {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.Merge {
		n += 2
	}
	return n
}

func (m *MsgCreateClawbackVestingAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgClawback) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FunderAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AccountAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DestAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClawbackResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateClawbackVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateClawbackVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateClawbackVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockupPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockupPeriods = append(m.LockupPeriods, types.Period{})
			if err := m.LockupPeriods[len(m.LockupPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, types.Period{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Merge", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Merge = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateClawbackVestingAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateClawbackVestingAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateClawbackVestingAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClawback) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClawback: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClawback: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FunderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClawbackResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClawbackResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClawbackResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)