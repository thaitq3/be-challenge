package middleware

import "net/http"

func LimitBody(h http.Handler, bodySizeLimit int64) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w,r.Body ,bodySizeLimit)

		h.ServeHTTP(w, r)
	})
}