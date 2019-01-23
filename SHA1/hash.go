package main


import (
    "crypto/sha1"
    "encoding/base64"
    "fmt"
)


func main() {
    s := "this is a string"

    hash := sha1.New()
    hash.Write([]byte(s))

    bs := hash.Sum(nil)
    fmt.Println(s)
    fmt.Printf("%x\n", bs)

    mail := "justzs@gmail.com"
    mailEnc := base64.StdEncoding.EncodeToString([]byte(mail))
    fmt.Println(mailEnc)

    uEnc, _:= base64.StdEncoding.DecodeString(mailEnc)
    fmt.Println(string(uEnc))
}
