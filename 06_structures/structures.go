package main

import "fmt"

type Book struct {
  title string
  author string
  id int
}

func main() {
  var Book1 Book
  var Book2 Book

  Book1.title = "Harry Potter"
  Book1.author = "J.K.Rowling"
  Book1.id = 1

  Book2.title = "Lord of the Rings"
  Book2.author = "J.R.R. Tolkien"
  Book2.id = 2

  fmt.Println("Structures\n")

  fmt.Printf("%s, %s\n", Book1.title, Book1.author)
  fmt.Printf("%s, %s\n", Book2.title, Book2.author)

  // passing as argument to the function
  printBook(Book1);
  printBook(Book2);

  // pointer to a Book object
  var structPointer *Book

  structPointer = &Book1
  structPointer.title = "X-Men"
  structPointer.author = "Various"

  structPointer = &Book2
  structPointer.title = "Hobbit"

  printBook(Book1)
  printBook(Book2)
}

func printBook(b Book) {
  fmt.Printf("%s, %s\n", b.title, b.author);
}
