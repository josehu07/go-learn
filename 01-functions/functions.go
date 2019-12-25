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


// Functions are also first-class types.
func action(fn func(int) int, num int) int {
    return fn(num)
}


// Closures: functions that references variables from outside its body.
// When a function is "bound" to an outside variable, it may access and
// modify the variable as if it is affliated to the function.
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}


// Execution starts at function `main`.
func main() {
    hello()
    fmt.Println(addInt(1, 23))
    fmt.Println(addFloat(1.1, 3.7))
    fmt.Println(swap("ho", "oh"))
    fmt.Println(split(572))

    // Functions are also first-class types.
    fadd := func(num int) int {     // Anonymous functions.
        return num + 1
    }
    fdec := func(num int) int {
        return num - 1
    }
    fmt.Println(action(fadd, 7), action(fdec, 9))

    // Function closures.
    accumulator := adder()
    fmt.Println(accumulator(2))
    fmt.Println(accumulator(3))
    fmt.Println(accumulator(9))
}


// NOTE: Multiple-line comments can also be written as:
/*
 Lalala.
 Hahaha.
 */
