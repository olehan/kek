package formatters

import (
    "github.com/olehan/kek/pool"
    "reflect"
)

func (f *FormatterUtils) StringifyValues(ps *pool.PoolState, v ...interface{}) {
    for _, n := range v {
        f.StringifyValue(ps, n)
        ps.Buffer.WriteSpace()
    }
}

func (f *FormatterUtils) StringifyValue(ps *pool.PoolState, v interface{}) {
    switch v.(type) {
    case int:
        ps.Buffer.WriteInt(int64(v.(int)))
    case int8:
        ps.Buffer.WriteInt(int64(v.(int8)))
    case int16:
        ps.Buffer.WriteInt(int64(v.(int16)))
    case int32:
        ps.Buffer.WriteInt(int64(v.(int32)))
    case int64:
        ps.Buffer.WriteInt(v.(int64))
    case uint:
        ps.Buffer.WriteUint(uint64(v.(uint)))
    case uint8:
        ps.Buffer.WriteUint(uint64(v.(uint8)))
    case uint16:
        ps.Buffer.WriteUint(uint64(v.(uint16)))
    case uint32:
        ps.Buffer.WriteUint(uint64(v.(uint32)))
    case uint64:
        ps.Buffer.WriteUint(v.(uint64))
    case uintptr:
        ps.Buffer.WriteUint(uint64(v.(uintptr)))
    case string:
        ps.Buffer.WriteString(v.(string))
    case bool:
        ps.Buffer.WriteBool(v.(bool))
    case float32:
        ps.Buffer.WriteFloat(float64(v.(float32)), 32)
    case float64:
        ps.Buffer.WriteFloat(v.(float64), 64)
    case nil:
        ps.Buffer.WriteString("nil")
    case complex64:
    case complex128:
    default:
        f.StringifyValueWithReflect(ps, v)
    }
}

func (f *FormatterUtils) StringifyValueWithReflect(ps *pool.PoolState, v interface{}) {
    typeOfValue := reflect.TypeOf(v)
    switch typeOfValue.Kind() {
    case reflect.Struct:
        f.StringifyStruct(ps, typeOfValue)
    default:
        ps.Buffer.WriteString(typeOfValue.Name())
        ps.Buffer.WriteString(typeOfValue.String())
    }
}

func (f *FormatterUtils) StringifyStruct(ps *pool.PoolState, t reflect.Type) {
    ps.Buffer.WriteNewLine()
    ps.Buffer.WriteString(t.Name())
    ps.Buffer.WriteSpace()
    ps.Buffer.WriteByte('{')
    ps.Buffer.WriteNewLine()

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        ps.Buffer.WriteByte('\t')
        ps.Buffer.WriteString(field.Name)
        ps.Buffer.WriteString(": ")
        ps.Buffer.WriteString(field.Type.String())
        ps.Buffer.WriteNewLine()
    }

    ps.Buffer.WriteByte('}')
    ps.Buffer.WriteNewLine()
}
