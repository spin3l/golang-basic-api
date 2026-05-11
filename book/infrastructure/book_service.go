package infrastructure

import (
	"test-backend/book/application"
	"test-backend/book/domain"
)

type service struct {
	store application.BookStore
}

func NewBookService(s application.BookStore) *service {
	return &service{store: s}
}

func (s *service) GetAll() ([]*domain.Book, error) {
	return s.store.GetAll()
}

func (s *service) GetById(id int) (*domain.Book, error) {
	return s.store.GetById(id)
}

func (s *service) Create(book *domain.Book) (*domain.Book, error) {
	return s.store.Create(book)
}

func (s *service) Update(id int, book *domain.Book) (*domain.Book, error) {
	return s.store.Update(id, book)
}

func (s *service) Delete(id int) error {
	return s.store.Delete(id)
}
