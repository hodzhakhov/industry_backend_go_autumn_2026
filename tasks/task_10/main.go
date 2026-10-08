package main

import (
	"log"
	"log/slog"
	"net/http"
	"time"
)

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(rec, r)

		slog.Info("request",
			"method", r.Method,
			"path", r.URL.RequestURI(),
			"status", rec.status,
			"duration", time.Since(start),
		)
	})
}

func main() {
	addr := ":8080"
	repo := NewInMemoryTaskRepo(realClock{})
	handler := loggingMiddleware(NewHTTPHandler(repo))

	srv := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
