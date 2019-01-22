/*
    channel 配合time.Afters实现channel的超时处理
*/
package main


import "fmt"
import "time"

func main() {
    channel := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        channel <- "res data"
    }()

    select {
    case msg := <-channel:
        fmt.Println(msg)
    case <- time.After(1 * time.Second):
        fmt.Println("timeout 1s")
    }

    channel2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        channel2 <- "two"
    }()

    select{
    case msg := <-channel2:
        fmt.Println(msg)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2s")
    }
}
