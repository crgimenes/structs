package main

import (
	"fmt"

	"crg.eti.br/go/structs"
)

func main() {
	// struct to map

	type Book struct {
		Title  string
		Author string
		Price  float64
	}

	type User struct {
		APIKey []byte
		Name   string
		Age    int
		Books  []Book
	}

	user := User{
		APIKey: []byte("1234567890"),
		Name:   "John Doe",
		Age:    30,
		Books: []Book{
			{
				Title:  "The Lord of the Rings",
				Author: "J. R. R. Tolkien",
				Price:  39.99,
			},
			{
				Title:  "The Hobbit",
				Author: "J. R. R. Tolkien",
				Price:  29.99,
			},
		},
	}

	m := structs.Map(user)

	fmt.Printf("struct to map: %#v\n", m)

}
