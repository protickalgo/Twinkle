package handler

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "twinkle/domain"
    "twinkle/service"
)

type ProductHandler struct {
    service service.ProductServiceInterface
}

func NewProductHandler(s service.ProductServiceInterface) *ProductHandler {
    return &ProductHandler{service: s}
}

// GET /products
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
    products, err := h.service.GetAllProducts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

// GET /products/{id}
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idInt, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    id := uint(idInt)

    product, err := h.service.GetProductByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}

// POST /products
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product domain.Product
    json.NewDecoder(r.Body).Decode(&product)

    err := h.service.CreateProduct(&product)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}

// PUT /products/{id}
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idInt, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    id := uint(idInt)

    var product domain.Product
    json.NewDecoder(r.Body).Decode(&product)
    product.ID = id

    err = h.service.UpdateProduct(&product)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}

// DELETE /products/{id}
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idInt, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    id := uint(idInt)

    err = h.service.DeleteProduct(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
