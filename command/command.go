package main


import "fmt"
import "flag"


func main() {
    wordPtr := flag.String("word", "foo", "a string")
    numPtr := flag.Int("numb", 41, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")
    
    flag.Parse()
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("total: ", flag.Args())
}
