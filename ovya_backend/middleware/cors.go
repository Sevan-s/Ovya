package middleware

import (
	"net/http"
	"os"
)

func WithCORS(next http.Handler) http.Handler {
	frontAdd := os.Getenv("FRONT_ADD")

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", frontAdd)
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if req.Method == "OPTIONS" {
			res.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(res, req)
	})
}
