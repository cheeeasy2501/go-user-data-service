package middleware

import "net/http"

func DefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		fn(writer, request)
	}
}
