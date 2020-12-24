package iterator

type Aggregate interface {
	Iterator() Iterator
}

type Bookshelf interface {
	GetBookAt(index int) (Book, error)
	AppendBook(book Book)
	Last() int
}

type bookshelf struct {
	books []Book
}

func NewBookshelf(maximum int) *bookshelf {
	return &bookshelf{}
}

func (bookshelf *bookshelf) GetBookAt(index int) (Book, error) {
	return bookshelf.books[index], nil
}

func (bookshelf *bookshelf) AppendBook(book Book) {
	bookshelf.books = append(bookshelf.books, book)
}

func (bookshelf *bookshelf) Last() int {
	return len(bookshelf.books)
}

func (bookshelf *bookshelf) Iterator() Iterator {
	return NewBookshelfIterator(bookshelf)
}