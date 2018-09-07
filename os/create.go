package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("make new file")
    file, err := os.Create("new.txt")

    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
    fmt.Println("complete!", file)
}
