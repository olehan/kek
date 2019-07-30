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
        ps.Buffer.WriteString("<nil>")
    case complex64:
    case complex128:
    default:
        f.StringifyValueWithReflect(ps, v)
    }
}

// StringifyValueWithReflect converts the given value using the reflect.
func (f *FormatterUtils) StringifyValueWithReflect(ps *pool.State, v interface{}) {
    typeOfValue := reflect.TypeOf(v)
    switch typeOfValue.Kind() {
    case reflect.Struct:
        f.StringifyStruct(ps, typeOfValue)
    default:
        ps.Buffer.WriteString(typeOfValue.Name())
        ps.Buffer.WriteString(typeOfValue.String())
    }
}

// StringifyStruct converts structure into a string.
func (f *FormatterUtils) StringifyStruct(ps *pool.State, t reflect.Type) {
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

// StringifyByTemplate converts value into a string by the given format/template.
func (f *FormatterUtils) StringifyByTemplate(ps *pool.State, template string, values ...interface{}) {
    templateLen := len(template)
    printedValues := 0
    valuesLen := len(values)

    for i := 0; i < templateLen; i++ {
        templateLenIndex := templateLen - 1


        if  i < templateLenIndex &&
            printedValues < valuesLen &&
            template[i] == templateValueStart &&
            template[i + 1] == templateValueEnd {

            f.StringifyValue(ps, values[printedValues])

            printedValues++
            i++
        } else {
            ps.Buffer.WriteByte(template[i])
        }

        if templateLenIndex < i {
            break
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
