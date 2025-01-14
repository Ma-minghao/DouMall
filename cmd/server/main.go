package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/go-ecommerce/internal/auth"
	"github.com/yourusername/go-ecommerce/internal/cart"
	"github.com/yourusername/go-ecommerce/internal/order"
	"github.com/yourusername/go-ecommerce/internal/payment"
	"github.com/yourusername/go-ecommerce/internal/product"
	"github.com/yourusername/go-ecommerce/internal/user"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	// API Version 1 Routes
	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Authentication routes
	v1.HandleFunc("/auth/login", auth.LoginHandler).Methods("POST")
	v1.HandleFunc("/auth/logout", auth.LogoutHandler).Methods("POST")
	v1.HandleFunc("/auth/validate", auth.ValidateTokenHandler).Methods("GET")

	// User routes
	v1.HandleFunc("/users", user.CreateUserHandler).Methods("POST")
	v1.HandleFunc("/users/{id}", user.GetUserHandler).Methods("GET")
	v1.HandleFunc("/users/{id}", user.UpdateUserHandler).Methods("PUT")
	v1.HandleFunc("/users/{id}", user.DeleteUserHandler).Methods("DELETE")

	// Product routes
	v1.HandleFunc("/products", product.CreateProductHandler).Methods("POST")
	v1.HandleFunc("/products/{id}", product.GetProductHandler).Methods("GET")
	v1.HandleFunc("/products/{id}", product.UpdateProductHandler).Methods("PUT")
	v1.HandleFunc("/products/{id}", product.DeleteProductHandler).Methods("DELETE")

	// Cart routes
	v1.HandleFunc("/cart", cart.CreateCartHandler).Methods("POST")
	v1.HandleFunc("/cart/{id}", cart.GetCartHandler).Methods("GET")
	v1.HandleFunc("/cart/{id}", cart.ClearCartHandler).Methods("DELETE")

	// Order routes
	v1.HandleFunc("/orders", order.CreateOrderHandler).Methods("POST")
	v1.HandleFunc("/orders/{id}", order.GetOrderHandler).Methods("GET")
	v1.HandleFunc("/orders/{id}", order.CancelOrderHandler).Methods("DELETE")

	// Payment routes
	v1.HandleFunc("/payments", payment.ProcessPaymentHandler).Methods("POST")
	v1.HandleFunc("/payments/{id}", payment.CancelPaymentHandler).Methods("DELETE")

	// Start the server
	port := ":8080"
	fmt.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
