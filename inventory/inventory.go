package inventory

import (
	"errors"

	"books.com/models"
)

type Inventory interface {
	AddBook(title string, stock int) error
	ListAllBooks() []models.Book
	GetBookByTitle(title string) *models.Book
	EditBook(title, newTitle string, stock int) error
}

type InventoryImpl struct {
	Books map[string]int
}

func NewInventory() Inventory {
	return InventoryImpl{
		Books: make(map[string]int),
	}
}

func (i InventoryImpl) AddBook(title string, stock int) error {
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	if _, ok := i.Books[title]; ok {
		i.Books[title] += stock

		return nil
	}

	i.Books[title] = stock

	return nil
}

func (i InventoryImpl) ListAllBooks() []models.Book {
	books := []models.Book{}

	for key, val := range i.Books {
		books = append(books, *i.generateBook(key, val))
	}

	return books
}

func (i InventoryImpl) GetBookByTitle(title string) *models.Book {
	if _, ok := i.Books[title]; ok {
		return i.generateBook(title, i.Books[title])
	}

	return nil
}

func (i InventoryImpl) EditBook(title, newTitle string, stock int) error {
	if _, ok := i.Books[title]; !ok {
		return errors.New("book not found with that title")
	}

	if stock < 0 {
		stock *= -1
		if i.Books[title] < stock {
			return errors.New("not enough books in the inventory")
		}

		i.Books[title] -= stock
	} else {
		i.Books[title] += stock

		if title != newTitle {
			currentStock := i.Books[title]

			delete(i.Books, title)

			i.Books[newTitle] = currentStock
		}
	}

	return nil
}

func (i InventoryImpl) generateBook(title string, stock int) *models.Book {
	return &models.Book{
		Title: title,
		Stock: stock,
	}
}
