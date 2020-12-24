package main

import (
	"fmt"
	"github.com/hirokikojima/study-algorithms/design_pattern/iterator"
)

func main() {
	bookshelf := iterator.NewBookshelf(10)

	bookshelf.AppendBook(iterator.NewBook("Hello"))
	bookshelf.AppendBook(iterator.NewBook("Iterator"))
	bookshelf.AppendBook(iterator.NewBook("Pattern"))
	
	iterator := bookshelf.Iterator()

	for iterator.HasNext() {
		book, err := iterator.Next()
		if (err != nil) {
			continue
		}

		fmt.Printf("%s\n", book.GetName())
	}
}