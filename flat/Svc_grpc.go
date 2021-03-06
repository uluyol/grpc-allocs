//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: data

package flat

import "github.com/google/flatbuffers/go"

import (
  context "golang.org/x/net/context"
  grpc "google.golang.org/grpc"
)

// Client API for Svc service
type SvcClient interface{
  Get(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* Resp, error)  
}

type svcClient struct {
  cc *grpc.ClientConn
}

func NewSvcClient(cc *grpc.ClientConn) SvcClient {
  return &svcClient{cc}
}

func (c *svcClient) Get(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* Resp, error) {
  out := new(Resp)
  err := grpc.Invoke(ctx, "/flat.Svc/Get", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

// Server API for Svc service
type SvcServer interface {
  Get(context.Context, *Req) (*flatbuffers.Builder, error)  
}

func RegisterSvcServer(s *grpc.Server, srv SvcServer) {
  s.RegisterService(&_Svc_serviceDesc, srv)
}

func _Svc_Get_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(Req)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(SvcServer).Get(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/flat.Svc/Get",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(SvcServer).Get(ctx, req.(* Req))
  }
  return interceptor(ctx, in, info, handler)
}


var _Svc_serviceDesc = grpc.ServiceDesc{
  ServiceName: "flat.Svc",
  HandlerType: (*SvcServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "Get",
      Handler: _Svc_Get_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
  },
}

