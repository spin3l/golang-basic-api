package infrastructure

import (
	"database/sql"
	"test-backend/book/application"
	"test-backend/book/domain"
)

type store struct {
	db *sql.DB
}

func (s *store) GetAll() ([]*domain.Book, error) {
	q := `SELECT id, title, author FROM books`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []*domain.Book{}
	for rows.Next() {
		b := domain.Book{}
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}

		books = append(books, &b)
	}

	return books, nil
}

func (s *store) GetById(id int) (*domain.Book, error) {
	q := `SELECT id, title, author FROM books WHERE id = ?`

	b := domain.Book{}
	err := s.db.QueryRow(q, id).Scan(&b.ID, &b.Title, &b.Author)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (s *store) Create(book *domain.Book) (*domain.Book, error) {
	q := `INSERT INTO books (title, author) VALUES (?, ?)`

	resp, err := s.db.Exec(q, book.Title, book.Author)
	if err != nil {
		return nil, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}

	book.ID = int(id)
	return book, nil
}

func (s *store) Update(id int, book *domain.Book) (*domain.Book, error) {
	q := `UPDATE books SET title = ?, author = ? WHERE id = ?`

	_, err := s.db.Exec(q, book.Title, book.Author, id)
	if err != nil {
		return nil, err
	}

	book.ID = id

	return book, nil
}

func (s *store) Delete(id int) error {
	q := `DELETE FROM books WHERE id = ?`

	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}

func NewBookStore(db *sql.DB) application.BookStore {
	return &store{db: db}
}
