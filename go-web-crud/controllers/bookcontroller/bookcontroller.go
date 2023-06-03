package bookcontroller

import (
	"go-web-crud/entities"
	"go-web-crud/models/bookmodel"
	"go-web-crud/models/categorymodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	books := bookmodel.Getall()
	data := map[string]any{
		"books": books,
	}

	temp, err := template.ParseFiles("views/book/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/book/create.html")
		if err != nil {
			panic(err)
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var book entities.Book

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		book.Name = r.FormValue("name")
		book.Category.Id = uint(categoryId)
		book.Stock = int64(stock)
		book.Description = r.FormValue("description")
		book.CreatedAt = time.Now()
		book.UpdatedAt = time.Now()

		if ok := bookmodel.Create(book); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	book := bookmodel.Detail(id)
	data := map[string]any{
		"book": book,
	}

	temp, err := template.ParseFiles("views/book/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/book/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		book := bookmodel.Detail(id)
		categories := categorymodel.GetAll()

		data := map[string]any{
			"book":       book,
			"categories": categories,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var book entities.Book

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		book.Name = r.FormValue("name")
		book.Category.Id = uint(categoryId)
		book.Stock = int64(stock)
		book.Description = r.FormValue("description")
		book.UpdatedAt = time.Now()

		if ok := bookmodel.Update(id, book); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := bookmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
