package main

import (
	"net/http"
	"scr/go-kit/endpoint"
	"scr/go-kit/service"
	"scr/go-kit/transport"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	edp := endpoint.GenUserEndPoint(&service.User{})

	s := kithttp.NewServer(edp, transport.DecUserReq, transport.EncUserResp)

	r := mux.NewRouter()
	r.Handle("/user/{uid:\\d+}", s)

	http.ListenAndServe(":8001", s)
}
