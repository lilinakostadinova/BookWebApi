package models

import (
	"BookWebApi/db"
	"fmt"
)

type Book struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	ISBN   string `json:"isbn"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func (b *Book) Save() error {
	statement, err := db.GetDb().Prepare(`
		INSERT INTO 
		books
			(title, isbn, author, year)
		VALUES
			(?, ?, ?, ?)
	`)
	defer statement.Close()

	if err != nil {
		return err
	}

	result, err := statement.Exec(b.Title, b.ISBN, b.Author, b.Year)
	if err != nil {
		return err
	}

	bookId, err := result.LastInsertId()
	b.Id = bookId

	return err
}

func GetAllBooks() ([]Book, error) {
	dbCursor, err := db.GetDb().Query(`SELECT * FROM books`)
	if err != nil {
		return nil, err
	}

	var bookCollection []Book
	for dbCursor.Next() {
		var bookObject Book
		err := dbCursor.Scan(
			&bookObject.Id,
			&bookObject.Title,
			&bookObject.ISBN,
			&bookObject.Author,
			&bookObject.Year,
		)

		if err != nil {
			return nil, err
		}

		bookCollection = append(bookCollection, bookObject)
	}

	return bookCollection, nil
}

func GetBooks(limit int) ([]Book, error) {
	query := `SELECT * FROM books`
	if limit > 0 {
		query = fmt.Sprintf(`%s LIMIT %d`, query, limit)
	}

	dbCursor, err := db.GetDb().Query(query)
	if err != nil {
		return nil, err
	}

	var books []Book
	for dbCursor.Next() {
		var bookObject Book
		err := dbCursor.Scan(
			&bookObject.Id,
			&bookObject.Title,
			&bookObject.ISBN,
			&bookObject.Author,
			&bookObject.Year,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, bookObject)
	}

	return books, nil
}

func GetBookById(id int64) (Book, error) {
	var book Book
	err := db.GetDb().QueryRow(`SELECT * FROM books WHERE id = ?`, id).Scan(
		&book.Id,
		&book.Title,
		&book.ISBN,
		&book.Author,
		&book.Year,
	)
	return book, err
}

func (b *Book) Update() error {
	statement, err := db.GetDb().Prepare(`
		UPDATE books SET title = ?, isbn = ?, author = ?, year = ? WHERE id = ?
	`)
	defer statement.Close()

	if err != nil {
		return err
	}

	_, err = statement.Exec(b.Title, b.ISBN, b.Author, b.Year, b.Id)
	return err
}

func DeleteBookById(id int64) error {
	statement, err := db.GetDb().Prepare(`DELETE FROM books WHERE id = ?`)
	defer statement.Close()

	if err != nil {
		return err
	}

	_, err = statement.Exec(id)
	return err
}
