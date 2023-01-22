package controllers

import "net/http"

type HealthCheck struct{}

func (h *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("app is running!"))
	}
}
