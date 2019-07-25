package buffer

// Buffer is a small byte slice, made to avoid usage of
// big dependencies as bytes.Buffer.
type Buffer []byte

// NewBuffer returns new buffer.
func NewBuffer() Buffer {
    return Buffer{}
}
