// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package injective_spot_exchange_rpcpb

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

// InjectiveSpotExchangeRPCClient is the client API for InjectiveSpotExchangeRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InjectiveSpotExchangeRPCClient interface {
	// Get a list of Spot Markets
	Markets(ctx context.Context, in *MarketsRequest, opts ...grpc.CallOption) (*MarketsResponse, error)
	// Get details of a single spot market
	Market(ctx context.Context, in *MarketRequest, opts ...grpc.CallOption) (*MarketResponse, error)
	// Stream live updates of selected spot markets
	StreamMarkets(ctx context.Context, in *StreamMarketsRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamMarketsClient, error)
	// Orderbook of a Spot Market
	Orderbook(ctx context.Context, in *OrderbookRequest, opts ...grpc.CallOption) (*OrderbookResponse, error)
	// Stream live updates of selected spot market orderbook
	StreamOrderbook(ctx context.Context, in *StreamOrderbookRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrderbookClient, error)
	// Orders of a Spot Market
	Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	// Stream updates to individual orders of a Spot Market
	StreamOrders(ctx context.Context, in *StreamOrdersRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrdersClient, error)
	// Trades of a Spot Market
	Trades(ctx context.Context, in *TradesRequest, opts ...grpc.CallOption) (*TradesResponse, error)
	// Stream newly executed trades from Spot Market
	StreamTrades(ctx context.Context, in *StreamTradesRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamTradesClient, error)
	// List orders posted from this subaccount
	SubaccountOrdersList(ctx context.Context, in *SubaccountOrdersListRequest, opts ...grpc.CallOption) (*SubaccountOrdersListResponse, error)
	// List trades executed by this subaccount
	SubaccountTradesList(ctx context.Context, in *SubaccountTradesListRequest, opts ...grpc.CallOption) (*SubaccountTradesListResponse, error)
}

type injectiveSpotExchangeRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewInjectiveSpotExchangeRPCClient(cc grpc.ClientConnInterface) InjectiveSpotExchangeRPCClient {
	return &injectiveSpotExchangeRPCClient{cc}
}

func (c *injectiveSpotExchangeRPCClient) Markets(ctx context.Context, in *MarketsRequest, opts ...grpc.CallOption) (*MarketsResponse, error) {
	out := new(MarketsResponse)
	err := c.cc.Invoke(ctx, "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/Markets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveSpotExchangeRPCClient) Market(ctx context.Context, in *MarketRequest, opts ...grpc.CallOption) (*MarketResponse, error) {
	out := new(MarketResponse)
	err := c.cc.Invoke(ctx, "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/Market", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveSpotExchangeRPCClient) StreamMarkets(ctx context.Context, in *StreamMarketsRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamMarketsClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveSpotExchangeRPC_ServiceDesc.Streams[0], "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/StreamMarkets", opts...)
	if err != nil {
		return nil, err
	}
	x := &injectiveSpotExchangeRPCStreamMarketsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveSpotExchangeRPC_StreamMarketsClient interface {
	Recv() (*StreamMarketsResponse, error)
	grpc.ClientStream
}

type injectiveSpotExchangeRPCStreamMarketsClient struct {
	grpc.ClientStream
}

func (x *injectiveSpotExchangeRPCStreamMarketsClient) Recv() (*StreamMarketsResponse, error) {
	m := new(StreamMarketsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *injectiveSpotExchangeRPCClient) Orderbook(ctx context.Context, in *OrderbookRequest, opts ...grpc.CallOption) (*OrderbookResponse, error) {
	out := new(OrderbookResponse)
	err := c.cc.Invoke(ctx, "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/Orderbook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveSpotExchangeRPCClient) StreamOrderbook(ctx context.Context, in *StreamOrderbookRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrderbookClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveSpotExchangeRPC_ServiceDesc.Streams[1], "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/StreamOrderbook", opts...)
	if err != nil {
		return nil, err
	}
	x := &injectiveSpotExchangeRPCStreamOrderbookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveSpotExchangeRPC_StreamOrderbookClient interface {
	Recv() (*StreamOrderbookResponse, error)
	grpc.ClientStream
}

type injectiveSpotExchangeRPCStreamOrderbookClient struct {
	grpc.ClientStream
}

func (x *injectiveSpotExchangeRPCStreamOrderbookClient) Recv() (*StreamOrderbookResponse, error) {
	m := new(StreamOrderbookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *injectiveSpotExchangeRPCClient) Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/Orders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveSpotExchangeRPCClient) StreamOrders(ctx context.Context, in *StreamOrdersRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveSpotExchangeRPC_ServiceDesc.Streams[2], "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/StreamOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &injectiveSpotExchangeRPCStreamOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveSpotExchangeRPC_StreamOrdersClient interface {
	Recv() (*StreamOrdersResponse, error)
	grpc.ClientStream
}

type injectiveSpotExchangeRPCStreamOrdersClient struct {
	grpc.ClientStream
}

func (x *injectiveSpotExchangeRPCStreamOrdersClient) Recv() (*StreamOrdersResponse, error) {
	m := new(StreamOrdersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *injectiveSpotExchangeRPCClient) Trades(ctx context.Context, in *TradesRequest, opts ...grpc.CallOption) (*TradesResponse, error) {
	out := new(TradesResponse)
	err := c.cc.Invoke(ctx, "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/Trades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveSpotExchangeRPCClient) StreamTrades(ctx context.Context, in *StreamTradesRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamTradesClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveSpotExchangeRPC_ServiceDesc.Streams[3], "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/StreamTrades", opts...)
	if err != nil {
		return nil, err
	}
	x := &injectiveSpotExchangeRPCStreamTradesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveSpotExchangeRPC_StreamTradesClient interface {
	Recv() (*StreamTradesResponse, error)
	grpc.ClientStream
}

type injectiveSpotExchangeRPCStreamTradesClient struct {
	grpc.ClientStream
}

func (x *injectiveSpotExchangeRPCStreamTradesClient) Recv() (*StreamTradesResponse, error) {
	m := new(StreamTradesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *injectiveSpotExchangeRPCClient) SubaccountOrdersList(ctx context.Context, in *SubaccountOrdersListRequest, opts ...grpc.CallOption) (*SubaccountOrdersListResponse, error) {
	out := new(SubaccountOrdersListResponse)
	err := c.cc.Invoke(ctx, "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/SubaccountOrdersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveSpotExchangeRPCClient) SubaccountTradesList(ctx context.Context, in *SubaccountTradesListRequest, opts ...grpc.CallOption) (*SubaccountTradesListResponse, error) {
	out := new(SubaccountTradesListResponse)
	err := c.cc.Invoke(ctx, "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/SubaccountTradesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InjectiveSpotExchangeRPCServer is the server API for InjectiveSpotExchangeRPC service.
// All implementations must embed UnimplementedInjectiveSpotExchangeRPCServer
// for forward compatibility
type InjectiveSpotExchangeRPCServer interface {
	// Get a list of Spot Markets
	Markets(context.Context, *MarketsRequest) (*MarketsResponse, error)
	// Get details of a single spot market
	Market(context.Context, *MarketRequest) (*MarketResponse, error)
	// Stream live updates of selected spot markets
	StreamMarkets(*StreamMarketsRequest, InjectiveSpotExchangeRPC_StreamMarketsServer) error
	// Orderbook of a Spot Market
	Orderbook(context.Context, *OrderbookRequest) (*OrderbookResponse, error)
	// Stream live updates of selected spot market orderbook
	StreamOrderbook(*StreamOrderbookRequest, InjectiveSpotExchangeRPC_StreamOrderbookServer) error
	// Orders of a Spot Market
	Orders(context.Context, *OrdersRequest) (*OrdersResponse, error)
	// Stream updates to individual orders of a Spot Market
	StreamOrders(*StreamOrdersRequest, InjectiveSpotExchangeRPC_StreamOrdersServer) error
	// Trades of a Spot Market
	Trades(context.Context, *TradesRequest) (*TradesResponse, error)
	// Stream newly executed trades from Spot Market
	StreamTrades(*StreamTradesRequest, InjectiveSpotExchangeRPC_StreamTradesServer) error
	// List orders posted from this subaccount
	SubaccountOrdersList(context.Context, *SubaccountOrdersListRequest) (*SubaccountOrdersListResponse, error)
	// List trades executed by this subaccount
	SubaccountTradesList(context.Context, *SubaccountTradesListRequest) (*SubaccountTradesListResponse, error)
	mustEmbedUnimplementedInjectiveSpotExchangeRPCServer()
}

// UnimplementedInjectiveSpotExchangeRPCServer must be embedded to have forward compatible implementations.
type UnimplementedInjectiveSpotExchangeRPCServer struct {
}

func (UnimplementedInjectiveSpotExchangeRPCServer) Markets(context.Context, *MarketsRequest) (*MarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Markets not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) Market(context.Context, *MarketRequest) (*MarketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Market not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) StreamMarkets(*StreamMarketsRequest, InjectiveSpotExchangeRPC_StreamMarketsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMarkets not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) Orderbook(context.Context, *OrderbookRequest) (*OrderbookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Orderbook not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) StreamOrderbook(*StreamOrderbookRequest, InjectiveSpotExchangeRPC_StreamOrderbookServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrderbook not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) Orders(context.Context, *OrdersRequest) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Orders not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) StreamOrders(*StreamOrdersRequest, InjectiveSpotExchangeRPC_StreamOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrders not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) Trades(context.Context, *TradesRequest) (*TradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trades not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) StreamTrades(*StreamTradesRequest, InjectiveSpotExchangeRPC_StreamTradesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTrades not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) SubaccountOrdersList(context.Context, *SubaccountOrdersListRequest) (*SubaccountOrdersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountOrdersList not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) SubaccountTradesList(context.Context, *SubaccountTradesListRequest) (*SubaccountTradesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountTradesList not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) mustEmbedUnimplementedInjectiveSpotExchangeRPCServer() {
}

// UnsafeInjectiveSpotExchangeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InjectiveSpotExchangeRPCServer will
// result in compilation errors.
type UnsafeInjectiveSpotExchangeRPCServer interface {
	mustEmbedUnimplementedInjectiveSpotExchangeRPCServer()
}

func RegisterInjectiveSpotExchangeRPCServer(s grpc.ServiceRegistrar, srv InjectiveSpotExchangeRPCServer) {
	s.RegisterService(&InjectiveSpotExchangeRPC_ServiceDesc, srv)
}

func _InjectiveSpotExchangeRPC_Markets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).Markets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/Markets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).Markets(ctx, req.(*MarketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_Market_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).Market(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/Market",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).Market(ctx, req.(*MarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_StreamMarkets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamMarketsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveSpotExchangeRPCServer).StreamMarkets(m, &injectiveSpotExchangeRPCStreamMarketsServer{stream})
}

type InjectiveSpotExchangeRPC_StreamMarketsServer interface {
	Send(*StreamMarketsResponse) error
	grpc.ServerStream
}

type injectiveSpotExchangeRPCStreamMarketsServer struct {
	grpc.ServerStream
}

func (x *injectiveSpotExchangeRPCStreamMarketsServer) Send(m *StreamMarketsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveSpotExchangeRPC_Orderbook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderbookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).Orderbook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/Orderbook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).Orderbook(ctx, req.(*OrderbookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_StreamOrderbook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamOrderbookRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveSpotExchangeRPCServer).StreamOrderbook(m, &injectiveSpotExchangeRPCStreamOrderbookServer{stream})
}

type InjectiveSpotExchangeRPC_StreamOrderbookServer interface {
	Send(*StreamOrderbookResponse) error
	grpc.ServerStream
}

type injectiveSpotExchangeRPCStreamOrderbookServer struct {
	grpc.ServerStream
}

func (x *injectiveSpotExchangeRPCStreamOrderbookServer) Send(m *StreamOrderbookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveSpotExchangeRPC_Orders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).Orders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/Orders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).Orders(ctx, req.(*OrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_StreamOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamOrdersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveSpotExchangeRPCServer).StreamOrders(m, &injectiveSpotExchangeRPCStreamOrdersServer{stream})
}

type InjectiveSpotExchangeRPC_StreamOrdersServer interface {
	Send(*StreamOrdersResponse) error
	grpc.ServerStream
}

type injectiveSpotExchangeRPCStreamOrdersServer struct {
	grpc.ServerStream
}

func (x *injectiveSpotExchangeRPCStreamOrdersServer) Send(m *StreamOrdersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveSpotExchangeRPC_Trades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).Trades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/Trades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).Trades(ctx, req.(*TradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_StreamTrades_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamTradesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveSpotExchangeRPCServer).StreamTrades(m, &injectiveSpotExchangeRPCStreamTradesServer{stream})
}

type InjectiveSpotExchangeRPC_StreamTradesServer interface {
	Send(*StreamTradesResponse) error
	grpc.ServerStream
}

type injectiveSpotExchangeRPCStreamTradesServer struct {
	grpc.ServerStream
}

func (x *injectiveSpotExchangeRPCStreamTradesServer) Send(m *StreamTradesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveSpotExchangeRPC_SubaccountOrdersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubaccountOrdersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).SubaccountOrdersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/SubaccountOrdersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).SubaccountOrdersList(ctx, req.(*SubaccountOrdersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_SubaccountTradesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubaccountTradesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).SubaccountTradesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_spot_exchange_rpc.InjectiveSpotExchangeRPC/SubaccountTradesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).SubaccountTradesList(ctx, req.(*SubaccountTradesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InjectiveSpotExchangeRPC_ServiceDesc is the grpc.ServiceDesc for InjectiveSpotExchangeRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InjectiveSpotExchangeRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "injective_spot_exchange_rpc.InjectiveSpotExchangeRPC",
	HandlerType: (*InjectiveSpotExchangeRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Markets",
			Handler:    _InjectiveSpotExchangeRPC_Markets_Handler,
		},
		{
			MethodName: "Market",
			Handler:    _InjectiveSpotExchangeRPC_Market_Handler,
		},
		{
			MethodName: "Orderbook",
			Handler:    _InjectiveSpotExchangeRPC_Orderbook_Handler,
		},
		{
			MethodName: "Orders",
			Handler:    _InjectiveSpotExchangeRPC_Orders_Handler,
		},
		{
			MethodName: "Trades",
			Handler:    _InjectiveSpotExchangeRPC_Trades_Handler,
		},
		{
			MethodName: "SubaccountOrdersList",
			Handler:    _InjectiveSpotExchangeRPC_SubaccountOrdersList_Handler,
		},
		{
			MethodName: "SubaccountTradesList",
			Handler:    _InjectiveSpotExchangeRPC_SubaccountTradesList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMarkets",
			Handler:       _InjectiveSpotExchangeRPC_StreamMarkets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrderbook",
			Handler:       _InjectiveSpotExchangeRPC_StreamOrderbook_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrders",
			Handler:       _InjectiveSpotExchangeRPC_StreamOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamTrades",
			Handler:       _InjectiveSpotExchangeRPC_StreamTrades_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "injective_spot_exchange_rpc.proto",
}