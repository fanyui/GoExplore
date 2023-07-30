//Create a struct to represent a Book, with fields for title, author, and price.
//Write a function that takes a slice of Book and returns the total price of all the books.

package main

import "fmt"

type Book struct {
 Title  string
 Author string
 Price  float64
}

func calculateTotalPrice(books []Book) float64 {
 totalPrice := 0.0

 for _, book := range books {
  totalPrice += book.Price
 }

 return totalPrice
}

func main() {
 // Example usage
 books := []Book{
  {
   Title:  "The Great Gatsby",
   Author: "F. Scott Fitzgerald",
   Price:  10.99,
  },
  {
   Title:  "To Kill a Mockingbird",
   Author: "Harper Lee",
   Price:  12.99,
  },
  {
   Title:  "1984",
   Author: "George Orwell",
   Price:  9.99,
  },
 }

 totalPrice := calculateTotalPrice(books)
 fmt.Println("Total Price:", totalPrice)
}