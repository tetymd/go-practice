package main

import (
    "fmt"
    "os"
    "time"
    "github.com/kr/pty"
    "github.com/nsf/termbox-go"
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

    // ptsに書き込み＆読み込み
    buf := make([]byte, 1024)
    p.Write([]byte("\x1b[2J"))
    n, err := p.Read(buf)
    fmt.Println(string(buf[:n]))

    row, col, err := pty.Getsize(t)
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
    fmt.Println(row, col)

    ws := &pty.Winsize{uint16(3), uint16(3), uint16(3), uint16(3)}

    err = pty.Setsize(t, ws)
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

    // fmt.Println("\x1b[2J")

    err = termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()

    termbox.SetCursor(0, 0)

    for {
        switch ev := termbox.PollEvent(); ev.Type {
            case termbox.EventKey:
                switch ev.Key {
                    case termbox.KeyEsc:
                        fmt.Println("ESC")
                        return
                    case termbox.KeyEnter:
                        fmt.Println("\n")
                    case termbox.KeyArrowUp:
                        fmt.Println("UP")
                    case termbox.KeyArrowDown:
                        fmt.Println("DOWN")
                    default:
                        fmt.Println("Other")
                }
            default:
        }
    }

    fmt.Println("\x1b[7A")
    time.Sleep(3 * time.Second)
    fmt.Println("\x1b[3B")
    time.Sleep(3 * time.Second)


}
