package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
)

func main() {
    counts := make(map[string]int)
    for _, file := range os.Args[1:] {
        data, err := ioutil.ReadFile(file)
        if err != nil {
            fmt.Fprintf(os.Stderr, "%v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {
            counts[line] ++
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}