package main

import "fmt"

func main() {
    arr := [3]int {1, 2, 3}
    reverse(arr[:])
    fmt.Println(arr)
}


func reverse(s []int) {
    for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
        swap(&s[i], &s[j])
    }
}


func swap(p1 *int, p2 *int) {
    *p1, *p2 = *p2, *p1
}
