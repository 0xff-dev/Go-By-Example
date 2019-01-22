/*
    默认情况channel是阻塞的, 通过设置select的default来进行non-blocking
*/
package main

import "fmt"
import "time"

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i<2; i++ {
        select {
        case msg := <- c1:
            fmt.Println("received from c1: ", msg)
        case msg := <- c2:
            fmt.Println("received from c2: ", msg)
        }
    }

    msgs := make(chan string)

    select {
    case msg := <-msgs:
        fmt.Println("received msgs: ", msg)
    default:
        fmt.Println("no msgs")
    }
}
