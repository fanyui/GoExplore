package main
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Square struct {
	side float64
}
func (s Square) Area() float64 {
	return s.side * s.side
}
func (s Square) Perimeter() float64 {
	return 4 * s.side
}

func main() {
	s := Square{side: 5}
	var area, perimeter float64 = s.Area(), s.Perimeter()
	println("Area:", area, "Perimeter:", perimeter)
}