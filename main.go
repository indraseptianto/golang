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

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc("/categories", categoryHandler.HandleCategories)

	fmt.Println("🚀 Server running on :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
