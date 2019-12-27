package main


import "fmt"


// One of Go's special characteristics: defering. Very useful.
var global_counter = 0

func incrementSharedCounter() {
	fmt.Println("Acquire a lock")

	// Defered actions happen automatically after host function exits.
	defer fmt.Println("Release the lock")

	// Do something that may trigger runtime errors or containing many
	// side conditions. No need for writing an unlock statement before
	// every early `return`. Safe & sound!
	global_counter += 1
}


// Multiple defers follow LIFO (stack) order.
func multipleDefers() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}


// Another interesting property: variable escaping. Go runtime automatically
// lets a variable "escape" from local function stack to global heap when
// it is still referenced after the function exits.
//
// An interesting way to avoid certain memory management errors (with some
// amount of runtime overhead).
var global_pointer *int

func letNumEscape() {
	var num int
	num = 37892
	global_pointer = &num
}


func main() {
	incrementSharedCounter()
	multipleDefers()

	letNumEscape()
	fmt.Println(*global_pointer)
}
