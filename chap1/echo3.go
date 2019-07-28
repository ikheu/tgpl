// echo2 prints command-line argument
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)
}
