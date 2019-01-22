/*
    单纯的等待, 可以使用time.Sleep, timer在结束之前是可以停止的
*/
package main

import "fmt"
import "time"

func main() {
    timer1 := time.NewTimer(2 * time.Second)

    <-timer1.C
    fmt.Println("Timer 1 expired")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stoped")
    }
}
