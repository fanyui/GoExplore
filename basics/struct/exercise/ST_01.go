package main

import "fmt"

type Rectangle struct {
 Length float64
 Width  float64
}
//Write a function that takes a struct representing a Rectangle (with fields for length and width) as input and calculates its area.
func calculateArea(rectangle Rectangle) float64 {
 return rectangle.Length * rectangle.Width
}

func main() {
 // Example usage
 rectangle := Rectangle{
  Length: 5.0,
  Width:  3.0,
 }

 area := calculateArea(rectangle)
 fmt.Println("Area:", area)
}