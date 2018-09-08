package main

import (
    "fmt"
)

func main() {
    var i int = 100
    // ipにiのアドレスを代入
    var ip *int = &i

    fmt.Println(i)
    // fmt.Println(i)と同じ結果
    fmt.Println(*ip)
}
