package main

import (
	"net"
	service "scr/grpc/services"

	"google.golang.org/grpc"
)

func main() {
	rpcS := grpc.NewServer()
	service.RegisterProdServiceServer(rpcS, &service.Prod{})

	lis, _ := net.Listen("tcp", ":8099")
	rpcS.Serve(lis)
}
