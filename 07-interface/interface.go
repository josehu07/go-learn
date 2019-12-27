package main


import (
	"fmt"
	"math"
)


// Interface: a set of method signatures. A value of interface type can
// hold any value that implements those methods. Kinda like traits in
// Rust but not as flexible and powerful.
type Abser interface {
	Abs() float64
}

type MyInt int

func (v MyInt) Abs() float64 {
	if v >= 0 {
		return float64(v)
	}
	return float64(-v)
}

type MyFloat float32

func (v MyFloat) Abs() float64 {
	if v >= 0 {
		return float64(v)
	}
	return float64(-v)
}

type MyPair struct {
	X, Y float64
}

func (p *MyPair) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}


func main() {
	var a Abser
	fmt.Printf("%v, %T\n", a, a)

	a = MyInt(-233)
	a = MyFloat(-1.3)
	a = &MyPair{3, 4}   // NOTE: *MyPair implements Abser, not MyPair.
	fmt.Printf("%v, %T\n", a.Abs(), a)

	// Empty interface may hold values of any type.
	var i interface{}
	i = 3
	i = "hello"
	fmt.Printf("%v, %T\n", i, i)

	// Type assertions on interface variables.
	val := i.(string)   // If `i` is not a string, panic!
	// val := i.(int)   // This panics!
	_, ok := i.(int)    // Get assertion result in `ok`.
	fmt.Println(val, ok)

	// Type switches.
	switch v := i.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v, len(v))
	case MyPair:
		fmt.Println(v.X, v.Y)
	default:
		fmt.Println("Unknown type!")
	}
}


// NOTE: The `Stringer` interface requires a `String()` method. We can
// implement a `String()` method for custom types to enable `fmt.Print*`
// on instances of these types.

// NOTE: The `error` interface requires an `Error()` method. We can
// implement an `Error()` method on customized error types.
// A `nil` error denotes success; non-nil denotes error.

// NOTE: Other interesting standard library interfaces.
//   - `io.Reader` for inputing
//   - `image.Image` for images
