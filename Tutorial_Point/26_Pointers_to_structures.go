package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go Programming"
	Book1.author = "Subhadeep Banerjee"
	Book1.subject = "Structures"
	Book1.book_id = 100

	Book2.title = "Python Programming"
	Book2.author = "Subhadeep Banerjee"
	Book2.subject = "Pep8"
	Book2.book_id = 1000

	PrintBook(&Book1)
	PrintBook(&Book2)
}

func PrintBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
