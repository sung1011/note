package endpoint

import (
	"context"
	"scr/go-kit/service"

	"github.com/go-kit/kit/endpoint"
)

type UserReq struct {
	Uid  int64  `json:"uid"`
	Name string `json:"name"`
}

type UserResp struct {
	Name string `json:"name"`
}

func GenUserEndPoint(userService service.IUser) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserReq)
		name := userService.GetName(r.Uid)
		return UserResp{Name: name}, nil
	}
}
