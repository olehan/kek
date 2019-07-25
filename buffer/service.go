package buffer

import (
    "fmt"
    "strconv"
)

// Write appends byte slice into buffer.
func (b *Buffer) Write(p []byte) {
    *b = append(*b, p...)
    fmt.Println()
}

// WriteString appends string into buffer.
func (b *Buffer) WriteString(s string) {
    *b = append(*b, s...)
}

// WriteByte appends a single byte into buffer.
func (b *Buffer) WriteByte(c byte) {
    *b = append(*b, c)
}

// WriteInt appends a single int into buffer.
func (b *Buffer) WriteInt(i int64) {
    *b = strconv.AppendInt(*b, i, 10)
}

// WriteUint appends a single uint into buffer.
func (b *Buffer) WriteUint(i uint64) {
    *b = strconv.AppendUint(*b, i, 10)
}

// WriteFloat appends a single float into buffer.
func (b *Buffer) WriteFloat(f float64, size int) {
    *b = strconv.AppendFloat(*b, f, 'f', -1, size)
}

// WriteBool appends bool into buffer.
func (b *Buffer) WriteBool(bl bool) {
    *b = strconv.AppendBool(*b, bl)
}

// WriteSpace appends space into buffer.
func (b *Buffer) WriteSpace() {
    b.WriteByte(' ')
}

// WriteNewLine appends new line space into buffer.
func (b *Buffer) WriteNewLine() {
    b.WriteByte('\n')
}

// Reset resets the byte slice.
func (b *Buffer) Reset() {
    *b = (*b)[:0]
}
