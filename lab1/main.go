package main

import (
	"log/slog"
	"net/http"
	"time"
)

func dayToNewYear(date time.Time) int {
	newYear := time.Date(date.Year()+1, 1, 1, 0, 0, 0, 0, date.Location())
	result := newYear.Sub(date).Hours() / 24
	return int(result)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("HTTP request",
			"method", r.Method,
			"path", r.URL.Path,
		)

		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/days", daysHandler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8080", loggingMiddleware(http.DefaultServeMux))
}
