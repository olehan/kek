package formatters

import (
    "github.com/olehan/kek/buffer"
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/pool"
    "strings"
    "testing"
)

var (
    _testFormatterPool = pool.NewPool()
    _testFormatter = NewFormatterUtils()
)

type _namedStruct struct {
    PublicField string
}

func (_namedStruct) method(arg int) string {
    return "a"
}

func TestFormatterUtils_StringifyValues(t *testing.T) {
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
    interfaceSlice := []interface{}{}

    state := _testFormatterPool.Get()

    _testFormatter.StringifyValues(
        state,
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
        interfaceSlice,
    )

    t.Log(string(state.Buffer))
    _testFormatterPool.Free(state)
}

func TestFormatterUtils_KeyValuesToMap(t *testing.T) {
    state := _testFormatterPool.Get()
    expectedKey := "key1"
    expectedValue := "value1"
    _testFormatter.KeyValuesToMap(state, expectedKey, expectedValue, 134, 134, "skips this")
    if state.Map[expectedKey] != expectedValue {
        t.Fail()
    }
    _testFormatterPool.Free(state)
}

func TestFormatterUtils_StringifyByTemplateMap(t *testing.T) {
    state := _testFormatterPool.Get()
    expectedKey := "key1"
    expectedValue := "value1"
    template := "key: {{key1}}"
    expectedOutput := strings.Replace(template, "{{" + expectedKey + "}}", expectedValue, 1)
    _testFormatter.StringifyByTemplateMap(
        state,
        template,
        ds.NewMap().Set(expectedKey, expectedValue),
    )
    if expectedOutput != string(state.Buffer) {
        t.Fail()
    }
    _testFormatterPool.Free(state)
}

func TestFormatterUtils_StringifyByTemplate(t *testing.T) {
    values := []interface{}{"value", 2345245, 245235624.234562453, false}

    buf := buffer.NewBuffer()
    buf.WriteString(values[0].(string))
    buf.WriteSpace()
    buf.WriteInt(int64(values[1].(int)))
    buf.WriteSpace()
    buf.WriteFloat(values[2].(float64), 64)
    buf.WriteSpace()
    buf.WriteBool(values[3].(bool))

    state := _testFormatterPool.Get()
    _testFormatter.StringifyByTemplate(state, "{} {} {} {}", values...)
    if string(state.Buffer) != string(buf) {
        t.Fail()
    }
    _testFormatterPool.Free(state)
}
