package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)


func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        data := strings.ToUpper(scanner.Text())
        fmt.Println(data)
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error: ", err)
        os.Exit(1)
    }
}
