package endpoint

import (
	"context"
	"scr/go-kit/service"
	"scr/go-kit/util"
	"time"

	kitlog "github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

type UserReq struct {
	Uid  int64  `json:"uid"`
	Name string `json:"name"`
}

type UserResp struct {
	Name string `json:"name"`
}

func Log(logger kitlog.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(UserReq)
			logger.Log("uid", r.Uid)
			return next(ctx, request)
		}
	}
}

func RateLimit(l *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !l.AllowN(time.Now(), 1) {
				return nil, util.NewErr(429, "too many req!~")
			}
			return next(ctx, request)
		}
	}
}

func GenUserEndPoint(userService service.IUser) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserReq)
		name := userService.GetName(r.Uid) + util.SvcPort
		return UserResp{Name: name}, nil
	}
}
