package main

import (
    "fmt"
    "time"
)


func main() {
    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC,
    )
    p(then)
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Location())
    p(then.Nanosecond())
    
    // 返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。
    // 要获取时间点t-d（d为Duration），可以使用t.Add(-d)。
    diss := then.Sub(then)
    p(diss)
    p(diss.Hours())
    p(then.Add(diss))
    p(then.Add(-diss))

    // Epoch
    fmt.Println("Epoxh")
    now = time.Now()
    secs := now.Unix()
    naons := now.UnixNano()
    fmt.Println(secs)
    fmt.Println(naons)
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, naons))
}
