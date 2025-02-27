// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/avstask/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// TaskContractInfo is the task info.
type TaskContractInfo struct {
	// contract address of avstask
	TaskContractAddress string `protobuf:"bytes,1,opt,name=task_contract_address,json=taskContractAddress,proto3" json:"task_contract_address,omitempty"`
	// name of task
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// meta_info of task
	MetaInfo string `protobuf:"bytes,3,opt,name=meta_info,json=metaInfo,proto3" json:"meta_info,omitempty"`
	// status of task
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// source_code of task
	SourceCode string `protobuf:"bytes,5,opt,name=source_code,json=sourceCode,proto3" json:"source_code,omitempty"`
}

func (m *TaskContractInfo) Reset()         { *m = TaskContractInfo{} }
func (m *TaskContractInfo) String() string { return proto.CompactTextString(m) }
func (*TaskContractInfo) ProtoMessage()    {}
func (*TaskContractInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_486b8c7a293e95a4, []int{0}
}
func (m *TaskContractInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskContractInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskContractInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskContractInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskContractInfo.Merge(m, src)
}
func (m *TaskContractInfo) XXX_Size() int {
	return m.Size()
}
func (m *TaskContractInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskContractInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TaskContractInfo proto.InternalMessageInfo

func (m *TaskContractInfo) GetTaskContractAddress() string {
	if m != nil {
		return m.TaskContractAddress
	}
	return ""
}

func (m *TaskContractInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskContractInfo) GetMetaInfo() string {
	if m != nil {
		return m.MetaInfo
	}
	return ""
}

func (m *TaskContractInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TaskContractInfo) GetSourceCode() string {
	if m != nil {
		return m.SourceCode
	}
	return ""
}

// RegisterAVSTaskReq is the request to register a new task for avs.
type RegisterAVSTaskReq struct {
	// from_address is the address of the avs (sdk.AccAddress).
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	// info is the task info.
	Task *TaskContractInfo `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (m *RegisterAVSTaskReq) Reset()         { *m = RegisterAVSTaskReq{} }
func (m *RegisterAVSTaskReq) String() string { return proto.CompactTextString(m) }
func (*RegisterAVSTaskReq) ProtoMessage()    {}
func (*RegisterAVSTaskReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_486b8c7a293e95a4, []int{1}
}
func (m *RegisterAVSTaskReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterAVSTaskReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterAVSTaskReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterAVSTaskReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAVSTaskReq.Merge(m, src)
}
func (m *RegisterAVSTaskReq) XXX_Size() int {
	return m.Size()
}
func (m *RegisterAVSTaskReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAVSTaskReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAVSTaskReq proto.InternalMessageInfo

// RegisterAVSTaskResponse is the response for register avs task
type RegisterAVSTaskResponse struct {
}

func (m *RegisterAVSTaskResponse) Reset()         { *m = RegisterAVSTaskResponse{} }
func (m *RegisterAVSTaskResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterAVSTaskResponse) ProtoMessage()    {}
func (*RegisterAVSTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_486b8c7a293e95a4, []int{2}
}
func (m *RegisterAVSTaskResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterAVSTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterAVSTaskResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterAVSTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAVSTaskResponse.Merge(m, src)
}
func (m *RegisterAVSTaskResponse) XXX_Size() int {
	return m.Size()
}
func (m *RegisterAVSTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAVSTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAVSTaskResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TaskContractInfo)(nil), "exocore.avstask.v1.TaskContractInfo")
	proto.RegisterType((*RegisterAVSTaskReq)(nil), "exocore.avstask.v1.RegisterAVSTaskReq")
	proto.RegisterType((*RegisterAVSTaskResponse)(nil), "exocore.avstask.v1.RegisterAVSTaskResponse")
}

func init() { proto.RegisterFile("exocore/avstask/v1/tx.proto", fileDescriptor_486b8c7a293e95a4) }

var fileDescriptor_486b8c7a293e95a4 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x6e, 0x13, 0x41,
	0x10, 0xc6, 0xef, 0x88, 0x13, 0x91, 0x35, 0x12, 0xb0, 0x04, 0x72, 0xb1, 0xa5, 0x33, 0xb2, 0x10,
	0x42, 0x46, 0xb9, 0x95, 0x4d, 0x13, 0x41, 0x95, 0x44, 0x20, 0x21, 0x04, 0xc5, 0x05, 0x51, 0xd0,
	0x9c, 0x36, 0x77, 0xeb, 0xe3, 0x64, 0xed, 0x8d, 0xd9, 0x59, 0x1b, 0xd3, 0x21, 0x2a, 0x44, 0xc5,
	0x23, 0xe4, 0x11, 0x5c, 0xf0, 0x0a, 0x48, 0x94, 0x11, 0x15, 0x25, 0xb2, 0x0b, 0xf3, 0x18, 0xe8,
	0x76, 0x37, 0xfc, 0xb1, 0x29, 0xd2, 0x9c, 0x76, 0xe6, 0xf7, 0xcd, 0xce, 0xb7, 0x73, 0x43, 0x9a,
	0x62, 0x02, 0x29, 0x28, 0xc1, 0xf8, 0x18, 0x35, 0xc7, 0x01, 0x1b, 0x77, 0x99, 0x9e, 0x44, 0x43,
	0x05, 0x1a, 0x28, 0x75, 0x30, 0x72, 0x30, 0x1a, 0x77, 0x1b, 0x57, 0xb9, 0x2c, 0x4a, 0x60, 0xe6,
	0x6b, 0x65, 0x8d, 0xed, 0x14, 0x50, 0x02, 0x32, 0x89, 0x79, 0x55, 0x2e, 0x31, 0x77, 0x60, 0xc7,
	0x82, 0xc4, 0x44, 0xcc, 0x06, 0x0e, 0x6d, 0xe5, 0x90, 0x83, 0xcd, 0x57, 0x27, 0x9b, 0x6d, 0x4f,
	0x7d, 0x72, 0xe5, 0x39, 0xc7, 0xc1, 0x21, 0x94, 0x5a, 0xf1, 0x54, 0x3f, 0x2e, 0xfb, 0x40, 0x7b,
	0xe4, 0x7a, 0xd5, 0x3c, 0x49, 0x5d, 0x32, 0xe1, 0x59, 0xa6, 0x04, 0x62, 0xe0, 0xdf, 0xf4, 0xef,
	0x6c, 0xc6, 0xd7, 0xf4, 0x5f, 0x05, 0xfb, 0x16, 0x51, 0x4a, 0x6a, 0x25, 0x97, 0x22, 0xb8, 0x60,
	0x24, 0xe6, 0x4c, 0x9b, 0x64, 0x53, 0x0a, 0xcd, 0x93, 0xa2, 0xec, 0x43, 0xb0, 0x66, 0xc0, 0xc5,
	0x2a, 0x61, 0x9a, 0xdc, 0x20, 0x1b, 0xa8, 0xb9, 0x1e, 0x61, 0x50, 0x33, 0xc4, 0x45, 0xb4, 0x45,
	0xea, 0x08, 0x23, 0x95, 0x8a, 0x24, 0x85, 0x4c, 0x04, 0xeb, 0x06, 0x12, 0x9b, 0x3a, 0x84, 0x4c,
	0xb4, 0xbf, 0xf8, 0x84, 0xc6, 0x22, 0x2f, 0x50, 0x0b, 0xb5, 0xff, 0xe2, 0xa8, 0x72, 0x1f, 0x8b,
	0xd7, 0xf4, 0x01, 0xb9, 0xd4, 0x57, 0x20, 0xff, 0xf5, 0x7a, 0x10, 0x7c, 0xfb, 0xbc, 0xbb, 0xe5,
	0xe6, 0xe0, 0xac, 0x1e, 0x69, 0x55, 0x94, 0x79, 0x5c, 0xaf, 0xd4, 0x67, 0xee, 0xf7, 0x48, 0xad,
	0x7a, 0x94, 0x71, 0x5f, 0xef, 0xdd, 0x8a, 0x56, 0x7f, 0x43, 0xb4, 0x3c, 0xa5, 0xd8, 0x54, 0xdc,
	0xdf, 0xfb, 0x70, 0xd2, 0xf2, 0x7e, 0x9e, 0xb4, 0xbc, 0xf7, 0x8b, 0x69, 0xa7, 0xfe, 0xe8, 0xcf,
	0x9d, 0x1f, 0x17, 0xd3, 0x4e, 0xd3, 0x76, 0xde, 0xc5, 0x6c, 0xc0, 0x96, 0xeb, 0xdb, 0x3b, 0x64,
	0x7b, 0xe5, 0x19, 0x38, 0x84, 0x12, 0x45, 0x4f, 0x93, 0xb5, 0xa7, 0x98, 0xd3, 0x3e, 0xb9, 0xbc,
	0xa4, 0xa0, 0xb7, 0xff, 0x67, 0x6d, 0x75, 0x1a, 0x8d, 0xbb, 0xe7, 0xd2, 0xd9, 0x76, 0x8d, 0xf5,
	0x77, 0x8b, 0x69, 0xc7, 0x3f, 0x78, 0xf2, 0x75, 0x16, 0xfa, 0xa7, 0xb3, 0xd0, 0xff, 0x31, 0x0b,
	0xfd, 0x4f, 0xf3, 0xd0, 0x3b, 0x9d, 0x87, 0xde, 0xf7, 0x79, 0xe8, 0xbd, 0xec, 0xe6, 0x85, 0x7e,
	0x35, 0x3a, 0x8e, 0x52, 0x90, 0xec, 0xa1, 0xbd, 0xf7, 0x99, 0xd0, 0x6f, 0x40, 0x0d, 0xd8, 0xd9,
	0x36, 0x4f, 0x7e, 0xef, 0xb3, 0x7e, 0x3b, 0x14, 0x78, 0xbc, 0x61, 0xf6, 0xeb, 0xde, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb5, 0x90, 0xcf, 0x54, 0xef, 0x02, 0x00, 0x00,
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
	// RegisterAVSTask registers a new task.
	RegisterAVSTask(ctx context.Context, in *RegisterAVSTaskReq, opts ...grpc.CallOption) (*RegisterAVSTaskResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterAVSTask(ctx context.Context, in *RegisterAVSTaskReq, opts ...grpc.CallOption) (*RegisterAVSTaskResponse, error) {
	out := new(RegisterAVSTaskResponse)
	err := c.cc.Invoke(ctx, "/exocore.avstask.v1.Msg/RegisterAVSTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RegisterAVSTask registers a new task.
	RegisterAVSTask(context.Context, *RegisterAVSTaskReq) (*RegisterAVSTaskResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterAVSTask(ctx context.Context, req *RegisterAVSTaskReq) (*RegisterAVSTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAVSTask not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterAVSTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAVSTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterAVSTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exocore.avstask.v1.Msg/RegisterAVSTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterAVSTask(ctx, req.(*RegisterAVSTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exocore.avstask.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAVSTask",
			Handler:    _Msg_RegisterAVSTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exocore/avstask/v1/tx.proto",
}

func (m *TaskContractInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskContractInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskContractInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SourceCode) > 0 {
		i -= len(m.SourceCode)
		copy(dAtA[i:], m.SourceCode)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceCode)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MetaInfo) > 0 {
		i -= len(m.MetaInfo)
		copy(dAtA[i:], m.MetaInfo)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MetaInfo)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TaskContractAddress) > 0 {
		i -= len(m.TaskContractAddress)
		copy(dAtA[i:], m.TaskContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TaskContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterAVSTaskReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterAVSTaskReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterAVSTaskReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Task != nil {
		{
			size, err := m.Task.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
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

func (m *RegisterAVSTaskResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterAVSTaskResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterAVSTaskResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *TaskContractInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TaskContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MetaInfo)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourceCode)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *RegisterAVSTaskReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Task != nil {
		l = m.Task.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *RegisterAVSTaskResponse) Size() (n int) {
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
func (m *TaskContractInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TaskContractInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskContractInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskContractAddress", wireType)
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
			m.TaskContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaInfo", wireType)
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
			m.MetaInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceCode", wireType)
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
			m.SourceCode = string(dAtA[iNdEx:postIndex])
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
func (m *RegisterAVSTaskReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RegisterAVSTaskReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterAVSTaskReq: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Task", wireType)
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
			if m.Task == nil {
				m.Task = &TaskContractInfo{}
			}
			if err := m.Task.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *RegisterAVSTaskResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RegisterAVSTaskResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterAVSTaskResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
