package util

import (
	"context"
	"net/http"
)

type Err struct {
	Code int
	Msg  string
}

func NewErr(code int, msg string) error {
	return &Err{Code: code, Msg: msg}
}

func (this *Err) Error() string {
	return this.Msg
}

func ErrEnc(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if e, ok := err.(*Err); ok {
		w.WriteHeader(e.Code)
	} else {
		w.WriteHeader(500)
	}
	w.Write([]byte(err.Error()))
}
