package main

import "fmt"
import "time"


func work(id int, jobs <- chan int, results chan <- int){
    for j := range jobs {
        fmt.Println("woker ", id, "start job", j)
        time.Sleep(time.Second)
        fmt.Println("woker ", id, "done job ", j)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    result := make(chan int, 100)

    for i := 0; i<3; i++ {
        go work(i, jobs, result)
    }

    for j := 1; j<5; j++{
        jobs <- j
    }
    close(jobs)
    for a := 1; a<5; a ++ {
        fmt.Println("res: ", <-result)
    }
}
