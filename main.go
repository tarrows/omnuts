package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/c10t/omnuts/pb"
)

type BacklogServer struct{}

func (s *BacklogServer) List(ctx context.Context, in *pb.ListBacklogsRequest) (*pb.ListBacklogsResponse, error) {
	return &pb.ListBacklogsResponse{
		Backlogs: []*pb.Backlog{},
	}, nil
}

func (s *BacklogServer) Get(ctx context.Context, in *pb.GetBacklogRequest) (*pb.Backlog, error) {
	return &pb.Backlog{
		Id: in.GetId(),
	}, nil
}

func startGRPCServer(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	backlogServer := &BacklogServer{}
	pb.RegisterBacklogServiceServer(s, backlogServer)

	reflection.Register(s)
	log.Println("Starting gRPC server on:", port)
	return s.Serve(lis)
}

func startHTTPServer(grpcPort, httpPort int) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	muxOpts := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		OrigName:     true,
		EmitDefaults: true,
	})
	gatewayMux := runtime.NewServeMux(muxOpts)

	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf("127.0.0.1:%d", grpcPort)
	if err := pb.RegisterBacklogServiceHandlerFromEndpoint(ctx, gatewayMux, endpoint, dialOpts); err != nil {
		return err
	}
	log.Println("Listening HTTP on:", httpPort)

	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), gatewayMux)
}

func main() {
	httpPort := 5000
	grpcPort := 5001

	errors := make(chan error)

	go func() {
		errors <- startGRPCServer(grpcPort)
	}()

	go func() {
		errors <- startHTTPServer(grpcPort, httpPort)
	}()

	for err := range errors {
		log.Fatalln(err)
	}
}
