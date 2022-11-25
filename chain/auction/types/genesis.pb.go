// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kaiju/auction/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the auction module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to auction.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// current auction round
	AuctionRound uint64 `protobuf:"varint,2,opt,name=auction_round,json=auctionRound,proto3" json:"auction_round,omitempty"`
	// current highest bid
	HighestBid *Bid `protobuf:"bytes,3,opt,name=highest_bid,json=highestBid,proto3" json:"highest_bid,omitempty"`
	// auction ending timestamp
	AuctionEndingTimestamp int64 `protobuf:"varint,4,opt,name=auction_ending_timestamp,json=auctionEndingTimestamp,proto3" json:"auction_ending_timestamp,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f095f457353f49, []int{0}
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

func (m *GenesisState) GetAuctionRound() uint64 {
	if m != nil {
		return m.AuctionRound
	}
	return 0
}

func (m *GenesisState) GetHighestBid() *Bid {
	if m != nil {
		return m.HighestBid
	}
	return nil
}

func (m *GenesisState) GetAuctionEndingTimestamp() int64 {
	if m != nil {
		return m.AuctionEndingTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "kaiju.auction.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("kaiju/auction/v1beta1/genesis.proto", fileDescriptor_56f095f457353f49)
}

var fileDescriptor_56f095f457353f49 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x5a, 0x75, 0x70, 0xcb, 0x12, 0x21, 0x14, 0x3a, 0x98, 0x02, 0x03, 0x5d, 0x88,
	0x55, 0x58, 0xd8, 0x2a, 0x45, 0x42, 0x08, 0x09, 0x24, 0x14, 0x98, 0x58, 0x2a, 0x27, 0xb1, 0x9c,
	0x43, 0xc4, 0x8e, 0x62, 0xa7, 0x12, 0x6f, 0xc1, 0x63, 0x75, 0xec, 0xc8, 0x84, 0x50, 0xfb, 0x00,
	0xbc, 0x02, 0xaa, 0xeb, 0x44, 0x2c, 0x74, 0x3b, 0xff, 0xff, 0x77, 0xf7, 0x9f, 0x7c, 0xf8, 0x1c,
	0xe4, 0x2b, 0x4f, 0x0d, 0xcc, 0x39, 0x65, 0x75, 0x6a, 0x40, 0x49, 0x3a, 0x9f, 0x24, 0xdc, 0xb0,
	0x09, 0x15, 0x5c, 0x72, 0x0d, 0x3a, 0x2c, 0x2b, 0x65, 0x94, 0x7f, 0xd4, 0x82, 0xa1, 0x03, 0x43,
	0x07, 0x0e, 0x77, 0xcc, 0x68, 0x50, 0x3b, 0x63, 0x78, 0x20, 0x94, 0x50, 0xb6, 0xa4, 0x9b, 0x6a,
	0xab, 0x9e, 0xfe, 0x20, 0x3c, 0xb8, 0xdd, 0x66, 0x3d, 0x19, 0x66, 0xb8, 0x3f, 0xc5, 0xbd, 0x92,
	0x55, 0xac, 0xd0, 0x01, 0x1a, 0xa1, 0x71, 0xff, 0xf2, 0x24, 0xfc, 0x37, 0x3b, 0x7c, 0xb4, 0x60,
	0xd4, 0x5d, 0x7c, 0x1d, 0x7b, 0xb1, 0x6b, 0xf3, 0xcf, 0xf0, 0xbe, 0xe3, 0x66, 0x95, 0xaa, 0x65,
	0x16, 0xec, 0x8d, 0xd0, 0xb8, 0x1b, 0x0f, 0x9c, 0x18, 0x6f, 0x34, 0x7f, 0x8a, 0xfb, 0x39, 0x88,
	0x9c, 0x6b, 0x33, 0x4b, 0x20, 0x0b, 0x3a, 0x36, 0x8a, 0xec, 0x88, 0x8a, 0x20, 0x8b, 0xb1, 0x6b,
	0x89, 0x20, 0xf3, 0xaf, 0x71, 0xd0, 0xa4, 0x70, 0x99, 0x81, 0x14, 0x33, 0x03, 0x05, 0xd7, 0x86,
	0x15, 0x65, 0xd0, 0x1d, 0xa1, 0x71, 0x27, 0x3e, 0x74, 0xfe, 0x8d, 0xb5, 0x9f, 0x1b, 0x37, 0x12,
	0x8b, 0x15, 0x41, 0xcb, 0x15, 0x41, 0xdf, 0x2b, 0x82, 0x3e, 0xd6, 0xc4, 0x5b, 0xae, 0x89, 0xf7,
	0xb9, 0x26, 0xde, 0xcb, 0x83, 0x00, 0x93, 0xd7, 0x49, 0x98, 0xaa, 0x82, 0xde, 0x35, 0x9b, 0xdc,
	0xb3, 0x44, 0xd3, 0x76, 0xaf, 0x8b, 0x54, 0x55, 0xfc, 0xef, 0x33, 0x67, 0x20, 0x69, 0xa1, 0xb2,
	0xfa, 0x8d, 0xeb, 0xf6, 0x00, 0xe6, 0xbd, 0xe4, 0x3a, 0xe9, 0xd9, 0x1f, 0xbe, 0xfa, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x2f, 0xc4, 0xb2, 0x63, 0xe6, 0x01, 0x00, 0x00,
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
	if m.AuctionEndingTimestamp != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AuctionEndingTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if m.HighestBid != nil {
		{
			size, err := m.HighestBid.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.AuctionRound != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AuctionRound))
		i--
		dAtA[i] = 0x10
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
	if m.AuctionRound != 0 {
		n += 1 + sovGenesis(uint64(m.AuctionRound))
	}
	if m.HighestBid != nil {
		l = m.HighestBid.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.AuctionEndingTimestamp != 0 {
		n += 1 + sovGenesis(uint64(m.AuctionEndingTimestamp))
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionRound", wireType)
			}
			m.AuctionRound = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionRound |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HighestBid", wireType)
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
			if m.HighestBid == nil {
				m.HighestBid = &Bid{}
			}
			if err := m.HighestBid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionEndingTimestamp", wireType)
			}
			m.AuctionEndingTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionEndingTimestamp |= int64(b&0x7F) << shift
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
