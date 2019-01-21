package main

import "fmt"

/*
    python角度
    def func():
        i = 0
        def innerFunc():
            nonlocal i
            i += 1
            return i
        return innerFunc
    闭包, 函数内部定义变量供内部函数进行局部使用.
*/
func clusore() func() int {
    i := 1
    return func() int {
        i ++
        return i
    }
}


func testInner() []func() int{
    funcs := make([]func() int, 0)
    for i := 0; i<3; i ++{
        funcs = append(funcs, func(i int) func() int {
            return func() int{
                return i * i
            }
        }(i))
    }
    return funcs
}

func main() {
    cl := clusore()
    for i := 0; i<4; i++{
        fmt.Println(cl())
    }
    funcs :=  testInner()
    fmt.Println(funcs)
    for _, f := range funcs {
        fmt.Println(f())
    }
}
