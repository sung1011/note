package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"runtime"
	"scr/go-kit/client/transport"
	"scr/go-kit/endpoint"
	"time"

	kitendpoint "github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	kitsdconsul "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	kithttp "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	PickOne()

	Direct()

	go RoundRobin()

	go RoundRandom()

	time.Sleep(5 * time.Second)
}

func PickOne() {
	eprs, err := sdEndpointer().Endpoints()
	if err != nil {
		log.Fatal(err)
	}
	doClient(eprs[0], funcName())
}

func Direct() {
	tgt, err := url.Parse("http://127.0.0.1:8001")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c := kithttp.NewClient("GET", tgt, transport.GetUserInfoReq, transport.GetUserInfoResp)
	doClient(c.Endpoint(), funcName())
}

// RoundRobin 同一个service 多个instance 进行轮询
func RoundRobin() {
	endpointer := sdEndpointer()
	c := lb.NewRoundRobin(endpointer)
	for {
		ep, err := c.Endpoint()
		if err != nil {
			log.Fatal(err)
		}
		doClient(ep, funcName())
		time.Sleep(1 * time.Second)
	}
}

func RoundRandom() {
	endpointr := sdEndpointer()
	seed := time.Now().UnixNano()
	c := lb.NewRandom(endpointr, seed)
	for {
		ep, err := c.Endpoint()
		if err != nil {
			log.Fatal(err)
		}
		doClient(ep, funcName())
		time.Sleep(1 * time.Second)
	}
}

func sdEndpointer() sd.Endpointer {
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
	serviceName := "sunji01"
	instancer := kitsdconsul.NewInstancer(client, logger, serviceName, tags, true)

	factory := func(svcURL string) (kitendpoint.Endpoint, io.Closer, error) {
		tgt, err := url.Parse("http://" + svcURL)
		if err != nil {
			return nil, nil, err
		}
		return kithttp.NewClient("GET", tgt, transport.GetUserInfoReq, transport.GetUserInfoResp).Endpoint(), nil, nil
	}
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	return endpointer
}

func doClient(ep kitendpoint.Endpoint, prefix string) {
	ctx := context.Background()
	resp, err := ep(ctx, endpoint.UserReq{Uid: 99})
	if err != nil {
		fmt.Println(err, resp)
		os.Exit(1)
	}
	userinfo := resp.(endpoint.UserResp)
	fmt.Println(prefix, userinfo.Name)
}

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
