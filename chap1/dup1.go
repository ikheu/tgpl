// counter repeated lines of stdin

package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for  input.Scan() {  // EOF 时, input.Scan() == false
        counts[input.Text()] ++
    }

    fmt.Println(counts)

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n ,line)
        }
    }
}