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
