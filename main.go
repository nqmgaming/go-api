package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	products = append(products, p)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedProduct Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, p := range products {
		if p.ID == updatedProduct.ID {
			products[i] = updatedProduct
			if err := json.NewEncoder(w).Encode(updatedProduct); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	products = append(products, Product{ID: "1", Name: "Product 1", Price: 100.0})
	products = append(products, Product{ID: "2", Name: "Product 2", Price: 200.0})

	router := http.NewServeMux()
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProducts(w, r)
		case http.MethodPost:
			createProduct(w, r)
		case http.MethodPut:
			updateProduct(w, r)
		case http.MethodDelete:
			deleteProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
