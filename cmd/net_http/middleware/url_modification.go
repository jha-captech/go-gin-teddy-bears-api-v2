package middleware

import (
	"net/http"
)

func AddTrailingBackSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path) > 0 && r.URL.Path[len(r.URL.Path)-1] != '/' {
			r.URL.Path = r.URL.Path + "/"
		}

		next.ServeHTTP(w, r)
	})
}
