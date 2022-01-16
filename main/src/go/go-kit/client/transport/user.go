package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"scr/go-kit/endpoint"
	"strconv"
)

func GetUserInfoReq(ctx context.Context, req *http.Request, r interface{}) error {
	user_req := r.(endpoint.UserReq)
	req.URL.Path += `/user/`
	req.URL.RawQuery = "uid=" + strconv.Itoa(int(user_req.Uid))
	return nil
}

func GetUserInfoResp(ctx context.Context, resp *http.Response) (response interface{}, err error) {
	if resp.StatusCode > 400 {
		return nil, errors.New(fmt.Sprintf("bad resp %v", resp.StatusCode))
	}

	var user endpoint.UserResp
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
