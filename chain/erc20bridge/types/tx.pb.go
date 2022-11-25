// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kaiju/erc20bridge/tx.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// MsgInitHub describes a message to init ERC20 parent contract inside bridge keeper
type MsgInitHub struct {
	HubAddress string                                        `protobuf:"bytes,1,opt,name=hub_address,json=hubAddress,proto3" json:"hub_address,omitempty" yaml:"hub_address"`
	Proposer   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty" yaml:"proposer"`
}

func (m *MsgInitHub) Reset()         { *m = MsgInitHub{} }
func (m *MsgInitHub) String() string { return proto.CompactTextString(m) }
func (*MsgInitHub) ProtoMessage()    {}
func (*MsgInitHub) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3d69c0f4664bc25, []int{0}
}
func (m *MsgInitHub) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitHub) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitHub.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitHub) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitHub.Merge(m, src)
}
func (m *MsgInitHub) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitHub) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitHub.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitHub proto.InternalMessageInfo

func (m *MsgInitHub) GetHubAddress() string {
	if m != nil {
		return m.HubAddress
	}
	return ""
}

func (m *MsgInitHub) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

type MsgERC20BridgeMint struct {
	// ID of token mapping registered on bridge
	MappingId string `protobuf:"bytes,1,opt,name=mapping_id,json=mappingId,proto3" json:"mapping_id,omitempty"`
	// amount of token to mint
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// iEVM address to receive ERC20 token
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// cosmos address
	Proposer github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,4,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty" yaml:"proposer"`
}

func (m *MsgERC20BridgeMint) Reset()         { *m = MsgERC20BridgeMint{} }
func (m *MsgERC20BridgeMint) String() string { return proto.CompactTextString(m) }
func (*MsgERC20BridgeMint) ProtoMessage()    {}
func (*MsgERC20BridgeMint) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3d69c0f4664bc25, []int{1}
}
func (m *MsgERC20BridgeMint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgERC20BridgeMint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgERC20BridgeMint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgERC20BridgeMint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgERC20BridgeMint.Merge(m, src)
}
func (m *MsgERC20BridgeMint) XXX_Size() int {
	return m.Size()
}
func (m *MsgERC20BridgeMint) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgERC20BridgeMint.DiscardUnknown(m)
}

var xxx_messageInfo_MsgERC20BridgeMint proto.InternalMessageInfo

func (m *MsgERC20BridgeMint) GetMappingId() string {
	if m != nil {
		return m.MappingId
	}
	return ""
}

func (m *MsgERC20BridgeMint) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *MsgERC20BridgeMint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgERC20BridgeMint) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgInitHub)(nil), "kaiju.erc20bridge.v1beta1.MsgInitHub")
	proto.RegisterType((*MsgERC20BridgeMint)(nil), "kaiju.erc20bridge.v1beta1.MsgERC20BridgeMint")
}

func init() { proto.RegisterFile("kaiju/erc20bridge/tx.proto", fileDescriptor_f3d69c0f4664bc25) }

var fileDescriptor_f3d69c0f4664bc25 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x3b, 0xf7, 0x5e, 0x7a, 0xed, 0x28, 0x0a, 0x83, 0x94, 0x22, 0x34, 0x29, 0x59, 0x75,
	0x61, 0x93, 0xb6, 0x2e, 0x04, 0x77, 0x8d, 0x0a, 0x16, 0xec, 0xc2, 0x2c, 0xdd, 0xd4, 0x4c, 0x32,
	0x24, 0x63, 0x9b, 0x4c, 0x98, 0x99, 0x14, 0xfb, 0x16, 0xbe, 0x85, 0xaf, 0xe2, 0xb2, 0x4b, 0x41,
	0x08, 0xd2, 0xbe, 0x41, 0x97, 0xae, 0xa4, 0xf9, 0x47, 0xfa, 0x00, 0xae, 0x66, 0xce, 0x7c, 0xe7,
	0x1c, 0xbe, 0x1f, 0xdf, 0x40, 0x85, 0x86, 0xcf, 0xc4, 0x91, 0x74, 0x41, 0x0c, 0xc2, 0x9d, 0x61,
	0x1f, 0x73, 0xea, 0x7a, 0xc4, 0x90, 0x2f, 0x7a, 0xc4, 0x99, 0x64, 0xa8, 0x5d, 0xea, 0x7a, 0x45,
	0xd7, 0x17, 0x03, 0x4c, 0xa4, 0x3d, 0x38, 0x3b, 0xf5, 0x98, 0xc7, 0xd2, 0x4e, 0x63, 0x77, 0xcb,
	0x86, 0xb4, 0x37, 0x00, 0xe1, 0x44, 0x78, 0xe3, 0x90, 0xca, 0xbb, 0x18, 0xa3, 0x4b, 0x78, 0xe8,
	0xc7, 0x78, 0x6a, 0xbb, 0x2e, 0x27, 0x42, 0xb4, 0x40, 0x07, 0x74, 0x1b, 0x66, 0x73, 0x9b, 0xa8,
	0x68, 0x69, 0x07, 0xf3, 0x2b, 0xad, 0x22, 0x6a, 0x16, 0xf4, 0x63, 0x3c, 0xca, 0x0a, 0xf4, 0x04,
	0x0f, 0x22, 0xce, 0x22, 0x26, 0x08, 0x6f, 0xfd, 0xe9, 0x80, 0xee, 0x91, 0x79, 0xb3, 0x4d, 0xd4,
	0x93, 0x6c, 0xaa, 0x50, 0xb4, 0xef, 0x44, 0xed, 0x79, 0x54, 0xfa, 0x31, 0xd6, 0x1d, 0x16, 0x18,
	0x0e, 0x13, 0x01, 0x13, 0xf9, 0xd1, 0x13, 0xee, 0xcc, 0x90, 0xcb, 0x88, 0x08, 0x7d, 0xe4, 0x38,
	0xf9, 0x5e, 0xab, 0xdc, 0xaa, 0x7d, 0x02, 0x88, 0x26, 0xc2, 0xbb, 0xb5, 0xae, 0x87, 0x7d, 0x33,
	0x45, 0x9b, 0xd0, 0x50, 0xa2, 0x36, 0x84, 0x81, 0x1d, 0x45, 0x34, 0xf4, 0xa6, 0xd4, 0xcd, 0x0c,
	0x5b, 0x8d, 0xfc, 0x65, 0xec, 0xa2, 0x26, 0xac, 0xdb, 0x01, 0x8b, 0x43, 0x99, 0xba, 0x6a, 0x58,
	0x79, 0x85, 0xce, 0xe1, 0xff, 0x02, 0xf2, 0x6f, 0x0a, 0x89, 0xb6, 0x89, 0x7a, 0x9c, 0xd9, 0x2d,
	0x01, 0x8b, 0x96, 0x3d, 0xba, 0x7f, 0xbf, 0x41, 0x67, 0xce, 0xde, 0xd7, 0x0a, 0x58, 0xad, 0x15,
	0xf0, 0xb5, 0x56, 0xc0, 0xeb, 0x46, 0xa9, 0xad, 0x36, 0x4a, 0xed, 0x63, 0xa3, 0xd4, 0x1e, 0x1f,
	0x2a, 0x2b, 0xc7, 0x45, 0xc2, 0xf7, 0x36, 0x16, 0x46, 0x99, 0x77, 0xcf, 0x61, 0x9c, 0x54, 0x4b,
	0xdf, 0xa6, 0xa1, 0x11, 0x30, 0x37, 0x9e, 0x13, 0xb1, 0xff, 0x59, 0x76, 0x0e, 0x70, 0x3d, 0xcd,
	0xfe, 0xe2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x61, 0x6a, 0x65, 0x85, 0x52, 0x02, 0x00, 0x00,
}

func (m *MsgInitHub) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitHub) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitHub) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.HubAddress) > 0 {
		i -= len(m.HubAddress)
		copy(dAtA[i:], m.HubAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.HubAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgERC20BridgeMint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgERC20BridgeMint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgERC20BridgeMint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MappingId) > 0 {
		i -= len(m.MappingId)
		copy(dAtA[i:], m.MappingId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MappingId)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *MsgInitHub) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HubAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgERC20BridgeMint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MappingId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInitHub) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgInitHub: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitHub: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HubAddress", wireType)
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
			m.HubAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
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
func (m *MsgERC20BridgeMint) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgERC20BridgeMint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgERC20BridgeMint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MappingId", wireType)
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
			m.MappingId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
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
