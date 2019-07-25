package buffer

import (
    "strconv"
    "testing"
)

var (
    _testBuffer = NewBuffer()
)

func TestBuffer_Write(t *testing.T) {
    bytes := []byte{1, 4, 6, 7, 1, 3}
    _testBuffer.Write(bytes)
    for i, n := range bytes {
        if _testBuffer[i] != n {
            panic("byte slice and buffer are not equal")
        }
    }
    _testBuffer.Reset()
}

func TestBuffer_WriteInt(t *testing.T) {
    expectedValue := 245824
    _testBuffer.WriteInt(int64(expectedValue))
    if string(_testBuffer) != strconv.Itoa(expectedValue) {
        panic("int value and buffer are not equal")
    }
    _testBuffer.Reset()
}

func TestBuffer_WriteUint(t *testing.T) {
    expectedValue := 245824
    _testBuffer.WriteUint(uint64(expectedValue))
    if string(_testBuffer) != strconv.Itoa(expectedValue) {
        panic("uint value and buffer are not equal")
    }
    _testBuffer.Reset()
}

func TestBuffer_WriteFloat(t *testing.T) {
    expectedValue := float64(245824.4534543)
    _testBuffer.WriteFloat(expectedValue, 64)
    if string(_testBuffer) != strconv.FormatFloat(expectedValue, 'f', -1, 64) {
        panic("float64 value and buffer are not equal")
    }
    _testBuffer.Reset()
}

func TestBuffer_WriteByte(t *testing.T) {
    expectedValue := byte(5)
    _testBuffer.WriteByte(expectedValue)
    if string(_testBuffer) != string(expectedValue) {
        panic("byte value and buffer are not equal")
    }
    _testBuffer.Reset()
}

func TestBuffer_WriteBool(t *testing.T) {
    expectedValue := true
    _testBuffer.WriteBool(expectedValue)
    if string(_testBuffer) != strconv.FormatBool(expectedValue) {
        panic("bool value and buffer are not equal")
    }
    _testBuffer.Reset()
}

func TestBuffer_WriteString(t *testing.T) {
    expectedValue := "not that long but very long string that will be tested"
    _testBuffer.WriteString(expectedValue)
    if string(_testBuffer) != expectedValue {
        panic("string value and buffer are not equal")
    }
    _testBuffer.Reset()
}

func TestBuffer_WriteNewLine(t *testing.T) {
    _testBuffer.WriteNewLine()
    if string(_testBuffer) != "\n" {
        panic("new line is not presented in the buffer")
    }
    _testBuffer.Reset()
}

func TestBuffer_WriteSpace(t *testing.T) {
    _testBuffer.WriteSpace()
    if string(_testBuffer) != " " {
        panic("space is not presented in the buffer")
    }
    _testBuffer.Reset()
}
