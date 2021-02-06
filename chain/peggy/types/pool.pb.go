// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/peggy/v1beta1/pool.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// OutgoingTx is a withdrawal on the bridged contract
// TODO: can this type be replaced by outgoing transfer tx
type OutgoingTx struct {
	Sender    string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	DestAddr  string     `protobuf:"bytes,2,opt,name=dest_addr,json=destAddr,proto3" json:"dest_addr,omitempty"`
	Amount    types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	BridgeFee types.Coin `protobuf:"bytes,4,opt,name=bridge_fee,json=bridgeFee,proto3" json:"bridge_fee"`
}

func (m *OutgoingTx) Reset()         { *m = OutgoingTx{} }
func (m *OutgoingTx) String() string { return proto.CompactTextString(m) }
func (*OutgoingTx) ProtoMessage()    {}
func (*OutgoingTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_0baba9e874638a68, []int{0}
}
func (m *OutgoingTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTx.Merge(m, src)
}
func (m *OutgoingTx) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingTx) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTx.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTx proto.InternalMessageInfo

func (m *OutgoingTx) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *OutgoingTx) GetDestAddr() string {
	if m != nil {
		return m.DestAddr
	}
	return ""
}

func (m *OutgoingTx) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *OutgoingTx) GetBridgeFee() types.Coin {
	if m != nil {
		return m.BridgeFee
	}
	return types.Coin{}
}

// IDSet represents a set of IDs
type IDSet struct {
	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (m *IDSet) Reset()         { *m = IDSet{} }
func (m *IDSet) String() string { return proto.CompactTextString(m) }
func (*IDSet) ProtoMessage()    {}
func (*IDSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0baba9e874638a68, []int{1}
}
func (m *IDSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IDSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IDSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IDSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDSet.Merge(m, src)
}
func (m *IDSet) XXX_Size() int {
	return m.Size()
}
func (m *IDSet) XXX_DiscardUnknown() {
	xxx_messageInfo_IDSet.DiscardUnknown(m)
}

var xxx_messageInfo_IDSet proto.InternalMessageInfo

func (m *IDSet) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*OutgoingTx)(nil), "injective.peggy.v1beta1.OutgoingTx")
	proto.RegisterType((*IDSet)(nil), "injective.peggy.v1beta1.IDSet")
}

func init() {
	proto.RegisterFile("injective/peggy/v1beta1/pool.proto", fileDescriptor_0baba9e874638a68)
}

var fileDescriptor_0baba9e874638a68 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4a, 0x03, 0x41,
	0x18, 0xc4, 0x6f, 0x4d, 0x0c, 0x66, 0x6d, 0xe4, 0x10, 0xbd, 0x44, 0x38, 0x43, 0xaa, 0x34, 0xee,
	0x12, 0x2d, 0xec, 0x04, 0xa3, 0x08, 0x41, 0x41, 0x88, 0x56, 0x36, 0xe1, 0xee, 0xf6, 0x73, 0xb3,
	0x92, 0xdb, 0xef, 0xb8, 0xdd, 0x0b, 0xe6, 0x2d, 0x7c, 0x1e, 0x9f, 0x20, 0x65, 0x4a, 0x2b, 0x91,
	0xe4, 0x45, 0xe4, 0xfe, 0x24, 0x58, 0xda, 0xcd, 0x37, 0x3b, 0x3f, 0x18, 0x66, 0x69, 0x57, 0xe9,
	0x37, 0x88, 0xac, 0x9a, 0x01, 0x4f, 0x40, 0xca, 0x39, 0x9f, 0xf5, 0x43, 0xb0, 0x41, 0x9f, 0x27,
	0x88, 0x53, 0x96, 0xa4, 0x68, 0xd1, 0x3d, 0xde, 0x66, 0x58, 0x91, 0x61, 0x55, 0xa6, 0xed, 0x47,
	0x68, 0x62, 0x34, 0x3c, 0x0c, 0x0c, 0x6c, 0xc1, 0x08, 0x95, 0x2e, 0xc1, 0xf6, 0xa1, 0x44, 0x89,
	0x85, 0xe4, 0xb9, 0x2a, 0xdd, 0xee, 0x27, 0xa1, 0xf4, 0x31, 0xb3, 0x12, 0x95, 0x96, 0xcf, 0xef,
	0xee, 0x11, 0x6d, 0x18, 0xd0, 0x02, 0x52, 0x8f, 0x74, 0x48, 0xaf, 0x39, 0xaa, 0x2e, 0xf7, 0x84,
	0x36, 0x05, 0x18, 0x3b, 0x0e, 0x84, 0x48, 0xbd, 0x9d, 0xe2, 0x69, 0x2f, 0x37, 0xae, 0x85, 0x48,
	0xdd, 0x4b, 0xda, 0x08, 0x62, 0xcc, 0xb4, 0xf5, 0x6a, 0x1d, 0xd2, 0xdb, 0x3f, 0x6f, 0xb1, 0xb2,
	0x0a, 0xcb, 0xab, 0x6c, 0xfa, 0xb1, 0x1b, 0x54, 0x7a, 0x50, 0x5f, 0x7c, 0x9f, 0x3a, 0xa3, 0x2a,
	0xee, 0x5e, 0x51, 0x1a, 0xa6, 0x4a, 0x48, 0x18, 0xbf, 0x02, 0x78, 0xf5, 0xff, 0xc1, 0xcd, 0x12,
	0xb9, 0x03, 0xe8, 0xb6, 0xe8, 0xee, 0xf0, 0xf6, 0x09, 0xac, 0x7b, 0x40, 0x6b, 0x4a, 0x18, 0x8f,
	0x74, 0x6a, 0xbd, 0xfa, 0x28, 0x97, 0x03, 0x58, 0xac, 0x7c, 0xb2, 0x5c, 0xf9, 0xe4, 0x67, 0xe5,
	0x93, 0x8f, 0xb5, 0xef, 0x2c, 0xd7, 0xbe, 0xf3, 0xb5, 0xf6, 0x9d, 0x97, 0x7b, 0xa9, 0xec, 0x24,
	0x0b, 0x59, 0x84, 0x31, 0x1f, 0x6e, 0xb6, 0x7c, 0x08, 0x42, 0xc3, 0xb7, 0xcb, 0x9e, 0x45, 0x98,
	0xc2, 0xdf, 0x73, 0x12, 0x28, 0xcd, 0x63, 0x14, 0xd9, 0x14, 0x4c, 0xf5, 0x35, 0x76, 0x9e, 0x80,
	0x09, 0x1b, 0xc5, 0x8a, 0x17, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x2b, 0xe7, 0xbd, 0xba,
	0x01, 0x00, 0x00,
}

func (m *OutgoingTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BridgeFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.DestAddr) > 0 {
		i -= len(m.DestAddr)
		copy(dAtA[i:], m.DestAddr)
		i = encodeVarintPool(dAtA, i, uint64(len(m.DestAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IDSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IDSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IDSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ids) > 0 {
		dAtA4 := make([]byte, len(m.Ids)*10)
		var j3 int
		for _, num := range m.Ids {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintPool(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutgoingTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.DestAddr)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.BridgeFee.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func (m *IDSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ids) > 0 {
		l = 0
		for _, e := range m.Ids {
			l += sovPool(uint64(e))
		}
		n += 1 + sovPool(uint64(l)) + l
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutgoingTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: OutgoingTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BridgeFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPool
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
func (m *IDSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: IDSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IDSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPool
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ids = append(m.Ids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPool
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthPool
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthPool
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Ids) == 0 {
					m.Ids = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPool
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ids = append(m.Ids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)