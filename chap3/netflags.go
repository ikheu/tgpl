package main 

import (
    "fmt"
    "math"
)

type Flags int

const (
    FlagUp Flags = 1 << iota
    FlagBroadcast
    FlagLookback
    FlagPointToPoint
    FlagMulticast
)

const pi = 3

func IsUp(v Flags) bool { return v&FlagUp == FlagUp }

func main() {
    fmt.Printf("%b\n", 8)
    fmt.Printf("%T\n", math.Pi)
}