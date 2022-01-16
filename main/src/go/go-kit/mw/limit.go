package mw

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

var r = rate.NewLimiter(1, 10)

func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if !r.AllowN(time.Now(), 1) {
			http.Error(w, "too many req", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, req)
	})
}
