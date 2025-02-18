// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: kaiju_accounts_rpc.proto

/*
Package kaiju_accounts_rpcpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package kaiju_accounts_rpcpb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_KaijuAccountsRPC_Portfolio_0(ctx context.Context, marshaler runtime.Marshaler, client KaijuAccountsRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PortfolioRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Portfolio(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_KaijuAccountsRPC_Portfolio_0(ctx context.Context, marshaler runtime.Marshaler, server KaijuAccountsRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PortfolioRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Portfolio(ctx, &protoReq)
	return msg, metadata, err

}

func request_KaijuAccountsRPC_OrderStates_0(ctx context.Context, marshaler runtime.Marshaler, client KaijuAccountsRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq OrderStatesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.OrderStates(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_KaijuAccountsRPC_OrderStates_0(ctx context.Context, marshaler runtime.Marshaler, server KaijuAccountsRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq OrderStatesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.OrderStates(ctx, &protoReq)
	return msg, metadata, err

}

func request_KaijuAccountsRPC_SubaccountsList_0(ctx context.Context, marshaler runtime.Marshaler, client KaijuAccountsRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SubaccountsListRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.SubaccountsList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_KaijuAccountsRPC_SubaccountsList_0(ctx context.Context, marshaler runtime.Marshaler, server KaijuAccountsRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SubaccountsListRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.SubaccountsList(ctx, &protoReq)
	return msg, metadata, err

}

func request_KaijuAccountsRPC_SubaccountBalancesList_0(ctx context.Context, marshaler runtime.Marshaler, client KaijuAccountsRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SubaccountBalancesListRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.SubaccountBalancesList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_KaijuAccountsRPC_SubaccountBalancesList_0(ctx context.Context, marshaler runtime.Marshaler, server KaijuAccountsRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SubaccountBalancesListRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.SubaccountBalancesList(ctx, &protoReq)
	return msg, metadata, err

}

func request_KaijuAccountsRPC_SubaccountBalanceEndpoint_0(ctx context.Context, marshaler runtime.Marshaler, client KaijuAccountsRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SubaccountBalanceRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.SubaccountBalanceEndpoint(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_KaijuAccountsRPC_SubaccountBalanceEndpoint_0(ctx context.Context, marshaler runtime.Marshaler, server KaijuAccountsRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SubaccountBalanceRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.SubaccountBalanceEndpoint(ctx, &protoReq)
	return msg, metadata, err

}

func request_KaijuAccountsRPC_StreamSubaccountBalance_0(ctx context.Context, marshaler runtime.Marshaler, client KaijuAccountsRPCClient, req *http.Request, pathParams map[string]string) (KaijuAccountsRPC_StreamSubaccountBalanceClient, runtime.ServerMetadata, error) {
	var protoReq StreamSubaccountBalanceRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.StreamSubaccountBalance(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

func request_KaijuAccountsRPC_SubaccountHistory_0(ctx context.Context, marshaler runtime.Marshaler, client KaijuAccountsRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SubaccountHistoryRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.SubaccountHistory(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_KaijuAccountsRPC_SubaccountHistory_0(ctx context.Context, marshaler runtime.Marshaler, server KaijuAccountsRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SubaccountHistoryRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.SubaccountHistory(ctx, &protoReq)
	return msg, metadata, err

}

func request_KaijuAccountsRPC_SubaccountOrderSummary_0(ctx context.Context, marshaler runtime.Marshaler, client KaijuAccountsRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SubaccountOrderSummaryRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.SubaccountOrderSummary(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_KaijuAccountsRPC_SubaccountOrderSummary_0(ctx context.Context, marshaler runtime.Marshaler, server KaijuAccountsRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SubaccountOrderSummaryRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.SubaccountOrderSummary(ctx, &protoReq)
	return msg, metadata, err

}

func request_KaijuAccountsRPC_Rewards_0(ctx context.Context, marshaler runtime.Marshaler, client KaijuAccountsRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RewardsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Rewards(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_KaijuAccountsRPC_Rewards_0(ctx context.Context, marshaler runtime.Marshaler, server KaijuAccountsRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RewardsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Rewards(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterKaijuAccountsRPCHandlerServer registers the http handlers for service KaijuAccountsRPC to "mux".
// UnaryRPC     :call KaijuAccountsRPCServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterKaijuAccountsRPCHandlerFromEndpoint instead.
func RegisterKaijuAccountsRPCHandlerServer(ctx context.Context, mux *runtime.ServeMux, server KaijuAccountsRPCServer) error {

	mux.Handle("POST", pattern_KaijuAccountsRPC_Portfolio_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/Portfolio", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/Portfolio"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_KaijuAccountsRPC_Portfolio_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_Portfolio_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_OrderStates_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/OrderStates", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/OrderStates"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_KaijuAccountsRPC_OrderStates_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_OrderStates_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_SubaccountsList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountsList", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountsList"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_KaijuAccountsRPC_SubaccountsList_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_SubaccountsList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_SubaccountBalancesList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountBalancesList", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountBalancesList"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_KaijuAccountsRPC_SubaccountBalancesList_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_SubaccountBalancesList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_SubaccountBalanceEndpoint_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountBalanceEndpoint", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountBalanceEndpoint"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_KaijuAccountsRPC_SubaccountBalanceEndpoint_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_SubaccountBalanceEndpoint_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_StreamSubaccountBalance_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_SubaccountHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountHistory", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountHistory"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_KaijuAccountsRPC_SubaccountHistory_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_SubaccountHistory_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_SubaccountOrderSummary_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountOrderSummary", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountOrderSummary"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_KaijuAccountsRPC_SubaccountOrderSummary_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_SubaccountOrderSummary_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_Rewards_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/Rewards", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/Rewards"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_KaijuAccountsRPC_Rewards_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_Rewards_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterKaijuAccountsRPCHandlerFromEndpoint is same as RegisterKaijuAccountsRPCHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterKaijuAccountsRPCHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterKaijuAccountsRPCHandler(ctx, mux, conn)
}

// RegisterKaijuAccountsRPCHandler registers the http handlers for service KaijuAccountsRPC to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterKaijuAccountsRPCHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterKaijuAccountsRPCHandlerClient(ctx, mux, NewKaijuAccountsRPCClient(conn))
}

// RegisterKaijuAccountsRPCHandlerClient registers the http handlers for service KaijuAccountsRPC
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "KaijuAccountsRPCClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "KaijuAccountsRPCClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "KaijuAccountsRPCClient" to call the correct interceptors.
func RegisterKaijuAccountsRPCHandlerClient(ctx context.Context, mux *runtime.ServeMux, client KaijuAccountsRPCClient) error {

	mux.Handle("POST", pattern_KaijuAccountsRPC_Portfolio_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/Portfolio", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/Portfolio"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KaijuAccountsRPC_Portfolio_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_Portfolio_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_OrderStates_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/OrderStates", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/OrderStates"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KaijuAccountsRPC_OrderStates_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_OrderStates_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_SubaccountsList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountsList", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountsList"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KaijuAccountsRPC_SubaccountsList_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_SubaccountsList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_SubaccountBalancesList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountBalancesList", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountBalancesList"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KaijuAccountsRPC_SubaccountBalancesList_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_SubaccountBalancesList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_SubaccountBalanceEndpoint_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountBalanceEndpoint", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountBalanceEndpoint"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KaijuAccountsRPC_SubaccountBalanceEndpoint_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_SubaccountBalanceEndpoint_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_StreamSubaccountBalance_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/StreamSubaccountBalance", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/StreamSubaccountBalance"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KaijuAccountsRPC_StreamSubaccountBalance_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_StreamSubaccountBalance_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_SubaccountHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountHistory", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountHistory"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KaijuAccountsRPC_SubaccountHistory_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_SubaccountHistory_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_SubaccountOrderSummary_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountOrderSummary", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/SubaccountOrderSummary"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KaijuAccountsRPC_SubaccountOrderSummary_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_SubaccountOrderSummary_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_KaijuAccountsRPC_Rewards_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/kaiju_accounts_rpc.KaijuAccountsRPC/Rewards", runtime.WithHTTPPathPattern("/kaiju_accounts_rpc.KaijuAccountsRPC/Rewards"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KaijuAccountsRPC_Rewards_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_KaijuAccountsRPC_Rewards_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_KaijuAccountsRPC_Portfolio_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"kaiju_accounts_rpc.KaijuAccountsRPC", "Portfolio"}, ""))

	pattern_KaijuAccountsRPC_OrderStates_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"kaiju_accounts_rpc.KaijuAccountsRPC", "OrderStates"}, ""))

	pattern_KaijuAccountsRPC_SubaccountsList_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"kaiju_accounts_rpc.KaijuAccountsRPC", "SubaccountsList"}, ""))

	pattern_KaijuAccountsRPC_SubaccountBalancesList_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"kaiju_accounts_rpc.KaijuAccountsRPC", "SubaccountBalancesList"}, ""))

	pattern_KaijuAccountsRPC_SubaccountBalanceEndpoint_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"kaiju_accounts_rpc.KaijuAccountsRPC", "SubaccountBalanceEndpoint"}, ""))

	pattern_KaijuAccountsRPC_StreamSubaccountBalance_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"kaiju_accounts_rpc.KaijuAccountsRPC", "StreamSubaccountBalance"}, ""))

	pattern_KaijuAccountsRPC_SubaccountHistory_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"kaiju_accounts_rpc.KaijuAccountsRPC", "SubaccountHistory"}, ""))

	pattern_KaijuAccountsRPC_SubaccountOrderSummary_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"kaiju_accounts_rpc.KaijuAccountsRPC", "SubaccountOrderSummary"}, ""))

	pattern_KaijuAccountsRPC_Rewards_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"kaiju_accounts_rpc.KaijuAccountsRPC", "Rewards"}, ""))
)

var (
	forward_KaijuAccountsRPC_Portfolio_0 = runtime.ForwardResponseMessage

	forward_KaijuAccountsRPC_OrderStates_0 = runtime.ForwardResponseMessage

	forward_KaijuAccountsRPC_SubaccountsList_0 = runtime.ForwardResponseMessage

	forward_KaijuAccountsRPC_SubaccountBalancesList_0 = runtime.ForwardResponseMessage

	forward_KaijuAccountsRPC_SubaccountBalanceEndpoint_0 = runtime.ForwardResponseMessage

	forward_KaijuAccountsRPC_StreamSubaccountBalance_0 = runtime.ForwardResponseStream

	forward_KaijuAccountsRPC_SubaccountHistory_0 = runtime.ForwardResponseMessage

	forward_KaijuAccountsRPC_SubaccountOrderSummary_0 = runtime.ForwardResponseMessage

	forward_KaijuAccountsRPC_Rewards_0 = runtime.ForwardResponseMessage
)
