package main

import (
	"go-web-crud/config"
	"go-web-crud/controllers/bookcontroller"
	"go-web-crud/controllers/categorycontroller"
	"go-web-crud/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// 3. Books
	http.HandleFunc("/books", bookcontroller.Index)
	http.HandleFunc("/books/add", bookcontroller.Add)
	http.HandleFunc("/books/detail", bookcontroller.Detail)
	http.HandleFunc("/books/edit", bookcontroller.Edit)
	http.HandleFunc("/books/delete", bookcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
