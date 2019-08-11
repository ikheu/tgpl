package main

import "fmt"

func main() {
	fmt.Println(comma("123456789试试看"))
}

func comma(s string) string {
    n := len(s)
    slen := 3
    if n <= slen {
        return s
    }
    return comma(s[:n-slen]) + "," + s[n-slen:]
}