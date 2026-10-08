package middleware

import (
	"Lab-Hacking/models"
	"log/slog"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := models.RecServer{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(&rec, r)

		durasi := time.Since(start)

		slog.Info("HTTP Request:",
			slog.String("Method", r.Method),
			slog.String("Path", r.URL.Path),
			slog.Int("StatusCode", rec.StatusCode),
			slog.Duration("Durasi:", durasi),
		)
	})
}
