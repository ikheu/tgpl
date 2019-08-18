package main


import "fmt"

func main() {
    // arr := [3]int {1, 2, 3}
    // sli := arr[:2]
    // sli = append(sli, 4)
    // fmt.Println(sli)
    // fmt.Println(arr)
    var x []int
    for i:=0; i < 10; i++ {
        x = appendInt(x, i)
        fmt.Printf("%d, cap=%d\t%v\n", i, cap(x), x)
    }
}

func appendInt(x []int, y int) []int {
    var z[]int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        z = x[:zlen]
    } else {
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2 * len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z, x)
    }
    z[len(x)] = y
    return z
}