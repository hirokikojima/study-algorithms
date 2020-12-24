package iterator

import (
	"errors"
)

type Iterator interface {
	HasNext() bool
	Next() (Book, error)
}

type bookshelfIterator  struct {
	bookshelf Bookshelf
	current int
}

func NewBookshelfIterator(bookshelf Bookshelf) Iterator {
	return &bookshelfIterator{
		bookshelf: bookshelf,
		current: 0,
	}
}

func (bookshelf_iterator *bookshelfIterator) HasNext() bool {
	return bookshelf_iterator.current < bookshelf_iterator.bookshelf.Last()
}

func (bookshelf_iterator *bookshelfIterator) Next() (Book, error) {
	book, err := bookshelf_iterator.bookshelf.GetBookAt(bookshelf_iterator.current)
	if (err != nil) {
		return nil, errors.New("Illegal index specified.")
	}
	
	bookshelf_iterator.current++
	return book, nil
}