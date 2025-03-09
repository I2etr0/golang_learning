package main

import "fmt"

type Book struct {
	Author string
	Name   string
	Pages  int
}

func main() {
	books := []Book{
		{Author: "Aboba", Name: "Chupapi", Pages: 100},
		{Author: "pepega", Name: "Pepelogia", Pages: 1},
		{Author: "mem", Name: "Memelogia", Pages: 10},
	}

	fmt.Println(books)

	golang := Book{
		Author: "Createor",
		Name:   "Golang",
		Pages:  1,
	}

	books = append(books, golang)

	for index, book := range books {
		fmt.Println(index, book.Author, book.Name, book.Pages)
	}

}
