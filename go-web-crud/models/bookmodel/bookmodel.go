package bookmodel

import (
	"go-web-crud/config"
	"go-web-crud/entities"
)

func Getall() []entities.Book {
	rows, err := config.DB.Query(`
		SELECT 
			books.id, 
			books.name, 
			categories.name as category_name,
			books.stock, 
			books.description, 
			books.created_at, 
			books.updated_at 
		FROM books
		JOIN categories ON books.category_id = categories.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var books []entities.Book

	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(
			&book.Id,
			&book.Name,
			&book.Category.Name,
			&book.Stock,
			&book.Description,
			&book.CreatedAt,
			&book.UpdatedAt,
		); err != nil {
			panic(err)
		}

		books = append(books, book)
	}

	return books
}

func Create(book entities.Book) bool {
	result, err := config.DB.Exec(`
		INSERT INTO books(
			name, category_id, stock, description, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?)`,
		book.Name,
		book.Category.Id,
		book.Stock,
		book.Description,
		book.CreatedAt,
		book.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Book {
	row := config.DB.QueryRow(`
		SELECT 
			books.id, 
			books.name, 
			categories.name as category_name,
			books.stock, 
			books.description, 
			books.created_at, 
			books.updated_at FROM books
		JOIN categories ON books.category_id = categories.id
		WHERE books.id = ?
	`, id)

	var book entities.Book

	err := row.Scan(
		&book.Id,
		&book.Name,
		&book.Category.Name,
		&book.Stock,
		&book.Description,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return book
}

func Update(id int, book entities.Book) bool {
	query, err := config.DB.Exec(`
		UPDATE books SET 
			name = ?, 
			category_id = ?,
			stock = ?,
			description = ?,
			updated_at = ?
		WHERE id = ?`,
		book.Name,
		book.Category.Id,
		book.Stock,
		book.Description,
		book.UpdatedAt,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
