package minified

import (
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/formatters"
)

func (m *FormatterModel) Print(fs *formatters.FormatterConfig, v ...interface{}) {
    WriteMetaInfo(fs)
    m.StringifyValues(fs.PoolState, v...)
}

func (m *FormatterModel) PrintTemplate(
    fs *formatters.FormatterConfig,
    template string,
    v ...interface{},
) {
    WriteMetaInfo(fs)
    m.StringifyByTemplate(fs.PoolState, template, v...)
}

func (m *FormatterModel) PrintTemplateMap(
    fs *formatters.FormatterConfig,
    template string,
    v ds.Map,
) {
    WriteMetaInfo(fs)
    m.StringifyByTemplateMap(fs.PoolState, template, v)
}

func (m *FormatterModel) PrintTemplateKeyValue(
    fs *formatters.FormatterConfig,
    template string,
    keyValues ...interface{},
) {
    m.KeyValuesToMap(fs.PoolState, keyValues...)
    WriteMetaInfo(fs)
    m.StringifyByTemplateMap(fs.PoolState, template, fs.PoolState.Map)
}

func (m *FormatterModel) PrintStructKeyValues(
    fs *formatters.FormatterConfig,
    message string,
    keyValues ...interface{},
) {
    WriteMetaInfo(fs)

    fs.PoolState.Buffer.WriteString(message)
    fs.PoolState.Buffer.WriteNewLine()

    keyValuesLen := len(keyValues)

    for i := 0; i < keyValuesLen; i += 2 {
        if keyValuesLen < i + 2 {
            break
        }

        key := keyValues[i]

        fs.PoolState.Buffer.WriteString("   ")

        if fs.WithColors {
            fs.PoolState.Buffer.WriteString(bold)
        }

        switch key.(type) {
        case string:
            fs.PoolState.Buffer.WriteString(key.(string))
        default:
            m.StringifyValue(fs.PoolState, key)
        }

        if fs.WithColors {
            fs.PoolState.Buffer.WriteString(reset)
        }

        fs.PoolState.Buffer.WriteString(": ")

        m.StringifyValue(fs.PoolState, keyValues[i + 1])
        fs.PoolState.Buffer.WriteNewLine()
    }
}
