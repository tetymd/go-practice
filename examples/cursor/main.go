package main

import (
    _ "fmt"
    "github.com/kr/pty"
	"io"
	"os"
	"os/exec"
    _ "github.com/ahmetalpbalkan/go-cursor"
)

func main() {
    // fmt.Print(cursor.ClearScreenUp())
    // fmt.Print(cursor.MoveTo(0, 0))
    // fmt.Print(cursor.MoveUp(4))
    // fmt.Print("\x1b", "[2J")
	c := exec.Command("grep", "--color=auto", "foo")
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	go func() {
		f.Write([]byte("foo\n"))
		f.Write([]byte("bar\n"))
		f.Write([]byte("baz\n"))
		f.Write([]byte{4}) // EOT
	}()
	io.Copy(os.Stdout, f)
}
