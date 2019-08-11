package main

import "fmt"

func main() {
	s := "hello"
	b := []byte(s)
	arr := []int {1, 2, 3}
	b[len(b)-1] = 'l'
	fmt.Println(s)
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", b)
}
