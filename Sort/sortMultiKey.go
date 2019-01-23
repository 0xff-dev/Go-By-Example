/*
    数据多直排序

    sort(iter, key=lambda x:(x0, x1, x2))
*/

package main

import (
    "fmt"
    "sort"
)


type Change struct {
    user string
    language string
    lines int
}


type lessFunc func(p1, p2 *Change) bool

type multiSorter struct {
    changes []Change
    less []lessFunc
}


func (ms *multiSorter) Sort(changes []Change) {
    ms.changes = changes
    sort.Sort(ms)
}

func OrderBy(less ...lessFunc) *multiSorter {
    return &multiSorter{
        less: less,
    }
}

func (ms *multiSorter) Len() int {
    return len(ms.changes)
}
func (ms *multiSorter) Swap(i, j int) {
    ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}
func (ms *multiSorter) Less(i, j int) bool {
    d1, d2 := &ms.changes[i], &ms.changes[j]
    index := 0
    for ;index < len(ms.less)-1; index ++{
        switch{
        case ms.less[index](d1, d2):
            return true
        case ms.less[index](d2, d1):
            return false
        }
    }
    return ms.less[index](d1, d2)
}

var cs = []Change{
    {"gri", "Go", 100},
    {"ken", "C", 150},
    {"glenda", "Go", 200},
    {"rsc", "Go", 200},
    {"r", "Go", 100},
    {"ken", "Go", 200},
    {"dmr", "C", 100},
    {"r", "C", 150},
    {"gri", "Smalltalk", 80},
}

func main() {
    user := func(p1, p2 *Change) bool {
        return p1.user < p2.user
    }
    language := func(p1, p2 *Change) bool {
        return p1.language < p2.language
    }
    lines := func(p1, p2 *Change) bool {
        return p1.lines < p2.lines
    }

    OrderBy(user, language, lines).Sort(cs)
    fmt.Println("res: ", cs)

    /*
        python
        data = [
            {"user": "a", "lines": 1},
            {"user": "B", "lines": 2}
        ]
        sort(data, lambda x: (x['user', x['lines']]))

    */
}
