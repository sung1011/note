package main

import (
	"context"
	"fmt"
	"log"
	service "scr/grpc/services"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8099", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := service.NewProdServiceClient(conn)
	resp, err := client.GetProdStock(context.Background(), &service.ProdRequest{Id: 999})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Stock)
}
