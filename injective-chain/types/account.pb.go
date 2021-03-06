// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/types/v1beta1/account.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// EthAccount implements the authtypes.AccountI interface and embeds an
// authtypes.BaseAccount type. It is compatible with the auth AccountKeeper.
type EthAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
	CodeHash           []byte `protobuf:"bytes,2,opt,name=code_hash,json=codeHash,proto3" json:"code_hash,omitempty" yaml:"code_hash"`
}

func (m *EthAccount) Reset()      { *m = EthAccount{} }
func (*EthAccount) ProtoMessage() {}
func (*EthAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e25f61138fdc8ede, []int{0}
}
func (m *EthAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthAccount.Merge(m, src)
}
func (m *EthAccount) XXX_Size() int {
	return m.Size()
}
func (m *EthAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_EthAccount.DiscardUnknown(m)
}

var xxx_messageInfo_EthAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EthAccount)(nil), "injective.types.v1beta1.EthAccount")
}

func init() {
	proto.RegisterFile("injective/types/v1beta1/account.proto", fileDescriptor_e25f61138fdc8ede)
}

var fileDescriptor_e25f61138fdc8ede = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0x31, 0x4f, 0x72, 0x31,
	0x14, 0x7d, 0xfd, 0x86, 0x2f, 0xfa, 0x60, 0x30, 0x48, 0x22, 0x62, 0xd2, 0x47, 0x5e, 0x62, 0xc2,
	0xf2, 0xda, 0x80, 0x1b, 0x9b, 0x2f, 0x31, 0x81, 0xc4, 0x89, 0xd1, 0x05, 0x6f, 0x4b, 0xa1, 0x4f,
	0x81, 0x12, 0xda, 0x47, 0xe4, 0x1f, 0x38, 0x3a, 0x3a, 0xf2, 0x23, 0xfc, 0x11, 0x8e, 0xc4, 0xc9,
	0x89, 0x18, 0x88, 0x89, 0x33, 0xbf, 0xc0, 0xf0, 0x5a, 0x81, 0xa9, 0xf7, 0xde, 0x73, 0xce, 0x3d,
	0xed, 0xa9, 0x7f, 0x99, 0x8c, 0x1e, 0x04, 0x37, 0xc9, 0x54, 0x50, 0x33, 0x1b, 0x0b, 0x4d, 0xa7,
	0x35, 0x26, 0x0c, 0xd4, 0x28, 0x70, 0xae, 0xd2, 0x91, 0x21, 0xe3, 0x89, 0x32, 0xaa, 0x70, 0xb6,
	0xa3, 0x91, 0x8c, 0x46, 0x1c, 0xad, 0x8c, 0xb9, 0xd2, 0x43, 0xa5, 0x29, 0xa4, 0x46, 0xee, 0xb5,
	0xa9, 0x91, 0x56, 0x58, 0x3e, 0xb7, 0x78, 0x27, 0xeb, 0xa8, 0x6d, 0x1c, 0x54, 0xec, 0xab, 0xbe,
	0xb2, 0xf3, 0x6d, 0x65, 0xa7, 0xe1, 0x37, 0xf2, 0xfd, 0x1b, 0x23, 0xaf, 0xad, 0x7d, 0xe1, 0xde,
	0xcf, 0x33, 0xd0, 0xa2, 0xe3, 0xae, 0x53, 0x42, 0x15, 0x54, 0xcd, 0xd5, 0x2b, 0xc4, 0x6d, 0xca,
	0x9c, 0x9c, 0x2d, 0x89, 0x41, 0x0b, 0xa7, 0x8b, 0x2f, 0x16, 0xcb, 0x00, 0x6d, 0x96, 0xc1, 0xe9,
	0x0c, 0x86, 0x83, 0x46, 0x78, 0xb8, 0x23, 0x6c, 0xe7, 0xd8, 0x9e, 0x59, 0xa8, 0xf9, 0xc7, 0x5c,
	0x75, 0x45, 0x47, 0x82, 0x96, 0xa5, 0x7f, 0x15, 0x54, 0xcd, 0xc7, 0xc5, 0xcd, 0x32, 0x38, 0xb1,
	0xc2, 0x1d, 0x14, 0xb6, 0x8f, 0xb6, 0x75, 0x13, 0xb4, 0x6c, 0xc4, 0xcf, 0xf3, 0xc0, 0x7b, 0x9d,
	0x07, 0xde, 0xcf, 0x3c, 0xf0, 0x3e, 0xde, 0xa2, 0x7a, 0x3f, 0x31, 0x32, 0x65, 0x84, 0xab, 0xa1,
	0x7b, 0xa2, 0x3b, 0x22, 0xdd, 0x7d, 0xa4, 0x4f, 0x36, 0x1c, 0x9b, 0x9b, 0x73, 0x6d, 0xc5, 0xec,
	0x7d, 0x85, 0xd1, 0x62, 0x85, 0xd1, 0xd7, 0x0a, 0xa3, 0x97, 0x35, 0xf6, 0x16, 0x6b, 0xec, 0x7d,
	0xae, 0xb1, 0x77, 0xd7, 0x3c, 0xd8, 0xd6, 0xfa, 0x8b, 0xfd, 0x16, 0x98, 0xa6, 0xbb, 0x4f, 0x88,
	0xd4, 0x04, 0xf8, 0x40, 0x44, 0x9a, 0x43, 0xaf, 0xa7, 0x06, 0xdd, 0x03, 0x84, 0x4b, 0x48, 0x46,
	0xd6, 0x8c, 0xfd, 0xcf, 0x22, 0xbd, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x22, 0x68, 0xbe, 0x80,
	0xe5, 0x01, 0x00, 0x00,
}

func (m *EthAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CodeHash) > 0 {
		i -= len(m.CodeHash)
		copy(dAtA[i:], m.CodeHash)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.CodeHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAccount(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.CodeHash)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
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
			return fmt.Errorf("proto: EthAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeHash = append(m.CodeHash[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeHash == nil {
				m.CodeHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
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
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)
