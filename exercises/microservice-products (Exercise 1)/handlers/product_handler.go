package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"microservice-products/models"
	"microservice-products/services"

	"github.com/gorilla/mux"
)

// ProductHandler maneja las solicitudes HTTP relacionadas con productos.
type ProductHandler struct {
	service *services.ProductService
}

// NewProductHandler crea una nueva instancia de ProductHandler.
func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// CreateProductHandler maneja la creación de un nuevo producto (POST /products).
func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Solicitud inválida: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error al decodificar la solicitud de creación de producto: %v", err)
		return
	}

	createdProduct, err := h.service.CreateProduct(product)
	if err != nil {
		http.Error(w, "Error al crear producto: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error en el servicio al crear producto: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
	log.Printf("Respuesta de creación de producto enviada: %+v", createdProduct)
}

// GetAllProductsHandler maneja la obtención de todos los productos (GET /products).
func (h *ProductHandler) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := h.service.GetAllProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
	log.Println("Respuesta de todos los productos enviada")
}

// GetProductByIDHandler maneja la obtención de un producto por ID (GET /products/{id}).
func (h *ProductHandler) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	product, err := h.service.GetProductByID(id)
	if err != nil {
		http.Error(w, "Producto no encontrado: "+err.Error(), http.StatusNotFound)
		log.Printf("Error al obtener producto por ID '%s': %v", id, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	log.Printf("Respuesta de producto por ID '%s' enviada: %+v", id, product)
}

// UpdateProductHandler maneja la actualización de un producto (PUT /products/{id}).
func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Solicitud inválida: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error al decodificar la solicitud de actualización de producto para ID '%s': %v", id, err)
		return
	}

	updatedProduct, err := h.service.UpdateProduct(id, product)
	if err != nil {
		http.Error(w, "Error al actualizar producto: "+err.Error(), http.StatusNotFound) // O 500 si es un error interno
		log.Printf("Error en el servicio al actualizar producto con ID '%s': %v", id, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProduct)
	log.Printf("Respuesta de actualización de producto enviada para ID '%s': %+v", id, updatedProduct)
}

// DeleteProductHandler maneja la eliminación de un producto (DELETE /products/{id}).
func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Error al eliminar producto: "+err.Error(), http.StatusNotFound)
		log.Printf("Error en el servicio al eliminar producto con ID '%s': %v", id, err)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content para eliminación exitosa
	log.Printf("Producto con ID '%s' eliminado exitosamente", id)
}
