# Book Store API

Test project to learn Golang for backend development. Simple Book store API, not production-ready.

Basic hexagonal architecture, SQLite for persistence, no auth.

## Execution

```bash
go run .
```

## Endpoints

`/books`
 - `GET`    Get all stored books
 - `POST`   Create a new book


`/books/{id}`
 - `GET`    Get a stored book
 - `PUT`    Update a stored book
 - `DELETE` Delete a stored book
