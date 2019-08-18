package main

import "fmt"

func main() {
    data := []string {"one", "", "three"}
    data = nonempty(data)
    fmt.Printf("%q\n", data)
}

func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i++
        }
    }
    return strings[:i]
}