package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"scr/go-kit/endpoint"
	"scr/go-kit/mw"
	"scr/go-kit/service"
	"scr/go-kit/transport"
	"scr/go-kit/util"
	"syscall"

	"github.com/go-kit/kit/log"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

// docker run -d --name=cs -p=8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0
// -bootstrap 指定自己为leader, 而不需要选举
// -ui 启动一个内置的web页面
// -client 指定客户端可以访问的IP, 0.0.0.0任意访问, 否则默认本机访问

func main() {
	flag.StringVar(&util.SvcName, "name", "", "service name")
	flag.StringVar(&util.SvcPort, "port", "", "service port")
	flag.Parse()

	r := httprouter.New()
	r.GET("/health", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	limiter := rate.NewLimiter(1, 3)

	var logger kitlog.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.WithPrefix(logger, "tickles", "v1.0")
	logger = log.With(logger, "TIME", log.DefaultTimestampUTC)
	logger = log.With(logger, "CALLER", log.DefaultCaller)

	ep := endpoint.GenUserEndPoint(&service.User{})
	ep = endpoint.RateLimit(limiter)(ep)
	ep = endpoint.Log(logger)(ep)
	op := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(util.ErrEnc),
	}
	r.Handler("GET", "/user", kithttp.NewServer(
		ep,
		transport.DecUserReq,
		transport.EncUserResp,
		op...,
	))

	var h http.Handler
	h = mw.Limit(r)

	errChan := make(chan error)
	go func() {
		err := util.RegService()
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		err := http.ListenAndServe(":"+util.SvcPort, h)
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sig)
	}()

	e := <-errChan
	util.UnregService()
	fmt.Println(e)
}
