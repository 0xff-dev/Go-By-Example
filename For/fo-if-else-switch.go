package main

import "fmt"
import "time"

func main() {
    for i := 1; i <= 3; i ++ {
        fmt.Println(i)
    }
    for {
        fmt.Println("string")
        break
    }
    for n := 2; n <= 10; n ++ {
        if n % 2 == 0 {
            continue
        }
        fmt.Println(n)
    }
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }
    
    i := 2
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Sunday, time.Saturday:
        fmt.Println("today is 周末")
    default:
        fmt.Println("工作日")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("before 12")
    default:
        fmt.Println("after 12")
    }

    whatType := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("Bool")
        case int:
            fmt.Println("int")
        default:
            fmt.Printf("Unknow %T\n", t)
        }
    }
    whatType(true)
    whatType(4)
    whatType("str")
}
