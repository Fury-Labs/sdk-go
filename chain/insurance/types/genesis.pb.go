// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kaiju/insurance/v1beta1/genesis.proto

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

// GenesisState defines the insurance module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to insurance.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// insurance_funds describes the insurance funds available for derivative markets
	InsuranceFunds []InsuranceFund `protobuf:"bytes,2,rep,name=insurance_funds,json=insuranceFunds,proto3" json:"insurance_funds"`
	// redemption_schedule describes the redemption requests pending
	RedemptionSchedule []RedemptionSchedule `protobuf:"bytes,3,rep,name=redemption_schedule,json=redemptionSchedule,proto3" json:"redemption_schedule"`
	// next_share_denom_id describes the next share denom id to be used for newly creating insurance fund
	// incremented by 1 per insurance fund creation
	NextShareDenomId uint64 `protobuf:"varint,4,opt,name=next_share_denom_id,json=nextShareDenomId,proto3" json:"next_share_denom_id,omitempty"`
	// next_redemption_schedule_id describes next redemption schedule id to be used for next schedule
	// incremented by 1 per redemption request
	NextRedemptionScheduleId uint64 `protobuf:"varint,5,opt,name=next_redemption_schedule_id,json=nextRedemptionScheduleId,proto3" json:"next_redemption_schedule_id,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_293324fee7d3f3b1, []int{0}
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

func (m *GenesisState) GetInsuranceFunds() []InsuranceFund {
	if m != nil {
		return m.InsuranceFunds
	}
	return nil
}

func (m *GenesisState) GetRedemptionSchedule() []RedemptionSchedule {
	if m != nil {
		return m.RedemptionSchedule
	}
	return nil
}

func (m *GenesisState) GetNextShareDenomId() uint64 {
	if m != nil {
		return m.NextShareDenomId
	}
	return 0
}

func (m *GenesisState) GetNextRedemptionScheduleId() uint64 {
	if m != nil {
		return m.NextRedemptionScheduleId
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "kaiju.insurance.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("kaiju/insurance/v1beta1/genesis.proto", fileDescriptor_293324fee7d3f3b1)
}

var fileDescriptor_293324fee7d3f3b1 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x93, 0xdb, 0xde, 0x2e, 0xd2, 0xcb, 0x55, 0x52, 0x17, 0xa1, 0x85, 0x58, 0x74, 0x53,
	0x95, 0x66, 0x68, 0x5d, 0xbb, 0xb0, 0x88, 0x12, 0x10, 0x94, 0x76, 0xa5, 0x9b, 0x30, 0xc9, 0x9c,
	0x26, 0x23, 0x66, 0x26, 0x64, 0x26, 0x45, 0xb7, 0x3e, 0x81, 0x8f, 0xd5, 0x65, 0x97, 0xae, 0x44,
	0xda, 0x17, 0x91, 0x4c, 0xd3, 0x58, 0xa8, 0x64, 0x77, 0x72, 0xfe, 0xef, 0xff, 0xcf, 0x39, 0x64,
	0x8c, 0x13, 0xca, 0x9e, 0x20, 0x90, 0x74, 0x06, 0x88, 0x32, 0x91, 0xa5, 0x98, 0x05, 0x80, 0x66,
	0x03, 0x1f, 0x24, 0x1e, 0xa0, 0x10, 0x18, 0x08, 0x2a, 0x9c, 0x24, 0xe5, 0x92, 0x9b, 0x9d, 0x12,
	0x75, 0x4a, 0xd4, 0x29, 0xd0, 0xf6, 0x59, 0x55, 0xce, 0x0f, 0xae, 0x92, 0xda, 0x07, 0x21, 0x0f,
	0xb9, 0x2a, 0x51, 0x5e, 0xad, 0xbb, 0x47, 0x6f, 0x35, 0xe3, 0xdf, 0xcd, 0x7a, 0xe2, 0x44, 0x62,
	0x09, 0xe6, 0xa5, 0xd1, 0x48, 0x70, 0x8a, 0x63, 0x61, 0xe9, 0x5d, 0xbd, 0xd7, 0x1c, 0x1e, 0x3b,
	0x15, 0x1b, 0x38, 0xf7, 0x0a, 0x1d, 0xd5, 0xe7, 0x9f, 0x87, 0xda, 0xb8, 0x30, 0x9a, 0x0f, 0xc6,
	0x5e, 0x49, 0x7a, 0xd3, 0x8c, 0x11, 0x61, 0xfd, 0xe9, 0xd6, 0x7a, 0xcd, 0xe1, 0x69, 0x65, 0x96,
	0xbb, 0xe9, 0x5c, 0x67, 0x8c, 0x14, 0x91, 0xff, 0xe9, 0x76, 0x53, 0x98, 0x53, 0xa3, 0x95, 0x02,
	0x81, 0x38, 0x91, 0x94, 0x33, 0x4f, 0x04, 0x11, 0x90, 0xec, 0x19, 0xac, 0x9a, 0x8a, 0x47, 0x95,
	0xf1, 0xe3, 0xd2, 0x37, 0x29, 0x6c, 0xc5, 0x0c, 0x33, 0xdd, 0x51, 0xcc, 0xbe, 0xd1, 0x62, 0xf0,
	0x22, 0x3d, 0x11, 0xe1, 0x14, 0x3c, 0x02, 0x8c, 0xc7, 0x1e, 0x25, 0x56, 0xbd, 0xab, 0xf7, 0xea,
	0xe3, 0xfd, 0x5c, 0x9a, 0xe4, 0xca, 0x55, 0x2e, 0xb8, 0xc4, 0xbc, 0x30, 0x3a, 0x0a, 0xff, 0x65,
	0xb7, 0xdc, 0xf6, 0x57, 0xd9, 0xac, 0x1c, 0xd9, 0xdd, 0xc2, 0x25, 0x23, 0x3a, 0x5f, 0xda, 0xfa,
	0x62, 0x69, 0xeb, 0x5f, 0x4b, 0x5b, 0x7f, 0x5f, 0xd9, 0xda, 0x62, 0x65, 0x6b, 0x1f, 0x2b, 0x5b,
	0x7b, 0xbc, 0x0b, 0xa9, 0x8c, 0x32, 0xdf, 0x09, 0x78, 0x8c, 0xdc, 0xcd, 0x71, 0xb7, 0xd8, 0x17,
	0xa8, 0x3c, 0xb5, 0x1f, 0xf0, 0x14, 0xb6, 0x3f, 0x23, 0x4c, 0x19, 0x8a, 0x79, 0x9e, 0x2d, 0xb6,
	0xde, 0x85, 0x7c, 0x4d, 0x40, 0xf8, 0x0d, 0xf5, 0xdb, 0xcf, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x41, 0xfa, 0xec, 0x9b, 0x83, 0x02, 0x00, 0x00,
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
	if m.NextRedemptionScheduleId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextRedemptionScheduleId))
		i--
		dAtA[i] = 0x28
	}
	if m.NextShareDenomId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextShareDenomId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.RedemptionSchedule) > 0 {
		for iNdEx := len(m.RedemptionSchedule) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RedemptionSchedule[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.InsuranceFunds) > 0 {
		for iNdEx := len(m.InsuranceFunds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InsuranceFunds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.InsuranceFunds) > 0 {
		for _, e := range m.InsuranceFunds {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RedemptionSchedule) > 0 {
		for _, e := range m.RedemptionSchedule {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NextShareDenomId != 0 {
		n += 1 + sovGenesis(uint64(m.NextShareDenomId))
	}
	if m.NextRedemptionScheduleId != 0 {
		n += 1 + sovGenesis(uint64(m.NextRedemptionScheduleId))
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
				return fmt.Errorf("proto: wrong wireType = %d for field InsuranceFunds", wireType)
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
			m.InsuranceFunds = append(m.InsuranceFunds, InsuranceFund{})
			if err := m.InsuranceFunds[len(m.InsuranceFunds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionSchedule", wireType)
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
			m.RedemptionSchedule = append(m.RedemptionSchedule, RedemptionSchedule{})
			if err := m.RedemptionSchedule[len(m.RedemptionSchedule)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextShareDenomId", wireType)
			}
			m.NextShareDenomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextShareDenomId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextRedemptionScheduleId", wireType)
			}
			m.NextRedemptionScheduleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextRedemptionScheduleId |= uint64(b&0x7F) << shift
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
