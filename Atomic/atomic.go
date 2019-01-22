package main


import "fmt"
import "time"
import "sync/atomic"


func main() {
    var ops int64 = 1
    for i := 0; i < 50; i++ {
        go func() {
            atomic.AddInt64(&ops, 2)
            time.Sleep(time.Millisecond)
        }()
    }
    time.Sleep(time.Millisecond)
    fmt.Println("ops: ", atomic.LoadInt64(&ops))
}
