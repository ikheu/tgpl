package main

import "fmt"

const boilingF = 212.0

func main() {
    // var i int = 10
    // var f1 float64 = 10.0
    fmt.Printf("%T\n", 10)  // int
    fmt.Printf("%T\n", 10.0) //float64
    c := 10
    // c = 11.1 // constant 11.1 truncated to integer
    fmt.Printf("%T\n", c) // int
    d := 10.0
    fmt.Printf("%T\n", d) // int
    d = 10 + 11.1
    var e float64 = 10 // ok
    fmt.Printf("%T\n", e) // float64
    // var f float64 = c // cannot use c (type int) as type float64 in assignment
    // fmt.Printf("%T\n", f)
    // fmt.Printf("%T\n", c == d) // invalid operation: c == d (mismatched types int and float64)
}