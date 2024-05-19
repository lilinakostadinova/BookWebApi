This is a simple web API for managing books using Go and SQLite.

## Endpoints

- **GET /books**: Returns a collection of all books in the system. Optionally, you can limit the number of books returned by using the `limit` query parameter.
- **GET /books/:id**: Returns data for a single book.
- **POST /books**: Creates a new book.
- **PUT /books/:id**: Updates information for an existing book.
- **DELETE /books/:id**: Deletes a book by its identifier.
