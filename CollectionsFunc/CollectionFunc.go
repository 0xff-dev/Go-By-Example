package main

import "fmt"
import "strings"


func Index(vs []string, t string) int {
    for index, val := range vs {
        if val == t {
            return index
        }
    }
    return -1
}

func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v){
            return true
        }
    }
    return false
}

func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

func Map(vs []string, f func(string) string) []string {
    res := make([]string, len(vs))
    for index, v := range vs {
        res[index] = f(v)
    }
    return res
}

func Filter(vs []string, f func(string) bool) []string {
    res := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            res = append(res, v)
        }
    }
    return res
}

func main() {
    var strs = []string{"peach", "apple", "pear", "plum"}
    fmt.Println(Index(strs, "pear"))
    fmt.Println(Include(strs, "grape"))
    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))
    fmt.Println(Map(strs, strings.ToUpper))
}
