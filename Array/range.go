package main


import "fmt"


func main() {
    fmt.Println("range slice")
    nums := []int{1, 2, 3, 4}
    sum := 0
    for _, v := range nums {
        sum += v
    }
    fmt.Println("sum of nums is: ", sum)
    
    fmt.Println("range map")
    kvs := map[string]string{"a": "apple", "b": "banana", "c": "cccc"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    
    fmt.Println("range string")
    for i, c := range "string" {
        fmt.Printf("%d -> %c\n", i, c)
    }
}
