package main

import (
    "fmt"
    "os"
    "github.com/kr/pty"
)

func main() {
    p, t, err := pty.Open()
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
    defer p.Close()
    defer t.Close()
    fmt.Println(p, t)

    row, col, err := pty.Getsize(t)
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
    fmt.Println(row, col)

    ws := &pty.Winsize{uint16(3), uint16(3), uint16(3), uint16(3)}

    pty.Setsize(t, ws)
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }

    row, col, err = pty.Getsize(t)
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
    fmt.Println(row, col)

}
