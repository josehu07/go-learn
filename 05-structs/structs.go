package main


import (
	"fmt"
	"math"
)


// Use the `type` keyword to define new type names.
type MyInt int


// Defining nnew composite types (structs).
type Point struct {
	X float32
	Y float32
}

type Vertex struct {
	X, Y float64    // Shortened.
}


// Use `iota` keyword to declare a C-like enum.
type Weekday int
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)


func main() {

	var a MyInt = 2
	b := Tuesday
	fmt.Println(a, b)

	// Instantiation.
	p1 := Point{2, 3}
	p2 := Point{
		X: 4,   // Comma here!
		Y: 5,
	}

	// Accessing fields.
	p1.X = 7
	pp := &p2
	pp.Y = 9    // Auto dereference.
	fmt.Println(p1, p1.X, p1.Y)
	fmt.Println(p2, p2.X, p2.Y)

	// Methods.
	v1 := Vertex{3, 4}
	fmt.Println(v1.Distance())
	v1.Move()   // Auto indirection.
	fmt.Println(v1)
}


// Methods on self-defined types (functions with a receiver, not
// necessarily a struct).
func (v Vertex) Distance() float64 {    // This acts on a copy.
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Move() {   // Pointer receiver: the method can modify.
	v.X += 1
}
