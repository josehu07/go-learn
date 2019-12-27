package main


import "fmt"


// Global variables.
var me = "A student"
var (   // Batched declaration.
	name = "Guanzhou"
	age = 100
)


// Constants.
const Pi = 3.14
const Days = 365
const (
	One = 1
	AlsoOne     // Omit `= xxx` means the same as previous one.
	Two = 2
)


func main() {

	// Basic types include:
	//   bool
	//   string
	//   int/uint (follows os bit-width)
	//   int/uint[8|16|32|64]
	//   uintptr (pointer type)
	//   byte (i.e., uint8), rune (i.e., int32)
	//   float[32|64]
	//   complex[64|128]

	// Uninitialized declaration.
	var i int
	var a, b, c float32
	i = 123
	a, b, c = 1, 2, 3
	fmt.Println(i, a, b, c)

	// Initialized declaration.
	var k int64 = 12345
	var j = 7   // Type is inferred.
	var b1, b2 = true, false
	fmt.Println(k, j, b1, b2)

	// Shortened declaration can be used inside a function.
	g := 9.8
	h := 0.38 + 0.5i
	fmt.Printf("g = %v, type is %T; h = %v, type is %T\n", g, g, h, h)

	// Zero values (default values) are safely handled. REMEMBER them.
	var iz int
	var fz float64
	var bz bool
	var sz string
	fmt.Printf("%v %v %v %q\n", iz, fz, bz, sz)

	// Type conversions must be done EXPLICITLY!
	var i0 int = 42
	ic := i0    // `ic` is of type `int`
	fc := float32(ic)
	fmt.Printf("fc = %.1f, type is %T\n", fc, fc)
	// var fe float32 = ic  // This won't work.

	// Pointers only bear simple semantics.
	n := 3
	var p1 *int = &n
	p2 := &n
	*p1 = 7
	fmt.Println(p2, *p2)

	// Untyped numeric constants can fit in different type contexts. Both
	// should succeed.
	needInt(Days)
	needFloat(Days)
}


func needInt(x int) int { return x }
func needFloat(x float32) float32 { return x }


// NOTE: Check the `fmt` package's doc for detailed formatting options.
