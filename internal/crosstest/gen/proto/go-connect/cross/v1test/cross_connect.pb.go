// Code generated by protoc-gen-go-connect. DO NOT EDIT.
// versions:
// - protoc-gen-go-connect v0.0.1
// - protoc              v3.17.3
// source: cross/v1test/cross.proto

package crossv1test

import (
	context "context"
	errors "errors"
	connect "github.com/bufconnect/connect"
	clientstream "github.com/bufconnect/connect/clientstream"
	protobuf "github.com/bufconnect/connect/codec/protobuf"
	handlerstream "github.com/bufconnect/connect/handlerstream"
	v1test "github.com/bufconnect/connect/internal/crosstest/gen/proto/go/cross/v1test"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the
// connect package are compatible. If you get a compiler error that this
// constant isn't defined, this code was generated with a version of connect
// newer than the one compiled into your binary. You can fix the problem by
// either regenerating this code with an older version of connect or updating
// the connect version compiled into your binary.
const _ = connect.SupportsCodeGenV0 // requires connect v0.0.1 or later

// WrappedCrossServiceClient is a client for the cross.v1test.CrossService
// service.
//
// It's a simplified wrapper around the full-featured API of
// UnwrappedCrossServiceClient.
type WrappedCrossServiceClient interface {
	Ping(context.Context, *v1test.PingRequest) (*v1test.PingResponse, error)
	Fail(context.Context, *v1test.FailRequest) (*v1test.FailResponse, error)
	Sum(context.Context) *clientstream.Client[v1test.SumRequest, v1test.SumResponse]
	CountUp(context.Context, *v1test.CountUpRequest) (*clientstream.Server[v1test.CountUpResponse], error)
	CumSum(context.Context) *clientstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]
}

// UnwrappedCrossServiceClient is a client for the cross.v1test.CrossService
// service. It's more complex than WrappedCrossServiceClient, but it gives
// callers more fine-grained control (e.g., sending and receiving headers).
type UnwrappedCrossServiceClient interface {
	Ping(context.Context, *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error)
	Fail(context.Context, *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error)
	Sum(context.Context) *clientstream.Client[v1test.SumRequest, v1test.SumResponse]
	CountUp(context.Context, *connect.Request[v1test.CountUpRequest]) (*clientstream.Server[v1test.CountUpResponse], error)
	CumSum(context.Context) *clientstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]
}

// CrossServiceClient is a client for the cross.v1test.CrossService service.
type CrossServiceClient struct {
	client unwrappedCrossServiceClient
}

var _ WrappedCrossServiceClient = (*CrossServiceClient)(nil)

// NewCrossServiceClient constructs a client for the cross.v1test.CrossService
// service. By default, it uses the binary protobuf codec.
//
// The URL supplied here should be the base URL for the gRPC server (e.g.,
// https://api.acme.com or https://acme.com/grpc).
func NewCrossServiceClient(baseURL string, doer connect.Doer, opts ...connect.ClientOption) (*CrossServiceClient, error) {
	baseURL = strings.TrimRight(baseURL, "/")
	opts = append([]connect.ClientOption{
		connect.Codec(protobuf.NameBinary, protobuf.NewBinary()),
	}, opts...)
	pingFunc, err := connect.NewClientFunc[v1test.PingRequest, v1test.PingResponse](
		doer,
		baseURL,
		"cross.v1test.CrossService/Ping",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	failFunc, err := connect.NewClientFunc[v1test.FailRequest, v1test.FailResponse](
		doer,
		baseURL,
		"cross.v1test.CrossService/Fail",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	sumFunc, err := connect.NewClientStream(
		doer,
		connect.StreamTypeClient,
		baseURL,
		"cross.v1test.CrossService/Sum",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	countUpFunc, err := connect.NewClientStream(
		doer,
		connect.StreamTypeServer,
		baseURL,
		"cross.v1test.CrossService/CountUp",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	cumSumFunc, err := connect.NewClientStream(
		doer,
		connect.StreamTypeBidirectional,
		baseURL,
		"cross.v1test.CrossService/CumSum",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return &CrossServiceClient{client: unwrappedCrossServiceClient{
		ping:    pingFunc,
		fail:    failFunc,
		sum:     sumFunc,
		countUp: countUpFunc,
		cumSum:  cumSumFunc,
	}}, nil
}

// Ping calls cross.v1test.CrossService.Ping.
func (c *CrossServiceClient) Ping(ctx context.Context, req *v1test.PingRequest) (*v1test.PingResponse, error) {
	res, err := c.client.Ping(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}

// Fail calls cross.v1test.CrossService.Fail.
func (c *CrossServiceClient) Fail(ctx context.Context, req *v1test.FailRequest) (*v1test.FailResponse, error) {
	res, err := c.client.Fail(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}

// Sum calls cross.v1test.CrossService.Sum.
func (c *CrossServiceClient) Sum(ctx context.Context) *clientstream.Client[v1test.SumRequest, v1test.SumResponse] {
	return c.client.Sum(ctx)
}

// CountUp calls cross.v1test.CrossService.CountUp.
func (c *CrossServiceClient) CountUp(ctx context.Context, req *v1test.CountUpRequest) (*clientstream.Server[v1test.CountUpResponse], error) {
	return c.client.CountUp(ctx, connect.NewRequest(req))
}

// CumSum calls cross.v1test.CrossService.CumSum.
func (c *CrossServiceClient) CumSum(ctx context.Context) *clientstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse] {
	return c.client.CumSum(ctx)
}

// Unwrap exposes the underlying generic client. Use it if you need finer
// control (e.g., sending and receiving headers).
func (c *CrossServiceClient) Unwrap() UnwrappedCrossServiceClient {
	return &c.client
}

type unwrappedCrossServiceClient struct {
	ping    func(context.Context, *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error)
	fail    func(context.Context, *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error)
	sum     func(context.Context) (connect.Sender, connect.Receiver)
	countUp func(context.Context) (connect.Sender, connect.Receiver)
	cumSum  func(context.Context) (connect.Sender, connect.Receiver)
}

var _ UnwrappedCrossServiceClient = (*unwrappedCrossServiceClient)(nil)

// Ping calls cross.v1test.CrossService.Ping.
func (c *unwrappedCrossServiceClient) Ping(ctx context.Context, req *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error) {
	return c.ping(ctx, req)
}

// Fail calls cross.v1test.CrossService.Fail.
func (c *unwrappedCrossServiceClient) Fail(ctx context.Context, req *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error) {
	return c.fail(ctx, req)
}

// Sum calls cross.v1test.CrossService.Sum.
func (c *unwrappedCrossServiceClient) Sum(ctx context.Context) *clientstream.Client[v1test.SumRequest, v1test.SumResponse] {
	sender, receiver := c.sum(ctx)
	return clientstream.NewClient[v1test.SumRequest, v1test.SumResponse](sender, receiver)
}

// CountUp calls cross.v1test.CrossService.CountUp.
func (c *unwrappedCrossServiceClient) CountUp(ctx context.Context, req *connect.Request[v1test.CountUpRequest]) (*clientstream.Server[v1test.CountUpResponse], error) {
	sender, receiver := c.countUp(ctx)
	if err := sender.Send(req.Msg); err != nil {
		_ = sender.Close(err)
		_ = receiver.Close()
		return nil, err
	}
	if err := sender.Close(nil); err != nil {
		_ = receiver.Close()
		return nil, err
	}
	return clientstream.NewServer[v1test.CountUpResponse](receiver), nil
}

// CumSum calls cross.v1test.CrossService.CumSum.
func (c *unwrappedCrossServiceClient) CumSum(ctx context.Context) *clientstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse] {
	sender, receiver := c.cumSum(ctx)
	return clientstream.NewBidirectional[v1test.CumSumRequest, v1test.CumSumResponse](sender, receiver)
}

// CrossService is an implementation of the cross.v1test.CrossService service.
//
// When writing your code, you can always implement the complete CrossService
// interface. However, if you don't need to work with headers, you can instead
// implement a simpler version of any or all of the unary methods. Where
// available, the simplified signatures are listed in comments.
//
// NewCrossService first tries to find the simplified version of each method,
// then falls back to the more complex version. If neither is implemented,
// connect.NewServeMux will return an error.
type CrossService interface {
	// Can also be implemented in a simplified form:
	// Ping(context.Context, *v1test.PingRequest) (*v1test.PingResponse, error)
	Ping(context.Context, *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error)
	// Can also be implemented in a simplified form:
	// Fail(context.Context, *v1test.FailRequest) (*v1test.FailResponse, error)
	Fail(context.Context, *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error)
	Sum(context.Context, *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error
	CountUp(context.Context, *connect.Request[v1test.CountUpRequest], *handlerstream.Server[v1test.CountUpResponse]) error
	CumSum(context.Context, *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error
}

// newUnwrappedCrossService wraps the service implementation in a
// connect.Service, which can then be passed to connect.NewServeMux.
//
// By default, services support the gRPC and gRPC-Web protocols with the binary
// protobuf and JSON codecs.
func newUnwrappedCrossService(svc CrossService, opts ...connect.HandlerOption) *connect.Service {
	handlers := make([]connect.Handler, 0, 5)
	opts = append([]connect.HandlerOption{
		connect.Codec(protobuf.NameBinary, protobuf.NewBinary()),
		connect.Codec(protobuf.NameJSON, protobuf.NewJSON()),
	}, opts...)

	ping, err := connect.NewUnaryHandler(
		"cross.v1test.CrossService/Ping", // procedure name
		"cross.v1test.CrossService",      // reflection name
		svc.Ping,
		opts...,
	)
	if err != nil {
		return connect.NewService(nil, err)
	}
	handlers = append(handlers, *ping)

	fail, err := connect.NewUnaryHandler(
		"cross.v1test.CrossService/Fail", // procedure name
		"cross.v1test.CrossService",      // reflection name
		svc.Fail,
		opts...,
	)
	if err != nil {
		return connect.NewService(nil, err)
	}
	handlers = append(handlers, *fail)

	sum, err := connect.NewStreamingHandler(
		connect.StreamTypeClient,
		"cross.v1test.CrossService/Sum", // procedure name
		"cross.v1test.CrossService",     // reflection name
		func(ctx context.Context, sender connect.Sender, receiver connect.Receiver) {
			typed := handlerstream.NewClient[v1test.SumRequest, v1test.SumResponse](sender, receiver)
			err := svc.Sum(ctx, typed)
			_ = receiver.Close()
			if err != nil {
				if _, ok := connect.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = connect.Wrap(connect.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = connect.Wrap(connect.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = sender.Close(err)
		},
		opts...,
	)
	if err != nil {
		return connect.NewService(nil, err)
	}
	handlers = append(handlers, *sum)

	countUp, err := connect.NewStreamingHandler(
		connect.StreamTypeServer,
		"cross.v1test.CrossService/CountUp", // procedure name
		"cross.v1test.CrossService",         // reflection name
		func(ctx context.Context, sender connect.Sender, receiver connect.Receiver) {
			typed := handlerstream.NewServer[v1test.CountUpResponse](sender)
			req, err := connect.ReceiveRequest[v1test.CountUpRequest](receiver)
			if err != nil {
				_ = receiver.Close()
				_ = sender.Close(err)
				return
			}
			if err = receiver.Close(); err != nil {
				_ = sender.Close(err)
				return
			}
			err = svc.CountUp(ctx, req, typed)
			if err != nil {
				if _, ok := connect.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = connect.Wrap(connect.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = connect.Wrap(connect.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = sender.Close(err)
		},
		opts...,
	)
	if err != nil {
		return connect.NewService(nil, err)
	}
	handlers = append(handlers, *countUp)

	cumSum, err := connect.NewStreamingHandler(
		connect.StreamTypeBidirectional,
		"cross.v1test.CrossService/CumSum", // procedure name
		"cross.v1test.CrossService",        // reflection name
		func(ctx context.Context, sender connect.Sender, receiver connect.Receiver) {
			typed := handlerstream.NewBidirectional[v1test.CumSumRequest, v1test.CumSumResponse](sender, receiver)
			err := svc.CumSum(ctx, typed)
			_ = receiver.Close()
			if err != nil {
				if _, ok := connect.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = connect.Wrap(connect.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = connect.Wrap(connect.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = sender.Close(err)
		},
		opts...,
	)
	if err != nil {
		return connect.NewService(nil, err)
	}
	handlers = append(handlers, *cumSum)

	return connect.NewService(handlers, nil)
}

type pluggableCrossServiceServer struct {
	ping    func(context.Context, *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error)
	fail    func(context.Context, *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error)
	sum     func(context.Context, *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error
	countUp func(context.Context, *connect.Request[v1test.CountUpRequest], *handlerstream.Server[v1test.CountUpResponse]) error
	cumSum  func(context.Context, *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error
}

func (s *pluggableCrossServiceServer) Ping(ctx context.Context, req *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error) {
	return s.ping(ctx, req)
}

func (s *pluggableCrossServiceServer) Fail(ctx context.Context, req *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error) {
	return s.fail(ctx, req)
}

func (s *pluggableCrossServiceServer) Sum(ctx context.Context, stream *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error {
	return s.sum(ctx, stream)
}

func (s *pluggableCrossServiceServer) CountUp(ctx context.Context, req *connect.Request[v1test.CountUpRequest], stream *handlerstream.Server[v1test.CountUpResponse]) error {
	return s.countUp(ctx, req, stream)
}

func (s *pluggableCrossServiceServer) CumSum(ctx context.Context, stream *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error {
	return s.cumSum(ctx, stream)
}

// NewCrossService wraps the service implementation in a connect.Service, ready
// for use with connect.NewServeMux. By default, services support the gRPC and
// gRPC-Web protocols with the binary protobuf and JSON codecs.
//
// The service implementation may mix and match the signatures of CrossService
// and the simplified signatures described in its comments. For each method,
// NewCrossService first tries to find a simplified implementation. If a simple
// implementation isn't available, it falls back to the more complex
// implementation. If neither is available, connect.NewServeMux will return an
// error.
//
// Taken together, this approach lets implementations embed
// UnimplementedCrossService and implement each method using whichever signature
// is most convenient.
func NewCrossService(svc any, opts ...connect.HandlerOption) *connect.Service {
	var impl pluggableCrossServiceServer

	// Find an implementation of Ping
	if pinger, ok := svc.(interface {
		Ping(context.Context, *v1test.PingRequest) (*v1test.PingResponse, error)
	}); ok {
		impl.ping = func(ctx context.Context, req *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error) {
			res, err := pinger.Ping(ctx, req.Msg)
			if err != nil {
				return nil, err
			}
			return connect.NewResponse(res), nil
		}
	} else if pinger, ok := svc.(interface {
		Ping(context.Context, *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error)
	}); ok {
		impl.ping = pinger.Ping
	} else {
		return connect.NewService(nil, errors.New("no Ping implementation found"))
	}

	// Find an implementation of Fail
	if failer, ok := svc.(interface {
		Fail(context.Context, *v1test.FailRequest) (*v1test.FailResponse, error)
	}); ok {
		impl.fail = func(ctx context.Context, req *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error) {
			res, err := failer.Fail(ctx, req.Msg)
			if err != nil {
				return nil, err
			}
			return connect.NewResponse(res), nil
		}
	} else if failer, ok := svc.(interface {
		Fail(context.Context, *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error)
	}); ok {
		impl.fail = failer.Fail
	} else {
		return connect.NewService(nil, errors.New("no Fail implementation found"))
	}

	// Find an implementation of Sum
	if sumer, ok := svc.(interface {
		Sum(context.Context, *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error
	}); ok {
		impl.sum = sumer.Sum
	} else {
		return connect.NewService(nil, errors.New("no Sum implementation found"))
	}

	// Find an implementation of CountUp
	if countUper, ok := svc.(interface {
		CountUp(context.Context, *v1test.CountUpRequest, *handlerstream.Server[v1test.CountUpResponse]) error
	}); ok {
		impl.countUp = func(ctx context.Context, req *connect.Request[v1test.CountUpRequest], stream *handlerstream.Server[v1test.CountUpResponse]) error {
			return countUper.CountUp(ctx, req.Msg, stream)
		}
	} else if countUper, ok := svc.(interface {
		CountUp(context.Context, *connect.Request[v1test.CountUpRequest], *handlerstream.Server[v1test.CountUpResponse]) error
	}); ok {
		impl.countUp = countUper.CountUp
	} else {
		return connect.NewService(nil, errors.New("no CountUp implementation found"))
	}

	// Find an implementation of CumSum
	if cumSumer, ok := svc.(interface {
		CumSum(context.Context, *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error
	}); ok {
		impl.cumSum = cumSumer.CumSum
	} else {
		return connect.NewService(nil, errors.New("no CumSum implementation found"))
	}

	return newUnwrappedCrossService(&impl, opts...)
}

var _ CrossService = (*UnimplementedCrossService)(nil) // verify interface implementation

// UnimplementedCrossService returns CodeUnimplemented from all methods.
type UnimplementedCrossService struct{}

func (UnimplementedCrossService) Ping(context.Context, *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error) {
	return nil, connect.Errorf(connect.CodeUnimplemented, "cross.v1test.CrossService.Ping isn't implemented")
}

func (UnimplementedCrossService) Fail(context.Context, *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error) {
	return nil, connect.Errorf(connect.CodeUnimplemented, "cross.v1test.CrossService.Fail isn't implemented")
}

func (UnimplementedCrossService) Sum(context.Context, *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error {
	return connect.Errorf(connect.CodeUnimplemented, "cross.v1test.CrossService.Sum isn't implemented")
}

func (UnimplementedCrossService) CountUp(context.Context, *connect.Request[v1test.CountUpRequest], *handlerstream.Server[v1test.CountUpResponse]) error {
	return connect.Errorf(connect.CodeUnimplemented, "cross.v1test.CrossService.CountUp isn't implemented")
}

func (UnimplementedCrossService) CumSum(context.Context, *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error {
	return connect.Errorf(connect.CodeUnimplemented, "cross.v1test.CrossService.CumSum isn't implemented")
}