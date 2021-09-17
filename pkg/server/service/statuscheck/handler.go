package statuscheck

import (
	"fmt"
	"net/http"
)

func AddHandlers(mux *http.ServeMux, liveness, readiness Monitor) {
	AddReadinessHandler(mux, readiness)
	AddLivenessHandler(mux, liveness)
}

func AddLivenessHandler(mux *http.ServeMux, monitor Monitor) {
	mux.HandleFunc("/health", makeHealthHandler(monitor))
}

func AddReadinessHandler(mux *http.ServeMux, monitor Monitor) {
	mux.HandleFunc("/health/readiness", makeHealthHandler(monitor))
}

func makeHealthHandler(monitor Monitor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("content-type", "application/json")

		if monitor.Status() == Good {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, `{"status":"%s"}`, string(monitor.Status()))
	}
}
