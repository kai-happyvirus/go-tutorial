package main

	
import (
	"fmt"
	"math"
)
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 39
	return &p
}

func structExample() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)
	fmt.Println(s.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

func main() {
	structExample()
	methods()
}

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
// Methods can be defined for either pointer or value receiver types.
// Here’s an example of a value receiver.
func (r *rect) area() int {

	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// Go automatically handles conversion between values and pointers for method calls.
// You may want to use a pointer receiver type to avoid copying on method calls or
// to allow the method to mutate the receiving struct.
func methods() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

