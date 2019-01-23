package main


import "fmt"
import "sort"


type str []string

func (s str) Len() int {
    return len(s)
}
func (s str) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s str) Less(i, j int) bool {
    return len(s[i]) <= len(s[j])
}

func main() {
    strs := []string{"v", "n", "d"}
    sort.Strings(strs)
    fmt.Println("Sorted strs: ", strs)

    ints := []int{10, 6, 99}
    sort.Ints(ints)
    fmt.Println("Sorted ints: ", ints)
    fmt.Printf("%v sorted: %v", ints, sort.IntsAreSorted(ints))

    // sort.Search
    GuessingGame()

    // 自定义排序函数, 实现 Swap(交换), Less(比较)
    s := []string{"apple", "pear", "lll"}
    sort.Sort(str(s))
    fmt.Println("Sort str type: ", s)
}

func GuessingGame() {
    var s string
    fmt.Println("Pick a number from 0-100\n")
    answer := sort.Search(100, func(i int) bool {
        fmt.Println("Your number <= %d\n?", i)
        fmt.Scanf("%s", &s)
        return s != "" && s[0] == 'y'
    })
    fmt.Println("You Number is: ", answer)
}
