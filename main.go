package main

import (
	"context"
	"log"
	"net"

	"github.com/shamhub/go_grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {

	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8089")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)

	err = serverRegistrar.Serve(listener)
	if err != nil {
		log.Fatalf("unable to serve: %s", err)
	}
}
