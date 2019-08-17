package main

import "crypto/sha256"
import "fmt"

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    fmt.Println([]byte("xdssds"))
    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
    arr := [3]int{1, 2, 3}
    fmt.Println(arr)
    test_param(arr)
    fmt.Println(arr)
    // ptr := *arr
    zero(&arr)
    fmt.Println(arr)
}


func test_param(arr [3]int) {
    arr[1] = 0
    arr = [3]int {2,3,4}
    // fmt.Println(arr)
}

func zero(ptr *[3]int) {
    for i := range ptr {
        ptr[i] = 0
    }
}