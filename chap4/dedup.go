package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    seen := make(mao[string]bool)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            seen[line] = true
            fmt.Println(line)
        }
    }
    if err := input.Err(); err != nil {
        fmt.Println(os.Stderr, "dedup: %v\n", err)
        os.exit(1)
    }
}