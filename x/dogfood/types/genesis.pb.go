// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/dogfood/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the dogfood module's genesis state.
type GenesisState struct {
	// params refers to the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// initial_val_set is the initial validator set.
	InitialValSet []GenesisValidator `protobuf:"bytes,2,rep,name=initial_val_set,json=initialValSet,proto3" json:"initial_val_set"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a9d908a27866b1b, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetInitialValSet() []GenesisValidator {
	if m != nil {
		return m.InitialValSet
	}
	return nil
}

// GenesisValidator defines a genesis validator. It is a helper struct
// used for serializing the genesis state.
type GenesisValidator struct {
	// public_key is the consensus public key of the validator. It should
	// be exactly 32 bytes, but this is not enforced in protobuf.
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// power is the voting power of the validator.
	Power int64 `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *GenesisValidator) Reset()         { *m = GenesisValidator{} }
func (m *GenesisValidator) String() string { return proto.CompactTextString(m) }
func (*GenesisValidator) ProtoMessage()    {}
func (*GenesisValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a9d908a27866b1b, []int{1}
}
func (m *GenesisValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisValidator.Merge(m, src)
}
func (m *GenesisValidator) XXX_Size() int {
	return m.Size()
}
func (m *GenesisValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisValidator.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisValidator proto.InternalMessageInfo

func (m *GenesisValidator) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *GenesisValidator) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "exocore.dogfood.v1.GenesisState")
	proto.RegisterType((*GenesisValidator)(nil), "exocore.dogfood.v1.GenesisValidator")
}

func init() { proto.RegisterFile("exocore/dogfood/v1/genesis.proto", fileDescriptor_1a9d908a27866b1b) }

var fileDescriptor_1a9d908a27866b1b = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x6a, 0x32, 0x41,
	0x18, 0x85, 0x77, 0xf4, 0xfb, 0x04, 0xc7, 0x84, 0x84, 0xc1, 0x42, 0x84, 0x8c, 0x8b, 0xa4, 0xb0,
	0x9a, 0x41, 0xd3, 0xa4, 0x16, 0x82, 0x85, 0x10, 0xc2, 0x0a, 0x16, 0x69, 0x64, 0xd4, 0x37, 0x9b,
	0xc1, 0xd5, 0x77, 0x99, 0x1d, 0xff, 0xee, 0x22, 0x17, 0x90, 0x0b, 0xb2, 0xb4, 0x4c, 0x15, 0x82,
	0xde, 0x48, 0xc8, 0xce, 0x24, 0x45, 0x62, 0x37, 0x3f, 0x0f, 0xcf, 0x39, 0x1c, 0x1a, 0xc2, 0x06,
	0x27, 0x68, 0x40, 0x4e, 0x31, 0x7e, 0x42, 0x9c, 0xca, 0x55, 0x5b, 0xc6, 0xb0, 0x80, 0x4c, 0x67,
	0x22, 0x35, 0x68, 0x91, 0x31, 0x4f, 0x08, 0x4f, 0x88, 0x55, 0xbb, 0x5e, 0x8d, 0x31, 0xc6, 0xfc,
	0x5b, 0x7e, 0x9d, 0x1c, 0x59, 0x6f, 0x9c, 0x70, 0xa5, 0xca, 0xa8, 0xb9, 0x57, 0x35, 0x5f, 0x09,
	0x3d, 0xeb, 0x39, 0xf9, 0xc0, 0x2a, 0x0b, 0xec, 0x96, 0x96, 0x1c, 0x50, 0x23, 0x21, 0x69, 0x55,
	0x3a, 0x75, 0xf1, 0x37, 0x4c, 0x3c, 0xe4, 0x44, 0xf7, 0xdf, 0xee, 0xbd, 0x11, 0x44, 0x9e, 0x67,
	0x11, 0xbd, 0xd0, 0x0b, 0x6d, 0xb5, 0x4a, 0x46, 0x2b, 0x95, 0x8c, 0x32, 0xb0, 0xb5, 0x42, 0x58,
	0x6c, 0x55, 0x3a, 0xd7, 0xa7, 0x14, 0x3e, 0x74, 0xa8, 0x12, 0x3d, 0x55, 0x16, 0x8d, 0x97, 0x9d,
	0x7b, 0xc5, 0x50, 0x25, 0x03, 0xb0, 0xcd, 0x1e, 0xbd, 0xfc, 0x0d, 0xb2, 0x2b, 0x4a, 0xd3, 0xe5,
	0x38, 0xd1, 0x93, 0xd1, 0x0c, 0xb6, 0x79, 0xcb, 0x72, 0x54, 0x76, 0x2f, 0x7d, 0xd8, 0xb2, 0x2a,
	0xfd, 0x9f, 0xe2, 0x1a, 0x4c, 0xad, 0x10, 0x92, 0x56, 0x31, 0x72, 0x97, 0x6e, 0x7f, 0x77, 0xe0,
	0x64, 0x7f, 0xe0, 0xe4, 0xe3, 0xc0, 0xc9, 0xcb, 0x91, 0x07, 0xfb, 0x23, 0x0f, 0xde, 0x8e, 0x3c,
	0x78, 0x6c, 0xc7, 0xda, 0x3e, 0x2f, 0xc7, 0x62, 0x82, 0x73, 0x79, 0xe7, 0x7a, 0xde, 0x83, 0x5d,
	0xa3, 0x99, 0xc9, 0xef, 0xf1, 0x36, 0x3f, 0xf3, 0xd9, 0x6d, 0x0a, 0xd9, 0xb8, 0x94, 0x6f, 0x77,
	0xf3, 0x19, 0x00, 0x00, 0xff, 0xff, 0x94, 0xc2, 0x5a, 0x03, 0xaa, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InitialValSet) > 0 {
		for iNdEx := len(m.InitialValSet) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialValSet[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.InitialValSet) > 0 {
		for _, e := range m.InitialValSet {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovGenesis(uint64(m.Power))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialValSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialValSet = append(m.InitialValSet, GenesisValidator{})
			if err := m.InitialValSet[len(m.InitialValSet)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
