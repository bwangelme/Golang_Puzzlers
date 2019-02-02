package os

import (
    "io"
    "fmt"
)


func Hello(w io.Writer, content string) {
    fmt.Fprint(w, content)
}