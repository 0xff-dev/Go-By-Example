package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    t := time.Now()
    p(t.Format(time.RFC3339)) // 时间格式化

    t1, err := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00",
    )
    p(t1)

    p(t.Format("3:40PM"))
}
