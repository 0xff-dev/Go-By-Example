package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            job, more := <-jobs
            if more {
                fmt.Println("received job")
            } else {
                fmt.Println("received all jobs: ", job)
                done <- true
                return
            }
        }
    }()

    for i := 0; i<3; i++ {
        jobs <- i
        fmt.Println("send ", i)
    }
    close(jobs) // 关闭后, 不能写, 可以读
    fmt.Println("send all jobs")
    <-done
}
