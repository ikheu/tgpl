package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}
    fmt.Println(arr)
    test_param(arr)
    fmt.Println(arr)
    zero(&arr)
    fmt.Println(arr)
}


func test_param(arr [3]int) {
    arr[1] = 0
    arr = [3]int {2,3,4}
}

func zero(ptr *[3]int) {
    for i := range ptr {
        ptr[i] = 0
    }
}