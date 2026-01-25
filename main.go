package main

import (
	"log"
	"net/http"

	"github.com/fatahnuram/learn-go-kasir-api/internal/handler"
	"github.com/fatahnuram/learn-go-kasir-api/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /healthz", handler.Healthz())

	// products
	mux.Handle("GET /api/products", handler.ListProducts())

	log.Println("running server on port 8080..")
	err := http.ListenAndServe(":8080", middleware.SimpleLogger(middleware.DefaultHeaders(mux)))
	if err != nil {
		log.Println("failed to run server:", err)
	}
}
