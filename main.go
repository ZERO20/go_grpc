package main

import (
	"context"
	"github.com/ZER020/go-grpc/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	list, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}
	serverRegister := grpc.NewServer()
	service := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegister, service)
	err = serverRegister.Serve(list)
	if err != nil {
		log.Fatal("Error serving:", err.Error())
	}
}
