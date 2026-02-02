package main

import (
	"fmt"
	"log"
	"net/http"

	"kasir-api/config"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
)

func main() {
	cfg := config.Load()

	fmt.Println("PORT =", cfg.Port)
	fmt.Println("DB_CONN =", cfg.DBConn)

	db, err := database.Init(cfg.DBConn)
	if err != nil {
		log.Fatal("DB error:", err)
	}
	defer db.Close()

<<<<<<< HEAD
=======
	// Category setup
>>>>>>> 5c36bd9 (perbaikan)
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

<<<<<<< HEAD
	http.HandleFunc("/categories", categoryHandler.HandleCategories)

	fmt.Println("🚀 Server running on :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
=======
	// Product setup
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/categories", categoryHandler.HandleCategories)
	mux.HandleFunc("/products/", productHandler.HandleProductDetail)

	fmt.Println("🚀 Server running on :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
>>>>>>> 5c36bd9 (perbaikan)
}
productRepo := repositories.NewProductRepository(db)
productService := services.NewProductService(productRepo)
productHandler := handlers.NewProductHandler(productService)

http.HandleFunc("/products", productHandler.HandleProducts)
http.HandleFunc("/products/", productHandler.HandleProductByID)
