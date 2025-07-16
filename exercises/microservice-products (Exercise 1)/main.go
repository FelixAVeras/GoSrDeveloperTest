package main

import (
	"log"
	"net/http"

	"microservice-products/handlers"
	"microservice-products/services"

	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Iniciando el microservicio de productos...")

	productService := services.NewProductService()
	productHandler := handlers.NewProductHandler(productService)

	router := mux.NewRouter()

	router.HandleFunc("/products", productHandler.CreateProductHandler).Methods("POST")
	router.HandleFunc("/products", productHandler.GetAllProductsHandler).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.GetProductByIDHandler).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.UpdateProductHandler).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteProductHandler).Methods("DELETE")

	port := ":8080"
	log.Printf("Servidor escuchando en el puerto %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
