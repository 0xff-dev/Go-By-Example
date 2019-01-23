package main


import "fmt"
import "os"


func main() {
    file := CreateFile("/tmp/res.txt")
    defer CloseFile(file)
    writeFile(file)
}

func CreateFile(path string) *os.File {
    fmt.Println("Create")
    file, err := os.Create(path)
    if err != nil {
        panic(err)
    }
    return file
}
func writeFile(file *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(file, "data")
}

func CloseFile(file *os.File){
    fmt.Println("close file")
    file.Close()
}
