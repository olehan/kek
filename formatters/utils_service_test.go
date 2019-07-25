package formatters

import (
    "github.com/olehan/kek/pool"
    "testing"
)

var (
    _testFormatterPoolState = pool.NewPool().Get()
    _testFormatter = NewFormatterUtils()
)

type _namedStruct struct {
    PublicField string
}

func (_namedStruct) method(arg int) string {
    return "a"
}

func TestFormatter_stringifyValues(t *testing.T) {
    intValue := 13453463
    int8Value := int8(127)
    int16Value := int16(3003)
    int32Value := int32(345693)
    int64Value := int64(2454562492)
    uintValue := uint(134134)
    uint8Value := uint8(134)
    uint16Value := uint16(13413)
    uint32Value := uint32(13411334)
    uint64Value := uint64(13524636364)
    uintptrValue := uintptr(346457456356)
    stringValue := string("oih24t9h334rad")
    boolValue := true
    float32Value := float32(245245.23423)
    float64Value := float64(2463621.351313)
    complex64Value := complex64(2346363623452.45235135134134)
    complex128Value := complex128(345346246134.13513513513413)
    structValue := struct {
        x int
        y int
        a func() interface{}
        slice []string
    }{
        x: 1,
        y: 2,
        a: func() interface{} {
            return ""
        },
        slice: []string{""},
    }

    _testFormatter.StringifyValues(
        _testFormatterPoolState,
        nil,
        _namedStruct{},
        intValue,
        int8Value,
        int16Value,
        int32Value,
        int64Value,
        uintValue,
        uint8Value,
        uint16Value,
        uint32Value,
        uint64Value,
        uintptrValue,
        stringValue,
        boolValue,
        float32Value,
        float64Value,
        complex64Value,
        complex128Value,
        structValue,
    )

    t.Log(string(_testFormatterPoolState.Buffer))
}
