package handler

import (
    "encoding/json"
    "net/http"
    "twinkle/domain"
    "twinkle/service"
)

type ProductHandler struct {
    Service service.ProductServiceInterface
}

func NewProductHandler(s service.ProductServiceInterface) *ProductHandler {
    return &ProductHandler{Service: s}
}

// GetProducts handles GET /products
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
    products, err := h.Service.GetProducts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(products); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// CreateProduct handles POST /products
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product domain.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, "invalid request body", http.StatusBadRequest)
        return
    }

    if err := h.Service.CreateProduct(&product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated) // 201 Created
    if err := json.NewEncoder(w).Encode(product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
