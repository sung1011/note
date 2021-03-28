package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"scr/go-kit/endpoint"
	"strconv"
)

func DecUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	uid := r.URL.Query().Get("uid")
	if uid == "" {
		return nil, errors.New("must params uid")
	}
	id, _ := strconv.ParseInt(uid, 10, 64)
	return endpoint.UserReq{Uid: id}, nil
}

func EncUserResp(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}
