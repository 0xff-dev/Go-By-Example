package main


import (
    "fmt"
    "sort"
)


type earthMass float64
type au float64


type Planet struct {
    name string
    mass earthMass
    disatance au
}

// less function
type By func(p1, p2 *Planet) bool


type planetSorter struct {
    planets []Planet
    by func(p1, p2 *Planet) bool
}

func (p *planetSorter) Len() int {
    return len(p.planets)
}
func (p *planetSorter) Swap(i, j int) {
    p.planets[i], p.planets[j] = p.planets[j], p.planets[i]
}
func (p *planetSorter) Less(i, j int) bool {
    return p.by(&p.planets[i], &p.planets[j])
}

func (by By) Sort(planets []Planet) {
    ps := &planetSorter{
        planets: planets,
        by: by,
    }
    sort.Sort(ps)
}

var pls = []Planet {
    {"Mercury", 0.055, 0.4},
    {"Venus", 0.815, 0.7},
    {"Earth", 1.0, 1.0},
    {"Mars", 0.107, 1.5},
}


func main() {
    name := func(p1, p2 *Planet) bool {
        return p1.name < p2.name
    }
    disatance := func(p1, p2 *Planet) bool {
        return p1.disatance < p2.disatance
    }
    By(name).Sort(pls)
    fmt.Println("By Name: ", pls)
    By(disatance).Sort(pls)
    fmt.Println("By Dis: ", pls)
}
