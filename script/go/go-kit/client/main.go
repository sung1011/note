package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"scr/go-kit/client/transport"
	"scr/go-kit/endpoint"

	kitendpoint "github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	kitsdconsul "github.com/go-kit/kit/sd/consul"
	kithttp "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	registry()

	// direct()
}

func registry() {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500" // 注册中心
	clientAPI, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	client := kitsdconsul.NewClient(clientAPI)

	var logger kitlog.Logger
	logger = kitlog.NewLogfmtLogger(os.Stdout)
	tags := []string{"primary"}
	serviceName := "kkk1"
	instancer := kitsdconsul.NewInstancer(client, logger, serviceName, tags, true)

	factory := func(svcURL string) (kitendpoint.Endpoint, io.Closer, error) {
		tgt, _ := url.Parse("http://" + svcURL)
		return kithttp.NewClient("GET", tgt, transport.GetUserInfoReq, transport.GetUserInfoResp).Endpoint(), nil, nil
	}
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	eps, err := endpointer.Endpoints()
	if err != nil {
		log.Fatal(err)
	}
	doClient(eps[0])
}

func direct() {
	tgt, err := url.Parse("http://127.0.0.1:8001")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c := kithttp.NewClient("GET", tgt, transport.GetUserInfoReq, transport.GetUserInfoResp)
	doClient(c.Endpoint())
}

func doClient(ep kitendpoint.Endpoint) {

	ctx := context.Background()
	resp, err := ep(ctx, endpoint.UserReq{Uid: 99})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	userinfo := resp.(endpoint.UserResp)
	fmt.Println(userinfo.Name)
}
