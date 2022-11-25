// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kaiju/oracle/v1beta1/genesis.proto

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

// GenesisState defines the oracle module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to oracle.
	Params                 Params                 `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	BandRelayers           []string               `protobuf:"bytes,2,rep,name=band_relayers,json=bandRelayers,proto3" json:"band_relayers,omitempty"`
	BandPriceStates        []*BandPriceState      `protobuf:"bytes,3,rep,name=band_price_states,json=bandPriceStates,proto3" json:"band_price_states,omitempty"`
	PriceFeedPriceStates   []*PriceFeedState      `protobuf:"bytes,4,rep,name=price_feed_price_states,json=priceFeedPriceStates,proto3" json:"price_feed_price_states,omitempty"`
	CoinbasePriceStates    []*CoinbasePriceState  `protobuf:"bytes,5,rep,name=coinbase_price_states,json=coinbasePriceStates,proto3" json:"coinbase_price_states,omitempty"`
	BandIbcPriceStates     []*BandPriceState      `protobuf:"bytes,6,rep,name=band_ibc_price_states,json=bandIbcPriceStates,proto3" json:"band_ibc_price_states,omitempty"`
	BandIbcOracleRequests  []*BandOracleRequest   `protobuf:"bytes,7,rep,name=band_ibc_oracle_requests,json=bandIbcOracleRequests,proto3" json:"band_ibc_oracle_requests,omitempty"`
	BandIbcParams          BandIBCParams          `protobuf:"bytes,8,opt,name=band_ibc_params,json=bandIbcParams,proto3" json:"band_ibc_params"`
	BandIbcLatestClientId  uint64                 `protobuf:"varint,9,opt,name=band_ibc_latest_client_id,json=bandIbcLatestClientId,proto3" json:"band_ibc_latest_client_id,omitempty"`
	CalldataRecords        []*CalldataRecord      `protobuf:"bytes,10,rep,name=calldata_records,json=calldataRecords,proto3" json:"calldata_records,omitempty"`
	BandIbcLatestRequestId uint64                 `protobuf:"varint,11,opt,name=band_ibc_latest_request_id,json=bandIbcLatestRequestId,proto3" json:"band_ibc_latest_request_id,omitempty"`
	ChainlinkPriceStates   []*ChainlinkPriceState `protobuf:"bytes,12,rep,name=chainlink_price_states,json=chainlinkPriceStates,proto3" json:"chainlink_price_states,omitempty"`
	HistoricalPriceRecords []*PriceRecords        `protobuf:"bytes,13,rep,name=historical_price_records,json=historicalPriceRecords,proto3" json:"historical_price_records,omitempty"`
	ProviderStates         []*ProviderState       `protobuf:"bytes,14,rep,name=provider_states,json=providerStates,proto3" json:"provider_states,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e14cf80151b4d2, []int{0}
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

func (m *GenesisState) GetBandRelayers() []string {
	if m != nil {
		return m.BandRelayers
	}
	return nil
}

func (m *GenesisState) GetBandPriceStates() []*BandPriceState {
	if m != nil {
		return m.BandPriceStates
	}
	return nil
}

func (m *GenesisState) GetPriceFeedPriceStates() []*PriceFeedState {
	if m != nil {
		return m.PriceFeedPriceStates
	}
	return nil
}

func (m *GenesisState) GetCoinbasePriceStates() []*CoinbasePriceState {
	if m != nil {
		return m.CoinbasePriceStates
	}
	return nil
}

func (m *GenesisState) GetBandIbcPriceStates() []*BandPriceState {
	if m != nil {
		return m.BandIbcPriceStates
	}
	return nil
}

func (m *GenesisState) GetBandIbcOracleRequests() []*BandOracleRequest {
	if m != nil {
		return m.BandIbcOracleRequests
	}
	return nil
}

func (m *GenesisState) GetBandIbcParams() BandIBCParams {
	if m != nil {
		return m.BandIbcParams
	}
	return BandIBCParams{}
}

func (m *GenesisState) GetBandIbcLatestClientId() uint64 {
	if m != nil {
		return m.BandIbcLatestClientId
	}
	return 0
}

func (m *GenesisState) GetCalldataRecords() []*CalldataRecord {
	if m != nil {
		return m.CalldataRecords
	}
	return nil
}

func (m *GenesisState) GetBandIbcLatestRequestId() uint64 {
	if m != nil {
		return m.BandIbcLatestRequestId
	}
	return 0
}

func (m *GenesisState) GetChainlinkPriceStates() []*ChainlinkPriceState {
	if m != nil {
		return m.ChainlinkPriceStates
	}
	return nil
}

func (m *GenesisState) GetHistoricalPriceRecords() []*PriceRecords {
	if m != nil {
		return m.HistoricalPriceRecords
	}
	return nil
}

func (m *GenesisState) GetProviderStates() []*ProviderState {
	if m != nil {
		return m.ProviderStates
	}
	return nil
}

type CalldataRecord struct {
	ClientId uint64 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Calldata []byte `protobuf:"bytes,2,opt,name=calldata,proto3" json:"calldata,omitempty"`
}

func (m *CalldataRecord) Reset()         { *m = CalldataRecord{} }
func (m *CalldataRecord) String() string { return proto.CompactTextString(m) }
func (*CalldataRecord) ProtoMessage()    {}
func (*CalldataRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e14cf80151b4d2, []int{1}
}
func (m *CalldataRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CalldataRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CalldataRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CalldataRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalldataRecord.Merge(m, src)
}
func (m *CalldataRecord) XXX_Size() int {
	return m.Size()
}
func (m *CalldataRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_CalldataRecord.DiscardUnknown(m)
}

var xxx_messageInfo_CalldataRecord proto.InternalMessageInfo

func (m *CalldataRecord) GetClientId() uint64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *CalldataRecord) GetCalldata() []byte {
	if m != nil {
		return m.Calldata
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "kaiju.oracle.v1beta1.GenesisState")
	proto.RegisterType((*CalldataRecord)(nil), "kaiju.oracle.v1beta1.CalldataRecord")
}

func init() {
	proto.RegisterFile("kaiju/oracle/v1beta1/genesis.proto", fileDescriptor_f7e14cf80151b4d2)
}

var fileDescriptor_f7e14cf80151b4d2 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0xe3, 0xb6, 0x4f, 0x9f, 0x66, 0x9b, 0x36, 0xb0, 0xb4, 0xc5, 0x04, 0xc9, 0x58, 0x45,
	0xb4, 0x96, 0xa0, 0xb6, 0x5a, 0x2e, 0x88, 0x03, 0x87, 0x44, 0x02, 0x59, 0xaa, 0x44, 0xe4, 0xc2,
	0x05, 0x0e, 0x66, 0xbd, 0xde, 0x26, 0x0b, 0x8e, 0xd7, 0xec, 0x6e, 0x22, 0xf5, 0x2b, 0x70, 0xe2,
	0x63, 0xf5, 0xd8, 0x23, 0x27, 0x84, 0x92, 0x2f, 0x82, 0xbc, 0x7e, 0x89, 0xdd, 0x2a, 0x89, 0xb8,
	0x79, 0xc6, 0x33, 0xbf, 0xf9, 0xcf, 0x8b, 0x0d, 0x8e, 0x68, 0xfc, 0x95, 0x60, 0x49, 0x27, 0xc4,
	0x61, 0x1c, 0xe1, 0x88, 0x38, 0x93, 0xd3, 0x80, 0x48, 0x74, 0xea, 0x0c, 0x48, 0x4c, 0x04, 0x15,
	0x76, 0xc2, 0x99, 0x64, 0x50, 0x2f, 0xe3, 0xec, 0x2c, 0xce, 0xce, 0xe3, 0x3a, 0xcf, 0x16, 0x12,
	0xf2, 0x40, 0x05, 0xe8, 0xec, 0x0d, 0xd8, 0x80, 0xa9, 0x47, 0x27, 0x7d, 0xca, 0xbc, 0x87, 0x3f,
	0x9a, 0xa0, 0xf5, 0x2e, 0x2b, 0x74, 0x21, 0x91, 0x24, 0xf0, 0x0d, 0xd8, 0x4c, 0x10, 0x47, 0x23,
	0xa1, 0x6b, 0xa6, 0x66, 0x6d, 0x9f, 0x99, 0xf6, 0xa2, 0xc2, 0x76, 0x5f, 0xc5, 0x75, 0x37, 0xae,
	0x7f, 0x3f, 0x69, 0x78, 0x79, 0x16, 0x7c, 0x0a, 0x76, 0x02, 0x14, 0x87, 0x3e, 0x27, 0x11, 0xba,
	0x22, 0x5c, 0xe8, 0x6b, 0xe6, 0xba, 0xd5, 0xf4, 0x5a, 0xa9, 0xd3, 0xcb, 0x7d, 0xf0, 0x03, 0xb8,
	0xaf, 0x82, 0x12, 0x4e, 0x31, 0xf1, 0x45, 0x5a, 0x58, 0xe8, 0xeb, 0xe6, 0xba, 0xb5, 0x7d, 0x66,
	0x2d, 0xae, 0xd7, 0x45, 0x71, 0xd8, 0x4f, 0x33, 0x94, 0x52, 0xaf, 0x1d, 0xd4, 0x6c, 0x01, 0x7d,
	0xf0, 0x30, 0x03, 0x5e, 0x12, 0x72, 0x8b, 0xbd, 0xb1, 0x8a, 0xad, 0x38, 0x6f, 0x09, 0x09, 0x33,
	0xf6, 0x5e, 0x52, 0xd8, 0xd5, 0x02, 0x5f, 0xc0, 0x3e, 0x66, 0x34, 0x0e, 0x90, 0x20, 0x75, 0xfc,
	0x7f, 0x0a, 0xff, 0x62, 0x31, 0xbe, 0x97, 0xa7, 0x55, 0xe4, 0x3f, 0xc0, 0x77, 0x7c, 0x02, 0x7e,
	0x06, 0xfb, 0x6a, 0x30, 0x34, 0xc0, 0xf5, 0x0a, 0x9b, 0xff, 0x38, 0x1c, 0x98, 0x62, 0xdc, 0x00,
	0x57, 0xe1, 0x21, 0xd0, 0x4b, 0x78, 0x96, 0xed, 0x73, 0xf2, 0x7d, 0x4c, 0x84, 0x14, 0xfa, 0xff,
	0x8a, 0xff, 0x7c, 0x39, 0xff, 0xbd, 0x72, 0x79, 0x59, 0x8e, 0xb7, 0x9f, 0x97, 0xa8, 0x79, 0x05,
	0xfc, 0x08, 0xda, 0xf3, 0x16, 0xb2, 0x4b, 0xda, 0x52, 0x97, 0x74, 0xbc, 0x1c, 0xee, 0x76, 0x7b,
	0xb5, 0x83, 0xda, 0x29, 0x3a, 0xc8, 0xee, 0xea, 0x15, 0x78, 0x54, 0x62, 0xa3, 0xb4, 0x1d, 0xe9,
	0xe3, 0x88, 0x92, 0x58, 0xfa, 0x34, 0xd4, 0x9b, 0xa6, 0x66, 0x6d, 0x94, 0x82, 0xce, 0xd5, 0xeb,
	0x9e, 0x7a, 0xeb, 0x86, 0xf0, 0x02, 0xdc, 0xc3, 0x28, 0x8a, 0x42, 0x24, 0x91, 0xcf, 0x09, 0x66,
	0x3c, 0x14, 0x3a, 0x58, 0x35, 0xce, 0x5e, 0x9e, 0xe1, 0xa9, 0x04, 0xaf, 0x8d, 0x6b, 0xb6, 0x80,
	0xaf, 0x41, 0xe7, 0xb6, 0x9c, 0x7c, 0x96, 0xa9, 0x9e, 0x6d, 0xa5, 0xe7, 0xa0, 0xa6, 0x27, 0x1f,
	0x90, 0x1b, 0x42, 0x0c, 0x0e, 0xf0, 0x10, 0xd1, 0x38, 0xa2, 0xf1, 0xb7, 0xfa, 0x96, 0x5b, 0x4a,
	0xd6, 0xc9, 0x12, 0x59, 0x45, 0x5e, 0x65, 0xd5, 0x7b, 0xf8, 0xae, 0x33, 0xbd, 0x55, 0x7d, 0x48,
	0x85, 0x64, 0x9c, 0x62, 0x14, 0xe5, 0x55, 0x8a, 0xee, 0x77, 0x54, 0x99, 0xa3, 0x15, 0x5f, 0x43,
	0xde, 0xaa, 0x77, 0x30, 0xe7, 0x54, 0xfd, 0xb0, 0x0f, 0xda, 0x09, 0x67, 0x13, 0x1a, 0x12, 0x5e,
	0xe8, 0xdf, 0x55, 0xe0, 0xe3, 0x65, 0xe0, 0x2c, 0x21, 0x53, 0xbe, 0x9b, 0x54, 0x4d, 0x71, 0xe8,
	0x82, 0xdd, 0xfa, 0xdc, 0xe1, 0x63, 0xd0, 0x9c, 0x6f, 0x59, 0x53, 0x53, 0xdd, 0xc2, 0xc5, 0x62,
	0x3b, 0x60, 0xab, 0x58, 0x8b, 0xbe, 0x66, 0x6a, 0x56, 0xcb, 0x2b, 0xed, 0xee, 0xe5, 0xf5, 0xd4,
	0xd0, 0x6e, 0xa6, 0x86, 0xf6, 0x67, 0x6a, 0x68, 0x3f, 0x67, 0x46, 0xe3, 0x66, 0x66, 0x34, 0x7e,
	0xcd, 0x8c, 0xc6, 0xa7, 0xf3, 0x01, 0x95, 0xc3, 0x71, 0x60, 0x63, 0x36, 0x72, 0xdc, 0x42, 0xe7,
	0x39, 0x0a, 0x84, 0x53, 0xaa, 0x3e, 0xc1, 0x8c, 0x93, 0xaa, 0x99, 0x0e, 0xd8, 0x19, 0xb1, 0x70,
	0x1c, 0x11, 0x51, 0xfc, 0x64, 0xe5, 0x55, 0x42, 0x44, 0xb0, 0xa9, 0x7e, 0xa3, 0x2f, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0x26, 0x5b, 0x31, 0xc7, 0x05, 0x00, 0x00,
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
	if len(m.ProviderStates) > 0 {
		for iNdEx := len(m.ProviderStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProviderStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.HistoricalPriceRecords) > 0 {
		for iNdEx := len(m.HistoricalPriceRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HistoricalPriceRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.ChainlinkPriceStates) > 0 {
		for iNdEx := len(m.ChainlinkPriceStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainlinkPriceStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if m.BandIbcLatestRequestId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BandIbcLatestRequestId))
		i--
		dAtA[i] = 0x58
	}
	if len(m.CalldataRecords) > 0 {
		for iNdEx := len(m.CalldataRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CalldataRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.BandIbcLatestClientId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BandIbcLatestClientId))
		i--
		dAtA[i] = 0x48
	}
	{
		size, err := m.BandIbcParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.BandIbcOracleRequests) > 0 {
		for iNdEx := len(m.BandIbcOracleRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BandIbcOracleRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.BandIbcPriceStates) > 0 {
		for iNdEx := len(m.BandIbcPriceStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BandIbcPriceStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.CoinbasePriceStates) > 0 {
		for iNdEx := len(m.CoinbasePriceStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoinbasePriceStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PriceFeedPriceStates) > 0 {
		for iNdEx := len(m.PriceFeedPriceStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PriceFeedPriceStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.BandPriceStates) > 0 {
		for iNdEx := len(m.BandPriceStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BandPriceStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.BandRelayers) > 0 {
		for iNdEx := len(m.BandRelayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BandRelayers[iNdEx])
			copy(dAtA[i:], m.BandRelayers[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.BandRelayers[iNdEx])))
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

func (m *CalldataRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CalldataRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CalldataRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Calldata) > 0 {
		i -= len(m.Calldata)
		copy(dAtA[i:], m.Calldata)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Calldata)))
		i--
		dAtA[i] = 0x12
	}
	if m.ClientId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ClientId))
		i--
		dAtA[i] = 0x8
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
	if len(m.BandRelayers) > 0 {
		for _, s := range m.BandRelayers {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BandPriceStates) > 0 {
		for _, e := range m.BandPriceStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PriceFeedPriceStates) > 0 {
		for _, e := range m.PriceFeedPriceStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CoinbasePriceStates) > 0 {
		for _, e := range m.CoinbasePriceStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BandIbcPriceStates) > 0 {
		for _, e := range m.BandIbcPriceStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BandIbcOracleRequests) > 0 {
		for _, e := range m.BandIbcOracleRequests {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.BandIbcParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.BandIbcLatestClientId != 0 {
		n += 1 + sovGenesis(uint64(m.BandIbcLatestClientId))
	}
	if len(m.CalldataRecords) > 0 {
		for _, e := range m.CalldataRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.BandIbcLatestRequestId != 0 {
		n += 1 + sovGenesis(uint64(m.BandIbcLatestRequestId))
	}
	if len(m.ChainlinkPriceStates) > 0 {
		for _, e := range m.ChainlinkPriceStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.HistoricalPriceRecords) > 0 {
		for _, e := range m.HistoricalPriceRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProviderStates) > 0 {
		for _, e := range m.ProviderStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *CalldataRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClientId != 0 {
		n += 1 + sovGenesis(uint64(m.ClientId))
	}
	l = len(m.Calldata)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field BandRelayers", wireType)
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
			m.BandRelayers = append(m.BandRelayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandPriceStates", wireType)
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
			m.BandPriceStates = append(m.BandPriceStates, &BandPriceState{})
			if err := m.BandPriceStates[len(m.BandPriceStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceFeedPriceStates", wireType)
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
			m.PriceFeedPriceStates = append(m.PriceFeedPriceStates, &PriceFeedState{})
			if err := m.PriceFeedPriceStates[len(m.PriceFeedPriceStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinbasePriceStates", wireType)
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
			m.CoinbasePriceStates = append(m.CoinbasePriceStates, &CoinbasePriceState{})
			if err := m.CoinbasePriceStates[len(m.CoinbasePriceStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandIbcPriceStates", wireType)
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
			m.BandIbcPriceStates = append(m.BandIbcPriceStates, &BandPriceState{})
			if err := m.BandIbcPriceStates[len(m.BandIbcPriceStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandIbcOracleRequests", wireType)
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
			m.BandIbcOracleRequests = append(m.BandIbcOracleRequests, &BandOracleRequest{})
			if err := m.BandIbcOracleRequests[len(m.BandIbcOracleRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandIbcParams", wireType)
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
			if err := m.BandIbcParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandIbcLatestClientId", wireType)
			}
			m.BandIbcLatestClientId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BandIbcLatestClientId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CalldataRecords", wireType)
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
			m.CalldataRecords = append(m.CalldataRecords, &CalldataRecord{})
			if err := m.CalldataRecords[len(m.CalldataRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandIbcLatestRequestId", wireType)
			}
			m.BandIbcLatestRequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BandIbcLatestRequestId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainlinkPriceStates", wireType)
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
			m.ChainlinkPriceStates = append(m.ChainlinkPriceStates, &ChainlinkPriceState{})
			if err := m.ChainlinkPriceStates[len(m.ChainlinkPriceStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoricalPriceRecords", wireType)
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
			m.HistoricalPriceRecords = append(m.HistoricalPriceRecords, &PriceRecords{})
			if err := m.HistoricalPriceRecords[len(m.HistoricalPriceRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderStates", wireType)
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
			m.ProviderStates = append(m.ProviderStates, &ProviderState{})
			if err := m.ProviderStates[len(m.ProviderStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *CalldataRecord) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CalldataRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CalldataRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			m.ClientId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Calldata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Calldata = append(m.Calldata[:0], dAtA[iNdEx:postIndex]...)
			if m.Calldata == nil {
				m.Calldata = []byte{}
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
