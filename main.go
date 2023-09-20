package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	port := ":8080"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/api", Home)
	r.Get("/api/get-encode", GetHandlerEncode)
	r.Get("/api/get-marshal", GetHandlerMarshal)
	r.Post("/api/convert/list", PostConversions)
	r.Post("/api/convert/weight-us", PostWeightUS)
	r.Post("/api/convert/volume-us", PostVolumeUS)
	r.Post("/api/convert/weight-metric", PostWeightMetric)
	r.Post("/api/convert/volume-metric", PostVolumeMetric)
	fmt.Println("Listening on port " + port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("failed to start server: %w", err)
	}
}
