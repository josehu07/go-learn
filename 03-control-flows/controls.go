package main


import (
    "fmt"
    "math/rand"
    "time"
    "runtime"
)


func main() {

    // If statements.
    x := rand.Intn(11) - 5
    if x < 0 {
        fmt.Println("Turn left!")
    } else if x > 0 {
        fmt.Println("Turn right!")
    } else {
        fmt.Println("Go straight!")
    }

    // `if` can take a prologue statement.
    if y := rand.Intn(10); y > 5 {  // `y` is inside the scope of if/else.
        fmt.Println("Lucky!")
    }
    // fmt.Println(y)   // This won't work.

    // C-style for loops.
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // The init & post statements are optional => `for` can act as `while`.
    sum := 1
    for sum < 100 {     // `for ; sum < 100; {` is obviously valid.
        sum *= 2
    }
    fmt.Println(sum)

    // Infinite loop is by taking off even the end condition.
    for { break }

    // Switch-case statements, different from C-style!
    // Purely equivalent to a bunch of if-else statements (no need for breaks)
    // and can switch on almost everything (or emtpy) besides integers.
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X")
        // break    // NO need for break here: exactly one branch executed.
    case "linux":
        fmt.Println("Linux")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s\n", os)
    }

    // Allows empty switch condition. Though still not as handy as rusty
    // pattern matching.
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}
