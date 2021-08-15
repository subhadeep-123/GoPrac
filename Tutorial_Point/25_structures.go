package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func PrintBook(book Books) {
	fmt.Printf("Book 1 title : %s\n", book.title)
	fmt.Printf("Book 1 author : %s\n", book.author)
	fmt.Printf("Book 1 subject : %s\n", book.subject)
	fmt.Printf("Book 1 book_id : %d\n", book.book_id)
}
func main() {

	var Book1 Books /* Declare Book1 of type Book */
	var Book2 Books /* Declare Book2 of type Book */

	Book1 = Books{title: "And Then There Were None",
		author:  "Agatha Christie",
		subject: "Fiction Story",
		book_id: 1}

	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	PrintBook(Book1)
	PrintBook(Book2)
}
