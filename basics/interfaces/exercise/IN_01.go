package main

import (
	"fmt"
	"sort"
)

type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
  }
type Person struct {
	Name string
	Age int
}
type Persons []Person

func (p Persons) Len() int {return len(p)}
func (p Persons) Less(i, j int) bool {return p[i].Name < p[j].Name}
func (p Persons) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
func main()  {


	people := []Person{
		{Age: 30, Name: "Alice"},
		{Age: 25, Name: "Bob"},
		{Age: 35, Name: "Charlie"},
	   }
	sort.Sort(Persons(people))
	for _, person := range people {
	  fmt.Println(person.Name, person.Age)
	}
}

