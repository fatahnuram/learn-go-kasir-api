package main

import (
	"log"
	"net/http"

	"github.com/fatahnuram/learn-go-kasir-api/internal/handler"
	"github.com/fatahnuram/learn-go-kasir-api/internal/middleware"
	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
)

var products = []model.Product{
	{
		ID:    1,
		Name:  "Indomie",
		Price: 3000,
		Stock: 3,
	},
	{
		ID:    2,
		Name:  "Lifeboy",
		Price: 1500,
		Stock: 5,
	},
	{
		ID:    3,
		Name:  "Kacang Garuda",
		Price: 500,
		Stock: 4,
	},
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /healthz", handler.Healthz())

	log.Println("running server on port 8080..")
	err := http.ListenAndServe(":8080", middleware.SimpleLogger(mux))
	if err != nil {
		log.Println("failed to run server:", err)
	}
}
