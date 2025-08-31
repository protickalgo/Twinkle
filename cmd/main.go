package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"twinkle/domain"
	"twinkle/handler"
	"twinkle/repository/mysqlrepo"
	"twinkle/service"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	user := "dbuser"
	password := "dbpass"
	host := "your-db-host:3306"
	dbName := "demo"


	// Connect without database first
	sqlDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", user, password, host))
	if err != nil {
		log.Fatal("Cannot connect to MySQL:", err)
	}
	defer sqlDB.Close()

	// Create database if not exists
	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal("Cannot create database:", err)
	}
	log.Println("Database checked/created successfully.")

	// Connect with GORM
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Auto migrate
	db.AutoMigrate(&domain.Product{})
	log.Println("Tables checked/created successfully.")

	// Repository, Service, Handler
	productRepo := mysqlrepo.NewMySQLProductRepo(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")

	// Start server
	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
