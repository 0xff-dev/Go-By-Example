package main

import (
    "encoding/json"
    "fmt"
    "os"
)


type testStruct struct {
    field int `json:", omitempty"` // 使用字段默认值, 如果为空忽略
}

type response1 struct {
    Page int
    Fruits []string
}

type response2 struct {
    Page int `json:"page"`
    Fruits [] string `json:"fruits, omitempty"` // 如果字段值为空, 忽略
}


func main() {
    boolB, _ := json.Marshal(true)
    fmt.Println(string(boolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("string")
    fmt.Println(string(strB))

    slcB, _ := json.Marshal([]string{"apple", "pear", "banana"})
    fmt.Println(string(slcB))

    mapB, _ := json.Marshal(map[int]string{1: "1", 2: "2"})
    fmt.Println(string(mapB))

    res1D := &response1{
        Page: 1,
        Fruits: []string{"a", "b", "c"},
    }
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))
    
    res2D := &response2{
        Page: 2,
        Fruits: []string{"a", "b", "c"},
    }
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    byt := []byte(`{"num": 6.13, "strs": ["a", "b"]}`)
    var dat map[string]interface{}

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)
    num := dat["num"].(float64)
    fmt.Println(num)

    strs := dat["strs"].([]interface{})
    fmt.Println(strs[0].(string))

    str := `{"page": 1, "fruits": ["a", "w", "s"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    file, _ := os.Create("./res.json")
    enc := json.NewEncoder(file)
    d := map[string]int{
        "apple": 1,
        "pear": 10,
    }
    enc.Encode(d)
}
