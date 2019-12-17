// Code starts at package `main`.
package main


// Imports syntax.
import "fmt"        // Single.
import (            // Factored.
    "time"
    "math"
    "math/rand"
)


// Exported names (i.e., public names) start with Capital letter. Unexported names
// (private names) cannot be accessed from outside those packages.
func main() {
    start := time.Now()
    fmt.Println("Roll a dice ->", rand.Intn(6)+1)
    fmt.Printf("PI = %.3f\n", math.Pi)
    fmt.Println("Time elapsed:", time.Since(start))
}
