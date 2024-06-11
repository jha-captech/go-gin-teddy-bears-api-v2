package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

const (
	blueBackground      = "\033[44m"
	lightBlueBackground = "\033[104m"
	greenBackground     = "\033[42m"
	yellowBackground    = "\033[43m"
	redBackground       = "\033[41m"
	resetColor          = "\033[0m"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		var (
			codeColorStart   string
			methodColorStart string
		)

		switch code := wrapped.statusCode; {
		case code >= 500:
			codeColorStart = redBackground
		case code >= 400:
			codeColorStart = yellowBackground
		case code >= 300:
			codeColorStart = blueBackground
		default:
			codeColorStart = greenBackground
		}

		switch r.Method {
		case http.MethodGet:
			methodColorStart = blueBackground
		case http.MethodPost:
			methodColorStart = lightBlueBackground
		case http.MethodPatch:
			methodColorStart = yellowBackground
		case http.MethodPut:
			methodColorStart = greenBackground
		case http.MethodDelete:
			methodColorStart = redBackground
		}

		slog.Info(
			fmt.Sprintf(
				"|%s %d %s| %12v | %14s |%s %6s %s| \"%s\"",
				codeColorStart,
				wrapped.statusCode,
				resetColor,
				time.Since(start),
				r.RemoteAddr,
				methodColorStart,
				r.Method,
				resetColor,
				r.URL.Path,
			),
		)
	})
}
