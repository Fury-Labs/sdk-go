// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kaiju/auction/v1beta1/auction.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	// auction_period_duration defines the auction period duration
	AuctionPeriod int64 `protobuf:"varint,1,opt,name=auction_period,json=auctionPeriod,proto3" json:"auction_period,omitempty"`
	// min_next_bid_increment_rate defines the minimum increment rate for new bids
	MinNextBidIncrementRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=min_next_bid_increment_rate,json=minNextBidIncrementRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_next_bid_increment_rate"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_49edfee5f1ef4b5a, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAuctionPeriod() int64 {
	if m != nil {
		return m.AuctionPeriod
	}
	return 0
}

type Bid struct {
	Bidder string                                  `protobuf:"bytes,1,opt,name=bidder,proto3" json:"bidder" yaml:"bidder"`
	Amount github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
}

func (m *Bid) Reset()         { *m = Bid{} }
func (m *Bid) String() string { return proto.CompactTextString(m) }
func (*Bid) ProtoMessage()    {}
func (*Bid) Descriptor() ([]byte, []int) {
	return fileDescriptor_49edfee5f1ef4b5a, []int{1}
}
func (m *Bid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bid.Merge(m, src)
}
func (m *Bid) XXX_Size() int {
	return m.Size()
}
func (m *Bid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bid.DiscardUnknown(m)
}

var xxx_messageInfo_Bid proto.InternalMessageInfo

func (m *Bid) GetBidder() string {
	if m != nil {
		return m.Bidder
	}
	return ""
}

type EventBid struct {
	// bidder describes the address of bidder
	Bidder string `protobuf:"bytes,1,opt,name=bidder,proto3" json:"bidder,omitempty"`
	// amount describes the amount the bidder put on the auction
	Amount github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
	// round defines the round number of auction
	Round uint64 `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
}

func (m *EventBid) Reset()         { *m = EventBid{} }
func (m *EventBid) String() string { return proto.CompactTextString(m) }
func (*EventBid) ProtoMessage()    {}
func (*EventBid) Descriptor() ([]byte, []int) {
	return fileDescriptor_49edfee5f1ef4b5a, []int{2}
}
func (m *EventBid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBid.Merge(m, src)
}
func (m *EventBid) XXX_Size() int {
	return m.Size()
}
func (m *EventBid) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBid.DiscardUnknown(m)
}

var xxx_messageInfo_EventBid proto.InternalMessageInfo

func (m *EventBid) GetBidder() string {
	if m != nil {
		return m.Bidder
	}
	return ""
}

func (m *EventBid) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

type EventAuctionResult struct {
	// winner describes the address of the winner
	Winner string `protobuf:"bytes,1,opt,name=winner,proto3" json:"winner,omitempty"`
	// amount describes the amount the winner get from the auction
	Amount github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
	// round defines the round number of auction
	Round uint64 `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
}

func (m *EventAuctionResult) Reset()         { *m = EventAuctionResult{} }
func (m *EventAuctionResult) String() string { return proto.CompactTextString(m) }
func (*EventAuctionResult) ProtoMessage()    {}
func (*EventAuctionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_49edfee5f1ef4b5a, []int{3}
}
func (m *EventAuctionResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAuctionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAuctionResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAuctionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAuctionResult.Merge(m, src)
}
func (m *EventAuctionResult) XXX_Size() int {
	return m.Size()
}
func (m *EventAuctionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAuctionResult.DiscardUnknown(m)
}

var xxx_messageInfo_EventAuctionResult proto.InternalMessageInfo

func (m *EventAuctionResult) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func (m *EventAuctionResult) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

type EventAuctionStart struct {
	// round defines the round number of auction
	Round uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	// ending_timestamp describes auction end time
	EndingTimestamp int64 `protobuf:"varint,2,opt,name=ending_timestamp,json=endingTimestamp,proto3" json:"ending_timestamp,omitempty"`
	// new_basket describes auction module balance at the time of new auction start
	NewBasket github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=new_basket,json=newBasket,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"new_basket"`
}

func (m *EventAuctionStart) Reset()         { *m = EventAuctionStart{} }
func (m *EventAuctionStart) String() string { return proto.CompactTextString(m) }
func (*EventAuctionStart) ProtoMessage()    {}
func (*EventAuctionStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_49edfee5f1ef4b5a, []int{4}
}
func (m *EventAuctionStart) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAuctionStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAuctionStart.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAuctionStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAuctionStart.Merge(m, src)
}
func (m *EventAuctionStart) XXX_Size() int {
	return m.Size()
}
func (m *EventAuctionStart) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAuctionStart.DiscardUnknown(m)
}

var xxx_messageInfo_EventAuctionStart proto.InternalMessageInfo

func (m *EventAuctionStart) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *EventAuctionStart) GetEndingTimestamp() int64 {
	if m != nil {
		return m.EndingTimestamp
	}
	return 0
}

func (m *EventAuctionStart) GetNewBasket() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.NewBasket
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "kaiju.auction.v1beta1.Params")
	proto.RegisterType((*Bid)(nil), "kaiju.auction.v1beta1.Bid")
	proto.RegisterType((*EventBid)(nil), "kaiju.auction.v1beta1.EventBid")
	proto.RegisterType((*EventAuctionResult)(nil), "kaiju.auction.v1beta1.EventAuctionResult")
	proto.RegisterType((*EventAuctionStart)(nil), "kaiju.auction.v1beta1.EventAuctionStart")
}

func init() {
	proto.RegisterFile("kaiju/auction/v1beta1/auction.proto", fileDescriptor_49edfee5f1ef4b5a)
}

var fileDescriptor_49edfee5f1ef4b5a = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x73, 0xb8, 0x44, 0xe4, 0x50, 0xf9, 0x73, 0xaa, 0x20, 0x6d, 0x25, 0x3b, 0xb2, 0x04,
	0x0d, 0x43, 0x6d, 0x4a, 0xb7, 0x6e, 0x18, 0x10, 0xaa, 0x04, 0xa8, 0x32, 0x4c, 0x2c, 0xd6, 0xd9,
	0xf7, 0x2a, 0xbd, 0x36, 0x77, 0x17, 0xf9, 0xce, 0x49, 0x3b, 0x22, 0x46, 0x16, 0x3e, 0x02, 0x12,
	0x1b, 0x9f, 0x82, 0xb1, 0x63, 0x47, 0xc4, 0x10, 0x50, 0xb2, 0x20, 0x46, 0x3e, 0x01, 0x8a, 0x7d,
	0x0e, 0x66, 0xeb, 0x00, 0x53, 0xf2, 0x3c, 0xf7, 0xdc, 0x73, 0xbf, 0x93, 0xef, 0xc5, 0x5b, 0x5c,
	0x1e, 0x41, 0x66, 0xf8, 0x18, 0x42, 0x5a, 0x64, 0x86, 0x2b, 0x19, 0x8e, 0x77, 0x52, 0x30, 0x74,
	0xa7, 0xd6, 0xc1, 0x28, 0x57, 0x46, 0x91, 0xf5, 0x65, 0x30, 0xa8, 0x17, 0x6c, 0x70, 0x63, 0x6d,
	0xa0, 0x06, 0xaa, 0x4c, 0x85, 0x8b, 0x7f, 0xd5, 0x86, 0x0d, 0x37, 0x53, 0x5a, 0x28, 0x1d, 0xa6,
	0x54, 0xc3, 0xb2, 0x33, 0x53, 0xdc, 0x16, 0xfa, 0x1f, 0x11, 0x6e, 0x1f, 0xd0, 0x9c, 0x0a, 0x4d,
	0xee, 0xe0, 0x6b, 0xb6, 0x33, 0x19, 0x41, 0xce, 0x15, 0xeb, 0xa2, 0x1e, 0xea, 0x3b, 0xf1, 0xaa,
	0x75, 0x0f, 0x4a, 0x93, 0x0c, 0xf1, 0xa6, 0xe0, 0x32, 0x91, 0x70, 0x62, 0x92, 0x94, 0xb3, 0x84,
	0xcb, 0x2c, 0x07, 0x01, 0xd2, 0x24, 0x39, 0x35, 0xd0, 0xbd, 0xd4, 0x43, 0xfd, 0x4e, 0x14, 0x9c,
	0x4d, 0xbd, 0xd6, 0xd7, 0xa9, 0x77, 0x77, 0xc0, 0xcd, 0x61, 0x91, 0x06, 0x99, 0x12, 0xa1, 0x25,
	0xa9, 0x7e, 0xb6, 0x35, 0x3b, 0x0e, 0xcd, 0xe9, 0x08, 0x74, 0xf0, 0x18, 0xb2, 0xf8, 0xb6, 0xe0,
	0xf2, 0x05, 0x9c, 0x98, 0x88, 0xb3, 0xfd, 0xba, 0x2f, 0xa6, 0x06, 0xf6, 0x56, 0x7e, 0x7c, 0xf0,
	0x90, 0xff, 0x16, 0x61, 0x27, 0xe2, 0x8c, 0xec, 0xe2, 0x76, 0xca, 0x19, 0x83, 0xbc, 0x44, 0xeb,
	0x44, 0x9b, 0x3f, 0xa7, 0x9e, 0x75, 0x7e, 0x4d, 0xbd, 0xd5, 0x53, 0x2a, 0x86, 0x7b, 0x7e, 0xa5,
	0xfd, 0xd8, 0x2e, 0x90, 0xa7, 0xb8, 0x4d, 0x85, 0x2a, 0xa4, 0xb1, 0x6c, 0xa1, 0x65, 0xdb, 0xba,
	0x00, 0xdb, 0x23, 0xc5, 0x65, 0x6c, 0xb7, 0xfb, 0x6f, 0x10, 0xbe, 0xf2, 0x64, 0x0c, 0x72, 0x41,
	0x49, 0x6e, 0xfd, 0x8d, 0xf2, 0xcf, 0x4f, 0x23, 0x6b, 0xf8, 0x72, 0xae, 0x0a, 0xc9, 0xba, 0x4e,
	0x0f, 0xf5, 0x57, 0xe2, 0x4a, 0xf8, 0xef, 0x10, 0x26, 0x25, 0xc3, 0xc3, 0xea, 0xa3, 0xc4, 0xa0,
	0x8b, 0xa1, 0x59, 0xd0, 0x4c, 0xb8, 0x94, 0x7f, 0x68, 0x2a, 0xf5, 0xbf, 0x69, 0x3e, 0x23, 0x7c,
	0xb3, 0x49, 0xf3, 0xd2, 0xd0, 0xbc, 0x91, 0x45, 0x8d, 0x2c, 0xb9, 0x87, 0x6f, 0x80, 0x64, 0x5c,
	0x0e, 0x12, 0xc3, 0x05, 0x68, 0x43, 0xc5, 0xa8, 0x84, 0x72, 0xe2, 0xeb, 0x95, 0xff, 0xaa, 0xb6,
	0xc9, 0x11, 0xc6, 0x12, 0x26, 0x49, 0x4a, 0xf5, 0x31, 0x98, 0xae, 0xd3, 0x73, 0xfa, 0x57, 0x1f,
	0xac, 0x07, 0x15, 0x60, 0xb0, 0x78, 0xc9, 0xf5, 0xa3, 0x2f, 0x19, 0xa3, 0xfb, 0x8b, 0x4b, 0x7d,
	0xfa, 0xe6, 0xf5, 0x2f, 0x78, 0x29, 0x1d, 0x77, 0x24, 0x4c, 0xa2, 0xb2, 0x3d, 0x1a, 0x9c, 0xcd,
	0x5c, 0x74, 0x3e, 0x73, 0xd1, 0xf7, 0x99, 0x8b, 0xde, 0xcf, 0xdd, 0xd6, 0xf9, 0xdc, 0x6d, 0x7d,
	0x99, 0xbb, 0xad, 0xd7, 0xcf, 0x1b, 0x75, 0xfb, 0xf5, 0xd8, 0x3d, 0xa3, 0xa9, 0x0e, 0x97, 0x43,
	0xb8, 0x9d, 0xa9, 0x1c, 0x9a, 0xf2, 0x90, 0x72, 0x19, 0x0a, 0xc5, 0x8a, 0x21, 0xe8, 0xe5, 0x28,
	0x97, 0x27, 0xa7, 0xed, 0x72, 0xe0, 0x76, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x6d, 0xcb,
	0x8b, 0xec, 0x03, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.AuctionPeriod != that1.AuctionPeriod {
		return false
	}
	if !this.MinNextBidIncrementRate.Equal(that1.MinNextBidIncrementRate) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MinNextBidIncrementRate.Size()
		i -= size
		if _, err := m.MinNextBidIncrementRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AuctionPeriod != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.AuctionPeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Bid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Round != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventAuctionResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAuctionResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAuctionResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Round != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventAuctionStart) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAuctionStart) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAuctionStart) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewBasket) > 0 {
		for iNdEx := len(m.NewBasket) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NewBasket[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuction(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.EndingTimestamp != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.EndingTimestamp))
		i--
		dAtA[i] = 0x10
	}
	if m.Round != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionPeriod != 0 {
		n += 1 + sovAuction(uint64(m.AuctionPeriod))
	}
	l = m.MinNextBidIncrementRate.Size()
	n += 1 + l + sovAuction(uint64(l))
	return n
}

func (m *Bid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAuction(uint64(l))
	return n
}

func (m *EventBid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAuction(uint64(l))
	if m.Round != 0 {
		n += 1 + sovAuction(uint64(m.Round))
	}
	return n
}

func (m *EventAuctionResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAuction(uint64(l))
	if m.Round != 0 {
		n += 1 + sovAuction(uint64(m.Round))
	}
	return n
}

func (m *EventAuctionStart) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovAuction(uint64(m.Round))
	}
	if m.EndingTimestamp != 0 {
		n += 1 + sovAuction(uint64(m.EndingTimestamp))
	}
	if len(m.NewBasket) > 0 {
		for _, e := range m.NewBasket {
			l = e.Size()
			n += 1 + l + sovAuction(uint64(l))
		}
	}
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionPeriod", wireType)
			}
			m.AuctionPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinNextBidIncrementRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinNextBidIncrementRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *Bid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Bid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
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
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *EventBid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: EventBid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *EventAuctionResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: EventAuctionResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAuctionResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *EventAuctionStart) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: EventAuctionStart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAuctionStart: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndingTimestamp", wireType)
			}
			m.EndingTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndingTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewBasket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewBasket = append(m.NewBasket, types.Coin{})
			if err := m.NewBasket[len(m.NewBasket)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)
