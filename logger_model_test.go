package kek

import (
    "github.com/olehan/kek/formatters/minified"
    "testing"
)

func TestLogger_SetFormatter(t *testing.T) {
    l := NewLogger("name")
    f := l.formatter
    if  f == l.SetFormatter(minified.NewMinifiedFormatter()).formatter {
        panic("formatter haven't change after SetFormatter execution")
    }
}
