package main


import (
    "fmt"
    "math"
)


type geometry interface {
    area() float64
    perim() float64
}


type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (r rect) area()float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}


func testInterface(g geometry) {
    fmt.Printf("area: %.4f, perim: %.4f\n", g.area(), g.perim())
}
func main() {
    r := rect{
        width: 5,
        height: 6,
    }
    c := circle {
        radius: 5,
    }
    testInterface(r)
    testInterface(c)
}
