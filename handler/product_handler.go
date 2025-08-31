package handler

import (
    "encoding/json"
    "net/http"
    // "strconv"
    // "github.com/gorilla/mux"
    "twinkle/domain"
    "twinkle/service"
)

type ProductHandler struct {
    Service service.ProductServiceInterface
}

func NewProductHandler(s service.ProductServiceInterface) *ProductHandler {
    return &ProductHandler{Service: s}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
    products, err := h.Service.GetProducts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product domain.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.Service.CreateProduct(&product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}
