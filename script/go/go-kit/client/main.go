package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"scr/go-kit/client/transport"
	"scr/go-kit/endpoint"

	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {
	tgt, err := url.Parse("http://127.0.0.1:8001")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c := kithttp.NewClient("GET", tgt, transport.GetUserInfoReq, transport.GetUserInfoResp)
	getUserInfo := c.Endpoint()

	ctx := context.Background()
	resp, err := getUserInfo(ctx, endpoint.UserReq{Uid: 99})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	userinfo := resp.(endpoint.UserResp)
	fmt.Println(userinfo.Name)
}
