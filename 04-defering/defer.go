package main


import "fmt"


// One of Go's special characteristics: defering.
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
func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
