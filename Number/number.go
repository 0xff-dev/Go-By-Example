package main


import (
    "fmt"
    _ "math/rand"
    "strconv"
)


func main() {
    format := "%v -> %T\n"
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Printf(format, f, f)

    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Printf(format, i, i)

    k, _ := strconv.Atoi("1234")
    fmt.Printf(format, k, k)

}
