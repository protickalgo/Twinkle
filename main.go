package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "twinkle/domain"
    "twinkle/handler"
    "twinkle/repository"
    "twinkle/service"
    "github.com/gorilla/mux"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    // Read environment variables
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASS")
    host := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")
    port := os.Getenv("PORT")

    if port == "" {
        port = "8080"
    }

    // Connect to PostgreSQL
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, user, password, dbName)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }

    // Auto-migrate Product table
    db.AutoMigrate(&domain.Product{})

    // Repository, Service, Handler
    productRepo := repository.NewProductRepo(db)
    productService := service.NewProductService(productRepo)
    productHandler := handler.NewProductHandler(productService)

    // Router and endpoints
    r := mux.NewRouter()
    r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
    r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")

    log.Println("Server running on port " + port)
    http.ListenAndServe(":"+port, r)
}
