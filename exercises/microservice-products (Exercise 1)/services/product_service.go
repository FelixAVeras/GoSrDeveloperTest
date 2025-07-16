package services

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"microservice-products/models"
)

// ProductService maneja la lógica de negocio y el almacenamiento de productos.
type ProductService struct {
	products map[string]models.Product // Almacenamiento en memoria
	mu       sync.RWMutex              // Mutex para concurrencia segura
	nextID   int                       // Para generar IDs únicos
}

// NewProductService crea una nueva instancia de ProductService.
func NewProductService() *ProductService {
	return &ProductService{
		products: make(map[string]models.Product),
		nextID:   1,
	}
}

// CreateProduct agrega un nuevo producto al almacenamiento.
func (s *ProductService) CreateProduct(product models.Product) (models.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validación básica
	if product.Name == "" {
		return models.Product{}, fmt.Errorf("el nombre del producto no puede estar vacío")
	}
	if product.Price <= 0 {
		return models.Product{}, fmt.Errorf("el precio del producto debe ser mayor que cero")
	}
	if product.Quantity < 0 {
		return models.Product{}, fmt.Errorf("la cantidad del producto no puede ser negativa")
	}

	product.ID = strconv.Itoa(s.nextID)
	s.nextID++
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	s.products[product.ID] = product
	log.Printf("Producto creado: %+v", product)
	return product, nil
}

// GetAllProducts devuelve todos los productos almacenados.
func (s *ProductService) GetAllProducts() []models.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	products := make([]models.Product, 0, len(s.products))
	for _, p := range s.products {
		products = append(products, p)
	}
	log.Println("Todos los productos recuperados")
	return products
}

// GetProductByID devuelve un producto por su ID.
func (s *ProductService) GetProductByID(id string) (models.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	product, ok := s.products[id]
	if !ok {
		return models.Product{}, fmt.Errorf("producto con ID '%s' no encontrado", id)
	}
	log.Printf("Producto recuperado por ID: %s", id)
	return product, nil
}

// UpdateProduct actualiza un producto existente.
func (s *ProductService) UpdateProduct(id string, updatedProduct models.Product) (models.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	existingProduct, ok := s.products[id]
	if !ok {
		return models.Product{}, fmt.Errorf("producto con ID '%s' no encontrado para actualizar", id)
	}

	// Validación básica para la actualización
	if updatedProduct.Name != "" {
		existingProduct.Name = updatedProduct.Name
	}
	if updatedProduct.Description != "" {
		existingProduct.Description = updatedProduct.Description
	}
	if updatedProduct.Price > 0 {
		existingProduct.Price = updatedProduct.Price
	}
	if updatedProduct.Quantity >= 0 {
		existingProduct.Quantity = updatedProduct.Quantity
	}

	existingProduct.UpdatedAt = time.Now()
	s.products[id] = existingProduct
	log.Printf("Producto actualizado: %+v", existingProduct)
	return existingProduct, nil
}

// DeleteProduct elimina un producto por su ID.
func (s *ProductService) DeleteProduct(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.products[id]
	if !ok {
		return fmt.Errorf("producto con ID '%s' no encontrado para eliminar", id)
	}
	delete(s.products, id)
	log.Printf("Producto eliminado con ID: %s", id)
	return nil
}
