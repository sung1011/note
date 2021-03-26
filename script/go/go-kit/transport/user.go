package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"scr/go-kit/endpoint"
	"strconv"

	"github.com/gorilla/mux"
)

func DecUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	uid, _ := vars["uid"]
	// uid := r.URL.Query().Get("uid")
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
