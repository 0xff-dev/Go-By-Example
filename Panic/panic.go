package main


import "os"

func main() {

    _, err := os.Create("/tmp/test.txt")
    if err != nil {
        panic(err)
    }
}
