package application

import "golang-basic-api/book/domain"

type BookService interface {
	GetAll() ([]*domain.Book, error)
	GetById(id int) (*domain.Book, error)
	Create(book *domain.Book) (*domain.Book, error)
	Update(id int, book *domain.Book) (*domain.Book, error)
	Delete(id int) error
}
