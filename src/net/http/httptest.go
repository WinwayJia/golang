package main

import (
    "fmt"
    "net/http"
    "os"
    "io"
)

func main() {
    resp, err := http.Get("http://baidu.com")
    if err != nil {
        os.Exit(1)
    } else {
        fmt.Println(resp, "\n\n")
    }
    defer resp.Body.Close()

    var c http.Client
    resp2, err := c.Get("http://baidu.com")
    if err != nil {
        os.Exit(1)
    } else {
        io.Copy(os.Stdout, resp2.Body)
    }
    defer resp2.Body.Close()

    fmt.Println("\n\nreps3\n")
    resp3, err := http.Head("http://baidu.com")
    if err != nil {
        fmt.Println("Error:")
        os.Exit(1)
    } else {
        io.Copy(os.Stdout, resp3.Body)
    }
    defer resp3.Body.Close()

}
