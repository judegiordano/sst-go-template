package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/judegiordano/gogetem/pkg/logger"
)

type Middleware func(http.Handler) http.Handler

type ErrorResponse struct {
	Error string `json:"error"`
}

func Stack(ms ...Middleware) Middleware {
	return Middleware(func(next http.Handler) http.Handler {
		for _, mid := range ms {
			next = mid(next)
		}
		return next
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func TransformJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				status := http.StatusInternalServerError
				w.WriteHeader(status)
				json.NewEncoder(w).Encode(ErrorResponse{
					Error: fmt.Sprint(err),
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("checking if is admin...")
		next.ServeHTTP(w, r)
	})
}
