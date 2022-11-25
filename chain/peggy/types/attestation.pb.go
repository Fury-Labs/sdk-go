// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kaiju/peggy/v1/attestation.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// ClaimType is the cosmos type of an event from the counterpart chain that can
// be handled
type ClaimType int32

const (
	CLAIM_TYPE_UNKNOWN        ClaimType = 0
	CLAIM_TYPE_DEPOSIT        ClaimType = 1
	CLAIM_TYPE_WITHDRAW       ClaimType = 2
	CLAIM_TYPE_ERC20_DEPLOYED ClaimType = 3
	CLAIM_TYPE_VALSET_UPDATED ClaimType = 4
)

var ClaimType_name = map[int32]string{
	0: "CLAIM_TYPE_UNKNOWN",
	1: "CLAIM_TYPE_DEPOSIT",
	2: "CLAIM_TYPE_WITHDRAW",
	3: "CLAIM_TYPE_ERC20_DEPLOYED",
	4: "CLAIM_TYPE_VALSET_UPDATED",
}

var ClaimType_value = map[string]int32{
	"CLAIM_TYPE_UNKNOWN":        0,
	"CLAIM_TYPE_DEPOSIT":        1,
	"CLAIM_TYPE_WITHDRAW":       2,
	"CLAIM_TYPE_ERC20_DEPLOYED": 3,
	"CLAIM_TYPE_VALSET_UPDATED": 4,
}

func (x ClaimType) String() string {
	return proto.EnumName(ClaimType_name, int32(x))
}

func (ClaimType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3022043570947e1e, []int{0}
}

// Attestation is an aggregate of `claims` that eventually becomes `observed` by
// all orchestrators
// EVENT_NONCE:
// EventNonce a nonce provided by the peggy contract that is unique per event fired
// These event nonces must be relayed in order. This is a correctness issue,
// if relaying out of order transaction replay attacks become possible
// OBSERVED:
// Observed indicates that >67% of validators have attested to the event,
// and that the event should be executed by the peggy state machine
//
// The actual content of the claims is passed in with the transaction making the claim
// and then passed through the call stack alongside the attestation while it is processed
// the key in which the attestation is stored is keyed on the exact details of the claim
// but there is no reason to store those exact details becuause the next message sender
// will kindly provide you with them.
type Attestation struct {
	Observed bool       `protobuf:"varint,1,opt,name=observed,proto3" json:"observed,omitempty"`
	Votes    []string   `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes,omitempty"`
	Height   uint64     `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Claim    *types.Any `protobuf:"bytes,4,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (m *Attestation) Reset()         { *m = Attestation{} }
func (m *Attestation) String() string { return proto.CompactTextString(m) }
func (*Attestation) ProtoMessage()    {}
func (*Attestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3022043570947e1e, []int{0}
}
func (m *Attestation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Attestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Attestation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Attestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attestation.Merge(m, src)
}
func (m *Attestation) XXX_Size() int {
	return m.Size()
}
func (m *Attestation) XXX_DiscardUnknown() {
	xxx_messageInfo_Attestation.DiscardUnknown(m)
}

var xxx_messageInfo_Attestation proto.InternalMessageInfo

func (m *Attestation) GetObserved() bool {
	if m != nil {
		return m.Observed
	}
	return false
}

func (m *Attestation) GetVotes() []string {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *Attestation) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Attestation) GetClaim() *types.Any {
	if m != nil {
		return m.Claim
	}
	return nil
}

// ERC20Token unique identifier for an Ethereum ERC20 token.
// CONTRACT:
// The contract address on ETH of the token, this could be a Cosmos
// originated token, if so it will be the ERC20 address of the representation
// (note: developers should look up the token symbol using the address on ETH to display for UI)
type ERC20Token struct {
	Contract string                                 `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Amount   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *ERC20Token) Reset()         { *m = ERC20Token{} }
func (m *ERC20Token) String() string { return proto.CompactTextString(m) }
func (*ERC20Token) ProtoMessage()    {}
func (*ERC20Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_3022043570947e1e, []int{1}
}
func (m *ERC20Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ERC20Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ERC20Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ERC20Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ERC20Token.Merge(m, src)
}
func (m *ERC20Token) XXX_Size() int {
	return m.Size()
}
func (m *ERC20Token) XXX_DiscardUnknown() {
	xxx_messageInfo_ERC20Token.DiscardUnknown(m)
}

var xxx_messageInfo_ERC20Token proto.InternalMessageInfo

func (m *ERC20Token) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func init() {
	proto.RegisterEnum("kaiju.peggy.v1.ClaimType", ClaimType_name, ClaimType_value)
	proto.RegisterType((*Attestation)(nil), "kaiju.peggy.v1.Attestation")
	proto.RegisterType((*ERC20Token)(nil), "kaiju.peggy.v1.ERC20Token")
}

func init() {
	proto.RegisterFile("kaiju/peggy/v1/attestation.proto", fileDescriptor_3022043570947e1e)
}

var fileDescriptor_3022043570947e1e = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x31, 0x6f, 0x9b, 0x40,
	0x1c, 0xc5, 0xc1, 0x76, 0xac, 0xf8, 0xb2, 0x58, 0x57, 0x2b, 0x25, 0x48, 0x25, 0x28, 0xaa, 0x2a,
	0x2b, 0x52, 0xee, 0x92, 0x74, 0xed, 0x42, 0x0c, 0x55, 0x51, 0x5c, 0xdb, 0x22, 0xa4, 0x56, 0xba,
	0x58, 0x80, 0xaf, 0x98, 0xc6, 0x70, 0xc8, 0x77, 0x46, 0x62, 0xee, 0x52, 0x79, 0xea, 0x17, 0xb0,
	0x3a, 0xf4, 0xcb, 0x64, 0xcc, 0x58, 0x75, 0x88, 0x2a, 0xfb, 0x8b, 0x54, 0x06, 0xe2, 0xba, 0x49,
	0x3a, 0xc1, 0xbb, 0xf7, 0xfb, 0x1f, 0x4f, 0x7f, 0x1e, 0x78, 0x19, 0x44, 0x9f, 0x89, 0xc7, 0x83,
	0x84, 0xe0, 0x98, 0xf8, 0x7e, 0x8a, 0x93, 0x13, 0xec, 0x70, 0x4e, 0x18, 0x77, 0x78, 0x40, 0x23,
	0x14, 0x4f, 0x28, 0xa7, 0x10, 0xae, 0x29, 0x94, 0x51, 0x28, 0x39, 0x91, 0x1b, 0x3e, 0xf5, 0x69,
	0x66, 0xe3, 0xd5, 0x5b, 0x4e, 0xca, 0x7b, 0x3e, 0xa5, 0xfe, 0x98, 0xe0, 0x4c, 0xb9, 0xd3, 0x4f,
	0xd8, 0x89, 0xd2, 0xdc, 0x3a, 0xf8, 0x22, 0x82, 0x1d, 0xed, 0xef, 0xd5, 0x50, 0x06, 0xdb, 0xd4,
	0x65, 0x64, 0x92, 0x90, 0xa1, 0x24, 0xaa, 0x62, 0x73, 0xdb, 0x5a, 0x6b, 0xd8, 0x00, 0x5b, 0x09,
	0xe5, 0x84, 0x49, 0x25, 0xb5, 0xdc, 0xac, 0x59, 0xb9, 0x80, 0xbb, 0xa0, 0x3a, 0x22, 0x81, 0x3f,
	0xe2, 0x52, 0x59, 0x15, 0x9b, 0x15, 0xab, 0x50, 0xf0, 0x10, 0x6c, 0x79, 0x63, 0x27, 0x08, 0xa5,
	0x8a, 0x2a, 0x36, 0x77, 0x4e, 0x1b, 0x28, 0x0f, 0x81, 0xee, 0x43, 0x20, 0x2d, 0x4a, 0xad, 0x1c,
	0x39, 0x88, 0x01, 0x30, 0xac, 0xd6, 0xe9, 0xb1, 0x4d, 0xaf, 0x49, 0x96, 0xc1, 0xa3, 0x11, 0x9f,
	0x38, 0x1e, 0xcf, 0x32, 0xd4, 0xac, 0xb5, 0x86, 0x6f, 0x41, 0xd5, 0x09, 0xe9, 0x34, 0xe2, 0x52,
	0x69, 0xe5, 0x9c, 0xa1, 0x9b, 0xbb, 0x7d, 0xe1, 0xd7, 0xdd, 0xfe, 0x2b, 0x3f, 0xe0, 0xa3, 0xa9,
	0x8b, 0x3c, 0x1a, 0x62, 0x8f, 0xb2, 0x90, 0xb2, 0xe2, 0x71, 0xc4, 0x86, 0xd7, 0x98, 0xa7, 0x31,
	0x61, 0xc8, 0x8c, 0xb8, 0x55, 0x4c, 0x1f, 0x7e, 0x2f, 0x81, 0x5a, 0x6b, 0xf5, 0x6d, 0x3b, 0x8d,
	0x09, 0x44, 0x00, 0xb6, 0xda, 0x9a, 0xf9, 0x7e, 0x60, 0x5f, 0xf5, 0x8c, 0xc1, 0x65, 0xe7, 0xbc,
	0xd3, 0xed, 0x77, 0xea, 0x82, 0xbc, 0x3b, 0x9b, 0xab, 0x4f, 0x38, 0x0f, 0x78, 0xdd, 0xe8, 0x75,
	0x2f, 0x4c, 0xbb, 0x2e, 0x3e, 0xe2, 0x0b, 0x07, 0x1e, 0x83, 0x67, 0x1b, 0xa7, 0x7d, 0xd3, 0x7e,
	0xa7, 0x5b, 0x5a, 0xbf, 0x5e, 0x92, 0x9f, 0xcf, 0xe6, 0xea, 0x53, 0x16, 0x7c, 0x03, 0xf6, 0x36,
	0x8e, 0xb3, 0xe5, 0xac, 0x6e, 0x6b, 0x77, 0xaf, 0x0c, 0xbd, 0x5e, 0x96, 0x5f, 0xcc, 0xe6, 0xea,
	0xff, 0x81, 0x07, 0xd3, 0x1f, 0xb4, 0xf6, 0x85, 0x61, 0x0f, 0x2e, 0x7b, 0xba, 0x66, 0x1b, 0x7a,
	0xbd, 0xf2, 0x68, 0xfa, 0x5f, 0x40, 0xae, 0x7c, 0xfd, 0xa1, 0x08, 0x67, 0xe4, 0x66, 0xa1, 0x88,
	0xb7, 0x0b, 0x45, 0xfc, 0xbd, 0x50, 0xc4, 0x6f, 0x4b, 0x45, 0xb8, 0x5d, 0x2a, 0xc2, 0xcf, 0xa5,
	0x22, 0x7c, 0x3c, 0xdf, 0xd8, 0xb5, 0x79, 0xdf, 0xc1, 0xb6, 0xe3, 0x32, 0xbc, 0x6e, 0xe4, 0x91,
	0x47, 0x27, 0x64, 0x53, 0x8e, 0x9c, 0x20, 0xc2, 0x21, 0x1d, 0x4e, 0xc7, 0x84, 0x15, 0xa5, 0xce,
	0x7e, 0x8a, 0x5b, 0xcd, 0xfa, 0xf0, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xa2, 0x82,
	0x9c, 0xf4, 0x02, 0x00, 0x00,
}

func (m *Attestation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Attestation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Attestation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAttestation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Height != 0 {
		i = encodeVarintAttestation(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Votes[iNdEx])
			copy(dAtA[i:], m.Votes[iNdEx])
			i = encodeVarintAttestation(dAtA, i, uint64(len(m.Votes[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Observed {
		i--
		if m.Observed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ERC20Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ERC20Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ERC20Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAttestation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintAttestation(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAttestation(dAtA []byte, offset int, v uint64) int {
	offset -= sovAttestation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Attestation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Observed {
		n += 2
	}
	if len(m.Votes) > 0 {
		for _, s := range m.Votes {
			l = len(s)
			n += 1 + l + sovAttestation(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovAttestation(uint64(m.Height))
	}
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovAttestation(uint64(l))
	}
	return n
}

func (m *ERC20Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAttestation(uint64(l))
	return n
}

func sovAttestation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAttestation(x uint64) (n int) {
	return sovAttestation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Attestation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: Attestation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attestation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Observed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
			m.Observed = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &types.Any{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttestation
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
func (m *ERC20Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: ERC20Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ERC20Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttestation
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
func skipAttestation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
				return 0, ErrInvalidLengthAttestation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAttestation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAttestation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAttestation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttestation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAttestation = fmt.Errorf("proto: unexpected end of group")
)
