package formatters

import (
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/pool"
    "reflect"
)

const (
    templateValueStart = '{'
    templateValueEnd = '}'
)

// StringifyValues loops through the given values, converts them into string
// and writes into a buffer, separating each value
func (f *FormatterUtils) StringifyValues(ps *pool.State, v ...interface{}) {
    for _, n := range v {
        f.StringifyValue(ps, n)
        ps.Buffer.WriteSpace()
    }
}

// StringifyValue writes into the buffer converted to string value that
// it receives.
func (f *FormatterUtils) StringifyValue(ps *pool.State, v interface{}) {
    switch t := v.(type) {
    case nil:
        ps.Buffer.WriteString("<nil>")
    case int:
        ps.Buffer.WriteInt(int64(t))
    case int8:
        ps.Buffer.WriteInt(int64(t))
    case int16:
        ps.Buffer.WriteInt(int64(t))
    case int32:
        ps.Buffer.WriteInt(int64(t))
    case int64:
        ps.Buffer.WriteInt(t)
    case uint:
        ps.Buffer.WriteUint(uint64(t))
    case uint8:
        ps.Buffer.WriteUint(uint64(t))
    case uint16:
        ps.Buffer.WriteUint(uint64(t))
    case uint32:
        ps.Buffer.WriteUint(uint64(t))
    case uint64:
        ps.Buffer.WriteUint(t)
    case uintptr:
        ps.Buffer.WriteUint(uint64(t))
    case string:
        ps.Buffer.WriteString(t)
    case bool:
        ps.Buffer.WriteBool(t)
    case float32:
        ps.Buffer.WriteFloat(float64(t), 32)
    case float64:
        ps.Buffer.WriteFloat(t, 64)
    case complex64:
        f.writeComplex(ps, complex128(t), 64)
    case complex128:
        f.writeComplex(ps, t, 128)
    default:
        f.StringifyValueWithReflect(ps, reflect.ValueOf(v))
    }
}

// StringifyValueWithReflect converts the given value using the reflect.
func (f *FormatterUtils) StringifyValueWithReflect(ps *pool.State, value reflect.Value) {
    switch t := value.Kind(); t {
    case reflect.Invalid:
        ps.Buffer.WriteString("<invalid reflect.Value>")
    case reflect.Bool:
        ps.Buffer.WriteBool(value.Bool())
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        ps.Buffer.WriteInt(value.Int())
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        ps.Buffer.WriteUint(value.Uint())
    case reflect.Float32:
        ps.Buffer.WriteFloat(value.Float(), 32)
    case reflect.Float64:
        ps.Buffer.WriteFloat(value.Float(), 64)
    case reflect.Complex64:
        f.writeComplex(ps, value.Complex(), 64)
    case reflect.Complex128:
        f.writeComplex(ps, value.Complex(), 128)
    case reflect.String:
        ps.Buffer.WriteString(value.String())
    case reflect.Map:
        ps.Buffer.WriteString(value.Type().String())
        if value.IsNil() {
            ps.Buffer.WriteString("<nil>")
            return
        }
        ps.Buffer.WriteByte('{')
        ps.Buffer.WriteSpace()
        keys := value.MapKeys()
        for i, key := range keys {
            if i > 0 {
                ps.Buffer.WriteString(", ")
            }
            f.StringifyValueWithReflect(ps, key)
            ps.Buffer.WriteString(": ")
            f.StringifyValueWithReflect(ps, value.MapIndex(key))
        }
        ps.Buffer.WriteSpace()
        ps.Buffer.WriteByte('}')
    case reflect.Struct:
        ps.Buffer.WriteString(value.Type().String())
        ps.Buffer.WriteByte('{')
        ps.Buffer.WriteSpace()

        for i := 0; i < value.NumField(); i++ {
            if i > 0 {
                ps.Buffer.WriteString(", ")
            }

            field := value.Field(i)
            ps.Buffer.WriteString(value.Type().Field(i).Name)
            ps.Buffer.WriteString(": ")
            f.StringifyValueWithReflect(ps, field)
        }

        ps.Buffer.WriteSpace()
        ps.Buffer.WriteByte('}')
    case reflect.Interface:
        v := value.Elem()
        if !v.IsValid() {
            ps.Buffer.WriteString(v.Type().String())
        } else {
            f.StringifyValue(ps, v.Interface())
        }
    case reflect.Array, reflect.Slice:
        ps.Buffer.WriteByte('[')
        for i := 0; i < value.Len(); i++ {
            if i > 0 {
                ps.Buffer.WriteByte(' ')
            }
            f.StringifyValueWithReflect(ps, value.Index(i))
        }
        ps.Buffer.WriteByte(']')
    case reflect.Ptr:
        if value.Pointer() != 0 {
            switch a := value.Elem(); value.Kind() {
            case reflect.Array, reflect.Slice, reflect.Struct, reflect.Map:
                ps.Buffer.WriteByte('&')
                f.StringifyValueWithReflect(ps, a)
                return
            }
        }
        fallthrough
    case reflect.Chan, reflect.Func, reflect.UnsafePointer:
        switch value.Kind() {
        case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
            ptr := value.Pointer()
            ps.Buffer.WriteByte('(')
            ps.Buffer.WriteString(value.Type().String())
            ps.Buffer.WriteString(")(")
            if ptr == 0 {
                ps.Buffer.WriteString("<nil>")
            } else {
                ps.Buffer.WriteUint(uint64(ptr))
            }
            ps.Buffer.WriteByte(')')
        }
    default:
        ps.Buffer.WriteString(value.String())
    }
}

// StringifyByTemplate converts value into a string by the given format/template.
func (f *FormatterUtils) StringifyByTemplate(ps *pool.State, template string, values ...interface{}) {
    templateLen := len(template)
    printedValues := 0
    valuesLen := len(values)

    i := 0
    for i < templateLen {
        templateLenIndex := templateLen - 1

        if  i < templateLenIndex &&
            printedValues < valuesLen &&
            template[i] == templateValueStart &&
            template[i + 1] == templateValueEnd {

            f.StringifyValue(ps, values[printedValues])

            printedValues++
            i += 2
        } else {
            ps.Buffer.WriteByte(template[i])
            i++
        }
    }
}

// StringifyByTemplateMap converts value that is given in the form of a map
// into a string by the given format/template
func (f *FormatterUtils) StringifyByTemplateMap(ps *pool.State, template string, v ds.Map) {
    templateLen := len(template)

    for i := 0; i < templateLen; i++ {
        templateLenIndex := templateLen - 1

        if  i < templateLenIndex &&
            template[i] == templateValueStart &&
            template[i + 1] == templateValueStart {

            skipCount := 2

            stopIndex := i + skipCount
            stopIndexEnd := stopIndex
            for j := stopIndex; j < templateLenIndex; j++ {
                if template[j] == templateValueEnd && template[j + 1] == templateValueEnd {
                    stopIndexEnd = j
                    break
                }
            }

            valueKey := template[stopIndex:stopIndexEnd]
            if value, ok := v[valueKey]; ok {
                f.StringifyValue(ps, value)
                i += stopIndexEnd + skipCount - i
            }
        }

        if templateLenIndex < i {
            break
        }

        ps.Buffer.WriteByte(template[i])
    }
}

// KeyValuesToMap loops through the key value pair array, setting them
// into a Map in the pool state.
func (f *FormatterUtils) KeyValuesToMap(ps *pool.State, keyValues ...interface{}) {
    keyValuesLen := len(keyValues)
    for i := 0; i < keyValuesLen; i += 2 {
        if keyValuesLen < i + 2 {
            break
        }

        key := keyValues[i]

        switch key.(type) {
        case string:
        default:
            continue
        }

        ps.Map.Set(key.(string), keyValues[i + 1])
    }
}

func (f *FormatterUtils) writeComplex(ps *pool.State, v complex128, size int) {
    ps.Buffer.WriteByte('(')
    ps.Buffer.WriteFloat(real(v), size / 2)
    ps.Buffer.WriteFloat(imag(v), size / 2)
    ps.Buffer.WriteString("i)")
}
