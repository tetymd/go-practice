package main

import (
    "fmt"
    "os"
    "bufio"
    "io/ioutil"
)

func main() {
    file_name := "output.txt"
    output := "Hello World!"

    // ファイルの存在有無を確認
    s, err := os.Stat(file_name)
    fmt.Println(s, err)

    if err != nil {
        // ファイルを作成
        file, err := os.Create(file_name)
        if err != nil {
            fmt.Println(err)
            os.Exit(0)
        }
        defer file.Close()
        fmt.Println("Create", file_name, *file)
    }

    // 追記モードでファイルに書き込み
    file, err := os.OpenFile(file_name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
    defer file.Close()
    file.Write(([]byte) (output))

    // 標準入力から読み込み
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    fmt.Println(scanner.Text())
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }

    // ファイルを読み込み
    data, err := ioutil.ReadFile(file_name)
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }

    fmt.Println(data)
    fmt.Println(string(data))
}
