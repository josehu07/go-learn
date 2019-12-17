package main


import "fmt"


// Functions are first class in Go. They can take zero or more arguments and
// return zero or more values.
func hello() {
    fmt.Println("Hello, world!")
}

func addInt(x int, y int) int {         // Type comes after the name.
    return x + y
}

func addFloat(x, y float32) float32 {   // Shortened.
    return x + y
}

func swap(x, y string) (string, string) {   // Multiple return values.
    return y, x
}


// A special specification is so-called named return values. They are like
// declared when entering this function and automatically returned when
// the `return` statement is reached.
// Avoid using them in complicated functions.
func split(num int) (x, y int) {
    y = num % 10
    x = num - y
    return
}


// Execution starts at function `main`.
func main() {
    hello()
    fmt.Println(addInt(1, 23))
    fmt.Println(addFloat(1.1, 3.7))
    fmt.Println(swap("ho", "oh"))
    fmt.Println(split(572))
}
