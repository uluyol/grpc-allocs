package main

import (
	"context"
	"net"
	"testing"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/uluyol/grpc-allocs/flat"
	"github.com/uluyol/grpc-allocs/pb2gogo"
	"github.com/uluyol/grpc-allocs/pb3"
	"google.golang.org/grpc"
)

type pb3s struct{}

var data = []byte("hElLo")

func (pb3s) Get(_ context.Context, req *pb3.Req) (*pb3.Resp, error) {
	resp := new(pb3.Resp)
	resp.Succ = true
	resp.F2 = &pb3.F2{AllVals: []string{req.Key, "1"}}
	resp.F1 = &pb3.F1{Seq: 2, Data: data}
	return resp, nil
}

var pb3Sink *pb3.Resp

type pb2gogos struct{}

func (pb2gogos) Get(_ context.Context, req *pb2gogo.Req) (*pb2gogo.Resp, error) {
	resp := new(pb2gogo.Resp)
	resp.Succ = true
	resp.F2 = pb2gogo.F2{AllVals: []string{req.Key, "1"}}
	resp.F1 = pb2gogo.F1{Seq: 2, Data: data}
	return resp, nil
}

var pb2gogoSink *pb2gogo.Resp

type fbtabs struct{}

func (fbtabs) Get(_ context.Context, req *flat.Req) (*flatbuffers.Builder, error) {
	b := flatbuffers.NewBuilder(256)
	dataOff := b.CreateByteVector(data)
	flat.F1Start(b)
	flat.F1AddSeq(b, 2)
	flat.F1AddData(b, dataOff)
	f1Off := flat.F1End(b)
	av1Off := b.CreateByteString(req.Key())
	av2Off := b.CreateString("1")
	flat.F2StartAllValsVector(b, 2)
	b.PrependUOffsetT(av1Off)
	b.PrependUOffsetT(av2Off)
	avOff := b.EndVector(2)
	flat.F2Start(b)
	flat.F2AddAllVals(b, avOff)
	f2Off := flat.F2End(b)
	flat.RespStart(b)
	flat.RespAddSucc(b, 1)
	flat.RespAddF1(b, f1Off)
	flat.RespAddF2(b, f2Off)
	b.Finish(flat.RespEnd(b))
	return b, nil
}

var fbtabSink *flat.Resp

func unixDialer(addr string, _ time.Duration) (net.Conn, error) { return net.Dial("unix", addr) }

func BenchmarkProto3(b *testing.B) {
	const addr = "/tmp/alloc-bench-proto3.socket"
	lis, err := net.Listen("unix", addr)
	if err != nil {
		b.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb3.RegisterSvcServer(s, &pb3s{})
	go func() {
		if err := s.Serve(lis); err != nil {
			b.Fatalf("Failed to serve: %v", err)
			b.FailNow()
		}
	}()
	defer s.Stop()
	defer lis.Close()
	conn, err := grpc.Dial(addr, grpc.WithDialer(unixDialer), grpc.WithInsecure())
	if err != nil {
		b.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb3.NewSvcClient(conn)

	req := &pb3.Req{Key: "google3"}
	pb3Sink, err = client.Get(context.Background(), req)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pb3Sink, err = client.Get(context.Background(), req)
	}
}

func BenchmarkProto2Gogo(b *testing.B) {
	const addr = "/tmp/alloc-bench-proto2gogo.socket"
	lis, err := net.Listen("unix", addr)
	if err != nil {
		b.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb2gogo.RegisterSvcServer(s, &pb2gogos{})
	go func() {
		if err := s.Serve(lis); err != nil {
			b.Fatalf("Failed to serve: %v", err)
			b.FailNow()
		}
	}()
	defer s.Stop()
	defer lis.Close()
	conn, err := grpc.Dial(addr, grpc.WithDialer(unixDialer), grpc.WithInsecure())
	if err != nil {
		b.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb2gogo.NewSvcClient(conn)

	req := &pb2gogo.Req{Key: "google3"}
	pb2gogoSink, err = client.Get(context.Background(), req)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pb2gogoSink, err = client.Get(context.Background(), req)
	}
}

func BenchmarkFlatbufferTable(b *testing.B) {
	const addr = "/tmp/alloc-bench-fbtab.socket"
	lis, err := net.Listen("unix", addr)
	if err != nil {
		b.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))
	flat.RegisterSvcServer(s, &fbtabs{})
	go func() {
		if err := s.Serve(lis); err != nil {
			b.Fatalf("Failed to serve: %v", err)
			b.FailNow()
		}
	}()
	defer s.Stop()
	defer lis.Close()
	conn, err := grpc.Dial(addr, grpc.WithDialer(unixDialer), grpc.WithInsecure(), grpc.WithCodec(flatbuffers.FlatbuffersCodec{}))
	if err != nil {
		b.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := flat.NewSvcClient(conn)

	fb := flatbuffers.NewBuilder(0)
	keyOff := fb.CreateString("google3")
	flat.ReqStart(fb)
	flat.ReqAddKey(fb, keyOff)
	fb.Finish(flat.ReqEnd(fb))
	fbtabSink, err = client.Get(context.Background(), fb)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fbtabSink, err = client.Get(context.Background(), fb)
	}
}
