// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kaiju_spot_exchange_rpc.proto

package kaiju_spot_exchange_rpcpb

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

// KaijuSpotExchangeRPCClient is the client API for KaijuSpotExchangeRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KaijuSpotExchangeRPCClient interface {
	// Get a list of Spot Markets
	Markets(ctx context.Context, in *MarketsRequest, opts ...grpc.CallOption) (*MarketsResponse, error)
	// Get details of a single spot market
	Market(ctx context.Context, in *MarketRequest, opts ...grpc.CallOption) (*MarketResponse, error)
	// Stream live updates of selected spot markets
	StreamMarkets(ctx context.Context, in *StreamMarketsRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamMarketsClient, error)
	// Orderbook of a Spot Market
	Orderbook(ctx context.Context, in *OrderbookRequest, opts ...grpc.CallOption) (*OrderbookResponse, error)
	// Orderbook of Spot Markets
	Orderbooks(ctx context.Context, in *OrderbooksRequest, opts ...grpc.CallOption) (*OrderbooksResponse, error)
	// Stream live snapshot updates of selected spot market orderbook
	StreamOrderbookSnapshot(ctx context.Context, in *StreamOrderbookSnapshotRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamOrderbookSnapshotClient, error)
	// Stream live level updates of selected spot market orderbook
	StreamOrderbookUpdate(ctx context.Context, in *StreamOrderbookUpdateRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamOrderbookUpdateClient, error)
	// Orders of a Spot Market
	Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	// Stream updates to individual orders of a Spot Market
	StreamOrders(ctx context.Context, in *StreamOrdersRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamOrdersClient, error)
	// Trades of a Spot Market
	Trades(ctx context.Context, in *TradesRequest, opts ...grpc.CallOption) (*TradesResponse, error)
	// Stream newly executed trades from Spot Market
	StreamTrades(ctx context.Context, in *StreamTradesRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamTradesClient, error)
	// List orders posted from this subaccount
	SubaccountOrdersList(ctx context.Context, in *SubaccountOrdersListRequest, opts ...grpc.CallOption) (*SubaccountOrdersListResponse, error)
	// List trades executed by this subaccount
	SubaccountTradesList(ctx context.Context, in *SubaccountTradesListRequest, opts ...grpc.CallOption) (*SubaccountTradesListResponse, error)
	// Lists history orders posted from this subaccount
	OrdersHistory(ctx context.Context, in *OrdersHistoryRequest, opts ...grpc.CallOption) (*OrdersHistoryResponse, error)
	// Stream updates to historical orders of a spot Market
	StreamOrdersHistory(ctx context.Context, in *StreamOrdersHistoryRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamOrdersHistoryClient, error)
}

type kaijuSpotExchangeRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewKaijuSpotExchangeRPCClient(cc grpc.ClientConnInterface) KaijuSpotExchangeRPCClient {
	return &kaijuSpotExchangeRPCClient{cc}
}

func (c *kaijuSpotExchangeRPCClient) Markets(ctx context.Context, in *MarketsRequest, opts ...grpc.CallOption) (*MarketsResponse, error) {
	out := new(MarketsResponse)
	err := c.cc.Invoke(ctx, "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Markets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuSpotExchangeRPCClient) Market(ctx context.Context, in *MarketRequest, opts ...grpc.CallOption) (*MarketResponse, error) {
	out := new(MarketResponse)
	err := c.cc.Invoke(ctx, "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Market", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuSpotExchangeRPCClient) StreamMarkets(ctx context.Context, in *StreamMarketsRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamMarketsClient, error) {
	stream, err := c.cc.NewStream(ctx, &KaijuSpotExchangeRPC_ServiceDesc.Streams[0], "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/StreamMarkets", opts...)
	if err != nil {
		return nil, err
	}
	x := &kaijuSpotExchangeRPCStreamMarketsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KaijuSpotExchangeRPC_StreamMarketsClient interface {
	Recv() (*StreamMarketsResponse, error)
	grpc.ClientStream
}

type kaijuSpotExchangeRPCStreamMarketsClient struct {
	grpc.ClientStream
}

func (x *kaijuSpotExchangeRPCStreamMarketsClient) Recv() (*StreamMarketsResponse, error) {
	m := new(StreamMarketsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kaijuSpotExchangeRPCClient) Orderbook(ctx context.Context, in *OrderbookRequest, opts ...grpc.CallOption) (*OrderbookResponse, error) {
	out := new(OrderbookResponse)
	err := c.cc.Invoke(ctx, "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Orderbook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuSpotExchangeRPCClient) Orderbooks(ctx context.Context, in *OrderbooksRequest, opts ...grpc.CallOption) (*OrderbooksResponse, error) {
	out := new(OrderbooksResponse)
	err := c.cc.Invoke(ctx, "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Orderbooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuSpotExchangeRPCClient) StreamOrderbookSnapshot(ctx context.Context, in *StreamOrderbookSnapshotRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamOrderbookSnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &KaijuSpotExchangeRPC_ServiceDesc.Streams[1], "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/StreamOrderbookSnapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &kaijuSpotExchangeRPCStreamOrderbookSnapshotClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KaijuSpotExchangeRPC_StreamOrderbookSnapshotClient interface {
	Recv() (*StreamOrderbookSnapshotResponse, error)
	grpc.ClientStream
}

type kaijuSpotExchangeRPCStreamOrderbookSnapshotClient struct {
	grpc.ClientStream
}

func (x *kaijuSpotExchangeRPCStreamOrderbookSnapshotClient) Recv() (*StreamOrderbookSnapshotResponse, error) {
	m := new(StreamOrderbookSnapshotResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kaijuSpotExchangeRPCClient) StreamOrderbookUpdate(ctx context.Context, in *StreamOrderbookUpdateRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamOrderbookUpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &KaijuSpotExchangeRPC_ServiceDesc.Streams[2], "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/StreamOrderbookUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &kaijuSpotExchangeRPCStreamOrderbookUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KaijuSpotExchangeRPC_StreamOrderbookUpdateClient interface {
	Recv() (*StreamOrderbookUpdateResponse, error)
	grpc.ClientStream
}

type kaijuSpotExchangeRPCStreamOrderbookUpdateClient struct {
	grpc.ClientStream
}

func (x *kaijuSpotExchangeRPCStreamOrderbookUpdateClient) Recv() (*StreamOrderbookUpdateResponse, error) {
	m := new(StreamOrderbookUpdateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kaijuSpotExchangeRPCClient) Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Orders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuSpotExchangeRPCClient) StreamOrders(ctx context.Context, in *StreamOrdersRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &KaijuSpotExchangeRPC_ServiceDesc.Streams[3], "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/StreamOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &kaijuSpotExchangeRPCStreamOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KaijuSpotExchangeRPC_StreamOrdersClient interface {
	Recv() (*StreamOrdersResponse, error)
	grpc.ClientStream
}

type kaijuSpotExchangeRPCStreamOrdersClient struct {
	grpc.ClientStream
}

func (x *kaijuSpotExchangeRPCStreamOrdersClient) Recv() (*StreamOrdersResponse, error) {
	m := new(StreamOrdersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kaijuSpotExchangeRPCClient) Trades(ctx context.Context, in *TradesRequest, opts ...grpc.CallOption) (*TradesResponse, error) {
	out := new(TradesResponse)
	err := c.cc.Invoke(ctx, "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Trades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuSpotExchangeRPCClient) StreamTrades(ctx context.Context, in *StreamTradesRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamTradesClient, error) {
	stream, err := c.cc.NewStream(ctx, &KaijuSpotExchangeRPC_ServiceDesc.Streams[4], "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/StreamTrades", opts...)
	if err != nil {
		return nil, err
	}
	x := &kaijuSpotExchangeRPCStreamTradesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KaijuSpotExchangeRPC_StreamTradesClient interface {
	Recv() (*StreamTradesResponse, error)
	grpc.ClientStream
}

type kaijuSpotExchangeRPCStreamTradesClient struct {
	grpc.ClientStream
}

func (x *kaijuSpotExchangeRPCStreamTradesClient) Recv() (*StreamTradesResponse, error) {
	m := new(StreamTradesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kaijuSpotExchangeRPCClient) SubaccountOrdersList(ctx context.Context, in *SubaccountOrdersListRequest, opts ...grpc.CallOption) (*SubaccountOrdersListResponse, error) {
	out := new(SubaccountOrdersListResponse)
	err := c.cc.Invoke(ctx, "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/SubaccountOrdersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuSpotExchangeRPCClient) SubaccountTradesList(ctx context.Context, in *SubaccountTradesListRequest, opts ...grpc.CallOption) (*SubaccountTradesListResponse, error) {
	out := new(SubaccountTradesListResponse)
	err := c.cc.Invoke(ctx, "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/SubaccountTradesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuSpotExchangeRPCClient) OrdersHistory(ctx context.Context, in *OrdersHistoryRequest, opts ...grpc.CallOption) (*OrdersHistoryResponse, error) {
	out := new(OrdersHistoryResponse)
	err := c.cc.Invoke(ctx, "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/OrdersHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuSpotExchangeRPCClient) StreamOrdersHistory(ctx context.Context, in *StreamOrdersHistoryRequest, opts ...grpc.CallOption) (KaijuSpotExchangeRPC_StreamOrdersHistoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &KaijuSpotExchangeRPC_ServiceDesc.Streams[5], "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/StreamOrdersHistory", opts...)
	if err != nil {
		return nil, err
	}
	x := &kaijuSpotExchangeRPCStreamOrdersHistoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KaijuSpotExchangeRPC_StreamOrdersHistoryClient interface {
	Recv() (*StreamOrdersHistoryResponse, error)
	grpc.ClientStream
}

type kaijuSpotExchangeRPCStreamOrdersHistoryClient struct {
	grpc.ClientStream
}

func (x *kaijuSpotExchangeRPCStreamOrdersHistoryClient) Recv() (*StreamOrdersHistoryResponse, error) {
	m := new(StreamOrdersHistoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KaijuSpotExchangeRPCServer is the server API for KaijuSpotExchangeRPC service.
// All implementations must embed UnimplementedKaijuSpotExchangeRPCServer
// for forward compatibility
type KaijuSpotExchangeRPCServer interface {
	// Get a list of Spot Markets
	Markets(context.Context, *MarketsRequest) (*MarketsResponse, error)
	// Get details of a single spot market
	Market(context.Context, *MarketRequest) (*MarketResponse, error)
	// Stream live updates of selected spot markets
	StreamMarkets(*StreamMarketsRequest, KaijuSpotExchangeRPC_StreamMarketsServer) error
	// Orderbook of a Spot Market
	Orderbook(context.Context, *OrderbookRequest) (*OrderbookResponse, error)
	// Orderbook of Spot Markets
	Orderbooks(context.Context, *OrderbooksRequest) (*OrderbooksResponse, error)
	// Stream live snapshot updates of selected spot market orderbook
	StreamOrderbookSnapshot(*StreamOrderbookSnapshotRequest, KaijuSpotExchangeRPC_StreamOrderbookSnapshotServer) error
	// Stream live level updates of selected spot market orderbook
	StreamOrderbookUpdate(*StreamOrderbookUpdateRequest, KaijuSpotExchangeRPC_StreamOrderbookUpdateServer) error
	// Orders of a Spot Market
	Orders(context.Context, *OrdersRequest) (*OrdersResponse, error)
	// Stream updates to individual orders of a Spot Market
	StreamOrders(*StreamOrdersRequest, KaijuSpotExchangeRPC_StreamOrdersServer) error
	// Trades of a Spot Market
	Trades(context.Context, *TradesRequest) (*TradesResponse, error)
	// Stream newly executed trades from Spot Market
	StreamTrades(*StreamTradesRequest, KaijuSpotExchangeRPC_StreamTradesServer) error
	// List orders posted from this subaccount
	SubaccountOrdersList(context.Context, *SubaccountOrdersListRequest) (*SubaccountOrdersListResponse, error)
	// List trades executed by this subaccount
	SubaccountTradesList(context.Context, *SubaccountTradesListRequest) (*SubaccountTradesListResponse, error)
	// Lists history orders posted from this subaccount
	OrdersHistory(context.Context, *OrdersHistoryRequest) (*OrdersHistoryResponse, error)
	// Stream updates to historical orders of a spot Market
	StreamOrdersHistory(*StreamOrdersHistoryRequest, KaijuSpotExchangeRPC_StreamOrdersHistoryServer) error
	mustEmbedUnimplementedKaijuSpotExchangeRPCServer()
}

// UnimplementedKaijuSpotExchangeRPCServer must be embedded to have forward compatible implementations.
type UnimplementedKaijuSpotExchangeRPCServer struct {
}

func (UnimplementedKaijuSpotExchangeRPCServer) Markets(context.Context, *MarketsRequest) (*MarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Markets not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) Market(context.Context, *MarketRequest) (*MarketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Market not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) StreamMarkets(*StreamMarketsRequest, KaijuSpotExchangeRPC_StreamMarketsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMarkets not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) Orderbook(context.Context, *OrderbookRequest) (*OrderbookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Orderbook not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) Orderbooks(context.Context, *OrderbooksRequest) (*OrderbooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Orderbooks not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) StreamOrderbookSnapshot(*StreamOrderbookSnapshotRequest, KaijuSpotExchangeRPC_StreamOrderbookSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrderbookSnapshot not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) StreamOrderbookUpdate(*StreamOrderbookUpdateRequest, KaijuSpotExchangeRPC_StreamOrderbookUpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrderbookUpdate not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) Orders(context.Context, *OrdersRequest) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Orders not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) StreamOrders(*StreamOrdersRequest, KaijuSpotExchangeRPC_StreamOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrders not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) Trades(context.Context, *TradesRequest) (*TradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trades not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) StreamTrades(*StreamTradesRequest, KaijuSpotExchangeRPC_StreamTradesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTrades not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) SubaccountOrdersList(context.Context, *SubaccountOrdersListRequest) (*SubaccountOrdersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountOrdersList not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) SubaccountTradesList(context.Context, *SubaccountTradesListRequest) (*SubaccountTradesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountTradesList not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) OrdersHistory(context.Context, *OrdersHistoryRequest) (*OrdersHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrdersHistory not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) StreamOrdersHistory(*StreamOrdersHistoryRequest, KaijuSpotExchangeRPC_StreamOrdersHistoryServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrdersHistory not implemented")
}
func (UnimplementedKaijuSpotExchangeRPCServer) mustEmbedUnimplementedKaijuSpotExchangeRPCServer() {
}

// UnsafeKaijuSpotExchangeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KaijuSpotExchangeRPCServer will
// result in compilation errors.
type UnsafeKaijuSpotExchangeRPCServer interface {
	mustEmbedUnimplementedKaijuSpotExchangeRPCServer()
}

func RegisterKaijuSpotExchangeRPCServer(s grpc.ServiceRegistrar, srv KaijuSpotExchangeRPCServer) {
	s.RegisterService(&KaijuSpotExchangeRPC_ServiceDesc, srv)
}

func _KaijuSpotExchangeRPC_Markets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuSpotExchangeRPCServer).Markets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Markets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuSpotExchangeRPCServer).Markets(ctx, req.(*MarketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuSpotExchangeRPC_Market_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuSpotExchangeRPCServer).Market(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Market",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuSpotExchangeRPCServer).Market(ctx, req.(*MarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuSpotExchangeRPC_StreamMarkets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamMarketsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KaijuSpotExchangeRPCServer).StreamMarkets(m, &kaijuSpotExchangeRPCStreamMarketsServer{stream})
}

type KaijuSpotExchangeRPC_StreamMarketsServer interface {
	Send(*StreamMarketsResponse) error
	grpc.ServerStream
}

type kaijuSpotExchangeRPCStreamMarketsServer struct {
	grpc.ServerStream
}

func (x *kaijuSpotExchangeRPCStreamMarketsServer) Send(m *StreamMarketsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KaijuSpotExchangeRPC_Orderbook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderbookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuSpotExchangeRPCServer).Orderbook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Orderbook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuSpotExchangeRPCServer).Orderbook(ctx, req.(*OrderbookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuSpotExchangeRPC_Orderbooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderbooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuSpotExchangeRPCServer).Orderbooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Orderbooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuSpotExchangeRPCServer).Orderbooks(ctx, req.(*OrderbooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuSpotExchangeRPC_StreamOrderbookSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamOrderbookSnapshotRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KaijuSpotExchangeRPCServer).StreamOrderbookSnapshot(m, &kaijuSpotExchangeRPCStreamOrderbookSnapshotServer{stream})
}

type KaijuSpotExchangeRPC_StreamOrderbookSnapshotServer interface {
	Send(*StreamOrderbookSnapshotResponse) error
	grpc.ServerStream
}

type kaijuSpotExchangeRPCStreamOrderbookSnapshotServer struct {
	grpc.ServerStream
}

func (x *kaijuSpotExchangeRPCStreamOrderbookSnapshotServer) Send(m *StreamOrderbookSnapshotResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KaijuSpotExchangeRPC_StreamOrderbookUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamOrderbookUpdateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KaijuSpotExchangeRPCServer).StreamOrderbookUpdate(m, &kaijuSpotExchangeRPCStreamOrderbookUpdateServer{stream})
}

type KaijuSpotExchangeRPC_StreamOrderbookUpdateServer interface {
	Send(*StreamOrderbookUpdateResponse) error
	grpc.ServerStream
}

type kaijuSpotExchangeRPCStreamOrderbookUpdateServer struct {
	grpc.ServerStream
}

func (x *kaijuSpotExchangeRPCStreamOrderbookUpdateServer) Send(m *StreamOrderbookUpdateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KaijuSpotExchangeRPC_Orders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuSpotExchangeRPCServer).Orders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Orders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuSpotExchangeRPCServer).Orders(ctx, req.(*OrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuSpotExchangeRPC_StreamOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamOrdersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KaijuSpotExchangeRPCServer).StreamOrders(m, &kaijuSpotExchangeRPCStreamOrdersServer{stream})
}

type KaijuSpotExchangeRPC_StreamOrdersServer interface {
	Send(*StreamOrdersResponse) error
	grpc.ServerStream
}

type kaijuSpotExchangeRPCStreamOrdersServer struct {
	grpc.ServerStream
}

func (x *kaijuSpotExchangeRPCStreamOrdersServer) Send(m *StreamOrdersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KaijuSpotExchangeRPC_Trades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuSpotExchangeRPCServer).Trades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/Trades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuSpotExchangeRPCServer).Trades(ctx, req.(*TradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuSpotExchangeRPC_StreamTrades_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamTradesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KaijuSpotExchangeRPCServer).StreamTrades(m, &kaijuSpotExchangeRPCStreamTradesServer{stream})
}

type KaijuSpotExchangeRPC_StreamTradesServer interface {
	Send(*StreamTradesResponse) error
	grpc.ServerStream
}

type kaijuSpotExchangeRPCStreamTradesServer struct {
	grpc.ServerStream
}

func (x *kaijuSpotExchangeRPCStreamTradesServer) Send(m *StreamTradesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KaijuSpotExchangeRPC_SubaccountOrdersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubaccountOrdersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuSpotExchangeRPCServer).SubaccountOrdersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/SubaccountOrdersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuSpotExchangeRPCServer).SubaccountOrdersList(ctx, req.(*SubaccountOrdersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuSpotExchangeRPC_SubaccountTradesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubaccountTradesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuSpotExchangeRPCServer).SubaccountTradesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/SubaccountTradesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuSpotExchangeRPCServer).SubaccountTradesList(ctx, req.(*SubaccountTradesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuSpotExchangeRPC_OrdersHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuSpotExchangeRPCServer).OrdersHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC/OrdersHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuSpotExchangeRPCServer).OrdersHistory(ctx, req.(*OrdersHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuSpotExchangeRPC_StreamOrdersHistory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamOrdersHistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KaijuSpotExchangeRPCServer).StreamOrdersHistory(m, &kaijuSpotExchangeRPCStreamOrdersHistoryServer{stream})
}

type KaijuSpotExchangeRPC_StreamOrdersHistoryServer interface {
	Send(*StreamOrdersHistoryResponse) error
	grpc.ServerStream
}

type kaijuSpotExchangeRPCStreamOrdersHistoryServer struct {
	grpc.ServerStream
}

func (x *kaijuSpotExchangeRPCStreamOrdersHistoryServer) Send(m *StreamOrdersHistoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// KaijuSpotExchangeRPC_ServiceDesc is the grpc.ServiceDesc for KaijuSpotExchangeRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KaijuSpotExchangeRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kaiju_spot_exchange_rpc.KaijuSpotExchangeRPC",
	HandlerType: (*KaijuSpotExchangeRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Markets",
			Handler:    _KaijuSpotExchangeRPC_Markets_Handler,
		},
		{
			MethodName: "Market",
			Handler:    _KaijuSpotExchangeRPC_Market_Handler,
		},
		{
			MethodName: "Orderbook",
			Handler:    _KaijuSpotExchangeRPC_Orderbook_Handler,
		},
		{
			MethodName: "Orderbooks",
			Handler:    _KaijuSpotExchangeRPC_Orderbooks_Handler,
		},
		{
			MethodName: "Orders",
			Handler:    _KaijuSpotExchangeRPC_Orders_Handler,
		},
		{
			MethodName: "Trades",
			Handler:    _KaijuSpotExchangeRPC_Trades_Handler,
		},
		{
			MethodName: "SubaccountOrdersList",
			Handler:    _KaijuSpotExchangeRPC_SubaccountOrdersList_Handler,
		},
		{
			MethodName: "SubaccountTradesList",
			Handler:    _KaijuSpotExchangeRPC_SubaccountTradesList_Handler,
		},
		{
			MethodName: "OrdersHistory",
			Handler:    _KaijuSpotExchangeRPC_OrdersHistory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMarkets",
			Handler:       _KaijuSpotExchangeRPC_StreamMarkets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrderbookSnapshot",
			Handler:       _KaijuSpotExchangeRPC_StreamOrderbookSnapshot_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrderbookUpdate",
			Handler:       _KaijuSpotExchangeRPC_StreamOrderbookUpdate_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrders",
			Handler:       _KaijuSpotExchangeRPC_StreamOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamTrades",
			Handler:       _KaijuSpotExchangeRPC_StreamTrades_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrdersHistory",
			Handler:       _KaijuSpotExchangeRPC_StreamOrdersHistory_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kaiju_spot_exchange_rpc.proto",
}
