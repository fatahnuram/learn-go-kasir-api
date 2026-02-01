package main

import (
	"log"
	"net/http"

	"github.com/fatahnuram/learn-go-kasir-api/internal/config"
	"github.com/fatahnuram/learn-go-kasir-api/internal/db"
	"github.com/fatahnuram/learn-go-kasir-api/internal/handler"
	"github.com/fatahnuram/learn-go-kasir-api/internal/middleware"
	"github.com/fatahnuram/learn-go-kasir-api/internal/repository"
	"github.com/fatahnuram/learn-go-kasir-api/internal/service"
)

func main() {
	conf := config.Init()
	db, err := db.InitDB(conf.DBConn)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	// healthcheck
	mux.Handle("GET /healthz", handler.Healthz())
	mux.Handle("GET /kaithhealth", handler.Healthz()) // Leapcell healthcheck

	// products
	productRepo := repository.NewProductRepo(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)
	mux.Handle("GET /api/products", productHandler.ListProducts())
	mux.Handle("POST /api/products", productHandler.CreateProduct())
	mux.Handle("GET /api/products/{id}", productHandler.GetProductById())
	mux.Handle("DELETE /api/products/{id}", productHandler.DeleteProductById())
	mux.Handle("PUT /api/products/{id}", productHandler.UpdateProductById())

	// categories
	categoryRepo := repository.NewCategoryRepo(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	mux.Handle("GET /api/categories", categoryHandler.ListCategories())
	mux.Handle("POST /api/categories", categoryHandler.CreateCategory())
	mux.Handle("GET /api/categories/{id}", categoryHandler.GetCategoryById())
	mux.Handle("DELETE /api/categories/{id}", categoryHandler.DeleteCategoryById())
	mux.Handle("PUT /api/categories/{id}", categoryHandler.UpdateCategoryById())

	// default route
	mux.Handle("/", handler.DefaultHandler())

	addr := "0.0.0.0:" + conf.Port
	log.Println("running server on", addr)
	err = http.ListenAndServe(addr, middleware.SimpleLogger(middleware.DefaultHeaders(mux)))
	if err != nil {
		log.Println("failed to run server:", err)
	}
}
