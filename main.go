package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "twinkle/domain"
    "twinkle/handler"
    "twinkle/repository"
    "twinkle/service"

    _ "github.com/go-sql-driver/mysql"
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

    // Step 1: Connect to MySQL without specifying a database
    sqlDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", user, password, host))
    if err != nil {
        log.Fatal("Cannot connect to MySQL:", err)
    }
    defer sqlDB.Close()

    // Step 2: Create database if not exists
    _, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
    if err != nil {
        log.Fatal("Cannot create database:", err)
    }
    log.Println("Database checked/created successfully.")

    // Step 3: Connect to the newly created database using GORM
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }

    // Step 4: Auto-migrate Product table
    db.AutoMigrate(&domain.Product{})
    log.Println("Tables checked/created successfully.")

    // Step 5: Repository, Service, Handler
    productRepo := repository.NewMySQLProductRepo(db)
    productService := service.NewProductService(productRepo)
    productHandler := handler.NewProductHandler(productService)

    // Step 6: Router and endpoints
    r := mux.NewRouter()
    r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
    r.HandleFunc("/products/{id}", productHandler.GetProductByID).Methods("GET")
    r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
    r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
    r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

    // Step 7: Start server
    log.Printf("Server running on http://0.0.0.0:%s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
