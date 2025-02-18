// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kaiju_explorer_rpc.proto

package kaiju_explorer_rpcpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KaijuExplorerRPCClient is the client API for KaijuExplorerRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KaijuExplorerRPCClient interface {
	// GetAccountTxs returns tranctions involving in an account based upon params.
	GetAccountTxs(ctx context.Context, in *GetAccountTxsRequest, opts ...grpc.CallOption) (*GetAccountTxsResponse, error)
	// GetContractTxs returns contract-related transactions
	GetContractTxs(ctx context.Context, in *GetContractTxsRequest, opts ...grpc.CallOption) (*GetContractTxsResponse, error)
	// GetBlocks returns blocks based upon the request params
	GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error)
	// GetBlock returns block based upon the height or hash
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error)
	// GetValidators returns validators on the active chain
	GetValidators(ctx context.Context, in *GetValidatorsRequest, opts ...grpc.CallOption) (*GetValidatorsResponse, error)
	// GetValidator returns validator information on the active chain
	GetValidator(ctx context.Context, in *GetValidatorRequest, opts ...grpc.CallOption) (*GetValidatorResponse, error)
	// GetValidatorUptime returns validator uptime information on the active chain
	GetValidatorUptime(ctx context.Context, in *GetValidatorUptimeRequest, opts ...grpc.CallOption) (*GetValidatorUptimeResponse, error)
	// GetTxs returns transactions based upon the request params
	GetTxs(ctx context.Context, in *GetTxsRequest, opts ...grpc.CallOption) (*GetTxsResponse, error)
	// GetTxByTxHash returns certain transaction information by its tx hash.
	GetTxByTxHash(ctx context.Context, in *GetTxByTxHashRequest, opts ...grpc.CallOption) (*GetTxByTxHashResponse, error)
	// GetPeggyDepositTxs returns the peggy deposit transactions based upon the
	// request params
	GetPeggyDepositTxs(ctx context.Context, in *GetPeggyDepositTxsRequest, opts ...grpc.CallOption) (*GetPeggyDepositTxsResponse, error)
	// GetPeggyWithdrawalTxs returns the peggy withdrawal transactions based upon
	// the request params
	GetPeggyWithdrawalTxs(ctx context.Context, in *GetPeggyWithdrawalTxsRequest, opts ...grpc.CallOption) (*GetPeggyWithdrawalTxsResponse, error)
	// GetIBCTransferTxs returns the ibc transfer transactions based upon the
	// request params
	GetIBCTransferTxs(ctx context.Context, in *GetIBCTransferTxsRequest, opts ...grpc.CallOption) (*GetIBCTransferTxsResponse, error)
	// GetWasmCodes lists all stored code
	GetWasmCodes(ctx context.Context, in *GetWasmCodesRequest, opts ...grpc.CallOption) (*GetWasmCodesResponse, error)
	// GetWasmCodeById list cosmwasm code infor by ID
	GetWasmCodeByID(ctx context.Context, in *GetWasmCodeByIDRequest, opts ...grpc.CallOption) (*GetWasmCodeByIDResponse, error)
	// GetWasmContracts lists all contracts
	GetWasmContracts(ctx context.Context, in *GetWasmContractsRequest, opts ...grpc.CallOption) (*GetWasmContractsResponse, error)
	// GetWasmContractByAddress list cosmwasm contract infor by its address
	GetWasmContractByAddress(ctx context.Context, in *GetWasmContractByAddressRequest, opts ...grpc.CallOption) (*GetWasmContractByAddressResponse, error)
	// GetCw20Balance lists all cw20 balances of an kaiju account
	GetCw20Balance(ctx context.Context, in *GetCw20BalanceRequest, opts ...grpc.CallOption) (*GetCw20BalanceResponse, error)
	// Request relayers infos by marketIDs. If no ids are provided, all market with
	// associated relayers are returned
	Relayers(ctx context.Context, in *RelayersRequest, opts ...grpc.CallOption) (*RelayersResponse, error)
	// StreamTxs returns transactions based upon the request params
	StreamTxs(ctx context.Context, in *StreamTxsRequest, opts ...grpc.CallOption) (KaijuExplorerRPC_StreamTxsClient, error)
	// StreamBlocks returns the latest blocks
	StreamBlocks(ctx context.Context, in *StreamBlocksRequest, opts ...grpc.CallOption) (KaijuExplorerRPC_StreamBlocksClient, error)
}

type kaijuExplorerRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewKaijuExplorerRPCClient(cc grpc.ClientConnInterface) KaijuExplorerRPCClient {
	return &kaijuExplorerRPCClient{cc}
}

func (c *kaijuExplorerRPCClient) GetAccountTxs(ctx context.Context, in *GetAccountTxsRequest, opts ...grpc.CallOption) (*GetAccountTxsResponse, error) {
	out := new(GetAccountTxsResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetAccountTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetContractTxs(ctx context.Context, in *GetContractTxsRequest, opts ...grpc.CallOption) (*GetContractTxsResponse, error) {
	out := new(GetContractTxsResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetContractTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error) {
	out := new(GetBlocksResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error) {
	out := new(GetBlockResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetValidators(ctx context.Context, in *GetValidatorsRequest, opts ...grpc.CallOption) (*GetValidatorsResponse, error) {
	out := new(GetValidatorsResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetValidator(ctx context.Context, in *GetValidatorRequest, opts ...grpc.CallOption) (*GetValidatorResponse, error) {
	out := new(GetValidatorResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetValidatorUptime(ctx context.Context, in *GetValidatorUptimeRequest, opts ...grpc.CallOption) (*GetValidatorUptimeResponse, error) {
	out := new(GetValidatorUptimeResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetValidatorUptime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetTxs(ctx context.Context, in *GetTxsRequest, opts ...grpc.CallOption) (*GetTxsResponse, error) {
	out := new(GetTxsResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetTxByTxHash(ctx context.Context, in *GetTxByTxHashRequest, opts ...grpc.CallOption) (*GetTxByTxHashResponse, error) {
	out := new(GetTxByTxHashResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetTxByTxHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetPeggyDepositTxs(ctx context.Context, in *GetPeggyDepositTxsRequest, opts ...grpc.CallOption) (*GetPeggyDepositTxsResponse, error) {
	out := new(GetPeggyDepositTxsResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetPeggyDepositTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetPeggyWithdrawalTxs(ctx context.Context, in *GetPeggyWithdrawalTxsRequest, opts ...grpc.CallOption) (*GetPeggyWithdrawalTxsResponse, error) {
	out := new(GetPeggyWithdrawalTxsResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetPeggyWithdrawalTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetIBCTransferTxs(ctx context.Context, in *GetIBCTransferTxsRequest, opts ...grpc.CallOption) (*GetIBCTransferTxsResponse, error) {
	out := new(GetIBCTransferTxsResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetIBCTransferTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetWasmCodes(ctx context.Context, in *GetWasmCodesRequest, opts ...grpc.CallOption) (*GetWasmCodesResponse, error) {
	out := new(GetWasmCodesResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetWasmCodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetWasmCodeByID(ctx context.Context, in *GetWasmCodeByIDRequest, opts ...grpc.CallOption) (*GetWasmCodeByIDResponse, error) {
	out := new(GetWasmCodeByIDResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetWasmCodeByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetWasmContracts(ctx context.Context, in *GetWasmContractsRequest, opts ...grpc.CallOption) (*GetWasmContractsResponse, error) {
	out := new(GetWasmContractsResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetWasmContracts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetWasmContractByAddress(ctx context.Context, in *GetWasmContractByAddressRequest, opts ...grpc.CallOption) (*GetWasmContractByAddressResponse, error) {
	out := new(GetWasmContractByAddressResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetWasmContractByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) GetCw20Balance(ctx context.Context, in *GetCw20BalanceRequest, opts ...grpc.CallOption) (*GetCw20BalanceResponse, error) {
	out := new(GetCw20BalanceResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/GetCw20Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) Relayers(ctx context.Context, in *RelayersRequest, opts ...grpc.CallOption) (*RelayersResponse, error) {
	out := new(RelayersResponse)
	err := c.cc.Invoke(ctx, "/kaiju_explorer_rpc.KaijuExplorerRPC/Relayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuExplorerRPCClient) StreamTxs(ctx context.Context, in *StreamTxsRequest, opts ...grpc.CallOption) (KaijuExplorerRPC_StreamTxsClient, error) {
	stream, err := c.cc.NewStream(ctx, &KaijuExplorerRPC_ServiceDesc.Streams[0], "/kaiju_explorer_rpc.KaijuExplorerRPC/StreamTxs", opts...)
	if err != nil {
		return nil, err
	}
	x := &kaijuExplorerRPCStreamTxsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KaijuExplorerRPC_StreamTxsClient interface {
	Recv() (*StreamTxsResponse, error)
	grpc.ClientStream
}

type kaijuExplorerRPCStreamTxsClient struct {
	grpc.ClientStream
}

func (x *kaijuExplorerRPCStreamTxsClient) Recv() (*StreamTxsResponse, error) {
	m := new(StreamTxsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kaijuExplorerRPCClient) StreamBlocks(ctx context.Context, in *StreamBlocksRequest, opts ...grpc.CallOption) (KaijuExplorerRPC_StreamBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &KaijuExplorerRPC_ServiceDesc.Streams[1], "/kaiju_explorer_rpc.KaijuExplorerRPC/StreamBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &kaijuExplorerRPCStreamBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KaijuExplorerRPC_StreamBlocksClient interface {
	Recv() (*StreamBlocksResponse, error)
	grpc.ClientStream
}

type kaijuExplorerRPCStreamBlocksClient struct {
	grpc.ClientStream
}

func (x *kaijuExplorerRPCStreamBlocksClient) Recv() (*StreamBlocksResponse, error) {
	m := new(StreamBlocksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KaijuExplorerRPCServer is the server API for KaijuExplorerRPC service.
// All implementations must embed UnimplementedKaijuExplorerRPCServer
// for forward compatibility
type KaijuExplorerRPCServer interface {
	// GetAccountTxs returns tranctions involving in an account based upon params.
	GetAccountTxs(context.Context, *GetAccountTxsRequest) (*GetAccountTxsResponse, error)
	// GetContractTxs returns contract-related transactions
	GetContractTxs(context.Context, *GetContractTxsRequest) (*GetContractTxsResponse, error)
	// GetBlocks returns blocks based upon the request params
	GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error)
	// GetBlock returns block based upon the height or hash
	GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error)
	// GetValidators returns validators on the active chain
	GetValidators(context.Context, *GetValidatorsRequest) (*GetValidatorsResponse, error)
	// GetValidator returns validator information on the active chain
	GetValidator(context.Context, *GetValidatorRequest) (*GetValidatorResponse, error)
	// GetValidatorUptime returns validator uptime information on the active chain
	GetValidatorUptime(context.Context, *GetValidatorUptimeRequest) (*GetValidatorUptimeResponse, error)
	// GetTxs returns transactions based upon the request params
	GetTxs(context.Context, *GetTxsRequest) (*GetTxsResponse, error)
	// GetTxByTxHash returns certain transaction information by its tx hash.
	GetTxByTxHash(context.Context, *GetTxByTxHashRequest) (*GetTxByTxHashResponse, error)
	// GetPeggyDepositTxs returns the peggy deposit transactions based upon the
	// request params
	GetPeggyDepositTxs(context.Context, *GetPeggyDepositTxsRequest) (*GetPeggyDepositTxsResponse, error)
	// GetPeggyWithdrawalTxs returns the peggy withdrawal transactions based upon
	// the request params
	GetPeggyWithdrawalTxs(context.Context, *GetPeggyWithdrawalTxsRequest) (*GetPeggyWithdrawalTxsResponse, error)
	// GetIBCTransferTxs returns the ibc transfer transactions based upon the
	// request params
	GetIBCTransferTxs(context.Context, *GetIBCTransferTxsRequest) (*GetIBCTransferTxsResponse, error)
	// GetWasmCodes lists all stored code
	GetWasmCodes(context.Context, *GetWasmCodesRequest) (*GetWasmCodesResponse, error)
	// GetWasmCodeById list cosmwasm code infor by ID
	GetWasmCodeByID(context.Context, *GetWasmCodeByIDRequest) (*GetWasmCodeByIDResponse, error)
	// GetWasmContracts lists all contracts
	GetWasmContracts(context.Context, *GetWasmContractsRequest) (*GetWasmContractsResponse, error)
	// GetWasmContractByAddress list cosmwasm contract infor by its address
	GetWasmContractByAddress(context.Context, *GetWasmContractByAddressRequest) (*GetWasmContractByAddressResponse, error)
	// GetCw20Balance lists all cw20 balances of an kaiju account
	GetCw20Balance(context.Context, *GetCw20BalanceRequest) (*GetCw20BalanceResponse, error)
	// Request relayers infos by marketIDs. If no ids are provided, all market with
	// associated relayers are returned
	Relayers(context.Context, *RelayersRequest) (*RelayersResponse, error)
	// StreamTxs returns transactions based upon the request params
	StreamTxs(*StreamTxsRequest, KaijuExplorerRPC_StreamTxsServer) error
	// StreamBlocks returns the latest blocks
	StreamBlocks(*StreamBlocksRequest, KaijuExplorerRPC_StreamBlocksServer) error
	mustEmbedUnimplementedKaijuExplorerRPCServer()
}

// UnimplementedKaijuExplorerRPCServer must be embedded to have forward compatible implementations.
type UnimplementedKaijuExplorerRPCServer struct {
}

func (UnimplementedKaijuExplorerRPCServer) GetAccountTxs(context.Context, *GetAccountTxsRequest) (*GetAccountTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountTxs not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetContractTxs(context.Context, *GetContractTxsRequest) (*GetContractTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContractTxs not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetValidators(context.Context, *GetValidatorsRequest) (*GetValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidators not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetValidator(context.Context, *GetValidatorRequest) (*GetValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidator not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetValidatorUptime(context.Context, *GetValidatorUptimeRequest) (*GetValidatorUptimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorUptime not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetTxs(context.Context, *GetTxsRequest) (*GetTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxs not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetTxByTxHash(context.Context, *GetTxByTxHashRequest) (*GetTxByTxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByTxHash not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetPeggyDepositTxs(context.Context, *GetPeggyDepositTxsRequest) (*GetPeggyDepositTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeggyDepositTxs not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetPeggyWithdrawalTxs(context.Context, *GetPeggyWithdrawalTxsRequest) (*GetPeggyWithdrawalTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeggyWithdrawalTxs not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetIBCTransferTxs(context.Context, *GetIBCTransferTxsRequest) (*GetIBCTransferTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIBCTransferTxs not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetWasmCodes(context.Context, *GetWasmCodesRequest) (*GetWasmCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWasmCodes not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetWasmCodeByID(context.Context, *GetWasmCodeByIDRequest) (*GetWasmCodeByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWasmCodeByID not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetWasmContracts(context.Context, *GetWasmContractsRequest) (*GetWasmContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWasmContracts not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetWasmContractByAddress(context.Context, *GetWasmContractByAddressRequest) (*GetWasmContractByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWasmContractByAddress not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) GetCw20Balance(context.Context, *GetCw20BalanceRequest) (*GetCw20BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCw20Balance not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) Relayers(context.Context, *RelayersRequest) (*RelayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Relayers not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) StreamTxs(*StreamTxsRequest, KaijuExplorerRPC_StreamTxsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTxs not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) StreamBlocks(*StreamBlocksRequest, KaijuExplorerRPC_StreamBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamBlocks not implemented")
}
func (UnimplementedKaijuExplorerRPCServer) mustEmbedUnimplementedKaijuExplorerRPCServer() {}

// UnsafeKaijuExplorerRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KaijuExplorerRPCServer will
// result in compilation errors.
type UnsafeKaijuExplorerRPCServer interface {
	mustEmbedUnimplementedKaijuExplorerRPCServer()
}

func RegisterKaijuExplorerRPCServer(s grpc.ServiceRegistrar, srv KaijuExplorerRPCServer) {
	s.RegisterService(&KaijuExplorerRPC_ServiceDesc, srv)
}

func _KaijuExplorerRPC_GetAccountTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetAccountTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetAccountTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetAccountTxs(ctx, req.(*GetAccountTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetContractTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetContractTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetContractTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetContractTxs(ctx, req.(*GetContractTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetBlocks(ctx, req.(*GetBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetValidators(ctx, req.(*GetValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetValidator(ctx, req.(*GetValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetValidatorUptime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorUptimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetValidatorUptime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetValidatorUptime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetValidatorUptime(ctx, req.(*GetValidatorUptimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetTxs(ctx, req.(*GetTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetTxByTxHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxByTxHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetTxByTxHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetTxByTxHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetTxByTxHash(ctx, req.(*GetTxByTxHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetPeggyDepositTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeggyDepositTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetPeggyDepositTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetPeggyDepositTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetPeggyDepositTxs(ctx, req.(*GetPeggyDepositTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetPeggyWithdrawalTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeggyWithdrawalTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetPeggyWithdrawalTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetPeggyWithdrawalTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetPeggyWithdrawalTxs(ctx, req.(*GetPeggyWithdrawalTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetIBCTransferTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIBCTransferTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetIBCTransferTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetIBCTransferTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetIBCTransferTxs(ctx, req.(*GetIBCTransferTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetWasmCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWasmCodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetWasmCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetWasmCodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetWasmCodes(ctx, req.(*GetWasmCodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetWasmCodeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWasmCodeByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetWasmCodeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetWasmCodeByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetWasmCodeByID(ctx, req.(*GetWasmCodeByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetWasmContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWasmContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetWasmContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetWasmContracts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetWasmContracts(ctx, req.(*GetWasmContractsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetWasmContractByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWasmContractByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetWasmContractByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetWasmContractByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetWasmContractByAddress(ctx, req.(*GetWasmContractByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_GetCw20Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCw20BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).GetCw20Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/GetCw20Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).GetCw20Balance(ctx, req.(*GetCw20BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_Relayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuExplorerRPCServer).Relayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_explorer_rpc.KaijuExplorerRPC/Relayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuExplorerRPCServer).Relayers(ctx, req.(*RelayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuExplorerRPC_StreamTxs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamTxsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KaijuExplorerRPCServer).StreamTxs(m, &kaijuExplorerRPCStreamTxsServer{stream})
}

type KaijuExplorerRPC_StreamTxsServer interface {
	Send(*StreamTxsResponse) error
	grpc.ServerStream
}

type kaijuExplorerRPCStreamTxsServer struct {
	grpc.ServerStream
}

func (x *kaijuExplorerRPCStreamTxsServer) Send(m *StreamTxsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KaijuExplorerRPC_StreamBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamBlocksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KaijuExplorerRPCServer).StreamBlocks(m, &kaijuExplorerRPCStreamBlocksServer{stream})
}

type KaijuExplorerRPC_StreamBlocksServer interface {
	Send(*StreamBlocksResponse) error
	grpc.ServerStream
}

type kaijuExplorerRPCStreamBlocksServer struct {
	grpc.ServerStream
}

func (x *kaijuExplorerRPCStreamBlocksServer) Send(m *StreamBlocksResponse) error {
	return x.ServerStream.SendMsg(m)
}

// KaijuExplorerRPC_ServiceDesc is the grpc.ServiceDesc for KaijuExplorerRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KaijuExplorerRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kaiju_explorer_rpc.KaijuExplorerRPC",
	HandlerType: (*KaijuExplorerRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountTxs",
			Handler:    _KaijuExplorerRPC_GetAccountTxs_Handler,
		},
		{
			MethodName: "GetContractTxs",
			Handler:    _KaijuExplorerRPC_GetContractTxs_Handler,
		},
		{
			MethodName: "GetBlocks",
			Handler:    _KaijuExplorerRPC_GetBlocks_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _KaijuExplorerRPC_GetBlock_Handler,
		},
		{
			MethodName: "GetValidators",
			Handler:    _KaijuExplorerRPC_GetValidators_Handler,
		},
		{
			MethodName: "GetValidator",
			Handler:    _KaijuExplorerRPC_GetValidator_Handler,
		},
		{
			MethodName: "GetValidatorUptime",
			Handler:    _KaijuExplorerRPC_GetValidatorUptime_Handler,
		},
		{
			MethodName: "GetTxs",
			Handler:    _KaijuExplorerRPC_GetTxs_Handler,
		},
		{
			MethodName: "GetTxByTxHash",
			Handler:    _KaijuExplorerRPC_GetTxByTxHash_Handler,
		},
		{
			MethodName: "GetPeggyDepositTxs",
			Handler:    _KaijuExplorerRPC_GetPeggyDepositTxs_Handler,
		},
		{
			MethodName: "GetPeggyWithdrawalTxs",
			Handler:    _KaijuExplorerRPC_GetPeggyWithdrawalTxs_Handler,
		},
		{
			MethodName: "GetIBCTransferTxs",
			Handler:    _KaijuExplorerRPC_GetIBCTransferTxs_Handler,
		},
		{
			MethodName: "GetWasmCodes",
			Handler:    _KaijuExplorerRPC_GetWasmCodes_Handler,
		},
		{
			MethodName: "GetWasmCodeByID",
			Handler:    _KaijuExplorerRPC_GetWasmCodeByID_Handler,
		},
		{
			MethodName: "GetWasmContracts",
			Handler:    _KaijuExplorerRPC_GetWasmContracts_Handler,
		},
		{
			MethodName: "GetWasmContractByAddress",
			Handler:    _KaijuExplorerRPC_GetWasmContractByAddress_Handler,
		},
		{
			MethodName: "GetCw20Balance",
			Handler:    _KaijuExplorerRPC_GetCw20Balance_Handler,
		},
		{
			MethodName: "Relayers",
			Handler:    _KaijuExplorerRPC_Relayers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTxs",
			Handler:       _KaijuExplorerRPC_StreamTxs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamBlocks",
			Handler:       _KaijuExplorerRPC_StreamBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kaiju_explorer_rpc.proto",
}
