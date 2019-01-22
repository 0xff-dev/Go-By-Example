package main


import "fmt"


func zeroval(val int) {
    val = -1
}

func zeroptr(val *int) {
    *val = -1
}

func main() {
    i := 1
    zeroval(i)
    fmt.Println("val -> : ", i)

    zeroptr(&i)
    fmt.Println("val ->: ", i)
}
