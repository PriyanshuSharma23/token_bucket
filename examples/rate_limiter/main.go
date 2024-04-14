package main

import (
	"log"
	"net/http"

	"github.com/PriyanshuSharma23/token_bucket/bucket"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	server := http.Server{
		Handler: rateLimiter(mux),
		Addr:    ":6060",
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func rateLimiter(next http.Handler) http.Handler {
	limiter := bucket.NewBucket(2, 6)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Check(1) {
			http.Error(w, "rate limit", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
