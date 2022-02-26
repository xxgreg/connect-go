// Code generated by protoc-gen-go-connect. DO NOT EDIT.
// versions:
// - protoc-gen-go-connect v0.0.1
// - protoc              v3.17.3
// source: grpc/health/v1/health.proto

package healthv1

import (
	context "context"
	errors "errors"
	connect "github.com/bufbuild/connect"
	v1 "github.com/bufbuild/connect/internal/gen/proto/go/grpc/health/v1"
	http "net/http"
	path "path"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the
// connect package are compatible. If you get a compiler error that this
// constant isn't defined, this code was generated with a version of connect
// newer than the one compiled into your binary. You can fix the problem by
// either regenerating this code with an older version of connect or updating
// the connect version compiled into your binary.
const _ = connect.IsAtLeastVersion0_0_1

// HealthClient is a client for the internal.health.v1.Health service.
type HealthClient interface {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check(context.Context, *connect.Envelope[v1.HealthCheckRequest]) (*connect.Envelope[v1.HealthCheckResponse], error)
	// Performs a watch for the serving status of the requested service.
	// The server will immediately send back a message indicating the current
	// serving status.  It will then subsequently send a new message whenever
	// the service's serving status changes.
	//
	// If the requested service is unknown when the call is received, the
	// server will send a message setting the serving status to
	// SERVICE_UNKNOWN but will *not* terminate the call.  If at some
	// future point, the serving status of the service becomes known, the
	// server will send a new message with the service's serving status.
	//
	// If the call terminates with status UNIMPLEMENTED, then clients
	// should assume this method is not supported and should not retry the
	// call.  If the call terminates with any other status (including OK),
	// clients should retry the call with appropriate exponential backoff.
	Watch(context.Context, *connect.Envelope[v1.HealthCheckRequest]) (*connect.ServerStreamForClient[v1.HealthCheckResponse], error)
}

// NewHealthClient constructs a client for the internal.health.v1.Health
// service. By default, it uses the binary protobuf codec.
//
// The URL supplied here should be the base URL for the gRPC server (e.g.,
// https://api.acme.com or https://acme.com/grpc).
func NewHealthClient(baseURL string, doer connect.Doer, opts ...connect.ClientOption) (HealthClient, error) {
	baseURL = strings.TrimRight(baseURL, "/")
	opts = append([]connect.ClientOption{
		connect.WithProtobuf(),
		connect.WithGzip(),
	}, opts...)
	var (
		client healthClient
		err    error
	)
	client.check, err = connect.NewUnaryClientImplementation[v1.HealthCheckRequest, v1.HealthCheckResponse](
		doer,
		baseURL,
		"internal.health.v1.Health/Check",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	client.watch, err = connect.NewStreamClientImplementation(
		doer,
		baseURL,
		"internal.health.v1.Health/Watch",
		connect.StreamTypeServer,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

// healthClient implements HealthClient.
type healthClient struct {
	check func(context.Context, *connect.Envelope[v1.HealthCheckRequest]) (*connect.Envelope[v1.HealthCheckResponse], error)
	watch func(context.Context) (connect.Sender, connect.Receiver)
}

var _ HealthClient = (*healthClient)(nil) // verify interface implementation

// Check calls internal.health.v1.Health.Check.
func (c *healthClient) Check(ctx context.Context, req *connect.Envelope[v1.HealthCheckRequest]) (*connect.Envelope[v1.HealthCheckResponse], error) {
	return c.check(ctx, req)
}

// Watch calls internal.health.v1.Health.Watch.
func (c *healthClient) Watch(ctx context.Context, req *connect.Envelope[v1.HealthCheckRequest]) (*connect.ServerStreamForClient[v1.HealthCheckResponse], error) {
	sender, receiver := c.watch(ctx)
	for key, values := range req.Header() {
		sender.Header()[key] = append(sender.Header()[key], values...)
	}
	for key, values := range req.Trailer() {
		sender.Trailer()[key] = append(sender.Trailer()[key], values...)
	}
	if err := sender.Send(req.Msg); err != nil {
		_ = sender.Close(err)
		_ = receiver.Close()
		return nil, err
	}
	if err := sender.Close(nil); err != nil {
		_ = receiver.Close()
		return nil, err
	}
	return connect.NewServerStreamForClient[v1.HealthCheckResponse](receiver), nil
}

// HealthHandler is an implementation of the internal.health.v1.Health service.
type HealthHandler interface {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check(context.Context, *connect.Envelope[v1.HealthCheckRequest]) (*connect.Envelope[v1.HealthCheckResponse], error)
	// Performs a watch for the serving status of the requested service.
	// The server will immediately send back a message indicating the current
	// serving status.  It will then subsequently send a new message whenever
	// the service's serving status changes.
	//
	// If the requested service is unknown when the call is received, the
	// server will send a message setting the serving status to
	// SERVICE_UNKNOWN but will *not* terminate the call.  If at some
	// future point, the serving status of the service becomes known, the
	// server will send a new message with the service's serving status.
	//
	// If the call terminates with status UNIMPLEMENTED, then clients
	// should assume this method is not supported and should not retry the
	// call.  If the call terminates with any other status (including OK),
	// clients should retry the call with appropriate exponential backoff.
	Watch(context.Context, *connect.Envelope[v1.HealthCheckRequest], *connect.ServerStream[v1.HealthCheckResponse]) error
}

// NewHealthHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the gRPC and gRPC-Web protocols with the binary
// protobuf and JSON codecs.
func NewHealthHandler(svc HealthHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	var lastHandlerPath string
	mux := http.NewServeMux()
	opts = append([]connect.HandlerOption{
		connect.WithProtobuf(),
		connect.WithProtobufJSON(),
		connect.WithGzip(),
	}, opts...)

	check := connect.NewUnaryHandler(
		"internal.health.v1.Health/Check", // procedure name
		"internal.health.v1.Health",       // reflection name
		svc.Check,
		opts...,
	)
	mux.Handle(check.Path(), check)
	lastHandlerPath = check.Path()

	watch := connect.NewStreamHandler(
		"internal.health.v1.Health/Watch", // procedure name
		"internal.health.v1.Health",       // reflection name
		connect.StreamTypeServer,
		func(ctx context.Context, sender connect.Sender, receiver connect.Receiver) {
			typed := connect.NewServerStream[v1.HealthCheckResponse](sender)
			req, err := connect.ReceiveUnaryEnvelope[v1.HealthCheckRequest](receiver)
			if err != nil {
				_ = receiver.Close()
				_ = sender.Close(err)
				return
			}
			if err = receiver.Close(); err != nil {
				_ = sender.Close(err)
				return
			}
			err = svc.Watch(ctx, req, typed)
			_ = sender.Close(err)
		},
		opts...,
	)
	mux.Handle(watch.Path(), watch)
	lastHandlerPath = watch.Path()

	return path.Dir(lastHandlerPath) + "/", mux
}

// UnimplementedHealthHandler returns CodeUnimplemented from all methods.
type UnimplementedHealthHandler struct{}

var _ HealthHandler = (*UnimplementedHealthHandler)(nil) // verify interface implementation

func (UnimplementedHealthHandler) Check(context.Context, *connect.Envelope[v1.HealthCheckRequest]) (*connect.Envelope[v1.HealthCheckResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("internal.health.v1.Health.Check isn't implemented"))
}

func (UnimplementedHealthHandler) Watch(context.Context, *connect.Envelope[v1.HealthCheckRequest], *connect.ServerStream[v1.HealthCheckResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("internal.health.v1.Health.Watch isn't implemented"))
}
