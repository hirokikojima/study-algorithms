package iterator

type Book interface {
	GetName() string
}

type book struct {
	Name string
}

func NewBook(name string) Book {
	return &book{
		Name: name,
	}
}

func (book *book) GetName() string {
	return book.Name
}