package main


import "fmt"

func main() {
    m := make(map[string]int)

    m["key1"] = 1
    m["key2"] = 2

    fmt.Println("map: ", m)
    v1 := m["key1"]
    fmt.Println("get: ", v1)

    fmt.Println("len: ", len(m))

    delete(m, "key2")
    fmt.Println("delete: ", m)

    n := map[string]int{"k1": 1, "k2": 2}
    fmt.Println("map init: ", n)
}
