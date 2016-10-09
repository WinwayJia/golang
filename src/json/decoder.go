package main

import (
    "encoding/json"
    "fmt"
)

type Employee struct {
    Name string "name"
    Age int "age"
    Gender int "gender"
}

func (e *Employee)Show() {
    fmt.Println("Name:", e.Name, "Age:", e.Age, "Gender:", e.Gender)
}

type Book struct {
    Title string "title"
    Authors []string
    Publisher string
    IsPublished bool
    Price float32
    TestMap map[string] int
}

func main() {
    var e = Employee{Name: "john", Age: 20, Gender: 0}
    ret, err := json.Marshal(e)
    if err != nil {
        return
    }
    fmt.Println(string(ret))

    str := "{\"Name\":\"john\",\"Age\":20,\"Gender\":0}"
    var ee Employee
    err = json.Unmarshal([]byte(str), &ee)
    if err != nil {
        return
    } else {
        ee.Show()
    }
    testmap := map[string] int {"Hello": 5}
    authors := []string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"}
    gobook := Book{"Go语言编程", authors[0:], "ituring.com.cn", true, 9.99, testmap}
    b, err := json.Marshal(gobook)
    fmt.Println(string(b))

    // 解码未知结构JSON数据
    b = []byte(`{
        "Title": "Go语言编程",
        "Authors": ["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan",
        "XuDaoli"],
        "Publisher": "ituring.com.cn",
        "IsPublished": true,
        "Price": 9.99,
        "Sales": 1000000
        }`)
    var r interface{}
    err = json.Unmarshal(b, &r)
    fmt.Println(r)
    book, ok := r.(map[string]interface{})
    if ok {
        fmt.Println(book["Sales"], book["Publisher"])
    }

}
