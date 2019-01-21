package main

import "fmt"

func puls(a, b int) int {
    return a + b
}

func testType(types ...interface{}) {
    format := "%v type is %T\n"
    for _, _type := range types {
        fmt.Printf(format, _type, _type)
    }
}

func testMultiReturnAndV() (s, t int) {
    s = 1
    t = 6
    return
}

func main() {
    fmt.Println("1 + 2", puls(1, 2))
    testType(1, 2, 2, "abc", []int{})
    s, t := testMultiReturnAndV()
    fmt.Println(s, t)
}
