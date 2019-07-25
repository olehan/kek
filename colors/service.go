package colors

import (
    "fmt"
    "strconv"
)

// String stringify's color and joins it with additional color options.
func (c Color) String(colors ...Color) string {
    return fmt.Sprintf("\x1b[%v%vm", strconv.Itoa(int(c)), c.JoinColors(colors...))
}

// Add adds color for the given string.
func (c Color) Add(v string, colors ...Color) string {
    return c.String(colors...) + v + Reset.String()
}

// JoinColors joins given colors together to get single ansi sequence.
func (c Color) JoinColors(colors ...Color) string {
    additionalColors := ""
    for _, n := range colors {
        additionalColors += ";" + strconv.Itoa(int(n))
    }
    return additionalColors
}
