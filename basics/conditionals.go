package main
var color string = "blue"
var age int = 20
func main()  {
	if color == "red"  {
		println("Color is red")
	} else if color == "blue" && age >= 18 {
		println("Color is blue and age is greater than or equal to 18")

	}else {
		println("Oops None of the above is true")
	}
}