package middleware

import "net/http"


var defaultHeaders = map[string]string{
	"Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,OPTIONS",
	"Access-Control-Allow-Origin":  "*",
	"ContentType": "application/json",
}

func BasicHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		injectDefaultHeader(w)

		h.ServeHTTP(w, r)
	})
}

func injectDefaultHeader(resp http.ResponseWriter) {
	for header, headerVal := range defaultHeaders {
		resp.Header().Set(header, headerVal)
	}
}
