package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var (
	port             int
	host             string
	enableHealth     bool
	enableReflection bool
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.LUTC | log.Lshortfile)

	flag.IntVar(&port, "port", 6565, "Server's listening port")
	flag.StringVar(&host, "address", "0.0.0.0", "Server's listening address")
	flag.BoolVar(&enableHealth, "enable-health", true, "Enable or disable health grpc service")
	flag.BoolVar(&enableReflection, "enable-reflection", true, "Enable or disable reflection service")

	flag.Parse()

	addr := fmt.Sprintf("%s:%d", host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("could not listen tcp on %s: %s", addr, err.Error())
	}

	srv := grpc.NewServer()

	RegisterGreeterServer(srv, &greeter{})

	if enableHealth {
		healthgrpc.RegisterHealthServer(srv, health.NewServer())
	}

	if enableReflection {
		reflection.Register(srv)
	}

    log.Printf("starting greeter server at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error serving grpc server: %s", err.Error())
	}
}

type greeter struct {
	UnimplementedGreeterServer
}

func (s *greeter) SayHello(_ context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("received: %s", in.GetName())
	return &HelloReply{Message: fmt.Sprintf("Hello, %s!", in.GetName())}, nil
}
