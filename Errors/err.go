package main

import "fmt"
import "errors"


func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }
    return arg+1, nil
}

type argError struct {
    arg int
    prob string
}

func (e *argError) Error() string{
    return fmt.Sprintf("%d - %s\n", e.arg, e.prob)
}

func main(){
    for _, v := range []int{1, 42, 54} {
        if res, err := f1(v); err != nil {
            fmt.Println("f1 failed")
        } else {
            fmt.Println("f1 worked -> ", res)
        }
    }
}
