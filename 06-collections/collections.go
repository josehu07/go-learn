package main


import "fmt"


func main() {

    //
    // Arrays.
    //

    // Primitive arrays.
    var a1 [10]float32  // Zero values by default.
    a1[3] = 3.14
    fmt.Println(a1[3], a1)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)

    // Slice is a dynamically-sized view into the elements of an
    // array. Indices follow Python flavor.
    var s1 []int = primes[:4]
    s2 := primes[1:5]
    s3 := primes[2:]
    s1[3] = 71  // Slices are references instead of copies.
    fmt.Println(s1, s2, s3, primes)

    // Use slice literals to build an array and create a slice that
    // references it. Go's "variable escape" mechanism ensures that
    // dangling pointers won't appear.
    s4 := []struct{
        name string
        age int
    }{
        {"Alice", 13},
        {"Bob", 17},
        {"Caman", 79},
    }
    fmt.Println(s4)

    // A slice's capacity is the number of elements in the underlying
    // array, counting from the first element of the slice. Its length
    // cannot exceed its capacity.
    a2 := [4]bool{true, false, true, false}
    s5 := a2[1:3]
    fmt.Println("Len:", len(s5), "Cap:", cap(s5))
    s5 = s5[:3]     // Reslicing: extend its length.
    fmt.Println("Len:", len(s5), "Cap:", cap(s5))
    s5 = s5[1:]     // Reslicing: drop elements.
    fmt.Println("Len:", len(s5), "Cap:", cap(s5))

    var s6 []int    // Zero value of a slice is `nil`.
    fmt.Println(s6==nil)

    // Use the `make` function to dynamically allocate runtime-length
    // slices with given length (and capacity). Use `append` to
    // conveniently append a value to slice.
    s7 := make([]int, 5)
    s8 := make([]int, 0, 5)
    s7 = append(s7, 123)
    fmt.Println(s7, s8)


    //
    // Maps.
    //

    // Initialize with `make`.
    m1 := make(map[string]bool)
    m1["I am a student"] = true
    fmt.Println(m1["I am a student"], m1)

    // Map literals.
    type Vertex struct {
        Lat, Long float64
    }

    var m2 = map[string]Vertex{
        "Bell Labs": {40.68433, -74.39967},     // Can omit `Vertex` here.
        "Google":    {37.42202, -122.08408},
    }

    // Deleting & retrivinng elements.
    delete(m2, "Bell Labs")
    elem1, ok1 := m2["Bell Labs"]   // Retriving with check.
    elem2, ok2 := m2["Google"]
    fmt.Println(elem1, ok1, elem2, ok2)


    //
    // Range syntax.
    //

    // Iterating using `range`. ALso supports:
    //   idx, _
    //   _, val
    //   i (only index)
    for idx, val := range primes[:] {
        fmt.Println(idx, val)
    }
}
