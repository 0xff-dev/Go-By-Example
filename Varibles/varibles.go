package main


import "fmt"


func main() {
    var str = "this is a string"
    fmt.Println(str)

    var int1, int2 int = 1, 2
    fmt.Println(int1, int2)

    var (
        s string
        i int
    )
    fmt.Println(s, i)

    f := "string"
    fmt.Println(f)
}
