package main

import (
	"log"
	"net/http"

	"github.com/fatahnuram/learn-go-kasir-api/internal/handler"
	"github.com/fatahnuram/learn-go-kasir-api/internal/middleware"
	"github.com/fatahnuram/learn-go-kasir-api/internal/repository"
	"github.com/fatahnuram/learn-go-kasir-api/internal/service"
)

func main() {
	mux := http.NewServeMux()

	// healthcheck
	mux.Handle("GET /healthz", handler.Healthz())
	mux.Handle("GET /kaithhealth", handler.Healthz()) // Leapcell healthcheck

	// products
	productRepo := repository.NewProductRepo()
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)
	mux.Handle("GET /api/products", productHandler.ListProducts())
	mux.Handle("POST /api/products", productHandler.CreateProduct())
	mux.Handle("GET /api/products/{id}", productHandler.GetProductById())
	mux.Handle("DELETE /api/products/{id}", productHandler.DeleteProductById())
	mux.Handle("PUT /api/products/{id}", productHandler.UpdateProductById())

	// categories
	mux.Handle("GET /api/categories", handler.ListCategories())
	mux.Handle("POST /api/categories", handler.CreateCategory())
	mux.Handle("GET /api/categories/{id}", handler.GetCategoryById())
	mux.Handle("DELETE /api/categories/{id}", handler.DeleteCategoryById())
	mux.Handle("PUT /api/categories/{id}", handler.UpdateCategoryById())

	// default route
	mux.Handle("/", handler.DefaultHandler())

	log.Println("running server on port 8080..")
	err := http.ListenAndServe(":8080", middleware.SimpleLogger(middleware.DefaultHeaders(mux)))
	if err != nil {
		log.Println("failed to run server:", err)
	}
}
