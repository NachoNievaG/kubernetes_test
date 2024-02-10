package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"math"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	slog.Info("Starting the application...")
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		a := 0.0001
		for i := 0; i < 1000000; i++ {
			math.Sqrt(a)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprint(a)))
	}).Methods("GET")

	r.HandleFunc("/api/v1/catfact", func(w http.ResponseWriter, r *http.Request) {
		client := http.Client{}
		req, err := http.NewRequest("GET", "https://catfact.ninja/fact?max_length=40", nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
			return
		}
		result, err := client.Do(req)
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if statusOK := result.StatusCode >= 200 && result.StatusCode < 300; !statusOK {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
		}
		b, err := io.ReadAll(result.Body)
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
		}
		w.Write(b)
	}).Methods("GET")

	// add /prom handler
	r.Handle("/metrics", promhttp.Handler())
	slog.Info("Booting server...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
