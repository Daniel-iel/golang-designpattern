package main

import "fmt"

type BookType int

const (
	HardeCover BookType = iota
	SoftCover
	PaperBack
	Ebook
)

type Book struct {
	Name      string
	Author    string
	pageCount int
	Type      BookType
}

type Libraty struct {
	Collection []Book
}

func (l *Library) InterateBook(f func(Book) error) {
	var err error
	for _, b := range l.Collection {
		err = f(b)
		if err != nil {
			fmt.Println("Error encountered:", err)
			break
		}
	}
}

var lib *Library = &Library{
	Collection: []Book{
		{
			Name:      "War and Peace",
			Author:    "Leo Tolstoy",
			pageCount: 864,
			Type:      HardeCover,
		},
		{
			Name:      "Crime and Punishment",
			Author:    "Leo Tolstoy",
			pageCount: 1225,
			Type:      SoftCover,
		},
		{
			Name:      "Brave New World",
			Author:    "Aldous Huxley",
			pageCount: 325,
			Type:      PaperBack,
		},
		{
			Name:      "Catcher in the Rye",
			Author:    "J.D Salinger",
			pageCount: 206,
			Type:      HardeCover,
		},
	},
}
