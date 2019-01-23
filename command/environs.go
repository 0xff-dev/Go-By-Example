package main


import (
    "os"
    "strings"
    "fmt"
)

func main() {
    os.Setenv("FOO", "foo")
    fmt.Println("FOO: ", os.Getenv("FOO"))
    for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Printf("%s -> %s\n", pair[0], pair[1])
    }
}
