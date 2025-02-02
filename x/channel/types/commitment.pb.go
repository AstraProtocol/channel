// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: channel/commitment.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Commitment struct {
	Index          string      `protobuf:"bytes,1,opt,name=Index,proto3" json:"Index,omitempty"`
	From           string      `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`
	CoinToCreator  *types.Coin `protobuf:"bytes,3,opt,name=CoinToCreator,proto3" json:"CoinToCreator,omitempty"`
	ToTimelockAddr string      `protobuf:"bytes,4,opt,name=ToTimelockAddr,proto3" json:"ToTimelockAddr,omitempty"`
	ToHashlockAddr string      `protobuf:"bytes,5,opt,name=ToHashlockAddr,proto3" json:"ToHashlockAddr,omitempty"`
	CoinToHtlc     *types.Coin `protobuf:"bytes,6,opt,name=CoinToHtlc,proto3" json:"CoinToHtlc,omitempty"`
	Timelock       uint64      `protobuf:"varint,7,opt,name=Timelock,proto3" json:"Timelock,omitempty"`
	Hashcode       string      `protobuf:"bytes,8,opt,name=Hashcode,proto3" json:"Hashcode,omitempty"`
	ChannelID      string      `protobuf:"bytes,9,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
}

func (m *Commitment) Reset()         { *m = Commitment{} }
func (m *Commitment) String() string { return proto.CompactTextString(m) }
func (*Commitment) ProtoMessage()    {}
func (*Commitment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f017e3d8b1468c, []int{0}
}
func (m *Commitment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Commitment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Commitment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Commitment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commitment.Merge(m, src)
}
func (m *Commitment) XXX_Size() int {
	return m.Size()
}
func (m *Commitment) XXX_DiscardUnknown() {
	xxx_messageInfo_Commitment.DiscardUnknown(m)
}

var xxx_messageInfo_Commitment proto.InternalMessageInfo

func (m *Commitment) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Commitment) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Commitment) GetCoinToCreator() *types.Coin {
	if m != nil {
		return m.CoinToCreator
	}
	return nil
}

func (m *Commitment) GetToTimelockAddr() string {
	if m != nil {
		return m.ToTimelockAddr
	}
	return ""
}

func (m *Commitment) GetToHashlockAddr() string {
	if m != nil {
		return m.ToHashlockAddr
	}
	return ""
}

func (m *Commitment) GetCoinToHtlc() *types.Coin {
	if m != nil {
		return m.CoinToHtlc
	}
	return nil
}

func (m *Commitment) GetTimelock() uint64 {
	if m != nil {
		return m.Timelock
	}
	return 0
}

func (m *Commitment) GetHashcode() string {
	if m != nil {
		return m.Hashcode
	}
	return ""
}

func (m *Commitment) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func init() {
	proto.RegisterType((*Commitment)(nil), "channel.channel.Commitment")
}

func init() { proto.RegisterFile("channel/commitment.proto", fileDescriptor_e0f017e3d8b1468c) }

var fileDescriptor_e0f017e3d8b1468c = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xeb, 0xfe, 0x6d, 0xff, 0xd6, 0x08, 0x90, 0x2c, 0x06, 0x53, 0x21, 0xab, 0x62, 0x40,
	0x9d, 0x1c, 0x0a, 0x13, 0x13, 0x2a, 0x41, 0xa8, 0x65, 0x42, 0x55, 0x27, 0xb6, 0xc4, 0xb1, 0x68,
	0x44, 0x9c, 0x5b, 0xd9, 0x06, 0x95, 0xb7, 0xe0, 0xb1, 0x18, 0x3b, 0x76, 0x44, 0xed, 0x8b, 0x20,
	0x3b, 0x4d, 0x28, 0x2c, 0x4c, 0xb9, 0xf7, 0xdc, 0x93, 0xfb, 0xd9, 0x3e, 0x98, 0x8a, 0x59, 0x94,
	0xe7, 0x32, 0x0b, 0x04, 0x28, 0x95, 0x5a, 0x25, 0x73, 0xcb, 0xe7, 0x1a, 0x2c, 0x90, 0xc3, 0xed,
	0x84, 0x6f, 0xbf, 0x5d, 0x26, 0xc0, 0x28, 0x30, 0x41, 0x1c, 0x19, 0x19, 0xbc, 0x0e, 0x62, 0x69,
	0xa3, 0x41, 0x20, 0x20, 0xcd, 0x8b, 0x1f, 0x4e, 0x57, 0x75, 0x8c, 0xc3, 0x6a, 0x0b, 0x39, 0xc2,
	0xcd, 0x71, 0x9e, 0xc8, 0x05, 0x45, 0x3d, 0xd4, 0xef, 0x4c, 0x8a, 0x86, 0x10, 0xdc, 0xb8, 0xd3,
	0xa0, 0x68, 0xdd, 0x8b, 0xbe, 0x26, 0xd7, 0x78, 0x3f, 0x84, 0x34, 0x9f, 0x42, 0xa8, 0x65, 0x64,
	0x41, 0xd3, 0x7f, 0x3d, 0xd4, 0xdf, 0xbb, 0x38, 0xe6, 0x05, 0x90, 0x3b, 0x20, 0xdf, 0x02, 0xb9,
	0x73, 0x4e, 0x7e, 0xfa, 0xc9, 0x19, 0x3e, 0x98, 0xc2, 0x34, 0x55, 0x32, 0x03, 0xf1, 0x3c, 0x4c,
	0x12, 0x4d, 0x1b, 0x7e, 0xfd, 0x2f, 0xb5, 0xf0, 0x8d, 0x22, 0x33, 0xab, 0x7c, 0xcd, 0xd2, 0xb7,
	0xab, 0x92, 0x2b, 0x77, 0x11, 0x07, 0x18, 0xd9, 0x4c, 0xd0, 0xd6, 0x5f, 0xa7, 0xd9, 0x31, 0x93,
	0x2e, 0x6e, 0x97, 0x48, 0xfa, 0xbf, 0x87, 0xfa, 0x8d, 0x49, 0xd5, 0xbb, 0x99, 0xc3, 0x08, 0x48,
	0x24, 0x6d, 0x7b, 0x70, 0xd5, 0x93, 0x13, 0xdc, 0x09, 0x8b, 0x77, 0x1e, 0xdf, 0xd2, 0x8e, 0x1f,
	0x7e, 0x0b, 0x37, 0xf7, 0x1f, 0x6b, 0x86, 0x96, 0x6b, 0x86, 0x3e, 0xd7, 0x0c, 0xbd, 0x6f, 0x58,
	0x6d, 0xb9, 0x61, 0xb5, 0xd5, 0x86, 0xd5, 0x1e, 0xcf, 0x9f, 0x52, 0x3b, 0x7b, 0x89, 0xb9, 0x00,
	0x15, 0x0c, 0x8d, 0xd5, 0xd1, 0x83, 0xcb, 0x42, 0x40, 0x16, 0x94, 0xc1, 0x2e, 0xaa, 0xca, 0xbe,
	0xcd, 0xa5, 0x89, 0x5b, 0x3e, 0xad, 0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x9b, 0x82,
	0x30, 0xfa, 0x01, 0x00, 0x00,
}

func (m *Commitment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Commitment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Commitment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintCommitment(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Hashcode) > 0 {
		i -= len(m.Hashcode)
		copy(dAtA[i:], m.Hashcode)
		i = encodeVarintCommitment(dAtA, i, uint64(len(m.Hashcode)))
		i--
		dAtA[i] = 0x42
	}
	if m.Timelock != 0 {
		i = encodeVarintCommitment(dAtA, i, uint64(m.Timelock))
		i--
		dAtA[i] = 0x38
	}
	if m.CoinToHtlc != nil {
		{
			size, err := m.CoinToHtlc.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommitment(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.ToHashlockAddr) > 0 {
		i -= len(m.ToHashlockAddr)
		copy(dAtA[i:], m.ToHashlockAddr)
		i = encodeVarintCommitment(dAtA, i, uint64(len(m.ToHashlockAddr)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ToTimelockAddr) > 0 {
		i -= len(m.ToTimelockAddr)
		copy(dAtA[i:], m.ToTimelockAddr)
		i = encodeVarintCommitment(dAtA, i, uint64(len(m.ToTimelockAddr)))
		i--
		dAtA[i] = 0x22
	}
	if m.CoinToCreator != nil {
		{
			size, err := m.CoinToCreator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommitment(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintCommitment(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintCommitment(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommitment(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommitment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Commitment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovCommitment(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovCommitment(uint64(l))
	}
	if m.CoinToCreator != nil {
		l = m.CoinToCreator.Size()
		n += 1 + l + sovCommitment(uint64(l))
	}
	l = len(m.ToTimelockAddr)
	if l > 0 {
		n += 1 + l + sovCommitment(uint64(l))
	}
	l = len(m.ToHashlockAddr)
	if l > 0 {
		n += 1 + l + sovCommitment(uint64(l))
	}
	if m.CoinToHtlc != nil {
		l = m.CoinToHtlc.Size()
		n += 1 + l + sovCommitment(uint64(l))
	}
	if m.Timelock != 0 {
		n += 1 + sovCommitment(uint64(m.Timelock))
	}
	l = len(m.Hashcode)
	if l > 0 {
		n += 1 + l + sovCommitment(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovCommitment(uint64(l))
	}
	return n
}

func sovCommitment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommitment(x uint64) (n int) {
	return sovCommitment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Commitment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommitment
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
			return fmt.Errorf("proto: Commitment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Commitment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitment
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
				return ErrInvalidLengthCommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitment
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
				return ErrInvalidLengthCommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinToCreator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitment
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
				return ErrInvalidLengthCommitment
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CoinToCreator == nil {
				m.CoinToCreator = &types.Coin{}
			}
			if err := m.CoinToCreator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToTimelockAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitment
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
				return ErrInvalidLengthCommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToTimelockAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToHashlockAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitment
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
				return ErrInvalidLengthCommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToHashlockAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinToHtlc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitment
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
				return ErrInvalidLengthCommitment
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CoinToHtlc == nil {
				m.CoinToHtlc = &types.Coin{}
			}
			if err := m.CoinToHtlc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timelock", wireType)
			}
			m.Timelock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timelock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashcode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitment
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
				return ErrInvalidLengthCommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashcode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitment
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
				return ErrInvalidLengthCommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommitment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommitment
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
func skipCommitment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommitment
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
					return 0, ErrIntOverflowCommitment
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
					return 0, ErrIntOverflowCommitment
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
				return 0, ErrInvalidLengthCommitment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommitment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommitment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommitment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommitment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommitment = fmt.Errorf("proto: unexpected end of group")
)
