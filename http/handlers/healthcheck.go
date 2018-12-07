package handlers

import (
	"net/http"
	"time"
)

// HealthCheck just checks whether server works or not. It returns simple text word OK
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plaint")
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")
	w.Write([]byte("OK"))
}
