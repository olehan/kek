package colors

import (
    "fmt"
    "strconv"
)

// JoinColors joins given colors together to get single ansi sequence.
func JoinColors(colors ...Color) string {
    additionalColors := ""
    for _, n := range colors {
        additionalColors += ";" + strconv.Itoa(int(n))
    }
    return additionalColors
}

// String stringify's color and joins it with additional color options.
func String(colors ...Color) string {
    return fmt.Sprintf("\x1b[%vm", JoinColors(colors...))
}

func (c Color) String(colors ...Color) string {
    return String(append(colors, c)...)
}

// Add adds color for the given string.
func (c Color) Add(v string, colors ...Color) string {
    return c.String(colors...) + v + Reset.String()
}
