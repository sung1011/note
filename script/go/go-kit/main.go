package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"scr/go-kit/endpoint"
	"scr/go-kit/service"
	"scr/go-kit/transport"
	"scr/go-kit/util"
	"syscall"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

// docker run -d --name=cs -p=8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0
// -bootstrap 指定自己为leader, 而不需要选举
// -ui 启动一个内置的web页面
// -client 指定客户端可以访问的IP, 0.0.0.0任意访问, 否则默认本机访问

func main() {
	// r := mux.NewRouter()
	// r.Handle("/user/{uid:\\d+}", s)

	r := httprouter.New()
	r.GET("/health", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})
	r.Handler("GET", "/user", kithttp.NewServer(
		endpoint.GenUserEndPoint(&service.User{}),
		transport.DecUserReq,
		transport.EncUserResp,
	))

	errChan := make(chan error)
	go func() {
		err := util.RegService()
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		err := http.ListenAndServe(":8001", r)
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sig)
	}()

	e := <-errChan
	util.UnregService()
	fmt.Println(e)
}
