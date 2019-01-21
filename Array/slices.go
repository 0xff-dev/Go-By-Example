package main

import "fmt"

func main() {
    s := make([]string, 3)
    fmt.Println("emp: ", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set: ", s)
    fmt.Println("get: ", s[2])
    fmt.Println("len: ", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Printf("apd: %v, len: %d, cap: %d", s, len(s), cap(s))
    c := make([]string, len(s))
    copy(c, s) // cp后, 不是同地址
    fmt.Println("cpy: ", c)
    c[1] = "abc"
    fmt.Println("cpc: ", s, c)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i ++{
        innterLen := i + 1
        twoD[i] = make([]int, innterLen)
        for j := 0; j < innterLen; j ++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("make create slice: ", twoD)
}
