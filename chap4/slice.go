package main

import "fmt"

func main() {
    months := [...]string{1: "January", "February", "March", "April", "May", "June", "July", "Auguest", "September", "October", "Norvember", "December"}
    fmt.Println(months)
    fmt.Printf("%T\n", months)

    Q2 := months[4:7]
    summer := months[6:9]
    fmt.Println(Q2)
    fmt.Println(summer)

    endlessSummer := summer[:5]
    fmt.Println(endlessSummer)

    s := "hello"
    m := s[1:2]
    // m[0] = 's'
    fmt.Printf("%T\n", m)
    fmt.Println(s)

    b:= []byte{'q', 's'}
    c := b[:]
    c[0] = 'a'
    fmt.Println(b)
    fmt.Println(c)


    arr1 := []int {}
    // arr2 := []int {2,3,4}
    fmt.Println(arr1 == nil)
}