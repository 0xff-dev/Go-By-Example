/*
    ticker断续器, 设定间隔, 在休眠睡觉/休息间隔 就是要执行的次数
*/
package main

import "fmt"
import "time"


func main() {
    ticker := time.NewTicker(500 * time.Millisecond)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at: ", t)
        }
    }()
    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    fmt.Println("Ticker stoped")
}
