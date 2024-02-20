package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/cod3rboy/demo-grpc/proto"
	"github.com/cod3rboy/demo-grpc/services"
	"google.golang.org/grpc"
)

var (
	port = flag.String("port", "8000", "server port")
)

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("error while starting listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register our service with gRPC server
	invoicerService := services.NewInvoicerService()
	pb.RegisterInvoicerServiceServer(grpcServer, invoicerService)

	log.Printf("gRPC server listening on address :%s", *port)
	// start gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("error while starting grpc server: %v", err)
	}
}
