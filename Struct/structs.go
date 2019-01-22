package main


import "fmt"


type person struct {
    name string
    age int
}

func (p *person) Hello () {
    fmt.Printf("person: %s say Hello\n", p.name)
}

func (p person) World() {
    fmt.Printf("person: %s say World\n", p.name)
}

func main() {
    fmt.Println(person{
        name: "Bob",
        age: 4,
    })
    fmt.Println(&person{
        name: "Ann",
        age: 13,
    })
    s := person{
        name: "Coco",
        age: 20,
    }
    ptr := &s
    fmt.Println(ptr.age, ptr.name)
    ptr.age = 100
    fmt.Println(ptr.age)

    ptr.Hello()
    ptr.World()
}
