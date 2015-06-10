package healthcheck

import "net/http"

// Pattern.
const Pattern = "/internal/health"

// Add handler to default serve mux.
func AddHandler() {
	http.HandleFunc(Pattern, health)
}

// Health check end-point.
func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", "0")
	w.WriteHeader(200)
}
