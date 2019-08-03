package main

import "fmt"

const boilingF = 212.0

func main() {
    var f = boilingF
    var c = (f - 32) * 5 / 9
    var i int = 10
    var f1 float64 = 10
    fmt.Printf("%g\n", f1)
    fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}