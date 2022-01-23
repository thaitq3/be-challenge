package middleware

import (
	"errors"
	"net/http"
	"crud-challenge/utils"
)

func RecoverWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			re := recover()
			if re != nil {
				var err error
				switch t := re.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("unknown error")
				}

				utils.OutputInternalServerError(w, err)
			}
		}()

		h.ServeHTTP(w, r)
	})
}

