package handlers

import "net/http"

// NotFoundHandler handles non-existing route
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"resource not found: " + r.URL.Path + "\"}"))
}
